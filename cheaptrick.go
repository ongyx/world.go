package world

// #include <stdlib.h>
// #include "world/cheaptrick.h"
import "C"

func CheapTrick(x, f0, tpos []float64, fs int, o *Options) (sp *Matrix) {
	FS := C.int(fs)
	xLength := C.int(len(x))
	f0Length := C.int(len(f0))

	var co C.CheapTrickOption
	C.InitializeCheapTrickOption(FS, &co)

	co.q1 = C.double(o.Q1)

	if o.FFTSize == nil {
		co.f0_floor = C.double(o.F0Floor)
		co.fft_size = C.GetFFTSizeForCheapTrick(FS, &co)
	} else {
		co.fft_size = C.int(*o.FFTSize)
	}

	n := int(f0Length)
	m := (int(co.fft_size) / 2) + 1

	sp = NewMatrix(n, m)

	C.CheapTrick(
		(*C.double)(&x[0]),
		xLength,
		FS,
		(*C.double)(&tpos[0]),
		(*C.double)(&f0[0]),
		f0Length,
		&co,
		sp.Pointer(),
	)

	return sp
}
