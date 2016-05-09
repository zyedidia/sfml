package sfml

// Used for types RenderWindow and RenderTexture
type RenderTarget interface {
	Draw(object Drawable)
	DrawWithRenderStates(object Drawable, states *RenderStates)
}
