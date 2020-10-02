package db

import (
	"database/sql"
	"fmt"
	"os"
	"time"
	"workspace-go/github-repo-crawler/model"

	"github.com/joho/godotenv"
)

const (
	connTimeoutSec = 5

	dbConnectionAttemts = 5
	retryInterval       = 2
)

type Database struct {
	Conn *sql.DB
}

type DBConnector interface {
	InitDB(configPath string) (*Database, error)
	GetRecentRepositories(username string) (model.Repositories, error)
	AddRecentRepositories(username string, repos model.Repositories) error
}

func InitDB(configPath string) (*Database, error) {

	err := godotenv.Load(configPath)
	if err != nil {
		return nil, fmt.Errorf("Unable to load DB configuration %v", err)
	}

	dbName := os.Getenv("POSTGRES_DB")
	dbUser := os.Getenv("POSTGRES_USER")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")

	var dbPool *sql.DB
	url := fmt.Sprintf("user=%v dbname=%v host=%v port=%v password=%v  connect_timeout=%v sslmode=disable", dbUser, dbName, dbHost, dbPort, dbPassword, connTimeoutSec)

	for i := 0; i < dbConnectionAttemts; i++ {

		dbPool, err = sql.Open("postgres", url)
		if err != nil {
			return nil, err
		}

		err = dbPool.Ping()
		if err == nil {

			dbPool.SetMaxOpenConns(7)
			dbPool.SetMaxIdleConns(5)

			fmt.Println("Database connection established")
			return &Database{dbPool}, nil
		}

		fmt.Printf("Connecting to DB try: %v", i+1)
		time.Sleep(retryInterval * time.Second)
	}

	return nil, fmt.Errorf("Failed to connect database. Server will be shut down. Error: %v", err)
}

func (db *Database) GetRecentRepositories(username string) (model.Repositories, error) {

	// TODO:
	return nil, nil
}

func (db *Database) AddRecentRepositories(username string, repos model.Repositories) error {

	// TODO:
	return nil
}
