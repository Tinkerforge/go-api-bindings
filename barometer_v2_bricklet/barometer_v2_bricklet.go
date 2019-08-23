/* ***********************************************************
 * This file was automatically generated on 2019-08-23.      *
 *                                                           *
 * Go Bindings Version 2.0.4                                 *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


// Measures air pressure and altitude changes.
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/BarometerV2_Bricklet_Go.html.
package barometer_v2_bricklet

import (
	"encoding/binary"
	"bytes"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionGetAirPressure Function = 1
	FunctionSetAirPressureCallbackConfiguration Function = 2
	FunctionGetAirPressureCallbackConfiguration Function = 3
	FunctionGetAltitude Function = 5
	FunctionSetAltitudeCallbackConfiguration Function = 6
	FunctionGetAltitudeCallbackConfiguration Function = 7
	FunctionGetTemperature Function = 9
	FunctionSetTemperatureCallbackConfiguration Function = 10
	FunctionGetTemperatureCallbackConfiguration Function = 11
	FunctionSetMovingAverageConfiguration Function = 13
	FunctionGetMovingAverageConfiguration Function = 14
	FunctionSetReferenceAirPressure Function = 15
	FunctionGetReferenceAirPressure Function = 16
	FunctionSetCalibration Function = 17
	FunctionGetCalibration Function = 18
	FunctionSetSensorConfiguration Function = 19
	FunctionGetSensorConfiguration Function = 20
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
	FunctionCallbackAirPressure Function = 4
	FunctionCallbackAltitude Function = 8
	FunctionCallbackTemperature Function = 12
)

type ThresholdOption = rune

const (
	ThresholdOptionOff ThresholdOption = 'x'
	ThresholdOptionOutside ThresholdOption = 'o'
	ThresholdOptionInside ThresholdOption = 'i'
	ThresholdOptionSmaller ThresholdOption = '<'
	ThresholdOptionGreater ThresholdOption = '>'
)

type DataRate = uint8

const (
	DataRateOff DataRate = 0
	DataRate1Hz DataRate = 1
	DataRate10Hz DataRate = 2
	DataRate25Hz DataRate = 3
	DataRate50Hz DataRate = 4
	DataRate75Hz DataRate = 5
)

type LowPassFilter = uint8

const (
	LowPassFilterOff LowPassFilter = 0
	LowPassFilter19th LowPassFilter = 1
	LowPassFilter120th LowPassFilter = 2
)

type BootloaderMode = uint8

const (
	BootloaderModeBootloader BootloaderMode = 0
	BootloaderModeFirmware BootloaderMode = 1
	BootloaderModeBootloaderWaitForReboot BootloaderMode = 2
	BootloaderModeFirmwareWaitForReboot BootloaderMode = 3
	BootloaderModeFirmwareWaitForEraseAndReboot BootloaderMode = 4
)

type BootloaderStatus = uint8

const (
	BootloaderStatusOK BootloaderStatus = 0
	BootloaderStatusInvalidMode BootloaderStatus = 1
	BootloaderStatusNoChange BootloaderStatus = 2
	BootloaderStatusEntryFunctionNotPresent BootloaderStatus = 3
	BootloaderStatusDeviceIdentifierIncorrect BootloaderStatus = 4
	BootloaderStatusCRCMismatch BootloaderStatus = 5
)

type StatusLEDConfig = uint8

const (
	StatusLEDConfigOff StatusLEDConfig = 0
	StatusLEDConfigOn StatusLEDConfig = 1
	StatusLEDConfigShowHeartbeat StatusLEDConfig = 2
	StatusLEDConfigShowStatus StatusLEDConfig = 3
)

type BarometerV2Bricklet struct {
	device Device
}
const DeviceIdentifier = 2117
const DeviceDisplayName = "Barometer Bricklet 2.0"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (BarometerV2Bricklet, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{ 2,0,0 }, uid, &internalIPCon, 0)
	if err != nil {
		return BarometerV2Bricklet{}, err
	}
	dev.ResponseExpected[FunctionGetAirPressure] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetAirPressureCallbackConfiguration] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetAirPressureCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetAltitude] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetAltitudeCallbackConfiguration] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetAltitudeCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetTemperature] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetTemperatureCallbackConfiguration] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetTemperatureCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetMovingAverageConfiguration] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetMovingAverageConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetReferenceAirPressure] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetReferenceAirPressure] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetCalibration] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetCalibration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetSensorConfiguration] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetSensorConfiguration] = ResponseExpectedFlagAlwaysTrue;
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
	return BarometerV2Bricklet{dev}, nil
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
func (device *BarometerV2Bricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *BarometerV2Bricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *BarometerV2Bricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *BarometerV2Bricklet) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// This callback is triggered periodically according to the configuration set by
// SetAirPressureCallbackConfiguration.
// 
// The parameter is the same as GetAirPressure.
func (device *BarometerV2Bricklet) RegisterAirPressureCallback(fn func(int32)) uint64 {
	wrapper := func(byteSlice []byte) {
		buf := bytes.NewBuffer(byteSlice[8:])
		var airPressure int32
		binary.Read(buf, binary.LittleEndian, &airPressure)
		fn(airPressure)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackAirPressure), wrapper)
}

// Remove a registered Air Pressure callback.
func (device *BarometerV2Bricklet) DeregisterAirPressureCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackAirPressure), registrationId)
}


// This callback is triggered periodically according to the configuration set by
// SetAltitudeCallbackConfiguration.
// 
// The parameter is the same as GetAltitude.
func (device *BarometerV2Bricklet) RegisterAltitudeCallback(fn func(int32)) uint64 {
	wrapper := func(byteSlice []byte) {
		buf := bytes.NewBuffer(byteSlice[8:])
		var altitude int32
		binary.Read(buf, binary.LittleEndian, &altitude)
		fn(altitude)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackAltitude), wrapper)
}

// Remove a registered Altitude callback.
func (device *BarometerV2Bricklet) DeregisterAltitudeCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackAltitude), registrationId)
}


// This callback is triggered periodically according to the configuration set by
// SetTemperatureCallbackConfiguration.
// 
// The parameter is the same as GetTemperature.
func (device *BarometerV2Bricklet) RegisterTemperatureCallback(fn func(int32)) uint64 {
	wrapper := func(byteSlice []byte) {
		buf := bytes.NewBuffer(byteSlice[8:])
		var temperature int32
		binary.Read(buf, binary.LittleEndian, &temperature)
		fn(temperature)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackTemperature), wrapper)
}

// Remove a registered Temperature callback.
func (device *BarometerV2Bricklet) DeregisterTemperatureCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackTemperature), registrationId)
}


// Returns the measured air pressure. The value has a range of
// 260000 to 1260000 and is given in mbar/1000, i.e. a value of
// 1001092 means that an air pressure of 1001.092 mbar is measured.
// 
// 
// If you want to get the value periodically, it is recommended to use the
// RegisterAirPressureCallback callback. You can set the callback configuration
// with SetAirPressureCallbackConfiguration.
func (device *BarometerV2Bricklet) GetAirPressure() (airPressure int32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetAirPressure), buf.Bytes())
	if err != nil {
		return airPressure, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return airPressure, DeviceError(header.ErrorCode)
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
func (device *BarometerV2Bricklet) SetAirPressureCallbackConfiguration(period uint32, valueHasToChange bool, option ThresholdOption, min int32, max int32) (err error) {
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
			return DeviceError(header.ErrorCode)
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
func (device *BarometerV2Bricklet) GetAirPressureCallbackConfiguration() (period uint32, valueHasToChange bool, option ThresholdOption, min int32, max int32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetAirPressureCallbackConfiguration), buf.Bytes())
	if err != nil {
		return period, valueHasToChange, option, min, max, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return period, valueHasToChange, option, min, max, DeviceError(header.ErrorCode)
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

// Returns the relative altitude of the air pressure sensor. The value
// is given in mm and is calculated based on the difference between the
// current air pressure and the reference air pressure that can be set
// with SetReferenceAirPressure.
// 
// 
// If you want to get the value periodically, it is recommended to use the
// RegisterAltitudeCallback callback. You can set the callback configuration
// with SetAltitudeCallbackConfiguration.
func (device *BarometerV2Bricklet) GetAltitude() (altitude int32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetAltitude), buf.Bytes())
	if err != nil {
		return altitude, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return altitude, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &altitude)

	}

	return altitude, nil
}

// The period in ms is the period with which the RegisterAltitudeCallback callback is triggered
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
// The `option`-parameter together with min/max sets a threshold for the RegisterAltitudeCallback callback.
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
func (device *BarometerV2Bricklet) SetAltitudeCallbackConfiguration(period uint32, valueHasToChange bool, option ThresholdOption, min int32, max int32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);
	binary.Write(&buf, binary.LittleEndian, valueHasToChange);
	binary.Write(&buf, binary.LittleEndian, option);
	binary.Write(&buf, binary.LittleEndian, min);
	binary.Write(&buf, binary.LittleEndian, max);

	resultBytes, err := device.device.Set(uint8(FunctionSetAltitudeCallbackConfiguration), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])
		
	}

	return nil
}

// Returns the callback configuration as set by SetAltitudeCallbackConfiguration.
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *BarometerV2Bricklet) GetAltitudeCallbackConfiguration() (period uint32, valueHasToChange bool, option ThresholdOption, min int32, max int32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetAltitudeCallbackConfiguration), buf.Bytes())
	if err != nil {
		return period, valueHasToChange, option, min, max, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return period, valueHasToChange, option, min, max, DeviceError(header.ErrorCode)
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

// Returns the temperature of the air pressure sensor. The value
// has a range of -4000 to 8500 and is given in °C/100, i.e. a value
// of 2007 means that a temperature of 20.07 °C is measured.
// 
// This temperature is used internally for temperature compensation
// of the air pressure measurement. It is not as accurate as the
// temperature measured by the `temperature_v2_bricklet` or the
// `temperature_ir_v2_bricklet`.
// 
// 
// If you want to get the value periodically, it is recommended to use the
// RegisterTemperatureCallback callback. You can set the callback configuration
// with SetTemperatureCallbackConfiguration.
func (device *BarometerV2Bricklet) GetTemperature() (temperature int32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetTemperature), buf.Bytes())
	if err != nil {
		return temperature, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return temperature, DeviceError(header.ErrorCode)
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
func (device *BarometerV2Bricklet) SetTemperatureCallbackConfiguration(period uint32, valueHasToChange bool, option ThresholdOption, min int32, max int32) (err error) {
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
			return DeviceError(header.ErrorCode)
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
func (device *BarometerV2Bricklet) GetTemperatureCallbackConfiguration() (period uint32, valueHasToChange bool, option ThresholdOption, min int32, max int32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetTemperatureCallbackConfiguration), buf.Bytes())
	if err != nil {
		return period, valueHasToChange, option, min, max, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return period, valueHasToChange, option, min, max, DeviceError(header.ErrorCode)
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

// Sets the length of a https://en.wikipedia.org/wiki/Moving_average
// for the air pressure and temperature measurements.
// 
// Setting the length to 1 will turn the averaging off. With less
// averaging, there is more noise on the data.
// 
// The range for the averaging is 1-1000.
// 
// If you want to do long term measurements the longest moving average will give
// the cleanest results.
// 
// The default value is 100.
func (device *BarometerV2Bricklet) SetMovingAverageConfiguration(movingAverageLengthAirPressure uint16, movingAverageLengthTemperature uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, movingAverageLengthAirPressure);
	binary.Write(&buf, binary.LittleEndian, movingAverageLengthTemperature);

	resultBytes, err := device.device.Set(uint8(FunctionSetMovingAverageConfiguration), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])
		
	}

	return nil
}

// Returns the moving average configuration as set by
// SetMovingAverageConfiguration.
func (device *BarometerV2Bricklet) GetMovingAverageConfiguration() (movingAverageLengthAirPressure uint16, movingAverageLengthTemperature uint16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetMovingAverageConfiguration), buf.Bytes())
	if err != nil {
		return movingAverageLengthAirPressure, movingAverageLengthTemperature, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return movingAverageLengthAirPressure, movingAverageLengthTemperature, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &movingAverageLengthAirPressure)
	binary.Read(resultBuf, binary.LittleEndian, &movingAverageLengthTemperature)

	}

	return movingAverageLengthAirPressure, movingAverageLengthTemperature, nil
}

// Sets the reference air pressure in mbar/1000 for the altitude calculation.
// Valid values are between 260000 and 1260000. Setting the reference to the
// current air pressure results in a calculated altitude of 0mm. Passing 0 is
// a shortcut for passing the current air pressure as reference.
// 
// Well known reference values are the Q codes
// https://en.wikipedia.org/wiki/QNH and
// https://en.wikipedia.org/wiki/Mean_sea_level_pressure#Mean_sea_level_pressure
// used in aviation.
// 
// The default value is 1013.25mbar.
func (device *BarometerV2Bricklet) SetReferenceAirPressure(airPressure int32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, airPressure);

	resultBytes, err := device.device.Set(uint8(FunctionSetReferenceAirPressure), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])
		
	}

	return nil
}

// Returns the reference air pressure as set by SetReferenceAirPressure.
func (device *BarometerV2Bricklet) GetReferenceAirPressure() (airPressure int32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetReferenceAirPressure), buf.Bytes())
	if err != nil {
		return airPressure, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return airPressure, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &airPressure)

	}

	return airPressure, nil
}

// Sets the one point calibration (OPC) values for the air pressure measurement.
// 
// Before the Bricklet can be calibrated any previous calibration has to be removed
// by setting ``measured air pressure`` and ``actual air pressure`` to 0.
// 
// Then the current air pressure has to be measured using the Bricklet
// (``measured air pressure``) and with and accurate reference barometer
// (``actual air pressure``) at the same time and passed to this function in
// mbar/1000.
// 
// After proper calibration the air pressure measurement can achieve an accuracy
// up to 0.2 mbar.
// 
// The calibration is saved in the EEPROM of the Bricklet and only needs to be
// configured once.
func (device *BarometerV2Bricklet) SetCalibration(measuredAirPressure int32, actualAirPressure int32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, measuredAirPressure);
	binary.Write(&buf, binary.LittleEndian, actualAirPressure);

	resultBytes, err := device.device.Set(uint8(FunctionSetCalibration), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])
		
	}

	return nil
}

// Returns the air pressure one point calibration values as set by
// SetCalibration.
func (device *BarometerV2Bricklet) GetCalibration() (measuredAirPressure int32, actualAirPressure int32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetCalibration), buf.Bytes())
	if err != nil {
		return measuredAirPressure, actualAirPressure, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return measuredAirPressure, actualAirPressure, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &measuredAirPressure)
	binary.Read(resultBuf, binary.LittleEndian, &actualAirPressure)

	}

	return measuredAirPressure, actualAirPressure, nil
}

// Configures the data rate and air pressure low pass filter. The low pass filter
// cut-off frequency (if enabled) can be set to 1/9th or 1/20th of the configure
// data rate to decrease the noise on the air pressure data.
// 
// The low pass filter configuration only applies to the air pressure measurement.
// There is no low pass filter for the temperature measurement.
// 
// A higher data rate will result in a less precise temperature because of
// self-heating of the sensor. If the accuracy of the temperature reading is
// important to you, we would recommend the 1Hz data rate.
// 
// The default values are 50Hz data rate and 1/9th low pass filter.
//
// Associated constants:
//
//	* DataRateOff
//	* DataRate1Hz
//	* DataRate10Hz
//	* DataRate25Hz
//	* DataRate50Hz
//	* DataRate75Hz
//	* LowPassFilterOff
//	* LowPassFilter19th
//	* LowPassFilter120th
func (device *BarometerV2Bricklet) SetSensorConfiguration(dataRate DataRate, airPressureLowPassFilter LowPassFilter) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, dataRate);
	binary.Write(&buf, binary.LittleEndian, airPressureLowPassFilter);

	resultBytes, err := device.device.Set(uint8(FunctionSetSensorConfiguration), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])
		
	}

	return nil
}

// Returns the sensor configuration as set by SetSensorConfiguration.
//
// Associated constants:
//
//	* DataRateOff
//	* DataRate1Hz
//	* DataRate10Hz
//	* DataRate25Hz
//	* DataRate50Hz
//	* DataRate75Hz
//	* LowPassFilterOff
//	* LowPassFilter19th
//	* LowPassFilter120th
func (device *BarometerV2Bricklet) GetSensorConfiguration() (dataRate DataRate, airPressureLowPassFilter LowPassFilter, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetSensorConfiguration), buf.Bytes())
	if err != nil {
		return dataRate, airPressureLowPassFilter, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return dataRate, airPressureLowPassFilter, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &dataRate)
	binary.Read(resultBuf, binary.LittleEndian, &airPressureLowPassFilter)

	}

	return dataRate, airPressureLowPassFilter, nil
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
func (device *BarometerV2Bricklet) GetSPITFPErrorCount() (errorCountAckChecksum uint32, errorCountMessageChecksum uint32, errorCountFrame uint32, errorCountOverflow uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetSPITFPErrorCount), buf.Bytes())
	if err != nil {
		return errorCountAckChecksum, errorCountMessageChecksum, errorCountFrame, errorCountOverflow, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return errorCountAckChecksum, errorCountMessageChecksum, errorCountFrame, errorCountOverflow, DeviceError(header.ErrorCode)
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
func (device *BarometerV2Bricklet) SetBootloaderMode(mode BootloaderMode) (status BootloaderStatus, err error) {
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
			return status, DeviceError(header.ErrorCode)
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
func (device *BarometerV2Bricklet) GetBootloaderMode() (mode BootloaderMode, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetBootloaderMode), buf.Bytes())
	if err != nil {
		return mode, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return mode, DeviceError(header.ErrorCode)
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
func (device *BarometerV2Bricklet) SetWriteFirmwarePointer(pointer uint32) (err error) {
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
			return DeviceError(header.ErrorCode)
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
func (device *BarometerV2Bricklet) WriteFirmware(data [64]uint8) (status uint8, err error) {
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
			return status, DeviceError(header.ErrorCode)
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
func (device *BarometerV2Bricklet) SetStatusLEDConfig(config StatusLEDConfig) (err error) {
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
			return DeviceError(header.ErrorCode)
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
func (device *BarometerV2Bricklet) GetStatusLEDConfig() (config StatusLEDConfig, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetStatusLEDConfig), buf.Bytes())
	if err != nil {
		return config, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return config, DeviceError(header.ErrorCode)
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
func (device *BarometerV2Bricklet) GetChipTemperature() (temperature int16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetChipTemperature), buf.Bytes())
	if err != nil {
		return temperature, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return temperature, DeviceError(header.ErrorCode)
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
func (device *BarometerV2Bricklet) Reset() (err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Set(uint8(FunctionReset), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
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
func (device *BarometerV2Bricklet) WriteUID(uid uint32) (err error) {
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
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])
		
	}

	return nil
}

// Returns the current UID as an integer. Encode as
// Base58 to get the usual string version.
func (device *BarometerV2Bricklet) ReadUID() (uid uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionReadUID), buf.Bytes())
	if err != nil {
		return uid, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return uid, DeviceError(header.ErrorCode)
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
func (device *BarometerV2Bricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetIdentity), buf.Bytes())
	if err != nil {
		return uid, connectedUid, position, hardwareVersion, firmwareVersion, deviceIdentifier, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return uid, connectedUid, position, hardwareVersion, firmwareVersion, deviceIdentifier, DeviceError(header.ErrorCode)
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
