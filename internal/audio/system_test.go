package audio

import (
	adapters "github.com/GomeBox/gome/adapters/audio"
	"github.com/GomeBox/gome/adapters/audio/mocks"
	"testing"
)

func TestSystem_LoadSong_PassesFilename(t *testing.T) {
	got := ""
	want := "test.file"
	songLoaderMock := mocks.SongLoader{OnLoad: func(fileName string) (adapters.Song, error) {
		got = fileName
		return nil, nil
	}}
	a := &mocks.Adapters{
		OnSongLoader: func() adapters.SongLoader {
			return songLoaderMock
		},
	}
	system := NewSystem(a)
	_, err := system.LoadSong(want)
	if err != nil {
		t.Errorf("expected err to be nil but was %v", err)
		return
	}
	if got != want {
		t.Errorf("fileName was not passed to songLoader. got %v, want %v", got, want)
		return
	}
}

func TestSystem_LoadSound_PassesFilename(t *testing.T) {
	got := ""
	want := "test.file"
	soundLoaderMock := mocks.SoundLoader{OnLoad: func(fileName string) (adapters.Sound, error) {
		got = fileName
		return nil, nil
	}}
	a := &mocks.Adapters{
		OnSoundLoader: func() adapters.SoundLoader {
			return soundLoaderMock
		},
	}
	system := NewSystem(a)
	_, err := system.LoadSound(want)
	if err != nil {
		t.Errorf("expected err to be nil but was %v", err)
		return
	}
	if got != want {
		t.Errorf("fileName was not passed to soundLoader. got %v, want %v", got, want)
		return
	}
}
