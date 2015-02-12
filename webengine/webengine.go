package webengine

// #cgo LDFLAGS: -L./ -lwebengine -lstdc++
// #include "cpp/webengine.h"
import "C"

import "github.com/chai2010/qml"

// Initializes the WebEngine extension.
func Initialize() {
	qml.RunMain(func() {
		C.webengineInitialize()
	})
}
