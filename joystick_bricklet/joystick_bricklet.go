/* ***********************************************************
 * This file was automatically generated on 2022-05-11.      *
 *                                                           *
 * Go Bindings Version 2.0.12                                *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


// 2-axis joystick with push-button.
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/Joystick_Bricklet_Go.html.
package joystick_bricklet

import (
	"encoding/binary"
	"bytes"
	"fmt"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionGetPosition Function = 1
	FunctionIsPressed Function = 2
	FunctionGetAnalogValue Function = 3
	FunctionCalibrate Function = 4
	FunctionSetPositionCallbackPeriod Function = 5
	FunctionGetPositionCallbackPeriod Function = 6
	FunctionSetAnalogValueCallbackPeriod Function = 7
	FunctionGetAnalogValueCallbackPeriod Function = 8
	FunctionSetPositionCallbackThreshold Function = 9
	FunctionGetPositionCallbackThreshold Function = 10
	FunctionSetAnalogValueCallbackThreshold Function = 11
	FunctionGetAnalogValueCallbackThreshold Function = 12
	FunctionSetDebouncePeriod Function = 13
	FunctionGetDebouncePeriod Function = 14
	FunctionGetIdentity Function = 255
	FunctionCallbackPosition Function = 15
	FunctionCallbackAnalogValue Function = 16
	FunctionCallbackPositionReached Function = 17
	FunctionCallbackAnalogValueReached Function = 18
	FunctionCallbackPressed Function = 19
	FunctionCallbackReleased Function = 20
)

type ThresholdOption = rune

const (
	ThresholdOptionOff ThresholdOption = 'x'
	ThresholdOptionOutside ThresholdOption = 'o'
	ThresholdOptionInside ThresholdOption = 'i'
	ThresholdOptionSmaller ThresholdOption = '<'
	ThresholdOptionGreater ThresholdOption = '>'
)

type JoystickBricklet struct {
	device Device
}
const DeviceIdentifier = 210
const DeviceDisplayName = "Joystick Bricklet"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (JoystickBricklet, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{ 2,0,0 }, uid, &internalIPCon, 0, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return JoystickBricklet{}, err
	}
	dev.ResponseExpected[FunctionGetPosition] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionIsPressed] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetAnalogValue] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionCalibrate] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionSetPositionCallbackPeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetPositionCallbackPeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetAnalogValueCallbackPeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetAnalogValueCallbackPeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetPositionCallbackThreshold] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetPositionCallbackThreshold] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetAnalogValueCallbackThreshold] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetAnalogValueCallbackThreshold] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetDebouncePeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetDebouncePeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetIdentity] = ResponseExpectedFlagAlwaysTrue;
	return JoystickBricklet{dev}, nil
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
func (device *JoystickBricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *JoystickBricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *JoystickBricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *JoystickBricklet) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// This callback is triggered periodically with the period that is set by
// SetPositionCallbackPeriod. The parameter is the position of the
// joystick.
// 
// The RegisterPositionCallback callback is only triggered if the position has changed since the
// last triggering.
func (device *JoystickBricklet) RegisterPositionCallback(fn func(int16, int16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 12 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var x int16
		var y int16
		binary.Read(buf, binary.LittleEndian, &x)
		binary.Read(buf, binary.LittleEndian, &y)
		fn(x, y)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackPosition), wrapper)
}

// Remove a registered Position callback.
func (device *JoystickBricklet) DeregisterPositionCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackPosition), registrationId)
}


// This callback is triggered periodically with the period that is set by
// SetAnalogValueCallbackPeriod. The parameters are the
// analog values of the joystick.
// 
// The RegisterAnalogValueCallback callback is only triggered if the values have changed
// since the last triggering.
func (device *JoystickBricklet) RegisterAnalogValueCallback(fn func(uint16, uint16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 12 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var x uint16
		var y uint16
		binary.Read(buf, binary.LittleEndian, &x)
		binary.Read(buf, binary.LittleEndian, &y)
		fn(x, y)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackAnalogValue), wrapper)
}

// Remove a registered Analog Value callback.
func (device *JoystickBricklet) DeregisterAnalogValueCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackAnalogValue), registrationId)
}


// This callback is triggered when the threshold as set by
// SetPositionCallbackThreshold is reached.
// The parameters are the position of the joystick.
// 
// If the threshold keeps being reached, the callback is triggered periodically
// with the period as set by SetDebouncePeriod.
func (device *JoystickBricklet) RegisterPositionReachedCallback(fn func(int16, int16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 12 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var x int16
		var y int16
		binary.Read(buf, binary.LittleEndian, &x)
		binary.Read(buf, binary.LittleEndian, &y)
		fn(x, y)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackPositionReached), wrapper)
}

// Remove a registered Position Reached callback.
func (device *JoystickBricklet) DeregisterPositionReachedCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackPositionReached), registrationId)
}


// This callback is triggered when the threshold as set by
// SetAnalogValueCallbackThreshold is reached.
// The parameters are the analog values of the joystick.
// 
// If the threshold keeps being reached, the callback is triggered periodically
// with the period as set by SetDebouncePeriod.
func (device *JoystickBricklet) RegisterAnalogValueReachedCallback(fn func(uint16, uint16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 12 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var x uint16
		var y uint16
		binary.Read(buf, binary.LittleEndian, &x)
		binary.Read(buf, binary.LittleEndian, &y)
		fn(x, y)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackAnalogValueReached), wrapper)
}

// Remove a registered Analog Value Reached callback.
func (device *JoystickBricklet) DeregisterAnalogValueReachedCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackAnalogValueReached), registrationId)
}


// This callback is triggered when the button is pressed.
func (device *JoystickBricklet) RegisterPressedCallback(fn func()) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 8 {
			return
		}
		
		
		
		fn()
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackPressed), wrapper)
}

// Remove a registered Pressed callback.
func (device *JoystickBricklet) DeregisterPressedCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackPressed), registrationId)
}


// This callback is triggered when the button is released.
func (device *JoystickBricklet) RegisterReleasedCallback(fn func()) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 8 {
			return
		}
		
		
		
		fn()
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackReleased), wrapper)
}

// Remove a registered Released callback.
func (device *JoystickBricklet) DeregisterReleasedCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackReleased), registrationId)
}


// Returns the position of the joystick. The middle position of the joystick is x=0, y=0.
// The returned values are averaged and calibrated (see Calibrate).
// 
// If you want to get the position periodically, it is recommended to use the
// RegisterPositionCallback callback and set the period with
// SetPositionCallbackPeriod.
func (device *JoystickBricklet) GetPosition() (x int16, y int16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetPosition), buf.Bytes())
	if err != nil {
		return x, y, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 12 {
			return x, y, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		if header.ErrorCode != 0 {
			return x, y, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &x)
		binary.Read(resultBuf, binary.LittleEndian, &y)

	}

	return x, y, nil
}

// Returns *true* if the button is pressed and *false* otherwise.
// 
// It is recommended to use the RegisterPressedCallback and RegisterReleasedCallback callbacks
// to handle the button.
func (device *JoystickBricklet) IsPressed() (pressed bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionIsPressed), buf.Bytes())
	if err != nil {
		return pressed, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return pressed, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return pressed, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &pressed)

	}

	return pressed, nil
}

// Returns the values as read by a 12-bit analog-to-digital converter.
// 
// Note
//  The values returned by GetPosition are averaged over several samples
//  to yield less noise, while GetAnalogValue gives back raw
//  unfiltered analog values. The only reason to use GetAnalogValue is,
//  if you need the full resolution of the analog-to-digital converter.
// 
// If you want the analog values periodically, it is recommended to use the
// RegisterAnalogValueCallback callback and set the period with
// SetAnalogValueCallbackPeriod.
func (device *JoystickBricklet) GetAnalogValue() (x uint16, y uint16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetAnalogValue), buf.Bytes())
	if err != nil {
		return x, y, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 12 {
			return x, y, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		if header.ErrorCode != 0 {
			return x, y, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &x)
		binary.Read(resultBuf, binary.LittleEndian, &y)

	}

	return x, y, nil
}

// Calibrates the middle position of the joystick. If your Joystick Bricklet
// does not return x=0 and y=0 in the middle position, call this function
// while the joystick is standing still in the middle position.
// 
// The resulting calibration will be saved on the EEPROM of the Joystick
// Bricklet, thus you only have to calibrate it once.
func (device *JoystickBricklet) Calibrate() (err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Set(uint8(FunctionCalibrate), buf.Bytes())
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

// Sets the period with which the RegisterPositionCallback callback is triggered
// periodically. A value of 0 turns the callback off.
// 
// The RegisterPositionCallback callback is only triggered if the position has changed since the
// last triggering.
func (device *JoystickBricklet) SetPositionCallbackPeriod(period uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);

	resultBytes, err := device.device.Set(uint8(FunctionSetPositionCallbackPeriod), buf.Bytes())
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

// Returns the period as set by SetPositionCallbackPeriod.
func (device *JoystickBricklet) GetPositionCallbackPeriod() (period uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetPositionCallbackPeriod), buf.Bytes())
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

// Sets the period with which the RegisterAnalogValueCallback callback is triggered
// periodically. A value of 0 turns the callback off.
// 
// The RegisterAnalogValueCallback callback is only triggered if the analog values have
// changed since the last triggering.
func (device *JoystickBricklet) SetAnalogValueCallbackPeriod(period uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);

	resultBytes, err := device.device.Set(uint8(FunctionSetAnalogValueCallbackPeriod), buf.Bytes())
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

// Returns the period as set by SetAnalogValueCallbackPeriod.
func (device *JoystickBricklet) GetAnalogValueCallbackPeriod() (period uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetAnalogValueCallbackPeriod), buf.Bytes())
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

// Sets the thresholds for the RegisterPositionReachedCallback callback.
// 
// The following options are possible:
// 
//  Option| Description
//  --- | --- 
//  'x'|    Callback is turned off
//  'o'|    Callback is triggered when the position is *outside* the min and max values
//  'i'|    Callback is triggered when the position is *inside* the min and max values
//  '<'|    Callback is triggered when the position is smaller than the min values (max is ignored)
//  '>'|    Callback is triggered when the position is greater than the min values (max is ignored)
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *JoystickBricklet) SetPositionCallbackThreshold(option ThresholdOption, minX int16, maxX int16, minY int16, maxY int16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, option);
	binary.Write(&buf, binary.LittleEndian, minX);
	binary.Write(&buf, binary.LittleEndian, maxX);
	binary.Write(&buf, binary.LittleEndian, minY);
	binary.Write(&buf, binary.LittleEndian, maxY);

	resultBytes, err := device.device.Set(uint8(FunctionSetPositionCallbackThreshold), buf.Bytes())
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

// Returns the threshold as set by SetPositionCallbackThreshold.
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *JoystickBricklet) GetPositionCallbackThreshold() (option ThresholdOption, minX int16, maxX int16, minY int16, maxY int16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetPositionCallbackThreshold), buf.Bytes())
	if err != nil {
		return option, minX, maxX, minY, maxY, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 17 {
			return option, minX, maxX, minY, maxY, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 17)
		}

		if header.ErrorCode != 0 {
			return option, minX, maxX, minY, maxY, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &option)
		binary.Read(resultBuf, binary.LittleEndian, &minX)
		binary.Read(resultBuf, binary.LittleEndian, &maxX)
		binary.Read(resultBuf, binary.LittleEndian, &minY)
		binary.Read(resultBuf, binary.LittleEndian, &maxY)

	}

	return option, minX, maxX, minY, maxY, nil
}

// Sets the thresholds for the RegisterAnalogValueReachedCallback callback.
// 
// The following options are possible:
// 
//  Option| Description
//  --- | --- 
//  'x'|    Callback is turned off
//  'o'|    Callback is triggered when the analog values are *outside* the min and max values
//  'i'|    Callback is triggered when the analog values are *inside* the min and max values
//  '<'|    Callback is triggered when the analog values are smaller than the min values (max is ignored)
//  '>'|    Callback is triggered when the analog values are greater than the min values (max is ignored)
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *JoystickBricklet) SetAnalogValueCallbackThreshold(option ThresholdOption, minX uint16, maxX uint16, minY uint16, maxY uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, option);
	binary.Write(&buf, binary.LittleEndian, minX);
	binary.Write(&buf, binary.LittleEndian, maxX);
	binary.Write(&buf, binary.LittleEndian, minY);
	binary.Write(&buf, binary.LittleEndian, maxY);

	resultBytes, err := device.device.Set(uint8(FunctionSetAnalogValueCallbackThreshold), buf.Bytes())
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

// Returns the threshold as set by SetAnalogValueCallbackThreshold.
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *JoystickBricklet) GetAnalogValueCallbackThreshold() (option ThresholdOption, minX uint16, maxX uint16, minY uint16, maxY uint16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetAnalogValueCallbackThreshold), buf.Bytes())
	if err != nil {
		return option, minX, maxX, minY, maxY, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 17 {
			return option, minX, maxX, minY, maxY, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 17)
		}

		if header.ErrorCode != 0 {
			return option, minX, maxX, minY, maxY, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &option)
		binary.Read(resultBuf, binary.LittleEndian, &minX)
		binary.Read(resultBuf, binary.LittleEndian, &maxX)
		binary.Read(resultBuf, binary.LittleEndian, &minY)
		binary.Read(resultBuf, binary.LittleEndian, &maxY)

	}

	return option, minX, maxX, minY, maxY, nil
}

// Sets the period with which the threshold callbacks
// 
// * RegisterPositionReachedCallback,
// * RegisterAnalogValueReachedCallback
// 
// are triggered, if the thresholds
// 
// * SetPositionCallbackThreshold,
// * SetAnalogValueCallbackThreshold
// 
// keep being reached.
func (device *JoystickBricklet) SetDebouncePeriod(debounce uint32) (err error) {
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
func (device *JoystickBricklet) GetDebouncePeriod() (debounce uint32, err error) {
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
func (device *JoystickBricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
