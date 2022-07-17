package database

import "fmt"

func GenerateDropExistingTableIfExistsSQL(tableName string) (statement string) {
	return fmt.Sprintf("DROP TABLE IF EXISTS %s;", tableName)
}

func GenerateCreateTableSQL(tableName string) (statement string) {
	return fmt.Sprintf("CREATE TABLE %s(barcode text, name text, stock text, price text, unitOfMeasure text, unitOfMeasureKoef decimal);", tableName)
}

func GenerateImportFromCSVToTableSQL(tableName string, pathToCSVUpdate string, csvUpdateFileName string, delimiter string) (statement string) {
	return fmt.Sprintf("COPY %s FROM '%s/%s' DELIMITER '%s' CSV HEADER;", tableName, pathToCSVUpdate, csvUpdateFileName, delimiter)
}

func GenerateQueryProductDataSQL(tableName string) (statement string) {
	return fmt.Sprintf("SELECT name, stock, price, unitOfMeasure, unitOfMeasureKoef FROM %s WHERE barcode = $1;", tableName)
}
