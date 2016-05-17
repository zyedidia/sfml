package sfml

//#include <SFML/Audio.h>
import "C"

const (
	Stopped SoundStatus = C.sfStopped
	Paused  SoundStatus = C.sfPaused
	Playing SoundStatus = C.sfPlaying
)

type SoundStatus int
