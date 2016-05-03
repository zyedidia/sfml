package sfml

//#include <SFML/Window.h>
import "C"

import "unicode/utf8"

func utf32(s string) *C.sfUint32 {
	s32 := make([]C.sfUint32, utf8.RuneCountInString(s)+1)
	for _, r := range s {
		s32 = append(s32, C.sfUint32(r))
	}
	// End with a NULL byte.
	s32 = append(s32, 0)
	return &s32[0]
}

func goVideoMode(vm C.sfVideoMode) *VideoMode {
	govm := new(VideoMode)
	govm.Width = uint(vm.width)
	govm.Height = uint(vm.height)
	govm.BitsPerPixel = uint(vm.bitsPerPixel)
	return govm
}

func cVideoMode(vm *VideoMode) C.sfVideoMode {
	var cvm C.sfVideoMode
	cvm.width = C.uint(vm.Width)
	cvm.height = C.uint(vm.Height)
	cvm.bitsPerPixel = C.uint(vm.BitsPerPixel)
	return cvm
}

func cBool(b bool) C.sfBool {
	if b {
		return C.sfTrue
	}
	return C.sfFalse
}

func goBool(b C.sfBool) bool {
	if b == C.sfTrue {
		return true
	}
	return false
}
