package audio

import (
	"github.com/GomeBox/gome/adapters/audio/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSound_Play(t *testing.T) {
	playCalled := false
	adapterMock := &mocks.Sound{OnPlay: func() error {
		playCalled = true
		return nil
	}}
	sound := sound{soundPlayer: adapterMock}
	_ = sound.Play()
	assert.True(t, playCalled)
}
