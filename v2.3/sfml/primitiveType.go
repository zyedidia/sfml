package sfml

//#include <SFML/Graphics.h>
import "C"

const (
	PrimitivePoints         PrimitiveType = C.sfPoints
	PrimitiveLines          PrimitiveType = C.sfLines
	PrimitiveLinesStrip     PrimitiveType = C.sfLinesStrip
	PrimitiveTriangles      PrimitiveType = C.sfTriangles
	PrimitiveTrianglesStrip PrimitiveType = C.sfTrianglesStrip
	PrimitiveTrianglesFan   PrimitiveType = C.sfTrianglesFan
	PrimitiveQuads          PrimitiveType = C.sfQuads
)

type PrimitiveType int
