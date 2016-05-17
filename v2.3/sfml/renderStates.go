package sfml

//#include <SFML/Graphics.h>
import "C"

type RenderStates struct {
	BlendMode BlendMode
	Transform Transform
	Texture   *Texture
	Shader    *Shader
}

var DefaultRenderState = &RenderStates{
	BlendAlpha,
	TransformIdentity,
	nil,
	nil,
}

func cRenderStates(r *RenderStates) C.sfRenderStates {
	return C.sfRenderStates{
		blendMode: cBlendMode(&r.BlendMode),
		transform: cTransform(&r.Transform),
		texture:   r.Texture.data,
		shader:    r.Shader.data,
	}
}
