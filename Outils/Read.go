package outils

import (
	"bufio"
	"fmt"
	"os"
)

func Read(s string) []string {
	numbers := []string{}
	InputFile, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("error open file")
	}
	defer InputFile.Close()
	scanner := bufio.NewScanner(InputFile)
	for scanner.Scan() {
		numb := scanner.Text()
		numbers = append(numbers, numb)
	}
	return numbers
}
