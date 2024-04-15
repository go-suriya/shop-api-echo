package controller

import (
	"net/http"

	custom "github.com/go-suriya/shop-api-echo/pkg/custom"
	_playerCoinModel "github.com/go-suriya/shop-api-echo/pkg/playerCoin/model"
	_playerCoinService "github.com/go-suriya/shop-api-echo/pkg/playerCoin/service"
	"github.com/go-suriya/shop-api-echo/pkg/validation"
	"github.com/labstack/echo/v4"
)

type playerCoinControllerImpl struct {
	playerCoinService _playerCoinService.PlayerCoinService
}

func NewPlayerCoinControllerImpl(playerCoinService _playerCoinService.PlayerCoinService) PlayerCoinController {
	return &playerCoinControllerImpl{
		playerCoinService: playerCoinService,
	}
}

func (c *playerCoinControllerImpl) CoinAdding(pctx echo.Context) error {
	playerID, err := validation.PlayerIDGetting(pctx)
	if err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}

	coinAddingReq := new(_playerCoinModel.CoinAddingReq)

	validatingContext := custom.NewCustomEchoRequest(pctx)

	if err := validatingContext.Bind(coinAddingReq); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}
	coinAddingReq.PlayerID = playerID

	playerCoin, err := c.playerCoinService.CoinAdding(coinAddingReq)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusCreated, playerCoin)
}

func (c *playerCoinControllerImpl) PlayerCoinShowing(pctx echo.Context) error {
	playerID, err := validation.PlayerIDGetting(pctx)
	if err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}

	playerCoin := c.playerCoinService.PlayerCoinShowing(playerID)

	return pctx.JSON(http.StatusOK, playerCoin)
}
