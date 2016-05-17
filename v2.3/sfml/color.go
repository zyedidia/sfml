package sfml

//#include <SFML/Graphics.h>
import "C"

var (
	ColorBlack       = Color{0, 0, 0, 255}
	ColorWhite       = Color{255, 255, 255, 255}
	ColorRed         = Color{255, 0, 0, 255}
	ColorGreen       = Color{0, 255, 0, 255}
	ColorBlue        = Color{0, 0, 255, 255}
	ColorYellow      = Color{255, 255, 0, 255}
	ColorMagenta     = Color{255, 0, 255, 255}
	ColorCyan        = Color{0, 255, 255, 255}
	ColorTransparent = Color{255, 255, 255, 0}
)

type Color struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

func NewColorFromInteger(color uint) Color {
	r := C.sfColor_fromInteger(C.sfUint32(color))
	return *goColor(&r)
}

func (c Color) Add(color Color) Color {
	r := C.sfColor_add(cColor(&c), cColor(&color))
	return *goColor(&r)
}

func (c Color) Subtract(color Color) Color {
	r := C.sfColor_subtract(cColor(&c), cColor(&color))
	return *goColor(&r)
}

func (c Color) Modulate(color Color) Color {
	r := C.sfColor_modulate(cColor(&c), cColor(&color))
	return *goColor(&r)
}

func (c Color) ToInteger() uint {
	return uint(C.sfColor_toInteger(cColor(&c)))
}

func (c Color) IsEqual(color Color) bool {
	return c.R == color.R && c.G == color.G && c.B == color.B && c.A == color.A
}

func goColor(c *C.sfColor) *Color {
	return &Color{
		uint8(c.r),
		uint8(c.g),
		uint8(c.b),
		uint8(c.a),
	}
}

func cColor(c *Color) C.sfColor {
	return C.sfColor{
		C.sfUint8(c.R),
		C.sfUint8(c.G),
		C.sfUint8(c.B),
		C.sfUint8(c.A),
	}
}
