package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-restapi-noframework/helper"
	"golang-restapi-noframework/model"
)

type ProductRepositoryImpl struct {
	Db *sql.DB
}

func NewProductRepository(Db *sql.DB) ProductRepository {
	return &ProductRepositoryImpl{Db: Db}
}

// Create implements ProductRepository.
func (p *ProductRepositoryImpl) Create(ctx context.Context, product model.Product) {
	tx, err := p.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "insert into product(name) values ($1)"
	_, err = tx.ExecContext(ctx, SQL, product.Name)
	helper.PanicIfError(err)
}

// Delete implements ProductRepository.
func (p *ProductRepositoryImpl) Delete(ctx context.Context, productId int) {
	tx, err := p.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "delete from product where id =$1"
	_, errExec := tx.ExecContext(ctx, SQL, productId)
	helper.PanicIfError(errExec)
}

// GetAll implements ProductRepository.
func (p *ProductRepositoryImpl) GetAll(ctx context.Context) []model.Product {
	tx, err := p.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "select id,name from product"
	result, errQuery := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(errQuery)
	defer result.Close()

	var products []model.Product

	for result.Next() {
		product := model.Product{}
		err := result.Scan(&product.Id, &product.Name)
		helper.PanicIfError(err)

		products = append(products, product)
	}

	return products
}

// GetById implements ProductRepository.
func (p *ProductRepositoryImpl) GetById(ctx context.Context, productId int) (model.Product, error) {
	tx, err := p.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "select id,name from product where id=$1"
	result, errQuery := tx.QueryContext(ctx, SQL, productId)
	helper.PanicIfError(errQuery)
	defer result.Close()

	product := model.Product{}

	if result.Next() {
		err := result.Scan(&product.Id, &product.Name)
		helper.PanicIfError(err)
		return product, nil
	} else {
		return product, errors.New("product id not found")
	}
}

// Update implements ProductRepository.
func (p *ProductRepositoryImpl) Update(ctx context.Context, product model.Product) {
	tx, err := p.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "update product set name=$1 where id=$2"
	_, err = tx.ExecContext(ctx, SQL, product.Name, product.Id)
	helper.PanicIfError(err)
}
