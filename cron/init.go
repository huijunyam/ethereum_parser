package cron

import (
	"ethereum_parser/dal"
	"ethereum_parser/helper"
	"ethereum_parser/json_rpc"
	"ethereum_parser/models"
	"ethereum_parser/subscription"
	"log"
	"sync"
	"time"
)

const (
	pollingBlockKey = "polling_block"
	user            = "cron_polling"
)

var AddressSubscription map[string]*subscription.AddressMonitor
var AddressSubscriptionLock sync.RWMutex

func Init() {
	AddressSubscription = make(map[string]*subscription.AddressMonitor)
	go func() {
		tick := time.Tick(10 * time.Second)
		for range tick {
			pollBlockInfo()
		}
		time.Sleep(10 * time.Second)
	}()
}

func pollBlockInfo() {
	ok, err := dal.StorageClient.DbClient.Client.SetNX(pollingBlockKey, 1, time.Second*10).Result()
	if err != nil {
		log.Printf("error in setting polling block key, err:%v", err.Error())
		return
	}
	if !ok {
		log.Printf("polling in progress, process aborted")
		return
	}
	log.Printf("getting curr block")
	startBlock, endBlock, err := getStartAndEndBlock()
	if err != nil {
		return
	}
	dal.StorageClient.DbClient.SaveCurrentPopulatedBlock(endBlock)
	populateTransaction(startBlock, endBlock)
}

func getStartAndEndBlock() (int64, int64, error) {
	var startBlock, endBlock int64
	currentBlockHex, err := json_rpc.GetCurrentBlock(user)
	if err != nil {
		log.Printf("error in getting current block, err:%v", err.Error())
		return startBlock, endBlock, err
	}
	endBlock, err = helper.ConvertHexStringToInt(*currentBlockHex)
	if err != nil {
		log.Printf("unable to convert to int, err:%v", err.Error())
		return startBlock, endBlock, err
	}

	startBlock, err = dal.StorageClient.DbClient.GetLastPopulatedBlock()
	if err != nil {
		log.Printf("last populated block is not found in redis, err:%v", err.Error())
	}
	if startBlock == 0 {
		startBlock = endBlock - 5
	} else {
		startBlock += 1
	}
	return startBlock, endBlock, nil
}

func populateTransaction(startBlock, endBlock int64) {
	log.Printf("populating process start for start:%d, end:%d", startBlock, endBlock)
	for i := startBlock; i <= endBlock; i++ {
		hexStr := helper.ConvertIntToHexString(i)
		blockInfo, err := json_rpc.GetBlockByNumber(user, hexStr, true)
		if err != nil {
			log.Printf("error in getting block info for block: %s, err:%v", hexStr, err.Error())
			//todo can consider handle in dead letter queue for reprocessing
			continue
		}
		saveTransactions(blockInfo.Transactions)
	}
	log.Printf("populating process end for start:%d, end:%d", startBlock, endBlock)
}

func saveTransactions(t []models.Transaction) {
	aToTMap := make(map[string][]models.Transaction)
	AddressSubscriptionLock.RLock()
	defer AddressSubscriptionLock.RUnlock()

	for _, l := range t {
		if _, ok := aToTMap[l.From]; !ok {
			aToTMap[l.From] = make([]models.Transaction, 0)
		}
		if _, ok := aToTMap[l.To]; !ok {
			aToTMap[l.To] = make([]models.Transaction, 0)
		}
		aToTMap[l.From] = append(aToTMap[l.From], l)
		aToTMap[l.To] = append(aToTMap[l.To], l)

		if m, ok := AddressSubscription[l.From]; ok {
			m.SetTransaction(l)
		}
		if m, ok := AddressSubscription[l.To]; ok {
			m.SetTransaction(l)
		}
	}
	for k, v := range aToTMap {
		dal.StorageClient.DbClient.SaveTransactions(k, v)
	}
}
