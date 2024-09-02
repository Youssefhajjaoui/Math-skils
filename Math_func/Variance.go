package Mathfunc

import "math"

func Variance(numbers []float64) int {
	var res float64
	var med float64
	lenght := float64(len(numbers))
	for i := 0; i < len(numbers); i++ {
		med += numbers[i]
	}
	med = med / lenght
	for i := 0; i < len(numbers); i++ {
		res += math.Pow((numbers[i] - med), 2)
	}
	return int(math.Round(res / (lenght)))
}
