package graphics

import "github.com/GomeBox/gome/primitives"

type Texture interface {
	TextureDrawer
	Dimensions() primitives.Dimensions
}
