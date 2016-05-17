package sfml

//#include <SFML/Audio.h>
import "C"

type SoundBuffer struct {
	data *C.sfSoundBuffer
}
