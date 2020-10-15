package graphics

import (
	"github.com/GomeBox/gome/adapters/graphics"
	"github.com/GomeBox/gome/primitives"
)

type Texture interface {
	Draw(source, dest *primitives.Rectangle) error
	DrawF(source *primitives.Rectangle, dest *primitives.RectangleF) error
}

func NewTexture(drawer graphics.TextureDrawer) Texture {
	texture := new(texture)
	texture.drawer = drawer
	return texture
}

type texture struct {
	drawer graphics.TextureDrawer
}

func (texture *texture) Draw(source, dest *primitives.Rectangle) error {
	err := texture.drawer.Draw(source, dest)
	if err != nil {
		return err
	}
	return nil
}

func (texture *texture) DrawF(source *primitives.Rectangle, dest *primitives.RectangleF) error {
	err := texture.drawer.DrawF(source, dest)
	if err != nil {
		return err
	}
	return nil
}
