package sfml

//#include <SFML/Audio.h>
//#include <stdlib.h>
import "C"

import (
	"runtime"
)

type Sound struct {
	data *C.sfSound
}

func destroySound(s *Sound) {
	C.sfSound_destroy(s.data)
}

func NewEmptySound() *Sound {
	f := C.sfSound_create()
	if f == nil {
		return nil
	}
	obj := &Sound{f}
	runtime.SetFinalizer(obj, destroySound)
	return obj
}

func NewSound(buffer *SoundBuffer) *Sound {
	f := C.sfSound_create()
	if f == nil {
		return nil
	}
	C.sfSound_setBuffer(f, buffer.data)
	obj := &Sound{f}
	runtime.SetFinalizer(obj, destroySound)
	return obj
}

func (s *Sound) Copy() *Sound {
	c := C.sfSound_copy(s.data)
	obj := &Sound{c}
	runtime.SetFinalizer(obj, destroySound)
	return obj
}

func (s *Sound) SetBuffer(buffer *SoundBuffer) {
	C.sfSound_setBuffer(s.data, buffer.data)
}

func (s *Sound) GetBuffer() *SoundBuffer {
	return &SoundBuffer{C.sfSound_getBuffer(s.data)}
}

func (s *Sound) SetLoop(loop bool) {
	C.sfSound_setLoop(s.data, cBool(loop))
}

func (s *Sound) GetLoop() bool {
	return goBool(C.sfSound_getLoop(s.data))
}

func (s *Sound) Play() {
	C.sfSound_play(s.data)
}

func (s *Sound) Pause() {
	C.sfSound_play(s.data)
}

func (s *Sound) Stop() {
	C.sfSound_play(s.data)
}

func (s *Sound) GetStatus() SoundStatus {
	return SoundStatus(C.sfSound_getStatus(s.data))
}

func (s *Sound) GetPlayingOffset() int64 {
	return int64(C.sfSound_getPlayingOffset(s.data).microseconds)
}

func (s *Sound) SetPitch(pitch float32) {
	C.sfSound_setPitch(s.data, C.float(pitch))
}

func (s *Sound) SetVolume(volume float32) {
	C.sfSound_setVolume(s.data, C.float(volume))
}

func (s *Sound) SetPosition(position Vector3f) {
	C.sfSound_setPosition(s.data, cVector3f(&position))
}

func (s *Sound) SetRelativeToListener(relative bool) {
	C.sfSound_setRelativeToListener(s.data, cBool(relative))
}

func (s *Sound) SetMinDistance(distance float32) {
	C.sfSound_setMinDistance(s.data, C.float(distance))
}

func (s *Sound) SetAttenuation(attenuation float32) {
	C.sfSound_setAttenuation(s.data, C.float(attenuation))
}

func (s *Sound) SetPlayingOffset(playingOffset int64) {
	var t C.sfTime
	t.microseconds = C.sfInt64(playingOffset)
	C.sfSound_setPlayingOffset(s.data, t)
}

func (s *Sound) GetPitch() float32 {
	return float32(C.sfSound_getPitch(s.data))
}

func (s *Sound) GetVolume() float32 {
	return float32(C.sfSound_getVolume(s.data))
}

func (s *Sound) GetPosition() Vector3f {
	p := C.sfSound_getPosition(s.data)
	return *goVector3f(&p)
}

func (s *Sound) IsRelativeToListener() bool {
	return goBool(C.sfSound_isRelativeToListener(s.data))
}

func (s *Sound) GetMinDistance() float32 {
	return float32(C.sfSound_getMinDistance(s.data))
}

func (s *Sound) GetAttenuation() float32 {
	return float32(C.sfSound_getAttenuation(s.data))
}
