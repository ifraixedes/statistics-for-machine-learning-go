package c01

import (
	"fmt"
	"sort"

	"gonum.org/v1/gonum/stat"
)

func Example_mean_median_mode() {
	var data = []float64{4, 5, 1, 2, 7, 2, 6, 9, 3}

	// Calculate mean
	var mean = stat.Mean(data, nil)
	fmt.Printf("Mean: %.6f\n", mean)

	// Calculate median
	var datasasc = make([]float64, len(data))
	copy(datasasc, data)
	sort.Float64s(datasasc)

	var median float64
	if len(datasasc)%2 == 0 {
		median = (datasasc[len(datasasc)/2-1] + datasasc[len(datasasc)/2]) / 2
	} else {
		median = datasasc[len(datasasc)/2]
	}
	fmt.Printf("Median: %.0f\n", median)

	// Calculate mode
	var mode, _ = stat.Mode(data, nil)
	fmt.Printf("Mode: %.0f\n", mode)

	// Output:
	// Mean: 4.333333
	// Median: 4
	// Mode: 2
}
