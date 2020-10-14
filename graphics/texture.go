package graphics

import "github.com/GomeBox/gome/adapters/graphics"

type Texture interface {
	Draw() error
}

func NewTexture(drawer graphics.TextureDrawer) Texture {
	texture := new(texture)
	texture.drawer = drawer
	return texture
}

type texture struct {
	drawer graphics.TextureDrawer
}

func (texture *texture) Draw() error {
	err := texture.drawer.Draw()
	if err != nil {
		return err
	}
	return nil
}
