package database

import (
	"database/sql"
	"fmt"

	"github.com/ilhamnyto/echo-fw/config"
	_ "github.com/lib/pq"
)

func ConnectDB() (*Database) {
	var (
		db_host = config.GetString(config.POSTGRES_HOST)
		db_port = config.GetString(config.POSTGRES_PORT)
		db_user = config.GetString(config.POSTGRES_USER)
		db_password = config.GetString(config.POSTGRES_PASSWORD)
		db_name = config.GetString(config.POSTGRES_DB)
	)

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		db_host, db_port, db_user, db_password, db_name,
	)

	dbsql, err := sql.Open("postgres", dsn)

	if err != nil {
		panic(err)
	}

	if err := dbsql.Ping(); err != nil {
		panic(err)
	}

	db := NewDatabase()
	db = db.SetDatabase(dbsql)
	
	return db
}