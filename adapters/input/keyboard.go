package input

type Keyboard interface {
	KeyPressed(keyType KeyType) (bool, error)
}

type KeyType int

const (
	Key_A KeyType = iota
	Key_B
	Key_C
	Key_D
	Key_E
	Key_F
	Key_G
	Key_H
	Key_I
	Key_J
	Key_K
	Key_L
	Key_M
	Key_N
	Key_O
	Key_P
	Key_Q
	Key_R
	Key_S
	Key_T
	Key_U
	Key_V
	Key_W
	Key_X
	Key_Y
	Key_Z
	Key_Up
	Key_Down
	Key_Left
	Key_Right
	Key_0
	Key_1
	Key_2
	Key_3
	Key_4
	Key_5
	Key_6
	Key_7
	Key_8
	Key_9
	Key_Ctrl
	Key_Alt
	Key_Space
	Key_Return
	Key_Esc
)
