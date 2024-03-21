package serializer

import (
	"context"
	"go_mall/conf"
	"go_mall/dao"
	"go_mall/model"
)

type FavoriteVO struct {
	UserId        uint   `json:"user_id"`
	ProductId     uint   `json:"product_id"`
	CreatedAt     int64  `json:"created_at"`
	Name          string `json:"name"`
	CategoryId    uint   `json:"category_id"`
	Title         string `json:"title"`
	Info          string `json:"info"`
	ImgPath       string `json:"img_path"`
	Price         string `json:"price"`
	DiscountPrice string `json:"discount_price"`
	BossId        uint   `json:"boss_id"`
	Num           int    `json:"num"`
	OnSale        bool   `json:"on_sale"`
}

func BuildFavorite(favorite *model.Favorite, product *model.Product) *FavoriteVO {
	return &FavoriteVO{
		UserId:        favorite.UserId,
		ProductId:     favorite.ProductId,
		CreatedAt:     favorite.CreatedAt.Unix(),
		Name:          product.Name,
		CategoryId:    product.CategoryId,
		Title:         product.Title,
		Info:          product.Info,
		ImgPath:       conf.Host + conf.HttpPort + conf.ProductPath + product.ImgPath,
		Price:         product.Price,
		DiscountPrice: product.DiscountPrice,
		BossId:        favorite.BossId,
		Num:           product.Num,
		OnSale:        product.OnSale,
	}
}

func BuildFavoriteList(ctx context.Context, items []*model.Favorite) (favorites []*FavoriteVO) {
	productDao := dao.NewProductDao(ctx)
	for _, item := range items {
		product, _ := productDao.GetProductById(item.ProductId)
		favorite := BuildFavorite(item, product)
		favorites = append(favorites, favorite)
	}
	return
}
