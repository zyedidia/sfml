package sfml

//#include <SFML/Window.h>
import "C"

import (
	"unicode/utf8"
)

func goString(s *C.sfUint32) string {
	//st := []uint32(unsafe.Pointer(s))
	//return string(st)
	return "FIXME"
}

func cString(s string) *C.sfUint32 {
	s32 := make([]C.sfUint32, utf8.RuneCountInString(s)+1)
	for _, r := range s {
		s32 = append(s32, C.sfUint32(r))
	}
	s32 = append(s32, 0)
	return &s32[0]
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
