package sfml

//#include <SFML/Graphics.h>
import "C"

import (
	"runtime"
)

type View struct {
	data *C.sfView
}

func destroyView(v *View) {
	C.sfView_destroy(v.data)
}

func CreateView() *View {
	v := C.sfView_create()
	obj := &View{v}
	runtime.SetFinalizer(obj, destroyView)
	return obj
}

func CreateViewFromRect(rect Rectf) *View {
	v := C.sfView_createFromRect(cRectf(&rect))
	obj := &View{v}
	runtime.SetFinalizer(obj, destroyView)
	return obj
}

func (v *View) Copy() *View {
	vv := C.sfView_copy(v.data)
	obj := &View{vv}
	runtime.SetFinalizer(obj, destroyView)
	return obj
}

func (v *View) SetCenter(center Vector2f) {
	C.sfView_setCenter(v.data, cVector2f(&center))
}

func (v *View) SetSize(size Vector2f) {
	C.sfView_setSize(v.data, cVector2f(&size))
}

func (v *View) SetRotation(rotation float32) {
	C.sfView_setRotation(v.data, C.float(rotation))
}

func (v *View) SetViewport(viewport Rectf) {
	C.sfView_setViewport(v.data, cRectf(&viewport))
}

func (v *View) Reset(rect Rectf) {
	C.sfView_reset(v.data, cRectf(&rect))
}

func (v *View) GetCenter() Vector2f {
	r := C.sfView_getCenter(v.data)
	return *goVector2f(&r)
}

func (v *View) GetSize() Vector2f {
	r := C.sfView_getSize(v.data)
	return *goVector2f(&r)
}

func (v *View) GetRotation() float32 {
	return float32(C.sfView_getRotation(v.data))
}

func (v *View) GetViewport() Rectf {
	r := C.sfView_getViewport(v.data)
	return *goRectf(&r)
}

func (v *View) Move(offset Vector2f) {
	C.sfView_move(v.data, cVector2f(&offset))
}

func (v *View) Rotate(angle float32) {
	C.sfView_rotate(v.data, C.float(angle))
}

func (v *View) Zoom(factor float32) {
	C.sfView_zoom(v.data, C.float(factor))
}
