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
