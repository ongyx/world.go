package world

// #cgo CFLAGS: -I ${SRCDIR}/World/src
// #cgo linux LDFLAGS: -L ${SRCDIR}/lib -lworld_linux -lm -lstdc++
// #cgo windows LDFLAGS: -L ${SRCDIR}/lib -lworld_windows -lm -lstdc++
import "C"

var DefaultOptions = &Options{
	AllowedRange:     0.1,
	ChannelsInOctave: 2.0,
	F0Floor:          71.0,
	F0Ceil:           800.0,
	FFTSize:          nil,
	FramePeriod:      5.0,
	Q1:               -0.15,
	Speed:            1,
	Threshold:        0.85,
}

type Options struct {
	AllowedRange     float64
	ChannelsInOctave float64
	F0Floor          float64
	F0Ceil           float64
	FFTSize          *int
	FramePeriod      float64
	Q1               float64
	Speed            int
	Threshold        float64
}
