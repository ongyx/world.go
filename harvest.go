package world

// #include "world/harvest.h"
import "C"

type HarvestOptions struct {
	F0Floor     float64
	F0Ceil      float64
	FramePeriod float64
}

func DefaultHarvestOptions() *HarvestOptions {
	return &HarvestOptions{
		F0Floor,
		F0Ceil,
		FramePeriod,
	}
}

func Harvest(x []float64, fs int, o *HarvestOptions) (f0, tpos []float64) {
	FS := C.int(fs)
	xLength := C.int(len(x))

	var co C.HarvestOption
	C.InitializeHarvestOption(&co)

	co.f0_floor = C.double(o.F0Floor)
	co.f0_ceil = C.double(o.F0Ceil)
	co.frame_period = C.double(o.FramePeriod)

	f0Length := int(C.GetSamplesForHarvest(FS, xLength, co.frame_period))
	f0 = make([]float64, f0Length)
	tpos = make([]float64, f0Length)

	C.Harvest(
		(*C.double)(&x[0]),
		xLength,
		FS,
		&co,
		(*C.double)(&tpos[0]),
		(*C.double)(&f0[0]),
	)

	return f0, tpos
}
