package service

import (
	"context"
	"go_mall/dao"
	"go_mall/model"
	"go_mall/serializer"
	"net/http"
	"strconv"
)

type FavoriteService struct {
	ProductId  uint `json:"product_id" form:"product_id"`
	BossId     uint `json:"boss_id" form:"boss_id"`
	FavoriteId uint `json:"favorite_id" form:"favorite_id"`
	model.BasePage
}

func (service *FavoriteService) List(ctx context.Context, uId uint) serializer.Response {
	favoriteDao := dao.NewFavoriteDao(ctx)
	favorites, err := favoriteDao.ListFavorite(uId)
	if err != nil {
		return serializer.Response{
			Status: http.StatusNotFound,
			Msg:    "Error",
			Error:  "获取收藏夹失败",
		}
	}
	return serializer.BuildListResponse(serializer.BuildFavoriteList(ctx, favorites), uint(len(favorites)))
}

func (service *FavoriteService) Create(ctx context.Context, uId uint) serializer.Response {
	favoriteDao := dao.NewFavoriteDao(ctx)
	exist, _ := favoriteDao.ExistOrNot(service.ProductId, uId)
	if exist {
		return serializer.Response{
			Status: 500,
			Msg:    "FavoriteAlreadyExist",
			Error:  "已收藏",
		}
	}
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(uId)
	if err != nil {
		return serializer.Response{
			Status: http.StatusNotFound,
			Msg:    "Error",
			Error:  "获取用户失败",
		}
	}
	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(service.ProductId)
	if err != nil {
		return serializer.Response{
			Status: http.StatusNotFound,
			Msg:    "Error",
			Error:  "获取商品失败",
		}
	}
	favorite := &model.Favorite{
		User:      *user,
		UserId:    uId,
		Product:   *product,
		ProductId: service.ProductId,
		BossId:    service.BossId,
	}
	err = favoriteDao.CreateFavorite(favorite)
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "CreateFavoriteError",
			Error:  "创建收藏失败",
		}
	}
	return serializer.Response{
		Status: http.StatusOK,
		Msg:    "ok",
		Data:   favorite,
	}
}

func (service *FavoriteService) Delete(ctx context.Context, uId uint, fId string) serializer.Response {
	favoriteDao := dao.NewFavoriteDao(ctx)
	favoriteId, _ := strconv.Atoi(fId)
	err := favoriteDao.DeleteFavorite(uId, uint(favoriteId))
	if err != nil {
		return serializer.Response{
			Status: http.StatusNotFound,
			Msg:    "Error",
			Error:  "获取收藏夹失败",
		}
	}
	return serializer.Response{
		Status: http.StatusOK,
		Msg:    "ok",
		Error:  "删除成功",
	}
}
