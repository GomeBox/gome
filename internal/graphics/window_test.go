package graphics

import (
	"github.com/GomeBox/gome/adapters/graphics"
	"github.com/GomeBox/gome/adapters/graphics/mocks"
	"github.com/GomeBox/gome/primitives"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWindow_newWindow(t *testing.T) {
	adapter := new(mocks.WindowAdapter)
	window := newWindow(adapter)
	if window == nil {
		t.Error("newWindow returned nil")
	}
}

func TestWindow_Dimensions(t *testing.T) {
	open := true
	want := primitives.Dimensions{
		Width:  135,
		Height: 67,
	}
	adapter := &mocks.WindowAdapter{
		OnIsOpen: func() bool {
			return open
		},
		OnSize: func() primitives.Dimensions {
			return want
		},
	}
	window := newWindow(adapter)
	got, err := window.Dimensions()
	if err != nil {
		t.Errorf("Dimensions returned unexpected error %v", err)
	}
	if got != want {
		t.Errorf("Dimensions returned unexpected value. Got %v, want %v", got, want)
	}
	open = false
	_, err = window.Dimensions()
	if err == nil {
		t.Error("Dimensions does not return error, when window is not open")
	}
}

func TestWindow_Open(t *testing.T) {
	want := WindowSettings{
		Fullscreen: true,
		Resolution: primitives.Dimensions{
			Width:  499,
			Height: 233,
		},
		Title: "test title",
	}
	var got *graphics.WindowSettings
	adapter := &mocks.WindowAdapter{
		OnOpenWindow: func(windowSettings *graphics.WindowSettings) error {
			got = windowSettings
			return nil
		},
	}
	window := newWindow(adapter)
	err := window.Open(want)
	assert.NoError(t, err)
	assert.Equal(t, want.Fullscreen, got.Fullscreen)
	assert.Equal(t, want.Resolution, got.Resolution)
	assert.Equal(t, want.Title, got.Title)
}
