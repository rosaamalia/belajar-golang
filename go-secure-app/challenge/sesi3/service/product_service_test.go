package service

import (
	"sesi3/entity"
	"sesi3/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productRepository = &repository.ProductRepositoryMock{Mock: mock.Mock{}}
var productService = ProductService{Repository: productRepository}

func TestProductServiceGetOneProductNotFound(t *testing.T) {
	productRepository.Mock.On("FindById", "1").Return(nil)

	product, err := productService.GetOneProduct("1")

	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, "Product not found", err.Error(), "Error response has to be 'Product not found'")
}

func TestProductServiceGetOneProduct(t *testing.T) {
	product := entity.Product{
		ID: 2,
		CreatedAt: "2023-04-14T23:28:52.9345909+07:00",
		UpdatedAt: "2023-04-14T23:28:52.9345909+07:00",
		Title: "Buku Novel",
		Description: "Buku novel The Gambler",
	}

	productRepository.Mock.On("FindById", "2").Return(product)

	result, err := productService.GetOneProduct("2")

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, product.ID, result.ID, "Result has to be '2'")
	assert.Equal(t, product.Title, result.Title, "Result has to be 'Buku Novel")
	assert.Equal(t, &product, result, "Result has to be a product data with id '2'")
}

func TestProductServiceGetAllProductNotFound(t *testing.T) {
	productRepository.Mock.On("GetAll").Return(nil)

	result, err := productService.GetAllProduct()

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, "No product saved", err.Error(), "Error response has to be 'No product saved'")
}

func TestProductServiceGetAllProduct(t *testing.T) {
	products := []*entity.Product{
		&entity.Product{
			ID: 2,
			CreatedAt: "2023-04-14T23:28:52.9345909+07:00",
			UpdatedAt: "2023-04-14T23:28:52.9345909+07:00",
			Title: "Buku Novel",
			Description: "Buku novel The Gambler",
		},
		&entity.Product{
			ID: 3,
			CreatedAt: "2023-04-14T23:28:52.9345909+07:00",
			UpdatedAt: "2023-04-14T23:28:52.9345909+07:00",
			Title: "Sarung Bantal",
			Description: "Sarung bantal bunga",
		},
		&entity.Product{
			ID: 4,
			CreatedAt: "2023-04-14T23:28:52.9345909+07:00",
			UpdatedAt: "2023-04-14T23:28:52.9345909+07:00",
			Title: "Tas Ransel",
			Description: "Tas ransel biru",
		},
	}

	productRepository.Mock.On("GetAll").Return(products)

	result, err := productService.GetAllProduct()

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, products, result)
}