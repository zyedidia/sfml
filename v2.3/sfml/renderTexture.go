package sfml

//#include <SFML/Graphics.h>
import "C"

import (
	"runtime"
)

type RenderTexture struct {
	data *C.sfRenderTexture
}

func destroyRenderTexture(t *Texture) {
	C.sfRenderTexture_destroy(t.data)
}

func NewRenderTexture(width, height uint, depthBuffer bool) *RenderTexture {
	r := C.sfRenderTexture_create(C.uint(width), C.uint(height), cBool(depthBuffer))
	if r == nil {
		return nil
	}
	obj := &RenderTexture{r}
	runtime.SetFinalizer(obj, destroyRenderTexture)
	return obj
}

func (r *RenderTexture) GetSize() Vector2u {
	v := C.sfRenderTexture_getSize(r.data)
	return *goVector2u(&v)
}

func (r *RenderTexture) SetActive(active bool) bool {
	return goBool(C.sfRenderTexture_setActive(r.data, cBool(active)))
}

func (r *RenderTexture) Display() {
	C.sfRenderTexture_display(r.data)
}

func (r *RenderTexture) Clear(color Color) {
	C.sfRenderTexture_clear(r.data, cColor(&color))
}

func (r *RenderTexture) SetView(view *View) {
	C.sfRenderTexture_setView(r.data, view.data)
}

func (r *RenderTexture) GetView() *View {
	return &View{C.sfRenderTexture_getView(r.data)}
}

func (r *RenderTexture) GetDefaultView() *View {
	return &View{C.sfRenderTexture_getDefaultView(r.data)}
}

func (r *RenderTexture) GetViewport(view *View) Recti {
	rr := C.sfRenderTexture_getViewport(r.data, view.data)
	return *goRecti(&rr)
}

func (r *RenderTexture) MapPixelToCoords(point Vector2i, view *View) Vector2f {
	rr := C.sfRenderTexture_mapPixelToCoords(r.data, cVector2i(&point), view.data)
	return *goVector2f(&rr)
}

func (r *RenderTexture) MapCoordsToPixel(point Vector2f, view *View) Vector2i {
	rr := C.sfRenderTexture_mapCoordsToPixel(r.data, cVector2f(&point), view.data)
	return *goVector2i(&rr)
}

func (r *RenderTexture) Draw(object Drawable) {
	switch object.(type) {
	case *Sprite:
		C.sfRenderTexture_drawSprite(r.data, object.(*Sprite).data, nil)
	case *Text:
		C.sfRenderTexture_drawText(r.data, object.(*Text).data, nil)
	/*case *Shape:
	C.sfRenderTexture_drawShape(r.data, object.(*Shape).data, nil)*/
	case *CircleShape:
		C.sfRenderTexture_drawCircleShape(r.data, object.(*CircleShape).data, nil)
	case *ConvexShape:
		C.sfRenderTexture_drawConvexShape(r.data, object.(*ConvexShape).data, nil)
	case *RectangleShape:
		C.sfRenderTexture_drawRectangleShape(r.data, object.(*RectangleShape).data, nil)
	}
}

func (r *RenderTexture) DrawWithRenderStates(object Drawable, states *RenderStates) {
	s := cRenderStates(states)
	switch object.(type) {
	case *Sprite:
		C.sfRenderTexture_drawSprite(r.data, object.(*Sprite).data, &s)
	case *Text:
		C.sfRenderTexture_drawText(r.data, object.(*Text).data, &s)
	/*case *Shape:
	C.sfRenderTexture_drawShape(r.data, object.(*Shape).data, &s)*/
	case *CircleShape:
		C.sfRenderTexture_drawCircleShape(r.data, object.(*CircleShape).data, &s)
	case *ConvexShape:
		C.sfRenderTexture_drawConvexShape(r.data, object.(*ConvexShape).data, &s)
	case *RectangleShape:
		C.sfRenderTexture_drawRectangleShape(r.data, object.(*RectangleShape).data, &s)
	}
}

func (r *RenderTexture) DrawPrimitives(object *VertexArray) {
	C.sfRenderTexture_drawVertexArray(r.data, object.data, nil)
}

func (r *RenderTexture) DrawPrimitivesWithRenderStates(object *VertexArray, states *RenderStates) {
	s := cRenderStates(states)
	C.sfRenderTexture_drawVertexArray(r.data, object.data, &s)
}

func (r *RenderTexture) PushGLStates() {
	C.sfRenderTexture_pushGLStates(r.data)
}

func (r *RenderTexture) PopGLStates() {
	C.sfRenderTexture_popGLStates(r.data)
}

func (r *RenderTexture) ResetGLStates() {
	C.sfRenderTexture_resetGLStates(r.data)
}

func (r *RenderTexture) SetSmooth(smooth bool) {
	C.sfRenderTexture_setSmooth(r.data, cBool(smooth))
}

func (r *RenderTexture) IsSmooth() bool {
	return goBool(C.sfRenderTexture_isSmooth(r.data))
}

func (r *RenderTexture) SetRepeated(repeated bool) {
	C.sfRenderTexture_setRepeated(r.data, cBool(repeated))
}

func (r *RenderTexture) IsRepeated() bool {
	return goBool(C.sfRenderTexture_isRepeated(r.data))
}
