package main

import (
	"context"
	"database/sql"
	"encoding/json"
	_ "github.com/mattn/go-sqlite3"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const DB_NAME = "sqlite-goexpert.db"

type Price struct {
	Body struct {
		VarBid string `json:"varBid"`
	} `json:"USDBRL"`
}

type SerializeResponse struct {
	Bid string `json:"bid"`
}

func main() {
	BootSQlite()
	mux := http.NewServeMux()
	mux.HandleFunc("/cotacao", CurrentUSDToBRLHandle)
	http.ListenAndServe(":8080", mux)

}

func CurrentUSDToBRLHandle(w http.ResponseWriter, r *http.Request) {

	jsonVia, err := fireApi()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonResp := SerializeResponse{Bid: jsonVia.Body.VarBid}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(jsonResp)
}

func fireApi() (*Price, error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer response.Body.Close()
	content, err := io.ReadAll(response.Body)

	var data Price
	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Fatal(err.Error())
	}
	sqliteDatabase := OpenConnBD()
	defer sqliteDatabase.Close()

	log.Println(data.Body.VarBid)
	floatValue, _ := strconv.ParseFloat(data.Body.VarBid, 64)
	insertPrice(sqliteDatabase, floatValue)
	return &data, nil
}

func BootSQlite() {
	os.Remove(DB_NAME)
	log.Println("Creating sqlite database")
	file, err := os.Create(DB_NAME)
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println(DB_NAME + " created")
	sqliteDatabase := OpenConnBD()
	defer sqliteDatabase.Close()
	createTable(sqliteDatabase)
}

// https://www.codeproject.com/Articles/5261771/Golang-SQLite-Simple-Example
func OpenConnBD() *sql.DB {
	sqliteDatabase, err := sql.Open("sqlite3", "./"+DB_NAME)
	if err != nil {
		log.Fatal(err.Error())
	}
	return sqliteDatabase
}

func createTable(db *sql.DB) {
	createStudentTableSQL := `CREATE TABLE prices (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"price" decimal(10,2),
		"created_at" date default current_timestamp
	  );`

	log.Println("Create prices table...")
	statement, err := db.Prepare(createStudentTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("price table created")
}

func insertPrice(db *sql.DB, price float64) {
	log.Println("Inserting new prices record ...")
	insertSQL := `INSERT INTO prices (price) VALUES (?)`
	statement, err := db.Prepare(insertSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(price)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
