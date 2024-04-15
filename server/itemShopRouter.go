package server

import (
	_itemShopController "github.com/go-suriya/shop-api-echo/pkg/itemShop/controller"
	_itemShopRepository "github.com/go-suriya/shop-api-echo/pkg/itemShop/repository"
	_itemShopService "github.com/go-suriya/shop-api-echo/pkg/itemShop/service"
)

func (s *echoServer) initItemShopRouter(m *authorizingMiddleware) {
	router := s.app.Group("/v1/item-shop")

	itemShopRepository := _itemShopRepository.NewItemShopRepositoryImpl(s.db, s.app.Logger)

	itemShopService := _itemShopService.NewItemShopServiceImpl(itemShopRepository)

	itemShopController := _itemShopController.NewItemShopControllerImpl(itemShopService)

	router.GET("", itemShopController.Listing)
}
