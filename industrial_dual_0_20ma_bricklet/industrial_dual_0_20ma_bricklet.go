/* ***********************************************************
 * This file was automatically generated on 2020-04-07.      *
 *                                                           *
 * Go Bindings Version 2.0.6                                 *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


// Measures two DC currents between 0mA and 20mA (IEC 60381-1).
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/IndustrialDual020mA_Bricklet_Go.html.
package industrial_dual_0_20ma_bricklet

import (
	"encoding/binary"
	"bytes"
	"fmt"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionGetCurrent Function = 1
	FunctionSetCurrentCallbackPeriod Function = 2
	FunctionGetCurrentCallbackPeriod Function = 3
	FunctionSetCurrentCallbackThreshold Function = 4
	FunctionGetCurrentCallbackThreshold Function = 5
	FunctionSetDebouncePeriod Function = 6
	FunctionGetDebouncePeriod Function = 7
	FunctionSetSampleRate Function = 8
	FunctionGetSampleRate Function = 9
	FunctionGetIdentity Function = 255
	FunctionCallbackCurrent Function = 10
	FunctionCallbackCurrentReached Function = 11
)

type ThresholdOption = rune

const (
	ThresholdOptionOff ThresholdOption = 'x'
	ThresholdOptionOutside ThresholdOption = 'o'
	ThresholdOptionInside ThresholdOption = 'i'
	ThresholdOptionSmaller ThresholdOption = '<'
	ThresholdOptionGreater ThresholdOption = '>'
)

type SampleRate = uint8

const (
	SampleRate240SPS SampleRate = 0
	SampleRate60SPS SampleRate = 1
	SampleRate15SPS SampleRate = 2
	SampleRate4SPS SampleRate = 3
)

type IndustrialDual020mABricklet struct {
	device Device
}
const DeviceIdentifier = 228
const DeviceDisplayName = "Industrial Dual 0-20mA Bricklet"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (IndustrialDual020mABricklet, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{ 2,0,0 }, uid, &internalIPCon, 0, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return IndustrialDual020mABricklet{}, err
	}
	dev.ResponseExpected[FunctionGetCurrent] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetCurrentCallbackPeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetCurrentCallbackPeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetCurrentCallbackThreshold] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetCurrentCallbackThreshold] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetDebouncePeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetDebouncePeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetSampleRate] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetSampleRate] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetIdentity] = ResponseExpectedFlagAlwaysTrue;
	return IndustrialDual020mABricklet{dev}, nil
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
func (device *IndustrialDual020mABricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *IndustrialDual020mABricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *IndustrialDual020mABricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *IndustrialDual020mABricklet) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// This callback is triggered periodically with the period that is set by
// SetCurrentCallbackPeriod. The parameter is the current of the
// sensor.
// 
// The RegisterCurrentCallback callback is only triggered if the current has changed since the
// last triggering.
func (device *IndustrialDual020mABricklet) RegisterCurrentCallback(fn func(uint8, int32)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 13 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var sensor uint8
		var current int32
		binary.Read(buf, binary.LittleEndian, &sensor)
		binary.Read(buf, binary.LittleEndian, &current)
		fn(sensor, current)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackCurrent), wrapper)
}

// Remove a registered Current callback.
func (device *IndustrialDual020mABricklet) DeregisterCurrentCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackCurrent), registrationId)
}


// This callback is triggered when the threshold as set by
// SetCurrentCallbackThreshold is reached.
// The parameter is the current of the sensor.
// 
// If the threshold keeps being reached, the callback is triggered periodically
// with the period as set by SetDebouncePeriod.
func (device *IndustrialDual020mABricklet) RegisterCurrentReachedCallback(fn func(uint8, int32)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 13 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var sensor uint8
		var current int32
		binary.Read(buf, binary.LittleEndian, &sensor)
		binary.Read(buf, binary.LittleEndian, &current)
		fn(sensor, current)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackCurrentReached), wrapper)
}

// Remove a registered Current Reached callback.
func (device *IndustrialDual020mABricklet) DeregisterCurrentReachedCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackCurrentReached), registrationId)
}


// Returns the current of the specified sensor.
// 
// It is possible to detect if an IEC 60381-1 compatible sensor is connected
// and if it works properly.
// 
// If the returned current is below 4mA, there is likely no sensor connected
// or the sensor may be defect. If the returned current is over 20mA, there might
// be a short circuit or the sensor may be defect.
// 
// If you want to get the current periodically, it is recommended to use the
// RegisterCurrentCallback callback and set the period with
// SetCurrentCallbackPeriod.
func (device *IndustrialDual020mABricklet) GetCurrent(sensor uint8) (current int32, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, sensor);

	resultBytes, err := device.device.Get(uint8(FunctionGetCurrent), buf.Bytes())
	if err != nil {
		return current, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 12 {
		return current, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
	}


	if header.ErrorCode != 0 {
		return current, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &current)


	return current, nil
}

// Sets the period with which the RegisterCurrentCallback callback is triggered
// periodically for the given sensor. A value of 0 turns the callback off.
// 
// The RegisterCurrentCallback callback is only triggered if the current has changed since the
// last triggering.
func (device *IndustrialDual020mABricklet) SetCurrentCallbackPeriod(sensor uint8, period uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, sensor);
	binary.Write(&buf, binary.LittleEndian, period);

	resultBytes, err := device.device.Set(uint8(FunctionSetCurrentCallbackPeriod), buf.Bytes())
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

// Returns the period as set by SetCurrentCallbackPeriod.
func (device *IndustrialDual020mABricklet) GetCurrentCallbackPeriod(sensor uint8) (period uint32, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, sensor);

	resultBytes, err := device.device.Get(uint8(FunctionGetCurrentCallbackPeriod), buf.Bytes())
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

// Sets the thresholds for the RegisterCurrentReachedCallback callback for the given
// sensor.
// 
// The following options are possible:
// 
//  Option| Description
//  --- | --- 
//  'x'|    Callback is turned off
//  'o'|    Callback is triggered when the current is *outside* the min and max values
//  'i'|    Callback is triggered when the current is *inside* the min and max values
//  '<'|    Callback is triggered when the current is smaller than the min value (max is ignored)
//  '>'|    Callback is triggered when the current is greater than the min value (max is ignored)
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *IndustrialDual020mABricklet) SetCurrentCallbackThreshold(sensor uint8, option ThresholdOption, min int32, max int32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, sensor);
	binary.Write(&buf, binary.LittleEndian, option);
	binary.Write(&buf, binary.LittleEndian, min);
	binary.Write(&buf, binary.LittleEndian, max);

	resultBytes, err := device.device.Set(uint8(FunctionSetCurrentCallbackThreshold), buf.Bytes())
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

// Returns the threshold as set by SetCurrentCallbackThreshold.
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *IndustrialDual020mABricklet) GetCurrentCallbackThreshold(sensor uint8) (option ThresholdOption, min int32, max int32, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, sensor);

	resultBytes, err := device.device.Get(uint8(FunctionGetCurrentCallbackThreshold), buf.Bytes())
	if err != nil {
		return option, min, max, err
	}

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


	return option, min, max, nil
}

// Sets the period with which the threshold callback
// 
// * RegisterCurrentReachedCallback
// 
// is triggered, if the threshold
// 
// * SetCurrentCallbackThreshold
// 
// keeps being reached.
func (device *IndustrialDual020mABricklet) SetDebouncePeriod(debounce uint32) (err error) {
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
func (device *IndustrialDual020mABricklet) GetDebouncePeriod() (debounce uint32, err error) {
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

// Sets the sample rate to either 240, 60, 15 or 4 samples per second.
// The resolution for the rates is 12, 14, 16 and 18 bit respectively.
// 
//  Value| Description
//  --- | --- 
//  0|    240 samples per second| 12 bit resolution
//  1|    60 samples per second| 14 bit resolution
//  2|    15 samples per second| 16 bit resolution
//  3|    4 samples per second| 18 bit resolution
//
// Associated constants:
//
//	* SampleRate240SPS
//	* SampleRate60SPS
//	* SampleRate15SPS
//	* SampleRate4SPS
func (device *IndustrialDual020mABricklet) SetSampleRate(rate SampleRate) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, rate);

	resultBytes, err := device.device.Set(uint8(FunctionSetSampleRate), buf.Bytes())
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

// Returns the sample rate as set by SetSampleRate.
//
// Associated constants:
//
//	* SampleRate240SPS
//	* SampleRate60SPS
//	* SampleRate15SPS
//	* SampleRate4SPS
func (device *IndustrialDual020mABricklet) GetSampleRate() (rate SampleRate, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetSampleRate), buf.Bytes())
	if err != nil {
		return rate, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 9 {
		return rate, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
	}


	if header.ErrorCode != 0 {
		return rate, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &rate)


	return rate, nil
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
func (device *IndustrialDual020mABricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
