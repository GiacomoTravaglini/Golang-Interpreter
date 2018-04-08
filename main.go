package main

import (
	"fmt"
	"github.com/GiacomoTravaglini/Golang-Interpreter/repl"
	"os"
)

func main() {
	fmt.Printf("Interpreter: MAIN")
	repl.Start(os.Stdin, os.Stdout)
}
