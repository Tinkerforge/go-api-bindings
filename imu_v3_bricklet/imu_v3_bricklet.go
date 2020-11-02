/* ***********************************************************
 * This file was automatically generated on 2020-11-02.      *
 *                                                           *
 * Go Bindings Version 2.0.9                                 *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


// Full fledged AHRS with 9 degrees of freedom.
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/IMUV3_Bricklet_Go.html.
package imu_v3_bricklet

import (
	"encoding/binary"
	"bytes"
	"fmt"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionGetAcceleration Function = 1
	FunctionGetMagneticField Function = 2
	FunctionGetAngularVelocity Function = 3
	FunctionGetTemperature Function = 4
	FunctionGetOrientation Function = 5
	FunctionGetLinearAcceleration Function = 6
	FunctionGetGravityVector Function = 7
	FunctionGetQuaternion Function = 8
	FunctionGetAllData Function = 9
	FunctionSaveCalibration Function = 10
	FunctionSetSensorConfiguration Function = 11
	FunctionGetSensorConfiguration Function = 12
	FunctionSetSensorFusionMode Function = 13
	FunctionGetSensorFusionMode Function = 14
	FunctionSetAccelerationCallbackConfiguration Function = 15
	FunctionGetAccelerationCallbackConfiguration Function = 16
	FunctionSetMagneticFieldCallbackConfiguration Function = 17
	FunctionGetMagneticFieldCallbackConfiguration Function = 18
	FunctionSetAngularVelocityCallbackConfiguration Function = 19
	FunctionGetAngularVelocityCallbackConfiguration Function = 20
	FunctionSetTemperatureCallbackConfiguration Function = 21
	FunctionGetTemperatureCallbackConfiguration Function = 22
	FunctionSetOrientationCallbackConfiguration Function = 23
	FunctionGetOrientationCallbackConfiguration Function = 24
	FunctionSetLinearAccelerationCallbackConfiguration Function = 25
	FunctionGetLinearAccelerationCallbackConfiguration Function = 26
	FunctionSetGravityVectorCallbackConfiguration Function = 27
	FunctionGetGravityVectorCallbackConfiguration Function = 28
	FunctionSetQuaternionCallbackConfiguration Function = 29
	FunctionGetQuaternionCallbackConfiguration Function = 30
	FunctionSetAllDataCallbackConfiguration Function = 31
	FunctionGetAllDataCallbackConfiguration Function = 32
	FunctionGetSPITFPErrorCount Function = 234
	FunctionSetBootloaderMode Function = 235
	FunctionGetBootloaderMode Function = 236
	FunctionSetWriteFirmwarePointer Function = 237
	FunctionWriteFirmware Function = 238
	FunctionSetStatusLEDConfig Function = 239
	FunctionGetStatusLEDConfig Function = 240
	FunctionGetChipTemperature Function = 242
	FunctionReset Function = 243
	FunctionWriteUID Function = 248
	FunctionReadUID Function = 249
	FunctionGetIdentity Function = 255
	FunctionCallbackAcceleration Function = 33
	FunctionCallbackMagneticField Function = 34
	FunctionCallbackAngularVelocity Function = 35
	FunctionCallbackTemperature Function = 36
	FunctionCallbackLinearAcceleration Function = 37
	FunctionCallbackGravityVector Function = 38
	FunctionCallbackOrientation Function = 39
	FunctionCallbackQuaternion Function = 40
	FunctionCallbackAllData Function = 41
)

type MagnetometerRate = uint8

const (
	MagnetometerRate2Hz MagnetometerRate = 0
	MagnetometerRate6Hz MagnetometerRate = 1
	MagnetometerRate8Hz MagnetometerRate = 2
	MagnetometerRate10Hz MagnetometerRate = 3
	MagnetometerRate15Hz MagnetometerRate = 4
	MagnetometerRate20Hz MagnetometerRate = 5
	MagnetometerRate25Hz MagnetometerRate = 6
	MagnetometerRate30Hz MagnetometerRate = 7
)

type GyroscopeRange = uint8

const (
	GyroscopeRange2000DPS GyroscopeRange = 0
	GyroscopeRange1000DPS GyroscopeRange = 1
	GyroscopeRange500DPS GyroscopeRange = 2
	GyroscopeRange250DPS GyroscopeRange = 3
	GyroscopeRange125DPS GyroscopeRange = 4
)

type GyroscopeBandwidth = uint8

const (
	GyroscopeBandwidth523Hz GyroscopeBandwidth = 0
	GyroscopeBandwidth230Hz GyroscopeBandwidth = 1
	GyroscopeBandwidth116Hz GyroscopeBandwidth = 2
	GyroscopeBandwidth47Hz GyroscopeBandwidth = 3
	GyroscopeBandwidth23Hz GyroscopeBandwidth = 4
	GyroscopeBandwidth12Hz GyroscopeBandwidth = 5
	GyroscopeBandwidth64Hz GyroscopeBandwidth = 6
	GyroscopeBandwidth32Hz GyroscopeBandwidth = 7
)

type AccelerometerRange = uint8

const (
	AccelerometerRange2G AccelerometerRange = 0
	AccelerometerRange4G AccelerometerRange = 1
	AccelerometerRange8G AccelerometerRange = 2
	AccelerometerRange16G AccelerometerRange = 3
)

type AccelerometerBandwidth = uint8

const (
	AccelerometerBandwidth7_81Hz AccelerometerBandwidth = 0
	AccelerometerBandwidth15_63Hz AccelerometerBandwidth = 1
	AccelerometerBandwidth31_25Hz AccelerometerBandwidth = 2
	AccelerometerBandwidth62_5Hz AccelerometerBandwidth = 3
	AccelerometerBandwidth125Hz AccelerometerBandwidth = 4
	AccelerometerBandwidth250Hz AccelerometerBandwidth = 5
	AccelerometerBandwidth500Hz AccelerometerBandwidth = 6
	AccelerometerBandwidth1000Hz AccelerometerBandwidth = 7
)

type SensorFusion = uint8

const (
	SensorFusionOff SensorFusion = 0
	SensorFusionOn SensorFusion = 1
	SensorFusionOnWithoutMagnetometer SensorFusion = 2
	SensorFusionOnWithoutFastMagnetometerCalibration SensorFusion = 3
)

type BootloaderMode = uint8

const (
	BootloaderModeBootloader BootloaderMode = 0
	BootloaderModeFirmware BootloaderMode = 1
	BootloaderModeBootloaderWaitForReboot BootloaderMode = 2
	BootloaderModeFirmwareWaitForReboot BootloaderMode = 3
	BootloaderModeFirmwareWaitForEraseAndReboot BootloaderMode = 4
)

type BootloaderStatus = uint8

const (
	BootloaderStatusOK BootloaderStatus = 0
	BootloaderStatusInvalidMode BootloaderStatus = 1
	BootloaderStatusNoChange BootloaderStatus = 2
	BootloaderStatusEntryFunctionNotPresent BootloaderStatus = 3
	BootloaderStatusDeviceIdentifierIncorrect BootloaderStatus = 4
	BootloaderStatusCRCMismatch BootloaderStatus = 5
)

type StatusLEDConfig = uint8

const (
	StatusLEDConfigOff StatusLEDConfig = 0
	StatusLEDConfigOn StatusLEDConfig = 1
	StatusLEDConfigShowHeartbeat StatusLEDConfig = 2
	StatusLEDConfigShowStatus StatusLEDConfig = 3
)

type IMUV3Bricklet struct {
	device Device
}
const DeviceIdentifier = 2161
const DeviceDisplayName = "IMU Bricklet 3.0"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (IMUV3Bricklet, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{ 2,0,0 }, uid, &internalIPCon, 0, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return IMUV3Bricklet{}, err
	}
	dev.ResponseExpected[FunctionGetAcceleration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetMagneticField] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetAngularVelocity] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetTemperature] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetOrientation] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetLinearAcceleration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetGravityVector] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetQuaternion] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetAllData] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSaveCalibration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetSensorConfiguration] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetSensorConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetSensorFusionMode] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetSensorFusionMode] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetAccelerationCallbackConfiguration] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetAccelerationCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetMagneticFieldCallbackConfiguration] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetMagneticFieldCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetAngularVelocityCallbackConfiguration] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetAngularVelocityCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetTemperatureCallbackConfiguration] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetTemperatureCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetOrientationCallbackConfiguration] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetOrientationCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetLinearAccelerationCallbackConfiguration] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetLinearAccelerationCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetGravityVectorCallbackConfiguration] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetGravityVectorCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetQuaternionCallbackConfiguration] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetQuaternionCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetAllDataCallbackConfiguration] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetAllDataCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetSPITFPErrorCount] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetBootloaderMode] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetBootloaderMode] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetWriteFirmwarePointer] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionWriteFirmware] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetStatusLEDConfig] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetStatusLEDConfig] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetChipTemperature] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionReset] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionWriteUID] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionReadUID] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetIdentity] = ResponseExpectedFlagAlwaysTrue;
	return IMUV3Bricklet{dev}, nil
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
// for this purpose. If this flag is disabled for a setter function then no response is sent
// and errors are silently ignored, because they cannot be detected.
//
// See SetResponseExpected for the list of function ID constants available for this function.
func (device *IMUV3Bricklet) GetResponseExpected(functionID Function) (bool, error) {
	return device.device.GetResponseExpected(uint8(functionID))
}

// Changes the response expected flag of the function specified by the function ID parameter.
// This flag can only be changed for setter (default value: false) and callback configuration
// functions (default value: true). For getter functions it is always enabled.
//
// Enabling the response expected flag for a setter function allows to detect timeouts and
// other error conditions calls of this setter as well. The device will then send a response
// for this purpose. If this flag is disabled for a setter function then no response is sent
// and errors are silently ignored, because they cannot be detected.
func (device *IMUV3Bricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *IMUV3Bricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *IMUV3Bricklet) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// This callback is triggered periodically with the period that is set by
// SetAccelerationCallbackConfiguration. The parameters are the acceleration
// for the x, y and z axis.
func (device *IMUV3Bricklet) RegisterAccelerationCallback(fn func(int16, int16, int16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 14 {
			return
		}
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
func (device *IMUV3Bricklet) DeregisterAccelerationCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackAcceleration), registrationId)
}


// This callback is triggered periodically with the period that is set by
// SetMagneticFieldCallbackConfiguration. The parameters are the magnetic
// field for the x, y and z axis.
func (device *IMUV3Bricklet) RegisterMagneticFieldCallback(fn func(int16, int16, int16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 14 {
			return
		}
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
func (device *IMUV3Bricklet) DeregisterMagneticFieldCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackMagneticField), registrationId)
}


// This callback is triggered periodically with the period that is set by
// SetAngularVelocityCallbackConfiguration. The parameters are the angular
// velocity for the x, y and z axis.
func (device *IMUV3Bricklet) RegisterAngularVelocityCallback(fn func(int16, int16, int16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 14 {
			return
		}
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
func (device *IMUV3Bricklet) DeregisterAngularVelocityCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackAngularVelocity), registrationId)
}


// This callback is triggered periodically with the period that is set by
// SetTemperatureCallbackConfiguration. The parameter is the temperature.
func (device *IMUV3Bricklet) RegisterTemperatureCallback(fn func(int8)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 9 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var temperature int8
		binary.Read(buf, binary.LittleEndian, &temperature)
		fn(temperature)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackTemperature), wrapper)
}

// Remove a registered Temperature callback.
func (device *IMUV3Bricklet) DeregisterTemperatureCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackTemperature), registrationId)
}


// This callback is triggered periodically with the period that is set by
// SetLinearAccelerationCallbackConfiguration. The parameters are the
// linear acceleration for the x, y and z axis.
func (device *IMUV3Bricklet) RegisterLinearAccelerationCallback(fn func(int16, int16, int16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 14 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var x int16
		var y int16
		var z int16
		binary.Read(buf, binary.LittleEndian, &x)
		binary.Read(buf, binary.LittleEndian, &y)
		binary.Read(buf, binary.LittleEndian, &z)
		fn(x, y, z)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackLinearAcceleration), wrapper)
}

// Remove a registered Linear Acceleration callback.
func (device *IMUV3Bricklet) DeregisterLinearAccelerationCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackLinearAcceleration), registrationId)
}


// This callback is triggered periodically with the period that is set by
// SetGravityVectorCallbackConfiguration. The parameters gravity vector
// for the x, y and z axis.
func (device *IMUV3Bricklet) RegisterGravityVectorCallback(fn func(int16, int16, int16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 14 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var x int16
		var y int16
		var z int16
		binary.Read(buf, binary.LittleEndian, &x)
		binary.Read(buf, binary.LittleEndian, &y)
		binary.Read(buf, binary.LittleEndian, &z)
		fn(x, y, z)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackGravityVector), wrapper)
}

// Remove a registered Gravity Vector callback.
func (device *IMUV3Bricklet) DeregisterGravityVectorCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackGravityVector), registrationId)
}


// This callback is triggered periodically with the period that is set by
// SetOrientationCallbackConfiguration. The parameters are the orientation
// (heading (yaw), roll, pitch) of the IMU Brick in Euler angles. See
// GetOrientation for details.
func (device *IMUV3Bricklet) RegisterOrientationCallback(fn func(int16, int16, int16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 14 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var heading int16
		var roll int16
		var pitch int16
		binary.Read(buf, binary.LittleEndian, &heading)
		binary.Read(buf, binary.LittleEndian, &roll)
		binary.Read(buf, binary.LittleEndian, &pitch)
		fn(heading, roll, pitch)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackOrientation), wrapper)
}

// Remove a registered Orientation callback.
func (device *IMUV3Bricklet) DeregisterOrientationCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackOrientation), registrationId)
}


// This callback is triggered periodically with the period that is set by
// SetQuaternionCallbackConfiguration. The parameters are the orientation
// (w, x, y, z) of the IMU Brick in quaternions. See GetQuaternion
// for details.
func (device *IMUV3Bricklet) RegisterQuaternionCallback(fn func(int16, int16, int16, int16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 16 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var w int16
		var x int16
		var y int16
		var z int16
		binary.Read(buf, binary.LittleEndian, &w)
		binary.Read(buf, binary.LittleEndian, &x)
		binary.Read(buf, binary.LittleEndian, &y)
		binary.Read(buf, binary.LittleEndian, &z)
		fn(w, x, y, z)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackQuaternion), wrapper)
}

// Remove a registered Quaternion callback.
func (device *IMUV3Bricklet) DeregisterQuaternionCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackQuaternion), registrationId)
}


// This callback is triggered periodically with the period that is set by
// SetAllDataCallbackConfiguration. The parameters are as for
// GetAllData.
func (device *IMUV3Bricklet) RegisterAllDataCallback(fn func([3]int16, [3]int16, [3]int16, [3]int16, [4]int16, [3]int16, [3]int16, int8, uint8)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 54 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var acceleration [3]int16
		var magneticField [3]int16
		var angularVelocity [3]int16
		var eulerAngle [3]int16
		var quaternion [4]int16
		var linearAcceleration [3]int16
		var gravityVector [3]int16
		var temperature int8
		var calibrationStatus uint8
		binary.Read(buf, binary.LittleEndian, &acceleration)
		binary.Read(buf, binary.LittleEndian, &magneticField)
		binary.Read(buf, binary.LittleEndian, &angularVelocity)
		binary.Read(buf, binary.LittleEndian, &eulerAngle)
		binary.Read(buf, binary.LittleEndian, &quaternion)
		binary.Read(buf, binary.LittleEndian, &linearAcceleration)
		binary.Read(buf, binary.LittleEndian, &gravityVector)
		binary.Read(buf, binary.LittleEndian, &temperature)
		binary.Read(buf, binary.LittleEndian, &calibrationStatus)
		fn(acceleration, magneticField, angularVelocity, eulerAngle, quaternion, linearAcceleration, gravityVector, temperature, calibrationStatus)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackAllData), wrapper)
}

// Remove a registered All Data callback.
func (device *IMUV3Bricklet) DeregisterAllDataCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackAllData), registrationId)
}


// Returns the calibrated acceleration from the accelerometer for the
// x, y and z axis. The acceleration is in the range configured with
// SetSensorConfiguration.
// 
// If you want to get the acceleration periodically, it is recommended
// to use the RegisterAccelerationCallback callback and set the period with
// SetAccelerationCallbackConfiguration.
func (device *IMUV3Bricklet) GetAcceleration() (x int16, y int16, z int16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetAcceleration), buf.Bytes())
	if err != nil {
		return x, y, z, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 14 {
			return x, y, z, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 14)
		}

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
// x, y and z axis.
// 
// If you want to get the magnetic field periodically, it is recommended
// to use the RegisterMagneticFieldCallback callback and set the period with
// SetMagneticFieldCallbackConfiguration.
func (device *IMUV3Bricklet) GetMagneticField() (x int16, y int16, z int16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetMagneticField), buf.Bytes())
	if err != nil {
		return x, y, z, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 14 {
			return x, y, z, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 14)
		}

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
// x, y and z axis. The angular velocity is in the range configured with
// SetSensorConfiguration.
// 
// If you want to get the angular velocity periodically, it is recommended
// to use the RegisterAngularVelocityCallback acallback nd set the period with
// SetAngularVelocityCallbackConfiguration.
func (device *IMUV3Bricklet) GetAngularVelocity() (x int16, y int16, z int16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetAngularVelocity), buf.Bytes())
	if err != nil {
		return x, y, z, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 14 {
			return x, y, z, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 14)
		}

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

// Returns the temperature of the IMU Brick.
// The temperature is measured in the core of the BNO055 IC, it is not the
// ambient temperature
func (device *IMUV3Bricklet) GetTemperature() (temperature int8, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetTemperature), buf.Bytes())
	if err != nil {
		return temperature, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return temperature, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return temperature, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &temperature)

	}

	return temperature, nil
}

// Returns the current orientation (heading, roll, pitch) of the IMU Brick as
// independent Euler angles. Note that Euler angles always
// experience a https://en.wikipedia.org/wiki/Gimbal_lock.
// We recommend that you use quaternions instead, if you need the absolute
// orientation.
// 
// If you want to get the orientation periodically, it is recommended
// to use the RegisterOrientationCallback callback and set the period with
// SetOrientationCallbackConfiguration.
func (device *IMUV3Bricklet) GetOrientation() (heading int16, roll int16, pitch int16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetOrientation), buf.Bytes())
	if err != nil {
		return heading, roll, pitch, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 14 {
			return heading, roll, pitch, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 14)
		}

		if header.ErrorCode != 0 {
			return heading, roll, pitch, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &heading)
		binary.Read(resultBuf, binary.LittleEndian, &roll)
		binary.Read(resultBuf, binary.LittleEndian, &pitch)

	}

	return heading, roll, pitch, nil
}

// Returns the linear acceleration of the IMU Brick for the
// x, y and z axis. The acceleration is in the range configured with
// SetSensorConfiguration.
// 
// The linear acceleration is the acceleration in each of the three
// axis of the IMU Brick with the influences of gravity removed.
// 
// It is also possible to get the gravity vector with the influence of linear
// acceleration removed, see GetGravityVector.
// 
// If you want to get the linear acceleration periodically, it is recommended
// to use the RegisterLinearAccelerationCallback callback and set the period with
// SetLinearAccelerationCallbackConfiguration.
func (device *IMUV3Bricklet) GetLinearAcceleration() (x int16, y int16, z int16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetLinearAcceleration), buf.Bytes())
	if err != nil {
		return x, y, z, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 14 {
			return x, y, z, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 14)
		}

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

// Returns the current gravity vector of the IMU Brick for the
// x, y and z axis.
// 
// The gravity vector is the acceleration that occurs due to gravity.
// Influences of additional linear acceleration are removed.
// 
// It is also possible to get the linear acceleration with the influence
// of gravity removed, see GetLinearAcceleration.
// 
// If you want to get the gravity vector periodically, it is recommended
// to use the RegisterGravityVectorCallback callback and set the period with
// SetGravityVectorCallbackConfiguration.
func (device *IMUV3Bricklet) GetGravityVector() (x int16, y int16, z int16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetGravityVector), buf.Bytes())
	if err != nil {
		return x, y, z, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 14 {
			return x, y, z, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 14)
		}

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

// Returns the current orientation (w, x, y, z) of the IMU Brick as
// https://en.wikipedia.org/wiki/Quaternions_and_spatial_rotation.
// 
// You have to divide the return values by 16383 (14 bit) to get
// the usual range of -1.0 to +1.0 for quaternions.
// 
// If you want to get the quaternions periodically, it is recommended
// to use the RegisterQuaternionCallback callback and set the period with
// SetQuaternionCallbackConfiguration.
func (device *IMUV3Bricklet) GetQuaternion() (w int16, x int16, y int16, z int16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetQuaternion), buf.Bytes())
	if err != nil {
		return w, x, y, z, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 16 {
			return w, x, y, z, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 16)
		}

		if header.ErrorCode != 0 {
			return w, x, y, z, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &w)
		binary.Read(resultBuf, binary.LittleEndian, &x)
		binary.Read(resultBuf, binary.LittleEndian, &y)
		binary.Read(resultBuf, binary.LittleEndian, &z)

	}

	return w, x, y, z, nil
}

// Return all of the available data of the IMU Brick.
// 
// * acceleration (see GetAcceleration)
// * magnetic field (see GetMagneticField)
// * angular velocity (see GetAngularVelocity)
// * Euler angles (see GetOrientation)
// * quaternion (see GetQuaternion)
// * linear acceleration (see GetLinearAcceleration)
// * gravity vector (see GetGravityVector)
// * temperature (see GetTemperature)
// * calibration status (see below)
// 
// The calibration status consists of four pairs of two bits. Each pair
// of bits represents the status of the current calibration.
// 
// * bit 0-1: Magnetometer
// * bit 2-3: Accelerometer
// * bit 4-5: Gyroscope
// * bit 6-7: System
// 
// A value of 0 means for not calibrated and a value of 3 means
// fully calibrated. In your program you should always be able to
// ignore the calibration status, it is used by the calibration
// window of the Brick Viewer and it can be ignored after the first
// calibration. See the documentation in the calibration window for
// more information regarding the calibration of the IMU Brick.
// 
// If you want to get the data periodically, it is recommended
// to use the RegisterAllDataCallback callback and set the period with
// SetAllDataCallbackConfiguration.
func (device *IMUV3Bricklet) GetAllData() (acceleration [3]int16, magneticField [3]int16, angularVelocity [3]int16, eulerAngle [3]int16, quaternion [4]int16, linearAcceleration [3]int16, gravityVector [3]int16, temperature int8, calibrationStatus uint8, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetAllData), buf.Bytes())
	if err != nil {
		return acceleration, magneticField, angularVelocity, eulerAngle, quaternion, linearAcceleration, gravityVector, temperature, calibrationStatus, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 54 {
			return acceleration, magneticField, angularVelocity, eulerAngle, quaternion, linearAcceleration, gravityVector, temperature, calibrationStatus, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 54)
		}

		if header.ErrorCode != 0 {
			return acceleration, magneticField, angularVelocity, eulerAngle, quaternion, linearAcceleration, gravityVector, temperature, calibrationStatus, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &acceleration)
		binary.Read(resultBuf, binary.LittleEndian, &magneticField)
		binary.Read(resultBuf, binary.LittleEndian, &angularVelocity)
		binary.Read(resultBuf, binary.LittleEndian, &eulerAngle)
		binary.Read(resultBuf, binary.LittleEndian, &quaternion)
		binary.Read(resultBuf, binary.LittleEndian, &linearAcceleration)
		binary.Read(resultBuf, binary.LittleEndian, &gravityVector)
		binary.Read(resultBuf, binary.LittleEndian, &temperature)
		binary.Read(resultBuf, binary.LittleEndian, &calibrationStatus)

	}

	return acceleration, magneticField, angularVelocity, eulerAngle, quaternion, linearAcceleration, gravityVector, temperature, calibrationStatus, nil
}

// A call of this function saves the current calibration to be used
// as a starting point for the next restart of continuous calibration
// of the IMU Brick.
// 
// A return value of *true* means that the calibration could be used and
// *false* means that it could not be used (this happens if the calibration
// status is not fully calibrated).
// 
// This function is used by the calibration window of the Brick Viewer, you
// should not need to call it in your program.
func (device *IMUV3Bricklet) SaveCalibration() (calibrationDone bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionSaveCalibration), buf.Bytes())
	if err != nil {
		return calibrationDone, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return calibrationDone, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return calibrationDone, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &calibrationDone)

	}

	return calibrationDone, nil
}

// Sets the available sensor configuration for the Magnetometer, Gyroscope and
// Accelerometer. The Accelerometer Range is user selectable in all fusion modes,
// all other configurations are auto-controlled in fusion mode.
//
// Associated constants:
//
//	* MagnetometerRate2Hz
//	* MagnetometerRate6Hz
//	* MagnetometerRate8Hz
//	* MagnetometerRate10Hz
//	* MagnetometerRate15Hz
//	* MagnetometerRate20Hz
//	* MagnetometerRate25Hz
//	* MagnetometerRate30Hz
//	* GyroscopeRange2000DPS
//	* GyroscopeRange1000DPS
//	* GyroscopeRange500DPS
//	* GyroscopeRange250DPS
//	* GyroscopeRange125DPS
//	* GyroscopeBandwidth523Hz
//	* GyroscopeBandwidth230Hz
//	* GyroscopeBandwidth116Hz
//	* GyroscopeBandwidth47Hz
//	* GyroscopeBandwidth23Hz
//	* GyroscopeBandwidth12Hz
//	* GyroscopeBandwidth64Hz
//	* GyroscopeBandwidth32Hz
//	* AccelerometerRange2G
//	* AccelerometerRange4G
//	* AccelerometerRange8G
//	* AccelerometerRange16G
//	* AccelerometerBandwidth781Hz
//	* AccelerometerBandwidth1563Hz
//	* AccelerometerBandwidth3125Hz
//	* AccelerometerBandwidth625Hz
//	* AccelerometerBandwidth125Hz
//	* AccelerometerBandwidth250Hz
//	* AccelerometerBandwidth500Hz
//	* AccelerometerBandwidth1000Hz
func (device *IMUV3Bricklet) SetSensorConfiguration(magnetometerRate MagnetometerRate, gyroscopeRange GyroscopeRange, gyroscopeBandwidth GyroscopeBandwidth, accelerometerRange AccelerometerRange, accelerometerBandwidth AccelerometerBandwidth) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, magnetometerRate);
	binary.Write(&buf, binary.LittleEndian, gyroscopeRange);
	binary.Write(&buf, binary.LittleEndian, gyroscopeBandwidth);
	binary.Write(&buf, binary.LittleEndian, accelerometerRange);
	binary.Write(&buf, binary.LittleEndian, accelerometerBandwidth);

	resultBytes, err := device.device.Set(uint8(FunctionSetSensorConfiguration), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])
		
	}

	return nil
}

// Returns the sensor configuration as set by SetSensorConfiguration.
//
// Associated constants:
//
//	* MagnetometerRate2Hz
//	* MagnetometerRate6Hz
//	* MagnetometerRate8Hz
//	* MagnetometerRate10Hz
//	* MagnetometerRate15Hz
//	* MagnetometerRate20Hz
//	* MagnetometerRate25Hz
//	* MagnetometerRate30Hz
//	* GyroscopeRange2000DPS
//	* GyroscopeRange1000DPS
//	* GyroscopeRange500DPS
//	* GyroscopeRange250DPS
//	* GyroscopeRange125DPS
//	* GyroscopeBandwidth523Hz
//	* GyroscopeBandwidth230Hz
//	* GyroscopeBandwidth116Hz
//	* GyroscopeBandwidth47Hz
//	* GyroscopeBandwidth23Hz
//	* GyroscopeBandwidth12Hz
//	* GyroscopeBandwidth64Hz
//	* GyroscopeBandwidth32Hz
//	* AccelerometerRange2G
//	* AccelerometerRange4G
//	* AccelerometerRange8G
//	* AccelerometerRange16G
//	* AccelerometerBandwidth781Hz
//	* AccelerometerBandwidth1563Hz
//	* AccelerometerBandwidth3125Hz
//	* AccelerometerBandwidth625Hz
//	* AccelerometerBandwidth125Hz
//	* AccelerometerBandwidth250Hz
//	* AccelerometerBandwidth500Hz
//	* AccelerometerBandwidth1000Hz
func (device *IMUV3Bricklet) GetSensorConfiguration() (magnetometerRate MagnetometerRate, gyroscopeRange GyroscopeRange, gyroscopeBandwidth GyroscopeBandwidth, accelerometerRange AccelerometerRange, accelerometerBandwidth AccelerometerBandwidth, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetSensorConfiguration), buf.Bytes())
	if err != nil {
		return magnetometerRate, gyroscopeRange, gyroscopeBandwidth, accelerometerRange, accelerometerBandwidth, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 13 {
			return magnetometerRate, gyroscopeRange, gyroscopeBandwidth, accelerometerRange, accelerometerBandwidth, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 13)
		}

		if header.ErrorCode != 0 {
			return magnetometerRate, gyroscopeRange, gyroscopeBandwidth, accelerometerRange, accelerometerBandwidth, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &magnetometerRate)
		binary.Read(resultBuf, binary.LittleEndian, &gyroscopeRange)
		binary.Read(resultBuf, binary.LittleEndian, &gyroscopeBandwidth)
		binary.Read(resultBuf, binary.LittleEndian, &accelerometerRange)
		binary.Read(resultBuf, binary.LittleEndian, &accelerometerBandwidth)

	}

	return magnetometerRate, gyroscopeRange, gyroscopeBandwidth, accelerometerRange, accelerometerBandwidth, nil
}

// If the fusion mode is turned off, the functions GetAcceleration,
// GetMagneticField and GetAngularVelocity return uncalibrated
// and uncompensated sensor data. All other sensor data getters return no data.
// 
// Since firmware version 2.0.6 you can also use a fusion mode without magnetometer.
// In this mode the calculated orientation is relative (with magnetometer it is
// absolute with respect to the earth). However, the calculation can't be influenced
// by spurious magnetic fields.
// 
// Since firmware version 2.0.13 you can also use a fusion mode without fast
// magnetometer calibration. This mode is the same as the normal fusion mode,
// but the fast magnetometer calibration is turned off. So to find the orientation
// the first time will likely take longer, but small magnetic influences might
// not affect the automatic calibration as much.
//
// Associated constants:
//
//	* SensorFusionOff
//	* SensorFusionOn
//	* SensorFusionOnWithoutMagnetometer
//	* SensorFusionOnWithoutFastMagnetometerCalibration
func (device *IMUV3Bricklet) SetSensorFusionMode(mode SensorFusion) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, mode);

	resultBytes, err := device.device.Set(uint8(FunctionSetSensorFusionMode), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])
		
	}

	return nil
}

// Returns the sensor fusion mode as set by SetSensorFusionMode.
//
// Associated constants:
//
//	* SensorFusionOff
//	* SensorFusionOn
//	* SensorFusionOnWithoutMagnetometer
//	* SensorFusionOnWithoutFastMagnetometerCalibration
func (device *IMUV3Bricklet) GetSensorFusionMode() (mode SensorFusion, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetSensorFusionMode), buf.Bytes())
	if err != nil {
		return mode, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return mode, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return mode, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &mode)

	}

	return mode, nil
}

// The period is the period with which the RegisterAccelerationCallback callback
// is triggered periodically. A value of 0 turns the callback off.
// 
// If the `value has to change`-parameter is set to true, the callback is only
// triggered after the value has changed. If the value didn't change within the
// period, the callback is triggered immediately on change.
// 
// If it is set to false, the callback is continuously triggered with the period,
// independent of the value.
func (device *IMUV3Bricklet) SetAccelerationCallbackConfiguration(period uint32, valueHasToChange bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);
	binary.Write(&buf, binary.LittleEndian, valueHasToChange);

	resultBytes, err := device.device.Set(uint8(FunctionSetAccelerationCallbackConfiguration), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])
		
	}

	return nil
}

// Returns the callback configuration as set by SetAccelerationCallbackConfiguration.
func (device *IMUV3Bricklet) GetAccelerationCallbackConfiguration() (period uint32, valueHasToChange bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetAccelerationCallbackConfiguration), buf.Bytes())
	if err != nil {
		return period, valueHasToChange, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 13 {
			return period, valueHasToChange, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 13)
		}

		if header.ErrorCode != 0 {
			return period, valueHasToChange, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &period)
		binary.Read(resultBuf, binary.LittleEndian, &valueHasToChange)

	}

	return period, valueHasToChange, nil
}

// The period is the period with which the RegisterMagneticFieldCallback callback
// is triggered periodically. A value of 0 turns the callback off.
// 
// If the `value has to change`-parameter is set to true, the callback is only
// triggered after the value has changed. If the value didn't change within the
// period, the callback is triggered immediately on change.
// 
// If it is set to false, the callback is continuously triggered with the period,
// independent of the value.
func (device *IMUV3Bricklet) SetMagneticFieldCallbackConfiguration(period uint32, valueHasToChange bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);
	binary.Write(&buf, binary.LittleEndian, valueHasToChange);

	resultBytes, err := device.device.Set(uint8(FunctionSetMagneticFieldCallbackConfiguration), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])
		
	}

	return nil
}

// Returns the callback configuration as set by SetMagneticFieldCallbackConfiguration.
func (device *IMUV3Bricklet) GetMagneticFieldCallbackConfiguration() (period uint32, valueHasToChange bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetMagneticFieldCallbackConfiguration), buf.Bytes())
	if err != nil {
		return period, valueHasToChange, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 13 {
			return period, valueHasToChange, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 13)
		}

		if header.ErrorCode != 0 {
			return period, valueHasToChange, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &period)
		binary.Read(resultBuf, binary.LittleEndian, &valueHasToChange)

	}

	return period, valueHasToChange, nil
}

// The period is the period with which the RegisterAngularVelocityCallback callback
// is triggered periodically. A value of 0 turns the callback off.
// 
// If the `value has to change`-parameter is set to true, the callback is only
// triggered after the value has changed. If the value didn't change within the
// period, the callback is triggered immediately on change.
// 
// If it is set to false, the callback is continuously triggered with the period,
// independent of the value.
func (device *IMUV3Bricklet) SetAngularVelocityCallbackConfiguration(period uint32, valueHasToChange bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);
	binary.Write(&buf, binary.LittleEndian, valueHasToChange);

	resultBytes, err := device.device.Set(uint8(FunctionSetAngularVelocityCallbackConfiguration), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])
		
	}

	return nil
}

// Returns the callback configuration as set by SetAngularVelocityCallbackConfiguration.
func (device *IMUV3Bricklet) GetAngularVelocityCallbackConfiguration() (period uint32, valueHasToChange bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetAngularVelocityCallbackConfiguration), buf.Bytes())
	if err != nil {
		return period, valueHasToChange, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 13 {
			return period, valueHasToChange, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 13)
		}

		if header.ErrorCode != 0 {
			return period, valueHasToChange, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &period)
		binary.Read(resultBuf, binary.LittleEndian, &valueHasToChange)

	}

	return period, valueHasToChange, nil
}

// The period is the period with which the RegisterTemperatureCallback callback
// is triggered periodically. A value of 0 turns the callback off.
// 
// If the `value has to change`-parameter is set to true, the callback is only
// triggered after the value has changed. If the value didn't change within the
// period, the callback is triggered immediately on change.
// 
// If it is set to false, the callback is continuously triggered with the period,
// independent of the value.
func (device *IMUV3Bricklet) SetTemperatureCallbackConfiguration(period uint32, valueHasToChange bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);
	binary.Write(&buf, binary.LittleEndian, valueHasToChange);

	resultBytes, err := device.device.Set(uint8(FunctionSetTemperatureCallbackConfiguration), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])
		
	}

	return nil
}

// Returns the callback configuration as set by SetTemperatureCallbackConfiguration.
func (device *IMUV3Bricklet) GetTemperatureCallbackConfiguration() (period uint32, valueHasToChange bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetTemperatureCallbackConfiguration), buf.Bytes())
	if err != nil {
		return period, valueHasToChange, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 13 {
			return period, valueHasToChange, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 13)
		}

		if header.ErrorCode != 0 {
			return period, valueHasToChange, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &period)
		binary.Read(resultBuf, binary.LittleEndian, &valueHasToChange)

	}

	return period, valueHasToChange, nil
}

// The period is the period with which the RegisterOrientationCallback callback
// is triggered periodically. A value of 0 turns the callback off.
// 
// If the `value has to change`-parameter is set to true, the callback is only
// triggered after the value has changed. If the value didn't change within the
// period, the callback is triggered immediately on change.
// 
// If it is set to false, the callback is continuously triggered with the period,
// independent of the value.
func (device *IMUV3Bricklet) SetOrientationCallbackConfiguration(period uint32, valueHasToChange bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);
	binary.Write(&buf, binary.LittleEndian, valueHasToChange);

	resultBytes, err := device.device.Set(uint8(FunctionSetOrientationCallbackConfiguration), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])
		
	}

	return nil
}

// Returns the callback configuration as set by SetOrientationCallbackConfiguration.
func (device *IMUV3Bricklet) GetOrientationCallbackConfiguration() (period uint32, valueHasToChange bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetOrientationCallbackConfiguration), buf.Bytes())
	if err != nil {
		return period, valueHasToChange, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 13 {
			return period, valueHasToChange, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 13)
		}

		if header.ErrorCode != 0 {
			return period, valueHasToChange, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &period)
		binary.Read(resultBuf, binary.LittleEndian, &valueHasToChange)

	}

	return period, valueHasToChange, nil
}

// The period is the period with which the RegisterLinearAccelerationCallback callback
// is triggered periodically. A value of 0 turns the callback off.
// 
// If the `value has to change`-parameter is set to true, the callback is only
// triggered after the value has changed. If the value didn't change within the
// period, the callback is triggered immediately on change.
// 
// If it is set to false, the callback is continuously triggered with the period,
// independent of the value.
func (device *IMUV3Bricklet) SetLinearAccelerationCallbackConfiguration(period uint32, valueHasToChange bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);
	binary.Write(&buf, binary.LittleEndian, valueHasToChange);

	resultBytes, err := device.device.Set(uint8(FunctionSetLinearAccelerationCallbackConfiguration), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])
		
	}

	return nil
}

// Returns the callback configuration as set by SetLinearAccelerationCallbackConfiguration.
func (device *IMUV3Bricklet) GetLinearAccelerationCallbackConfiguration() (period uint32, valueHasToChange bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetLinearAccelerationCallbackConfiguration), buf.Bytes())
	if err != nil {
		return period, valueHasToChange, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 13 {
			return period, valueHasToChange, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 13)
		}

		if header.ErrorCode != 0 {
			return period, valueHasToChange, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &period)
		binary.Read(resultBuf, binary.LittleEndian, &valueHasToChange)

	}

	return period, valueHasToChange, nil
}

// The period is the period with which the RegisterGravityVectorCallback callback
// is triggered periodically. A value of 0 turns the callback off.
// 
// If the `value has to change`-parameter is set to true, the callback is only
// triggered after the value has changed. If the value didn't change within the
// period, the callback is triggered immediately on change.
// 
// If it is set to false, the callback is continuously triggered with the period,
// independent of the value.
func (device *IMUV3Bricklet) SetGravityVectorCallbackConfiguration(period uint32, valueHasToChange bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);
	binary.Write(&buf, binary.LittleEndian, valueHasToChange);

	resultBytes, err := device.device.Set(uint8(FunctionSetGravityVectorCallbackConfiguration), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])
		
	}

	return nil
}

// Returns the callback configuration as set by SetGravityVectorCallbackConfiguration.
func (device *IMUV3Bricklet) GetGravityVectorCallbackConfiguration() (period uint32, valueHasToChange bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetGravityVectorCallbackConfiguration), buf.Bytes())
	if err != nil {
		return period, valueHasToChange, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 13 {
			return period, valueHasToChange, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 13)
		}

		if header.ErrorCode != 0 {
			return period, valueHasToChange, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &period)
		binary.Read(resultBuf, binary.LittleEndian, &valueHasToChange)

	}

	return period, valueHasToChange, nil
}

// The period is the period with which the RegisterQuaternionCallback callback
// is triggered periodically. A value of 0 turns the callback off.
// 
// If the `value has to change`-parameter is set to true, the callback is only
// triggered after the value has changed. If the value didn't change within the
// period, the callback is triggered immediately on change.
// 
// If it is set to false, the callback is continuously triggered with the period,
// independent of the value.
func (device *IMUV3Bricklet) SetQuaternionCallbackConfiguration(period uint32, valueHasToChange bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);
	binary.Write(&buf, binary.LittleEndian, valueHasToChange);

	resultBytes, err := device.device.Set(uint8(FunctionSetQuaternionCallbackConfiguration), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])
		
	}

	return nil
}

// Returns the callback configuration as set by SetQuaternionCallbackConfiguration.
func (device *IMUV3Bricklet) GetQuaternionCallbackConfiguration() (period uint32, valueHasToChange bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetQuaternionCallbackConfiguration), buf.Bytes())
	if err != nil {
		return period, valueHasToChange, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 13 {
			return period, valueHasToChange, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 13)
		}

		if header.ErrorCode != 0 {
			return period, valueHasToChange, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &period)
		binary.Read(resultBuf, binary.LittleEndian, &valueHasToChange)

	}

	return period, valueHasToChange, nil
}

// The period is the period with which the RegisterAllDataCallback callback
// is triggered periodically. A value of 0 turns the callback off.
// 
// If the `value has to change`-parameter is set to true, the callback is only
// triggered after the value has changed. If the value didn't change within the
// period, the callback is triggered immediately on change.
// 
// If it is set to false, the callback is continuously triggered with the period,
// independent of the value.
func (device *IMUV3Bricklet) SetAllDataCallbackConfiguration(period uint32, valueHasToChange bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);
	binary.Write(&buf, binary.LittleEndian, valueHasToChange);

	resultBytes, err := device.device.Set(uint8(FunctionSetAllDataCallbackConfiguration), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])
		
	}

	return nil
}

// Returns the callback configuration as set by SetAllDataCallbackConfiguration.
func (device *IMUV3Bricklet) GetAllDataCallbackConfiguration() (period uint32, valueHasToChange bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetAllDataCallbackConfiguration), buf.Bytes())
	if err != nil {
		return period, valueHasToChange, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 13 {
			return period, valueHasToChange, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 13)
		}

		if header.ErrorCode != 0 {
			return period, valueHasToChange, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &period)
		binary.Read(resultBuf, binary.LittleEndian, &valueHasToChange)

	}

	return period, valueHasToChange, nil
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
// The errors counts are for errors that occur on the Bricklet side. All
// Bricks have a similar function that returns the errors on the Brick side.
func (device *IMUV3Bricklet) GetSPITFPErrorCount() (errorCountAckChecksum uint32, errorCountMessageChecksum uint32, errorCountFrame uint32, errorCountOverflow uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetSPITFPErrorCount), buf.Bytes())
	if err != nil {
		return errorCountAckChecksum, errorCountMessageChecksum, errorCountFrame, errorCountOverflow, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 24 {
			return errorCountAckChecksum, errorCountMessageChecksum, errorCountFrame, errorCountOverflow, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 24)
		}

		if header.ErrorCode != 0 {
			return errorCountAckChecksum, errorCountMessageChecksum, errorCountFrame, errorCountOverflow, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &errorCountAckChecksum)
		binary.Read(resultBuf, binary.LittleEndian, &errorCountMessageChecksum)
		binary.Read(resultBuf, binary.LittleEndian, &errorCountFrame)
		binary.Read(resultBuf, binary.LittleEndian, &errorCountOverflow)

	}

	return errorCountAckChecksum, errorCountMessageChecksum, errorCountFrame, errorCountOverflow, nil
}

// Sets the bootloader mode and returns the status after the requested
// mode change was instigated.
// 
// You can change from bootloader mode to firmware mode and vice versa. A change
// from bootloader mode to firmware mode will only take place if the entry function,
// device identifier and CRC are present and correct.
// 
// This function is used by Brick Viewer during flashing. It should not be
// necessary to call it in a normal user program.
//
// Associated constants:
//
//	* BootloaderModeBootloader
//	* BootloaderModeFirmware
//	* BootloaderModeBootloaderWaitForReboot
//	* BootloaderModeFirmwareWaitForReboot
//	* BootloaderModeFirmwareWaitForEraseAndReboot
//	* BootloaderStatusOK
//	* BootloaderStatusInvalidMode
//	* BootloaderStatusNoChange
//	* BootloaderStatusEntryFunctionNotPresent
//	* BootloaderStatusDeviceIdentifierIncorrect
//	* BootloaderStatusCRCMismatch
func (device *IMUV3Bricklet) SetBootloaderMode(mode BootloaderMode) (status BootloaderStatus, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, mode);

	resultBytes, err := device.device.Get(uint8(FunctionSetBootloaderMode), buf.Bytes())
	if err != nil {
		return status, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return status, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return status, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &status)

	}

	return status, nil
}

// Returns the current bootloader mode, see SetBootloaderMode.
//
// Associated constants:
//
//	* BootloaderModeBootloader
//	* BootloaderModeFirmware
//	* BootloaderModeBootloaderWaitForReboot
//	* BootloaderModeFirmwareWaitForReboot
//	* BootloaderModeFirmwareWaitForEraseAndReboot
func (device *IMUV3Bricklet) GetBootloaderMode() (mode BootloaderMode, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetBootloaderMode), buf.Bytes())
	if err != nil {
		return mode, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return mode, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return mode, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &mode)

	}

	return mode, nil
}

// Sets the firmware pointer for WriteFirmware. The pointer has
// to be increased by chunks of size 64. The data is written to flash
// every 4 chunks (which equals to one page of size 256).
// 
// This function is used by Brick Viewer during flashing. It should not be
// necessary to call it in a normal user program.
func (device *IMUV3Bricklet) SetWriteFirmwarePointer(pointer uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, pointer);

	resultBytes, err := device.device.Set(uint8(FunctionSetWriteFirmwarePointer), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])
		
	}

	return nil
}

// Writes 64 Bytes of firmware at the position as written by
// SetWriteFirmwarePointer before. The firmware is written
// to flash every 4 chunks.
// 
// You can only write firmware in bootloader mode.
// 
// This function is used by Brick Viewer during flashing. It should not be
// necessary to call it in a normal user program.
func (device *IMUV3Bricklet) WriteFirmware(data [64]uint8) (status uint8, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, data);

	resultBytes, err := device.device.Get(uint8(FunctionWriteFirmware), buf.Bytes())
	if err != nil {
		return status, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return status, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return status, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &status)

	}

	return status, nil
}

// Sets the status LED configuration. By default the LED shows
// communication traffic between Brick and Bricklet, it flickers once
// for every 10 received data packets.
// 
// You can also turn the LED permanently on/off or show a heartbeat.
// 
// If the Bricklet is in bootloader mode, the LED is will show heartbeat by default.
//
// Associated constants:
//
//	* StatusLEDConfigOff
//	* StatusLEDConfigOn
//	* StatusLEDConfigShowHeartbeat
//	* StatusLEDConfigShowStatus
func (device *IMUV3Bricklet) SetStatusLEDConfig(config StatusLEDConfig) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, config);

	resultBytes, err := device.device.Set(uint8(FunctionSetStatusLEDConfig), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])
		
	}

	return nil
}

// Returns the configuration as set by SetStatusLEDConfig
//
// Associated constants:
//
//	* StatusLEDConfigOff
//	* StatusLEDConfigOn
//	* StatusLEDConfigShowHeartbeat
//	* StatusLEDConfigShowStatus
func (device *IMUV3Bricklet) GetStatusLEDConfig() (config StatusLEDConfig, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetStatusLEDConfig), buf.Bytes())
	if err != nil {
		return config, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return config, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return config, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &config)

	}

	return config, nil
}

// Returns the temperature as measured inside the microcontroller. The
// value returned is not the ambient temperature!
// 
// The temperature is only proportional to the real temperature and it has bad
// accuracy. Practically it is only useful as an indicator for
// temperature changes.
func (device *IMUV3Bricklet) GetChipTemperature() (temperature int16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetChipTemperature), buf.Bytes())
	if err != nil {
		return temperature, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 10 {
			return temperature, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		if header.ErrorCode != 0 {
			return temperature, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &temperature)

	}

	return temperature, nil
}

// Calling this function will reset the Bricklet. All configurations
// will be lost.
// 
// After a reset you have to create new device objects,
// calling functions on the existing ones will result in
// undefined behavior!
func (device *IMUV3Bricklet) Reset() (err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Set(uint8(FunctionReset), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])
		
	}

	return nil
}

// Writes a new UID into flash. If you want to set a new UID
// you have to decode the Base58 encoded UID string into an
// integer first.
// 
// We recommend that you use Brick Viewer to change the UID.
func (device *IMUV3Bricklet) WriteUID(uid uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, uid);

	resultBytes, err := device.device.Set(uint8(FunctionWriteUID), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])
		
	}

	return nil
}

// Returns the current UID as an integer. Encode as
// Base58 to get the usual string version.
func (device *IMUV3Bricklet) ReadUID() (uid uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionReadUID), buf.Bytes())
	if err != nil {
		return uid, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 12 {
			return uid, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		if header.ErrorCode != 0 {
			return uid, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &uid)

	}

	return uid, nil
}

// Returns the UID, the UID where the Bricklet is connected to,
// the position, the hardware and firmware version as well as the
// device identifier.
// 
// The position can be 'a', 'b', 'c', 'd', 'e', 'f', 'g' or 'h' (Bricklet Port).
// A Bricklet connected to an `Isolator Bricklet <isolator_bricklet>` is always at
// position 'z'.
// 
// The device identifier numbers can be found `here <device_identifier>`.
// |device_identifier_constant|
func (device *IMUV3Bricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetIdentity), buf.Bytes())
	if err != nil {
		return uid, connectedUid, position, hardwareVersion, firmwareVersion, deviceIdentifier, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 33 {
			return uid, connectedUid, position, hardwareVersion, firmwareVersion, deviceIdentifier, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 33)
		}

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
