package sfml

//#include <SFML/Graphics.h>
import "C"

type CircleShape struct {
	data *C.sfCircleShape
}
