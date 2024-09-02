package main

import (
	"fmt"
	"os"

	outils "mathfunc/Outils"
)

func main() {
	Args := os.Args[1:]
	if len(Args) != 1 {
		fmt.Println("error arguments")
	}
	outils.Process(Args[0])
}
