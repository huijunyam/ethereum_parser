package dal

import "ethereum_parser/models"

type DataStorage interface {
	GetTransactionList(string, int64, int64) ([]models.Transaction, int64, error)
	SaveTransactions(string, []models.Transaction) error
	GetLastPopulatedBlock() (int64, error)
	SaveCurrentPopulatedBlock(int64) error
}

type Storage struct {
	DbClient RedisCache
}

var StorageClient Storage

func Init() {
	redisStorage := NewRedis("localhost", "6379", "")
	StorageClient = Storage{DbClient: redisStorage}
}

func (a *Storage) GetTransactionList(address string, offset, count int64) ([]models.Transaction, int64, error) {
	return a.DbClient.GetTransactionList(address, offset, count)
}

func (a *Storage) SaveTransactions(address string, transactions []models.Transaction) error {
	return a.DbClient.SaveTransactions(address, transactions)
}

func (a *Storage) GetLastPopulatedBlock() (int64, error) {
	return a.DbClient.GetLastPopulatedBlock()
}

func (a *Storage) SaveCurrentPopulatedBlock(blockNum int64) error {
	return a.DbClient.SaveCurrentPopulatedBlock(blockNum)
}
