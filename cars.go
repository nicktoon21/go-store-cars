package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Car struct {
	ID    int    `json:"id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Owner string `json:"owner"`
	Year  int    `json:"year"`
}

func routeGetCars(c *gin.Context) {
	c.JSON(http.StatusOK, Cars)

	c.Done()
}

func routePostCars(c *gin.Context) {
	var cars Car

	err := c.Bind(&cars)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message_error": "Não foi possível receber o payload",
		})
		return
	}
	cars.ID = len(Cars) + 1
	Cars = append(Cars, cars)

	c.JSON(http.StatusCreated, cars)
}
