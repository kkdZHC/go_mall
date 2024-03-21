package dao

import (
	"context"
	"go_mall/model"

	"gorm.io/gorm"
)

type CartDao struct {
	*gorm.DB
}

func NewCartDao(ctx context.Context) *CartDao {
	return &CartDao{NewDBClient(ctx)}
}

func NewCartDaoByDB(db *gorm.DB) *CartDao {
	return &CartDao{db}
}

func (dao *CartDao) CreateCart(cart *model.Cart) error {
	err := dao.DB.Model(&model.Cart{}).Create(&cart).Error
	return err
}

func (dao *CartDao) ListCartByUid(uId uint) (cart []*model.Cart, err error) {
	err = dao.DB.Model(&model.Cart{}).Where("user_id=?", uId).Find(&cart).Error
	return
}
func (dao *CartDao) UpdateCartByUid(uId, cId uint, num uint) (err error) {
	err = dao.DB.Model(&model.Cart{}).Where("id=? AND user_id=?", cId, uId).Update("num", num).Error
	return
}
func (dao *CartDao) DeleteCartByAid(uId, cId uint) error {
	err := dao.DB.Model(&model.Cart{}).Where("id=? AND user_id=?", cId, uId).Delete(&model.Cart{}).Error
	return err
}
