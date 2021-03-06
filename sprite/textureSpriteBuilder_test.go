package sprite

import (
	"github.com/GomeBox/gome/internal/graphics/mocks"
	"github.com/GomeBox/gome/primitives"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFromTexture_Get(t *testing.T) {
	textureMock := new(mocks.Texture)
	s := FromTexture(textureMock).Get()
	if s == nil {
		t.Error("sprite not properly created")
	}
}

func TestFromTexture_SetPosition(t *testing.T) {
	textureMock := new(mocks.Texture)
	wantX := float32(123.24)
	wantY := float32(234.56)
	s := FromTexture(textureMock).SetPosition(wantX, wantY).Get()
	gotX := s.Position().X
	gotY := s.Position().Y
	if wantY != gotY {
		t.Errorf("Position not set (Y): got: %f, want %f", gotY, wantY)
	}
	if wantX != gotX {
		t.Errorf("Position not set (X): got: %f, want %f", gotX, wantX)
	}
}

func TestFromTexture_DrawerProperlySet(t *testing.T) {
	textureMock := new(mocks.Texture)
	drawCalled := false
	textureMock.OnDraw = func(source *primitives.Rectangle, dest primitives.RectangleF) error {
		drawCalled = true
		return nil
	}
	s := FromTexture(textureMock).Get()
	err := s.Draw()
	if err != nil {
		t.Errorf("err was expected to be nil but was %v", err)
	}
	if !drawCalled {
		t.Error("Texture.DrawF was not called")
	}
}

func TestTextureSpriteBuilder_SetSourceRect(t *testing.T) {
	textureSpriteBuilder := textureSpriteBuilder{
		sourceRect: nil,
	}
	wantSourceRect := primitives.NewRectangle(0, 0, 10, 10)
	returnedBuilder := textureSpriteBuilder.SetSourceRect(&wantSourceRect)
	assert.Same(t, &wantSourceRect, textureSpriteBuilder.sourceRect)
	assert.Same(t, &textureSpriteBuilder, returnedBuilder)
}
