package audio

import (
	"errors"
	"github.com/GomeBox/gome/internal/audio/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlayer_Play(t *testing.T) {
	retErr := false
	playerMock := &mocks.Player{
		OnPlay: func() error {
			if retErr {
				return errors.New("test")
			}
			return nil
		},
	}
	player := player{internalPlayer: playerMock}
	err := player.Play()
	assert.NoError(t, err)
	assert.Equal(t, 1, playerMock.CallCntPlay, "Play was not called expected number of times")
	retErr = true
	err = player.Play()
	assert.Error(t, err)
}
