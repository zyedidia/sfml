package sfml

//#include <SFML/Graphics.h>
import "C"

import (
	"runtime"
)

type ConvexShape struct {
	data *C.sfConvexShape
	tex  *Texture
}

func destroyConvexShape(c *ConvexShape) {
	C.sfConvexShape_destroy(c.data)
}

func NewEmptyConvexShape() *ConvexShape {
	c := C.sfConvexShape_create()
	if c == nil {
		return nil
	}
	obj := &ConvexShape{data: c}
	runtime.SetFinalizer(obj, destroyConvexShape)
	return obj
}

func NewConvexShape(pointCount uint) *ConvexShape {
	c := C.sfConvexShape_create()
	if c == nil {
		return nil
	}
	C.sfConvexShape_setPointCount(c, C.size_t(pointCount))
	obj := &ConvexShape{data: c}
	runtime.SetFinalizer(obj, destroyConvexShape)
	return obj
}

func (c *ConvexShape) Copy() *ConvexShape {
	cs := C.sfConvexShape_copy(c.data)
	obj := &ConvexShape{cs, c.tex}
	runtime.SetFinalizer(obj, destroyConvexShape)
	return obj
}

func (c *ConvexShape) SetPosition(position Vector2f) {
	C.sfConvexShape_setPosition(c.data, cVector2f(&position))
}

func (c *ConvexShape) SetRotation(angle float32) {
	C.sfConvexShape_setRotation(c.data, C.float(angle))
}

func (c *ConvexShape) SetScale(scale Vector2f) {
	C.sfConvexShape_setScale(c.data, cVector2f(&scale))
}

func (c *ConvexShape) SetOrigin(origin Vector2f) {
	C.sfConvexShape_setOrigin(c.data, cVector2f(&origin))
}

func (c *ConvexShape) GetPosition() Vector2f {
	r := C.sfConvexShape_getPosition(c.data)
	return *goVector2f(&r)
}

func (c *ConvexShape) GetRotation() float32 {
	return float32(C.sfConvexShape_getRotation(c.data))
}

func (c *ConvexShape) GetScale() Vector2f {
	r := C.sfConvexShape_getScale(c.data)
	return *goVector2f(&r)
}

func (c *ConvexShape) GetOrigin() Vector2f {
	r := C.sfConvexShape_getOrigin(c.data)
	return *goVector2f(&r)
}

func (c *ConvexShape) Move(offset Vector2f) {
	C.sfConvexShape_move(c.data, cVector2f(&offset))
}

func (c *ConvexShape) Rotate(angle float32) {
	C.sfConvexShape_rotate(c.data, C.float(angle))
}

func (c *ConvexShape) Scale(factors Vector2f) {
	C.sfConvexShape_scale(c.data, cVector2f(&factors))
}

func (c *ConvexShape) GetTransform() Transform {
	r := C.sfConvexShape_getTransform(c.data)
	return *goTransform(&r)
}

func (c *ConvexShape) GetInverseTransform() Transform {
	r := C.sfConvexShape_getInverseTransform(c.data)
	return *goTransform(&r)
}

func (c *ConvexShape) SetTexture(texture *Texture, resetRect bool) {
	C.sfConvexShape_setTexture(c.data, texture.data, cBool(resetRect))
	c.tex = texture
}

func (c *ConvexShape) GetTexture() *Texture {
	return c.tex
}

func (c *ConvexShape) GetTextureRect() Recti {
	r := C.sfConvexShape_getTextureRect(c.data)
	return *goRecti(&r)
}

func (c *ConvexShape) SetTextureRect(rect Recti) {
	C.sfConvexShape_setTextureRect(c.data, cRecti(&rect))
}

func (c *ConvexShape) SetFillColor(color Color) {
	C.sfConvexShape_setFillColor(c.data, cColor(&color))
}

func (c *ConvexShape) GetFillColor() Color {
	r := C.sfConvexShape_getFillColor(c.data)
	return *goColor(&r)
}

func (c *ConvexShape) SetOutlineColor(color Color) {
	C.sfConvexShape_setOutlineColor(c.data, cColor(&color))
}

func (c *ConvexShape) GetOutlineColor() Color {
	r := C.sfConvexShape_getOutlineColor(c.data)
	return *goColor(&r)
}

func (c *ConvexShape) SetOutlineThickness(thickness float32) {
	C.sfConvexShape_setOutlineThickness(c.data, C.float(thickness))
}

func (c *ConvexShape) GetOutlineThickness() float32 {
	return float32(C.sfConvexShape_getOutlineThickness(c.data))
}

func (c *ConvexShape) SetPointCount(count int) {
	C.sfConvexShape_setPointCount(c.data, C.size_t(count))
}

func (c *ConvexShape) GetPointCount() int {
	return int(C.sfConvexShape_getPointCount(c.data))
}

func (c *ConvexShape) SetPoint(index int, point Vector2f) {
	C.sfConvexShape_setPoint(c.data, C.size_t(index), cVector2f(&point))
}

func (c *ConvexShape) GetPoint(index int) Vector2f {
	p := C.sfConvexShape_getPoint(c.data, C.size_t(index))
	return *goVector2f(&p)
}

func (c *ConvexShape) GetLocalBounds() Rectf {
	r := C.sfConvexShape_getLocalBounds(c.data)
	return *goRectf(&r)
}

func (c *ConvexShape) GetGlobalBounds() Rectf {
	r := C.sfConvexShape_getGlobalBounds(c.data)
	return *goRectf(&r)
}

func (c *ConvexShape) Draw(target RenderTarget) {
	switch target.(type) {
	case *RenderTexture:
		C.sfRenderTexture_drawConvexShape(target.(*RenderTexture).data, c.data, nil)
	case *RenderWindow:
		C.sfRenderWindow_drawConvexShape(target.(*RenderWindow).data, c.data, nil)
	}
}

func (c *ConvexShape) DrawWithRenderStates(target RenderTarget, states *RenderStates) {
	ss := cRenderStates(states)
	switch target.(type) {
	case *RenderTexture:
		C.sfRenderTexture_drawConvexShape(target.(*RenderTexture).data, c.data, &ss)
	case *RenderWindow:
		C.sfRenderWindow_drawConvexShape(target.(*RenderWindow).data, c.data, &ss)
	}
}
