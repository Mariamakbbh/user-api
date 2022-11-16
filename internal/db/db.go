package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

func Init(DB_User, DB_Pwd, DB_Host, DB_Port, DB_Name string) (*pgxpool.Pool, error) {
	db_url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", DB_User, DB_Pwd, DB_Host, DB_Port, DB_Name)

	// returns connection pool
	dbPool, err := pgxpool.Connect(context.Background(), db_url)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		log.Fatalln(err)
		return nil, err
	}

	// to close DB pool
	defer dbPool.Close()

	return dbPool, nil
}
