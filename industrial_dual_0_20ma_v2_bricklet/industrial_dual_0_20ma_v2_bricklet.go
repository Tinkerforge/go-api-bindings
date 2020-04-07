/* ***********************************************************
 * This file was automatically generated on 2020-04-07.      *
 *                                                           *
 * Go Bindings Version 2.0.6                                 *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


// Measures two DC currents between 0mA and 20mA (IEC 60381-1).
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/IndustrialDual020mAV2_Bricklet_Go.html.
package industrial_dual_0_20ma_v2_bricklet

import (
	"encoding/binary"
	"bytes"
	"fmt"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionGetCurrent Function = 1
	FunctionSetCurrentCallbackConfiguration Function = 2
	FunctionGetCurrentCallbackConfiguration Function = 3
	FunctionSetSampleRate Function = 5
	FunctionGetSampleRate Function = 6
	FunctionSetGain Function = 7
	FunctionGetGain Function = 8
	FunctionSetChannelLEDConfig Function = 9
	FunctionGetChannelLEDConfig Function = 10
	FunctionSetChannelLEDStatusConfig Function = 11
	FunctionGetChannelLEDStatusConfig Function = 12
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
	FunctionCallbackCurrent Function = 4
)

type ThresholdOption = rune

const (
	ThresholdOptionOff ThresholdOption = 'x'
	ThresholdOptionOutside ThresholdOption = 'o'
	ThresholdOptionInside ThresholdOption = 'i'
	ThresholdOptionSmaller ThresholdOption = '<'
	ThresholdOptionGreater ThresholdOption = '>'
)

type SampleRate = uint8

const (
	SampleRate240SPS SampleRate = 0
	SampleRate60SPS SampleRate = 1
	SampleRate15SPS SampleRate = 2
	SampleRate4SPS SampleRate = 3
)

type Gain = uint8

const (
	Gain1x Gain = 0
	Gain2x Gain = 1
	Gain4x Gain = 2
	Gain8x Gain = 3
)

type ChannelLEDConfig = uint8

const (
	ChannelLEDConfigOff ChannelLEDConfig = 0
	ChannelLEDConfigOn ChannelLEDConfig = 1
	ChannelLEDConfigShowHeartbeat ChannelLEDConfig = 2
	ChannelLEDConfigShowChannelStatus ChannelLEDConfig = 3
)

type ChannelLEDStatusConfig = uint8

const (
	ChannelLEDStatusConfigThreshold ChannelLEDStatusConfig = 0
	ChannelLEDStatusConfigIntensity ChannelLEDStatusConfig = 1
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

type IndustrialDual020mAV2Bricklet struct {
	device Device
}
const DeviceIdentifier = 2120
const DeviceDisplayName = "Industrial Dual 0-20mA Bricklet 2.0"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (IndustrialDual020mAV2Bricklet, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{ 2,0,0 }, uid, &internalIPCon, 0, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return IndustrialDual020mAV2Bricklet{}, err
	}
	dev.ResponseExpected[FunctionGetCurrent] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetCurrentCallbackConfiguration] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetCurrentCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetSampleRate] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetSampleRate] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetGain] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetGain] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetChannelLEDConfig] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetChannelLEDConfig] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetChannelLEDStatusConfig] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetChannelLEDStatusConfig] = ResponseExpectedFlagAlwaysTrue;
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
	return IndustrialDual020mAV2Bricklet{dev}, nil
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
func (device *IndustrialDual020mAV2Bricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *IndustrialDual020mAV2Bricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *IndustrialDual020mAV2Bricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *IndustrialDual020mAV2Bricklet) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// This callback is triggered periodically according to the configuration set by
// SetCurrentCallbackConfiguration.
// 
// The parameter is the same as GetCurrent.
func (device *IndustrialDual020mAV2Bricklet) RegisterCurrentCallback(fn func(uint8, int32)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 13 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var channel uint8
		var current int32
		binary.Read(buf, binary.LittleEndian, &channel)
		binary.Read(buf, binary.LittleEndian, &current)
		fn(channel, current)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackCurrent), wrapper)
}

// Remove a registered Current callback.
func (device *IndustrialDual020mAV2Bricklet) DeregisterCurrentCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackCurrent), registrationId)
}


// Returns the current of the specified channel.
// 
// It is possible to detect if an IEC 60381-1 compatible sensor is connected
// and if it works probably.
// 
// If the returned current is below 4mA, there is likely no sensor connected
// or the connected sensor is defective. If the returned current is over 20mA,
// there might be a short circuit or the sensor is defective.
// 
// 
// If you want to get the value periodically, it is recommended to use the
// RegisterCurrentCallback callback. You can set the callback configuration
// with SetCurrentCallbackConfiguration.
func (device *IndustrialDual020mAV2Bricklet) GetCurrent(channel uint8) (current int32, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, channel);

	resultBytes, err := device.device.Get(uint8(FunctionGetCurrent), buf.Bytes())
	if err != nil {
		return current, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 12 {
		return current, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
	}


	if header.ErrorCode != 0 {
		return current, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &current)


	return current, nil
}

// The period is the period with which the RegisterCurrentCallback callback is triggered
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
// The `option`-parameter together with min/max sets a threshold for the RegisterCurrentCallback callback.
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
func (device *IndustrialDual020mAV2Bricklet) SetCurrentCallbackConfiguration(channel uint8, period uint32, valueHasToChange bool, option ThresholdOption, min int32, max int32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, channel);
	binary.Write(&buf, binary.LittleEndian, period);
	binary.Write(&buf, binary.LittleEndian, valueHasToChange);
	binary.Write(&buf, binary.LittleEndian, option);
	binary.Write(&buf, binary.LittleEndian, min);
	binary.Write(&buf, binary.LittleEndian, max);

	resultBytes, err := device.device.Set(uint8(FunctionSetCurrentCallbackConfiguration), buf.Bytes())
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

// Returns the callback configuration as set by SetCurrentCallbackConfiguration.
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *IndustrialDual020mAV2Bricklet) GetCurrentCallbackConfiguration(channel uint8) (period uint32, valueHasToChange bool, option ThresholdOption, min int32, max int32, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, channel);

	resultBytes, err := device.device.Get(uint8(FunctionGetCurrentCallbackConfiguration), buf.Bytes())
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

// Sets the sample rate to either 240, 60, 15 or 4 samples per second.
// The resolution for the rates is 12, 14, 16 and 18 bit respectively.
// 
//  Value| Description
//  --- | --- 
//  0|    240 samples per second| 12 bit resolution
//  1|    60 samples per second| 14 bit resolution
//  2|    15 samples per second| 16 bit resolution
//  3|    4 samples per second| 18 bit resolution
//
// Associated constants:
//
//	* SampleRate240SPS
//	* SampleRate60SPS
//	* SampleRate15SPS
//	* SampleRate4SPS
func (device *IndustrialDual020mAV2Bricklet) SetSampleRate(rate SampleRate) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, rate);

	resultBytes, err := device.device.Set(uint8(FunctionSetSampleRate), buf.Bytes())
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

// Returns the gain as set by SetSampleRate.
//
// Associated constants:
//
//	* SampleRate240SPS
//	* SampleRate60SPS
//	* SampleRate15SPS
//	* SampleRate4SPS
func (device *IndustrialDual020mAV2Bricklet) GetSampleRate() (rate SampleRate, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetSampleRate), buf.Bytes())
	if err != nil {
		return rate, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 9 {
		return rate, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
	}


	if header.ErrorCode != 0 {
		return rate, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &rate)


	return rate, nil
}

// Sets a gain between 1x and 8x. If you want to measure a very small current,
// you can increase the gain to get some more resolution.
// 
// Example: If you measure 0.5mA with a gain of 8x the return value will be
// 4mA.
//
// Associated constants:
//
//	* Gain1x
//	* Gain2x
//	* Gain4x
//	* Gain8x
func (device *IndustrialDual020mAV2Bricklet) SetGain(gain Gain) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, gain);

	resultBytes, err := device.device.Set(uint8(FunctionSetGain), buf.Bytes())
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

// Returns the gain as set by SetGain.
//
// Associated constants:
//
//	* Gain1x
//	* Gain2x
//	* Gain4x
//	* Gain8x
func (device *IndustrialDual020mAV2Bricklet) GetGain() (gain Gain, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetGain), buf.Bytes())
	if err != nil {
		return gain, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 9 {
		return gain, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
	}


	if header.ErrorCode != 0 {
		return gain, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &gain)


	return gain, nil
}

// Each channel has a corresponding LED. You can turn the LED off, on or show a
// heartbeat. You can also set the LED to Channel Status. In this mode the
// LED can either be turned on with a pre-defined threshold or the intensity
// of the LED can change with the measured value.
// 
// You can configure the channel status behavior with SetChannelLEDStatusConfig.
//
// Associated constants:
//
//	* ChannelLEDConfigOff
//	* ChannelLEDConfigOn
//	* ChannelLEDConfigShowHeartbeat
//	* ChannelLEDConfigShowChannelStatus
func (device *IndustrialDual020mAV2Bricklet) SetChannelLEDConfig(channel uint8, config ChannelLEDConfig) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, channel);
	binary.Write(&buf, binary.LittleEndian, config);

	resultBytes, err := device.device.Set(uint8(FunctionSetChannelLEDConfig), buf.Bytes())
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

// Returns the channel LED configuration as set by SetChannelLEDConfig
//
// Associated constants:
//
//	* ChannelLEDConfigOff
//	* ChannelLEDConfigOn
//	* ChannelLEDConfigShowHeartbeat
//	* ChannelLEDConfigShowChannelStatus
func (device *IndustrialDual020mAV2Bricklet) GetChannelLEDConfig(channel uint8) (config ChannelLEDConfig, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, channel);

	resultBytes, err := device.device.Get(uint8(FunctionGetChannelLEDConfig), buf.Bytes())
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

// Sets the channel LED status config. This config is used if the channel LED is
// configured as Channel Status, see SetChannelLEDConfig.
// 
// For each channel you can choose between threshold and intensity mode.
// 
// In threshold mode you can define a positive or a negative threshold.
// For a positive threshold set the min parameter to the threshold value in nA
// above which the LED should turn on and set the max parameter to 0. Example:
// If you set a positive threshold of 10mA, the LED will turn on as soon as the
// current exceeds 10mA and turn off again if it goes below 10mA.
// For a negative threshold set the max parameter to the threshold value in nA
// below which the LED should turn on and set the min parameter to 0. Example:
// If you set a negative threshold of 10mA, the LED will turn on as soon as the
// current goes below 10mA and the LED will turn off when the current exceeds 10mA.
// 
// In intensity mode you can define a range in nA that is used to scale the brightness
// of the LED. Example with min=4mA and max=20mA: The LED is off at 4mA and below,
// on at 20mA and above and the brightness is linearly scaled between the values
// 4mA and 20mA. If the min value is greater than the max value, the LED brightness
// is scaled the other way around.
//
// Associated constants:
//
//	* ChannelLEDStatusConfigThreshold
//	* ChannelLEDStatusConfigIntensity
func (device *IndustrialDual020mAV2Bricklet) SetChannelLEDStatusConfig(channel uint8, min int32, max int32, config ChannelLEDStatusConfig) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, channel);
	binary.Write(&buf, binary.LittleEndian, min);
	binary.Write(&buf, binary.LittleEndian, max);
	binary.Write(&buf, binary.LittleEndian, config);

	resultBytes, err := device.device.Set(uint8(FunctionSetChannelLEDStatusConfig), buf.Bytes())
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

// Returns the channel LED status configuration as set by
// SetChannelLEDStatusConfig.
//
// Associated constants:
//
//	* ChannelLEDStatusConfigThreshold
//	* ChannelLEDStatusConfigIntensity
func (device *IndustrialDual020mAV2Bricklet) GetChannelLEDStatusConfig(channel uint8) (min int32, max int32, config ChannelLEDStatusConfig, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, channel);

	resultBytes, err := device.device.Get(uint8(FunctionGetChannelLEDStatusConfig), buf.Bytes())
	if err != nil {
		return min, max, config, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 17 {
		return min, max, config, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 17)
	}


	if header.ErrorCode != 0 {
		return min, max, config, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &min)
	binary.Read(resultBuf, binary.LittleEndian, &max)
	binary.Read(resultBuf, binary.LittleEndian, &config)


	return min, max, config, nil
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
func (device *IndustrialDual020mAV2Bricklet) GetSPITFPErrorCount() (errorCountAckChecksum uint32, errorCountMessageChecksum uint32, errorCountFrame uint32, errorCountOverflow uint32, err error) {
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
func (device *IndustrialDual020mAV2Bricklet) SetBootloaderMode(mode BootloaderMode) (status BootloaderStatus, err error) {
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
func (device *IndustrialDual020mAV2Bricklet) GetBootloaderMode() (mode BootloaderMode, err error) {
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
func (device *IndustrialDual020mAV2Bricklet) SetWriteFirmwarePointer(pointer uint32) (err error) {
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
func (device *IndustrialDual020mAV2Bricklet) WriteFirmware(data [64]uint8) (status uint8, err error) {
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
func (device *IndustrialDual020mAV2Bricklet) SetStatusLEDConfig(config StatusLEDConfig) (err error) {
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
func (device *IndustrialDual020mAV2Bricklet) GetStatusLEDConfig() (config StatusLEDConfig, err error) {
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
func (device *IndustrialDual020mAV2Bricklet) GetChipTemperature() (temperature int16, err error) {
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
func (device *IndustrialDual020mAV2Bricklet) Reset() (err error) {
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
func (device *IndustrialDual020mAV2Bricklet) WriteUID(uid uint32) (err error) {
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
func (device *IndustrialDual020mAV2Bricklet) ReadUID() (uid uint32, err error) {
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
func (device *IndustrialDual020mAV2Bricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
