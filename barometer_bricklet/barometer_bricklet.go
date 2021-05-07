/* ***********************************************************
 * This file was automatically generated on 2021-05-06.      *
 *                                                           *
 * Go Bindings Version 2.0.11                                *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


// Measures air pressure and altitude changes.
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/Barometer_Bricklet_Go.html.
package barometer_bricklet

import (
	"encoding/binary"
	"bytes"
	"fmt"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionGetAirPressure Function = 1
	FunctionGetAltitude Function = 2
	FunctionSetAirPressureCallbackPeriod Function = 3
	FunctionGetAirPressureCallbackPeriod Function = 4
	FunctionSetAltitudeCallbackPeriod Function = 5
	FunctionGetAltitudeCallbackPeriod Function = 6
	FunctionSetAirPressureCallbackThreshold Function = 7
	FunctionGetAirPressureCallbackThreshold Function = 8
	FunctionSetAltitudeCallbackThreshold Function = 9
	FunctionGetAltitudeCallbackThreshold Function = 10
	FunctionSetDebouncePeriod Function = 11
	FunctionGetDebouncePeriod Function = 12
	FunctionSetReferenceAirPressure Function = 13
	FunctionGetChipTemperature Function = 14
	FunctionGetReferenceAirPressure Function = 19
	FunctionSetAveraging Function = 20
	FunctionGetAveraging Function = 21
	FunctionSetI2CMode Function = 22
	FunctionGetI2CMode Function = 23
	FunctionGetIdentity Function = 255
	FunctionCallbackAirPressure Function = 15
	FunctionCallbackAltitude Function = 16
	FunctionCallbackAirPressureReached Function = 17
	FunctionCallbackAltitudeReached Function = 18
)

type ThresholdOption = rune

const (
	ThresholdOptionOff ThresholdOption = 'x'
	ThresholdOptionOutside ThresholdOption = 'o'
	ThresholdOptionInside ThresholdOption = 'i'
	ThresholdOptionSmaller ThresholdOption = '<'
	ThresholdOptionGreater ThresholdOption = '>'
)

type I2CMode = uint8

const (
	I2CModeFast I2CMode = 0
	I2CModeSlow I2CMode = 1
)

type BarometerBricklet struct {
	device Device
}
const DeviceIdentifier = 221
const DeviceDisplayName = "Barometer Bricklet"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (BarometerBricklet, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{ 2,0,2 }, uid, &internalIPCon, 0, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return BarometerBricklet{}, err
	}
	dev.ResponseExpected[FunctionGetAirPressure] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetAltitude] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetAirPressureCallbackPeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetAirPressureCallbackPeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetAltitudeCallbackPeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetAltitudeCallbackPeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetAirPressureCallbackThreshold] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetAirPressureCallbackThreshold] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetAltitudeCallbackThreshold] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetAltitudeCallbackThreshold] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetDebouncePeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetDebouncePeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetReferenceAirPressure] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetChipTemperature] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetReferenceAirPressure] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetAveraging] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetAveraging] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetI2CMode] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetI2CMode] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetIdentity] = ResponseExpectedFlagAlwaysTrue;
	return BarometerBricklet{dev}, nil
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
func (device *BarometerBricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *BarometerBricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *BarometerBricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *BarometerBricklet) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// This callback is triggered periodically with the period that is set by
// SetAirPressureCallbackPeriod. The parameter is the air
// pressure of the air pressure sensor.
// 
// The RegisterAirPressureCallback callback is only triggered if the air pressure has
// changed since the last triggering.
func (device *BarometerBricklet) RegisterAirPressureCallback(fn func(int32)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 12 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var airPressure int32
		binary.Read(buf, binary.LittleEndian, &airPressure)
		fn(airPressure)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackAirPressure), wrapper)
}

// Remove a registered Air Pressure callback.
func (device *BarometerBricklet) DeregisterAirPressureCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackAirPressure), registrationId)
}


// This callback is triggered periodically with the period that is set by
// SetAltitudeCallbackPeriod. The parameter is the altitude of
// the air pressure sensor.
// 
// The RegisterAltitudeCallback callback is only triggered if the altitude has changed since
// the last triggering.
func (device *BarometerBricklet) RegisterAltitudeCallback(fn func(int32)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 12 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var altitude int32
		binary.Read(buf, binary.LittleEndian, &altitude)
		fn(altitude)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackAltitude), wrapper)
}

// Remove a registered Altitude callback.
func (device *BarometerBricklet) DeregisterAltitudeCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackAltitude), registrationId)
}


// This callback is triggered when the threshold as set by
// SetAirPressureCallbackThreshold is reached.
// The parameter is the air pressure of the air pressure sensor.
// 
// If the threshold keeps being reached, the callback is triggered periodically
// with the period as set by SetDebouncePeriod.
func (device *BarometerBricklet) RegisterAirPressureReachedCallback(fn func(int32)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 12 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var airPressure int32
		binary.Read(buf, binary.LittleEndian, &airPressure)
		fn(airPressure)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackAirPressureReached), wrapper)
}

// Remove a registered Air Pressure Reached callback.
func (device *BarometerBricklet) DeregisterAirPressureReachedCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackAirPressureReached), registrationId)
}


// This callback is triggered when the threshold as set by
// SetAltitudeCallbackThreshold is reached.
// The parameter is the altitude of the air pressure sensor.
// 
// If the threshold keeps being reached, the callback is triggered periodically
// with the period as set by SetDebouncePeriod.
func (device *BarometerBricklet) RegisterAltitudeReachedCallback(fn func(int32)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 12 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var altitude int32
		binary.Read(buf, binary.LittleEndian, &altitude)
		fn(altitude)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackAltitudeReached), wrapper)
}

// Remove a registered Altitude Reached callback.
func (device *BarometerBricklet) DeregisterAltitudeReachedCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackAltitudeReached), registrationId)
}


// Returns the air pressure of the air pressure sensor.
// 
// If you want to get the air pressure periodically, it is recommended to use the
// RegisterAirPressureCallback callback and set the period with
// SetAirPressureCallbackPeriod.
func (device *BarometerBricklet) GetAirPressure() (airPressure int32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetAirPressure), buf.Bytes())
	if err != nil {
		return airPressure, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 12 {
			return airPressure, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		if header.ErrorCode != 0 {
			return airPressure, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &airPressure)

	}

	return airPressure, nil
}

// Returns the relative altitude of the air pressure sensor. The value is
// calculated based on the difference between the current air pressure
// and the reference air pressure that can be set with SetReferenceAirPressure.
// 
// If you want to get the altitude periodically, it is recommended to use the
// RegisterAltitudeCallback callback and set the period with
// SetAltitudeCallbackPeriod.
func (device *BarometerBricklet) GetAltitude() (altitude int32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetAltitude), buf.Bytes())
	if err != nil {
		return altitude, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 12 {
			return altitude, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		if header.ErrorCode != 0 {
			return altitude, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &altitude)

	}

	return altitude, nil
}

// Sets the period with which the RegisterAirPressureCallback callback is triggered
// periodically. A value of 0 turns the callback off.
// 
// The RegisterAirPressureCallback callback is only triggered if the air pressure has
// changed since the last triggering.
func (device *BarometerBricklet) SetAirPressureCallbackPeriod(period uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);

	resultBytes, err := device.device.Set(uint8(FunctionSetAirPressureCallbackPeriod), buf.Bytes())
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

// Returns the period as set by SetAirPressureCallbackPeriod.
func (device *BarometerBricklet) GetAirPressureCallbackPeriod() (period uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetAirPressureCallbackPeriod), buf.Bytes())
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

// Sets the period with which the RegisterAltitudeCallback callback is triggered
// periodically. A value of 0 turns the callback off.
// 
// The RegisterAltitudeCallback callback is only triggered if the altitude has changed since
// the last triggering.
func (device *BarometerBricklet) SetAltitudeCallbackPeriod(period uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);

	resultBytes, err := device.device.Set(uint8(FunctionSetAltitudeCallbackPeriod), buf.Bytes())
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

// Returns the period as set by SetAltitudeCallbackPeriod.
func (device *BarometerBricklet) GetAltitudeCallbackPeriod() (period uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetAltitudeCallbackPeriod), buf.Bytes())
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

// Sets the thresholds for the RegisterAirPressureReachedCallback callback.
// 
// The following options are possible:
// 
//  Option| Description
//  --- | --- 
//  'x'|    Callback is turned off
//  'o'|    Callback is triggered when the air pressure is *outside* the min and max values
//  'i'|    Callback is triggered when the air pressure is *inside* the min and max values
//  '<'|    Callback is triggered when the air pressure is smaller than the min value (max is ignored)
//  '>'|    Callback is triggered when the air pressure is greater than the min value (max is ignored)
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *BarometerBricklet) SetAirPressureCallbackThreshold(option ThresholdOption, min int32, max int32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, option);
	binary.Write(&buf, binary.LittleEndian, min);
	binary.Write(&buf, binary.LittleEndian, max);

	resultBytes, err := device.device.Set(uint8(FunctionSetAirPressureCallbackThreshold), buf.Bytes())
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

// Returns the threshold as set by SetAirPressureCallbackThreshold.
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *BarometerBricklet) GetAirPressureCallbackThreshold() (option ThresholdOption, min int32, max int32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetAirPressureCallbackThreshold), buf.Bytes())
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

// Sets the thresholds for the RegisterAltitudeReachedCallback callback.
// 
// The following options are possible:
// 
//  Option| Description
//  --- | --- 
//  'x'|    Callback is turned off
//  'o'|    Callback is triggered when the altitude is *outside* the min and max values
//  'i'|    Callback is triggered when the altitude is *inside* the min and max values
//  '<'|    Callback is triggered when the altitude is smaller than the min value (max is ignored)
//  '>'|    Callback is triggered when the altitude is greater than the min value (max is ignored)
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *BarometerBricklet) SetAltitudeCallbackThreshold(option ThresholdOption, min int32, max int32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, option);
	binary.Write(&buf, binary.LittleEndian, min);
	binary.Write(&buf, binary.LittleEndian, max);

	resultBytes, err := device.device.Set(uint8(FunctionSetAltitudeCallbackThreshold), buf.Bytes())
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

// Returns the threshold as set by SetAltitudeCallbackThreshold.
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *BarometerBricklet) GetAltitudeCallbackThreshold() (option ThresholdOption, min int32, max int32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetAltitudeCallbackThreshold), buf.Bytes())
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

// Sets the period with which the threshold callbacks
// 
// * RegisterAirPressureReachedCallback,
// * RegisterAltitudeReachedCallback
// 
// are triggered, if the thresholds
// 
// * SetAirPressureCallbackThreshold,
// * SetAltitudeCallbackThreshold
// 
// keep being reached.
func (device *BarometerBricklet) SetDebouncePeriod(debounce uint32) (err error) {
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
func (device *BarometerBricklet) GetDebouncePeriod() (debounce uint32, err error) {
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

// Sets the reference air pressure for the altitude calculation.
// Setting the reference to the current air pressure results in a calculated
// altitude of 0cm. Passing 0 is a shortcut for passing the current air pressure as
// reference.
// 
// Well known reference values are the Q codes
// https://en.wikipedia.org/wiki/QNH and
// https://en.wikipedia.org/wiki/Mean_sea_level_pressure#Mean_sea_level_pressure
// used in aviation.
func (device *BarometerBricklet) SetReferenceAirPressure(airPressure int32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, airPressure);

	resultBytes, err := device.device.Set(uint8(FunctionSetReferenceAirPressure), buf.Bytes())
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

// Returns the temperature of the air pressure sensor.
// 
// This temperature is used internally for temperature compensation of the air
// pressure measurement. It is not as accurate as the temperature measured by the
// `temperature_bricklet` or the `temperature_ir_bricklet`.
func (device *BarometerBricklet) GetChipTemperature() (temperature int16, err error) {
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

// Returns the reference air pressure as set by SetReferenceAirPressure.
func (device *BarometerBricklet) GetReferenceAirPressure() (airPressure int32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetReferenceAirPressure), buf.Bytes())
	if err != nil {
		return airPressure, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 12 {
			return airPressure, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		if header.ErrorCode != 0 {
			return airPressure, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &airPressure)

	}

	return airPressure, nil
}

// Sets the different averaging parameters. It is possible to set
// the length of a normal averaging for the temperature and pressure,
// as well as an additional length of a
// https://en.wikipedia.org/wiki/Moving_average
// for the pressure. The moving average is calculated from the normal
// averages.  There is no moving average for the temperature.
// 
// Setting the all three parameters to 0 will turn the averaging
// completely off. If the averaging is off, there is lots of noise
// on the data, but the data is without delay. Thus we recommend
// to turn the averaging off if the Barometer Bricklet data is
// to be used for sensor fusion with other sensors.
// 
// .. versionadded:: 2.0.1$nbsp;(Plugin)
func (device *BarometerBricklet) SetAveraging(movingAveragePressure uint8, averagePressure uint8, averageTemperature uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, movingAveragePressure);
	binary.Write(&buf, binary.LittleEndian, averagePressure);
	binary.Write(&buf, binary.LittleEndian, averageTemperature);

	resultBytes, err := device.device.Set(uint8(FunctionSetAveraging), buf.Bytes())
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

// Returns the averaging configuration as set by SetAveraging.
// 
// .. versionadded:: 2.0.1$nbsp;(Plugin)
func (device *BarometerBricklet) GetAveraging() (movingAveragePressure uint8, averagePressure uint8, averageTemperature uint8, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetAveraging), buf.Bytes())
	if err != nil {
		return movingAveragePressure, averagePressure, averageTemperature, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 11 {
			return movingAveragePressure, averagePressure, averageTemperature, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 11)
		}

		if header.ErrorCode != 0 {
			return movingAveragePressure, averagePressure, averageTemperature, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &movingAveragePressure)
		binary.Read(resultBuf, binary.LittleEndian, &averagePressure)
		binary.Read(resultBuf, binary.LittleEndian, &averageTemperature)

	}

	return movingAveragePressure, averagePressure, averageTemperature, nil
}

// Sets the I2C mode. Possible modes are:
// 
// * 0: Fast (400kHz)
// * 1: Slow (100kHz)
// 
// If you have problems with obvious outliers in the
// Barometer Bricklet measurements, they may be caused by EMI issues.
// In this case it may be helpful to lower the I2C speed.
// 
// It is however not recommended to lower the I2C speed in applications where
// a high throughput needs to be achieved.
// 
// .. versionadded:: 2.0.3$nbsp;(Plugin)
//
// Associated constants:
//
//	* I2CModeFast
//	* I2CModeSlow
func (device *BarometerBricklet) SetI2CMode(mode I2CMode) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, mode);

	resultBytes, err := device.device.Set(uint8(FunctionSetI2CMode), buf.Bytes())
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

// Returns the I2C mode as set by SetI2CMode.
// 
// .. versionadded:: 2.0.3$nbsp;(Plugin)
//
// Associated constants:
//
//	* I2CModeFast
//	* I2CModeSlow
func (device *BarometerBricklet) GetI2CMode() (mode I2CMode, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetI2CMode), buf.Bytes())
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
func (device *BarometerBricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
