package serializer

import (
	"go_mall/model"
	"go_mall/pkg/util"
)

type Money struct {
	UserId    uint   `json:"user_id" form:"user_id"`
	UserName  string `json:"user_name" form:"user_name"`
	UserMoney string `json:"user_money" form:"user_money"`
}

func BuildMoney(user *model.User, key string) Money {
	util.Encrypt.SetKey(key)
	return Money{
		UserId:    user.ID,
		UserName:  user.UserName,
		UserMoney: util.Encrypt.AesDecoding(user.Money),
	}
}
