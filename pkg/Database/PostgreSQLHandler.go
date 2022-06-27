package Database

import (
	"database/sql"
)

const DropExistingTableSQL = `DROP TABLE IF EXISTS products;`
const CreateTableSQL = `CREATE TABLE products(barcode text, name text, stock text, price text, mj text, mjkoef decimal);`
const ImportFromCSVToTableSQL = `COPY products FROM '/' DELIMITER ';' CSV HEADER;`
const QueryProductDataSQL = `SELECT name, stock, price, mj, mjkoef FROM products WHERE barcode = $1;`

func Connect(db *sql.DB, config string) {
	db, err := sql.Open("postgres", config)
	if err != nil {
		panic(err)
	}
}

func ExecuteStatement(db *sql.DB, statement string) {
	_, err := db.Exec(statement)
	if err != nil {
		return
	}
}

func QueryProductData(db *sql.DB, barcode string) (name string, stock string, price string, mj string, mjkoef float64) {
	row := db.QueryRow(QueryProductDataSQL, barcode)
	if row.Scan(&name, &stock, &price, &mj, &mjkoef) == sql.ErrNoRows {
		return "", "", "", "", 0
	} else {
		return name, stock, price, mj, mjkoef
	}
}
