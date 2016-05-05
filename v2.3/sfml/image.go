package sfml

//#include <SFML/Graphics.h>
import "C"

type Image struct {
	data *C.sfImage
}
