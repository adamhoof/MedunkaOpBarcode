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

	switch value.(type) {
	case int32:
		return strconv.FormatInt(int64(value.(int32)), 10)
	case float64:
		return fmt.Sprintf("%.2f", value.(float64))
	case int64:
		return strconv.FormatInt(value.(int64), 10)
	default:
		return ""
	}
}

func StringToFloat(value interface{}) float64 {
	switch value.(type) {
	case string:
		floatValue, err := strconv.ParseFloat(value.(string), 8)
		if err != nil {
			return 0
		}
		return floatValue
	default:
		return 0
	}
}
