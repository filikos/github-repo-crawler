package db

import (
	"database/sql"
	"fmt"
	"os"
	"time"
	"workspace-go/github-repo-crawler/model"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/prometheus/common/log"
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
	GetRecentRepositories(username string) (model.DBRepositories, error)
	ReplaceRecentRepositories(username string, repos model.DBRepositories) error
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

			dbPool.SetMaxOpenConns(20)
			dbPool.SetMaxIdleConns(5)

			fmt.Println("Database connection established")
			return &Database{dbPool}, nil
		}

		fmt.Printf("Connecting to DB try: %v", i+1)
		time.Sleep(retryInterval * time.Second)
	}

	return nil, fmt.Errorf("Failed to connect database. Server will be shut down. Error: %v", err)
}

func (db *Database) GetRecentRepositories(username string) (model.DBRepositories, error) {

	sqlStatement := `SELECT * FROM Repositories WHERE username = $1`
	rows, err := db.Conn.Query(sqlStatement, username)
	if err != nil {
		fmt.Printf("Query: Unable to get repositories: %v", err)
		return nil, err
	}

	defer rows.Close()

	var repositories model.DBRepositories
	for rows.Next() {

		var repository model.DBRepository
		err := rows.Scan(&repository.ID, &repository.Username, &repository.Name)
		if err != nil {
			log.Info(fmt.Sprintf("Unable to scan row:%v", err))
			continue
		}

		repositories = append(repositories, repository)
	}

	return repositories, nil
}

func (db *Database) ReplaceRecentRepositories(username string, repos model.Repositories) error {

	sqlDeleteStatement := `DELETE FROM Repositories;`
	_, err := db.Conn.Exec(sqlDeleteStatement)
	if err != nil {
		return fmt.Errorf("unable to remove last entries for replacing recent repositories: %v", err)
	}

	sqlStatement := `INSERT INTO Repositories(id, username, name) VALUES($1, $2, $3);`
	for _, repo := range repos {
		row, err := db.Conn.Query(sqlStatement, repo.ID, username, repo.Name)
		if err != nil {
			return fmt.Errorf("unable to replace recent repositories: %v", err)
		}
		defer row.Close()
	}

	return nil
}
