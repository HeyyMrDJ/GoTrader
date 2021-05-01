package db_functions

import (
	"database/sql"
	"fmt"
)

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

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
