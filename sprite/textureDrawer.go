package sprite

import (
	"github.com/GomeBox/gome/graphics"
	"github.com/GomeBox/gome/primitives"
)

type textureDrawer struct {
	texture    graphics.Texture
	dimensions primitives.Dimensions
	sourceRect *primitives.Rectangle
}

func newTextureDrawer(texture graphics.Texture) *textureDrawer {
	drawer := new(textureDrawer)
	drawer.texture = texture
	drawer.dimensions = primitives.Dimensions{
		Width:  texture.Dimensions().Width,
		Height: texture.Dimensions().Height,
	}
	return drawer
}

func (drawer *textureDrawer) DrawTo(pos *primitives.PointF) error {
	destRect := &primitives.RectangleF{
		PointF:      *pos,
		DimensionsF: drawer.dimensions.ToDimensionsF()}
	return drawer.texture.DrawF(drawer.sourceRect, destRect)
}

func (drawer *textureDrawer) Dimensions() primitives.Dimensions {
	return drawer.dimensions
}
