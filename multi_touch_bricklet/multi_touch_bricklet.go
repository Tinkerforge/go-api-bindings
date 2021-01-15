/* ***********************************************************
 * This file was automatically generated on 2021-01-15.      *
 *                                                           *
 * Go Bindings Version 2.0.10                                *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


// Capacitive touch sensor for 12 electrodes.
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/MultiTouch_Bricklet_Go.html.
package multi_touch_bricklet

import (
	"encoding/binary"
	"bytes"
	"fmt"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionGetTouchState Function = 1
	FunctionRecalibrate Function = 2
	FunctionSetElectrodeConfig Function = 3
	FunctionGetElectrodeConfig Function = 4
	FunctionSetElectrodeSensitivity Function = 6
	FunctionGetElectrodeSensitivity Function = 7
	FunctionGetIdentity Function = 255
	FunctionCallbackTouchState Function = 5
)

type MultiTouchBricklet struct {
	device Device
}
const DeviceIdentifier = 234
const DeviceDisplayName = "Multi Touch Bricklet"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (MultiTouchBricklet, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{ 2,0,0 }, uid, &internalIPCon, 0, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return MultiTouchBricklet{}, err
	}
	dev.ResponseExpected[FunctionGetTouchState] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionRecalibrate] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionSetElectrodeConfig] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetElectrodeConfig] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetElectrodeSensitivity] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetElectrodeSensitivity] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetIdentity] = ResponseExpectedFlagAlwaysTrue;
	return MultiTouchBricklet{dev}, nil
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
func (device *MultiTouchBricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *MultiTouchBricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *MultiTouchBricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *MultiTouchBricklet) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// Returns the current touch state, see GetTouchState for
// information about the state.
// 
// This callback is triggered every time the touch state changes.
func (device *MultiTouchBricklet) RegisterTouchStateCallback(fn func(uint16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 10 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var state uint16
		binary.Read(buf, binary.LittleEndian, &state)
		fn(state)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackTouchState), wrapper)
}

// Remove a registered Touch State callback.
func (device *MultiTouchBricklet) DeregisterTouchStateCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackTouchState), registrationId)
}


// Returns the current touch state. The state is given as a bitfield.
// 
// Bits 0 to 11 represent the 12 electrodes and bit 12 represents
// the proximity.
// 
// If an electrode is touched, the corresponding bit is *true*. If
// a hand or similar is in proximity to the electrodes, bit 12 is
// *true*.
// 
// Example: The state 4103 = 0x1007 = 0b1000000000111 means that
// electrodes 0, 1 and 2 are touched and that something is in the
// proximity of the electrodes.
// 
// The proximity is activated with a distance of 1-2cm. An electrode
// is already counted as touched if a finger is nearly touching the
// electrode. This means that you can put a piece of paper or foil
// or similar on top of a electrode to build a touch panel with
// a professional look.
func (device *MultiTouchBricklet) GetTouchState() (state uint16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetTouchState), buf.Bytes())
	if err != nil {
		return state, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 10 {
			return state, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		if header.ErrorCode != 0 {
			return state, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &state)

	}

	return state, nil
}

// Recalibrates the electrodes. Call this function whenever you changed
// or moved you electrodes.
func (device *MultiTouchBricklet) Recalibrate() (err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Set(uint8(FunctionRecalibrate), buf.Bytes())
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

// Enables/disables electrodes with a bitfield (see GetTouchState).
// 
// *True* enables the electrode, *false* disables the electrode. A
// disabled electrode will always return *false* as its state. If you
// don't need all electrodes you can disable the electrodes that are
// not needed.
// 
// It is recommended that you disable the proximity bit (bit 12) if
// the proximity feature is not needed. This will reduce the amount of
// traffic that is produced by the RegisterTouchStateCallback callback.
// 
// Disabling electrodes will also reduce power consumption.
// 
// Default: 8191 = 0x1FFF = 0b1111111111111 (all electrodes and proximity feature enabled)
func (device *MultiTouchBricklet) SetElectrodeConfig(enabledElectrodes uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, enabledElectrodes);

	resultBytes, err := device.device.Set(uint8(FunctionSetElectrodeConfig), buf.Bytes())
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

// Returns the electrode configuration, as set by SetElectrodeConfig.
func (device *MultiTouchBricklet) GetElectrodeConfig() (enabledElectrodes uint16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetElectrodeConfig), buf.Bytes())
	if err != nil {
		return enabledElectrodes, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 10 {
			return enabledElectrodes, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		if header.ErrorCode != 0 {
			return enabledElectrodes, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &enabledElectrodes)

	}

	return enabledElectrodes, nil
}

// Sets the sensitivity of the electrodes. An electrode with a high sensitivity
// will register a touch earlier then an electrode with a low sensitivity.
// 
// If you build a big electrode you might need to decrease the sensitivity, since
// the area that can be charged will get bigger. If you want to be able to
// activate an electrode from further away you need to increase the sensitivity.
// 
// After a new sensitivity is set, you likely want to call Recalibrate
// to calibrate the electrodes with the newly defined sensitivity.
func (device *MultiTouchBricklet) SetElectrodeSensitivity(sensitivity uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, sensitivity);

	resultBytes, err := device.device.Set(uint8(FunctionSetElectrodeSensitivity), buf.Bytes())
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

// Returns the current sensitivity, as set by SetElectrodeSensitivity.
func (device *MultiTouchBricklet) GetElectrodeSensitivity() (sensitivity uint8, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetElectrodeSensitivity), buf.Bytes())
	if err != nil {
		return sensitivity, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return sensitivity, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return sensitivity, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &sensitivity)

	}

	return sensitivity, nil
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
func (device *MultiTouchBricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
