package repository

import (
	"database/sql"
	"fmt"

	"gopkg.in/gorp.v1"
)

var dbmap *gorp.DbMap

func init() {
	var err error
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		"admin", "password", "ulife-share-db:3306", "ulifeshare",
	)
	Db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	dbmap = &gorp.DbMap{Db: Db, Dialect: gorp.MySQLDialect{Engine: "InnoDB", Encoding: "utf8mb4"}}
}
