package controller

import (
	"net/http"

	custom "github.com/go-suriya/shop-api-echo/pkg/custom"
	_inventoryService "github.com/go-suriya/shop-api-echo/pkg/inventory/service"
	"github.com/go-suriya/shop-api-echo/pkg/validation"
	"github.com/labstack/echo/v4"
)

type inventoryControllerImpl struct {
	inventoryService _inventoryService.InventoryService
	logger           echo.Logger
}

func NewInventoryControllerImpl(
	inventoryService _inventoryService.InventoryService,
	logger echo.Logger,
) InventoryController {
	return &inventoryControllerImpl{
		inventoryService: inventoryService,
		logger:           logger,
	}
}

func (c *inventoryControllerImpl) Listing(pctx echo.Context) error {
	playerID, err := validation.PlayerIDGetting(pctx)
	if err != nil {
		c.logger.Error("Failed to get playerID", err.Error())
		return custom.Error(pctx, http.StatusUnauthorized, err)
	}

	inventoryListing, err := c.inventoryService.Listing(playerID)
	if err != nil {
		return pctx.JSON(http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusOK, inventoryListing)
}
