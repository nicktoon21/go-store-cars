package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
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
	r.rg = r.eng.Group("")
}
