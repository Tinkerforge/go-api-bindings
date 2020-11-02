/* ***********************************************************
 * This file was automatically generated on 2020-11-02.      *
 *                                                           *
 * Go Bindings Version 2.0.9                                 *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


// Measures temperature with thermocouples.
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/Thermocouple_Bricklet_Go.html.
package thermocouple_bricklet

import (
	"encoding/binary"
	"bytes"
	"fmt"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionGetTemperature Function = 1
	FunctionSetTemperatureCallbackPeriod Function = 2
	FunctionGetTemperatureCallbackPeriod Function = 3
	FunctionSetTemperatureCallbackThreshold Function = 4
	FunctionGetTemperatureCallbackThreshold Function = 5
	FunctionSetDebouncePeriod Function = 6
	FunctionGetDebouncePeriod Function = 7
	FunctionSetConfiguration Function = 10
	FunctionGetConfiguration Function = 11
	FunctionGetErrorState Function = 12
	FunctionGetIdentity Function = 255
	FunctionCallbackTemperature Function = 8
	FunctionCallbackTemperatureReached Function = 9
	FunctionCallbackErrorState Function = 13
)

type ThresholdOption = rune

const (
	ThresholdOptionOff ThresholdOption = 'x'
	ThresholdOptionOutside ThresholdOption = 'o'
	ThresholdOptionInside ThresholdOption = 'i'
	ThresholdOptionSmaller ThresholdOption = '<'
	ThresholdOptionGreater ThresholdOption = '>'
)

type Averaging = uint8

const (
	Averaging1 Averaging = 1
	Averaging2 Averaging = 2
	Averaging4 Averaging = 4
	Averaging8 Averaging = 8
	Averaging16 Averaging = 16
)

type Type = uint8

const (
	TypeB Type = 0
	TypeE Type = 1
	TypeJ Type = 2
	TypeK Type = 3
	TypeN Type = 4
	TypeR Type = 5
	TypeS Type = 6
	TypeT Type = 7
	TypeG8 Type = 8
	TypeG32 Type = 9
)

type FilterOption = uint8

const (
	FilterOption50Hz FilterOption = 0
	FilterOption60Hz FilterOption = 1
)

type ThermocoupleBricklet struct {
	device Device
}
const DeviceIdentifier = 266
const DeviceDisplayName = "Thermocouple Bricklet"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (ThermocoupleBricklet, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{ 2,0,0 }, uid, &internalIPCon, 0, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return ThermocoupleBricklet{}, err
	}
	dev.ResponseExpected[FunctionGetTemperature] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetTemperatureCallbackPeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetTemperatureCallbackPeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetTemperatureCallbackThreshold] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetTemperatureCallbackThreshold] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetDebouncePeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetDebouncePeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetConfiguration] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetErrorState] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetIdentity] = ResponseExpectedFlagAlwaysTrue;
	return ThermocoupleBricklet{dev}, nil
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
func (device *ThermocoupleBricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *ThermocoupleBricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *ThermocoupleBricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *ThermocoupleBricklet) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// This callback is triggered periodically with the period that is set by
// SetTemperatureCallbackPeriod. The parameter is the
// temperature of the thermocouple.
// 
// The RegisterTemperatureCallback callback is only triggered if the temperature has
// changed since the last triggering.
func (device *ThermocoupleBricklet) RegisterTemperatureCallback(fn func(int32)) uint64 {
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
func (device *ThermocoupleBricklet) DeregisterTemperatureCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackTemperature), registrationId)
}


// This callback is triggered when the threshold as set by
// SetTemperatureCallbackThreshold is reached.
// The parameter is the temperature of the thermocouple.
// 
// If the threshold keeps being reached, the callback is triggered periodically
// with the period as set by SetDebouncePeriod.
func (device *ThermocoupleBricklet) RegisterTemperatureReachedCallback(fn func(int32)) uint64 {
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
	return device.device.RegisterCallback(uint8(FunctionCallbackTemperatureReached), wrapper)
}

// Remove a registered Temperature Reached callback.
func (device *ThermocoupleBricklet) DeregisterTemperatureReachedCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackTemperatureReached), registrationId)
}


// This Callback is triggered every time the error state changes
// (see GetErrorState).
func (device *ThermocoupleBricklet) RegisterErrorStateCallback(fn func(bool, bool)) uint64 {
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
func (device *ThermocoupleBricklet) DeregisterErrorStateCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackErrorState), registrationId)
}


// Returns the temperature of the thermocouple.
// 
// If you want to get the temperature periodically, it is recommended
// to use the RegisterTemperatureCallback callback and set the period with
// SetTemperatureCallbackPeriod.
func (device *ThermocoupleBricklet) GetTemperature() (temperature int32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetTemperature), buf.Bytes())
	if err != nil {
		return temperature, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 12 {
			return temperature, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		if header.ErrorCode != 0 {
			return temperature, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &temperature)

	}

	return temperature, nil
}

// Sets the period with which the RegisterTemperatureCallback callback is triggered
// periodically. A value of 0 turns the callback off.
// 
// The RegisterTemperatureCallback callback is only triggered if the temperature has changed
// since the last triggering.
func (device *ThermocoupleBricklet) SetTemperatureCallbackPeriod(period uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);

	resultBytes, err := device.device.Set(uint8(FunctionSetTemperatureCallbackPeriod), buf.Bytes())
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

// Returns the period as set by SetTemperatureCallbackPeriod.
func (device *ThermocoupleBricklet) GetTemperatureCallbackPeriod() (period uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetTemperatureCallbackPeriod), buf.Bytes())
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

// Sets the thresholds for the RegisterTemperatureReachedCallback callback.
// 
// The following options are possible:
// 
//  Option| Description
//  --- | --- 
//  'x'|    Callback is turned off
//  'o'|    Callback is triggered when the temperature is *outside* the min and max values
//  'i'|    Callback is triggered when the temperature is *inside* the min and max values
//  '<'|    Callback is triggered when the temperature is smaller than the min value (max is ignored)
//  '>'|    Callback is triggered when the temperature is greater than the min value (max is ignored)
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *ThermocoupleBricklet) SetTemperatureCallbackThreshold(option ThresholdOption, min int32, max int32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, option);
	binary.Write(&buf, binary.LittleEndian, min);
	binary.Write(&buf, binary.LittleEndian, max);

	resultBytes, err := device.device.Set(uint8(FunctionSetTemperatureCallbackThreshold), buf.Bytes())
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

// Returns the threshold as set by SetTemperatureCallbackThreshold.
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *ThermocoupleBricklet) GetTemperatureCallbackThreshold() (option ThresholdOption, min int32, max int32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetTemperatureCallbackThreshold), buf.Bytes())
	if err != nil {
		return option, min, max, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 17 {
			return option, min, max, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 17)
		}

		if header.ErrorCode != 0 {
			return option, min, max, DeviceError(header.ErrorCode)
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
// * RegisterTemperatureReachedCallback
// 
// is triggered, if the threshold
// 
// * SetTemperatureCallbackThreshold
// 
// keeps being reached.
func (device *ThermocoupleBricklet) SetDebouncePeriod(debounce uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, debounce);

	resultBytes, err := device.device.Set(uint8(FunctionSetDebouncePeriod), buf.Bytes())
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

// Returns the debounce period as set by SetDebouncePeriod.
func (device *ThermocoupleBricklet) GetDebouncePeriod() (debounce uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetDebouncePeriod), buf.Bytes())
	if err != nil {
		return debounce, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 12 {
			return debounce, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		if header.ErrorCode != 0 {
			return debounce, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &debounce)

	}

	return debounce, nil
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
func (device *ThermocoupleBricklet) SetConfiguration(averaging Averaging, thermocoupleType Type, filter FilterOption) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, averaging);
	binary.Write(&buf, binary.LittleEndian, thermocoupleType);
	binary.Write(&buf, binary.LittleEndian, filter);

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
func (device *ThermocoupleBricklet) GetConfiguration() (averaging Averaging, thermocoupleType Type, filter FilterOption, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetConfiguration), buf.Bytes())
	if err != nil {
		return averaging, thermocoupleType, filter, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 11 {
			return averaging, thermocoupleType, filter, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 11)
		}

		if header.ErrorCode != 0 {
			return averaging, thermocoupleType, filter, DeviceError(header.ErrorCode)
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
func (device *ThermocoupleBricklet) GetErrorState() (overUnder bool, openCircuit bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetErrorState), buf.Bytes())
	if err != nil {
		return overUnder, openCircuit, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 10 {
			return overUnder, openCircuit, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		if header.ErrorCode != 0 {
			return overUnder, openCircuit, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &overUnder)
		binary.Read(resultBuf, binary.LittleEndian, &openCircuit)

	}

	return overUnder, openCircuit, nil
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
func (device *ThermocoupleBricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
