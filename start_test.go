package gome

import (
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
	createAdapters := func() adapters.System {
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
		}
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
}
