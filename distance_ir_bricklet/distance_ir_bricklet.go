/* ***********************************************************
 * This file was automatically generated on 2021-05-06.      *
 *                                                           *
 * Go Bindings Version 2.0.11                                *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


// Measures distance up to 150cm with infrared light.
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/DistanceIR_Bricklet_Go.html.
package distance_ir_bricklet

import (
	"encoding/binary"
	"bytes"
	"fmt"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionGetDistance Function = 1
	FunctionGetAnalogValue Function = 2
	FunctionSetSamplingPoint Function = 3
	FunctionGetSamplingPoint Function = 4
	FunctionSetDistanceCallbackPeriod Function = 5
	FunctionGetDistanceCallbackPeriod Function = 6
	FunctionSetAnalogValueCallbackPeriod Function = 7
	FunctionGetAnalogValueCallbackPeriod Function = 8
	FunctionSetDistanceCallbackThreshold Function = 9
	FunctionGetDistanceCallbackThreshold Function = 10
	FunctionSetAnalogValueCallbackThreshold Function = 11
	FunctionGetAnalogValueCallbackThreshold Function = 12
	FunctionSetDebouncePeriod Function = 13
	FunctionGetDebouncePeriod Function = 14
	FunctionGetIdentity Function = 255
	FunctionCallbackDistance Function = 15
	FunctionCallbackAnalogValue Function = 16
	FunctionCallbackDistanceReached Function = 17
	FunctionCallbackAnalogValueReached Function = 18
)

type ThresholdOption = rune

const (
	ThresholdOptionOff ThresholdOption = 'x'
	ThresholdOptionOutside ThresholdOption = 'o'
	ThresholdOptionInside ThresholdOption = 'i'
	ThresholdOptionSmaller ThresholdOption = '<'
	ThresholdOptionGreater ThresholdOption = '>'
)

type DistanceIRBricklet struct {
	device Device
}
const DeviceIdentifier = 25
const DeviceDisplayName = "Distance IR Bricklet"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (DistanceIRBricklet, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{ 2,0,1 }, uid, &internalIPCon, 0, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return DistanceIRBricklet{}, err
	}
	dev.ResponseExpected[FunctionGetDistance] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetAnalogValue] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetSamplingPoint] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetSamplingPoint] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetDistanceCallbackPeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetDistanceCallbackPeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetAnalogValueCallbackPeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetAnalogValueCallbackPeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetDistanceCallbackThreshold] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetDistanceCallbackThreshold] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetAnalogValueCallbackThreshold] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetAnalogValueCallbackThreshold] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetDebouncePeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetDebouncePeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetIdentity] = ResponseExpectedFlagAlwaysTrue;
	return DistanceIRBricklet{dev}, nil
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
func (device *DistanceIRBricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *DistanceIRBricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *DistanceIRBricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *DistanceIRBricklet) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// This callback is triggered periodically with the period that is set by
// SetDistanceCallbackPeriod. The parameter is the distance of the
// sensor.
// 
// The RegisterDistanceCallback callback is only triggered if the distance has changed since the
// last triggering.
func (device *DistanceIRBricklet) RegisterDistanceCallback(fn func(uint16)) uint64 {
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
func (device *DistanceIRBricklet) DeregisterDistanceCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackDistance), registrationId)
}


// This callback is triggered periodically with the period that is set by
// SetAnalogValueCallbackPeriod. The parameter is the analog value of the
// sensor.
// 
// The RegisterAnalogValueCallback callback is only triggered if the analog value has changed
// since the last triggering.
func (device *DistanceIRBricklet) RegisterAnalogValueCallback(fn func(uint16)) uint64 {
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
func (device *DistanceIRBricklet) DeregisterAnalogValueCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackAnalogValue), registrationId)
}


// This callback is triggered when the threshold as set by
// SetDistanceCallbackThreshold is reached.
// The parameter is the distance of the sensor.
// 
// If the threshold keeps being reached, the callback is triggered periodically
// with the period as set by SetDebouncePeriod.
func (device *DistanceIRBricklet) RegisterDistanceReachedCallback(fn func(uint16)) uint64 {
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
func (device *DistanceIRBricklet) DeregisterDistanceReachedCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackDistanceReached), registrationId)
}


// This callback is triggered when the threshold as set by
// SetAnalogValueCallbackThreshold is reached.
// The parameter is the analog value of the sensor.
// 
// If the threshold keeps being reached, the callback is triggered periodically
// with the period as set by SetDebouncePeriod.
func (device *DistanceIRBricklet) RegisterAnalogValueReachedCallback(fn func(uint16)) uint64 {
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
func (device *DistanceIRBricklet) DeregisterAnalogValueReachedCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackAnalogValueReached), registrationId)
}


// Returns the distance measured by the sensor. Possible
// distance ranges are 40 to 300, 100 to 800 and 200 to 1500, depending on the
// selected IR sensor.
// 
// If you want to get the distance periodically, it is recommended to use the
// RegisterDistanceCallback callback and set the period with
// SetDistanceCallbackPeriod.
func (device *DistanceIRBricklet) GetDistance() (distance uint16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetDistance), buf.Bytes())
	if err != nil {
		return distance, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 10 {
			return distance, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		if header.ErrorCode != 0 {
			return distance, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &distance)

	}

	return distance, nil
}

// Returns the value as read by a 12-bit analog-to-digital converter.
// 
// Note
//  The value returned by GetDistance is averaged over several samples
//  to yield less noise, while GetAnalogValue gives back raw
//  unfiltered analog values. The only reason to use GetAnalogValue is,
//  if you need the full resolution of the analog-to-digital converter.
// 
// If you want the analog value periodically, it is recommended to use the
// RegisterAnalogValueCallback callback and set the period with
// SetAnalogValueCallbackPeriod.
func (device *DistanceIRBricklet) GetAnalogValue() (value uint16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetAnalogValue), buf.Bytes())
	if err != nil {
		return value, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 10 {
			return value, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		if header.ErrorCode != 0 {
			return value, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &value)

	}

	return value, nil
}

// Sets a sampling point value to a specific position of the lookup table.
// The lookup table comprises 128 equidistant analog values with
// corresponding distances.
// 
// If you measure a distance of 50cm at the analog value 2048, you
// should call this function with (64, 5000). The utilized analog-to-digital
// converter has a resolution of 12 bit. With 128 sampling points on the
// whole range, this means that every sampling point has a size of 32
// analog values. Thus the analog value 2048 has the corresponding sampling
// point 64 = 2048/32.
// 
// Sampling points are saved on the EEPROM of the Distance IR Bricklet and
// loaded again on startup.
// 
// Note
//  An easy way to calibrate the sampling points of the Distance IR Bricklet is
//  implemented in the Brick Viewer. If you want to calibrate your Bricklet it is
//  highly recommended to use this implementation.
func (device *DistanceIRBricklet) SetSamplingPoint(position uint8, distance uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, position);
	binary.Write(&buf, binary.LittleEndian, distance);

	resultBytes, err := device.device.Set(uint8(FunctionSetSamplingPoint), buf.Bytes())
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

// Returns the distance to a sampling point position as set by
// SetSamplingPoint.
func (device *DistanceIRBricklet) GetSamplingPoint(position uint8) (distance uint16, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, position);

	resultBytes, err := device.device.Get(uint8(FunctionGetSamplingPoint), buf.Bytes())
	if err != nil {
		return distance, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 10 {
			return distance, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		if header.ErrorCode != 0 {
			return distance, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &distance)

	}

	return distance, nil
}

// Sets the period with which the RegisterDistanceCallback callback is triggered
// periodically. A value of 0 turns the callback off.
// 
// The RegisterDistanceCallback callback is only triggered if the distance has changed since the
// last triggering.
func (device *DistanceIRBricklet) SetDistanceCallbackPeriod(period uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);

	resultBytes, err := device.device.Set(uint8(FunctionSetDistanceCallbackPeriod), buf.Bytes())
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

// Returns the period as set by SetDistanceCallbackPeriod.
func (device *DistanceIRBricklet) GetDistanceCallbackPeriod() (period uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetDistanceCallbackPeriod), buf.Bytes())
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
// The RegisterAnalogValueCallback callback is only triggered if the analog value has
// changed since the last triggering.
func (device *DistanceIRBricklet) SetAnalogValueCallbackPeriod(period uint32) (err error) {
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
func (device *DistanceIRBricklet) GetAnalogValueCallbackPeriod() (period uint32, err error) {
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

// Sets the thresholds for the RegisterDistanceReachedCallback callback.
// 
// The following options are possible:
// 
//  Option| Description
//  --- | --- 
//  'x'|    Callback is turned off
//  'o'|    Callback is triggered when the distance is *outside* the min and max values
//  'i'|    Callback is triggered when the distance is *inside* the min and max values
//  '<'|    Callback is triggered when the distance is smaller than the min value (max is ignored)
//  '>'|    Callback is triggered when the distance is greater than the min value (max is ignored)
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *DistanceIRBricklet) SetDistanceCallbackThreshold(option ThresholdOption, min uint16, max uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, option);
	binary.Write(&buf, binary.LittleEndian, min);
	binary.Write(&buf, binary.LittleEndian, max);

	resultBytes, err := device.device.Set(uint8(FunctionSetDistanceCallbackThreshold), buf.Bytes())
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

// Returns the threshold as set by SetDistanceCallbackThreshold.
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *DistanceIRBricklet) GetDistanceCallbackThreshold() (option ThresholdOption, min uint16, max uint16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetDistanceCallbackThreshold), buf.Bytes())
	if err != nil {
		return option, min, max, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 13 {
			return option, min, max, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 13)
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
func (device *DistanceIRBricklet) SetAnalogValueCallbackThreshold(option ThresholdOption, min uint16, max uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, option);
	binary.Write(&buf, binary.LittleEndian, min);
	binary.Write(&buf, binary.LittleEndian, max);

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
func (device *DistanceIRBricklet) GetAnalogValueCallbackThreshold() (option ThresholdOption, min uint16, max uint16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetAnalogValueCallbackThreshold), buf.Bytes())
	if err != nil {
		return option, min, max, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 13 {
			return option, min, max, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 13)
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
// * RegisterDistanceReachedCallback,
// * RegisterAnalogValueReachedCallback
// 
// are triggered, if the thresholds
// 
// * SetDistanceCallbackThreshold,
// * SetAnalogValueCallbackThreshold
// 
// keep being reached.
func (device *DistanceIRBricklet) SetDebouncePeriod(debounce uint32) (err error) {
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
func (device *DistanceIRBricklet) GetDebouncePeriod() (debounce uint32, err error) {
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
func (device *DistanceIRBricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
