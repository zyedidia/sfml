package sfml

//#include <SFML/Graphics.h>
import "C"

import (
	"runtime"
)

type Transformable struct {
	data *C.sfTransformable
}

func destroyTransformable(t *Transformable) {
	C.sfTransformable_destroy(t.data)
}

func NewTransformable() *Transformable {
	r := C.sfTransformable_create()
	obj := &Transformable{r}
	runtime.SetFinalizer(obj, destroyTransformable)
	return obj
}

func (t *Transformable) Copy() *Transformable {
	r := C.sfTransformable_copy(t.data)
	obj := &Transformable{r}
	runtime.SetFinalizer(obj, destroyTransformable)
	return obj
}

func (t *Transformable) SetPosition(position Vector2f) {
	C.sfTransformable_setPosition(t.data, cVector2f(&position))
}

func (t *Transformable) SetRotation(angle float32) {
	C.sfTransformable_setRotation(t.data, C.float(angle))
}

func (t *Transformable) SetScale(scale Vector2f) {
	C.sfTransformable_setScale(t.data, cVector2f(&scale))
}

func (t *Transformable) SetOrigin(origin Vector2f) {
	C.sfTransformable_setOrigin(t.data, cVector2f(&origin))
}

func (t *Transformable) GetPosition() Vector2f {
	r := C.sfTransformable_getPosition(t.data)
	return *goVector2f(&r)
}

func (t *Transformable) GetRotation() float32 {
	return float32(C.sfTransformable_getRotation(t.data))
}

func (t *Transformable) GetScale() Vector2f {
	r := C.sfTransformable_getScale(t.data)
	return *goVector2f(&r)
}

func (t *Transformable) GetOrigin() Vector2f {
	r := C.sfTransformable_getOrigin(t.data)
	return *goVector2f(&r)
}

func (t *Transformable) Move(offset Vector2f) {
	C.sfTransformable_move(t.data, cVector2f(&offset))
}

func (t *Transformable) Rotate(angle float32) {
	C.sfTransformable_rotate(t.data, C.float(angle))
}

func (t *Transformable) Scale(factors Vector2f) {
	C.sfTransformable_scale(t.data, cVector2f(&factors))
}

func (t *Transformable) GetTransform() Transform {
	r := C.sfTransformable_getTransform(t.data)
	return *goTransform(&r)
}

func (t *Transformable) GetInverseTransform() Transform {
	r := C.sfTransformable_getInverseTransform(t.data)
	return *goTransform(&r)
}
