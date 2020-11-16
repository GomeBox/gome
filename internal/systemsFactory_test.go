package internal

import (
	"github.com/GomeBox/gome/adapters/audio"
	audioMocks "github.com/GomeBox/gome/adapters/audio/mocks"
	"github.com/GomeBox/gome/adapters/graphics"
	graphicsMocks "github.com/GomeBox/gome/adapters/graphics/mocks"
	"github.com/GomeBox/gome/adapters/input"
	inputMocks "github.com/GomeBox/gome/adapters/input/mocks"
	"github.com/GomeBox/gome/adapters/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSystemsFactory_CreateAudioSystem(t *testing.T) {
	adapterSystem := &mocks.System{
		OnAudio: func() audio.Adapters {
			return &audioMocks.Adapters{}
		},
	}
	sysFact := NewSystemsFactory(adapterSystem)
	audioSystem := sysFact.CreateAudioSystem()
	assert.NotNil(t, audioSystem)
}

func TestSystemsFactory_CreateGraphicsSystem(t *testing.T) {
	adapterSystem := &mocks.System{
		OnGraphics: func() graphics.Adapters {
			return &graphicsMocks.Adapters{}
		},
	}
	sysFact := NewSystemsFactory(adapterSystem)
	graphicsSystem := sysFact.CreateGraphicsSystem()
	assert.NotNil(t, graphicsSystem)
}

func TestSystemsFactory_CreateInputSystem(t *testing.T) {
	adapterSystem := &mocks.System{
		OnInput: func() input.Adapters {
			return &inputMocks.Adapters{}
		},
	}
	sysFact := NewSystemsFactory(adapterSystem)
	inputSystem := sysFact.CreateInputSystem()
	assert.NotNil(t, inputSystem)
}
