package tests

import (
	adapters "github.com/GomeBox/gome/adapters/audio"
	"github.com/GomeBox/gome/adapters/audio/mocks"
	"github.com/GomeBox/gome/audio"
	internalAudio "github.com/GomeBox/gome/internal/audio"
)

func createMockSystem(soundMock *mocks.Sound, songMock *mocks.Song) audio.System {
	soundLoaderMock := mocks.SoundLoader{OnLoad: func(fileName string) (adapters.Sound, error) {
		return soundMock, nil
	}}
	songLoaderMock := mocks.SongLoader{OnLoad: func(fileName string) (adapters.Song, error) {
		return songMock, nil
	}}
	a := internalAudio.Adapters{
		SoundLoader: soundLoaderMock,
		SongLoader:  songLoaderMock,
	}
	return internalAudio.NewSystem(a)
}
