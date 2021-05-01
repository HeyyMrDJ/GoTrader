package http_server

import (
	"GoTrader/pkg/http/http_handler"
	"fmt"
	"log"
	"net/http"
)

func StartServer() {
	fileServer := http.FileServer(http.Dir("../web/static/"))
	http.Handle("/", fileServer)
	http.HandleFunc("/trades/", http_handler.TradeHandler)
	http.HandleFunc("/trades/newtrade", http_handler.PosttradeHandler)
	http.HandleFunc("/trades/deletetrade", http_handler.DeletetradeHandler)
	http.HandleFunc("/trades/updatetrade", http_handler.UpdatetradeHandler)
	http.HandleFunc("/dashboard/", http_handler.DashHandler)
	http.HandleFunc("/home/", http_handler.HomeHandler)
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
