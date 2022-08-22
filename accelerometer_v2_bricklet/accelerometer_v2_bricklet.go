/* ***********************************************************
 * This file was automatically generated on 2022-08-22.      *
 *                                                           *
 * Go Bindings Version 2.0.14                                *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/

// Measures acceleration in three axis.
//
//
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/AccelerometerV2_Bricklet_Go.html.
package accelerometer_v2_bricklet

import (
	"bytes"
	"encoding/binary"
	"fmt"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionGetAcceleration                        Function = 1
	FunctionSetConfiguration                       Function = 2
	FunctionGetConfiguration                       Function = 3
	FunctionSetAccelerationCallbackConfiguration   Function = 4
	FunctionGetAccelerationCallbackConfiguration   Function = 5
	FunctionSetInfoLEDConfig                       Function = 6
	FunctionGetInfoLEDConfig                       Function = 7
	FunctionSetContinuousAccelerationConfiguration Function = 9
	FunctionGetContinuousAccelerationConfiguration Function = 10
	FunctionSetFilterConfiguration                 Function = 13
	FunctionGetFilterConfiguration                 Function = 14
	FunctionGetSPITFPErrorCount                    Function = 234
	FunctionSetBootloaderMode                      Function = 235
	FunctionGetBootloaderMode                      Function = 236
	FunctionSetWriteFirmwarePointer                Function = 237
	FunctionWriteFirmware                          Function = 238
	FunctionSetStatusLEDConfig                     Function = 239
	FunctionGetStatusLEDConfig                     Function = 240
	FunctionGetChipTemperature                     Function = 242
	FunctionReset                                  Function = 243
	FunctionWriteUID                               Function = 248
	FunctionReadUID                                Function = 249
	FunctionGetIdentity                            Function = 255
	FunctionCallbackAcceleration                   Function = 8
	FunctionCallbackContinuousAcceleration16Bit    Function = 11
	FunctionCallbackContinuousAcceleration8Bit     Function = 12
)

type DataRate = uint8

const (
	//Deprecated: Use 0_781Hz instead.
	DataRate0781Hz  DataRate = 0
	DataRate0_781Hz DataRate = 0
	//Deprecated: Use 1_563Hz instead.
	DataRate1563Hz  DataRate = 1
	DataRate1_563Hz DataRate = 1
	//Deprecated: Use 3_125Hz instead.
	DataRate3125Hz  DataRate = 2
	DataRate3_125Hz DataRate = 2
	//Deprecated: Use 6_2512Hz instead.
	DataRate62512Hz  DataRate = 3
	DataRate6_2512Hz DataRate = 3
	//Deprecated: Use 12_5Hz instead.
	DataRate125Hz   DataRate = 4
	DataRate12_5Hz  DataRate = 4
	DataRate25Hz    DataRate = 5
	DataRate50Hz    DataRate = 6
	DataRate100Hz   DataRate = 7
	DataRate200Hz   DataRate = 8
	DataRate400Hz   DataRate = 9
	DataRate800Hz   DataRate = 10
	DataRate1600Hz  DataRate = 11
	DataRate3200Hz  DataRate = 12
	DataRate6400Hz  DataRate = 13
	DataRate12800Hz DataRate = 14
	DataRate25600Hz DataRate = 15
)

type FullScale = uint8

const (
	FullScale2g FullScale = 0
	FullScale4g FullScale = 1
	FullScale8g FullScale = 2
)

type InfoLEDConfig = uint8

const (
	InfoLEDConfigOff           InfoLEDConfig = 0
	InfoLEDConfigOn            InfoLEDConfig = 1
	InfoLEDConfigShowHeartbeat InfoLEDConfig = 2
)

type Resolution = uint8

const (
	Resolution8bit  Resolution = 0
	Resolution16bit Resolution = 1
)

type IIRBypass = uint8

const (
	IIRBypassApplied  IIRBypass = 0
	IIRBypassBypassed IIRBypass = 1
)

type LowPassFilter = uint8

const (
	LowPassFilterNinth LowPassFilter = 0
	LowPassFilterHalf  LowPassFilter = 1
)

type BootloaderMode = uint8

const (
	BootloaderModeBootloader                    BootloaderMode = 0
	BootloaderModeFirmware                      BootloaderMode = 1
	BootloaderModeBootloaderWaitForReboot       BootloaderMode = 2
	BootloaderModeFirmwareWaitForReboot         BootloaderMode = 3
	BootloaderModeFirmwareWaitForEraseAndReboot BootloaderMode = 4
)

type BootloaderStatus = uint8

const (
	BootloaderStatusOK                        BootloaderStatus = 0
	BootloaderStatusInvalidMode               BootloaderStatus = 1
	BootloaderStatusNoChange                  BootloaderStatus = 2
	BootloaderStatusEntryFunctionNotPresent   BootloaderStatus = 3
	BootloaderStatusDeviceIdentifierIncorrect BootloaderStatus = 4
	BootloaderStatusCRCMismatch               BootloaderStatus = 5
)

type StatusLEDConfig = uint8

const (
	StatusLEDConfigOff           StatusLEDConfig = 0
	StatusLEDConfigOn            StatusLEDConfig = 1
	StatusLEDConfigShowHeartbeat StatusLEDConfig = 2
	StatusLEDConfigShowStatus    StatusLEDConfig = 3
)

type AccelerometerV2Bricklet struct {
	device Device
}

const DeviceIdentifier = 2130
const DeviceDisplayName = "Accelerometer Bricklet 2.0"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (AccelerometerV2Bricklet, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{2, 0, 1}, uid, &internalIPCon, 0, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return AccelerometerV2Bricklet{}, err
	}
	dev.ResponseExpected[FunctionGetAcceleration] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetConfiguration] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetConfiguration] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetAccelerationCallbackConfiguration] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionGetAccelerationCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetInfoLEDConfig] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetInfoLEDConfig] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetContinuousAccelerationConfiguration] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionGetContinuousAccelerationConfiguration] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetFilterConfiguration] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetFilterConfiguration] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetSPITFPErrorCount] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetBootloaderMode] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetBootloaderMode] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetWriteFirmwarePointer] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionWriteFirmware] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetStatusLEDConfig] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetStatusLEDConfig] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetChipTemperature] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionReset] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionWriteUID] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionReadUID] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetIdentity] = ResponseExpectedFlagAlwaysTrue
	return AccelerometerV2Bricklet{dev}, nil
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
func (device *AccelerometerV2Bricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *AccelerometerV2Bricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *AccelerometerV2Bricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *AccelerometerV2Bricklet) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// This callback is triggered periodically according to the configuration set by
// SetAccelerationCallbackConfiguration.
//
// The parameters are the same as GetAcceleration.
func (device *AccelerometerV2Bricklet) RegisterAccelerationCallback(fn func(int32, int32, int32)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 20 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var x int32
		var y int32
		var z int32
		binary.Read(buf, binary.LittleEndian, &x)
		binary.Read(buf, binary.LittleEndian, &y)
		binary.Read(buf, binary.LittleEndian, &z)
		fn(x, y, z)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackAcceleration), wrapper)
}

// Remove a registered Acceleration callback.
func (device *AccelerometerV2Bricklet) DeregisterAccelerationCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackAcceleration), registrationId)
}

// Returns 30 acceleration values with 16 bit resolution. The data rate can
// be configured with SetConfiguration and this callback can be
// enabled with SetContinuousAccelerationConfiguration.
//
// The returned values are raw ADC data. If you want to put this data into
// a FFT to determine the occurrences of specific frequencies we recommend
// that you use the data as is. It has all of the ADC noise in it. This noise
// looks like pure noise at first glance, but it might still have some frequnecy
// information in it that can be utilized by the FFT.
//
// Otherwise you have to use the following formulas that depend on the
// full scale range (see SetConfiguration) to calculate
// the data in gₙ/10000 (same unit that is returned by GetAcceleration):
//
// * Full scale 2g: acceleration = value * 625 / 1024
// * Full scale 4g: acceleration = value * 1250 / 1024
// * Full scale 8g: acceleration = value * 2500 / 1024
//
// The data is formated in the sequence x, y, z, x, y, z, ... depending on
// the enabled axis. Examples:
//
// * x, y, z enabled: x, y, z, ... 10x repeated
// * x, z enabled: x, z, ... 15x repeated
// * y enabled: y, ... 30x repeated
func (device *AccelerometerV2Bricklet) RegisterContinuousAcceleration16BitCallback(fn func([30]int16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 68 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var acceleration [30]int16
		binary.Read(buf, binary.LittleEndian, &acceleration)
		fn(acceleration)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackContinuousAcceleration16Bit), wrapper)
}

// Remove a registered Continuous Acceleration 16 Bit callback.
func (device *AccelerometerV2Bricklet) DeregisterContinuousAcceleration16BitCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackContinuousAcceleration16Bit), registrationId)
}

// Returns 60 acceleration values with 8 bit resolution. The data rate can
// be configured with SetConfiguration and this callback can be
// enabled with SetContinuousAccelerationConfiguration.
//
// The returned values are raw ADC data. If you want to put this data into
// a FFT to determine the occurrences of specific frequencies we recommend
// that you use the data as is. It has all of the ADC noise in it. This noise
// looks like pure noise at first glance, but it might still have some frequnecy
// information in it that can be utilized by the FFT.
//
// Otherwise you have to use the following formulas that depend on the
// full scale range (see SetConfiguration) to calculate
// the data in gₙ/10000 (same unit that is returned by GetAcceleration):
//
// * Full scale 2g: acceleration = value * 256 * 625 / 1024
// * Full scale 4g: acceleration = value * 256 * 1250 / 1024
// * Full scale 8g: acceleration = value * 256 * 2500 / 1024
//
// The data is formated in the sequence x, y, z, x, y, z, ... depending on
// the enabled axis. Examples:
//
// * x, y, z enabled: x, y, z, ... 20x repeated
// * x, z enabled: x, z, ... 30x repeated
// * y enabled: y, ... 60x repeated
func (device *AccelerometerV2Bricklet) RegisterContinuousAcceleration8BitCallback(fn func([60]int8)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 68 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var acceleration [60]int8
		binary.Read(buf, binary.LittleEndian, &acceleration)
		fn(acceleration)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackContinuousAcceleration8Bit), wrapper)
}

// Remove a registered Continuous Acceleration 8 Bit callback.
func (device *AccelerometerV2Bricklet) DeregisterContinuousAcceleration8BitCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackContinuousAcceleration8Bit), registrationId)
}

// Returns the acceleration in x, y and z direction. The values
// are given in gₙ/10000 (1gₙ = 9.80665m/s²). The range is
// configured with SetConfiguration.
//
// If you want to get the acceleration periodically, it is recommended
// to use the RegisterAccelerationCallback callback and set the period with
// SetAccelerationCallbackConfiguration.
func (device *AccelerometerV2Bricklet) GetAcceleration() (x int32, y int32, z int32, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetAcceleration), buf.Bytes())
	if err != nil {
		return x, y, z, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 20 {
			return x, y, z, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 20)
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

// Configures the data rate and full scale range.
// Possible values are:
//
// * Data rate of 0.781Hz to 25600Hz.
// * Full scale range of ±2g up to ±8g.
//
// Decreasing data rate or full scale range will also decrease the noise on
// the data.
//
// Associated constants:
//
//	* DataRate0781Hz
//	* DataRate1563Hz
//	* DataRate3125Hz
//	* DataRate62512Hz
//	* DataRate125Hz
//	* DataRate25Hz
//	* DataRate50Hz
//	* DataRate100Hz
//	* DataRate200Hz
//	* DataRate400Hz
//	* DataRate800Hz
//	* DataRate1600Hz
//	* DataRate3200Hz
//	* DataRate6400Hz
//	* DataRate12800Hz
//	* DataRate25600Hz
//	* FullScale2g
//	* FullScale4g
//	* FullScale8g
func (device *AccelerometerV2Bricklet) SetConfiguration(dataRate DataRate, fullScale FullScale) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, dataRate)
	binary.Write(&buf, binary.LittleEndian, fullScale)

	resultBytes, err := device.device.Set(uint8(FunctionSetConfiguration), buf.Bytes())
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

// Returns the configuration as set by SetConfiguration.
//
// Associated constants:
//
//	* DataRate0781Hz
//	* DataRate1563Hz
//	* DataRate3125Hz
//	* DataRate62512Hz
//	* DataRate125Hz
//	* DataRate25Hz
//	* DataRate50Hz
//	* DataRate100Hz
//	* DataRate200Hz
//	* DataRate400Hz
//	* DataRate800Hz
//	* DataRate1600Hz
//	* DataRate3200Hz
//	* DataRate6400Hz
//	* DataRate12800Hz
//	* DataRate25600Hz
//	* FullScale2g
//	* FullScale4g
//	* FullScale8g
func (device *AccelerometerV2Bricklet) GetConfiguration() (dataRate DataRate, fullScale FullScale, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetConfiguration), buf.Bytes())
	if err != nil {
		return dataRate, fullScale, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 10 {
			return dataRate, fullScale, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		if header.ErrorCode != 0 {
			return dataRate, fullScale, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &dataRate)
		binary.Read(resultBuf, binary.LittleEndian, &fullScale)

	}

	return dataRate, fullScale, nil
}

// The period is the period with which the RegisterAccelerationCallback
// callback is triggered periodically. A value of 0 turns the callback off.
//
// If the `value has to change`-parameter is set to true, the callback is only
// triggered after the value has changed. If the value didn't change within the
// period, the callback is triggered immediately on change.
//
// If it is set to false, the callback is continuously triggered with the period,
// independent of the value.
//
// If this callback is enabled, the RegisterContinuousAcceleration16BitCallback callback
// and RegisterContinuousAcceleration8BitCallback callback will automatically be disabled.
func (device *AccelerometerV2Bricklet) SetAccelerationCallbackConfiguration(period uint32, valueHasToChange bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period)
	binary.Write(&buf, binary.LittleEndian, valueHasToChange)

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

// Returns the callback configuration as set by
// SetAccelerationCallbackConfiguration.
func (device *AccelerometerV2Bricklet) GetAccelerationCallbackConfiguration() (period uint32, valueHasToChange bool, err error) {
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

// Configures the info LED (marked as Force on the Bricklet) to be either turned off,
// turned on, or blink in heartbeat mode.
//
// Associated constants:
//
//	* InfoLEDConfigOff
//	* InfoLEDConfigOn
//	* InfoLEDConfigShowHeartbeat
func (device *AccelerometerV2Bricklet) SetInfoLEDConfig(config InfoLEDConfig) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, config)

	resultBytes, err := device.device.Set(uint8(FunctionSetInfoLEDConfig), buf.Bytes())
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

// Returns the LED configuration as set by SetInfoLEDConfig
//
// Associated constants:
//
//	* InfoLEDConfigOff
//	* InfoLEDConfigOn
//	* InfoLEDConfigShowHeartbeat
func (device *AccelerometerV2Bricklet) GetInfoLEDConfig() (config InfoLEDConfig, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetInfoLEDConfig), buf.Bytes())
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

// For high throughput of acceleration data (> 1000Hz) you have to use the
// RegisterContinuousAcceleration16BitCallback or RegisterContinuousAcceleration8BitCallback
// callbacks.
//
// You can enable the callback for each axis (x, y, z) individually and choose a
// resolution of 8 bit or 16 bit.
//
// If at least one of the axis is enabled and the resolution is set to 8 bit,
// the RegisterContinuousAcceleration8BitCallback callback is activated. If at least
// one of the axis is enabled and the resolution is set to 16 bit,
// the RegisterContinuousAcceleration16BitCallback callback is activated.
//
// The returned values are raw ADC data. If you want to put this data into
// a FFT to determine the occurrences of specific frequencies we recommend
// that you use the data as is. It has all of the ADC noise in it. This noise
// looks like pure noise at first glance, but it might still have some frequnecy
// information in it that can be utilized by the FFT.
//
// Otherwise you have to use the following formulas that depend on the configured
// resolution (8/16 bit) and the full scale range (see SetConfiguration) to calculate
// the data in gₙ/10000 (same unit that is returned by GetAcceleration):
//
// * 16 bit, full scale 2g: acceleration = value * 625 / 1024
// * 16 bit, full scale 4g: acceleration = value * 1250 / 1024
// * 16 bit, full scale 8g: acceleration = value * 2500 / 1024
//
// If a resolution of 8 bit is used, only the 8 most significant bits will be
// transferred, so you can use the following formulas:
//
// * 8 bit, full scale 2g: acceleration = value * 256 * 625 / 1024
// * 8 bit, full scale 4g: acceleration = value * 256 * 1250 / 1024
// * 8 bit, full scale 8g: acceleration = value * 256 * 2500 / 1024
//
// If no axis is enabled, both callbacks are disabled. If one of the continuous
// callbacks is enabled, the RegisterAccelerationCallback callback is disabled.
//
// The maximum throughput depends on the exact configuration:
//
//  Number of axis enabled| Throughput 8 bit| Throughout 16 bit
//  --- | --- | ---
//  1| 25600Hz| 25600Hz
//  2| 25600Hz| 15000Hz
//  3| 20000Hz| 10000Hz
//
// Associated constants:
//
//	* Resolution8bit
//	* Resolution16bit
func (device *AccelerometerV2Bricklet) SetContinuousAccelerationConfiguration(enableX bool, enableY bool, enableZ bool, resolution Resolution) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, enableX)
	binary.Write(&buf, binary.LittleEndian, enableY)
	binary.Write(&buf, binary.LittleEndian, enableZ)
	binary.Write(&buf, binary.LittleEndian, resolution)

	resultBytes, err := device.device.Set(uint8(FunctionSetContinuousAccelerationConfiguration), buf.Bytes())
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

// Returns the continuous acceleration configuration as set by
// SetContinuousAccelerationConfiguration.
//
// Associated constants:
//
//	* Resolution8bit
//	* Resolution16bit
func (device *AccelerometerV2Bricklet) GetContinuousAccelerationConfiguration() (enableX bool, enableY bool, enableZ bool, resolution Resolution, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetContinuousAccelerationConfiguration), buf.Bytes())
	if err != nil {
		return enableX, enableY, enableZ, resolution, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 12 {
			return enableX, enableY, enableZ, resolution, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		if header.ErrorCode != 0 {
			return enableX, enableY, enableZ, resolution, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &enableX)
		binary.Read(resultBuf, binary.LittleEndian, &enableY)
		binary.Read(resultBuf, binary.LittleEndian, &enableZ)
		binary.Read(resultBuf, binary.LittleEndian, &resolution)

	}

	return enableX, enableY, enableZ, resolution, nil
}

// Configures IIR Bypass filter mode and low pass filter roll off corner frequency.
//
// The filter can be applied or bypassed and the corner frequency can be
// half or a ninth of the output data rate.
//
// .. image:: /Images/Bricklets/bricklet_accelerometer_v2_filter.png
//    :scale: 100 %
//    :alt: Accelerometer filter
//    :align: center
//    :target: ../../_images/Bricklets/bricklet_accelerometer_v2_filter.png
//
// .. versionadded:: 2.0.2$nbsp;(Plugin)
//
// Associated constants:
//
//	* IIRBypassApplied
//	* IIRBypassBypassed
//	* LowPassFilterNinth
//	* LowPassFilterHalf
func (device *AccelerometerV2Bricklet) SetFilterConfiguration(iirBypass IIRBypass, lowPassFilter LowPassFilter) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, iirBypass)
	binary.Write(&buf, binary.LittleEndian, lowPassFilter)

	resultBytes, err := device.device.Set(uint8(FunctionSetFilterConfiguration), buf.Bytes())
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

// Returns the configuration as set by SetFilterConfiguration.
//
// .. versionadded:: 2.0.2$nbsp;(Plugin)
//
// Associated constants:
//
//	* IIRBypassApplied
//	* IIRBypassBypassed
//	* LowPassFilterNinth
//	* LowPassFilterHalf
func (device *AccelerometerV2Bricklet) GetFilterConfiguration() (iirBypass IIRBypass, lowPassFilter LowPassFilter, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetFilterConfiguration), buf.Bytes())
	if err != nil {
		return iirBypass, lowPassFilter, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 10 {
			return iirBypass, lowPassFilter, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		if header.ErrorCode != 0 {
			return iirBypass, lowPassFilter, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &iirBypass)
		binary.Read(resultBuf, binary.LittleEndian, &lowPassFilter)

	}

	return iirBypass, lowPassFilter, nil
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
func (device *AccelerometerV2Bricklet) GetSPITFPErrorCount() (errorCountAckChecksum uint32, errorCountMessageChecksum uint32, errorCountFrame uint32, errorCountOverflow uint32, err error) {
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
func (device *AccelerometerV2Bricklet) SetBootloaderMode(mode BootloaderMode) (status BootloaderStatus, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, mode)

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
func (device *AccelerometerV2Bricklet) GetBootloaderMode() (mode BootloaderMode, err error) {
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
func (device *AccelerometerV2Bricklet) SetWriteFirmwarePointer(pointer uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, pointer)

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
func (device *AccelerometerV2Bricklet) WriteFirmware(data [64]uint8) (status uint8, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, data)

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
func (device *AccelerometerV2Bricklet) SetStatusLEDConfig(config StatusLEDConfig) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, config)

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
func (device *AccelerometerV2Bricklet) GetStatusLEDConfig() (config StatusLEDConfig, err error) {
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
func (device *AccelerometerV2Bricklet) GetChipTemperature() (temperature int16, err error) {
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
func (device *AccelerometerV2Bricklet) Reset() (err error) {
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
func (device *AccelerometerV2Bricklet) WriteUID(uid uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, uid)

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
func (device *AccelerometerV2Bricklet) ReadUID() (uid uint32, err error) {
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
func (device *AccelerometerV2Bricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
