package sfml

//#include <SFML/Window.h>
import "C"

import (
	"runtime"
)

type Context struct {
	data *C.sfContext
}

func destroyContext(c *Context) {
	C.sfContext_destroy(c.data)
}

func NewContext() *Context {
	c := C.sfContext_create()
	obj := &Context{c}
	runtime.SetFinalizer(obj, destroyContext)
	return obj
}

func (c *Context) SetActive(active bool) {
	C.sfContext_setActive(c.data, cBool(active))
}
