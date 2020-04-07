/* ***********************************************************
 * This file was automatically generated on 2020-04-07.      *
 *                                                           *
 * Go Bindings Version 2.0.6                                 *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


// Measures UV-A, UV-B and UV index.
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/UVLightV2_Bricklet_Go.html.
package uv_light_v2_bricklet

import (
	"encoding/binary"
	"bytes"
	"fmt"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionGetUVA Function = 1
	FunctionSetUVACallbackConfiguration Function = 2
	FunctionGetUVACallbackConfiguration Function = 3
	FunctionGetUVB Function = 5
	FunctionSetUVBCallbackConfiguration Function = 6
	FunctionGetUVBCallbackConfiguration Function = 7
	FunctionGetUVI Function = 9
	FunctionSetUVICallbackConfiguration Function = 10
	FunctionGetUVICallbackConfiguration Function = 11
	FunctionSetConfiguration Function = 13
	FunctionGetConfiguration Function = 14
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
	FunctionCallbackUVA Function = 4
	FunctionCallbackUVB Function = 8
	FunctionCallbackUVI Function = 12
)

type ThresholdOption = rune

const (
	ThresholdOptionOff ThresholdOption = 'x'
	ThresholdOptionOutside ThresholdOption = 'o'
	ThresholdOptionInside ThresholdOption = 'i'
	ThresholdOptionSmaller ThresholdOption = '<'
	ThresholdOptionGreater ThresholdOption = '>'
)

type IntegrationTime = uint8

const (
	IntegrationTime50ms IntegrationTime = 0
	IntegrationTime100ms IntegrationTime = 1
	IntegrationTime200ms IntegrationTime = 2
	IntegrationTime400ms IntegrationTime = 3
	IntegrationTime800ms IntegrationTime = 4
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

type UVLightV2Bricklet struct {
	device Device
}
const DeviceIdentifier = 2118
const DeviceDisplayName = "UV Light Bricklet 2.0"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (UVLightV2Bricklet, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{ 2,0,0 }, uid, &internalIPCon, 0, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return UVLightV2Bricklet{}, err
	}
	dev.ResponseExpected[FunctionGetUVA] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetUVACallbackConfiguration] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetUVACallbackConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetUVB] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetUVBCallbackConfiguration] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetUVBCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetUVI] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetUVICallbackConfiguration] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetUVICallbackConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetConfiguration] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetConfiguration] = ResponseExpectedFlagAlwaysTrue;
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
	return UVLightV2Bricklet{dev}, nil
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
// for this purpose. If this flag is disabled for a setter function then no response is sent
// and errors are silently ignored, because they cannot be detected.
//
// See SetResponseExpected for the list of function ID constants available for this function.
func (device *UVLightV2Bricklet) GetResponseExpected(functionID Function) (bool, error) {
	return device.device.GetResponseExpected(uint8(functionID))
}

// Changes the response expected flag of the function specified by the function ID parameter.
// This flag can only be changed for setter (default value: false) and callback configuration
// functions (default value: true). For getter functions it is always enabled.
//
// Enabling the response expected flag for a setter function allows to detect timeouts and
// other error conditions calls of this setter as well. The device will then send a response
// for this purpose. If this flag is disabled for a setter function then no response is sent
// and errors are silently ignored, because they cannot be detected.
func (device *UVLightV2Bricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *UVLightV2Bricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *UVLightV2Bricklet) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// This callback is triggered periodically according to the configuration set by
// SetUVACallbackConfiguration.
// 
// The parameter is the same as GetUVA.
func (device *UVLightV2Bricklet) RegisterUVACallback(fn func(int32)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 12 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var uva int32
		binary.Read(buf, binary.LittleEndian, &uva)
		fn(uva)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackUVA), wrapper)
}

// Remove a registered UVA callback.
func (device *UVLightV2Bricklet) DeregisterUVACallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackUVA), registrationId)
}


// This callback is triggered periodically according to the configuration set by
// SetUVBCallbackConfiguration.
// 
// The parameter is the same as GetUVB.
func (device *UVLightV2Bricklet) RegisterUVBCallback(fn func(int32)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 12 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var uvb int32
		binary.Read(buf, binary.LittleEndian, &uvb)
		fn(uvb)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackUVB), wrapper)
}

// Remove a registered UVB callback.
func (device *UVLightV2Bricklet) DeregisterUVBCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackUVB), registrationId)
}


// This callback is triggered periodically according to the configuration set by
// SetUVICallbackConfiguration.
// 
// The parameter is the same as GetUVI.
func (device *UVLightV2Bricklet) RegisterUVICallback(fn func(int32)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 12 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var uvi int32
		binary.Read(buf, binary.LittleEndian, &uvi)
		fn(uvi)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackUVI), wrapper)
}

// Remove a registered UVI callback.
func (device *UVLightV2Bricklet) DeregisterUVICallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackUVI), registrationId)
}


// Returns the UVA intensity of the sensor.
// The sensor has not weighted the intensity with the erythemal
// action spectrum to get the skin-affecting irradiation. Therefore, you cannot
// just divide the value by 250 to get the UVA index. To get the UV index use
// GetUVI.
// 
// If the sensor is saturated, then -1 is returned, see SetConfiguration.
// 
// If you want to get the intensity periodically, it is recommended to use the
// RegisterUVACallback callback and set the period with
// SetUVACallbackConfiguration.
// 
// 
// If you want to get the value periodically, it is recommended to use the
// RegisterUVACallback callback. You can set the callback configuration
// with SetUVACallbackConfiguration.
func (device *UVLightV2Bricklet) GetUVA() (uva int32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetUVA), buf.Bytes())
	if err != nil {
		return uva, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 12 {
		return uva, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
	}


	if header.ErrorCode != 0 {
		return uva, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &uva)


	return uva, nil
}

// The period is the period with which the RegisterUVACallback callback is triggered
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
// The `option`-parameter together with min/max sets a threshold for the RegisterUVACallback callback.
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
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *UVLightV2Bricklet) SetUVACallbackConfiguration(period uint32, valueHasToChange bool, option ThresholdOption, min int32, max int32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);
	binary.Write(&buf, binary.LittleEndian, valueHasToChange);
	binary.Write(&buf, binary.LittleEndian, option);
	binary.Write(&buf, binary.LittleEndian, min);
	binary.Write(&buf, binary.LittleEndian, max);

	resultBytes, err := device.device.Set(uint8(FunctionSetUVACallbackConfiguration), buf.Bytes())
	if err != nil {
		return err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 8 {
		return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
	}


	if header.ErrorCode != 0 {
		return DeviceError(header.ErrorCode)
	}

	bytes.NewBuffer(resultBytes[8:])
	

	return nil
}

// Returns the callback configuration as set by SetUVACallbackConfiguration.
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *UVLightV2Bricklet) GetUVACallbackConfiguration() (period uint32, valueHasToChange bool, option ThresholdOption, min int32, max int32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetUVACallbackConfiguration), buf.Bytes())
	if err != nil {
		return period, valueHasToChange, option, min, max, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 22 {
		return period, valueHasToChange, option, min, max, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 22)
	}


	if header.ErrorCode != 0 {
		return period, valueHasToChange, option, min, max, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &period)
	binary.Read(resultBuf, binary.LittleEndian, &valueHasToChange)
	binary.Read(resultBuf, binary.LittleEndian, &option)
	binary.Read(resultBuf, binary.LittleEndian, &min)
	binary.Read(resultBuf, binary.LittleEndian, &max)


	return period, valueHasToChange, option, min, max, nil
}

// Returns the UVB intensity of the sensor.
// The sensor has not weighted the intensity with the erythemal
// action spectrum to get the skin-affecting irradiation. Therefore, you cannot
// just divide the value by 250 to get the UVB index. To get the UV index use
// GetUVI.
// 
// If the sensor is saturated, then -1 is returned, see SetConfiguration.
// 
// If you want to get the intensity periodically, it is recommended to use the
// RegisterUVBCallback callback and set the period with
// SetUVBCallbackConfiguration.
// 
// 
// If you want to get the value periodically, it is recommended to use the
// RegisterUVBCallback callback. You can set the callback configuration
// with SetUVBCallbackConfiguration.
func (device *UVLightV2Bricklet) GetUVB() (uvb int32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetUVB), buf.Bytes())
	if err != nil {
		return uvb, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 12 {
		return uvb, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
	}


	if header.ErrorCode != 0 {
		return uvb, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &uvb)


	return uvb, nil
}

// The period is the period with which the RegisterUVBCallback callback is triggered
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
// The `option`-parameter together with min/max sets a threshold for the RegisterUVBCallback callback.
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
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *UVLightV2Bricklet) SetUVBCallbackConfiguration(period uint32, valueHasToChange bool, option ThresholdOption, min int32, max int32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);
	binary.Write(&buf, binary.LittleEndian, valueHasToChange);
	binary.Write(&buf, binary.LittleEndian, option);
	binary.Write(&buf, binary.LittleEndian, min);
	binary.Write(&buf, binary.LittleEndian, max);

	resultBytes, err := device.device.Set(uint8(FunctionSetUVBCallbackConfiguration), buf.Bytes())
	if err != nil {
		return err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 8 {
		return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
	}


	if header.ErrorCode != 0 {
		return DeviceError(header.ErrorCode)
	}

	bytes.NewBuffer(resultBytes[8:])
	

	return nil
}

// Returns the callback configuration as set by SetUVBCallbackConfiguration.
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *UVLightV2Bricklet) GetUVBCallbackConfiguration() (period uint32, valueHasToChange bool, option ThresholdOption, min int32, max int32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetUVBCallbackConfiguration), buf.Bytes())
	if err != nil {
		return period, valueHasToChange, option, min, max, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 22 {
		return period, valueHasToChange, option, min, max, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 22)
	}


	if header.ErrorCode != 0 {
		return period, valueHasToChange, option, min, max, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &period)
	binary.Read(resultBuf, binary.LittleEndian, &valueHasToChange)
	binary.Read(resultBuf, binary.LittleEndian, &option)
	binary.Read(resultBuf, binary.LittleEndian, &min)
	binary.Read(resultBuf, binary.LittleEndian, &max)


	return period, valueHasToChange, option, min, max, nil
}

// Returns the UV index of the sensor, the index is given in 1/10.
// 
// If the sensor is saturated, then -1 is returned, see SetConfiguration.
// 
// If you want to get the intensity periodically, it is recommended to use the
// RegisterUVICallback callback and set the period with
// SetUVICallbackConfiguration.
// 
// 
// If you want to get the value periodically, it is recommended to use the
// RegisterUVICallback callback. You can set the callback configuration
// with SetUVICallbackConfiguration.
func (device *UVLightV2Bricklet) GetUVI() (uvi int32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetUVI), buf.Bytes())
	if err != nil {
		return uvi, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 12 {
		return uvi, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
	}


	if header.ErrorCode != 0 {
		return uvi, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &uvi)


	return uvi, nil
}

// The period is the period with which the RegisterUVICallback callback is triggered
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
// The `option`-parameter together with min/max sets a threshold for the RegisterUVICallback callback.
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
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *UVLightV2Bricklet) SetUVICallbackConfiguration(period uint32, valueHasToChange bool, option ThresholdOption, min int32, max int32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);
	binary.Write(&buf, binary.LittleEndian, valueHasToChange);
	binary.Write(&buf, binary.LittleEndian, option);
	binary.Write(&buf, binary.LittleEndian, min);
	binary.Write(&buf, binary.LittleEndian, max);

	resultBytes, err := device.device.Set(uint8(FunctionSetUVICallbackConfiguration), buf.Bytes())
	if err != nil {
		return err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 8 {
		return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
	}


	if header.ErrorCode != 0 {
		return DeviceError(header.ErrorCode)
	}

	bytes.NewBuffer(resultBytes[8:])
	

	return nil
}

// Returns the callback configuration as set by SetUVICallbackConfiguration.
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *UVLightV2Bricklet) GetUVICallbackConfiguration() (period uint32, valueHasToChange bool, option ThresholdOption, min int32, max int32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetUVICallbackConfiguration), buf.Bytes())
	if err != nil {
		return period, valueHasToChange, option, min, max, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 22 {
		return period, valueHasToChange, option, min, max, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 22)
	}


	if header.ErrorCode != 0 {
		return period, valueHasToChange, option, min, max, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &period)
	binary.Read(resultBuf, binary.LittleEndian, &valueHasToChange)
	binary.Read(resultBuf, binary.LittleEndian, &option)
	binary.Read(resultBuf, binary.LittleEndian, &min)
	binary.Read(resultBuf, binary.LittleEndian, &max)


	return period, valueHasToChange, option, min, max, nil
}

// Sets the configuration of the sensor. The integration time can be configured
// between 50 and 800 ms. With a shorter integration time the sensor reading updates
// more often but contains more noise. With a longer integration the sensor reading
// contains less noise but updates less often.
// 
// With a longer integration time (especially 800 ms) and a higher UV intensity the
// sensor can be saturated. If this happens the UVA/UVB/UVI readings are all -1.
// In this case you need to choose a shorter integration time.
//
// Associated constants:
//
//	* IntegrationTime50ms
//	* IntegrationTime100ms
//	* IntegrationTime200ms
//	* IntegrationTime400ms
//	* IntegrationTime800ms
func (device *UVLightV2Bricklet) SetConfiguration(integrationTime IntegrationTime) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, integrationTime);

	resultBytes, err := device.device.Set(uint8(FunctionSetConfiguration), buf.Bytes())
	if err != nil {
		return err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 8 {
		return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
	}


	if header.ErrorCode != 0 {
		return DeviceError(header.ErrorCode)
	}

	bytes.NewBuffer(resultBytes[8:])
	

	return nil
}

// Returns the configuration as set by SetConfiguration.
//
// Associated constants:
//
//	* IntegrationTime50ms
//	* IntegrationTime100ms
//	* IntegrationTime200ms
//	* IntegrationTime400ms
//	* IntegrationTime800ms
func (device *UVLightV2Bricklet) GetConfiguration() (integrationTime IntegrationTime, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetConfiguration), buf.Bytes())
	if err != nil {
		return integrationTime, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 9 {
		return integrationTime, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
	}


	if header.ErrorCode != 0 {
		return integrationTime, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &integrationTime)


	return integrationTime, nil
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
func (device *UVLightV2Bricklet) GetSPITFPErrorCount() (errorCountAckChecksum uint32, errorCountMessageChecksum uint32, errorCountFrame uint32, errorCountOverflow uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetSPITFPErrorCount), buf.Bytes())
	if err != nil {
		return errorCountAckChecksum, errorCountMessageChecksum, errorCountFrame, errorCountOverflow, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 24 {
		return errorCountAckChecksum, errorCountMessageChecksum, errorCountFrame, errorCountOverflow, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 24)
	}


	if header.ErrorCode != 0 {
		return errorCountAckChecksum, errorCountMessageChecksum, errorCountFrame, errorCountOverflow, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &errorCountAckChecksum)
	binary.Read(resultBuf, binary.LittleEndian, &errorCountMessageChecksum)
	binary.Read(resultBuf, binary.LittleEndian, &errorCountFrame)
	binary.Read(resultBuf, binary.LittleEndian, &errorCountOverflow)


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
func (device *UVLightV2Bricklet) SetBootloaderMode(mode BootloaderMode) (status BootloaderStatus, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, mode);

	resultBytes, err := device.device.Get(uint8(FunctionSetBootloaderMode), buf.Bytes())
	if err != nil {
		return status, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 9 {
		return status, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
	}


	if header.ErrorCode != 0 {
		return status, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &status)


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
func (device *UVLightV2Bricklet) GetBootloaderMode() (mode BootloaderMode, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetBootloaderMode), buf.Bytes())
	if err != nil {
		return mode, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 9 {
		return mode, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
	}


	if header.ErrorCode != 0 {
		return mode, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &mode)


	return mode, nil
}

// Sets the firmware pointer for WriteFirmware. The pointer has
// to be increased by chunks of size 64. The data is written to flash
// every 4 chunks (which equals to one page of size 256).
// 
// This function is used by Brick Viewer during flashing. It should not be
// necessary to call it in a normal user program.
func (device *UVLightV2Bricklet) SetWriteFirmwarePointer(pointer uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, pointer);

	resultBytes, err := device.device.Set(uint8(FunctionSetWriteFirmwarePointer), buf.Bytes())
	if err != nil {
		return err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 8 {
		return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
	}


	if header.ErrorCode != 0 {
		return DeviceError(header.ErrorCode)
	}

	bytes.NewBuffer(resultBytes[8:])
	

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
func (device *UVLightV2Bricklet) WriteFirmware(data [64]uint8) (status uint8, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, data);

	resultBytes, err := device.device.Get(uint8(FunctionWriteFirmware), buf.Bytes())
	if err != nil {
		return status, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 9 {
		return status, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
	}


	if header.ErrorCode != 0 {
		return status, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &status)


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
func (device *UVLightV2Bricklet) SetStatusLEDConfig(config StatusLEDConfig) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, config);

	resultBytes, err := device.device.Set(uint8(FunctionSetStatusLEDConfig), buf.Bytes())
	if err != nil {
		return err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 8 {
		return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
	}


	if header.ErrorCode != 0 {
		return DeviceError(header.ErrorCode)
	}

	bytes.NewBuffer(resultBytes[8:])
	

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
func (device *UVLightV2Bricklet) GetStatusLEDConfig() (config StatusLEDConfig, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetStatusLEDConfig), buf.Bytes())
	if err != nil {
		return config, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 9 {
		return config, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
	}


	if header.ErrorCode != 0 {
		return config, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &config)


	return config, nil
}

// Returns the temperature as measured inside the microcontroller. The
// value returned is not the ambient temperature!
// 
// The temperature is only proportional to the real temperature and it has bad
// accuracy. Practically it is only useful as an indicator for
// temperature changes.
func (device *UVLightV2Bricklet) GetChipTemperature() (temperature int16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetChipTemperature), buf.Bytes())
	if err != nil {
		return temperature, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 10 {
		return temperature, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
	}


	if header.ErrorCode != 0 {
		return temperature, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &temperature)


	return temperature, nil
}

// Calling this function will reset the Bricklet. All configurations
// will be lost.
// 
// After a reset you have to create new device objects,
// calling functions on the existing ones will result in
// undefined behavior!
func (device *UVLightV2Bricklet) Reset() (err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Set(uint8(FunctionReset), buf.Bytes())
	if err != nil {
		return err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 8 {
		return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
	}


	if header.ErrorCode != 0 {
		return DeviceError(header.ErrorCode)
	}

	bytes.NewBuffer(resultBytes[8:])
	

	return nil
}

// Writes a new UID into flash. If you want to set a new UID
// you have to decode the Base58 encoded UID string into an
// integer first.
// 
// We recommend that you use Brick Viewer to change the UID.
func (device *UVLightV2Bricklet) WriteUID(uid uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, uid);

	resultBytes, err := device.device.Set(uint8(FunctionWriteUID), buf.Bytes())
	if err != nil {
		return err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 8 {
		return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
	}


	if header.ErrorCode != 0 {
		return DeviceError(header.ErrorCode)
	}

	bytes.NewBuffer(resultBytes[8:])
	

	return nil
}

// Returns the current UID as an integer. Encode as
// Base58 to get the usual string version.
func (device *UVLightV2Bricklet) ReadUID() (uid uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionReadUID), buf.Bytes())
	if err != nil {
		return uid, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 12 {
		return uid, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
	}


	if header.ErrorCode != 0 {
		return uid, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &uid)


	return uid, nil
}

// Returns the UID, the UID where the Bricklet is connected to,
// the position, the hardware and firmware version as well as the
// device identifier.
// 
// The position can be 'a', 'b', 'c', 'd', 'e', 'f', 'g' or 'h' (Bricklet Port).
// The Raspberry Pi HAT (Zero) Brick is always at position 'i' and the Bricklet
// connected to an `Isolator Bricklet <isolator_bricklet>` is always as
// position 'z'.
// 
// The device identifier numbers can be found `here <device_identifier>`.
// |device_identifier_constant|
func (device *UVLightV2Bricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetIdentity), buf.Bytes())
	if err != nil {
		return uid, connectedUid, position, hardwareVersion, firmwareVersion, deviceIdentifier, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 33 {
		return uid, connectedUid, position, hardwareVersion, firmwareVersion, deviceIdentifier, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 33)
	}


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


	return uid, connectedUid, position, hardwareVersion, firmwareVersion, deviceIdentifier, nil
}
