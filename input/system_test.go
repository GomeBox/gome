package input

import (
	"github.com/GomeBox/gome/internal/input/interfaces"
	"github.com/GomeBox/gome/internal/input/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSystem(t *testing.T) {
	sysMock := new(mocks.System)
	sys := NewSystem(sysMock)
	assert.NotNil(t, sys)
	assert.Same(t, sysMock, sys.(*system).internalSystem)
}

func TestSystem_Keyboard(t *testing.T) {
	sysMock := new(mocks.System)
	kbMock := new(mocks.Keyboard)
	sysMock.OnKeyboard = func() interfaces.Keyboard {
		return kbMock
	}
	sys := system{internalSystem: sysMock}
	kb := sys.Keyboard()
	assert.NotNil(t, kb)
	assert.Same(t, kbMock, kb.(*keyboard).internalKeyboard)
}
