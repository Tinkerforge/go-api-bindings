/* ***********************************************************
 * This file was automatically generated on 2024-02-27.      *
 *                                                           *
 * Go Bindings Version 2.0.15                                *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/

// 3.3cm (1.3") OLED display with 128x64 pixels.
//
//
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/OLED128x64V2_Bricklet_Go.html.
package oled_128x64_v2_bricklet

import (
	"bytes"
	"encoding/binary"
	"fmt"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionWritePixelsLowLevel     Function = 1
	FunctionReadPixelsLowLevel      Function = 2
	FunctionClearDisplay            Function = 3
	FunctionSetDisplayConfiguration Function = 4
	FunctionGetDisplayConfiguration Function = 5
	FunctionWriteLine               Function = 6
	FunctionDrawBufferedFrame       Function = 7
	FunctionGetSPITFPErrorCount     Function = 234
	FunctionSetBootloaderMode       Function = 235
	FunctionGetBootloaderMode       Function = 236
	FunctionSetWriteFirmwarePointer Function = 237
	FunctionWriteFirmware           Function = 238
	FunctionSetStatusLEDConfig      Function = 239
	FunctionGetStatusLEDConfig      Function = 240
	FunctionGetChipTemperature      Function = 242
	FunctionReset                   Function = 243
	FunctionWriteUID                Function = 248
	FunctionReadUID                 Function = 249
	FunctionGetIdentity             Function = 255
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

type OLED128x64V2Bricklet struct {
	device Device
}

const DeviceIdentifier = 2112
const DeviceDisplayName = "OLED 128x64 Bricklet 2.0"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (OLED128x64V2Bricklet, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{2, 0, 0}, uid, &internalIPCon, 2, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return OLED128x64V2Bricklet{}, err
	}
	dev.ResponseExpected[FunctionWritePixelsLowLevel] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionReadPixelsLowLevel] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionClearDisplay] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionSetDisplayConfiguration] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetDisplayConfiguration] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionWriteLine] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionDrawBufferedFrame] = ResponseExpectedFlagFalse
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
	return OLED128x64V2Bricklet{dev}, nil
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
func (device *OLED128x64V2Bricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *OLED128x64V2Bricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *OLED128x64V2Bricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *OLED128x64V2Bricklet) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// Writes pixels to the specified window.
//
// The pixels are written into the window line by line top to bottom
// and each line is written from left to right.
//
// If automatic draw is enabled (default) the pixels are directly written to
// the screen. Only pixels that have actually changed are updated on the screen,
// the rest stays the same.
//
// If automatic draw is disabled the pixels are written to an internal buffer and
// the buffer is transferred to the display only after DrawBufferedFrame
// is called. This can be used to avoid flicker when drawing a complex frame in
// multiple steps.
//
// Automatic draw can be configured with the SetDisplayConfiguration
// function.
func (device *OLED128x64V2Bricklet) WritePixelsLowLevel(xStart uint8, yStart uint8, xEnd uint8, yEnd uint8, pixelsLength uint16, pixelsChunkOffset uint16, pixelsChunkData [448]bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, xStart)
	binary.Write(&buf, binary.LittleEndian, yStart)
	binary.Write(&buf, binary.LittleEndian, xEnd)
	binary.Write(&buf, binary.LittleEndian, yEnd)
	binary.Write(&buf, binary.LittleEndian, pixelsLength)
	binary.Write(&buf, binary.LittleEndian, pixelsChunkOffset)
	buf.Write(BoolSliceToByteSlice(pixelsChunkData[:]))

	resultBytes, err := device.device.Set(uint8(FunctionWritePixelsLowLevel), buf.Bytes())
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

// Writes pixels to the specified window.
//
// The pixels are written into the window line by line top to bottom
// and each line is written from left to right.
//
// If automatic draw is enabled (default) the pixels are directly written to
// the screen. Only pixels that have actually changed are updated on the screen,
// the rest stays the same.
//
// If automatic draw is disabled the pixels are written to an internal buffer and
// the buffer is transferred to the display only after DrawBufferedFrame
// is called. This can be used to avoid flicker when drawing a complex frame in
// multiple steps.
//
// Automatic draw can be configured with the SetDisplayConfiguration
// function.
func (device *OLED128x64V2Bricklet) WritePixels(xStart uint8, yStart uint8, xEnd uint8, yEnd uint8, pixels []bool) (err error) {
	_, err = device.device.SetHighLevel(func(pixelsLength uint64, pixelsChunkOffset uint64, pixelsChunkData []byte) (LowLevelWriteResult, error) {
		arr := [448]bool{}
		copy(arr[:], ByteSliceToBoolSlice(pixelsChunkData))

		err := device.WritePixelsLowLevel(xStart, yStart, xEnd, yEnd, uint16(pixelsLength), uint16(pixelsChunkOffset), arr)

		var lowLevelResults bytes.Buffer

		return LowLevelWriteResult{
			uint64(448),
			lowLevelResults.Bytes()}, err
	}, 0, 1, 448, BoolSliceToByteSlice(pixels))

	if err != nil {
		return
	}

	return
}

// Reads pixels from the specified window.
//
// The pixels are read from the window line by line top to bottom
// and each line is read from left to right.
//
// If automatic draw is enabled (default) the pixels that are read are always the
// same that are shown on the display.
//
// If automatic draw is disabled the pixels are read from the internal buffer
// (see DrawBufferedFrame).
//
// Automatic draw can be configured with the SetDisplayConfiguration
// function.
func (device *OLED128x64V2Bricklet) ReadPixelsLowLevel(xStart uint8, yStart uint8, xEnd uint8, yEnd uint8) (pixelsLength uint16, pixelsChunkOffset uint16, pixelsChunkData [480]bool, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, xStart)
	binary.Write(&buf, binary.LittleEndian, yStart)
	binary.Write(&buf, binary.LittleEndian, xEnd)
	binary.Write(&buf, binary.LittleEndian, yEnd)

	resultBytes, err := device.device.Get(uint8(FunctionReadPixelsLowLevel), buf.Bytes())
	if err != nil {
		return pixelsLength, pixelsChunkOffset, pixelsChunkData, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return pixelsLength, pixelsChunkOffset, pixelsChunkData, DeviceError(header.ErrorCode)
		}

		if header.Length != 72 {
			return pixelsLength, pixelsChunkOffset, pixelsChunkData, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 72)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &pixelsLength)
		binary.Read(resultBuf, binary.LittleEndian, &pixelsChunkOffset)
		copy(pixelsChunkData[:], ByteSliceToBoolSlice(resultBuf.Next(1*480/8)))

	}

	return pixelsLength, pixelsChunkOffset, pixelsChunkData, nil
}

// Reads pixels from the specified window.
//
// The pixels are read from the window line by line top to bottom
// and each line is read from left to right.
//
// If automatic draw is enabled (default) the pixels that are read are always the
// same that are shown on the display.
//
// If automatic draw is disabled the pixels are read from the internal buffer
// (see DrawBufferedFrame).
//
// Automatic draw can be configured with the SetDisplayConfiguration
// function.
func (device *OLED128x64V2Bricklet) ReadPixels(xStart uint8, yStart uint8, xEnd uint8, yEnd uint8) (pixels []bool, err error) {
	buf, _, err := device.device.GetHighLevel(func() (LowLevelResult, error) {
		pixelsLength, pixelsChunkOffset, pixelsChunkData, err := device.ReadPixelsLowLevel(xStart, yStart, xEnd, yEnd)

		if err != nil {
			return LowLevelResult{}, err
		}

		var lowLevelResults bytes.Buffer

		return LowLevelResult{
			uint64(pixelsLength),
			uint64(pixelsChunkOffset),
			BoolSliceToByteSlice(pixelsChunkData[:]),
			lowLevelResults.Bytes()}, nil
	},
		1,
		1)
	if err != nil {
		return ByteSliceToBoolSlice(buf), err
	}

	return ByteSliceToBoolSlice(buf), nil
}

// Clears the complete content of the display.
//
// If automatic draw is enabled (default) the pixels are directly cleared.
//
// If automatic draw is disabled the the internal buffer is cleared and
// the buffer is transferred to the display only after DrawBufferedFrame
// is called. This can be used to avoid flicker when drawing a complex frame in
// multiple steps.
//
// Automatic draw can be configured with the SetDisplayConfiguration
// function.
func (device *OLED128x64V2Bricklet) ClearDisplay() (err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Set(uint8(FunctionClearDisplay), buf.Bytes())
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

// Sets the configuration of the display.
//
// You can set a contrast value from 0 to 255 and you can invert the color
// (white/black) of the display.
//
// If automatic draw is set to *true*, the display is automatically updated with every
// call of WritePixels or WriteLine. If it is set to false, the
// changes are written into an internal buffer and only shown on the display after
// a call of DrawBufferedFrame.
func (device *OLED128x64V2Bricklet) SetDisplayConfiguration(contrast uint8, invert bool, automaticDraw bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, contrast)
	binary.Write(&buf, binary.LittleEndian, invert)
	binary.Write(&buf, binary.LittleEndian, automaticDraw)

	resultBytes, err := device.device.Set(uint8(FunctionSetDisplayConfiguration), buf.Bytes())
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

// Returns the configuration as set by SetDisplayConfiguration.
func (device *OLED128x64V2Bricklet) GetDisplayConfiguration() (contrast uint8, invert bool, automaticDraw bool, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetDisplayConfiguration), buf.Bytes())
	if err != nil {
		return contrast, invert, automaticDraw, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return contrast, invert, automaticDraw, DeviceError(header.ErrorCode)
		}

		if header.Length != 11 {
			return contrast, invert, automaticDraw, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 11)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &contrast)
		binary.Read(resultBuf, binary.LittleEndian, &invert)
		binary.Read(resultBuf, binary.LittleEndian, &automaticDraw)

	}

	return contrast, invert, automaticDraw, nil
}

// Writes text to a specific line with a specific position.
// The text can have a maximum of 22 characters.
//
// For example: (1, 10, Hello) will write *Hello* in the middle of the
// second line of the display.
//
// The display uses a special 5x7 pixel charset. You can view the characters
// of the charset in Brick Viewer.
//
// If automatic draw is enabled (default) the text is directly written to
// the screen. Only pixels that have actually changed are updated on the screen,
// the rest stays the same.
//
// If automatic draw is disabled the text is written to an internal buffer and
// the buffer is transferred to the display only after DrawBufferedFrame
// is called. This can be used to avoid flicker when drawing a complex frame in
// multiple steps.
//
// Automatic draw can be configured with the SetDisplayConfiguration
// function.
//
// The font conforms to code page 437.
func (device *OLED128x64V2Bricklet) WriteLine(line uint8, position uint8, text string) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, line)
	binary.Write(&buf, binary.LittleEndian, position)
	text_byte_slice, err := StringToByteSlice(text, 22)
	if err != nil {
		return
	}
	buf.Write(text_byte_slice)

	resultBytes, err := device.device.Set(uint8(FunctionWriteLine), buf.Bytes())
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

// Draws the currently buffered frame. Normally each call of WritePixels and
// WriteLine draws directly onto the display. If you turn automatic draw off
// (SetDisplayConfiguration), the data is written in an internal buffer and
// only transferred to the display by calling this function. This can be used to
// avoid flicker when drawing a complex frame in multiple steps.
//
// Set the `force complete redraw` to *true* to redraw the whole display
// instead of only the changed parts. Normally it should not be necessary to set this to
// *true*. It may only become necessary in case of stuck pixels because of errors.
func (device *OLED128x64V2Bricklet) DrawBufferedFrame(forceCompleteRedraw bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, forceCompleteRedraw)

	resultBytes, err := device.device.Set(uint8(FunctionDrawBufferedFrame), buf.Bytes())
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
func (device *OLED128x64V2Bricklet) GetSPITFPErrorCount() (errorCountAckChecksum uint32, errorCountMessageChecksum uint32, errorCountFrame uint32, errorCountOverflow uint32, err error) {
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
func (device *OLED128x64V2Bricklet) SetBootloaderMode(mode BootloaderMode) (status BootloaderStatus, err error) {
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
func (device *OLED128x64V2Bricklet) GetBootloaderMode() (mode BootloaderMode, err error) {
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
func (device *OLED128x64V2Bricklet) SetWriteFirmwarePointer(pointer uint32) (err error) {
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
func (device *OLED128x64V2Bricklet) WriteFirmware(data [64]uint8) (status uint8, err error) {
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
func (device *OLED128x64V2Bricklet) SetStatusLEDConfig(config StatusLEDConfig) (err error) {
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
func (device *OLED128x64V2Bricklet) GetStatusLEDConfig() (config StatusLEDConfig, err error) {
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
func (device *OLED128x64V2Bricklet) GetChipTemperature() (temperature int16, err error) {
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
func (device *OLED128x64V2Bricklet) Reset() (err error) {
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
func (device *OLED128x64V2Bricklet) WriteUID(uid uint32) (err error) {
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
func (device *OLED128x64V2Bricklet) ReadUID() (uid uint32, err error) {
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
func (device *OLED128x64V2Bricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
