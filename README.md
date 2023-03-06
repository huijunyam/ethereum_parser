# Ethereum Parser

The goal is to create Ethereum blockchain parser that will allow to query transactions for subscribed addresses.

## Overview
The parser has a cron job that runs at an interval of 20 seconds to fetch all inbound/outbound transactions
in the blocks and store the transactions in in-memory storage. At the same time, each of the transaction will check 
`From` and `To` address. If the address are present in the subscription list, it can be linked up to the notification
service to inform users regarding the transaction. There are a couple of APIs that are available to interact with parser
as shown in [API Requests Samples](#api-requests-samples) section below. 

Due to the in-memory storage, all the data (transactions and subscription) are not persistence, but it is extensible to 
support other storage like relational database. Id currently are randomly generated via rand function, but it should 
eventually be replaced with user app-key for better usage monitoring and pub/sub messaging. Notifying user is currently
not implemented, it is just a simple logging now

## Prerequisites 
- local redis or remote redis server. Based on your redis, you can configure the host and port in `conf.yaml` accordingly. 
The default redis is `localhost:6379`

## Steps
1. Start your local redis server using `redis-server`
2. Start the server using `go run main.go`
3. The server port is configurable in `conf.yaml`, default port number is `3333`

## API Requests Samples
### Server health check
```bash
curl --location --request GET 'localhost:3333/ping'
```
### Get current block
```bash
curl --location --request GET 'localhost:3333/v1/eth-mainnet/blocknumber/current'
```
### Subscribe address for observer
```bash
curl --location --request POST 'localhost:3333/v1/eth-mainnet/subscribe' \
--header 'Content-Type: application/json' \
--data-raw '{"address": "INPUT ADDRESS HERE", "user":"INPUT USER HERE"}'
```
### Get list of inbound or outbound transactions based on address
```bash
curl --location --request POST 'localhost:3333/v1/eth-mainnet/transactions' \
--header 'Content-Type: application/json' \
--data-raw '{
    "address": "INPUT ADDRESS HERE",
    "page_info": {
        "page": 1,
        "page_size": 10
    }
}'
```