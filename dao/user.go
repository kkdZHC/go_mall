package dao

import (
	"context"
	"go_mall/model"

	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDBClient(ctx)}
}

func NewUserDaoByDB(db *gorm.DB) *UserDao {
	return &UserDao{db}
}

// 根据username判断是否存在
func (dao *UserDao) ExistOrNotByUserName(userName string) (user *model.User, exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.User{}).Where("user_name=?", userName).Find(&user).Count(&count).Error
	if count == 0 || err != nil {
		return nil, false, err
	}
	return user, true, nil
}

func (dao *UserDao) CreateUser(user *model.User) error {
	err := dao.DB.Model(&model.User{}).Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

// 通过uid找user
func (dao *UserDao) GetUserById(uId uint) (*model.User, error) {
	var user *model.User
	err := dao.DB.Model(&model.User{}).Where("id=?", uId).First(&user).Error
	return user, err
}

// 通过uid更新user
func (dao *UserDao) UpdateUserById(uId uint, user *model.User) error {
	err := dao.DB.Model(&model.User{}).Where("id=?", uId).Updates(&user).Error
	return err
}
