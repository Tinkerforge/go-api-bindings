/* ***********************************************************
 * This file was automatically generated on 2020-11-02.      *
 *                                                           *
 * Go Bindings Version 2.0.9                                 *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


// 4 channel counter up to 4MHz.
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/IndustrialCounter_Bricklet_Go.html.
package industrial_counter_bricklet

import (
	"encoding/binary"
	"bytes"
	"fmt"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionGetCounter Function = 1
	FunctionGetAllCounter Function = 2
	FunctionSetCounter Function = 3
	FunctionSetAllCounter Function = 4
	FunctionGetSignalData Function = 5
	FunctionGetAllSignalData Function = 6
	FunctionSetCounterActive Function = 7
	FunctionSetAllCounterActive Function = 8
	FunctionGetCounterActive Function = 9
	FunctionGetAllCounterActive Function = 10
	FunctionSetCounterConfiguration Function = 11
	FunctionGetCounterConfiguration Function = 12
	FunctionSetAllCounterCallbackConfiguration Function = 13
	FunctionGetAllCounterCallbackConfiguration Function = 14
	FunctionSetAllSignalDataCallbackConfiguration Function = 15
	FunctionGetAllSignalDataCallbackConfiguration Function = 16
	FunctionSetChannelLEDConfig Function = 17
	FunctionGetChannelLEDConfig Function = 18
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
	FunctionCallbackAllCounter Function = 19
	FunctionCallbackAllSignalData Function = 20
)

type Channel = uint8

const (
	Channel0 Channel = 0
	Channel1 Channel = 1
	Channel2 Channel = 2
	Channel3 Channel = 3
)

type CountEdge = uint8

const (
	CountEdgeRising CountEdge = 0
	CountEdgeFalling CountEdge = 1
	CountEdgeBoth CountEdge = 2
)

type CountDirection = uint8

const (
	CountDirectionUp CountDirection = 0
	CountDirectionDown CountDirection = 1
	CountDirectionExternalUp CountDirection = 2
	CountDirectionExternalDown CountDirection = 3
)

type DutyCyclePrescaler = uint8

const (
	DutyCyclePrescaler1 DutyCyclePrescaler = 0
	DutyCyclePrescaler2 DutyCyclePrescaler = 1
	DutyCyclePrescaler4 DutyCyclePrescaler = 2
	DutyCyclePrescaler8 DutyCyclePrescaler = 3
	DutyCyclePrescaler16 DutyCyclePrescaler = 4
	DutyCyclePrescaler32 DutyCyclePrescaler = 5
	DutyCyclePrescaler64 DutyCyclePrescaler = 6
	DutyCyclePrescaler128 DutyCyclePrescaler = 7
	DutyCyclePrescaler256 DutyCyclePrescaler = 8
	DutyCyclePrescaler512 DutyCyclePrescaler = 9
	DutyCyclePrescaler1024 DutyCyclePrescaler = 10
	DutyCyclePrescaler2048 DutyCyclePrescaler = 11
	DutyCyclePrescaler4096 DutyCyclePrescaler = 12
	DutyCyclePrescaler8192 DutyCyclePrescaler = 13
	DutyCyclePrescaler16384 DutyCyclePrescaler = 14
	DutyCyclePrescaler32768 DutyCyclePrescaler = 15
)

type FrequencyIntegrationTime = uint8

const (
	FrequencyIntegrationTime128MS FrequencyIntegrationTime = 0
	FrequencyIntegrationTime256MS FrequencyIntegrationTime = 1
	FrequencyIntegrationTime512MS FrequencyIntegrationTime = 2
	FrequencyIntegrationTime1024MS FrequencyIntegrationTime = 3
	FrequencyIntegrationTime2048MS FrequencyIntegrationTime = 4
	FrequencyIntegrationTime4096MS FrequencyIntegrationTime = 5
	FrequencyIntegrationTime8192MS FrequencyIntegrationTime = 6
	FrequencyIntegrationTime16384MS FrequencyIntegrationTime = 7
	FrequencyIntegrationTime32768MS FrequencyIntegrationTime = 8
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

type IndustrialCounterBricklet struct {
	device Device
}
const DeviceIdentifier = 293
const DeviceDisplayName = "Industrial Counter Bricklet"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (IndustrialCounterBricklet, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{ 2,0,0 }, uid, &internalIPCon, 0, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return IndustrialCounterBricklet{}, err
	}
	dev.ResponseExpected[FunctionGetCounter] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetAllCounter] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetCounter] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionSetAllCounter] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetSignalData] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetAllSignalData] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetCounterActive] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionSetAllCounterActive] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetCounterActive] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetAllCounterActive] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetCounterConfiguration] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetCounterConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetAllCounterCallbackConfiguration] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetAllCounterCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetAllSignalDataCallbackConfiguration] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetAllSignalDataCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue;
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
	return IndustrialCounterBricklet{dev}, nil
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
func (device *IndustrialCounterBricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *IndustrialCounterBricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *IndustrialCounterBricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *IndustrialCounterBricklet) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// This callback is triggered periodically according to the configuration set by
// SetAllCounterCallbackConfiguration.
// 
// The parameters are the same as GetAllCounter.
func (device *IndustrialCounterBricklet) RegisterAllCounterCallback(fn func([4]int64)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 40 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var counter [4]int64
		binary.Read(buf, binary.LittleEndian, &counter)
		fn(counter)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackAllCounter), wrapper)
}

// Remove a registered All Counter callback.
func (device *IndustrialCounterBricklet) DeregisterAllCounterCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackAllCounter), registrationId)
}


// This callback is triggered periodically according to the configuration set by
// SetAllSignalDataCallbackConfiguration.
// 
// The parameters are the same as GetAllSignalData.
func (device *IndustrialCounterBricklet) RegisterAllSignalDataCallback(fn func([4]uint16, [4]uint64, [4]uint32, [4]bool)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 65 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var dutyCycle [4]uint16
		var period [4]uint64
		var frequency [4]uint32
		var value [4]bool
		binary.Read(buf, binary.LittleEndian, &dutyCycle)
		binary.Read(buf, binary.LittleEndian, &period)
		binary.Read(buf, binary.LittleEndian, &frequency)
		binary.Read(buf, binary.LittleEndian, &value)
		fn(dutyCycle, period, frequency, value)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackAllSignalData), wrapper)
}

// Remove a registered All Signal Data callback.
func (device *IndustrialCounterBricklet) DeregisterAllSignalDataCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackAllSignalData), registrationId)
}


// Returns the current counter value for the given channel.
//
// Associated constants:
//
//	* Channel0
//	* Channel1
//	* Channel2
//	* Channel3
func (device *IndustrialCounterBricklet) GetCounter(channel Channel) (counter int64, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, channel);

	resultBytes, err := device.device.Get(uint8(FunctionGetCounter), buf.Bytes())
	if err != nil {
		return counter, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 16 {
			return counter, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 16)
		}

		if header.ErrorCode != 0 {
			return counter, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &counter)

	}

	return counter, nil
}

// Returns the current counter values for all four channels.
func (device *IndustrialCounterBricklet) GetAllCounter() (counter [4]int64, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetAllCounter), buf.Bytes())
	if err != nil {
		return counter, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 40 {
			return counter, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 40)
		}

		if header.ErrorCode != 0 {
			return counter, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &counter)

	}

	return counter, nil
}

// Sets the counter value for the given channel.
// 
// The default value for the counters on startup is 0.
//
// Associated constants:
//
//	* Channel0
//	* Channel1
//	* Channel2
//	* Channel3
func (device *IndustrialCounterBricklet) SetCounter(channel Channel, counter int64) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, channel);
	binary.Write(&buf, binary.LittleEndian, counter);

	resultBytes, err := device.device.Set(uint8(FunctionSetCounter), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])
		
	}

	return nil
}

// Sets the counter values for all four channels.
// 
// The default value for the counters on startup is 0.
func (device *IndustrialCounterBricklet) SetAllCounter(counter [4]int64) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, counter);

	resultBytes, err := device.device.Set(uint8(FunctionSetAllCounter), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])
		
	}

	return nil
}

// Returns the signal data (duty cycle, period, frequency and value) for the
// given channel.
//
// Associated constants:
//
//	* Channel0
//	* Channel1
//	* Channel2
//	* Channel3
func (device *IndustrialCounterBricklet) GetSignalData(channel Channel) (dutyCycle uint16, period uint64, frequency uint32, value bool, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, channel);

	resultBytes, err := device.device.Get(uint8(FunctionGetSignalData), buf.Bytes())
	if err != nil {
		return dutyCycle, period, frequency, value, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 23 {
			return dutyCycle, period, frequency, value, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 23)
		}

		if header.ErrorCode != 0 {
			return dutyCycle, period, frequency, value, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &dutyCycle)
		binary.Read(resultBuf, binary.LittleEndian, &period)
		binary.Read(resultBuf, binary.LittleEndian, &frequency)
		binary.Read(resultBuf, binary.LittleEndian, &value)

	}

	return dutyCycle, period, frequency, value, nil
}

// Returns the signal data (duty cycle, period, frequency and value) for all four
// channels.
func (device *IndustrialCounterBricklet) GetAllSignalData() (dutyCycle [4]uint16, period [4]uint64, frequency [4]uint32, value [4]bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetAllSignalData), buf.Bytes())
	if err != nil {
		return dutyCycle, period, frequency, value, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 65 {
			return dutyCycle, period, frequency, value, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 65)
		}

		if header.ErrorCode != 0 {
			return dutyCycle, period, frequency, value, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &dutyCycle)
		binary.Read(resultBuf, binary.LittleEndian, &period)
		binary.Read(resultBuf, binary.LittleEndian, &frequency)
		binary.Read(resultBuf, binary.LittleEndian, &value)

	}

	return dutyCycle, period, frequency, value, nil
}

// Activates/deactivates the counter of the given channel.
// 
// true = activate, false = deactivate.
// 
// By default all channels are activated.
//
// Associated constants:
//
//	* Channel0
//	* Channel1
//	* Channel2
//	* Channel3
func (device *IndustrialCounterBricklet) SetCounterActive(channel Channel, active bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, channel);
	binary.Write(&buf, binary.LittleEndian, active);

	resultBytes, err := device.device.Set(uint8(FunctionSetCounterActive), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])
		
	}

	return nil
}

// Activates/deactivates the counter of all four channels.
// 
// true = activate, false = deactivate.
// 
// By default all channels are activated.
func (device *IndustrialCounterBricklet) SetAllCounterActive(active [4]bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, active);

	resultBytes, err := device.device.Set(uint8(FunctionSetAllCounterActive), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])
		
	}

	return nil
}

// Returns the activation state of the given channel.
// 
// true = activated, false = deactivated.
//
// Associated constants:
//
//	* Channel0
//	* Channel1
//	* Channel2
//	* Channel3
func (device *IndustrialCounterBricklet) GetCounterActive(channel Channel) (active bool, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, channel);

	resultBytes, err := device.device.Get(uint8(FunctionGetCounterActive), buf.Bytes())
	if err != nil {
		return active, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return active, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return active, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &active)

	}

	return active, nil
}

// Returns the activation state of all four channels.
// 
// true = activated, false = deactivated.
func (device *IndustrialCounterBricklet) GetAllCounterActive() (active [4]bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetAllCounterActive), buf.Bytes())
	if err != nil {
		return active, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return active, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return active, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &active)

	}

	return active, nil
}

// Sets the counter configuration for the given channel.
// 
// * Count Edge: Counter can count on rising, falling or both edges.
// * Count Direction: Counter can count up or down. You can also use
//   another channel as direction input, see
//   https://www.tinkerforge.com/en/doc/Hardware/Bricklets/Industrial_Counter.html#external-count-direction
//   for details.
// * Duty Cycle Prescaler: Sets a divider for the internal clock. See
//   https://www.tinkerforge.com/en/doc/Hardware/Bricklets/Industrial_Counter.html#duty-cycle-prescaler-and-frequency-integration-time
//   for details.
// * Frequency Integration Time: Sets the integration time for the
//   frequency measurement. See
//   https://www.tinkerforge.com/en/doc/Hardware/Bricklets/Industrial_Counter.html#duty-cycle-prescaler-and-frequency-integration-time
//   for details.
//
// Associated constants:
//
//	* Channel0
//	* Channel1
//	* Channel2
//	* Channel3
//	* CountEdgeRising
//	* CountEdgeFalling
//	* CountEdgeBoth
//	* CountDirectionUp
//	* CountDirectionDown
//	* CountDirectionExternalUp
//	* CountDirectionExternalDown
//	* DutyCyclePrescaler1
//	* DutyCyclePrescaler2
//	* DutyCyclePrescaler4
//	* DutyCyclePrescaler8
//	* DutyCyclePrescaler16
//	* DutyCyclePrescaler32
//	* DutyCyclePrescaler64
//	* DutyCyclePrescaler128
//	* DutyCyclePrescaler256
//	* DutyCyclePrescaler512
//	* DutyCyclePrescaler1024
//	* DutyCyclePrescaler2048
//	* DutyCyclePrescaler4096
//	* DutyCyclePrescaler8192
//	* DutyCyclePrescaler16384
//	* DutyCyclePrescaler32768
//	* FrequencyIntegrationTime128MS
//	* FrequencyIntegrationTime256MS
//	* FrequencyIntegrationTime512MS
//	* FrequencyIntegrationTime1024MS
//	* FrequencyIntegrationTime2048MS
//	* FrequencyIntegrationTime4096MS
//	* FrequencyIntegrationTime8192MS
//	* FrequencyIntegrationTime16384MS
//	* FrequencyIntegrationTime32768MS
func (device *IndustrialCounterBricklet) SetCounterConfiguration(channel Channel, countEdge CountEdge, countDirection CountDirection, dutyCyclePrescaler DutyCyclePrescaler, frequencyIntegrationTime FrequencyIntegrationTime) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, channel);
	binary.Write(&buf, binary.LittleEndian, countEdge);
	binary.Write(&buf, binary.LittleEndian, countDirection);
	binary.Write(&buf, binary.LittleEndian, dutyCyclePrescaler);
	binary.Write(&buf, binary.LittleEndian, frequencyIntegrationTime);

	resultBytes, err := device.device.Set(uint8(FunctionSetCounterConfiguration), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])
		
	}

	return nil
}

// Returns the counter configuration as set by SetCounterConfiguration.
//
// Associated constants:
//
//	* Channel0
//	* Channel1
//	* Channel2
//	* Channel3
//	* CountEdgeRising
//	* CountEdgeFalling
//	* CountEdgeBoth
//	* CountDirectionUp
//	* CountDirectionDown
//	* CountDirectionExternalUp
//	* CountDirectionExternalDown
//	* DutyCyclePrescaler1
//	* DutyCyclePrescaler2
//	* DutyCyclePrescaler4
//	* DutyCyclePrescaler8
//	* DutyCyclePrescaler16
//	* DutyCyclePrescaler32
//	* DutyCyclePrescaler64
//	* DutyCyclePrescaler128
//	* DutyCyclePrescaler256
//	* DutyCyclePrescaler512
//	* DutyCyclePrescaler1024
//	* DutyCyclePrescaler2048
//	* DutyCyclePrescaler4096
//	* DutyCyclePrescaler8192
//	* DutyCyclePrescaler16384
//	* DutyCyclePrescaler32768
//	* FrequencyIntegrationTime128MS
//	* FrequencyIntegrationTime256MS
//	* FrequencyIntegrationTime512MS
//	* FrequencyIntegrationTime1024MS
//	* FrequencyIntegrationTime2048MS
//	* FrequencyIntegrationTime4096MS
//	* FrequencyIntegrationTime8192MS
//	* FrequencyIntegrationTime16384MS
//	* FrequencyIntegrationTime32768MS
func (device *IndustrialCounterBricklet) GetCounterConfiguration(channel Channel) (countEdge CountEdge, countDirection CountDirection, dutyCyclePrescaler DutyCyclePrescaler, frequencyIntegrationTime FrequencyIntegrationTime, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, channel);

	resultBytes, err := device.device.Get(uint8(FunctionGetCounterConfiguration), buf.Bytes())
	if err != nil {
		return countEdge, countDirection, dutyCyclePrescaler, frequencyIntegrationTime, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 12 {
			return countEdge, countDirection, dutyCyclePrescaler, frequencyIntegrationTime, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		if header.ErrorCode != 0 {
			return countEdge, countDirection, dutyCyclePrescaler, frequencyIntegrationTime, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &countEdge)
		binary.Read(resultBuf, binary.LittleEndian, &countDirection)
		binary.Read(resultBuf, binary.LittleEndian, &dutyCyclePrescaler)
		binary.Read(resultBuf, binary.LittleEndian, &frequencyIntegrationTime)

	}

	return countEdge, countDirection, dutyCyclePrescaler, frequencyIntegrationTime, nil
}

// The period is the period with which the RegisterAllCounterCallback
// callback is triggered periodically. A value of 0 turns the callback off.
// 
// If the `value has to change`-parameter is set to true, the callback is only
// triggered after the value has changed. If the value didn't change within the
// period, the callback is triggered immediately on change.
// 
// If it is set to false, the callback is continuously triggered with the period,
// independent of the value.
func (device *IndustrialCounterBricklet) SetAllCounterCallbackConfiguration(period uint32, valueHasToChange bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);
	binary.Write(&buf, binary.LittleEndian, valueHasToChange);

	resultBytes, err := device.device.Set(uint8(FunctionSetAllCounterCallbackConfiguration), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])
		
	}

	return nil
}

// Returns the callback configuration as set by
// SetAllCounterCallbackConfiguration.
func (device *IndustrialCounterBricklet) GetAllCounterCallbackConfiguration() (period uint32, valueHasToChange bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetAllCounterCallbackConfiguration), buf.Bytes())
	if err != nil {
		return period, valueHasToChange, err
	}
	if len(resultBytes) > 0 {
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

	}

	return period, valueHasToChange, nil
}

// The period is the period with which the RegisterAllSignalDataCallback
// callback is triggered periodically. A value of 0 turns the callback off.
// 
// If the `value has to change`-parameter is set to true, the callback is only
// triggered after the value has changed. If the value didn't change within the
// period, the callback is triggered immediately on change.
// 
// If it is set to false, the callback is continuously triggered with the period,
// independent of the value.
func (device *IndustrialCounterBricklet) SetAllSignalDataCallbackConfiguration(period uint32, valueHasToChange bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);
	binary.Write(&buf, binary.LittleEndian, valueHasToChange);

	resultBytes, err := device.device.Set(uint8(FunctionSetAllSignalDataCallbackConfiguration), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])
		
	}

	return nil
}

// Returns the callback configuration as set by
// SetAllSignalDataCallbackConfiguration.
func (device *IndustrialCounterBricklet) GetAllSignalDataCallbackConfiguration() (period uint32, valueHasToChange bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetAllSignalDataCallbackConfiguration), buf.Bytes())
	if err != nil {
		return period, valueHasToChange, err
	}
	if len(resultBytes) > 0 {
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

	}

	return period, valueHasToChange, nil
}

// Each channel has a corresponding LED. You can turn the LED off, on or show a
// heartbeat. You can also set the LED to Channel Status. In this mode the
// LED is on if the channel is high and off otherwise.
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
func (device *IndustrialCounterBricklet) SetChannelLEDConfig(channel Channel, config ChannelLEDConfig) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, channel);
	binary.Write(&buf, binary.LittleEndian, config);

	resultBytes, err := device.device.Set(uint8(FunctionSetChannelLEDConfig), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])
		
	}

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
func (device *IndustrialCounterBricklet) GetChannelLEDConfig(channel Channel) (config ChannelLEDConfig, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, channel);

	resultBytes, err := device.device.Get(uint8(FunctionGetChannelLEDConfig), buf.Bytes())
	if err != nil {
		return config, err
	}
	if len(resultBytes) > 0 {
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

	}

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
func (device *IndustrialCounterBricklet) GetSPITFPErrorCount() (errorCountAckChecksum uint32, errorCountMessageChecksum uint32, errorCountFrame uint32, errorCountOverflow uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetSPITFPErrorCount), buf.Bytes())
	if err != nil {
		return errorCountAckChecksum, errorCountMessageChecksum, errorCountFrame, errorCountOverflow, err
	}
	if len(resultBytes) > 0 {
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
func (device *IndustrialCounterBricklet) SetBootloaderMode(mode BootloaderMode) (status BootloaderStatus, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, mode);

	resultBytes, err := device.device.Get(uint8(FunctionSetBootloaderMode), buf.Bytes())
	if err != nil {
		return status, err
	}
	if len(resultBytes) > 0 {
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
func (device *IndustrialCounterBricklet) GetBootloaderMode() (mode BootloaderMode, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetBootloaderMode), buf.Bytes())
	if err != nil {
		return mode, err
	}
	if len(resultBytes) > 0 {
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

	}

	return mode, nil
}

// Sets the firmware pointer for WriteFirmware. The pointer has
// to be increased by chunks of size 64. The data is written to flash
// every 4 chunks (which equals to one page of size 256).
// 
// This function is used by Brick Viewer during flashing. It should not be
// necessary to call it in a normal user program.
func (device *IndustrialCounterBricklet) SetWriteFirmwarePointer(pointer uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, pointer);

	resultBytes, err := device.device.Set(uint8(FunctionSetWriteFirmwarePointer), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

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
func (device *IndustrialCounterBricklet) WriteFirmware(data [64]uint8) (status uint8, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, data);

	resultBytes, err := device.device.Get(uint8(FunctionWriteFirmware), buf.Bytes())
	if err != nil {
		return status, err
	}
	if len(resultBytes) > 0 {
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
func (device *IndustrialCounterBricklet) SetStatusLEDConfig(config StatusLEDConfig) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, config);

	resultBytes, err := device.device.Set(uint8(FunctionSetStatusLEDConfig), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

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
func (device *IndustrialCounterBricklet) GetStatusLEDConfig() (config StatusLEDConfig, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetStatusLEDConfig), buf.Bytes())
	if err != nil {
		return config, err
	}
	if len(resultBytes) > 0 {
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

	}

	return config, nil
}

// Returns the temperature as measured inside the microcontroller. The
// value returned is not the ambient temperature!
// 
// The temperature is only proportional to the real temperature and it has bad
// accuracy. Practically it is only useful as an indicator for
// temperature changes.
func (device *IndustrialCounterBricklet) GetChipTemperature() (temperature int16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetChipTemperature), buf.Bytes())
	if err != nil {
		return temperature, err
	}
	if len(resultBytes) > 0 {
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

	}

	return temperature, nil
}

// Calling this function will reset the Bricklet. All configurations
// will be lost.
// 
// After a reset you have to create new device objects,
// calling functions on the existing ones will result in
// undefined behavior!
func (device *IndustrialCounterBricklet) Reset() (err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Set(uint8(FunctionReset), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

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
func (device *IndustrialCounterBricklet) WriteUID(uid uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, uid);

	resultBytes, err := device.device.Set(uint8(FunctionWriteUID), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])
		
	}

	return nil
}

// Returns the current UID as an integer. Encode as
// Base58 to get the usual string version.
func (device *IndustrialCounterBricklet) ReadUID() (uid uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionReadUID), buf.Bytes())
	if err != nil {
		return uid, err
	}
	if len(resultBytes) > 0 {
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

	}

	return uid, nil
}

// Returns the UID, the UID where the Bricklet is connected to,
// the position, the hardware and firmware version as well as the
// device identifier.
// 
// The position can be 'a', 'b', 'c', 'd', 'e', 'f', 'g' or 'h' (Bricklet Port).
// A Bricklet connected to an `Isolator Bricklet <isolator_bricklet>` is always at
// position 'z'.
// 
// The device identifier numbers can be found `here <device_identifier>`.
// |device_identifier_constant|
func (device *IndustrialCounterBricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetIdentity), buf.Bytes())
	if err != nil {
		return uid, connectedUid, position, hardwareVersion, firmwareVersion, deviceIdentifier, err
	}
	if len(resultBytes) > 0 {
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

	}

	return uid, connectedUid, position, hardwareVersion, firmwareVersion, deviceIdentifier, nil
}
