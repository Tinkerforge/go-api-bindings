/* ***********************************************************
 * This file was automatically generated on 2025-10-07.      *
 *                                                           *
 * Go Bindings Version 2.0.17                                *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/

// Measures distance up to 40m with laser light.
//
//
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/LaserRangeFinder_Bricklet_Go.html.
package laser_range_finder_bricklet

import (
	"bytes"
	"encoding/binary"
	"fmt"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionGetDistance                  Function = 1
	FunctionGetVelocity                  Function = 2
	FunctionSetDistanceCallbackPeriod    Function = 3
	FunctionGetDistanceCallbackPeriod    Function = 4
	FunctionSetVelocityCallbackPeriod    Function = 5
	FunctionGetVelocityCallbackPeriod    Function = 6
	FunctionSetDistanceCallbackThreshold Function = 7
	FunctionGetDistanceCallbackThreshold Function = 8
	FunctionSetVelocityCallbackThreshold Function = 9
	FunctionGetVelocityCallbackThreshold Function = 10
	FunctionSetDebouncePeriod            Function = 11
	FunctionGetDebouncePeriod            Function = 12
	FunctionSetMovingAverage             Function = 13
	FunctionGetMovingAverage             Function = 14
	FunctionSetMode                      Function = 15
	FunctionGetMode                      Function = 16
	FunctionEnableLaser                  Function = 17
	FunctionDisableLaser                 Function = 18
	FunctionIsLaserEnabled               Function = 19
	FunctionGetSensorHardwareVersion     Function = 24
	FunctionSetConfiguration             Function = 25
	FunctionGetConfiguration             Function = 26
	FunctionGetIdentity                  Function = 255
	FunctionCallbackDistance             Function = 20
	FunctionCallbackVelocity             Function = 21
	FunctionCallbackDistanceReached      Function = 22
	FunctionCallbackVelocityReached      Function = 23
)

type ThresholdOption = rune

const (
	ThresholdOptionOff     ThresholdOption = 'x'
	ThresholdOptionOutside ThresholdOption = 'o'
	ThresholdOptionInside  ThresholdOption = 'i'
	ThresholdOptionSmaller ThresholdOption = '<'
	ThresholdOptionGreater ThresholdOption = '>'
)

type Mode = uint8

const (
	ModeDistance         Mode = 0
	ModeVelocityMax13ms  Mode = 1
	ModeVelocityMax32ms  Mode = 2
	ModeVelocityMax64ms  Mode = 3
	ModeVelocityMax127ms Mode = 4
)

type Version = uint8

const (
	Version1 Version = 1
	Version3 Version = 3
)

type LaserRangeFinderBricklet struct {
	device Device
}

const DeviceIdentifier = 255
const DeviceDisplayName = "Laser Range Finder Bricklet"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (LaserRangeFinderBricklet, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{2, 0, 1}, uid, &internalIPCon, 0, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return LaserRangeFinderBricklet{}, err
	}
	dev.ResponseExpected[FunctionGetDistance] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetVelocity] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetDistanceCallbackPeriod] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionGetDistanceCallbackPeriod] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetVelocityCallbackPeriod] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionGetVelocityCallbackPeriod] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetDistanceCallbackThreshold] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionGetDistanceCallbackThreshold] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetVelocityCallbackThreshold] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionGetVelocityCallbackThreshold] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetDebouncePeriod] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionGetDebouncePeriod] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetMovingAverage] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetMovingAverage] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetMode] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetMode] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionEnableLaser] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionDisableLaser] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionIsLaserEnabled] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetSensorHardwareVersion] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetConfiguration] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetConfiguration] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetIdentity] = ResponseExpectedFlagAlwaysTrue
	return LaserRangeFinderBricklet{dev}, nil
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
func (device *LaserRangeFinderBricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *LaserRangeFinderBricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *LaserRangeFinderBricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *LaserRangeFinderBricklet) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// This callback is triggered periodically with the period that is set by
// SetDistanceCallbackPeriod. The parameter is the distance
// value of the sensor.
//
// The RegisterDistanceCallback callback is only triggered if the distance value has changed
// since the last triggering.
func (device *LaserRangeFinderBricklet) RegisterDistanceCallback(fn func(uint16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 10 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var distance uint16
		binary.Read(buf, binary.LittleEndian, &distance)
		fn(distance)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackDistance), wrapper)
}

// Remove a registered Distance callback.
func (device *LaserRangeFinderBricklet) DeregisterDistanceCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackDistance), registrationId)
}

// This callback is triggered periodically with the period that is set by
// SetVelocityCallbackPeriod. The parameter is the velocity
// value of the sensor.
//
// The RegisterVelocityCallback callback is only triggered if the velocity has changed since
// the last triggering.
func (device *LaserRangeFinderBricklet) RegisterVelocityCallback(fn func(int16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 10 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var velocity int16
		binary.Read(buf, binary.LittleEndian, &velocity)
		fn(velocity)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackVelocity), wrapper)
}

// Remove a registered Velocity callback.
func (device *LaserRangeFinderBricklet) DeregisterVelocityCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackVelocity), registrationId)
}

// This callback is triggered when the threshold as set by
// SetDistanceCallbackThreshold is reached.
// The parameter is the distance value of the sensor.
//
// If the threshold keeps being reached, the callback is triggered periodically
// with the period as set by SetDebouncePeriod.
func (device *LaserRangeFinderBricklet) RegisterDistanceReachedCallback(fn func(uint16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 10 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var distance uint16
		binary.Read(buf, binary.LittleEndian, &distance)
		fn(distance)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackDistanceReached), wrapper)
}

// Remove a registered Distance Reached callback.
func (device *LaserRangeFinderBricklet) DeregisterDistanceReachedCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackDistanceReached), registrationId)
}

// This callback is triggered when the threshold as set by
// SetVelocityCallbackThreshold is reached.
// The parameter is the velocity value of the sensor.
//
// If the threshold keeps being reached, the callback is triggered periodically
// with the period as set by SetDebouncePeriod.
func (device *LaserRangeFinderBricklet) RegisterVelocityReachedCallback(fn func(int16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 10 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var velocity int16
		binary.Read(buf, binary.LittleEndian, &velocity)
		fn(velocity)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackVelocityReached), wrapper)
}

// Remove a registered Velocity Reached callback.
func (device *LaserRangeFinderBricklet) DeregisterVelocityReachedCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackVelocityReached), registrationId)
}

// Returns the measured distance.
//
// Sensor hardware version 1 (see GetSensorHardwareVersion) cannot
// measure distance and velocity at the same time. Therefore, the distance mode
// has to be enabled using SetMode.
// Sensor hardware version 3 can measure distance and velocity at the same
// time. Also the laser has to be enabled, see EnableLaser.
//
// If you want to get the distance periodically, it is recommended to
// use the RegisterDistanceCallback callback and set the period with
// SetDistanceCallbackPeriod.
func (device *LaserRangeFinderBricklet) GetDistance() (distance uint16, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetDistance), buf.Bytes())
	if err != nil {
		return distance, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return distance, DeviceError(header.ErrorCode)
		}

		if header.Length != 10 {
			return distance, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &distance)

	}

	return distance, nil
}

// Returns the measured velocity.
//
// Sensor hardware version 1 (see GetSensorHardwareVersion) cannot
// measure distance and velocity at the same time. Therefore, the velocity mode
// has to be enabled using SetMode.
// Sensor hardware version 3 can measure distance and velocity at the same
// time, but the velocity measurement only produces stables results if a fixed
// measurement rate (see SetConfiguration) is configured. Also the laser
// has to be enabled, see EnableLaser.
//
// If you want to get the velocity periodically, it is recommended to
// use the RegisterVelocityCallback callback and set the period with
// SetVelocityCallbackPeriod.
func (device *LaserRangeFinderBricklet) GetVelocity() (velocity int16, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetVelocity), buf.Bytes())
	if err != nil {
		return velocity, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return velocity, DeviceError(header.ErrorCode)
		}

		if header.Length != 10 {
			return velocity, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &velocity)

	}

	return velocity, nil
}

// Sets the period with which the RegisterDistanceCallback callback is triggered
// periodically. A value of 0 turns the callback off.
//
// The RegisterDistanceCallback callback is only triggered if the distance value has
// changed since the last triggering.
func (device *LaserRangeFinderBricklet) SetDistanceCallbackPeriod(period uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period)

	resultBytes, err := device.device.Set(uint8(FunctionSetDistanceCallbackPeriod), buf.Bytes())
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

// Returns the period as set by SetDistanceCallbackPeriod.
func (device *LaserRangeFinderBricklet) GetDistanceCallbackPeriod() (period uint32, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetDistanceCallbackPeriod), buf.Bytes())
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

// Sets the period with which the RegisterVelocityCallback callback is triggered
// periodically. A value of 0 turns the callback off.
//
// The RegisterVelocityCallback callback is only triggered if the velocity value has
// changed since the last triggering.
func (device *LaserRangeFinderBricklet) SetVelocityCallbackPeriod(period uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period)

	resultBytes, err := device.device.Set(uint8(FunctionSetVelocityCallbackPeriod), buf.Bytes())
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

// Returns the period as set by SetVelocityCallbackPeriod.
func (device *LaserRangeFinderBricklet) GetVelocityCallbackPeriod() (period uint32, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetVelocityCallbackPeriod), buf.Bytes())
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

// Sets the thresholds for the RegisterDistanceReachedCallback callback.
//
// The following options are possible:
//
//  Option| Description
//  --- | ---
//  'x'|    Callback is turned off
//  'o'|    Callback is triggered when the distance value is *outside* the min and max values
//  'i'|    Callback is triggered when the distance value is *inside* the min and max values
//  '<'|    Callback is triggered when the distance value is smaller than the min value (max is ignored)
//  '>'|    Callback is triggered when the distance value is greater than the min value (max is ignored)
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *LaserRangeFinderBricklet) SetDistanceCallbackThreshold(option ThresholdOption, min uint16, max uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, option)
	binary.Write(&buf, binary.LittleEndian, min)
	binary.Write(&buf, binary.LittleEndian, max)

	resultBytes, err := device.device.Set(uint8(FunctionSetDistanceCallbackThreshold), buf.Bytes())
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

// Returns the threshold as set by SetDistanceCallbackThreshold.
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *LaserRangeFinderBricklet) GetDistanceCallbackThreshold() (option ThresholdOption, min uint16, max uint16, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetDistanceCallbackThreshold), buf.Bytes())
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

// Sets the thresholds for the RegisterVelocityReachedCallback callback.
//
// The following options are possible:
//
//  Option| Description
//  --- | ---
//  'x'|    Callback is turned off
//  'o'|    Callback is triggered when the velocity is *outside* the min and max values
//  'i'|    Callback is triggered when the velocity is *inside* the min and max values
//  '<'|    Callback is triggered when the velocity is smaller than the min value (max is ignored)
//  '>'|    Callback is triggered when the velocity is greater than the min value (max is ignored)
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *LaserRangeFinderBricklet) SetVelocityCallbackThreshold(option ThresholdOption, min int16, max int16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, option)
	binary.Write(&buf, binary.LittleEndian, min)
	binary.Write(&buf, binary.LittleEndian, max)

	resultBytes, err := device.device.Set(uint8(FunctionSetVelocityCallbackThreshold), buf.Bytes())
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

// Returns the threshold as set by SetVelocityCallbackThreshold.
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *LaserRangeFinderBricklet) GetVelocityCallbackThreshold() (option ThresholdOption, min int16, max int16, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetVelocityCallbackThreshold), buf.Bytes())
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
// * RegisterDistanceReachedCallback,
// * RegisterVelocityReachedCallback,
//
// are triggered, if the thresholds
//
// * SetDistanceCallbackThreshold,
// * SetVelocityCallbackThreshold,
//
// keep being reached.
func (device *LaserRangeFinderBricklet) SetDebouncePeriod(debounce uint32) (err error) {
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
func (device *LaserRangeFinderBricklet) GetDebouncePeriod() (debounce uint32, err error) {
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
// for the distance and velocity.
//
// Setting the length to 0 will turn the averaging completely off. With less
// averaging, there is more noise on the data.
func (device *LaserRangeFinderBricklet) SetMovingAverage(distanceAverageLength uint8, velocityAverageLength uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, distanceAverageLength)
	binary.Write(&buf, binary.LittleEndian, velocityAverageLength)

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

// Returns the length moving average as set by SetMovingAverage.
func (device *LaserRangeFinderBricklet) GetMovingAverage() (distanceAverageLength uint8, velocityAverageLength uint8, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetMovingAverage), buf.Bytes())
	if err != nil {
		return distanceAverageLength, velocityAverageLength, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return distanceAverageLength, velocityAverageLength, DeviceError(header.ErrorCode)
		}

		if header.Length != 10 {
			return distanceAverageLength, velocityAverageLength, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &distanceAverageLength)
		binary.Read(resultBuf, binary.LittleEndian, &velocityAverageLength)

	}

	return distanceAverageLength, velocityAverageLength, nil
}

// Note
//  This function is only available if you have a LIDAR-Lite sensor with hardware
//  version 1. Use SetConfiguration for hardware version 3. You can check
//  the sensor hardware version using GetSensorHardwareVersion.
//
// The LIDAR-Lite sensor (hardware version 1) has five different modes. One mode is
// for distance measurements and four modes are for velocity measurements with
// different ranges.
//
// The following modes are available:
//
// * 0: Distance is measured with resolution 1.0 cm and range 0-4000 cm
// * 1: Velocity is measured with resolution 0.1 m/s and range is 0-12.7 m/s
// * 2: Velocity is measured with resolution 0.25 m/s and range is 0-31.75 m/s
// * 3: Velocity is measured with resolution 0.5 m/s and range is 0-63.5 m/s
// * 4: Velocity is measured with resolution 1.0 m/s and range is 0-127 m/s
//
// Associated constants:
//
//	* ModeDistance
//	* ModeVelocityMax13ms
//	* ModeVelocityMax32ms
//	* ModeVelocityMax64ms
//	* ModeVelocityMax127ms
func (device *LaserRangeFinderBricklet) SetMode(mode Mode) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, mode)

	resultBytes, err := device.device.Set(uint8(FunctionSetMode), buf.Bytes())
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

// Returns the mode as set by SetMode.
//
// Associated constants:
//
//	* ModeDistance
//	* ModeVelocityMax13ms
//	* ModeVelocityMax32ms
//	* ModeVelocityMax64ms
//	* ModeVelocityMax127ms
func (device *LaserRangeFinderBricklet) GetMode() (mode Mode, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetMode), buf.Bytes())
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

// Activates the laser of the LIDAR.
//
// We recommend that you wait 250ms after enabling the laser before
// the first call of GetDistance to ensure stable measurements.
func (device *LaserRangeFinderBricklet) EnableLaser() (err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Set(uint8(FunctionEnableLaser), buf.Bytes())
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

// Deactivates the laser of the LIDAR.
func (device *LaserRangeFinderBricklet) DisableLaser() (err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Set(uint8(FunctionDisableLaser), buf.Bytes())
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

// Returns *true* if the laser is enabled, *false* otherwise.
func (device *LaserRangeFinderBricklet) IsLaserEnabled() (laserEnabled bool, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionIsLaserEnabled), buf.Bytes())
	if err != nil {
		return laserEnabled, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return laserEnabled, DeviceError(header.ErrorCode)
		}

		if header.Length != 9 {
			return laserEnabled, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &laserEnabled)

	}

	return laserEnabled, nil
}

// Returns the LIDAR-Lite hardware version.
//
// .. versionadded:: 2.0.3$nbsp;(Plugin)
//
// Associated constants:
//
//	* Version1
//	* Version3
func (device *LaserRangeFinderBricklet) GetSensorHardwareVersion() (version Version, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetSensorHardwareVersion), buf.Bytes())
	if err != nil {
		return version, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return version, DeviceError(header.ErrorCode)
		}

		if header.Length != 9 {
			return version, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &version)

	}

	return version, nil
}

// Note
//  This function is only available if you have a LIDAR-Lite sensor with hardware
//  version 3. Use SetMode for hardware version 1. You can check
//  the sensor hardware version using GetSensorHardwareVersion.
//
// The **Acquisition Count** defines the number of times the Laser Range Finder Bricklet
// will integrate acquisitions to find a correlation record peak. With a higher count,
// the Bricklet can measure longer distances. With a lower count, the rate increases. The
// allowed values are 1-255.
//
// If you set **Enable Quick Termination** to true, the distance measurement will be terminated
// early if a high peak was already detected. This means that a higher measurement rate can be achieved
// and long distances can be measured at the same time. However, the chance of false-positive
// distance measurements increases.
//
// Normally the distance is calculated with a detection algorithm that uses peak value,
// signal strength and noise. You can however also define a fixed **Threshold Value**.
// Set this to a low value if you want to measure the distance to something that has
// very little reflection (e.g. glass) and set it to a high value if you want to measure
// the distance to something with a very high reflection (e.g. mirror). Set this to 0 to
// use the default algorithm. The other allowed values are 1-255.
//
// Set the **Measurement Frequency** to force a fixed measurement rate. If set to 0,
// the Laser Range Finder Bricklet will use the optimal frequency according to the other
// configurations and the actual measured distance. Since the rate is not fixed in this case,
// the velocity measurement is not stable. For a stable velocity measurement you should
// set a fixed measurement frequency. The lower the frequency, the higher is the resolution
// of the calculated velocity. The allowed values are 10Hz-500Hz (and 0 to turn the fixed
// frequency off).
//
// .. versionadded:: 2.0.3$nbsp;(Plugin)
func (device *LaserRangeFinderBricklet) SetConfiguration(acquisitionCount uint8, enableQuickTermination bool, thresholdValue uint8, measurementFrequency uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, acquisitionCount)
	binary.Write(&buf, binary.LittleEndian, enableQuickTermination)
	binary.Write(&buf, binary.LittleEndian, thresholdValue)
	binary.Write(&buf, binary.LittleEndian, measurementFrequency)

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
// .. versionadded:: 2.0.3$nbsp;(Plugin)
func (device *LaserRangeFinderBricklet) GetConfiguration() (acquisitionCount uint8, enableQuickTermination bool, thresholdValue uint8, measurementFrequency uint16, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetConfiguration), buf.Bytes())
	if err != nil {
		return acquisitionCount, enableQuickTermination, thresholdValue, measurementFrequency, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return acquisitionCount, enableQuickTermination, thresholdValue, measurementFrequency, DeviceError(header.ErrorCode)
		}

		if header.Length != 13 {
			return acquisitionCount, enableQuickTermination, thresholdValue, measurementFrequency, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 13)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &acquisitionCount)
		binary.Read(resultBuf, binary.LittleEndian, &enableQuickTermination)
		binary.Read(resultBuf, binary.LittleEndian, &thresholdValue)
		binary.Read(resultBuf, binary.LittleEndian, &measurementFrequency)

	}

	return acquisitionCount, enableQuickTermination, thresholdValue, measurementFrequency, nil
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
func (device *LaserRangeFinderBricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
