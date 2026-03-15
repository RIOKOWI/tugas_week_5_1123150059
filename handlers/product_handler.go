package handlers

import (
	"net/http"
	"strconv"

	"github.com/RIOKOWI/tugas_week_5_1123150059/models"
	"github.com/RIOKOWI/tugas_week_5_1123150059/services"
	"github.com/gin-gonic/gin"
)


type ProductHandler struct {
	productService *services.ProductService
}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{productService: services.NewProductService()}
}

func (h *ProductHandler) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	category := c.Query("category")
	products, total, err := h.productService.GetAll(page, limit, category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false, "message": "Gagal mengambil data produk",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": products,
		"meta": gin.H{
			"total": total,
			"page": page,
			"limit": limit,
			"per_page": limit,
		},
	})
}

