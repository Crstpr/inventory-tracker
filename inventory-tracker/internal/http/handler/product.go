package handler

import (
	"errors"
	"net/http"

	"github.com/Crstpr/inventory-tracker/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type ProductHandler struct {
	productService service.ProductServiceInterface
}

func NewProductHandler(
	productService service.ProductServiceInterface,
) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

func (h *ProductHandler) ListProducts(c *gin.Context) {
	products, err := h.productService.ListProducts(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get products",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": products,
	})
}

func (h *ProductHandler) GetProductByID(c *gin.Context) {
	idParam := c.Param("id")

	parsedID, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "invalid product id",
		})
		return
	}

	productID := pgtype.UUID{
		Bytes: parsedID,
		Valid: true,
	}

	product, err := h.productService.GetProductByID(
		c.Request.Context(),
		productID,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "product not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get product",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": product,
	})
}