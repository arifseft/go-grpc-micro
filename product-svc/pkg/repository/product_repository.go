package repository

import (
	"github.com/arifseft/go-grpc-micro/product-svc/pkg/db"
	"github.com/arifseft/go-grpc-micro/product-svc/pkg/models"
)

type IProductRepository interface {
    FindOne(id int64) (*models.Product, error)
    UpdateProductStock(product *models.Product, stock int64) (*models.Product, error)
    CreateProduct(product *models.Product) (*models.Product, error)
}

type ProductRepository struct {
    H db.Handler
}

func (r *ProductRepository) FindOne(id int64) (*models.Product, error) {
    var product models.Product

    result := r.H.DB.First(&product, id)

    if result.Error != nil {
        return &product, result.Error
    }

    return &product, nil
}

func (r *ProductRepository) CreateProduct(product *models.Product) (*models.Product, error) {
    result := r.H.DB.Create(&product)

    if result.Error != nil {
        return nil, result.Error
    }

    return product, nil
}

func (r *ProductRepository) UpdateProductStock(product *models.Product, stock int64) (*models.Product, error) {
    product.Stock = stock

    result := r.H.DB.Save(&product)

    if result.Error != nil {
        return nil, result.Error
    }

    return product, nil
}
