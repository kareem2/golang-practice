package main

import(
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args {
		fmt.Print(i)
		fmt.Print(" ")
		fmt.Println(arg)
	}
}
