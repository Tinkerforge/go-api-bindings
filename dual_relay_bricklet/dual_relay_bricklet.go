/* ***********************************************************
 * This file was automatically generated on 2021-05-06.      *
 *                                                           *
 * Go Bindings Version 2.0.11                                *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


// Two relays to switch AC/DC devices.
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/DualRelay_Bricklet_Go.html.
package dual_relay_bricklet

import (
	"encoding/binary"
	"bytes"
	"fmt"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionSetState Function = 1
	FunctionGetState Function = 2
	FunctionSetMonoflop Function = 3
	FunctionGetMonoflop Function = 4
	FunctionSetSelectedState Function = 6
	FunctionGetIdentity Function = 255
	FunctionCallbackMonoflopDone Function = 5
)

type DualRelayBricklet struct {
	device Device
}
const DeviceIdentifier = 26
const DeviceDisplayName = "Dual Relay Bricklet"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (DualRelayBricklet, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{ 2,0,0 }, uid, &internalIPCon, 0, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return DualRelayBricklet{}, err
	}
	dev.ResponseExpected[FunctionSetState] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetState] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetMonoflop] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetMonoflop] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetSelectedState] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetIdentity] = ResponseExpectedFlagAlwaysTrue;
	return DualRelayBricklet{dev}, nil
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
func (device *DualRelayBricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *DualRelayBricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *DualRelayBricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *DualRelayBricklet) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// This callback is triggered whenever a monoflop timer reaches 0. The
// parameter contain the relay (1 or 2) and the current state of the relay
// (the state after the monoflop).
func (device *DualRelayBricklet) RegisterMonoflopDoneCallback(fn func(uint8, bool)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 10 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var relay uint8
		var state bool
		binary.Read(buf, binary.LittleEndian, &relay)
		binary.Read(buf, binary.LittleEndian, &state)
		fn(relay, state)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackMonoflopDone), wrapper)
}

// Remove a registered Monoflop Done callback.
func (device *DualRelayBricklet) DeregisterMonoflopDoneCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackMonoflopDone), registrationId)
}


// Sets the state of the relays, *true* means on and *false* means off.
// For example: (true, false) turns relay 1 on and relay 2 off.
// 
// If you just want to set one of the relays and don't know the current state
// of the other relay, you can get the state with GetState or you
// can use SetSelectedState.
// 
// All running monoflop timers will be aborted if this function is called.
func (device *DualRelayBricklet) SetState(relay1 bool, relay2 bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, relay1);
	binary.Write(&buf, binary.LittleEndian, relay2);

	resultBytes, err := device.device.Set(uint8(FunctionSetState), buf.Bytes())
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

// Returns the state of the relays, *true* means on and *false* means off.
func (device *DualRelayBricklet) GetState() (relay1 bool, relay2 bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetState), buf.Bytes())
	if err != nil {
		return relay1, relay2, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 10 {
			return relay1, relay2, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		if header.ErrorCode != 0 {
			return relay1, relay2, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &relay1)
		binary.Read(resultBuf, binary.LittleEndian, &relay2)

	}

	return relay1, relay2, nil
}

// The first parameter can be 1 or 2 (relay 1 or relay 2). The second parameter
// is the desired state of the relay (*true* means on and *false* means off).
// The third parameter indicates the time that the relay should hold
// the state.
// 
// If this function is called with the parameters (1, true, 1500):
// Relay 1 will turn on and in 1.5s it will turn off again.
// 
// A monoflop can be used as a failsafe mechanism. For example: Lets assume you
// have a RS485 bus and a Dual Relay Bricklet connected to one of the slave
// stacks. You can now call this function every second, with a time parameter
// of two seconds. The relay will be on all the time. If now the RS485
// connection is lost, the relay will turn off in at most two seconds.
func (device *DualRelayBricklet) SetMonoflop(relay uint8, state bool, time uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, relay);
	binary.Write(&buf, binary.LittleEndian, state);
	binary.Write(&buf, binary.LittleEndian, time);

	resultBytes, err := device.device.Set(uint8(FunctionSetMonoflop), buf.Bytes())
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

// Returns (for the given relay) the current state and the time as set by
// SetMonoflop as well as the remaining time until the state flips.
// 
// If the timer is not running currently, the remaining time will be returned
// as 0.
func (device *DualRelayBricklet) GetMonoflop(relay uint8) (state bool, time uint32, timeRemaining uint32, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, relay);

	resultBytes, err := device.device.Get(uint8(FunctionGetMonoflop), buf.Bytes())
	if err != nil {
		return state, time, timeRemaining, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 17 {
			return state, time, timeRemaining, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 17)
		}

		if header.ErrorCode != 0 {
			return state, time, timeRemaining, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &state)
		binary.Read(resultBuf, binary.LittleEndian, &time)
		binary.Read(resultBuf, binary.LittleEndian, &timeRemaining)

	}

	return state, time, timeRemaining, nil
}

// Sets the state of the selected relay (1 or 2), *true* means on and *false* means off.
// 
// A running monoflop timer for the selected relay will be aborted if this function is called.
// 
// The other relay remains untouched.
func (device *DualRelayBricklet) SetSelectedState(relay uint8, state bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, relay);
	binary.Write(&buf, binary.LittleEndian, state);

	resultBytes, err := device.device.Set(uint8(FunctionSetSelectedState), buf.Bytes())
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
func (device *DualRelayBricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
