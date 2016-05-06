package sfml

//#include <SFML/Graphics.h>
//#include <stdlib.h>
import "C"

import (
	"runtime"
	"unsafe"
)

type Texture struct {
	data *C.sfTexture
}

func destroyTexture(t *Texture) {
	C.sfTexture_destroy(t.data)
}

func CreateTexture(width, height uint) *Texture {
	c := C.sfTexture_create(C.uint(width), C.uint(height))
	if c == nil {
		return nil
	}
	obj := &Texture{c}
	runtime.SetFinalizer(obj, destroyTexture)
	return obj
}

func CreateTextureFromFile(filename string) *Texture {
	file := C.CString(filename)
	defer C.free(unsafe.Pointer(file))
	c := C.sfTexture_createFromFile(file, nil)
	if c == nil {
		return nil
	}
	obj := &Texture{c}
	runtime.SetFinalizer(obj, destroyTexture)
	return obj
}

func CreateTextureFromMemory(data []byte) *Texture {
	c := C.sfTexture_createFromMemory(unsafe.Pointer(&data[0]), C.size_t(len(data)), nil)
	if c == nil {
		return nil
	}
	obj := &Texture{c}
	runtime.SetFinalizer(obj, destroyTexture)
	return obj
}

func CreateTextureFromImage(image *Image) *Texture {
	c := C.sfTexture_createFromImage(image.data, nil)
	if c == nil {
		return nil
	}
	obj := &Texture{c}
	runtime.SetFinalizer(obj, destroyTexture)
	return obj
}

func (t *Texture) Copy() *Texture {
	c := C.sfTexture_copy(t.data)
	obj := &Texture{c}
	runtime.SetFinalizer(obj, destroyTexture)
	return obj
}

func (t *Texture) GetSize() Vector2u {
	r := C.sfTexture_getSize(t.data)
	return *goVector2u(&r)
}

func (t *Texture) CopyToImage() *Image {
	return &Image{C.sfTexture_copyToImage(t.data)}
}

func (t *Texture) UpdateFromPixels(pixels []byte, width, height uint, offset Vector2u) {
	C.sfTexture_updateFromPixels(t.data, (*C.sfUint8)(&pixels[0]), C.uint(width), C.uint(height), C.uint(offset.X), C.uint(offset.Y))
}

func (t *Texture) UpdateFromImage(image *Image, offset Vector2u) {
	C.sfTexture_updateFromImage(t.data, image.data, C.uint(offset.X), C.uint(offset.Y))
}

func (t *Texture) UpdateFromWindow(window SystemWindow, offset Vector2u) {
	switch window.(type) {
	case *RenderWindow:
		C.sfTexture_updateFromRenderWindow(t.data, window.(*RenderWindow).data, C.uint(offset.X), C.uint(offset.Y))
	case *Window:
		C.sfTexture_updateFromWindow(t.data, window.(*Window).data, C.uint(offset.X), C.uint(offset.Y))
	}
}

func (t *Texture) SetSmooth(smooth bool) {
	C.sfTexture_setSmooth(t.data, cBool(smooth))
}

func (t *Texture) IsSmooth() bool {
	return goBool(C.sfTexture_isSmooth(t.data))
}

func (t *Texture) SetRepeated(repeated bool) {
	C.sfTexture_setRepeated(t.data, cBool(repeated))
}

func (t *Texture) IsRepeated() bool {
	return goBool(C.sfTexture_isRepeated(t.data))
}

func (t *Texture) GetNativeGandle() uint {
	return uint(C.sfTexture_getNativeHandle(t.data))
}

func BindTexture(texture *Texture) {
	C.sfTexture_bind(texture.data)
}

func GetMaximumTextureSize() uint {
	return uint(C.sfTexture_getMaximumSize())
}
