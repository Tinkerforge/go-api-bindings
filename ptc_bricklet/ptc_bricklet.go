/* ***********************************************************
 * This file was automatically generated on 2025-08-20.      *
 *                                                           *
 * Go Bindings Version 2.0.16                                *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/

// Reads temperatures from Pt100 und Pt1000 sensors.
//
//
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/PTC_Bricklet_Go.html.
package ptc_bricklet

import (
	"bytes"
	"encoding/binary"
	"fmt"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionGetTemperature                          Function = 1
	FunctionGetResistance                           Function = 2
	FunctionSetTemperatureCallbackPeriod            Function = 3
	FunctionGetTemperatureCallbackPeriod            Function = 4
	FunctionSetResistanceCallbackPeriod             Function = 5
	FunctionGetResistanceCallbackPeriod             Function = 6
	FunctionSetTemperatureCallbackThreshold         Function = 7
	FunctionGetTemperatureCallbackThreshold         Function = 8
	FunctionSetResistanceCallbackThreshold          Function = 9
	FunctionGetResistanceCallbackThreshold          Function = 10
	FunctionSetDebouncePeriod                       Function = 11
	FunctionGetDebouncePeriod                       Function = 12
	FunctionSetNoiseRejectionFilter                 Function = 17
	FunctionGetNoiseRejectionFilter                 Function = 18
	FunctionIsSensorConnected                       Function = 19
	FunctionSetWireMode                             Function = 20
	FunctionGetWireMode                             Function = 21
	FunctionSetSensorConnectedCallbackConfiguration Function = 22
	FunctionGetSensorConnectedCallbackConfiguration Function = 23
	FunctionGetIdentity                             Function = 255
	FunctionCallbackTemperature                     Function = 13
	FunctionCallbackTemperatureReached              Function = 14
	FunctionCallbackResistance                      Function = 15
	FunctionCallbackResistanceReached               Function = 16
	FunctionCallbackSensorConnected                 Function = 24
)

type ThresholdOption = rune

const (
	ThresholdOptionOff     ThresholdOption = 'x'
	ThresholdOptionOutside ThresholdOption = 'o'
	ThresholdOptionInside  ThresholdOption = 'i'
	ThresholdOptionSmaller ThresholdOption = '<'
	ThresholdOptionGreater ThresholdOption = '>'
)

type FilterOption = uint8

const (
	FilterOption50Hz FilterOption = 0
	FilterOption60Hz FilterOption = 1
)

type WireMode = uint8

const (
	WireMode2 WireMode = 2
	WireMode3 WireMode = 3
	WireMode4 WireMode = 4
)

type PTCBricklet struct {
	device Device
}

const DeviceIdentifier = 226
const DeviceDisplayName = "PTC Bricklet"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (PTCBricklet, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{2, 0, 1}, uid, &internalIPCon, 0, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return PTCBricklet{}, err
	}
	dev.ResponseExpected[FunctionGetTemperature] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetResistance] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetTemperatureCallbackPeriod] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionGetTemperatureCallbackPeriod] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetResistanceCallbackPeriod] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionGetResistanceCallbackPeriod] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetTemperatureCallbackThreshold] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionGetTemperatureCallbackThreshold] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetResistanceCallbackThreshold] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionGetResistanceCallbackThreshold] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetDebouncePeriod] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionGetDebouncePeriod] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetNoiseRejectionFilter] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetNoiseRejectionFilter] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionIsSensorConnected] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetWireMode] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetWireMode] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetSensorConnectedCallbackConfiguration] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionGetSensorConnectedCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetIdentity] = ResponseExpectedFlagAlwaysTrue
	return PTCBricklet{dev}, nil
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
func (device *PTCBricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *PTCBricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *PTCBricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *PTCBricklet) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// This callback is triggered periodically with the period that is set by
// SetTemperatureCallbackPeriod. The parameter is the
// temperature of the connected sensor.
//
// The RegisterTemperatureCallback callback is only triggered if the temperature has changed
// since the last triggering.
func (device *PTCBricklet) RegisterTemperatureCallback(fn func(int32)) uint64 {
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
func (device *PTCBricklet) DeregisterTemperatureCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackTemperature), registrationId)
}

// This callback is triggered when the threshold as set by
// SetTemperatureCallbackThreshold is reached.
// The parameter is the temperature of the connected sensor.
//
// If the threshold keeps being reached, the callback is triggered periodically
// with the period as set by SetDebouncePeriod.
func (device *PTCBricklet) RegisterTemperatureReachedCallback(fn func(int32)) uint64 {
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
func (device *PTCBricklet) DeregisterTemperatureReachedCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackTemperatureReached), registrationId)
}

// This callback is triggered periodically with the period that is set by
// SetResistanceCallbackPeriod. The parameter is the resistance
// of the connected sensor.
//
// The RegisterResistanceCallback callback is only triggered if the resistance has changed
// since the last triggering.
func (device *PTCBricklet) RegisterResistanceCallback(fn func(int32)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 12 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var resistance int32
		binary.Read(buf, binary.LittleEndian, &resistance)
		fn(resistance)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackResistance), wrapper)
}

// Remove a registered Resistance callback.
func (device *PTCBricklet) DeregisterResistanceCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackResistance), registrationId)
}

// This callback is triggered when the threshold as set by
// SetResistanceCallbackThreshold is reached.
// The parameter is the resistance of the connected sensor.
//
// If the threshold keeps being reached, the callback is triggered periodically
// with the period as set by SetDebouncePeriod.
func (device *PTCBricklet) RegisterResistanceReachedCallback(fn func(int32)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 12 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var resistance int32
		binary.Read(buf, binary.LittleEndian, &resistance)
		fn(resistance)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackResistanceReached), wrapper)
}

// Remove a registered Resistance Reached callback.
func (device *PTCBricklet) DeregisterResistanceReachedCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackResistanceReached), registrationId)
}

// This callback is triggered periodically according to the configuration set by
// SetSensorConnectedCallbackConfiguration.
//
// The parameter is the same as IsSensorConnected.
//
// .. versionadded:: 2.0.2$nbsp;(Plugin)
func (device *PTCBricklet) RegisterSensorConnectedCallback(fn func(bool)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 9 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var connected bool
		binary.Read(buf, binary.LittleEndian, &connected)
		fn(connected)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackSensorConnected), wrapper)
}

// Remove a registered Sensor Connected callback.
func (device *PTCBricklet) DeregisterSensorConnectedCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackSensorConnected), registrationId)
}

// Returns the temperature of connected sensor.
//
// If you want to get the temperature periodically, it is recommended
// to use the RegisterTemperatureCallback callback and set the period with
// SetTemperatureCallbackPeriod.
func (device *PTCBricklet) GetTemperature() (temperature int32, err error) {
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

// Returns the value as measured by the MAX31865 precision delta-sigma ADC.
//
// The value can be converted with the following formulas:
//
// * Pt100:  resistance = (value * 390) / 32768
// * Pt1000: resistance = (value * 3900) / 32768
//
// If you want to get the resistance periodically, it is recommended
// to use the RegisterResistanceCallback callback and set the period with
// SetResistanceCallbackPeriod.
func (device *PTCBricklet) GetResistance() (resistance int32, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetResistance), buf.Bytes())
	if err != nil {
		return resistance, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return resistance, DeviceError(header.ErrorCode)
		}

		if header.Length != 12 {
			return resistance, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &resistance)

	}

	return resistance, nil
}

// Sets the period with which the RegisterTemperatureCallback callback is triggered
// periodically. A value of 0 turns the callback off.
//
// The RegisterTemperatureCallback callback is only triggered if the temperature has
// changed since the last triggering.
func (device *PTCBricklet) SetTemperatureCallbackPeriod(period uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period)

	resultBytes, err := device.device.Set(uint8(FunctionSetTemperatureCallbackPeriod), buf.Bytes())
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

// Returns the period as set by SetTemperatureCallbackPeriod.
func (device *PTCBricklet) GetTemperatureCallbackPeriod() (period uint32, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetTemperatureCallbackPeriod), buf.Bytes())
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

// Sets the period with which the RegisterResistanceCallback callback is triggered
// periodically. A value of 0 turns the callback off.
//
// The RegisterResistanceCallback callback is only triggered if the resistance has changed
// since the last triggering.
func (device *PTCBricklet) SetResistanceCallbackPeriod(period uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period)

	resultBytes, err := device.device.Set(uint8(FunctionSetResistanceCallbackPeriod), buf.Bytes())
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

// Returns the period as set by SetResistanceCallbackPeriod.
func (device *PTCBricklet) GetResistanceCallbackPeriod() (period uint32, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetResistanceCallbackPeriod), buf.Bytes())
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
func (device *PTCBricklet) SetTemperatureCallbackThreshold(option ThresholdOption, min int32, max int32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, option)
	binary.Write(&buf, binary.LittleEndian, min)
	binary.Write(&buf, binary.LittleEndian, max)

	resultBytes, err := device.device.Set(uint8(FunctionSetTemperatureCallbackThreshold), buf.Bytes())
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

// Returns the threshold as set by SetTemperatureCallbackThreshold.
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *PTCBricklet) GetTemperatureCallbackThreshold() (option ThresholdOption, min int32, max int32, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetTemperatureCallbackThreshold), buf.Bytes())
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

// Sets the thresholds for the RegisterResistanceReachedCallback callback.
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
func (device *PTCBricklet) SetResistanceCallbackThreshold(option ThresholdOption, min int32, max int32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, option)
	binary.Write(&buf, binary.LittleEndian, min)
	binary.Write(&buf, binary.LittleEndian, max)

	resultBytes, err := device.device.Set(uint8(FunctionSetResistanceCallbackThreshold), buf.Bytes())
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

// Returns the threshold as set by SetResistanceCallbackThreshold.
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *PTCBricklet) GetResistanceCallbackThreshold() (option ThresholdOption, min int32, max int32, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetResistanceCallbackThreshold), buf.Bytes())
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
// * RegisterTemperatureReachedCallback,
// * RegisterResistanceReachedCallback
//
// is triggered, if the threshold
//
// * SetTemperatureCallbackThreshold,
// * SetResistanceCallbackThreshold
//
// keeps being reached.
func (device *PTCBricklet) SetDebouncePeriod(debounce uint32) (err error) {
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
func (device *PTCBricklet) GetDebouncePeriod() (debounce uint32, err error) {
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

// Sets the noise rejection filter to either 50Hz (0) or 60Hz (1).
// Noise from 50Hz or 60Hz power sources (including
// harmonics of the AC power's fundamental frequency) is
// attenuated by 82dB.
//
// Associated constants:
//
//	* FilterOption50Hz
//	* FilterOption60Hz
func (device *PTCBricklet) SetNoiseRejectionFilter(filter FilterOption) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, filter)

	resultBytes, err := device.device.Set(uint8(FunctionSetNoiseRejectionFilter), buf.Bytes())
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

// Returns the noise rejection filter option as set by
// SetNoiseRejectionFilter
//
// Associated constants:
//
//	* FilterOption50Hz
//	* FilterOption60Hz
func (device *PTCBricklet) GetNoiseRejectionFilter() (filter FilterOption, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetNoiseRejectionFilter), buf.Bytes())
	if err != nil {
		return filter, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return filter, DeviceError(header.ErrorCode)
		}

		if header.Length != 9 {
			return filter, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &filter)

	}

	return filter, nil
}

// Returns *true* if the sensor is connected correctly.
//
// If this function
// returns *false*, there is either no Pt100 or Pt1000 sensor connected,
// the sensor is connected incorrectly or the sensor itself is faulty.
func (device *PTCBricklet) IsSensorConnected() (connected bool, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionIsSensorConnected), buf.Bytes())
	if err != nil {
		return connected, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return connected, DeviceError(header.ErrorCode)
		}

		if header.Length != 9 {
			return connected, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &connected)

	}

	return connected, nil
}

// Sets the wire mode of the sensor. Possible values are 2, 3 and 4 which
// correspond to 2-, 3- and 4-wire sensors. The value has to match the jumper
// configuration on the Bricklet.
//
// Associated constants:
//
//	* WireMode2
//	* WireMode3
//	* WireMode4
func (device *PTCBricklet) SetWireMode(mode WireMode) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, mode)

	resultBytes, err := device.device.Set(uint8(FunctionSetWireMode), buf.Bytes())
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

// Returns the wire mode as set by SetWireMode
//
// Associated constants:
//
//	* WireMode2
//	* WireMode3
//	* WireMode4
func (device *PTCBricklet) GetWireMode() (mode WireMode, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetWireMode), buf.Bytes())
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

// If you enable this callback, the RegisterSensorConnectedCallback callback is triggered
// every time a Pt sensor is connected/disconnected.
//
// .. versionadded:: 2.0.2$nbsp;(Plugin)
func (device *PTCBricklet) SetSensorConnectedCallbackConfiguration(enabled bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, enabled)

	resultBytes, err := device.device.Set(uint8(FunctionSetSensorConnectedCallbackConfiguration), buf.Bytes())
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

// Returns the configuration as set by SetSensorConnectedCallbackConfiguration.
//
// .. versionadded:: 2.0.2$nbsp;(Plugin)
func (device *PTCBricklet) GetSensorConnectedCallbackConfiguration() (enabled bool, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetSensorConnectedCallbackConfiguration), buf.Bytes())
	if err != nil {
		return enabled, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return enabled, DeviceError(header.ErrorCode)
		}

		if header.Length != 9 {
			return enabled, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &enabled)

	}

	return enabled, nil
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
func (device *PTCBricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
