package graphics

import (
	"errors"
	"github.com/GomeBox/gome/adapters/graphics"
	"github.com/GomeBox/gome/adapters/graphics/mocks"
	"github.com/GomeBox/gome/primitives"
	"testing"
)

func TestNewSystem(t *testing.T) {
	adapters := createAdapters()
	sys := NewSystem(adapters)
	if sys == nil {
		t.Error("NewSystem() returned nil")
	}
}

func TestSystem_LoadTexture(t *testing.T) {
	var gotFileName string
	testFileName := "myFile.file"
	adapters := createAdapters()
	firstCall := true
	textureLoader := &mocks.TextureLoader{
		OnLoad: func(fileName string) (graphics.Texture, error) {
			gotFileName = fileName
			if firstCall {
				firstCall = false
				return nil, nil
			} else {
				return nil, errors.New("test")
			}
		}}
	adapters.OnTextureLoader = func() graphics.TextureLoader {
		return textureLoader
	}
	sys := NewSystem(adapters)
	texture, err := sys.LoadTexture(testFileName)
	if err != nil {
		t.Errorf("LoadTexture returned unexpected error, %v", err)
	}
	if texture == nil {
		t.Error("LoadTexture returned texture = nil")
	}
	if gotFileName != testFileName {
		t.Errorf("fileName was not passed to textureLoader. got %v, want %v", gotFileName, testFileName)
	}
	_, err = sys.LoadTexture(testFileName)
	if err == nil {
		t.Error("LoadTexture did not return error of textureLoader")
	}
}

func TestSystem_LoadFont(t *testing.T) {
	var gotFileName string
	var gotSize int
	testFileName := "myFont.font"
	testSize := 44
	adapters := createAdapters()
	firstCall := true
	fontLoader := &mocks.FontLoader{
		OnLoad: func(fileName string, size int) (graphics.TextCreator, error) {
			gotFileName = fileName
			gotSize = size
			if firstCall {
				firstCall = false
				return nil, nil
			} else {
				return nil, errors.New("test")
			}
		}}
	adapters.OnFontLoader = func() graphics.FontLoader {
		return fontLoader
	}
	sys := NewSystem(adapters)
	font, err := sys.LoadFont(testFileName, testSize)
	if err != nil {
		t.Errorf("LoadFont returned unexpected error, %v", err)
	}
	if font == nil {
		t.Error("LoadFont returned font = nil")
	}
	if gotFileName != testFileName {
		t.Errorf("fileName was not passed to fontLoader. got %v, want %v", gotFileName, testFileName)
	}
	if gotSize != testSize {
		t.Errorf("size was not passed to fontLoader. got %v, want %v", gotSize, testSize)
	}
	_, err = sys.LoadFont(testFileName, testSize)
	if err == nil {
		t.Error("LoadFont did not return error of fontLoader")
	}
}

func TestSystem_CreateTexture(t *testing.T) {
	var gotDimensions primitives.Dimensions
	var gotColor primitives.Color
	testDimensions := primitives.Dimensions{
		Width:  123,
		Height: 345,
	}
	testColor := primitives.Colors().White()
	adapters := createAdapters()
	firstCall := true
	textureCreator := &mocks.TextureCreator{
		OnCreate: func(dimensions primitives.Dimensions, color primitives.Color) (graphics.Texture, error) {
			gotDimensions = dimensions
			gotColor = color
			if firstCall {
				firstCall = false
				return nil, nil
			} else {
				return nil, errors.New("test")
			}
		}}
	adapters.OnTextureCreator = func() graphics.TextureCreator {
		return textureCreator
	}
	sys := NewSystem(adapters)
	font, err := sys.CreateTexture(testDimensions, testColor)
	if err != nil {
		t.Errorf("CreateTexture returned unexpected error, %v", err)
	}
	if font == nil {
		t.Error("CreateTexture returned font = nil")
	}
	if gotColor != testColor {
		t.Errorf("color was not passed to textureCreator. got %v, want %v", gotColor, testColor)
	}
	if gotDimensions != testDimensions {
		t.Errorf("dimensions were not passed to textureCreator. got %v, want %v", gotDimensions, testDimensions)
	}
	_, err = sys.CreateTexture(testDimensions, testColor)
	if err == nil {
		t.Error("CreateTexture did not return error of textureCreator")
	}
}

func TestSystem_Clear(t *testing.T) {
	adapters := createAdapters()
	screenPresenter := new(mocks.ScreenPresenter)
	adapters.OnScreenPresenter = func() graphics.ScreenPresenter {
		return screenPresenter
	}
	sys := NewSystem(adapters)
	_ = sys.Clear()
	if screenPresenter.CallCntClear != 1 {
		t.Errorf("screenPresenter.Clear() was not called expected number of times. got %v, want %v", screenPresenter.CallCntClear, 1)
	}
}

func TestSystem_Present(t *testing.T) {
	adapters := createAdapters()
	screenPresenter := new(mocks.ScreenPresenter)
	adapters.OnScreenPresenter = func() graphics.ScreenPresenter {
		return screenPresenter
	}
	sys := NewSystem(adapters)
	_ = sys.Present()
	if screenPresenter.CallCntPresent != 1 {
		t.Errorf("screenPresenter.Present() was not called expected number of times. got %v, want %v", screenPresenter.CallCntPresent, 1)
	}
}

func TestSystem_Window(t *testing.T) {
	adapters := createAdapters()
	system := NewSystem(adapters)
	win := system.Window()
	if win == nil {
		t.Error("Window() returned nil")
	}
}

func createAdapters() *mocks.Adapters {
	return &mocks.Adapters{
		OnTextureLoader: func() graphics.TextureLoader {
			return new(mocks.TextureLoader)
		},
		OnTextureCreator: func() graphics.TextureCreator {
			return new(mocks.TextureCreator)
		},
		OnFontLoader: func() graphics.FontLoader {
			return new(mocks.FontLoader)
		},
		OnWindowAdapter: func() graphics.WindowAdapter {
			return new(mocks.WindowAdapter)
		},
		OnScreenPresenter: func() graphics.ScreenPresenter {
			return new(mocks.ScreenPresenter)
		},
	}
}
