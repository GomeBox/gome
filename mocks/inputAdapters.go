package mocks

import "github.com/GomeBox/gome/adapters/input"

type InputAdapters struct {
}

func (input *InputAdapters) Keyboard() input.KeyboardAdapter {
	return nil
}
