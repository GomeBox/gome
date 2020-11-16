package audio

import (
	"github.com/GomeBox/gome/adapters/audio/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSong_Play(t *testing.T) {
	playCalled := false
	adapterMock := &mocks.Song{OnPlay: func() error {
		playCalled = true
		return nil
	}}
	song := song{songPlayer: adapterMock}
	_ = song.Play()
	assert.True(t, playCalled)
}
