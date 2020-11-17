package graphics

import (
	"errors"
	"github.com/GomeBox/gome/adapters/graphics/mocks"
	"github.com/GomeBox/gome/primitives"
	"testing"
)

func TestText_newText(t *testing.T) {
	drawer := &mocks.TextDrawer{}
	text := newText(drawer)
	if text == nil {
		t.Error("newText returned nil")
	}
}

func TestText_Draw(t *testing.T) {
	testPos := primitives.PointF{
		X: 100,
		Y: 200,
	}
	var gotPos primitives.PointF
	firstCall := true
	drawer := &mocks.TextDrawer{
		OnDraw: func(position primitives.PointF) error {
			gotPos = position
			if firstCall {
				firstCall = false
				return nil
			}
			return errors.New("test")
		},
	}
	text := newText(drawer)
	err := text.Draw(testPos)
	if err != nil {
		t.Errorf("Draw returned unexpected error, %v", err)
	}
	if gotPos != testPos {
		t.Errorf("position was not correctly passed to drawer")
	}
	err = text.Draw(testPos)
	if err == nil {
		t.Error("Draw did not return expected error")
	}
}

func TestText_Dimensions(t *testing.T) {
	want := primitives.Dimensions{
		Width:  123,
		Height: 234,
	}
	drawer := &mocks.TextDrawer{
		OnDimensions: func() primitives.Dimensions {
			return want
		},
	}
	text := newText(drawer)
	got := text.Dimensions()
	if got != want {
		t.Errorf("Dimensions() did not return expected value. Got %v, want %v", got, want)
	}
}
