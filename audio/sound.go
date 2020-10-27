package audio

//Sound is used to Play a sound effect
type Sound interface {
	//Play plays the sound effect once
	Play() error
}
