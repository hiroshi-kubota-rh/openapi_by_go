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
	impressionTag struct {
		db *sqlx.DB
	}
)

// Create accout repository
func NewImpressionTag(db *sqlx.DB) repository.ImpressionTag {
	return &impressionTag{db: db}
}

// CreatNewUser　: ユーザーを登録
func (r *impressionTag) CreatNewImpressionTag(ctx context.Context, impressionTag *object.ImpressionTag) (sql.Result, error) {
	result, err := r.db.ExecContext(ctx, `insert into impression_tags (name) values (?)`, impressionTag.Name)
	if err != nil {
		//usernameがuniquでない場合に警告を返すことが必要
		return nil, fmt.Errorf("%v", err)
	}
	return result, nil
}

//FindById : idからimpressionを取得
func (r *impressionTag) FindById(ctx context.Context, id uint64) (*object.ImpressionTag, error) {
	entity := new(object.ImpressionTag)
	err := r.db.QueryRowxContext(ctx, "select * from impression_tags where id = ?", id).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("%w", err)
	}
	return entity, nil
}

func (r *impressionTag) DeleteById(ctx context.Context, id uint64) (*object.ImpressionTag, error) {
	entity := new(object.ImpressionTag)
	//deleteの命令はexec関数を用いるのが良いのでは？
	err := r.db.QueryRowContext(ctx, "delete from impression_tags where id = ?", id)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	return entity, nil
}
