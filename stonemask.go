package world

// #include "world/stonemask.h"
import "C"

func StoneMask(x, f0, tpos []float64, fs int) (f0Refined []float64) {
	FS := C.int(fs)
	xLength := C.int(len(x))
	f0Length := C.int(len(f0))

	f0Refined = make([]float64, f0Length)

	C.StoneMask(
		(*C.double)(&x[0]),
		xLength,
		FS,
		(*C.double)(&tpos[0]),
		(*C.double)(&f0[0]),
		f0Length,
		(*C.double)(&f0Refined[0]),
	)

	return f0Refined
}
