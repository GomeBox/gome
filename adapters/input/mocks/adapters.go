package mocks

import "github.com/GomeBox/gome/adapters/input"

type Adapters struct {
	OnKeyboard      func() input.KeyboardAdapter
	CallCntKeyboard int
}

func (a *Adapters) Keyboard() input.KeyboardAdapter {
	a.CallCntKeyboard++
	if a.OnKeyboard != nil {
		return a.OnKeyboard()
	}
	return nil
}
