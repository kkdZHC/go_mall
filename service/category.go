package service

import (
	"context"
	"go_mall/dao"
	"go_mall/serializer"
	"net/http"
)

type CategoryService struct {
}

func (service *CategoryService) List(ctx context.Context) serializer.Response {
	categoryDao := dao.NewCategoryDao(ctx)
	categories, err := categoryDao.ListCategory()
	if err != nil {
		return serializer.Response{
			Status: http.StatusNotFound,
			Msg:    "Error",
			Error:  "获取轮播图失败",
		}
	}
	return serializer.BuildListResponse(serializer.BuildCategoryList(categories), uint(len(categories)))
}
