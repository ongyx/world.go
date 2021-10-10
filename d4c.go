package world

// #include "world/cheaptrick.h"
// #include "world/d4c.h"
import "C"

func D4C(x, f0, tpos []float64, fs int, o *Options) (ap *Matrix) {
	FS := C.int(fs)
	xLength := C.int(len(x))
	f0Length := C.int(len(f0))

	var fftSize C.int
	if o.FFTSize == nil {
		fftSize = C.GetFFTSizeForCheapTrick(FS, &C.CheapTrickOption{f0_floor: C.double(DefaultOptions.F0Floor)})
	} else {
		fftSize = C.int(*o.FFTSize)
	}

	var co C.D4COption
	C.InitializeD4COption(&co)
	co.threshold = C.double(o.Threshold)

	n := int(f0Length)
	m := (int(fftSize) / 2) + 1

	ap = NewMatrix(n, m)

	C.D4C(
		(*C.double)(&x[0]),
		xLength,
		FS,
		(*C.double)(&tpos[0]),
		(*C.double)(&f0[0]),
		f0Length,
		fftSize,
		&co,
		ap.Pointer(),
	)

	return ap
}
