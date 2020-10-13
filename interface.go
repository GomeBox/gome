package gome

import (
	"github.com/GomeBox/gome/adapters"
)

type Interface interface {
	Initialize() (adapters.System, error)
	Update(c Context) error
	Draw(c Context) error
}
