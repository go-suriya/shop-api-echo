package repository

import (
	"github.com/go-suriya/shop-api-echo/databases"
	entities "github.com/go-suriya/shop-api-echo/entities"
	_adminExpception "github.com/go-suriya/shop-api-echo/pkg/admin/exception"
	"github.com/labstack/echo/v4"
)

type adminRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewAdminRepositoryImpl(db databases.Database, logger echo.Logger) AdminRepository {
	return &adminRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (r *adminRepositoryImpl) Creating(adminEntity *entities.Admin) (string, error) {
	tx := r.db.Connect().Create(adminEntity)

	if tx.Error != nil {
		r.logger.Errorf("Error inserting player: %s", tx.Error.Error())
		return "", &_adminExpception.AdminCreating{AdminID: adminEntity.ID}
	}

	return adminEntity.ID, nil
}

func (r *adminRepositoryImpl) FindByID(adminID string) (*entities.Admin, error) {
	admin := new(entities.Admin)

	if err := r.db.Connect().Where("id = ?", adminID).First(admin).Error; err != nil {
		r.logger.Errorf("Error finding player: %s", err.Error())
		return nil, &_adminExpception.AdminNotFound{AdminID: adminID}
	}

	return admin, nil
}
