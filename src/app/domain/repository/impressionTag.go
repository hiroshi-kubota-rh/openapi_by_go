package repository

import (
	"context"
	"database/sql"

	"newview-backend-go/app/domain/object"
)

type ImpressionTag interface {
	// Creat new user accpunt.
	CreatNewImpressionTag(ctx context.Context, account *object.ImpressionTag) (sql.Result, error)

	// Fetch account which has spwcified id
	FindById(ctx context.Context, id uint64) (*object.ImpressionTag, error)
}
