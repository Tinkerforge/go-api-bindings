/* ***********************************************************
 * This file was automatically generated on 2025-10-07.      *
 *                                                           *
 * Go Bindings Version 2.0.17                                *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/

// Controls AC and DC Solid State Relays.
//
//
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/SolidStateRelay_Bricklet_Go.html.
package solid_state_relay_bricklet

import (
	"bytes"
	"encoding/binary"
	"fmt"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionSetState             Function = 1
	FunctionGetState             Function = 2
	FunctionSetMonoflop          Function = 3
	FunctionGetMonoflop          Function = 4
	FunctionGetIdentity          Function = 255
	FunctionCallbackMonoflopDone Function = 5
)

type SolidStateRelayBricklet struct {
	device Device
}

const DeviceIdentifier = 244
const DeviceDisplayName = "Solid State Relay Bricklet"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (SolidStateRelayBricklet, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{2, 0, 0}, uid, &internalIPCon, 0, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return SolidStateRelayBricklet{}, err
	}
	dev.ResponseExpected[FunctionSetState] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetState] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetMonoflop] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetMonoflop] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetIdentity] = ResponseExpectedFlagAlwaysTrue
	return SolidStateRelayBricklet{dev}, nil
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
func (device *SolidStateRelayBricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *SolidStateRelayBricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *SolidStateRelayBricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *SolidStateRelayBricklet) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// This callback is triggered whenever the monoflop timer reaches 0.
// The parameter is the current state of the relay
// (the state after the monoflop).
func (device *SolidStateRelayBricklet) RegisterMonoflopDoneCallback(fn func(bool)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 9 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var state bool
		binary.Read(buf, binary.LittleEndian, &state)
		fn(state)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackMonoflopDone), wrapper)
}

// Remove a registered Monoflop Done callback.
func (device *SolidStateRelayBricklet) DeregisterMonoflopDoneCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackMonoflopDone), registrationId)
}

// Sets the state of the relays *true* means on and *false* means off.
//
// A running monoflop timer will be aborted if this function is called.
func (device *SolidStateRelayBricklet) SetState(state bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, state)

	resultBytes, err := device.device.Set(uint8(FunctionSetState), buf.Bytes())
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

// Returns the state of the relay, *true* means on and *false* means off.
func (device *SolidStateRelayBricklet) GetState() (state bool, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetState), buf.Bytes())
	if err != nil {
		return state, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return state, DeviceError(header.ErrorCode)
		}

		if header.Length != 9 {
			return state, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &state)

	}

	return state, nil
}

// The first parameter  is the desired state of the relay (*true* means on
// and *false* means off). The second parameter indicates the time that
// the relay should hold the state.
//
// If this function is called with the parameters (true, 1500):
// The relay will turn on and in 1.5s it will turn off again.
//
// A monoflop can be used as a failsafe mechanism. For example: Lets assume you
// have a RS485 bus and a Solid State Relay Bricklet connected to one of the slave
// stacks. You can now call this function every second, with a time parameter
// of two seconds. The relay will be on all the time. If now the RS485
// connection is lost, the relay will turn off in at most two seconds.
func (device *SolidStateRelayBricklet) SetMonoflop(state bool, time uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, state)
	binary.Write(&buf, binary.LittleEndian, time)

	resultBytes, err := device.device.Set(uint8(FunctionSetMonoflop), buf.Bytes())
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

// Returns the current state and the time as set by
// SetMonoflop as well as the remaining time until the state flips.
//
// If the timer is not running currently, the remaining time will be returned
// as 0.
func (device *SolidStateRelayBricklet) GetMonoflop() (state bool, time uint32, timeRemaining uint32, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetMonoflop), buf.Bytes())
	if err != nil {
		return state, time, timeRemaining, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return state, time, timeRemaining, DeviceError(header.ErrorCode)
		}

		if header.Length != 17 {
			return state, time, timeRemaining, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 17)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &state)
		binary.Read(resultBuf, binary.LittleEndian, &time)
		binary.Read(resultBuf, binary.LittleEndian, &timeRemaining)

	}

	return state, time, timeRemaining, nil
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
func (device *SolidStateRelayBricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
