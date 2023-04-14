package repository

import (
	"sesi3/entity"

	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	Mock mock.Mock
}

func (repository *ProductRepositoryMock) FindById(id string) *entity.Product {
	arguments := repository.Mock.Called(id)

	if arguments.Get(0) == nil {
		return nil
	}

	product := arguments.Get(0).(entity.Product)

	return &product
}

func (repository *ProductRepositoryMock) GetAll() []*entity.Product {
	arguments := repository.Mock.Called()

	if arguments.Get(0) == nil {
		return nil
	}

	products := arguments.Get(0).([]*entity.Product)

	return products
}