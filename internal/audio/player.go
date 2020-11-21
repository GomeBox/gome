package audio

import "github.com/GomeBox/gome/internal/shared"

func newPlayer(adapter playerAdapter) *player {
	p := new(player)
	p.adapter = adapter
	p.Unloader.Adapter = adapter
	return p
}

type player struct {
	adapter playerAdapter
	shared.Unloader
}

func (p *player) Play() error {
	return p.adapter.Play()
}
