// Herfindahl index of concentration

package div

import (
	. "go-eco.googlecode.com/hg/eco"
	"math"
)

// Herfindahl index of concentration
// F A Cowell: Measurement of Inequality, 2000, in A B Atkinson & F Bourguignon (Eds): Handbook of Income Distribution. Amsterdam.
// F A Cowell: Measuring Inequality, 1995 Prentice Hall/Harvester Wheatshef.
// M Hall & N Tidemann: Measures of Concentration, 1967, JASA 62, 162-168.
func Herfindahl_D(data *Matrix, m float64) *Vector {
	rows := data.R
	cols := data.C
	out := NewVector(rows)

	for i := 0; i < rows; i++ {
		s := 0.0    // number of species
		sumX := 0.0 // total number of all individuals in the sample

		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			if x > 0.0 {
				s++
				sumX += x
			}
		}

		v := 0.0
		for j := 0; j < cols; j++ {
			x := data.Get(i, j)
			if x > 0.0 {
				y := x / sumX
				y = math.Pow(y, m+1)
				y = x * math.Log(y)
				v += y
			}
		}
		v = math.Pow(v, 1/s)
		out.Set(i, v)
	}
	return out
}