package repository

import (
	"context"
	"golang-restapi-noframework/model"
)

type ProductRepository interface {
	Create(ctx context.Context, product model.Product)
	Update(ctx context.Context, product model.Product)
	Delete(ctx context.Context, productId int)
	GetById(ctx context.Context, productId int) (model.Product, error)
	GetAll(ctx context.Context) []model.Product
}
