package service_test

import (
	"context"
	"erajaya/internal/model"
	"erajaya/internal/repository/mocks"
	"erajaya/internal/service"
	"erajaya/pkg/utils"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddProduct(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(mocks.ProductRepository)
	svc := service.NewService(mockRepo)

	t.Run("Invalid name (too short)", func(t *testing.T) {
		p := &model.Product{Name: "a"}
		err := svc.AddProduct(ctx, p)
		assert.ErrorIs(t, err, utils.ErrInvalidName)
	})

	t.Run("Invalid name (non-alphanumeric)", func(t *testing.T) {
		p := &model.Product{Name: "Invalid@Name"}
		err := svc.AddProduct(ctx, p)
		assert.ErrorIs(t, err, utils.ErrNameAlphaNumeric)
	})

	t.Run("Duplicate product", func(t *testing.T) {
		p := &model.Product{Name: "ValidName"}
		mockRepo.On("ExistsByName", "ValidName").Return(true, nil).Once()

		err := svc.AddProduct(ctx, p)
		assert.ErrorIs(t, err, utils.ErrProductExists)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Description too long", func(t *testing.T) {
		p := &model.Product{
			Name:        "ValidName",
			Description: string(make([]byte, utils.ProductDescriptionMaxLength+1)),
		}

		mockRepo.On("ExistsByName", "ValidName").Return(false, nil).Once()

		err := svc.AddProduct(ctx, p)
		assert.ErrorIs(t, err, utils.ErrDescriptionTooLong)

		mockRepo.AssertExpectations(t)
	})

	t.Run("Success", func(t *testing.T) {
		p := &model.Product{Name: "ValidName", Description: "A description"}

		mockRepo.On("ExistsByName", "ValidName").Return(false, nil).Once()
		mockRepo.On("Create", p).Return(nil).Once()
		mockRepo.On("DeleteCachedProducts", ctx).Return(nil).Once()

		err := svc.AddProduct(ctx, p)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestListProduct(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(mocks.ProductRepository)
	svc := service.NewService(mockRepo)

	validOrder := []utils.OrderParam{{Key: "price", Direction: utils.ASC}}
	invalidKeyOrder := []utils.OrderParam{{Key: "invalid", Direction: utils.ASC}}
	invalidDirOrder := []utils.OrderParam{{Key: "price", Direction: "sideways"}}

	mockProducts := []model.Product{{ID: 1, Name: "Test"}}

	t.Run("Invalid sort field", func(t *testing.T) {
		_, err := svc.ListProduct(ctx, invalidKeyOrder)
		assert.ErrorIs(t, err, utils.ErrInvalidSortField)
	})

	t.Run("Invalid sort direction", func(t *testing.T) {
		_, err := svc.ListProduct(ctx, invalidDirOrder)
		assert.ErrorIs(t, err, utils.ErrInvalidSortDirection)
	})

	t.Run("Return from cache", func(t *testing.T) {
		mockRepo.On("GetCachedProducts", ctx, "price:asc").Return(mockProducts, nil).Once()

		products, err := svc.ListProduct(ctx, validOrder)
		assert.NoError(t, err)
		assert.Equal(t, mockProducts, products)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Return from DB when cache miss", func(t *testing.T) {
		mockRepo.On("GetCachedProducts", ctx, "price:asc").Return(nil, errors.New("cache miss")).Once()
		mockRepo.On("GetAllByOrder", validOrder).Return(mockProducts, nil).Once()
		mockRepo.On("SetCachedProducts", ctx, "price:asc", mockProducts).Return(nil).Once()

		products, err := svc.ListProduct(ctx, validOrder)
		assert.NoError(t, err)
		assert.Equal(t, mockProducts, products)
		mockRepo.AssertExpectations(t)
	})
}
