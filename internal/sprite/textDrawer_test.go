package sprite

import (
	"github.com/GomeBox/gome/internal/graphics/mocks"
	"github.com/GomeBox/gome/primitives"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTextDrawer_DrawTo(t *testing.T) {
	pos := primitives.PointF{
		X: 123.23,
		Y: 234.34,
	}
	var drawPos primitives.PointF
	text := mocks.Text{
		OnDraw: func(f primitives.PointF) error {
			drawPos = f
			return nil
		},
	}
	drawer := textDrawer{text: text}
	err := drawer.DrawTo(pos)
	assert.NoError(t, err)
	assert.Equal(t, pos, drawPos)
}

func TestTextDrawer_Dimensions(t *testing.T) {
	dimensions := primitives.Dimensions{
		Width:  100,
		Height: 390,
	}
	text := mocks.Text{
		OnDimensions: func() primitives.Dimensions {
			return dimensions
		},
	}
	drawer := textDrawer{text: text}
	assert.Equal(t, dimensions, drawer.Dimensions())
}
