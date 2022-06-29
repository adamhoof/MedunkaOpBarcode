package Formatting

import (
	"fmt"
	"strconv"
)

func StringToInt(s string) (i int) {
	var err error
	i, err = strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	return i
}

func ByteArrayToString(arrayOfBytes []byte) string {
	parseInt, err := strconv.ParseInt(string(arrayOfBytes), 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	return strconv.FormatInt(parseInt, 10)
}

func FloatToString(value interface{}) string {
	return fmt.Sprintf("%.2f", value.(float64))
}

func StringToFloat(value interface{}) float64 {
	floatValue, err := strconv.ParseFloat(value.(string), 8)
	if err != nil {
		return 0
	}
	return floatValue
}
