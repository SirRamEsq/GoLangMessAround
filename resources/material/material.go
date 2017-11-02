package material

import (
	"lengine/resources/shader"
	"lengine/resources/texture"
)

type Material struct {
	Diffuse texture.Texture
	Shader  shader.Shader
}
