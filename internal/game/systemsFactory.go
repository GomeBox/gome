package game

import (
	"github.com/GomeBox/gome/adapters"
	"github.com/GomeBox/gome/internal/audio"
	audioInterfaces "github.com/GomeBox/gome/internal/audio/interfaces"
	"github.com/GomeBox/gome/internal/graphics"
	graphicsInterfaces "github.com/GomeBox/gome/internal/graphics/interfaces"
	"github.com/GomeBox/gome/internal/input"
	inputInterfaces "github.com/GomeBox/gome/internal/input/interfaces"
)

type systemsFactory interface {
	CreateGraphicsSystem() graphicsInterfaces.System
	CreateInputSystem() inputInterfaces.System
	CreateAudioSystem() audioInterfaces.System
}

func newSystemsFactory(adapterSystem adapters.System) systemsFactory {
	factory := new(systemsFactoryImpl)
	factory.adapterSystem = adapterSystem
	return factory
}

type systemsFactoryImpl struct {
	adapterSystem adapters.System
}

func (factory systemsFactoryImpl) CreateGraphicsSystem() graphicsInterfaces.System {
	textureLoader := factory.adapterSystem.Graphics().TextureLoader()
	textureCreator := factory.adapterSystem.Graphics().TextureCreator()
	windowAdapter := factory.adapterSystem.Graphics().WindowAdapter()
	fontLoader := factory.adapterSystem.Graphics().FontLoader()
	screenPresenter := factory.adapterSystem.Graphics().ScreenPresenter()
	return graphics.NewSystem(
		graphics.Adapters{
			TextureCreator:  textureCreator,
			TextureLoader:   textureLoader,
			WindowAdapter:   windowAdapter,
			FontLoader:      fontLoader,
			ScreenPresenter: screenPresenter})
}

func (factory systemsFactoryImpl) CreateInputSystem() inputInterfaces.System {
	return input.NewSystem(
		input.Adapters{
			Keyboard: factory.adapterSystem.Input().Keyboard()})
}

func (factory systemsFactoryImpl) CreateAudioSystem() audioInterfaces.System {
	songLoader := factory.adapterSystem.Audio().SongLoader()
	soundLoader := factory.adapterSystem.Audio().SoundLoader()
	return audio.NewSystem(
		audio.Adapters{
			SoundLoader: soundLoader,
			SongLoader:  songLoader})
}
