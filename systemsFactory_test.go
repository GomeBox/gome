package gome

import (
	audioMocks "github.com/GomeBox/gome/adapters/audio/mocks"
	graphicsMocks "github.com/GomeBox/gome/adapters/graphics/mocks"
	inputMocks "github.com/GomeBox/gome/adapters/input/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSystemsFactory_CreateAudioSystem(t *testing.T) {
	adapterMock := audioMocks.Adapters{}
	sysFact := systemsFactory{}
	audioSystem := sysFact.CreateAudioSystem(&adapterMock)
	assert.NotNil(t, audioSystem)
}

func TestSystemsFactory_CreateGraphicsSystem(t *testing.T) {
	adapterMock := graphicsMocks.Adapters{}
	sysFact := systemsFactory{}
	graphicsSystem := sysFact.CreateGraphicsSystem(&adapterMock)
	assert.NotNil(t, graphicsSystem)
}

func TestSystemsFactory_CreateInputSystem(t *testing.T) {
	adapterMock := inputMocks.Adapters{}
	sysFact := systemsFactory{}
	inputSystem := sysFact.CreateInputSystem(&adapterMock)
	assert.NotNil(t, inputSystem)
}
