package sfml

//#include <SFML/Graphics.h>
import "C"

type Sprite struct {
	data *C.sfSprite
}
