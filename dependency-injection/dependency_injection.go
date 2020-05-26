package main

import (
	"fmt"
	"os"
	"io"
)

func Greet(write io.Writer, name string) {
	fmt.Fprintf(write, "Hello, %s", name)
}

func main() {
	Greet(os.Stdout, "Elodie")
}