package http

import (
	"fmt"
	"log"
	"net/http"
)

func StartServer() {
	fileServer := http.FileServer(http.Dir("../web/static/"))
	http.Handle("/", fileServer)
	http.HandleFunc("/trades/", tradeHandler)
	http.HandleFunc("/trades/newtrade", posttradeHandler)
	http.HandleFunc("/trades/deletetrade", deletetradeHandler)
	http.HandleFunc("/trades/updatetrade", updatetradeHandler)
	http.HandleFunc("/dashboard/", dashHandler)
	http.HandleFunc("/home/", homeHandler)
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
