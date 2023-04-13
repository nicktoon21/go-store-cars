package main

import (
	"github.com/gin-gonic/gin"
)

var Cars = []Car{
	{ID: 1, Brand: "Volks", Model: "Gol", Owner: "Alonso", Year: 1998},
	{ID: 2, Brand: "Chevrolet", Model: "Celta", Owner: "Jussara", Year: 2005},
	{ID: 3, Brand: "Fiat", Model: "Palio", Owner: "Cleuda", Year: 2003},
}

func main() {
	service := gin.Default()
	getRoutes(service)
	service.Run()
}
