/* ***********************************************************
 * This file was automatically generated on 2024-02-27.      *
 *                                                           *
 * Go Bindings Version 2.0.15                                *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/

// Measures power, DC voltage and DC current up to 720W/36V/20A‍.
//
//
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/VoltageCurrentV2_Bricklet_Go.html.
package voltage_current_v2_bricklet

import (
	"bytes"
	"encoding/binary"
	"fmt"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionGetCurrent                      Function = 1
	FunctionSetCurrentCallbackConfiguration Function = 2
	FunctionGetCurrentCallbackConfiguration Function = 3
	FunctionGetVoltage                      Function = 5
	FunctionSetVoltageCallbackConfiguration Function = 6
	FunctionGetVoltageCallbackConfiguration Function = 7
	FunctionGetPower                        Function = 9
	FunctionSetPowerCallbackConfiguration   Function = 10
	FunctionGetPowerCallbackConfiguration   Function = 11
	FunctionSetConfiguration                Function = 13
	FunctionGetConfiguration                Function = 14
	FunctionSetCalibration                  Function = 15
	FunctionGetCalibration                  Function = 16
	FunctionGetSPITFPErrorCount             Function = 234
	FunctionSetBootloaderMode               Function = 235
	FunctionGetBootloaderMode               Function = 236
	FunctionSetWriteFirmwarePointer         Function = 237
	FunctionWriteFirmware                   Function = 238
	FunctionSetStatusLEDConfig              Function = 239
	FunctionGetStatusLEDConfig              Function = 240
	FunctionGetChipTemperature              Function = 242
	FunctionReset                           Function = 243
	FunctionWriteUID                        Function = 248
	FunctionReadUID                         Function = 249
	FunctionGetIdentity                     Function = 255
	FunctionCallbackCurrent                 Function = 4
	FunctionCallbackVoltage                 Function = 8
	FunctionCallbackPower                   Function = 12
)

type ThresholdOption = rune

const (
	ThresholdOptionOff     ThresholdOption = 'x'
	ThresholdOptionOutside ThresholdOption = 'o'
	ThresholdOptionInside  ThresholdOption = 'i'
	ThresholdOptionSmaller ThresholdOption = '<'
	ThresholdOptionGreater ThresholdOption = '>'
)

type Averaging = uint8

const (
	Averaging1    Averaging = 0
	Averaging4    Averaging = 1
	Averaging16   Averaging = 2
	Averaging64   Averaging = 3
	Averaging128  Averaging = 4
	Averaging256  Averaging = 5
	Averaging512  Averaging = 6
	Averaging1024 Averaging = 7
)

type ConversionTime = uint8

const (
	ConversionTime140us ConversionTime = 0
	ConversionTime204us ConversionTime = 1
	ConversionTime332us ConversionTime = 2
	ConversionTime588us ConversionTime = 3
	//Deprecated: Use 1_1ms instead.
	ConversionTime11ms  ConversionTime = 4
	ConversionTime1_1ms ConversionTime = 4
	//Deprecated: Use 2_116ms instead.
	ConversionTime2116ms  ConversionTime = 5
	ConversionTime2_116ms ConversionTime = 5
	//Deprecated: Use 4_156ms instead.
	ConversionTime4156ms  ConversionTime = 6
	ConversionTime4_156ms ConversionTime = 6
	//Deprecated: Use 8_244ms instead.
	ConversionTime8244ms  ConversionTime = 7
	ConversionTime8_244ms ConversionTime = 7
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

type VoltageCurrentV2Bricklet struct {
	device Device
}

const DeviceIdentifier = 2105
const DeviceDisplayName = "Voltage/Current Bricklet 2.0"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (VoltageCurrentV2Bricklet, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{2, 0, 0}, uid, &internalIPCon, 0, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return VoltageCurrentV2Bricklet{}, err
	}
	dev.ResponseExpected[FunctionGetCurrent] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetCurrentCallbackConfiguration] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionGetCurrentCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetVoltage] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetVoltageCallbackConfiguration] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionGetVoltageCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetPower] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetPowerCallbackConfiguration] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionGetPowerCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetConfiguration] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetConfiguration] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetCalibration] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetCalibration] = ResponseExpectedFlagAlwaysTrue
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
	return VoltageCurrentV2Bricklet{dev}, nil
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
func (device *VoltageCurrentV2Bricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *VoltageCurrentV2Bricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *VoltageCurrentV2Bricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *VoltageCurrentV2Bricklet) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// This callback is triggered periodically according to the configuration set by
// SetCurrentCallbackConfiguration.
//
// The parameter is the same as GetCurrent.
func (device *VoltageCurrentV2Bricklet) RegisterCurrentCallback(fn func(int32)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 12 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var current int32
		binary.Read(buf, binary.LittleEndian, &current)
		fn(current)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackCurrent), wrapper)
}

// Remove a registered Current callback.
func (device *VoltageCurrentV2Bricklet) DeregisterCurrentCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackCurrent), registrationId)
}

// This callback is triggered periodically according to the configuration set by
// SetVoltageCallbackConfiguration.
//
// The parameter is the same as GetVoltage.
func (device *VoltageCurrentV2Bricklet) RegisterVoltageCallback(fn func(int32)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 12 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var voltage int32
		binary.Read(buf, binary.LittleEndian, &voltage)
		fn(voltage)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackVoltage), wrapper)
}

// Remove a registered Voltage callback.
func (device *VoltageCurrentV2Bricklet) DeregisterVoltageCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackVoltage), registrationId)
}

// This callback is triggered periodically according to the configuration set by
// SetPowerCallbackConfiguration.
//
// The parameter is the same as GetPower.
func (device *VoltageCurrentV2Bricklet) RegisterPowerCallback(fn func(int32)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 12 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var power int32
		binary.Read(buf, binary.LittleEndian, &power)
		fn(power)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackPower), wrapper)
}

// Remove a registered Power callback.
func (device *VoltageCurrentV2Bricklet) DeregisterPowerCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackPower), registrationId)
}

// Returns the current.
//
//
// If you want to get the value periodically, it is recommended to use the
// RegisterCurrentCallback callback. You can set the callback configuration
// with SetCurrentCallbackConfiguration.
func (device *VoltageCurrentV2Bricklet) GetCurrent() (current int32, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetCurrent), buf.Bytes())
	if err != nil {
		return current, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return current, DeviceError(header.ErrorCode)
		}

		if header.Length != 12 {
			return current, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &current)

	}

	return current, nil
}

// The period is the period with which the RegisterCurrentCallback callback is triggered
// periodically. A value of 0 turns the callback off.
//
// If the `value has to change`-parameter is set to true, the callback is only
// triggered after the value has changed. If the value didn't change
// within the period, the callback is triggered immediately on change.
//
// If it is set to false, the callback is continuously triggered with the period,
// independent of the value.
//
// It is furthermore possible to constrain the callback with thresholds.
//
// The `option`-parameter together with min/max sets a threshold for the RegisterCurrentCallback callback.
//
// The following options are possible:
//
//  Option| Description
//  --- | ---
//  'x'|    Threshold is turned off
//  'o'|    Threshold is triggered when the value is *outside* the min and max values
//  'i'|    Threshold is triggered when the value is *inside* or equal to the min and max values
//  '<'|    Threshold is triggered when the value is smaller than the min value (max is ignored)
//  '>'|    Threshold is triggered when the value is greater than the min value (max is ignored)
//
// If the option is set to 'x' (threshold turned off) the callback is triggered with the fixed period.
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *VoltageCurrentV2Bricklet) SetCurrentCallbackConfiguration(period uint32, valueHasToChange bool, option ThresholdOption, min int32, max int32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period)
	binary.Write(&buf, binary.LittleEndian, valueHasToChange)
	binary.Write(&buf, binary.LittleEndian, option)
	binary.Write(&buf, binary.LittleEndian, min)
	binary.Write(&buf, binary.LittleEndian, max)

	resultBytes, err := device.device.Set(uint8(FunctionSetCurrentCallbackConfiguration), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		bytes.NewBuffer(resultBytes[8:])

	}

	return nil
}

// Returns the callback configuration as set by SetCurrentCallbackConfiguration.
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *VoltageCurrentV2Bricklet) GetCurrentCallbackConfiguration() (period uint32, valueHasToChange bool, option ThresholdOption, min int32, max int32, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetCurrentCallbackConfiguration), buf.Bytes())
	if err != nil {
		return period, valueHasToChange, option, min, max, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return period, valueHasToChange, option, min, max, DeviceError(header.ErrorCode)
		}

		if header.Length != 22 {
			return period, valueHasToChange, option, min, max, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 22)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &period)
		binary.Read(resultBuf, binary.LittleEndian, &valueHasToChange)
		binary.Read(resultBuf, binary.LittleEndian, &option)
		binary.Read(resultBuf, binary.LittleEndian, &min)
		binary.Read(resultBuf, binary.LittleEndian, &max)

	}

	return period, valueHasToChange, option, min, max, nil
}

// Returns the voltage.
//
//
// If you want to get the value periodically, it is recommended to use the
// RegisterVoltageCallback callback. You can set the callback configuration
// with SetVoltageCallbackConfiguration.
func (device *VoltageCurrentV2Bricklet) GetVoltage() (voltage int32, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetVoltage), buf.Bytes())
	if err != nil {
		return voltage, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return voltage, DeviceError(header.ErrorCode)
		}

		if header.Length != 12 {
			return voltage, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &voltage)

	}

	return voltage, nil
}

// The period is the period with which the RegisterVoltageCallback callback is triggered
// periodically. A value of 0 turns the callback off.
//
// If the `value has to change`-parameter is set to true, the callback is only
// triggered after the value has changed. If the value didn't change
// within the period, the callback is triggered immediately on change.
//
// If it is set to false, the callback is continuously triggered with the period,
// independent of the value.
//
// It is furthermore possible to constrain the callback with thresholds.
//
// The `option`-parameter together with min/max sets a threshold for the RegisterVoltageCallback callback.
//
// The following options are possible:
//
//  Option| Description
//  --- | ---
//  'x'|    Threshold is turned off
//  'o'|    Threshold is triggered when the value is *outside* the min and max values
//  'i'|    Threshold is triggered when the value is *inside* or equal to the min and max values
//  '<'|    Threshold is triggered when the value is smaller than the min value (max is ignored)
//  '>'|    Threshold is triggered when the value is greater than the min value (max is ignored)
//
// If the option is set to 'x' (threshold turned off) the callback is triggered with the fixed period.
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *VoltageCurrentV2Bricklet) SetVoltageCallbackConfiguration(period uint32, valueHasToChange bool, option ThresholdOption, min int32, max int32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period)
	binary.Write(&buf, binary.LittleEndian, valueHasToChange)
	binary.Write(&buf, binary.LittleEndian, option)
	binary.Write(&buf, binary.LittleEndian, min)
	binary.Write(&buf, binary.LittleEndian, max)

	resultBytes, err := device.device.Set(uint8(FunctionSetVoltageCallbackConfiguration), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		bytes.NewBuffer(resultBytes[8:])

	}

	return nil
}

// Returns the callback configuration as set by SetVoltageCallbackConfiguration.
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *VoltageCurrentV2Bricklet) GetVoltageCallbackConfiguration() (period uint32, valueHasToChange bool, option ThresholdOption, min int32, max int32, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetVoltageCallbackConfiguration), buf.Bytes())
	if err != nil {
		return period, valueHasToChange, option, min, max, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return period, valueHasToChange, option, min, max, DeviceError(header.ErrorCode)
		}

		if header.Length != 22 {
			return period, valueHasToChange, option, min, max, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 22)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &period)
		binary.Read(resultBuf, binary.LittleEndian, &valueHasToChange)
		binary.Read(resultBuf, binary.LittleEndian, &option)
		binary.Read(resultBuf, binary.LittleEndian, &min)
		binary.Read(resultBuf, binary.LittleEndian, &max)

	}

	return period, valueHasToChange, option, min, max, nil
}

// Returns the power.
//
//
// If you want to get the value periodically, it is recommended to use the
// RegisterPowerCallback callback. You can set the callback configuration
// with SetPowerCallbackConfiguration.
func (device *VoltageCurrentV2Bricklet) GetPower() (power int32, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetPower), buf.Bytes())
	if err != nil {
		return power, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return power, DeviceError(header.ErrorCode)
		}

		if header.Length != 12 {
			return power, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &power)

	}

	return power, nil
}

// The period is the period with which the RegisterPowerCallback callback is triggered
// periodically. A value of 0 turns the callback off.
//
// If the `value has to change`-parameter is set to true, the callback is only
// triggered after the value has changed. If the value didn't change
// within the period, the callback is triggered immediately on change.
//
// If it is set to false, the callback is continuously triggered with the period,
// independent of the value.
//
// It is furthermore possible to constrain the callback with thresholds.
//
// The `option`-parameter together with min/max sets a threshold for the RegisterPowerCallback callback.
//
// The following options are possible:
//
//  Option| Description
//  --- | ---
//  'x'|    Threshold is turned off
//  'o'|    Threshold is triggered when the value is *outside* the min and max values
//  'i'|    Threshold is triggered when the value is *inside* or equal to the min and max values
//  '<'|    Threshold is triggered when the value is smaller than the min value (max is ignored)
//  '>'|    Threshold is triggered when the value is greater than the min value (max is ignored)
//
// If the option is set to 'x' (threshold turned off) the callback is triggered with the fixed period.
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *VoltageCurrentV2Bricklet) SetPowerCallbackConfiguration(period uint32, valueHasToChange bool, option ThresholdOption, min int32, max int32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period)
	binary.Write(&buf, binary.LittleEndian, valueHasToChange)
	binary.Write(&buf, binary.LittleEndian, option)
	binary.Write(&buf, binary.LittleEndian, min)
	binary.Write(&buf, binary.LittleEndian, max)

	resultBytes, err := device.device.Set(uint8(FunctionSetPowerCallbackConfiguration), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		bytes.NewBuffer(resultBytes[8:])

	}

	return nil
}

// Returns the callback configuration as set by SetPowerCallbackConfiguration.
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *VoltageCurrentV2Bricklet) GetPowerCallbackConfiguration() (period uint32, valueHasToChange bool, option ThresholdOption, min int32, max int32, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetPowerCallbackConfiguration), buf.Bytes())
	if err != nil {
		return period, valueHasToChange, option, min, max, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return period, valueHasToChange, option, min, max, DeviceError(header.ErrorCode)
		}

		if header.Length != 22 {
			return period, valueHasToChange, option, min, max, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 22)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &period)
		binary.Read(resultBuf, binary.LittleEndian, &valueHasToChange)
		binary.Read(resultBuf, binary.LittleEndian, &option)
		binary.Read(resultBuf, binary.LittleEndian, &min)
		binary.Read(resultBuf, binary.LittleEndian, &max)

	}

	return period, valueHasToChange, option, min, max, nil
}

// Sets the configuration of the Voltage/Current Bricklet 2.0. It is
// possible to configure number of averages as well as
// voltage and current conversion time.
//
// Associated constants:
//
//	* Averaging1
//	* Averaging4
//	* Averaging16
//	* Averaging64
//	* Averaging128
//	* Averaging256
//	* Averaging512
//	* Averaging1024
//	* ConversionTime140us
//	* ConversionTime204us
//	* ConversionTime332us
//	* ConversionTime588us
//	* ConversionTime11ms
//	* ConversionTime2116ms
//	* ConversionTime4156ms
//	* ConversionTime8244ms
func (device *VoltageCurrentV2Bricklet) SetConfiguration(averaging Averaging, voltageConversionTime ConversionTime, currentConversionTime ConversionTime) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, averaging)
	binary.Write(&buf, binary.LittleEndian, voltageConversionTime)
	binary.Write(&buf, binary.LittleEndian, currentConversionTime)

	resultBytes, err := device.device.Set(uint8(FunctionSetConfiguration), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		bytes.NewBuffer(resultBytes[8:])

	}

	return nil
}

// Returns the configuration as set by SetConfiguration.
//
// Associated constants:
//
//	* Averaging1
//	* Averaging4
//	* Averaging16
//	* Averaging64
//	* Averaging128
//	* Averaging256
//	* Averaging512
//	* Averaging1024
//	* ConversionTime140us
//	* ConversionTime204us
//	* ConversionTime332us
//	* ConversionTime588us
//	* ConversionTime11ms
//	* ConversionTime2116ms
//	* ConversionTime4156ms
//	* ConversionTime8244ms
func (device *VoltageCurrentV2Bricklet) GetConfiguration() (averaging Averaging, voltageConversionTime ConversionTime, currentConversionTime ConversionTime, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetConfiguration), buf.Bytes())
	if err != nil {
		return averaging, voltageConversionTime, currentConversionTime, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return averaging, voltageConversionTime, currentConversionTime, DeviceError(header.ErrorCode)
		}

		if header.Length != 11 {
			return averaging, voltageConversionTime, currentConversionTime, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 11)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &averaging)
		binary.Read(resultBuf, binary.LittleEndian, &voltageConversionTime)
		binary.Read(resultBuf, binary.LittleEndian, &currentConversionTime)

	}

	return averaging, voltageConversionTime, currentConversionTime, nil
}

// Since the ADC and the shunt resistor used for the measurements
// are not perfect they need to be calibrated by a multiplier and
// a divisor if a very precise reading is needed.
//
// For example, if you are expecting a current of 1000mA and you
// are measuring 1023mA, you can calibrate the Voltage/Current Bricklet
// by setting the current multiplier to 1000 and the divisor to 1023.
// The same applies for the voltage.
//
// The calibration will be saved on the EEPROM of the Voltage/Current
// Bricklet and only needs to be done once.
func (device *VoltageCurrentV2Bricklet) SetCalibration(voltageMultiplier uint16, voltageDivisor uint16, currentMultiplier uint16, currentDivisor uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, voltageMultiplier)
	binary.Write(&buf, binary.LittleEndian, voltageDivisor)
	binary.Write(&buf, binary.LittleEndian, currentMultiplier)
	binary.Write(&buf, binary.LittleEndian, currentDivisor)

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

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		bytes.NewBuffer(resultBytes[8:])

	}

	return nil
}

// Returns the calibration as set by SetCalibration.
func (device *VoltageCurrentV2Bricklet) GetCalibration() (voltageMultiplier uint16, voltageDivisor uint16, currentMultiplier uint16, currentDivisor uint16, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetCalibration), buf.Bytes())
	if err != nil {
		return voltageMultiplier, voltageDivisor, currentMultiplier, currentDivisor, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return voltageMultiplier, voltageDivisor, currentMultiplier, currentDivisor, DeviceError(header.ErrorCode)
		}

		if header.Length != 16 {
			return voltageMultiplier, voltageDivisor, currentMultiplier, currentDivisor, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 16)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &voltageMultiplier)
		binary.Read(resultBuf, binary.LittleEndian, &voltageDivisor)
		binary.Read(resultBuf, binary.LittleEndian, &currentMultiplier)
		binary.Read(resultBuf, binary.LittleEndian, &currentDivisor)

	}

	return voltageMultiplier, voltageDivisor, currentMultiplier, currentDivisor, nil
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
func (device *VoltageCurrentV2Bricklet) GetSPITFPErrorCount() (errorCountAckChecksum uint32, errorCountMessageChecksum uint32, errorCountFrame uint32, errorCountOverflow uint32, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetSPITFPErrorCount), buf.Bytes())
	if err != nil {
		return errorCountAckChecksum, errorCountMessageChecksum, errorCountFrame, errorCountOverflow, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return errorCountAckChecksum, errorCountMessageChecksum, errorCountFrame, errorCountOverflow, DeviceError(header.ErrorCode)
		}

		if header.Length != 24 {
			return errorCountAckChecksum, errorCountMessageChecksum, errorCountFrame, errorCountOverflow, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 24)
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
func (device *VoltageCurrentV2Bricklet) SetBootloaderMode(mode BootloaderMode) (status BootloaderStatus, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, mode)

	resultBytes, err := device.device.Get(uint8(FunctionSetBootloaderMode), buf.Bytes())
	if err != nil {
		return status, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return status, DeviceError(header.ErrorCode)
		}

		if header.Length != 9 {
			return status, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
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
func (device *VoltageCurrentV2Bricklet) GetBootloaderMode() (mode BootloaderMode, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetBootloaderMode), buf.Bytes())
	if err != nil {
		return mode, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return mode, DeviceError(header.ErrorCode)
		}

		if header.Length != 9 {
			return mode, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
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
func (device *VoltageCurrentV2Bricklet) SetWriteFirmwarePointer(pointer uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, pointer)

	resultBytes, err := device.device.Set(uint8(FunctionSetWriteFirmwarePointer), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
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
func (device *VoltageCurrentV2Bricklet) WriteFirmware(data [64]uint8) (status uint8, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, data)

	resultBytes, err := device.device.Get(uint8(FunctionWriteFirmware), buf.Bytes())
	if err != nil {
		return status, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return status, DeviceError(header.ErrorCode)
		}

		if header.Length != 9 {
			return status, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
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
func (device *VoltageCurrentV2Bricklet) SetStatusLEDConfig(config StatusLEDConfig) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, config)

	resultBytes, err := device.device.Set(uint8(FunctionSetStatusLEDConfig), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
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
func (device *VoltageCurrentV2Bricklet) GetStatusLEDConfig() (config StatusLEDConfig, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetStatusLEDConfig), buf.Bytes())
	if err != nil {
		return config, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return config, DeviceError(header.ErrorCode)
		}

		if header.Length != 9 {
			return config, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
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
func (device *VoltageCurrentV2Bricklet) GetChipTemperature() (temperature int16, err error) {
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

		if header.Length != 10 {
			return temperature, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
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
func (device *VoltageCurrentV2Bricklet) Reset() (err error) {
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

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
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
func (device *VoltageCurrentV2Bricklet) WriteUID(uid uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, uid)

	resultBytes, err := device.device.Set(uint8(FunctionWriteUID), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		bytes.NewBuffer(resultBytes[8:])

	}

	return nil
}

// Returns the current UID as an integer. Encode as
// Base58 to get the usual string version.
func (device *VoltageCurrentV2Bricklet) ReadUID() (uid uint32, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionReadUID), buf.Bytes())
	if err != nil {
		return uid, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return uid, DeviceError(header.ErrorCode)
		}

		if header.Length != 12 {
			return uid, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
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
func (device *VoltageCurrentV2Bricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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

		if header.Length != 33 {
			return uid, connectedUid, position, hardwareVersion, firmwareVersion, deviceIdentifier, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 33)
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
