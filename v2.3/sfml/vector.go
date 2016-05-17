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

func (v Vector2i) IsEqual(vector Vector2i) bool {
	return v.X == vector.X && v.Y == vector.Y
}

func (v Vector2i) Multiply(vector Vector2i) Vector2i {
	return Vector2i{
		v.X * vector.X,
		v.Y * vector.Y,
	}
}

func (v Vector2i) Divide(vector Vector2i) Vector2i {
	return Vector2i{
		v.X / vector.X,
		v.Y / vector.Y,
	}
}

func (v Vector2i) Add(vector Vector2i) Vector2i {
	return Vector2i{
		v.X + vector.X,
		v.Y + vector.Y,
	}
}

func (v Vector2i) Subtract(vector Vector2i) Vector2i {
	return Vector2i{
		v.X - vector.X,
		v.Y - vector.Y,
	}
}

func (v Vector2u) IsEqual(vector Vector2u) bool {
	return v.X == vector.X && v.Y == vector.Y
}

func (v Vector2u) Multiply(vector Vector2u) Vector2u {
	return Vector2u{
		v.X * vector.X,
		v.Y * vector.Y,
	}
}

func (v Vector2u) Divide(vector Vector2u) Vector2u {
	return Vector2u{
		v.X / vector.X,
		v.Y / vector.Y,
	}
}

func (v Vector2u) Add(vector Vector2u) Vector2u {
	return Vector2u{
		v.X + vector.X,
		v.Y + vector.Y,
	}
}

func (v Vector2u) Subtract(vector Vector2u) Vector2u {
	return Vector2u{
		v.X - vector.X,
		v.Y - vector.Y,
	}
}

func (v Vector2f) IsEqual(vector Vector2f) bool {
	return v.X == vector.X && v.Y == vector.Y
}

func (v Vector2f) Multiply(vector Vector2f) Vector2f {
	return Vector2f{
		v.X * vector.X,
		v.Y * vector.Y,
	}
}

func (v Vector2f) Divide(vector Vector2f) Vector2f {
	return Vector2f{
		v.X / vector.X,
		v.Y / vector.Y,
	}
}

func (v Vector2f) Add(vector Vector2f) Vector2f {
	return Vector2f{
		v.X + vector.X,
		v.Y + vector.Y,
	}
}

func (v Vector2f) Subtract(vector Vector2f) Vector2f {
	return Vector2f{
		v.X - vector.X,
		v.Y - vector.Y,
	}
}

func (v Vector3f) IsEqual(vector Vector3f) bool {
	return v.X == vector.X && v.Y == vector.Y && v.Z == vector.Z
}

func (v Vector3f) Multiply(vector Vector3f) Vector3f {
	return Vector3f{
		v.X * vector.X,
		v.Y * vector.Y,
		v.Z * vector.Z,
	}
}

func (v Vector3f) Divide(vector Vector3f) Vector3f {
	return Vector3f{
		v.X / vector.X,
		v.Y / vector.Y,
		v.Z / vector.Z,
	}
}

func (v Vector3f) Add(vector Vector3f) Vector3f {
	return Vector3f{
		v.X + vector.X,
		v.Y + vector.Y,
		v.Z + vector.Z,
	}
}

func (v Vector3f) Subtract(vector Vector3f) Vector3f {
	return Vector3f{
		v.X - vector.X,
		v.Y - vector.Y,
		v.Z - vector.Z,
	}
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
