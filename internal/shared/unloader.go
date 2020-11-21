package shared

import (
	"errors"
	"github.com/GomeBox/gome/adapters/shared"
)

type Unloader struct {
	Adapter  shared.Unloader
	unloaded bool
}

func (u *Unloader) Unload() error {
	if u.unloaded {
		return errors.New("resource already unloaded")
	}
	err := u.Adapter.Unload()
	if err != nil {
		return err
	}
	u.unloaded = true
	return nil
}

func (u *Unloader) Unloaded() bool {
	return u.unloaded
}
