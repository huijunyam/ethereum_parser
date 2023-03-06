package main

import (
	"ethereum_parser/apis"
	"ethereum_parser/config"
	"ethereum_parser/cron"
	"ethereum_parser/dal"
)

func main() {
	config.Init()
	dal.Init()
	cron.Init()
	apis.Init()
}
