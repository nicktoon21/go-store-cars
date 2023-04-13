package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getRoutes(c *gin.Engine) *gin.Engine {
	c.GET("/home", routeStart)
	groupCars := c.Group("/cars")
	groupCars.GET("/", routeGetCars)
	groupCars.POST("/", routePostCars)

	return c
}

func routeStart(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "API is running!!!",
	})
}
