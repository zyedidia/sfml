#include <SFML/Window/Event.h>

typedef struct eventWrapper {
	sfEventType             type_;
	sfSizeEvent             *size;
	sfKeyEvent              *key;
	sfTextEvent             *text;
	sfMouseMoveEvent        *mouseMove;
	sfMouseButtonEvent      *mouseButton;
	sfMouseWheelScrollEvent *mouseWheelScroll;
	sfJoystickMoveEvent     *joystickMove;
	sfJoystickButtonEvent   *joystickButton;
	sfJoystickConnectEvent  *joystickConnect;
	sfTouchEvent            *touch;
	sfSensorEvent           *sensor;
} eventWrapper;

eventWrapper GetSFMLEvent(sfEvent *e) {
	eventWrapper ev = {
		e->type,
		&(e->size),
		&(e->key),
		&(e->text),
		&(e->mouseMove),
		&(e->mouseButton),
		&(e->mouseWheelScroll),
		&(e->joystickMove),
		&(e->joystickButton),
		&(e->joystickConnect),
		&(e->touch),
		&(e->sensor)
	};
	return ev;
}
