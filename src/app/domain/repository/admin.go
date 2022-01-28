package repository

import (
	"context"

	"newview-backend-go/app/domain/object"
)

type Admin interface {
	// Fetch account which has spwcified id
	FindById(ctx context.Context, id int64) (*object.Admin, error)

	FindByUsername(ctx context.Context, username string) (*object.Admin, error)
}
