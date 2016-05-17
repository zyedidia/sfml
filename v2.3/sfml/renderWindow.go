package sfml

//#include <SFML/Graphics.h>
import "C"

import (
	"runtime"
)

type RenderWindow struct {
	data *C.sfRenderWindow
}

func destroyRenderWindow(r *RenderWindow) {
	C.sfRenderWindow_destroy(r.data)
}

func NewRenderWindow(mode VideoMode, title string, style WindowStyle, settings *ContextSettings) *RenderWindow {
	cs := new(C.sfContextSettings)
	if settings == nil {
		cs = nil
	} else {
		*cs = cContextSettings(settings)
	}
	r := C.sfRenderWindow_createUnicode(cVideoMode(&mode), cString(title), C.sfUint32(style), cs)
	obj := &RenderWindow{r}
	runtime.SetFinalizer(obj, destroyRenderWindow)
	return obj
}

func NewRenderWindowFromHandle(handle WindowHandle, settings *ContextSettings) *RenderWindow {
	cs := new(C.sfContextSettings)
	if settings == nil {
		cs = nil
	} else {
		*cs = cContextSettings(settings)
	}
	r := C.sfRenderWindow_createFromHandle(C.sfWindowHandle(handle), cs)
	obj := &RenderWindow{r}
	runtime.SetFinalizer(obj, destroyRenderWindow)
	return obj
}

func (r *RenderWindow) Close() {
	C.sfRenderWindow_close(r.data)
}

func (r *RenderWindow) IsOpen() bool {
	return goBool(C.sfRenderWindow_isOpen(r.data))
}

func (r *RenderWindow) GetSettings() *ContextSettings {
	cs := C.sfRenderWindow_getSettings(r.data)
	return goContextSettings(&cs)
}

func (r *RenderWindow) PollEvent() *Event {
	var ev C.sfEvent
	if C.sfRenderWindow_pollEvent(r.data, &ev) == C.sfTrue {
		return goEvent(&ev)
	}
	return nil
}

func (r *RenderWindow) WaitEvent() *Event {
	var ev C.sfEvent
	if C.sfRenderWindow_waitEvent(r.data, &ev) == C.sfTrue {
		return goEvent(&ev)
	}
	return nil
}

func (r *RenderWindow) GetPosition() Vector2i {
	v := C.sfRenderWindow_getPosition(r.data)
	return *goVector2i(&v)
}

func (r *RenderWindow) SetPosition(position Vector2i) {
	C.sfRenderWindow_setPosition(r.data, cVector2i(&position))
}

func (r *RenderWindow) GetSize() Vector2u {
	v := C.sfRenderWindow_getSize(r.data)
	return *goVector2u(&v)
}

func (r *RenderWindow) SetSize(size Vector2u) {
	C.sfRenderWindow_setSize(r.data, cVector2u(&size))
}

func (r *RenderWindow) SetTitle(title string) {
	C.sfRenderWindow_setUnicodeTitle(r.data, cString(title))
}

func (r *RenderWindow) SetIcon(width, height uint, pixels []byte) {
	C.sfRenderWindow_setIcon(r.data, C.uint(width), C.uint(height), (*C.sfUint8)(&pixels[0]))
}

func (r *RenderWindow) SetVisible(visible bool) {
	C.sfRenderWindow_setVisible(r.data, cBool(visible))
}

func (r *RenderWindow) SetMouseCursorVisible(visible bool) {
	C.sfRenderWindow_setMouseCursorVisible(r.data, cBool(visible))
}

func (r *RenderWindow) SetVerticalSyncEnabled(enabled bool) {
	C.sfRenderWindow_setVerticalSyncEnabled(r.data, cBool(enabled))
}

func (r *RenderWindow) SetKeyRepeatEnabled(enabled bool) {
	C.sfRenderWindow_setKeyRepeatEnabled(r.data, cBool(enabled))
}

func (r *RenderWindow) SetActive(active bool) bool {
	return goBool(C.sfRenderWindow_setActive(r.data, cBool(active)))
}

func (r *RenderWindow) RequestFocus() {
	C.sfRenderWindow_requestFocus(r.data)
}

func (r *RenderWindow) HasFocus() bool {
	return goBool(C.sfRenderWindow_hasFocus(r.data))
}

