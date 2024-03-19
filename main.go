package main

import (
	"BudLang/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Println("Hey %s!! You are in the budlang console!", user.Username)
	fmt.Println("Type in commands")

	repl.Start(os.Stdin, os.Stdout)

}
