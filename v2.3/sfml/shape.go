package sfml

//#include <SFML/Graphics.h>
import "C"

import (
	"runtime"
)

type Shape struct {
	data *C.sfShape
	tex  *Texture
}

func destroyShape(s *Shape) {
	C.sfShape_destroy(s.data)
}

// FIXME
func CreateShape() *Shape {
	s := C.sfShape_create(nil, nil, nil)
	if s == nil {
		return nil
	}
	obj := &Shape{data: s}
	runtime.SetFinalizer(obj, destroyShape)
	return obj
}

func (s *Shape) SetPosition(position Vector2f) {
	C.sfShape_setPosition(s.data, cVector2f(&position))
}

func (s *Shape) SetRotation(angle float32) {
	C.sfShape_setRotation(s.data, C.float(angle))
}

func (s *Shape) SetScale(scale Vector2f) {
	C.sfShape_setScale(s.data, cVector2f(&scale))
}

func (s *Shape) SetOrigin(origin Vector2f) {
	C.sfShape_setOrigin(s.data, cVector2f(&origin))
}

func (s *Shape) GetPosition() Vector2f {
	r := C.sfShape_getPosition(s.data)
	return *goVector2f(&r)
}

func (s *Shape) GetRotation() float32 {
	return float32(C.sfShape_getRotation(s.data))
}

func (s *Shape) GetScale() Vector2f {
	r := C.sfShape_getScale(s.data)
	return *goVector2f(&r)
}

func (s *Shape) GetOrigin() Vector2f {
	r := C.sfShape_getOrigin(s.data)
	return *goVector2f(&r)
}

func (s *Shape) Move(offset Vector2f) {
	C.sfShape_move(s.data, cVector2f(&offset))
}

func (s *Shape) Rotate(angle float32) {
	C.sfShape_rotate(s.data, C.float(angle))
}

func (s *Shape) Scale(factors Vector2f) {
	C.sfShape_scale(s.data, cVector2f(&factors))
}

func (s *Shape) GetTransform() Transform {
	r := C.sfShape_getTransform(s.data)
	return *goTransform(&r)
}

func (s *Shape) GetInverseTransform() Transform {
	r := C.sfShape_getInverseTransform(s.data)
	return *goTransform(&r)
}

func (s *Shape) SetTexture(texture *Texture, resetRect bool) {
	C.sfShape_setTexture(s.data, texture.data, cBool(resetRect))
	s.tex = texture
}

func (s *Shape) GetTexture() *Texture {
	return s.tex
}

func (s *Shape) GetTextureRect() Recti {
	r := C.sfShape_getTextureRect(s.data)
	return *goRecti(&r)
}

func (s *Shape) SetTextureRect(rect Recti) {
	C.sfShape_setTextureRect(s.data, cRecti(&rect))
}

func (s *Shape) SetFillColor(color Color) {
	C.sfShape_setFillColor(s.data, cColor(&color))
}

func (s *Shape) GetFillColor() Color {
	r := C.sfShape_getFillColor(s.data)
	return *goColor(&r)
}

func (s *Shape) SetOutlineColor(color Color) {
	C.sfShape_setOutlineColor(s.data, cColor(&color))
}

func (s *Shape) GetOutlineColor() Color {
	r := C.sfShape_getOutlineColor(s.data)
	return *goColor(&r)
}

func (s *Shape) SetOutlineThickness(thickness float32) {
	C.sfShape_setOutlineThickness(s.data, C.float(thickness))
}

func (s *Shape) GetOutlineThickness() float32 {
	return float32(C.sfShape_getOutlineThickness(s.data))
}

func (s *Shape) GetPointCount() int {
	return int(C.sfShape_getPointCount(s.data))
}

func (s *Shape) GetPoint(index int) Vector2f {
	p := C.sfShape_getPoint(s.data, C.size_t(index))
	return *goVector2f(&p)
}

func (s *Shape) GetLocalBounds() Rectf {
	r := C.sfShape_getLocalBounds(s.data)
	return *goRectf(&r)
}

func (s *Shape) GetGlobalBounds() Rectf {
	r := C.sfShape_getGlobalBounds(s.data)
	return *goRectf(&r)
}

func (s *Shape) Update() {
	C.sfShape_update(s.data)
}

func (s *Shape) Draw(target RenderTarget) {
	switch target.(type) {
	case *RenderTexture:
		C.sfRenderTexture_drawShape(target.(*RenderTexture).data, s.data, nil)
	case *RenderWindow:
		C.sfRenderWindow_drawShape(target.(*RenderWindow).data, s.data, nil)
	}
}

func (s *Shape) DrawWithRenderStates(target RenderTarget, states *RenderStates) {
	ss := cRenderStates(states)
	switch target.(type) {
	case *RenderTexture:
		C.sfRenderTexture_drawShape(target.(*RenderTexture).data, s.data, &ss)
	case *RenderWindow:
		C.sfRenderWindow_drawShape(target.(*RenderWindow).data, s.data, &ss)
	}
}
