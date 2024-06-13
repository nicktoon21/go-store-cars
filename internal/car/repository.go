package car

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/nicktoon21/go-store-cars/internal/domain"
)

type Repository interface {
	Create(ctx context.Context, c domain.Car) (domain.Car, error)
}

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(ctx context.Context, c domain.Car) (domain.Car, error) {
	return c, nil
}
