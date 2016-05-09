package sfml

// Used for types Sprite, Text, Shape, CircleShape, ConvexShape, RectangleShape, VertexArray
type Drawable interface {
	Draw(target RenderTarget)
	DrawWithRenderStates(target RenderTarget, states *RenderStates)
}
