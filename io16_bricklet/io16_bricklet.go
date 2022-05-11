/* ***********************************************************
 * This file was automatically generated on 2022-05-11.      *
 *                                                           *
 * Go Bindings Version 2.0.12                                *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


// 16-channel digital input/output.
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/IO16_Bricklet_Go.html.
package io16_bricklet

import (
	"encoding/binary"
	"bytes"
	"fmt"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionSetPort Function = 1
	FunctionGetPort Function = 2
	FunctionSetPortConfiguration Function = 3
	FunctionGetPortConfiguration Function = 4
	FunctionSetDebouncePeriod Function = 5
	FunctionGetDebouncePeriod Function = 6
	FunctionSetPortInterrupt Function = 7
	FunctionGetPortInterrupt Function = 8
	FunctionSetPortMonoflop Function = 10
	FunctionGetPortMonoflop Function = 11
	FunctionSetSelectedValues Function = 13
	FunctionGetEdgeCount Function = 14
	FunctionSetEdgeCountConfig Function = 15
	FunctionGetEdgeCountConfig Function = 16
	FunctionGetIdentity Function = 255
	FunctionCallbackInterrupt Function = 9
	FunctionCallbackMonoflopDone Function = 12
)

type Direction = rune

const (
	DirectionIn Direction = 'i'
	DirectionOut Direction = 'o'
)

type EdgeType = uint8

const (
	EdgeTypeRising EdgeType = 0
	EdgeTypeFalling EdgeType = 1
	EdgeTypeBoth EdgeType = 2
)

type IO16Bricklet struct {
	device Device
}
const DeviceIdentifier = 28
const DeviceDisplayName = "IO-16 Bricklet"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (IO16Bricklet, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{ 2,0,1 }, uid, &internalIPCon, 0, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return IO16Bricklet{}, err
	}
	dev.ResponseExpected[FunctionSetPort] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetPort] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetPortConfiguration] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetPortConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetDebouncePeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetDebouncePeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetPortInterrupt] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetPortInterrupt] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetPortMonoflop] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetPortMonoflop] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetSelectedValues] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetEdgeCount] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetEdgeCountConfig] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetEdgeCountConfig] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetIdentity] = ResponseExpectedFlagAlwaysTrue;
	return IO16Bricklet{dev}, nil
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
func (device *IO16Bricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *IO16Bricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *IO16Bricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *IO16Bricklet) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// This callback is triggered whenever a change of the voltage level is detected
// on pins where the interrupt was activated with SetPortInterrupt.
// 
// The values are the port, a bitmask that specifies which interrupts occurred
// and the current value bitmask of the port.
// 
// For example:
// 
// * ('a', 1, 1) or ('a', 0b00000001, 0b00000001) means that on port A an
//   interrupt on pin 0 occurred and currently pin 0 is high and pins 1-7 are low.
// * ('b', 129, 254) or ('b', 0b10000001, 0b11111110) means that on port B
//   interrupts on pins 0 and 7 occurred and currently pin 0 is low and pins 1-7
//   are high.
func (device *IO16Bricklet) RegisterInterruptCallback(fn func(rune, uint8, uint8)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 11 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var port rune
		var interruptMask uint8
		var valueMask uint8
		port = rune(buf.Next(1)[0])
		binary.Read(buf, binary.LittleEndian, &interruptMask)
		binary.Read(buf, binary.LittleEndian, &valueMask)
		fn(port, interruptMask, valueMask)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackInterrupt), wrapper)
}

// Remove a registered Interrupt callback.
func (device *IO16Bricklet) DeregisterInterruptCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackInterrupt), registrationId)
}


// This callback is triggered whenever a monoflop timer reaches 0. The
// parameters contain the port, the involved pins and the current value of
// the pins (the value after the monoflop).
func (device *IO16Bricklet) RegisterMonoflopDoneCallback(fn func(rune, uint8, uint8)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 11 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var port rune
		var selectionMask uint8
		var valueMask uint8
		port = rune(buf.Next(1)[0])
		binary.Read(buf, binary.LittleEndian, &selectionMask)
		binary.Read(buf, binary.LittleEndian, &valueMask)
		fn(port, selectionMask, valueMask)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackMonoflopDone), wrapper)
}

// Remove a registered Monoflop Done callback.
func (device *IO16Bricklet) DeregisterMonoflopDoneCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackMonoflopDone), registrationId)
}


// Sets the output value (high or low) for a port (a or b) with a bitmask
// (8bit). A 1 in the bitmask means high and a 0 in the bitmask means low.
// 
// For example: The value 15 or 0b00001111 will turn the pins 0-3 high and the
// pins 4-7 low for the specified port.
// 
// All running monoflop timers of the given port will be aborted if this function
// is called.
// 
// Note
//  This function does nothing for pins that are configured as input.
//  Pull-up resistors can be switched on with SetPortConfiguration.
func (device *IO16Bricklet) SetPort(port rune, valueMask uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, port);
	binary.Write(&buf, binary.LittleEndian, valueMask);

	resultBytes, err := device.device.Set(uint8(FunctionSetPort), buf.Bytes())
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

// Returns a bitmask of the values that are currently measured on the
// specified port. This function works if the pin is configured to input
// as well as if it is configured to output.
func (device *IO16Bricklet) GetPort(port rune) (valueMask uint8, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, port);

	resultBytes, err := device.device.Get(uint8(FunctionGetPort), buf.Bytes())
	if err != nil {
		return valueMask, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return valueMask, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return valueMask, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &valueMask)

	}

	return valueMask, nil
}

// Configures the value and direction of a specified port. Possible directions
// are 'i' and 'o' for input and output.
// 
// If the direction is configured as output, the value is either high or low
// (set as *true* or *false*).
// 
// If the direction is configured as input, the value is either pull-up or
// default (set as *true* or *false*).
// 
// For example:
// 
// * ('a', 255, 'i', true) or ('a', 0b11111111, 'i', true) will set all pins of port A as input pull-up.
// * ('a', 128, 'i', false) or ('a', 0b10000000, 'i', false) will set pin 7 of port A as input default (floating if nothing is connected).
// * ('b', 3, 'o', false) or ('b', 0b00000011, 'o', false) will set pins 0 and 1 of port B as output low.
// * ('b', 4, 'o', true) or ('b', 0b00000100, 'o', true) will set pin 2 of port B as output high.
// 
// Running monoflop timers for the selected pins will be aborted if this
// function is called.
//
// Associated constants:
//
//	* DirectionIn
//	* DirectionOut
func (device *IO16Bricklet) SetPortConfiguration(port rune, selectionMask uint8, direction Direction, value bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, port);
	binary.Write(&buf, binary.LittleEndian, selectionMask);
	binary.Write(&buf, binary.LittleEndian, direction);
	binary.Write(&buf, binary.LittleEndian, value);

	resultBytes, err := device.device.Set(uint8(FunctionSetPortConfiguration), buf.Bytes())
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

// Returns a direction bitmask and a value bitmask for the specified port. A 1 in
// the direction bitmask means input and a 0 in the bitmask means output.
// 
// For example: A return value of (15, 51) or (0b00001111, 0b00110011) for
// direction and value means that:
// 
// * pins 0 and 1 are configured as input pull-up,
// * pins 2 and 3 are configured as input default,
// * pins 4 and 5 are configured as output high
// * and pins 6 and 7 are configured as output low.
func (device *IO16Bricklet) GetPortConfiguration(port rune) (directionMask uint8, valueMask uint8, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, port);

	resultBytes, err := device.device.Get(uint8(FunctionGetPortConfiguration), buf.Bytes())
	if err != nil {
		return directionMask, valueMask, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 10 {
			return directionMask, valueMask, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		if header.ErrorCode != 0 {
			return directionMask, valueMask, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &directionMask)
		binary.Read(resultBuf, binary.LittleEndian, &valueMask)

	}

	return directionMask, valueMask, nil
}

// Sets the debounce period of the RegisterInterruptCallback callback.
// 
// For example: If you set this value to 100, you will get the interrupt
// maximal every 100ms. This is necessary if something that bounces is
// connected to the IO-16 Bricklet, such as a button.
func (device *IO16Bricklet) SetDebouncePeriod(debounce uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, debounce);

	resultBytes, err := device.device.Set(uint8(FunctionSetDebouncePeriod), buf.Bytes())
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

// Returns the debounce period as set by SetDebouncePeriod.
func (device *IO16Bricklet) GetDebouncePeriod() (debounce uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetDebouncePeriod), buf.Bytes())
	if err != nil {
		return debounce, err
	}
	if len(resultBytes) > 0 {
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

	}

	return debounce, nil
}

// Sets the pins on which an interrupt is activated with a bitmask.
// Interrupts are triggered on changes of the voltage level of the pin,
// i.e. changes from high to low and low to high.
// 
// For example: ('a', 129) or ('a', 0b10000001) will enable the interrupt for
// pins 0 and 7 of port a.
// 
// The interrupt is delivered with the RegisterInterruptCallback callback.
func (device *IO16Bricklet) SetPortInterrupt(port rune, interruptMask uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, port);
	binary.Write(&buf, binary.LittleEndian, interruptMask);

	resultBytes, err := device.device.Set(uint8(FunctionSetPortInterrupt), buf.Bytes())
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

// Returns the interrupt bitmask for the specified port as set by
// SetPortInterrupt.
func (device *IO16Bricklet) GetPortInterrupt(port rune) (interruptMask uint8, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, port);

	resultBytes, err := device.device.Get(uint8(FunctionGetPortInterrupt), buf.Bytes())
	if err != nil {
		return interruptMask, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return interruptMask, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return interruptMask, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &interruptMask)

	}

	return interruptMask, nil
}

// Configures a monoflop of the pins specified by the second parameter as 8 bit
// long bitmask. The specified pins must be configured for output. Non-output
// pins will be ignored.
// 
// The third parameter is a bitmask with the desired value of the specified
// output pins. A 1 in the bitmask means high and a 0 in the bitmask means low.
// 
// The forth parameter indicates the time that the pins should hold
// the value.
// 
// If this function is called with the parameters ('a', 9, 1, 1500) or
// ('a', 0b00001001, 0b00000001, 1500): Pin 0 will get high and pin 3 will get
// low on port 'a'. In 1.5s pin 0 will get low and pin 3 will get high again.
// 
// A monoflop can be used as a fail-safe mechanism. For example: Lets assume you
// have a RS485 bus and an IO-16 Bricklet connected to one of the slave
// stacks. You can now call this function every second, with a time parameter
// of two seconds and pin 0 set to high. Pin 0 will be high all the time. If now
// the RS485 connection is lost, then pin 0 will get low in at most two seconds.
func (device *IO16Bricklet) SetPortMonoflop(port rune, selectionMask uint8, valueMask uint8, time uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, port);
	binary.Write(&buf, binary.LittleEndian, selectionMask);
	binary.Write(&buf, binary.LittleEndian, valueMask);
	binary.Write(&buf, binary.LittleEndian, time);

	resultBytes, err := device.device.Set(uint8(FunctionSetPortMonoflop), buf.Bytes())
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
// SetPortMonoflop as well as the remaining time until the value flips.
// 
// If the timer is not running currently, the remaining time will be returned
// as 0.
func (device *IO16Bricklet) GetPortMonoflop(port rune, pin uint8) (value uint8, time uint32, timeRemaining uint32, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, port);
	binary.Write(&buf, binary.LittleEndian, pin);

	resultBytes, err := device.device.Get(uint8(FunctionGetPortMonoflop), buf.Bytes())
	if err != nil {
		return value, time, timeRemaining, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 17 {
			return value, time, timeRemaining, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 17)
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

// Sets the output value (high or low) for a port (a or b with a bitmask,
// according to the selection mask. The bitmask is 8 bit long and a 1 in the
// bitmask means high and a 0 in the bitmask means low.
// 
// For example: The parameters ('a', 192, 128) or ('a', 0b11000000, 0b10000000)
// will turn pin 7 high and pin 6 low on port A, pins 0-6 will remain untouched.
// 
// Running monoflop timers for the selected pins will be aborted if this
// function is called.
// 
// Note
//  This function does nothing for pins that are configured as input.
//  Pull-up resistors can be switched on with SetPortConfiguration.
func (device *IO16Bricklet) SetSelectedValues(port rune, selectionMask uint8, valueMask uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, port);
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

// Returns the current value of the edge counter for the selected pin on port A.
// You can configure the edges that are counted with SetEdgeCountConfig.
// 
// If you set the reset counter to *true*, the count is set back to 0
// directly after it is read.
// 
// .. versionadded:: 2.0.3$nbsp;(Plugin)
func (device *IO16Bricklet) GetEdgeCount(pin uint8, resetCounter bool) (count uint32, err error) {
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

		if header.Length != 12 {
			return count, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		if header.ErrorCode != 0 {
			return count, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &count)

	}

	return count, nil
}

// Configures the edge counter for the selected pin of port A. Pins 0 and 1
// are available for edge counting.
// 
// The edge type parameter configures if rising edges, falling edges or
// both are counted if the pin is configured for input. Possible edge types are:
// 
// * 0 = rising
// * 1 = falling
// * 2 = both
// 
// Configuring an edge counter resets its value to 0.
// 
// If you don't know what any of this means, just leave it at default. The
// default configuration is very likely OK for you.
// 
// .. versionadded:: 2.0.3$nbsp;(Plugin)
//
// Associated constants:
//
//	* EdgeTypeRising
//	* EdgeTypeFalling
//	* EdgeTypeBoth
func (device *IO16Bricklet) SetEdgeCountConfig(pin uint8, edgeType EdgeType, debounce uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, pin);
	binary.Write(&buf, binary.LittleEndian, edgeType);
	binary.Write(&buf, binary.LittleEndian, debounce);

	resultBytes, err := device.device.Set(uint8(FunctionSetEdgeCountConfig), buf.Bytes())
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

// Returns the edge type and debounce time for the selected pin of port A as set by
// SetEdgeCountConfig.
// 
// .. versionadded:: 2.0.3$nbsp;(Plugin)
//
// Associated constants:
//
//	* EdgeTypeRising
//	* EdgeTypeFalling
//	* EdgeTypeBoth
func (device *IO16Bricklet) GetEdgeCountConfig(pin uint8) (edgeType EdgeType, debounce uint8, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, pin);

	resultBytes, err := device.device.Get(uint8(FunctionGetEdgeCountConfig), buf.Bytes())
	if err != nil {
		return edgeType, debounce, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 10 {
			return edgeType, debounce, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

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
// The position can be 'a', 'b', 'c', 'd', 'e', 'f', 'g' or 'h' (Bricklet Port).
// A Bricklet connected to an `Isolator Bricklet <isolator_bricklet>` is always at
// position 'z'.
// 
// The device identifier numbers can be found `here <device_identifier>`.
// |device_identifier_constant|
func (device *IO16Bricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
