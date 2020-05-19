/* ***********************************************************
 * This file was automatically generated on 2020-05-19.      *
 *                                                           *
 * Go Bindings Version 2.0.8                                 *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


// 4 galvanically isolated solid state relays.
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/IndustrialQuadRelay_Bricklet_Go.html.
package industrial_quad_relay_bricklet

import (
	"encoding/binary"
	"bytes"
	"fmt"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionSetValue Function = 1
	FunctionGetValue Function = 2
	FunctionSetMonoflop Function = 3
	FunctionGetMonoflop Function = 4
	FunctionSetGroup Function = 5
	FunctionGetGroup Function = 6
	FunctionGetAvailableForGroup Function = 7
	FunctionSetSelectedValues Function = 9
	FunctionGetIdentity Function = 255
	FunctionCallbackMonoflopDone Function = 8
)

type IndustrialQuadRelayBricklet struct {
	device Device
}
const DeviceIdentifier = 225
const DeviceDisplayName = "Industrial Quad Relay Bricklet"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (IndustrialQuadRelayBricklet, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{ 2,0,0 }, uid, &internalIPCon, 0, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return IndustrialQuadRelayBricklet{}, err
	}
	dev.ResponseExpected[FunctionSetValue] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetValue] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetMonoflop] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetMonoflop] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetGroup] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetGroup] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetAvailableForGroup] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetSelectedValues] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetIdentity] = ResponseExpectedFlagAlwaysTrue;
	return IndustrialQuadRelayBricklet{dev}, nil
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
func (device *IndustrialQuadRelayBricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *IndustrialQuadRelayBricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *IndustrialQuadRelayBricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *IndustrialQuadRelayBricklet) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// This callback is triggered whenever a monoflop timer reaches 0. The
// parameters contain the involved pins and the current value of the pins
// (the value after the monoflop).
func (device *IndustrialQuadRelayBricklet) RegisterMonoflopDoneCallback(fn func(uint16, uint16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 12 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var selectionMask uint16
		var valueMask uint16
		binary.Read(buf, binary.LittleEndian, &selectionMask)
		binary.Read(buf, binary.LittleEndian, &valueMask)
		fn(selectionMask, valueMask)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackMonoflopDone), wrapper)
}

// Remove a registered Monoflop Done callback.
func (device *IndustrialQuadRelayBricklet) DeregisterMonoflopDoneCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackMonoflopDone), registrationId)
}


// Sets the output value with a bitmask (16bit). A 1 in the bitmask means relay
// closed and a 0 means relay open.
// 
// For example: The value 3 or 0b0011 will close the relay of pins 0-1 and open
// the other pins.
// 
// If no groups are used (see SetGroup), the pins correspond to the
// markings on the Industrial Quad Relay Bricklet.
// 
// If groups are used, the pins correspond to the element in the group.
// Element 1 in the group will get pins 0-3, element 2 pins 4-7, element 3
// pins 8-11 and element 4 pins 12-15.
// 
// All running monoflop timers will be aborted if this function is called.
func (device *IndustrialQuadRelayBricklet) SetValue(valueMask uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, valueMask);

	resultBytes, err := device.device.Set(uint8(FunctionSetValue), buf.Bytes())
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

// Returns the bitmask as set by SetValue.
func (device *IndustrialQuadRelayBricklet) GetValue() (valueMask uint16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetValue), buf.Bytes())
	if err != nil {
		return valueMask, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 10 {
			return valueMask, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		if header.ErrorCode != 0 {
			return valueMask, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &valueMask)

	}

	return valueMask, nil
}

// Configures a monoflop of the pins specified by the first parameter
// bitmask.
// 
// The second parameter is a bitmask with the desired value of the specified
// pins. A 1 in the bitmask means relay closed and a 0 means relay open.
// 
// The third parameter indicates the time that the pins should hold
// the value.
// 
// If this function is called with the parameters (9, 1, 1500) or
// (0b1001, 0b0001, 1500): Pin 0 will close and pin 3 will open. In 1.5s pin 0
// will open and pin 3 will close again.
// 
// A monoflop can be used as a fail-safe mechanism. For example: Lets assume you
// have a RS485 bus and a Quad Relay Bricklet connected to one of the slave
// stacks. You can now call this function every second, with a time parameter
// of two seconds and pin 0 closed. Pin 0 will be closed all the time. If now
// the RS485 connection is lost, then pin 0 will be opened in at most two seconds.
func (device *IndustrialQuadRelayBricklet) SetMonoflop(selectionMask uint16, valueMask uint16, time uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, selectionMask);
	binary.Write(&buf, binary.LittleEndian, valueMask);
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

// Returns (for the given pin) the current value and the time as set by
// SetMonoflop as well as the remaining time until the value flips.
// 
// If the timer is not running currently, the remaining time will be returned
// as 0.
func (device *IndustrialQuadRelayBricklet) GetMonoflop(pin uint8) (value uint16, time uint32, timeRemaining uint32, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, pin);

	resultBytes, err := device.device.Get(uint8(FunctionGetMonoflop), buf.Bytes())
	if err != nil {
		return value, time, timeRemaining, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 18 {
			return value, time, timeRemaining, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 18)
		}

		if header.ErrorCode != 0 {
			return value, time, timeRemaining, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &value)
		binary.Read(resultBuf, binary.LittleEndian, &time)
		binary.Read(resultBuf, binary.LittleEndian, &timeRemaining)

	}

	return value, time, timeRemaining, nil
}

// Sets a group of Quad Relay Bricklets that should work together. You can
// find Bricklets that can be grouped together with GetAvailableForGroup.
// 
// The group consists of 4 elements. Element 1 in the group will get pins 0-3,
// element 2 pins 4-7, element 3 pins 8-11 and element 4 pins 12-15.
// 
// Each element can either be one of the ports ('a' to 'd') or 'n' if it should
// not be used.
// 
// For example: If you have two Quad Relay Bricklets connected to port A and
// port B respectively, you could call with ``['a', 'b', 'n', 'n']``.
// 
// Now the pins on the Quad Relay on port A are assigned to 0-3 and the
// pins on the Quad Relay on port B are assigned to 4-7. It is now possible
// to call SetValue and control two Bricklets at the same time.
func (device *IndustrialQuadRelayBricklet) SetGroup(group [4]rune) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, group);

	resultBytes, err := device.device.Set(uint8(FunctionSetGroup), buf.Bytes())
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

// Returns the group as set by SetGroup
func (device *IndustrialQuadRelayBricklet) GetGroup() (group [4]rune, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetGroup), buf.Bytes())
	if err != nil {
		return group, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 12 {
			return group, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		if header.ErrorCode != 0 {
			return group, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		copy(group[:], ByteSliceToRuneSlice(resultBuf.Next(4)))

	}

	return group, nil
}

// Returns a bitmask of ports that are available for grouping. For example the
// value 5 or 0b0101 means: Port A and port C are connected to Bricklets that
// can be grouped together.
func (device *IndustrialQuadRelayBricklet) GetAvailableForGroup() (available uint8, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetAvailableForGroup), buf.Bytes())
	if err != nil {
		return available, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return available, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return available, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &available)

	}

	return available, nil
}

// Sets the output value with a bitmask, according to the selection mask.
// The bitmask is 16 bit long, *true* refers to a closed relay and
// *false* refers to an open relay.
// 
// For example: The values (3, 1) or (0b0011, 0b0001) will close the relay of
// pin 0, open the relay of pin 1 and leave the others untouched.
// 
// If no groups are used (see SetGroup), the pins correspond to the
// markings on the Industrial Quad Relay Bricklet.
// 
// If groups are used, the pins correspond to the element in the group.
// Element 1 in the group will get pins 0-3, element 2 pins 4-7, element 3
// pins 8-11 and element 4 pins 12-15.
// 
// Running monoflop timers for the selected relays will be aborted if this function
// is called.
func (device *IndustrialQuadRelayBricklet) SetSelectedValues(selectionMask uint16, valueMask uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, selectionMask);
	binary.Write(&buf, binary.LittleEndian, valueMask);

	resultBytes, err := device.device.Set(uint8(FunctionSetSelectedValues), buf.Bytes())
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
// The Raspberry Pi HAT (Zero) Brick is always at position 'i' and the Bricklet
// connected to an `Isolator Bricklet <isolator_bricklet>` is always as
// position 'z'.
// 
// The device identifier numbers can be found `here <device_identifier>`.
// |device_identifier_constant|
func (device *IndustrialQuadRelayBricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
