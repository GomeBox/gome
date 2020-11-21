package audio

import (
	"errors"
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
	song := player{adapter: adapterMock}
	_ = song.Play()
	assert.True(t, playCalled)
}

func TestSong_Unload(t *testing.T) {
	unloadCalled := false
	retErr := false
	adapterMock := &mocks.Song{
		OnUnload: func() error {
			unloadCalled = true
			if retErr {
				return errors.New("test")
			}
			return nil
		},
	}
	song := newPlayer(adapterMock)
	err := song.Unload()
	assert.NoError(t, err)
	assert.True(t, unloadCalled)
	retErr = true
	err = song.Unload()
	assert.Error(t, err)
	retErr = false
	err = song.Unload()
	assert.Error(t, err, "already unloaded. Should return error")
}

func TestSong_Unloaded(t *testing.T) {
	retErr := false
	adapterMock := &mocks.Song{
		OnUnload: func() error {
			if retErr {
				return errors.New("test")
			}
			return nil
		},
	}
	s := newPlayer(adapterMock)
	assert.False(t, s.Unloaded())
	_ = s.Unload()
	assert.True(t, s.Unloaded())

	s = newPlayer(adapterMock)
	retErr = true
	_ = s.Unload()
	assert.False(t, s.Unloaded())
}
