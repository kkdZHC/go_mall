package service

import (
	"context"
	"go_mall/dao"
	"go_mall/model"
	"go_mall/serializer"
	"strconv"
)

type AddressService struct {
	Name    string `json:"name" form:"name"`
	Phone   string `json:"phone" form:"phone"`
	Address string `json:"address" form:"address"`
}

func (service *AddressService) Create(ctx context.Context, uId uint) serializer.Response {
	var address *model.Address
	addressDao := dao.NewAddressDao(ctx)
	address = &model.Address{
		UserID:  uId,
		Name:    service.Name,
		Phone:   service.Phone,
		Address: service.Address,
	}
	err := addressDao.CreateAddress(address)
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "CreateAddressError",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "ok",
	}
}

func (service *AddressService) Show(ctx context.Context, aId string) serializer.Response {
	var address *model.Address
	addressDao := dao.NewAddressDao(ctx)
	addressId, _ := strconv.Atoi(aId)
	address, err := addressDao.GetAddressByAid(uint(addressId))
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "GetAddressByUidError",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "ok",
		Data:   serializer.BuildAddress(address),
	}
}

func (service *AddressService) List(ctx context.Context, uId uint) serializer.Response {
	addressDao := dao.NewAddressDao(ctx)
	addresses, err := addressDao.ListAddressByUid(uId)
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "ListAddressError",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "ok",
		Data:   serializer.BuildAddressList(addresses),
	}
}

func (service *AddressService) Update(ctx context.Context, uId uint, aId string) serializer.Response {
	var address *model.Address
	addressDao := dao.NewAddressDao(ctx)
	address = &model.Address{
		Name:    service.Name,
		Phone:   service.Phone,
		Address: service.Address,
	}
	addressId, _ := strconv.Atoi(aId)
	err := addressDao.UpdateAddressByUid(uId, uint(addressId), address)
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "UpdateAddressError",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "ok",
	}
}

func (service *AddressService) Delete(ctx context.Context, uId uint, aId string) serializer.Response {
	addressDao := dao.NewAddressDao(ctx)
	addressId, _ := strconv.Atoi(aId)
	err := addressDao.DeleteAddressByAid(uId, uint(addressId))
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "DeleteAddressByAidError",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "ok",
	}
}
