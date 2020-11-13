package graphics

import (
	"errors"
	"github.com/GomeBox/gome/internal/graphics/interfaces"
	"github.com/GomeBox/gome/internal/graphics/mocks"
	"github.com/GomeBox/gome/primitives"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFont_CreateText(t *testing.T) {
	var gotText string
	var gotColor primitives.Color
	wantText := "test text"
	wantColor := primitives.Colors().White()
	wantInternalTextInstance := new(mocks.Text)
	retErr := false
	internalFont := &mocks.Font{
		OnCreateText: func(text string, color primitives.Color) (interfaces.Text, error) {
			gotText = text
			gotColor = color
			if retErr {
				return nil, errors.New("test")
			}
			return wantInternalTextInstance, nil
		},
	}
	font := font{internalFont: internalFont}
	gotTextInstance, err := font.CreateText(wantText, wantColor)
	assert.NoError(t, err)
	assert.Same(t, wantInternalTextInstance, gotTextInstance.(*text).internalText)
	assert.Equal(t, wantText, gotText)
	assert.Equal(t, wantColor, gotColor)
	retErr = true
	_, err = font.CreateText(wantText, wantColor)
	assert.Error(t, err)

}
