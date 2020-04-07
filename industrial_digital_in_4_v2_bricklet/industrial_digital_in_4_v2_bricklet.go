/* ***********************************************************
 * This file was automatically generated on 2020-04-07.      *
 *                                                           *
 * Go Bindings Version 2.0.6                                 *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


// 4 galvanically isolated digital inputs.
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/IndustrialDigitalIn4V2_Bricklet_Go.html.
package industrial_digital_in_4_v2_bricklet

import (
	"encoding/binary"
	"bytes"
	"fmt"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionGetValue Function = 1
	FunctionSetValueCallbackConfiguration Function = 2
	FunctionGetValueCallbackConfiguration Function = 3
	FunctionSetAllValueCallbackConfiguration Function = 4
	FunctionGetAllValueCallbackConfiguration Function = 5
	FunctionGetEdgeCount Function = 6
	FunctionSetEdgeCountConfiguration Function = 7
	FunctionGetEdgeCountConfiguration Function = 8
	FunctionSetChannelLEDConfig Function = 9
	FunctionGetChannelLEDConfig Function = 10
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
	FunctionCallbackValue Function = 11
	FunctionCallbackAllValue Function = 12
)

type Channel = uint8

const (
	Channel0 Channel = 0
	Channel1 Channel = 1
	Channel2 Channel = 2
	Channel3 Channel = 3
)

type EdgeType = uint8

const (
	EdgeTypeRising EdgeType = 0
	EdgeTypeFalling EdgeType = 1
	EdgeTypeBoth EdgeType = 2
)

type ChannelLEDConfig = uint8

const (
	ChannelLEDConfigOff ChannelLEDConfig = 0
	ChannelLEDConfigOn ChannelLEDConfig = 1
	ChannelLEDConfigShowHeartbeat ChannelLEDConfig = 2
	ChannelLEDConfigShowChannelStatus ChannelLEDConfig = 3
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

type IndustrialDigitalIn4V2Bricklet struct {
	device Device
}
const DeviceIdentifier = 2100
const DeviceDisplayName = "Industrial Digital In 4 Bricklet 2.0"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (IndustrialDigitalIn4V2Bricklet, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{ 2,0,0 }, uid, &internalIPCon, 0, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return IndustrialDigitalIn4V2Bricklet{}, err
	}
	dev.ResponseExpected[FunctionGetValue] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetValueCallbackConfiguration] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetValueCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetAllValueCallbackConfiguration] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetAllValueCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetEdgeCount] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetEdgeCountConfiguration] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetEdgeCountConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetChannelLEDConfig] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetChannelLEDConfig] = ResponseExpectedFlagAlwaysTrue;
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
	return IndustrialDigitalIn4V2Bricklet{dev}, nil
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
func (device *IndustrialDigitalIn4V2Bricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *IndustrialDigitalIn4V2Bricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *IndustrialDigitalIn4V2Bricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *IndustrialDigitalIn4V2Bricklet) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// This callback is triggered periodically according to the configuration set by
// SetValueCallbackConfiguration.
// 
// The parameters are the channel, a value-changed indicator and the actual
// value for the channel. The `changed` parameter is true if the value has changed
// since the last callback.
func (device *IndustrialDigitalIn4V2Bricklet) RegisterValueCallback(fn func(Channel, bool, bool)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 11 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var channel Channel
		var changed bool
		var value bool
		binary.Read(buf, binary.LittleEndian, &channel)
		binary.Read(buf, binary.LittleEndian, &changed)
		binary.Read(buf, binary.LittleEndian, &value)
		fn(channel, changed, value)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackValue), wrapper)
}

// Remove a registered Value callback.
func (device *IndustrialDigitalIn4V2Bricklet) DeregisterValueCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackValue), registrationId)
}


// This callback is triggered periodically according to the configuration set by
// SetAllValueCallbackConfiguration.
// 
// The parameters are the same as GetValue. Additional the
// `changed` parameter is true if the value has changed since
// the last callback.
func (device *IndustrialDigitalIn4V2Bricklet) RegisterAllValueCallback(fn func([4]bool, [4]bool)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 10 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var changed [4]bool
		var value [4]bool
		binary.Read(buf, binary.LittleEndian, &changed)
		binary.Read(buf, binary.LittleEndian, &value)
		fn(changed, value)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackAllValue), wrapper)
}

// Remove a registered All Value callback.
func (device *IndustrialDigitalIn4V2Bricklet) DeregisterAllValueCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackAllValue), registrationId)
}


// Returns the input value as bools, *true* refers to high and *false* refers to low.
func (device *IndustrialDigitalIn4V2Bricklet) GetValue() (value [4]bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetValue), buf.Bytes())
	if err != nil {
		return value, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 9 {
		return value, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
	}


	if header.ErrorCode != 0 {
		return value, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &value)


	return value, nil
}

// This callback can be configured per channel.
// 
// The period is the period with which the RegisterValueCallback
// callback is triggered periodically. A value of 0 turns the callback off.
// 
// If the `value has to change`-parameter is set to true, the callback is only
// triggered after the value has changed. If the value didn't change within the
// period, the callback is triggered immediately on change.
// 
// If it is set to false, the callback is continuously triggered with the period,
// independent of the value.
//
// Associated constants:
//
//	* Channel0
//	* Channel1
//	* Channel2
//	* Channel3
func (device *IndustrialDigitalIn4V2Bricklet) SetValueCallbackConfiguration(channel Channel, period uint32, valueHasToChange bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, channel);
	binary.Write(&buf, binary.LittleEndian, period);
	binary.Write(&buf, binary.LittleEndian, valueHasToChange);

	resultBytes, err := device.device.Set(uint8(FunctionSetValueCallbackConfiguration), buf.Bytes())
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

// Returns the callback configuration for the given channel as set by
// SetValueCallbackConfiguration.
//
// Associated constants:
//
//	* Channel0
//	* Channel1
//	* Channel2
//	* Channel3
func (device *IndustrialDigitalIn4V2Bricklet) GetValueCallbackConfiguration(channel Channel) (period uint32, valueHasToChange bool, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, channel);

	resultBytes, err := device.device.Get(uint8(FunctionGetValueCallbackConfiguration), buf.Bytes())
	if err != nil {
		return period, valueHasToChange, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 13 {
		return period, valueHasToChange, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 13)
	}


	if header.ErrorCode != 0 {
		return period, valueHasToChange, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &period)
	binary.Read(resultBuf, binary.LittleEndian, &valueHasToChange)


	return period, valueHasToChange, nil
}

// The period is the period with which the RegisterAllValueCallback
// callback is triggered periodically. A value of 0 turns the callback off.
// 
// If the `value has to change`-parameter is set to true, the callback is only
// triggered after the value has changed. If the value didn't change within the
// period, the callback is triggered immediately on change.
// 
// If it is set to false, the callback is continuously triggered with the period,
// independent of the value.
func (device *IndustrialDigitalIn4V2Bricklet) SetAllValueCallbackConfiguration(period uint32, valueHasToChange bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);
	binary.Write(&buf, binary.LittleEndian, valueHasToChange);

	resultBytes, err := device.device.Set(uint8(FunctionSetAllValueCallbackConfiguration), buf.Bytes())
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

// Returns the callback configuration as set by
// SetAllValueCallbackConfiguration.
func (device *IndustrialDigitalIn4V2Bricklet) GetAllValueCallbackConfiguration() (period uint32, valueHasToChange bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetAllValueCallbackConfiguration), buf.Bytes())
	if err != nil {
		return period, valueHasToChange, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 13 {
		return period, valueHasToChange, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 13)
	}


	if header.ErrorCode != 0 {
		return period, valueHasToChange, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &period)
	binary.Read(resultBuf, binary.LittleEndian, &valueHasToChange)


	return period, valueHasToChange, nil
}

// Returns the current value of the edge counter for the selected channel. You can
// configure the edges that are counted with SetEdgeCountConfiguration.
// 
// If you set the reset counter to *true*, the count is set back to 0
// directly after it is read.
//
// Associated constants:
//
//	* Channel0
//	* Channel1
//	* Channel2
//	* Channel3
func (device *IndustrialDigitalIn4V2Bricklet) GetEdgeCount(channel Channel, resetCounter bool) (count uint32, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, channel);
	binary.Write(&buf, binary.LittleEndian, resetCounter);

	resultBytes, err := device.device.Get(uint8(FunctionGetEdgeCount), buf.Bytes())
	if err != nil {
		return count, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 12 {
		return count, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
	}


	if header.ErrorCode != 0 {
		return count, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &count)


	return count, nil
}

// Configures the edge counter for a specific channel.
// 
// The edge type parameter configures if rising edges, falling edges or both are
// counted. Possible edge types are:
// 
// * 0 = rising
// * 1 = falling
// * 2 = both
// 
// Configuring an edge counter resets its value to 0.
// 
// If you don't know what any of this means, just leave it at default. The
// default configuration is very likely OK for you.
//
// Associated constants:
//
//	* Channel0
//	* Channel1
//	* Channel2
//	* Channel3
//	* EdgeTypeRising
//	* EdgeTypeFalling
//	* EdgeTypeBoth
func (device *IndustrialDigitalIn4V2Bricklet) SetEdgeCountConfiguration(channel Channel, edgeType EdgeType, debounce uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, channel);
	binary.Write(&buf, binary.LittleEndian, edgeType);
	binary.Write(&buf, binary.LittleEndian, debounce);

	resultBytes, err := device.device.Set(uint8(FunctionSetEdgeCountConfiguration), buf.Bytes())
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

// Returns the edge type and debounce time for the selected channel as set by
// SetEdgeCountConfiguration.
//
// Associated constants:
//
//	* Channel0
//	* Channel1
//	* Channel2
//	* Channel3
//	* EdgeTypeRising
//	* EdgeTypeFalling
//	* EdgeTypeBoth
func (device *IndustrialDigitalIn4V2Bricklet) GetEdgeCountConfiguration(channel Channel) (edgeType EdgeType, debounce uint8, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, channel);

	resultBytes, err := device.device.Get(uint8(FunctionGetEdgeCountConfiguration), buf.Bytes())
	if err != nil {
		return edgeType, debounce, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 10 {
		return edgeType, debounce, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
	}


	if header.ErrorCode != 0 {
		return edgeType, debounce, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &edgeType)
	binary.Read(resultBuf, binary.LittleEndian, &debounce)


	return edgeType, debounce, nil
}

// Each channel has a corresponding LED. You can turn the LED off, on or show a
// heartbeat. You can also set the LED to Channel Status. In this mode the
// LED is on if the channel is high and off otherwise.
// 
// By default all channel LEDs are configured as Channel Status.
//
// Associated constants:
//
//	* Channel0
//	* Channel1
//	* Channel2
//	* Channel3
//	* ChannelLEDConfigOff
//	* ChannelLEDConfigOn
//	* ChannelLEDConfigShowHeartbeat
//	* ChannelLEDConfigShowChannelStatus
func (device *IndustrialDigitalIn4V2Bricklet) SetChannelLEDConfig(channel Channel, config ChannelLEDConfig) (err error) {
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
//	* Channel0
//	* Channel1
//	* Channel2
//	* Channel3
//	* ChannelLEDConfigOff
//	* ChannelLEDConfigOn
//	* ChannelLEDConfigShowHeartbeat
//	* ChannelLEDConfigShowChannelStatus
func (device *IndustrialDigitalIn4V2Bricklet) GetChannelLEDConfig(channel Channel) (config ChannelLEDConfig, err error) {
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
func (device *IndustrialDigitalIn4V2Bricklet) GetSPITFPErrorCount() (errorCountAckChecksum uint32, errorCountMessageChecksum uint32, errorCountFrame uint32, errorCountOverflow uint32, err error) {
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
func (device *IndustrialDigitalIn4V2Bricklet) SetBootloaderMode(mode BootloaderMode) (status BootloaderStatus, err error) {
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
func (device *IndustrialDigitalIn4V2Bricklet) GetBootloaderMode() (mode BootloaderMode, err error) {
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
func (device *IndustrialDigitalIn4V2Bricklet) SetWriteFirmwarePointer(pointer uint32) (err error) {
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
func (device *IndustrialDigitalIn4V2Bricklet) WriteFirmware(data [64]uint8) (status uint8, err error) {
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
func (device *IndustrialDigitalIn4V2Bricklet) SetStatusLEDConfig(config StatusLEDConfig) (err error) {
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
func (device *IndustrialDigitalIn4V2Bricklet) GetStatusLEDConfig() (config StatusLEDConfig, err error) {
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
func (device *IndustrialDigitalIn4V2Bricklet) GetChipTemperature() (temperature int16, err error) {
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
func (device *IndustrialDigitalIn4V2Bricklet) Reset() (err error) {
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
func (device *IndustrialDigitalIn4V2Bricklet) WriteUID(uid uint32) (err error) {
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
func (device *IndustrialDigitalIn4V2Bricklet) ReadUID() (uid uint32, err error) {
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
func (device *IndustrialDigitalIn4V2Bricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
