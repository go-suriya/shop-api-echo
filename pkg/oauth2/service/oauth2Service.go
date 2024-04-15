package service

import (
	_adminModel "github.com/go-suriya/shop-api-echo/pkg/admin/model"
	_playerModel "github.com/go-suriya/shop-api-echo/pkg/player/model"
)

type OAuth2Service interface {
	PlayerAccountCreating(playerCreatingReq *_playerModel.PlayerCreatingReq) error
	AdminAccountCreating(adminCreatingInfoReq *_adminModel.AdminCreatingReq) error
	IsThisGuyIsReallyPlayer(playerID string) bool
	IsThisGuyIsReallyAdmin(adminID string) bool
}
