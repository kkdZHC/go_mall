package serializer

import (
	"go_mall/conf"
	"go_mall/model"
)

type ProductImgVO struct {
	ProductId uint   `json:"product_id"`
	ImgPath   string `json:"img_path"`
}

func BuildProductImg(item *model.ProductImg) *ProductImgVO {
	return &ProductImgVO{
		ProductId: item.ProductId,
		ImgPath:   conf.Host + conf.HttpPort + conf.ProductPath + item.ImgPath,
	}
}

func BuildProductImgs(items []*model.ProductImg) (productImgVOs []*ProductImgVO) {
	for _, item := range items {
		productImgVO := BuildProductImg(item)
		productImgVOs = append(productImgVOs, productImgVO)
	}
	return
}
