package server

import (
	_playerCoinController "github.com/go-suriya/shop-api-echo/pkg/playerCoin/controller"
	_playerCoinRepository "github.com/go-suriya/shop-api-echo/pkg/playerCoin/repository"
	_playerCoinService "github.com/go-suriya/shop-api-echo/pkg/playerCoin/service"
)

func (s *echoServer) initPlayerCoinRouter(m *authorizingMiddleware) {
	router := s.app.Group("/v1/player-coin")

	playerCoinRepository := _playerCoinRepository.NewPlayerCoinRepositoryImpl(s.db, s.app.Logger)

	playerCoinService := _playerCoinService.NewPlayerCoinServiceImpl(
		playerCoinRepository,
	)
	playerCoinController := _playerCoinController.NewPlayerCoinControllerImpl(playerCoinService)

	router.POST("", playerCoinController.CoinAdding, m.PlayerAuthorizing)
	router.GET("", playerCoinController.PlayerCoinShowing, m.PlayerAuthorizing)
}
