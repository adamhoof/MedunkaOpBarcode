package Database

import (
	"database/sql"
	"fmt"
)

type PostgresDBHandler struct {
	db     *sql.DB
	config string
}

const DropExistingTableSQL = `DROP TABLE IF EXISTS products;`
const CreateTableSQL = `CREATE TABLE products(barcode text, name text, stock text, price text, mj text, mjkoef decimal);`
const ImportFromCSVToTableSQL = `COPY products FROM '/' DELIMITER ';' CSV HEADER;`
const QueryProductDataSQL = `SELECT name, stock, price, mj, mjkoef FROM products WHERE barcode = $1;`

func (handler *PostgresDBHandler) GrabConfig(config *DBConfig) {
	handler.config = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName)
}

func (handler *PostgresDBHandler) Connect() {
	var err error
	handler.db, err = sql.Open("postgres", handler.config)
	if err != nil {
		panic(err)
	}
}

func (handler *PostgresDBHandler) ExecuteStatement(statement string) {
	_, err := handler.db.Exec(statement)
	if err != nil {
		return
	}
}

func (handler *PostgresDBHandler) QueryProductData(barcode string) (name string, stock string, price string, mj string, mjkoef float64) {
	row := handler.db.QueryRow(QueryProductDataSQL, barcode)
	if row.Scan(&name, &stock, &price, &mj, &mjkoef) == sql.ErrNoRows {
		return "", "", "", "", 0
	} else {
		return name, stock, price, mj, mjkoef
	}
}
