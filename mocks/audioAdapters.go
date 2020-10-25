package mocks

import "github.com/GomeBox/gome/adapters/audio"

type AudioAdapters struct {
}

func (adapters *AudioAdapters) SoundLoader() audio.SoundLoader {
	return nil
}

func (adapters *AudioAdapters) SongLoader() audio.SongLoader {
	return nil
}
