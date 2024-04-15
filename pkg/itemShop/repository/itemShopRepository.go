package repository

import (
	"github.com/go-suriya/shop-api-echo/entities"
	_itemShopModel "github.com/go-suriya/shop-api-echo/pkg/itemShop/model"
	"gorm.io/gorm"
)

type ItemShopRepository interface {
	BeginTransaction() *gorm.DB
	RollbackTransaction(tx *gorm.DB) error
	CommitTransaction(tx *gorm.DB) error
	Listing(itemFilterDto *_itemShopModel.ItemFilter) ([]*entities.Item, error)
	FindByID(itemID uint64) (*entities.Item, error)
	FindByIDList(itemIDs []uint64) ([]*entities.Item, error)
	Counting(itemFilterDto *_itemShopModel.ItemFilter) (int64, error)
	PurchaseHistoryRecording(
		purchasingEntity *entities.PurchaseHistory,
		tx *gorm.DB,
	) (*entities.PurchaseHistory, error)
}
