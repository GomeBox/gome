package graphics

import (
	"errors"
	"github.com/GomeBox/gome/adapters/graphics/mocks"
	"github.com/GomeBox/gome/primitives"
	"testing"
)

func TestTexture_newTexture(t *testing.T) {
	adapter := &mocks.Texture{}
	tex := newTexture(adapter)
	if tex == nil {
		t.Error("newTexture returned nil")
	}
}

func TestTexture_Draw(t *testing.T) {
	firstCall := true
	wantSource := primitives.NewRectangle(100, 200, 34, 45)
	wantDest := primitives.NewRectangleF(5, 6, 7, 8)
	var gotSource *primitives.Rectangle
	var gotDest primitives.RectangleF
	adapter := &mocks.Texture{
		OnDraw: func(source *primitives.Rectangle, dest primitives.RectangleF) error {
			gotSource, gotDest = source, dest
			if firstCall {
				firstCall = false
				return nil
			}
			return errors.New("test")
		},
	}
	tex := newTexture(adapter)
	err := tex.Draw(&wantSource, wantDest)
	if err != nil {
		t.Error("DrawF returned unexpected error")
	}
	if &wantSource != gotSource {
		t.Errorf("source not as expected. Got %v, want %v", gotSource, wantSource)
	}
	if wantDest != gotDest {
		t.Errorf("dest not as expected. Got %v, want %v", gotDest, wantDest)
	}
	err = tex.Draw(&wantSource, wantDest)
	if err == nil {
		t.Error("DrawF did not return expected error of adapter")
	}
}

func TestTexture_Dimensions(t *testing.T) {
	want := primitives.Dimensions{
		Width:  123,
		Height: 45,
	}
	adapter := &mocks.Texture{
		OnDimensions: func() primitives.Dimensions {
			return want
		},
	}
	tex := newTexture(adapter)
	got := tex.Dimensions()
	if got != want {
		t.Errorf("Dimenions did not return expected value. Got %v, want %v", got, want)
	}
}
