package TextFormatting

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

func (formatter *Formatter) ToString(value interface{}) string {

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

func (formatter *Formatter) ToFloat(value interface{}) float64 {
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
