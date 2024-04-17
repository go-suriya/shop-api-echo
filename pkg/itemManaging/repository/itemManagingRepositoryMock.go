package repository

import (
	entities "github.com/go-suriya/shop-api-echo/entities"
	_itemManagingModel "github.com/go-suriya/shop-api-echo/pkg/itemManaging/model"

	"github.com/stretchr/testify/mock"
)

type ItemManagingRepositoryMock struct {
	mock.Mock
}

func (m *ItemManagingRepositoryMock) Creating(itemEntity *entities.Item) (*entities.Item, error) {
	args := m.Called(itemEntity)
	return args.Get(0).(*entities.Item), args.Error(1)
}

func (m *ItemManagingRepositoryMock) Editing(itemID uint64, itemEditingReq *_itemManagingModel.ItemEditingReq) (uint64, error) {
	args := m.Called(itemID, itemEditingReq)
	return args.Get(0).(uint64), args.Error(1)
}

func (m *ItemManagingRepositoryMock) Archiving(itemID uint64) error {
	args := m.Called(itemID)
	return args.Error(0)
}
