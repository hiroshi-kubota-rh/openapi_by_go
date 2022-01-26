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
	account struct {
		db *sqlx.DB
	}
)

// Create accout repository
func NewAccount(db *sqlx.DB) repository.Account {
	return &account{db: db}
}

// CreatNewUser　: ユーザーを登録
func (r *account) CreatNewUser(ctx context.Context, account *object.Account) (sql.Result, error) {
	result, err := r.db.ExecContext(ctx, `insert into account (username, password_hash) values (?, ?)`, account.Username, account.PasswordHash)
	if err != nil {
		//usernameがuniquでない場合に警告を返すことが必要
		return nil, fmt.Errorf("%v", err)
	}
	return result, nil
}

//FindById : idからユーザーを取得
func (r *account) FindById(ctx context.Context, id int64) (*object.Account, error) {
	entity := new(object.Account)
	err := r.db.QueryRowxContext(ctx, "select * from account where id = ?", id).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("%w", err)
	}
	return entity, nil
}
