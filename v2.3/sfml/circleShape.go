package sfml

//#include <SFML/Graphics.h>
import "C"

import (
	"runtime"
)

type CircleShape struct {
	data *C.sfCircleShape
	tex  *Texture
}

func destroyCircleShape(c *CircleShape) {
	C.sfCircleShape_destroy(c.data)
}

func CreateEmptyCircleShape() *CircleShape {
	c := C.sfCircleShape_create()
	if c == nil {
		return nil
	}
	obj := &CircleShape{data: c}
	runtime.SetFinalizer(obj, destroyCircleShape)
	return obj
}

func CreateCircleShape(radius float32, pointCount uint) *CircleShape {
	c := C.sfCircleShape_create()
	if c == nil {
		return nil
	}
	C.sfCircleShape_setRadius(c, C.float(radius))
	C.sfCircleShape_setPointCount(c, C.size_t(pointCount))
	obj := &CircleShape{data: c}
	runtime.SetFinalizer(obj, destroyCircleShape)
	return obj
}

func (c *CircleShape) Copy() *CircleShape {
	cs := C.sfCircleShape_copy(c.data)
	obj := &CircleShape{cs, c.tex}
	runtime.SetFinalizer(obj, destroyCircleShape)
	return obj
}

func (c *CircleShape) SetPosition(position Vector2f) {
	C.sfCircleShape_setPosition(c.data, cVector2f(&position))
}

func (c *CircleShape) SetRotation(angle float32) {
	C.sfCircleShape_setRotation(c.data, C.float(angle))
}

func (c *CircleShape) SetScale(scale Vector2f) {
	C.sfCircleShape_setScale(c.data, cVector2f(&scale))
}

func (c *CircleShape) SetOrigin(origin Vector2f) {
	C.sfCircleShape_setOrigin(c.data, cVector2f(&origin))
}

func (c *CircleShape) GetPosition() Vector2f {
	r := C.sfCircleShape_getPosition(c.data)
	return *goVector2f(&r)
}

func (c *CircleShape) GetRotation() float32 {
	return float32(C.sfCircleShape_getRotation(c.data))
}

func (c *CircleShape) GetScale() Vector2f {
	r := C.sfCircleShape_getScale(c.data)
	return *goVector2f(&r)
}

func (c *CircleShape) GetOrigin() Vector2f {
	r := C.sfCircleShape_getOrigin(c.data)
	return *goVector2f(&r)
}

func (c *CircleShape) Move(offset Vector2f) {
	C.sfCircleShape_move(c.data, cVector2f(&offset))
}

func (c *CircleShape) Rotate(angle float32) {
	C.sfCircleShape_rotate(c.data, C.float(angle))
}

func (c *CircleShape) Scale(factors Vector2f) {
	C.sfCircleShape_scale(c.data, cVector2f(&factors))
}

func (c *CircleShape) GetTransform() Transform {
	r := C.sfCircleShape_getTransform(c.data)
	return *goTransform(&r)
}

func (c *CircleShape) GetInverseTransform() Transform {
	r := C.sfCircleShape_getInverseTransform(c.data)
	return *goTransform(&r)
}

func (c *CircleShape) SetTexture(texture *Texture, resetRect bool) {
	C.sfCircleShape_setTexture(c.data, texture.data, cBool(resetRect))
	c.tex = texture
}

func (c *CircleShape) GetTexture() *Texture {
	return c.tex
}

func (c *CircleShape) GetTextureRect() Recti {
	r := C.sfCircleShape_getTextureRect(c.data)
	return *goRecti(&r)
}

func (c *CircleShape) SetTextureRect(rect Recti) {
	C.sfCircleShape_setTextureRect(c.data, cRecti(&rect))
}

func (c *CircleShape) SetFillColor(color Color) {
	C.sfCircleShape_setFillColor(c.data, cColor(&color))
}

func (c *CircleShape) GetFillColor() Color {
	r := C.sfCircleShape_getFillColor(c.data)
	return *goColor(&r)
}

func (c *CircleShape) SetOutlineColor(color Color) {
	C.sfCircleShape_setOutlineColor(c.data, cColor(&color))
}

func (c *CircleShape) GetOutlineColor() Color {
	r := C.sfCircleShape_getOutlineColor(c.data)
	return *goColor(&r)
}

func (c *CircleShape) SetOutlineThickness(thickness float32) {
	C.sfCircleShape_setOutlineThickness(c.data, C.float(thickness))
}

func (c *CircleShape) GetOutlineThickness() float32 {
	return float32(C.sfCircleShape_getOutlineThickness(c.data))
}

func (c *CircleShape) SetPointCount(count int) {
	C.sfCircleShape_setPointCount(c.data, C.size_t(count))
}

func (c *CircleShape) GetPointCount() int {
	return int(C.sfCircleShape_getPointCount(c.data))
}

func (c *CircleShape) SetRadius(radius float32) {
	C.sfCircleShape_setRadius(c.data, C.float(radius))
}

func (c *CircleShape) GetRadius() float32 {
	return float32(C.sfCircleShape_getRadius(c.data))
}

func (c *CircleShape) GetPoint(index int) Vector2f {
	p := C.sfCircleShape_getPoint(c.data, C.size_t(index))
	return *goVector2f(&p)
}

func (c *CircleShape) GetLocalBounds() Rectf {
	r := C.sfCircleShape_getLocalBounds(c.data)
	return *goRectf(&r)
}

func (c *CircleShape) GetGlobalBounds() Rectf {
	r := C.sfCircleShape_getGlobalBounds(c.data)
	return *goRectf(&r)
}

func (c *CircleShape) Draw(target RenderTarget) {
	switch target.(type) {
	case *RenderTexture:
		C.sfRenderTexture_drawCircleShape(target.(*RenderTexture).data, c.data, nil)
	case *RenderWindow:
		C.sfRenderWindow_drawCircleShape(target.(*RenderWindow).data, c.data, nil)
	}
}

func (c *CircleShape) DrawWithRenderStates(target RenderTarget, states *RenderStates) {
	ss := cRenderStates(states)
	switch target.(type) {
	case *RenderTexture:
		C.sfRenderTexture_drawCircleShape(target.(*RenderTexture).data, c.data, &ss)
	case *RenderWindow:
		C.sfRenderWindow_drawCircleShape(target.(*RenderWindow).data, c.data, &ss)
	}
}
