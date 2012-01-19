package eco

import (
	. "gomatrix.googlecode.com/hg/matrix"
)

func AverageSqBool_D(data *DenseMatrix) *DenseMatrix {
	var (
		b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.Rows()
	cols := data.Cols()
	dis := Zeros(rows, rows)
	for i := 0; i < rows; i++ {
		dis.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			_, b, c, _ = getABCD(data, i, j)
			delta := (b+c) / float64(cols)
			dis.Set(i, j, delta)
			dis.Set(j, i, delta)
		}
	}
	return dis
}
