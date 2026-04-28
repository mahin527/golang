package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/mahin527/golang/auth"
	"github.com/mahin527/golang/user"
)

func main() {
	auth.LoginWithCredentials("mahin", "123456")

	session := auth.GetSession()
	fmt.Println("session:", session)

	user := user.User{
		Email: "example@gmail.com",
		Name:  "Mahin",
	}

	// fmt.Println(user)

	color.Green(user.Email)
	color.Red(user.Name)
}
