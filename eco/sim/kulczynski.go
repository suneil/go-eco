// Copyright 2012 - 2013 The Eco Authors. All rights reserved. See the LICENSE file.

package sim

// Kulczynski similarity and distance matrix
// Kulczynski (1928)
// Oosting (1956), Southwood (1978)

import (
	"code.google.com/p/go-eco/eco/aux"
	. "math"
)

// Kulczynski1Bool_S returns a Kulczynski similarity matrix #1 for boolean data. 
// Legendre & Legendre (1998): 257, eq. 7.15 (S12 index). 
func Kulczynski1Bool_S(data *aux.Matrix) *aux.Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = aux.GetABCD(data, i, j)
			v := a / (b + c)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// Kulczynski2Bool_S returns a Kulczynski similarity matrix #2 for boolean data. 
// From Boyce and Ellison 2001: 713, Tab. 3.
func Kulczynski2Bool_S(data *aux.Matrix) *aux.Matrix {
	var (
		a, b, c float64 // these are actually counts, but float64 simplifies the formulas
	)

	rows := data.R
	out := aux.NewMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := i; j < rows; j++ {
			a, b, c, _ = aux.GetABCD(data, i, j)
			//			v := ((a / 2) * ((2 * a) + b + c)) / ((a + b) * (a + c))
			v := 0.5 * (a/(a+b) + a/(a+c))
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// Kulczynski_D returns a Kulczynski distance matrix for floating-point data. 
func Kulczynski_D(data *aux.Matrix) *aux.Matrix {
	rows := data.R
	cols := data.C
	out := aux.NewMatrix(rows, rows)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			sumMin := 0.0
			sumX := 0.0
			sumY := 0.0
			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				sumMin += Min(x, y)
				sumX += x
				sumY += x
			}
			v := 1 - 0.5*(sumMin/sumX+sumMin/sumY)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}

// Kulczynski_S returns a Kulczynski similarity matrix for  for count or interval data. 
// Legendre & Legendre (1998): 265, eq. 7.25 (S18 index). 
func Kulczynski_S(data *aux.Matrix) *aux.Matrix {
	rows := data.R
	cols := data.C
	out := aux.NewMatrix(rows, rows)

	for i := 0; i < rows; i++ {
		out.Set(i, i, 0.0)
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			sumMin := 0.0
			sumX := 0.0
			sumY := 0.0
			for k := 0; k < cols; k++ {
				x := data.Get(i, k)
				y := data.Get(j, k)
				sumMin += Min(x, y)
				sumX += x
				sumY += x
			}
			v := 0.5 * (sumMin/sumX + sumMin/sumY)
			out.Set(i, j, v)
			out.Set(j, i, v)
		}
	}
	return out
}
