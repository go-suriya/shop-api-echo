package repository

import (
	"github.com/go-suriya/shop-api-echo/databases"
	"github.com/go-suriya/shop-api-echo/entities"
	"github.com/labstack/echo/v4"

	_itemManagingException "github.com/go-suriya/shop-api-echo/pkg/itemManaging/exception"
	"github.com/go-suriya/shop-api-echo/pkg/itemManaging/model"
)

type itemMangingRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewItemManagingRepositoryImpl(db databases.Database, logger echo.Logger) ItemManagingRepository {
	return &itemMangingRepositoryImpl{db, logger}
}

func (r *itemMangingRepositoryImpl) Creating(itemEntity *entities.Item) (*entities.Item, error) {
	item := new(entities.Item)

	if err := r.db.Connect().Create(itemEntity).Scan(item).Error; err != nil {
		r.logger.Error("Item creating failed:", err.Error())
		return nil, &_itemManagingException.ItemCreating{}
	}

	return item, nil
}

func (r *itemMangingRepositoryImpl) Editing(itemID uint64, itemEditingReq *model.ItemEditingReq) (uint64, error) {
	if err := r.db.Connect().Model(&entities.Item{}).Where(
		"id = ?", itemID,
	).Updates(
		itemEditingReq,
	).Error; err != nil {
		r.logger.Errorf("Editing item failed: %s", err.Error())
		return 0, &_itemManagingException.ItemEditing{}
	}

	return itemID, nil
}

func (r *itemMangingRepositoryImpl) Archiving(itemID uint64) error {
	if err := r.db.Connect().Table("items").Where(
		"id = ?", itemID,
	).Update(
		"is_archive", true,
	).Error; err != nil {
		r.logger.Error("Archiving item failed:", err.Error())
		return &_itemManagingException.ItemArchiving{ItemID: itemID}
	}

	return nil
}
