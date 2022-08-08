/* ***********************************************************
 * This file was automatically generated on 2022-08-08.      *
 *                                                           *
 * Go Bindings Version 2.0.13                                *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/

// 3.3cm (1.3") OLED display with 128x64 pixels.
//
//
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/OLED128x64_Bricklet_Go.html.
package oled_128x64_bricklet

import (
	"bytes"
	"encoding/binary"
	"fmt"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionWrite                   Function = 1
	FunctionNewWindow               Function = 2
	FunctionClearDisplay            Function = 3
	FunctionSetDisplayConfiguration Function = 4
	FunctionGetDisplayConfiguration Function = 5
	FunctionWriteLine               Function = 6
	FunctionGetIdentity             Function = 255
)

type OLED128x64Bricklet struct {
	device Device
}

const DeviceIdentifier = 263
const DeviceDisplayName = "OLED 128x64 Bricklet"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (OLED128x64Bricklet, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{2, 0, 0}, uid, &internalIPCon, 0, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return OLED128x64Bricklet{}, err
	}
	dev.ResponseExpected[FunctionWrite] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionNewWindow] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionClearDisplay] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionSetDisplayConfiguration] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetDisplayConfiguration] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionWriteLine] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetIdentity] = ResponseExpectedFlagAlwaysTrue
	return OLED128x64Bricklet{dev}, nil
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
func (device *OLED128x64Bricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *OLED128x64Bricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *OLED128x64Bricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *OLED128x64Bricklet) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// Appends 64 byte of data to the window as set by NewWindow.
//
// Each row has a height of 8 pixels which corresponds to one byte of data.
//
// Example: if you call NewWindow with column from 0 to 127 and row
// from 0 to 7 (the whole display) each call of Write (red arrow) will
// write half of a row.
//
// .. image:: /Images/Bricklets/bricklet_oled_128x64_display.png
//    :scale: 100 %
//    :alt: Display pixel order
//    :align: center
//    :target: ../../_images/Bricklets/bricklet_oled_128x64_display.png
//
// The LSB (D0) of each data byte is at the top and the MSB (D7) is at the
// bottom of the row.
//
// The next call of Write will write the second half of the row
// and the next two the second row and so on. To fill the whole display
// you need to call Write 16 times.
func (device *OLED128x64Bricklet) Write(data [64]uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, data)

	resultBytes, err := device.device.Set(uint8(FunctionWrite), buf.Bytes())
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

// Sets the window in which you can write with Write. One row
// has a height of 8 pixels.
func (device *OLED128x64Bricklet) NewWindow(columnFrom uint8, columnTo uint8, rowFrom uint8, rowTo uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, columnFrom)
	binary.Write(&buf, binary.LittleEndian, columnTo)
	binary.Write(&buf, binary.LittleEndian, rowFrom)
	binary.Write(&buf, binary.LittleEndian, rowTo)

	resultBytes, err := device.device.Set(uint8(FunctionNewWindow), buf.Bytes())
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

// Clears the current content of the window as set by NewWindow.
func (device *OLED128x64Bricklet) ClearDisplay() (err error) {
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
// You can set a contrast value from 0 to 255 and you can invert the color
// (black/white) of the display.
func (device *OLED128x64Bricklet) SetDisplayConfiguration(contrast uint8, invert bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, contrast)
	binary.Write(&buf, binary.LittleEndian, invert)

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
func (device *OLED128x64Bricklet) GetDisplayConfiguration() (contrast uint8, invert bool, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetDisplayConfiguration), buf.Bytes())
	if err != nil {
		return contrast, invert, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 10 {
			return contrast, invert, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		if header.ErrorCode != 0 {
			return contrast, invert, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &contrast)
		binary.Read(resultBuf, binary.LittleEndian, &invert)

	}

	return contrast, invert, nil
}

// Writes text to a specific line with a specific position.
// The text can have a maximum of 26 characters.
//
// For example: (1, 10, Hello) will write *Hello* in the middle of the
// second line of the display.
//
// You can draw to the display with Write and then add text to it
// afterwards.
//
// The display uses a special 5x7 pixel charset. You can view the characters
// of the charset in Brick Viewer.
//
// The font conforms to code page 437.
func (device *OLED128x64Bricklet) WriteLine(line uint8, position uint8, text string) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, line)
	binary.Write(&buf, binary.LittleEndian, position)
	text_byte_slice, err := StringToByteSlice(text, 26)
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
func (device *OLED128x64Bricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
