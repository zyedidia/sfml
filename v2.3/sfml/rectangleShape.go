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

func destroyRectangleShape(c *RectangleShape) {
	C.sfRectangleShape_destroy(c.data)
}

func CreateEmptyRectangleShape() *RectangleShape {
	c := C.sfRectangleShape_create()
	if c == nil {
		return nil
	}
	obj := &RectangleShape{data: c}
	runtime.SetFinalizer(obj, destroyRectangleShape)
	return obj
}

func CreateRectangleShape(size Vector2f) *RectangleShape {
	c := C.sfRectangleShape_create()
	if c == nil {
		return nil
	}
	C.sfRectangleShape_setSize(c, cVector2f(&size))
	obj := &RectangleShape{data: c}
	runtime.SetFinalizer(obj, destroyRectangleShape)
	return obj
}

func (c *RectangleShape) Copy() *RectangleShape {
	cs := C.sfRectangleShape_copy(c.data)
	obj := &RectangleShape{cs, c.tex}
	runtime.SetFinalizer(obj, destroyRectangleShape)
	return obj
}

func (c *RectangleShape) SetPosition(position Vector2f) {
	C.sfRectangleShape_setPosition(c.data, cVector2f(&position))
}

func (c *RectangleShape) SetRotation(angle float32) {
	C.sfRectangleShape_setRotation(c.data, C.float(angle))
}

func (c *RectangleShape) SetScale(scale Vector2f) {
	C.sfRectangleShape_setScale(c.data, cVector2f(&scale))
}

func (c *RectangleShape) SetOrigin(origin Vector2f) {
	C.sfRectangleShape_setOrigin(c.data, cVector2f(&origin))
}

func (c *RectangleShape) GetPosition() Vector2f {
	r := C.sfRectangleShape_getPosition(c.data)
	return *goVector2f(&r)
}

func (c *RectangleShape) GetRotation() float32 {
	return float32(C.sfRectangleShape_getRotation(c.data))
}

func (c *RectangleShape) GetScale() Vector2f {
	r := C.sfRectangleShape_getScale(c.data)
	return *goVector2f(&r)
}

func (c *RectangleShape) GetOrigin() Vector2f {
	r := C.sfRectangleShape_getOrigin(c.data)
	return *goVector2f(&r)
}

func (c *RectangleShape) Move(offset Vector2f) {
	C.sfRectangleShape_move(c.data, cVector2f(&offset))
}

func (c *RectangleShape) Rotate(angle float32) {
	C.sfRectangleShape_rotate(c.data, C.float(angle))
}

func (c *RectangleShape) Scale(factors Vector2f) {
	C.sfRectangleShape_scale(c.data, cVector2f(&factors))
}

func (c *RectangleShape) GetTransform() Transform {
	r := C.sfRectangleShape_getTransform(c.data)
	return *goTransform(&r)
}

func (c *RectangleShape) GetInverseTransform() Transform {
	r := C.sfRectangleShape_getInverseTransform(c.data)
	return *goTransform(&r)
}

func (c *RectangleShape) SetTexture(texture *Texture, resetRect bool) {
	C.sfRectangleShape_setTexture(c.data, texture.data, cBool(resetRect))
	c.tex = texture
}

func (c *RectangleShape) GetTexture() *Texture {
	return c.tex
}

func (c *RectangleShape) GetTextureRect() Recti {
	r := C.sfRectangleShape_getTextureRect(c.data)
	return *goRecti(&r)
}

func (c *RectangleShape) SetTextureRect(rect Recti) {
	C.sfRectangleShape_setTextureRect(c.data, cRecti(&rect))
}

func (c *RectangleShape) SetFillColor(color Color) {
	C.sfRectangleShape_setFillColor(c.data, cColor(&color))
}

func (c *RectangleShape) GetFillColor() Color {
	r := C.sfRectangleShape_getFillColor(c.data)
	return *goColor(&r)
}

func (c *RectangleShape) SetOutlineColor(color Color) {
	C.sfRectangleShape_setOutlineColor(c.data, cColor(&color))
}

func (c *RectangleShape) GetOutlineColor() Color {
	r := C.sfRectangleShape_getOutlineColor(c.data)
	return *goColor(&r)
}

func (c *RectangleShape) SetOutlineThickness(thickness float32) {
	C.sfRectangleShape_setOutlineThickness(c.data, C.float(thickness))
}

func (c *RectangleShape) GetOutlineThickness() float32 {
	return float32(C.sfRectangleShape_getOutlineThickness(c.data))
}

func (c *RectangleShape) SetPointCount(count int) {
	C.sfRectangleShape_setPointCount(c.data, C.size_t(count))
}

func (c *RectangleShape) GetPointCount() int {
	return int(C.sfRectangleShape_getPointCount(c.data))
}

func (c *RectangleShape) SetPoint(index int, point Vector2f) {
	C.sfRectangleShape_setPoint(c.data, C.size_t(index), cVector2f(&point))
}

func (c *RectangleShape) GetPoint(index int) Vector2f {
	p := C.sfRectangleShape_getPoint(c.data, C.size_t(index))
	return *goVector2f(&p)
}

func (c *RectangleShape) SetSize(Size Vector2f) {
	C.sfRectangleShape_setSize(c.data, cVector2f(&Size))
}

func (c *RectangleShape) GetSize() Vector2f {
	p := C.sfRectangleShape_getSize(c.data)
	return *goVector2f(&p)
}

func (c *RectangleShape) GetLocalBounds() Rectf {
	r := C.sfRectangleShape_getLocalBounds(c.data)
	return *goRectf(&r)
}

func (c *RectangleShape) GetGlobalBounds() Rectf {
	r := C.sfRectangleShape_getGlobalBounds(c.data)
	return *goRectf(&r)
}

func (c *RectangleShape) Draw(target RenderTarget) {
	switch target.(type) {
	case *RenderTexture:
		C.sfRenderTexture_drawRectangleShape(target.(*RenderTexture).data, c.data, nil)
	case *RenderWindow:
		C.sfRenderWindow_drawRectangleShape(target.(*RenderWindow).data, c.data, nil)
	}
}

func (c *RectangleShape) DrawWithRenderStates(target RenderTarget, states *RenderStates) {
	ss := cRenderStates(states)
	switch target.(type) {
	case *RenderTexture:
		C.sfRenderTexture_drawRectangleShape(target.(*RenderTexture).data, c.data, &ss)
	case *RenderWindow:
		C.sfRenderWindow_drawRectangleShape(target.(*RenderWindow).data, c.data, &ss)
	}
}
