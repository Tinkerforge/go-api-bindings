/* ***********************************************************
 * This file was automatically generated on 2025-10-07.      *
 *                                                           *
 * Go Bindings Version 2.0.17                                *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/

// Galvanically isolates any Bricklet from any Brick.
//
//
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/Isolator_Bricklet_Go.html.
package isolator_bricklet

import (
	"bytes"
	"encoding/binary"
	"fmt"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionGetStatistics                      Function = 1
	FunctionSetSPITFPBaudrateConfig            Function = 2
	FunctionGetSPITFPBaudrateConfig            Function = 3
	FunctionSetSPITFPBaudrate                  Function = 4
	FunctionGetSPITFPBaudrate                  Function = 5
	FunctionGetIsolatorSPITFPErrorCount        Function = 6
	FunctionSetStatisticsCallbackConfiguration Function = 7
	FunctionGetStatisticsCallbackConfiguration Function = 8
	FunctionGetSPITFPErrorCount                Function = 234
	FunctionSetBootloaderMode                  Function = 235
	FunctionGetBootloaderMode                  Function = 236
	FunctionSetWriteFirmwarePointer            Function = 237
	FunctionWriteFirmware                      Function = 238
	FunctionSetStatusLEDConfig                 Function = 239
	FunctionGetStatusLEDConfig                 Function = 240
	FunctionGetChipTemperature                 Function = 242
	FunctionReset                              Function = 243
	FunctionWriteUID                           Function = 248
	FunctionReadUID                            Function = 249
	FunctionGetIdentity                        Function = 255
	FunctionCallbackStatistics                 Function = 9
)

type BootloaderMode = uint8

const (
	BootloaderModeBootloader                    BootloaderMode = 0
	BootloaderModeFirmware                      BootloaderMode = 1
	BootloaderModeBootloaderWaitForReboot       BootloaderMode = 2
	BootloaderModeFirmwareWaitForReboot         BootloaderMode = 3
	BootloaderModeFirmwareWaitForEraseAndReboot BootloaderMode = 4
)

type BootloaderStatus = uint8

const (
	BootloaderStatusOK                        BootloaderStatus = 0
	BootloaderStatusInvalidMode               BootloaderStatus = 1
	BootloaderStatusNoChange                  BootloaderStatus = 2
	BootloaderStatusEntryFunctionNotPresent   BootloaderStatus = 3
	BootloaderStatusDeviceIdentifierIncorrect BootloaderStatus = 4
	BootloaderStatusCRCMismatch               BootloaderStatus = 5
)

type StatusLEDConfig = uint8

const (
	StatusLEDConfigOff           StatusLEDConfig = 0
	StatusLEDConfigOn            StatusLEDConfig = 1
	StatusLEDConfigShowHeartbeat StatusLEDConfig = 2
	StatusLEDConfigShowStatus    StatusLEDConfig = 3
)

type IsolatorBricklet struct {
	device Device
}

const DeviceIdentifier = 2122
const DeviceDisplayName = "Isolator Bricklet"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (IsolatorBricklet, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{2, 0, 1}, uid, &internalIPCon, 0, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return IsolatorBricklet{}, err
	}
	dev.ResponseExpected[FunctionGetStatistics] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetSPITFPBaudrateConfig] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetSPITFPBaudrateConfig] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetSPITFPBaudrate] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetSPITFPBaudrate] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetIsolatorSPITFPErrorCount] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetStatisticsCallbackConfiguration] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionGetStatisticsCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetSPITFPErrorCount] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetBootloaderMode] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetBootloaderMode] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetWriteFirmwarePointer] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionWriteFirmware] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetStatusLEDConfig] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetStatusLEDConfig] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetChipTemperature] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionReset] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionWriteUID] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionReadUID] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetIdentity] = ResponseExpectedFlagAlwaysTrue
	return IsolatorBricklet{dev}, nil
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
func (device *IsolatorBricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *IsolatorBricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *IsolatorBricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *IsolatorBricklet) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// This callback is triggered periodically according to the configuration set by
// SetStatisticsCallbackConfiguration.
//
// The parameters are the same as GetStatistics.
//
// .. versionadded:: 2.0.2$nbsp;(Plugin)
func (device *IsolatorBricklet) RegisterStatisticsCallback(fn func(uint32, uint32, uint16, string)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 26 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var messagesFromBrick uint32
		var messagesFromBricklet uint32
		var connectedBrickletDeviceIdentifier uint16
		var connectedBrickletUID string
		binary.Read(buf, binary.LittleEndian, &messagesFromBrick)
		binary.Read(buf, binary.LittleEndian, &messagesFromBricklet)
		binary.Read(buf, binary.LittleEndian, &connectedBrickletDeviceIdentifier)
		connectedBrickletUID = ByteSliceToString(buf.Next(8))
		fn(messagesFromBrick, messagesFromBricklet, connectedBrickletDeviceIdentifier, connectedBrickletUID)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackStatistics), wrapper)
}

// Remove a registered Statistics callback.
func (device *IsolatorBricklet) DeregisterStatisticsCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackStatistics), registrationId)
}

// Returns statistics for the Isolator Bricklet.
func (device *IsolatorBricklet) GetStatistics() (messagesFromBrick uint32, messagesFromBricklet uint32, connectedBrickletDeviceIdentifier uint16, connectedBrickletUID string, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetStatistics), buf.Bytes())
	if err != nil {
		return messagesFromBrick, messagesFromBricklet, connectedBrickletDeviceIdentifier, connectedBrickletUID, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return messagesFromBrick, messagesFromBricklet, connectedBrickletDeviceIdentifier, connectedBrickletUID, DeviceError(header.ErrorCode)
		}

		if header.Length != 26 {
			return messagesFromBrick, messagesFromBricklet, connectedBrickletDeviceIdentifier, connectedBrickletUID, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 26)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &messagesFromBrick)
		binary.Read(resultBuf, binary.LittleEndian, &messagesFromBricklet)
		binary.Read(resultBuf, binary.LittleEndian, &connectedBrickletDeviceIdentifier)
		connectedBrickletUID = ByteSliceToString(resultBuf.Next(8))

	}

	return messagesFromBrick, messagesFromBricklet, connectedBrickletDeviceIdentifier, connectedBrickletUID, nil
}

// The SPITF protocol can be used with a dynamic baudrate. If the dynamic baudrate is
// enabled, the Isolator Bricklet will try to adapt the baudrate for the communication
// between Bricks and Bricklets according to the amount of data that is transferred.
//
// The baudrate for communication config between
// Brick and Isolator Bricklet can be set through the API of the Brick.
//
// The baudrate will be increased exponentially if lots of data is sent/received and
// decreased linearly if little data is sent/received.
//
// This lowers the baudrate in applications where little data is transferred (e.g.
// a weather station) and increases the robustness. If there is lots of data to transfer
// (e.g. Thermal Imaging Bricklet) it automatically increases the baudrate as needed.
//
// In cases where some data has to transferred as fast as possible every few seconds
// (e.g. RS485 Bricklet with a high baudrate but small payload) you may want to turn
// the dynamic baudrate off to get the highest possible performance.
//
// The maximum value of the baudrate can be set per port with the function
// SetSPITFPBaudrate. If the dynamic baudrate is disabled, the baudrate
// as set by SetSPITFPBaudrate will be used statically.
func (device *IsolatorBricklet) SetSPITFPBaudrateConfig(enableDynamicBaudrate bool, minimumDynamicBaudrate uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, enableDynamicBaudrate)
	binary.Write(&buf, binary.LittleEndian, minimumDynamicBaudrate)

	resultBytes, err := device.device.Set(uint8(FunctionSetSPITFPBaudrateConfig), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		bytes.NewBuffer(resultBytes[8:])

	}

	return nil
}

// Returns the baudrate config, see SetSPITFPBaudrateConfig.
func (device *IsolatorBricklet) GetSPITFPBaudrateConfig() (enableDynamicBaudrate bool, minimumDynamicBaudrate uint32, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetSPITFPBaudrateConfig), buf.Bytes())
	if err != nil {
		return enableDynamicBaudrate, minimumDynamicBaudrate, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return enableDynamicBaudrate, minimumDynamicBaudrate, DeviceError(header.ErrorCode)
		}

		if header.Length != 13 {
			return enableDynamicBaudrate, minimumDynamicBaudrate, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 13)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &enableDynamicBaudrate)
		binary.Read(resultBuf, binary.LittleEndian, &minimumDynamicBaudrate)

	}

	return enableDynamicBaudrate, minimumDynamicBaudrate, nil
}

// Sets the baudrate for a the communication between Isolator Bricklet
// and the connected Bricklet. The baudrate for communication between
// Brick and Isolator Bricklet can be set through the API of the Brick.
//
// If you want to increase the throughput of Bricklets you can increase
// the baudrate. If you get a high error count because of high
// interference (see GetSPITFPErrorCount) you can decrease the
// baudrate.
//
// If the dynamic baudrate feature is enabled, the baudrate set by this
// function corresponds to the maximum baudrate (see SetSPITFPBaudrateConfig).
//
// Regulatory testing is done with the default baudrate. If CE compatibility
// or similar is necessary in your applications we recommend to not change
// the baudrate.
func (device *IsolatorBricklet) SetSPITFPBaudrate(baudrate uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, baudrate)

	resultBytes, err := device.device.Set(uint8(FunctionSetSPITFPBaudrate), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		bytes.NewBuffer(resultBytes[8:])

	}

	return nil
}

// Returns the baudrate, see SetSPITFPBaudrate.
func (device *IsolatorBricklet) GetSPITFPBaudrate() (baudrate uint32, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetSPITFPBaudrate), buf.Bytes())
	if err != nil {
		return baudrate, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return baudrate, DeviceError(header.ErrorCode)
		}

		if header.Length != 12 {
			return baudrate, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &baudrate)

	}

	return baudrate, nil
}

// Returns the error count for the communication between Isolator Bricklet and
// the connected Bricklet. Call GetSPITFPErrorCount to get the
// error count between Isolator Bricklet and Brick.
//
// The errors are divided into
//
// * ACK checksum errors,
// * message checksum errors,
// * framing errors and
// * overflow errors.
func (device *IsolatorBricklet) GetIsolatorSPITFPErrorCount() (errorCountACKChecksum uint32, errorCountMessageChecksum uint32, errorCountFrame uint32, errorCountOverflow uint32, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetIsolatorSPITFPErrorCount), buf.Bytes())
	if err != nil {
		return errorCountACKChecksum, errorCountMessageChecksum, errorCountFrame, errorCountOverflow, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return errorCountACKChecksum, errorCountMessageChecksum, errorCountFrame, errorCountOverflow, DeviceError(header.ErrorCode)
		}

		if header.Length != 24 {
			return errorCountACKChecksum, errorCountMessageChecksum, errorCountFrame, errorCountOverflow, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 24)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &errorCountACKChecksum)
		binary.Read(resultBuf, binary.LittleEndian, &errorCountMessageChecksum)
		binary.Read(resultBuf, binary.LittleEndian, &errorCountFrame)
		binary.Read(resultBuf, binary.LittleEndian, &errorCountOverflow)

	}

	return errorCountACKChecksum, errorCountMessageChecksum, errorCountFrame, errorCountOverflow, nil
}

// The period is the period with which the RegisterStatisticsCallback
// callback is triggered periodically. A value of 0 turns the callback off.
//
// If the `value has to change`-parameter is set to true, the callback is only
// triggered after the value has changed. If the value didn't change within the
// period, the callback is triggered immediately on change.
//
// If it is set to false, the callback is continuously triggered with the period,
// independent of the value.
//
// .. versionadded:: 2.0.2$nbsp;(Plugin)
func (device *IsolatorBricklet) SetStatisticsCallbackConfiguration(period uint32, valueHasToChange bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period)
	binary.Write(&buf, binary.LittleEndian, valueHasToChange)

	resultBytes, err := device.device.Set(uint8(FunctionSetStatisticsCallbackConfiguration), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		bytes.NewBuffer(resultBytes[8:])

	}

	return nil
}

// Returns the callback configuration as set by
// SetStatisticsCallbackConfiguration.
//
// .. versionadded:: 2.0.2$nbsp;(Plugin)
func (device *IsolatorBricklet) GetStatisticsCallbackConfiguration() (period uint32, valueHasToChange bool, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetStatisticsCallbackConfiguration), buf.Bytes())
	if err != nil {
		return period, valueHasToChange, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return period, valueHasToChange, DeviceError(header.ErrorCode)
		}

		if header.Length != 13 {
			return period, valueHasToChange, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 13)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &period)
		binary.Read(resultBuf, binary.LittleEndian, &valueHasToChange)

	}

	return period, valueHasToChange, nil
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
func (device *IsolatorBricklet) GetSPITFPErrorCount() (errorCountAckChecksum uint32, errorCountMessageChecksum uint32, errorCountFrame uint32, errorCountOverflow uint32, err error) {
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

		if header.Length != 24 {
			return errorCountAckChecksum, errorCountMessageChecksum, errorCountFrame, errorCountOverflow, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 24)
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
func (device *IsolatorBricklet) SetBootloaderMode(mode BootloaderMode) (status BootloaderStatus, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, mode)

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

		if header.Length != 9 {
			return status, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
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
func (device *IsolatorBricklet) GetBootloaderMode() (mode BootloaderMode, err error) {
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

		if header.Length != 9 {
			return mode, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
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
func (device *IsolatorBricklet) SetWriteFirmwarePointer(pointer uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, pointer)

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

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
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
func (device *IsolatorBricklet) WriteFirmware(data [64]uint8) (status uint8, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, data)

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

		if header.Length != 9 {
			return status, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
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
func (device *IsolatorBricklet) SetStatusLEDConfig(config StatusLEDConfig) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, config)

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

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
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
func (device *IsolatorBricklet) GetStatusLEDConfig() (config StatusLEDConfig, err error) {
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

		if header.Length != 9 {
			return config, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
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
func (device *IsolatorBricklet) GetChipTemperature() (temperature int16, err error) {
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

		if header.Length != 10 {
			return temperature, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
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
func (device *IsolatorBricklet) Reset() (err error) {
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

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
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
func (device *IsolatorBricklet) WriteUID(uid uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, uid)

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

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		bytes.NewBuffer(resultBytes[8:])

	}

	return nil
}

// Returns the current UID as an integer. Encode as
// Base58 to get the usual string version.
func (device *IsolatorBricklet) ReadUID() (uid uint32, err error) {
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

		if header.Length != 12 {
			return uid, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
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
func (device *IsolatorBricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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

		if header.Length != 33 {
			return uid, connectedUid, position, hardwareVersion, firmwareVersion, deviceIdentifier, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 33)
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
