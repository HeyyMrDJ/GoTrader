package main

import (
	"GoTrader/pkg/database/db_functions"
	"GoTrader/pkg/http/http_server"
)

func main() {
	db_functions.CreateTable()
	http_server.StartServer()

}
