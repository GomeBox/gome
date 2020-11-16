package game

import (
	gomeInterfaces "github.com/GomeBox/gome/interfaces"
	"github.com/GomeBox/gome/internal/game/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContext_Graphics(t *testing.T) {
	graphics := new(mocks.Graphics)
	context := createContext(graphics, nil, nil)
	assert.Same(t, graphics, context.Graphics())
}

func TestContext_Audio(t *testing.T) {
	audio := new(mocks.Audio)
	context := createContext(nil, nil, audio)
	assert.Same(t, audio, context.Audio())
}

func TestContext_Input(t *testing.T) {
	input := new(mocks.Input)
	context := createContext(nil, input, nil)
	assert.Same(t, input, context.Input())
}

func createContext(
	graphics gomeInterfaces.Graphics,
	input gomeInterfaces.Input,
	audio gomeInterfaces.Audio) gomeInterfaces.Context {
	return newContext(graphics, input, audio)
}
