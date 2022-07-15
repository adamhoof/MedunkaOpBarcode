package database

type Database interface {
	Connect(config *string)
	ExecuteStatement(statement string)
	QueryProductData(query string, barcode string) (name string, stock string, price string, mj string, mjKoef float64)
}
