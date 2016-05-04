package sfml

//#include <SFML/Graphics.h>
import "C"

import (
	"runtime"
)

type RenderWindow struct {
	data *C.sfRenderWindow
}

func destroyRenderWindow(w *C.sfRenderWindow) {
	C.sfRenderWindow_destroy(w)
}

func CreateRenderWindow(mode *VideoMode, title string, style WindowStyle, settings *ContextSettings) *RenderWindow {
	cs := cContextSettings(settings)
	w := C.sfRenderWindow_createUnicode(cVideoMode(mode), cString(title), C.sfUint32(style), &cs)
	runtime.SetFinalizer(w, destroyRenderWindow)
	return &RenderWindow{w}
}

func CreateRenderWindowFromHandle(handle WindowHandle, settings *ContextSettings) *RenderWindow {
	cs := cContextSettings(settings)
	w := C.sfRenderWindow_createFromHandle(C.sfWindowHandle(handle), &cs)
	runtime.SetFinalizer(w, destroyWindow)
	return &RenderWindow{w}
}

func (w *RenderWindow) Close() {
	C.sfRenderWindow_close(w.data)
}

func (w *RenderWindow) IsOpen() bool {
	return goBool(C.sfRenderWindow_isOpen(w.data))
}

func (w *RenderWindow) GetSettings() *ContextSettings {
	cs := C.sfRenderWindow_getSettings(w.data)
	return goContextSettings(&cs)
}

func (w *RenderWindow) PollEvent() *Event {
	var ev C.sfEvent
	if C.sfRenderWindow_pollEvent(w.data, &ev) == C.sfTrue {
		return goEvent(&ev)
	}
	return nil
}

func (w *RenderWindow) WaitEvent() *Event {
	var ev C.sfEvent
	if C.sfRenderWindow_waitEvent(w.data, &ev) == C.sfTrue {
		return goEvent(&ev)
	}
	return nil
}

func (w *RenderWindow) GetPosition() Vector2i {
	v := C.sfRenderWindow_getPosition(w.data)
	return *goVector2i(&v)
}

func (w *RenderWindow) SetPosition(position Vector2i) {
	C.sfRenderWindow_setPosition(w.data, cVector2i(&position))
}

func (w *RenderWindow) GetSize() Vector2u {
	v := C.sfRenderWindow_getSize(w.data)
	return *goVector2u(&v)
}

func (w *RenderWindow) SetSize(size Vector2u) {
	C.sfRenderWindow_setSize(w.data, cVector2u(&size))
}

func (w *RenderWindow) SetTitle(title string) {
	C.sfRenderWindow_setUnicodeTitle(w.data, cString(title))
}

func (w *RenderWindow) SetIcon(width, height uint, pixels []uint8) {
	C.sfRenderWindow_setIcon(w.data, C.uint(width), C.uint(height), (*C.sfUint8)(&pixels[0]))
}

func (w *RenderWindow) SetVisible(visible bool) {
	C.sfRenderWindow_setVisible(w.data, cBool(visible))
}

func (w *RenderWindow) SetMouseCursorVisible(visible bool) {
	C.sfRenderWindow_setMouseCursorVisible(w.data, cBool(visible))
}

func (w *RenderWindow) SetVerticalSyncEnabled(enabled bool) {
	C.sfRenderWindow_setVerticalSyncEnabled(w.data, cBool(enabled))
}

func (w *RenderWindow) SetKeyRepeatEnabled(enabled bool) {
	C.sfRenderWindow_setKeyRepeatEnabled(w.data, cBool(enabled))
}

func (w *RenderWindow) SetActive(active bool) bool {
	return goBool(C.sfRenderWindow_setActive(w.data, cBool(active)))
}

func (w *RenderWindow) RequestFocus() {
	C.sfRenderWindow_requestFocus(w.data)
}

func (w *RenderWindow) HasFocus() bool {
	return goBool(C.sfRenderWindow_hasFocus(w.data))
}

func (w *RenderWindow) Display() {
	C.sfRenderWindow_display(w.data)
}

func (w *RenderWindow) SetFramerateLimit(limit uint) {
	C.sfRenderWindow_setFramerateLimit(w.data, C.uint(limit))
}

func (w *RenderWindow) SetJoystickThreshold(treshold float32) {
	C.sfRenderWindow_setJoystickThreshold(w.data, C.float(treshold))
}

func (w *RenderWindow) GetSystemHandle() WindowHandle {
	return WindowHandle(C.sfRenderWindow_getSystemHandle(w.data))
}
