package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/nicktoon21/go-store-cars/internal/brand"
)

type BrandController interface {
	Create() gin.HandlerFunc
}

type brandController struct {
	service brand.Service
}

func NewBrandController(service brand.Service) BrandController {
	return &brandController{
		service: service,
	}
}

func (c *brandController) Create() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
