package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/nicktoon21/go-store-cars/cmd/handler"
	"github.com/nicktoon21/go-store-cars/internal/brand"
	"github.com/nicktoon21/go-store-cars/internal/car"
	"github.com/nicktoon21/go-store-cars/internal/model"
)

type Router interface {
	MapRoutes()
}

type router struct {
	eng *gin.Engine
	rg  *gin.RouterGroup
	db  *pgxpool.Pool
}

func NewRouter(eng *gin.Engine, db *pgxpool.Pool) Router {
	return &router{
		eng: eng,
		db:  db,
	}
}

func (r *router) MapRoutes() {
	r.setGroup()
	r.createCarRoutes()
	r.createBrandRoutes()
	r.createModelRoutes()
}

func (r *router) setGroup() {
	r.rg = r.eng.Group("")
}

func (r *router) createCarRoutes() {
	repository := car.NewRepository(r.db)
	service := car.NewService(repository)
	handler := handler.NewCarController(service)

	r.rg.POST("/car", handler.Create())
}

func (r *router) createBrandRoutes() {
	repository := brand.NewRepository(r.db)
	service := brand.NewService(repository)
	handler := handler.NewBrandController(service)

	r.rg.POST("/brand", handler.Create())
}

func (r *router) createModelRoutes() {
	repository := model.NewRepository(r.db)
	service := model.NewService(repository)
	handler := handler.NewModelController(service)

	r.rg.POST("/model", handler.Create())
}
