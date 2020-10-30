package game

import (
	"github.com/GomeBox/gome/adapters/audio"
	"github.com/GomeBox/gome/adapters/graphics"
	"github.com/GomeBox/gome/adapters/input"
	adapterMocks "github.com/GomeBox/gome/adapters/mocks"
)

func createAdapterSystemMock(
	inputAdapter input.Adapters,
	graphicsAdapter graphics.Adapters,
	audioAdapter audio.Adapters) *adapterMocks.System {
	return &adapterMocks.System{
		OnInput:    func() input.Adapters { return inputAdapter },
		OnGraphics: func() graphics.Adapters { return graphicsAdapter },
		OnAudio:    func() audio.Adapters { return audioAdapter },
	}
}
