package input

//Keyboard gives access to the computers keyboard
type Keyboard interface {
	KeyPressed(keyType KeyType) (bool, error)
}

//KeyType defines a keyboards key
type KeyType int

//goland:noinspection GoUnusedConst,GoUnusedConst,GoUnusedConst,GoUnusedConst,GoUnusedConst,GoUnusedConst,GoUnusedConst,GoUnusedConst,GoUnusedConst,GoUnusedConst,GoUnusedConst,GoUnusedConst,GoUnusedConst,GoUnusedConst,GoUnusedConst,GoUnusedConst,GoUnusedConst,GoUnusedConst,GoUnusedConst,GoUnusedConst,GoUnusedConst,GoUnusedConst,GoUnusedConst,GoUnusedConst,GoUnusedConst,GoUnusedConst,GoUnusedConst,GoUnusedConst,GoUnusedConst,GoUnusedConst,GoUnusedConst,GoUnusedConst,GoUnusedConst,GoUnusedConst,GoUnusedConst,GoUnusedConst,GoUnusedConst,GoUnusedConst,GoUnusedConst,GoUnusedConst,GoUnusedConst,GoUnusedConst,GoUnusedConst,GoUnusedConst,GoUnusedConst
const (
	KeyA KeyType = iota
	KeyB
	KeyC
	KeyD
	KeyE
	KeyF
	KeyG
	KeyH
	KeyI
	KeyJ
	KeyK
	KeyL
	KeyM
	KeyN
	KeyO
	KeyP
	KeyQ
	KeyR
	KeyS
	KeyT
	KeyU
	KeyV
	KeyW
	KeyX
	KeyY
	KeyZ
	KeyUp
	KeyDown
	KeyLeft
	KeyRight
	Key0
	Key1
	Key2
	Key3
	Key4
	Key5
	Key6
	Key7
	Key8
	Key9
	KeyCtrl
	KeyAlt
	KeySpace
	KeyReturn
	KeyEsc
)
