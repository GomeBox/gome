package sprite

import (
	"github.com/GomeBox/gome/primitives"
)

func New(drawer drawer, position primitives.PointF) *Implementation {
	s := new(Implementation)
	s.drawer = drawer
	s.pos = position
	return s
}

type Implementation struct {
	drawer drawer
	pos    primitives.PointF
}

func (sprite *Implementation) Dimensions() primitives.Dimensions {
	return sprite.Dimensions()
}

func (sprite *Implementation) Position() primitives.PointF {
	return sprite.pos
}

func (sprite *Implementation) SetPosition(x, y float32) error {
	sprite.pos.X = x
	sprite.pos.Y = y
	return nil
}

func (sprite *Implementation) Draw() error {
	return sprite.drawer.DrawTo(&sprite.pos)
}
