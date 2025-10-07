/* ***********************************************************
 * This file was automatically generated on 2025-10-07.      *
 *                                                           *
 * Go Bindings Version 2.0.17                                *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/

// Measures temperature with thermocouples.
//
//
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/ThermocoupleV2_Bricklet_Go.html.
package thermocouple_v2_bricklet

import (
	"bytes"
	"encoding/binary"
	"fmt"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionGetTemperature                      Function = 1
	FunctionSetTemperatureCallbackConfiguration Function = 2
	FunctionGetTemperatureCallbackConfiguration Function = 3
	FunctionSetConfiguration                    Function = 5
	FunctionGetConfiguration                    Function = 6
	FunctionGetErrorState                       Function = 7
	FunctionGetSPITFPErrorCount                 Function = 234
	FunctionSetBootloaderMode                   Function = 235
	FunctionGetBootloaderMode                   Function = 236
	FunctionSetWriteFirmwarePointer             Function = 237
	FunctionWriteFirmware                       Function = 238
	FunctionSetStatusLEDConfig                  Function = 239
	FunctionGetStatusLEDConfig                  Function = 240
	FunctionGetChipTemperature                  Function = 242
	FunctionReset                               Function = 243
	FunctionWriteUID                            Function = 248
	FunctionReadUID                             Function = 249
	FunctionGetIdentity                         Function = 255
	FunctionCallbackTemperature                 Function = 4
	FunctionCallbackErrorState                  Function = 8
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
	Averaging1  Averaging = 1
	Averaging2  Averaging = 2
	Averaging4  Averaging = 4
	Averaging8  Averaging = 8
	Averaging16 Averaging = 16
)

type Type = uint8

const (
	TypeB   Type = 0
	TypeE   Type = 1
	TypeJ   Type = 2
	TypeK   Type = 3
	TypeN   Type = 4
	TypeR   Type = 5
	TypeS   Type = 6
	TypeT   Type = 7
	TypeG8  Type = 8
	TypeG32 Type = 9
)

type FilterOption = uint8

const (
	FilterOption50Hz FilterOption = 0
	FilterOption60Hz FilterOption = 1
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

type ThermocoupleV2Bricklet struct {
	device Device
}

const DeviceIdentifier = 2109
const DeviceDisplayName = "Thermocouple Bricklet 2.0"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (ThermocoupleV2Bricklet, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{2, 0, 0}, uid, &internalIPCon, 0, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return ThermocoupleV2Bricklet{}, err
	}
	dev.ResponseExpected[FunctionGetTemperature] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetTemperatureCallbackConfiguration] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionGetTemperatureCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetConfiguration] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetConfiguration] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetErrorState] = ResponseExpectedFlagAlwaysTrue
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
	return ThermocoupleV2Bricklet{dev}, nil
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
func (device *ThermocoupleV2Bricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *ThermocoupleV2Bricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *ThermocoupleV2Bricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *ThermocoupleV2Bricklet) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// This callback is triggered periodically according to the configuration set by
// SetTemperatureCallbackConfiguration.
//
// The parameter is the same as GetTemperature.
func (device *ThermocoupleV2Bricklet) RegisterTemperatureCallback(fn func(int32)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 12 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var temperature int32
		binary.Read(buf, binary.LittleEndian, &temperature)
		fn(temperature)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackTemperature), wrapper)
}

// Remove a registered Temperature callback.
func (device *ThermocoupleV2Bricklet) DeregisterTemperatureCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackTemperature), registrationId)
}

// This Callback is triggered every time the error state changes
// (see GetErrorState).
func (device *ThermocoupleV2Bricklet) RegisterErrorStateCallback(fn func(bool, bool)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 10 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var overUnder bool
		var openCircuit bool
		binary.Read(buf, binary.LittleEndian, &overUnder)
		binary.Read(buf, binary.LittleEndian, &openCircuit)
		fn(overUnder, openCircuit)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackErrorState), wrapper)
}

// Remove a registered Error State callback.
func (device *ThermocoupleV2Bricklet) DeregisterErrorStateCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackErrorState), registrationId)
}

// Returns the temperature of the thermocouple. The value is given in °C/100,
// e.g. a value of 4223 means that a temperature of 42.23 °C is measured.
//
// If you want to get the temperature periodically, it is recommended
// to use the RegisterTemperatureCallback callback and set the period with
// SetTemperatureCallbackConfiguration.
//
//
// If you want to get the value periodically, it is recommended to use the
// RegisterTemperatureCallback callback. You can set the callback configuration
// with SetTemperatureCallbackConfiguration.
func (device *ThermocoupleV2Bricklet) GetTemperature() (temperature int32, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetTemperature), buf.Bytes())
	if err != nil {
		return temperature, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return temperature, DeviceError(header.ErrorCode)
		}

		if header.Length != 12 {
			return temperature, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &temperature)

	}

	return temperature, nil
}

// The period is the period with which the RegisterTemperatureCallback callback is triggered
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
// The `option`-parameter together with min/max sets a threshold for the RegisterTemperatureCallback callback.
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
func (device *ThermocoupleV2Bricklet) SetTemperatureCallbackConfiguration(period uint32, valueHasToChange bool, option ThresholdOption, min int32, max int32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period)
	binary.Write(&buf, binary.LittleEndian, valueHasToChange)
	binary.Write(&buf, binary.LittleEndian, option)
	binary.Write(&buf, binary.LittleEndian, min)
	binary.Write(&buf, binary.LittleEndian, max)

	resultBytes, err := device.device.Set(uint8(FunctionSetTemperatureCallbackConfiguration), buf.Bytes())
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

// Returns the callback configuration as set by SetTemperatureCallbackConfiguration.
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *ThermocoupleV2Bricklet) GetTemperatureCallbackConfiguration() (period uint32, valueHasToChange bool, option ThresholdOption, min int32, max int32, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetTemperatureCallbackConfiguration), buf.Bytes())
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

// You can configure averaging size, thermocouple type and frequency
// filtering.
//
// Available averaging sizes are 1, 2, 4, 8 and 16 samples.
//
// As thermocouple type you can use B, E, J, K, N, R, S and T. If you have a
// different thermocouple or a custom thermocouple you can also use
// G8 and G32. With these types the returned value will not be in °C/100,
// it will be calculated by the following formulas:
//
// * G8: ``value = 8 * 1.6 * 2^17 * Vin``
// * G32: ``value = 32 * 1.6 * 2^17 * Vin``
//
// where Vin is the thermocouple input voltage.
//
// The frequency filter can be either configured to 50Hz or to 60Hz. You should
// configure it according to your utility frequency.
//
// The conversion time depends on the averaging and filter configuration, it can
// be calculated as follows:
//
// * 60Hz: ``time = 82 + (samples - 1) * 16.67``
// * 50Hz: ``time = 98 + (samples - 1) * 20``
//
// Associated constants:
//
//	* Averaging1
//	* Averaging2
//	* Averaging4
//	* Averaging8
//	* Averaging16
//	* TypeB
//	* TypeE
//	* TypeJ
//	* TypeK
//	* TypeN
//	* TypeR
//	* TypeS
//	* TypeT
//	* TypeG8
//	* TypeG32
//	* FilterOption50Hz
//	* FilterOption60Hz
func (device *ThermocoupleV2Bricklet) SetConfiguration(averaging Averaging, thermocoupleType Type, filter FilterOption) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, averaging)
	binary.Write(&buf, binary.LittleEndian, thermocoupleType)
	binary.Write(&buf, binary.LittleEndian, filter)

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
//	* Averaging2
//	* Averaging4
//	* Averaging8
//	* Averaging16
//	* TypeB
//	* TypeE
//	* TypeJ
//	* TypeK
//	* TypeN
//	* TypeR
//	* TypeS
//	* TypeT
//	* TypeG8
//	* TypeG32
//	* FilterOption50Hz
//	* FilterOption60Hz
func (device *ThermocoupleV2Bricklet) GetConfiguration() (averaging Averaging, thermocoupleType Type, filter FilterOption, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetConfiguration), buf.Bytes())
	if err != nil {
		return averaging, thermocoupleType, filter, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return averaging, thermocoupleType, filter, DeviceError(header.ErrorCode)
		}

		if header.Length != 11 {
			return averaging, thermocoupleType, filter, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 11)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &averaging)
		binary.Read(resultBuf, binary.LittleEndian, &thermocoupleType)
		binary.Read(resultBuf, binary.LittleEndian, &filter)

	}

	return averaging, thermocoupleType, filter, nil
}

// Returns the current error state. There are two possible errors:
//
// * Over/Under Voltage and
// * Open Circuit.
//
// Over/Under Voltage happens for voltages below 0V or above 3.3V. In this case
// it is very likely that your thermocouple is defective. An Open Circuit error
// indicates that there is no thermocouple connected.
//
// You can use the RegisterErrorStateCallback callback to automatically get triggered
// when the error state changes.
func (device *ThermocoupleV2Bricklet) GetErrorState() (overUnder bool, openCircuit bool, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetErrorState), buf.Bytes())
	if err != nil {
		return overUnder, openCircuit, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return overUnder, openCircuit, DeviceError(header.ErrorCode)
		}

		if header.Length != 10 {
			return overUnder, openCircuit, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &overUnder)
		binary.Read(resultBuf, binary.LittleEndian, &openCircuit)

	}

	return overUnder, openCircuit, nil
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
func (device *ThermocoupleV2Bricklet) GetSPITFPErrorCount() (errorCountAckChecksum uint32, errorCountMessageChecksum uint32, errorCountFrame uint32, errorCountOverflow uint32, err error) {
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
func (device *ThermocoupleV2Bricklet) SetBootloaderMode(mode BootloaderMode) (status BootloaderStatus, err error) {
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
func (device *ThermocoupleV2Bricklet) GetBootloaderMode() (mode BootloaderMode, err error) {
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
func (device *ThermocoupleV2Bricklet) SetWriteFirmwarePointer(pointer uint32) (err error) {
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
func (device *ThermocoupleV2Bricklet) WriteFirmware(data [64]uint8) (status uint8, err error) {
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
func (device *ThermocoupleV2Bricklet) SetStatusLEDConfig(config StatusLEDConfig) (err error) {
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
func (device *ThermocoupleV2Bricklet) GetStatusLEDConfig() (config StatusLEDConfig, err error) {
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
func (device *ThermocoupleV2Bricklet) GetChipTemperature() (temperature int16, err error) {
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
func (device *ThermocoupleV2Bricklet) Reset() (err error) {
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
func (device *ThermocoupleV2Bricklet) WriteUID(uid uint32) (err error) {
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
func (device *ThermocoupleV2Bricklet) ReadUID() (uid uint32, err error) {
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
func (device *ThermocoupleV2Bricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
