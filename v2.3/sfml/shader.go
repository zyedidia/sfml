package sfml

//#include <SFML/Graphics.h>
import "C"

type Shader struct {
	data *C.sfShader
}
