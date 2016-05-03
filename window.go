package sfml

//#include <SFML/Window.h>
import "C"

import (
	"runtime"
)

// Window styles
type WindowStyle uint

const (
	StyleNone        WindowStyle = C.sfNone
	StyleTitlebar    WindowStyle = C.sfTitlebar
	StyleResize      WindowStyle = C.sfResize
	StyleClose       WindowStyle = C.sfClose
	StyleFullscreen  WindowStyle = C.sfFullscreen
	StyleDefaultStye WindowStyle = C.sfDefaultStyle
)

// Context attributes
const (
	ContextDefault uint = C.sfContextDefault
	ContextCore    uint = C.sfContextCore
	ContextDebug   uint = C.sfContextDebug
)

// Context settings
type ContextSettings struct {
	DepthBits         uint
	StencilBits       uint
	AntialiasingLevel uint
	MajorVersion      uint
	MinorVersion      uint
	AttributeFlags    uint
}

type Window struct {
	data *C.sfWindow
}

func destroyWindow(w *C.sfWindow) {
	C.sfWindow_destroy(w)
}

func CreateWindow(mode *VideoMode, title string, style WindowStyle, settings *ContextSettings) *Window {
	var cs C.sfContextSettings
	cs.depthBits = C.uint(settings.DepthBits)
	cs.stencilBits = C.uint(settings.StencilBits)
	cs.antialiasingLevel = C.uint(settings.AntialiasingLevel)
	cs.majorVersion = C.uint(settings.MajorVersion)
	cs.minorVersion = C.uint(settings.MinorVersion)
	cs.attributeFlags = C.sfUint32(settings.AttributeFlags)
	w := C.sfWindow_createUnicode(cVideoMode(mode), utf32(title), C.sfUint32(style), &cs)
	runtime.SetFinalizer(w, destroyWindow)
	return &Window{w}
}

func (w *Window) Close() {
	C.sfWindow_close(w.data)
}

func (w *Window) IsOpen() bool {
	return goBool(C.sfWindow_isOpen(w.data))
}

func (w *Window) GetSettings() *ContextSettings {
	cs := new(ContextSettings)
	ccs := C.sfWindow_getSettings(w.data)
	cs.DepthBits = uint(ccs.depthBits)
	cs.StencilBits = uint(ccs.stencilBits)
	cs.AntialiasingLevel = uint(ccs.antialiasingLevel)
	cs.MajorVersion = uint(ccs.majorVersion)
	cs.MinorVersion = uint(ccs.minorVersion)
	cs.AttributeFlags = uint(ccs.attributeFlags)
	return cs
}

func (w *Window) PollEvent() *Event {

}
