package webengine

// #cgo CPPFLAGS: -I./
// #cgo CXXFLAGS: -std=c++0x -pedantic-errors -Wall -fno-strict-aliasing
// #cgo LDFLAGS: -lstdc++
// #cgo !windows pkg-config: Qt5WebEngine
//
// #include "cpp/webengine.h"
import "C"

import "github.com/chai2010/qml"

// Initializes the WebEngine extension.
func Initialize() {
	qml.RunMain(func() {
		C.webengineInitialize()
	})
}
