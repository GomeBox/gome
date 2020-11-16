package game

import (
	"errors"
	adapterMocks "github.com/GomeBox/gome/adapters/mocks"
	"github.com/GomeBox/gome/internal/game/graphics"
	"github.com/GomeBox/gome/internal/game/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSystem(t *testing.T) {
	adapters := new(adapterMocks.System)
	graphicsMock := new(mocks.Graphics)
	audioMock := new(mocks.Audio)
	inputMock := new(mocks.Input)
	sys := NewSystem(adapters, graphicsMock, audioMock, inputMock)
	assert.NotNil(t, sys)
	s := sys.(*system)
	assert.Same(t, adapters, s.adapterSystem)
	assert.Same(t, graphicsMock, s.graphics)
	assert.Same(t, inputMock, s.input)
	assert.Same(t, audioMock, s.audio)
}

func TestSystem_Initialize(t *testing.T) {
	adapters := new(adapterMocks.System)
	retErr := false
	adapters.OnInitialize = func() error {
		if retErr {
			return errors.New("test")
		}
		return nil
	}
	sys := NewSystem(adapters, nil, nil, nil)
	err := sys.Initialize()
	assert.NoError(t, err)
	assert.Equal(t, 1, adapters.CallCntInitialize)
	retErr = true
	err = sys.Initialize()
	assert.Error(t, err)
}

func TestSystem_Update(t *testing.T) {
	adapters := new(adapterMocks.System)
	retErr := false
	adapters.OnUpdate = func() error {
		if retErr {
			return errors.New("test")
		}
		return nil
	}
	retErrInput := false
	input := new(mocks.Input)
	input.OnUpdate = func() error {
		if retErrInput {
			return errors.New("test")
		}
		return nil
	}
	sys := NewSystem(adapters, nil, nil, input)
	err := sys.Update()
	assert.NoError(t, err)
	assert.Equal(t, 1, adapters.CallCntUpdate, "adapters")
	assert.Equal(t, 1, input.CallCntUpdate, "input")
	retErr = true
	err = sys.Update()
	assert.Error(t, err, "adapters")
	retErr = false
	retErrInput = true
	err = sys.Update()
	assert.Error(t, err, "input")
}

func TestSystem_Graphics(t *testing.T) {
	graphicsMock := new(mocks.Graphics)
	sys := NewSystem(nil, graphicsMock, nil, nil)
	assert.Same(t, graphicsMock, sys.Graphics())
}

func TestSystem_Context(t *testing.T) {
	adapters := new(adapterMocks.System)
	graphicsMock := new(mocks.Graphics)
	audioMock := new(mocks.Audio)
	inputMock := new(mocks.Input)
	sys := NewSystem(adapters, graphicsMock, audioMock, inputMock)
	context := sys.Context()
	context2 := sys.Context()
	assert.NotNil(t, context)
	assert.Same(t, graphicsMock, context.Graphics())
	assert.Same(t, inputMock, context.Input())
	assert.Same(t, audioMock, context.Audio())
	assert.Same(t, context, context2)
}

func TestSystem_OpenGameWindow(t *testing.T) {
	graphicsMock := new(mocks.Graphics)
	var gotSettings, wantSettings graphics.WindowSettings
	wantSettings = graphics.WindowSettings{
		Fullscreen: true,
		Title:      "test",
	}
	retErr := false
	graphicsMock.OnOpenWindow = func(settings graphics.WindowSettings) error {
		gotSettings = settings
		if retErr {
			return errors.New("test")
		}
		return nil
	}
	sys := NewSystem(nil, graphicsMock, nil, nil)
	err := sys.OpenGameWindow(wantSettings)
	assert.NoError(t, err)
	assert.Equal(t, wantSettings, gotSettings)
	retErr = true
	err = sys.OpenGameWindow(wantSettings)
	assert.Error(t, err)
}
