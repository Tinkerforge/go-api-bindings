/* ***********************************************************
 * This file was automatically generated on 2018-12-21.      *
 *                                                           *
 * Go Bindings Version 2.0.0                                 *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


//Measures DC voltage between 0V and 45V‍.
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/AnalogIn_Bricklet_Go.html.
package analog_in_bricklet

import (
	"encoding/binary"
	"bytes"
    . "github.com/Tinkerforge/go-api-bindings/internal"
    "github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function uint8

const (
    FunctionGetVoltage Function = 1
	FunctionGetAnalogValue Function = 2
	FunctionSetVoltageCallbackPeriod Function = 3
	FunctionGetVoltageCallbackPeriod Function = 4
	FunctionSetAnalogValueCallbackPeriod Function = 5
	FunctionGetAnalogValueCallbackPeriod Function = 6
	FunctionSetVoltageCallbackThreshold Function = 7
	FunctionGetVoltageCallbackThreshold Function = 8
	FunctionSetAnalogValueCallbackThreshold Function = 9
	FunctionGetAnalogValueCallbackThreshold Function = 10
	FunctionSetDebouncePeriod Function = 11
	FunctionGetDebouncePeriod Function = 12
	FunctionSetRange Function = 17
	FunctionGetRange Function = 18
	FunctionSetAveraging Function = 19
	FunctionGetAveraging Function = 20
	FunctionGetIdentity Function = 255
	FunctionCallbackVoltage Function = 13
	FunctionCallbackAnalogValue Function = 14
	FunctionCallbackVoltageReached Function = 15
	FunctionCallbackAnalogValueReached Function = 16
)

type ThresholdOption rune

const (
    ThresholdOptionOff ThresholdOption = 'x'
	ThresholdOptionOutside ThresholdOption = 'o'
	ThresholdOptionInside ThresholdOption = 'i'
	ThresholdOptionSmaller ThresholdOption = '<'
	ThresholdOptionGreater ThresholdOption = '>'
)

type Range uint8

const (
    RangeAutomatic Range = 0
	RangeUpTo6V Range = 1
	RangeUpTo10V Range = 2
	RangeUpTo36V Range = 3
	RangeUpTo45V Range = 4
	RangeUpTo3V Range = 5
)

type AnalogInBricklet struct{
	device Device
}
const DeviceIdentifier = 219
const DeviceDisplayName = "Analog In Bricklet"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (AnalogInBricklet, error) {
    internalIPCon := ipcon.GetInternalHandle().(IPConnection)
    dev, err := NewDevice([3]uint8{ 2,0,3 }, uid, &internalIPCon, 0)
    if err != nil {
        return AnalogInBricklet{}, err
    }
    dev.ResponseExpected[FunctionGetVoltage] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetAnalogValue] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetVoltageCallbackPeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetVoltageCallbackPeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetAnalogValueCallbackPeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetAnalogValueCallbackPeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetVoltageCallbackThreshold] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetVoltageCallbackThreshold] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetAnalogValueCallbackThreshold] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetAnalogValueCallbackThreshold] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetDebouncePeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetDebouncePeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetRange] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetRange] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetAveraging] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetAveraging] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetIdentity] = ResponseExpectedFlagAlwaysTrue;
    return AnalogInBricklet{dev}, nil
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
func (device *AnalogInBricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *AnalogInBricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
    return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *AnalogInBricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *AnalogInBricklet) GetAPIVersion() [3]uint8 {
    return device.device.GetAPIVersion()
}

// This callback is triggered periodically with the period that is set by
	// SetVoltageCallbackPeriod. The parameter is the voltage of the
	// sensor.
	// 
	// The RegisterVoltageCallback callback is only triggered if the voltage has changed since
	// the last triggering.
func (device *AnalogInBricklet) RegisterVoltageCallback(fn func(uint16)) uint64 {
            wrapper := func(byteSlice []byte) {
                buf := bytes.NewBuffer(byteSlice[8:])
                var voltage uint16
                binary.Read(buf, binary.LittleEndian, &voltage)
                fn(voltage)
            }
    return device.device.RegisterCallback(uint8(FunctionCallbackVoltage), wrapper)
}

//Remove a registered Voltage callback.
func (device *AnalogInBricklet) DeregisterVoltageCallback(callbackID uint64) {
    device.device.DeregisterCallback(uint8(FunctionCallbackVoltage), callbackID)
}


// This callback is triggered periodically with the period that is set by
	// SetAnalogValueCallbackPeriod. The parameter is the analog
	// value of the sensor.
	// 
	// The RegisterAnalogValueCallback callback is only triggered if the voltage has changed
	// since the last triggering.
func (device *AnalogInBricklet) RegisterAnalogValueCallback(fn func(uint16)) uint64 {
            wrapper := func(byteSlice []byte) {
                buf := bytes.NewBuffer(byteSlice[8:])
                var value uint16
                binary.Read(buf, binary.LittleEndian, &value)
                fn(value)
            }
    return device.device.RegisterCallback(uint8(FunctionCallbackAnalogValue), wrapper)
}

//Remove a registered Analog Value callback.
func (device *AnalogInBricklet) DeregisterAnalogValueCallback(callbackID uint64) {
    device.device.DeregisterCallback(uint8(FunctionCallbackAnalogValue), callbackID)
}


// This callback is triggered when the threshold as set by
	// SetVoltageCallbackThreshold is reached.
	// The parameter is the voltage of the sensor.
	// 
	// If the threshold keeps being reached, the callback is triggered periodically
	// with the period as set by SetDebouncePeriod.
func (device *AnalogInBricklet) RegisterVoltageReachedCallback(fn func(uint16)) uint64 {
            wrapper := func(byteSlice []byte) {
                buf := bytes.NewBuffer(byteSlice[8:])
                var voltage uint16
                binary.Read(buf, binary.LittleEndian, &voltage)
                fn(voltage)
            }
    return device.device.RegisterCallback(uint8(FunctionCallbackVoltageReached), wrapper)
}

//Remove a registered Voltage Reached callback.
func (device *AnalogInBricklet) DeregisterVoltageReachedCallback(callbackID uint64) {
    device.device.DeregisterCallback(uint8(FunctionCallbackVoltageReached), callbackID)
}


// This callback is triggered when the threshold as set by
	// SetAnalogValueCallbackThreshold is reached.
	// The parameter is the analog value of the sensor.
	// 
	// If the threshold keeps being reached, the callback is triggered periodically
	// with the period as set by SetDebouncePeriod.
func (device *AnalogInBricklet) RegisterAnalogValueReachedCallback(fn func(uint16)) uint64 {
            wrapper := func(byteSlice []byte) {
                buf := bytes.NewBuffer(byteSlice[8:])
                var value uint16
                binary.Read(buf, binary.LittleEndian, &value)
                fn(value)
            }
    return device.device.RegisterCallback(uint8(FunctionCallbackAnalogValueReached), wrapper)
}

//Remove a registered Analog Value Reached callback.
func (device *AnalogInBricklet) DeregisterAnalogValueReachedCallback(callbackID uint64) {
    device.device.DeregisterCallback(uint8(FunctionCallbackAnalogValueReached), callbackID)
}


// Returns the voltage of the sensor. The value is in mV and
	// between 0V and 45V. The resolution between 0 and 6V is about 2mV.
	// Between 6 and 45V the resolution is about 10mV.
	// 
	// If you want to get the voltage periodically, it is recommended to use the
	// RegisterVoltageCallback callback and set the period with
	// SetVoltageCallbackPeriod.
func (device *AnalogInBricklet) GetVoltage() (voltage uint16, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetVoltage), buf.Bytes())
    if err != nil {
        return voltage, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return voltage, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &voltage)

    }
    
    return voltage, nil
}

// Returns the value as read by a 12-bit analog-to-digital converter.
	// The value is between 0 and 4095.
	// 
	// Note
	//  The value returned by GetVoltage is averaged over several samples
	//  to yield less noise, while GetAnalogValue gives back raw
	//  unfiltered analog values. The only reason to use GetAnalogValue is,
	//  if you need the full resolution of the analog-to-digital converter.
	// 
	// If you want the analog value periodically, it is recommended to use the
	// RegisterAnalogValueCallback callback and set the period with
	// SetAnalogValueCallbackPeriod.
func (device *AnalogInBricklet) GetAnalogValue() (value uint16, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetAnalogValue), buf.Bytes())
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

// Sets the period in ms with which the RegisterVoltageCallback callback is triggered
	// periodically. A value of 0 turns the callback off.
	// 
	// The RegisterVoltageCallback callback is only triggered if the voltage has changed since
	// the last triggering.
	// 
	// The default value is 0.
func (device *AnalogInBricklet) SetVoltageCallbackPeriod(period uint32) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, period);

    resultBytes, err := device.device.Set(uint8(FunctionSetVoltageCallbackPeriod), buf.Bytes())
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

// Returns the period as set by SetVoltageCallbackPeriod.
func (device *AnalogInBricklet) GetVoltageCallbackPeriod() (period uint32, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetVoltageCallbackPeriod), buf.Bytes())
    if err != nil {
        return period, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return period, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &period)

    }
    
    return period, nil
}

// Sets the period in ms with which the RegisterAnalogValueCallback callback is triggered
	// periodically. A value of 0 turns the callback off.
	// 
	// The RegisterAnalogValueCallback callback is only triggered if the analog value has
	// changed since the last triggering.
	// 
	// The default value is 0.
func (device *AnalogInBricklet) SetAnalogValueCallbackPeriod(period uint32) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, period);

    resultBytes, err := device.device.Set(uint8(FunctionSetAnalogValueCallbackPeriod), buf.Bytes())
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

// Returns the period as set by SetAnalogValueCallbackPeriod.
func (device *AnalogInBricklet) GetAnalogValueCallbackPeriod() (period uint32, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetAnalogValueCallbackPeriod), buf.Bytes())
    if err != nil {
        return period, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return period, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &period)

    }
    
    return period, nil
}

// Sets the thresholds for the RegisterVoltageReachedCallback callback.
	// 
	// The following options are possible:
	// 
	//  Option| Description
	//  --- | --- 
	//  'x'|    Callback is turned off
	//  'o'|    Callback is triggered when the voltage is *outside* the min and max values
	//  'i'|    Callback is triggered when the voltage is *inside* the min and max values
	//  '<'|    Callback is triggered when the voltage is smaller than the min value (max is ignored)
	//  '>'|    Callback is triggered when the voltage is greater than the min value (max is ignored)
	// 
	// The default value is ('x', 0, 0).
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *AnalogInBricklet) SetVoltageCallbackThreshold(option ThresholdOption, min uint16, max uint16) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, option);
	binary.Write(&buf, binary.LittleEndian, min);
	binary.Write(&buf, binary.LittleEndian, max);

    resultBytes, err := device.device.Set(uint8(FunctionSetVoltageCallbackThreshold), buf.Bytes())
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

// Returns the threshold as set by SetVoltageCallbackThreshold.
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *AnalogInBricklet) GetVoltageCallbackThreshold() (option ThresholdOption, min uint16, max uint16, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetVoltageCallbackThreshold), buf.Bytes())
    if err != nil {
        return option, min, max, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return option, min, max, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &option)
	binary.Read(resultBuf, binary.LittleEndian, &min)
	binary.Read(resultBuf, binary.LittleEndian, &max)

    }
    
    return option, min, max, nil
}

// Sets the thresholds for the RegisterAnalogValueReachedCallback callback.
	// 
	// The following options are possible:
	// 
	//  Option| Description
	//  --- | --- 
	//  'x'|    Callback is turned off
	//  'o'|    Callback is triggered when the analog value is *outside* the min and max values
	//  'i'|    Callback is triggered when the analog value is *inside* the min and max values
	//  '<'|    Callback is triggered when the analog value is smaller than the min value (max is ignored)
	//  '>'|    Callback is triggered when the analog value is greater than the min value (max is ignored)
	// 
	// The default value is ('x', 0, 0).
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *AnalogInBricklet) SetAnalogValueCallbackThreshold(option ThresholdOption, min uint16, max uint16) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, option);
	binary.Write(&buf, binary.LittleEndian, min);
	binary.Write(&buf, binary.LittleEndian, max);

    resultBytes, err := device.device.Set(uint8(FunctionSetAnalogValueCallbackThreshold), buf.Bytes())
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

// Returns the threshold as set by SetAnalogValueCallbackThreshold.
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *AnalogInBricklet) GetAnalogValueCallbackThreshold() (option ThresholdOption, min uint16, max uint16, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetAnalogValueCallbackThreshold), buf.Bytes())
    if err != nil {
        return option, min, max, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return option, min, max, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &option)
	binary.Read(resultBuf, binary.LittleEndian, &min)
	binary.Read(resultBuf, binary.LittleEndian, &max)

    }
    
    return option, min, max, nil
}

// Sets the period in ms with which the threshold callbacks
	// 
	// * RegisterVoltageReachedCallback,
	// * RegisterAnalogValueReachedCallback
	// 
	// are triggered, if the thresholds
	// 
	// * SetVoltageCallbackThreshold,
	// * SetAnalogValueCallbackThreshold
	// 
	// keep being reached.
	// 
	// The default value is 100.
func (device *AnalogInBricklet) SetDebouncePeriod(debounce uint32) (err error) {    
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
func (device *AnalogInBricklet) GetDebouncePeriod() (debounce uint32, err error) {    
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

// Sets the measurement range. Possible ranges:
	// 
	// * 0: Automatically switched
	// * 1: 0V - 6.05V, ~1.48mV resolution
	// * 2: 0V - 10.32V, ~2.52mV resolution
	// * 3: 0V - 36.30V, ~8.86mV resolution
	// * 4: 0V - 45.00V, ~11.25mV resolution
	// * 5: 0V - 3.3V, ~0.81mV resolution, new in version 2.0.3$nbsp;(Plugin)
	// 
	// The default measurement range is 0.
	// 
	// .. versionadded:: 2.0.1$nbsp;(Plugin)
//
// Associated constants:
//
//	* RangeAutomatic
//	* RangeUpTo6V
//	* RangeUpTo10V
//	* RangeUpTo36V
//	* RangeUpTo45V
//	* RangeUpTo3V
func (device *AnalogInBricklet) SetRange(range_ Range) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, range_);

    resultBytes, err := device.device.Set(uint8(FunctionSetRange), buf.Bytes())
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

// Returns the measurement range as set by SetRange.
	// 
	// .. versionadded:: 2.0.1$nbsp;(Plugin)
//
// Associated constants:
//
//	* RangeAutomatic
//	* RangeUpTo6V
//	* RangeUpTo10V
//	* RangeUpTo36V
//	* RangeUpTo45V
//	* RangeUpTo3V
func (device *AnalogInBricklet) GetRange() (range_ Range, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetRange), buf.Bytes())
    if err != nil {
        return range_, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return range_, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &range_)

    }
    
    return range_, nil
}

// Set the length of a averaging for the voltage value.
	// 
	// Setting the length to 0 will turn the averaging completely off. If the
	// averaging is off, there is more noise on the data, but the data is without
	// delay.
	// 
	// The default value is 50.
	// 
	// .. versionadded:: 2.0.3$nbsp;(Plugin)
func (device *AnalogInBricklet) SetAveraging(average uint8) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, average);

    resultBytes, err := device.device.Set(uint8(FunctionSetAveraging), buf.Bytes())
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

// Returns the averaging configuration as set by SetAveraging.
	// 
	// .. versionadded:: 2.0.3$nbsp;(Plugin)
func (device *AnalogInBricklet) GetAveraging() (average uint8, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetAveraging), buf.Bytes())
    if err != nil {
        return average, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return average, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &average)

    }
    
    return average, nil
}

// Returns the UID, the UID where the Bricklet is connected to,
	// the position, the hardware and firmware version as well as the
	// device identifier.
	// 
	// The position can be 'a', 'b', 'c' or 'd'.
	// 
	// The device identifier numbers can be found `here <device_identifier>`.
	// |device_identifier_constant|
func (device *AnalogInBricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {    
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
