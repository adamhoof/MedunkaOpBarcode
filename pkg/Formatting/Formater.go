package Formatting

import (
	"fmt"
	"gopkg.in/gookit/color.v1"
	"strconv"
)

var ActualPriceStyle = color.Style{color.FgRed, color.OpBold}
var DefaultStyle = color.Style{color.FgLightWhite, color.OpItalic}

func PrintStyledText(style color.Style, text string) {
	style.Println(text)
}

func ToString(value interface{}) string {

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

func ToFloat(value interface{}) float64 {
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
