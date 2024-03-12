package dao

import (
	"context"
	"go_mall/model"

	"gorm.io/gorm"
)

type ProductImgDao struct {
	*gorm.DB
}

func NewProductImgDao(ctx context.Context) *ProductImgDao {
	return &ProductImgDao{NewDBClient(ctx)}
}

func NewProductImgDaoByDB(db *gorm.DB) *ProductImgDao {
	return &ProductImgDao{db}
}

func (dao *ProductImgDao) CreateProductImg(productImg *model.ProductImg) error {
	err := dao.DB.Model(&model.ProductImg{}).Create(&productImg).Error
	if err != nil {
		return err
	}
	return nil
}
func (dao *ProductImgDao) ListProductImg(id uint) (productImg []*model.ProductImg, err error) {
	err = dao.DB.Model(&model.ProductImg{}).Where("product_id=?", id).Find(&productImg).Error
	return
}
