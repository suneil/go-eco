// Fager similarity matrix
// Fager (1957), Shi (1993)

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
	"math"
)

// Fager similarity matrix
func FagerBool_S(data *DenseMatrix) *DenseMatrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	sim := Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = getABCD(data, i, j)
			s:= (a / math.Sqrt(math.Min(a+b, a+c)*math.Max(a+b, a+c))) - (1/(2*math.Sqrt(math.Min(a+b, a+c))))
			sim.Set(i, j, s)
			sim.Set(j, i, s)
		}
	}
	return sim
}
