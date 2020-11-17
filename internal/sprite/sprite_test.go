package sprite

import (
	"github.com/GomeBox/gome/internal/sprite/mocks"
	"github.com/GomeBox/gome/primitives"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	drawer := new(mocks.Drawer)
	pos := primitives.PointF{
		X: 100,
		Y: 222,
	}
	sprite := New(drawer, pos)
	assert.NotEmpty(t, sprite, "New returned nil")
	s := sprite.(*spriteImpl)
	assert.Same(t, drawer, s.drawer)
	assert.Equal(t, pos, s.pos)
}

func TestSpriteImpl_Dimensions(t *testing.T) {
	dimensions := primitives.Dimensions{
		Width:  129,
		Height: 234,
	}
	drawer := mocks.Drawer{
		OnDimensions: func() primitives.Dimensions {
			return dimensions
		},
	}
	sprite := spriteImpl{
		drawer: &drawer,
	}
	assert.Equal(t, dimensions, sprite.Dimensions())
}

func TestSpriteImpl_Position(t *testing.T) {
	pos := primitives.PointF{
		X: 123,
		Y: 345,
	}
	sprite := spriteImpl{
		pos: pos,
	}
	assert.Equal(t, pos, sprite.Position())
}

func TestSpriteImpl_SetPosition(t *testing.T) {
	pos := primitives.PointF{
		X: 123,
		Y: 345,
	}
	sprite := spriteImpl{
		pos: primitives.PointF{},
	}
	sprite.SetPosition(pos.X, pos.Y)
	assert.Equal(t, pos, sprite.pos)
}

func TestSpriteImpl_Draw(t *testing.T) {
	var drawPos primitives.PointF
	drawer := mocks.Drawer{
		OnDrawTo: func(pos primitives.PointF) error {
			drawPos = pos
			return nil
		},
	}
	pos := primitives.PointF{
		X: 123,
		Y: 345,
	}
	sprite := spriteImpl{
		drawer: &drawer,
		pos:    pos,
	}
	err := sprite.Draw()
	assert.NoError(t, err)
	assert.Equal(t, pos, drawPos)
}
