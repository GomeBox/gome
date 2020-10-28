package game

import (
	"github.com/GomeBox/gome/adapters"
	"github.com/GomeBox/gome/internal/audio"
	audioInterfaces "github.com/GomeBox/gome/internal/audio/interfaces"
	"github.com/GomeBox/gome/internal/game/interfaces"
	"github.com/GomeBox/gome/internal/graphics"
	graphicsInterfaces "github.com/GomeBox/gome/internal/graphics/interfaces"
	"github.com/GomeBox/gome/internal/input"
	inputInterfaces "github.com/GomeBox/gome/internal/input/interfaces"
)

func newSystemsFactory(adapterSystem adapters.System) interfaces.SystemsFactory {
	factory := new(systemsFactory)
	factory.adapterSystem = adapterSystem
	return factory
}

type systemsFactory struct {
	adapterSystem adapters.System
}

func (factory systemsFactory) CreateGraphicsSystem() graphicsInterfaces.System {
	textureLoader := factory.adapterSystem.Graphics().TextureLoader()
	textureCreator := factory.adapterSystem.Graphics().TextureCreator()
	windowAdapter := factory.adapterSystem.Graphics().WindowAdapter()
	fontLoader := factory.adapterSystem.Graphics().FontLoader()
	return graphics.NewSystem(
		graphics.Adapters{
			TextureCreator: textureCreator,
			TextureLoader:  textureLoader,
			WindowAdapter:  windowAdapter,
			FontLoader:     fontLoader})
}

func (factory systemsFactory) CreateInputSystem() inputInterfaces.System {
	return input.NewSystem(
		input.Adapters{
			Keyboard: factory.adapterSystem.Input().Keyboard()})
}

func (factory systemsFactory) CreateAudioSystem() audioInterfaces.System {
	songLoader := factory.adapterSystem.Audio().SongLoader()
	soundLoader := factory.adapterSystem.Audio().SoundLoader()
	return audio.NewSystem(
		audio.Adapters{
			SoundLoader: soundLoader,
			SongLoader:  songLoader})
}
