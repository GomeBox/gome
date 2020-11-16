package internal

import (
	"errors"
	"github.com/GomeBox/gome/adapters/audio"
	audioAdapterMocks "github.com/GomeBox/gome/adapters/audio/mocks"
	"github.com/GomeBox/gome/adapters/graphics"
	graphicsAdapterMocks "github.com/GomeBox/gome/adapters/graphics/mocks"
	"github.com/GomeBox/gome/adapters/input"
	inputAdapterMocks "github.com/GomeBox/gome/adapters/input/mocks"
	adapterMocks "github.com/GomeBox/gome/adapters/mocks"
	internalGame "github.com/GomeBox/gome/internal/game"
	gameGraphics "github.com/GomeBox/gome/internal/game/graphics"
	gameMocks "github.com/GomeBox/gome/internal/game/mocks"
	"github.com/GomeBox/gome/internal/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStart(t *testing.T) {
	game := &mocks.Game{}
	adapterSystem := &adapterMocks.System{
		OnInput: func() input.Adapters {
			return &inputAdapterMocks.Adapters{}
		},
		OnGraphics: func() graphics.Adapters {
			return &graphicsAdapterMocks.Adapters{}
		},
		OnAudio: func() audio.Adapters {
			return &audioAdapterMocks.Adapters{}
		},
	}
	systemsFactory := &mocks.SystemsFactory{
		OnCreateGraphicsSystem: func() gameGraphics.System {
			return &gameMocks.Graphics{}
		},
	}
	loopCalled := false
	loopErr := false
	loop := func(data *internalGame.LoopData) error {
		loopCalled = true
		if loopErr {
			return errors.New("test")
		}
		return nil
	}
	err := Start(game, adapterSystem, systemsFactory, loop)
	assert.NoError(t, err)
	assert.True(t, loopCalled)
	assert.Equal(t, 1, systemsFactory.CallCntCreateInputSystem)
	assert.Equal(t, 1, systemsFactory.CallCntCreateGraphicsSystem)
	assert.Equal(t, 1, systemsFactory.CallCntCreateAudioSystem)

	loopErr = true
	err = Start(game, adapterSystem, systemsFactory, loop)
	assert.Error(t, err)
}
