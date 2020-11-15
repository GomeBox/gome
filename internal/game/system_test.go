package game

import (
	"testing"
)

func TestSystem_NewSystem(t *testing.T) {
	//adapters := new(adapterMocks.System)
	//systemsFactory := new(mocks.SystemsFactory)
	//System := NewSystem(adapters, systemsFactory)
	//if System == nil {
	//	t.Fatal("NewSystem() returned nil")
	//}
	//want := 1
	//got := systemsFactory.CallCntCreateGraphicsSystem
	//if got != want {
	//	t.Errorf("CreateGraphicsSystem was not called expected number of times. Got: %d, want: %d", got, want)
	//}
	//got = systemsFactory.CallCntCreateInputSystem
	//if got != want {
	//	t.Errorf("CreateInputSystem was not called expected number of times. Got: %d, want: %d", got, want)
	//}
	//got = systemsFactory.CallCntCreateAudioSystem
	//if got != want {
	//	t.Errorf("CreateAudioSystem was not called expected number of times. Got: %d, want: %d", got, want)
	//}
}

func TestSystem_Initialize(t *testing.T) {
	//adapters := new(adapterMocks.System)
	//systemsFactory := new(mocks.SystemsFactory)
	//System := NewSystem(adapters, systemsFactory)
	//_ = System.Initialize()
	//want := 1
	//got := adapters.CallCntInitialize
	//if got != want {
	//	t.Errorf("adapterSystem.Initialize was not called expected number of times. Got: %d, wawnt: %d", got, want)
	//}
}

func TestSystem_Update(t *testing.T) {
	//adapters := new(adapterMocks.System)
	//systemsFactory := new(mocks.SystemsFactory)
	//inputSystem := new(inputMocks.System)
	//systemsFactory.OnCreateInputSystem = func() input.System {
	//	return inputSystem
	//}
	//System := NewSystem(adapters, systemsFactory)
	//_ = System.Update()
	//want := 1
	//got := adapters.CallCntUpdate
	//if got != want {
	//	t.Errorf("adapterSystem.Update was not called expected number of times. Got: %d, want: %d", got, want)
	//}
	//got = inputSystem.CallCntUpdate
	//if got != want {
	//	t.Errorf("inputSystem.Update was not called expected number of times. Got: %d, want: %d", got, want)
	//}
}

func TestSystem_UpdateErrOnAdapterSystemErr(t *testing.T) {
	//adapters := new(adapterMocks.System)
	//adapters.OnUpdate = func() error {
	//	return errors.New("test")
	//}
	//systemsFactory := new(mocks.SystemsFactory)
	//System := NewSystem(adapters, systemsFactory)
	//err := System.Update()
	//if err == nil {
	//	t.Error("Update did not return expected error")
	//}
}

func TestSystem_Input(t *testing.T) {
	//adapters := new(adapterMocks.System)
	//systemsFactory := new(mocks.SystemsFactory)
	//inputSystem := new(inputMocks.System)
	//systemsFactory.OnCreateInputSystem = func() input.System {
	//	return inputSystem
	//}
	//System := NewSystem(adapters, systemsFactory)
	//if System.Input() != inputSystem {
	//	t.Error("Input() does not return created input System")
	//}
}

func TestSystem_Audio(t *testing.T) {
	//adapters := new(adapterMocks.Graphics)
	//systemsFactory := new(mocks.SystemsFactory)
	//audioSystem := new(audioMocks.Graphics)
	//systemsFactory.OnCreateAudioSystem = func() audio.Graphics {
	//	return audioSystem
	//}
	//System := NewSystem(adapters, systemsFactory)
	//if System.Audio() != audioSystem {
	//	t.Error("Audio() does not return created input System")
	//}
}

func TestSystem_Graphics(t *testing.T) {
	//adapters := new(adapterMocks.System)
	//systemsFactory := new(mocks.SystemsFactory)
	//graphicsSystem := new(graphicsMocks.System)
	//systemsFactory.OnCreateGraphicsSystem = func() graphics.System {
	//	return graphicsSystem
	//}
	//System := NewSystem(adapters, systemsFactory)
	//if System.Graphics() != graphicsSystem {
	//	t.Error("Audio() does not return created input System")
	//}
}
