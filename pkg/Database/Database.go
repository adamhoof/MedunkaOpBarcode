package database

type Database interface {
	GrabConfig(config *DBConfig)
	Connect()
	ExecuteStatement(statement string)
	QueryProductData(query string, barcode string) (name string, stock string, price string, mj string, mjKoef float64)
}
