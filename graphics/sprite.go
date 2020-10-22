package graphics

import (
	"errors"
	"github.com/GomeBox/gome/primitives"
)

type Sprite interface {
	Dimensions() primitives.Dimensions
	Position() primitives.PointF
	SetPosition(pos *primitives.PointF) error
	Draw() error
}

func NewTextureSprite(texture Texture, position primitives.PointF) Sprite {
	return &textureSprite{
		texture: texture,
		destRect: primitives.RectangleF{
			PointF: position,
			DimensionsF: primitives.DimensionsF{
				Width:  float32(texture.Dimensions().Width),
				Height: float32(texture.Dimensions().Height),
			},
		},
	}
}

type textureSprite struct {
	texture  Texture
	destRect primitives.RectangleF
}

func (t *textureSprite) Dimensions() primitives.Dimensions {
	return t.Dimensions()
}

func (t *textureSprite) Position() primitives.PointF {
	return primitives.PointF{
		X: t.destRect.X,
		Y: t.destRect.Y,
	}
}

func (t *textureSprite) SetPosition(pos *primitives.PointF) error {
	if pos == nil {
		return errors.New("argument pos may not be nil")
	}
	t.destRect.X = pos.X
	t.destRect.Y = pos.Y
	return nil
}

func (t *textureSprite) Draw() error {
	return t.texture.DrawF(nil, &t.destRect)
}
