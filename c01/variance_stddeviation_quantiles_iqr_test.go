package c01

import (
	"fmt"
	"sort"

	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/stat"
)

func Example_variance_stddeviation_quantiles_iqr() {
	var gamePoints = []float64{
		35, 56, 43, 59, 63, 79, 35, 41, 64, 43, 93, 60, 77, 24, 82,
	}

	var variance = stat.Variance(gamePoints, nil)
	fmt.Printf("Sample variance: %.2f\n", variance)

	var stddev = stat.StdDev(gamePoints, nil)
	fmt.Printf("Sample std.dev: %.2f\n", stddev)

	var rang = floats.Max(gamePoints) - floats.Min(gamePoints)
	fmt.Printf("Range: %.0f\n", rang)

	var gamePointssasc = make([]float64, len(gamePoints))
	copy(gamePointssasc, gamePoints)
	sort.Float64s(gamePointssasc)
	fmt.Printf("%d - %v\n", len(gamePointssasc), gamePointssasc)

	fmt.Println("Quantiles:")
	// Unfortunately, gonum doesn't provide the cumulant type 7 that python and
	// R uses by default so the output values here differ of the ones present in
	// the book
	// see https://github.com/gonum/gonum/issues/253
	for _, p := range [3]float64{0.2, 0.8, 1} {
		var q = stat.Quantile(p, stat.Empirical, gamePointssasc, nil)
		fmt.Printf("\t%d%% %.0f\n", int(p*100), q)
	}

	var iqr = stat.Quantile(0.75, stat.Empirical, gamePointssasc, nil) -
		stat.Quantile(0.25, stat.Empirical, gamePointssasc, nil)
	fmt.Printf("Inter quartile range: %.1f\n", iqr)

	// Output:
	// Sample variance: 400.64
	// Sample std.dev: 20.02
	// Range: 69
	// Quantiles:
	// 	20% 35
	// 	%80 77
	//	%100 93
	// Inter quartile range: 36
}
