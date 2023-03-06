package main

import (
	"ethereum_parser/apis"
	"ethereum_parser/cron"
	"ethereum_parser/dal"
)

func main() {
	dal.Init()
	cron.Init()
	apis.Init()
}
