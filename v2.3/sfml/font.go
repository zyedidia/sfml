package sfml

//#include <SFML/Graphics.h>
import "C"

type Font struct {
	data *C.sfFont
}
