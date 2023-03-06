package apis

import (
	"errors"
	"ethereum_parser/config"
	"fmt"
	"log"
	"net/http"
	"os"
)

func Init() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", getPing)
	mux.HandleFunc("/v1/eth-mainnet/blocknumber/current", getCurrentBlock)
	mux.HandleFunc("/v1/eth-mainnet/blocknumber/info", getBlockInfo)
	mux.HandleFunc("/v1/eth-mainnet/subscribe", subscribe)
	mux.HandleFunc("/v1/eth-mainnet/transactions", getTransactionsByAddress)

	err := http.ListenAndServe(fmt.Sprintf(":%s", config.Conf.ServerPort), mux)

	if errors.Is(err, http.ErrServerClosed) {
		log.Printf("server closed\n")
	} else if err != nil {
		log.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
