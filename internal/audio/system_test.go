package audio

import (
	"errors"
	adapters "github.com/GomeBox/gome/adapters/audio"
	"github.com/GomeBox/gome/adapters/audio/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSystem_LoadSong(t *testing.T) {
	got := ""
	want := "test.file"
	retErr := false
	songLoaderMock := mocks.SongLoader{OnLoad: func(fileName string) (adapters.Song, error) {
		got = fileName
		if retErr {
			return nil, errors.New("test")
		}
		return nil, nil
	}}
	sys := system{
		songLoader: songLoaderMock,
	}
	_, err := sys.LoadSong(want)
	assert.NoError(t, err)
	assert.Equal(t, want, got)
	retErr = true
	_, err = sys.LoadSong(want)
	assert.Error(t, err)
}

func TestSystem_LoadSound(t *testing.T) {
	got := ""
	want := "test.file"
	retErr := false
	soundLoaderMock := mocks.SoundLoader{OnLoad: func(fileName string) (adapters.Sound, error) {
		got = fileName
		if retErr {
			return nil, errors.New("test")
		}
		return nil, nil
	}}
	sys := system{
		soundLoader: soundLoaderMock,
	}
	_, err := sys.LoadSound(want)
	assert.NoError(t, err)
	assert.Equal(t, want, got)
	retErr = true
	_, err = sys.LoadSound(want)
	assert.Error(t, err)
}
