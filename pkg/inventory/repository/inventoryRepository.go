package repository

import (
	entities "github.com/go-suriya/shop-api-echo/entities"
	"gorm.io/gorm"
)

type InventoryRepository interface {
	Filling(playerID string, itemID uint64, qty int, tx *gorm.DB) ([]*entities.Inventory, error)
	Listing(playerID string) ([]*entities.Inventory, error)
	Removing(playerID string, itemID uint64, limit int, tx *gorm.DB) error
	PlayerItemCounting(playerID string, itemID uint64) int64
}
