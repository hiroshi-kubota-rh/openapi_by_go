package app

import (
	"newview-backend-go/app/config"
	"newview-backend-go/app/dao"
)

type App struct {
	Dao dao.Dao
}

// Create dependency manager
func NewApp() (*App, error) {
	// panic if lacking something
	daoCfg := config.MySQLConfig()

	dao, err := dao.New(daoCfg)
	if err != nil {
		return nil, err
	}

	return &App{Dao: dao}, nil
}
