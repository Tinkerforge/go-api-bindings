/* ***********************************************************
 * This file was automatically generated on 2019-08-23.      *
 *                                                           *
 * Go Bindings Version 2.0.4                                 *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


// 4 galvanically isolated digital inputs.
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/IndustrialDigitalIn4_Bricklet_Go.html.
package industrial_digital_in_4_bricklet

import (
	"encoding/binary"
	"bytes"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionGetValue Function = 1
	FunctionSetGroup Function = 2
	FunctionGetGroup Function = 3
	FunctionGetAvailableForGroup Function = 4
	FunctionSetDebouncePeriod Function = 5
	FunctionGetDebouncePeriod Function = 6
	FunctionSetInterrupt Function = 7
	FunctionGetInterrupt Function = 8
	FunctionGetEdgeCount Function = 10
	FunctionSetEdgeCountConfig Function = 11
	FunctionGetEdgeCountConfig Function = 12
	FunctionGetIdentity Function = 255
	FunctionCallbackInterrupt Function = 9
)

type EdgeType = uint8

const (
	EdgeTypeRising EdgeType = 0
	EdgeTypeFalling EdgeType = 1
	EdgeTypeBoth EdgeType = 2
)

type IndustrialDigitalIn4Bricklet struct {
	device Device
}
const DeviceIdentifier = 223
const DeviceDisplayName = "Industrial Digital In 4 Bricklet"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (IndustrialDigitalIn4Bricklet, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{ 2,0,1 }, uid, &internalIPCon, 0)
	if err != nil {
		return IndustrialDigitalIn4Bricklet{}, err
	}
	dev.ResponseExpected[FunctionGetValue] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetGroup] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetGroup] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetAvailableForGroup] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetDebouncePeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetDebouncePeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetInterrupt] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetInterrupt] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetEdgeCount] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetEdgeCountConfig] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetEdgeCountConfig] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetIdentity] = ResponseExpectedFlagAlwaysTrue;
	return IndustrialDigitalIn4Bricklet{dev}, nil
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
// for this purpose. If this flag is disabled for a setter function then no response is send
// and errors are silently ignored, because they cannot be detected.
//
// See SetResponseExpected for the list of function ID constants available for this function.
func (device *IndustrialDigitalIn4Bricklet) GetResponseExpected(functionID Function) (bool, error) {
	return device.device.GetResponseExpected(uint8(functionID))
}

// Changes the response expected flag of the function specified by the function ID parameter.
// This flag can only be changed for setter (default value: false) and callback configuration
// functions (default value: true). For getter functions it is always enabled.
//
// Enabling the response expected flag for a setter function allows to detect timeouts and
// other error conditions calls of this setter as well. The device will then send a response
// for this purpose. If this flag is disabled for a setter function then no response is send
// and errors are silently ignored, because they cannot be detected.
func (device *IndustrialDigitalIn4Bricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *IndustrialDigitalIn4Bricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *IndustrialDigitalIn4Bricklet) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// This callback is triggered whenever a change of the voltage level is detected
// on pins where the interrupt was activated with SetInterrupt.
// 
// The values are a bitmask that specifies which interrupts occurred
// and the current value bitmask.
// 
// For example:
// 
// * (1, 1) or (0b0001, 0b0001) means that an interrupt on pin 0 occurred and
//   currently pin 0 is high and pins 1-3 are low.
// * (9, 14) or (0b1001, 0b1110) means that interrupts on pins 0 and 3
//   occurred and currently pin 0 is low and pins 1-3 are high.
func (device *IndustrialDigitalIn4Bricklet) RegisterInterruptCallback(fn func(uint16, uint16)) uint64 {
	wrapper := func(byteSlice []byte) {
		buf := bytes.NewBuffer(byteSlice[8:])
		var interruptMask uint16
		var valueMask uint16
		binary.Read(buf, binary.LittleEndian, &interruptMask)
		binary.Read(buf, binary.LittleEndian, &valueMask)
		fn(interruptMask, valueMask)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackInterrupt), wrapper)
}

// Remove a registered Interrupt callback.
func (device *IndustrialDigitalIn4Bricklet) DeregisterInterruptCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackInterrupt), registrationId)
}


// Returns the input value with a bitmask. The bitmask is 16bit long, *true*
// refers to high and *false* refers to low.
// 
// For example: The value 3 or 0b0011 means that pins 0-1 are high and the other
// pins are low.
// 
// If no groups are used (see SetGroup), the pins correspond to the
// markings on the IndustrialDigital In 4 Bricklet.
// 
// If groups are used, the pins correspond to the element in the group.
// Element 1 in the group will get pins 0-3, element 2 pins 4-7, element 3
// pins 8-11 and element 4 pins 12-15.
func (device *IndustrialDigitalIn4Bricklet) GetValue() (valueMask uint16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetValue), buf.Bytes())
	if err != nil {
		return valueMask, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return valueMask, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &valueMask)

	}

	return valueMask, nil
}

// Sets a group of Digital In 4 Bricklets that should work together. You can
// find Bricklets that can be grouped together with GetAvailableForGroup.
// 
// The group consists of 4 elements. Element 1 in the group will get pins 0-3,
// element 2 pins 4-7, element 3 pins 8-11 and element 4 pins 12-15.
// 
// Each element can either be one of the ports ('a' to 'd') or 'n' if it should
// not be used.
// 
// For example: If you have two Digital In 4 Bricklets connected to port A and
// port B respectively, you could call with ``['a', 'b', 'n', 'n']``.
// 
// Now the pins on the Digital In 4 on port A are assigned to 0-3 and the
// pins on the Digital In 4 on port B are assigned to 4-7. It is now possible
// to call GetValue and read out two Bricklets at the same time.
// 
// Changing the group configuration resets all edge counter configurations
// and values.
func (device *IndustrialDigitalIn4Bricklet) SetGroup(group [4]rune) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, group);

	resultBytes, err := device.device.Set(uint8(FunctionSetGroup), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])
		
	}

	return nil
}

// Returns the group as set by SetGroup
func (device *IndustrialDigitalIn4Bricklet) GetGroup() (group [4]rune, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetGroup), buf.Bytes())
	if err != nil {
		return group, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
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
func (device *IndustrialDigitalIn4Bricklet) GetAvailableForGroup() (available uint8, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetAvailableForGroup), buf.Bytes())
	if err != nil {
		return available, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return available, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &available)

	}

	return available, nil
}

// Sets the debounce period of the RegisterInterruptCallback callback in ms.
// 
// For example: If you set this value to 100, you will get the interrupt
// maximal every 100ms. This is necessary if something that bounces is
// connected to the Digital In 4 Bricklet, such as a button.
// 
// The default value is 100.
func (device *IndustrialDigitalIn4Bricklet) SetDebouncePeriod(debounce uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, debounce);

	resultBytes, err := device.device.Set(uint8(FunctionSetDebouncePeriod), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])
		
	}

	return nil
}

// Returns the debounce period as set by SetDebouncePeriod.
func (device *IndustrialDigitalIn4Bricklet) GetDebouncePeriod() (debounce uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetDebouncePeriod), buf.Bytes())
	if err != nil {
		return debounce, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return debounce, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &debounce)

	}

	return debounce, nil
}

// Sets the pins on which an interrupt is activated with a bitmask.
// Interrupts are triggered on changes of the voltage level of the pin,
// i.e. changes from high to low and low to high.
// 
// For example: An interrupt bitmask of 9 or 0b1001 will enable the interrupt for
// pins 0 and 3.
// 
// The interrupts use the grouping as set by SetGroup.
// 
// The interrupt is delivered with the RegisterInterruptCallback callback.
func (device *IndustrialDigitalIn4Bricklet) SetInterrupt(interruptMask uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, interruptMask);

	resultBytes, err := device.device.Set(uint8(FunctionSetInterrupt), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])
		
	}

	return nil
}

// Returns the interrupt bitmask as set by SetInterrupt.
func (device *IndustrialDigitalIn4Bricklet) GetInterrupt() (interruptMask uint16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetInterrupt), buf.Bytes())
	if err != nil {
		return interruptMask, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return interruptMask, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &interruptMask)

	}

	return interruptMask, nil
}

// Returns the current value of the edge counter for the selected pin. You can
// configure the edges that are counted with SetEdgeCountConfig.
// 
// If you set the reset counter to *true*, the count is set back to 0
// directly after it is read.
// 
// .. versionadded:: 2.0.1$nbsp;(Plugin)
func (device *IndustrialDigitalIn4Bricklet) GetEdgeCount(pin uint8, resetCounter bool) (count uint32, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, pin);
	binary.Write(&buf, binary.LittleEndian, resetCounter);

	resultBytes, err := device.device.Get(uint8(FunctionGetEdgeCount), buf.Bytes())
	if err != nil {
		return count, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return count, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &count)

	}

	return count, nil
}

// Configures the edge counter for the selected pins. A bitmask of 9 or 0b1001 will
// enable the edge counter for pins 0 and 3.
// 
// The edge type parameter configures if rising edges, falling edges or
// both are counted if the pin is configured for input. Possible edge types are:
// 
// * 0 = rising (default)
// * 1 = falling
// * 2 = both
// 
// The debounce time is given in ms.
// 
// Configuring an edge counter resets its value to 0.
// 
// If you don't know what any of this means, just leave it at default. The
// default configuration is very likely OK for you.
// 
// Default values: 0 (edge type) and 100ms (debounce time)
// 
// .. versionadded:: 2.0.1$nbsp;(Plugin)
//
// Associated constants:
//
//	* EdgeTypeRising
//	* EdgeTypeFalling
//	* EdgeTypeBoth
func (device *IndustrialDigitalIn4Bricklet) SetEdgeCountConfig(selectionMask uint16, edgeType EdgeType, debounce uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, selectionMask);
	binary.Write(&buf, binary.LittleEndian, edgeType);
	binary.Write(&buf, binary.LittleEndian, debounce);

	resultBytes, err := device.device.Set(uint8(FunctionSetEdgeCountConfig), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])
		
	}

	return nil
}

// Returns the edge type and debounce time for the selected pin as set by
// SetEdgeCountConfig.
// 
// .. versionadded:: 2.0.1$nbsp;(Plugin)
//
// Associated constants:
//
//	* EdgeTypeRising
//	* EdgeTypeFalling
//	* EdgeTypeBoth
func (device *IndustrialDigitalIn4Bricklet) GetEdgeCountConfig(pin uint8) (edgeType EdgeType, debounce uint8, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, pin);

	resultBytes, err := device.device.Get(uint8(FunctionGetEdgeCountConfig), buf.Bytes())
	if err != nil {
		return edgeType, debounce, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return edgeType, debounce, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &edgeType)
	binary.Read(resultBuf, binary.LittleEndian, &debounce)

	}

	return edgeType, debounce, nil
}

// Returns the UID, the UID where the Bricklet is connected to,
// the position, the hardware and firmware version as well as the
// device identifier.
// 
// The position can be 'a', 'b', 'c' or 'd'.
// 
// The device identifier numbers can be found `here <device_identifier>`.
// |device_identifier_constant|
func (device *IndustrialDigitalIn4Bricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
