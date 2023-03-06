package dal

import (
	"encoding/json"
	"ethereum_parser/models"
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"strconv"
)

const (
	lastPopulatedBlock = "last_populated_block"
)

type RedisCache struct {
	Client *redis.Client
}

func NewRedis(host, port, password string) RedisCache {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       0,
	})

	return RedisCache{
		Client: client,
	}
}

func (c *RedisCache) GetTransactionList(address string, offset, count int64) ([]models.Transaction, int64, error) {
	var total int64
	transactions := make([]models.Transaction, 0)
	list, err := c.GetList(address, offset, count)
	if err != nil {
		return transactions, total, err
	}
	for _, l := range list {
		var t models.Transaction
		err = json.Unmarshal([]byte(l), &t)
		if err != nil {
			log.Printf("error in unmarshaling, err:%v", err.Error())
			continue
		}
		transactions = append(transactions, t)
	}
	total, err = c.GetListLen(address)
	if err != nil {
		log.Printf("unable to get total transactions, err:%v", err.Error())
	}
	return transactions, total, nil
}

//
//func (c *RedisCache) SaveTransactions(key string, transactions []models.Transaction) error {
//	//total, err := c.GetListLen(key)
//	//if err != nil {
//	//	return err
//	//}
//	//val, err := json.Marshal(transactions)
//	//if err != nil {
//	//	fmt.Printf("error in marshalling transaction, err:%v", err.Error())
//	//	return err
//	//}
//	//if total != 0 {
//	//	_, err = c.AddToList(key, string(val))
//	//	if err != nil {
//	//		fmt.Printf("error in saving transaction, err:%v", err.Error())
//	//		return err
//	//	}
//	//	return nil
//	//}
//	_, err := c.CreateAndAddToList(key, transactions)
//	if err != nil {
//		fmt.Printf("error in saving transaction, err:%v", err.Error())
//		return err
//	}
//	return nil
//}

func (c *RedisCache) SaveTransactions(key string, transactions []models.Transaction) error {
	transactionsStr := make([]string, 0)
	for _, t := range transactions {
		val, err := json.Marshal(t)
		if err != nil {
			log.Printf("error in marshalling transaction, err:%v", err.Error())
			return err
		}
		transactionsStr = append(transactionsStr, string(val))
	}

	_, err := c.CreateAndAddToList(key, transactionsStr)
	if err != nil {
		log.Printf("error in saving transaction, err:%v", err.Error())
		return err
	}
	c.TrimList(key, 100)
	return nil
}

func (c *RedisCache) GetLastPopulatedBlock() (int64, error) {
	str, err := c.Client.Get(lastPopulatedBlock).Result()
	if err != nil {
		log.Printf("error in getting last populated block, err:%v", err.Error())
		return 0, err
	}
	v, err := strconv.Atoi(str)
	if err != nil {
		log.Printf("error in converting str to int, err:%v", err.Error())
		return 0, err
	}
	return int64(v), nil
}

func (c *RedisCache) SaveCurrentPopulatedBlock(val int64) error {
	_, err := c.Client.Set(lastPopulatedBlock, val, 0).Result()
	if err != nil {
		log.Printf("error in setting last populated block, err:%v", err.Error())
		return err
	}
	return nil
}

func (c *RedisCache) CreateAndAddToList(key string, val interface{}) (int64, error) {
	return c.Client.LPush(key, val).Result()
}

func (c *RedisCache) AddToList(key string, val interface{}) (int64, error) {
	return c.Client.LPushX(key, val).Result()
}

func (c *RedisCache) GetList(key string, offset, total int64) ([]string, error) {
	return c.Client.LRange(key, offset, offset+total-1).Result()
}

func (c *RedisCache) GetListLen(key string) (int64, error) {
	return c.Client.LLen(key).Result()
}

func (c *RedisCache) TrimList(key string, len int64) {
	c.Client.LTrim(key, 0, len-1)
}
