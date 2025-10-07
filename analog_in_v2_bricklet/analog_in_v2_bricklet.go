/* ***********************************************************
 * This file was automatically generated on 2025-10-07.      *
 *                                                           *
 * Go Bindings Version 2.0.17                                *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/

// Measures DC voltage between 0V and 42V‍.
//
//
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/AnalogInV2_Bricklet_Go.html.
package analog_in_v2_bricklet

import (
	"bytes"
	"encoding/binary"
	"fmt"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionGetVoltage                      Function = 1
	FunctionGetAnalogValue                  Function = 2
	FunctionSetVoltageCallbackPeriod        Function = 3
	FunctionGetVoltageCallbackPeriod        Function = 4
	FunctionSetAnalogValueCallbackPeriod    Function = 5
	FunctionGetAnalogValueCallbackPeriod    Function = 6
	FunctionSetVoltageCallbackThreshold     Function = 7
	FunctionGetVoltageCallbackThreshold     Function = 8
	FunctionSetAnalogValueCallbackThreshold Function = 9
	FunctionGetAnalogValueCallbackThreshold Function = 10
	FunctionSetDebouncePeriod               Function = 11
	FunctionGetDebouncePeriod               Function = 12
	FunctionSetMovingAverage                Function = 13
	FunctionGetMovingAverage                Function = 14
	FunctionGetIdentity                     Function = 255
	FunctionCallbackVoltage                 Function = 15
	FunctionCallbackAnalogValue             Function = 16
	FunctionCallbackVoltageReached          Function = 17
	FunctionCallbackAnalogValueReached      Function = 18
)

type ThresholdOption = rune

const (
	ThresholdOptionOff     ThresholdOption = 'x'
	ThresholdOptionOutside ThresholdOption = 'o'
	ThresholdOptionInside  ThresholdOption = 'i'
	ThresholdOptionSmaller ThresholdOption = '<'
	ThresholdOptionGreater ThresholdOption = '>'
)

type AnalogInV2Bricklet struct {
	device Device
}

const DeviceIdentifier = 251
const DeviceDisplayName = "Analog In Bricklet 2.0"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (AnalogInV2Bricklet, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{2, 0, 1}, uid, &internalIPCon, 0, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return AnalogInV2Bricklet{}, err
	}
	dev.ResponseExpected[FunctionGetVoltage] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetAnalogValue] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetVoltageCallbackPeriod] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionGetVoltageCallbackPeriod] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetAnalogValueCallbackPeriod] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionGetAnalogValueCallbackPeriod] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetVoltageCallbackThreshold] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionGetVoltageCallbackThreshold] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetAnalogValueCallbackThreshold] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionGetAnalogValueCallbackThreshold] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetDebouncePeriod] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionGetDebouncePeriod] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetMovingAverage] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetMovingAverage] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetIdentity] = ResponseExpectedFlagAlwaysTrue
	return AnalogInV2Bricklet{dev}, nil
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
func (device *AnalogInV2Bricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *AnalogInV2Bricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *AnalogInV2Bricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *AnalogInV2Bricklet) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// This callback is triggered periodically with the period that is set by
// SetVoltageCallbackPeriod. The parameter is the voltage of the
// sensor.
//
// Der RegisterVoltageCallback callback is only triggered if the voltage has changed since
// the last triggering.
func (device *AnalogInV2Bricklet) RegisterVoltageCallback(fn func(uint16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 10 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var voltage uint16
		binary.Read(buf, binary.LittleEndian, &voltage)
		fn(voltage)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackVoltage), wrapper)
}

// Remove a registered Voltage callback.
func (device *AnalogInV2Bricklet) DeregisterVoltageCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackVoltage), registrationId)
}

// This callback is triggered periodically with the period that is set by
// SetAnalogValueCallbackPeriod. The parameter is the analog
// value of the sensor.
//
// The RegisterAnalogValueCallback callback is only triggered if the voltage has changed
// since the last triggering.
func (device *AnalogInV2Bricklet) RegisterAnalogValueCallback(fn func(uint16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 10 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var value uint16
		binary.Read(buf, binary.LittleEndian, &value)
		fn(value)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackAnalogValue), wrapper)
}

// Remove a registered Analog Value callback.
func (device *AnalogInV2Bricklet) DeregisterAnalogValueCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackAnalogValue), registrationId)
}

// This callback is triggered when the threshold as set by
// SetVoltageCallbackThreshold is reached.
// The parameter is the voltage of the sensor.
//
// If the threshold keeps being reached, the callback is triggered periodically
// with the period as set by SetDebouncePeriod.
func (device *AnalogInV2Bricklet) RegisterVoltageReachedCallback(fn func(uint16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 10 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var voltage uint16
		binary.Read(buf, binary.LittleEndian, &voltage)
		fn(voltage)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackVoltageReached), wrapper)
}

// Remove a registered Voltage Reached callback.
func (device *AnalogInV2Bricklet) DeregisterVoltageReachedCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackVoltageReached), registrationId)
}

// This callback is triggered when the threshold as set by
// SetAnalogValueCallbackThreshold is reached.
// The parameter is the analog value of the sensor.
//
// If the threshold keeps being reached, the callback is triggered periodically
// with the period as set by SetDebouncePeriod.
func (device *AnalogInV2Bricklet) RegisterAnalogValueReachedCallback(fn func(uint16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 10 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var value uint16
		binary.Read(buf, binary.LittleEndian, &value)
		fn(value)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackAnalogValueReached), wrapper)
}

// Remove a registered Analog Value Reached callback.
func (device *AnalogInV2Bricklet) DeregisterAnalogValueReachedCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackAnalogValueReached), registrationId)
}

// Returns the measured voltage. The resolution is approximately 10mV.
//
// If you want to get the voltage periodically, it is recommended to use the
// RegisterVoltageCallback callback and set the period with
// SetVoltageCallbackPeriod.
func (device *AnalogInV2Bricklet) GetVoltage() (voltage uint16, err error) {
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

		if header.Length != 10 {
			return voltage, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &voltage)

	}

	return voltage, nil
}

// Returns the value as read by a 12-bit analog-to-digital converter.
//
// If you want the analog value periodically, it is recommended to use the
// RegisterAnalogValueCallback callback and set the period with
// SetAnalogValueCallbackPeriod.
func (device *AnalogInV2Bricklet) GetAnalogValue() (value uint16, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetAnalogValue), buf.Bytes())
	if err != nil {
		return value, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return value, DeviceError(header.ErrorCode)
		}

		if header.Length != 10 {
			return value, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &value)

	}

	return value, nil
}

// Sets the period with which the RegisterVoltageCallback callback is triggered
// periodically. A value of 0 turns the callback off.
//
// The RegisterVoltageCallback callback is only triggered if the voltage has changed since
// the last triggering.
func (device *AnalogInV2Bricklet) SetVoltageCallbackPeriod(period uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period)

	resultBytes, err := device.device.Set(uint8(FunctionSetVoltageCallbackPeriod), buf.Bytes())
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

// Returns the period as set by SetVoltageCallbackPeriod.
func (device *AnalogInV2Bricklet) GetVoltageCallbackPeriod() (period uint32, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetVoltageCallbackPeriod), buf.Bytes())
	if err != nil {
		return period, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return period, DeviceError(header.ErrorCode)
		}

		if header.Length != 12 {
			return period, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &period)

	}

	return period, nil
}

// Sets the period with which the RegisterAnalogValueCallback callback is triggered
// periodically. A value of 0 turns the callback off.
//
// The RegisterAnalogValueCallback callback is only triggered if the analog value has
// changed since the last triggering.
func (device *AnalogInV2Bricklet) SetAnalogValueCallbackPeriod(period uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period)

	resultBytes, err := device.device.Set(uint8(FunctionSetAnalogValueCallbackPeriod), buf.Bytes())
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

// Returns the period as set by SetAnalogValueCallbackPeriod.
func (device *AnalogInV2Bricklet) GetAnalogValueCallbackPeriod() (period uint32, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetAnalogValueCallbackPeriod), buf.Bytes())
	if err != nil {
		return period, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return period, DeviceError(header.ErrorCode)
		}

		if header.Length != 12 {
			return period, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &period)

	}

	return period, nil
}

// Sets the thresholds for the RegisterVoltageReachedCallback callback.
//
// The following options are possible:
//
//  Option| Description
//  --- | ---
//  'x'|    Callback is turned off
//  'o'|    Callback is triggered when the voltage is *outside* the min and max values
//  'i'|    Callback is triggered when the voltage is *inside* the min and max values
//  '<'|    Callback is triggered when the voltage is smaller than the min value (max is ignored)
//  '>'|    Callback is triggered when the voltage is greater than the min value (max is ignored)
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *AnalogInV2Bricklet) SetVoltageCallbackThreshold(option ThresholdOption, min uint16, max uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, option)
	binary.Write(&buf, binary.LittleEndian, min)
	binary.Write(&buf, binary.LittleEndian, max)

	resultBytes, err := device.device.Set(uint8(FunctionSetVoltageCallbackThreshold), buf.Bytes())
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

// Returns the threshold as set by SetVoltageCallbackThreshold.
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *AnalogInV2Bricklet) GetVoltageCallbackThreshold() (option ThresholdOption, min uint16, max uint16, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetVoltageCallbackThreshold), buf.Bytes())
	if err != nil {
		return option, min, max, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return option, min, max, DeviceError(header.ErrorCode)
		}

		if header.Length != 13 {
			return option, min, max, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 13)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &option)
		binary.Read(resultBuf, binary.LittleEndian, &min)
		binary.Read(resultBuf, binary.LittleEndian, &max)

	}

	return option, min, max, nil
}

// Sets the thresholds for the RegisterAnalogValueReachedCallback callback.
//
// The following options are possible:
//
//  Option| Description
//  --- | ---
//  'x'|    Callback is turned off
//  'o'|    Callback is triggered when the analog value is *outside* the min and max values
//  'i'|    Callback is triggered when the analog value is *inside* the min and max values
//  '<'|    Callback is triggered when the analog value is smaller than the min value (max is ignored)
//  '>'|    Callback is triggered when the analog value is greater than the min value (max is ignored)
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *AnalogInV2Bricklet) SetAnalogValueCallbackThreshold(option ThresholdOption, min uint16, max uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, option)
	binary.Write(&buf, binary.LittleEndian, min)
	binary.Write(&buf, binary.LittleEndian, max)

	resultBytes, err := device.device.Set(uint8(FunctionSetAnalogValueCallbackThreshold), buf.Bytes())
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

// Returns the threshold as set by SetAnalogValueCallbackThreshold.
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *AnalogInV2Bricklet) GetAnalogValueCallbackThreshold() (option ThresholdOption, min uint16, max uint16, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetAnalogValueCallbackThreshold), buf.Bytes())
	if err != nil {
		return option, min, max, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return option, min, max, DeviceError(header.ErrorCode)
		}

		if header.Length != 13 {
			return option, min, max, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 13)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &option)
		binary.Read(resultBuf, binary.LittleEndian, &min)
		binary.Read(resultBuf, binary.LittleEndian, &max)

	}

	return option, min, max, nil
}

// Sets the period with which the threshold callbacks
//
// * RegisterVoltageReachedCallback,
// * RegisterAnalogValueReachedCallback
//
// are triggered, if the thresholds
//
// * SetVoltageCallbackThreshold,
// * SetAnalogValueCallbackThreshold
//
// keep being reached.
func (device *AnalogInV2Bricklet) SetDebouncePeriod(debounce uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, debounce)

	resultBytes, err := device.device.Set(uint8(FunctionSetDebouncePeriod), buf.Bytes())
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

// Returns the debounce period as set by SetDebouncePeriod.
func (device *AnalogInV2Bricklet) GetDebouncePeriod() (debounce uint32, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetDebouncePeriod), buf.Bytes())
	if err != nil {
		return debounce, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return debounce, DeviceError(header.ErrorCode)
		}

		if header.Length != 12 {
			return debounce, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &debounce)

	}

	return debounce, nil
}

// Sets the length of a https://en.wikipedia.org/wiki/Moving_average
// for the voltage.
//
// Setting the length to 1 will turn the averaging off. With less
// averaging, there is more noise on the data.
func (device *AnalogInV2Bricklet) SetMovingAverage(average uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, average)

	resultBytes, err := device.device.Set(uint8(FunctionSetMovingAverage), buf.Bytes())
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

// Returns the length of the moving average as set by SetMovingAverage.
func (device *AnalogInV2Bricklet) GetMovingAverage() (average uint8, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetMovingAverage), buf.Bytes())
	if err != nil {
		return average, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return average, DeviceError(header.ErrorCode)
		}

		if header.Length != 9 {
			return average, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &average)

	}

	return average, nil
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
func (device *AnalogInV2Bricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
