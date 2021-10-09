package world

// #include "world/cheaptrick.h"
// #include "world/d4c.h"
import "C"

type D4COptions struct {
	Threshold float64
	FFTSize   *int
}

func DefaultD4COptions() *D4COptions {
	return &D4COptions{
		0.85,
		nil,
	}
}

func D4C(x, f0, tpos []float64, fs int, o *D4COptions) [][]float64 {
	FS := C.int(fs)
	xLength := C.int(len(x))
	f0Length := C.int(len(f0))

	var fftSize C.int
	if o.FFTSize == nil {
		fftSize = C.GetFFTSizeForCheapTrick(FS, &C.CheapTrickOption{f0_floor: F0Floor})
	} else {
		fftSize = C.int(*o.FFTSize)
	}

	var co C.D4COption
	C.InitializeD4COption(&co)
	co.threshold = C.double(o.Threshold)

	n := int(f0Length)
	m := (int(fftSize) / 2) + 1

	ap := make([][]float64, n)
	ptrs := make([]*C.double, n)

	for i := range ap {
		ap[i] = make([]float64, m)
		ptrs[i] = (*C.double)(&ap[i][0])
	}

	C.D4C(
		(*C.double)(&x[0]),
		xLength,
		FS,
		(*C.double)(&tpos[0]),
		(*C.double)(&f0[0]),
		f0Length,
		fftSize,
		&co,
		(**C.double)(&ptrs[0]),
	)

	return ap
}
