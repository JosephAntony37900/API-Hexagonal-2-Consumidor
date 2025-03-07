package helpers

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func NewMySQLConnection() (*sql.DB, error){
	dsn := "root:chocolate200201614@tcp(127.0.0.1:3306)/API_Productor"
	db, err := sql.Open("mysql", dsn)
	if err != nil{
		return nil, fmt.Errorf("error conectando con MySQÑ: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pineando MySQL: %w", err)
	}
	return db, nil
}