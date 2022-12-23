package main

import "fmt"

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func main() {
	minFunc, err1 := operation(minimum)
	averageFunc, err2 := operation(average)
	maxFunc, err3 := operation(maximum)

	fmt.Print(err1, err2, err3)

	minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Printf("Max: %.2f,\nMin: %.2f, \nAverage: %.2f\n", maxValue, minValue, averageValue)

}

func operation(operacion string) (func(valores ...float64) float64, error) {
	switch operacion {
	case minimum:
		return MinFunc, nil
	case average:
		return AverageFunc, nil
	case maximum:
		return MaxFunc, nil
	}
	return nil, nil
}

func MinFunc(valores ...float64) float64 {
	min := valores[0]
	for _, value := range valores {
		if value < min {
			min = value
		}
	}
	return min
}

func AverageFunc(valores ...float64) float64 {
	var promedio float64
	for _, v := range valores {
		promedio += v
	}
	return promedio / float64(len(valores))
}

func MaxFunc(valores ...float64) float64 {
	max := valores[0]
	for _, value := range valores {
		if value > max {
			max = value
		}
	}
	return max
}
