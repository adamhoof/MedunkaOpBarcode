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

const dropExistingTable = `DROP TABLE IF EXISTS products;`
const createTable = `CREATE TABLE products(barcode bigint, price smallint, mj varchar(5), mjkoef decimal);`
const importFromCSVToTable = `COPY products FROM '/home/pi/MedunkaOpBarcode/products.csv' DELIMITER ';' CSV HEADER;`

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
	_, err := postgreHandler.db.Exec(createTable)
	if err != nil {
		panic(err)
	}
}

func (postgreHandler *PostgreSQLHandler) ImportFromCSV(){
	_, err := postgreHandler.db.Exec(importFromCSVToTable)
	if err != nil {
		panic(err)
	}
}

func (postgreHandler *PostgreSQLHandler) DropTableIfExists(){
	_, err := postgreHandler.db.Exec(dropExistingTable)
	if err != nil {
		panic(err)
	}
}
