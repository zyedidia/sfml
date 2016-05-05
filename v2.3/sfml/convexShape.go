package sfml

//#include <SFML/Graphics.h>
import "C"

type ConvexShape struct {
	data *C.sfConvexShape
}
