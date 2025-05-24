package repository

import (
	"context"
	"encoding/json"
	"erajaya/internal/model"
	"erajaya/pkg/utils"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type productRepository struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewProductRepository(db *gorm.DB, rdb *redis.Client) ProductRepository {
	return &productRepository{db: db, rdb: rdb}
}

func (r *productRepository) Create(product *model.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) GetAllByOrder(order []utils.OrderParam) ([]model.Product, error) {
	var products []model.Product
	tx := r.db.Model(&model.Product{})

	for _, param := range order {
		tx = tx.Order(fmt.Sprintf("%s %s", param.Key, param.Direction))
	}

	err := tx.Find(&products).Error
	return products, err
}

func (r *productRepository) ExistsByName(name string) (bool, error) {
	var count int64
	err := r.db.Model(&model.Product{}).Where("name = ?", name).Count(&count).Error
	return count > 0, err
}

func (r *productRepository) GetCachedProducts(ctx context.Context, sortParam string) (products []model.Product, err error) {
	cacheKey := fmt.Sprintf("products:sort=%s", sortParam)
	cached, err := r.rdb.Get(ctx, cacheKey).Result()
	json.Unmarshal([]byte(cached), &products)
	return products, err
}

func (r *productRepository) SetCachedProducts(ctx context.Context, sortParam string, products []model.Product) error {
	cacheKey := fmt.Sprintf("products:sort=%s", sortParam)
	data, err := json.Marshal(products)
	if err != nil {
		return err
	}
	err = r.rdb.Set(ctx, cacheKey, data, 5*time.Minute).Err()
	return err
}

func (r *productRepository) DeleteCachedProducts(ctx context.Context) error {
	cacheKey := "products:*"
	return r.rdb.Del(ctx, cacheKey).Err()
}
