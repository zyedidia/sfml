package sfml

//#include <SFML/Graphics.h>
import "C"

import (
	"runtime"
)

type Sprite struct {
	data *C.sfSprite
	tex  *Texture
}

func destroySprite(s *Sprite) {
	C.sfSprite_destroy(s.data)
}

func CreateEmptySprite() *Sprite {
	s := C.sfSprite_create()
	if s == nil {
		return nil
	}
	obj := &Sprite{data: s}
	runtime.SetFinalizer(obj, destroySprite)
	return obj
}

func CreateSprite(texture *Texture) *Sprite {
	s := C.sfSprite_create()
	if s == nil {
		return nil
	}
	C.sfSprite_setTexture(s, texture.data, C.sfTrue)
	obj := &Sprite{s, texture}
	runtime.SetFinalizer(obj, destroySprite)
	return obj
}

func CreateSpriteFromRect(texture *Texture, rectangle Recti) *Sprite {
	s := C.sfSprite_create()
	if s == nil {
		return nil
	}
	C.sfSprite_setTexture(s, texture.data, C.sfTrue)
	C.sfSprite_setTextureRect(s, cRecti(&rectangle))
	obj := &Sprite{s, texture}
	runtime.SetFinalizer(obj, destroySprite)
	return obj
}

func (s *Sprite) Copy() *Sprite {
	cs := C.sfSprite_copy(s.data)
	obj := &Sprite{cs, s.tex}
	runtime.SetFinalizer(obj, destroySprite)
	return obj
}

func (s *Sprite) SetPosition(position Vector2f) {
	C.sfSprite_setPosition(s.data, cVector2f(&position))
}

func (s *Sprite) SetRotation(angle float32) {
	C.sfSprite_setRotation(s.data, C.float(angle))
}

func (s *Sprite) SetScale(scale Vector2f) {
	C.sfSprite_setScale(s.data, cVector2f(&scale))
}

func (s *Sprite) SetOrigin(origin Vector2f) {
	C.sfSprite_setOrigin(s.data, cVector2f(&origin))
}

func (s *Sprite) GetPosition() Vector2f {
	r := C.sfSprite_getPosition(s.data)
	return *goVector2f(&r)
}

func (s *Sprite) GetRotation() float32 {
	return float32(C.sfSprite_getRotation(s.data))
}

func (s *Sprite) GetScale() Vector2f {
	r := C.sfSprite_getScale(s.data)
	return *goVector2f(&r)
}

func (s *Sprite) GetOrigin() Vector2f {
	r := C.sfSprite_getOrigin(s.data)
	return *goVector2f(&r)
}

func (s *Sprite) Move(offset Vector2f) {
	C.sfSprite_move(s.data, cVector2f(&offset))
}

func (s *Sprite) Rotate(angle float32) {
	C.sfSprite_rotate(s.data, C.float(angle))
}

func (s *Sprite) Scale(factors Vector2f) {
	C.sfSprite_scale(s.data, cVector2f(&factors))
}

func (s *Sprite) GetTransform() Transform {
	r := C.sfSprite_getTransform(s.data)
	return *goTransform(&r)
}

func (s *Sprite) GetInverseTransform() Transform {
	r := C.sfSprite_getInverseTransform(s.data)
	return *goTransform(&r)
}

func (s *Sprite) SetTexture(texture *Texture, resetRect bool) {
	C.sfSprite_setTexture(s.data, texture.data, cBool(resetRect))
	s.tex = texture
}

func (s *Sprite) SetTextureRect(rect Recti) {
	C.sfSprite_setTextureRect(s.data, cRecti(&rect))
}

func (s *Sprite) SetColor(color Color) {
	C.sfSprite_setColor(s.data, cColor(&color))
}

func (s *Sprite) GetColor() Color {
	r := C.sfSprite_getColor(s.data)
	return *goColor(&r)
}

func (s *Sprite) GetTexture() *Texture {
	return s.tex
}

func (s *Sprite) GetTextureRect() Recti {
	r := C.sfSprite_getTextureRect(s.data)
	return *goRecti(&r)
}

func (s *Sprite) GetLocalBounds() Rectf {
	r := C.sfSprite_getLocalBounds(s.data)
	return *goRectf(&r)
}

func (s *Sprite) GetGlobalBounds() Rectf {
	r := C.sfSprite_getGlobalBounds(s.data)
	return *goRectf(&r)
}

func (s *Sprite) Draw(target RenderTarget) {
	switch target.(type) {
	case *RenderTexture:
		C.sfRenderTexture_drawSprite(target.(*RenderTexture).data, s.data, nil)
	case *RenderWindow:
		C.sfRenderWindow_drawSprite(target.(*RenderWindow).data, s.data, nil)
	}
}

func (s *Sprite) DrawWithRenderStates(target RenderTarget, states *RenderStates) {
	ss := cRenderStates(states)
	switch target.(type) {
	case *RenderTexture:
		C.sfRenderTexture_drawSprite(target.(*RenderTexture).data, s.data, &ss)
	case *RenderWindow:
		C.sfRenderWindow_drawSprite(target.(*RenderWindow).data, s.data, &ss)
	}
}
