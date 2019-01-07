/* ***********************************************************
 * This file was automatically generated on 2019-01-07.      *
 *                                                           *
 * Go Bindings Version 2.0.1                                 *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


//Measures UV light.
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/UVLight_Bricklet_Go.html.
package uv_light_bricklet

import (
	"encoding/binary"
	"bytes"
    . "github.com/Tinkerforge/go-api-bindings/internal"
    "github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function uint8

const (
    FunctionGetUVLight Function = 1
	FunctionSetUVLightCallbackPeriod Function = 2
	FunctionGetUVLightCallbackPeriod Function = 3
	FunctionSetUVLightCallbackThreshold Function = 4
	FunctionGetUVLightCallbackThreshold Function = 5
	FunctionSetDebouncePeriod Function = 6
	FunctionGetDebouncePeriod Function = 7
	FunctionGetIdentity Function = 255
	FunctionCallbackUVLight Function = 8
	FunctionCallbackUVLightReached Function = 9
)

type ThresholdOption rune

const (
    ThresholdOptionOff ThresholdOption = 'x'
	ThresholdOptionOutside ThresholdOption = 'o'
	ThresholdOptionInside ThresholdOption = 'i'
	ThresholdOptionSmaller ThresholdOption = '<'
	ThresholdOptionGreater ThresholdOption = '>'
)

type UVLightBricklet struct{
	device Device
}
const DeviceIdentifier = 265
const DeviceDisplayName = "UV Light Bricklet"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (UVLightBricklet, error) {
    internalIPCon := ipcon.GetInternalHandle().(IPConnection)
    dev, err := NewDevice([3]uint8{ 2,0,0 }, uid, &internalIPCon, 0)
    if err != nil {
        return UVLightBricklet{}, err
    }
    dev.ResponseExpected[FunctionGetUVLight] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetUVLightCallbackPeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetUVLightCallbackPeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetUVLightCallbackThreshold] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetUVLightCallbackThreshold] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetDebouncePeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetDebouncePeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetIdentity] = ResponseExpectedFlagAlwaysTrue;
    return UVLightBricklet{dev}, nil
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
func (device *UVLightBricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *UVLightBricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
    return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *UVLightBricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *UVLightBricklet) GetAPIVersion() [3]uint8 {
    return device.device.GetAPIVersion()
}

// This callback is triggered periodically with the period that is set by
	// SetUVLightCallbackPeriod. The parameter is the UV Light
	// intensity of the sensor.
	// 
	// The RegisterUVLightCallback callback is only triggered if the intensity has changed
	// since the last triggering.
func (device *UVLightBricklet) RegisterUVLightCallback(fn func(uint32)) uint64 {
            wrapper := func(byteSlice []byte) {
                buf := bytes.NewBuffer(byteSlice[8:])
                var uvLight uint32
                binary.Read(buf, binary.LittleEndian, &uvLight)
                fn(uvLight)
            }
    return device.device.RegisterCallback(uint8(FunctionCallbackUVLight), wrapper)
}

//Remove a registered UV Light callback.
func (device *UVLightBricklet) DeregisterUVLightCallback(callbackID uint64) {
    device.device.DeregisterCallback(uint8(FunctionCallbackUVLight), callbackID)
}


// This callback is triggered when the threshold as set by
	// SetUVLightCallbackThreshold is reached.
	// The parameter is the UV Light intensity of the sensor.
	// 
	// If the threshold keeps being reached, the callback is triggered periodically
	// with the period as set by SetDebouncePeriod.
func (device *UVLightBricklet) RegisterUVLightReachedCallback(fn func(uint32)) uint64 {
            wrapper := func(byteSlice []byte) {
                buf := bytes.NewBuffer(byteSlice[8:])
                var uvLight uint32
                binary.Read(buf, binary.LittleEndian, &uvLight)
                fn(uvLight)
            }
    return device.device.RegisterCallback(uint8(FunctionCallbackUVLightReached), wrapper)
}

//Remove a registered UV Light Reached callback.
func (device *UVLightBricklet) DeregisterUVLightReachedCallback(callbackID uint64) {
    device.device.DeregisterCallback(uint8(FunctionCallbackUVLightReached), callbackID)
}


// Returns the UV light intensity of the sensor, the intensity is given
	// in 1/10 mW/m². The sensor has already weighted the intensity with the erythemal
	// action spectrum to get the skin-affecting irradiation.
	// 
	// To get UV index you just have to divide the value by 250. For example, a UV
	// light intensity of 500 is equivalent to an UV index of 2.
	// 
	// If you want to get the intensity periodically, it is recommended to use the
	// RegisterUVLightCallback callback and set the period with
	// SetUVLightCallbackPeriod.
func (device *UVLightBricklet) GetUVLight() (uvLight uint32, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetUVLight), buf.Bytes())
    if err != nil {
        return uvLight, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return uvLight, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &uvLight)

    }
    
    return uvLight, nil
}

// Sets the period in ms with which the RegisterUVLightCallback callback is triggered
	// periodically. A value of 0 turns the callback off.
	// 
	// The RegisterUVLightCallback callback is only triggered if the intensity has changed since
	// the last triggering.
	// 
	// The default value is 0.
func (device *UVLightBricklet) SetUVLightCallbackPeriod(period uint32) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, period);

    resultBytes, err := device.device.Set(uint8(FunctionSetUVLightCallbackPeriod), buf.Bytes())
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

// Returns the period as set by SetUVLightCallbackPeriod.
func (device *UVLightBricklet) GetUVLightCallbackPeriod() (period uint32, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetUVLightCallbackPeriod), buf.Bytes())
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

// Sets the thresholds for the RegisterUVLightReachedCallback callback.
	// 
	// The following options are possible:
	// 
	//  Option| Description
	//  --- | --- 
	//  'x'|    Callback is turned off
	//  'o'|    Callback is triggered when the intensity is *outside* the min and max values
	//  'i'|    Callback is triggered when the intensity is *inside* the min and max values
	//  '<'|    Callback is triggered when the intensity is smaller than the min value (max is ignored)
	//  '>'|    Callback is triggered when the intensity is greater than the min value (max is ignored)
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
func (device *UVLightBricklet) SetUVLightCallbackThreshold(option ThresholdOption, min uint32, max uint32) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, option);
	binary.Write(&buf, binary.LittleEndian, min);
	binary.Write(&buf, binary.LittleEndian, max);

    resultBytes, err := device.device.Set(uint8(FunctionSetUVLightCallbackThreshold), buf.Bytes())
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

// Returns the threshold as set by SetUVLightCallbackThreshold.
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *UVLightBricklet) GetUVLightCallbackThreshold() (option ThresholdOption, min uint32, max uint32, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetUVLightCallbackThreshold), buf.Bytes())
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
	// * RegisterUVLightReachedCallback,
	// 
	// are triggered, if the thresholds
	// 
	// * SetUVLightCallbackThreshold,
	// 
	// keep being reached.
	// 
	// The default value is 100.
func (device *UVLightBricklet) SetDebouncePeriod(debounce uint32) (err error) {    
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
func (device *UVLightBricklet) GetDebouncePeriod() (debounce uint32, err error) {    
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
func (device *UVLightBricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {    
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
