/* ***********************************************************
 * This file was automatically generated on 2019-05-21.      *
 *                                                           *
 * Go Bindings Version 2.0.3                                 *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


//16-channel digital input/output.
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/IO16V2_Bricklet_Go.html.
package io16_v2_bricklet

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
	FunctionSetSelectedValue Function = 3
	FunctionSetConfiguration Function = 4
	FunctionGetConfiguration Function = 5
	FunctionSetInputValueCallbackConfiguration Function = 6
	FunctionGetInputValueCallbackConfiguration Function = 7
	FunctionSetAllInputValueCallbackConfiguration Function = 8
	FunctionGetAllInputValueCallbackConfiguration Function = 9
	FunctionSetMonoflop Function = 10
	FunctionGetMonoflop Function = 11
	FunctionGetEdgeCount Function = 12
	FunctionSetEdgeCountConfiguration Function = 13
	FunctionGetEdgeCountConfiguration Function = 14
	FunctionGetSPITFPErrorCount Function = 234
	FunctionSetBootloaderMode Function = 235
	FunctionGetBootloaderMode Function = 236
	FunctionSetWriteFirmwarePointer Function = 237
	FunctionWriteFirmware Function = 238
	FunctionSetStatusLEDConfig Function = 239
	FunctionGetStatusLEDConfig Function = 240
	FunctionGetChipTemperature Function = 242
	FunctionReset Function = 243
	FunctionWriteUID Function = 248
	FunctionReadUID Function = 249
	FunctionGetIdentity Function = 255
	FunctionCallbackInputValue Function = 15
	FunctionCallbackAllInputValue Function = 16
	FunctionCallbackMonoflopDone Function = 17
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

type BootloaderMode uint8

const (
    BootloaderModeBootloader BootloaderMode = 0
	BootloaderModeFirmware BootloaderMode = 1
	BootloaderModeBootloaderWaitForReboot BootloaderMode = 2
	BootloaderModeFirmwareWaitForReboot BootloaderMode = 3
	BootloaderModeFirmwareWaitForEraseAndReboot BootloaderMode = 4
)

type BootloaderStatus uint8

const (
    BootloaderStatusOK BootloaderStatus = 0
	BootloaderStatusInvalidMode BootloaderStatus = 1
	BootloaderStatusNoChange BootloaderStatus = 2
	BootloaderStatusEntryFunctionNotPresent BootloaderStatus = 3
	BootloaderStatusDeviceIdentifierIncorrect BootloaderStatus = 4
	BootloaderStatusCRCMismatch BootloaderStatus = 5
)

type StatusLEDConfig uint8

const (
    StatusLEDConfigOff StatusLEDConfig = 0
	StatusLEDConfigOn StatusLEDConfig = 1
	StatusLEDConfigShowHeartbeat StatusLEDConfig = 2
	StatusLEDConfigShowStatus StatusLEDConfig = 3
)

type IO16V2Bricklet struct{
	device Device
}
const DeviceIdentifier = 2114
const DeviceDisplayName = "IO-16 Bricklet 2.0"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (IO16V2Bricklet, error) {
    internalIPCon := ipcon.GetInternalHandle().(IPConnection)
    dev, err := NewDevice([3]uint8{ 2,0,0 }, uid, &internalIPCon, 0)
    if err != nil {
        return IO16V2Bricklet{}, err
    }
    dev.ResponseExpected[FunctionSetValue] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetValue] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetSelectedValue] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionSetConfiguration] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetInputValueCallbackConfiguration] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetInputValueCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetAllInputValueCallbackConfiguration] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetAllInputValueCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetMonoflop] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetMonoflop] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetEdgeCount] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetEdgeCountConfiguration] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetEdgeCountConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetSPITFPErrorCount] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetBootloaderMode] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetBootloaderMode] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetWriteFirmwarePointer] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionWriteFirmware] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetStatusLEDConfig] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetStatusLEDConfig] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetChipTemperature] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionReset] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionWriteUID] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionReadUID] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetIdentity] = ResponseExpectedFlagAlwaysTrue;
    return IO16V2Bricklet{dev}, nil
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
func (device *IO16V2Bricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *IO16V2Bricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
    return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *IO16V2Bricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *IO16V2Bricklet) GetAPIVersion() [3]uint8 {
    return device.device.GetAPIVersion()
}

// This callback is triggered periodically according to the configuration set by
// SetInputValueCallbackConfiguration.
// 
// The parameters are the channel, a value-changed indicator and the actual value
// for the channel. The `changed` parameter is true if the value has changed since
// the last callback.
func (device *IO16V2Bricklet) RegisterInputValueCallback(fn func(uint8, bool, bool)) uint64 {
            wrapper := func(byteSlice []byte) {
                buf := bytes.NewBuffer(byteSlice[8:])
                var channel uint8
var changed bool
var value bool
                binary.Read(buf, binary.LittleEndian, &channel)
binary.Read(buf, binary.LittleEndian, &changed)
binary.Read(buf, binary.LittleEndian, &value)
                fn(channel, changed, value)
            }
    return device.device.RegisterCallback(uint8(FunctionCallbackInputValue), wrapper)
}

//Remove a registered Input Value callback.
func (device *IO16V2Bricklet) DeregisterInputValueCallback(registrationID uint64) {
    device.device.DeregisterCallback(uint8(FunctionCallbackInputValue), registrationID)
}


// This callback is triggered periodically according to the configuration set by
// SetAllInputValueCallbackConfiguration.
// 
// The parameters are the same as GetValue. Additional the
// `changed` parameter is true if the value has changed since
// the last callback.
func (device *IO16V2Bricklet) RegisterAllInputValueCallback(fn func([16]bool, [16]bool)) uint64 {
            wrapper := func(byteSlice []byte) {
                buf := bytes.NewBuffer(byteSlice[8:])
                var changed [16]bool
var value [16]bool
                binary.Read(buf, binary.LittleEndian, &changed)
binary.Read(buf, binary.LittleEndian, &value)
                fn(changed, value)
            }
    return device.device.RegisterCallback(uint8(FunctionCallbackAllInputValue), wrapper)
}

//Remove a registered All Input Value callback.
func (device *IO16V2Bricklet) DeregisterAllInputValueCallback(registrationID uint64) {
    device.device.DeregisterCallback(uint8(FunctionCallbackAllInputValue), registrationID)
}


// This callback is triggered whenever a monoflop timer reaches 0. The
// parameters contain the channel and the current value of the channel
// (the value after the monoflop).
func (device *IO16V2Bricklet) RegisterMonoflopDoneCallback(fn func(uint8, bool)) uint64 {
            wrapper := func(byteSlice []byte) {
                buf := bytes.NewBuffer(byteSlice[8:])
                var channel uint8
var value bool
                binary.Read(buf, binary.LittleEndian, &channel)
binary.Read(buf, binary.LittleEndian, &value)
                fn(channel, value)
            }
    return device.device.RegisterCallback(uint8(FunctionCallbackMonoflopDone), wrapper)
}

//Remove a registered Monoflop Done callback.
func (device *IO16V2Bricklet) DeregisterMonoflopDoneCallback(registrationID uint64) {
    device.device.DeregisterCallback(uint8(FunctionCallbackMonoflopDone), registrationID)
}


// Sets the output value of all sixteen channels. A value of *true* or *false* outputs
// logic 1 or logic 0 respectively on the corresponding channel.
// 
// Use SetSelectedValue to change only one output channel state.
// 
// For example: (True, True, False, False, ..., False) will turn the channels 0-1
// high and the channels 2-15 low.
// 
// All running monoflop timers will be aborted if this function is called.
// 
// Note
//  This function does nothing for channels that are configured as input. Pull-up
//  resistors can be switched on with SetConfiguration.
func (device *IO16V2Bricklet) SetValue(value [16]bool) (err error) {
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, value);

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

// Returns the logic levels that are currently measured on the channels.
// This function works if the channel is configured as input as well as if it is
// configured as output.
func (device *IO16V2Bricklet) GetValue() (value [16]bool, err error) {
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetValue), buf.Bytes())
    if err != nil {
        return value, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader

        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return value, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &value)

    }

    return value, nil
}

// Sets the output value of a specific channel without affecting the other channels.
// 
// A running monoflop timer for the specific channel will be aborted if this
// function is called.
// 
// Note
//  This function does nothing for channels that are configured as input. Pull-up
//  resistors can be switched on with SetConfiguration.
func (device *IO16V2Bricklet) SetSelectedValue(channel uint8, value bool) (err error) {
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, channel);
	binary.Write(&buf, binary.LittleEndian, value);

    resultBytes, err := device.device.Set(uint8(FunctionSetSelectedValue), buf.Bytes())
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

// Configures the value and direction of a specific channel. Possible directions
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
// * (0, 'i', true) will set channel-0 as input pull-up.
// * (1, 'i', false) will set channel-1 as input default (floating if nothing is connected).
// * (2, 'o', true) will set channel-2 as output high.
// * (3, 'o', false) will set channel-3 as output low.
// 
// A running monoflop timer for the specific channel will be aborted if this
// function is called.
// 
// The default configuration is input with pull-up.
//
// Associated constants:
//
//	* DirectionIn
//	* DirectionOut
func (device *IO16V2Bricklet) SetConfiguration(channel uint8, direction Direction, value bool) (err error) {
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, channel);
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

// Returns the channel configuration as set by SetConfiguration.
//
// Associated constants:
//
//	* DirectionIn
//	* DirectionOut
func (device *IO16V2Bricklet) GetConfiguration(channel uint8) (direction Direction, value bool, err error) {
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, channel);

    resultBytes, err := device.device.Get(uint8(FunctionGetConfiguration), buf.Bytes())
    if err != nil {
        return direction, value, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader

        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return direction, value, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &direction)
	binary.Read(resultBuf, binary.LittleEndian, &value)

    }

    return direction, value, nil
}

// This callback can be configured per channel.
// 
// The period in ms is the period with which the RegisterInputValueCallback
// callback is triggered periodically. A value of 0 turns the callback off.
// 
// If the `value has to change`-parameter is set to true, the callback is only
// triggered after the value has changed. If the value didn't change within the
// period, the callback is triggered immediately on change.
// 
// If it is set to false, the callback is continuously triggered with the period,
// independent of the value.
// 
// The default value is (0, false).
func (device *IO16V2Bricklet) SetInputValueCallbackConfiguration(channel uint8, period uint32, valueHasToChange bool) (err error) {
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, channel);
	binary.Write(&buf, binary.LittleEndian, period);
	binary.Write(&buf, binary.LittleEndian, valueHasToChange);

    resultBytes, err := device.device.Set(uint8(FunctionSetInputValueCallbackConfiguration), buf.Bytes())
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

// Returns the callback configuration as set by
// SetInputValueCallbackConfiguration.
func (device *IO16V2Bricklet) GetInputValueCallbackConfiguration(channel uint8) (period uint32, valueHasToChange bool, err error) {
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, channel);

    resultBytes, err := device.device.Get(uint8(FunctionGetInputValueCallbackConfiguration), buf.Bytes())
    if err != nil {
        return period, valueHasToChange, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader

        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return period, valueHasToChange, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &period)
	binary.Read(resultBuf, binary.LittleEndian, &valueHasToChange)

    }

    return period, valueHasToChange, nil
}

// The period in ms is the period with which the RegisterAllInputValueCallback
// callback is triggered periodically. A value of 0 turns the callback off.
// 
// If the `value has to change`-parameter is set to true, the callback is only
// triggered after the value has changed. If the value didn't change within the
// period, the callback is triggered immediately on change.
// 
// If it is set to false, the callback is continuously triggered with the period,
// independent of the value.
// 
// The default value is (0, false).
func (device *IO16V2Bricklet) SetAllInputValueCallbackConfiguration(period uint32, valueHasToChange bool) (err error) {
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, period);
	binary.Write(&buf, binary.LittleEndian, valueHasToChange);

    resultBytes, err := device.device.Set(uint8(FunctionSetAllInputValueCallbackConfiguration), buf.Bytes())
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

// Returns the callback configuration as set by
// SetAllInputValueCallbackConfiguration.
func (device *IO16V2Bricklet) GetAllInputValueCallbackConfiguration() (period uint32, valueHasToChange bool, err error) {
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetAllInputValueCallbackConfiguration), buf.Bytes())
    if err != nil {
        return period, valueHasToChange, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader

        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return period, valueHasToChange, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &period)
	binary.Read(resultBuf, binary.LittleEndian, &valueHasToChange)

    }

    return period, valueHasToChange, nil
}

// Configures a monoflop of the specified channel.
// 
// The second parameter is the desired value of the specified
// channel. A *true* means relay closed and a *false* means relay open.
// 
// The third parameter indicates the time (in ms) that the channels should hold
// the value.
// 
// If this function is called with the parameters (0, 1, 1500) channel 0 will
// close and in 1.5s channel 0 will open again
// 
// A monoflop can be used as a fail-safe mechanism. For example: Lets assume you
// have a RS485 bus and a IO-16 Bricklet 2.0 connected to one of
// the slave stacks. You can now call this function every second, with a time
// parameter of two seconds and channel 0 closed. Channel 0 will be closed all the
// time. If now the RS485 connection is lost, then channel 0 will be opened in at
// most two seconds.
func (device *IO16V2Bricklet) SetMonoflop(channel uint8, value bool, time uint32) (err error) {
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, channel);
	binary.Write(&buf, binary.LittleEndian, value);
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

// Returns (for the given channel) the current value and the time as set by
// SetMonoflop as well as the remaining time until the value flips.
// 
// If the timer is not running currently, the remaining time will be returned
// as 0.
func (device *IO16V2Bricklet) GetMonoflop(channel uint8) (value bool, time uint32, timeRemaining uint32, err error) {
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, channel);

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

// Returns the current value of the edge counter for the selected channel. You can
// configure the edges that are counted with SetEdgeCountConfiguration.
// 
// If you set the reset counter to *true*, the count is set back to 0
// directly after it is read.
func (device *IO16V2Bricklet) GetEdgeCount(channel uint8, resetCounter bool) (count uint32, err error) {
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, channel);
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

// Configures the edge counter for a specific channel.
// 
// The edge type parameter configures if rising edges, falling edges or
// both are counted if the channel is configured for input. Possible edge types are:
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
// Associated constants:
//
//	* EdgeTypeRising
//	* EdgeTypeFalling
//	* EdgeTypeBoth
func (device *IO16V2Bricklet) SetEdgeCountConfiguration(channel uint8, edgeType EdgeType, debounce uint8) (err error) {
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, channel);
	binary.Write(&buf, binary.LittleEndian, edgeType);
	binary.Write(&buf, binary.LittleEndian, debounce);

    resultBytes, err := device.device.Set(uint8(FunctionSetEdgeCountConfiguration), buf.Bytes())
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

// Returns the edge type and debounce time for the selected channel as set by
// SetEdgeCountConfiguration.
//
// Associated constants:
//
//	* EdgeTypeRising
//	* EdgeTypeFalling
//	* EdgeTypeBoth
func (device *IO16V2Bricklet) GetEdgeCountConfiguration(channel uint8) (edgeType EdgeType, debounce uint8, err error) {
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, channel);

    resultBytes, err := device.device.Get(uint8(FunctionGetEdgeCountConfiguration), buf.Bytes())
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

// Returns the error count for the communication between Brick and Bricklet.
// 
// The errors are divided into
// 
// * ACK checksum errors,
// * message checksum errors,
// * framing errors and
// * overflow errors.
// 
// The errors counts are for errors that occur on the Bricklet side. All
// Bricks have a similar function that returns the errors on the Brick side.
func (device *IO16V2Bricklet) GetSPITFPErrorCount() (errorCountAckChecksum uint32, errorCountMessageChecksum uint32, errorCountFrame uint32, errorCountOverflow uint32, err error) {
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetSPITFPErrorCount), buf.Bytes())
    if err != nil {
        return errorCountAckChecksum, errorCountMessageChecksum, errorCountFrame, errorCountOverflow, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader

        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return errorCountAckChecksum, errorCountMessageChecksum, errorCountFrame, errorCountOverflow, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &errorCountAckChecksum)
	binary.Read(resultBuf, binary.LittleEndian, &errorCountMessageChecksum)
	binary.Read(resultBuf, binary.LittleEndian, &errorCountFrame)
	binary.Read(resultBuf, binary.LittleEndian, &errorCountOverflow)

    }

    return errorCountAckChecksum, errorCountMessageChecksum, errorCountFrame, errorCountOverflow, nil
}

// Sets the bootloader mode and returns the status after the requested
// mode change was instigated.
// 
// You can change from bootloader mode to firmware mode and vice versa. A change
// from bootloader mode to firmware mode will only take place if the entry function,
// device identifier and CRC are present and correct.
// 
// This function is used by Brick Viewer during flashing. It should not be
// necessary to call it in a normal user program.
//
// Associated constants:
//
//	* BootloaderModeBootloader
//	* BootloaderModeFirmware
//	* BootloaderModeBootloaderWaitForReboot
//	* BootloaderModeFirmwareWaitForReboot
//	* BootloaderModeFirmwareWaitForEraseAndReboot
//	* BootloaderStatusOK
//	* BootloaderStatusInvalidMode
//	* BootloaderStatusNoChange
//	* BootloaderStatusEntryFunctionNotPresent
//	* BootloaderStatusDeviceIdentifierIncorrect
//	* BootloaderStatusCRCMismatch
func (device *IO16V2Bricklet) SetBootloaderMode(mode BootloaderMode) (status BootloaderStatus, err error) {
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, mode);

    resultBytes, err := device.device.Get(uint8(FunctionSetBootloaderMode), buf.Bytes())
    if err != nil {
        return status, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader

        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return status, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &status)

    }

    return status, nil
}

// Returns the current bootloader mode, see SetBootloaderMode.
//
// Associated constants:
//
//	* BootloaderModeBootloader
//	* BootloaderModeFirmware
//	* BootloaderModeBootloaderWaitForReboot
//	* BootloaderModeFirmwareWaitForReboot
//	* BootloaderModeFirmwareWaitForEraseAndReboot
func (device *IO16V2Bricklet) GetBootloaderMode() (mode BootloaderMode, err error) {
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetBootloaderMode), buf.Bytes())
    if err != nil {
        return mode, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader

        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return mode, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &mode)

    }

    return mode, nil
}

// Sets the firmware pointer for WriteFirmware. The pointer has
// to be increased by chunks of size 64. The data is written to flash
// every 4 chunks (which equals to one page of size 256).
// 
// This function is used by Brick Viewer during flashing. It should not be
// necessary to call it in a normal user program.
func (device *IO16V2Bricklet) SetWriteFirmwarePointer(pointer uint32) (err error) {
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, pointer);

    resultBytes, err := device.device.Set(uint8(FunctionSetWriteFirmwarePointer), buf.Bytes())
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

// Writes 64 Bytes of firmware at the position as written by
// SetWriteFirmwarePointer before. The firmware is written
// to flash every 4 chunks.
// 
// You can only write firmware in bootloader mode.
// 
// This function is used by Brick Viewer during flashing. It should not be
// necessary to call it in a normal user program.
func (device *IO16V2Bricklet) WriteFirmware(data [64]uint8) (status uint8, err error) {
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, data);

    resultBytes, err := device.device.Get(uint8(FunctionWriteFirmware), buf.Bytes())
    if err != nil {
        return status, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader

        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return status, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &status)

    }

    return status, nil
}

// Sets the status LED configuration. By default the LED shows
// communication traffic between Brick and Bricklet, it flickers once
// for every 10 received data packets.
// 
// You can also turn the LED permanently on/off or show a heartbeat.
// 
// If the Bricklet is in bootloader mode, the LED is will show heartbeat by default.
//
// Associated constants:
//
//	* StatusLEDConfigOff
//	* StatusLEDConfigOn
//	* StatusLEDConfigShowHeartbeat
//	* StatusLEDConfigShowStatus
func (device *IO16V2Bricklet) SetStatusLEDConfig(config StatusLEDConfig) (err error) {
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, config);

    resultBytes, err := device.device.Set(uint8(FunctionSetStatusLEDConfig), buf.Bytes())
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

// Returns the configuration as set by SetStatusLEDConfig
//
// Associated constants:
//
//	* StatusLEDConfigOff
//	* StatusLEDConfigOn
//	* StatusLEDConfigShowHeartbeat
//	* StatusLEDConfigShowStatus
func (device *IO16V2Bricklet) GetStatusLEDConfig() (config StatusLEDConfig, err error) {
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetStatusLEDConfig), buf.Bytes())
    if err != nil {
        return config, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader

        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return config, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &config)

    }

    return config, nil
}

// Returns the temperature in °C as measured inside the microcontroller. The
// value returned is not the ambient temperature!
// 
// The temperature is only proportional to the real temperature and it has bad
// accuracy. Practically it is only useful as an indicator for
// temperature changes.
func (device *IO16V2Bricklet) GetChipTemperature() (temperature int16, err error) {
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetChipTemperature), buf.Bytes())
    if err != nil {
        return temperature, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader

        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return temperature, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &temperature)

    }

    return temperature, nil
}

// Calling this function will reset the Bricklet. All configurations
// will be lost.
// 
// After a reset you have to create new device objects,
// calling functions on the existing ones will result in
// undefined behavior!
func (device *IO16V2Bricklet) Reset() (err error) {
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Set(uint8(FunctionReset), buf.Bytes())
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

// Writes a new UID into flash. If you want to set a new UID
// you have to decode the Base58 encoded UID string into an
// integer first.
// 
// We recommend that you use Brick Viewer to change the UID.
func (device *IO16V2Bricklet) WriteUID(uid uint32) (err error) {
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, uid);

    resultBytes, err := device.device.Set(uint8(FunctionWriteUID), buf.Bytes())
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

// Returns the current UID as an integer. Encode as
// Base58 to get the usual string version.
func (device *IO16V2Bricklet) ReadUID() (uid uint32, err error) {
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionReadUID), buf.Bytes())
    if err != nil {
        return uid, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader

        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return uid, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &uid)

    }

    return uid, nil
}

// Returns the UID, the UID where the Bricklet is connected to,
// the position, the hardware and firmware version as well as the
// device identifier.
// 
// The position can be 'a', 'b', 'c' or 'd'.
// 
// The device identifier numbers can be found `here <device_identifier>`.
// |device_identifier_constant|
func (device *IO16V2Bricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
