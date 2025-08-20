/* ***********************************************************
 * This file was automatically generated on 2025-08-20.      *
 *                                                           *
 * Go Bindings Version 2.0.16                                *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/

// HAT for Raspberry Pi with 8 Bricklets ports and real-time clock.
//
//
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricks/HAT_Brick_Go.html.
package hat_brick

import (
	"bytes"
	"encoding/binary"
	"fmt"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionSetSleepMode                     Function = 1
	FunctionGetSleepMode                     Function = 2
	FunctionSetBrickletPower                 Function = 3
	FunctionGetBrickletPower                 Function = 4
	FunctionGetVoltages                      Function = 5
	FunctionSetVoltagesCallbackConfiguration Function = 6
	FunctionGetVoltagesCallbackConfiguration Function = 7
	FunctionSetRTCDriver                     Function = 9
	FunctionGetRTCDriver                     Function = 10
	FunctionGetSPITFPErrorCount              Function = 234
	FunctionSetBootloaderMode                Function = 235
	FunctionGetBootloaderMode                Function = 236
	FunctionSetWriteFirmwarePointer          Function = 237
	FunctionWriteFirmware                    Function = 238
	FunctionSetStatusLEDConfig               Function = 239
	FunctionGetStatusLEDConfig               Function = 240
	FunctionGetChipTemperature               Function = 242
	FunctionReset                            Function = 243
	FunctionWriteUID                         Function = 248
	FunctionReadUID                          Function = 249
	FunctionGetIdentity                      Function = 255
	FunctionCallbackVoltages                 Function = 8
)

type RTCDriver = uint8

const (
	RTCDriverPCF8523 RTCDriver = 0
	RTCDriverDS1338  RTCDriver = 1
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

type HATBrick struct {
	device Device
}

const DeviceIdentifier = 111
const DeviceDisplayName = "HAT Brick"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (HATBrick, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{2, 0, 2}, uid, &internalIPCon, 0, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return HATBrick{}, err
	}
	dev.ResponseExpected[FunctionSetSleepMode] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetSleepMode] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetBrickletPower] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetBrickletPower] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetVoltages] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetVoltagesCallbackConfiguration] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionGetVoltagesCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetRTCDriver] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetRTCDriver] = ResponseExpectedFlagAlwaysTrue
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
	return HATBrick{dev}, nil
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
func (device *HATBrick) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *HATBrick) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *HATBrick) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *HATBrick) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// This callback is triggered periodically according to the configuration set by
// SetVoltagesCallbackConfiguration.
//
// The parameters are the same as GetVoltages.
//
// .. versionadded:: 2.0.1$nbsp;(Firmware)
func (device *HATBrick) RegisterVoltagesCallback(fn func(uint16, uint16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 12 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var voltageUSB uint16
		var voltageDC uint16
		binary.Read(buf, binary.LittleEndian, &voltageUSB)
		binary.Read(buf, binary.LittleEndian, &voltageDC)
		fn(voltageUSB, voltageDC)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackVoltages), wrapper)
}

// Remove a registered Voltages callback.
func (device *HATBrick) DeregisterVoltagesCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackVoltages), registrationId)
}

// Sets the sleep mode.
//
// Note
//  Calling this function will cut the Raspberry Pi's power after Power Off Delay seconds.
//  You have to shut down the operating system yourself, e.g. by calling 'sudo shutdown -h now'.
//
// Parameters:
//
// * Power Off Delay: Time before the RPi/Bricklets are powered off.
// * Power Off Duration: Duration that the RPi/Bricklets stay powered off.
// * Raspberry Pi Off: RPi is powered off if set to true.
// * Bricklets Off: Bricklets are powered off if set to true.
// * Enable Sleep Indicator: If set to true, the status LED will blink in a 1s interval
//   during the whole power off duration. This will draw additional 0.3mA.
//
// Example: To turn RPi and Bricklets off in 5 seconds for 10 minutes with sleep
// indicator enabled, call (5, 60*10, true, true, true).
//
// This function can also be used to implement a watchdog. To do this you can
// write a program that calls this function once per second in a loop with
// (10, 2, true, false, false). If the RPi crashes or gets stuck
// the HAT will reset the RPi after 10 seconds.
func (device *HATBrick) SetSleepMode(powerOffDelay uint32, powerOffDuration uint32, raspberryPiOff bool, brickletsOff bool, enableSleepIndicator bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, powerOffDelay)
	binary.Write(&buf, binary.LittleEndian, powerOffDuration)
	binary.Write(&buf, binary.LittleEndian, raspberryPiOff)
	binary.Write(&buf, binary.LittleEndian, brickletsOff)
	binary.Write(&buf, binary.LittleEndian, enableSleepIndicator)

	resultBytes, err := device.device.Set(uint8(FunctionSetSleepMode), buf.Bytes())
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

// Returns the sleep mode settings as set by SetSleepMode.
func (device *HATBrick) GetSleepMode() (powerOffDelay uint32, powerOffDuration uint32, raspberryPiOff bool, brickletsOff bool, enableSleepIndicator bool, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetSleepMode), buf.Bytes())
	if err != nil {
		return powerOffDelay, powerOffDuration, raspberryPiOff, brickletsOff, enableSleepIndicator, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return powerOffDelay, powerOffDuration, raspberryPiOff, brickletsOff, enableSleepIndicator, DeviceError(header.ErrorCode)
		}

		if header.Length != 19 {
			return powerOffDelay, powerOffDuration, raspberryPiOff, brickletsOff, enableSleepIndicator, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 19)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &powerOffDelay)
		binary.Read(resultBuf, binary.LittleEndian, &powerOffDuration)
		binary.Read(resultBuf, binary.LittleEndian, &raspberryPiOff)
		binary.Read(resultBuf, binary.LittleEndian, &brickletsOff)
		binary.Read(resultBuf, binary.LittleEndian, &enableSleepIndicator)

	}

	return powerOffDelay, powerOffDuration, raspberryPiOff, brickletsOff, enableSleepIndicator, nil
}

// Set to true/false to turn the power supply of the connected Bricklets on/off.
func (device *HATBrick) SetBrickletPower(brickletPower bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, brickletPower)

	resultBytes, err := device.device.Set(uint8(FunctionSetBrickletPower), buf.Bytes())
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

// Returns the power status of the connected Bricklets as set by SetBrickletPower.
func (device *HATBrick) GetBrickletPower() (brickletPower bool, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetBrickletPower), buf.Bytes())
	if err != nil {
		return brickletPower, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return brickletPower, DeviceError(header.ErrorCode)
		}

		if header.Length != 9 {
			return brickletPower, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &brickletPower)

	}

	return brickletPower, nil
}

// Returns the USB supply voltage and the DC input supply voltage.
//
// There are three possible combinations:
//
// * Only USB connected: The USB supply voltage will be fed back to the
//   DC input connector. You will read the USB voltage and a slightly lower
//   voltage on the DC input.
// * Only DC input connected: The DC voltage will not be fed back to the
//   USB connector. You will read the DC input voltage and the USB voltage
//   will be 0.
// * USB and DC input connected: You will read both voltages. In this case
//   the USB supply will be without load, but it will work as backup if you
//   disconnect the DC input (or if the DC input voltage falls below the
//   USB voltage).
func (device *HATBrick) GetVoltages() (voltageUSB uint16, voltageDC uint16, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetVoltages), buf.Bytes())
	if err != nil {
		return voltageUSB, voltageDC, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return voltageUSB, voltageDC, DeviceError(header.ErrorCode)
		}

		if header.Length != 12 {
			return voltageUSB, voltageDC, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &voltageUSB)
		binary.Read(resultBuf, binary.LittleEndian, &voltageDC)

	}

	return voltageUSB, voltageDC, nil
}

// The period is the period with which the RegisterVoltagesCallback
// callback is triggered periodically. A value of 0 turns the callback off.
//
// If the `value has to change`-parameter is set to true, the callback is only
// triggered after the value has changed. If the value didn't change within the
// period, the callback is triggered immediately on change.
//
// If it is set to false, the callback is continuously triggered with the period,
// independent of the value.
//
// .. versionadded:: 2.0.1$nbsp;(Firmware)
func (device *HATBrick) SetVoltagesCallbackConfiguration(period uint32, valueHasToChange bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period)
	binary.Write(&buf, binary.LittleEndian, valueHasToChange)

	resultBytes, err := device.device.Set(uint8(FunctionSetVoltagesCallbackConfiguration), buf.Bytes())
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
// SetVoltagesCallbackConfiguration.
//
// .. versionadded:: 2.0.1$nbsp;(Firmware)
func (device *HATBrick) GetVoltagesCallbackConfiguration() (period uint32, valueHasToChange bool, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetVoltagesCallbackConfiguration), buf.Bytes())
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

// Configures the RTC driver that is given to the Raspberry Pi to be used.
// Currently there are two different RTCs used:
//
// * Hardware version <= 1.5: PCF8523
// * Hardware version 1.6: DS1338
//
// The correct driver will be set during factory flashing by Tinkerforge.
//
// .. versionadded:: 2.0.3$nbsp;(Firmware)
//
// Associated constants:
//
//	* RTCDriverPCF8523
//	* RTCDriverDS1338
func (device *HATBrick) SetRTCDriver(rtcDriver RTCDriver) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, rtcDriver)

	resultBytes, err := device.device.Set(uint8(FunctionSetRTCDriver), buf.Bytes())
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

// Returns the RTC driver as set by SetRTCDriver.
//
// .. versionadded:: 2.0.3$nbsp;(Firmware)
//
// Associated constants:
//
//	* RTCDriverPCF8523
//	* RTCDriverDS1338
func (device *HATBrick) GetRTCDriver() (rtcDriver RTCDriver, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetRTCDriver), buf.Bytes())
	if err != nil {
		return rtcDriver, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return rtcDriver, DeviceError(header.ErrorCode)
		}

		if header.Length != 9 {
			return rtcDriver, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &rtcDriver)

	}

	return rtcDriver, nil
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
func (device *HATBrick) GetSPITFPErrorCount() (errorCountAckChecksum uint32, errorCountMessageChecksum uint32, errorCountFrame uint32, errorCountOverflow uint32, err error) {
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
func (device *HATBrick) SetBootloaderMode(mode BootloaderMode) (status BootloaderStatus, err error) {
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
func (device *HATBrick) GetBootloaderMode() (mode BootloaderMode, err error) {
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
func (device *HATBrick) SetWriteFirmwarePointer(pointer uint32) (err error) {
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
func (device *HATBrick) WriteFirmware(data [64]uint8) (status uint8, err error) {
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
func (device *HATBrick) SetStatusLEDConfig(config StatusLEDConfig) (err error) {
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
func (device *HATBrick) GetStatusLEDConfig() (config StatusLEDConfig, err error) {
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
func (device *HATBrick) GetChipTemperature() (temperature int16, err error) {
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
func (device *HATBrick) Reset() (err error) {
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
func (device *HATBrick) WriteUID(uid uint32) (err error) {
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
func (device *HATBrick) ReadUID() (uid uint32, err error) {
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

// Returns the UID, the UID where the HAT is connected to
// (typically '0' as the HAT is the root device in the topology),
// the position, the hardware and firmware version as well as the
// device identifier.
//
// The HAT (Zero) Brick is always at position 'i'.
//
// The device identifier numbers can be found `here <device_identifier>`.
// |device_identifier_constant|
func (device *HATBrick) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
