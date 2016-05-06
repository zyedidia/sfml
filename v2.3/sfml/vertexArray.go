package sfml

//#include <SFML/Graphics.h>
import "C"

import (
	"runtime"
)

type VertexArray struct {
	data *C.sfVertexArray
}

func CreateVertexArray(pType PrimitiveType) *VertexArray {
	r := C.sfVertexArray_create()
	runtime.SetFinalizer(r, C.sfVertexArray_destroy)
	C.sfVertexArray_setPrimitiveType(r, C.sfPrimitiveType(pType))
	return &VertexArray{r}
}

func CreateVertexArrayFromSlice(vertices []Vertex, pType PrimitiveType) *VertexArray {
	r := C.sfVertexArray_create()
	runtime.SetFinalizer(r, C.sfVertexArray_destroy)
	for _, v := range vertices {
		C.sfVertexArray_append(r, cVertex(&v))
	}
	C.sfVertexArray_setPrimitiveType(r, C.sfPrimitiveType(pType))
	return &VertexArray{r}
}

func (v *VertexArray) Copy() *VertexArray {
	r := C.sfVertexArray_copy(v.data)
	runtime.SetFinalizer(r, C.sfVertexArray_destroy)
	return &VertexArray{r}
}

func (v *VertexArray) GetVertexCount() int {
	return int(C.sfVertexArray_getVertexCount(v.data))
}

func (v *VertexArray) GetVertex(index int) Vertex {
	r := C.sfVertexArray_getVertex(v.data, C.size_t(index))
	return *goVertex(r)
}

func (v *VertexArray) Clear() {
	C.sfVertexArray_clear(v.data)
}

func (v *VertexArray) Resize(vertexCount int) {
	C.sfVertexArray_resize(v.data, C.size_t(vertexCount))
}

func (v *VertexArray) Append(vertex Vertex) {
	C.sfVertexArray_append(v.data, cVertex(&vertex))
}

func (v *VertexArray) SetPrimitiveType(pType PrimitiveType) {
	C.sfVertexArray_setPrimitiveType(v.data, C.sfPrimitiveType(pType))
}

func (v *VertexArray) GetPrimitiveType() PrimitiveType {
	return PrimitiveType(C.sfVertexArray_getPrimitiveType(v.data))
}

func (v *VertexArray) GetBounds() Rectf {
	r := C.sfVertexArray_getBounds(v.data)
	return *goRectf(&r)
}

func (v *VertexArray) Draw(target RenderTarget, states *RenderStates) {
	s := new(C.sfRenderStates)
	if states == nil {
		s = nil
	} else {
		*s = cRenderStates(states)
	}
	switch target.(type) {
	case *RenderTexture:
		C.sfRenderTexture_drawSprite(target.(*RenderTexture).data, v.data, s)
	case *RenderWindow:
		C.sfRenderWindow_drawSprite(target.(*RenderWindow).data, v.data, s)
	}
}
