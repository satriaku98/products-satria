package service

import (
	"context"
	"erajaya/internal/model"
	"erajaya/pkg/utils"
)

type Service interface {
	AddProduct(ctx context.Context, product *model.Product) error
	ListProduct(ctx context.Context, order []utils.OrderParam) ([]model.Product, error)
}
