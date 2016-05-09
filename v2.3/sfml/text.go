package sfml

//#include <SFML/Graphics.h>
import "C"

import (
	"runtime"
)

type Text struct {
	data *C.sfText
	font *Font
}

const (
	TextRegular       TextStyle = C.sfTextRegular
	TextBold          TextStyle = C.sfTextBold
	TextItalic        TextStyle = C.sfTextItalic
	TextUnderlined    TextStyle = C.sfTextUnderlined
	TextStrikeThrough TextStyle = C.sfTextStrikeThrough
)

type TextStyle int

func destroyText(t *Text) {
	C.sfText_destroy(t.data)
}

func CreateEmptyText() *Text {
	c := C.sfText_create()
	if c == nil {
		return nil
	}
	obj := &Text{data: c}
	runtime.SetFinalizer(obj, destroyText)
	return obj
}

func CreateText(text string, font *Font, size uint) *Text {
	c := C.sfText_create()
	if c == nil {
		return nil
	}
	C.sfText_setFont(c, font.data)
	C.sfText_setUnicodeString(c, cString(text))
	C.sfText_setCharacterSize(c, C.uint(size))
	obj := &Text{c, font}
	runtime.SetFinalizer(obj, destroyText)
	return obj
}

func (t *Text) Copy() *Text {
	c := C.sfText_copy(t.data)
	obj := &Text{c, t.font}
	runtime.SetFinalizer(obj, destroyText)
	return obj
}

func (t *Text) SetPosition(position Vector2f) {
	C.sfText_setPosition(t.data, cVector2f(&position))
}

func (t *Text) SetRotation(angle float32) {
	C.sfText_setRotation(t.data, C.float(angle))
}

func (t *Text) SetScale(scale Vector2f) {
	C.sfText_setScale(t.data, cVector2f(&scale))
}

func (t *Text) SetOrigin(origin Vector2f) {
	C.sfText_setOrigin(t.data, cVector2f(&origin))
}

func (t *Text) GetPosition() Vector2f {
	r := C.sfText_getPosition(t.data)
	return *goVector2f(&r)
}

func (t *Text) GetRotation() float32 {
	return float32(C.sfText_getRotation(t.data))
}

func (t *Text) GetScale() Vector2f {
	r := C.sfText_getScale(t.data)
	return *goVector2f(&r)
}

func (t *Text) GetOrigin() Vector2f {
	r := C.sfText_getOrigin(t.data)
	return *goVector2f(&r)
}

func (t *Text) Move(offset Vector2f) {
	C.sfText_move(t.data, cVector2f(&offset))
}

func (t *Text) Rotate(angle float32) {
	C.sfText_rotate(t.data, C.float(angle))
}

func (t *Text) Scale(factors Vector2f) {
	C.sfText_scale(t.data, cVector2f(&factors))
}

func (t *Text) GetTransform() Transform {
	r := C.sfText_getTransform(t.data)
	return *goTransform(&r)
}

func (t *Text) GetInverseTransform() Transform {
	r := C.sfText_getInverseTransform(t.data)
	return *goTransform(&r)
}

func (t *Text) SetString(text string) {
	C.sfText_setUnicodeString(t.data, cString(text))
}

func (t *Text) GetString() string {
	r := C.sfText_getString(t.data)
	return C.GoString(r)
}

func (t *Text) SetFont(font *Font) {
	C.sfText_setFont(t.data, font.data)
	t.font = font
}

func (t *Text) GetFont() *Font {
	return &Font{C.sfText_getFont(t.data)}
}

func (t *Text) SetCharacterSize(size uint) {
	C.sfText_setCharacterSize(t.data, C.uint(size))
}

func (t *Text) GetCharacterSize() uint {
	return uint(C.sfText_getCharacterSize(t.data))
}

func (t *Text) GetStyle() TextStyle {
	return TextStyle(C.sfText_getStyle(t.data))
}

func (t *Text) SetStyle(style TextStyle) {
	C.sfText_setStyle(t.data, C.sfUint32(style))
}

func (t *Text) SetColor(color Color) {
	C.sfText_setColor(t.data, cColor(&color))
}

func (t *Text) GetColor() Color {
	r := C.sfText_getColor(t.data)
	return *goColor(&r)
}

func (t *Text) FindCharacterPos(index int) Vector2f {
	r := C.sfText_findCharacterPos(t.data, C.size_t(index))
	return *goVector2f(&r)
}

func (t *Text) GetLocalBounds() Rectf {
	r := C.sfText_getLocalBounds(t.data)
	return *goRectf(&r)
}

func (t *Text) GetGlobalBounds() Rectf {
	r := C.sfText_getGlobalBounds(t.data)
	return *goRectf(&r)
}

func (t *Text) Draw(target RenderTarget, states *RenderStates) {
	ss := new(C.sfRenderStates)
	if states == nil {
		ss = nil
	} else {
		*ss = cRenderStates(states)
	}
	switch target.(type) {
	case *RenderTexture:
		C.sfRenderTexture_drawText(target.(*RenderTexture).data, t.data, ss)
	case *RenderWindow:
		C.sfRenderWindow_drawText(target.(*RenderWindow).data, t.data, ss)
	}
}
