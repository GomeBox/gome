package graphics

import (
	adapters "github.com/GomeBox/gome/adapters/graphics"
	"github.com/GomeBox/gome/graphics"
	"github.com/GomeBox/gome/primitives"
)

func newTexture(tex adapters.Texture) graphics.Texture {
	texture := new(texture)
	texture.tex = tex
	return texture
}

type texture struct {
	tex graphics.Texture
}

func (texture *texture) Draw(source, dest *primitives.Rectangle) error {
	err := texture.tex.Draw(source, dest)
	if err != nil {
		return err
	}
	return nil
}

func (texture *texture) DrawF(source *primitives.Rectangle, dest *primitives.RectangleF) error {
	err := texture.tex.DrawF(source, dest)
	if err != nil {
		return err
	}
	return nil
}

func (texture *texture) Dimensions() primitives.Dimensions {
	return texture.tex.Dimensions()
}
