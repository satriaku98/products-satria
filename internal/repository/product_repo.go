package repository

import (
	"context"
	"erajaya/internal/model"
	"erajaya/pkg/utils"
)

type ProductRepository interface {
	Create(product *model.Product) error
	GetAllByOrder(order []utils.OrderParam) ([]model.Product, error)
	ExistsByName(name string) (bool, error)
	GetCachedProducts(ctx context.Context, sortParam string) (products []model.Product, err error)
	SetCachedProducts(ctx context.Context, sortParam string, products []model.Product) error
	DeleteCachedProducts(ctx context.Context) error
}
