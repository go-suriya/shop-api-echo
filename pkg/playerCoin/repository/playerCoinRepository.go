package repository

import (
	"github.com/go-suriya/shop-api-echo/entities"
	_playerCoinModel "github.com/go-suriya/shop-api-echo/pkg/playerCoin/model"
	"gorm.io/gorm"
)

type PlayerCoinRepository interface {
	CoinAdding(playerCoinEntity *entities.PlayerCoin, tx *gorm.DB) (*entities.PlayerCoin, error)
	Showing(playerID string) (*_playerCoinModel.PlayerCoinShowing, error)
}
