package service

import (
	"context"
	"go_mall/conf"
	"go_mall/dao"
	"go_mall/model"
	"go_mall/pkg/util"
	"go_mall/serializer"
	"mime/multipart"
	"net/http"
	"strings"
	"time"

	"gopkg.in/mail.v2"
)

type UserService struct {
	NickName string `json:"nick_name" form:"nick_name"`
	UserName string `json:"user_name" form:"user_name"`
	Password string `json:"password" form:"password"`
	Key      string `json:"key" form:"key"` //前端验证
}

type SendEmailService struct {
	Email         string `json:"email" form:"email"`
	Password      string `json:"password" form:"password"`
	OperationType uint   `json:"operation_type" form:"operation_type"`
	//1：绑定邮箱2：解绑邮箱3：改密码
}

type ValidEmailService struct {
}

type ShowMoneyService struct {
	Key string `json:"key" form:"key"`
}

func (service *UserService) Register(ctx context.Context) serializer.Response {
	var user model.User

	if service.Key == "" || len(service.Key) != 16 { //密钥错误
		return serializer.Response{
			Status: http.StatusOK,
			Msg:    "ok",
			Error:  "密钥出错！",
		}
	}
	//10000初始金额 --> 密文存储余额 对称加密操作
	util.Encrypt.SetKey(service.Key)

	userDao := dao.NewUserDao(ctx)
	_, exist, err := userDao.ExistOrNotByUserName(service.UserName)
	if err != nil {
		return serializer.Response{
			Status: http.StatusInternalServerError,
			Msg:    "error",
			Error:  "查询用户名是否存在出错",
		}
	}
	if exist {
		return serializer.Response{
			Status: http.StatusConflict,
			Msg:    "ErrorExistUser",
			Error:  "用户名已存在",
		}
	}
	user = model.User{
		UserName: service.UserName,
		//Email: "",
		//PasswordDigest: "",
		NickName: service.NickName,
		Status:   model.Active,
		Avatar:   "avatar.jpg",
		Money:    util.Encrypt.AesEncoding("10000"), //初始10000元
	}
	//密码加密
	err = user.SetPassword(service.Password)
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "ErrorFailEncryption",
			Error:  "密码加密失败",
		}
	}
	//创建用户
	err = userDao.CreateUser(&user)
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "ErrorCreateUser",
			Error:  "创建用户失败",
		}
	}
	return serializer.Response{
		Status: http.StatusOK,
		Msg:    "OK",
	}
}

func (service *UserService) Login(ctx context.Context) serializer.Response {
	var user *model.User
	userDao := dao.NewUserDao(ctx)
	//判断是否存在用户
	user, exist, err := userDao.ExistOrNotByUserName(service.UserName)
	if !exist || err != nil {
		return serializer.Response{
			Status: http.StatusNotAcceptable,
			Data:   "用户不存在请先注册",
			Msg:    "ErrorUserNotExist",
			Error:  "查找用户失败",
		}
	}
	//校验密码
	if !user.CheckPassword(service.Password) {
		return serializer.Response{
			Status: http.StatusNotFound,
			Data:   "密码错误, 还可以输入5次",
			Msg:    "ErrorPasswordNotMatch",
			Error:  "密码错误",
		}
	}
	//携带token返回, token签发
	token, err := util.GenerateToken(user.ID, service.UserName, 0)
	if err != nil {
		return serializer.Response{
			Status: http.StatusBadRequest,
			Msg:    "ErrorAuthToken",
			Error:  "token签名出错",
		}
	}
	return serializer.Response{
		Status: http.StatusOK,
		Data: serializer.TokenData{
			User:  serializer.BuildUserVO(user),
			Token: token,
		},
		Msg: "ok",
	}
}

