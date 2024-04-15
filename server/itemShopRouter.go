package server

import (
	_inventoryRepository "github.com/go-suriya/shop-api-echo/pkg/inventory/repository"
	_itemShopController "github.com/go-suriya/shop-api-echo/pkg/itemShop/controller"
	_itemShopRepository "github.com/go-suriya/shop-api-echo/pkg/itemShop/repository"
	_itemShopService "github.com/go-suriya/shop-api-echo/pkg/itemShop/service"
	_playerCoinRepository "github.com/go-suriya/shop-api-echo/pkg/playerCoin/repository"
)

func (s *echoServer) initItemShopRouter(m *authorizingMiddleware) {
	router := s.app.Group("/v1/item-shop")

	playerCoinRepository := _playerCoinRepository.NewPlayerCoinRepositoryImpl(s.db, s.app.Logger)
	inventoryRepository := _inventoryRepository.NewInventoryRepositoryImpl(s.db, s.app.Logger)
	itemShopRepository := _itemShopRepository.NewItemShopRepositoryImpl(s.db, s.app.Logger)

	itemShopService := _itemShopService.NewItemShopServiceImpl(
		itemShopRepository,
		playerCoinRepository,
		inventoryRepository,
		s.app.Logger,
	)

	itemShopController := _itemShopController.NewItemShopControllerImpl(itemShopService)

	router.GET("", itemShopController.Listing)
	router.POST("/buying", itemShopController.Buying, m.PlayerAuthorizing)
	router.POST("/selling", itemShopController.Selling, m.PlayerAuthorizing)
}
