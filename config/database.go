package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Connect membuat koneksi ke database MySQL dan memastikan koneksinya dapat digunakan.
func Connect() (*sql.DB, error) {
	user := "root"
	password := "@Tsan020922"
	host := "127.0.0.1"
	port := "3306"
	database := "beverages_db"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
