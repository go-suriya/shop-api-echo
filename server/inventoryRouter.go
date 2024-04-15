package server

import (
	_inventoryController "github.com/go-suriya/shop-api-echo/pkg/inventory/controller"
	_inventoryRepository "github.com/go-suriya/shop-api-echo/pkg/inventory/repository"
	_inventoryService "github.com/go-suriya/shop-api-echo/pkg/inventory/service"
	_itemShopRepository "github.com/go-suriya/shop-api-echo/pkg/itemShop/repository"
)

func (s *echoServer) initInventoryRouter(m *authorizingMiddleware) {
	router := s.app.Group("/v1/inventory")

	itemRepository := _itemShopRepository.NewItemShopRepositoryImpl(s.db, s.app.Logger)
	inventoryRepository := _inventoryRepository.NewInventoryRepositoryImpl(s.db, s.app.Logger)

	inventoryService := _inventoryService.NewInventoryServiceImpl(
		inventoryRepository,
		itemRepository,
	)

	inventoryController := _inventoryController.NewInventoryControllerImpl(inventoryService, s.app.Logger)

	router.GET("", inventoryController.Listing, m.PlayerAuthorizing)
}
