package sfml

//#include <SFML/Graphics.h>
import "C"

type Glyph struct {
	Advance     float32
	Bounds      Rectf
	TextureRect Recti
}

func goGlyph(v *C.sfGlyph) *Glyph {
	return &Glyph{
		float32(v.advance),
		*goRectf(&v.bounds),
		*goRecti(&v.textureRect),
	}
}

func cGlyph(v *Glyph) C.sfGlyph {
	return C.sfGlyph{
		C.float(v.Advance),
		cRectf(&v.Bounds),
		cRecti(&v.TextureRect),
	}
}
