package gobreezylidar

// #cgo LDFLAGS: -L. -lstdc++ -lbreezylidar
// #cgo CXXFLAGS: -std=c++14 -I. -I../cpp
// #include "gobreezylidar.h"
import "C"

type GoBreezylidar struct {
	gobreezylidar C.Breezylidar
}

func New() GoBreezylidar {
	var ret GoBreezylidar
	return ret
}

func (n GoBreezylidar) getScan() {
	C.breezy()
}
