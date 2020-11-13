package graphics

import (
	"errors"
	"github.com/GomeBox/gome/internal/graphics/mocks"
	"github.com/GomeBox/gome/primitives"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestText_Draw(t *testing.T) {
	var gotPosition *primitives.Point
	wantPosition := &primitives.Point{
		X: 340,
		Y: 2134,
	}
	retErr := false
	internalText := mocks.Text{
		OnDraw: func(position *primitives.Point) error {
			gotPosition = position
			if retErr {
				return errors.New("test")
			}
			return nil
		},
	}
	text := text{internalText: internalText}
	err := text.Draw(wantPosition)
	assert.NoError(t, err)
	assert.Equal(t, wantPosition, gotPosition)
	retErr = true
	err = text.Draw(wantPosition)
	assert.Error(t, err)
}

func TestText_DrawF(t *testing.T) {
	var gotPosition *primitives.PointF
	wantPosition := &primitives.PointF{
		X: 340.3,
		Y: 2134.01,
	}
	retErr := false
	internalText := mocks.Text{
		OnDrawF: func(position *primitives.PointF) error {
			gotPosition = position
			if retErr {
				return errors.New("test")
			}
			return nil
		},
	}
	text := text{internalText: internalText}
	err := text.DrawF(wantPosition)
	assert.NoError(t, err)
	assert.Equal(t, wantPosition, gotPosition)
	retErr = true
	err = text.DrawF(wantPosition)
	assert.Error(t, err)
}

func TestText_Dimensions(t *testing.T) {
	wantDimensions := primitives.Dimensions{
		Width:  239,
		Height: 34,
	}
	internalText := mocks.Text{
		OnDimensions: func() primitives.Dimensions {
			return wantDimensions
		},
	}
	text := text{internalText: internalText}
	gotDimensions := text.Dimensions()
	assert.Equal(t, wantDimensions, gotDimensions)
}
