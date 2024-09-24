package main

import (
	"fmt"
	"github.com/nam9nine/interpreter/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is The Nam9 programming language!\n", user.Username)

	err = repl.Start(os.Stdin, os.Stdout)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
	}
}
