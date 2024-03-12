package dao

import (
	"context"
	"go_mall/model"

	"gorm.io/gorm"
)

type CarouselDao struct {
	*gorm.DB
}

func NewCarouselDao(ctx context.Context) *CarouselDao {
	return &CarouselDao{NewDBClient(ctx)}
}

func NewCarouselDaoByDB(db *gorm.DB) *CarouselDao {
	return &CarouselDao{db}
}

// 通过id找Carousel
func (dao *CarouselDao) ListCarousel() ([]model.Carousel, error) {
	var carousel []model.Carousel
	err := dao.DB.Model(&model.Carousel{}).Find(&carousel).Error
	return carousel, err
}
