package sprite

import (
	"github.com/GomeBox/gome/internal/graphics/mocks"
	"github.com/GomeBox/gome/primitives"
	"testing"
)

func TestFromText_Get(t *testing.T) {
	textMock := new(mocks.Text)
	s := FromText(textMock).Get()
	if s == nil {
		t.Error("sprite not properly created")
	}
}

func TestFromText_SetPosition(t *testing.T) {
	textMock := new(mocks.Text)
	wantX := float32(123.24)
	wantY := float32(234.56)
	s := FromText(textMock).SetPosition(wantX, wantY).Get()
	gotX := s.Position().X
	gotY := s.Position().Y
	if wantY != gotY {
		t.Errorf("Position not set (Y): got: %f, want %f", gotY, wantY)
	}
	if wantX != gotX {
		t.Errorf("Position not set (X): got: %f, want %f", gotX, wantX)
	}
}

func TestFromText_DrawerProperlySet(t *testing.T) {
	textMock := new(mocks.Text)
	drawCalled := false
	textMock.OnDrawF = func(f *primitives.PointF) error {
		drawCalled = true
		return nil
	}
	s := FromText(textMock).Get()
	err := s.Draw()
	if err != nil {
		t.Errorf("err was expected to be nil but was %v", err)
	}
	if !drawCalled {
		t.Error("Text.DrawF was not called")
	}
}
