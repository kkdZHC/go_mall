package dao

import (
	"context"
	"go_mall/model"

	"gorm.io/gorm"
)

type AddressDao struct {
	*gorm.DB
}

func NewAddressDao(ctx context.Context) *AddressDao {
	return &AddressDao{NewDBClient(ctx)}
}

func NewAddressDaoByDB(db *gorm.DB) *AddressDao {
	return &AddressDao{db}
}

func (dao *AddressDao) CreateAddress(address *model.Address) error {
	err := dao.DB.Model(&model.Address{}).Create(&address).Error
	return err
}

func (dao *AddressDao) GetAddressByAid(aId uint) (address *model.Address, err error) {
	err = dao.DB.Model(&model.Address{}).Where("id=?", aId).First(&address).Error
	return
}
func (dao *AddressDao) ListAddressByUid(uId uint) (addresses []*model.Address, err error) {
	err = dao.DB.Model(&model.Address{}).Where("user_id=?", uId).Find(&addresses).Error
	return
}
func (dao *AddressDao) UpdateAddressByUid(uId, aId uint, address *model.Address) (err error) {
	err = dao.DB.Model(&model.Address{}).Where("id=? AND user_id=?", aId, uId).Updates(address).Error
	return
}
func (dao *AddressDao) DeleteAddressByAid(uId, aId uint) error {
	err := dao.DB.Model(&model.Address{}).Where("id=? AND user_id=?", aId, uId).Delete(&model.Address{}).Error
	return err
}
