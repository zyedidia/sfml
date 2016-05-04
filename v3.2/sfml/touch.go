package sfml

//#include <SFML/Window.h>
import "C"

func TouchIsDown(finger uint) bool {
	return goBool(C.sfTouch_isDown(C.uint(finger)))
}

func TouchGetPosition(finger uint, relativeTo *Window) Vector2i {
	v := C.sfTouch_getPosition(C.uint(finger), relativeTo.data)
	return *goVector2i(&v)
}
