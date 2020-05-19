/* ***********************************************************
 * This file was automatically generated on 2020-05-19.      *
 *                                                           *
 * Go Bindings Version 2.0.8                                 *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


// Measures dust density.
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/DustDetector_Bricklet_Go.html.
package dust_detector_bricklet

import (
	"encoding/binary"
	"bytes"
	"fmt"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionGetDustDensity Function = 1
	FunctionSetDustDensityCallbackPeriod Function = 2
	FunctionGetDustDensityCallbackPeriod Function = 3
	FunctionSetDustDensityCallbackThreshold Function = 4
	FunctionGetDustDensityCallbackThreshold Function = 5
	FunctionSetDebouncePeriod Function = 6
	FunctionGetDebouncePeriod Function = 7
	FunctionSetMovingAverage Function = 10
	FunctionGetMovingAverage Function = 11
	FunctionGetIdentity Function = 255
	FunctionCallbackDustDensity Function = 8
	FunctionCallbackDustDensityReached Function = 9
)

type ThresholdOption = rune

const (
	ThresholdOptionOff ThresholdOption = 'x'
	ThresholdOptionOutside ThresholdOption = 'o'
	ThresholdOptionInside ThresholdOption = 'i'
	ThresholdOptionSmaller ThresholdOption = '<'
	ThresholdOptionGreater ThresholdOption = '>'
)

type DustDetectorBricklet struct {
	device Device
}
const DeviceIdentifier = 260
const DeviceDisplayName = "Dust Detector Bricklet"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (DustDetectorBricklet, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{ 2,0,0 }, uid, &internalIPCon, 0, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return DustDetectorBricklet{}, err
	}
	dev.ResponseExpected[FunctionGetDustDensity] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetDustDensityCallbackPeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetDustDensityCallbackPeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetDustDensityCallbackThreshold] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetDustDensityCallbackThreshold] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetDebouncePeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetDebouncePeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetMovingAverage] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetMovingAverage] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetIdentity] = ResponseExpectedFlagAlwaysTrue;
	return DustDetectorBricklet{dev}, nil
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
func (device *DustDetectorBricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *DustDetectorBricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *DustDetectorBricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *DustDetectorBricklet) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// This callback is triggered periodically with the period that is set by
// SetDustDensityCallbackPeriod. The parameter is the
// dust density of the sensor.
// 
// Der RegisterDustDensityCallback callback is only triggered if the dust density value has
// changed since the last triggering.
func (device *DustDetectorBricklet) RegisterDustDensityCallback(fn func(uint16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 10 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var dustDensity uint16
		binary.Read(buf, binary.LittleEndian, &dustDensity)
		fn(dustDensity)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackDustDensity), wrapper)
}

// Remove a registered Dust Density callback.
func (device *DustDetectorBricklet) DeregisterDustDensityCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackDustDensity), registrationId)
}


// This callback is triggered when the threshold as set by
// SetDustDensityCallbackThreshold is reached.
// The parameter is the dust density of the sensor.
// 
// If the threshold keeps being reached, the callback is triggered periodically
// with the period as set by SetDebouncePeriod.
func (device *DustDetectorBricklet) RegisterDustDensityReachedCallback(fn func(uint16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 10 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var dustDensity uint16
		binary.Read(buf, binary.LittleEndian, &dustDensity)
		fn(dustDensity)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackDustDensityReached), wrapper)
}

// Remove a registered Dust Density Reached callback.
func (device *DustDetectorBricklet) DeregisterDustDensityReachedCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackDustDensityReached), registrationId)
}


// Returns the dust density.
// 
// If you want to get the dust density periodically, it is recommended
// to use the RegisterDustDensityCallback callback and set the period with
// SetDustDensityCallbackPeriod.
func (device *DustDetectorBricklet) GetDustDensity() (dustDensity uint16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetDustDensity), buf.Bytes())
	if err != nil {
		return dustDensity, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 10 {
			return dustDensity, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		if header.ErrorCode != 0 {
			return dustDensity, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &dustDensity)

	}

	return dustDensity, nil
}

// Sets the period with which the RegisterDustDensityCallback callback is triggered
// periodically. A value of 0 turns the callback off.
// 
// The RegisterDustDensityCallback callback is only triggered if the dust density has
// changed since the last triggering.
func (device *DustDetectorBricklet) SetDustDensityCallbackPeriod(period uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);

	resultBytes, err := device.device.Set(uint8(FunctionSetDustDensityCallbackPeriod), buf.Bytes())
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

// Returns the period as set by SetDustDensityCallbackPeriod.
func (device *DustDetectorBricklet) GetDustDensityCallbackPeriod() (period uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetDustDensityCallbackPeriod), buf.Bytes())
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

// Sets the thresholds for the RegisterDustDensityReachedCallback callback.
// 
// The following options are possible:
// 
//  Option| Description
//  --- | --- 
//  'x'|    Callback is turned off
//  'o'|    Callback is triggered when the dust density value is *outside* the min and max values
//  'i'|    Callback is triggered when the dust density value is *inside* the min and max values
//  '<'|    Callback is triggered when the dust density value is smaller than the min value (max is ignored)
//  '>'|    Callback is triggered when the dust density value is greater than the min value (max is ignored)
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *DustDetectorBricklet) SetDustDensityCallbackThreshold(option ThresholdOption, min uint16, max uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, option);
	binary.Write(&buf, binary.LittleEndian, min);
	binary.Write(&buf, binary.LittleEndian, max);

	resultBytes, err := device.device.Set(uint8(FunctionSetDustDensityCallbackThreshold), buf.Bytes())
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

// Returns the threshold as set by SetDustDensityCallbackThreshold.
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *DustDetectorBricklet) GetDustDensityCallbackThreshold() (option ThresholdOption, min uint16, max uint16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetDustDensityCallbackThreshold), buf.Bytes())
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

// Sets the period with which the threshold callback
// 
// * RegisterDustDensityReachedCallback
// 
// is triggered, if the threshold
// 
// * SetDustDensityCallbackThreshold
// 
// keeps being reached.
func (device *DustDetectorBricklet) SetDebouncePeriod(debounce uint32) (err error) {
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
func (device *DustDetectorBricklet) GetDebouncePeriod() (debounce uint32, err error) {
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

// Sets the length of a https://en.wikipedia.org/wiki/Moving_average
// for the dust density.
// 
// Setting the length to 0 will turn the averaging completely off. With less
// averaging, there is more noise on the data.
func (device *DustDetectorBricklet) SetMovingAverage(average uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, average);

	resultBytes, err := device.device.Set(uint8(FunctionSetMovingAverage), buf.Bytes())
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

// Returns the length moving average as set by SetMovingAverage.
func (device *DustDetectorBricklet) GetMovingAverage() (average uint8, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetMovingAverage), buf.Bytes())
	if err != nil {
		return average, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return average, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return average, DeviceError(header.ErrorCode)
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
// The Raspberry Pi HAT (Zero) Brick is always at position 'i' and the Bricklet
// connected to an `Isolator Bricklet <isolator_bricklet>` is always as
// position 'z'.
// 
// The device identifier numbers can be found `here <device_identifier>`.
// |device_identifier_constant|
func (device *DustDetectorBricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
