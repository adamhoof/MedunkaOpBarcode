package Database

import (
	"database/sql"
	"fmt"
)

type PostgresDBHandler struct {
	db     *sql.DB
	config string
}

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

func (handler *PostgresDBHandler) QueryProductData(query string, barcode string) (name string, stock string, price string, mj string, mjKoef float64) {
	row := handler.db.QueryRow(query, barcode)
	if row.Scan(&name, &stock, &price, &mj, &mjKoef) == sql.ErrNoRows {
		return "", "", "", "", 0
	} else {
		return name, stock, price, mj, mjKoef
	}
}
