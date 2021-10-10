package world

import "C"

type Matrix struct {
	n, m int
	S    [][]float64
	p    []*C.double
}

func NewMatrix(n, m int) *Matrix {
	s := make([][]float64, n)
	for i := range s {
		s[i] = make([]float64, m)
	}

	return &Matrix{n, m, s, nil}
}

func (mat *Matrix) Clone() *Matrix {
	newMat := NewMatrix(mat.n, mat.m)

	for i := range mat.S {
		copy(newMat.S[i], mat.S[i])
	}

	return newMat
}

func (mat *Matrix) Pointer() **C.double {
	if mat.p == nil {
		mat.p = make([]*C.double, mat.n)

		for i := range mat.S {
			mat.p[i] = (*C.double)(&mat.S[i][0])
		}
	}

	return (**C.double)(&mat.p[0])
}

func (mat *Matrix) Shape() [2]int {
	return [2]int{mat.n, mat.m}
}
