package sfml

//#include <SFML/Window.h>
//#include <SFML/Graphics.h>
import "C"

type MouseButton int
type MouseWheel int

const (
	MouseLeft        MouseButton = C.sfMouseLeft
	MouseRight       MouseButton = C.sfMouseRight
	MouseMiddle      MouseButton = C.sfMouseMiddle
	MouseXButton1    MouseButton = C.sfMouseXButton1
	MouseXButton2    MouseButton = C.sfMouseXButton2
	MouseButtonCount MouseButton = C.sfMouseButtonCount
)

const (
	MouseVerticalWheel   MouseWheel = C.sfMouseVerticalWheel
	MouseHorizontalWheel MouseWheel = C.sfMouseHorizontalWheel
)

func IsMouseButtonPressed(button MouseButton) bool {
	return goBool(C.sfMouse_isButtonPressed(C.sfMouseButton(button)))
}

func MouseGetPosition(relativeTo SystemWindow) Vector2i {
	var v C.sfVector2i
	switch relativeTo.(type) {
	case *RenderWindow:
		v = C.sfMouse_getPositionRenderWindow(relativeTo.(*RenderWindow).data)
	case *Window:
		v = C.sfMouse_getPosition(relativeTo.(*Window).data)
	}
	return *goVector2i(&v)
}

func MouseSetPosition(position Vector2i, relativeTo SystemWindow) {
	switch relativeTo.(type) {
	case *RenderWindow:
		C.sfMouse_setPositionRenderWindow(cVector2i(&position), relativeTo.(*RenderWindow).data)
	case *Window:
		C.sfMouse_setPosition(cVector2i(&position), relativeTo.(*Window).data)
	}
}
