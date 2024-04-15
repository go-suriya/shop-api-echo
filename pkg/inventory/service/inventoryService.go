package service

import (
	_inventoryModel "github.com/go-suriya/shop-api-echo/pkg/inventory/model"
)

type InventoryService interface {
	Listing(playerID string) ([]*_inventoryModel.Inventory, error)
}
