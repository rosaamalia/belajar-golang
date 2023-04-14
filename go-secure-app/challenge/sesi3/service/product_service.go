package service

import (
	"errors"
	"sesi3/entity"
	"sesi3/repository"
)

type ProductService struct {
	Repository repository.ProductRepository
}

func (service ProductService) GetOneProduct(id string) (*entity.Product, error) {
	product := service.Repository.FindById(id)
	if product == nil {
		return nil, errors.New("Product not found")
	}

	return product, nil
}

func (service ProductService) GetAllProduct() ([]*entity.Product, error) {
	products := service.Repository.GetAll()

	if products == nil {
		return nil, errors.New("No product saved")
	}

	return products, nil
}