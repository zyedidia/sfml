package sfml

//#include <SFML/Window.h>
//#include "event.h"
import "C"

import "unicode/utf8"

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

func goVideoMode(vm *C.sfVideoMode) *VideoMode {
	govm := new(VideoMode)
	govm.Width = uint(vm.width)
	govm.Height = uint(vm.height)
	govm.BitsPerPixel = uint(vm.bitsPerPixel)
	return govm
}

func cVideoMode(vm *VideoMode) C.sfVideoMode {
	var cvm C.sfVideoMode
	cvm.width = C.uint(vm.Width)
	cvm.height = C.uint(vm.Height)
	cvm.bitsPerPixel = C.uint(vm.BitsPerPixel)
	return cvm
}

func goContextSettings(cs *C.sfContextSettings) *ContextSettings {
	r := new(ContextSettings)
	r.DepthBits = uint(cs.depthBits)
	r.StencilBits = uint(cs.stencilBits)
	r.AntialiasingLevel = uint(cs.antialiasingLevel)
	r.MajorVersion = uint(cs.majorVersion)
	r.MinorVersion = uint(cs.minorVersion)
	r.AttributeFlags = uint(cs.attributeFlags)
	return r
}

func cContextSettings(cs *ContextSettings) C.sfContextSettings {
	var r C.sfContextSettings
	r.depthBits = C.uint(cs.DepthBits)
	r.stencilBits = C.uint(cs.StencilBits)
	r.antialiasingLevel = C.uint(cs.AntialiasingLevel)
	r.majorVersion = C.uint(cs.MajorVersion)
	r.minorVersion = C.uint(cs.MinorVersion)
	r.attributeFlags = C.sfUint32(cs.AttributeFlags)
	return r
}

func goEvent(e *C.sfEvent) *Event {
	ev := C.GetSFMLEvent(e)
	r := new(Event)
	r.Type = EventType(ev.type_)
	r.Size.Width = uint(ev.size.width)
	r.Size.Height = uint(ev.size.height)
	r.Key.Code = KeyCode(ev.key.code)
	r.Key.Alt = goBool(ev.key.alt)
	r.Key.Control = goBool(ev.key.control)
	r.Key.Shift = goBool(ev.key.shift)
	r.Key.System = goBool(ev.key.system)
	r.Text.Unicode = rune(ev.text.unicode)
	r.MouseMove.X = int(ev.mouseMove.x)
	r.MouseMove.Y = int(ev.mouseMove.y)
	r.MouseButton.Button = MouseButton(ev.mouseButton.button)
	r.MouseButton.X = int(ev.mouseButton.x)
	r.MouseButton.Y = int(ev.mouseButton.y)
	r.MouseWheelScroll.Wheel = MouseWheel(ev.mouseWheelScroll.wheel)
	r.MouseWheelScroll.Delta = float32(ev.mouseWheelScroll.delta)
	r.MouseWheelScroll.X = int(ev.mouseWheelScroll.x)
	r.MouseWheelScroll.Y = int(ev.mouseWheelScroll.y)
	r.JoystickMove.JoystickId = uint(ev.joystickMove.joystickId)
	r.JoystickMove.Axis = JoystickAxis(ev.joystickMove.axis)
	r.JoystickMove.Position = float32(ev.joystickMove.position)
	r.JoystickButton.JoystickId = uint(ev.joystickButton.joystickId)
	r.JoystickButton.Button = uint(ev.joystickButton.button)
	r.JoystickConnect.JoystickId = uint(ev.joystickConnect.joystickId)
	r.Touch.Finger = uint(ev.touch.finger)
	r.Touch.X = int(ev.touch.x)
	r.Touch.Y = int(ev.touch.y)
	r.Sensor.Sensor = SensorType(ev.sensor.sensorType)
	r.Sensor.X = float32(ev.sensor.x)
	r.Sensor.Y = float32(ev.sensor.y)
	r.Sensor.Z = float32(ev.sensor.z)
	return r
}

func goVector2i(v *C.sfVector2i) *Vector2i {
	return &Vector2i{int(v.x), int(v.y)}
}

func cVector2i(v *Vector2i) C.sfVector2i {
	var r C.sfVector2i
	r.x = C.int(v.X)
	r.y = C.int(v.Y)
	return r
}

func goVector2u(v *C.sfVector2u) *Vector2u {
	return &Vector2u{uint(v.x), uint(v.y)}
}

func cVector2u(v *Vector2u) C.sfVector2u {
	var r C.sfVector2u
	r.x = C.uint(v.X)
	r.y = C.uint(v.Y)
	return r
}

func goVector2f(v *C.sfVector2f) *Vector2f {
	return &Vector2f{float32(v.x), float32(v.y)}
}

func cVector2f(v *Vector2f) C.sfVector2f {
	var r C.sfVector2f
	r.x = C.float(v.X)
	r.y = C.float(v.Y)
	return r
}

func goVector3f(v *C.sfVector3f) *Vector3f {
	return &Vector3f{float32(v.x), float32(v.y), float32(v.z)}
}

func cVector3f(v *Vector3f) C.sfVector3f {
	var r C.sfVector3f
	r.x = C.float(v.X)
	r.y = C.float(v.Y)
	r.z = C.float(v.Z)
	return r
}

func goJoystickIdentification(j *C.sfJoystickIdentification) *JoystickIdentification {
	r := new(JoystickIdentification)
	r.Name = C.GoString(j.name)
	r.VendorId = uint(j.vendorId)
	r.ProductId = uint(j.productId)
	return r
}
