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

func MigrateDB (db *sql.DB) {
	fmt.Println("Start Migrating DB.")
	var (
		queryCreateUserTable = `
			CREATE TABLE IF NOT EXISTS users (
				id serial primary key,
				username varchar(100) NOT NULL,
				first_name varchar(100),
				last_name varchar(100),
				email varchar(100) NOT NULL,
				phone_number varchar(100),
				location varchar(100),
				password varchar(100) NOT NULL,
				salt varchar(100) NOT NULL,
				created_at timestamp,
				updated_at timestamp
			)
		`

		queryCreatePostTable = `
			CREATE TABLE IF NOT EXISTS posts (
				id serial primary key,
				user_id int NOT NULL,
				body text NOT NULL,
				created_at timestamp,
				deleted_at timestamp,
				CONSTRAINT fk_posts
				FOREIGN KEY(user_id)
				REFERENCES users(id)
			)
		`
	)

	stmt, err := db.Prepare(queryCreateUserTable)

	if err != nil {
		panic(err)
	}

	if _, err := stmt.Exec(); err != nil {
		panic(err)
	}
	
	stmt, err = db.Prepare(queryCreatePostTable)
	
	if err != nil {
		panic(err)
	}

	if _, err := stmt.Exec(); err != nil {
		panic(err)
	}

	fmt.Println("Migrate Success.")
}