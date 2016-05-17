package sfml

//#include <SFML/Graphics.h>
//#include "transform.h"
import "C"

var TransformIdentity = Transform{1, 0, 0, 0, 1, 0, 0, 0, 1}

type Transform [9]float32
type Matrix [16]float32

func (t Transform) GetMatrix() Matrix {
	var m Matrix
	ct := cTransform(&t)
	C.sfTransform_getMatrix(&ct, (*C.float)(&m[0]))
	return m
}

func (t Transform) GetInverse() Transform {
	ct := cTransform(&t)
	r := C.sfTransform_getInverse(&ct)
	return *goTransform(&r)
}

func (t Transform) TransformPoint(point Vector2f) Vector2f {
	ct := cTransform(&t)
	r := C.sfTransform_transformPoint(&ct, cVector2f(&point))
	return *goVector2f(&r)
}

func (t Transform) TransformRect(rect Rectf) Rectf {
	ct := cTransform(&t)
	r := C.sfTransform_transformRect(&ct, cRectf(&rect))
	return *goRectf(&r)
}

func (t Transform) Combine(other Transform) {
	ct := cTransform(&t)
	ot := cTransform(&other)
	C.sfTransform_combine(&ct, &ot)
}

func (t Transform) Translate(offset Vector2f) {
	ct := cTransform(&t)
	C.sfTransform_translate(&ct, C.float(offset.X), C.float(offset.Y))
}

func (t Transform) Rotate(angle float32) {
	ct := cTransform(&t)
	C.sfTransform_rotate(&ct, C.float(angle))
}

func (t Transform) RotateWithCenter(angle float32, center Vector2f) {
	ct := cTransform(&t)
	C.sfTransform_rotateWithCenter(&ct, C.float(angle), C.float(center.X), C.float(center.Y))
}

func (t Transform) Scale(scale Vector2f) {
	ct := cTransform(&t)
	C.sfTransform_scale(&ct, C.float(scale.X), C.float(scale.Y))
}

func (t Transform) ScaleWithCenter(scale, center Vector2f) {
	ct := cTransform(&t)
	C.sfTransform_scaleWithCenter(&ct, C.float(scale.X), C.float(scale.Y), C.float(center.X), C.float(center.Y))
}

func goTransform(t *C.sfTransform) *Transform {
	r := new(Transform)
	for i, _ := range r {
		r[i] = float32(C.GetTransformAtIndex(t, C.int(i)))
	}
	return r
}

func cTransform(t *Transform) C.sfTransform {
	var r C.sfTransform
	for i, v := range t {
		C.SetTransformAtIndex(&r, C.int(i), C.float(v))
	}
	return r
}
