/* ***********************************************************
 * This file was automatically generated on 2018-12-21.      *
 *                                                           *
 * Go Bindings Version 2.0.0                                 *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


//Measures relative humidity.
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/HumidityV2_Bricklet_Go.html.
package humidity_v2_bricklet

import (
	"encoding/binary"
	"bytes"
    . "github.com/Tinkerforge/go-api-bindings/internal"
    "github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function uint8

const (
    FunctionGetHumidity Function = 1
	FunctionSetHumidityCallbackConfiguration Function = 2
	FunctionGetHumidityCallbackConfiguration Function = 3
	FunctionGetTemperature Function = 5
	FunctionSetTemperatureCallbackConfiguration Function = 6
	FunctionGetTemperatureCallbackConfiguration Function = 7
	FunctionSetHeaterConfiguration Function = 9
	FunctionGetHeaterConfiguration Function = 10
	FunctionSetMovingAverageConfiguration Function = 11
	FunctionGetMovingAverageConfiguration Function = 12
	FunctionSetSamplesPerSecond Function = 13
	FunctionGetSamplesPerSecond Function = 14
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
	FunctionCallbackHumidity Function = 4
	FunctionCallbackTemperature Function = 8
)

type ThresholdOption rune

const (
    ThresholdOptionOff ThresholdOption = 'x'
	ThresholdOptionOutside ThresholdOption = 'o'
	ThresholdOptionInside ThresholdOption = 'i'
	ThresholdOptionSmaller ThresholdOption = '<'
	ThresholdOptionGreater ThresholdOption = '>'
)

type HeaterConfig uint8

const (
    HeaterConfigDisabled HeaterConfig = 0
	HeaterConfigEnabled HeaterConfig = 1
)

type SPS uint8

const (
    SPS20 SPS = 0
	SPS10 SPS = 1
	SPS5 SPS = 2
	SPS1 SPS = 3
	SPS02 SPS = 4
	SPS01 SPS = 5
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

type HumidityV2Bricklet struct{
	device Device
}
const DeviceIdentifier = 283
const DeviceDisplayName = "Humidity Bricklet 2.0"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (HumidityV2Bricklet, error) {
    internalIPCon := ipcon.GetInternalHandle().(IPConnection)
    dev, err := NewDevice([3]uint8{ 2,0,2 }, uid, &internalIPCon, 0)
    if err != nil {
        return HumidityV2Bricklet{}, err
    }
    dev.ResponseExpected[FunctionGetHumidity] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetHumidityCallbackConfiguration] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetHumidityCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetTemperature] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetTemperatureCallbackConfiguration] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetTemperatureCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetHeaterConfiguration] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetHeaterConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetMovingAverageConfiguration] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetMovingAverageConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetSamplesPerSecond] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetSamplesPerSecond] = ResponseExpectedFlagAlwaysTrue;
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
    return HumidityV2Bricklet{dev}, nil
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
func (device *HumidityV2Bricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *HumidityV2Bricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
    return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *HumidityV2Bricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *HumidityV2Bricklet) GetAPIVersion() [3]uint8 {
    return device.device.GetAPIVersion()
}

// This callback is triggered periodically according to the configuration set by
	// SetHumidityCallbackConfiguration.
	// 
	// The parameter is the same as GetHumidity.
func (device *HumidityV2Bricklet) RegisterHumidityCallback(fn func(uint16)) uint64 {
            wrapper := func(byteSlice []byte) {
                buf := bytes.NewBuffer(byteSlice[8:])
                var humidity uint16
                binary.Read(buf, binary.LittleEndian, &humidity)
                fn(humidity)
            }
    return device.device.RegisterCallback(uint8(FunctionCallbackHumidity), wrapper)
}

//Remove a registered Humidity callback.
func (device *HumidityV2Bricklet) DeregisterHumidityCallback(callbackID uint64) {
    device.device.DeregisterCallback(uint8(FunctionCallbackHumidity), callbackID)
}


// This callback is triggered periodically according to the configuration set by
	// SetTemperatureCallbackConfiguration.
	// 
	// The parameter is the same as GetTemperature.
func (device *HumidityV2Bricklet) RegisterTemperatureCallback(fn func(int16)) uint64 {
            wrapper := func(byteSlice []byte) {
                buf := bytes.NewBuffer(byteSlice[8:])
                var temperature int16
                binary.Read(buf, binary.LittleEndian, &temperature)
                fn(temperature)
            }
    return device.device.RegisterCallback(uint8(FunctionCallbackTemperature), wrapper)
}

//Remove a registered Temperature callback.
func (device *HumidityV2Bricklet) DeregisterTemperatureCallback(callbackID uint64) {
    device.device.DeregisterCallback(uint8(FunctionCallbackTemperature), callbackID)
}


// Returns the humidity measured by the sensor. The value
	// has a range of 0 to 10000 and is given in %RH/100 (Relative Humidity),
	// i.e. a value of 4223 means that a humidity of 42.23 %RH is measured.
	// 
	// 
	// If you want to get the value periodically, it is recommended to use the
	// RegisterHumidityCallback callback. You can set the callback configuration
	// with SetHumidityCallbackConfiguration.
func (device *HumidityV2Bricklet) GetHumidity() (humidity uint16, err error) {    
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

// The period in ms is the period with which the RegisterHumidityCallback callback is triggered
	// periodically. A value of 0 turns the callback off.
	// 
	// If the `value has to change`-parameter is set to true, the callback is only
	// triggered after the value has changed. If the value didn't change
	// within the period, the callback is triggered immediately on change.
	// 
	// If it is set to false, the callback is continuously triggered with the period,
	// independent of the value.
	// 
	// It is furthermore possible to constrain the callback with thresholds.
	// 
	// The `option`-parameter together with min/max sets a threshold for the RegisterHumidityCallback callback.
	// 
	// The following options are possible:
	// 
	//  Option| Description
	//  --- | --- 
	//  'x'|    Threshold is turned off
	//  'o'|    Threshold is triggered when the value is *outside* the min and max values
	//  'i'|    Threshold is triggered when the value is *inside* or equal to the min and max values
	//  '<'|    Threshold is triggered when the value is smaller than the min value (max is ignored)
	//  '>'|    Threshold is triggered when the value is greater than the min value (max is ignored)
	// 
	// If the option is set to 'x' (threshold turned off) the callback is triggered with the fixed period.
	// 
	// The default value is (0, false, 'x', 0, 0).
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *HumidityV2Bricklet) SetHumidityCallbackConfiguration(period uint32, valueHasToChange bool, option ThresholdOption, min uint16, max uint16) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, period);
	binary.Write(&buf, binary.LittleEndian, valueHasToChange);
	binary.Write(&buf, binary.LittleEndian, option);
	binary.Write(&buf, binary.LittleEndian, min);
	binary.Write(&buf, binary.LittleEndian, max);

    resultBytes, err := device.device.Set(uint8(FunctionSetHumidityCallbackConfiguration), buf.Bytes())
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

// Returns the callback configuration as set by SetHumidityCallbackConfiguration.
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *HumidityV2Bricklet) GetHumidityCallbackConfiguration() (period uint32, valueHasToChange bool, option ThresholdOption, min uint16, max uint16, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetHumidityCallbackConfiguration), buf.Bytes())
    if err != nil {
        return period, valueHasToChange, option, min, max, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return period, valueHasToChange, option, min, max, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &period)
	binary.Read(resultBuf, binary.LittleEndian, &valueHasToChange)
	binary.Read(resultBuf, binary.LittleEndian, &option)
	binary.Read(resultBuf, binary.LittleEndian, &min)
	binary.Read(resultBuf, binary.LittleEndian, &max)

    }
    
    return period, valueHasToChange, option, min, max, nil
}

// Returns the temperature measured by the sensor. The value
	// has a range of -4000 to 16500 and is given in °C/100,
	// i.e. a value of 3200 means that a temperature of 32.00 °C is measured.
	// 
	// 
	// If you want to get the value periodically, it is recommended to use the
	// RegisterTemperatureCallback callback. You can set the callback configuration
	// with SetTemperatureCallbackConfiguration.
func (device *HumidityV2Bricklet) GetTemperature() (temperature int16, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetTemperature), buf.Bytes())
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

// The period in ms is the period with which the RegisterTemperatureCallback callback is triggered
	// periodically. A value of 0 turns the callback off.
	// 
	// If the `value has to change`-parameter is set to true, the callback is only
	// triggered after the value has changed. If the value didn't change
	// within the period, the callback is triggered immediately on change.
	// 
	// If it is set to false, the callback is continuously triggered with the period,
	// independent of the value.
	// 
	// It is furthermore possible to constrain the callback with thresholds.
	// 
	// The `option`-parameter together with min/max sets a threshold for the RegisterTemperatureCallback callback.
	// 
	// The following options are possible:
	// 
	//  Option| Description
	//  --- | --- 
	//  'x'|    Threshold is turned off
	//  'o'|    Threshold is triggered when the value is *outside* the min and max values
	//  'i'|    Threshold is triggered when the value is *inside* or equal to the min and max values
	//  '<'|    Threshold is triggered when the value is smaller than the min value (max is ignored)
	//  '>'|    Threshold is triggered when the value is greater than the min value (max is ignored)
	// 
	// If the option is set to 'x' (threshold turned off) the callback is triggered with the fixed period.
	// 
	// The default value is (0, false, 'x', 0, 0).
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *HumidityV2Bricklet) SetTemperatureCallbackConfiguration(period uint32, valueHasToChange bool, option ThresholdOption, min int16, max int16) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, period);
	binary.Write(&buf, binary.LittleEndian, valueHasToChange);
	binary.Write(&buf, binary.LittleEndian, option);
	binary.Write(&buf, binary.LittleEndian, min);
	binary.Write(&buf, binary.LittleEndian, max);

    resultBytes, err := device.device.Set(uint8(FunctionSetTemperatureCallbackConfiguration), buf.Bytes())
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

// Returns the callback configuration as set by SetTemperatureCallbackConfiguration.
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *HumidityV2Bricklet) GetTemperatureCallbackConfiguration() (period uint32, valueHasToChange bool, option ThresholdOption, min int16, max int16, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetTemperatureCallbackConfiguration), buf.Bytes())
    if err != nil {
        return period, valueHasToChange, option, min, max, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return period, valueHasToChange, option, min, max, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &period)
	binary.Read(resultBuf, binary.LittleEndian, &valueHasToChange)
	binary.Read(resultBuf, binary.LittleEndian, &option)
	binary.Read(resultBuf, binary.LittleEndian, &min)
	binary.Read(resultBuf, binary.LittleEndian, &max)

    }
    
    return period, valueHasToChange, option, min, max, nil
}

// Enables/disables the heater. The heater can be used to dry the sensor in
	// extremely wet conditions.
	// 
	// By default the heater is disabled.
//
// Associated constants:
//
//	* HeaterConfigDisabled
//	* HeaterConfigEnabled
func (device *HumidityV2Bricklet) SetHeaterConfiguration(heaterConfig HeaterConfig) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, heaterConfig);

    resultBytes, err := device.device.Set(uint8(FunctionSetHeaterConfiguration), buf.Bytes())
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

// Returns the heater configuration as set by SetHeaterConfiguration.
//
// Associated constants:
//
//	* HeaterConfigDisabled
//	* HeaterConfigEnabled
func (device *HumidityV2Bricklet) GetHeaterConfiguration() (heaterConfig HeaterConfig, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetHeaterConfiguration), buf.Bytes())
    if err != nil {
        return heaterConfig, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return heaterConfig, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &heaterConfig)

    }
    
    return heaterConfig, nil
}

// Sets the length of a https://en.wikipedia.org/wiki/Moving_average
	// for the humidity and temperature.
	// 
	// Setting the length to 1 will turn the averaging off. With less
	// averaging, there is more noise on the data.
	// 
	// The range for the averaging is 1-1000.
	// 
	// New data is gathered every 50ms*. With a moving average of length 1000 the resulting
	// averaging window has a length of 50s. If you want to do long term measurements the longest
	// moving average will give the cleanest results.
	// 
	// The default value is 5.
	// 
	// \* In firmware version 2.0.3 we added the SetSamplesPerSecond function. It
	// configures the measurement frequency. Since high frequencies can result in self-heating
	// of th IC, changed the default value from 20 samples per second to 1. With 1 sample per second
	// a moving average length of 1000 would result in an averaging window of 1000 seconds!
func (device *HumidityV2Bricklet) SetMovingAverageConfiguration(movingAverageLengthHumidity uint16, movingAverageLengthTemperature uint16) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, movingAverageLengthHumidity);
	binary.Write(&buf, binary.LittleEndian, movingAverageLengthTemperature);

    resultBytes, err := device.device.Set(uint8(FunctionSetMovingAverageConfiguration), buf.Bytes())
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

// Returns the moving average configuration as set by SetMovingAverageConfiguration.
func (device *HumidityV2Bricklet) GetMovingAverageConfiguration() (movingAverageLengthHumidity uint16, movingAverageLengthTemperature uint16, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetMovingAverageConfiguration), buf.Bytes())
    if err != nil {
        return movingAverageLengthHumidity, movingAverageLengthTemperature, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return movingAverageLengthHumidity, movingAverageLengthTemperature, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &movingAverageLengthHumidity)
	binary.Read(resultBuf, binary.LittleEndian, &movingAverageLengthTemperature)

    }
    
    return movingAverageLengthHumidity, movingAverageLengthTemperature, nil
}

// Sets the samples per second that are gathered by the humidity/temperature sensor HDC1080.
	// 
	// We added this function since we found out that a high measurement frequency can lead to
	// self-heating of the sensor. Which can distort the temperature measurement. 
	// 
	// If you don't need a lot of measurements, you can use the lowest available measurement
	// frequency of 0.1 samples per second for the least amount of self-heating.
	// 
	// Before version 2.0.3 the default was 20 samples per second. The new default is 1 sample per second.
	// 
	// .. versionadded:: 2.0.3$nbsp;(Plugin)
//
// Associated constants:
//
//	* SPS20
//	* SPS10
//	* SPS5
//	* SPS1
//	* SPS02
//	* SPS01
func (device *HumidityV2Bricklet) SetSamplesPerSecond(sps SPS) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, sps);

    resultBytes, err := device.device.Set(uint8(FunctionSetSamplesPerSecond), buf.Bytes())
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

// Returnes the samples per second, as set by SetSamplesPerSecond.
	// 
	// .. versionadded:: 2.0.3$nbsp;(Plugin)
//
// Associated constants:
//
//	* SPS20
//	* SPS10
//	* SPS5
//	* SPS1
//	* SPS02
//	* SPS01
func (device *HumidityV2Bricklet) GetSamplesPerSecond() (sps SPS, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetSamplesPerSecond), buf.Bytes())
    if err != nil {
        return sps, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return sps, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &sps)

    }
    
    return sps, nil
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
func (device *HumidityV2Bricklet) GetSPITFPErrorCount() (errorCountAckChecksum uint32, errorCountMessageChecksum uint32, errorCountFrame uint32, errorCountOverflow uint32, err error) {    
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
func (device *HumidityV2Bricklet) SetBootloaderMode(mode BootloaderMode) (status BootloaderStatus, err error) {    
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
func (device *HumidityV2Bricklet) GetBootloaderMode() (mode BootloaderMode, err error) {    
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
func (device *HumidityV2Bricklet) SetWriteFirmwarePointer(pointer uint32) (err error) {    
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
func (device *HumidityV2Bricklet) WriteFirmware(data [64]uint8) (status uint8, err error) {    
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
func (device *HumidityV2Bricklet) SetStatusLEDConfig(config StatusLEDConfig) (err error) {    
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
func (device *HumidityV2Bricklet) GetStatusLEDConfig() (config StatusLEDConfig, err error) {    
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
func (device *HumidityV2Bricklet) GetChipTemperature() (temperature int16, err error) {    
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
func (device *HumidityV2Bricklet) Reset() (err error) {    
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
func (device *HumidityV2Bricklet) WriteUID(uid uint32) (err error) {    
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
func (device *HumidityV2Bricklet) ReadUID() (uid uint32, err error) {    
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
func (device *HumidityV2Bricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {    
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
