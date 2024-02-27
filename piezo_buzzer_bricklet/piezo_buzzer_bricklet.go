/* ***********************************************************
 * This file was automatically generated on 2024-02-27.      *
 *                                                           *
 * Go Bindings Version 2.0.15                                *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/

// Creates 1kHz beep.
//
//
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/PiezoBuzzer_Bricklet_Go.html.
package piezo_buzzer_bricklet

import (
	"bytes"
	"encoding/binary"
	"fmt"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionBeep                      Function = 1
	FunctionMorseCode                 Function = 2
	FunctionGetIdentity               Function = 255
	FunctionCallbackBeepFinished      Function = 3
	FunctionCallbackMorseCodeFinished Function = 4
)

type PiezoBuzzerBricklet struct {
	device Device
}

const DeviceIdentifier = 214
const DeviceDisplayName = "Piezo Buzzer Bricklet"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (PiezoBuzzerBricklet, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{2, 0, 0}, uid, &internalIPCon, 0, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return PiezoBuzzerBricklet{}, err
	}
	dev.ResponseExpected[FunctionBeep] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionMorseCode] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetIdentity] = ResponseExpectedFlagAlwaysTrue
	return PiezoBuzzerBricklet{dev}, nil
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
func (device *PiezoBuzzerBricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *PiezoBuzzerBricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *PiezoBuzzerBricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *PiezoBuzzerBricklet) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// This callback is triggered if a beep set by Beep is finished
func (device *PiezoBuzzerBricklet) RegisterBeepFinishedCallback(fn func()) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 8 {
			return
		}

		fn()
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackBeepFinished), wrapper)
}

// Remove a registered Beep Finished callback.
func (device *PiezoBuzzerBricklet) DeregisterBeepFinishedCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackBeepFinished), registrationId)
}

// This callback is triggered if the playback of the morse code set by
// MorseCode is finished.
func (device *PiezoBuzzerBricklet) RegisterMorseCodeFinishedCallback(fn func()) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 8 {
			return
		}

		fn()
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackMorseCodeFinished), wrapper)
}

// Remove a registered Morse Code Finished callback.
func (device *PiezoBuzzerBricklet) DeregisterMorseCodeFinishedCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackMorseCodeFinished), registrationId)
}

// Beeps for the given duration.
func (device *PiezoBuzzerBricklet) Beep(duration uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, duration)

	resultBytes, err := device.device.Set(uint8(FunctionBeep), buf.Bytes())
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

// Sets morse code that will be played by the piezo buzzer. The morse code
// is given as a string consisting of . (dot), - (minus) and   (space)
// for *dits*, *dahs* and *pauses*. Every other character is ignored.
//
// For example: If you set the string ...---..., the piezo buzzer will beep
// nine times with the durations short short short long long long short
// short short.
func (device *PiezoBuzzerBricklet) MorseCode(morse string) (err error) {
	var buf bytes.Buffer
	morse_byte_slice, err := StringToByteSlice(morse, 60)
	if err != nil {
		return
	}
	buf.Write(morse_byte_slice)

	resultBytes, err := device.device.Set(uint8(FunctionMorseCode), buf.Bytes())
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
func (device *PiezoBuzzerBricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
