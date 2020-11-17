package sprite

import (
	"github.com/GomeBox/gome/internal/graphics/mocks"
	"github.com/GomeBox/gome/primitives"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTextureDrawer(t *testing.T) {
	sourceRect := primitives.NewRectangle(1, 2, 3, 4)
	dimensions := primitives.Dimensions{
		Width:  123,
		Height: 234,
	}
	texture := &mocks.Texture{
		OnDimensions: func() primitives.Dimensions {
			return dimensions
		},
	}
	drawer := NewTextureDrawer(texture, &sourceRect)
	if !assert.NotEmpty(t, drawer, "NewTextureDrawer returned nil") {
		return
	}
	d := drawer.(*textureDrawer)
	assert.Same(t, texture, d.texture)
	assert.Equal(t, &dimensions, d.dimensions, "dimensions do not match")
	assert.Equal(t, &sourceRect, d.sourceRect, "sourceRect does not match")
}

func TestTextureDrawer_DrawTo(t *testing.T) {
	sourceRect := primitives.NewRectangle(1, 2, 3, 4)
	dimensions := primitives.Dimensions{
		Width:  123,
		Height: 234,
	}
	pos := primitives.PointF{
		X: 234.5,
		Y: 145.34,
	}
	destRect := primitives.RectangleF{
		PointF:      pos,
		DimensionsF: dimensions.ToDimensionsF()}
	var drawSourceRect *primitives.Rectangle
	var drawDestRect primitives.RectangleF
	texture := &mocks.Texture{
		OnDraw: func(source *primitives.Rectangle, dest primitives.RectangleF) error {
			drawSourceRect = source
			drawDestRect = dest
			return nil
		},
	}
	drawer := textureDrawer{
		texture:    texture,
		dimensions: &dimensions,
		sourceRect: &sourceRect,
	}
	err := drawer.DrawTo(pos)
	assert.NoError(t, err)
	assert.Equal(t, drawSourceRect, &sourceRect, "sourceRect not equal")
	assert.Equal(t, destRect, drawDestRect, "destRect not equal")
}

func TestTextureDrawer_Dimensions(t *testing.T) {
	dimensions := primitives.Dimensions{
		Width:  123,
		Height: 234,
	}
	drawer := textureDrawer{
		dimensions: &dimensions,
	}
	assert.Equal(t, dimensions, drawer.Dimensions())
}
