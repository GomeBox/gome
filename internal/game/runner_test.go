package game

import (
	"github.com/GomeBox/gome/adapters"
	"github.com/GomeBox/gome/adapters/audio"
	"github.com/GomeBox/gome/adapters/graphics"
	"github.com/GomeBox/gome/adapters/input"
	adapterMocks "github.com/GomeBox/gome/adapters/mocks"
	gameMocks "github.com/GomeBox/gome/game/mocks"
	"github.com/GomeBox/gome/mocks"
	"testing"
)

func TestRunner_NewRunner(t *testing.T) {
	runner := NewRunner()
	if runner == nil {
		t.Error("returned nil")
	}
}

func TestRunner_Init(t *testing.T) {
	runner := NewRunner()
	createAdaptersCalled := false
	settings := new(gameMocks.Settings)
	adapterSystem := createAdapterSystemMock(
		new(mocks.InputAdapters),
		new(mocks.GraphicsAdapters),
		new(mocks.AudioAdapters))
	createAdapters := func() (adapters.System, error) {
		createAdaptersCalled = true
		return adapterSystem, nil
	}
	err := runner.Init(createAdapters, settings)
	if err != nil {
		t.Error("Init returned an error")
		return
	}
	if !createAdaptersCalled {
		t.Error("createAdapters was not called")
		return
	}
	got := adapterSystem.CallCntInitialize
	if got != 1 {
		t.Errorf("adapterSystem.Initialize was not called expected numbe of times. Got %d, want: %d", got, 1)
	}
	if runner.settings != settings {
		t.Errorf("runner.settings was not set")
	}
	if runner.quit != false {
		t.Errorf("runner.quit was not set to false")
	}
}

func TestRunner_InitChecksAdapters(t *testing.T) {
	settings := new(gameMocks.Settings)
	inputAdapter := new(mocks.InputAdapters)
	graphicsAdapter := new(mocks.GraphicsAdapters)
	audioAdapter := new(mocks.AudioAdapters)
	adapterSystems := make(map[string]adapters.System, 3)
	adapterSystems["input"] = createAdapterSystemMock(nil, graphicsAdapter, audioAdapter)
	adapterSystems["graphics"] = createAdapterSystemMock(inputAdapter, nil, audioAdapter)
	adapterSystems["audio"] = createAdapterSystemMock(inputAdapter, graphicsAdapter, nil)
	for adapterName, sys := range adapterSystems {
		runner := NewRunner()
		createAdapters := func() (adapters.System, error) { return sys, nil }
		err := runner.Init(createAdapters, settings)
		if err == nil {
			t.Errorf("Init did not return error but adapter %v was nil", adapterName)
		}
	}
}

func createAdapterSystemMock(
	inputAdapter input.Adapters,
	graphicsAdapter graphics.Adapters,
	audioAdapter audio.Adapters) *adapterMocks.System {
	return &adapterMocks.System{
		OnInput:    func() input.Adapters { return inputAdapter },
		OnGraphics: func() graphics.Adapters { return graphicsAdapter },
		OnAudio:    func() audio.Adapters { return audioAdapter },
	}
}
