package sfml

//#include <SFML/Window.h>
//#include "utils.h"
import "C"

import (
	//"unicode/utf8"
	"unsafe"
)

func goString(s *C.sfUint32) string {
	//st := []uint32(unsafe.Pointer(s))
	//return string(st)
	return "FIXME"
}

func strToRunes(str string) []rune {
	return append([]rune(str), rune(0))
}

func cString(s string) *C.sfUint32 {
	/*length := utf8.RuneCountInString(s)
	cstr := make([]uint32, length)
	for _, v := range(s) {
		cstr = append(cstr, uint32(v))
	}
	cstr[length] = 0
	return (*C.sfUint32)(unsafe.Pointer(&cstr[0]))*/
	utf32 := strToRunes(s)
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
