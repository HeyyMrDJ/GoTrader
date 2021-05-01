package http_handler

import (
	"GoTrader/pkg/database/db_functions"
	"GoTrader/pkg/database/db_types"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

func UpdatetradeHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	amount, _ := strconv.ParseFloat(r.Form.Get("amount"), 64)
	price, _ := strconv.ParseFloat(r.Form.Get("price"), 64)
	id, _ := strconv.Atoi(r.Form.Get("id"))
	name := r.Form.Get("name")
	fmt.Println(r.Form.Get("id"))
	fmt.Println(r.Form.Get("name"))
	fmt.Println(r.Form.Get("amount"))
	fmt.Println(r.Form.Get("price"))
	db_functions.UpdateTrade(id, name, amount, price)
	http.Redirect(w, r, "/trades", http.StatusFound)

}

func DeletetradeHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	trade, _ := strconv.Atoi(r.Form.Get("delete"))
	db_functions.DeleteTrade(trade)
	http.Redirect(w, r, "/trades", http.StatusFound)

}

func PosttradeHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	amount, _ := strconv.ParseFloat(r.Form.Get("amount"), 64)
	price, _ := strconv.ParseFloat(r.Form.Get("price"), 64)
	name := r.Form.Get("name")
	fmt.Println(r.Form.Get("name"))
	fmt.Println(r.Form.Get("amount"))
	fmt.Println(r.Form.Get("price"))
	db_functions.NewTrade(name, amount, price)
	http.Redirect(w, r, "/trades", http.StatusFound)

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"../web/static/template/home/index.html",
		"../web/static/template/header.html",
		"../web/static/template/footer.html",
	}

	tmpl := template.Must(template.ParseFiles(files...))

	tmpl.Execute(w, nil)
}

func DashHandler(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"../web/static/template/dashboard/index.html",
		"../web/static/template/header.html",
		"../web/static/template/footer.html",
	}

	tmpl := template.Must(template.ParseFiles(files...))

	tmpl.Execute(w, nil)
}

func TradeHandler(w http.ResponseWriter, r *http.Request) {
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

	var tradeQuery db_types.Trade
	var listtrades []db_types.Trade
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
