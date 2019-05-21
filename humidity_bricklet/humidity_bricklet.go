/* ***********************************************************
 * This file was automatically generated on 2019-05-21.      *
 *                                                           *
 * Go Bindings Version 2.0.3                                 *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


//Measures relative humidity.
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/Humidity_Bricklet_Go.html.
package humidity_bricklet

import (
	"encoding/binary"
	"bytes"
    . "github.com/Tinkerforge/go-api-bindings/internal"
    "github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function uint8

const (
    FunctionGetHumidity Function = 1
	FunctionGetAnalogValue Function = 2
	FunctionSetHumidityCallbackPeriod Function = 3
	FunctionGetHumidityCallbackPeriod Function = 4
	FunctionSetAnalogValueCallbackPeriod Function = 5
	FunctionGetAnalogValueCallbackPeriod Function = 6
	FunctionSetHumidityCallbackThreshold Function = 7
	FunctionGetHumidityCallbackThreshold Function = 8
	FunctionSetAnalogValueCallbackThreshold Function = 9
	FunctionGetAnalogValueCallbackThreshold Function = 10
	FunctionSetDebouncePeriod Function = 11
	FunctionGetDebouncePeriod Function = 12
	FunctionGetIdentity Function = 255
	FunctionCallbackHumidity Function = 13
	FunctionCallbackAnalogValue Function = 14
	FunctionCallbackHumidityReached Function = 15
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

type HumidityBricklet struct{
	device Device
}
const DeviceIdentifier = 27
const DeviceDisplayName = "Humidity Bricklet"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (HumidityBricklet, error) {
    internalIPCon := ipcon.GetInternalHandle().(IPConnection)
    dev, err := NewDevice([3]uint8{ 2,0,1 }, uid, &internalIPCon, 0)
    if err != nil {
        return HumidityBricklet{}, err
    }
    dev.ResponseExpected[FunctionGetHumidity] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetAnalogValue] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetHumidityCallbackPeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetHumidityCallbackPeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetAnalogValueCallbackPeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetAnalogValueCallbackPeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetHumidityCallbackThreshold] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetHumidityCallbackThreshold] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetAnalogValueCallbackThreshold] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetAnalogValueCallbackThreshold] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetDebouncePeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetDebouncePeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetIdentity] = ResponseExpectedFlagAlwaysTrue;
    return HumidityBricklet{dev}, nil
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
func (device *HumidityBricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *HumidityBricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
    return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *HumidityBricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *HumidityBricklet) GetAPIVersion() [3]uint8 {
    return device.device.GetAPIVersion()
}

// This callback is triggered periodically with the period that is set by
// SetHumidityCallbackPeriod. The parameter is the humidity of
// the sensor.
// 
// The RegisterHumidityCallback callback is only triggered if the humidity has changed since
// the last triggering.
func (device *HumidityBricklet) RegisterHumidityCallback(fn func(uint16)) uint64 {
            wrapper := func(byteSlice []byte) {
                buf := bytes.NewBuffer(byteSlice[8:])
                var humidity uint16
                binary.Read(buf, binary.LittleEndian, &humidity)
                fn(humidity)
            }
    return device.device.RegisterCallback(uint8(FunctionCallbackHumidity), wrapper)
}

//Remove a registered Humidity callback.
func (device *HumidityBricklet) DeregisterHumidityCallback(registrationID uint64) {
    device.device.DeregisterCallback(uint8(FunctionCallbackHumidity), registrationID)
}


// This callback is triggered periodically with the period that is set by
// SetAnalogValueCallbackPeriod. The parameter is the analog
// value of the sensor.
// 
// The RegisterAnalogValueCallback callback is only triggered if the humidity has changed
// since the last triggering.
func (device *HumidityBricklet) RegisterAnalogValueCallback(fn func(uint16)) uint64 {
            wrapper := func(byteSlice []byte) {
                buf := bytes.NewBuffer(byteSlice[8:])
                var value uint16
                binary.Read(buf, binary.LittleEndian, &value)
                fn(value)
            }
    return device.device.RegisterCallback(uint8(FunctionCallbackAnalogValue), wrapper)
}

//Remove a registered Analog Value callback.
func (device *HumidityBricklet) DeregisterAnalogValueCallback(registrationID uint64) {
    device.device.DeregisterCallback(uint8(FunctionCallbackAnalogValue), registrationID)
}


// This callback is triggered when the threshold as set by
// SetHumidityCallbackThreshold is reached.
// The parameter is the humidity of the sensor.
// 
// If the threshold keeps being reached, the callback is triggered periodically
// with the period as set by SetDebouncePeriod.
func (device *HumidityBricklet) RegisterHumidityReachedCallback(fn func(uint16)) uint64 {
            wrapper := func(byteSlice []byte) {
                buf := bytes.NewBuffer(byteSlice[8:])
                var humidity uint16
                binary.Read(buf, binary.LittleEndian, &humidity)
                fn(humidity)
            }
    return device.device.RegisterCallback(uint8(FunctionCallbackHumidityReached), wrapper)
}

//Remove a registered Humidity Reached callback.
func (device *HumidityBricklet) DeregisterHumidityReachedCallback(registrationID uint64) {
    device.device.DeregisterCallback(uint8(FunctionCallbackHumidityReached), registrationID)
}


// This callback is triggered when the threshold as set by
// SetAnalogValueCallbackThreshold is reached.
// The parameter is the analog value of the sensor.
// 
// If the threshold keeps being reached, the callback is triggered periodically
// with the period as set by SetDebouncePeriod.
func (device *HumidityBricklet) RegisterAnalogValueReachedCallback(fn func(uint16)) uint64 {
            wrapper := func(byteSlice []byte) {
                buf := bytes.NewBuffer(byteSlice[8:])
                var value uint16
                binary.Read(buf, binary.LittleEndian, &value)
                fn(value)
            }
    return device.device.RegisterCallback(uint8(FunctionCallbackAnalogValueReached), wrapper)
}

//Remove a registered Analog Value Reached callback.
func (device *HumidityBricklet) DeregisterAnalogValueReachedCallback(registrationID uint64) {
    device.device.DeregisterCallback(uint8(FunctionCallbackAnalogValueReached), registrationID)
}


// Returns the humidity of the sensor. The value
// has a range of 0 to 1000 and is given in %RH/10 (Relative Humidity),
// i.e. a value of 421 means that a humidity of 42.1 %RH is measured.
// 
// If you want to get the humidity periodically, it is recommended to use the
// RegisterHumidityCallback callback and set the period with
// SetHumidityCallbackPeriod.
func (device *HumidityBricklet) GetHumidity() (humidity uint16, err error) {
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetHumidity), buf.Bytes())
    if err != nil {
        return humidity, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader

        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return humidity, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &humidity)

    }

    return humidity, nil
}

// Returns the value as read by a 12-bit analog-to-digital converter.
// The value is between 0 and 4095.
// 
// Note
//  The value returned by :func:https://github.com/Tinkerforge/humidity-bricklet/raw/master/datasheets/hih-5030.pdf.
// 
// If you want the analog value periodically, it is recommended to use the
// RegisterAnalogValueCallback callback and set the period with
// SetAnalogValueCallbackPeriod.
func (device *HumidityBricklet) GetAnalogValue() (value uint16, err error) {
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

// Sets the period in ms with which the RegisterHumidityCallback callback is triggered
// periodically. A value of 0 turns the callback off.
// 
// The RegisterHumidityCallback callback is only triggered if the humidity has changed
// since the last triggering.
// 
// The default value is 0.
func (device *HumidityBricklet) SetHumidityCallbackPeriod(period uint32) (err error) {
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, period);

    resultBytes, err := device.device.Set(uint8(FunctionSetHumidityCallbackPeriod), buf.Bytes())
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

// Returns the period as set by SetHumidityCallbackPeriod.
func (device *HumidityBricklet) GetHumidityCallbackPeriod() (period uint32, err error) {
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetHumidityCallbackPeriod), buf.Bytes())
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
func (device *HumidityBricklet) SetAnalogValueCallbackPeriod(period uint32) (err error) {
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
func (device *HumidityBricklet) GetAnalogValueCallbackPeriod() (period uint32, err error) {
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

// Sets the thresholds for the RegisterHumidityReachedCallback callback.
// 
// The following options are possible:
// 
//  Option| Description
//  --- | --- 
//  'x'|    Callback is turned off
//  'o'|    Callback is triggered when the humidity is *outside* the min and max values
//  'i'|    Callback is triggered when the humidity is *inside* the min and max values
//  '<'|    Callback is triggered when the humidity is smaller than the min value (max is ignored)
//  '>'|    Callback is triggered when the humidity is greater than the min value (max is ignored)
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
func (device *HumidityBricklet) SetHumidityCallbackThreshold(option ThresholdOption, min uint16, max uint16) (err error) {
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, option);
	binary.Write(&buf, binary.LittleEndian, min);
	binary.Write(&buf, binary.LittleEndian, max);

    resultBytes, err := device.device.Set(uint8(FunctionSetHumidityCallbackThreshold), buf.Bytes())
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

// Returns the threshold as set by SetHumidityCallbackThreshold.
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *HumidityBricklet) GetHumidityCallbackThreshold() (option ThresholdOption, min uint16, max uint16, err error) {
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetHumidityCallbackThreshold), buf.Bytes())
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
func (device *HumidityBricklet) SetAnalogValueCallbackThreshold(option ThresholdOption, min uint16, max uint16) (err error) {
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
func (device *HumidityBricklet) GetAnalogValueCallbackThreshold() (option ThresholdOption, min uint16, max uint16, err error) {
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
// * RegisterHumidityReachedCallback,
// * RegisterAnalogValueReachedCallback
// 
// are triggered, if the thresholds
// 
// * SetHumidityCallbackThreshold,
// * SetAnalogValueCallbackThreshold
// 
// keep being reached.
// 
// The default value is 100.
func (device *HumidityBricklet) SetDebouncePeriod(debounce uint32) (err error) {
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
func (device *HumidityBricklet) GetDebouncePeriod() (debounce uint32, err error) {
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
func (device *HumidityBricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
