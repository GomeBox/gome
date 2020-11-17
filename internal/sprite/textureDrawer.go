package sprite

import (
	gomeInterfaces "github.com/GomeBox/gome/interfaces"
	"github.com/GomeBox/gome/internal/sprite/interfaces"
	"github.com/GomeBox/gome/primitives"
)

type textureDrawer struct {
	texture    gomeInterfaces.Texture
	dimensions *primitives.Dimensions
	sourceRect *primitives.Rectangle
}

func NewTextureDrawer(texture gomeInterfaces.Texture, sourceRect *primitives.Rectangle) interfaces.Drawer {
	drawer := new(textureDrawer)
	drawer.texture = texture
	drawer.dimensions = &primitives.Dimensions{
		Width:  texture.Dimensions().Width,
		Height: texture.Dimensions().Height,
	}
	drawer.sourceRect = sourceRect
	return drawer
}

func (drawer *textureDrawer) DrawTo(pos primitives.PointF) error {
	destRect := primitives.RectangleF{
		PointF:      pos,
		DimensionsF: drawer.dimensions.ToDimensionsF()}
	return drawer.texture.Draw(drawer.sourceRect, destRect)
}

func (drawer *textureDrawer) Dimensions() primitives.Dimensions {
	return primitives.Dimensions{
		Width:  drawer.dimensions.Width,
		Height: drawer.dimensions.Height,
	}
}
