package sfml

//#include <SFML/Graphics.h>
import "C"

type Texture struct {
	data *C.sfTexture
}
