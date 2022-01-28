package repository

import (
	"context"
	"database/sql"

	"newview-backend-go/app/domain/object"
)

type ImpressionTag interface {
	// Creat new user impression tag.
	CreatNewImpressionTag(ctx context.Context, account *object.ImpressionTag) (sql.Result, error)

	// Fetch impression tag which has spwcified id
	FindById(ctx context.Context, id uint64) (*object.ImpressionTag, error)

	// Delete impression tag whick has spwcified id
	DeleteById(ctx context.Context, id uint64) (*object.ImpressionTag, error)
}
