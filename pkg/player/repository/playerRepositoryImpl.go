package repository

import (
	"github.com/go-suriya/shop-api-echo/databases"
	entities "github.com/go-suriya/shop-api-echo/entities"
	_player "github.com/go-suriya/shop-api-echo/pkg/player/exception"

	"github.com/labstack/echo/v4"
)

type playerRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewPlayerRepositoryImpl(db databases.Database, logger echo.Logger) PlayerRepository {
	return &playerRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (r *playerRepositoryImpl) Creating(playerEntity *entities.Player) (*entities.Player, error) {
	insertedPlayer := new(entities.Player)

	if err := r.db.Connect().Create(playerEntity).Scan(insertedPlayer).Error; err != nil {
		r.logger.Error("Creating player failed", err.Error())
		return nil, &_player.PlayerCreating{PlayerID: playerEntity.ID}
	}

	return insertedPlayer, nil
}

func (r *playerRepositoryImpl) FindByID(playerID string) (*entities.Player, error) {
	player := new(entities.Player)

	if err := r.db.Connect().Where("id = ?", playerID).First(player).Error; err != nil {
		r.logger.Errorf("Finding player failed: %s", err.Error())
		return nil, &_player.PlayerNotFound{PlayerID: playerID}
	}

	return player, nil
}
