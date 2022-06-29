package Formatting

import (
	"fmt"
	"gopkg.in/gookit/color.v1"
)

var BoldRed = color.Style{color.FgRed, color.OpBold}
var Default = color.Style{color.FgLightWhite, color.OpItalic}

func PrintStyledText(style color.Style, text string) {
	style.Println(text)
}

func PrintSpaces(numSpaces uint8) {
	var i uint8 = 0
	for ; i < numSpaces; i++ {
		fmt.Println("")
	}
}
