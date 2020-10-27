package game

import (
	"github.com/GomeBox/gome/adapters"
	"github.com/GomeBox/gome/internal/audio"
	"github.com/GomeBox/gome/internal/graphics"
	"github.com/GomeBox/gome/internal/input"
)

func newSystemsFactory(adapterSystem adapters.System) *systemsFactory {
	factory := new(systemsFactory)
	factory.adapterSystem = adapterSystem
	return factory
}

type systemsFactory struct {
	adapterSystem adapters.System
}

func (factory systemsFactory) CreateGraphicsSystem() *graphics.System {
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

func (factory systemsFactory) CreateInputSystem() *input.System {
	return input.NewSystem(
		input.Adapters{
			Keyboard: factory.adapterSystem.Input().Keyboard()})
}

func (factory systemsFactory) CreateAudioSystem() *audio.System {
	songLoader := factory.adapterSystem.Audio().SongLoader()
	soundLoader := factory.adapterSystem.Audio().SoundLoader()
	return audio.NewSystem(
		audio.Adapters{
			SoundLoader: soundLoader,
			SongLoader:  songLoader})
}
