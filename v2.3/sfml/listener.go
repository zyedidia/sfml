package sfml

//#include <SFML/Audio.h>
import "C"

func SetGlobalVolume(volume float32) {
	C.sfListener_setGlobalVolume(C.float(volume))
}

func GetGlobalVolume() float32 {
	return float32(C.sfListener_getGlobalVolume())
}

func SetPosition(position Vector3f) {
	C.sfListener_setPosition(cVector3f(&position))
}

func GetPosition() Vector3f {
	p := C.sfListener_getPosition()
	return *goVector3f(&p)
}

func SetDirection(direction Vector3f) {
	C.sfListener_setDirection(cVector3f(&direction))
}

func GetDirection() Vector3f {
	p := C.sfListener_getDirection()
	return *goVector3f(&p)
}

func SetUpVector(upVector Vector3f) {
	C.sfListener_setUpVector(cVector3f(&upVector))
}

func GetUpVector() Vector3f {
	p := C.sfListener_getUpVector()
	return *goVector3f(&p)
}
