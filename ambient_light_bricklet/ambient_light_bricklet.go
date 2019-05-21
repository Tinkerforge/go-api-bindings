/* ***********************************************************
 * This file was automatically generated on 2019-05-21.      *
 *                                                           *
 * Go Bindings Version 2.0.3                                 *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


//Measures ambient light up to 900lux.
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/AmbientLight_Bricklet_Go.html.
package ambient_light_bricklet

import (
	"encoding/binary"
	"bytes"
    . "github.com/Tinkerforge/go-api-bindings/internal"
    "github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function uint8

const (
    FunctionGetIlluminance Function = 1
	FunctionGetAnalogValue Function = 2
	FunctionSetIlluminanceCallbackPeriod Function = 3
	FunctionGetIlluminanceCallbackPeriod Function = 4
	FunctionSetAnalogValueCallbackPeriod Function = 5
	FunctionGetAnalogValueCallbackPeriod Function = 6
	FunctionSetIlluminanceCallbackThreshold Function = 7
	FunctionGetIlluminanceCallbackThreshold Function = 8
	FunctionSetAnalogValueCallbackThreshold Function = 9
	FunctionGetAnalogValueCallbackThreshold Function = 10
	FunctionSetDebouncePeriod Function = 11
	FunctionGetDebouncePeriod Function = 12
	FunctionGetIdentity Function = 255
	FunctionCallbackIlluminance Function = 13
	FunctionCallbackAnalogValue Function = 14
	FunctionCallbackIlluminanceReached Function = 15
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

type AmbientLightBricklet struct{
	device Device
}
const DeviceIdentifier = 21
const DeviceDisplayName = "Ambient Light Bricklet"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (AmbientLightBricklet, error) {
    internalIPCon := ipcon.GetInternalHandle().(IPConnection)
    dev, err := NewDevice([3]uint8{ 2,0,1 }, uid, &internalIPCon, 0)
    if err != nil {
        return AmbientLightBricklet{}, err
    }
    dev.ResponseExpected[FunctionGetIlluminance] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetAnalogValue] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetIlluminanceCallbackPeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetIlluminanceCallbackPeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetAnalogValueCallbackPeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetAnalogValueCallbackPeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetIlluminanceCallbackThreshold] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetIlluminanceCallbackThreshold] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetAnalogValueCallbackThreshold] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetAnalogValueCallbackThreshold] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetDebouncePeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetDebouncePeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetIdentity] = ResponseExpectedFlagAlwaysTrue;
    return AmbientLightBricklet{dev}, nil
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
func (device *AmbientLightBricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *AmbientLightBricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
    return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *AmbientLightBricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *AmbientLightBricklet) GetAPIVersion() [3]uint8 {
    return device.device.GetAPIVersion()
}

// This callback is triggered periodically with the period that is set by
// SetIlluminanceCallbackPeriod. The parameter is the illuminance of the
// ambient light sensor.
// 
// The RegisterIlluminanceCallback callback is only triggered if the illuminance has changed since the
// last triggering.
func (device *AmbientLightBricklet) RegisterIlluminanceCallback(fn func(uint16)) uint64 {
            wrapper := func(byteSlice []byte) {
                buf := bytes.NewBuffer(byteSlice[8:])
                var illuminance uint16
                binary.Read(buf, binary.LittleEndian, &illuminance)
                fn(illuminance)
            }
    return device.device.RegisterCallback(uint8(FunctionCallbackIlluminance), wrapper)
}

//Remove a registered Illuminance callback.
func (device *AmbientLightBricklet) DeregisterIlluminanceCallback(registrationID uint64) {
    device.device.DeregisterCallback(uint8(FunctionCallbackIlluminance), registrationID)
}


// This callback is triggered periodically with the period that is set by
// SetAnalogValueCallbackPeriod. The parameter is the analog value of the
// ambient light sensor.
// 
// The RegisterAnalogValueCallback callback is only triggered if the analog value has changed since the
// last triggering.
func (device *AmbientLightBricklet) RegisterAnalogValueCallback(fn func(uint16)) uint64 {
            wrapper := func(byteSlice []byte) {
                buf := bytes.NewBuffer(byteSlice[8:])
                var value uint16
                binary.Read(buf, binary.LittleEndian, &value)
                fn(value)
            }
    return device.device.RegisterCallback(uint8(FunctionCallbackAnalogValue), wrapper)
}

//Remove a registered Analog Value callback.
func (device *AmbientLightBricklet) DeregisterAnalogValueCallback(registrationID uint64) {
    device.device.DeregisterCallback(uint8(FunctionCallbackAnalogValue), registrationID)
}


// This callback is triggered when the threshold as set by
// SetIlluminanceCallbackThreshold is reached.
// The parameter is the illuminance of the ambient light sensor.
// 
// If the threshold keeps being reached, the callback is triggered periodically
// with the period as set by SetDebouncePeriod.
func (device *AmbientLightBricklet) RegisterIlluminanceReachedCallback(fn func(uint16)) uint64 {
            wrapper := func(byteSlice []byte) {
                buf := bytes.NewBuffer(byteSlice[8:])
                var illuminance uint16
                binary.Read(buf, binary.LittleEndian, &illuminance)
                fn(illuminance)
            }
    return device.device.RegisterCallback(uint8(FunctionCallbackIlluminanceReached), wrapper)
}

//Remove a registered Illuminance Reached callback.
func (device *AmbientLightBricklet) DeregisterIlluminanceReachedCallback(registrationID uint64) {
    device.device.DeregisterCallback(uint8(FunctionCallbackIlluminanceReached), registrationID)
}


// This callback is triggered when the threshold as set by
// SetAnalogValueCallbackThreshold is reached.
// The parameter is the analog value of the ambient light sensor.
// 
// If the threshold keeps being reached, the callback is triggered periodically
// with the period as set by SetDebouncePeriod.
func (device *AmbientLightBricklet) RegisterAnalogValueReachedCallback(fn func(uint16)) uint64 {
            wrapper := func(byteSlice []byte) {
                buf := bytes.NewBuffer(byteSlice[8:])
                var value uint16
                binary.Read(buf, binary.LittleEndian, &value)
                fn(value)
            }
    return device.device.RegisterCallback(uint8(FunctionCallbackAnalogValueReached), wrapper)
}

//Remove a registered Analog Value Reached callback.
func (device *AmbientLightBricklet) DeregisterAnalogValueReachedCallback(registrationID uint64) {
    device.device.DeregisterCallback(uint8(FunctionCallbackAnalogValueReached), registrationID)
}


// Returns the illuminance of the ambient light sensor. The value
// has a range of 0 to 9000 and is given in lux/10, i.e. a value
// of 4500 means that an illuminance of 450lux is measured.
// 
// If you want to get the illuminance periodically, it is recommended to use the
// RegisterIlluminanceCallback callback and set the period with
// SetIlluminanceCallbackPeriod.
func (device *AmbientLightBricklet) GetIlluminance() (illuminance uint16, err error) {
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetIlluminance), buf.Bytes())
    if err != nil {
        return illuminance, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader

        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return illuminance, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &illuminance)

    }

    return illuminance, nil
}

// Returns the value as read by a 12-bit analog-to-digital converter.
// The value is between 0 and 4095.
// 
// Note
//  The value returned by GetIlluminance is averaged over several samples
//  to yield less noise, while GetAnalogValue gives back raw
//  unfiltered analog values. The only reason to use GetAnalogValue is,
//  if you need the full resolution of the analog-to-digital converter.
// 
//  Also, the analog-to-digital converter covers three different ranges that are
//  set dynamically depending on the light intensity. It is impossible to
//  distinguish between these ranges with the analog value.
// 
// If you want the analog value periodically, it is recommended to use the
// RegisterAnalogValueCallback callback and set the period with
// SetAnalogValueCallbackPeriod.
func (device *AmbientLightBricklet) GetAnalogValue() (value uint16, err error) {
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

// Sets the period in ms with which the RegisterIlluminanceCallback callback is triggered
// periodically. A value of 0 turns the callback off.
// 
// The RegisterIlluminanceCallback callback is only triggered if the illuminance has changed since the
// last triggering.
// 
// The default value is 0.
func (device *AmbientLightBricklet) SetIlluminanceCallbackPeriod(period uint32) (err error) {
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, period);

    resultBytes, err := device.device.Set(uint8(FunctionSetIlluminanceCallbackPeriod), buf.Bytes())
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

// Returns the period as set by SetIlluminanceCallbackPeriod.
func (device *AmbientLightBricklet) GetIlluminanceCallbackPeriod() (period uint32, err error) {
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetIlluminanceCallbackPeriod), buf.Bytes())
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
// The RegisterAnalogValueCallback callback is only triggered if the analog value has changed since the
// last triggering.
// 
// The default value is 0.
func (device *AmbientLightBricklet) SetAnalogValueCallbackPeriod(period uint32) (err error) {
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
func (device *AmbientLightBricklet) GetAnalogValueCallbackPeriod() (period uint32, err error) {
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

// Sets the thresholds for the RegisterIlluminanceReachedCallback callback.
// 
// The following options are possible:
// 
//  Option| Description
//  --- | --- 
//  'x'|    Callback is turned off
//  'o'|    Callback is triggered when the illuminance is *outside* the min and max values
//  'i'|    Callback is triggered when the illuminance is *inside* the min and max values
//  '<'|    Callback is triggered when the illuminance is smaller than the min value (max is ignored)
//  '>'|    Callback is triggered when the illuminance is greater than the min value (max is ignored)
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
func (device *AmbientLightBricklet) SetIlluminanceCallbackThreshold(option ThresholdOption, min uint16, max uint16) (err error) {
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, option);
	binary.Write(&buf, binary.LittleEndian, min);
	binary.Write(&buf, binary.LittleEndian, max);

    resultBytes, err := device.device.Set(uint8(FunctionSetIlluminanceCallbackThreshold), buf.Bytes())
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

// Returns the threshold as set by SetIlluminanceCallbackThreshold.
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *AmbientLightBricklet) GetIlluminanceCallbackThreshold() (option ThresholdOption, min uint16, max uint16, err error) {
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetIlluminanceCallbackThreshold), buf.Bytes())
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
func (device *AmbientLightBricklet) SetAnalogValueCallbackThreshold(option ThresholdOption, min uint16, max uint16) (err error) {
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
func (device *AmbientLightBricklet) GetAnalogValueCallbackThreshold() (option ThresholdOption, min uint16, max uint16, err error) {
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
// * RegisterIlluminanceReachedCallback,
// * RegisterAnalogValueReachedCallback
// 
// are triggered, if the thresholds
// 
// * SetIlluminanceCallbackThreshold,
// * SetAnalogValueCallbackThreshold
// 
// keep being reached.
// 
// The default value is 100.
func (device *AmbientLightBricklet) SetDebouncePeriod(debounce uint32) (err error) {
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
func (device *AmbientLightBricklet) GetDebouncePeriod() (debounce uint32, err error) {
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
func (device *AmbientLightBricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
