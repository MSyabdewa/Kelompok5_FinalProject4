package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/MSyabdewa/Kelompok5_FinalProject4/infra/config"
	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	err error
)

func handleDatabaseConnection() {
	appConfig := config.GetAppConfig()
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		appConfig.DBHost, appConfig.DBPort, appConfig.DBUser, appConfig.DBPassword, appConfig.DBName,
	)

	db, err = sql.Open(appConfig.DBDialect, psqlInfo)

	if err != nil {
		log.Panic("error occured while trying to validate database arguments:", err)
	}

	err = db.Ping()

	if err != nil {
		log.Panic("error occured while trying to connect to database:", err)
	}

}

func createTables() {

	usersTable := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			full_name text NOT NULL,
			email text NOT NULL UNIQUE,
			password text NOT NULL,
			role text NOT NULL DEFAULT 'customer',
			balance integer NOT NULL DEFAULT 0,
			created_at timestamptz DEFAULT current_timestamp,
	 		updated_at timestamptz DEFAULT current_timestamp
		)
		`

	categoriesTable := `
		CREATE TABLE IF NOT EXISTS categories (
			id SERIAL PRIMARY KEY,
			type text NOT NULL,
			sold_product_amount integer DEFAULT 0,
			created_at timestamptz DEFAULT current_timestamp,
	 		updated_at timestamptz DEFAULT current_timestamp
		)
	`

	productsTable := `
		CREATE TABLE IF NOT EXISTS products (
			id SERIAL PRIMARY KEY,
			title text NOT NULL,
			price integer NOT NULL,
			stock integer NOT NULL,
			category_id integer REFERENCES categories(id)
				ON DELETE CASCADE,
			created_at timestamptz DEFAULT current_timestamp,
	 		updated_at timestamptz DEFAULT current_timestamp
		)
	`

	transactionHistoriesTable := `
		CREATE TABLE IF NOT EXISTS transaction_histories (
			id SERIAL PRIMARY KEY,
			product_id integer REFERENCES products(id)
				ON DELETE CASCADE,
			user_id integer REFERENCES users(id)
				ON DELETE CASCADE,
			quantity integer NOT NULL,
			total_price integer NOT NULL,
			created_at timestamptz DEFAULT current_timestamp,
	 		updated_at timestamptz DEFAULT current_timestamp
		)
	`
	_, err = db.Exec(usersTable)
	if err != nil {
		log.Panic("error occured while trying to create order table:", err)
	}
	_, err = db.Exec(categoriesTable)
	if err != nil {
		log.Panic("error occured while trying to create order table:", err)
	}
	_, err = db.Exec(productsTable)
	if err != nil {
		log.Panic("error occured while trying to create order table:", err)
	}
	_, err = db.Exec(transactionHistoriesTable)
	if err != nil {
		log.Panic("error occured while trying to create order table:", err)
	}

}

func InitiliazeDatabase() {
	handleDatabaseConnection()
	createTables()
}

func GetDatabaseInstance() *sql.DB {
	if db == nil {
		log.Panic("database instance is still nill!!!")
	}

	return db
}
