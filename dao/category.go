package dao

import (
	"context"
	"go_mall/model"

	"gorm.io/gorm"
)

type CategoryDao struct {
	*gorm.DB
}

func NewCategoryDao(ctx context.Context) *CategoryDao {
	return &CategoryDao{NewDBClient(ctx)}
}

func NewCategoryDaoByDB(db *gorm.DB) *CategoryDao {
	return &CategoryDao{db}
}

// 通过id找Category
func (dao *CategoryDao) ListCategory() ([]*model.Category, error) {
	var category []*model.Category
	err := dao.DB.Model(&model.Category{}).Find(&category).Error
	return category, err
}
