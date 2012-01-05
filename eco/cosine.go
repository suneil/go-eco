// Cosine distance and similarity
// Algorithm taken from: Carbonell, J.G.& al. 1997 Translingual Information
// Retrieval: A comparative evaluation. IJCAI'97. See also Salton, G. 1989
// Automatic text processing: The transformation, Analysis, and retrieval of
// information by computer. Addison-Wesley, Reading, Pennsylvania.
// Jongman, et. al., 1995, page 178)--"More emphasis is given to qualitative
// aspects by not considering a site as point but as a vector.Understandably,
// the direction of this vector tells us something about the relative
// abundances of species. The similarity of two sites can be expressed as some 
// function of the angle between the vector of these sites. Quite common is
// the use of the cosine (or Ochiai coefficient):
// cos=OS=sigma(k)Y(ki)Y(kj)/sqrt{[sigma(k)(Y(ki)^2)][sigma(k)(Y(kj))^2)]}"
// <-- this is obviously for disance between data->cols, not data->rows (++pac). 

package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
	. "math"
)

// Cosine distance matrix
func Cosine_D(data *DenseMatrix) *DenseMatrix {
	var (
		dis *DenseMatrix
	)
	rows := data.Rows()
	cols := data.Cols()
	dis = Zeros(rows, rows)

	for i := 0; i < rows; i++ {
		dis.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			sum1 := 0.0
			sum2 := 0.0
			sum3 := 0.0
			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				sum1 += x * y
				sum2 += x * x
				sum3 += y * y
			}
			d := sum1 / (Sqrt(sum2) * Sqrt(sum3))
			dis.Set(i, j, d)
			dis.Set(j, i, d)
		}
	}
	return dis
}

// Cosine similarity matrix
// If d denotes Cosine distance, similarity is s=1.00/(d+1), so that it is in [0, 1]
func Cosine_S(data *DenseMatrix) *DenseMatrix {
	var (
		sim, dis *DenseMatrix
	)

	dis = Cosine_D(data)
	rows := data.Rows()
	sim = Zeros(rows, rows)

	for i := 0; i < rows; i++ {
		sim.Set(i, i, 1.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			s := 1.00 / (dis.Get(i, j) + 1.0)
			sim.Set(i, j, s)
			sim.Set(j, i, s)
		}
	}
	return sim
}
