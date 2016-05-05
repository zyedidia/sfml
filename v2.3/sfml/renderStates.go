package sfml

//#include <SFML/Graphics.h>
import "C"

type RenderStates struct {
	BlendMode BlendMode
	Transform Transform
	Texture   *Texture
	Shader    *Shader
}

func CreateDefaultRenderStates() *RenderStates {
	return &RenderStates{
		BlendAlpha,
		TransformIdentity,
		nil,
		nil,
	}
}

func cRenderStates(r *RenderStates) C.sfRenderStates {
	return C.sfRenderStates{
		cBlendMode(&r.BlendMode),
		cTransform(&r.Transform),
		r.Texture.data,
		r.Shader.data,
	}
}
