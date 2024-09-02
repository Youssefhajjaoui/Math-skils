package outils

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func Convert(numbers []string) []float64 {
	result := []float64{}
	for i := 0; i < len(numbers); i++ {
		num, err := strconv.ParseFloat(numbers[i], 64)
		if err != nil {
			fmt.Println("error convert")
			os.Exit(1)
		}
		result = append(result, num)
	}
	sort.Float64s(result)
	return result
}
