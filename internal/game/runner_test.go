package game

import (
	"errors"
	"github.com/GomeBox/gome/adapters"
	audioAdapters "github.com/GomeBox/gome/adapters/audio"
	audioAdapterMocks "github.com/GomeBox/gome/adapters/audio/mocks"
	graphicsAdapters "github.com/GomeBox/gome/adapters/graphics"
	graphicsAdapterMocks "github.com/GomeBox/gome/adapters/graphics/mocks"
	inputAdapters "github.com/GomeBox/gome/adapters/input"
	inputAdapterMocks "github.com/GomeBox/gome/adapters/input/mocks"
	adapterMocks "github.com/GomeBox/gome/adapters/mocks"
	"github.com/GomeBox/gome/internal/game/interfaces"
	"github.com/GomeBox/gome/internal/game/mocks"
	"github.com/GomeBox/gome/internal/graphics"
	graphicsInterfaces "github.com/GomeBox/gome/internal/graphics/interfaces"
	graphicsMocks "github.com/GomeBox/gome/internal/graphics/mocks"
	"testing"
)

func TestRunner_NewRunner(t *testing.T) {
	initCallback := func(gameSystem interfaces.System) error {
		return nil
	}
	updateCallback := func(timeDelta float32) (bool, error) {
		return true, nil
	}
	drawCallback := func(timeDelta float32) error {
		return nil
	}
	createAdapters := func() (adapters.System, error) {
		graphicAdapterMock := new(graphicsAdapterMocks.Adapters)
		inputAdapterMock := new(inputAdapterMocks.Adapters)
		audioAdapterMock := new(audioAdapterMocks.Adapters)
		adapterSystem := &adapterMocks.System{
			OnInput:    func() inputAdapters.Adapters { return inputAdapterMock },
			OnGraphics: func() graphicsAdapters.Adapters { return graphicAdapterMock },
			OnAudio:    func() audioAdapters.Adapters { return audioAdapterMock },
		}
		return adapterSystem, nil
	}
	getSettings := func() interfaces.Settings {
		return new(mocks.Settings)
	}
	callbacks := &interfaces.Callbacks{
		Init:           initCallback,
		Update:         updateCallback,
		Draw:           drawCallback,
		CreateAdapters: createAdapters,
		GetSettings:    getSettings,
	}
	runner, err := NewRunner(callbacks)
	if err != nil {
		t.Fatalf("NewRunner() returned unexpected error %v", err)
	}
	if runner == nil {
		t.Fatal("NewRunner() returned nil")
	}
}

func TestRunner_initialize_InitializesGameSystem(t *testing.T) {
	initCallBack := func(gameSystem interfaces.System) error {
		return nil
	}
	gameSystem := new(mocks.System)
	gameSystem.OnGraphics = createGraphicsMock(nil)
	settings := new(mocks.Settings)
	err := initialize(gameSystem, settings, initCallBack)
	if err != nil {
		t.Fatalf("initialize returned error %v", err)
	}
	got := gameSystem.CallCntInitialize
	want := 1
	if got != want {
		t.Errorf("gameSystem.Initialize() was called incorrenct number of times. Got %v, want %v", got, want)
	}
}

func TestRunner_initialize_ReturnsErrorOfGameSystemInitialize(t *testing.T) {
	gameSystem := new(mocks.System)
	gameSystem.OnInitialize = func() error {
		return errors.New("test")
	}
	err := initialize(gameSystem, nil, nil)
	if err == nil {
		t.Error("initialize did not return expected error")
	}
}

func TestRunner_initialize_OpenWindow(t *testing.T) {
	initCallBack := func(gameSystem interfaces.System) error {
		return nil
	}
	gameSystem := new(mocks.System)
	var wantWindowSettings *graphicsAdapters.WindowSettings
	callCntOpenWindow := 0
	onOpenWindow := func(settings *graphicsAdapters.WindowSettings) error {
		callCntOpenWindow++
		wantWindowSettings = settings
		return nil
	}
	gameSystem.OnGraphics = createGraphicsMock(onOpenWindow)
	settings := new(mocks.Settings)
	windowSettings := new(graphicsAdapters.WindowSettings)
	settings.OnWindowSettings = func() *graphicsAdapters.WindowSettings {
		return windowSettings
	}
	err := initialize(gameSystem, settings, initCallBack)
	if err != nil {
		t.Fatalf("initialize returned error %v", err)
	}
	if callCntOpenWindow != 1 {
		t.Errorf("Window.Open() was not called once. Got %v, want %v", callCntOpenWindow, 1)
	}
	if wantWindowSettings != windowSettings {
		t.Error("windowSettings were not passed to Window.Open() correctly")
	}
}

func TestRunner_initialize_ReturnsErrorOfOpenWindow(t *testing.T) {
	gameSystem := new(mocks.System)
	onOpenWindow := func(settings *graphicsAdapters.WindowSettings) error {
		return errors.New("test")
	}
	graphicsMock := createGraphicsMock(onOpenWindow)
	gameSystem.OnGraphics = graphicsMock
	settings := new(mocks.Settings)
	err := initialize(gameSystem, settings, nil)
	if err == nil {
		t.Fatalf("initialize did not return expected error")
	}
}

