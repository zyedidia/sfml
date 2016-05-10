package sfml

//#include <SFML/Graphics.h>
//#include <stdlib.h>
import "C"

import (
	"runtime"
	"unsafe"
)

type Shader struct {
	data *C.sfShader
}

func destroyShader(s *Shader) {
	C.sfShader_destroy(s.data)
}

func CreateShader(vertexShaderFilename, fragmentShaderFilename string) *Shader {
	vshader := C.CString(vertexShaderFilename)
	defer C.free(unsafe.Pointer(vshader))
	fshader := C.CString(fragmentShaderFilename)
	defer C.free(unsafe.Pointer(fshader))
	s := C.sfShader_createFromFile(vshader, fshader)
	if s == nil {
		return nil
	}
	obj := &Shader{s}
	runtime.SetFinalizer(obj, destroyShader)
	return obj
}

func CreateShaderFromMemory(vertexShader, fragmentShader string) *Shader {
	vshader := C.CString(vertexShader)
	defer C.free(unsafe.Pointer(vshader))
	fshader := C.CString(fragmentShader)
	defer C.free(unsafe.Pointer(fshader))
	s := C.sfShader_createFromMemory(vshader, fshader)
	if s == nil {
		return nil
	}
	obj := &Shader{s}
	runtime.SetFinalizer(obj, destroyShader)
	return obj
}

func (s *Shader) SetFloatParameter(name string, xyzw ...float32) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	switch len(xyzw) {
	case 1:
		C.sfShader_setFloatParameter(s.data, cname, C.float(xyzw[0]))
	case 2:
		C.sfShader_setFloat2Parameter(s.data, cname, C.float(xyzw[0]), C.float(xyzw[1]))
	case 3:
		C.sfShader_setFloat3Parameter(s.data, cname, C.float(xyzw[0]), C.float(xyzw[1]), C.float(xyzw[2]))
	case 4:
		C.sfShader_setFloat4Parameter(s.data, cname, C.float(xyzw[0]), C.float(xyzw[1]), C.float(xyzw[2]), C.float(xyzw[3]))
	}
}

func (s *Shader) SetVector2Parameter(name string, vector Vector2f) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	C.sfShader_setVector2Parameter(s.data, cname, cVector2f(&vector))
}

func (s *Shader) SetVector3Parameter(name string, vector Vector3f) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	C.sfShader_setVector3Parameter(s.data, cname, cVector3f(&vector))
}

func (s *Shader) SetColorParameter(name string, color Color) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	C.sfShader_setColorParameter(s.data, cname, cColor(&color))
}

func (s *Shader) SetTransformParameter(name string, transform Transform) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	C.sfShader_setTransformParameter(s.data, cname, cTransform(&transform))
}

func (s *Shader) SetTextureParameter(name string, texture *Texture) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	C.sfShader_setTextureParameter(s.data, cname, texture.data)
}

func (s *Shader) SetCurrentTextureParameter(name string) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	C.sfShader_setCurrentTextureParameter(s.data, cname)
}

func BindShader(shader *Shader) {
	C.sfShader_bind(shader.data)
}

func ShaderIsAvailable() bool {
	return goBool(C.sfShader_isAvailable())
}