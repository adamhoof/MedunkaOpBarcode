package Env

import (
	"fmt"
	"os"
)

func setEnv() {
	err := os.Setenv("")
	if err != nil {
		fmt.Println("unable to set env for ")
	}
}


