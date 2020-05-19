/* ***********************************************************
 * This file was automatically generated on 2020-05-19.      *
 *                                                           *
 * Go Bindings Version 2.0.8                                 *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


// Full fledged AHRS with 9 degrees of freedom.
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricks/IMUV2_Brick_Go.html.
package imu_v2_brick

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
	FunctionLedsOn Function = 10
	FunctionLedsOff Function = 11
	FunctionAreLedsOn Function = 12
	FunctionSaveCalibration Function = 13
	FunctionSetAccelerationPeriod Function = 14
	FunctionGetAccelerationPeriod Function = 15
	FunctionSetMagneticFieldPeriod Function = 16
	FunctionGetMagneticFieldPeriod Function = 17
	FunctionSetAngularVelocityPeriod Function = 18
	FunctionGetAngularVelocityPeriod Function = 19
	FunctionSetTemperaturePeriod Function = 20
	FunctionGetTemperaturePeriod Function = 21
	FunctionSetOrientationPeriod Function = 22
	FunctionGetOrientationPeriod Function = 23
	FunctionSetLinearAccelerationPeriod Function = 24
	FunctionGetLinearAccelerationPeriod Function = 25
	FunctionSetGravityVectorPeriod Function = 26
	FunctionGetGravityVectorPeriod Function = 27
	FunctionSetQuaternionPeriod Function = 28
	FunctionGetQuaternionPeriod Function = 29
	FunctionSetAllDataPeriod Function = 30
	FunctionGetAllDataPeriod Function = 31
	FunctionSetSensorConfiguration Function = 41
	FunctionGetSensorConfiguration Function = 42
	FunctionSetSensorFusionMode Function = 43
	FunctionGetSensorFusionMode Function = 44
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
	FunctionWriteBrickletPlugin Function = 246
	FunctionReadBrickletPlugin Function = 247
	FunctionGetIdentity Function = 255
	FunctionCallbackAcceleration Function = 32
	FunctionCallbackMagneticField Function = 33
	FunctionCallbackAngularVelocity Function = 34
	FunctionCallbackTemperature Function = 35
	FunctionCallbackLinearAcceleration Function = 36
	FunctionCallbackGravityVector Function = 37
	FunctionCallbackOrientation Function = 38
	FunctionCallbackQuaternion Function = 39
	FunctionCallbackAllData Function = 40
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
	//Deprecated: Use 7_81Hz instead.
	AccelerometerBandwidth781Hz AccelerometerBandwidth = 0
	AccelerometerBandwidth7_81Hz AccelerometerBandwidth = 0
	//Deprecated: Use 15_63Hz instead.
	AccelerometerBandwidth1563Hz AccelerometerBandwidth = 1
	AccelerometerBandwidth15_63Hz AccelerometerBandwidth = 1
	//Deprecated: Use 31_25Hz instead.
	AccelerometerBandwidth3125Hz AccelerometerBandwidth = 2
	AccelerometerBandwidth31_25Hz AccelerometerBandwidth = 2
	//Deprecated: Use 62_5Hz instead.
	AccelerometerBandwidth625Hz AccelerometerBandwidth = 3
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

type IMUV2Brick struct {
	device Device
}
const DeviceIdentifier = 18
const DeviceDisplayName = "IMU Brick 2.0"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (IMUV2Brick, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{ 2,0,3 }, uid, &internalIPCon, 0, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return IMUV2Brick{}, err
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
	dev.ResponseExpected[FunctionLedsOn] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionLedsOff] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionAreLedsOn] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSaveCalibration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetAccelerationPeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetAccelerationPeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetMagneticFieldPeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetMagneticFieldPeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetAngularVelocityPeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetAngularVelocityPeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetTemperaturePeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetTemperaturePeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetOrientationPeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetOrientationPeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetLinearAccelerationPeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetLinearAccelerationPeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetGravityVectorPeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetGravityVectorPeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetQuaternionPeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetQuaternionPeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetAllDataPeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetAllDataPeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetSensorConfiguration] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetSensorConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetSensorFusionMode] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetSensorFusionMode] = ResponseExpectedFlagAlwaysTrue;
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
	dev.ResponseExpected[FunctionWriteBrickletPlugin] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionReadBrickletPlugin] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetIdentity] = ResponseExpectedFlagAlwaysTrue;
	return IMUV2Brick{dev}, nil
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
func (device *IMUV2Brick) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *IMUV2Brick) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *IMUV2Brick) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *IMUV2Brick) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// This callback is triggered periodically with the period that is set by
// SetAccelerationPeriod. The parameters are the acceleration
// for the x, y and z axis.
func (device *IMUV2Brick) RegisterAccelerationCallback(fn func(int16, int16, int16)) uint64 {
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
func (device *IMUV2Brick) DeregisterAccelerationCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackAcceleration), registrationId)
}


// This callback is triggered periodically with the period that is set by
// SetMagneticFieldPeriod. The parameters are the magnetic
// field for the x, y and z axis.
func (device *IMUV2Brick) RegisterMagneticFieldCallback(fn func(int16, int16, int16)) uint64 {
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
func (device *IMUV2Brick) DeregisterMagneticFieldCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackMagneticField), registrationId)
}


// This callback is triggered periodically with the period that is set by
// SetAngularVelocityPeriod. The parameters are the angular
// velocity for the x, y and z axis.
func (device *IMUV2Brick) RegisterAngularVelocityCallback(fn func(int16, int16, int16)) uint64 {
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
func (device *IMUV2Brick) DeregisterAngularVelocityCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackAngularVelocity), registrationId)
}


// This callback is triggered periodically with the period that is set by
// SetTemperaturePeriod. The parameter is the temperature.
func (device *IMUV2Brick) RegisterTemperatureCallback(fn func(int8)) uint64 {
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
func (device *IMUV2Brick) DeregisterTemperatureCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackTemperature), registrationId)
}


// This callback is triggered periodically with the period that is set by
// SetLinearAccelerationPeriod. The parameters are the
// linear acceleration for the x, y and z axis.
func (device *IMUV2Brick) RegisterLinearAccelerationCallback(fn func(int16, int16, int16)) uint64 {
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
func (device *IMUV2Brick) DeregisterLinearAccelerationCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackLinearAcceleration), registrationId)
}


// This callback is triggered periodically with the period that is set by
// SetGravityVectorPeriod. The parameters gravity vector
// for the x, y and z axis.
func (device *IMUV2Brick) RegisterGravityVectorCallback(fn func(int16, int16, int16)) uint64 {
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
func (device *IMUV2Brick) DeregisterGravityVectorCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackGravityVector), registrationId)
}


// This callback is triggered periodically with the period that is set by
// SetOrientationPeriod. The parameters are the orientation
// (heading (yaw), roll, pitch) of the IMU Brick in Euler angles. See
// GetOrientation for details.
func (device *IMUV2Brick) RegisterOrientationCallback(fn func(int16, int16, int16)) uint64 {
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
func (device *IMUV2Brick) DeregisterOrientationCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackOrientation), registrationId)
}


// This callback is triggered periodically with the period that is set by
// SetQuaternionPeriod. The parameters are the orientation
// (w, x, y, z) of the IMU Brick in quaternions. See GetQuaternion
// for details.
func (device *IMUV2Brick) RegisterQuaternionCallback(fn func(int16, int16, int16, int16)) uint64 {
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
func (device *IMUV2Brick) DeregisterQuaternionCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackQuaternion), registrationId)
}


// This callback is triggered periodically with the period that is set by
// SetAllDataPeriod. The parameters are as for
// GetAllData.
func (device *IMUV2Brick) RegisterAllDataCallback(fn func([3]int16, [3]int16, [3]int16, [3]int16, [4]int16, [3]int16, [3]int16, int8, uint8)) uint64 {
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
func (device *IMUV2Brick) DeregisterAllDataCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackAllData), registrationId)
}


// Returns the calibrated acceleration from the accelerometer for the
// x, y and z axis. The acceleration is in the range configured with
// SetSensorConfiguration.
// 
// If you want to get the acceleration periodically, it is recommended
// to use the RegisterAccelerationCallback callback and set the period with
// SetAccelerationPeriod.
func (device *IMUV2Brick) GetAcceleration() (x int16, y int16, z int16, err error) {
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
// SetMagneticFieldPeriod.
func (device *IMUV2Brick) GetMagneticField() (x int16, y int16, z int16, err error) {
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
// SetAngularVelocityPeriod.
func (device *IMUV2Brick) GetAngularVelocity() (x int16, y int16, z int16, err error) {
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
func (device *IMUV2Brick) GetTemperature() (temperature int8, err error) {
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
// SetOrientationPeriod.
func (device *IMUV2Brick) GetOrientation() (heading int16, roll int16, pitch int16, err error) {
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
// SetLinearAccelerationPeriod.
func (device *IMUV2Brick) GetLinearAcceleration() (x int16, y int16, z int16, err error) {
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
// SetGravityVectorPeriod.
func (device *IMUV2Brick) GetGravityVector() (x int16, y int16, z int16, err error) {
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
// SetQuaternionPeriod.
func (device *IMUV2Brick) GetQuaternion() (w int16, x int16, y int16, z int16, err error) {
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
// SetAllDataPeriod.
func (device *IMUV2Brick) GetAllData() (acceleration [3]int16, magneticField [3]int16, angularVelocity [3]int16, eulerAngle [3]int16, quaternion [4]int16, linearAcceleration [3]int16, gravityVector [3]int16, temperature int8, calibrationStatus uint8, err error) {
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

// Turns the orientation and direction LEDs of the IMU Brick on.
func (device *IMUV2Brick) LedsOn() (err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Set(uint8(FunctionLedsOn), buf.Bytes())
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

// Turns the orientation and direction LEDs of the IMU Brick off.
func (device *IMUV2Brick) LedsOff() (err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Set(uint8(FunctionLedsOff), buf.Bytes())
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

// Returns *true* if the orientation and direction LEDs of the IMU Brick
// are on, *false* otherwise.
func (device *IMUV2Brick) AreLedsOn() (leds bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionAreLedsOn), buf.Bytes())
	if err != nil {
		return leds, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return leds, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return leds, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &leds)

	}

	return leds, nil
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
func (device *IMUV2Brick) SaveCalibration() (calibrationDone bool, err error) {
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

// Sets the period with which the RegisterAccelerationCallback callback is triggered
// periodically. A value of 0 turns the callback off.
func (device *IMUV2Brick) SetAccelerationPeriod(period uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);

	resultBytes, err := device.device.Set(uint8(FunctionSetAccelerationPeriod), buf.Bytes())
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

// Returns the period as set by SetAccelerationPeriod.
func (device *IMUV2Brick) GetAccelerationPeriod() (period uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetAccelerationPeriod), buf.Bytes())
	if err != nil {
		return period, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 12 {
			return period, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		if header.ErrorCode != 0 {
			return period, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &period)

	}

	return period, nil
}

// Sets the period with which the RegisterMagneticFieldCallback callback is triggered
// periodically. A value of 0 turns the callback off.
func (device *IMUV2Brick) SetMagneticFieldPeriod(period uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);

	resultBytes, err := device.device.Set(uint8(FunctionSetMagneticFieldPeriod), buf.Bytes())
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

// Returns the period as set by SetMagneticFieldPeriod.
func (device *IMUV2Brick) GetMagneticFieldPeriod() (period uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetMagneticFieldPeriod), buf.Bytes())
	if err != nil {
		return period, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 12 {
			return period, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

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
func (device *IMUV2Brick) SetAngularVelocityPeriod(period uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);

	resultBytes, err := device.device.Set(uint8(FunctionSetAngularVelocityPeriod), buf.Bytes())
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

// Returns the period as set by SetAngularVelocityPeriod.
func (device *IMUV2Brick) GetAngularVelocityPeriod() (period uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetAngularVelocityPeriod), buf.Bytes())
	if err != nil {
		return period, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 12 {
			return period, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		if header.ErrorCode != 0 {
			return period, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &period)

	}

	return period, nil
}

// Sets the period with which the RegisterTemperatureCallback callback is triggered
// periodically. A value of 0 turns the callback off.
func (device *IMUV2Brick) SetTemperaturePeriod(period uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);

	resultBytes, err := device.device.Set(uint8(FunctionSetTemperaturePeriod), buf.Bytes())
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

// Returns the period as set by SetTemperaturePeriod.
func (device *IMUV2Brick) GetTemperaturePeriod() (period uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetTemperaturePeriod), buf.Bytes())
	if err != nil {
		return period, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 12 {
			return period, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

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
func (device *IMUV2Brick) SetOrientationPeriod(period uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);

	resultBytes, err := device.device.Set(uint8(FunctionSetOrientationPeriod), buf.Bytes())
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

// Returns the period as set by SetOrientationPeriod.
func (device *IMUV2Brick) GetOrientationPeriod() (period uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetOrientationPeriod), buf.Bytes())
	if err != nil {
		return period, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 12 {
			return period, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		if header.ErrorCode != 0 {
			return period, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &period)

	}

	return period, nil
}

// Sets the period with which the RegisterLinearAccelerationCallback callback is
// triggered periodically. A value of 0 turns the callback off.
func (device *IMUV2Brick) SetLinearAccelerationPeriod(period uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);

	resultBytes, err := device.device.Set(uint8(FunctionSetLinearAccelerationPeriod), buf.Bytes())
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

// Returns the period as set by SetLinearAccelerationPeriod.
func (device *IMUV2Brick) GetLinearAccelerationPeriod() (period uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetLinearAccelerationPeriod), buf.Bytes())
	if err != nil {
		return period, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 12 {
			return period, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		if header.ErrorCode != 0 {
			return period, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &period)

	}

	return period, nil
}

// Sets the period with which the RegisterGravityVectorCallback callback is triggered
// periodically. A value of 0 turns the callback off.
func (device *IMUV2Brick) SetGravityVectorPeriod(period uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);

	resultBytes, err := device.device.Set(uint8(FunctionSetGravityVectorPeriod), buf.Bytes())
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

// Returns the period as set by SetGravityVectorPeriod.
func (device *IMUV2Brick) GetGravityVectorPeriod() (period uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetGravityVectorPeriod), buf.Bytes())
	if err != nil {
		return period, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 12 {
			return period, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

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
func (device *IMUV2Brick) SetQuaternionPeriod(period uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);

	resultBytes, err := device.device.Set(uint8(FunctionSetQuaternionPeriod), buf.Bytes())
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

// Returns the period as set by SetQuaternionPeriod.
func (device *IMUV2Brick) GetQuaternionPeriod() (period uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetQuaternionPeriod), buf.Bytes())
	if err != nil {
		return period, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 12 {
			return period, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

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
func (device *IMUV2Brick) SetAllDataPeriod(period uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);

	resultBytes, err := device.device.Set(uint8(FunctionSetAllDataPeriod), buf.Bytes())
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

// Returns the period as set by SetAllDataPeriod.
func (device *IMUV2Brick) GetAllDataPeriod() (period uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetAllDataPeriod), buf.Bytes())
	if err != nil {
		return period, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 12 {
			return period, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		if header.ErrorCode != 0 {
			return period, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &period)

	}

	return period, nil
}

// Sets the available sensor configuration for the Magnetometer, Gyroscope and
// Accelerometer. The Accelerometer Range is user selectable in all fusion modes,
// all other configurations are auto-controlled in fusion mode.
// 
// .. versionadded:: 2.0.5$nbsp;(Firmware)
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
func (device *IMUV2Brick) SetSensorConfiguration(magnetometerRate MagnetometerRate, gyroscopeRange GyroscopeRange, gyroscopeBandwidth GyroscopeBandwidth, accelerometerRange AccelerometerRange, accelerometerBandwidth AccelerometerBandwidth) (err error) {
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
// .. versionadded:: 2.0.5$nbsp;(Firmware)
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
func (device *IMUV2Brick) GetSensorConfiguration() (magnetometerRate MagnetometerRate, gyroscopeRange GyroscopeRange, gyroscopeBandwidth GyroscopeBandwidth, accelerometerRange AccelerometerRange, accelerometerBandwidth AccelerometerBandwidth, err error) {
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
// .. versionadded:: 2.0.5$nbsp;(Firmware)
//
// Associated constants:
//
//	* SensorFusionOff
//	* SensorFusionOn
//	* SensorFusionOnWithoutMagnetometer
//	* SensorFusionOnWithoutFastMagnetometerCalibration
func (device *IMUV2Brick) SetSensorFusionMode(mode SensorFusion) (err error) {
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
// .. versionadded:: 2.0.5$nbsp;(Firmware)
//
// Associated constants:
//
//	* SensorFusionOff
//	* SensorFusionOn
//	* SensorFusionOnWithoutMagnetometer
//	* SensorFusionOnWithoutFastMagnetometerCalibration
func (device *IMUV2Brick) GetSensorFusionMode() (mode SensorFusion, err error) {
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

// The SPITF protocol can be used with a dynamic baudrate. If the dynamic baudrate is
// enabled, the Brick will try to adapt the baudrate for the communication
// between Bricks and Bricklets according to the amount of data that is transferred.
// 
// The baudrate will be increased exponentially if lots of data is sent/received and
// decreased linearly if little data is sent/received.
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
// .. versionadded:: 2.0.10$nbsp;(Firmware)
func (device *IMUV2Brick) SetSPITFPBaudrateConfig(enableDynamicBaudrate bool, minimumDynamicBaudrate uint32) (err error) {
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

// Returns the baudrate config, see SetSPITFPBaudrateConfig.
// 
// .. versionadded:: 2.0.10$nbsp;(Firmware)
func (device *IMUV2Brick) GetSPITFPBaudrateConfig() (enableDynamicBaudrate bool, minimumDynamicBaudrate uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetSPITFPBaudrateConfig), buf.Bytes())
	if err != nil {
		return enableDynamicBaudrate, minimumDynamicBaudrate, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 13 {
			return enableDynamicBaudrate, minimumDynamicBaudrate, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 13)
		}

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
// .. versionadded:: 2.0.7$nbsp;(Firmware)
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
func (device *IMUV2Brick) GetSendTimeoutCount(communicationMethod CommunicationMethod) (timeoutCount uint32, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, communicationMethod);

	resultBytes, err := device.device.Get(uint8(FunctionGetSendTimeoutCount), buf.Bytes())
	if err != nil {
		return timeoutCount, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 12 {
			return timeoutCount, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		if header.ErrorCode != 0 {
			return timeoutCount, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &timeoutCount)

	}

	return timeoutCount, nil
}

// Sets the baudrate for a specific Bricklet port.
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
// .. versionadded:: 2.0.5$nbsp;(Firmware)
func (device *IMUV2Brick) SetSPITFPBaudrate(brickletPort rune, baudrate uint32) (err error) {
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

// Returns the baudrate for a given Bricklet port, see SetSPITFPBaudrate.
// 
// .. versionadded:: 2.0.5$nbsp;(Firmware)
func (device *IMUV2Brick) GetSPITFPBaudrate(brickletPort rune) (baudrate uint32, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, brickletPort);

	resultBytes, err := device.device.Get(uint8(FunctionGetSPITFPBaudrate), buf.Bytes())
	if err != nil {
		return baudrate, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 12 {
			return baudrate, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

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
// .. versionadded:: 2.0.5$nbsp;(Firmware)
func (device *IMUV2Brick) GetSPITFPErrorCount(brickletPort rune) (errorCountACKChecksum uint32, errorCountMessageChecksum uint32, errorCountFrame uint32, errorCountOverflow uint32, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, brickletPort);

	resultBytes, err := device.device.Get(uint8(FunctionGetSPITFPErrorCount), buf.Bytes())
	if err != nil {
		return errorCountACKChecksum, errorCountMessageChecksum, errorCountFrame, errorCountOverflow, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 24 {
			return errorCountACKChecksum, errorCountMessageChecksum, errorCountFrame, errorCountOverflow, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 24)
		}

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
func (device *IMUV2Brick) EnableStatusLED() (err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Set(uint8(FunctionEnableStatusLED), buf.Bytes())
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

// Disables the status LED.
// 
// The status LED is the blue LED next to the USB connector. If enabled is is
// on and it flickers if data is transfered. If disabled it is always off.
// 
// The default state is enabled.
func (device *IMUV2Brick) DisableStatusLED() (err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Set(uint8(FunctionDisableStatusLED), buf.Bytes())
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

// Returns *true* if the status LED is enabled, *false* otherwise.
func (device *IMUV2Brick) IsStatusLEDEnabled() (enabled bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionIsStatusLEDEnabled), buf.Bytes())
	if err != nil {
		return enabled, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return enabled, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

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
func (device *IMUV2Brick) GetProtocol1BrickletName(port rune) (protocolVersion uint8, firmwareVersion [3]uint8, name string, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, port);

	resultBytes, err := device.device.Get(uint8(FunctionGetProtocol1BrickletName), buf.Bytes())
	if err != nil {
		return protocolVersion, firmwareVersion, name, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 52 {
			return protocolVersion, firmwareVersion, name, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 52)
		}

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

// Returns the temperature as measured inside the microcontroller. The
// value returned is not the ambient temperature!
// 
// The temperature is only proportional to the real temperature and it has an
// accuracy of ±15%. Practically it is only useful as an indicator for
// temperature changes.
func (device *IMUV2Brick) GetChipTemperature() (temperature int16, err error) {
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

// Calling this function will reset the Brick. Calling this function
// on a Brick inside of a stack will reset the whole stack.
// 
// After a reset you have to create new device objects,
// calling functions on the existing ones will result in
// undefined behavior!
func (device *IMUV2Brick) Reset() (err error) {
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

// Writes 32 bytes of firmware to the bricklet attached at the given port.
// The bytes are written to the position offset * 32.
// 
// This function is used by Brick Viewer during flashing. It should not be
// necessary to call it in a normal user program.
func (device *IMUV2Brick) WriteBrickletPlugin(port rune, offset uint8, chunk [32]uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, port);
	binary.Write(&buf, binary.LittleEndian, offset);
	binary.Write(&buf, binary.LittleEndian, chunk);

	resultBytes, err := device.device.Set(uint8(FunctionWriteBrickletPlugin), buf.Bytes())
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

// Reads 32 bytes of firmware from the bricklet attached at the given port.
// The bytes are read starting at the position offset * 32.
// 
// This function is used by Brick Viewer during flashing. It should not be
// necessary to call it in a normal user program.
func (device *IMUV2Brick) ReadBrickletPlugin(port rune, offset uint8) (chunk [32]uint8, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, port);
	binary.Write(&buf, binary.LittleEndian, offset);

	resultBytes, err := device.device.Get(uint8(FunctionReadBrickletPlugin), buf.Bytes())
	if err != nil {
		return chunk, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 40 {
			return chunk, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 40)
		}

		if header.ErrorCode != 0 {
			return chunk, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &chunk)

	}

	return chunk, nil
}

// Returns the UID, the UID where the Brick is connected to,
// the position, the hardware and firmware version as well as the
// device identifier.
// 
// The position is the position in the stack from '0' (bottom) to '8' (top).
// 
// The device identifier numbers can be found `here <device_identifier>`.
// |device_identifier_constant|
func (device *IMUV2Brick) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
