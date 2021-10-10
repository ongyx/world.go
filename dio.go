package world

// #include "world/dio.h"
import "C"

func Dio(x []float64, fs int, o *Options) (f0, tpos []float64) {
	FS := C.int(fs)
	xLength := C.int(len(x))

	var co C.DioOption
	C.InitializeDioOption(&co)

	co.f0_floor = C.double(o.F0Floor)
	co.f0_ceil = C.double(o.F0Ceil)
	co.channels_in_octave = C.double(o.ChannelsInOctave)
	co.frame_period = C.double(o.FramePeriod)
	co.speed = C.int(o.Speed)
	co.allowed_range = C.double(o.AllowedRange)

	f0Length := int(C.GetSamplesForDIO(FS, xLength, co.frame_period))
	f0 = make([]float64, f0Length)
	tpos = make([]float64, f0Length)

	C.Dio(
		(*C.double)(&x[0]),
		xLength,
		FS,
		&co,
		(*C.double)(&tpos[0]),
		(*C.double)(&f0[0]),
	)

	return f0, tpos
}
