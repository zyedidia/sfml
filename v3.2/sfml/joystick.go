package sfml

//#include <SFML/Window.h>
import "C"

const (
	JoystickCount       = C.sfJoystickCount
	JoystickButtonCount = C.sfJoystickButtonCount
	JoystickAxisCount   = C.sfJoystickAxisCount
)

const (
	JoystickX    JoystickAxis = C.sfJoystickX
	JoystickY    JoystickAxis = C.sfJoystickY
	JoystickZ    JoystickAxis = C.sfJoystickZ
	JoystickR    JoystickAxis = C.sfJoystickR
	JoystickU    JoystickAxis = C.sfJoystickU
	JoystickV    JoystickAxis = C.sfJoystickV
	JoystickPovX JoystickAxis = C.sfJoystickPovX
	JoystickPovY JoystickAxis = C.sfJoystickPovY
)

type JoystickAxis int

type JoystickIdentification struct {
	Name      string
	VendorId  uint
	ProductId uint
}

func JoystickIsConnected(joystick uint) bool {
	return goBool(C.sfJoystick_isConnected(C.uint(joystick)))
}

func JoystickGetButtonCount(joystick uint) uint {
	return uint(C.sfJoystick_getButtonCount(C.uint(joystick)))
}

func JoystickHasAxis(joystick uint, axis JoystickAxis) bool {
	return goBool(C.sfJoystick_hasAxis(C.uint(joystick), C.sfJoystickAxis(axis)))
}

func JoystickIsButtonPressed(joystick uint, button uint) bool {
	return goBool(C.sfJoystick_isButtonPressed(C.uint(joystick), C.uint(button)))
}

func JoystickGetAxisPosition(joystick uint, axis JoystickAxis) float32 {
	return float32(C.sfJoystick_getAxisPosition(C.uint(joystick), C.sfJoystickAxis(axis)))
}

func JoystickGetIdentification(joystick uint) *JoystickIdentification {
	j := C.sfJoystick_getIdentification(C.uint(joystick))
	return goJoystickIdentification(&j)
}

func JoystickUpdate() {
	C.sfJoystick_update()
}
