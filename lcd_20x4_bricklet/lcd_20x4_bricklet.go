/* ***********************************************************
 * This file was automatically generated on 2020-04-20.      *
 *                                                           *
 * Go Bindings Version 2.0.7                                 *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


// 20x4 character alphanumeric display with blue backlight.
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/LCD20x4_Bricklet_Go.html.
package lcd_20x4_bricklet

import (
	"encoding/binary"
	"bytes"
	"fmt"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionWriteLine Function = 1
	FunctionClearDisplay Function = 2
	FunctionBacklightOn Function = 3
	FunctionBacklightOff Function = 4
	FunctionIsBacklightOn Function = 5
	FunctionSetConfig Function = 6
	FunctionGetConfig Function = 7
	FunctionIsButtonPressed Function = 8
	FunctionSetCustomCharacter Function = 11
	FunctionGetCustomCharacter Function = 12
	FunctionSetDefaultText Function = 13
	FunctionGetDefaultText Function = 14
	FunctionSetDefaultTextCounter Function = 15
	FunctionGetDefaultTextCounter Function = 16
	FunctionGetIdentity Function = 255
	FunctionCallbackButtonPressed Function = 9
	FunctionCallbackButtonReleased Function = 10
)

type LCD20x4Bricklet struct {
	device Device
}
const DeviceIdentifier = 212
const DeviceDisplayName = "LCD 20x4 Bricklet"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (LCD20x4Bricklet, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{ 2,0,2 }, uid, &internalIPCon, 0, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return LCD20x4Bricklet{}, err
	}
	dev.ResponseExpected[FunctionWriteLine] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionClearDisplay] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionBacklightOn] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionBacklightOff] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionIsBacklightOn] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetConfig] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetConfig] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionIsButtonPressed] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetCustomCharacter] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetCustomCharacter] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetDefaultText] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetDefaultText] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetDefaultTextCounter] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetDefaultTextCounter] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetIdentity] = ResponseExpectedFlagAlwaysTrue;
	return LCD20x4Bricklet{dev}, nil
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
func (device *LCD20x4Bricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *LCD20x4Bricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *LCD20x4Bricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *LCD20x4Bricklet) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// This callback is triggered when a button is pressed. The parameter is
// the number of the button (0 to 2 or 0 to 3 since hardware version 1.2).
func (device *LCD20x4Bricklet) RegisterButtonPressedCallback(fn func(uint8)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 9 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var button uint8
		binary.Read(buf, binary.LittleEndian, &button)
		fn(button)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackButtonPressed), wrapper)
}

// Remove a registered Button Pressed callback.
func (device *LCD20x4Bricklet) DeregisterButtonPressedCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackButtonPressed), registrationId)
}


// This callback is triggered when a button is released. The parameter is
// the number of the button (0 to 2 or 0 to 3 since hardware version 1.2).
func (device *LCD20x4Bricklet) RegisterButtonReleasedCallback(fn func(uint8)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 9 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var button uint8
		binary.Read(buf, binary.LittleEndian, &button)
		fn(button)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackButtonReleased), wrapper)
}

// Remove a registered Button Released callback.
func (device *LCD20x4Bricklet) DeregisterButtonReleasedCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackButtonReleased), registrationId)
}


// Writes text to a specific line with a specific position.
// The text can have a maximum of 20 characters.
// 
// For example: (0, 7, Hello) will write *Hello* in the middle of the
// first line of the display.
// 
// The display uses a special charset that includes all ASCII characters except
// backslash and tilde. The LCD charset also includes several other non-ASCII characters, see
// the https://github.com/Tinkerforge/lcd-20x4-bricklet/raw/master/datasheets/standard_charset.pdf
// for details. The Unicode example above shows how to specify non-ASCII characters
// and how to translate from Unicode to the LCD charset.
func (device *LCD20x4Bricklet) WriteLine(line uint8, position uint8, text string) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, line);
	binary.Write(&buf, binary.LittleEndian, position);
	text_byte_slice, err := StringToByteSlice(text, 20)
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

// Deletes all characters from the display.
func (device *LCD20x4Bricklet) ClearDisplay() (err error) {
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

// Turns the backlight on.
func (device *LCD20x4Bricklet) BacklightOn() (err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Set(uint8(FunctionBacklightOn), buf.Bytes())
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

// Turns the backlight off.
func (device *LCD20x4Bricklet) BacklightOff() (err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Set(uint8(FunctionBacklightOff), buf.Bytes())
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

// Returns *true* if the backlight is on and *false* otherwise.
func (device *LCD20x4Bricklet) IsBacklightOn() (backlight bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionIsBacklightOn), buf.Bytes())
	if err != nil {
		return backlight, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return backlight, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return backlight, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &backlight)

	}

	return backlight, nil
}

// Configures if the cursor (shown as _) should be visible and if it
// should be blinking (shown as a blinking block). The cursor position
// is one character behind the the last text written with
// WriteLine.
func (device *LCD20x4Bricklet) SetConfig(cursor bool, blinking bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, cursor);
	binary.Write(&buf, binary.LittleEndian, blinking);

	resultBytes, err := device.device.Set(uint8(FunctionSetConfig), buf.Bytes())
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

// Returns the configuration as set by SetConfig.
func (device *LCD20x4Bricklet) GetConfig() (cursor bool, blinking bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetConfig), buf.Bytes())
	if err != nil {
		return cursor, blinking, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 10 {
			return cursor, blinking, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		if header.ErrorCode != 0 {
			return cursor, blinking, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &cursor)
		binary.Read(resultBuf, binary.LittleEndian, &blinking)

	}

	return cursor, blinking, nil
}

// Returns *true* if the button (0 to 2 or 0 to 3 since hardware version 1.2)
// is pressed.
// 
// If you want to react on button presses and releases it is recommended to use
// the RegisterButtonPressedCallback and RegisterButtonReleasedCallback callbacks.
func (device *LCD20x4Bricklet) IsButtonPressed(button uint8) (pressed bool, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, button);

	resultBytes, err := device.device.Get(uint8(FunctionIsButtonPressed), buf.Bytes())
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

// The LCD 20x4 Bricklet can store up to 8 custom characters. The characters
// consist of 5x8 pixels and can be addressed with the index 0-7. To describe
// the pixels, the first 5 bits of 8 bytes are used. For example, to make
// a custom character H, you should transfer the following:
// 
// * ``character[0] = 0b00010001`` (decimal value 17)
// * ``character[1] = 0b00010001`` (decimal value 17)
// * ``character[2] = 0b00010001`` (decimal value 17)
// * ``character[3] = 0b00011111`` (decimal value 31)
// * ``character[4] = 0b00010001`` (decimal value 17)
// * ``character[5] = 0b00010001`` (decimal value 17)
// * ``character[6] = 0b00010001`` (decimal value 17)
// * ``character[7] = 0b00000000`` (decimal value 0)
// 
// The characters can later be written with WriteLine by using the
// characters with the byte representation 8 (\\x08 or \\u0008) to 15
// (\\x0F or \\u000F).
// 
// You can play around with the custom characters in Brick Viewer version
// since 2.0.1.
// 
// Custom characters are stored by the LCD in RAM, so they have to be set
// after each startup.
// 
// .. versionadded:: 2.0.1$nbsp;(Plugin)
func (device *LCD20x4Bricklet) SetCustomCharacter(index uint8, character [8]uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, index);
	binary.Write(&buf, binary.LittleEndian, character);

	resultBytes, err := device.device.Set(uint8(FunctionSetCustomCharacter), buf.Bytes())
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

// Returns the custom character for a given index, as set with
// SetCustomCharacter.
// 
// .. versionadded:: 2.0.1$nbsp;(Plugin)
func (device *LCD20x4Bricklet) GetCustomCharacter(index uint8) (character [8]uint8, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, index);

	resultBytes, err := device.device.Get(uint8(FunctionGetCustomCharacter), buf.Bytes())
	if err != nil {
		return character, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 16 {
			return character, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 16)
		}

		if header.ErrorCode != 0 {
			return character, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &character)

	}

	return character, nil
}

// Sets the default text for lines 0-3. The max number of characters
// per line is 20.
// 
// The default text is shown on the LCD, if the default text counter
// expires, see SetDefaultTextCounter.
// 
// .. versionadded:: 2.0.2$nbsp;(Plugin)
func (device *LCD20x4Bricklet) SetDefaultText(line uint8, text string) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, line);
	text_byte_slice, err := StringToByteSlice(text, 20)
	if err != nil { return }
	buf.Write(text_byte_slice)

	resultBytes, err := device.device.Set(uint8(FunctionSetDefaultText), buf.Bytes())
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

// Returns the default text for a given line (0-3) as set by
// SetDefaultText.
// 
// .. versionadded:: 2.0.2$nbsp;(Plugin)
func (device *LCD20x4Bricklet) GetDefaultText(line uint8) (text string, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, line);

	resultBytes, err := device.device.Get(uint8(FunctionGetDefaultText), buf.Bytes())
	if err != nil {
		return text, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 28 {
			return text, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 28)
		}

		if header.ErrorCode != 0 {
			return text, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		text = ByteSliceToString(resultBuf.Next(20))

	}

	return text, nil
}

// Sets the default text counter. This counter is decremented each
// ms by the LCD firmware. If the counter reaches 0, the default text
// (see SetDefaultText) is shown on the LCD.
// 
// This functionality can be used to show a default text if the controlling
// program crashes or the connection is interrupted.
// 
// A possible approach is to call SetDefaultTextCounter every
// minute with the parameter 1000*60*2 (2 minutes). In this case the
// default text will be shown no later than 2 minutes after the
// controlling program crashes.
// 
// A negative counter turns the default text functionality off.
// 
// .. versionadded:: 2.0.2$nbsp;(Plugin)
func (device *LCD20x4Bricklet) SetDefaultTextCounter(counter int32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, counter);

	resultBytes, err := device.device.Set(uint8(FunctionSetDefaultTextCounter), buf.Bytes())
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

// Returns the current value of the default text counter.
// 
// .. versionadded:: 2.0.2$nbsp;(Plugin)
func (device *LCD20x4Bricklet) GetDefaultTextCounter() (counter int32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetDefaultTextCounter), buf.Bytes())
	if err != nil {
		return counter, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 12 {
			return counter, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		if header.ErrorCode != 0 {
			return counter, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &counter)

	}

	return counter, nil
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
func (device *LCD20x4Bricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
