package sprite

import (
	"github.com/GomeBox/gome/primitives"
)

type drawer interface {
	DrawTo(pos *primitives.PointF) error
	Dimensions() primitives.Dimensions
}

type sprite struct {
	drawer   drawer
	position primitives.PointF
}

func (sprite *sprite) Dimensions() primitives.Dimensions {
	return sprite.Dimensions()
}

func (sprite *sprite) Position() primitives.PointF {
	return sprite.position
}

func (sprite *sprite) SetPosition(x, y float32) error {
	sprite.position.X = x
	sprite.position.Y = y
	return nil
}

func (sprite *sprite) Draw() error {
	return sprite.drawer.DrawTo(&sprite.position)
}
