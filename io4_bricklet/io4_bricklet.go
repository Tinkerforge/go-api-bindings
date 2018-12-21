/* ***********************************************************
 * This file was automatically generated on 2018-12-21.      *
 *                                                           *
 * Go Bindings Version 2.0.0                                 *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


//4-channel digital input/output.
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/IO4_Bricklet_Go.html.
package io4_bricklet

import (
	"encoding/binary"
	"bytes"
    . "github.com/Tinkerforge/go-api-bindings/internal"
    "github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function uint8

const (
    FunctionSetValue Function = 1
	FunctionGetValue Function = 2
	FunctionSetConfiguration Function = 3
	FunctionGetConfiguration Function = 4
	FunctionSetDebouncePeriod Function = 5
	FunctionGetDebouncePeriod Function = 6
	FunctionSetInterrupt Function = 7
	FunctionGetInterrupt Function = 8
	FunctionSetMonoflop Function = 10
	FunctionGetMonoflop Function = 11
	FunctionSetSelectedValues Function = 13
	FunctionGetEdgeCount Function = 14
	FunctionSetEdgeCountConfig Function = 15
	FunctionGetEdgeCountConfig Function = 16
	FunctionGetIdentity Function = 255
	FunctionCallbackInterrupt Function = 9
	FunctionCallbackMonoflopDone Function = 12
)

type Direction rune

const (
    DirectionIn Direction = 'i'
	DirectionOut Direction = 'o'
)

type EdgeType uint8

const (
    EdgeTypeRising EdgeType = 0
	EdgeTypeFalling EdgeType = 1
	EdgeTypeBoth EdgeType = 2
)

type IO4Bricklet struct{
	device Device
}
const DeviceIdentifier = 29
const DeviceDisplayName = "IO-4 Bricklet"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (IO4Bricklet, error) {
    internalIPCon := ipcon.GetInternalHandle().(IPConnection)
    dev, err := NewDevice([3]uint8{ 2,0,1 }, uid, &internalIPCon, 0)
    if err != nil {
        return IO4Bricklet{}, err
    }
    dev.ResponseExpected[FunctionSetValue] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetValue] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetConfiguration] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetDebouncePeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetDebouncePeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetInterrupt] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetInterrupt] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetMonoflop] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetMonoflop] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetSelectedValues] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetEdgeCount] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetEdgeCountConfig] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetEdgeCountConfig] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetIdentity] = ResponseExpectedFlagAlwaysTrue;
    return IO4Bricklet{dev}, nil
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
func (device *IO4Bricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *IO4Bricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
    return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *IO4Bricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *IO4Bricklet) GetAPIVersion() [3]uint8 {
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
func (device *IO4Bricklet) RegisterInterruptCallback(fn func(uint8, uint8)) uint64 {
            wrapper := func(byteSlice []byte) {
                buf := bytes.NewBuffer(byteSlice[8:])
                var interruptMask uint8
var valueMask uint8
                binary.Read(buf, binary.LittleEndian, &interruptMask)
binary.Read(buf, binary.LittleEndian, &valueMask)
                fn(interruptMask, valueMask)
            }
    return device.device.RegisterCallback(uint8(FunctionCallbackInterrupt), wrapper)
}

//Remove a registered Interrupt callback.
func (device *IO4Bricklet) DeregisterInterruptCallback(callbackID uint64) {
    device.device.DeregisterCallback(uint8(FunctionCallbackInterrupt), callbackID)
}


// This callback is triggered whenever a monoflop timer reaches 0. The
	// parameters contain the involved pins and the current value of the pins
	// (the value after the monoflop).
func (device *IO4Bricklet) RegisterMonoflopDoneCallback(fn func(uint8, uint8)) uint64 {
            wrapper := func(byteSlice []byte) {
                buf := bytes.NewBuffer(byteSlice[8:])
                var selectionMask uint8
var valueMask uint8
                binary.Read(buf, binary.LittleEndian, &selectionMask)
binary.Read(buf, binary.LittleEndian, &valueMask)
                fn(selectionMask, valueMask)
            }
    return device.device.RegisterCallback(uint8(FunctionCallbackMonoflopDone), wrapper)
}

//Remove a registered Monoflop Done callback.
func (device *IO4Bricklet) DeregisterMonoflopDoneCallback(callbackID uint64) {
    device.device.DeregisterCallback(uint8(FunctionCallbackMonoflopDone), callbackID)
}


// Sets the output value (high or low) with a bitmask (4bit). A 1 in the bitmask
	// means high and a 0 in the bitmask means low.
	// 
	// For example: The value 3 or 0b0011 will turn the pins 0-1 high and the
	// pins 2-3 low.
	// 
	// Note
	//  This function does nothing for pins that are configured as input.
	//  Pull-up resistors can be switched on with SetConfiguration.
func (device *IO4Bricklet) SetValue(valueMask uint8) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, valueMask);

    resultBytes, err := device.device.Set(uint8(FunctionSetValue), buf.Bytes())
    if err != nil {
        return err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return BrickletError(header.ErrorCode)
        }

        bytes.NewBuffer(resultBytes[8:])
        
    }
    
    return nil
}

// Returns a bitmask of the values that are currently measured.
	// This function works if the pin is configured to input
	// as well as if it is configured to output.
func (device *IO4Bricklet) GetValue() (valueMask uint8, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetValue), buf.Bytes())
    if err != nil {
        return valueMask, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return valueMask, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &valueMask)

    }
    
    return valueMask, nil
}

// Configures the value and direction of the specified pins. Possible directions
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
	// * (15, 'i', true) or (0b1111, 'i', true) will set all pins of as input pull-up.
	// * (8, 'i', false) or (0b1000, 'i', false) will set pin 3 of as input default (floating if nothing is connected).
	// * (3, 'o', false) or (0b0011, 'o', false) will set pins 0 and 1 as output low.
	// * (4, 'o', true) or (0b0100, 'o', true) will set pin 2 of as output high.
	// 
	// The default configuration is input with pull-up.
//
// Associated constants:
//
//	* DirectionIn
//	* DirectionOut
func (device *IO4Bricklet) SetConfiguration(selectionMask uint8, direction Direction, value bool) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, selectionMask);
	binary.Write(&buf, binary.LittleEndian, direction);
	binary.Write(&buf, binary.LittleEndian, value);

    resultBytes, err := device.device.Set(uint8(FunctionSetConfiguration), buf.Bytes())
    if err != nil {
        return err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return BrickletError(header.ErrorCode)
        }

        bytes.NewBuffer(resultBytes[8:])
        
    }
    
    return nil
}

// Returns a value bitmask and a direction bitmask. A 1 in the direction bitmask
	// means input and a 0 in the bitmask means output.
	// 
	// For example: A return value of (3, 5) or (0b0011, 0b0101) for direction and
	// value means that:
	// 
	// * pin 0 is configured as input pull-up,
	// * pin 1 is configured as input default,
	// * pin 2 is configured as output high and
	// * pin 3 is are configured as output low.
func (device *IO4Bricklet) GetConfiguration() (directionMask uint8, valueMask uint8, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetConfiguration), buf.Bytes())
    if err != nil {
        return directionMask, valueMask, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return directionMask, valueMask, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &directionMask)
	binary.Read(resultBuf, binary.LittleEndian, &valueMask)

    }
    
    return directionMask, valueMask, nil
}

// Sets the debounce period of the RegisterInterruptCallback callback in ms.
	// 
	// For example: If you set this value to 100, you will get the interrupt
	// maximal every 100ms. This is necessary if something that bounces is
	// connected to the IO-4 Bricklet, such as a button.
	// 
	// The default value is 100.
func (device *IO4Bricklet) SetDebouncePeriod(debounce uint32) (err error) {    
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
            return BrickletError(header.ErrorCode)
        }

        bytes.NewBuffer(resultBytes[8:])
        
    }
    
    return nil
}

// Returns the debounce period as set by SetDebouncePeriod.
func (device *IO4Bricklet) GetDebouncePeriod() (debounce uint32, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetDebouncePeriod), buf.Bytes())
    if err != nil {
        return debounce, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return debounce, BrickletError(header.ErrorCode)
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
	// For example: An interrupt bitmask of 10 or 0b1010 will enable the interrupt for
	// pins 1 and 3.
	// 
	// The interrupt is delivered with the RegisterInterruptCallback callback.
func (device *IO4Bricklet) SetInterrupt(interruptMask uint8) (err error) {    
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
            return BrickletError(header.ErrorCode)
        }

        bytes.NewBuffer(resultBytes[8:])
        
    }
    
    return nil
}

// Returns the interrupt bitmask as set by SetInterrupt.
func (device *IO4Bricklet) GetInterrupt() (interruptMask uint8, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetInterrupt), buf.Bytes())
    if err != nil {
        return interruptMask, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return interruptMask, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &interruptMask)

    }
    
    return interruptMask, nil
}

// Configures a monoflop of the pins specified by the first parameter as 4 bit
	// long bitmask. The specified pins must be configured for output. Non-output
	// pins will be ignored.
	// 
	// The second parameter is a bitmask with the desired value of the specified
	// output pins. A 1 in the bitmask means high and a 0 in the bitmask means low.
	// 
	// The third parameter indicates the time (in ms) that the pins should hold
	// the value.
	// 
	// If this function is called with the parameters (9, 1, 1500) or
	// (0b1001, 0b0001, 1500): Pin 0 will get high and pin 3 will get low. In 1.5s pin
	// 0 will get low and pin 3 will get high again.
	// 
	// A monoflop can be used as a fail-safe mechanism. For example: Lets assume you
	// have a RS485 bus and an IO-4 Bricklet connected to one of the slave
	// stacks. You can now call this function every second, with a time parameter
	// of two seconds and pin 0 set to high. Pin 0 will be high all the time. If now
	// the RS485 connection is lost, then pin 0 will get low in at most two seconds.
func (device *IO4Bricklet) SetMonoflop(selectionMask uint8, valueMask uint8, time uint32) (err error) {    
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
        if header.ErrorCode != 0 {
            return BrickletError(header.ErrorCode)
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
func (device *IO4Bricklet) GetMonoflop(pin uint8) (value uint8, time uint32, timeRemaining uint32, err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, pin);

    resultBytes, err := device.device.Get(uint8(FunctionGetMonoflop), buf.Bytes())
    if err != nil {
        return value, time, timeRemaining, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return value, time, timeRemaining, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &value)
	binary.Read(resultBuf, binary.LittleEndian, &time)
	binary.Read(resultBuf, binary.LittleEndian, &timeRemaining)

    }
    
    return value, time, timeRemaining, nil
}

// Sets the output value (high or low) with a bitmask, according to
	// the selection mask. The bitmask is 4 bit long, *true* refers to high
	// and *false* refers to low.
	// 
	// For example: The parameters (9, 4) or (0b0110, 0b0100) will turn
	// pin 1 low and pin 2 high, pin 0 and 3 will remain untouched.
	// 
	// Note
	//  This function does nothing for pins that are configured as input.
	//  Pull-up resistors can be switched on with SetConfiguration.
func (device *IO4Bricklet) SetSelectedValues(selectionMask uint8, valueMask uint8) (err error) {    
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
        if header.ErrorCode != 0 {
            return BrickletError(header.ErrorCode)
        }

        bytes.NewBuffer(resultBytes[8:])
        
    }
    
    return nil
}

// Returns the current value of the edge counter for the selected pin. You can
	// configure the edges that are counted with SetEdgeCountConfig.
	// 
	// If you set the reset counter to *true*, the count is set back to 0
	// directly after it is read.
	// 
	// .. versionadded:: 2.0.1$nbsp;(Plugin)
func (device *IO4Bricklet) GetEdgeCount(pin uint8, resetCounter bool) (count uint32, err error) {    
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
            return count, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &count)

    }
    
    return count, nil
}

// Configures the edge counter for the selected pins.
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
func (device *IO4Bricklet) SetEdgeCountConfig(selectionMask uint8, edgeType EdgeType, debounce uint8) (err error) {    
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
            return BrickletError(header.ErrorCode)
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
func (device *IO4Bricklet) GetEdgeCountConfig(pin uint8) (edgeType EdgeType, debounce uint8, err error) {    
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
            return edgeType, debounce, BrickletError(header.ErrorCode)
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
func (device *IO4Bricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetIdentity), buf.Bytes())
    if err != nil {
        return uid, connectedUid, position, hardwareVersion, firmwareVersion, deviceIdentifier, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return uid, connectedUid, position, hardwareVersion, firmwareVersion, deviceIdentifier, BrickletError(header.ErrorCode)
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
