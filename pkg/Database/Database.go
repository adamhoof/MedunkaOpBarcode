package database

type Database interface {
	Connect(config *string) (err error)
	ExecuteStatement(statement string) (err error)
	QueryProductData(query string, barcode string) (name string, stock string, price string, mj string, mjKoef float64)
}
