/* ***********************************************************
 * This file was automatically generated on 2019-01-29.      *
 *                                                           *
 * Go Bindings Version 2.0.2                                 *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


//Measures AC/DC current between -25A and +25A‍.
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/Current25_Bricklet_Go.html.
package current25_bricklet

import (
	"encoding/binary"
	"bytes"
    . "github.com/Tinkerforge/go-api-bindings/internal"
    "github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function uint8

const (
    FunctionGetCurrent Function = 1
	FunctionCalibrate Function = 2
	FunctionIsOverCurrent Function = 3
	FunctionGetAnalogValue Function = 4
	FunctionSetCurrentCallbackPeriod Function = 5
	FunctionGetCurrentCallbackPeriod Function = 6
	FunctionSetAnalogValueCallbackPeriod Function = 7
	FunctionGetAnalogValueCallbackPeriod Function = 8
	FunctionSetCurrentCallbackThreshold Function = 9
	FunctionGetCurrentCallbackThreshold Function = 10
	FunctionSetAnalogValueCallbackThreshold Function = 11
	FunctionGetAnalogValueCallbackThreshold Function = 12
	FunctionSetDebouncePeriod Function = 13
	FunctionGetDebouncePeriod Function = 14
	FunctionGetIdentity Function = 255
	FunctionCallbackCurrent Function = 15
	FunctionCallbackAnalogValue Function = 16
	FunctionCallbackCurrentReached Function = 17
	FunctionCallbackAnalogValueReached Function = 18
	FunctionCallbackOverCurrent Function = 19
)

type ThresholdOption rune

const (
    ThresholdOptionOff ThresholdOption = 'x'
	ThresholdOptionOutside ThresholdOption = 'o'
	ThresholdOptionInside ThresholdOption = 'i'
	ThresholdOptionSmaller ThresholdOption = '<'
	ThresholdOptionGreater ThresholdOption = '>'
)

type Current25Bricklet struct{
	device Device
}
const DeviceIdentifier = 24
const DeviceDisplayName = "Current25 Bricklet"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (Current25Bricklet, error) {
    internalIPCon := ipcon.GetInternalHandle().(IPConnection)
    dev, err := NewDevice([3]uint8{ 2,0,0 }, uid, &internalIPCon, 0)
    if err != nil {
        return Current25Bricklet{}, err
    }
    dev.ResponseExpected[FunctionGetCurrent] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionCalibrate] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionIsOverCurrent] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetAnalogValue] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetCurrentCallbackPeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetCurrentCallbackPeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetAnalogValueCallbackPeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetAnalogValueCallbackPeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetCurrentCallbackThreshold] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetCurrentCallbackThreshold] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetAnalogValueCallbackThreshold] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetAnalogValueCallbackThreshold] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetDebouncePeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetDebouncePeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetIdentity] = ResponseExpectedFlagAlwaysTrue;
    return Current25Bricklet{dev}, nil
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
func (device *Current25Bricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *Current25Bricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
    return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *Current25Bricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *Current25Bricklet) GetAPIVersion() [3]uint8 {
    return device.device.GetAPIVersion()
}

// This callback is triggered periodically with the period that is set by
	// SetCurrentCallbackPeriod. The parameter is the current of the
	// sensor.
	// 
	// The RegisterCurrentCallback callback is only triggered if the current has changed since the
	// last triggering.
func (device *Current25Bricklet) RegisterCurrentCallback(fn func(int16)) uint64 {
            wrapper := func(byteSlice []byte) {
                buf := bytes.NewBuffer(byteSlice[8:])
                var current int16
                binary.Read(buf, binary.LittleEndian, &current)
                fn(current)
            }
    return device.device.RegisterCallback(uint8(FunctionCallbackCurrent), wrapper)
}

//Remove a registered Current callback.
func (device *Current25Bricklet) DeregisterCurrentCallback(callbackID uint64) {
    device.device.DeregisterCallback(uint8(FunctionCallbackCurrent), callbackID)
}


// This callback is triggered periodically with the period that is set by
	// SetAnalogValueCallbackPeriod. The parameter is the analog value of the
	// sensor.
	// 
	// The RegisterAnalogValueCallback callback is only triggered if the current has changed since the
	// last triggering.
func (device *Current25Bricklet) RegisterAnalogValueCallback(fn func(uint16)) uint64 {
            wrapper := func(byteSlice []byte) {
                buf := bytes.NewBuffer(byteSlice[8:])
                var value uint16
                binary.Read(buf, binary.LittleEndian, &value)
                fn(value)
            }
    return device.device.RegisterCallback(uint8(FunctionCallbackAnalogValue), wrapper)
}

//Remove a registered Analog Value callback.
func (device *Current25Bricklet) DeregisterAnalogValueCallback(callbackID uint64) {
    device.device.DeregisterCallback(uint8(FunctionCallbackAnalogValue), callbackID)
}


// This callback is triggered when the threshold as set by
	// SetCurrentCallbackThreshold is reached.
	// The parameter is the current of the sensor.
	// 
	// If the threshold keeps being reached, the callback is triggered periodically
	// with the period as set by SetDebouncePeriod.
func (device *Current25Bricklet) RegisterCurrentReachedCallback(fn func(int16)) uint64 {
            wrapper := func(byteSlice []byte) {
                buf := bytes.NewBuffer(byteSlice[8:])
                var current int16
                binary.Read(buf, binary.LittleEndian, &current)
                fn(current)
            }
    return device.device.RegisterCallback(uint8(FunctionCallbackCurrentReached), wrapper)
}

//Remove a registered Current Reached callback.
func (device *Current25Bricklet) DeregisterCurrentReachedCallback(callbackID uint64) {
    device.device.DeregisterCallback(uint8(FunctionCallbackCurrentReached), callbackID)
}


// This callback is triggered when the threshold as set by
	// SetAnalogValueCallbackThreshold is reached.
	// The parameter is the analog value of the sensor.
	// 
	// If the threshold keeps being reached, the callback is triggered periodically
	// with the period as set by SetDebouncePeriod.
func (device *Current25Bricklet) RegisterAnalogValueReachedCallback(fn func(uint16)) uint64 {
            wrapper := func(byteSlice []byte) {
                buf := bytes.NewBuffer(byteSlice[8:])
                var value uint16
                binary.Read(buf, binary.LittleEndian, &value)
                fn(value)
            }
    return device.device.RegisterCallback(uint8(FunctionCallbackAnalogValueReached), wrapper)
}

//Remove a registered Analog Value Reached callback.
func (device *Current25Bricklet) DeregisterAnalogValueReachedCallback(callbackID uint64) {
    device.device.DeregisterCallback(uint8(FunctionCallbackAnalogValueReached), callbackID)
}


// This callback is triggered when an over current is measured
	// (see IsOverCurrent).
func (device *Current25Bricklet) RegisterOverCurrentCallback(fn func()) uint64 {
            wrapper := func(byteSlice []byte) {
                
                
                
                fn()
            }
    return device.device.RegisterCallback(uint8(FunctionCallbackOverCurrent), wrapper)
}

//Remove a registered Over Current callback.
func (device *Current25Bricklet) DeregisterOverCurrentCallback(callbackID uint64) {
    device.device.DeregisterCallback(uint8(FunctionCallbackOverCurrent), callbackID)
}


// Returns the current of the sensor. The value is in mA
	// and between -25000mA and 25000mA.
	// 
	// If you want to get the current periodically, it is recommended to use the
	// RegisterCurrentCallback callback and set the period with
	// SetCurrentCallbackPeriod.
func (device *Current25Bricklet) GetCurrent() (current int16, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetCurrent), buf.Bytes())
    if err != nil {
        return current, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return current, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &current)

    }
    
    return current, nil
}

// Calibrates the 0 value of the sensor. You have to call this function
	// when there is no current present.
	// 
	// The zero point of the current sensor
	// is depending on the exact properties of the analog-to-digital converter,
	// the length of the Bricklet cable and the temperature. Thus, if you change
	// the Brick or the environment in which the Bricklet is used, you might
	// have to recalibrate.
	// 
	// The resulting calibration will be saved on the EEPROM of the Current
	// Bricklet.
func (device *Current25Bricklet) Calibrate() (err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Set(uint8(FunctionCalibrate), buf.Bytes())
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

// Returns *true* if more than 25A were measured.
	// 
	// Note
	//  To reset this value you have to power cycle the Bricklet.
func (device *Current25Bricklet) IsOverCurrent() (over bool, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionIsOverCurrent), buf.Bytes())
    if err != nil {
        return over, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return over, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &over)

    }
    
    return over, nil
}

// Returns the value as read by a 12-bit analog-to-digital converter.
	// The value is between 0 and 4095.
	// 
	// Note
	//  The value returned by GetCurrent is averaged over several samples
	//  to yield less noise, while GetAnalogValue gives back raw
	//  unfiltered analog values. The only reason to use GetAnalogValue is,
	//  if you need the full resolution of the analog-to-digital converter.
	// 
	// If you want the analog value periodically, it is recommended to use the
	// RegisterAnalogValueCallback callback and set the period with
	// SetAnalogValueCallbackPeriod.
func (device *Current25Bricklet) GetAnalogValue() (value uint16, err error) {    
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

// Sets the period in ms with which the RegisterCurrentCallback callback is triggered
	// periodically. A value of 0 turns the callback off.
	// 
	// The RegisterCurrentCallback callback is only triggered if the current has changed since
	// the last triggering.
	// 
	// The default value is 0.
func (device *Current25Bricklet) SetCurrentCallbackPeriod(period uint32) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, period);

    resultBytes, err := device.device.Set(uint8(FunctionSetCurrentCallbackPeriod), buf.Bytes())
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

// Returns the period as set by SetCurrentCallbackPeriod.
func (device *Current25Bricklet) GetCurrentCallbackPeriod() (period uint32, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetCurrentCallbackPeriod), buf.Bytes())
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
func (device *Current25Bricklet) SetAnalogValueCallbackPeriod(period uint32) (err error) {    
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
func (device *Current25Bricklet) GetAnalogValueCallbackPeriod() (period uint32, err error) {    
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

// Sets the thresholds for the RegisterCurrentReachedCallback callback.
	// 
	// The following options are possible:
	// 
	//  Option| Description
	//  --- | --- 
	//  'x'|    Callback is turned off
	//  'o'|    Callback is triggered when the current is *outside* the min and max values
	//  'i'|    Callback is triggered when the current is *inside* the min and max values
	//  '<'|    Callback is triggered when the current is smaller than the min value (max is ignored)
	//  '>'|    Callback is triggered when the current is greater than the min value (max is ignored)
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
func (device *Current25Bricklet) SetCurrentCallbackThreshold(option ThresholdOption, min int16, max int16) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, option);
	binary.Write(&buf, binary.LittleEndian, min);
	binary.Write(&buf, binary.LittleEndian, max);

    resultBytes, err := device.device.Set(uint8(FunctionSetCurrentCallbackThreshold), buf.Bytes())
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

// Returns the threshold as set by SetCurrentCallbackThreshold.
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *Current25Bricklet) GetCurrentCallbackThreshold() (option ThresholdOption, min int16, max int16, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetCurrentCallbackThreshold), buf.Bytes())
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
func (device *Current25Bricklet) SetAnalogValueCallbackThreshold(option ThresholdOption, min uint16, max uint16) (err error) {    
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
func (device *Current25Bricklet) GetAnalogValueCallbackThreshold() (option ThresholdOption, min uint16, max uint16, err error) {    
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
	// * RegisterCurrentReachedCallback,
	// * RegisterAnalogValueReachedCallback
	// 
	// are triggered, if the thresholds
	// 
	// * SetCurrentCallbackThreshold,
	// * SetAnalogValueCallbackThreshold
	// 
	// keep being reached.
	// 
	// The default value is 100.
func (device *Current25Bricklet) SetDebouncePeriod(debounce uint32) (err error) {    
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
func (device *Current25Bricklet) GetDebouncePeriod() (debounce uint32, err error) {    
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

// Returns the UID, the UID where the Bricklet is connected to,
	// the position, the hardware and firmware version as well as the
	// device identifier.
	// 
	// The position can be 'a', 'b', 'c' or 'd'.
	// 
	// The device identifier numbers can be found `here <device_identifier>`.
	// |device_identifier_constant|
func (device *Current25Bricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {    
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
