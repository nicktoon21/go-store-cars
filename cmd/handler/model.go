package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/nicktoon21/go-store-cars/internal/model"
)

type ModelController interface {
	Create() gin.HandlerFunc
}

type modelController struct {
	service model.Service
}

func NewModelController(service model.Service) ModelController {
	return &modelController{
		service: service,
	}
}

func (c *modelController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
