package app

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mmaruf23/golang-rest-api/config"
	"github.com/mmaruf23/golang-rest-api/internal/helper"
)

func NewMySQLConnection(config *config.AppConfig) *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.DBUser, config.DBPass, config.DBHost, config.DBPort, config.DBName)
	db, err := sql.Open("mysql", dsn)

	helper.PanicIfError(err)

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(5 * time.Minute)

	err = db.Ping()
	helper.PanicIfError(err)

	fmt.Println("Successfully connected to MySQL Database")

	return db
}
