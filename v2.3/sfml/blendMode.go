package sfml

//#include <SFML/Graphics.h>
import "C"

const (
	FactorZero             BlendFactor = C.sfBlendFactorZero
	FactorOne              BlendFactor = C.sfBlendFactorOne
	FactorSrcColor         BlendFactor = C.sfBlendFactorSrcColor
	FactorOneMinusSrcColor BlendFactor = C.sfBlendFactorOneMinusSrcColor
	FactorDstColor         BlendFactor = C.sfBlendFactorDstColor
	FactorOneMinusDstColor BlendFactor = C.sfBlendFactorOneMinusDstColor
	FactorSrcAlpha         BlendFactor = C.sfBlendFactorSrcAlpha
	FactorOneMinusSrcAlpha BlendFactor = C.sfBlendFactorOneMinusSrcAlpha
	FactorDstAlpha         BlendFactor = C.sfBlendFactorDstAlpha
	FactorOneMinusDstAlpha BlendFactor = C.sfBlendFactorOneMinusDstAlpha
)

const (
	EquationAdd      BlendEquation = C.sfBlendEquationAdd
	EquationSubtract BlendEquation = C.sfBlendEquationSubtract
)

var (
	BlendAlpha = BlendMode{
		FactorSrcAlpha,
		FactorOneMinusSrcAlpha,
		EquationAdd,
		FactorOne,
		FactorOneMinusSrcAlpha,
		EquationAdd,
	}
	BlendAdd = BlendMode{
		FactorSrcAlpha,
		FactorOne,
		EquationAdd,
		FactorOne,
		FactorOne,
		EquationAdd,
	}
	BlendMultiply = BlendMode{
		FactorDstColor,
		FactorZero,
		EquationAdd,
		FactorDstColor,
		FactorZero,
		EquationAdd,
	}
	BlendNone = BlendMode{
		FactorOne,
		FactorZero,
		EquationAdd,
		FactorOne,
		FactorZero,
		EquationAdd,
	}
)

type BlendFactor int
type BlendEquation int

type BlendMode struct {
	ColorSrcFactor BlendFactor
	ColorDstFactor BlendFactor
	ColorEquation  BlendEquation
	AlphaSrcFactor BlendFactor
	AlphaDstFactor BlendFactor
	AlphaEquation  BlendEquation
}

func goBlendMode(b *C.sfBlendMode) *BlendMode {
	return &BlendMode{
		BlendFactor(b.colorSrcFactor),
		BlendFactor(b.colorDstFactor),
		BlendEquation(b.colorEquation),
		BlendFactor(b.alphaSrcFactor),
		BlendFactor(b.alphaDstFactor),
		BlendEquation(b.alphaEquation),
	}
}

func cBlendMode(b *BlendMode) C.sfBlendMode {
	return C.sfBlendMode{
		C.sfBlendFactor(b.ColorSrcFactor),
		C.sfBlendFactor(b.ColorDstFactor),
		C.sfBlendEquation(b.ColorEquation),
		C.sfBlendFactor(b.AlphaSrcFactor),
		C.sfBlendFactor(b.AlphaDstFactor),
		C.sfBlendEquation(b.AlphaEquation),
	}
}
