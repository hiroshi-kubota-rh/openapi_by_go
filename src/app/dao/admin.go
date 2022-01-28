package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"newview-backend-go/app/domain/object"
	"newview-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	// Implementation for repository.Account
	admin struct {
		db *sqlx.DB
	}
)

// Create accout repository
func NewAdmin(db *sqlx.DB) repository.Admin {
	return &admin{db: db}
}

//FindById : idからユーザーを取得
func (r *admin) FindById(ctx context.Context, id int64) (*object.Admin, error) {
	entity := new(object.Admin)
	err := r.db.QueryRowxContext(ctx, "select * from users where id = ?", id).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("%w", err)
	}
	return entity, nil
}

func (r *admin) FindByUsername(ctx context.Context, username string) (*object.Admin, error) {
	entity := new(object.Admin)
	err := r.db.QueryRowxContext(ctx, "select * from users where name = ?", username).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("%w", err)
	}
	return entity, nil
}
