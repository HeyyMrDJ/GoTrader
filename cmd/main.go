package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	//GetTradeInfo()
	//CreateTable()
	//UpdateTrade(1, "DOGE", 1, 1)
	StartServer()

}

func updatetradeHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// logic part of log in
	//var amount float32
	//name := r.Form["name"][0]
	//name := r.Form["name"][0]
	amount, _ := strconv.ParseFloat(r.Form.Get("amount"), 64)
	price, _ := strconv.ParseFloat(r.Form.Get("price"), 64)
	id, _ := strconv.Atoi(r.Form.Get("id"))
	name := r.Form.Get("name")
	fmt.Println(r.Form.Get("id"))
	fmt.Println(r.Form.Get("name"))
	fmt.Println(r.Form.Get("amount"))
	fmt.Println(r.Form.Get("price"))
	UpdateTrade(id, name, amount, price)
	http.Redirect(w, r, "/trades", 302)

}

func deletetradeHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// logic part of log in
	//var amount float32
	//name := r.Form["name"][0]
	//name := r.Form["name"][0]
	trade, _ := strconv.Atoi(r.Form.Get("delete"))
	DeleteTrade(trade)
	http.Redirect(w, r, "/trades", 302)

}

func posttradeHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// logic part of log in
	//var amount float32
	//name := r.Form["name"][0]
	//name := r.Form["name"][0]
	amount, _ := strconv.ParseFloat(r.Form.Get("amount"), 64)
	price, _ := strconv.ParseFloat(r.Form.Get("price"), 64)
	fmt.Println(r.Form.Get("name"))
	fmt.Println(r.Form.Get("amount"))
	fmt.Println(r.Form.Get("price"))
	NewTrade(r.Form.Get("name"), price, amount)
	http.Redirect(w, r, "/trades", 302)

}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"../web/static/template/home/index.html",
		"../web/static/template/header.html",
		"../web/static/template/footer.html",
	}

	tmpl := template.Must(template.ParseFiles(files...))

	tmpl.Execute(w, nil)
}

func dashHandler(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"../web/static/template/dashboard/index.html",
		"../web/static/template/header.html",
		"../web/static/template/footer.html",
	}

	tmpl := template.Must(template.ParseFiles(files...))

	tmpl.Execute(w, nil)
}

func tradeHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "../trades.db")
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
		"../web/static/template/trades/index.html",
		"../web/static/template/header.html",
		"../web/static/template/footer.html",
		"../web/static/template/trades.html",
	}

	tmpl := template.Must(template.ParseFiles(files...))
	data := listtrades
	fmt.Println(data)

	tmpl.Execute(w, listtrades)
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

func NewTrade(name string, amount, price float64) {
	db, err := sql.Open("sqlite3", "../trades.db")
	if err != nil {
		fmt.Println(err)
	}
	stmt, err2 := db.Prepare(`
		INSERT INTO trades (Name, Amount, Price) values (?,?,?)
	`)
	stmt.Exec(name, amount, price)
	if err2 != nil {
		fmt.Println(err2)
	}
	db.Close()
}

func CreateTable() {
	db, err := sql.Open("sqlite3", "../trades.db")
	if err != nil {
		fmt.Println(err)
	}
	stmt, err2 := db.Prepare(`
		CREATE TABLE "trades" (
			"ID"	INTEGER NOT NULL,
			"Name"	TEXT NOT NULL,
			"Amount"	INTEGER NOT NULL,
			"Price"	INTEGER NOT NULL,
			PRIMARY KEY("ID" AUTOINCREMENT)d
		);
	`)
	stmt.Exec()
	if err2 != nil {
		fmt.Println(err2)
	}

	db.Close()
}

func DeleteTrade(ID int) {
	db, err := sql.Open("sqlite3", "../trades.db")
	if err != nil {
		fmt.Println(err)
	}
	//stmt, err2 := db.Prepare(`
	//	delete from trades where ID=12
	//`)
	//stmt.Exec(12)
	//
	//if err2 != nil {
	//	fmt.Println(err2)
	//}

	stmt, err := db.Prepare("DELETE FROM trades WHERE ID=?")
	checkErr(err)

	res, err := stmt.Exec(ID)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	db.Close()
}

func UpdateTrade(ID int, name string, amount, price float64) {
	db, err := sql.Open("sqlite3", "../trades.db")
	if err != nil {
		fmt.Println(err)
	}
	//stmt, err2 := db.Prepare(`
	//	delete from trades where ID=12
	//`)
	//stmt.Exec(12)
	//
	//if err2 != nil {
	//	fmt.Println(err2)
	//}

	stmt, err := db.Prepare("UPDATE trades set Name=?, Amount=?, Price=? where ID=?")
	checkErr(err)

	res, err := stmt.Exec(name, amount, price, ID)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
	db.Close()
}

func GetTrades() {
	db, err := sql.Open("sqlite3", "../trades.db")
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

type Trade struct {
	ID     int
	Name   string
	Amount float64
	Price  float64
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
