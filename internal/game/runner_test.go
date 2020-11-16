package game

import (
	"errors"
	"github.com/GomeBox/gome/interfaces"
	"github.com/GomeBox/gome/internal/game/graphics"
	"github.com/GomeBox/gome/internal/game/mocks"
	"github.com/GomeBox/gome/primitives"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRunner(t *testing.T) {
	loopCalled := false
	loop := func(loopData *LoopData) error {
		loopCalled = true
		return nil
	}
	runner, sysMock := createRunner(loop)
	assert.NotNil(t, runner)
	assert.Same(t, sysMock, runner.gameSystem)
	_ = runner.loop(nil)
	assert.True(t, loopCalled)
}

func createRunner(loop Loop) (*runner, *mocks.System) {
	if loop == nil {
		loop = func(loopData *LoopData) error {
			return nil
		}
	}
	sysMock := new(mocks.System)
	testRunner := NewRunner(sysMock, loop)
	return testRunner.(*runner), sysMock
}

func TestRunner_initialize_InitializesGameSystem(t *testing.T) {
	runner, sysMock := createRunner(nil)
	retErr := false
	sysMock.OnInitialize = func() error {
		if retErr {
			return errors.New("test")
		}
		return nil
	}
	gameMock := new(mocks.Game)
	err := runner.initialize(gameMock, new(settings))
	assert.NoError(t, err)
	assert.Equal(t, 1, sysMock.CallCntInitialize)
}

func TestRunner_initialize_OpenWindow(t *testing.T) {
	runner, sysMock := createRunner(nil)
	var gotSettings graphics.WindowSettings
	retErr := false
	sysMock.OnOpenGameWindow = func(settings graphics.WindowSettings) error {
		if retErr {
			return errors.New("test")
		}
		gotSettings = settings
		return nil
	}
	wantSettings := graphics.WindowSettings{
		Fullscreen: true,
		Resolution: primitives.Dimensions{
			Width:  456,
			Height: 345,
		},
		Title: "test title",
	}
	settings := settings{windowSettings: wantSettings}
	err := runner.initialize(new(mocks.Game), &settings)
	assert.NoError(t, err)
	assert.Equal(t, gotSettings, wantSettings)
	retErr = true
	err = runner.initialize(new(mocks.Game), &settings)
	assert.Error(t, err)
}

func TestRunner_initialize_InitializesGame(t *testing.T) {
	context := new(context)
	game := new(mocks.Game)
	retErr := false
	var gotContext interfaces.Context
	game.OnInitialize = func(context interfaces.Context) error {
		if retErr {
			return errors.New("test")
		}
		gotContext = context
		return nil
	}
	runner, sys := createRunner(nil)
	sys.OnContext = func() interfaces.Context {
		return context
	}
	err := runner.initialize(game, new(settings))
	assert.NoError(t, err)
	assert.Same(t, context, gotContext)
}

func TestRunner_Run(t *testing.T) {
	var gotSettings interfaces.Settings
	game := &mocks.Game{
		OnSetup: func(settings interfaces.Settings) {
			gotSettings = settings
		},
	}
	var gotLoopData *LoopData
	loopCalled := false
	loop := func(data *LoopData) error {
		loopCalled = true
		gotLoopData = data
		return nil
	}
	runner, sysMock := createRunner(loop)
	sysMock.OnOpenGameWindow = func(settings graphics.WindowSettings) error {
		return nil
	}
	err := runner.Run(game)
	assert.NoError(t, err)
	assert.Equal(t, newSettings(), gotSettings)
	assert.Same(t, sysMock, gotLoopData.gameSystem)
	assert.True(t, loopCalled)

	testUpdate(t, sysMock, game, gotLoopData)
}

func testUpdate(t *testing.T, gameSystem *mocks.System, game *mocks.Game, loopData *LoopData) {
	var gotDelta float32
	var retError error
	game.OnUpdate = func(timeDelta float32, context interfaces.Context) (keepRunning bool, error error) {
		gotDelta = timeDelta
		return true, retError
	}
	var wantDelta float32
	wantDelta = 123.345
	keepRunning, err := loopData.update(wantDelta)
	assert.NoError(t, err)
	assert.True(t, keepRunning)
	assert.Equal(t, wantDelta, gotDelta)
	assert.Equal(t, 1, gameSystem.CallCntUpdate)
}
