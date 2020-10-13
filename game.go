package gome

import (
	"github.com/GomeBox/gome/adapters"
)

type Game interface {
	Initialize() (adapters.System, error)
	Update(c Context) error
	Draw(c Context) error
}
