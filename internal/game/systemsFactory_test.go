package game

import (
	"github.com/GomeBox/gome/adapters/audio"
	audioMocks "github.com/GomeBox/gome/adapters/audio/mocks"
	"github.com/GomeBox/gome/adapters/graphics"
	graphicsMocks "github.com/GomeBox/gome/adapters/graphics/mocks"
	"github.com/GomeBox/gome/adapters/input"
	inputMocks "github.com/GomeBox/gome/adapters/input/mocks"
	"github.com/GomeBox/gome/adapters/mocks"
	"testing"
)

func TestSystemsFactory_newSystemsFactory(t *testing.T) {
	factory := NewSystemsFactory(nil)
	if factory == nil {
		t.Error("NewSystemsFactory returned nil")
	}
}

func TestSystemsFactory_CreateGraphicsSystem(t *testing.T) {
	textureLoader := new(graphicsMocks.TextureLoader)
	textureCreator := new(graphicsMocks.TextureCreator)
	fontLoader := new(graphicsMocks.FontLoader)
	windowAdapter := new(graphicsMocks.WindowAdapter)
	graphicsAdapter := &graphicsMocks.Adapters{
		OnTextureLoader:  func() graphics.TextureLoader { return textureLoader },
		OnTextureCreator: func() graphics.TextureCreator { return textureCreator },
		OnFontLoader:     func() graphics.FontLoader { return fontLoader },
		OnWindowAdapter:  func() graphics.WindowAdapter { return windowAdapter },
	}
	adapterSystem := mocks.System{
		OnGraphics: func() graphics.Adapters {
			return graphicsAdapter
		},
	}
	systemsFactory := NewSystemsFactory(&adapterSystem)
	_ = systemsFactory.CreateGraphicsSystem()
	want := 1
	got := graphicsAdapter.CallCntTextureLoader
	if got != want {
		t.Errorf("TextureLoader() was not called expected number of times. Got: %d, want: %d", got, want)
	}
	got = graphicsAdapter.CallCntTextureCreator
	if got != want {
		t.Errorf("TextureCreator() was not called expected number of times. Got: %d, want: %d", got, want)
	}
	got = graphicsAdapter.CallCntWindowAdapter
	if got != want {
		t.Errorf("WindowAdapter() was not called expected number of times. Got: %d, want: %d", got, want)
	}
	got = graphicsAdapter.CallCntFontLoader
	if got != want {
		t.Errorf("FontLoader() was not called expected number of times. Got: %d, want: %d", got, want)
	}
}

func TestSystemsFactory_CreateInputSystem(t *testing.T) {
	keyboardAdapter := new(inputMocks.KeyboardAdapter)
	inputAdapter := &inputMocks.Adapters{
		OnKeyboard:      func() input.KeyboardAdapter { return keyboardAdapter },
		CallCntKeyboard: 0,
	}
	adapterSystem := &mocks.System{
		OnInput: func() input.Adapters { return inputAdapter },
	}
	systemsFactory := NewSystemsFactory(adapterSystem)
	_ = systemsFactory.CreateInputSystem()
	want := 1
	got := inputAdapter.CallCntKeyboard
	if got != want {
		t.Errorf("KeyboardMock() was not called expected number of times. Got: %d, want: %d", got, want)
	}
}

func TestSystemsFactory_CreateAudioSystem(t *testing.T) {
	songLoader := new(audioMocks.SongLoader)
	soundLoader := new(audioMocks.SoundLoader)
	audioAdapter := &audioMocks.Adapters{
		OnSoundLoader:      func() audio.SoundLoader { return soundLoader },
		OnSongLoader:       func() audio.SongLoader { return songLoader },
		CallCntSoundLoader: 0,
		CallCntSongLoader:  0,
	}
	adapterSystem := &mocks.System{
		OnAudio: func() audio.Adapters {
			return audioAdapter
		},
	}
	systemsFactory := NewSystemsFactory(adapterSystem)
	_ = systemsFactory.CreateAudioSystem()
	want := 1
	got := audioAdapter.CallCntSongLoader
	if got != want {
		t.Errorf("SongLoader() not called expected number of times. Got: %d, want: %d", got, want)
	}
	got = audioAdapter.CallCntSoundLoader
	if got != want {
		t.Errorf("SoundLoader() not called expected number of times. Got: %d, want: %d", got, want)
	}
}
