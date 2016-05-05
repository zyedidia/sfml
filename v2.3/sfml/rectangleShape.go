package sfml

//#include <SFML/Graphics.h>
import "C"

type RectangleShape struct {
	data *C.sfRectangleShape
}
