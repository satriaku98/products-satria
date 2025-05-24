package service

import (
	"context"
	"erajaya/internal/model"
	"erajaya/internal/repository"
	"erajaya/pkg/utils"
	"fmt"
	"regexp"
	"strings"
)

var allowedSortKeys = map[string]bool{
	"price":      true,
	"created_at": true,
}

type service struct {
	repo repository.ProductRepository
}

func NewService(repo repository.ProductRepository) Service {
	return &service{repo}
}

func (s *service) AddProduct(ctx context.Context, p *model.Product) error {
	// Validasi name
	p.Name = strings.TrimSpace(p.Name)
	if len(p.Name) < utils.ProductNameMinLength || len(p.Name) > utils.ProductNameMaxLength {
		return utils.ErrInvalidName
	}
	if !regexp.MustCompile(`^[a-zA-Z0-9\s\-]+$`).MatchString(p.Name) {
		return utils.ErrNameAlphaNumeric
	}

	// Cek duplikat
	exists, err := s.repo.ExistsByName(p.Name)
	if err != nil {
		return err
	}
	if exists {
		return utils.ErrProductExists
	}

	// Validasi panjang deskripsi
	if len(p.Description) > utils.ProductDescriptionMaxLength {
		return utils.ErrDescriptionTooLong
	}

	// Insert ke DB
	err = s.repo.Create(p)
	if err != nil {
		return err
	}

	// Hapus cache
	err = s.repo.DeleteCachedProducts(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) ListProduct(ctx context.Context, order []utils.OrderParam) ([]model.Product, error) {
	orderParams := ""
	// Validasi order params
	for _, o := range order {
		if !allowedSortKeys[o.Key] {
			return nil, utils.ErrInvalidSortField
		}
		if o.Direction != utils.ASC && o.Direction != utils.DESC {
			return nil, utils.ErrInvalidSortDirection
		}

		orderParams += fmt.Sprintf("%s:%s,", o.Key, o.Direction)
	}
	orderParams = strings.TrimSuffix(orderParams, ",")

	// Ambil data dari Redis
	products, err := s.repo.GetCachedProducts(ctx, orderParams)
	if err == nil || len(products) > 0 {
		return products, nil
	}

	// jika tidak ada di Redis, ambil dari DB
	products, err = s.repo.GetAllByOrder(order)
	if err != nil {
		return nil, err
	}

	// simpan ke Redis
	err = s.repo.SetCachedProducts(ctx, orderParams, products)
	if err != nil {
		return nil, err
	}

	return products, err
}
