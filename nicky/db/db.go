package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type DbEngine struct {
	*xorm.Engine
}

func NewDbEngine(dbHost, user, pass, dbName string) (*DbEngine, error) {
	engine, err := xorm.NewEngine("mysql",
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", user, pass, dbHost, dbName))
	if err != nil {
		return nil, err
	}
	if err := engine.Ping(); err != nil {
		return nil, err
	}

	return &DbEngine{engine}, nil
}
