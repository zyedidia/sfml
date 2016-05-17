package sfml

//#include <SFML/Audio.h>
//#include <stdlib.h>
//#include "soundBuffer.h"
import "C"

import (
	"runtime"
	"unsafe"
)

type SoundBuffer struct {
	data *C.sfSoundBuffer
}

func destroySoundBuffer(m *SoundBuffer) {
	C.sfSoundBuffer_destroy(m.data)
}

func NewSoundBuffer(filename string) *SoundBuffer {
	file := C.CString(filename)
	defer C.free(unsafe.Pointer(file))
	f := C.sfSoundBuffer_createFromFile(file)
	if f == nil {
		return nil
	}
	obj := &SoundBuffer{f}
	runtime.SetFinalizer(obj, destroySoundBuffer)
	return obj
}

func NewSoundBufferFromMemory(data []byte) *SoundBuffer {
	f := C.sfSoundBuffer_createFromMemory(unsafe.Pointer(&data[0]), C.size_t(len(data)))
	if f == nil {
		return nil
	}
	obj := &SoundBuffer{f}
	runtime.SetFinalizer(obj, destroySoundBuffer)
	return obj
}

func NewSoundBufferFromSamples(samples []int16, channelCount uint, sampleRate uint) *SoundBuffer {
	f := C.sfSoundBuffer_createFromSamples((*C.sfInt16)(&samples[0]), C.sfUint64(len(samples)), C.uint(channelCount), C.uint(sampleRate))
	if f == nil {
		return nil
	}
	obj := &SoundBuffer{f}
	runtime.SetFinalizer(obj, destroySoundBuffer)
	return obj
}

func (s *SoundBuffer) Copy() *SoundBuffer {
	c := C.sfSoundBuffer_copy(s.data)
	obj := &SoundBuffer{c}
	runtime.SetFinalizer(obj, destroySoundBuffer)
	return obj
}

func (s *SoundBuffer) SaveToFile(filename string) {
	file := C.CString(filename)
	defer C.free(unsafe.Pointer(file))
	C.sfSoundBuffer_saveToFile(s.data, file)
}

func (s *SoundBuffer) GetSamples() []int16 {
	count := uint64(C.sfSoundBuffer_getSampleCount(s.data))
	ptr := C.sfSoundBuffer_getSamples(s.data)
	var samples []int16
	for i := uint64(0); i < count; i++ {
		samples = append(samples, int16(*ptr))
		C.nextSample(&ptr)
	}
	return samples
}

func (s *SoundBuffer) GetSampleCount() uint64 {
	return uint64(C.sfSoundBuffer_getSampleCount(s.data))
}

func (s *SoundBuffer) GetSampleRate() uint {
	return uint(C.sfSoundBuffer_getSampleRate(s.data))
}

func (s *SoundBuffer) GetChannelCount() uint {
	return uint(C.sfSoundBuffer_getChannelCount(s.data))
}

func (s *SoundBuffer) GetDuration() int64 {
	return int64(C.sfSoundBuffer_getDuration(s.data).microseconds)
}
