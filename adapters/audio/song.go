package audio

import (
	"github.com/GomeBox/gome/adapters/shared"
)

type Song interface {
	shared.Unloader
	Play() error
}
