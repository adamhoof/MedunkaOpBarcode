package EssentialConfig

type Config struct {
	PathToCSVUpdateFile                string
	CsvUpdateFileName                  string
	CsvDelimiter                       string
	DatabaseTableName                  string
	SerialPortName                     string
	SerialPortBaudRate                 int
	BarcodeReadingTerminationDelimiter byte
	DbHost                             string
	DbName                             string
	DbPort                             int
	DbUser                             string
	DbUserPassword                     string
	TelegramBotToken                   string
	TelegramBotOwner                   string
}
