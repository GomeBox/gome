package interfaces

import (
	"github.com/GomeBox/gome/primitives"
)

type Keyboard interface {
	RegisterKey(keyType primitives.KeyType) (Key, error)
	UnregisterKey(keyType primitives.KeyType)
}
