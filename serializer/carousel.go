package serializer

import "go_mall/model"

type CarouselVO struct {
	Id        uint   `json:"id"`
	ImgPath   string `json:"img_path"`
	ProductId uint   `json:"product_id"`
	CreateAt  int64  `json:"create_at"`
}

func BuildCarousel(item *model.Carousel) CarouselVO {
	return CarouselVO{
		Id:        item.ID,
		ImgPath:   item.ImgPath,
		ProductId: item.ProductId,
		CreateAt:  item.CreatedAt.Unix(),
	}
}

func BuildCarouselList(items []model.Carousel) (carousels []CarouselVO) {
	for _, item := range items {
		carousel := BuildCarousel(&item)
		carousels = append(carousels, carousel)
	}
	return carousels
}
