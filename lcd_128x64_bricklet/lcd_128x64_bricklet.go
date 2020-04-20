/* ***********************************************************
 * This file was automatically generated on 2020-04-20.      *
 *                                                           *
 * Go Bindings Version 2.0.7                                 *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


// 7.1cm (2.8") display with 128x64 pixel and touch screen.
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/LCD128x64_Bricklet_Go.html.
package lcd_128x64_bricklet

import (
	"encoding/binary"
	"bytes"
	"fmt"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionWritePixelsLowLevel Function = 1
	FunctionReadPixelsLowLevel Function = 2
	FunctionClearDisplay Function = 3
	FunctionSetDisplayConfiguration Function = 4
	FunctionGetDisplayConfiguration Function = 5
	FunctionWriteLine Function = 6
	FunctionDrawBufferedFrame Function = 7
	FunctionGetTouchPosition Function = 8
	FunctionSetTouchPositionCallbackConfiguration Function = 9
	FunctionGetTouchPositionCallbackConfiguration Function = 10
	FunctionGetTouchGesture Function = 12
	FunctionSetTouchGestureCallbackConfiguration Function = 13
	FunctionGetTouchGestureCallbackConfiguration Function = 14
	FunctionDrawLine Function = 16
	FunctionDrawBox Function = 17
	FunctionDrawText Function = 18
	FunctionSetGUIButton Function = 19
	FunctionGetGUIButton Function = 20
	FunctionRemoveGUIButton Function = 21
	FunctionSetGUIButtonPressedCallbackConfiguration Function = 22
	FunctionGetGUIButtonPressedCallbackConfiguration Function = 23
	FunctionGetGUIButtonPressed Function = 24
	FunctionSetGUISlider Function = 26
	FunctionGetGUISlider Function = 27
	FunctionRemoveGUISlider Function = 28
	FunctionSetGUISliderValueCallbackConfiguration Function = 29
	FunctionGetGUISliderValueCallbackConfiguration Function = 30
	FunctionGetGUISliderValue Function = 31
	FunctionSetGUITabConfiguration Function = 33
	FunctionGetGUITabConfiguration Function = 34
	FunctionSetGUITabText Function = 35
	FunctionGetGUITabText Function = 36
	FunctionSetGUITabIcon Function = 37
	FunctionGetGUITabIcon Function = 38
	FunctionRemoveGUITab Function = 39
	FunctionSetGUITabSelected Function = 40
	FunctionSetGUITabSelectedCallbackConfiguration Function = 41
	FunctionGetGUITabSelectedCallbackConfiguration Function = 42
	FunctionGetGUITabSelected Function = 43
	FunctionSetGUIGraphConfiguration Function = 45
	FunctionGetGUIGraphConfiguration Function = 46
	FunctionSetGUIGraphDataLowLevel Function = 47
	FunctionGetGUIGraphDataLowLevel Function = 48
	FunctionRemoveGUIGraph Function = 49
	FunctionRemoveAllGUI Function = 50
	FunctionSetTouchLEDConfig Function = 51
	FunctionGetTouchLEDConfig Function = 52
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
	FunctionCallbackTouchPosition Function = 11
	FunctionCallbackTouchGesture Function = 15
	FunctionCallbackGUIButtonPressed Function = 25
	FunctionCallbackGUISliderValue Function = 32
	FunctionCallbackGUITabSelected Function = 44
)

type Gesture = uint8

const (
	GestureLeftToRight Gesture = 0
	GestureRightToLeft Gesture = 1
	GestureTopToBottom Gesture = 2
	GestureBottomToTop Gesture = 3
)

type Color = bool

const (
	ColorWhite Color = false
	ColorBlack Color = true
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

type Direction = uint8

const (
	DirectionHorizontal Direction = 0
	DirectionVertical Direction = 1
)

type ChangeTabOn = uint8

const (
	ChangeTabOnClick ChangeTabOn = 1
	ChangeTabOnSwipe ChangeTabOn = 2
	ChangeTabOnClickAndSwipe ChangeTabOn = 3
)

type GraphType = uint8

const (
	GraphTypeDot GraphType = 0
	GraphTypeLine GraphType = 1
	GraphTypeBar GraphType = 2
)

type TouchLEDConfig = uint8

const (
	TouchLEDConfigOff TouchLEDConfig = 0
	TouchLEDConfigOn TouchLEDConfig = 1
	TouchLEDConfigShowHeartbeat TouchLEDConfig = 2
	TouchLEDConfigShowTouch TouchLEDConfig = 3
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

type LCD128x64Bricklet struct {
	device Device
}
const DeviceIdentifier = 298
const DeviceDisplayName = "LCD 128x64 Bricklet"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (LCD128x64Bricklet, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{ 2,0,1 }, uid, &internalIPCon, 4, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return LCD128x64Bricklet{}, err
	}
	dev.ResponseExpected[FunctionWritePixelsLowLevel] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionReadPixelsLowLevel] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionClearDisplay] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionSetDisplayConfiguration] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetDisplayConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionWriteLine] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionDrawBufferedFrame] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetTouchPosition] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetTouchPositionCallbackConfiguration] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetTouchPositionCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetTouchGesture] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetTouchGestureCallbackConfiguration] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetTouchGestureCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionDrawLine] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionDrawBox] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionDrawText] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionSetGUIButton] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetGUIButton] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionRemoveGUIButton] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionSetGUIButtonPressedCallbackConfiguration] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetGUIButtonPressedCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetGUIButtonPressed] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetGUISlider] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetGUISlider] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionRemoveGUISlider] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionSetGUISliderValueCallbackConfiguration] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetGUISliderValueCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetGUISliderValue] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetGUITabConfiguration] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetGUITabConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetGUITabText] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetGUITabText] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetGUITabIcon] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetGUITabIcon] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionRemoveGUITab] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionSetGUITabSelected] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionSetGUITabSelectedCallbackConfiguration] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetGUITabSelectedCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetGUITabSelected] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetGUIGraphConfiguration] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetGUIGraphConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetGUIGraphDataLowLevel] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetGUIGraphDataLowLevel] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionRemoveGUIGraph] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionRemoveAllGUI] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionSetTouchLEDConfig] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetTouchLEDConfig] = ResponseExpectedFlagAlwaysTrue;
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
	return LCD128x64Bricklet{dev}, nil
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
func (device *LCD128x64Bricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *LCD128x64Bricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *LCD128x64Bricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *LCD128x64Bricklet) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// This callback is triggered periodically with the period that is set by
// SetTouchPositionCallbackConfiguration. The parameters are the
// same as for GetTouchPosition.
func (device *LCD128x64Bricklet) RegisterTouchPositionCallback(fn func(uint16, uint16, uint16, uint32)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 18 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var pressure uint16
		var x uint16
		var y uint16
		var age uint32
		binary.Read(buf, binary.LittleEndian, &pressure)
		binary.Read(buf, binary.LittleEndian, &x)
		binary.Read(buf, binary.LittleEndian, &y)
		binary.Read(buf, binary.LittleEndian, &age)
		fn(pressure, x, y, age)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackTouchPosition), wrapper)
}

// Remove a registered Touch Position callback.
func (device *LCD128x64Bricklet) DeregisterTouchPositionCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackTouchPosition), registrationId)
}


// This callback is triggered periodically with the period that is set by
// SetTouchGestureCallbackConfiguration. The parameters are the
// same as for GetTouchGesture.
func (device *LCD128x64Bricklet) RegisterTouchGestureCallback(fn func(Gesture, uint32, uint16, uint16, uint16, uint16, uint16, uint32)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 27 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var gesture Gesture
		var duration uint32
		var pressureMax uint16
		var xStart uint16
		var yStart uint16
		var xEnd uint16
		var yEnd uint16
		var age uint32
		binary.Read(buf, binary.LittleEndian, &gesture)
		binary.Read(buf, binary.LittleEndian, &duration)
		binary.Read(buf, binary.LittleEndian, &pressureMax)
		binary.Read(buf, binary.LittleEndian, &xStart)
		binary.Read(buf, binary.LittleEndian, &yStart)
		binary.Read(buf, binary.LittleEndian, &xEnd)
		binary.Read(buf, binary.LittleEndian, &yEnd)
		binary.Read(buf, binary.LittleEndian, &age)
		fn(gesture, duration, pressureMax, xStart, yStart, xEnd, yEnd, age)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackTouchGesture), wrapper)
}

// Remove a registered Touch Gesture callback.
func (device *LCD128x64Bricklet) DeregisterTouchGestureCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackTouchGesture), registrationId)
}


// This callback is triggered periodically with the period that is set by
// SetGUIButtonPressedCallbackConfiguration. The parameters are the
// same as for GetGUIButtonPressed.
// 
// .. versionadded:: 2.0.2$nbsp;(Plugin)
func (device *LCD128x64Bricklet) RegisterGUIButtonPressedCallback(fn func(uint8, bool)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 10 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var index uint8
		var pressed bool
		binary.Read(buf, binary.LittleEndian, &index)
		binary.Read(buf, binary.LittleEndian, &pressed)
		fn(index, pressed)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackGUIButtonPressed), wrapper)
}

// Remove a registered GUI Button Pressed callback.
func (device *LCD128x64Bricklet) DeregisterGUIButtonPressedCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackGUIButtonPressed), registrationId)
}


// This callback is triggered periodically with the period that is set by
// SetGUISliderValueCallbackConfiguration. The parameters are the
// same as for GetGUISliderValue.
// 
// .. versionadded:: 2.0.2$nbsp;(Plugin)
func (device *LCD128x64Bricklet) RegisterGUISliderValueCallback(fn func(uint8, uint8)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 10 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var index uint8
		var value uint8
		binary.Read(buf, binary.LittleEndian, &index)
		binary.Read(buf, binary.LittleEndian, &value)
		fn(index, value)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackGUISliderValue), wrapper)
}

// Remove a registered GUI Slider Value callback.
func (device *LCD128x64Bricklet) DeregisterGUISliderValueCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackGUISliderValue), registrationId)
}


// This callback is triggered periodically with the period that is set by
// SetGUITabSelectedCallbackConfiguration. The parameters are the
// same as for GetGUITabSelected.
// 
// .. versionadded:: 2.0.2$nbsp;(Plugin)
func (device *LCD128x64Bricklet) RegisterGUITabSelectedCallback(fn func(int8)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 9 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var index int8
		binary.Read(buf, binary.LittleEndian, &index)
		fn(index)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackGUITabSelected), wrapper)
}

// Remove a registered GUI Tab Selected callback.
func (device *LCD128x64Bricklet) DeregisterGUITabSelectedCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackGUITabSelected), registrationId)
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
func (device *LCD128x64Bricklet) WritePixelsLowLevel(xStart uint8, yStart uint8, xEnd uint8, yEnd uint8, pixelsLength uint16, pixelsChunkOffset uint16, pixelsChunkData [448]bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, xStart);
	binary.Write(&buf, binary.LittleEndian, yStart);
	binary.Write(&buf, binary.LittleEndian, xEnd);
	binary.Write(&buf, binary.LittleEndian, yEnd);
	binary.Write(&buf, binary.LittleEndian, pixelsLength);
	binary.Write(&buf, binary.LittleEndian, pixelsChunkOffset);
	buf.Write(BoolSliceToByteSlice(pixelsChunkData[:]))

	resultBytes, err := device.device.Set(uint8(FunctionWritePixelsLowLevel), buf.Bytes())
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
	func (device *LCD128x64Bricklet) WritePixels(xStart uint8, yStart uint8, xEnd uint8, yEnd uint8, pixels []bool) (err error) {
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
func (device *LCD128x64Bricklet) ReadPixelsLowLevel(xStart uint8, yStart uint8, xEnd uint8, yEnd uint8) (pixelsLength uint16, pixelsChunkOffset uint16, pixelsChunkData [480]bool, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, xStart);
	binary.Write(&buf, binary.LittleEndian, yStart);
	binary.Write(&buf, binary.LittleEndian, xEnd);
	binary.Write(&buf, binary.LittleEndian, yEnd);

	resultBytes, err := device.device.Get(uint8(FunctionReadPixelsLowLevel), buf.Bytes())
	if err != nil {
		return pixelsLength, pixelsChunkOffset, pixelsChunkData, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 72 {
			return pixelsLength, pixelsChunkOffset, pixelsChunkData, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 72)
		}

		if header.ErrorCode != 0 {
			return pixelsLength, pixelsChunkOffset, pixelsChunkData, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &pixelsLength)
		binary.Read(resultBuf, binary.LittleEndian, &pixelsChunkOffset)
		copy(pixelsChunkData[:], ByteSliceToBoolSlice(resultBuf.Next(1 * 480/8)))

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
	func (device *LCD128x64Bricklet) ReadPixels(xStart uint8, yStart uint8, xEnd uint8, yEnd uint8) (pixels []bool, err error) {
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
func (device *LCD128x64Bricklet) ClearDisplay() (err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Set(uint8(FunctionClearDisplay), buf.Bytes())
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

// Sets the configuration of the display.
// 
// If automatic draw is set to *true*, the display is automatically updated with every
// call of WritePixels and WriteLine. If it is set to false, the
// changes are written into an internal buffer and only shown on the display after
// a call of DrawBufferedFrame.
func (device *LCD128x64Bricklet) SetDisplayConfiguration(contrast uint8, backlight uint8, invert bool, automaticDraw bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, contrast);
	binary.Write(&buf, binary.LittleEndian, backlight);
	binary.Write(&buf, binary.LittleEndian, invert);
	binary.Write(&buf, binary.LittleEndian, automaticDraw);

	resultBytes, err := device.device.Set(uint8(FunctionSetDisplayConfiguration), buf.Bytes())
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

// Returns the configuration as set by SetDisplayConfiguration.
func (device *LCD128x64Bricklet) GetDisplayConfiguration() (contrast uint8, backlight uint8, invert bool, automaticDraw bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetDisplayConfiguration), buf.Bytes())
	if err != nil {
		return contrast, backlight, invert, automaticDraw, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 12 {
			return contrast, backlight, invert, automaticDraw, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		if header.ErrorCode != 0 {
			return contrast, backlight, invert, automaticDraw, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &contrast)
		binary.Read(resultBuf, binary.LittleEndian, &backlight)
		binary.Read(resultBuf, binary.LittleEndian, &invert)
		binary.Read(resultBuf, binary.LittleEndian, &automaticDraw)

	}

	return contrast, backlight, invert, automaticDraw, nil
}

// Writes text to a specific line with a specific position.
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
// This function is a 1:1 replacement for the function with the same name
// in the LCD 20x4 Bricklet. You can draw text at a specific pixel position
// and with different font sizes with the DrawText function.
func (device *LCD128x64Bricklet) WriteLine(line uint8, position uint8, text string) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, line);
	binary.Write(&buf, binary.LittleEndian, position);
	text_byte_slice, err := StringToByteSlice(text, 22)
	if err != nil { return }
	buf.Write(text_byte_slice)

	resultBytes, err := device.device.Set(uint8(FunctionWriteLine), buf.Bytes())
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

// Draws the currently buffered frame. Normally each call of WritePixels and
// WriteLine draws directly onto the display. If you turn automatic draw off
// (SetDisplayConfiguration), the data is written in an internal buffer and
// only transferred to the display by calling this function. This can be used to
// avoid flicker when drawing a complex frame in multiple steps.
// 
// Set the `force complete redraw` to *true* to redraw the whole display
// instead of only the changed parts. Normally it should not be necessary to set this to
// *true*. It may only become necessary in case of stuck pixels because of errors.
func (device *LCD128x64Bricklet) DrawBufferedFrame(forceCompleteRedraw bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, forceCompleteRedraw);

	resultBytes, err := device.device.Set(uint8(FunctionDrawBufferedFrame), buf.Bytes())
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

// Returns the last valid touch position:
// 
// * Pressure: Amount of pressure applied by the user
// * X: Touch position on x-axis
// * Y: Touch position on y-axis
// * Age: Age of touch press (how long ago it was)
func (device *LCD128x64Bricklet) GetTouchPosition() (pressure uint16, x uint16, y uint16, age uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetTouchPosition), buf.Bytes())
	if err != nil {
		return pressure, x, y, age, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 18 {
			return pressure, x, y, age, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 18)
		}

		if header.ErrorCode != 0 {
			return pressure, x, y, age, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &pressure)
		binary.Read(resultBuf, binary.LittleEndian, &x)
		binary.Read(resultBuf, binary.LittleEndian, &y)
		binary.Read(resultBuf, binary.LittleEndian, &age)

	}

	return pressure, x, y, age, nil
}

// The period is the period with which the RegisterTouchPositionCallback callback
// is triggered periodically. A value of 0 turns the callback off.
// 
// If the `value has to change`-parameter is set to true, the callback is only
// triggered after the value has changed. If the value didn't change within the
// period, the callback is triggered immediately on change.
// 
// If it is set to false, the callback is continuously triggered with the period,
// independent of the value.
func (device *LCD128x64Bricklet) SetTouchPositionCallbackConfiguration(period uint32, valueHasToChange bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);
	binary.Write(&buf, binary.LittleEndian, valueHasToChange);

	resultBytes, err := device.device.Set(uint8(FunctionSetTouchPositionCallbackConfiguration), buf.Bytes())
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
// SetTouchPositionCallbackConfiguration.
func (device *LCD128x64Bricklet) GetTouchPositionCallbackConfiguration() (period uint32, valueHasToChange bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetTouchPositionCallbackConfiguration), buf.Bytes())
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

// Returns one of four touch gestures that can be automatically detected by the Bricklet.
// 
// The gestures are swipes from left to right, right to left, top to bottom and bottom to top.
// 
// Additionally to the gestures a vector with a start and end position of the gesture is
// provided. You can use this vector do determine a more exact location of the gesture (e.g.
// the swipe from top to bottom was on the left or right part of the screen).
// 
// The age parameter corresponds to the age of gesture (how long ago it was).
//
// Associated constants:
//
//	* GestureLeftToRight
//	* GestureRightToLeft
//	* GestureTopToBottom
//	* GestureBottomToTop
func (device *LCD128x64Bricklet) GetTouchGesture() (gesture Gesture, duration uint32, pressureMax uint16, xStart uint16, yStart uint16, xEnd uint16, yEnd uint16, age uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetTouchGesture), buf.Bytes())
	if err != nil {
		return gesture, duration, pressureMax, xStart, yStart, xEnd, yEnd, age, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 27 {
			return gesture, duration, pressureMax, xStart, yStart, xEnd, yEnd, age, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 27)
		}

		if header.ErrorCode != 0 {
			return gesture, duration, pressureMax, xStart, yStart, xEnd, yEnd, age, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &gesture)
		binary.Read(resultBuf, binary.LittleEndian, &duration)
		binary.Read(resultBuf, binary.LittleEndian, &pressureMax)
		binary.Read(resultBuf, binary.LittleEndian, &xStart)
		binary.Read(resultBuf, binary.LittleEndian, &yStart)
		binary.Read(resultBuf, binary.LittleEndian, &xEnd)
		binary.Read(resultBuf, binary.LittleEndian, &yEnd)
		binary.Read(resultBuf, binary.LittleEndian, &age)

	}

	return gesture, duration, pressureMax, xStart, yStart, xEnd, yEnd, age, nil
}

// The period is the period with which the RegisterTouchGestureCallback callback
// is triggered periodically. A value of 0 turns the callback off.
// 
// If the `value has to change`-parameter is set to true, the callback is only
// triggered after the value has changed. If the value didn't change within the
// period, the callback is triggered immediately on change.
// 
// If it is set to false, the callback is continuously triggered with the period,
// independent of the value.
func (device *LCD128x64Bricklet) SetTouchGestureCallbackConfiguration(period uint32, valueHasToChange bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);
	binary.Write(&buf, binary.LittleEndian, valueHasToChange);

	resultBytes, err := device.device.Set(uint8(FunctionSetTouchGestureCallbackConfiguration), buf.Bytes())
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
// SetTouchGestureCallbackConfiguration.
func (device *LCD128x64Bricklet) GetTouchGestureCallbackConfiguration() (period uint32, valueHasToChange bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetTouchGestureCallbackConfiguration), buf.Bytes())
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

// Draws a white or black line from (x, y)-start to (x, y)-end.
// 
// .. versionadded:: 2.0.2$nbsp;(Plugin)
//
// Associated constants:
//
//	* ColorWhite
//	* ColorBlack
func (device *LCD128x64Bricklet) DrawLine(positionXStart uint8, positionYStart uint8, positionXEnd uint8, positionYEnd uint8, color Color) (err error) {
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

// Draws a white or black box from (x, y)-start to (x, y)-end.
// 
// If you set fill to true, the box will be filled with the
// color. Otherwise only the outline will be drawn.
// 
// .. versionadded:: 2.0.2$nbsp;(Plugin)
//
// Associated constants:
//
//	* ColorWhite
//	* ColorBlack
func (device *LCD128x64Bricklet) DrawBox(positionXStart uint8, positionYStart uint8, positionXEnd uint8, positionYEnd uint8, fill bool, color Color) (err error) {
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

// Draws a text at the pixel position (x, y).
// 
// You can use one of 9 different font sizes and draw the text in white or black.
// 
// .. versionadded:: 2.0.2$nbsp;(Plugin)
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
//	* ColorWhite
//	* ColorBlack
func (device *LCD128x64Bricklet) DrawText(positionX uint8, positionY uint8, font Font, color Color, text string) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, positionX);
	binary.Write(&buf, binary.LittleEndian, positionY);
	binary.Write(&buf, binary.LittleEndian, font);
	binary.Write(&buf, binary.LittleEndian, color);
	text_byte_slice, err := StringToByteSlice(text, 22)
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

// Draws a clickable button at position (x, y) with the given text.
// 
// You can use up to 12 buttons.
// 
// The x position + width has to be within the range of 1 to 128 and the y
// position + height has to be within the range of 1 to 64.
// 
// The minimum useful width/height of a button is 3.
// 
// You can enable a callback for a button press with
// SetGUIButtonPressedCallbackConfiguration. The callback will
// be triggered for press and release-events.
// 
// The button is drawn in a separate GUI buffer and the button-frame will
// always stay on top of the graphics drawn with WritePixels. To
// remove the button use RemoveGUIButton.
// 
// If you want an icon instead of text, you can draw the icon inside of the
// button with WritePixels.
// 
// .. versionadded:: 2.0.2$nbsp;(Plugin)
func (device *LCD128x64Bricklet) SetGUIButton(index uint8, positionX uint8, positionY uint8, width uint8, height uint8, text string) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, index);
	binary.Write(&buf, binary.LittleEndian, positionX);
	binary.Write(&buf, binary.LittleEndian, positionY);
	binary.Write(&buf, binary.LittleEndian, width);
	binary.Write(&buf, binary.LittleEndian, height);
	text_byte_slice, err := StringToByteSlice(text, 16)
	if err != nil { return }
	buf.Write(text_byte_slice)

	resultBytes, err := device.device.Set(uint8(FunctionSetGUIButton), buf.Bytes())
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

// Returns the button properties for a given `Index` as set by SetGUIButton.
// 
// Additionally the `Active` parameter shows if a button is currently active/visible
// or not.
// 
// .. versionadded:: 2.0.2$nbsp;(Plugin)
func (device *LCD128x64Bricklet) GetGUIButton(index uint8) (active bool, positionX uint8, positionY uint8, width uint8, height uint8, text string, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, index);

	resultBytes, err := device.device.Get(uint8(FunctionGetGUIButton), buf.Bytes())
	if err != nil {
		return active, positionX, positionY, width, height, text, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 29 {
			return active, positionX, positionY, width, height, text, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 29)
		}

		if header.ErrorCode != 0 {
			return active, positionX, positionY, width, height, text, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &active)
		binary.Read(resultBuf, binary.LittleEndian, &positionX)
		binary.Read(resultBuf, binary.LittleEndian, &positionY)
		binary.Read(resultBuf, binary.LittleEndian, &width)
		binary.Read(resultBuf, binary.LittleEndian, &height)
		text = ByteSliceToString(resultBuf.Next(16))

	}

	return active, positionX, positionY, width, height, text, nil
}

// Removes the button with the given index.
// 
// You can use index 255 to remove all buttons.
// 
// .. versionadded:: 2.0.2$nbsp;(Plugin)
func (device *LCD128x64Bricklet) RemoveGUIButton(index uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, index);

	resultBytes, err := device.device.Set(uint8(FunctionRemoveGUIButton), buf.Bytes())
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

// The period is the period with which the RegisterGUIButtonPressedCallback callback
// is triggered periodically. A value of 0 turns the callback off.
// 
// If the `value has to change`-parameter is set to true, the callback is only
// triggered after the value has changed. If the value didn't change within the
// period, the callback is triggered immediately on change.
// 
// If it is set to false, the callback is continuously triggered with the period,
// independent of the value.
// 
// .. versionadded:: 2.0.2$nbsp;(Plugin)
func (device *LCD128x64Bricklet) SetGUIButtonPressedCallbackConfiguration(period uint32, valueHasToChange bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);
	binary.Write(&buf, binary.LittleEndian, valueHasToChange);

	resultBytes, err := device.device.Set(uint8(FunctionSetGUIButtonPressedCallbackConfiguration), buf.Bytes())
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
// SetGUIButtonPressedCallbackConfiguration.
// 
// .. versionadded:: 2.0.2$nbsp;(Plugin)
func (device *LCD128x64Bricklet) GetGUIButtonPressedCallbackConfiguration() (period uint32, valueHasToChange bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetGUIButtonPressedCallbackConfiguration), buf.Bytes())
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

// Returns the state of the button for the given index.
// 
// The state can either be pressed (true) or released (false).
// 
// .. versionadded:: 2.0.2$nbsp;(Plugin)
func (device *LCD128x64Bricklet) GetGUIButtonPressed(index uint8) (pressed bool, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, index);

	resultBytes, err := device.device.Get(uint8(FunctionGetGUIButtonPressed), buf.Bytes())
	if err != nil {
		return pressed, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return pressed, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return pressed, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &pressed)

	}

	return pressed, nil
}

// Draws a slider at position (x, y) with the given length.
// 
// You can use up to 6 sliders.
// 
// If you use the horizontal direction, the x position + length has to be
// within the range of 1 to 128 and the y position has to be within
// the range of 0 to 46.
// 
// If you use the vertical direction, the y position + length has to be
// within the range of 1 to 64 and the x position has to be within
// the range of 0 to 110.
// 
// The minimum length of a slider is 8.
// 
// The parameter value is the start-position of the slider, it can
// be between 0 and length-8.
// 
// You can enable a callback for the slider value with
// SetGUISliderValueCallbackConfiguration.
// 
// The slider is drawn in a separate GUI buffer and it will
// always stay on top of the graphics drawn with WritePixels. To
// remove the button use RemoveGUISlider.
// 
// .. versionadded:: 2.0.2$nbsp;(Plugin)
//
// Associated constants:
//
//	* DirectionHorizontal
//	* DirectionVertical
func (device *LCD128x64Bricklet) SetGUISlider(index uint8, positionX uint8, positionY uint8, length uint8, direction Direction, value uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, index);
	binary.Write(&buf, binary.LittleEndian, positionX);
	binary.Write(&buf, binary.LittleEndian, positionY);
	binary.Write(&buf, binary.LittleEndian, length);
	binary.Write(&buf, binary.LittleEndian, direction);
	binary.Write(&buf, binary.LittleEndian, value);

	resultBytes, err := device.device.Set(uint8(FunctionSetGUISlider), buf.Bytes())
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

// Returns the slider properties for a given `Index` as set by SetGUISlider.
// 
// Additionally the `Active` parameter shows if a button is currently active/visible
// or not.
// 
// .. versionadded:: 2.0.2$nbsp;(Plugin)
//
// Associated constants:
//
//	* DirectionHorizontal
//	* DirectionVertical
func (device *LCD128x64Bricklet) GetGUISlider(index uint8) (active bool, positionX uint8, positionY uint8, length uint8, direction Direction, value uint8, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, index);

	resultBytes, err := device.device.Get(uint8(FunctionGetGUISlider), buf.Bytes())
	if err != nil {
		return active, positionX, positionY, length, direction, value, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 14 {
			return active, positionX, positionY, length, direction, value, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 14)
		}

		if header.ErrorCode != 0 {
			return active, positionX, positionY, length, direction, value, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &active)
		binary.Read(resultBuf, binary.LittleEndian, &positionX)
		binary.Read(resultBuf, binary.LittleEndian, &positionY)
		binary.Read(resultBuf, binary.LittleEndian, &length)
		binary.Read(resultBuf, binary.LittleEndian, &direction)
		binary.Read(resultBuf, binary.LittleEndian, &value)

	}

	return active, positionX, positionY, length, direction, value, nil
}

// Removes the slider with the given index.
// 
// You can use index 255 to remove all slider.
// 
// .. versionadded:: 2.0.2$nbsp;(Plugin)
func (device *LCD128x64Bricklet) RemoveGUISlider(index uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, index);

	resultBytes, err := device.device.Set(uint8(FunctionRemoveGUISlider), buf.Bytes())
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

// The period is the period with which the RegisterGUISliderValueCallback callback
// is triggered periodically. A value of 0 turns the callback off.
// 
// If the `value has to change`-parameter is set to true, the callback is only
// triggered after the value has changed. If the value didn't change within the
// period, the callback is triggered immediately on change.
// 
// If it is set to false, the callback is continuously triggered with the period,
// independent of the value.
// 
// .. versionadded:: 2.0.2$nbsp;(Plugin)
func (device *LCD128x64Bricklet) SetGUISliderValueCallbackConfiguration(period uint32, valueHasToChange bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);
	binary.Write(&buf, binary.LittleEndian, valueHasToChange);

	resultBytes, err := device.device.Set(uint8(FunctionSetGUISliderValueCallbackConfiguration), buf.Bytes())
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
// SetGUISliderValueCallbackConfiguration.
// 
// .. versionadded:: 2.0.2$nbsp;(Plugin)
func (device *LCD128x64Bricklet) GetGUISliderValueCallbackConfiguration() (period uint32, valueHasToChange bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetGUISliderValueCallbackConfiguration), buf.Bytes())
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

// Returns the current slider value for the given index.
// 
// .. versionadded:: 2.0.2$nbsp;(Plugin)
func (device *LCD128x64Bricklet) GetGUISliderValue(index uint8) (value uint8, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, index);

	resultBytes, err := device.device.Get(uint8(FunctionGetGUISliderValue), buf.Bytes())
	if err != nil {
		return value, err
	}
	if len(resultBytes) > 0 {
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

	}

	return value, nil
}

// Sets the general configuration for tabs. You can configure the tabs to only
// accept clicks or only swipes (gesture left/right and right/left) or both.
// 
// Additionally, if you set `Clear GUI` to true, all of the GUI elements (buttons,
// slider, graphs) will automatically be removed on every tab change.
// 
// .. versionadded:: 2.0.2$nbsp;(Plugin)
//
// Associated constants:
//
//	* ChangeTabOnClick
//	* ChangeTabOnSwipe
//	* ChangeTabOnClickAndSwipe
func (device *LCD128x64Bricklet) SetGUITabConfiguration(changeTabConfig ChangeTabOn, clearGUI bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, changeTabConfig);
	binary.Write(&buf, binary.LittleEndian, clearGUI);

	resultBytes, err := device.device.Set(uint8(FunctionSetGUITabConfiguration), buf.Bytes())
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

// Returns the tab configuration as set by SetGUITabConfiguration.
// 
// .. versionadded:: 2.0.2$nbsp;(Plugin)
//
// Associated constants:
//
//	* ChangeTabOnClick
//	* ChangeTabOnSwipe
//	* ChangeTabOnClickAndSwipe
func (device *LCD128x64Bricklet) GetGUITabConfiguration() (changeTabConfig ChangeTabOn, clearGUI bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetGUITabConfiguration), buf.Bytes())
	if err != nil {
		return changeTabConfig, clearGUI, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 10 {
			return changeTabConfig, clearGUI, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		if header.ErrorCode != 0 {
			return changeTabConfig, clearGUI, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &changeTabConfig)
		binary.Read(resultBuf, binary.LittleEndian, &clearGUI)

	}

	return changeTabConfig, clearGUI, nil
}

// Adds a text-tab with the given index.
// 
// You can use up to 10 tabs.
// 
// A text-tab with the same index as a icon-tab will overwrite the icon-tab.
// 
// .. versionadded:: 2.0.2$nbsp;(Plugin)
func (device *LCD128x64Bricklet) SetGUITabText(index uint8, text string) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, index);
	text_byte_slice, err := StringToByteSlice(text, 5)
	if err != nil { return }
	buf.Write(text_byte_slice)

	resultBytes, err := device.device.Set(uint8(FunctionSetGUITabText), buf.Bytes())
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

// Returns the text for a given index as set by SetGUITabText.
// 
// Additionally the `Active` parameter shows if the tab is currently active/visible
// or not.
// 
// .. versionadded:: 2.0.2$nbsp;(Plugin)
func (device *LCD128x64Bricklet) GetGUITabText(index uint8) (active bool, text string, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, index);

	resultBytes, err := device.device.Get(uint8(FunctionGetGUITabText), buf.Bytes())
	if err != nil {
		return active, text, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 14 {
			return active, text, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 14)
		}

		if header.ErrorCode != 0 {
			return active, text, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &active)
		text = ByteSliceToString(resultBuf.Next(5))

	}

	return active, text, nil
}

// Adds a icon-tab with the given index. The icon can have a width of 28 pixels
// with a height of 6 pixels. It is drawn line-by-line from left to right.
// 
// You can use up to 10 tabs.
// 
// A icon-tab with the same index as a text-tab will overwrite the text-tab.
// 
// .. versionadded:: 2.0.2$nbsp;(Plugin)
func (device *LCD128x64Bricklet) SetGUITabIcon(index uint8, icon [168]bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, index);
	binary.Write(&buf, binary.LittleEndian, icon);

	resultBytes, err := device.device.Set(uint8(FunctionSetGUITabIcon), buf.Bytes())
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

// Returns the icon for a given index as set by SetGUITabIcon.
// 
// Additionally the `Active` parameter shows if the tab is currently active/visible
// or not.
// 
// .. versionadded:: 2.0.2$nbsp;(Plugin)
func (device *LCD128x64Bricklet) GetGUITabIcon(index uint8) (active bool, icon [168]bool, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, index);

	resultBytes, err := device.device.Get(uint8(FunctionGetGUITabIcon), buf.Bytes())
	if err != nil {
		return active, icon, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 30 {
			return active, icon, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 30)
		}

		if header.ErrorCode != 0 {
			return active, icon, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &active)
		binary.Read(resultBuf, binary.LittleEndian, &icon)

	}

	return active, icon, nil
}

// Removes the tab with the given index.
// 
// You can use index 255 to remove all tabs.
// 
// .. versionadded:: 2.0.2$nbsp;(Plugin)
func (device *LCD128x64Bricklet) RemoveGUITab(index uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, index);

	resultBytes, err := device.device.Set(uint8(FunctionRemoveGUITab), buf.Bytes())
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

// Sets the tab with the given index as selected (drawn as selected on the display).
// 
// .. versionadded:: 2.0.2$nbsp;(Plugin)
func (device *LCD128x64Bricklet) SetGUITabSelected(index uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, index);

	resultBytes, err := device.device.Set(uint8(FunctionSetGUITabSelected), buf.Bytes())
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

// The period is the period with which the RegisterGUITabSelectedCallback callback
// is triggered periodically. A value of 0 turns the callback off.
// 
// If the `value has to change`-parameter is set to true, the callback is only
// triggered after the value has changed. If the value didn't change within the
// period, the callback is triggered immediately on change.
// 
// If it is set to false, the callback is continuously triggered with the period,
// independent of the value.
// 
// .. versionadded:: 2.0.2$nbsp;(Plugin)
func (device *LCD128x64Bricklet) SetGUITabSelectedCallbackConfiguration(period uint32, valueHasToChange bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);
	binary.Write(&buf, binary.LittleEndian, valueHasToChange);

	resultBytes, err := device.device.Set(uint8(FunctionSetGUITabSelectedCallbackConfiguration), buf.Bytes())
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
// SetGUITabSelectedCallbackConfiguration.
// 
// .. versionadded:: 2.0.2$nbsp;(Plugin)
func (device *LCD128x64Bricklet) GetGUITabSelectedCallbackConfiguration() (period uint32, valueHasToChange bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetGUITabSelectedCallbackConfiguration), buf.Bytes())
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

// Returns the index of the currently selected tab.
// If there are not tabs, the returned index is -1.
// 
// .. versionadded:: 2.0.2$nbsp;(Plugin)
func (device *LCD128x64Bricklet) GetGUITabSelected() (index int8, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetGUITabSelected), buf.Bytes())
	if err != nil {
		return index, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return index, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return index, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &index)

	}

	return index, nil
}

// Sets the configuration for up to four graphs.
// 
// The graph type can be dot-, line- or bar-graph.
// 
// The x and y position are pixel positions.
// 
// You can add a text for the x and y axis.
// The text is drawn at the inside of the graph and it can overwrite some
// of the graph data. If you need the text outside of the graph you can
// leave this text here empty and use DrawText to draw the caption
// outside of the graph.
// 
// The data of the graph can be set and updated with SetGUIGraphData.
// 
// The graph is drawn in a separate GUI buffer and the graph-frame and data will
// always stay on top of the graphics drawn with WritePixels. To
// remove the graph use RemoveGUIGraph.
// 
// .. versionadded:: 2.0.2$nbsp;(Plugin)
//
// Associated constants:
//
//	* GraphTypeDot
//	* GraphTypeLine
//	* GraphTypeBar
func (device *LCD128x64Bricklet) SetGUIGraphConfiguration(index uint8, graphType GraphType, positionX uint8, positionY uint8, width uint8, height uint8, textX string, textY string) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, index);
	binary.Write(&buf, binary.LittleEndian, graphType);
	binary.Write(&buf, binary.LittleEndian, positionX);
	binary.Write(&buf, binary.LittleEndian, positionY);
	binary.Write(&buf, binary.LittleEndian, width);
	binary.Write(&buf, binary.LittleEndian, height);
	textX_byte_slice, err := StringToByteSlice(textX, 4)
	if err != nil { return }
	buf.Write(textX_byte_slice)
	textY_byte_slice, err := StringToByteSlice(textY, 4)
	if err != nil { return }
	buf.Write(textY_byte_slice)

	resultBytes, err := device.device.Set(uint8(FunctionSetGUIGraphConfiguration), buf.Bytes())
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

// Returns the graph properties for a given `Index` as set by SetGUIGraphConfiguration.
// 
// Additionally the `Active` parameter shows if a graph is currently active/visible
// or not.
// 
// .. versionadded:: 2.0.2$nbsp;(Plugin)
//
// Associated constants:
//
//	* GraphTypeDot
//	* GraphTypeLine
//	* GraphTypeBar
func (device *LCD128x64Bricklet) GetGUIGraphConfiguration(index uint8) (active bool, graphType GraphType, positionX uint8, positionY uint8, width uint8, height uint8, textX string, textY string, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, index);

	resultBytes, err := device.device.Get(uint8(FunctionGetGUIGraphConfiguration), buf.Bytes())
	if err != nil {
		return active, graphType, positionX, positionY, width, height, textX, textY, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 22 {
			return active, graphType, positionX, positionY, width, height, textX, textY, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 22)
		}

		if header.ErrorCode != 0 {
			return active, graphType, positionX, positionY, width, height, textX, textY, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &active)
		binary.Read(resultBuf, binary.LittleEndian, &graphType)
		binary.Read(resultBuf, binary.LittleEndian, &positionX)
		binary.Read(resultBuf, binary.LittleEndian, &positionY)
		binary.Read(resultBuf, binary.LittleEndian, &width)
		binary.Read(resultBuf, binary.LittleEndian, &height)
		textX = ByteSliceToString(resultBuf.Next(4))
		textY = ByteSliceToString(resultBuf.Next(4))

	}

	return active, graphType, positionX, positionY, width, height, textX, textY, nil
}

// Sets the data for a graph with the given index. You have to configure the graph with
// SetGUIGraphConfiguration before you can set the first data.
// 
// The graph will show the first n values of the data that you set, where
// n is the width set with SetGUIGraphConfiguration. If you set
// less then n values it will show the rest of the values as zero.
// 
// The maximum number of data-points you can set is 118 (which also corresponds to the
// maximum width of the graph).
// 
// You have to scale your values to be between 0 and 255. 0 will be shown
// at the bottom of the graph and 255 at the top.
// 
// .. versionadded:: 2.0.2$nbsp;(Plugin)
func (device *LCD128x64Bricklet) SetGUIGraphDataLowLevel(index uint8, dataLength uint16, dataChunkOffset uint16, dataChunkData [59]uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, index);
	binary.Write(&buf, binary.LittleEndian, dataLength);
	binary.Write(&buf, binary.LittleEndian, dataChunkOffset);
	buf.Write(Uint8SliceToByteSlice(dataChunkData[:]))

	resultBytes, err := device.device.Set(uint8(FunctionSetGUIGraphDataLowLevel), buf.Bytes())
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

// Sets the data for a graph with the given index. You have to configure the graph with
// SetGUIGraphConfiguration before you can set the first data.
// 
// The graph will show the first n values of the data that you set, where
// n is the width set with SetGUIGraphConfiguration. If you set
// less then n values it will show the rest of the values as zero.
// 
// The maximum number of data-points you can set is 118 (which also corresponds to the
// maximum width of the graph).
// 
// You have to scale your values to be between 0 and 255. 0 will be shown
// at the bottom of the graph and 255 at the top.
// 
// .. versionadded:: 2.0.2$nbsp;(Plugin)
	func (device *LCD128x64Bricklet) SetGUIGraphData(index uint8, data []uint8) (err error) {
		_, err = device.device.SetHighLevel(func(dataLength uint64, dataChunkOffset uint64, dataChunkData []byte) (LowLevelWriteResult, error) {
			arr := [59]uint8{}
			copy(arr[:], ByteSliceToUint8Slice(dataChunkData))

			err := device.SetGUIGraphDataLowLevel(index, uint16(dataLength), uint16(dataChunkOffset), arr)

			var lowLevelResults bytes.Buffer
			

			return LowLevelWriteResult{
				uint64(59),
				lowLevelResults.Bytes()}, err
		}, 2, 8, 472, Uint8SliceToByteSlice(data))

		if err != nil {
			return
		}

		
		
		
		return
	}

// Returns the graph data for a given index as set by SetGUIGraphData.
// 
// .. versionadded:: 2.0.2$nbsp;(Plugin)
func (device *LCD128x64Bricklet) GetGUIGraphDataLowLevel(index uint8) (dataLength uint16, dataChunkOffset uint16, dataChunkData [59]uint8, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, index);

	resultBytes, err := device.device.Get(uint8(FunctionGetGUIGraphDataLowLevel), buf.Bytes())
	if err != nil {
		return dataLength, dataChunkOffset, dataChunkData, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 71 {
			return dataLength, dataChunkOffset, dataChunkData, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 71)
		}

		if header.ErrorCode != 0 {
			return dataLength, dataChunkOffset, dataChunkData, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &dataLength)
		binary.Read(resultBuf, binary.LittleEndian, &dataChunkOffset)
		copy(dataChunkData[:], ByteSliceToUint8Slice(resultBuf.Next(8 * 59/8)))

	}

	return dataLength, dataChunkOffset, dataChunkData, nil
}

// Returns the graph data for a given index as set by SetGUIGraphData.
// 
// .. versionadded:: 2.0.2$nbsp;(Plugin)
	func (device *LCD128x64Bricklet) GetGUIGraphData(index uint8) (data []uint8, err error) {
		buf, _, err := device.device.GetHighLevel(func() (LowLevelResult, error) {
			dataLength, dataChunkOffset, dataChunkData, err := device.GetGUIGraphDataLowLevel(index)

			if err != nil {
				return LowLevelResult{}, err
			}

			var lowLevelResults bytes.Buffer
			

			return LowLevelResult{
				uint64(dataLength),
				uint64(dataChunkOffset),
				Uint8SliceToByteSlice(dataChunkData[:]),
				lowLevelResults.Bytes()}, nil
		},
			3,
			8)
		if err != nil {
			return ByteSliceToUint8Slice(buf), err
		}
		
		
		return ByteSliceToUint8Slice(buf), nil
	}

// Removes the graph with the given index.
// 
// You can use index 255 to remove all graphs.
// 
// .. versionadded:: 2.0.2$nbsp;(Plugin)
func (device *LCD128x64Bricklet) RemoveGUIGraph(index uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, index);

	resultBytes, err := device.device.Set(uint8(FunctionRemoveGUIGraph), buf.Bytes())
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

// Removes all GUI elements (buttons, slider, graphs, tabs).
// 
// .. versionadded:: 2.0.2$nbsp;(Plugin)
func (device *LCD128x64Bricklet) RemoveAllGUI() (err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Set(uint8(FunctionRemoveAllGUI), buf.Bytes())
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

// Sets the touch LED configuration. By default the LED is on if the
// LCD is touched.
// 
// You can also turn the LED permanently on/off or show a heartbeat.
// 
// If the Bricklet is in bootloader mode, the LED is off.
// 
// .. versionadded:: 2.0.2$nbsp;(Plugin)
//
// Associated constants:
//
//	* TouchLEDConfigOff
//	* TouchLEDConfigOn
//	* TouchLEDConfigShowHeartbeat
//	* TouchLEDConfigShowTouch
func (device *LCD128x64Bricklet) SetTouchLEDConfig(config TouchLEDConfig) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, config);

	resultBytes, err := device.device.Set(uint8(FunctionSetTouchLEDConfig), buf.Bytes())
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

// Returns the configuration as set by SetTouchLEDConfig
// 
// .. versionadded:: 2.0.2$nbsp;(Plugin)
//
// Associated constants:
//
//	* TouchLEDConfigOff
//	* TouchLEDConfigOn
//	* TouchLEDConfigShowHeartbeat
//	* TouchLEDConfigShowTouch
func (device *LCD128x64Bricklet) GetTouchLEDConfig() (config TouchLEDConfig, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetTouchLEDConfig), buf.Bytes())
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
func (device *LCD128x64Bricklet) GetSPITFPErrorCount() (errorCountAckChecksum uint32, errorCountMessageChecksum uint32, errorCountFrame uint32, errorCountOverflow uint32, err error) {
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
func (device *LCD128x64Bricklet) SetBootloaderMode(mode BootloaderMode) (status BootloaderStatus, err error) {
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
func (device *LCD128x64Bricklet) GetBootloaderMode() (mode BootloaderMode, err error) {
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
func (device *LCD128x64Bricklet) SetWriteFirmwarePointer(pointer uint32) (err error) {
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
func (device *LCD128x64Bricklet) WriteFirmware(data [64]uint8) (status uint8, err error) {
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
func (device *LCD128x64Bricklet) SetStatusLEDConfig(config StatusLEDConfig) (err error) {
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
func (device *LCD128x64Bricklet) GetStatusLEDConfig() (config StatusLEDConfig, err error) {
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
func (device *LCD128x64Bricklet) GetChipTemperature() (temperature int16, err error) {
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
func (device *LCD128x64Bricklet) Reset() (err error) {
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
func (device *LCD128x64Bricklet) WriteUID(uid uint32) (err error) {
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
func (device *LCD128x64Bricklet) ReadUID() (uid uint32, err error) {
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
// The Raspberry Pi HAT (Zero) Brick is always at position 'i' and the Bricklet
// connected to an `Isolator Bricklet <isolator_bricklet>` is always as
// position 'z'.
// 
// The device identifier numbers can be found `here <device_identifier>`.
// |device_identifier_constant|
func (device *LCD128x64Bricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
