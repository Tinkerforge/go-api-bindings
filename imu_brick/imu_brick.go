/* ***********************************************************
 * This file was automatically generated on 2019-11-25.      *
 *                                                           *
 * Go Bindings Version 2.0.5                                 *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


// Full fledged AHRS with 9 degrees of freedom.
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricks/IMU_Brick_Go.html.
package imu_brick

import (
	"encoding/binary"
	"bytes"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionGetAcceleration Function = 1
	FunctionGetMagneticField Function = 2
	FunctionGetAngularVelocity Function = 3
	FunctionGetAllData Function = 4
	FunctionGetOrientation Function = 5
	FunctionGetQuaternion Function = 6
	FunctionGetIMUTemperature Function = 7
	FunctionLedsOn Function = 8
	FunctionLedsOff Function = 9
	FunctionAreLedsOn Function = 10
	FunctionSetAccelerationRange Function = 11
	FunctionGetAccelerationRange Function = 12
	FunctionSetMagnetometerRange Function = 13
	FunctionGetMagnetometerRange Function = 14
	FunctionSetConvergenceSpeed Function = 15
	FunctionGetConvergenceSpeed Function = 16
	FunctionSetCalibration Function = 17
	FunctionGetCalibration Function = 18
	FunctionSetAccelerationPeriod Function = 19
	FunctionGetAccelerationPeriod Function = 20
	FunctionSetMagneticFieldPeriod Function = 21
	FunctionGetMagneticFieldPeriod Function = 22
	FunctionSetAngularVelocityPeriod Function = 23
	FunctionGetAngularVelocityPeriod Function = 24
	FunctionSetAllDataPeriod Function = 25
	FunctionGetAllDataPeriod Function = 26
	FunctionSetOrientationPeriod Function = 27
	FunctionGetOrientationPeriod Function = 28
	FunctionSetQuaternionPeriod Function = 29
	FunctionGetQuaternionPeriod Function = 30
	FunctionOrientationCalculationOn Function = 37
	FunctionOrientationCalculationOff Function = 38
	FunctionIsOrientationCalculationOn Function = 39
	FunctionSetSPITFPBaudrateConfig Function = 231
	FunctionGetSPITFPBaudrateConfig Function = 232
	FunctionGetSendTimeoutCount Function = 233
	FunctionSetSPITFPBaudrate Function = 234
	FunctionGetSPITFPBaudrate Function = 235
	FunctionGetSPITFPErrorCount Function = 237
	FunctionEnableStatusLED Function = 238
	FunctionDisableStatusLED Function = 239
	FunctionIsStatusLEDEnabled Function = 240
	FunctionGetProtocol1BrickletName Function = 241
	FunctionGetChipTemperature Function = 242
	FunctionReset Function = 243
	FunctionGetIdentity Function = 255
	FunctionCallbackAcceleration Function = 31
	FunctionCallbackMagneticField Function = 32
	FunctionCallbackAngularVelocity Function = 33
	FunctionCallbackAllData Function = 34
	FunctionCallbackOrientation Function = 35
	FunctionCallbackQuaternion Function = 36
)

type CalibrationType = uint8

const (
	CalibrationTypeAccelerometerGain CalibrationType = 0
	CalibrationTypeAccelerometerBias CalibrationType = 1
	CalibrationTypeMagnetometerGain CalibrationType = 2
	CalibrationTypeMagnetometerBias CalibrationType = 3
	CalibrationTypeGyroscopeGain CalibrationType = 4
	CalibrationTypeGyroscopeBias CalibrationType = 5
)

type CommunicationMethod = uint8

const (
	CommunicationMethodNone CommunicationMethod = 0
	CommunicationMethodUSB CommunicationMethod = 1
	CommunicationMethodSPIStack CommunicationMethod = 2
	CommunicationMethodChibi CommunicationMethod = 3
	CommunicationMethodRS485 CommunicationMethod = 4
	CommunicationMethodWIFI CommunicationMethod = 5
	CommunicationMethodEthernet CommunicationMethod = 6
	CommunicationMethodWIFIV2 CommunicationMethod = 7
)

type IMUBrick struct {
	device Device
}
const DeviceIdentifier = 16
const DeviceDisplayName = "IMU Brick"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (IMUBrick, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{ 2,0,4 }, uid, &internalIPCon, 0)
	if err != nil {
		return IMUBrick{}, err
	}
	dev.ResponseExpected[FunctionGetAcceleration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetMagneticField] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetAngularVelocity] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetAllData] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetOrientation] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetQuaternion] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetIMUTemperature] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionLedsOn] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionLedsOff] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionAreLedsOn] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetAccelerationRange] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetAccelerationRange] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetMagnetometerRange] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetMagnetometerRange] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetConvergenceSpeed] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetConvergenceSpeed] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetCalibration] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetCalibration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetAccelerationPeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetAccelerationPeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetMagneticFieldPeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetMagneticFieldPeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetAngularVelocityPeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetAngularVelocityPeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetAllDataPeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetAllDataPeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetOrientationPeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetOrientationPeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetQuaternionPeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetQuaternionPeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionOrientationCalculationOn] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionOrientationCalculationOff] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionIsOrientationCalculationOn] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetSPITFPBaudrateConfig] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetSPITFPBaudrateConfig] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetSendTimeoutCount] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetSPITFPBaudrate] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetSPITFPBaudrate] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetSPITFPErrorCount] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionEnableStatusLED] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionDisableStatusLED] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionIsStatusLEDEnabled] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetProtocol1BrickletName] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetChipTemperature] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionReset] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetIdentity] = ResponseExpectedFlagAlwaysTrue;
	return IMUBrick{dev}, nil
}

// Returns the response expected flag for the function specified by the function ID parameter.
// It is true if the function is expected to send a response, false otherwise.
//
// For getter functions this is enabled by default and cannot be disabled, because those
// functions will always send a response. For callback configuration functions it is enabled
// by default too, but can be disabled by SetResponseExpected.
// For setter functions it is disabled by default and can be enabled.
//
// Enabling the response expected flag for a setter function allows to detect timeouts
// and other error conditions calls of this setter as well. The device will then send a response
// for this purpose. If this flag is disabled for a setter function then no response is send
// and errors are silently ignored, because they cannot be detected.
//
// See SetResponseExpected for the list of function ID constants available for this function.
func (device *IMUBrick) GetResponseExpected(functionID Function) (bool, error) {
	return device.device.GetResponseExpected(uint8(functionID))
}

// Changes the response expected flag of the function specified by the function ID parameter.
// This flag can only be changed for setter (default value: false) and callback configuration
// functions (default value: true). For getter functions it is always enabled.
//
// Enabling the response expected flag for a setter function allows to detect timeouts and
// other error conditions calls of this setter as well. The device will then send a response
// for this purpose. If this flag is disabled for a setter function then no response is send
// and errors are silently ignored, because they cannot be detected.
func (device *IMUBrick) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *IMUBrick) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *IMUBrick) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// This callback is triggered periodically with the period that is set by
// SetAccelerationPeriod. The parameters are the acceleration
// for the x, y and z axis.
func (device *IMUBrick) RegisterAccelerationCallback(fn func(int16, int16, int16)) uint64 {
	wrapper := func(byteSlice []byte) {
		buf := bytes.NewBuffer(byteSlice[8:])
		var x int16
		var y int16
		var z int16
		binary.Read(buf, binary.LittleEndian, &x)
		binary.Read(buf, binary.LittleEndian, &y)
		binary.Read(buf, binary.LittleEndian, &z)
		fn(x, y, z)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackAcceleration), wrapper)
}

// Remove a registered Acceleration callback.
func (device *IMUBrick) DeregisterAccelerationCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackAcceleration), registrationId)
}


// This callback is triggered periodically with the period that is set by
// SetMagneticFieldPeriod. The parameters are the magnetic
// field for the x, y and z axis.
func (device *IMUBrick) RegisterMagneticFieldCallback(fn func(int16, int16, int16)) uint64 {
	wrapper := func(byteSlice []byte) {
		buf := bytes.NewBuffer(byteSlice[8:])
		var x int16
		var y int16
		var z int16
		binary.Read(buf, binary.LittleEndian, &x)
		binary.Read(buf, binary.LittleEndian, &y)
		binary.Read(buf, binary.LittleEndian, &z)
		fn(x, y, z)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackMagneticField), wrapper)
}

// Remove a registered Magnetic Field callback.
func (device *IMUBrick) DeregisterMagneticFieldCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackMagneticField), registrationId)
}


// This callback is triggered periodically with the period that is set by
// SetAngularVelocityPeriod. The parameters are the angular
// velocity for the x, y and z axis.
func (device *IMUBrick) RegisterAngularVelocityCallback(fn func(int16, int16, int16)) uint64 {
	wrapper := func(byteSlice []byte) {
		buf := bytes.NewBuffer(byteSlice[8:])
		var x int16
		var y int16
		var z int16
		binary.Read(buf, binary.LittleEndian, &x)
		binary.Read(buf, binary.LittleEndian, &y)
		binary.Read(buf, binary.LittleEndian, &z)
		fn(x, y, z)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackAngularVelocity), wrapper)
}

// Remove a registered Angular Velocity callback.
func (device *IMUBrick) DeregisterAngularVelocityCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackAngularVelocity), registrationId)
}


// This callback is triggered periodically with the period that is set by
// SetAllDataPeriod. The parameters are the acceleration,
// the magnetic field and the angular velocity for the x, y and z axis as
// well as the temperature of the IMU Brick.
func (device *IMUBrick) RegisterAllDataCallback(fn func(int16, int16, int16, int16, int16, int16, int16, int16, int16, int16)) uint64 {
	wrapper := func(byteSlice []byte) {
		buf := bytes.NewBuffer(byteSlice[8:])
		var accX int16
		var accY int16
		var accZ int16
		var magX int16
		var magY int16
		var magZ int16
		var angX int16
		var angY int16
		var angZ int16
		var temperature int16
		binary.Read(buf, binary.LittleEndian, &accX)
		binary.Read(buf, binary.LittleEndian, &accY)
		binary.Read(buf, binary.LittleEndian, &accZ)
		binary.Read(buf, binary.LittleEndian, &magX)
		binary.Read(buf, binary.LittleEndian, &magY)
		binary.Read(buf, binary.LittleEndian, &magZ)
		binary.Read(buf, binary.LittleEndian, &angX)
		binary.Read(buf, binary.LittleEndian, &angY)
		binary.Read(buf, binary.LittleEndian, &angZ)
		binary.Read(buf, binary.LittleEndian, &temperature)
		fn(accX, accY, accZ, magX, magY, magZ, angX, angY, angZ, temperature)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackAllData), wrapper)
}

// Remove a registered All Data callback.
func (device *IMUBrick) DeregisterAllDataCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackAllData), registrationId)
}


// This callback is triggered periodically with the period that is set by
// SetOrientationPeriod. The parameters are the orientation
// (roll, pitch and yaw) of the IMU Brick in Euler angles. See
// GetOrientation for details.
func (device *IMUBrick) RegisterOrientationCallback(fn func(int16, int16, int16)) uint64 {
	wrapper := func(byteSlice []byte) {
		buf := bytes.NewBuffer(byteSlice[8:])
		var roll int16
		var pitch int16
		var yaw int16
		binary.Read(buf, binary.LittleEndian, &roll)
		binary.Read(buf, binary.LittleEndian, &pitch)
		binary.Read(buf, binary.LittleEndian, &yaw)
		fn(roll, pitch, yaw)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackOrientation), wrapper)
}

// Remove a registered Orientation callback.
func (device *IMUBrick) DeregisterOrientationCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackOrientation), registrationId)
}


// This callback is triggered periodically with the period that is set by
// SetQuaternionPeriod. The parameters are the orientation
// (x, y, z, w) of the IMU Brick in quaternions. See GetQuaternion
// for details.
func (device *IMUBrick) RegisterQuaternionCallback(fn func(float32, float32, float32, float32)) uint64 {
	wrapper := func(byteSlice []byte) {
		buf := bytes.NewBuffer(byteSlice[8:])
		var x float32
		var y float32
		var z float32
		var w float32
		binary.Read(buf, binary.LittleEndian, &x)
		binary.Read(buf, binary.LittleEndian, &y)
		binary.Read(buf, binary.LittleEndian, &z)
		binary.Read(buf, binary.LittleEndian, &w)
		fn(x, y, z, w)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackQuaternion), wrapper)
}

// Remove a registered Quaternion callback.
func (device *IMUBrick) DeregisterQuaternionCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackQuaternion), registrationId)
}


// Returns the calibrated acceleration from the accelerometer for the
// x, y and z axis in g/1000 (1g = 9.80665m/s²).
// 
// If you want to get the acceleration periodically, it is recommended
// to use the RegisterAccelerationCallback callback and set the period with
// SetAccelerationPeriod.
func (device *IMUBrick) GetAcceleration() (x int16, y int16, z int16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetAcceleration), buf.Bytes())
	if err != nil {
		return x, y, z, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return x, y, z, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &x)
	binary.Read(resultBuf, binary.LittleEndian, &y)
	binary.Read(resultBuf, binary.LittleEndian, &z)

	}

	return x, y, z, nil
}

// Returns the calibrated magnetic field from the magnetometer for the
// x, y and z axis in mG (Milligauss or Nanotesla).
// 
// If you want to get the magnetic field periodically, it is recommended
// to use the RegisterMagneticFieldCallback callback and set the period with
// SetMagneticFieldPeriod.
func (device *IMUBrick) GetMagneticField() (x int16, y int16, z int16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetMagneticField), buf.Bytes())
	if err != nil {
		return x, y, z, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return x, y, z, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &x)
	binary.Read(resultBuf, binary.LittleEndian, &y)
	binary.Read(resultBuf, binary.LittleEndian, &z)

	}

	return x, y, z, nil
}

// Returns the calibrated angular velocity from the gyroscope for the
// x, y and z axis in °/14.375s (you have to divide by 14.375 to
// get the value in °/s).
// 
// If you want to get the angular velocity periodically, it is recommended
// to use the RegisterAngularVelocityCallback callback and set the period with
// SetAngularVelocityPeriod.
func (device *IMUBrick) GetAngularVelocity() (x int16, y int16, z int16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetAngularVelocity), buf.Bytes())
	if err != nil {
		return x, y, z, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return x, y, z, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &x)
	binary.Read(resultBuf, binary.LittleEndian, &y)
	binary.Read(resultBuf, binary.LittleEndian, &z)

	}

	return x, y, z, nil
}

// Returns the data from GetAcceleration, GetMagneticField
// and GetAngularVelocity as well as the temperature of the IMU Brick.
// 
// The temperature is given in °C/100.
// 
// If you want to get the data periodically, it is recommended
// to use the RegisterAllDataCallback callback and set the period with
// SetAllDataPeriod.
func (device *IMUBrick) GetAllData() (accX int16, accY int16, accZ int16, magX int16, magY int16, magZ int16, angX int16, angY int16, angZ int16, temperature int16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetAllData), buf.Bytes())
	if err != nil {
		return accX, accY, accZ, magX, magY, magZ, angX, angY, angZ, temperature, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return accX, accY, accZ, magX, magY, magZ, angX, angY, angZ, temperature, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &accX)
	binary.Read(resultBuf, binary.LittleEndian, &accY)
	binary.Read(resultBuf, binary.LittleEndian, &accZ)
	binary.Read(resultBuf, binary.LittleEndian, &magX)
	binary.Read(resultBuf, binary.LittleEndian, &magY)
	binary.Read(resultBuf, binary.LittleEndian, &magZ)
	binary.Read(resultBuf, binary.LittleEndian, &angX)
	binary.Read(resultBuf, binary.LittleEndian, &angY)
	binary.Read(resultBuf, binary.LittleEndian, &angZ)
	binary.Read(resultBuf, binary.LittleEndian, &temperature)

	}

	return accX, accY, accZ, magX, magY, magZ, angX, angY, angZ, temperature, nil
}

// Returns the current orientation (roll, pitch, yaw) of the IMU Brick as Euler
// angles in one-hundredth degree. Note that Euler angles always experience a
// https://en.wikipedia.org/wiki/Gimbal_lock.
// 
// We recommend that you use quaternions instead.
// 
// The order to sequence in which the orientation values should be applied is
// roll, yaw, pitch.
// 
// If you want to get the orientation periodically, it is recommended
// to use the RegisterOrientationCallback callback and set the period with
// SetOrientationPeriod.
func (device *IMUBrick) GetOrientation() (roll int16, pitch int16, yaw int16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetOrientation), buf.Bytes())
	if err != nil {
		return roll, pitch, yaw, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return roll, pitch, yaw, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &roll)
	binary.Read(resultBuf, binary.LittleEndian, &pitch)
	binary.Read(resultBuf, binary.LittleEndian, &yaw)

	}

	return roll, pitch, yaw, nil
}

// Returns the current orientation (x, y, z, w) of the IMU as
// https://en.wikipedia.org/wiki/Quaternions_and_spatial_rotation.
// 
// You can go from quaternions to Euler angles with the following formula::
// 
//  xAngle = atan2(2*y*w - 2*x*z, 1 - 2*y*y - 2*z*z)
//  yAngle = atan2(2*x*w - 2*y*z, 1 - 2*x*x - 2*z*z)
//  zAngle =  asin(2*x*y + 2*z*w)
// 
// This process is not reversible, because of the
// https://en.wikipedia.org/wiki/Gimbal_lock.
// 
// It is also possible to calculate independent angles. You can calculate
// yaw, pitch and roll in a right-handed vehicle coordinate system according to
// DIN70000 with::
// 
//  yaw   =  atan2(2*x*y + 2*w*z, w*w + x*x - y*y - z*z)
//  pitch = -asin(2*w*y - 2*x*z)
//  roll  = -atan2(2*y*z + 2*w*x, -w*w + x*x + y*y - z*z))
// 
// Converting the quaternions to an OpenGL transformation matrix is
// possible with the following formula::
// 
//  matrix = [[1 - 2*(y*y + z*z),     2*(x*y - w*z),     2*(x*z + w*y), 0],
//            [    2*(x*y + w*z), 1 - 2*(x*x + z*z),     2*(y*z - w*x), 0],
//            [    2*(x*z - w*y),     2*(y*z + w*x), 1 - 2*(x*x + y*y), 0],
//            [                0,                 0,                 0, 1]]
// 
// If you want to get the quaternions periodically, it is recommended
// to use the RegisterQuaternionCallback callback and set the period with
// SetQuaternionPeriod.
func (device *IMUBrick) GetQuaternion() (x float32, y float32, z float32, w float32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetQuaternion), buf.Bytes())
	if err != nil {
		return x, y, z, w, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return x, y, z, w, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &x)
	binary.Read(resultBuf, binary.LittleEndian, &y)
	binary.Read(resultBuf, binary.LittleEndian, &z)
	binary.Read(resultBuf, binary.LittleEndian, &w)

	}

	return x, y, z, w, nil
}

// Returns the temperature of the IMU Brick. The temperature is given in
// °C/100.
func (device *IMUBrick) GetIMUTemperature() (temperature int16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetIMUTemperature), buf.Bytes())
	if err != nil {
		return temperature, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return temperature, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &temperature)

	}

	return temperature, nil
}

// Turns the orientation and direction LEDs of the IMU Brick on.
func (device *IMUBrick) LedsOn() (err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Set(uint8(FunctionLedsOn), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])
		
	}

	return nil
}

// Turns the orientation and direction LEDs of the IMU Brick off.
func (device *IMUBrick) LedsOff() (err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Set(uint8(FunctionLedsOff), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])
		
	}

	return nil
}

// Returns *true* if the orientation and direction LEDs of the IMU Brick
// are on, *false* otherwise.
func (device *IMUBrick) AreLedsOn() (leds bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionAreLedsOn), buf.Bytes())
	if err != nil {
		return leds, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return leds, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &leds)

	}

	return leds, nil
}

// Not implemented yet.
func (device *IMUBrick) SetAccelerationRange(range_ uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, range_);

	resultBytes, err := device.device.Set(uint8(FunctionSetAccelerationRange), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])
		
	}

	return nil
}

// Not implemented yet.
func (device *IMUBrick) GetAccelerationRange() (range_ uint8, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetAccelerationRange), buf.Bytes())
	if err != nil {
		return range_, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return range_, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &range_)

	}

	return range_, nil
}

// Not implemented yet.
func (device *IMUBrick) SetMagnetometerRange(range_ uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, range_);

	resultBytes, err := device.device.Set(uint8(FunctionSetMagnetometerRange), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])
		
	}

	return nil
}

// Not implemented yet.
func (device *IMUBrick) GetMagnetometerRange() (range_ uint8, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetMagnetometerRange), buf.Bytes())
	if err != nil {
		return range_, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return range_, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &range_)

	}

	return range_, nil
}

// Sets the convergence speed of the IMU Brick in °/s. The convergence speed
// determines how the different sensor measurements are fused.
// 
// If the orientation of the IMU Brick is off by 10° and the convergence speed is
// set to 20°/s, it will take 0.5s until the orientation is corrected. However,
// if the correct orientation is reached and the convergence speed is too high,
// the orientation will fluctuate with the fluctuations of the accelerometer and
// the magnetometer.
// 
// If you set the convergence speed to 0, practically only the gyroscope is used
// to calculate the orientation. This gives very smooth movements, but errors of the
// gyroscope will not be corrected. If you set the convergence speed to something
// above 500, practically only the magnetometer and the accelerometer are used to
// calculate the orientation. In this case the movements are abrupt and the values
// will fluctuate, but there won't be any errors that accumulate over time.
// 
// In an application with high angular velocities, we recommend a high convergence
// speed, so the errors of the gyroscope can be corrected fast. In applications with
// only slow movements we recommend a low convergence speed. You can change the
// convergence speed on the fly. So it is possible (and recommended) to increase
// the convergence speed before an abrupt movement and decrease it afterwards
// again.
// 
// You might want to play around with the convergence speed in the Brick Viewer to
// get a feeling for a good value for your application.
// 
// The default value is 30.
func (device *IMUBrick) SetConvergenceSpeed(speed uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, speed);

	resultBytes, err := device.device.Set(uint8(FunctionSetConvergenceSpeed), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])
		
	}

	return nil
}

// Returns the convergence speed as set by SetConvergenceSpeed.
func (device *IMUBrick) GetConvergenceSpeed() (speed uint16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetConvergenceSpeed), buf.Bytes())
	if err != nil {
		return speed, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return speed, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &speed)

	}

	return speed, nil
}

// There are several different types that can be calibrated:
// 
//  Type| Description| Values
//  --- | --- | --- 
//  0|    Accelerometer Gain| ``[mul x| mul y| mul z| div x| div y| div z| 0| 0| 0| 0]``
//  1|    Accelerometer Bias| ``[bias x| bias y| bias z| 0| 0| 0| 0| 0| 0| 0]``
//  2|    Magnetometer Gain|  ``[mul x| mul y| mul z| div x| div y| div z| 0| 0| 0| 0]``
//  3|    Magnetometer Bias|  ``[bias x| bias y| bias z| 0| 0| 0| 0| 0| 0| 0]``
//  4|    Gyroscope Gain|     ``[mul x| mul y| mul z| div x| div y| div z| 0| 0| 0| 0]``
//  5|    Gyroscope Bias|     ``[bias xl| bias yl| bias zl| temp l| bias xh| bias yh| bias zh| temp h| 0| 0]``
// 
// The calibration via gain and bias is done with the following formula::
// 
//  new_value = (bias + orig_value) * gain_mul / gain_div
// 
// If you really want to write your own calibration software, please keep
// in mind that you first have to undo the old calibration (set bias to 0 and
// gain to 1/1) and that you have to average over several thousand values
// to obtain a usable result in the end.
// 
// The gyroscope bias is highly dependent on the temperature, so you have to
// calibrate the bias two times with different temperatures. The values ``xl``,
// ``yl``, ``zl`` and ``temp l`` are the bias for ``x``, ``y``, ``z`` and the
// corresponding temperature for a low temperature. The values ``xh``, ``yh``,
// ``zh`` and ``temp h`` are the same for a high temperatures. The temperature
// difference should be at least 5°C. If you have a temperature where the
// IMU Brick is mostly used, you should use this temperature for one of the
// sampling points.
// 
// Note
//  We highly recommend that you use the Brick Viewer to calibrate your
//  IMU Brick.
//
// Associated constants:
//
//	* CalibrationTypeAccelerometerGain
//	* CalibrationTypeAccelerometerBias
//	* CalibrationTypeMagnetometerGain
//	* CalibrationTypeMagnetometerBias
//	* CalibrationTypeGyroscopeGain
//	* CalibrationTypeGyroscopeBias
func (device *IMUBrick) SetCalibration(typ CalibrationType, data [10]int16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, typ);
	binary.Write(&buf, binary.LittleEndian, data);

	resultBytes, err := device.device.Set(uint8(FunctionSetCalibration), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])
		
	}

	return nil
}

// Returns the calibration for a given type as set by SetCalibration.
//
// Associated constants:
//
//	* CalibrationTypeAccelerometerGain
//	* CalibrationTypeAccelerometerBias
//	* CalibrationTypeMagnetometerGain
//	* CalibrationTypeMagnetometerBias
//	* CalibrationTypeGyroscopeGain
//	* CalibrationTypeGyroscopeBias
func (device *IMUBrick) GetCalibration(typ CalibrationType) (data [10]int16, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, typ);

	resultBytes, err := device.device.Get(uint8(FunctionGetCalibration), buf.Bytes())
	if err != nil {
		return data, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return data, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &data)

	}

	return data, nil
}

// Sets the period with which the RegisterAccelerationCallback callback is triggered
// periodically. A value of 0 turns the callback off.
func (device *IMUBrick) SetAccelerationPeriod(period uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);

	resultBytes, err := device.device.Set(uint8(FunctionSetAccelerationPeriod), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])
		
	}

	return nil
}

// Returns the period as set by SetAccelerationPeriod.
func (device *IMUBrick) GetAccelerationPeriod() (period uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetAccelerationPeriod), buf.Bytes())
	if err != nil {
		return period, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return period, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &period)

	}

	return period, nil
}

// Sets the period with which the RegisterMagneticFieldCallback callback is
// triggered periodically. A value of 0 turns the callback off.
func (device *IMUBrick) SetMagneticFieldPeriod(period uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);

	resultBytes, err := device.device.Set(uint8(FunctionSetMagneticFieldPeriod), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])
		
	}

	return nil
}

// Returns the period as set by SetMagneticFieldPeriod.
func (device *IMUBrick) GetMagneticFieldPeriod() (period uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetMagneticFieldPeriod), buf.Bytes())
	if err != nil {
		return period, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return period, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &period)

	}

	return period, nil
}

// Sets the period with which the RegisterAngularVelocityCallback callback is
// triggered periodically. A value of 0 turns the callback off.
func (device *IMUBrick) SetAngularVelocityPeriod(period uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);

	resultBytes, err := device.device.Set(uint8(FunctionSetAngularVelocityPeriod), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])
		
	}

	return nil
}

// Returns the period as set by SetAngularVelocityPeriod.
func (device *IMUBrick) GetAngularVelocityPeriod() (period uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetAngularVelocityPeriod), buf.Bytes())
	if err != nil {
		return period, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return period, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &period)

	}

	return period, nil
}

// Sets the period with which the RegisterAllDataCallback callback is triggered
// periodically. A value of 0 turns the callback off.
func (device *IMUBrick) SetAllDataPeriod(period uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);

	resultBytes, err := device.device.Set(uint8(FunctionSetAllDataPeriod), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])
		
	}

	return nil
}

// Returns the period as set by SetAllDataPeriod.
func (device *IMUBrick) GetAllDataPeriod() (period uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetAllDataPeriod), buf.Bytes())
	if err != nil {
		return period, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return period, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &period)

	}

	return period, nil
}

// Sets the period with which the RegisterOrientationCallback callback is triggered
// periodically. A value of 0 turns the callback off.
func (device *IMUBrick) SetOrientationPeriod(period uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);

	resultBytes, err := device.device.Set(uint8(FunctionSetOrientationPeriod), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])
		
	}

	return nil
}

// Returns the period as set by SetOrientationPeriod.
func (device *IMUBrick) GetOrientationPeriod() (period uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetOrientationPeriod), buf.Bytes())
	if err != nil {
		return period, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return period, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &period)

	}

	return period, nil
}

// Sets the period with which the RegisterQuaternionCallback callback is triggered
// periodically. A value of 0 turns the callback off.
func (device *IMUBrick) SetQuaternionPeriod(period uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);

	resultBytes, err := device.device.Set(uint8(FunctionSetQuaternionPeriod), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])
		
	}

	return nil
}

// Returns the period as set by SetQuaternionPeriod.
func (device *IMUBrick) GetQuaternionPeriod() (period uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetQuaternionPeriod), buf.Bytes())
	if err != nil {
		return period, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return period, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &period)

	}

	return period, nil
}

// Turns the orientation calculation of the IMU Brick on.
// 
// As default the calculation is on.
// 
// .. versionadded:: 2.0.2$nbsp;(Firmware)
func (device *IMUBrick) OrientationCalculationOn() (err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Set(uint8(FunctionOrientationCalculationOn), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])
		
	}

	return nil
}

// Turns the orientation calculation of the IMU Brick off.
// 
// If the calculation is off, GetOrientation will return
// the last calculated value until the calculation is turned on again.
// 
// The trigonometric functions that are needed to calculate the orientation
// are very expensive. We recommend to turn the orientation calculation
// off if the orientation is not needed, to free calculation time for the
// sensor fusion algorithm.
// 
// As default the calculation is on.
// 
// .. versionadded:: 2.0.2$nbsp;(Firmware)
func (device *IMUBrick) OrientationCalculationOff() (err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Set(uint8(FunctionOrientationCalculationOff), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])
		
	}

	return nil
}

// Returns *true* if the orientation calculation of the IMU Brick
// is on, *false* otherwise.
// 
// .. versionadded:: 2.0.2$nbsp;(Firmware)
func (device *IMUBrick) IsOrientationCalculationOn() (orientationCalculationOn bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionIsOrientationCalculationOn), buf.Bytes())
	if err != nil {
		return orientationCalculationOn, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return orientationCalculationOn, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &orientationCalculationOn)

	}

	return orientationCalculationOn, nil
}

// The SPITF protocol can be used with a dynamic baudrate. If the dynamic baudrate is
// enabled, the Brick will try to adapt the baudrate for the communication
// between Bricks and Bricklets according to the amount of data that is transferred.
// 
// The baudrate will be increased exponentially if lots of data is send/received and
// decreased linearly if little data is send/received.
// 
// This lowers the baudrate in applications where little data is transferred (e.g.
// a weather station) and increases the robustness. If there is lots of data to transfer
// (e.g. Thermal Imaging Bricklet) it automatically increases the baudrate as needed.
// 
// In cases where some data has to transferred as fast as possible every few seconds
// (e.g. RS485 Bricklet with a high baudrate but small payload) you may want to turn
// the dynamic baudrate off to get the highest possible performance.
// 
// The maximum value of the baudrate can be set per port with the function
// SetSPITFPBaudrate. If the dynamic baudrate is disabled, the baudrate
// as set by SetSPITFPBaudrate will be used statically.
// 
// The minimum dynamic baudrate has a value range of 400000 to 2000000 baud.
// 
// By default dynamic baudrate is enabled and the minimum dynamic baudrate is 400000.
// 
// .. versionadded:: 2.3.5$nbsp;(Firmware)
func (device *IMUBrick) SetSPITFPBaudrateConfig(enableDynamicBaudrate bool, minimumDynamicBaudrate uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, enableDynamicBaudrate);
	binary.Write(&buf, binary.LittleEndian, minimumDynamicBaudrate);

	resultBytes, err := device.device.Set(uint8(FunctionSetSPITFPBaudrateConfig), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])
		
	}

	return nil
}

// Returns the baudrate config, see SetSPITFPBaudrateConfig.
// 
// .. versionadded:: 2.3.5$nbsp;(Firmware)
func (device *IMUBrick) GetSPITFPBaudrateConfig() (enableDynamicBaudrate bool, minimumDynamicBaudrate uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetSPITFPBaudrateConfig), buf.Bytes())
	if err != nil {
		return enableDynamicBaudrate, minimumDynamicBaudrate, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return enableDynamicBaudrate, minimumDynamicBaudrate, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &enableDynamicBaudrate)
	binary.Read(resultBuf, binary.LittleEndian, &minimumDynamicBaudrate)

	}

	return enableDynamicBaudrate, minimumDynamicBaudrate, nil
}

// Returns the timeout count for the different communication methods.
// 
// The methods 0-2 are available for all Bricks, 3-7 only for Master Bricks.
// 
// This function is mostly used for debugging during development, in normal operation
// the counters should nearly always stay at 0.
// 
// .. versionadded:: 2.3.3$nbsp;(Firmware)
//
// Associated constants:
//
//	* CommunicationMethodNone
//	* CommunicationMethodUSB
//	* CommunicationMethodSPIStack
//	* CommunicationMethodChibi
//	* CommunicationMethodRS485
//	* CommunicationMethodWIFI
//	* CommunicationMethodEthernet
//	* CommunicationMethodWIFIV2
func (device *IMUBrick) GetSendTimeoutCount(communicationMethod CommunicationMethod) (timeoutCount uint32, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, communicationMethod);

	resultBytes, err := device.device.Get(uint8(FunctionGetSendTimeoutCount), buf.Bytes())
	if err != nil {
		return timeoutCount, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return timeoutCount, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &timeoutCount)

	}

	return timeoutCount, nil
}

// Sets the baudrate for a specific Bricklet port ('a' - 'd'). The
// baudrate can be in the range 400000 to 2000000.
// 
// If you want to increase the throughput of Bricklets you can increase
// the baudrate. If you get a high error count because of high
// interference (see GetSPITFPErrorCount) you can decrease the
// baudrate.
// 
// If the dynamic baudrate feature is enabled, the baudrate set by this
// function corresponds to the maximum baudrate (see SetSPITFPBaudrateConfig).
// 
// Regulatory testing is done with the default baudrate. If CE compatibility
// or similar is necessary in you applications we recommend to not change
// the baudrate.
// 
// The default baudrate for all ports is 1400000.
// 
// .. versionadded:: 2.3.3$nbsp;(Firmware)
func (device *IMUBrick) SetSPITFPBaudrate(brickletPort rune, baudrate uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, brickletPort);
	binary.Write(&buf, binary.LittleEndian, baudrate);

	resultBytes, err := device.device.Set(uint8(FunctionSetSPITFPBaudrate), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])
		
	}

	return nil
}

// Returns the baudrate for a given Bricklet port, see SetSPITFPBaudrate.
// 
// .. versionadded:: 2.3.3$nbsp;(Firmware)
func (device *IMUBrick) GetSPITFPBaudrate(brickletPort rune) (baudrate uint32, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, brickletPort);

	resultBytes, err := device.device.Get(uint8(FunctionGetSPITFPBaudrate), buf.Bytes())
	if err != nil {
		return baudrate, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return baudrate, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &baudrate)

	}

	return baudrate, nil
}

// Returns the error count for the communication between Brick and Bricklet.
// 
// The errors are divided into
// 
// * ACK checksum errors,
// * message checksum errors,
// * framing errors and
// * overflow errors.
// 
// The errors counts are for errors that occur on the Brick side. All
// Bricklets have a similar function that returns the errors on the Bricklet side.
// 
// .. versionadded:: 2.3.3$nbsp;(Firmware)
func (device *IMUBrick) GetSPITFPErrorCount(brickletPort rune) (errorCountACKChecksum uint32, errorCountMessageChecksum uint32, errorCountFrame uint32, errorCountOverflow uint32, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, brickletPort);

	resultBytes, err := device.device.Get(uint8(FunctionGetSPITFPErrorCount), buf.Bytes())
	if err != nil {
		return errorCountACKChecksum, errorCountMessageChecksum, errorCountFrame, errorCountOverflow, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return errorCountACKChecksum, errorCountMessageChecksum, errorCountFrame, errorCountOverflow, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &errorCountACKChecksum)
	binary.Read(resultBuf, binary.LittleEndian, &errorCountMessageChecksum)
	binary.Read(resultBuf, binary.LittleEndian, &errorCountFrame)
	binary.Read(resultBuf, binary.LittleEndian, &errorCountOverflow)

	}

	return errorCountACKChecksum, errorCountMessageChecksum, errorCountFrame, errorCountOverflow, nil
}

// Enables the status LED.
// 
// The status LED is the blue LED next to the USB connector. If enabled is is
// on and it flickers if data is transfered. If disabled it is always off.
// 
// The default state is enabled.
// 
// .. versionadded:: 2.3.1$nbsp;(Firmware)
func (device *IMUBrick) EnableStatusLED() (err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Set(uint8(FunctionEnableStatusLED), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])
		
	}

	return nil
}

// Disables the status LED.
// 
// The status LED is the blue LED next to the USB connector. If enabled is is
// on and it flickers if data is transfered. If disabled it is always off.
// 
// The default state is enabled.
// 
// .. versionadded:: 2.3.1$nbsp;(Firmware)
func (device *IMUBrick) DisableStatusLED() (err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Set(uint8(FunctionDisableStatusLED), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])
		
	}

	return nil
}

// Returns *true* if the status LED is enabled, *false* otherwise.
// 
// .. versionadded:: 2.3.1$nbsp;(Firmware)
func (device *IMUBrick) IsStatusLEDEnabled() (enabled bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionIsStatusLEDEnabled), buf.Bytes())
	if err != nil {
		return enabled, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return enabled, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &enabled)

	}

	return enabled, nil
}

// Returns the firmware and protocol version and the name of the Bricklet for a
// given port.
// 
// This functions sole purpose is to allow automatic flashing of v1.x.y Bricklet
// plugins.
func (device *IMUBrick) GetProtocol1BrickletName(port rune) (protocolVersion uint8, firmwareVersion [3]uint8, name string, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, port);

	resultBytes, err := device.device.Get(uint8(FunctionGetProtocol1BrickletName), buf.Bytes())
	if err != nil {
		return protocolVersion, firmwareVersion, name, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return protocolVersion, firmwareVersion, name, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &protocolVersion)
	binary.Read(resultBuf, binary.LittleEndian, &firmwareVersion)
	name = ByteSliceToString(resultBuf.Next(40))

	}

	return protocolVersion, firmwareVersion, name, nil
}

// Returns the temperature in °C/10 as measured inside the microcontroller. The
// value returned is not the ambient temperature!
// 
// The temperature is only proportional to the real temperature and it has an
// accuracy of ±15%. Practically it is only useful as an indicator for
// temperature changes.
func (device *IMUBrick) GetChipTemperature() (temperature int16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetChipTemperature), buf.Bytes())
	if err != nil {
		return temperature, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return temperature, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &temperature)

	}

	return temperature, nil
}

// Calling this function will reset the Brick. Calling this function
// on a Brick inside of a stack will reset the whole stack.
// 
// After a reset you have to create new device objects,
// calling functions on the existing ones will result in
// undefined behavior!
func (device *IMUBrick) Reset() (err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Set(uint8(FunctionReset), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])
		
	}

	return nil
}

// Returns the UID, the UID where the Brick is connected to,
// the position, the hardware and firmware version as well as the
// device identifier.
// 
// The position can be '0'-'8' (stack position).
// 
// The device identifier numbers can be found `here <device_identifier>`.
// |device_identifier_constant|
func (device *IMUBrick) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetIdentity), buf.Bytes())
	if err != nil {
		return uid, connectedUid, position, hardwareVersion, firmwareVersion, deviceIdentifier, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return uid, connectedUid, position, hardwareVersion, firmwareVersion, deviceIdentifier, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		uid = ByteSliceToString(resultBuf.Next(8))
	connectedUid = ByteSliceToString(resultBuf.Next(8))
	position = rune(resultBuf.Next(1)[0])
	binary.Read(resultBuf, binary.LittleEndian, &hardwareVersion)
	binary.Read(resultBuf, binary.LittleEndian, &firmwareVersion)
	binary.Read(resultBuf, binary.LittleEndian, &deviceIdentifier)

	}

	return uid, connectedUid, position, hardwareVersion, firmwareVersion, deviceIdentifier, nil
}
