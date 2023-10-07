package service

import (
	"context"
	"golang-restapi-noframework/dto/request"
	"golang-restapi-noframework/dto/response"
)

type ProductService interface {
	Create(ctx context.Context, request request.CreateProductRequest)
	Update(ctx context.Context, request request.UpdateProductRequest)
	Delete(ctx context.Context, productId int)
	GetById(ctx context.Context, productId int) response.ProductResponse
	GetAll(ctx context.Context) []response.ProductResponse
}
