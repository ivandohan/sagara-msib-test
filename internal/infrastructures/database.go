package infrastructures

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func NewDatabase() (db *sql.DB, err error) {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err = sql.Open("postgres", dsn)

	log.Printf("Connecting to DB: host=%s port=%s user=%s dbname=%s",
		dbHost, dbPort, dbUser, dbName)

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	if err := ensureTableExists(db); err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected to PostgreSQL!")
	return db, nil
}

func ensureTableExists(db *sql.DB) error {
	query := `
    CREATE TABLE IF NOT EXISTS baju (
        id SERIAL PRIMARY KEY,
		nama VARCHAR(100),
        brand VARCHAR(100),
        warna VARCHAR(50),
        ukuran VARCHAR(10),
        harga NUMERIC(10, 2),
        stok INT
    );`

	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to ensure table exists: %w", err)
	}

	fmt.Println("Table 'baju' checked and created if not exists.")
	return nil
}
