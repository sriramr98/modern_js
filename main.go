package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/sriramr98/modern_js/repl"
)

func main() {
    user, err := user.Current()
    if err != nil {
        panic(err)
    }
    println("Hello " + user.Username)
    fmt.Printf("Write code here to execute\n")
    repl.Start(os.Stdin, os.Stdout)
}
