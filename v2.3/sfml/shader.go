package sfml

//#include <SFML/Graphics.h>
import "C"

type Shader struct {
	data *C.sfShader
}

func CreateShaderFromFile(vertexShaderFilename, fragmentShaderFilename string) *Shader {
	return nil
}
