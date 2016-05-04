package sfml

//#include <SFML/Window.h>
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
	Width  uint //< New width, in pixels
	Height uint //< New height, in pixels
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
