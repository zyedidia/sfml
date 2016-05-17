package sfml

//#include <SFML/Audio.h>
//#include <stdlib.h>
import "C"

import (
	"runtime"
	"unsafe"
)

type Music struct {
	data *C.sfMusic
}

func destroyMusic(m *Music) {
	C.sfMusic_destroy(m.data)
}

func NewMusic(filename string) *Music {
	file := C.CString(filename)
	defer C.free(unsafe.Pointer(file))
	f := C.sfMusic_createFromFile(file)
	if f == nil {
		return nil
	}
	obj := &Music{f}
	runtime.SetFinalizer(obj, destroyMusic)
	return obj
}

func NewMusicFromMemory(data []byte) *Music {
	f := C.sfMusic_createFromMemory(unsafe.Pointer(&data[0]), C.size_t(len(data)))
	if f == nil {
		return nil
	}
	obj := &Music{f}
	runtime.SetFinalizer(obj, destroyMusic)
	return obj
}

func (m *Music) SetLoop(loop bool) {
	C.sfMusic_setLoop(m.data, cBool(loop))
}

func (m *Music) GetLoop() bool {
	return goBool(C.sfMusic_getLoop(m.data))
}

func (m *Music) GetDuration() int64 {
	return int64(C.sfMusic_getDuration(m.data).microseconds)
}

func (m *Music) Play() {
	C.sfMusic_play(m.data)
}

func (m *Music) Pause() {
	C.sfMusic_play(m.data)
}

func (m *Music) Stop() {
	C.sfMusic_play(m.data)
}

func (m *Music) GetChannelCount() uint {
	return uint(C.sfMusic_getChannelCount(m.data))
}

func (m *Music) GetSampleRate() int64 {
	return int64(C.sfMusic_getSampleRate(m.data))
}

func (m *Music) GetStatus() SoundStatus {
	return SoundStatus(C.sfMusic_getStatus(m.data))
}

func (m *Music) GetPlayingOffset() int64 {
	return int64(C.sfMusic_getPlayingOffset(m.data).microseconds)
}

func (m *Music) SetPitch(pitch float32) {
	C.sfMusic_setPitch(m.data, C.float(pitch))
}

func (m *Music) SetVolume(volume float32) {
	C.sfMusic_setVolume(m.data, C.float(volume))
}

func (m *Music) SetPosition(position Vector3f) {
	C.sfMusic_setPosition(m.data, cVector3f(&position))
}

func (m *Music) SetRelativeToListener(relative bool) {
	C.sfMusic_setRelativeToListener(m.data, cBool(relative))
}

func (m *Music) SetMinDistance(distance float32) {
	C.sfMusic_setMinDistance(m.data, C.float(distance))
}

func (m *Music) SetAttenuation(attenuation float32) {
	C.sfMusic_setAttenuation(m.data, C.float(attenuation))
}

func (m *Music) SetPlayingOffset(playingOffset int64) {
	var t C.sfTime
	t.microseconds = C.sfInt64(playingOffset)
	C.sfMusic_setPlayingOffset(m.data, t)
}

func (m *Music) GetPitch() float32 {
	return float32(C.sfMusic_getPitch(m.data))
}

func (m *Music) GetVolume() float32 {
	return float32(C.sfMusic_getVolume(m.data))
}

func (m *Music) GetPosition() Vector3f {
	p := C.sfMusic_getPosition(m.data)
	return *goVector3f(&p)
}

func (m *Music) IsRelativeToListener() bool {
	return goBool(C.sfMusic_isRelativeToListener(m.data))
}

func (m *Music) GetMinDistance() float32 {
	return float32(C.sfMusic_getMinDistance(m.data))
}

func (m *Music) GetAttenuation() float32 {
	return float32(C.sfMusic_getAttenuation(m.data))
}
