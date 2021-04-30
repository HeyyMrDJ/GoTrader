package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	//GetTradeInfo()
	//CreateTable()
	StartServer()

	//t := trade{name: "BTC", amount: 1, price: 30000}
	//fmt.Println(t)
}

//func helloHandler(w http.ResponseWriter, r *http.Request) {
//	if r.URL.Path != "/trades" {
//		http.Error(w, "404 not found.", http.StatusNotFound)
//		return
//	}
//
//	if r.Method != "GET" {
//		http.Error(w, "Method is not supported.", http.StatusNotFound)
//		return
//	}
//
//	db, err := sql.Open("sqlite3", "./trades.db")
//	if err != nil {
//		fmt.Println(err)
//	}
//	stmt, err2 := db.Query(`
//		SELECT Name, Amount, Price FROM trades
//	`)
//	if err2 != nil {
//		fmt.Println(err2)
//	}
//	var name string
//	var amount int
//	var price int
//	fmt.Println("Amount", "Name", "Price")
//	for stmt.Next() {
//		stmt.Scan(&name, &amount, &price)
//		fmt.Fprintf(w, strconv.Itoa(amount))
//		fmt.Fprintf(w, " ")
//		fmt.Fprintf(w, name)
//		fmt.Fprintf(w, " @ $")
//		fmt.Fprintf(w, strconv.Itoa(price))
//		fmt.Fprintf(w, "\n")
//		fmt.Println(name, amount)
//	}
//	db.Close()
//
//}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"./static/template/home/index.html",
		"./static/template/header.html",
		"./static/template/footer.html",
	}

	tmpl := template.Must(template.ParseFiles(files...))

	tmpl.Execute(w, nil)
}

func dashHandler(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"./static/template/dashboard/index.html",
		"./static/template/header.html",
		"./static/template/footer.html",
	}

	tmpl := template.Must(template.ParseFiles(files...))

	tmpl.Execute(w, nil)
}

func tradeHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "./trades.db")
	if err != nil {
		fmt.Println(err)
	}
	stmt, err2 := db.Query(`
		SELECT ID, Name, Amount, Price FROM trades
	`)
	if err2 != nil {
		fmt.Println(err2)
	}

	var tradeQuery Trade
	var listtrades []Trade
	fmt.Println("ID", "Amount", "Name", "Price")
	for stmt.Next() {
		stmt.Scan(&tradeQuery.ID, &tradeQuery.Name, &tradeQuery.Amount, &tradeQuery.Price)
		listtrades = append(listtrades, tradeQuery)
	}
	db.Close()

	files := []string{
		"./static/template/trades/index.html",
		"./static/template/header.html",
		"./static/template/footer.html",
		"./static/template/trades.html",
	}

	tmpl := template.Must(template.ParseFiles(files...))
	//trade1 := Trade{"BTC", 1, 30000}
	//trade2 := Trade{"BTC", 1, 60000}
	data := listtrades
	fmt.Println(data)

	tmpl.Execute(w, listtrades)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to GoTrader")
	fmt.Println("Endpoint hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	fmt.Println("Serving at http://localhost:10000")
	log.Fatal(http.ListenAndServe(":10000", nil))

}

func GetTradeInfo() (float64, float64) {
	fmt.Println("What are you trading?")
	var trade string
	fmt.Scanln(&trade)
	fmt.Println("Enter your entry price?")
	var price int
	fmt.Scanln(&price)
	fmt.Println("What is your Stoploss percentage?")
	var SLPercent float64
	fmt.Scanln(&SLPercent)
	PortValue := 24000.00
	TradeSize := TradeAmount(PortValue, SLPercent)
	fmt.Println("Your Trade size is:", TradeSize)
	return PortValue, SLPercent
}

func TradeAmount(PortValue, SLPercent float64) float64 {
	Risk := PortValue / 100
	TradeSize := (Risk / (SLPercent / 100))
	return TradeSize
}

func NewTrade() {
	db, err := sql.Open("sqlite3", "./trades.db")
	if err != nil {
		fmt.Println(err)
	}
	stmt, err2 := db.Prepare(`
		INSERT INTO trades (Name, Amount, Price) values (?,?,?)
	`)
	stmt.Exec("BTC", 5, 15000)
	if err2 != nil {
		fmt.Println(err2)
	}
	db.Close()
}

func CreateTable() {
	db, err := sql.Open("sqlite3", "./trades.db")
	if err != nil {
		fmt.Println(err)
	}
	stmt, err2 := db.Prepare(`
		CREATE TABLE "trades" (
			"ID"	INTEGER NOT NULL,
			"Name"	TEXT NOT NULL,
			"Amount"	INTEGER NOT NULL,
			"Price"	INTEGER NOT NULL,
			PRIMARY KEY("ID" AUTOINCREMENT)
		);
	`)
	stmt.Exec()
	if err2 != nil {
		fmt.Println(err2)
	}
	db.Close()
}

func GetTrades() {
	db, err := sql.Open("sqlite3", "./trades.db")
	if err != nil {
		fmt.Println(err)
	}
	stmt, err2 := db.Query(`
		SELECT Name, Amount FROM trades
	`)
	if err2 != nil {
		fmt.Println(err2)
	}
	var name string
	var amount int
	fmt.Println("Name", "Amount")
	for stmt.Next() {
		stmt.Scan(&name, &amount)
		fmt.Println(name, amount)
	}
	db.Close()
}

func StartServer() {
	fileServer := http.FileServer(http.Dir("./static")) // New code
	http.Handle("/", fileServer)                        // New code
	http.HandleFunc("/trades/", tradeHandler)
	http.HandleFunc("/dashboard/", dashHandler)
	http.HandleFunc("/home/", homeHandler)
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

type Trade struct {
	ID     int
	Name   string
	Amount int
	Price  float32
}
