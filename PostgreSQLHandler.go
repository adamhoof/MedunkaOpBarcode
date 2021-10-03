package main

import (
	"database/sql"
	"fmt"
)

type PostgreSQLHandler struct {
	db *sql.DB
}

const (
	host     = "192.168.1.116"
	port     = 5432
	user     = "pi"
	password = "medprodsdb"
	dbname   = "medunkaproducts"
)

const dropExistingTableSQL = `DROP TABLE IF EXISTS products;`
const createTableSQL = `CREATE TABLE products(barcode bigint, price smallint, mj varchar(5), mjkoef decimal);`
const importFromCSVToTableSQL = `COPY products FROM '/home/pi/MedunkaOpBarcode/products.csv' DELIMITER ';' CSV HEADER;`
const queryProductInfoSQL = `SELECT (barcode = $1) FROM products;`

func (postgreHandler *PostgreSQLHandler) Connect() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s",
		host, port, user, password, dbname)

	var err error
	postgreHandler.db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
}

func (postgreHandler *PostgreSQLHandler) TestConnection() {
	result := postgreHandler.db.Ping()
	if result != nil {
		panic(result)
	}
}

func (postgreHandler *PostgreSQLHandler) Disconnect() {
	err := postgreHandler.db.Close()
	if err != nil {
		panic(err)
	}
}

func (postgreHandler *PostgreSQLHandler) CreateTable(){
	_, err := postgreHandler.db.Exec(createTableSQL)
	if err != nil {
		panic(err)
	}
}

func (postgreHandler *PostgreSQLHandler) ImportFromCSV(){
	_, err := postgreHandler.db.Exec(importFromCSVToTableSQL)
	if err != nil {
		panic(err)
	}
}

func (postgreHandler *PostgreSQLHandler) DropTableIfExists(){
	_, err := postgreHandler.db.Exec(dropExistingTableSQL)
	if err != nil {
		panic(err)
	}
}

func (postgreHandler *PostgreSQLHandler) QueryProductData(barcode int64) (
	price int32, mj string, mjkoef float32) {

	var returnBarcode int64

	if err := postgreHandler.db.QueryRow(queryProductInfoSQL, barcode).Scan(returnBarcode, price, mj, mjkoef); err != nil {
		if err == sql.ErrNoRows {
			return 0, "", 0
		}
		return
	}

	return price, mj, mjkoef
}
