package world

// #include "world/synthesis.h"
import "C"

import (
	"errors"
	"fmt"
)

func Synthesize(f0 []float64, sp, ap *Matrix, fs int, o *Options) (y []float64, err error) {
	f0Length := len(f0)
	spShape := sp.Shape()
	apShape := ap.Shape()

	if (f0Length != spShape[0]) || (f0Length != apShape[0]) {
		return nil, errors.New(fmt.Sprintf("frame mismatch between f0 (%d), sp (%d) and ap (%d)", f0Length, spShape[0], apShape[0]))
	}

	if spShape[1] != apShape[1] {
		return nil, errors.New(fmt.Sprintf("shape mismatch between sp (%d) and ap (%d)", spShape[1], apShape[1]))
	}

	yLength := int(float64(f0Length) * o.FramePeriod * (float64(fs) / 1000))
	fftSize := (spShape[1] - 1) * 2

	y = make([]float64, yLength)

	C.Synthesis(
		(*C.double)(&f0[0]),
		C.int(f0Length),
		sp.Pointer(),
		ap.Pointer(),
		C.int(fftSize),
		C.double(o.FramePeriod),
		C.int(fs),
		C.int(yLength),
		(*C.double)(&y[0]),
	)

	return y, nil
}

func WavToWorld(x []float64, fs int, o *Options) (f0 []float64, sp, ap *Matrix) {
	if o == nil {
		o = DefaultOptions
	}

	rawF0, t := Dio(x, fs, o)
	f0 = StoneMask(x, rawF0, t, fs)
	sp = CheapTrick(x, f0, t, fs, o)
	ap = D4C(x, f0, t, fs, o)

	return f0, sp, ap
}
