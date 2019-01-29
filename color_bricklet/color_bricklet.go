/* ***********************************************************
 * This file was automatically generated on 2019-01-29.      *
 *                                                           *
 * Go Bindings Version 2.0.2                                 *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


//Measures color (RGB value), illuminance and color temperature.
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/Color_Bricklet_Go.html.
package color_bricklet

import (
	"encoding/binary"
	"bytes"
    . "github.com/Tinkerforge/go-api-bindings/internal"
    "github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function uint8

const (
    FunctionGetColor Function = 1
	FunctionSetColorCallbackPeriod Function = 2
	FunctionGetColorCallbackPeriod Function = 3
	FunctionSetColorCallbackThreshold Function = 4
	FunctionGetColorCallbackThreshold Function = 5
	FunctionSetDebouncePeriod Function = 6
	FunctionGetDebouncePeriod Function = 7
	FunctionLightOn Function = 10
	FunctionLightOff Function = 11
	FunctionIsLightOn Function = 12
	FunctionSetConfig Function = 13
	FunctionGetConfig Function = 14
	FunctionGetIlluminance Function = 15
	FunctionGetColorTemperature Function = 16
	FunctionSetIlluminanceCallbackPeriod Function = 17
	FunctionGetIlluminanceCallbackPeriod Function = 18
	FunctionSetColorTemperatureCallbackPeriod Function = 19
	FunctionGetColorTemperatureCallbackPeriod Function = 20
	FunctionGetIdentity Function = 255
	FunctionCallbackColor Function = 8
	FunctionCallbackColorReached Function = 9
	FunctionCallbackIlluminance Function = 21
	FunctionCallbackColorTemperature Function = 22
)

type ThresholdOption rune

const (
    ThresholdOptionOff ThresholdOption = 'x'
	ThresholdOptionOutside ThresholdOption = 'o'
	ThresholdOptionInside ThresholdOption = 'i'
	ThresholdOptionSmaller ThresholdOption = '<'
	ThresholdOptionGreater ThresholdOption = '>'
)

type Light uint8

const (
    LightOn Light = 0
	LightOff Light = 1
)

type Gain uint8

const (
    Gain1x Gain = 0
	Gain4x Gain = 1
	Gain16x Gain = 2
	Gain60x Gain = 3
)

type IntegrationTime uint8

const (
    IntegrationTime2ms IntegrationTime = 0
	IntegrationTime24ms IntegrationTime = 1
	IntegrationTime101ms IntegrationTime = 2
	IntegrationTime154ms IntegrationTime = 3
	IntegrationTime700ms IntegrationTime = 4
)

type ColorBricklet struct{
	device Device
}
const DeviceIdentifier = 243
const DeviceDisplayName = "Color Bricklet"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (ColorBricklet, error) {
    internalIPCon := ipcon.GetInternalHandle().(IPConnection)
    dev, err := NewDevice([3]uint8{ 2,0,0 }, uid, &internalIPCon, 0)
    if err != nil {
        return ColorBricklet{}, err
    }
    dev.ResponseExpected[FunctionGetColor] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetColorCallbackPeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetColorCallbackPeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetColorCallbackThreshold] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetColorCallbackThreshold] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetDebouncePeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetDebouncePeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionLightOn] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionLightOff] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionIsLightOn] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetConfig] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetConfig] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetIlluminance] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetColorTemperature] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetIlluminanceCallbackPeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetIlluminanceCallbackPeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetColorTemperatureCallbackPeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetColorTemperatureCallbackPeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetIdentity] = ResponseExpectedFlagAlwaysTrue;
    return ColorBricklet{dev}, nil
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
func (device *ColorBricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *ColorBricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
    return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *ColorBricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *ColorBricklet) GetAPIVersion() [3]uint8 {
    return device.device.GetAPIVersion()
}

// This callback is triggered periodically with the period that is set by
	// SetColorCallbackPeriod. The parameter is the color
	// of the sensor as RGBC.
	// 
	// The RegisterColorCallback callback is only triggered if the color has changed since the
	// last triggering.
func (device *ColorBricklet) RegisterColorCallback(fn func(uint16, uint16, uint16, uint16)) uint64 {
            wrapper := func(byteSlice []byte) {
                buf := bytes.NewBuffer(byteSlice[8:])
                var r uint16
var g uint16
var b uint16
var c uint16
                binary.Read(buf, binary.LittleEndian, &r)
binary.Read(buf, binary.LittleEndian, &g)
binary.Read(buf, binary.LittleEndian, &b)
binary.Read(buf, binary.LittleEndian, &c)
                fn(r, g, b, c)
            }
    return device.device.RegisterCallback(uint8(FunctionCallbackColor), wrapper)
}

//Remove a registered Color callback.
func (device *ColorBricklet) DeregisterColorCallback(callbackID uint64) {
    device.device.DeregisterCallback(uint8(FunctionCallbackColor), callbackID)
}


// This callback is triggered when the threshold as set by
	// SetColorCallbackThreshold is reached.
	// The parameter is the color
	// of the sensor as RGBC.
	// 
	// If the threshold keeps being reached, the callback is triggered periodically
	// with the period as set by SetDebouncePeriod.
func (device *ColorBricklet) RegisterColorReachedCallback(fn func(uint16, uint16, uint16, uint16)) uint64 {
            wrapper := func(byteSlice []byte) {
                buf := bytes.NewBuffer(byteSlice[8:])
                var r uint16
var g uint16
var b uint16
var c uint16
                binary.Read(buf, binary.LittleEndian, &r)
binary.Read(buf, binary.LittleEndian, &g)
binary.Read(buf, binary.LittleEndian, &b)
binary.Read(buf, binary.LittleEndian, &c)
                fn(r, g, b, c)
            }
    return device.device.RegisterCallback(uint8(FunctionCallbackColorReached), wrapper)
}

//Remove a registered Color Reached callback.
func (device *ColorBricklet) DeregisterColorReachedCallback(callbackID uint64) {
    device.device.DeregisterCallback(uint8(FunctionCallbackColorReached), callbackID)
}


// This callback is triggered periodically with the period that is set by
	// SetIlluminanceCallbackPeriod. The parameter is the illuminance.
	// See GetIlluminance for how to interpret this value.
	// 
	// The RegisterIlluminanceCallback callback is only triggered if the illuminance has changed
	// since the last triggering.
func (device *ColorBricklet) RegisterIlluminanceCallback(fn func(uint32)) uint64 {
            wrapper := func(byteSlice []byte) {
                buf := bytes.NewBuffer(byteSlice[8:])
                var illuminance uint32
                binary.Read(buf, binary.LittleEndian, &illuminance)
                fn(illuminance)
            }
    return device.device.RegisterCallback(uint8(FunctionCallbackIlluminance), wrapper)
}

//Remove a registered Illuminance callback.
func (device *ColorBricklet) DeregisterIlluminanceCallback(callbackID uint64) {
    device.device.DeregisterCallback(uint8(FunctionCallbackIlluminance), callbackID)
}


// This callback is triggered periodically with the period that is set by
	// SetColorTemperatureCallbackPeriod. The parameter is the
	// color temperature in Kelvin.
	// 
	// The RegisterColorTemperatureCallback callback is only triggered if the color temperature
	// has changed since the last triggering.
func (device *ColorBricklet) RegisterColorTemperatureCallback(fn func(uint16)) uint64 {
            wrapper := func(byteSlice []byte) {
                buf := bytes.NewBuffer(byteSlice[8:])
                var colorTemperature uint16
                binary.Read(buf, binary.LittleEndian, &colorTemperature)
                fn(colorTemperature)
            }
    return device.device.RegisterCallback(uint8(FunctionCallbackColorTemperature), wrapper)
}

//Remove a registered Color Temperature callback.
func (device *ColorBricklet) DeregisterColorTemperatureCallback(callbackID uint64) {
    device.device.DeregisterCallback(uint8(FunctionCallbackColorTemperature), callbackID)
}


// Returns the measured color of the sensor. The values
	// have a range of 0 to 65535.
	// 
	// The red (r), green (g), blue (b) and clear (c) colors are measured
	// with four different photodiodes that are responsive at different
	// wavelengths:
	// 
	// .. image:: /Images/Bricklets/bricklet_color_wavelength_chart_600.jpg
	//    :scale: 100 %
	//    :alt: Chart Responsivity / Wavelength
	//    :align: center
	//    :target: ../../_images/Bricklets/bricklet_color_wavelength_chart_600.jpg
	// 
	// If you want to get the color periodically, it is recommended
	// to use the RegisterColorCallback callback and set the period with
	// SetColorCallbackPeriod.
func (device *ColorBricklet) GetColor() (r uint16, g uint16, b uint16, c uint16, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetColor), buf.Bytes())
    if err != nil {
        return r, g, b, c, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return r, g, b, c, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &r)
	binary.Read(resultBuf, binary.LittleEndian, &g)
	binary.Read(resultBuf, binary.LittleEndian, &b)
	binary.Read(resultBuf, binary.LittleEndian, &c)

    }
    
    return r, g, b, c, nil
}

// Sets the period in ms with which the RegisterColorCallback callback is triggered
	// periodically. A value of 0 turns the callback off.
	// 
	// The RegisterColorCallback callback is only triggered if the color has changed since the
	// last triggering.
	// 
	// The default value is 0.
func (device *ColorBricklet) SetColorCallbackPeriod(period uint32) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, period);

    resultBytes, err := device.device.Set(uint8(FunctionSetColorCallbackPeriod), buf.Bytes())
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

// Returns the period as set by SetColorCallbackPeriod.
func (device *ColorBricklet) GetColorCallbackPeriod() (period uint32, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetColorCallbackPeriod), buf.Bytes())
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

// Sets the thresholds for the RegisterColorReachedCallback callback.
	// 
	// The following options are possible:
	// 
	//  Option| Description
	//  --- | --- 
	//  'x'|    Callback is turned off
	//  'o'|    Callback is triggered when the temperature is *outside* the min and max values
	//  'i'|    Callback is triggered when the temperature is *inside* the min and max values
	//  '<'|    Callback is triggered when the temperature is smaller than the min value (max is ignored)
	//  '>'|    Callback is triggered when the temperature is greater than the min value (max is ignored)
	// 
	// The default value is ('x', 0, 0, 0, 0, 0, 0, 0, 0).
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *ColorBricklet) SetColorCallbackThreshold(option ThresholdOption, minR uint16, maxR uint16, minG uint16, maxG uint16, minB uint16, maxB uint16, minC uint16, maxC uint16) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, option);
	binary.Write(&buf, binary.LittleEndian, minR);
	binary.Write(&buf, binary.LittleEndian, maxR);
	binary.Write(&buf, binary.LittleEndian, minG);
	binary.Write(&buf, binary.LittleEndian, maxG);
	binary.Write(&buf, binary.LittleEndian, minB);
	binary.Write(&buf, binary.LittleEndian, maxB);
	binary.Write(&buf, binary.LittleEndian, minC);
	binary.Write(&buf, binary.LittleEndian, maxC);

    resultBytes, err := device.device.Set(uint8(FunctionSetColorCallbackThreshold), buf.Bytes())
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

// Returns the threshold as set by SetColorCallbackThreshold.
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *ColorBricklet) GetColorCallbackThreshold() (option ThresholdOption, minR uint16, maxR uint16, minG uint16, maxG uint16, minB uint16, maxB uint16, minC uint16, maxC uint16, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetColorCallbackThreshold), buf.Bytes())
    if err != nil {
        return option, minR, maxR, minG, maxG, minB, maxB, minC, maxC, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return option, minR, maxR, minG, maxG, minB, maxB, minC, maxC, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &option)
	binary.Read(resultBuf, binary.LittleEndian, &minR)
	binary.Read(resultBuf, binary.LittleEndian, &maxR)
	binary.Read(resultBuf, binary.LittleEndian, &minG)
	binary.Read(resultBuf, binary.LittleEndian, &maxG)
	binary.Read(resultBuf, binary.LittleEndian, &minB)
	binary.Read(resultBuf, binary.LittleEndian, &maxB)
	binary.Read(resultBuf, binary.LittleEndian, &minC)
	binary.Read(resultBuf, binary.LittleEndian, &maxC)

    }
    
    return option, minR, maxR, minG, maxG, minB, maxB, minC, maxC, nil
}

// Sets the period in ms with which the threshold callback
	// 
	// * RegisterColorReachedCallback
	// 
	// is triggered, if the threshold
	// 
	// * SetColorCallbackThreshold
	// 
	// keeps being reached.
	// 
	// The default value is 100.
func (device *ColorBricklet) SetDebouncePeriod(debounce uint32) (err error) {    
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
func (device *ColorBricklet) GetDebouncePeriod() (debounce uint32, err error) {    
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

// Turns the LED on.
func (device *ColorBricklet) LightOn() (err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Set(uint8(FunctionLightOn), buf.Bytes())
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

// Turns the LED off.
func (device *ColorBricklet) LightOff() (err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Set(uint8(FunctionLightOff), buf.Bytes())
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

// Returns the state of the LED. Possible values are:
	// 
	// * 0: On
	// * 1: Off
//
// Associated constants:
//
//	* LightOn
//	* LightOff
func (device *ColorBricklet) IsLightOn() (light Light, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionIsLightOn), buf.Bytes())
    if err != nil {
        return light, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return light, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &light)

    }
    
    return light, nil
}

// Sets the configuration of the sensor. Gain and integration time
	// can be configured in this way.
	// 
	// For configuring the gain:
	// 
	// * 0: 1x Gain
	// * 1: 4x Gain
	// * 2: 16x Gain
	// * 3: 60x Gain
	// 
	// For configuring the integration time:
	// 
	// * 0: 2.4ms
	// * 1: 24ms
	// * 2: 101ms
	// * 3: 154ms
	// * 4: 700ms
	// 
	// Increasing the gain enables the sensor to detect a
	// color from a higher distance.
	// 
	// The integration time provides a trade-off between conversion time
	// and accuracy. With a longer integration time the values read will
	// be more accurate but it will take longer time to get the conversion
	// results.
	// 
	// The default values are 60x gain and 154ms integration time.
//
// Associated constants:
//
//	* Gain1x
//	* Gain4x
//	* Gain16x
//	* Gain60x
//	* IntegrationTime2ms
//	* IntegrationTime24ms
//	* IntegrationTime101ms
//	* IntegrationTime154ms
//	* IntegrationTime700ms
func (device *ColorBricklet) SetConfig(gain Gain, integrationTime IntegrationTime) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, gain);
	binary.Write(&buf, binary.LittleEndian, integrationTime);

    resultBytes, err := device.device.Set(uint8(FunctionSetConfig), buf.Bytes())
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

// Returns the configuration as set by SetConfig.
//
// Associated constants:
//
//	* Gain1x
//	* Gain4x
//	* Gain16x
//	* Gain60x
//	* IntegrationTime2ms
//	* IntegrationTime24ms
//	* IntegrationTime101ms
//	* IntegrationTime154ms
//	* IntegrationTime700ms
func (device *ColorBricklet) GetConfig() (gain Gain, integrationTime IntegrationTime, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetConfig), buf.Bytes())
    if err != nil {
        return gain, integrationTime, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return gain, integrationTime, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &gain)
	binary.Read(resultBuf, binary.LittleEndian, &integrationTime)

    }
    
    return gain, integrationTime, nil
}

// Returns the illuminance affected by the gain and integration time as
	// set by SetConfig. To get the illuminance in Lux apply this formula::
	// 
	//  lux = illuminance * 700 / gain / integration_time
	// 
	// To get a correct illuminance measurement make sure that the color
	// values themself are not saturated. The color value (R, G or B)
	// is saturated if it is equal to the maximum value of 65535.
	// In that case you have to reduce the gain, see SetConfig.
func (device *ColorBricklet) GetIlluminance() (illuminance uint32, err error) {    
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

// Returns the color temperature in Kelvin.
	// 
	// To get a correct color temperature measurement make sure that the color
	// values themself are not saturated. The color value (R, G or B)
	// is saturated if it is equal to the maximum value of 65535.
	// In that case you have to reduce the gain, see SetConfig.
func (device *ColorBricklet) GetColorTemperature() (colorTemperature uint16, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetColorTemperature), buf.Bytes())
    if err != nil {
        return colorTemperature, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return colorTemperature, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &colorTemperature)

    }
    
    return colorTemperature, nil
}

// Sets the period in ms with which the RegisterIlluminanceCallback callback is triggered
	// periodically. A value of 0 turns the callback off.
	// 
	// The RegisterIlluminanceCallback callback is only triggered if the illuminance has changed
	// since the last triggering.
	// 
	// The default value is 0.
func (device *ColorBricklet) SetIlluminanceCallbackPeriod(period uint32) (err error) {    
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
func (device *ColorBricklet) GetIlluminanceCallbackPeriod() (period uint32, err error) {    
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

// Sets the period in ms with which the RegisterColorTemperatureCallback callback is
	// triggered periodically. A value of 0 turns the callback off.
	// 
	// The RegisterColorTemperatureCallback callback is only triggered if the color temperature
	// has changed since the last triggering.
	// 
	// The default value is 0.
func (device *ColorBricklet) SetColorTemperatureCallbackPeriod(period uint32) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, period);

    resultBytes, err := device.device.Set(uint8(FunctionSetColorTemperatureCallbackPeriod), buf.Bytes())
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

// Returns the period as set by SetColorTemperatureCallbackPeriod.
func (device *ColorBricklet) GetColorTemperatureCallbackPeriod() (period uint32, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetColorTemperatureCallbackPeriod), buf.Bytes())
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

// Returns the UID, the UID where the Bricklet is connected to,
	// the position, the hardware and firmware version as well as the
	// device identifier.
	// 
	// The position can be 'a', 'b', 'c' or 'd'.
	// 
	// The device identifier numbers can be found `here <device_identifier>`.
	// |device_identifier_constant|
func (device *ColorBricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {    
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
