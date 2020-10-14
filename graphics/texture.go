package graphics

import (
	"github.com/GomeBox/gome/adapters/graphics"
	"github.com/GomeBox/gome/primitives"
)

type Texture interface {
	Draw(source, dest *primitives.Rectangle) error
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
