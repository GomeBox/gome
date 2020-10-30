package game

import (
	audio "github.com/GomeBox/gome/internal/audio/interfaces"
	"github.com/GomeBox/gome/internal/game/mocks"
	graphics "github.com/GomeBox/gome/internal/graphics/interfaces"
	input "github.com/GomeBox/gome/internal/input/interfaces"
	"testing"
)

func TestContext_Graphics(t *testing.T) {
	systemMock := new(mocks.System)
	graphicsCalled := false
	systemMock.OnGraphics = func() graphics.System {
		graphicsCalled = true
		return nil
	}
	context := newContext(systemMock)
	_ = context.Graphics()
	if !graphicsCalled {
		t.Error("context.Graphics() does not call system.Graphics()")
	}
}

func TestContext_Audio(t *testing.T) {
	systemMock := new(mocks.System)
	audioCalled := false
	systemMock.OnAudio = func() audio.System {
		audioCalled = true
		return nil
	}
	context := newContext(systemMock)
	_ = context.Audio()
	if !audioCalled {
		t.Error("context.Audio() does not call system.Audio()")
	}
}

func TestContext_Input(t *testing.T) {
	systemMock := new(mocks.System)
	inputCalled := false
	systemMock.OnInput = func() input.System {
		inputCalled = true
		return nil
	}
	context := newContext(systemMock)
	_ = context.Input()
	if !inputCalled {
		t.Error("context.Input() does not call system.Input()")
	}
}
