/* ***********************************************************
 * This file was automatically generated on 2020-04-07.      *
 *                                                           *
 * Go Bindings Version 2.0.6                                 *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


// Measures sound intensity.
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/SoundIntensity_Bricklet_Go.html.
package sound_intensity_bricklet

import (
	"encoding/binary"
	"bytes"
	"fmt"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionGetIntensity Function = 1
	FunctionSetIntensityCallbackPeriod Function = 2
	FunctionGetIntensityCallbackPeriod Function = 3
	FunctionSetIntensityCallbackThreshold Function = 4
	FunctionGetIntensityCallbackThreshold Function = 5
	FunctionSetDebouncePeriod Function = 6
	FunctionGetDebouncePeriod Function = 7
	FunctionGetIdentity Function = 255
	FunctionCallbackIntensity Function = 8
	FunctionCallbackIntensityReached Function = 9
)

type ThresholdOption = rune

const (
	ThresholdOptionOff ThresholdOption = 'x'
	ThresholdOptionOutside ThresholdOption = 'o'
	ThresholdOptionInside ThresholdOption = 'i'
	ThresholdOptionSmaller ThresholdOption = '<'
	ThresholdOptionGreater ThresholdOption = '>'
)

type SoundIntensityBricklet struct {
	device Device
}
const DeviceIdentifier = 238
const DeviceDisplayName = "Sound Intensity Bricklet"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (SoundIntensityBricklet, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{ 2,0,0 }, uid, &internalIPCon, 0, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return SoundIntensityBricklet{}, err
	}
	dev.ResponseExpected[FunctionGetIntensity] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetIntensityCallbackPeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetIntensityCallbackPeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetIntensityCallbackThreshold] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetIntensityCallbackThreshold] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetDebouncePeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetDebouncePeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetIdentity] = ResponseExpectedFlagAlwaysTrue;
	return SoundIntensityBricklet{dev}, nil
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
func (device *SoundIntensityBricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *SoundIntensityBricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *SoundIntensityBricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *SoundIntensityBricklet) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// This callback is triggered periodically with the period that is set by
// SetIntensityCallbackPeriod. The parameter is the intensity
// of the sensor.
// 
// The RegisterIntensityCallback callback is only triggered if the intensity has changed
// since the last triggering.
func (device *SoundIntensityBricklet) RegisterIntensityCallback(fn func(uint16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 10 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var intensity uint16
		binary.Read(buf, binary.LittleEndian, &intensity)
		fn(intensity)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackIntensity), wrapper)
}

// Remove a registered Intensity callback.
func (device *SoundIntensityBricklet) DeregisterIntensityCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackIntensity), registrationId)
}


// This callback is triggered when the threshold as set by
// SetIntensityCallbackThreshold is reached.
// The parameter is the intensity of the encoder.
// 
// If the threshold keeps being reached, the callback is triggered periodically
// with the period as set by SetDebouncePeriod.
func (device *SoundIntensityBricklet) RegisterIntensityReachedCallback(fn func(uint16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 10 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var intensity uint16
		binary.Read(buf, binary.LittleEndian, &intensity)
		fn(intensity)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackIntensityReached), wrapper)
}

// Remove a registered Intensity Reached callback.
func (device *SoundIntensityBricklet) DeregisterIntensityReachedCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackIntensityReached), registrationId)
}


// Returns the current sound intensity.
// 
// The value corresponds to the
// https://en.wikipedia.org/wiki/Envelope_(waves)
// of the signal of the microphone capsule.
// 
// If you want to get the intensity periodically, it is recommended to use the
// RegisterIntensityCallback callback and set the period with
// SetIntensityCallbackPeriod.
func (device *SoundIntensityBricklet) GetIntensity() (intensity uint16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetIntensity), buf.Bytes())
	if err != nil {
		return intensity, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 10 {
		return intensity, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
	}


	if header.ErrorCode != 0 {
		return intensity, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &intensity)


	return intensity, nil
}

// Sets the period with which the RegisterIntensityCallback callback is triggered
// periodically. A value of 0 turns the callback off.
// 
// The RegisterIntensityCallback callback is only triggered if the intensity has changed
// since the last triggering.
func (device *SoundIntensityBricklet) SetIntensityCallbackPeriod(period uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);

	resultBytes, err := device.device.Set(uint8(FunctionSetIntensityCallbackPeriod), buf.Bytes())
	if err != nil {
		return err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 8 {
		return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
	}


	if header.ErrorCode != 0 {
		return DeviceError(header.ErrorCode)
	}

	bytes.NewBuffer(resultBytes[8:])
	

	return nil
}

// Returns the period as set by SetIntensityCallbackPeriod.
func (device *SoundIntensityBricklet) GetIntensityCallbackPeriod() (period uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetIntensityCallbackPeriod), buf.Bytes())
	if err != nil {
		return period, err
	}

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


	return period, nil
}

// Sets the thresholds for the RegisterIntensityReachedCallback callback.
// 
// The following options are possible:
// 
//  Option| Description
//  --- | --- 
//  'x'|    Callback is turned off
//  'o'|    Callback is triggered when the intensity is *outside* the min and max values
//  'i'|    Callback is triggered when the intensity is *inside* the min and max values
//  '<'|    Callback is triggered when the intensity is smaller than the min value (max is ignored)
//  '>'|    Callback is triggered when the intensity is greater than the min value (max is ignored)
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *SoundIntensityBricklet) SetIntensityCallbackThreshold(option ThresholdOption, min uint16, max uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, option);
	binary.Write(&buf, binary.LittleEndian, min);
	binary.Write(&buf, binary.LittleEndian, max);

	resultBytes, err := device.device.Set(uint8(FunctionSetIntensityCallbackThreshold), buf.Bytes())
	if err != nil {
		return err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 8 {
		return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
	}


	if header.ErrorCode != 0 {
		return DeviceError(header.ErrorCode)
	}

	bytes.NewBuffer(resultBytes[8:])
	

	return nil
}

// Returns the threshold as set by SetIntensityCallbackThreshold.
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *SoundIntensityBricklet) GetIntensityCallbackThreshold() (option ThresholdOption, min uint16, max uint16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetIntensityCallbackThreshold), buf.Bytes())
	if err != nil {
		return option, min, max, err
	}

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


	return option, min, max, nil
}

// Sets the period with which the threshold callback
// 
// * RegisterIntensityReachedCallback
// 
// is triggered, if the thresholds
// 
// * SetIntensityCallbackThreshold
// 
// keeps being reached.
func (device *SoundIntensityBricklet) SetDebouncePeriod(debounce uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, debounce);

	resultBytes, err := device.device.Set(uint8(FunctionSetDebouncePeriod), buf.Bytes())
	if err != nil {
		return err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 8 {
		return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
	}


	if header.ErrorCode != 0 {
		return DeviceError(header.ErrorCode)
	}

	bytes.NewBuffer(resultBytes[8:])
	

	return nil
}

// Returns the debounce period as set by SetDebouncePeriod.
func (device *SoundIntensityBricklet) GetDebouncePeriod() (debounce uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetDebouncePeriod), buf.Bytes())
	if err != nil {
		return debounce, err
	}

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


	return debounce, nil
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
func (device *SoundIntensityBricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetIdentity), buf.Bytes())
	if err != nil {
		return uid, connectedUid, position, hardwareVersion, firmwareVersion, deviceIdentifier, err
	}

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


	return uid, connectedUid, position, hardwareVersion, firmwareVersion, deviceIdentifier, nil
}
