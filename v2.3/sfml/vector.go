package sfml

//#include <SFML/Window.h>
import "C"

type Vector2i struct {
	X int
	Y int
}

type Vector2u struct {
	X uint
	Y uint
}

type Vector2f struct {
	X float32
	Y float32
}

type Vector3f struct {
	X float32
	Y float32
	Z float32
}

func goVector2i(v *C.sfVector2i) *Vector2i {
	return &Vector2i{
		int(v.x),
		int(v.y),
	}
}

func cVector2i(v *Vector2i) C.sfVector2i {
	return C.sfVector2i{
		C.int(v.X),
		C.int(v.Y),
	}
}

func goVector2u(v *C.sfVector2u) *Vector2u {
	return &Vector2u{
		uint(v.x),
		uint(v.y),
	}
}

func cVector2u(v *Vector2u) C.sfVector2u {
	return C.sfVector2u{
		C.uint(v.X),
		C.uint(v.Y),
	}
}

func goVector2f(v *C.sfVector2f) *Vector2f {
	return &Vector2f{
		float32(v.x),
		float32(v.y),
	}
}

func cVector2f(v *Vector2f) C.sfVector2f {
	return C.sfVector2f{
		C.float(v.X),
		C.float(v.Y),
	}
}

func goVector3f(v *C.sfVector3f) *Vector3f {
	return &Vector3f{
		float32(v.x),
		float32(v.y),
		float32(v.z),
	}
}

func cVector3f(v *Vector3f) C.sfVector3f {
	return C.sfVector3f{
		C.float(v.X),
		C.float(v.Y),
		C.float(v.Z),
	}
}
