package mocks

import "github.com/GomeBox/gome/adapters/audio"

type Adapters struct {
	OnSoundLoader                         func() audio.SoundLoader
	OnSongLoader                          func() audio.SongLoader
	CallCntSoundLoader, CallCntSongLoader int
}

func (a *Adapters) SoundLoader() audio.SoundLoader {
	a.CallCntSoundLoader++
	if a.OnSoundLoader != nil {
		return a.OnSoundLoader()
	}
	return nil
}

func (a *Adapters) SongLoader() audio.SongLoader {
	a.CallCntSongLoader++
	if a.OnSongLoader != nil {
		return a.OnSongLoader()
	}
	return nil
}
