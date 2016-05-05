package sfml

//#include <SFML/Graphics.h>
import "C"

type Vertex struct {
	Position  Vector2f
	Color     Color
	TexCoords Vector2f
}

func goVertex(v *C.sfVertex) *Vertex {
	return &Vertex{
		*goVector2f(&v.position),
		*goColor(&v.color),
		*goVector2f(&v.texCoords),
	}
}

func cVertex(v *Vertex) C.sfVertex {
	return C.sfVertex{
		cVector2f(&v.Position),
		cColor(&v.Color),
		cVector2f(&v.TexCoords),
	}
}
