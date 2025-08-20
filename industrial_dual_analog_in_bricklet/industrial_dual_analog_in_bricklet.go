/* ***********************************************************
 * This file was automatically generated on 2025-08-20.      *
 *                                                           *
 * Go Bindings Version 2.0.16                                *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/

// Measures two DC voltages between -35V and +35V with 24bit resolution each.
//
//
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/IndustrialDualAnalogIn_Bricklet_Go.html.
package industrial_dual_analog_in_bricklet

import (
	"bytes"
	"encoding/binary"
	"fmt"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionGetVoltage                  Function = 1
	FunctionSetVoltageCallbackPeriod    Function = 2
	FunctionGetVoltageCallbackPeriod    Function = 3
	FunctionSetVoltageCallbackThreshold Function = 4
	FunctionGetVoltageCallbackThreshold Function = 5
	FunctionSetDebouncePeriod           Function = 6
	FunctionGetDebouncePeriod           Function = 7
	FunctionSetSampleRate               Function = 8
	FunctionGetSampleRate               Function = 9
	FunctionSetCalibration              Function = 10
	FunctionGetCalibration              Function = 11
	FunctionGetADCValues                Function = 12
	FunctionGetIdentity                 Function = 255
	FunctionCallbackVoltage             Function = 13
	FunctionCallbackVoltageReached      Function = 14
)

type ThresholdOption = rune

const (
	ThresholdOptionOff     ThresholdOption = 'x'
	ThresholdOptionOutside ThresholdOption = 'o'
	ThresholdOptionInside  ThresholdOption = 'i'
	ThresholdOptionSmaller ThresholdOption = '<'
	ThresholdOptionGreater ThresholdOption = '>'
)

type SampleRate = uint8

const (
	SampleRate976SPS SampleRate = 0
	SampleRate488SPS SampleRate = 1
	SampleRate244SPS SampleRate = 2
	SampleRate122SPS SampleRate = 3
	SampleRate61SPS  SampleRate = 4
	SampleRate4SPS   SampleRate = 5
	SampleRate2SPS   SampleRate = 6
	SampleRate1SPS   SampleRate = 7
)

type IndustrialDualAnalogInBricklet struct {
	device Device
}

const DeviceIdentifier = 249
const DeviceDisplayName = "Industrial Dual Analog In Bricklet"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (IndustrialDualAnalogInBricklet, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{2, 0, 0}, uid, &internalIPCon, 0, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return IndustrialDualAnalogInBricklet{}, err
	}
	dev.ResponseExpected[FunctionGetVoltage] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetVoltageCallbackPeriod] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionGetVoltageCallbackPeriod] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetVoltageCallbackThreshold] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionGetVoltageCallbackThreshold] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetDebouncePeriod] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionGetDebouncePeriod] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetSampleRate] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetSampleRate] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetCalibration] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetCalibration] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetADCValues] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetIdentity] = ResponseExpectedFlagAlwaysTrue
	return IndustrialDualAnalogInBricklet{dev}, nil
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
func (device *IndustrialDualAnalogInBricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *IndustrialDualAnalogInBricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *IndustrialDualAnalogInBricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *IndustrialDualAnalogInBricklet) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// This callback is triggered periodically with the period that is set by
// SetVoltageCallbackPeriod. The parameter is the voltage of the
// channel.
//
// The RegisterVoltageCallback callback is only triggered if the voltage has changed since the
// last triggering.
func (device *IndustrialDualAnalogInBricklet) RegisterVoltageCallback(fn func(uint8, int32)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 13 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var channel uint8
		var voltage int32
		binary.Read(buf, binary.LittleEndian, &channel)
		binary.Read(buf, binary.LittleEndian, &voltage)
		fn(channel, voltage)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackVoltage), wrapper)
}

// Remove a registered Voltage callback.
func (device *IndustrialDualAnalogInBricklet) DeregisterVoltageCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackVoltage), registrationId)
}

// This callback is triggered when the threshold as set by
// SetVoltageCallbackThreshold is reached.
// The parameter is the voltage of the channel.
//
// If the threshold keeps being reached, the callback is triggered periodically
// with the period as set by SetDebouncePeriod.
func (device *IndustrialDualAnalogInBricklet) RegisterVoltageReachedCallback(fn func(uint8, int32)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 13 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var channel uint8
		var voltage int32
		binary.Read(buf, binary.LittleEndian, &channel)
		binary.Read(buf, binary.LittleEndian, &voltage)
		fn(channel, voltage)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackVoltageReached), wrapper)
}

// Remove a registered Voltage Reached callback.
func (device *IndustrialDualAnalogInBricklet) DeregisterVoltageReachedCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackVoltageReached), registrationId)
}

// Returns the voltage for the given channel.
//
// If you want to get the voltage periodically, it is recommended to use the
// RegisterVoltageCallback callback and set the period with
// SetVoltageCallbackPeriod.
func (device *IndustrialDualAnalogInBricklet) GetVoltage(channel uint8) (voltage int32, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, channel)

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

// Sets the period with which the RegisterVoltageCallback callback is triggered
// periodically for the given channel. A value of 0 turns the callback off.
//
// The RegisterVoltageCallback callback is only triggered if the voltage has changed since the
// last triggering.
func (device *IndustrialDualAnalogInBricklet) SetVoltageCallbackPeriod(channel uint8, period uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, channel)
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
func (device *IndustrialDualAnalogInBricklet) GetVoltageCallbackPeriod(channel uint8) (period uint32, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, channel)

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

// Sets the thresholds for the RegisterVoltageReachedCallback callback for the given
// channel.
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
func (device *IndustrialDualAnalogInBricklet) SetVoltageCallbackThreshold(channel uint8, option ThresholdOption, min int32, max int32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, channel)
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
func (device *IndustrialDualAnalogInBricklet) GetVoltageCallbackThreshold(channel uint8) (option ThresholdOption, min int32, max int32, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, channel)

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

		if header.Length != 17 {
			return option, min, max, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 17)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &option)
		binary.Read(resultBuf, binary.LittleEndian, &min)
		binary.Read(resultBuf, binary.LittleEndian, &max)

	}

	return option, min, max, nil
}

// Sets the period with which the threshold callback
//
// * RegisterVoltageReachedCallback
//
// is triggered, if the threshold
//
// * SetVoltageCallbackThreshold
//
// keeps being reached.
func (device *IndustrialDualAnalogInBricklet) SetDebouncePeriod(debounce uint32) (err error) {
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
func (device *IndustrialDualAnalogInBricklet) GetDebouncePeriod() (debounce uint32, err error) {
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

// Sets the sample rate. The sample rate can be between 1 sample per second
// and 976 samples per second. Decreasing the sample rate will also decrease the
// noise on the data.
//
// Associated constants:
//
//	* SampleRate976SPS
//	* SampleRate488SPS
//	* SampleRate244SPS
//	* SampleRate122SPS
//	* SampleRate61SPS
//	* SampleRate4SPS
//	* SampleRate2SPS
//	* SampleRate1SPS
func (device *IndustrialDualAnalogInBricklet) SetSampleRate(rate SampleRate) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, rate)

	resultBytes, err := device.device.Set(uint8(FunctionSetSampleRate), buf.Bytes())
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

// Returns the sample rate as set by SetSampleRate.
//
// Associated constants:
//
//	* SampleRate976SPS
//	* SampleRate488SPS
//	* SampleRate244SPS
//	* SampleRate122SPS
//	* SampleRate61SPS
//	* SampleRate4SPS
//	* SampleRate2SPS
//	* SampleRate1SPS
func (device *IndustrialDualAnalogInBricklet) GetSampleRate() (rate SampleRate, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetSampleRate), buf.Bytes())
	if err != nil {
		return rate, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return rate, DeviceError(header.ErrorCode)
		}

		if header.Length != 9 {
			return rate, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &rate)

	}

	return rate, nil
}

// Sets offset and gain of MCP3911 internal calibration registers.
//
// See MCP3911 datasheet 7.7 and 7.8. The Industrial Dual Analog In Bricklet
// is already factory calibrated by Tinkerforge. It should not be necessary
// for you to use this function
func (device *IndustrialDualAnalogInBricklet) SetCalibration(offset [2]int32, gain [2]int32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, offset)
	binary.Write(&buf, binary.LittleEndian, gain)

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
func (device *IndustrialDualAnalogInBricklet) GetCalibration() (offset [2]int32, gain [2]int32, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetCalibration), buf.Bytes())
	if err != nil {
		return offset, gain, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return offset, gain, DeviceError(header.ErrorCode)
		}

		if header.Length != 24 {
			return offset, gain, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 24)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &offset)
		binary.Read(resultBuf, binary.LittleEndian, &gain)

	}

	return offset, gain, nil
}

// Returns the ADC values as given by the MCP3911 IC. This function
// is needed for proper calibration, see SetCalibration.
func (device *IndustrialDualAnalogInBricklet) GetADCValues() (value [2]int32, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetADCValues), buf.Bytes())
	if err != nil {
		return value, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return value, DeviceError(header.ErrorCode)
		}

		if header.Length != 16 {
			return value, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 16)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &value)

	}

	return value, nil
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
func (device *IndustrialDualAnalogInBricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
