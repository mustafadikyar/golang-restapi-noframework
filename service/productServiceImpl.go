package service

import (
	"context"
	"golang-restapi-noframework/dto/request"
	"golang-restapi-noframework/dto/response"
	"golang-restapi-noframework/helper"
	"golang-restapi-noframework/model"
	"golang-restapi-noframework/repository"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
}

func NewProductServiceImpl(productRepository repository.ProductRepository) ProductService {
	return &ProductServiceImpl{ProductRepository: productRepository}
}

// Create implements ProductService.
func (p *ProductServiceImpl) Create(ctx context.Context, request request.CreateProductRequest) {
	product := model.Product{
		Name: request.Name,
	}
	p.ProductRepository.Create(ctx, product)
}

// Delete implements ProductService.
func (p *ProductServiceImpl) Delete(ctx context.Context, productId int) {
	product, err := p.ProductRepository.GetById(ctx, productId)
	helper.PanicIfError(err)
	p.ProductRepository.Delete(ctx, product.Id)
}

// GetAll implements ProductService.
func (p *ProductServiceImpl) GetAll(ctx context.Context) []response.ProductResponse {
	products := p.ProductRepository.GetAll(ctx)

	var productResponse []response.ProductResponse

	for _, value := range products {
		product := response.ProductResponse{Id: value.Id, Name: value.Name}
		productResponse = append(productResponse, product)
	}
	return productResponse
}

// GetById implements ProductService.
func (p *ProductServiceImpl) GetById(ctx context.Context, productId int) response.ProductResponse {
	product, err := p.ProductRepository.GetById(ctx, productId)
	helper.PanicIfError(err)
	return response.ProductResponse(product)
}

// Update implements ProductService.
func (p *ProductServiceImpl) Update(ctx context.Context, request request.UpdateProductRequest) {
	product, err := p.ProductRepository.GetById(ctx, request.Id)
	helper.PanicIfError(err)

	product.Name = request.Name
	p.ProductRepository.Update(ctx, product)
}
