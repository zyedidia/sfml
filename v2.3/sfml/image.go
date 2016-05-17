package sfml

//#include <SFML/Graphics.h>
//#include <stdlib.h>
import "C"

import (
	"runtime"
	"unsafe"
)

type Image struct {
	data *C.sfImage
}

func destroyImage(i *Image) {
	C.sfImage_destroy(i.data)
}

func NewImage(filename string) *Image {
	file := C.CString(filename)
	defer C.free(unsafe.Pointer(file))
	i := C.sfImage_createFromFile(file)
	if i == nil {
		return nil
	}
	obj := &Image{i}
	runtime.SetFinalizer(obj, destroyImage)
	return obj
}

func NewEmptyImage(width, height uint) *Image {
	i := C.sfImage_create(C.uint(width), C.uint(height))
	if i == nil {
		return nil
	}
	obj := &Image{i}
	runtime.SetFinalizer(obj, destroyImage)
	return obj
}

func NewEmptyImageFromColor(width, height uint, color Color) *Image {
	i := C.sfImage_createFromColor(C.uint(width), C.uint(height), cColor(&color))
	if i == nil {
		return nil
	}
	obj := &Image{i}
	runtime.SetFinalizer(obj, destroyImage)
	return obj
}

func NewEmptyImageFromPixels(width, height uint, data []byte) *Image {
	i := C.sfImage_createFromPixels(C.uint(width), C.uint(height), (*C.sfUint8)(&data[0]))
	if i == nil {
		return nil
	}
	obj := &Image{i}
	runtime.SetFinalizer(obj, destroyImage)
	return obj
}

func NewImageFromMemory(data []byte) *Image {
	i := C.sfImage_createFromMemory(unsafe.Pointer(&data[0]), C.size_t(len(data)))
	if i == nil {
		return nil
	}
	obj := &Image{i}
	runtime.SetFinalizer(obj, destroyImage)
	return obj
}

func (i *Image) Copy() *Image {
	img := C.sfImage_copy(i.data)
	obj := &Image{img}
	runtime.SetFinalizer(obj, destroyImage)
	return obj
}

func (i *Image) SaveToFile(filename string) bool {
	file := C.CString(filename)
	defer C.free(unsafe.Pointer(file))
	return goBool(C.sfImage_saveToFile(i.data, file))
}

func (i *Image) GetSize() Vector2u {
	r := C.sfImage_getSize(i.data)
	return *goVector2u(&r)
}

func (i *Image) CreateMaskFromColor(color Color, alpha uint8) {
	C.sfImage_createMaskFromColor(i.data, cColor(&color), C.sfUint8(alpha))
}

func (i *Image) CopyImage(source *Image, destX, destY uint, sourceRect Recti, applyAlpha bool) {
	C.sfImage_copyImage(i.data, source.data, C.uint(destX), C.uint(destY), cRecti(&sourceRect), cBool(applyAlpha))
}

func (i *Image) SetPixel(x, y uint, color Color) {
	C.sfImage_setPixel(i.data, C.uint(x), C.uint(y), cColor(&color))
}

func (i *Image) GetPixel(x, y uint) Color {
	r := C.sfImage_getPixel(i.data, C.uint(x), C.uint(y))
	return *goColor(&r)
}

func (i *Image) FlipHorizontally() {
	C.sfImage_flipHorizontally(i.data)
}

func (i *Image) FlipVertically() {
	C.sfImage_flipVertically(i.data)
}
