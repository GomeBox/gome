package audio

import (
	"errors"
	"github.com/GomeBox/gome/adapters/audio/mocks"
	"testing"
)

func TestSound_Play(t *testing.T) {
	playCalled := false
	soundMock := mocks.Sound{OnPlay: func() error {
		playCalled = true
		return nil
	}}
	system := createMockSystem(&soundMock, nil)
	sound, err := system.LoadSound("test")
	if err != nil {
		t.Errorf("err was expected to be nil but was %q", err)
		return
	}
	err = sound.Play()
	if err != nil {
		t.Errorf("err was expected to be nil but was %q", err)
		return
	}
	if !playCalled {
		t.Error("Play was not called")
	}
}

func TestSound_PlayReturnsError(t *testing.T) {
	soundMock := mocks.Sound{OnPlay: func() error {
		return errors.New("test")
	}}
	system := createMockSystem(&soundMock, nil)
	sound, err := system.LoadSound("test")
	if err != nil {
		t.Errorf("err was expected to be nil but was %q", err)
		return
	}
	err = sound.Play()
	if err == nil {
		t.Errorf("play did not return error")
		return
	}
}
