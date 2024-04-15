package service

import (
	_itemShopModel "github.com/go-suriya/shop-api-echo/pkg/itemShop/model"
)

type ItemShopService interface {
	Listing(itemFilter *_itemShopModel.ItemFilter) (*_itemShopModel.ItemResult, error)
}
