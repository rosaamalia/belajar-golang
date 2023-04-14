package repository

import (
	"sesi3/entity"
)

type ProductRepository interface {
	GetAll() []*entity.Product
	FindById(id string) *entity.Product
}