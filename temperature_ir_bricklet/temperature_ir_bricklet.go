/* ***********************************************************
 * This file was automatically generated on 2025-08-20.      *
 *                                                           *
 * Go Bindings Version 2.0.16                                *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/

// Measures contactless object temperature between -70°C and +380°C‍.
//
//
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/TemperatureIR_Bricklet_Go.html.
package temperature_ir_bricklet

import (
	"bytes"
	"encoding/binary"
	"fmt"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionGetAmbientTemperature                  Function = 1
	FunctionGetObjectTemperature                   Function = 2
	FunctionSetEmissivity                          Function = 3
	FunctionGetEmissivity                          Function = 4
	FunctionSetAmbientTemperatureCallbackPeriod    Function = 5
	FunctionGetAmbientTemperatureCallbackPeriod    Function = 6
	FunctionSetObjectTemperatureCallbackPeriod     Function = 7
	FunctionGetObjectTemperatureCallbackPeriod     Function = 8
	FunctionSetAmbientTemperatureCallbackThreshold Function = 9
	FunctionGetAmbientTemperatureCallbackThreshold Function = 10
	FunctionSetObjectTemperatureCallbackThreshold  Function = 11
	FunctionGetObjectTemperatureCallbackThreshold  Function = 12
	FunctionSetDebouncePeriod                      Function = 13
	FunctionGetDebouncePeriod                      Function = 14
	FunctionGetIdentity                            Function = 255
	FunctionCallbackAmbientTemperature             Function = 15
	FunctionCallbackObjectTemperature              Function = 16
	FunctionCallbackAmbientTemperatureReached      Function = 17
	FunctionCallbackObjectTemperatureReached       Function = 18
)

type ThresholdOption = rune

const (
	ThresholdOptionOff     ThresholdOption = 'x'
	ThresholdOptionOutside ThresholdOption = 'o'
	ThresholdOptionInside  ThresholdOption = 'i'
	ThresholdOptionSmaller ThresholdOption = '<'
	ThresholdOptionGreater ThresholdOption = '>'
)

type TemperatureIRBricklet struct {
	device Device
}

const DeviceIdentifier = 217
const DeviceDisplayName = "Temperature IR Bricklet"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (TemperatureIRBricklet, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{2, 0, 0}, uid, &internalIPCon, 0, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return TemperatureIRBricklet{}, err
	}
	dev.ResponseExpected[FunctionGetAmbientTemperature] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetObjectTemperature] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetEmissivity] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetEmissivity] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetAmbientTemperatureCallbackPeriod] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionGetAmbientTemperatureCallbackPeriod] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetObjectTemperatureCallbackPeriod] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionGetObjectTemperatureCallbackPeriod] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetAmbientTemperatureCallbackThreshold] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionGetAmbientTemperatureCallbackThreshold] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetObjectTemperatureCallbackThreshold] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionGetObjectTemperatureCallbackThreshold] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetDebouncePeriod] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionGetDebouncePeriod] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetIdentity] = ResponseExpectedFlagAlwaysTrue
	return TemperatureIRBricklet{dev}, nil
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
func (device *TemperatureIRBricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *TemperatureIRBricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *TemperatureIRBricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *TemperatureIRBricklet) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// This callback is triggered periodically with the period that is set by
// SetAmbientTemperatureCallbackPeriod. The parameter is the
// ambient temperature of the sensor.
//
// The RegisterAmbientTemperatureCallback callback is only triggered if the ambient
// temperature has changed since the last triggering.
func (device *TemperatureIRBricklet) RegisterAmbientTemperatureCallback(fn func(int16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 10 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var temperature int16
		binary.Read(buf, binary.LittleEndian, &temperature)
		fn(temperature)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackAmbientTemperature), wrapper)
}

// Remove a registered Ambient Temperature callback.
func (device *TemperatureIRBricklet) DeregisterAmbientTemperatureCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackAmbientTemperature), registrationId)
}

// This callback is triggered periodically with the period that is set by
// SetObjectTemperatureCallbackPeriod. The parameter is the
// object temperature of the sensor.
//
// The RegisterObjectTemperatureCallback callback is only triggered if the object
// temperature has changed since the last triggering.
func (device *TemperatureIRBricklet) RegisterObjectTemperatureCallback(fn func(int16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 10 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var temperature int16
		binary.Read(buf, binary.LittleEndian, &temperature)
		fn(temperature)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackObjectTemperature), wrapper)
}

// Remove a registered Object Temperature callback.
func (device *TemperatureIRBricklet) DeregisterObjectTemperatureCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackObjectTemperature), registrationId)
}

// This callback is triggered when the threshold as set by
// SetAmbientTemperatureCallbackThreshold is reached.
// The parameter is the ambient temperature of the sensor.
//
// If the threshold keeps being reached, the callback is triggered periodically
// with the period as set by SetDebouncePeriod.
func (device *TemperatureIRBricklet) RegisterAmbientTemperatureReachedCallback(fn func(int16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 10 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var temperature int16
		binary.Read(buf, binary.LittleEndian, &temperature)
		fn(temperature)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackAmbientTemperatureReached), wrapper)
}

// Remove a registered Ambient Temperature Reached callback.
func (device *TemperatureIRBricklet) DeregisterAmbientTemperatureReachedCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackAmbientTemperatureReached), registrationId)
}

// This callback is triggered when the threshold as set by
// SetObjectTemperatureCallbackThreshold is reached.
// The parameter is the object temperature of the sensor.
//
// If the threshold keeps being reached, the callback is triggered periodically
// with the period as set by SetDebouncePeriod.
func (device *TemperatureIRBricklet) RegisterObjectTemperatureReachedCallback(fn func(int16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 10 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var temperature int16
		binary.Read(buf, binary.LittleEndian, &temperature)
		fn(temperature)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackObjectTemperatureReached), wrapper)
}

// Remove a registered Object Temperature Reached callback.
func (device *TemperatureIRBricklet) DeregisterObjectTemperatureReachedCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackObjectTemperatureReached), registrationId)
}

// Returns the ambient temperature of the sensor.
//
// If you want to get the ambient temperature periodically, it is recommended
// to use the RegisterAmbientTemperatureCallback callback and set the period with
// SetAmbientTemperatureCallbackPeriod.
func (device *TemperatureIRBricklet) GetAmbientTemperature() (temperature int16, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetAmbientTemperature), buf.Bytes())
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

// Returns the object temperature of the sensor, i.e. the temperature
// of the surface of the object the sensor is aimed at.
//
// The temperature of different materials is dependent on their `emissivity
// <https://en.wikipedia.org/wiki/Emissivity>`__. The emissivity of the material
// can be set with SetEmissivity.
//
// If you want to get the object temperature periodically, it is recommended
// to use the RegisterObjectTemperatureCallback callback and set the period with
// SetObjectTemperatureCallbackPeriod.
func (device *TemperatureIRBricklet) GetObjectTemperature() (temperature int16, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetObjectTemperature), buf.Bytes())
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

// Sets the https://en.wikipedia.org/wiki/Emissivity that is
// used to calculate the surface temperature as returned by
// :func:https://www.infrared-thermography.com/material.htm.
//
// The parameter of SetEmissivity has to be given with a factor of
// 65535 (16-bit). For example: An emissivity of 0.1 can be set with the
// value 6553, an emissivity of 0.5 with the value 32767 and so on.
//
// Note
//  If you need a precise measurement for the object temperature, it is
//  absolutely crucial that you also provide a precise emissivity.
//
// The emissivity is stored in non-volatile memory and will still be used after a restart or power cycle of the Bricklet.
func (device *TemperatureIRBricklet) SetEmissivity(emissivity uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, emissivity)

	resultBytes, err := device.device.Set(uint8(FunctionSetEmissivity), buf.Bytes())
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

// Returns the emissivity as set by SetEmissivity.
func (device *TemperatureIRBricklet) GetEmissivity() (emissivity uint16, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetEmissivity), buf.Bytes())
	if err != nil {
		return emissivity, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return emissivity, DeviceError(header.ErrorCode)
		}

		if header.Length != 10 {
			return emissivity, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &emissivity)

	}

	return emissivity, nil
}

// Sets the period with which the RegisterAmbientTemperatureCallback callback is
// triggered periodically. A value of 0 turns the callback off.
//
// The RegisterAmbientTemperatureCallback callback is only triggered if the temperature has
// changed since the last triggering.
func (device *TemperatureIRBricklet) SetAmbientTemperatureCallbackPeriod(period uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period)

	resultBytes, err := device.device.Set(uint8(FunctionSetAmbientTemperatureCallbackPeriod), buf.Bytes())
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

// Returns the period as set by SetAmbientTemperatureCallbackPeriod.
func (device *TemperatureIRBricklet) GetAmbientTemperatureCallbackPeriod() (period uint32, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetAmbientTemperatureCallbackPeriod), buf.Bytes())
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

// Sets the period with which the RegisterObjectTemperatureCallback callback is
// triggered periodically. A value of 0 turns the callback off.
//
// The RegisterObjectTemperatureCallback callback is only triggered if the temperature
// has changed since the last triggering.
func (device *TemperatureIRBricklet) SetObjectTemperatureCallbackPeriod(period uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period)

	resultBytes, err := device.device.Set(uint8(FunctionSetObjectTemperatureCallbackPeriod), buf.Bytes())
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

// Returns the period as set by SetObjectTemperatureCallbackPeriod.
func (device *TemperatureIRBricklet) GetObjectTemperatureCallbackPeriod() (period uint32, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetObjectTemperatureCallbackPeriod), buf.Bytes())
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

// Sets the thresholds for the RegisterAmbientTemperatureReachedCallback callback.
//
// The following options are possible:
//
//  Option| Description
//  --- | ---
//  'x'|    Callback is turned off
//  'o'|    Callback is triggered when the ambient temperature is *outside* the min and max values
//  'i'|    Callback is triggered when the ambient temperature is *inside* the min and max values
//  '<'|    Callback is triggered when the ambient temperature is smaller than the min value (max is ignored)
//  '>'|    Callback is triggered when the ambient temperature is greater than the min value (max is ignored)
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *TemperatureIRBricklet) SetAmbientTemperatureCallbackThreshold(option ThresholdOption, min int16, max int16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, option)
	binary.Write(&buf, binary.LittleEndian, min)
	binary.Write(&buf, binary.LittleEndian, max)

	resultBytes, err := device.device.Set(uint8(FunctionSetAmbientTemperatureCallbackThreshold), buf.Bytes())
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

// Returns the threshold as set by SetAmbientTemperatureCallbackThreshold.
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *TemperatureIRBricklet) GetAmbientTemperatureCallbackThreshold() (option ThresholdOption, min int16, max int16, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetAmbientTemperatureCallbackThreshold), buf.Bytes())
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

// Sets the thresholds for the RegisterObjectTemperatureReachedCallback callback.
//
// The following options are possible:
//
//  Option| Description
//  --- | ---
//  'x'|    Callback is turned off
//  'o'|    Callback is triggered when the object temperature is *outside* the min and max values
//  'i'|    Callback is triggered when the object temperature is *inside* the min and max values
//  '<'|    Callback is triggered when the object temperature is smaller than the min value (max is ignored)
//  '>'|    Callback is triggered when the object temperature is greater than the min value (max is ignored)
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *TemperatureIRBricklet) SetObjectTemperatureCallbackThreshold(option ThresholdOption, min int16, max int16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, option)
	binary.Write(&buf, binary.LittleEndian, min)
	binary.Write(&buf, binary.LittleEndian, max)

	resultBytes, err := device.device.Set(uint8(FunctionSetObjectTemperatureCallbackThreshold), buf.Bytes())
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

// Returns the threshold as set by SetObjectTemperatureCallbackThreshold.
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *TemperatureIRBricklet) GetObjectTemperatureCallbackThreshold() (option ThresholdOption, min int16, max int16, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetObjectTemperatureCallbackThreshold), buf.Bytes())
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
// * RegisterAmbientTemperatureReachedCallback,
// * RegisterObjectTemperatureReachedCallback
//
// are triggered, if the thresholds
//
// * SetAmbientTemperatureCallbackThreshold,
// * SetObjectTemperatureCallbackThreshold
//
// keep being reached.
func (device *TemperatureIRBricklet) SetDebouncePeriod(debounce uint32) (err error) {
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
func (device *TemperatureIRBricklet) GetDebouncePeriod() (debounce uint32, err error) {
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
func (device *TemperatureIRBricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
