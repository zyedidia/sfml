package sfml

//#include <SFML/Window.h>
//#include "utils.h"
import "C"

import (
	"unsafe"
)

func goString(s *C.sfUint32) string {
	var str string
	for ptr := s; *ptr != 0; C.nextChar(&ptr) {
		str += string(rune(uint32(*ptr)))
	}
	return str
}

func cString(s string) *C.sfUint32 {
	utf32 := append([]rune(s), rune(0))
	return (*C.sfUint32)(unsafe.Pointer(&utf32[0]))
}

func goBool(b C.sfBool) bool {
	if b == C.sfTrue {
		return true
	}
	return false
}

func cBool(b bool) C.sfBool {
	if b {
		return C.sfTrue
	}
	return C.sfFalse
}
