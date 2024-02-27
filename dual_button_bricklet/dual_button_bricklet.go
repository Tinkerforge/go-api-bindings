/* ***********************************************************
 * This file was automatically generated on 2024-02-27.      *
 *                                                           *
 * Go Bindings Version 2.0.15                                *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/

// Two tactile buttons with built-in blue LEDs.
//
//
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/DualButton_Bricklet_Go.html.
package dual_button_bricklet

import (
	"bytes"
	"encoding/binary"
	"fmt"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionSetLEDState          Function = 1
	FunctionGetLEDState          Function = 2
	FunctionGetButtonState       Function = 3
	FunctionSetSelectedLEDState  Function = 5
	FunctionGetIdentity          Function = 255
	FunctionCallbackStateChanged Function = 4
)

type LEDState = uint8

const (
	LEDStateAutoToggleOn  LEDState = 0
	LEDStateAutoToggleOff LEDState = 1
	LEDStateOn            LEDState = 2
	LEDStateOff           LEDState = 3
)

type ButtonState = uint8

const (
	ButtonStatePressed  ButtonState = 0
	ButtonStateReleased ButtonState = 1
)

type LED = uint8

const (
	LEDLeft  LED = 0
	LEDRight LED = 1
)

type DualButtonBricklet struct {
	device Device
}

const DeviceIdentifier = 230
const DeviceDisplayName = "Dual Button Bricklet"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (DualButtonBricklet, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{2, 0, 0}, uid, &internalIPCon, 0, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return DualButtonBricklet{}, err
	}
	dev.ResponseExpected[FunctionSetLEDState] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetLEDState] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetButtonState] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetSelectedLEDState] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetIdentity] = ResponseExpectedFlagAlwaysTrue
	return DualButtonBricklet{dev}, nil
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
func (device *DualButtonBricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *DualButtonBricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *DualButtonBricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *DualButtonBricklet) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// This callback is called whenever a button is pressed.
//
// Possible states for buttons are:
//
// * 0 = pressed
// * 1 = released
//
// Possible states for LEDs are:
//
// * 0 = AutoToggleOn: Auto toggle enabled and LED on.
// * 1 = AutoToggleOff: Auto toggle enabled and LED off.
// * 2 = On: LED on (auto toggle is disabled).
// * 3 = Off: LED off (auto toggle is disabled).
func (device *DualButtonBricklet) RegisterStateChangedCallback(fn func(ButtonState, ButtonState, LEDState, LEDState)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 12 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var buttonL ButtonState
		var buttonR ButtonState
		var ledL LEDState
		var ledR LEDState
		binary.Read(buf, binary.LittleEndian, &buttonL)
		binary.Read(buf, binary.LittleEndian, &buttonR)
		binary.Read(buf, binary.LittleEndian, &ledL)
		binary.Read(buf, binary.LittleEndian, &ledR)
		fn(buttonL, buttonR, ledL, ledR)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackStateChanged), wrapper)
}

// Remove a registered State Changed callback.
func (device *DualButtonBricklet) DeregisterStateChangedCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackStateChanged), registrationId)
}

// Sets the state of the LEDs. Possible states are:
//
// * 0 = AutoToggleOn: Enables auto toggle with initially enabled LED.
// * 1 = AutoToggleOff: Activates auto toggle with initially disabled LED.
// * 2 = On: Enables LED (auto toggle is disabled).
// * 3 = Off: Disables LED (auto toggle is disabled).
//
// In auto toggle mode the LED is toggled automatically at each press of a button.
//
// If you just want to set one of the LEDs and don't know the current state
// of the other LED, you can get the state with GetLEDState or you
// can use SetSelectedLEDState.
//
// Associated constants:
//
//	* LEDStateAutoToggleOn
//	* LEDStateAutoToggleOff
//	* LEDStateOn
//	* LEDStateOff
func (device *DualButtonBricklet) SetLEDState(ledL LEDState, ledR LEDState) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, ledL)
	binary.Write(&buf, binary.LittleEndian, ledR)

	resultBytes, err := device.device.Set(uint8(FunctionSetLEDState), buf.Bytes())
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

// Returns the current state of the LEDs, as set by SetLEDState.
//
// Associated constants:
//
//	* LEDStateAutoToggleOn
//	* LEDStateAutoToggleOff
//	* LEDStateOn
//	* LEDStateOff
func (device *DualButtonBricklet) GetLEDState() (ledL LEDState, ledR LEDState, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetLEDState), buf.Bytes())
	if err != nil {
		return ledL, ledR, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return ledL, ledR, DeviceError(header.ErrorCode)
		}

		if header.Length != 10 {
			return ledL, ledR, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &ledL)
		binary.Read(resultBuf, binary.LittleEndian, &ledR)

	}

	return ledL, ledR, nil
}

// Returns the current state for both buttons. Possible states are:
//
// * 0 = pressed
// * 1 = released
//
// Associated constants:
//
//	* ButtonStatePressed
//	* ButtonStateReleased
func (device *DualButtonBricklet) GetButtonState() (buttonL ButtonState, buttonR ButtonState, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetButtonState), buf.Bytes())
	if err != nil {
		return buttonL, buttonR, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return buttonL, buttonR, DeviceError(header.ErrorCode)
		}

		if header.Length != 10 {
			return buttonL, buttonR, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &buttonL)
		binary.Read(resultBuf, binary.LittleEndian, &buttonR)

	}

	return buttonL, buttonR, nil
}

// Sets the state of the selected LED (0 or 1).
//
// The other LED remains untouched.
//
// Associated constants:
//
//	* LEDLeft
//	* LEDRight
//	* LEDStateAutoToggleOn
//	* LEDStateAutoToggleOff
//	* LEDStateOn
//	* LEDStateOff
func (device *DualButtonBricklet) SetSelectedLEDState(led LED, state LEDState) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, led)
	binary.Write(&buf, binary.LittleEndian, state)

	resultBytes, err := device.device.Set(uint8(FunctionSetSelectedLEDState), buf.Bytes())
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
func (device *DualButtonBricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
