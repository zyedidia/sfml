package sfml

//#include <SFML/Graphics.h>
//#include <stdlib.h>
import "C"

import (
	"runtime"
	"unsafe"
)

type Font struct {
	data *C.sfFont
}

func destroyFont(f *Font) {
	C.sfFont_destroy(f.data)
}

func CreateFont(filename string) *Font {
	file := C.CString(filename)
	defer C.free(unsafe.Pointer(file))
	f := C.sfFont_createFromFile(file)
	if f == nil {
		return nil
	}
	obj := &Font{f}
	runtime.SetFinalizer(obj, destroyFont)
	return obj
}

func CreateFontFromMemory(data []byte) *Font {
	f := C.sfFont_createFromMemory(unsafe.Pointer(&data[0]), C.size_t(len(data)))
	if f == nil {
		return nil
	}
	obj := &Font{f}
	runtime.SetFinalizer(obj, destroyFont)
	return obj
}

func (f *Font) Copy() *Font {
	cf := C.sfFont_copy(f.data)
	obj := &Font{cf}
	runtime.SetFinalizer(obj, destroyFont)
	return obj
}

func (f *Font) GetGlyph(codePoint rune, characterSize uint, bold bool) Glyph {
	r := C.sfFont_getGlyph(f.data, C.sfUint32(codePoint), C.uint(characterSize), cBool(bold))
	return *goGlyph(&r)
}

func (f *Font) GetKerning(first, second rune, characterSize uint) float32 {
	return float32(C.sfFont_getKerning(f.data, C.sfUint32(first), C.sfUint32(second), C.uint(characterSize)))
}

func (f *Font) GetLineSpacing(characterSize uint) float32 {
	return float32(C.sfFont_getLineSpacing(f.data, C.uint(characterSize)))
}

func (f *Font) GetUnderlinePosition(characterSize uint) float32 {
	return float32(C.sfFont_getUnderlinePosition(f.data, C.uint(characterSize)))
}

func (f *Font) GetUnderlineThickness(characterSize uint) float32 {
	return float32(C.sfFont_getUnderlineThickness(f.data, C.uint(characterSize)))
}

func (f *Font) GetTexture(characterSize uint) *Texture {
	return &Texture{C.sfFont_getTexture(f.data, C.uint(characterSize))}
}

func (f *Font) GetInfo() string {
	r := C.sfFont_getInfo(f.data)
	return C.GoString(r.family)
}
