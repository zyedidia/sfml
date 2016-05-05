package sfml

//#include <SFML/Window.h>
//#include <SFML/Graphics.h>
import "C"

func TouchIsDown(finger uint) bool {
	return goBool(C.sfTouch_isDown(C.uint(finger)))
}

func TouchGetPosition(finger uint, relativeTo SystemWindow) Vector2i {
	var v C.sfVector2i
	if relativeTo == nil {
		v = C.sfTouch_getPosition(C.uint(finger), nil)
	} else {
		switch relativeTo.(type) {
		case *RenderWindow:
			v = C.sfTouch_getPositionRenderWindow(C.uint(finger), relativeTo.(*RenderWindow).data)
		case *Window:
			v = C.sfTouch_getPosition(C.uint(finger), relativeTo.(*Window).data)
		}
	}
	return *goVector2i(&v)
}
