package sfml

//#include <SFML/Graphics.h>
import "C"

var (
	ColorBlack       = goColor(&C.sfBlack)
	ColorWhite       = goColor(&C.sfWhite)
	ColorRed         = goColor(&C.sfRed)
	ColorGreen       = goColor(&C.sfGreen)
	ColorBlue        = goColor(&C.sfBlue)
	ColorYellow      = goColor(&C.sfYellow)
	ColorMagenta     = goColor(&C.sfMagenta)
	ColorCyan        = goColor(&C.sfCyan)
	ColorTransparent = goColor(&C.sfTransparent)
)

type Color struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

func CreateColorFromInteger(color uint) Color {
	r := C.sfColor_fromInteger(C.sfUint32(color))
	return *goColor(&r)
}

func (c Color) Add(cc Color) Color {
	r := C.sfColor_add(cColor(&c), cColor(&cc))
	return *goColor(&r)
}

func (c Color) Subtract(cc Color) Color {
	r := C.sfColor_subtract(cColor(&c), cColor(&cc))
	return *goColor(&r)
}

func (c Color) Modulate(cc Color) Color {
	r := C.sfColor_modulate(cColor(&c), cColor(&cc))
	return *goColor(&r)
}

func (c Color) ToInteger() uint {
	return uint(C.sfColor_toInteger(cColor(&c)))
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
