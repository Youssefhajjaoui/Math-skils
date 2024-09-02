package Mathfunc

import "math"

func Average(numbers []float64) int64 {
	var somme float64
	for i := 0; i < len(numbers); i++ {
		somme += numbers[i]
	}
	lenght := float64(len(numbers))
	return int64(math.Round(somme / lenght))
}