// 修改用户昵称
func (service *UserService) Update(ctx context.Context, uId uint) serializer.Response {
	var user *model.User
	var err error

	userDao := dao.NewUserDao(ctx)
	user, err = userDao.GetUserById(uId)
	if err != nil {
		return serializer.Response{
			Status: http.StatusNotFound,
			Msg:    "Error",
			Error:  "获取用户失败",
		}
	}
	//修改昵称
	if service.NickName != "" {
		user.NickName = service.NickName
	}
	err = userDao.UpdateUserById(uId, user)
	if err != nil {
		return serializer.Response{
			Status: http.StatusBadRequest,
			Msg:    "Error",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: http.StatusOK,
		Msg:    "ok",
		Data:   serializer.BuildUserVO(user),
	}
}

// 上传头像
func (service *UserService) Post(ctx context.Context, uId uint, file multipart.File, fileSize int64) serializer.Response {

	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(uId)
	if err != nil {
		return serializer.Response{
			Status: http.StatusNotFound,
			Msg:    "Error",
			Error:  "获取用户失败",
		}
	}
	// 保存图片到本地
	path, err := UploadAvatarToLocalStatic(file, uId, user.UserName)
	if err != nil {
		return serializer.Response{
			Status: 30001,
			Msg:    "ErrorUploadFile",
			Error:  err.Error(),
		}
	}
	user.Avatar = path
	err = userDao.UpdateUserById(uId, user)
	if err != nil {
		return serializer.Response{
			Status: http.StatusBadRequest,
			Msg:    "Error",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: http.StatusOK,
		Msg:    "ok",
		Data:   serializer.BuildUserVO(user),
	}
}

// 发邮件
func (service *SendEmailService) Send(ctx context.Context, uId uint) serializer.Response {
	var address string
	var notice *model.Notice //绑定邮箱 修改密码 模板通知
	token, err := util.GenerateEmailToken(uId, service.OperationType, service.Email, service.Password)
	if err != nil {
		return serializer.Response{
			Status: http.StatusFailedDependency,
			Msg:    "GenerateEmailTokenError",
			Error:  err.Error(),
		}
	}
	noticeDao := dao.NewNoticeDao(ctx)
	notice, err = noticeDao.GetNoticeById(service.OperationType)
	if err != nil {
		return serializer.Response{
			Status: http.StatusNotFound,
			Msg:    "Error",
			Error:  "获取Notice失败",
		}
	}

	address = conf.ValidEmail + token //发送方
	mailStr := notice.Text
	mailTex := strings.Replace(mailStr, "Email", address, -1)
	m := mail.NewMessage()
	m.SetHeader("From", conf.SmtpEmail)
	m.SetHeader("To", service.Email)
	m.SetHeader("Subject", "ZHC")
	m.SetBody("text/html", mailTex)
	d := mail.NewDialer(conf.SmtpHost, 465, conf.SmtpEmail, conf.SmtpPass)
	d.StartTLSPolicy = mail.MandatoryStartTLS
	err = d.DialAndSend(m)
	if err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "ErrorSendEmail",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: http.StatusOK,
		Msg:    "ok",
	}
}

// 验证邮箱
func (service *ValidEmailService) Valid(ctx context.Context, token string) serializer.Response {
	var userId uint
	var email string
	var password string
	var operationType uint
	if token == "" {
		return serializer.Response{
			Status: http.StatusNotFound,
			Msg:    "ErrorTokenMissing",
		}
	} else {
		claims, err := util.ParseEmailToken(token)
		if err != nil {
			return serializer.Response{
				Status: http.StatusFailedDependency,
				Msg:    "ParseEmailTokenError",
				Error:  err.Error(),
			}
		} else if time.Now().Unix() > claims.ExpiresAt {
			return serializer.Response{
				Status: http.StatusRequestTimeout,
				Msg:    "EmailTokenTimeOut",
			}
		} else {
			userId = claims.UserID
			email = claims.Email
			password = claims.Password
			operationType = claims.OperationType
		}
	}
	//获取该用户信息
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(userId)
	if err != nil {
		return serializer.Response{
			Status: http.StatusNotFound,
			Msg:    "Error",
			Error:  "获取用户失败",
		}
	}
	switch operationType {
	case 1:
		//绑定邮箱
		user.Email = email
	case 2:
		//解绑邮箱
		user.Email = ""
	case 3:
		err = user.SetPassword(password)
		return serializer.Response{
			Status: 500,
			Msg:    "ErrorFailEncryption",
			Data:   "密码加密失败",
			Error:  err.Error(),
		}
	}
	err = userDao.UpdateUserById(userId, user)
	if err != nil {
		return serializer.Response{
			Status: http.StatusBadRequest,
			Msg:    "Error",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: http.StatusOK,
		Msg:    "ok",
		Data:   serializer.BuildUserVO(user),
	}
}

// 展示用户金钱
func (service *ShowMoneyService) Show(ctx context.Context, uId uint) serializer.Response {
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(uId)
	if err != nil {
		return serializer.Response{
			Status: http.StatusNotFound,
			Msg:    "Error",
			Error:  "获取用户失败",
		}
	}
	return serializer.Response{
		Status: http.StatusOK,
		Data:   serializer.BuildMoney(user, service.Key),
		Msg:    "ok",
	}
}
