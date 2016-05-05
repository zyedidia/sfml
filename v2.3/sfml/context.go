package sfml

//#include <SFML/Window.h>
import "C"

import (
	"runtime"
)

type Context struct {
	data *C.sfContext
}

func CreateContext() *Context {
	c := C.sfContext_create()
	runtime.SetFinalizer(c, C.sfContext_destroy(c))
	return &Context{c}
}

func (c *Context) SetActive(active bool) {
	C.sfContext_setActive(c.data, cBool(active))
}
