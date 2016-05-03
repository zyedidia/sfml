package sfml

//#include <SFML/Window.h>
import "C"

// Window styles
type WindowStyle uint

const (
	None        WindowStyle = C.sfNone
	Titlebar                = C.sfTitlebar
	Resize                  = C.sfResize
	Close                   = C.sfClose
	Fullscreen              = C.sfFullscreen
	DefaultStye             = C.sfDefaultStyle
)

// Context attributes
const (
	ContextDefault uint = C.sfContextDefault
	ContextCore         = C.sfContextCore
	ContextDebug        = C.sfContextDebug
)

// Context settings
type ContextSettings struct {
	DepthBits         uint
	StencilBits       uint
	AntialiasingLevel uint
	MajorVersion      uint
	MinorVersion      uint
	AttributeFlags    uint
}

type Window struct {
	data *C.sfWindow
}

func CreateWindow(mode *VideoMode, title string, style WindowStyle, settings *ContextSettings) *Window {
	var cs C.sfContextSettings
	cs.depthBits = C.uint(settings.DepthBits)
	cs.stencilBits = C.uint(settings.StencilBits)
	cs.antialiasingLevel = C.uint(settings.AntialiasingLevel)
	cs.majorVersion = C.uint(settings.MajorVersion)
	cs.minorVersion = C.uint(settings.MinorVersion)
	cs.attributeFlags = C.sfUint32(settings.AttributeFlags)
	w := C.sfWindow_createUnicode(CVideoMode(mode), utf32(title), C.sfUint32(style), &cs)
	return &Window{w}
}
