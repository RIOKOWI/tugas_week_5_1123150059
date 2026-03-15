package handlers 

import(
	"github.com/RIOKOWI/tugas_week_5_1123150059/models"
	"github.com/RIOKOWI/tugas_week_5_1123150059/services"
)


type ProductHandler struct {
	productService *services.ProductService
}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{productService: services.NewProductService()}
}

