/* ***********************************************************
 * This file was automatically generated on 2024-02-27.      *
 *                                                           *
 * Go Bindings Version 2.0.15                                *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/

// Four 7-segment displays with switchable colon.
//
//
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/SegmentDisplay4x7_Bricklet_Go.html.
package segment_display_4x7_bricklet

import (
	"bytes"
	"encoding/binary"
	"fmt"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionSetSegments             Function = 1
	FunctionGetSegments             Function = 2
	FunctionStartCounter            Function = 3
	FunctionGetCounterValue         Function = 4
	FunctionGetIdentity             Function = 255
	FunctionCallbackCounterFinished Function = 5
)

type SegmentDisplay4x7Bricklet struct {
	device Device
}

const DeviceIdentifier = 237
const DeviceDisplayName = "Segment Display 4x7 Bricklet"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (SegmentDisplay4x7Bricklet, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{2, 0, 0}, uid, &internalIPCon, 0, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return SegmentDisplay4x7Bricklet{}, err
	}
	dev.ResponseExpected[FunctionSetSegments] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetSegments] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionStartCounter] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetCounterValue] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetIdentity] = ResponseExpectedFlagAlwaysTrue
	return SegmentDisplay4x7Bricklet{dev}, nil
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
func (device *SegmentDisplay4x7Bricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *SegmentDisplay4x7Bricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *SegmentDisplay4x7Bricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *SegmentDisplay4x7Bricklet) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// This callback is triggered when the counter (see StartCounter) is
// finished.
func (device *SegmentDisplay4x7Bricklet) RegisterCounterFinishedCallback(fn func()) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 8 {
			return
		}

		fn()
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackCounterFinished), wrapper)
}

// Remove a registered Counter Finished callback.
func (device *SegmentDisplay4x7Bricklet) DeregisterCounterFinishedCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackCounterFinished), registrationId)
}

// The 7-segment display can be set with bitmaps. Every bit controls one
// segment:
//
// .. image:: /Images/Bricklets/bricklet_segment_display_4x7_bit_order.png
//    :scale: 100 %
//    :alt: Bit order of one segment
//    :align: center
//
// For example to set a 5 you would want to activate segments 0, 2, 3, 5 and 6.
// This is represented by the number 0b01101101 = 0x6d = 109.
//
// The brightness can be set between 0 (dark) and 7 (bright). The colon
// parameter turns the colon of the display on or off.
func (device *SegmentDisplay4x7Bricklet) SetSegments(segments [4]uint8, brightness uint8, colon bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, segments)
	binary.Write(&buf, binary.LittleEndian, brightness)
	binary.Write(&buf, binary.LittleEndian, colon)

	resultBytes, err := device.device.Set(uint8(FunctionSetSegments), buf.Bytes())
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

// Returns the segment, brightness and color data as set by
// SetSegments.
func (device *SegmentDisplay4x7Bricklet) GetSegments() (segments [4]uint8, brightness uint8, colon bool, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetSegments), buf.Bytes())
	if err != nil {
		return segments, brightness, colon, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return segments, brightness, colon, DeviceError(header.ErrorCode)
		}

		if header.Length != 14 {
			return segments, brightness, colon, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 14)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &segments)
		binary.Read(resultBuf, binary.LittleEndian, &brightness)
		binary.Read(resultBuf, binary.LittleEndian, &colon)

	}

	return segments, brightness, colon, nil
}

// Starts a counter with the *from* value that counts to the *to*
// value with the each step incremented by *increment*.
// *length* is the pause between each increment.
//
// Example: If you set *from* to 0, *to* to 100, *increment* to 1 and
// *length* to 1000, a counter that goes from 0 to 100 with one second
// pause between each increment will be started.
//
// Using a negative increment allows to count backwards.
//
// You can stop the counter at every time by calling SetSegments.
func (device *SegmentDisplay4x7Bricklet) StartCounter(valueFrom int16, valueTo int16, increment int16, length uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, valueFrom)
	binary.Write(&buf, binary.LittleEndian, valueTo)
	binary.Write(&buf, binary.LittleEndian, increment)
	binary.Write(&buf, binary.LittleEndian, length)

	resultBytes, err := device.device.Set(uint8(FunctionStartCounter), buf.Bytes())
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

// Returns the counter value that is currently shown on the display.
//
// If there is no counter running a 0 will be returned.
func (device *SegmentDisplay4x7Bricklet) GetCounterValue() (value uint16, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetCounterValue), buf.Bytes())
	if err != nil {
		return value, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return value, DeviceError(header.ErrorCode)
		}

		if header.Length != 10 {
			return value, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &value)

	}

	return value, nil
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
func (device *SegmentDisplay4x7Bricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
