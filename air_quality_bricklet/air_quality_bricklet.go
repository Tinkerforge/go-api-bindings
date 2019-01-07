/* ***********************************************************
 * This file was automatically generated on 2019-01-07.      *
 *                                                           *
 * Go Bindings Version 2.0.1                                 *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


//Measures IAQ index, temperature, humidity and air pressure.
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/AirQuality_Bricklet_Go.html.
package air_quality_bricklet

import (
	"encoding/binary"
	"bytes"
    . "github.com/Tinkerforge/go-api-bindings/internal"
    "github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function uint8

const (
    FunctionGetAllValues Function = 1
	FunctionSetTemperatureOffset Function = 2
	FunctionGetTemperatureOffset Function = 3
	FunctionSetAllValuesCallbackConfiguration Function = 4
	FunctionGetAllValuesCallbackConfiguration Function = 5
	FunctionGetIAQIndex Function = 7
	FunctionSetIAQIndexCallbackConfiguration Function = 8
	FunctionGetIAQIndexCallbackConfiguration Function = 9
	FunctionGetTemperature Function = 11
	FunctionSetTemperatureCallbackConfiguration Function = 12
	FunctionGetTemperatureCallbackConfiguration Function = 13
	FunctionGetHumidity Function = 15
	FunctionSetHumidityCallbackConfiguration Function = 16
	FunctionGetHumidityCallbackConfiguration Function = 17
	FunctionGetAirPressure Function = 19
	FunctionSetAirPressureCallbackConfiguration Function = 20
	FunctionGetAirPressureCallbackConfiguration Function = 21
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
	FunctionCallbackAllValues Function = 6
	FunctionCallbackIAQIndex Function = 10
	FunctionCallbackTemperature Function = 14
	FunctionCallbackHumidity Function = 18
	FunctionCallbackAirPressure Function = 22
)

type Accuracy uint8

const (
    AccuracyUnreliable Accuracy = 0
	AccuracyLow Accuracy = 1
	AccuracyMedium Accuracy = 2
	AccuracyHigh Accuracy = 3
)

type ThresholdOption rune

const (
    ThresholdOptionOff ThresholdOption = 'x'
	ThresholdOptionOutside ThresholdOption = 'o'
	ThresholdOptionInside ThresholdOption = 'i'
	ThresholdOptionSmaller ThresholdOption = '<'
	ThresholdOptionGreater ThresholdOption = '>'
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

type AirQualityBricklet struct{
	device Device
}
const DeviceIdentifier = 297
const DeviceDisplayName = "Air Quality Bricklet"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (AirQualityBricklet, error) {
    internalIPCon := ipcon.GetInternalHandle().(IPConnection)
    dev, err := NewDevice([3]uint8{ 2,0,0 }, uid, &internalIPCon, 0)
    if err != nil {
        return AirQualityBricklet{}, err
    }
    dev.ResponseExpected[FunctionGetAllValues] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetTemperatureOffset] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetTemperatureOffset] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetAllValuesCallbackConfiguration] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetAllValuesCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetIAQIndex] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetIAQIndexCallbackConfiguration] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetIAQIndexCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetTemperature] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetTemperatureCallbackConfiguration] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetTemperatureCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetHumidity] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetHumidityCallbackConfiguration] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetHumidityCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetAirPressure] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetAirPressureCallbackConfiguration] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetAirPressureCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue;
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
    return AirQualityBricklet{dev}, nil
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
func (device *AirQualityBricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *AirQualityBricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
    return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *AirQualityBricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *AirQualityBricklet) GetAPIVersion() [3]uint8 {
    return device.device.GetAPIVersion()
}

// This callback is triggered periodically according to the configuration set by
	// SetAllValuesCallbackConfiguration.
	// 
	// The parameters are the same as GetAllValues.
func (device *AirQualityBricklet) RegisterAllValuesCallback(fn func(int32, Accuracy, int32, int32, int32)) uint64 {
            wrapper := func(byteSlice []byte) {
                buf := bytes.NewBuffer(byteSlice[8:])
                var iaqIndex int32
var iaqIndexAccuracy Accuracy
var temperature int32
var humidity int32
var airPressure int32
                binary.Read(buf, binary.LittleEndian, &iaqIndex)
binary.Read(buf, binary.LittleEndian, &iaqIndexAccuracy)
binary.Read(buf, binary.LittleEndian, &temperature)
binary.Read(buf, binary.LittleEndian, &humidity)
binary.Read(buf, binary.LittleEndian, &airPressure)
                fn(iaqIndex, iaqIndexAccuracy, temperature, humidity, airPressure)
            }
    return device.device.RegisterCallback(uint8(FunctionCallbackAllValues), wrapper)
}

//Remove a registered All Values callback.
func (device *AirQualityBricklet) DeregisterAllValuesCallback(callbackID uint64) {
    device.device.DeregisterCallback(uint8(FunctionCallbackAllValues), callbackID)
}


// This callback is triggered periodically according to the configuration set by
	// SetIAQIndexCallbackConfiguration.
	// 
	// The parameters are the same as GetIAQIndex.
func (device *AirQualityBricklet) RegisterIAQIndexCallback(fn func(int32, Accuracy)) uint64 {
            wrapper := func(byteSlice []byte) {
                buf := bytes.NewBuffer(byteSlice[8:])
                var iaqIndex int32
var iaqIndexAccuracy Accuracy
                binary.Read(buf, binary.LittleEndian, &iaqIndex)
binary.Read(buf, binary.LittleEndian, &iaqIndexAccuracy)
                fn(iaqIndex, iaqIndexAccuracy)
            }
    return device.device.RegisterCallback(uint8(FunctionCallbackIAQIndex), wrapper)
}

//Remove a registered IAQ Index callback.
func (device *AirQualityBricklet) DeregisterIAQIndexCallback(callbackID uint64) {
    device.device.DeregisterCallback(uint8(FunctionCallbackIAQIndex), callbackID)
}


// This callback is triggered periodically according to the configuration set by
	// SetTemperatureCallbackConfiguration.
	// 
	// The parameter is the same as GetTemperature.
func (device *AirQualityBricklet) RegisterTemperatureCallback(fn func(int32)) uint64 {
            wrapper := func(byteSlice []byte) {
                buf := bytes.NewBuffer(byteSlice[8:])
                var temperature int32
                binary.Read(buf, binary.LittleEndian, &temperature)
                fn(temperature)
            }
    return device.device.RegisterCallback(uint8(FunctionCallbackTemperature), wrapper)
}

//Remove a registered Temperature callback.
func (device *AirQualityBricklet) DeregisterTemperatureCallback(callbackID uint64) {
    device.device.DeregisterCallback(uint8(FunctionCallbackTemperature), callbackID)
}


// This callback is triggered periodically according to the configuration set by
	// SetHumidityCallbackConfiguration.
	// 
	// The parameter is the same as GetHumidity.
func (device *AirQualityBricklet) RegisterHumidityCallback(fn func(int32)) uint64 {
            wrapper := func(byteSlice []byte) {
                buf := bytes.NewBuffer(byteSlice[8:])
                var humidity int32
                binary.Read(buf, binary.LittleEndian, &humidity)
                fn(humidity)
            }
    return device.device.RegisterCallback(uint8(FunctionCallbackHumidity), wrapper)
}

//Remove a registered Humidity callback.
func (device *AirQualityBricklet) DeregisterHumidityCallback(callbackID uint64) {
    device.device.DeregisterCallback(uint8(FunctionCallbackHumidity), callbackID)
}


// This callback is triggered periodically according to the configuration set by
	// SetAirPressureCallbackConfiguration.
	// 
	// The parameter is the same as GetAirPressure.
func (device *AirQualityBricklet) RegisterAirPressureCallback(fn func(int32)) uint64 {
            wrapper := func(byteSlice []byte) {
                buf := bytes.NewBuffer(byteSlice[8:])
                var airPressure int32
                binary.Read(buf, binary.LittleEndian, &airPressure)
                fn(airPressure)
            }
    return device.device.RegisterCallback(uint8(FunctionCallbackAirPressure), wrapper)
}

//Remove a registered Air Pressure callback.
func (device *AirQualityBricklet) DeregisterAirPressureCallback(callbackID uint64) {
    device.device.DeregisterCallback(uint8(FunctionCallbackAirPressure), callbackID)
}


// Returns all values measured by the Air Quality Bricklet. The values are
	// IAQ (Indoor Air Quality) Index, IAQ Index Accuracy, Temperature, Humidity and
	// Air Pressure.
	// 
	// .. image:: /Images/Misc/bricklet_air_quality_iaq_index.png
	//    :scale: 100 %
	//    :alt: Air Quality Index description
	//    :align: center
	//    :target: ../../_images/Misc/bricklet_air_quality_iaq_index.png
	// 
	// The values have these ranges and units:
	// 
	// * IAQ Index: 0 to 500, higher value means greater level of air pollution
	// * IAQ Index Accuracy: 0 = unreliable to 3 = high
	// * Temperature: in steps of 0.01 °C
	// * Humidity: in steps of 0.01 %RH
	// * Air Pressure: in steps of 0.01 mbar
//
// Associated constants:
//
//	* AccuracyUnreliable
//	* AccuracyLow
//	* AccuracyMedium
//	* AccuracyHigh
func (device *AirQualityBricklet) GetAllValues() (iaqIndex int32, iaqIndexAccuracy Accuracy, temperature int32, humidity int32, airPressure int32, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetAllValues), buf.Bytes())
    if err != nil {
        return iaqIndex, iaqIndexAccuracy, temperature, humidity, airPressure, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return iaqIndex, iaqIndexAccuracy, temperature, humidity, airPressure, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &iaqIndex)
	binary.Read(resultBuf, binary.LittleEndian, &iaqIndexAccuracy)
	binary.Read(resultBuf, binary.LittleEndian, &temperature)
	binary.Read(resultBuf, binary.LittleEndian, &humidity)
	binary.Read(resultBuf, binary.LittleEndian, &airPressure)

    }
    
    return iaqIndex, iaqIndexAccuracy, temperature, humidity, airPressure, nil
}

// Sets a temperature offset with resolution 1/100°C. A offset of 10 will decrease 
	// the measured temperature by 0.1°C.
	// 
	// If you install this Bricklet into an enclosure and you want to measure the ambient
	// temperature, you may have to decrease the measured temperature by some value to
	// compensate for the error because of the heating inside of the enclosure.
	// 
	// We recommend that you leave the parts in the enclosure running for at least
	// 24 hours such that a temperature equilibrium can be reached. After that you can measure
	// the temperature directly outside of enclosure and set the difference as offset.
	// 
	// This temperature offset is used to calculate the relative humidity and
	// IAQ index measurements. In case the Bricklet is installed in an enclosure, we
	// recommend to measure and set the temperature offset to imporve the accuracy of
	// the measurements.
func (device *AirQualityBricklet) SetTemperatureOffset(offset int32) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, offset);

    resultBytes, err := device.device.Set(uint8(FunctionSetTemperatureOffset), buf.Bytes())
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

// Returns the temperature offset as set by
	// SetTemperatureOffset.
func (device *AirQualityBricklet) GetTemperatureOffset() (offset int32, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetTemperatureOffset), buf.Bytes())
    if err != nil {
        return offset, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return offset, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &offset)

    }
    
    return offset, nil
}

// The period in ms is the period with which the RegisterAllValuesCallback
	// callback is triggered periodically. A value of 0 turns the callback off.
	// 
	// If the `value has to change`-parameter is set to true, the callback is only
	// triggered after at least one of the values has changed. If the values didn't
	// change within the period, the callback is triggered immediately on change.
	// 
	// If it is set to false, the callback is continuously triggered with the period,
	// independent of the value.
	// 
	// The default value is (0, false).
func (device *AirQualityBricklet) SetAllValuesCallbackConfiguration(period uint32, valueHasToChange bool) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, period);
	binary.Write(&buf, binary.LittleEndian, valueHasToChange);

    resultBytes, err := device.device.Set(uint8(FunctionSetAllValuesCallbackConfiguration), buf.Bytes())
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
	// SetAllValuesCallbackConfiguration.
func (device *AirQualityBricklet) GetAllValuesCallbackConfiguration() (period uint32, valueHasToChange bool, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetAllValuesCallbackConfiguration), buf.Bytes())
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

// Returns the IAQ index and accuracy. The IAQ index goes from
	// 0 to 500. The higher the IAQ index, the greater the level of air pollution.
	// 
	// .. image:: /Images/Misc/bricklet_air_quality_iaq_index.png
	//    :scale: 100 %
	//    :alt: IAQ index description
	//    :align: center
	//    :target: ../../_images/Misc/bricklet_air_quality_iaq_index.png
	// 
	// If you want to get the value periodically, it is recommended to use the
	// RegisterIAQIndexCallback callback. You can set the callback configuration
	// with SetIAQIndexCallbackConfiguration.
//
// Associated constants:
//
//	* AccuracyUnreliable
//	* AccuracyLow
//	* AccuracyMedium
//	* AccuracyHigh
func (device *AirQualityBricklet) GetIAQIndex() (iaqIndex int32, iaqIndexAccuracy Accuracy, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetIAQIndex), buf.Bytes())
    if err != nil {
        return iaqIndex, iaqIndexAccuracy, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return iaqIndex, iaqIndexAccuracy, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &iaqIndex)
	binary.Read(resultBuf, binary.LittleEndian, &iaqIndexAccuracy)

    }
    
    return iaqIndex, iaqIndexAccuracy, nil
}

// The period in ms is the period with which the RegisterIAQIndexCallback
	// callback is triggered periodically. A value of 0 turns the callback off.
	// 
	// If the `value has to change`-parameter is set to true, the callback is only
	// triggered after at least one of the values has changed. If the values didn't
	// change within the period, the callback is triggered immediately on change.
	// 
	// If it is set to false, the callback is continuously triggered with the period,
	// independent of the value.
	// 
	// The default value is (0, false).
func (device *AirQualityBricklet) SetIAQIndexCallbackConfiguration(period uint32, valueHasToChange bool) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, period);
	binary.Write(&buf, binary.LittleEndian, valueHasToChange);

    resultBytes, err := device.device.Set(uint8(FunctionSetIAQIndexCallbackConfiguration), buf.Bytes())
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
	// SetIAQIndexCallbackConfiguration.
func (device *AirQualityBricklet) GetIAQIndexCallbackConfiguration() (period uint32, valueHasToChange bool, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetIAQIndexCallbackConfiguration), buf.Bytes())
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

// Returns temperature in steps of 0.01 °C.
	// 
	// 
	// If you want to get the value periodically, it is recommended to use the
	// RegisterTemperatureCallback callback. You can set the callback configuration
	// with SetTemperatureCallbackConfiguration.
func (device *AirQualityBricklet) GetTemperature() (temperature int32, err error) {    
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
func (device *AirQualityBricklet) SetTemperatureCallbackConfiguration(period uint32, valueHasToChange bool, option ThresholdOption, min int32, max int32) (err error) {    
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
func (device *AirQualityBricklet) GetTemperatureCallbackConfiguration() (period uint32, valueHasToChange bool, option ThresholdOption, min int32, max int32, err error) {    
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

// Returns relative humidity in steps of 0.01 %RH.
	// 
	// 
	// If you want to get the value periodically, it is recommended to use the
	// RegisterHumidityCallback callback. You can set the callback configuration
	// with SetHumidityCallbackConfiguration.
func (device *AirQualityBricklet) GetHumidity() (humidity int32, err error) {    
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
func (device *AirQualityBricklet) SetHumidityCallbackConfiguration(period uint32, valueHasToChange bool, option ThresholdOption, min int32, max int32) (err error) {    
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
func (device *AirQualityBricklet) GetHumidityCallbackConfiguration() (period uint32, valueHasToChange bool, option ThresholdOption, min int32, max int32, err error) {    
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

// Returns air pressure in steps of 0.01 mbar.
	// 
	// 
	// If you want to get the value periodically, it is recommended to use the
	// RegisterAirPressureCallback callback. You can set the callback configuration
	// with SetAirPressureCallbackConfiguration.
func (device *AirQualityBricklet) GetAirPressure() (airPressure int32, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetAirPressure), buf.Bytes())
    if err != nil {
        return airPressure, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return airPressure, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &airPressure)

    }
    
    return airPressure, nil
}

// The period in ms is the period with which the RegisterAirPressureCallback callback is triggered
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
	// The `option`-parameter together with min/max sets a threshold for the RegisterAirPressureCallback callback.
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
func (device *AirQualityBricklet) SetAirPressureCallbackConfiguration(period uint32, valueHasToChange bool, option ThresholdOption, min int32, max int32) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, period);
	binary.Write(&buf, binary.LittleEndian, valueHasToChange);
	binary.Write(&buf, binary.LittleEndian, option);
	binary.Write(&buf, binary.LittleEndian, min);
	binary.Write(&buf, binary.LittleEndian, max);

    resultBytes, err := device.device.Set(uint8(FunctionSetAirPressureCallbackConfiguration), buf.Bytes())
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

// Returns the callback configuration as set by SetAirPressureCallbackConfiguration.
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *AirQualityBricklet) GetAirPressureCallbackConfiguration() (period uint32, valueHasToChange bool, option ThresholdOption, min int32, max int32, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetAirPressureCallbackConfiguration), buf.Bytes())
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
func (device *AirQualityBricklet) GetSPITFPErrorCount() (errorCountAckChecksum uint32, errorCountMessageChecksum uint32, errorCountFrame uint32, errorCountOverflow uint32, err error) {    
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
func (device *AirQualityBricklet) SetBootloaderMode(mode BootloaderMode) (status BootloaderStatus, err error) {    
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
func (device *AirQualityBricklet) GetBootloaderMode() (mode BootloaderMode, err error) {    
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
func (device *AirQualityBricklet) SetWriteFirmwarePointer(pointer uint32) (err error) {    
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
func (device *AirQualityBricklet) WriteFirmware(data [64]uint8) (status uint8, err error) {    
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
func (device *AirQualityBricklet) SetStatusLEDConfig(config StatusLEDConfig) (err error) {    
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
func (device *AirQualityBricklet) GetStatusLEDConfig() (config StatusLEDConfig, err error) {    
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
func (device *AirQualityBricklet) GetChipTemperature() (temperature int16, err error) {    
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
func (device *AirQualityBricklet) Reset() (err error) {    
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
func (device *AirQualityBricklet) WriteUID(uid uint32) (err error) {    
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
func (device *AirQualityBricklet) ReadUID() (uid uint32, err error) {    
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
func (device *AirQualityBricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {    
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
