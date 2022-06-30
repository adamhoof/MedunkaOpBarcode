package Formatting

import (
	"fmt"
	"strconv"
)

func StringToInt(str string) int {
	integer, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
	}
	return integer
}

func ByteArrayToString(arrayOfBytes []byte) string {
	str, err := strconv.ParseInt(string(arrayOfBytes), 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	return strconv.FormatInt(str, 10)
}

func FloatToString(f float64) string {
	return fmt.Sprintf("%.2f", f)
}

func StringToFloat(str string) float64 {
	floatValue, err := strconv.ParseFloat(str, 8)
	if err != nil {
		return 0
	}
	return floatValue
}
