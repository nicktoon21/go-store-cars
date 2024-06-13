package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/nicktoon21/go-store-cars/internal/car"
)

type CarController interface {
	Create() gin.HandlerFunc
}

type carController struct {
	service car.Service
}

func NewCarController(service car.Service) CarController {
	return &carController{
		service: service,
	}
}

func (c *carController) Create() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
