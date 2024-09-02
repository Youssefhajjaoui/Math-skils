package Mathfunc

import "math"

func Average(numbers []float64) int {
	var somme float64
	for i := 0; i < len(numbers); i++ {
		somme += numbers[i]
	}
	lenght := float64(len(numbers))
	return int(math.Round(somme / lenght))
}
