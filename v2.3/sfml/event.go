package sfml

//#include <SFML/Window.h>
//#include "event.h"
import "C"

const (
	EventClosed                 EventType = C.sfEvtClosed
	EventResized                EventType = C.sfEvtResized
	EventLostFocus              EventType = C.sfEvtLostFocus
	EventGainedFocus            EventType = C.sfEvtGainedFocus
	EventTextEntered            EventType = C.sfEvtTextEntered
	EventKeyPressed             EventType = C.sfEvtKeyPressed
	EventKeyReleased            EventType = C.sfEvtKeyReleased
	EventMouseWheelMoved        EventType = C.sfEvtMouseWheelMoved
	EventMouseWheelScrolled     EventType = C.sfEvtMouseWheelScrolled
	EventMouseButtonPressed     EventType = C.sfEvtMouseButtonPressed
	EventMouseButtonReleased    EventType = C.sfEvtMouseButtonReleased
	EventMouseMoved             EventType = C.sfEvtMouseMoved
	EventMouseEntered           EventType = C.sfEvtMouseEntered
	EventMouseLeft              EventType = C.sfEvtMouseLeft
	EventJoystickButtonPressed  EventType = C.sfEvtJoystickButtonPressed
	EventJoystickButtonReleased EventType = C.sfEvtJoystickButtonReleased
	EventJoystickMoved          EventType = C.sfEvtJoystickMoved
	EventJoystickConnected      EventType = C.sfEvtJoystickConnected
	EventJoystickDisconnected   EventType = C.sfEvtJoystickDisconnected
	EventTouchBegan             EventType = C.sfEvtTouchBegan
	EventTouchMoved             EventType = C.sfEvtTouchMoved
	EventTouchEnded             EventType = C.sfEvtTouchEnded
	EventSensorChanged          EventType = C.sfEvtSensorChanged
	EventCount                  EventType = C.sfEvtCount
)

type EventType int

type KeyEvent struct {
	Code    KeyCode
	Alt     bool
	Control bool
	Shift   bool
	System  bool
}

type TextEvent struct {
	Unicode rune
}

type MouseMoveEvent struct {
	X int
	Y int
}

type MouseButtonEvent struct {
	Button MouseButton
	X      int
	Y      int
}

type MouseWheelScrollEvent struct {
	Wheel MouseWheel
	Delta float32
	X     int
	Y     int
}

type JoystickMoveEvent struct {
	JoystickId uint
	Axis       JoystickAxis
	Position   float32
}

type JoystickButtonEvent struct {
	JoystickId uint
	Button     uint
}

type JoystickConnectEvent struct {
	JoystickId uint
}

type SizeEvent struct {
	Width  uint
	Height uint
}

type TouchEvent struct {
	Finger uint
	X      int
	Y      int
}

type SensorEvent struct {
	Sensor SensorType
	X      float32
	Y      float32
	Z      float32
}

type Event struct {
	Type             EventType
	Size             SizeEvent
	Key              KeyEvent
	Text             TextEvent
	MouseMove        MouseMoveEvent
	MouseButton      MouseButtonEvent
	MouseWheelScroll MouseWheelScrollEvent
	JoystickMove     JoystickMoveEvent
	JoystickButton   JoystickButtonEvent
	JoystickConnect  JoystickConnectEvent
	Touch            TouchEvent
	Sensor           SensorEvent
}

func goEvent(event *C.sfEvent) *Event {
	ev := C.GetSFMLEvent(event)
	e := new(Event)
	e.Type = EventType(ev.type_)
	switch e.Type {
	case EventResized:
		e.Size = SizeEvent{
			uint(ev.size.width),
			uint(ev.size.height),
		}
	case EventKeyPressed, EventKeyReleased:
		e.Key = KeyEvent{
			KeyCode(ev.key.code),
			goBool(ev.key.alt),
			goBool(ev.key.control),
			goBool(ev.key.shift),
			goBool(ev.key.system),
		}
	case EventTextEntered:
		e.Text = TextEvent{
			rune(ev.text.unicode),
		}
	case EventMouseMoved:
		e.MouseMove = MouseMoveEvent{
			int(ev.mouseMove.x),
			int(ev.mouseMove.y),
		}
	case EventMouseButtonPressed, EventMouseButtonReleased:
		e.MouseButton = MouseButtonEvent{
			MouseButton(ev.mouseButton.button),
			int(ev.mouseButton.x),
			int(ev.mouseButton.y),
		}
	case EventMouseWheelScrolled:
		e.MouseWheelScroll = MouseWheelScrollEvent{
			MouseWheel(ev.mouseWheelScroll.wheel),
			float32(ev.mouseWheelScroll.delta),
			int(ev.mouseWheelScroll.x),
			int(ev.mouseWheelScroll.y),
		}
	case EventJoystickMoved:
		e.JoystickMove = JoystickMoveEvent{
			uint(ev.joystickMove.joystickId),
			JoystickAxis(ev.joystickMove.axis),
			float32(ev.joystickMove.position),
		}
	case EventJoystickButtonPressed, EventJoystickButtonReleased:
		e.JoystickButton = JoystickButtonEvent{
			uint(ev.joystickButton.joystickId),
			uint(ev.joystickButton.button),
		}
	case EventJoystickConnected, EventJoystickDisconnected:
		e.JoystickConnect = JoystickConnectEvent{
			uint(ev.joystickConnect.joystickId),
		}
	case EventTouchBegan, EventTouchMoved, EventTouchEnded:
		e.Touch = TouchEvent{
			uint(ev.touch.finger),
			int(ev.touch.x),
			int(ev.touch.y),
		}
	case EventSensorChanged:
		e.Sensor = SensorEvent{
			SensorType(ev.sensor.sensorType),
			float32(ev.sensor.x),
			float32(ev.sensor.y),
			float32(ev.sensor.z),
		}
	}
	return e
}
