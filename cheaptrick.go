package world

// #include <stdlib.h>
// #include "world/cheaptrick.h"
import "C"

type CheapTrickOptions struct {
	Q1      float64
	F0Floor float64
	FFTSize *int
}

func DefaultCheapTrickOptions() *CheapTrickOptions {
	return &CheapTrickOptions{
		-0.15,
		F0Floor,
		nil,
	}
}

func CheapTrick(x, f0, tpos []float64, fs int, o *CheapTrickOptions) [][]float64 {
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

	spectrogram := make([][]float64, n)
	// This contains a slice of pointers to the first element of each slice.
	ptrs := make([]*C.double, n)

	for i := range spectrogram {
		spectrogram[i] = make([]float64, m)
		ptrs[i] = (*C.double)(&spectrogram[i][0])
	}

	C.CheapTrick(
		(*C.double)(&x[0]),
		xLength,
		FS,
		(*C.double)(&tpos[0]),
		(*C.double)(&f0[0]),
		f0Length,
		&co,
		(**C.double)(&ptrs[0]),
	)

	return spectrogram
}
