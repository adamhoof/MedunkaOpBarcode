package artist

import (
	"fmt"
	"gopkg.in/gookit/color.v1"
)

func PrintStyledText(style color.Style, text string) {
	style.Println(text)
}

func PrintSpaces(numSpaces uint8) {
	var i uint8 = 0
	for ; i < numSpaces; i++ {
		fmt.Println("")
	}
}

func ClearTerminal() {
	fmt.Print("\033[H\033[2J")
}
