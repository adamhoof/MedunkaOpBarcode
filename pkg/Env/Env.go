package env

import (
	"fmt"
	"os"
)

func SetEnv() {
	err := os.Setenv("host", "10.0.0.2")
	if err != nil {
		fmt.Println("unable to set env for: ", err)
	}
	err = os.Setenv("dbPort", "5432")
	if err != nil {
		fmt.Println("unable to set env for: ", err)
	}
	err = os.Setenv("user", "pi")
	if err != nil {
		fmt.Println("unable to set env for: ", err)
	}
	err = os.Setenv("password", "medprodsdb")
	if err != nil {
		fmt.Println("unable to set env for: ", err)
	}
	err = os.Setenv("dbname", "medunkaproducts")
	if err != nil {
		fmt.Println("unable to set env for: ", err)
	}
	err = os.Setenv("serialPort", "/dev/ttyAMA0")
	if err != nil {
		fmt.Println("unable to set env for: ", err)
	}
	err = os.Setenv("botToken", "5586681713:AAERYtqOJ-0MfPENpOOgCYG5zh_aXh0Maig")
	if err != nil {
		fmt.Println("unable to set env for: ", err)
	}
	err = os.Setenv("botOwner", "-1001671432440")
	if err != nil {
		fmt.Println("unable to set env for: ", err)
	}

}
