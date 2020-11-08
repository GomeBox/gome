package graphics

import (
	"errors"
	"github.com/GomeBox/gome/adapters/graphics"
	"github.com/GomeBox/gome/adapters/graphics/mocks"
	"github.com/GomeBox/gome/primitives"
	"testing"
)

func TestFont_newFont(t *testing.T) {
	creator := new(mocks.TextCreator)
	font := newFont(creator)
	if font == nil {
		t.Error("newFont() returned nil")
	}
}

func TestFont_CreateText(t *testing.T) {
	wantText := "Test text"
	wantColor := primitives.Colors().White()
	var gotText string
	var gotColor primitives.Color
	creator := new(mocks.TextCreator)
	firstCall := true
	creator.OnCreate = func(text string, color primitives.Color) (graphics.TextDrawer, error) {
		if firstCall {
			gotText = text
			gotColor = color
			firstCall = false
			return nil, nil
		} else {
			return nil, errors.New("test")
		}
	}
	font := newFont(creator)
	text, err := font.CreateText(wantText, wantColor)
	if err != nil {
		t.Errorf("CreateText() returned unexpected error %v", err)
	}
	if text == nil {
		t.Error("CreateText() returned text == nil")
	}
	if gotText != wantText {
		t.Errorf("text was not passed to creator. got %v, want %v", gotText, wantText)
	}
	if gotColor != wantColor {
		t.Errorf("color was not passed to creator. got %v, want %v", gotColor, wantColor)
	}
	_, err = font.CreateText(wantText, wantColor)
	if err == nil {
		t.Error("CreateText did not return expected error")
	}
}
