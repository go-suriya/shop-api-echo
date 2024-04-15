package service

import (
	_itemShopModel "github.com/go-suriya/shop-api-echo/pkg/itemShop/model"
	_playerCoinModel "github.com/go-suriya/shop-api-echo/pkg/playerCoin/model"
)

type ItemShopService interface {
	Listing(itemFilter *_itemShopModel.ItemFilter) (*_itemShopModel.ItemResult, error)
	Buying(buyingReq *_itemShopModel.BuyingReq) (*_playerCoinModel.PlayerCoin, error)
	Selling(sellingReq *_itemShopModel.SellingReq) (*_playerCoinModel.PlayerCoin, error)
}