func TestRunner_initialize_CallsInitCallback(t *testing.T) {
	gameSystem := new(mocks.System)
	graphicsMock := createGraphicsMock(nil)
	gameSystem.OnGraphics = graphicsMock
	settings := new(mocks.Settings)
	initCallbackCallCnt := 0
	var wantGameSystem interfaces.System
	initCallback := func(gameSystem interfaces.System) error {
		initCallbackCallCnt++
		wantGameSystem = gameSystem
		return nil
	}
	err := initialize(gameSystem, settings, initCallback)
	if err != nil {
		t.Fatalf("initialize returned error %v", err)
	}
	if initCallbackCallCnt != 1 {
		t.Errorf("InitCallback was not called correct number of times. Got %v, want %v", initCallbackCallCnt, 1)
	}
	if wantGameSystem != gameSystem {
		t.Error("gameSystem was not passed to initCallback")
	}
}

func TestRunner_initialize_ReturnsErrorOfInitCallback(t *testing.T) {
	gameSystem := new(mocks.System)
	graphicsMock := createGraphicsMock(nil)
	gameSystem.OnGraphics = graphicsMock
	settings := new(mocks.Settings)
	initCallback := func(gameSystem interfaces.System) error {
		return errors.New("test")
	}
	err := initialize(gameSystem, settings, initCallback)
	if err == nil {
		t.Error("did not return error")
	}
}

func createGraphicsMock(onWindowOpen func(settings *graphicsAdapters.WindowSettings) error) func() graphicsInterfaces.System {
	return func() graphicsInterfaces.System {
		graphicsSystem := new(graphics.SystemMock)
		graphicsSystem.OnWindow = func() graphicsInterfaces.Window {
			window := new(graphicsMocks.Window)
			window.OnOpen = onWindowOpen
			return window
		}
		return graphicsSystem
	}
}

func TestRunner_run(t *testing.T) {
	loopCallCnt := 0
	var gotLoopData *loopData
	loop := func(data *loopData) error {
		loopCallCnt++
		gotLoopData = data
		return nil
	}
	err := run(nil, nil, nil, loop)
	if err != nil {
		t.Error("run returned unexpected error")
	}
	if loopCallCnt != 1 {
		t.Errorf("loop function not called expected number of times. Got %v, want %v", loopCallCnt, 1)
	}
	if gotLoopData == nil {
		t.Error("loopData was nil")
	}
	loop = func(data *loopData) error {
		return errors.New("test")
	}
	err = run(nil, nil, nil, loop)
	if err == nil {
		t.Error("run did not return expected error")
	}
}

func TestRunner_createLoopData(t *testing.T) {
	var updateCalled, drawCalled int
	gameSystem := new(mocks.System)
	updateCallback := func(timeDelta float32) (bool, error) {
		updateCalled++
		return true, nil
	}
	drawCallback := func(timeDelta float32) error {
		drawCalled++
		return nil
	}
	loopData := createLoopData(gameSystem, updateCallback, drawCallback)
	if loopData.gameSystem != gameSystem {
		t.Error("loopData.gameSystem not set correctly")
	}
	if loopData.update == nil {
		t.Error("loopData.update is nil")
	} else {
		_, _ = loopData.update(0)
		if updateCalled != 1 {
			t.Error("loopData.update not set correctly")
		}
	}
	if loopData.draw == nil {
		t.Error("loopData.draw is nil")
	} else {
		_ = loopData.draw(0)
		if drawCalled != 1 {
			t.Error("loopData.draw not set correctly")
		}
	}
}

func TestRunner_update_UpdatesGameSystem(t *testing.T) {
	updateCallBack := func(timeDelta float32) (bool, error) {
		return true, nil
	}
	gameSystem := new(mocks.System)
	_, err := update(updateCallBack, gameSystem, 0)
	if err != nil {
		t.Fatalf("update returned error %v", err)
	}
	got := gameSystem.CallCntUpdate
	want := 1
	if got != want {
		t.Errorf("gameSystem.Update() was called incorrenct number of times. Got %v, want %v", got, want)
	}
}

func TestRunner_update_ReturnsErrorOfGameSystemUpdate(t *testing.T) {
	updateCallBack := func(timeDelta float32) (bool, error) {
		return true, nil
	}
	gameSystem := new(mocks.System)
	gameSystem.OnUpdate = func() error {
		return errors.New("test")
	}
	keepRunning, err := update(updateCallBack, gameSystem, 0)
	if err == nil {
		t.Error("update did not return expected error")
		return
	}
	if keepRunning {
		t.Error("update returned keepRunning == true on error")
	}
}

func TestRunner_update_UpdateCallbackIsCalled(t *testing.T) {
	var got, want float32
	updateCallBack := func(timeDelta float32) (bool, error) {
		got = timeDelta
		return true, nil
	}
	gameSystem := new(mocks.System)
	want = 12.34
	keepRunning, err := update(updateCallBack, gameSystem, want)
	if err != nil {
		t.Fatalf("update returned error %v", err)
	}
	if !keepRunning {
		t.Fatalf("update returned keepRunning == false unexpectedly")
	}
	if got != want {
		t.Errorf("updateCallback was not called or timeDelta was not passed correctly. Got %v, want %v", got, want)
	}
}

func TestRunner_update_ReturnsErrorOfUpdateCallback(t *testing.T) {
	updateCallBack := func(timeDelta float32) (bool, error) {
		return false, errors.New("test")
	}
	gameSystem := new(mocks.System)
	keepRunning, err := update(updateCallBack, gameSystem, 0)
	if err == nil {
		t.Error("update did not return expected error")
		return
	}
	if keepRunning {
		t.Error("update returned keepRunning == true on error")
	}
}
