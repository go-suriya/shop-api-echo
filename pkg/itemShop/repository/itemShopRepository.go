package repository

import (
	"github.com/go-suriya/shop-api-echo/entities"
	_itemShopModel "github.com/go-suriya/shop-api-echo/pkg/itemShop/model"
)

type ItemShopRepository interface {
	Listing(itemFilter *_itemShopModel.ItemFilter) ([]*entities.Item, error)
	Counting(itemFilter *_itemShopModel.ItemFilter) (int64, error)
	FindByID(itemID uint64) (*entities.Item, error)
}
