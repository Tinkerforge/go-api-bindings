/* ***********************************************************
 * This file was automatically generated on 2021-01-15.      *
 *                                                           *
 * Go Bindings Version 2.0.10                                *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


// Three color 296x128 e-paper display.
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/EPaper296x128_Bricklet_Go.html.
package e_paper_296x128_bricklet

import (
	"encoding/binary"
	"bytes"
	"fmt"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionDraw Function = 1
	FunctionGetDrawStatus Function = 2
	FunctionWriteBlackWhiteLowLevel Function = 3
	FunctionReadBlackWhiteLowLevel Function = 4
	FunctionWriteColorLowLevel Function = 5
	FunctionReadColorLowLevel Function = 6
	FunctionFillDisplay Function = 7
	FunctionDrawText Function = 8
	FunctionDrawLine Function = 9
	FunctionDrawBox Function = 10
	FunctionSetUpdateMode Function = 12
	FunctionGetUpdateMode Function = 13
	FunctionSetDisplayType Function = 14
	FunctionGetDisplayType Function = 15
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
	FunctionCallbackDrawStatus Function = 11
)

type DrawStatus = uint8

const (
	DrawStatusIdle DrawStatus = 0
	DrawStatusCopying DrawStatus = 1
	DrawStatusDrawing DrawStatus = 2
)

type Color = uint8

const (
	ColorBlack Color = 0
	ColorWhite Color = 1
	ColorRed Color = 2
	ColorGray Color = 2
)

type Font = uint8

const (
	Font6x8 Font = 0
	Font6x16 Font = 1
	Font6x24 Font = 2
	Font6x32 Font = 3
	Font12x16 Font = 4
	Font12x24 Font = 5
	Font12x32 Font = 6
	Font18x24 Font = 7
	Font18x32 Font = 8
	Font24x32 Font = 9
)

type Orientation = uint8

const (
	OrientationHorizontal Orientation = 0
	OrientationVertical Orientation = 1
)

type UpdateMode = uint8

const (
	UpdateModeDefault UpdateMode = 0
	UpdateModeBlackWhite UpdateMode = 1
	UpdateModeDelta UpdateMode = 2
)

type DisplayType = uint8

const (
	DisplayTypeBlackWhiteRed DisplayType = 0
	DisplayTypeBlackWhiteGray DisplayType = 1
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

type EPaper296x128Bricklet struct {
	device Device
}
const DeviceIdentifier = 2146
const DeviceDisplayName = "E-Paper 296x128 Bricklet"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (EPaper296x128Bricklet, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{ 2,0,0 }, uid, &internalIPCon, 4, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return EPaper296x128Bricklet{}, err
	}
	dev.ResponseExpected[FunctionDraw] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetDrawStatus] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionWriteBlackWhiteLowLevel] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionReadBlackWhiteLowLevel] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionWriteColorLowLevel] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionReadColorLowLevel] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionFillDisplay] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionDrawText] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionDrawLine] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionDrawBox] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionSetUpdateMode] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetUpdateMode] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetDisplayType] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetDisplayType] = ResponseExpectedFlagAlwaysTrue;
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
	return EPaper296x128Bricklet{dev}, nil
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
func (device *EPaper296x128Bricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *EPaper296x128Bricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *EPaper296x128Bricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *EPaper296x128Bricklet) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// Callback for the current draw status. Will be called every time the
// draw status changes (see GetDrawStatus).
func (device *EPaper296x128Bricklet) RegisterDrawStatusCallback(fn func(DrawStatus)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 9 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var drawStatus DrawStatus
		binary.Read(buf, binary.LittleEndian, &drawStatus)
		fn(drawStatus)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackDrawStatus), wrapper)
}

// Remove a registered Draw Status callback.
func (device *EPaper296x128Bricklet) DeregisterDrawStatusCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackDrawStatus), registrationId)
}


// Draws the current black/white and red or gray buffer to the e-paper display.
// 
// The Bricklet does not have any double-buffering. You should not call
// this function while writing to the buffer. See GetDrawStatus.
func (device *EPaper296x128Bricklet) Draw() (err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Set(uint8(FunctionDraw), buf.Bytes())
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

// Returns one of three draw statuses:
// 
// * Idle
// * Copying: Data is being copied from the buffer of the Bricklet to the buffer of the display.
// * Drawing: The display is updating its content (during this phase the flickering etc happens).
// 
// You can write to the buffer (through one of the write or draw functions) when the status is
// either *idle* or *drawing*. You should not write to the buffer while it is being *copied* to the
// display. There is no double-buffering.
//
// Associated constants:
//
//	* DrawStatusIdle
//	* DrawStatusCopying
//	* DrawStatusDrawing
func (device *EPaper296x128Bricklet) GetDrawStatus() (drawStatus DrawStatus, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetDrawStatus), buf.Bytes())
	if err != nil {
		return drawStatus, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return drawStatus, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return drawStatus, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &drawStatus)

	}

	return drawStatus, nil
}

// Writes black/white pixels to the specified window into the buffer.
// 
// The pixels are written into the window line by line top to bottom
// and each line is written from left to right.
// 
// The value 0 (false) corresponds to a black pixel and the value 1 (true) to a
// white pixel.
// 
// This function writes the pixels into the black/white pixel buffer, to draw the
// buffer to the display use Draw.
// 
// Use WriteColor to write red or gray pixels.
func (device *EPaper296x128Bricklet) WriteBlackWhiteLowLevel(xStart uint16, yStart uint8, xEnd uint16, yEnd uint8, pixelsLength uint16, pixelsChunkOffset uint16, pixelsChunkData [432]bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, xStart);
	binary.Write(&buf, binary.LittleEndian, yStart);
	binary.Write(&buf, binary.LittleEndian, xEnd);
	binary.Write(&buf, binary.LittleEndian, yEnd);
	binary.Write(&buf, binary.LittleEndian, pixelsLength);
	binary.Write(&buf, binary.LittleEndian, pixelsChunkOffset);
	buf.Write(BoolSliceToByteSlice(pixelsChunkData[:]))

	resultBytes, err := device.device.Set(uint8(FunctionWriteBlackWhiteLowLevel), buf.Bytes())
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

// Writes black/white pixels to the specified window into the buffer.
// 
// The pixels are written into the window line by line top to bottom
// and each line is written from left to right.
// 
// The value 0 (false) corresponds to a black pixel and the value 1 (true) to a
// white pixel.
// 
// This function writes the pixels into the black/white pixel buffer, to draw the
// buffer to the display use Draw.
// 
// Use WriteColor to write red or gray pixels.
	func (device *EPaper296x128Bricklet) WriteBlackWhite(xStart uint16, yStart uint8, xEnd uint16, yEnd uint8, pixels []bool) (err error) {
		_, err = device.device.SetHighLevel(func(pixelsLength uint64, pixelsChunkOffset uint64, pixelsChunkData []byte) (LowLevelWriteResult, error) {
			arr := [432]bool{}
			copy(arr[:], ByteSliceToBoolSlice(pixelsChunkData))

			err := device.WriteBlackWhiteLowLevel(xStart, yStart, xEnd, yEnd, uint16(pixelsLength), uint16(pixelsChunkOffset), arr)

			var lowLevelResults bytes.Buffer
			

			return LowLevelWriteResult{
				uint64(432),
				lowLevelResults.Bytes()}, err
		}, 0, 1, 432, BoolSliceToByteSlice(pixels))

		if err != nil {
			return
		}

		
		
		
		return
	}

// Returns the current content of the black/white pixel buffer for the specified window.
// 
// The pixels are read into the window line by line top to bottom and
// each line is read from left to right.
// 
// The current content of the buffer does not have to be the current content of the display.
// It is possible that the data was not drawn to the display yet and after a restart of
// the Bricklet the buffer will be reset to black, while the display retains its content.
func (device *EPaper296x128Bricklet) ReadBlackWhiteLowLevel(xStart uint16, yStart uint8, xEnd uint16, yEnd uint8) (pixelsLength uint16, pixelsChunkOffset uint16, pixelsChunkData [464]bool, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, xStart);
	binary.Write(&buf, binary.LittleEndian, yStart);
	binary.Write(&buf, binary.LittleEndian, xEnd);
	binary.Write(&buf, binary.LittleEndian, yEnd);

	resultBytes, err := device.device.Get(uint8(FunctionReadBlackWhiteLowLevel), buf.Bytes())
	if err != nil {
		return pixelsLength, pixelsChunkOffset, pixelsChunkData, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 70 {
			return pixelsLength, pixelsChunkOffset, pixelsChunkData, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 70)
		}

		if header.ErrorCode != 0 {
			return pixelsLength, pixelsChunkOffset, pixelsChunkData, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &pixelsLength)
		binary.Read(resultBuf, binary.LittleEndian, &pixelsChunkOffset)
		copy(pixelsChunkData[:], ByteSliceToBoolSlice(resultBuf.Next(1 * 464/8)))

	}

	return pixelsLength, pixelsChunkOffset, pixelsChunkData, nil
}

// Returns the current content of the black/white pixel buffer for the specified window.
// 
// The pixels are read into the window line by line top to bottom and
// each line is read from left to right.
// 
// The current content of the buffer does not have to be the current content of the display.
// It is possible that the data was not drawn to the display yet and after a restart of
// the Bricklet the buffer will be reset to black, while the display retains its content.
	func (device *EPaper296x128Bricklet) ReadBlackWhite(xStart uint16, yStart uint8, xEnd uint16, yEnd uint8) (pixels []bool, err error) {
		buf, _, err := device.device.GetHighLevel(func() (LowLevelResult, error) {
			pixelsLength, pixelsChunkOffset, pixelsChunkData, err := device.ReadBlackWhiteLowLevel(xStart, yStart, xEnd, yEnd)

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

// The E-Paper 296x128 Bricklet is available with the colors black/white/red and
// black/white/gray. Depending on the model this function writes either red or
// gray pixels to the specified window into the buffer.
// 
// The pixels are written into the window line by line top to bottom
// and each line is written from left to right.
// 
// The value 0 (false) means that this pixel does not have color. It will be either black
// or white (see WriteBlackWhite). The value 1 (true) corresponds to a red or gray
// pixel, depending on the Bricklet model.
// 
// This function writes the pixels into the red or gray pixel buffer, to draw the buffer
// to the display use Draw.
// 
// Use WriteBlackWhite to write black/white pixels.
func (device *EPaper296x128Bricklet) WriteColorLowLevel(xStart uint16, yStart uint8, xEnd uint16, yEnd uint8, pixelsLength uint16, pixelsChunkOffset uint16, pixelsChunkData [432]bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, xStart);
	binary.Write(&buf, binary.LittleEndian, yStart);
	binary.Write(&buf, binary.LittleEndian, xEnd);
	binary.Write(&buf, binary.LittleEndian, yEnd);
	binary.Write(&buf, binary.LittleEndian, pixelsLength);
	binary.Write(&buf, binary.LittleEndian, pixelsChunkOffset);
	buf.Write(BoolSliceToByteSlice(pixelsChunkData[:]))

	resultBytes, err := device.device.Set(uint8(FunctionWriteColorLowLevel), buf.Bytes())
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

// The E-Paper 296x128 Bricklet is available with the colors black/white/red and
// black/white/gray. Depending on the model this function writes either red or
// gray pixels to the specified window into the buffer.
// 
// The pixels are written into the window line by line top to bottom
// and each line is written from left to right.
// 
// The value 0 (false) means that this pixel does not have color. It will be either black
// or white (see WriteBlackWhite). The value 1 (true) corresponds to a red or gray
// pixel, depending on the Bricklet model.
// 
// This function writes the pixels into the red or gray pixel buffer, to draw the buffer
// to the display use Draw.
// 
// Use WriteBlackWhite to write black/white pixels.
	func (device *EPaper296x128Bricklet) WriteColor(xStart uint16, yStart uint8, xEnd uint16, yEnd uint8, pixels []bool) (err error) {
		_, err = device.device.SetHighLevel(func(pixelsLength uint64, pixelsChunkOffset uint64, pixelsChunkData []byte) (LowLevelWriteResult, error) {
			arr := [432]bool{}
			copy(arr[:], ByteSliceToBoolSlice(pixelsChunkData))

			err := device.WriteColorLowLevel(xStart, yStart, xEnd, yEnd, uint16(pixelsLength), uint16(pixelsChunkOffset), arr)

			var lowLevelResults bytes.Buffer
			

			return LowLevelWriteResult{
				uint64(432),
				lowLevelResults.Bytes()}, err
		}, 2, 1, 432, BoolSliceToByteSlice(pixels))

		if err != nil {
			return
		}

		
		
		
		return
	}

// Returns the current content of the red or gray pixel buffer for the specified window.
// 
// The pixels are written into the window line by line top to bottom
// and each line is written from left to right.
// 
// The current content of the buffer does not have to be the current content of the display.
// It is possible that the data was not drawn to the display yet and after a restart of
// the Bricklet the buffer will be reset to black, while the display retains its content.
func (device *EPaper296x128Bricklet) ReadColorLowLevel(xStart uint16, yStart uint8, xEnd uint16, yEnd uint8) (pixelsLength uint16, pixelsChunkOffset uint16, pixelsChunkData [464]bool, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, xStart);
	binary.Write(&buf, binary.LittleEndian, yStart);
	binary.Write(&buf, binary.LittleEndian, xEnd);
	binary.Write(&buf, binary.LittleEndian, yEnd);

	resultBytes, err := device.device.Get(uint8(FunctionReadColorLowLevel), buf.Bytes())
	if err != nil {
		return pixelsLength, pixelsChunkOffset, pixelsChunkData, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 70 {
			return pixelsLength, pixelsChunkOffset, pixelsChunkData, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 70)
		}

		if header.ErrorCode != 0 {
			return pixelsLength, pixelsChunkOffset, pixelsChunkData, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &pixelsLength)
		binary.Read(resultBuf, binary.LittleEndian, &pixelsChunkOffset)
		copy(pixelsChunkData[:], ByteSliceToBoolSlice(resultBuf.Next(1 * 464/8)))

	}

	return pixelsLength, pixelsChunkOffset, pixelsChunkData, nil
}

// Returns the current content of the red or gray pixel buffer for the specified window.
// 
// The pixels are written into the window line by line top to bottom
// and each line is written from left to right.
// 
// The current content of the buffer does not have to be the current content of the display.
// It is possible that the data was not drawn to the display yet and after a restart of
// the Bricklet the buffer will be reset to black, while the display retains its content.
	func (device *EPaper296x128Bricklet) ReadColor(xStart uint16, yStart uint8, xEnd uint16, yEnd uint8) (pixels []bool, err error) {
		buf, _, err := device.device.GetHighLevel(func() (LowLevelResult, error) {
			pixelsLength, pixelsChunkOffset, pixelsChunkData, err := device.ReadColorLowLevel(xStart, yStart, xEnd, yEnd)

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
			3,
			1)
		if err != nil {
			return ByteSliceToBoolSlice(buf), err
		}
		
		
		return ByteSliceToBoolSlice(buf), nil
	}

// Fills the complete content of the display with the given color.
// 
// This function writes the pixels into the black/white/red|gray pixel buffer, to draw the buffer
// to the display use Draw.
//
// Associated constants:
//
//	* ColorBlack
//	* ColorWhite
//	* ColorRed
//	* ColorGray
func (device *EPaper296x128Bricklet) FillDisplay(color Color) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, color);

	resultBytes, err := device.device.Set(uint8(FunctionFillDisplay), buf.Bytes())
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

// Draws a text with up to 50 characters at the pixel position (x, y).
// 
// You can use one of 9 different font sizes and draw the text in
// black/white/red|gray. The text can be drawn horizontal or vertical.
// 
// This function writes the pixels into the black/white/red|gray pixel buffer, to draw the buffer
// to the display use Draw.
//
// Associated constants:
//
//	* Font6x8
//	* Font6x16
//	* Font6x24
//	* Font6x32
//	* Font12x16
//	* Font12x24
//	* Font12x32
//	* Font18x24
//	* Font18x32
//	* Font24x32
//	* ColorBlack
//	* ColorWhite
//	* ColorRed
//	* ColorGray
//	* OrientationHorizontal
//	* OrientationVertical
func (device *EPaper296x128Bricklet) DrawText(positionX uint16, positionY uint8, font Font, color Color, orientation Orientation, text string) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, positionX);
	binary.Write(&buf, binary.LittleEndian, positionY);
	binary.Write(&buf, binary.LittleEndian, font);
	binary.Write(&buf, binary.LittleEndian, color);
	binary.Write(&buf, binary.LittleEndian, orientation);
	text_byte_slice, err := StringToByteSlice(text, 50)
	if err != nil { return }
	buf.Write(text_byte_slice)

	resultBytes, err := device.device.Set(uint8(FunctionDrawText), buf.Bytes())
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

// Draws a line from (x, y)-start to (x, y)-end in the given color.
// 
// This function writes the pixels into the black/white/red|gray pixel buffer, to draw the buffer
// to the display use Draw.
//
// Associated constants:
//
//	* ColorBlack
//	* ColorWhite
//	* ColorRed
//	* ColorGray
func (device *EPaper296x128Bricklet) DrawLine(positionXStart uint16, positionYStart uint8, positionXEnd uint16, positionYEnd uint8, color Color) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, positionXStart);
	binary.Write(&buf, binary.LittleEndian, positionYStart);
	binary.Write(&buf, binary.LittleEndian, positionXEnd);
	binary.Write(&buf, binary.LittleEndian, positionYEnd);
	binary.Write(&buf, binary.LittleEndian, color);

	resultBytes, err := device.device.Set(uint8(FunctionDrawLine), buf.Bytes())
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

// Draws a box from (x, y)-start to (x, y)-end in the given color.
// 
// If you set fill to true, the box will be filled with the
// color. Otherwise only the outline will be drawn.
// 
// This function writes the pixels into the black/white/red|gray pixel buffer, to draw the buffer
// to the display use Draw.
//
// Associated constants:
//
//	* ColorBlack
//	* ColorWhite
//	* ColorRed
//	* ColorGray
func (device *EPaper296x128Bricklet) DrawBox(positionXStart uint16, positionYStart uint8, positionXEnd uint16, positionYEnd uint8, fill bool, color Color) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, positionXStart);
	binary.Write(&buf, binary.LittleEndian, positionYStart);
	binary.Write(&buf, binary.LittleEndian, positionXEnd);
	binary.Write(&buf, binary.LittleEndian, positionYEnd);
	binary.Write(&buf, binary.LittleEndian, fill);
	binary.Write(&buf, binary.LittleEndian, color);

	resultBytes, err := device.device.Set(uint8(FunctionDrawBox), buf.Bytes())
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

// Note
//  The default update mode corresponds to the default e-paper display
//  manufacturer settings. All of the other modes are experimental and
//  will result in increased ghosting and possibly other long-term
//  side effects.
// 
//  If you want to know more about the inner workings of an e-paper display
//  take a look at this excellent video from Ben Krasnow:
//  https://www.youtube.com/watch?v=MsbiO8EAsGw.
// 
//  If you are not sure about this option, leave the update mode at default.
// 
// Currently there are three update modes available:
// 
// * Default: Settings as given by the manufacturer. An update will take about
//   7.5 seconds and during the update the screen will flicker several times.
// * Black/White: This will only update the black/white pixel. It uses the manufacturer
//   settings for black/white and ignores the red or gray pixel buffer. With this mode the
//   display will flicker once and it takes about 2.5 seconds. Compared to the default settings
//   there is more ghosting.
// * Delta: This will only update the black/white pixel. It uses an aggressive method where
//   the changes are not applied for a whole buffer but only for the delta between the last
//   and the next buffer. With this mode the display will not flicker during an update and
//   it takes about 900-950ms. Compared to the other two settings there is more ghosting. This
//   mode can be used for something like a flicker-free live update of a text.
// 
// With the black/white/red display if you use either the black/white or the delta mode,
// after a while of going back and forth between black and white the white color will
// start to appear red-ish or pink-ish.
// 
// If you use the aggressive delta mode and rapidly change the content, we recommend that you
// change back to the default mode every few hours and in the default mode cycle between the
// three available colors a few times. This will get rid of the ghosting and after that you can
// go back to the delta mode with flicker-free updates.
//
// Associated constants:
//
//	* UpdateModeDefault
//	* UpdateModeBlackWhite
//	* UpdateModeDelta
func (device *EPaper296x128Bricklet) SetUpdateMode(updateMode UpdateMode) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, updateMode);

	resultBytes, err := device.device.Set(uint8(FunctionSetUpdateMode), buf.Bytes())
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

// Returns the update mode as set by SetUpdateMode.
//
// Associated constants:
//
//	* UpdateModeDefault
//	* UpdateModeBlackWhite
//	* UpdateModeDelta
func (device *EPaper296x128Bricklet) GetUpdateMode() (updateMode UpdateMode, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetUpdateMode), buf.Bytes())
	if err != nil {
		return updateMode, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return updateMode, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return updateMode, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &updateMode)

	}

	return updateMode, nil
}

// Sets the type of the display. The e-paper display is available
// in black/white/red and black/white/gray. This will be factory set
// during the flashing and testing phase. The value is saved in
// non-volatile memory and will stay after a power cycle.
//
// Associated constants:
//
//	* DisplayTypeBlackWhiteRed
//	* DisplayTypeBlackWhiteGray
func (device *EPaper296x128Bricklet) SetDisplayType(displayType DisplayType) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, displayType);

	resultBytes, err := device.device.Set(uint8(FunctionSetDisplayType), buf.Bytes())
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

// Returns the type of the e-paper display. It can either be
// black/white/red or black/white/gray.
//
// Associated constants:
//
//	* DisplayTypeBlackWhiteRed
//	* DisplayTypeBlackWhiteGray
func (device *EPaper296x128Bricklet) GetDisplayType() (displayType DisplayType, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetDisplayType), buf.Bytes())
	if err != nil {
		return displayType, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return displayType, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return displayType, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &displayType)

	}

	return displayType, nil
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
func (device *EPaper296x128Bricklet) GetSPITFPErrorCount() (errorCountAckChecksum uint32, errorCountMessageChecksum uint32, errorCountFrame uint32, errorCountOverflow uint32, err error) {
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
func (device *EPaper296x128Bricklet) SetBootloaderMode(mode BootloaderMode) (status BootloaderStatus, err error) {
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
func (device *EPaper296x128Bricklet) GetBootloaderMode() (mode BootloaderMode, err error) {
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
func (device *EPaper296x128Bricklet) SetWriteFirmwarePointer(pointer uint32) (err error) {
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
func (device *EPaper296x128Bricklet) WriteFirmware(data [64]uint8) (status uint8, err error) {
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
func (device *EPaper296x128Bricklet) SetStatusLEDConfig(config StatusLEDConfig) (err error) {
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
func (device *EPaper296x128Bricklet) GetStatusLEDConfig() (config StatusLEDConfig, err error) {
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
func (device *EPaper296x128Bricklet) GetChipTemperature() (temperature int16, err error) {
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
func (device *EPaper296x128Bricklet) Reset() (err error) {
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
func (device *EPaper296x128Bricklet) WriteUID(uid uint32) (err error) {
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
func (device *EPaper296x128Bricklet) ReadUID() (uid uint32, err error) {
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
func (device *EPaper296x128Bricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
