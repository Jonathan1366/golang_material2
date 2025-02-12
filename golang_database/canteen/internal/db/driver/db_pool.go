	package driver

	import (
		"context"
		"fmt"
		"log"
		"time"

		"github.com/jackc/pgx/v5/pgxpool"
	)

	// Global pool variable
	var dbPool *pgxpool.Pool

	func InitDBPool() {

		urldb := "postgresql://postgres.bhbrhddnnkbdyveehwqk:tkym3fky12345@aws-0-ap-southeast-1.pooler.supabase.com:6543/postgres"

		// Parse the database URL
		config, err := pgxpool.ParseConfig(urldb)
		if err != nil {
			log.Fatalf("Unable to parse configuration: %v\n", err)
		}

		// Set pool configurations
		config.MaxConns = 100
		config.MinConns = 5
		config.MaxConnIdleTime = 10 * time.Minute

		// Create the pool
		dbPool, err = pgxpool.NewWithConfig(context.Background(), config)
		if err != nil {
			log.Fatalf("Unable to create connection pool: %v\n", err)
		}

		fmt.Println("Successfully connected to the database")	
	}

	// Function to close the pool, should be called when the application shuts down
	func CloseDBPool() {
		if dbPool != nil {
			dbPool.Close()
		}
	}

	func GetDbpool() *pgxpool.Pool  {
		return dbPool
	}



