package dao

import (
	"context"
	"go_mall/model"

	"gorm.io/gorm"
)

type FavoriteDao struct {
	*gorm.DB
}

func NewFavoriteDao(ctx context.Context) *FavoriteDao {
	return &FavoriteDao{NewDBClient(ctx)}
}

func NewFavoriteDaoByDB(db *gorm.DB) *FavoriteDao {
	return &FavoriteDao{db}
}

func (dao *FavoriteDao) ListFavorite(uId uint) (favorites []*model.Favorite, err error) {
	err = dao.DB.Model(&model.Favorite{}).Where("user_id=?", uId).Find(&favorites).Error
	return
}
func (dao *FavoriteDao) ExistOrNot(pId, uId uint) (exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.Favorite{}).Where("product_id=? AND user_id=?", pId, uId).Count(&count).Error
	if err != nil || count == 0 {
		return false, err
	}
	return true, err
}

func (dao *FavoriteDao) CreateFavorite(item *model.Favorite) error {
	err := dao.DB.Model(&model.Favorite{}).Create(&item).Error
	return err
}
