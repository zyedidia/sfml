package sfml

//#include <SFML/Window.h>
//#include "videoMode.h"
import "C"

type VideoMode struct {
	Width        uint
	Height       uint
	BitsPerPixel uint
}

func GetDesktopMode() *VideoMode {
	v := C.sfVideoMode_getDesktopMode()
	return goVideoMode(&v)
}

func GetFullscreenModes() []*VideoMode {
	var count C.size_t
	modes := C.sfVideoMode_getFullscreenModes(&count)
	vms := make([]*VideoMode, count)
	for i := 0; i < int(count); i++ {
		v := C.GetVideoModeAtIndex(modes, C.int(i))
		vms = append(vms, goVideoMode(&v))
	}
	return vms
}

func (v *VideoMode) IsValid() bool {
	return goBool(C.sfVideoMode_isValid(cVideoMode(v)))
}

func goVideoMode(vm *C.sfVideoMode) *VideoMode {
	return &VideoMode{
		uint(vm.width),
		uint(vm.height),
		uint(vm.bitsPerPixel),
	}
}

func cVideoMode(vm *VideoMode) C.sfVideoMode {
	return C.sfVideoMode{
		C.uint(vm.Width),
		C.uint(vm.Height),
		C.uint(vm.BitsPerPixel),
	}
}
