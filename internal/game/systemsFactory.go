package game

import (
	"github.com/GomeBox/gome/adapters"
	gomeInterfaces "github.com/GomeBox/gome/interfaces"
	"github.com/GomeBox/gome/internal/audio"
	"github.com/GomeBox/gome/internal/game/interfaces"
	"github.com/GomeBox/gome/internal/graphics"
	"github.com/GomeBox/gome/internal/input"
	inputInterfaces "github.com/GomeBox/gome/internal/input/interfaces"
)

func NewSystemsFactory(adapterSystem adapters.System) interfaces.SystemsFactory {
	factory := new(systemsFactoryImpl)
	factory.adapterSystem = adapterSystem
	return factory
}

type systemsFactoryImpl struct {
	adapterSystem adapters.System
}

func (factory systemsFactoryImpl) CreateGraphicsSystem() graphics.System {
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

func (factory systemsFactoryImpl) CreateAudioSystem() gomeInterfaces.Audio {
	songLoader := factory.adapterSystem.Audio().SongLoader()
	soundLoader := factory.adapterSystem.Audio().SoundLoader()
	return audio.NewSystem(
		audio.Adapters{
			SoundLoader: soundLoader,
			SongLoader:  songLoader})
}
