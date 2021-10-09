package world

// #cgo CFLAGS: -I ${SRCDIR}/World/src
// #cgo LDFLAGS: -L ${SRCDIR}/World/build -lworld -lm -lstdc++
import "C"

const (
	FramePeriod = 5.0
	F0Floor     = 71.0
	F0Ceil      = 800.0
)
