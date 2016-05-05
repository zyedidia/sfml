package sfml

//#include <SFML/Window.h>
import "C"

import (
	"runtime"
)

const (
	StyleNone        WindowStyle = C.sfNone
	StyleTitlebar    WindowStyle = C.sfTitlebar
	StyleResize      WindowStyle = C.sfResize
	StyleClose       WindowStyle = C.sfClose
	StyleFullscreen  WindowStyle = C.sfFullscreen
	StyleDefaultStye WindowStyle = C.sfDefaultStyle
)

const (
	ContextDefault uint = C.sfContextDefault
	ContextCore    uint = C.sfContextCore
	ContextDebug   uint = C.sfContextDebug
)

type WindowStyle uint

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

type WindowHandle C.sfWindowHandle

func CreateWindow(mode *VideoMode, title string, style WindowStyle, settings *ContextSettings) *Window {
	var cs *C.sfContextSettings
	if settings == nil {
		cs = nil
	} else {
		c := cContextSettings(settings)
		cs = &c
	}
	w := C.sfWindow_createUnicode(cVideoMode(mode), cString(title), C.sfUint32(style), cs)
	runtime.SetFinalizer(w, C.sfWindow_destroy)
	return &Window{w}
}

func CreateWindowFromHandle(handle WindowHandle, settings *ContextSettings) *Window {
	var cs *C.sfContextSettings
	if settings == nil {
		cs = nil
	} else {
		c := cContextSettings(settings)
		cs = &c
	}
	w := C.sfWindow_createFromHandle(C.sfWindowHandle(handle), cs)
	runtime.SetFinalizer(w, C.sfWindow_destroy)
	return &Window{w}
}

func (w *Window) Close() {
	C.sfWindow_close(w.data)
}

func (w *Window) IsOpen() bool {
	return goBool(C.sfWindow_isOpen(w.data))
}

func (w *Window) GetSettings() *ContextSettings {
	cs := C.sfWindow_getSettings(w.data)
	return goContextSettings(&cs)
}

func (w *Window) PollEvent() *Event {
	var ev C.sfEvent
	if C.sfWindow_pollEvent(w.data, &ev) == C.sfTrue {
		return goEvent(&ev)
	}
	return nil
}

func (w *Window) WaitEvent() *Event {
	var ev C.sfEvent
	if C.sfWindow_waitEvent(w.data, &ev) == C.sfTrue {
		return goEvent(&ev)
	}
	return nil
}

func (w *Window) GetPosition() Vector2i {
	v := C.sfWindow_getPosition(w.data)
	return *goVector2i(&v)
}

func (w *Window) SetPosition(position Vector2i) {
	C.sfWindow_setPosition(w.data, cVector2i(&position))
}

func (w *Window) GetSize() Vector2u {
	v := C.sfWindow_getSize(w.data)
	return *goVector2u(&v)
}

func (w *Window) SetSize(size Vector2u) {
	C.sfWindow_setSize(w.data, cVector2u(&size))
}

func (w *Window) SetTitle(title string) {
	C.sfWindow_setUnicodeTitle(w.data, cString(title))
}

func (w *Window) SetIcon(width, height uint, pixels []uint8) {
	C.sfWindow_setIcon(w.data, C.uint(width), C.uint(height), (*C.sfUint8)(&pixels[0]))
}

func (w *Window) SetVisible(visible bool) {
	C.sfWindow_setVisible(w.data, cBool(visible))
}

func (w *Window) SetMouseCursorVisible(visible bool) {
	C.sfWindow_setMouseCursorVisible(w.data, cBool(visible))
}

func (w *Window) SetVerticalSyncEnabled(enabled bool) {
	C.sfWindow_setVerticalSyncEnabled(w.data, cBool(enabled))
}

func (w *Window) SetKeyRepeatEnabled(enabled bool) {
	C.sfWindow_setKeyRepeatEnabled(w.data, cBool(enabled))
}

func (w *Window) SetActive(active bool) bool {
	return goBool(C.sfWindow_setActive(w.data, cBool(active)))
}

func (w *Window) RequestFocus() {
	C.sfWindow_requestFocus(w.data)
}

func (w *Window) HasFocus() bool {
	return goBool(C.sfWindow_hasFocus(w.data))
}

func (w *Window) Display() {
	C.sfWindow_display(w.data)
}

func (w *Window) SetFramerateLimit(limit uint) {
	C.sfWindow_setFramerateLimit(w.data, C.uint(limit))
}

func (w *Window) SetJoystickThreshold(treshold float32) {
	C.sfWindow_setJoystickThreshold(w.data, C.float(treshold))
}

func (w *Window) GetSystemHandle() WindowHandle {
	return WindowHandle(C.sfWindow_getSystemHandle(w.data))
}

func goContextSettings(cs *C.sfContextSettings) *ContextSettings {
	return &ContextSettings{
		uint(cs.depthBits),
		uint(cs.stencilBits),
		uint(cs.antialiasingLevel),
		uint(cs.majorVersion),
		uint(cs.minorVersion),
		uint(cs.attributeFlags),
	}
}

func cContextSettings(cs *ContextSettings) C.sfContextSettings {
	return C.sfContextSettings{
		C.uint(cs.DepthBits),
		C.uint(cs.StencilBits),
		C.uint(cs.AntialiasingLevel),
		C.uint(cs.MajorVersion),
		C.uint(cs.MinorVersion),
		C.sfUint32(cs.AttributeFlags),
	}
}
