package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/nicktoon21/go-store-cars/cmd/handler"
	"github.com/nicktoon21/go-store-cars/internal/car"
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
