package repository

import (
	entities "github.com/go-suriya/shop-api-echo/entities"
)

type AdminRepository interface {
	Creating(adminEntity *entities.Admin) (string, error)
	FindByID(adminID string) (*entities.Admin, error)
}
