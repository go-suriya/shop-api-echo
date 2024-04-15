package service

import (
	_itemManagingModel "github.com/go-suriya/shop-api-echo/pkg/itemManaging/model"
	_itemShopModel "github.com/go-suriya/shop-api-echo/pkg/itemShop/model"
)

type ItemManagingService interface {
	Creating(itemCreatingReq *_itemManagingModel.ItemCreatingReq) (*_itemShopModel.Item, error)
	Editing(itemID uint64, editItemReq *_itemManagingModel.ItemEditingReq) (*_itemShopModel.Item, error)
	Archiving(itemID uint64) error
}
