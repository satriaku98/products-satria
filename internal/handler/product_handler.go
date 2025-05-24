package handler

import (
	"net/http"
	"strings"

	"erajaya/internal/model"
	"erajaya/internal/service"
	"erajaya/pkg/utils"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service service.Service
}

func NewProductHandler(svc service.Service) *ProductHandler {
	return &ProductHandler{service: svc}
}

// AddProduct
// @Summary Add a new product
// @Description Add a new product
// @Schemes
// @Tags Product
// @Accept json
// @Produce json
// @Param product body request.ProductRequest true "Product details"
// @Success 201 {object} model.Product
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /products [post]
func (h *ProductHandler) AddProduct(c *gin.Context) {
	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.AddProduct(c.Request.Context(), &product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, product)
}

// ListProduct
// @Summary List products
// @Description Get a list of products
// @Schemes
// @Tags Product
// @Accept json
// @Produce json
// @Param sort query string false "Sort parameters in the format 'column:direction[,column2:direction]'. Example: created_at:desc,price:asc"
// @Success 200 {array} model.Product
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /products [get]
func (h *ProductHandler) ListProduct(c *gin.Context) {
	sortParam := c.DefaultQuery("sort", "created_at:desc")
	sortParts := strings.Split(sortParam, ",")

	var orderClauses []utils.OrderParam
	for _, part := range sortParts {
		items := strings.Split(part, ":")
		if len(items) != 2 {
			c.JSON(400, gin.H{"error": "invalid sort format"})
			return
		}

		column := items[0]
		direction := strings.ToLower(items[1])

		orderClauses = append(orderClauses, utils.OrderParam{
			Key:       column,
			Direction: direction,
		})
	}

	products, err := h.service.ListProduct(c.Request.Context(), orderClauses)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}
