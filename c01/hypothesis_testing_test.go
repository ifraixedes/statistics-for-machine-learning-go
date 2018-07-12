package c01

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/stat/distuv"
)

func Example_hypothesis_testing() {
	var (
		xbar float64 = 990
		mu0  float64 = 1000
		n    float64 = 30
		s            = 12.5
	)

	var tSmpl = (xbar - mu0) / (s / math.Sqrt(n))
	fmt.Printf("Test Statistics: %.2f\n", tSmpl)

	var alpha = 0.05
	var stdentsTDist = distuv.StudentsT{
		Mu:    0,
		Sigma: 1,
		Nu:    n - 1,
	}
	var tAlpha = stdentsTDist.Quantile(alpha)
	fmt.Printf("Critical value from t-table: %.3f\n", tAlpha)

	var pVal = stdentsTDist.CDF(tSmpl)
	fmt.Printf("Lower tail p-value from t-table: %.11e", pVal)

	// Output:
	// Test Statistics: -4.38
	// Critical value from t-table: -1.699
	// Lower tail p-value from t-table: 7.03502572901e-05
}
