package sfml

//#include <SFML/Window.h>
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

func MouseGetPosition(relativeTo *Window) Vector2i {
	v := C.sfMouse_getPosition(relativeTo.data)
	return *goVector2i(&v)
}

func MouseSetPosition(position Vector2i, relativeTo *Window) {
	C.sfMouse_setPosition(cVector2i(&position), relativeTo.data)
}
