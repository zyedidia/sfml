package sfml

// Used for types RenderWindow and RenderTexture
type RenderTarget interface {
	GetSize() Vector2u
	SetView(view *View)
	GetView() *View
	GetDefaultView() *View
	GetViewport(view *View) Recti
	MapPixelToCoords(point Vector2i, view *View) Vector2f
	MapCoordsToPixel(point Vector2f, view *View) Vector2i
	PushGLStates()
	PopGLStates()
	ResetGLStates()
	Draw(object Drawable)
	DrawWithRenderStates(object Drawable, states *RenderStates)
	DrawPrimitives(object *VertexArray)
	DrawPrimitivesWithRenderStates(object *VertexArray, states *RenderStates)
}
