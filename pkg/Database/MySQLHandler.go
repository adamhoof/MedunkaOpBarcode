package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type MySQLHandler struct {
	db *sql.DB
}

func (handler *MySQLHandler) Connect(config *string) {
	var err error
	handler.db, err = sql.Open("mysql", *config)
	if err != nil {
		fmt.Println(err)
	}
}

func (handler *MySQLHandler) ExecuteStatement(statement string) {
	_, err := handler.db.Exec(statement)
	if err != nil {
		fmt.Println(err)
	}
}

func (handler *MySQLHandler) QueryProductData(query string, barcode string) (name string, stock string, price string, mj string, mjKoef float64) {
	row := handler.db.QueryRow(query, barcode)
	if row.Scan(&name, &stock, &price, &mj, &mjKoef) == sql.ErrNoRows {
		return "", "", "", "", 0
	} else {
		return name, stock, price, mj, mjKoef
	}
}
