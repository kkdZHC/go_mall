package dao

import (
	"context"
	"go_mall/model"

	"gorm.io/gorm"
)

type OrderDao struct {
	*gorm.DB
}

func NewOrderDao(ctx context.Context) *OrderDao {
	return &OrderDao{NewDBClient(ctx)}
}

func NewOrderDaoByDB(db *gorm.DB) *OrderDao {
	return &OrderDao{db}
}

func (dao *OrderDao) CreateOrder(order *model.Order) error {
	err := dao.DB.Model(&model.Order{}).Create(&order).Error
	return err
}

func (dao *OrderDao) GetOrderByOid(oId, uId uint) (order *model.Order, err error) {
	err = dao.DB.Model(&model.Order{}).Where("id=? AND user_id=?", oId, uId).First(&order).Error
	return
}
func (dao *OrderDao) ListOrderByUid(uId uint) (orderes []*model.Order, err error) {
	err = dao.DB.Model(&model.Order{}).Where("user_id=?", uId).Find(&orderes).Error
	return
}
func (dao *OrderDao) UpdateOrderByUid(uId, aId uint, order *model.Order) (err error) {
	err = dao.DB.Model(&model.Order{}).Where("id=? AND user_id=?", aId, uId).Updates(order).Error
	return
}
func (dao *OrderDao) DeleteOrderByOid(oId, uId uint) error {
	err := dao.DB.Model(&model.Order{}).Where("id=? AND user_id=?", oId, uId).Delete(&model.Order{}).Error
	return err
}

func (dao *OrderDao) ListOrderByCondition(condition map[string]interface{}, page model.BasePage) (orders []*model.Order, total int64, err error) {
	err = dao.DB.Model(&model.Order{}).Where(condition).Count(&total).Error
	if err != nil {
		return
	}
	err = dao.DB.Model(&model.Order{}).Where(condition).Offset((page.PageNum - 1) * (page.PageSize)).Limit(page.PageSize).Find(&orders).Error
	return
}
