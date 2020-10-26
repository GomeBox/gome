package tests

import (
	"errors"
	"github.com/GomeBox/gome/adapters/audio/mocks"
	"testing"
)

func TestSong_Play(t *testing.T) {
	playCalled := false
	songMock := mocks.Song{OnPlay: func() error {
		playCalled = true
		return nil
	}}
	system := createMockSystem(nil, &songMock)
	song, err := system.LoadSong("test")
	if err != nil {
		t.Errorf("err was expected to be nil but was %q", err)
		return
	}
	err = song.Play()
	if err != nil {
		t.Errorf("err was expected to be nil but was %q", err)
		return
	}
	if !playCalled {
		t.Error("Play was not called")
	}
}

func TestSong_PlayReturnsError(t *testing.T) {
	songMock := mocks.Song{OnPlay: func() error {
		return errors.New("test")
	}}
	system := createMockSystem(nil, &songMock)
	song, err := system.LoadSong("test")
	if err != nil {
		t.Errorf("err was expected to be nil but was %q", err)
		return
	}
	err = song.Play()
	if err == nil {
		t.Errorf("play did not return error")
		return
	}
}
