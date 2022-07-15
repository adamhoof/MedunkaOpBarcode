package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type PostgresDBHandler struct {
	db *sql.DB
}

const DropExistingTableSQL = `DROP TABLE IF EXISTS products;`
const CreateTableSQL = `CREATE TABLE products(barcode text, name text, stock text, price text, unitOfMeasure text, unitOfMeasureKoef decimal);`
const ImportFromCSVToTableSQL = `COPY products FROM '/tmp/Products/update.csv' DELIMITER ';' CSV HEADER;`
const QueryProductDataSQL = `SELECT name, stock, price, unitOfMeasure, unitOfMeasureKoef FROM products WHERE barcode = $1;`

func (handler *PostgresDBHandler) Connect(config *string) {
	var err error
	handler.db, err = sql.Open("postgres", *config)
	if err != nil {
		fmt.Println(err)
	}
}

func (handler *PostgresDBHandler) ExecuteStatement(statement string) {
	_, err := handler.db.Exec(statement)
	if err != nil {
		fmt.Println(err)
	}
}

func (handler *PostgresDBHandler) QueryProductData(query string, barcode string) (name string, stock string, price string, mj string, mjKoef float64) {
	row := handler.db.QueryRow(query, barcode)
	if row.Scan(&name, &stock, &price, &mj, &mjKoef) == sql.ErrNoRows {
		return "", "", "", "", 0
	} else {
		return name, stock, price, mj, mjKoef
	}
}
