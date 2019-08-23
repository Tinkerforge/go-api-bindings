/* ***********************************************************
 * This file was automatically generated on 2019-08-23.      *
 *                                                           *
 * Go Bindings Version 2.0.4                                 *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


// Measures distance up to 150cm with infrared light.
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/DistanceIRV2_Bricklet_Go.html.
package distance_ir_v2_bricklet

import (
	"encoding/binary"
	"bytes"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionGetDistance Function = 1
	FunctionSetDistanceCallbackConfiguration Function = 2
	FunctionGetDistanceCallbackConfiguration Function = 3
	FunctionGetAnalogValue Function = 5
	FunctionSetAnalogValueCallbackConfiguration Function = 6
	FunctionGetAnalogValueCallbackConfiguration Function = 7
	FunctionSetMovingAverageConfiguration Function = 9
	FunctionGetMovingAverageConfiguration Function = 10
	FunctionSetDistanceLEDConfig Function = 11
	FunctionGetDistanceLEDConfig Function = 12
	FunctionSetSensorType Function = 13
	FunctionGetSensorType Function = 14
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
	FunctionCallbackDistance Function = 4
	FunctionCallbackAnalogValue Function = 8
)

type ThresholdOption = rune

const (
	ThresholdOptionOff ThresholdOption = 'x'
	ThresholdOptionOutside ThresholdOption = 'o'
	ThresholdOptionInside ThresholdOption = 'i'
	ThresholdOptionSmaller ThresholdOption = '<'
	ThresholdOptionGreater ThresholdOption = '>'
)

type DistanceLEDConfig = uint8

const (
	DistanceLEDConfigOff DistanceLEDConfig = 0
	DistanceLEDConfigOn DistanceLEDConfig = 1
	DistanceLEDConfigShowHeartbeat DistanceLEDConfig = 2
	DistanceLEDConfigShowDistance DistanceLEDConfig = 3
)

type SensorType = uint8

const (
	SensorType2Y0A41 SensorType = 0
	SensorType2Y0A21 SensorType = 1
	SensorType2Y0A02 SensorType = 2
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

type DistanceIRV2Bricklet struct {
	device Device
}
const DeviceIdentifier = 2125
const DeviceDisplayName = "Distance IR Bricklet 2.0"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (DistanceIRV2Bricklet, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{ 2,0,1 }, uid, &internalIPCon, 0)
	if err != nil {
		return DistanceIRV2Bricklet{}, err
	}
	dev.ResponseExpected[FunctionGetDistance] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetDistanceCallbackConfiguration] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetDistanceCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetAnalogValue] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetAnalogValueCallbackConfiguration] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetAnalogValueCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetMovingAverageConfiguration] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetMovingAverageConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetDistanceLEDConfig] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetDistanceLEDConfig] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetSensorType] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetSensorType] = ResponseExpectedFlagAlwaysTrue;
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
	return DistanceIRV2Bricklet{dev}, nil
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
func (device *DistanceIRV2Bricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *DistanceIRV2Bricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *DistanceIRV2Bricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *DistanceIRV2Bricklet) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// This callback is triggered periodically according to the configuration set by
// SetDistanceCallbackConfiguration.
// 
// The parameter is the same as GetDistance.
func (device *DistanceIRV2Bricklet) RegisterDistanceCallback(fn func(uint16)) uint64 {
	wrapper := func(byteSlice []byte) {
		buf := bytes.NewBuffer(byteSlice[8:])
		var distance uint16
		binary.Read(buf, binary.LittleEndian, &distance)
		fn(distance)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackDistance), wrapper)
}

// Remove a registered Distance callback.
func (device *DistanceIRV2Bricklet) DeregisterDistanceCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackDistance), registrationId)
}


// This callback is triggered periodically according to the configuration set by
// SetAnalogValueCallbackConfiguration.
// 
// The parameter is the same as GetAnalogValue.
func (device *DistanceIRV2Bricklet) RegisterAnalogValueCallback(fn func(uint32)) uint64 {
	wrapper := func(byteSlice []byte) {
		buf := bytes.NewBuffer(byteSlice[8:])
		var analogValue uint32
		binary.Read(buf, binary.LittleEndian, &analogValue)
		fn(analogValue)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackAnalogValue), wrapper)
}

// Remove a registered Analog Value callback.
func (device *DistanceIRV2Bricklet) DeregisterAnalogValueCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackAnalogValue), registrationId)
}


// Returns the distance measured by the sensor. The value is in mm and possible
// distance ranges are 40 to 300, 100 to 800 and 200 to 1500, depending on the
// selected IR sensor.
// 
// 
// If you want to get the value periodically, it is recommended to use the
// RegisterDistanceCallback callback. You can set the callback configuration
// with SetDistanceCallbackConfiguration.
func (device *DistanceIRV2Bricklet) GetDistance() (distance uint16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetDistance), buf.Bytes())
	if err != nil {
		return distance, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return distance, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &distance)

	}

	return distance, nil
}

// The period in ms is the period with which the RegisterDistanceCallback callback is triggered
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
// The `option`-parameter together with min/max sets a threshold for the RegisterDistanceCallback callback.
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
func (device *DistanceIRV2Bricklet) SetDistanceCallbackConfiguration(period uint32, valueHasToChange bool, option ThresholdOption, min uint16, max uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);
	binary.Write(&buf, binary.LittleEndian, valueHasToChange);
	binary.Write(&buf, binary.LittleEndian, option);
	binary.Write(&buf, binary.LittleEndian, min);
	binary.Write(&buf, binary.LittleEndian, max);

	resultBytes, err := device.device.Set(uint8(FunctionSetDistanceCallbackConfiguration), buf.Bytes())
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

// Returns the callback configuration as set by SetDistanceCallbackConfiguration.
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *DistanceIRV2Bricklet) GetDistanceCallbackConfiguration() (period uint32, valueHasToChange bool, option ThresholdOption, min uint16, max uint16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetDistanceCallbackConfiguration), buf.Bytes())
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

// Returns the analog value as read by a analog-to-digital converter.
// The value has 21 bit with a range of 0 to 2097151.
// 
// This is unfiltered raw data. We made sure that the integration time
// of the ADC is shorter then the measurement interval of the sensor
// (10ms vs 16.5ms). So there is no information lost.
// 
// If you want to do your own calibration or create your own lookup table
// you can use this value.
// 
// 
// If you want to get the value periodically, it is recommended to use the
// RegisterAnalogValueCallback callback. You can set the callback configuration
// with SetAnalogValueCallbackConfiguration.
func (device *DistanceIRV2Bricklet) GetAnalogValue() (analogValue uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetAnalogValue), buf.Bytes())
	if err != nil {
		return analogValue, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return analogValue, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &analogValue)

	}

	return analogValue, nil
}

// The period in ms is the period with which the RegisterAnalogValueCallback callback is triggered
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
// The `option`-parameter together with min/max sets a threshold for the RegisterAnalogValueCallback callback.
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
func (device *DistanceIRV2Bricklet) SetAnalogValueCallbackConfiguration(period uint32, valueHasToChange bool, option ThresholdOption, min uint32, max uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);
	binary.Write(&buf, binary.LittleEndian, valueHasToChange);
	binary.Write(&buf, binary.LittleEndian, option);
	binary.Write(&buf, binary.LittleEndian, min);
	binary.Write(&buf, binary.LittleEndian, max);

	resultBytes, err := device.device.Set(uint8(FunctionSetAnalogValueCallbackConfiguration), buf.Bytes())
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

// Returns the callback configuration as set by SetAnalogValueCallbackConfiguration.
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *DistanceIRV2Bricklet) GetAnalogValueCallbackConfiguration() (period uint32, valueHasToChange bool, option ThresholdOption, min uint32, max uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetAnalogValueCallbackConfiguration), buf.Bytes())
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
// for the distance.
// 
// Setting the length to 1 will turn the averaging off. With less averaging, there
// is more noise on the data.
// 
// The range for the averaging is 1-1000.
// 
// New data is gathered every ~10ms. With a moving average of length 1000 the
// resulting averaging window has a length of approximately 10s. If you want to do
// long term measurements the longest moving average will give the cleanest results.
// 
// The default value is 25.
func (device *DistanceIRV2Bricklet) SetMovingAverageConfiguration(movingAverageLength uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, movingAverageLength);

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

// Returns the moving average configuration as set by SetMovingAverageConfiguration.
func (device *DistanceIRV2Bricklet) GetMovingAverageConfiguration() (movingAverageLength uint16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetMovingAverageConfiguration), buf.Bytes())
	if err != nil {
		return movingAverageLength, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return movingAverageLength, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &movingAverageLength)

	}

	return movingAverageLength, nil
}

// Configures the distance LED to be either turned off, turned on, blink in
// heartbeat mode or show the distance (brighter = object is nearer).
// 
// The default value is 3 (show distance).
//
// Associated constants:
//
//	* DistanceLEDConfigOff
//	* DistanceLEDConfigOn
//	* DistanceLEDConfigShowHeartbeat
//	* DistanceLEDConfigShowDistance
func (device *DistanceIRV2Bricklet) SetDistanceLEDConfig(config DistanceLEDConfig) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, config);

	resultBytes, err := device.device.Set(uint8(FunctionSetDistanceLEDConfig), buf.Bytes())
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

// Returns the LED configuration as set by SetDistanceLEDConfig
//
// Associated constants:
//
//	* DistanceLEDConfigOff
//	* DistanceLEDConfigOn
//	* DistanceLEDConfigShowHeartbeat
//	* DistanceLEDConfigShowDistance
func (device *DistanceIRV2Bricklet) GetDistanceLEDConfig() (config DistanceLEDConfig, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetDistanceLEDConfig), buf.Bytes())
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

// Sets the sensor type.
// 
// The Bricklet comes configured with the correct sensor type
// and the type is saved in flash (i.e. the Bricklet retains the information if
// power is lost).
// 
// If you want to change the sensor you can set the type in Brick Viewer,
// you will likely never need to call this function from your program.
//
// Associated constants:
//
//	* SensorType2Y0A41
//	* SensorType2Y0A21
//	* SensorType2Y0A02
func (device *DistanceIRV2Bricklet) SetSensorType(sensor SensorType) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, sensor);

	resultBytes, err := device.device.Set(uint8(FunctionSetSensorType), buf.Bytes())
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

// Returns the sensor type as set by SetSensorType.
//
// Associated constants:
//
//	* SensorType2Y0A41
//	* SensorType2Y0A21
//	* SensorType2Y0A02
func (device *DistanceIRV2Bricklet) GetSensorType() (sensor SensorType, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetSensorType), buf.Bytes())
	if err != nil {
		return sensor, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return sensor, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &sensor)

	}

	return sensor, nil
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
func (device *DistanceIRV2Bricklet) GetSPITFPErrorCount() (errorCountAckChecksum uint32, errorCountMessageChecksum uint32, errorCountFrame uint32, errorCountOverflow uint32, err error) {
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
func (device *DistanceIRV2Bricklet) SetBootloaderMode(mode BootloaderMode) (status BootloaderStatus, err error) {
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
func (device *DistanceIRV2Bricklet) GetBootloaderMode() (mode BootloaderMode, err error) {
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
func (device *DistanceIRV2Bricklet) SetWriteFirmwarePointer(pointer uint32) (err error) {
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
func (device *DistanceIRV2Bricklet) WriteFirmware(data [64]uint8) (status uint8, err error) {
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
func (device *DistanceIRV2Bricklet) SetStatusLEDConfig(config StatusLEDConfig) (err error) {
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
func (device *DistanceIRV2Bricklet) GetStatusLEDConfig() (config StatusLEDConfig, err error) {
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
func (device *DistanceIRV2Bricklet) GetChipTemperature() (temperature int16, err error) {
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
func (device *DistanceIRV2Bricklet) Reset() (err error) {
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
func (device *DistanceIRV2Bricklet) WriteUID(uid uint32) (err error) {
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
func (device *DistanceIRV2Bricklet) ReadUID() (uid uint32, err error) {
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
func (device *DistanceIRV2Bricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
