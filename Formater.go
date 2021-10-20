package main

import (
	"fmt"
	"gopkg.in/gookit/color.v1"
	"strconv"
)

type Formatter struct {
}

func (formatter *Formatter) PrintColoredText(style color.Style, text string) {
	style.Println(text)
}

func (formatter *Formatter) ReturnAsString(value interface{}) string {

	switch value.(type) {
	case int32:
		return strconv.FormatInt(int64(value.(int32)), 10)
	case float32:
		return fmt.Sprintf("%.2f", value.(float32))
	default:
		return ""
	}
}
