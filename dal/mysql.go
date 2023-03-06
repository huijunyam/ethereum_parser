package dal

import (
	"database/sql"
	"ethereum_parser/models"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Db struct {
	db *sql.DB
}

func NewMySql(username, password, host, port, databaseName string) (Db, error) {
	db, err := sql.Open("mysql", fmt.Sprintf(
		"%s:%s@%s:%s/%s",
		username, password, host, port, databaseName,
	))
	if err != nil {
		return Db{}, fmt.Errorf("connection error: sql.Open: %s:***@%s:%s/%s",
			username, host, port, databaseName)
	}
	defer db.Close()

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return Db{
		db: db,
	}, nil
}

// GetTransactionList function to be built when mysql db is used
func (d *Db) GetTransactionList(address string, offset, count int64) ([]models.Transaction, int64, error) {
	return nil, 0, nil
}

// SaveTransactions function to be built when mysql db is used
func (d *Db) SaveTransactions(key string, transactions []models.Transaction) error {
	return nil
}

// GetLastPopulatedBlock function to be built when mysql db is used
func (d *Db) GetLastPopulatedBlock() (int64, error) {
	return 0, nil
}

// SaveCurrentPopulatedBlock function to be built when mysql db is used
func (d *Db) SaveCurrentPopulatedBlock(val int64) error {
	return nil
}
