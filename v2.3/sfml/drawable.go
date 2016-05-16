package sfml

// Used for types Sprite, Text, Shape, CircleShape, ConvexShape, RectangleShape, VertexArray
type Drawable interface {
	SetPosition(position Vector2f)
	SetRotation(angle float32)
	SetScale(scale Vector2f)
	SetOrigin(origin Vector2f)
	GetPosition() Vector2f
	GetRotation() float32
	GetScale() Vector2f
	GetOrigin() Vector2f
	Move(offset Vector2f)
	Rotate(angle float32)
	Scale(factors Vector2f)
	GetTransform() Transform
	GetInverseTransform() Transform
	Draw(target RenderTarget)
	DrawWithRenderStates(target RenderTarget, states *RenderStates)
}
