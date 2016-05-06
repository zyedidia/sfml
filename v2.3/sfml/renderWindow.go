package sfml

//#include <SFML/Graphics.h>
import "C"

import (
	"runtime"
)

type RenderWindow struct {
	data *C.sfRenderWindow
}

func CreateRenderWindow(mode *VideoMode, title string, style WindowStyle, settings *ContextSettings) *RenderWindow {
	cs := new(C.sfContextSettings)
	if settings == nil {
		cs = nil
	} else {
		*cs = cContextSettings(settings)
	}
	w := C.sfRenderWindow_createUnicode(cVideoMode(mode), cString(title), C.sfUint32(style), cs)
	runtime.SetFinalizer(w, C.sfRenderWindow_destroy)
	return &RenderWindow{w}
}

func CreateRenderWindowFromHandle(handle WindowHandle, settings *ContextSettings) *RenderWindow {
	cs := new(C.sfContextSettings)
	if settings == nil {
		cs = nil
	} else {
		*cs = cContextSettings(settings)
	}
	w := C.sfRenderWindow_createFromHandle(C.sfWindowHandle(handle), cs)
	runtime.SetFinalizer(w, C.sfRenderWindow_destroy)
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

func (w *RenderWindow) Clear(color Color) {
	C.sfRenderWindow_clear(w.data, cColor(&color))
}

func (w *RenderWindow) SetView(view *View) {
	C.sfRenderWindow_setView(w.data, view.data)
}

func (w *RenderWindow) GetView() *View {
	return &View{C.sfRenderWindow_getView(w.data)}
}

func (w *RenderWindow) GetDefaultView() *View {
	return &View{C.sfRenderWindow_getDefaultView(w.data)}
}

func (w *RenderWindow) GetViewport(view *View) Recti {
	r := C.sfRenderWindow_getViewport(w.data, view.data)
	return *goRecti(&r)
}

func (w *RenderWindow) MapPixelToCoords(point Vector2i, view *View) Vector2f {
	r := C.sfRenderWindow_mapPixelToCoords(w.data, cVector2i(&point), view.data)
	return *goVector2f(&r)
}

func (w *RenderWindow) MapCoordsToPixel(point Vector2f, view *View) Vector2i {
	r := C.sfRenderWindow_mapCoordsToPixel(w.data, cVector2f(&point), view.data)
	return *goVector2i(&r)
}

func (w *RenderWindow) Draw(object Drawable, states *RenderStates) {
	s := new(C.sfRenderStates)
	if states == nil {
		s = nil
	} else {
		*s = cRenderStates(states)
	}
	switch object.(type) {
	case *Sprite:
		C.sfRenderWindow_drawSprite(w.data, object.(*Sprite).data, s)
	case *Text:
		C.sfRenderWindow_drawText(w.data, object.(*Text).data, s)
	case *Shape:
		C.sfRenderWindow_drawShape(w.data, object.(*Shape).data, s)
	case *CircleShape:
		C.sfRenderWindow_drawCircleShape(w.data, object.(*CircleShape).data, s)
	case *ConvexShape:
		C.sfRenderWindow_drawConvexShape(w.data, object.(*ConvexShape).data, s)
	case *RectangleShape:
		C.sfRenderWindow_drawRectangleShape(w.data, object.(*RectangleShape).data, s)
	case *VertexArray:
		C.sfRenderWindow_drawVertexArray(w.data, object.(*VertexArray).data, s)
	}
}

func (w *RenderWindow) PushGLStates() {
	C.sfRenderWindow_pushGLStates(w.data)
}

func (w *RenderWindow) PopGLStates() {
	C.sfRenderWindow_popGLStates(w.data)
}

func (w *RenderWindow) ResetGLStates() {
	C.sfRenderWindow_resetGLStates(w.data)
}

func (w *RenderWindow) Capture() *Image {
	return &Image{C.sfRenderWindow_capture(w.data)}
}
