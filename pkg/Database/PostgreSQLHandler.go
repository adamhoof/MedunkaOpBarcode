package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type PostgresDBHandler struct {
	db *sql.DB
}

func (handler *PostgresDBHandler) Connect(config *string) (err error) {
	handler.db, err = sql.Open("postgres", *config)
	if err != nil {
		return fmt.Errorf("could not open connection %s", err)
	}
	return handler.db.Ping()
}

func (handler *PostgresDBHandler) ExecuteStatement(statement string) (err error) {
	_, err = handler.db.Exec(statement)
	if err != nil {
		return fmt.Errorf("failed to execute db statement %s", err)
	}
	return err
}

func (handler *PostgresDBHandler) QueryProductData(query string, barcode string) (name string, stock string, price string, mj string, mjKoef float64) {
	row := handler.db.QueryRow(query, barcode)
	if row.Scan(&name, &stock, &price, &mj, &mjKoef) == sql.ErrNoRows {
		return "", "", "", "", 0
	} else {
		return name, stock, price, mj, mjKoef
	}
}
