package service

import (
	"context"
	"go_mall/dao"
	"go_mall/serializer"
	"strconv"
)

type ListProductImg struct {
}

func (service *ListProductImg) List(ctx context.Context, id string) serializer.Response {
	productImgDao := dao.NewProductImgDao(ctx)
	productId, _ := strconv.Atoi(id)
	productImgs, _ := productImgDao.ListProductImg(uint(productId))
	return serializer.BuildListResponse(
		serializer.BuildProductImgs(productImgs),
		uint(len(productImgs)),
	)
}
