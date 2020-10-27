package audio

import (
	adapters "github.com/GomeBox/gome/adapters/audio"
	"github.com/GomeBox/gome/adapters/audio/mocks"
	"github.com/GomeBox/gome/audio"
	"testing"
)

func TestSystem_LoadSong_PassesFilename(t *testing.T) {
	got := ""
	want := "test.file"
	songLoaderMock := mocks.SongLoader{OnLoad: func(fileName string) (adapters.Song, error) {
		got = fileName
		return nil, nil
	}}
	a := Adapters{
		SoundLoader: nil,
		SongLoader:  songLoaderMock,
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
	a := Adapters{
		SoundLoader: soundLoaderMock,
		SongLoader:  nil,
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

func createMockSystem(soundMock *mocks.Sound, songMock *mocks.Song) audio.System {
	soundLoaderMock := mocks.SoundLoader{OnLoad: func(fileName string) (adapters.Sound, error) {
		return soundMock, nil
	}}
	songLoaderMock := mocks.SongLoader{OnLoad: func(fileName string) (adapters.Song, error) {
		return songMock, nil
	}}
	a := Adapters{
		SoundLoader: soundLoaderMock,
		SongLoader:  songLoaderMock,
	}
	return NewSystem(a)
}
