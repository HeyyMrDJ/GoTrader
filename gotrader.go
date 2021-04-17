package main

import(
	"fmt"
	"database/sql"
	_"github.com/mattn/go-sqlite3"
)

func main() {

	CreateTable()
	GetTrades()
	//NewTrade()
	//PortValue, SLPercent := GetTradeInfo()
	//TradeSize:= TradeAmount(PortValue, SLPercent)
	//fmt.Println("Your Trade size is:", TradeSize)
}

func GetTradeInfo() (float64, float64){
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
	return PortValue, SLPercent
}

func TradeAmount(PortValue, SLPercent float64) float64 {
	Risk := PortValue / 100
	TradeSize := (Risk/(SLPercent/100))
	return TradeSize
}

func NewTrade() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		fmt.Println(err)
	}
	stmt, err2 := db.Prepare(`
		INSERT INTO trades (Name, Amount) values (?,?)
	`)
	stmt.Exec("BTC", 65000)
	if err2 != nil {
		fmt.Println(err2)
	}
	db.Close()
}

func CreateTable(){
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		fmt.Println(err)
	}
	stmt, err2 := db.Prepare(`
		CREATE TABLE IF NOT EXISTS "trades" (
			"ID"	INTEGER NOT NULL,
			"Name"	TEXT NOT NULL,
			"Amount"	INTEGER NOT NULL,
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
	db, err := sql.Open("sqlite3", "./test.db")
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
	for stmt.Next(){
		stmt.Scan(&name, &amount)
		fmt.Println(name, amount)
	}
	db.Close()
}