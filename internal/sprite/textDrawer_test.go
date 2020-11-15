package sprite

import (
	"github.com/GomeBox/gome/internal/graphics"
	"github.com/GomeBox/gome/primitives"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTextDrawer(t *testing.T) {
	//text := new(mocks.TextMock)
	//drawer := NewTextDrawer(text)
	//d := drawer.(*textDrawer)
	//assert.Same(t, text, d.text)
}

func TestTextDrawer_DrawTo(t *testing.T) {
	pos := &primitives.PointF{
		X: 123.23,
		Y: 234.34,
	}
	var drawPos *primitives.PointF
	text := graphics.TextMock{
		OnDrawF: func(f *primitives.PointF) error {
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
	text := graphics.TextMock{
		OnDimensions: func() primitives.Dimensions {
			return dimensions
		},
	}
	drawer := textDrawer{text: text}
	assert.Equal(t, dimensions, drawer.Dimensions())
}
