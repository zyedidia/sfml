package sfml

//#include <SFML/Graphics.h>
import "C"

import (
	"runtime"
)

type RectangleShape struct {
	data *C.sfRectangleShape
	tex  *Texture
}

func destroyRectangleShape(r *RectangleShape) {
	C.sfRectangleShape_destroy(r.data)
}

func CreateEmptyRectangleShape() *RectangleShape {
	r := C.sfRectangleShape_create()
	if r == nil {
		return nil
	}
	obj := &RectangleShape{data: r}
	runtime.SetFinalizer(obj, destroyRectangleShape)
	return obj
}

func CreateRectangleShape(size Vector2f) *RectangleShape {
	r := C.sfRectangleShape_create()
	if r == nil {
		return nil
	}
	C.sfRectangleShape_setSize(r, cVector2f(&size))
	obj := &RectangleShape{data: r}
	runtime.SetFinalizer(obj, destroyRectangleShape)
	return obj
}

func (r *RectangleShape) Copy() *RectangleShape {
	cs := C.sfRectangleShape_copy(r.data)
	obj := &RectangleShape{cs, r.tex}
	runtime.SetFinalizer(obj, destroyRectangleShape)
	return obj
}

func (r *RectangleShape) SetPosition(position Vector2f) {
	C.sfRectangleShape_setPosition(r.data, cVector2f(&position))
}

func (r *RectangleShape) SetRotation(angle float32) {
	C.sfRectangleShape_setRotation(r.data, C.float(angle))
}

func (r *RectangleShape) SetScale(scale Vector2f) {
	C.sfRectangleShape_setScale(r.data, cVector2f(&scale))
}

func (r *RectangleShape) SetOrigin(origin Vector2f) {
	C.sfRectangleShape_setOrigin(r.data, cVector2f(&origin))
}

func (r *RectangleShape) GetPosition() Vector2f {
	p := C.sfRectangleShape_getPosition(r.data)
	return *goVector2f(&p)
}

func (r *RectangleShape) GetRotation() float32 {
	return float32(C.sfRectangleShape_getRotation(r.data))
}

func (r *RectangleShape) GetScale() Vector2f {
	s := C.sfRectangleShape_getScale(r.data)
	return *goVector2f(&s)
}

func (r *RectangleShape) GetOrigin() Vector2f {
	o := C.sfRectangleShape_getOrigin(r.data)
	return *goVector2f(&o)
}

func (r *RectangleShape) Move(offset Vector2f) {
	C.sfRectangleShape_move(r.data, cVector2f(&offset))
}

func (r *RectangleShape) Rotate(angle float32) {
	C.sfRectangleShape_rotate(r.data, C.float(angle))
}

func (r *RectangleShape) Scale(factors Vector2f) {
	C.sfRectangleShape_scale(r.data, cVector2f(&factors))
}

func (r *RectangleShape) GetTransform() Transform {
	t := C.sfRectangleShape_getTransform(r.data)
	return *goTransform(&t)
}

func (r *RectangleShape) GetInverseTransform() Transform {
	it := C.sfRectangleShape_getInverseTransform(r.data)
	return *goTransform(&it)
}

func (r *RectangleShape) SetTexture(texture *Texture, resetRect bool) {
	C.sfRectangleShape_setTexture(r.data, texture.data, cBool(resetRect))
	r.tex = texture
}

func (r *RectangleShape) GetTexture() *Texture {
	return r.tex
}

func (r *RectangleShape) GetTextureRect() Recti {
	tr := C.sfRectangleShape_getTextureRect(r.data)
	return *goRecti(&tr)
}

func (r *RectangleShape) SetTextureRect(rect Recti) {
	C.sfRectangleShape_setTextureRect(r.data, cRecti(&rect))
}

func (r *RectangleShape) SetFillColor(color Color) {
	C.sfRectangleShape_setFillColor(r.data, cColor(&color))
}

func (r *RectangleShape) GetFillColor() Color {
	fc := C.sfRectangleShape_getFillColor(r.data)
	return *goColor(&fc)
}

func (r *RectangleShape) SetOutlineColor(color Color) {
	C.sfRectangleShape_setOutlineColor(r.data, cColor(&color))
}

func (r *RectangleShape) GetOutlineColor() Color {
	oc := C.sfRectangleShape_getOutlineColor(r.data)
	return *goColor(&oc)
}

func (r *RectangleShape) SetOutlineThickness(thickness float32) {
	C.sfRectangleShape_setOutlineThickness(r.data, C.float(thickness))
}

func (r *RectangleShape) GetOutlineThickness() float32 {
	return float32(C.sfRectangleShape_getOutlineThickness(r.data))
}

func (r *RectangleShape) GetPointCount() int {
	return int(C.sfRectangleShape_getPointCount(r.data))
}

func (r *RectangleShape) GetPoint(index int) Vector2f {
	p := C.sfRectangleShape_getPoint(r.data, C.size_t(index))
	return *goVector2f(&p)
}

func (r *RectangleShape) SetSize(Size Vector2f) {
	C.sfRectangleShape_setSize(r.data, cVector2f(&Size))
}

func (r *RectangleShape) GetSize() Vector2f {
	s := C.sfRectangleShape_getSize(r.data)
	return *goVector2f(&s)
}

func (r *RectangleShape) GetLocalBounds() Rectf {
	lb := C.sfRectangleShape_getLocalBounds(r.data)
	return *goRectf(&lb)
}

func (r *RectangleShape) GetGlobalBounds() Rectf {
	gb := C.sfRectangleShape_getGlobalBounds(r.data)
	return *goRectf(&gb)
}

func (r *RectangleShape) Draw(target RenderTarget) {
	switch target.(type) {
	case *RenderTexture:
		C.sfRenderTexture_drawRectangleShape(target.(*RenderTexture).data, r.data, nil)
	case *RenderWindow:
		C.sfRenderWindow_drawRectangleShape(target.(*RenderWindow).data, r.data, nil)
	}
}

func (r *RectangleShape) DrawWithRenderStates(target RenderTarget, states *RenderStates) {
	ss := cRenderStates(states)
	switch target.(type) {
	case *RenderTexture:
		C.sfRenderTexture_drawRectangleShape(target.(*RenderTexture).data, r.data, &ss)
	case *RenderWindow:
		C.sfRenderWindow_drawRectangleShape(target.(*RenderWindow).data, r.data, &ss)
	}
}
