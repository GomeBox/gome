package sprite

import (
	"github.com/GomeBox/gome/internal/sprite/interfaces"
	"github.com/GomeBox/gome/primitives"
)

func New(drawer drawer, position primitives.PointF) interfaces.Sprite {
	s := new(spriteImpl)
	s.drawer = drawer
	s.pos = position
	return s
}

type spriteImpl struct {
	drawer drawer
	pos    primitives.PointF
}

func (sprite *spriteImpl) Dimensions() primitives.Dimensions {
	return sprite.drawer.Dimensions()
}

func (sprite *spriteImpl) Position() primitives.PointF {
	return sprite.pos
}

func (sprite *spriteImpl) SetPosition(x, y float32) {
	sprite.pos.X = x
	sprite.pos.Y = y
}

func (sprite *spriteImpl) Draw() error {
	return sprite.drawer.DrawTo(sprite.pos)
}
