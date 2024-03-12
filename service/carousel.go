package service

import (
	"context"
	"go_mall/dao"
	"go_mall/serializer"
	"net/http"
)

type CarouselService struct {
}

func (service *CarouselService) List(ctx context.Context) serializer.Response {
	carouselDao := dao.NewCarouselDao(ctx)
	carousels, err := carouselDao.ListCarousel()
	if err != nil {
		return serializer.Response{
			Status: http.StatusNotFound,
			Msg:    "Error",
			Error:  "获取轮播图失败",
		}
	}
	return serializer.BuildListResponse(serializer.BuildCarouselList(carousels), uint(len(carousels)))
}
