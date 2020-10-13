package gome

import (
	"github.com/GomeBox/gome/adapters"
)

type Game interface {
	Initialize() (adapters.System, error)
	Update(context Context) error
	Draw(context Context) error
}
