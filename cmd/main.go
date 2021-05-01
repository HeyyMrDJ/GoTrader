package main

import (
	"GoTrader/pkg/http/http_server"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	http_server.StartServer()

}