func (r *RenderWindow) Display() {
	C.sfRenderWindow_display(r.data)
}

func (r *RenderWindow) SetFramerateLimit(limit uint) {
	C.sfRenderWindow_setFramerateLimit(r.data, C.uint(limit))
}

func (r *RenderWindow) SetJoystickThreshold(treshold float32) {
	C.sfRenderWindow_setJoystickThreshold(r.data, C.float(treshold))
}

func (r *RenderWindow) GetSystemHandle() WindowHandle {
	return WindowHandle(C.sfRenderWindow_getSystemHandle(r.data))
}

func (r *RenderWindow) Clear(color Color) {
	C.sfRenderWindow_clear(r.data, cColor(&color))
}

func (r *RenderWindow) SetView(view *View) {
	C.sfRenderWindow_setView(r.data, view.data)
}

func (r *RenderWindow) GetView() *View {
	return &View{C.sfRenderWindow_getView(r.data)}
}

func (r *RenderWindow) GetDefaultView() *View {
	return &View{C.sfRenderWindow_getDefaultView(r.data)}
}

func (r *RenderWindow) GetViewport(view *View) Recti {
	v := C.sfRenderWindow_getViewport(r.data, view.data)
	return *goRecti(&v)
}

func (r *RenderWindow) MapPixelToCoords(point Vector2i, view *View) Vector2f {
	p := C.sfRenderWindow_mapPixelToCoords(r.data, cVector2i(&point), view.data)
	return *goVector2f(&p)
}

func (r *RenderWindow) MapCoordsToPixel(point Vector2f, view *View) Vector2i {
	p := C.sfRenderWindow_mapCoordsToPixel(r.data, cVector2f(&point), view.data)
	return *goVector2i(&p)
}

func (r *RenderWindow) Draw(object Drawable) {
	switch object.(type) {
	case *Sprite:
		C.sfRenderWindow_drawSprite(r.data, object.(*Sprite).data, nil)
	case *Text:
		C.sfRenderWindow_drawText(r.data, object.(*Text).data, nil)
	/*case *Shape:
	C.sfRenderWindow_drawShape(r.data, object.(*Shape).data, nil)*/
	case *CircleShape:
		C.sfRenderWindow_drawCircleShape(r.data, object.(*CircleShape).data, nil)
	case *ConvexShape:
		C.sfRenderWindow_drawConvexShape(r.data, object.(*ConvexShape).data, nil)
	case *RectangleShape:
		C.sfRenderWindow_drawRectangleShape(r.data, object.(*RectangleShape).data, nil)
	}
}

func (r *RenderWindow) DrawWithRenderStates(object Drawable, states *RenderStates) {
	s := cRenderStates(states)
	switch object.(type) {
	case *Sprite:
		C.sfRenderWindow_drawSprite(r.data, object.(*Sprite).data, &s)
	case *Text:
		C.sfRenderWindow_drawText(r.data, object.(*Text).data, &s)
	/*case *Shape:
	C.sfRenderWindow_drawShape(r.data, object.(*Shape).data, &s)*/
	case *CircleShape:
		C.sfRenderWindow_drawCircleShape(r.data, object.(*CircleShape).data, &s)
	case *ConvexShape:
		C.sfRenderWindow_drawConvexShape(r.data, object.(*ConvexShape).data, &s)
	case *RectangleShape:
		C.sfRenderWindow_drawRectangleShape(r.data, object.(*RectangleShape).data, &s)
	}

}

func (r *RenderWindow) DrawVertexArray(object *VertexArray) {
	C.sfRenderWindow_drawVertexArray(r.data, object.data, nil)
}

func (r *RenderWindow) DrawVertexArrayWithRenderStates(object *VertexArray, states *RenderStates) {
	s := cRenderStates(states)
	C.sfRenderWindow_drawVertexArray(r.data, object.data, &s)
}

func (r *RenderWindow) PushGLStates() {
	C.sfRenderWindow_pushGLStates(r.data)
}

func (r *RenderWindow) PopGLStates() {
	C.sfRenderWindow_popGLStates(r.data)
}

func (r *RenderWindow) ResetGLStates() {
	C.sfRenderWindow_resetGLStates(r.data)
}

func (r *RenderWindow) Capture() *Image {
	return &Image{C.sfRenderWindow_capture(r.data)}
}
