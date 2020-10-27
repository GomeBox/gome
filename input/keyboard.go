package input

import (
	"github.com/GomeBox/gome/primitives"
)

//Keyboard contains functions to access the computer's keyboard
type Keyboard interface {
	//RegisterKey registers a key to track it's state and returns a Key instance
	RegisterKey(keyType primitives.KeyType) (Key, error)
	//UnregisterKey unregisters a key. The key-state will not be updated anymore
	UnregisterKey(keyType primitives.KeyType)
}
