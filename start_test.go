package internal

import (
	"errors"
	"github.com/GomeBox/gome/adapters"
	"github.com/GomeBox/gome/adapters/audio"
	audioMocks "github.com/GomeBox/gome/adapters/audio/mocks"
	"github.com/GomeBox/gome/adapters/graphics"
	graphicsMocks "github.com/GomeBox/gome/adapters/graphics/mocks"
	"github.com/GomeBox/gome/adapters/input"
	inputMocks "github.com/GomeBox/gome/adapters/input/mocks"
	"github.com/GomeBox/gome/adapters/mocks"
	"github.com/GomeBox/gome/interfaces"
	internalMocks "github.com/GomeBox/gome/internal/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStart(t *testing.T) {
	retErr := false
	createAdapters := func() (adapters.System, error) {
		if retErr {
			return nil, errors.New("test")
		}
		return &mocks.System{
			OnInput: func() input.Adapters {
				return &inputMocks.Adapters{}
			},
			OnGraphics: func() graphics.Adapters {
				return &graphicsMocks.Adapters{
					OnWindowAdapter: func() graphics.WindowAdapter {
						return &graphicsMocks.WindowAdapter{}
					},
					OnScreenPresenter: func() graphics.ScreenPresenter {
						return &graphicsMocks.ScreenPresenter{}
					},
				}
			},
			OnAudio: func() audio.Adapters {
				return &audioMocks.Adapters{}
			},
		}, nil
	}
	updateCalled := false
	game := &internalMocks.Game{
		OnUpdate: func(timeDelta float32, context interfaces.Context) (keepRunning bool, error error) {
			updateCalled = true
			return false, nil
		},
	}
	err := Start(game, createAdapters)
	assert.NoError(t, err)
	assert.True(t, updateCalled)

	retErr = true
	err = Start(game, createAdapters)
	assert.Error(t, err)
}
