package repository

import (
	"context"
	"database/sql"

	"newview-backend-go/app/domain/object"
)

type Account interface {
	// Creat new user accpunt.
	CreatNewUser(ctx context.Context, account *object.Account) (sql.Result, error)

	// Fetch account which has spwcified id
	FindById(ctx context.Context, id int64) (*object.Account, error)
}
