package graphics

import (
	"github.com/GomeBox/gome/adapters/shared"
	"github.com/GomeBox/gome/primitives"
)

type Texture interface {
	TextureDrawer
	shared.Unloader
	Dimensions() primitives.Dimensions
}
