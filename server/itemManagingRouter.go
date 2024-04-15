package server

import (
	_itemManagingController "github.com/go-suriya/shop-api-echo/pkg/itemManaging/controller"
	_itemManagingRepository "github.com/go-suriya/shop-api-echo/pkg/itemManaging/repository"
	_itemManagingService "github.com/go-suriya/shop-api-echo/pkg/itemManaging/service"
	_itemShopRepository "github.com/go-suriya/shop-api-echo/pkg/itemShop/repository"
)

func (s *echoServer) initItemManagingRouter(m *authorizingMiddleware) {
	router := s.app.Group("/v1/item-managing")

	itemShopRepository := _itemShopRepository.NewItemShopRepositoryImpl(s.db, s.app.Logger)
	itemMangingRepository := _itemManagingRepository.NewItemManagingRepositoryImpl(s.db, s.app.Logger)

	itemManagingService := _itemManagingService.NewItemManagingServiceImpl(itemMangingRepository, itemShopRepository)

	itemManaging := _itemManagingController.NewItemManagingControllerImpl(itemManagingService)

	router.POST("", itemManaging.Creating)
	router.PATCH("/:itemID", itemManaging.Editing)
	router.DELETE("/:itemID", itemManaging.Archiving)
}
