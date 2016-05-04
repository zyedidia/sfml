package sfml

//#include <SFML/Window.h>
import "C"

const (
	SensorAccelerometer    SensorType = C.sfSensorAccelerometer
	SensorGyroscope        SensorType = C.sfSensorGyroscope
	SensorMagnetometer     SensorType = C.sfSensorMagnetometer
	SensorGravity          SensorType = C.sfSensorGravity
	SensorUserAcceleration SensorType = C.sfSensorUserAcceleration
	SensorOrientation      SensorType = C.sfSensorOrientation
	SensorCount            SensorType = C.sfSensorCount
)

type SensorType int

func SensorIsAvailable(sensor SensorType) bool {
	return goBool(C.sfSensor_isAvailable(C.sfSensorType(sensor)))
}

func SensorSetEnabled(sensor SensorType, enabled bool) {
	C.sfSensor_setEnabled(C.sfSensorType(sensor), cBool(enabled))
}

func SensorGetValue(sensor SensorType) Vector3f {
	v := C.sfSensor_getValue(C.sfSensorType(sensor))
	return *goVector3f(&v)
}
