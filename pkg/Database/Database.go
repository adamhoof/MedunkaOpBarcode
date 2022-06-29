package Database

type Database interface {
	GrabConfig(config *DBConfig)
	Connect()
	ExecuteStatement(statement string)
}
