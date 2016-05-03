package sfml

//#include <SFML/Window.h>
//sfVideoMode GetVideoModeAtIndex(sfVideoMode *modes, int index);
import "C"

type VideoMode struct {
	Width        uint
	Height       uint
	BitsPerPixel uint
}

func GetDesktopMode() *VideoMode {
	vm := C.sfVideoMode_getDesktopMode()
	return goVideoMode(vm)
}

func GetFullscreenModes() []*VideoMode {
	var count C.size_t
	modes := C.sfVideoMode_getFullscreenModes(&count)
	vms := make([]*VideoMode, count)
	for i := 0; i < int(count); i++ {
		vms = append(vms, goVideoMode(C.GetVideoModeAtIndex(modes, C.int(i))))
	}
	return vms
}

func (v *VideoMode) IsValid() bool {
	return goBool(C.sfVideoMode_isValid(cVideoMode(v)))
}
