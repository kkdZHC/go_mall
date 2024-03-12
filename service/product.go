package service

import (
	"context"
	"fmt"
	"go_mall/dao"
	"go_mall/model"
	"go_mall/serializer"
	"mime/multipart"
	"net/http"
	"strconv"
	"sync"
)

type ProductService struct {
	ID             uint   `form:"id" json:"id"`
	Name           string `form:"name" json:"name"`
	CategoryID     uint   `form:"category_id" json:"category_id"`
	Title          string `form:"title" json:"title" `
	Info           string `form:"info" json:"info" `
	ImgPath        string `form:"img_path" json:"img_path"`
	Price          string `form:"price" json:"price"`
	DiscountPrice  string `form:"discount_price" json:"discount_price"`
	OnSale         bool   `form:"on_sale" json:"on_sale"`
	Num            int    `form:"num" json:"num"`
	model.BasePage        //分页
}

func (service *ProductService) Create(ctx context.Context, uId uint, files []*multipart.FileHeader) serializer.Response {
	var boss *model.User
	userDao := dao.NewUserDao(ctx)
	boss, err := userDao.GetUserById(uId)
	if err != nil {
		return serializer.Response{
			Status: http.StatusNotFound,
			Msg:    "Error",
			Error:  "获取用户失败",
		}
	}
	//多张图片以第一张为封面
	fmt.Println(files)
	//
	tmp, _ := files[0].Open()
	path, err := UploadProductToLocalStatic(tmp, uId, service.Name)
	if err != nil {
		return serializer.Response{
			Status: 30001,
			Msg:    "ErrorUploadProductImg",
			Error:  err.Error(),
		}
	}
	product := &model.Product{
		Name:          service.Name,
		CategoryId:    service.CategoryID,
		Title:         service.Title,
		Info:          service.Info,
		ImgPath:       path,
		Price:         service.Price,
		DiscountPrice: service.DiscountPrice,
		OnSale:        true,
		Num:           service.Num,
		BossId:        boss.ID,
		BossName:      boss.UserName,
		BossAvatar:    boss.Avatar,
	}
	productDao := dao.NewProductDao(ctx)
	err = productDao.CreateProduct(product)
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "ErrorCreateProduct",
			Error:  "创建product失败",
		}
	}
	//并发创建
	wg := new(sync.WaitGroup)
	wg.Add(len(files))
	errChan := make(chan error)
	defer close(errChan)
	for index, file := range files {
		num := strconv.Itoa(index)
		go service.createProductImg(productDao, file, uId, num, product, wg, errChan)
	}
	wg.Wait()
	select {
	case err = <-errChan:
		return serializer.Response{
			Status: http.StatusInternalServerError,
			Msg:    "createProductImgError",
			//Data:   serializer.BuildProduct(product),
			Error: err.Error(),
		}
	default:
		return serializer.Response{
			Status: http.StatusOK,
			Msg:    "ok",
			Data:   serializer.BuildProductVO(product),
		}
	}

}

func (service *ProductService) createProductImg(productDao *dao.ProductDao, file *multipart.FileHeader, uId uint, num string, product *model.Product, wg *sync.WaitGroup, errChan chan<- error) {
	defer wg.Done()
	productImgDao := dao.NewProductImgDaoByDB(productDao.DB)
	tmp, _ := file.Open()
	path, err := UploadProductToLocalStatic(tmp, uId, service.Name+num)
	if err != nil {
		errChan <- err
	}
	productImg := &model.ProductImg{
		ProductId: product.ID,
		ImgPath:   path,
	}
	err = productImgDao.CreateProductImg(productImg)
	if err != nil {
		errChan <- err
	}
}

func (service *ProductService) List(ctx context.Context) serializer.Response {
	var products []*model.Product
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	condition := make(map[string]interface{})
	if service.CategoryID != 0 {
		condition["category_id"] = service.CategoryID
	}
	productDao := dao.NewProductDao(ctx)

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		productDao = dao.NewProductDaoByDB(productDao.DB)
		products, _ = productDao.ListProductByCondition(condition, service.BasePage)
		wg.Done()
	}()

	total, err := productDao.CountProductByCondition(condition)
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "countProductError",
			Error:  err.Error(),
		}
	}

	wg.Wait()
	return serializer.BuildListResponse(serializer.BuildProductVOs(products), uint(total))
}

func (service *ProductService) Search(ctx context.Context) serializer.Response {
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	productDao := dao.NewProductDao(ctx)
	products, count, err := productDao.SearchProduct(service.Info, service.BasePage)
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "SearchProductError",
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildProductVOs(products), uint(count))
}

func (service *ProductService) Show(ctx context.Context, id string) serializer.Response {
	pId, _ := strconv.Atoi(id)
	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(uint(pId))
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "GetProductByIdError",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: http.StatusOK,
		Msg:    "ok",
		Data:   serializer.BuildProductVO(product),
	}
}
