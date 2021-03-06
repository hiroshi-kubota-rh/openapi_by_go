package dao

import (
	"newview-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	// DAO interface
	Dao interface {
		Admin() repository.Admin

		ImpressionTag() repository.ImpressionTag
	}

	// Implementation for DAO
	dao struct {
		db *sqlx.DB
	}
)

// Create DAO
func New(config DBConfig) (Dao, error) {
	db, err := initDb(config)
	if err != nil {
		return nil, err
	}

	return &dao{db: db}, nil
}

func (d *dao) Admin() repository.Admin {
	return NewAdmin(d.db)
}

func (d *dao) ImpressionTag() repository.ImpressionTag {
	return NewImpressionTag(d.db)
}
