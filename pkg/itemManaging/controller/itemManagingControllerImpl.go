package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/go-suriya/shop-api-echo/pkg/custom"
	_itemManagingModel "github.com/go-suriya/shop-api-echo/pkg/itemManaging/model"
	_itemManagingService "github.com/go-suriya/shop-api-echo/pkg/itemManaging/service"
)

type itemManagingControllerImpl struct {
	itemMangingService _itemManagingService.ItemManagingService
}

func NewItemManagingControllerImpl(itemMangingService _itemManagingService.ItemManagingService) ItemManagingController {
	return &itemManagingControllerImpl{itemMangingService}
}

func (c *itemManagingControllerImpl) Creating(pctx echo.Context) error {
	itemCreatingReq := new(_itemManagingModel.ItemCreatingReq)

	customEchoRequest := custom.NewCustomEchoRequest(pctx)

	if err := customEchoRequest.Bind(itemCreatingReq); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}

	item, err := c.itemMangingService.Creating(itemCreatingReq)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusCreated, item)
}

func (c *itemManagingControllerImpl) Editing(pctx echo.Context) error {
	itemID, err := c.getItemID(pctx)
	if err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}

	editItemReq := new(_itemManagingModel.ItemEditingReq)

	validatingContext := custom.NewCustomEchoRequest(pctx)

	if err := validatingContext.Bind(editItemReq); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}

	item, err := c.itemMangingService.Editing(itemID, editItemReq)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusOK, item)
}

func (c *itemManagingControllerImpl) Archiving(pctx echo.Context) error {
	itemID, err := c.getItemID(pctx)
	if err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}

	err = c.itemMangingService.Archiving(itemID)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}

	return pctx.NoContent(http.StatusNoContent)
}

func (c *itemManagingControllerImpl) getItemID(pctx echo.Context) (uint64, error) {
	itemID := pctx.Param("itemID")
	itemIDUint64, err := strconv.ParseUint(itemID, 10, 64)
	if err != nil {
		return 0, err
	}

	return itemIDUint64, nil
}
