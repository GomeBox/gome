package gome

import (
	"errors"
	"github.com/GomeBox/gome/adapters"
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
	"github.com/GomeBox/gome/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStart(t *testing.T) {
	game := &mocks.Game{}
	retErr := false
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
	createAdapters := func() (adapters.System, error) {
		if retErr {
			return nil, errors.New("test")
		}
		return adapterSystem, nil
	}
	systemsFactory := &mocks.SystemsFactory{
		OnCreateGraphicsSystem: func(adapters graphics.Adapters) gameGraphics.System {
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
	err := Start(game, createAdapters, systemsFactory, loop)
	assert.NoError(t, err)
	assert.True(t, loopCalled)
	assert.Equal(t, 1, systemsFactory.CallCntCreateInputSystem)
	assert.Equal(t, 1, systemsFactory.CallCntCreateGraphicsSystem)
	assert.Equal(t, 1, systemsFactory.CallCntCreateAudioSystem)

	retErr = true
	err = Start(game, createAdapters, systemsFactory, loop)
	assert.Error(t, err)

	retErr = false
	loopErr = true
	err = Start(game, createAdapters, systemsFactory, loop)
	assert.Error(t, err)
}
