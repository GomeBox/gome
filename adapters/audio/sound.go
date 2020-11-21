package audio

import (
	"github.com/GomeBox/gome/adapters/shared"
)

type Sound interface {
	shared.Unloader
	Play() error
}
