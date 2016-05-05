package sfml

//#include <SFML/Graphics.h>
import "C"

type Shape struct {
	data *C.sfShape
}
