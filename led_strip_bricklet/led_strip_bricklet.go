/* ***********************************************************
 * This file was automatically generated on 2020-05-19.      *
 *                                                           *
 * Go Bindings Version 2.0.8                                 *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


// Controls up to 320 RGB LEDs.
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/LEDStrip_Bricklet_Go.html.
package led_strip_bricklet

import (
	"encoding/binary"
	"bytes"
	"fmt"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionSetRGBValues Function = 1
	FunctionGetRGBValues Function = 2
	FunctionSetFrameDuration Function = 3
	FunctionGetFrameDuration Function = 4
	FunctionGetSupplyVoltage Function = 5
	FunctionSetClockFrequency Function = 7
	FunctionGetClockFrequency Function = 8
	FunctionSetChipType Function = 9
	FunctionGetChipType Function = 10
	FunctionSetRGBWValues Function = 11
	FunctionGetRGBWValues Function = 12
	FunctionSetChannelMapping Function = 13
	FunctionGetChannelMapping Function = 14
	FunctionEnableFrameRenderedCallback Function = 15
	FunctionDisableFrameRenderedCallback Function = 16
	FunctionIsFrameRenderedCallbackEnabled Function = 17
	FunctionGetIdentity Function = 255
	FunctionCallbackFrameRendered Function = 6
)

type ChipType = uint16

const (
	ChipTypeWS2801 ChipType = 2801
	ChipTypeWS2811 ChipType = 2811
	ChipTypeWS2812 ChipType = 2812
	ChipTypeLPD8806 ChipType = 8806
	ChipTypeAPA102 ChipType = 102
)

type ChannelMapping = uint8

const (
	ChannelMappingRGB ChannelMapping = 6
	ChannelMappingRBG ChannelMapping = 9
	ChannelMappingBRG ChannelMapping = 33
	ChannelMappingBGR ChannelMapping = 36
	ChannelMappingGRB ChannelMapping = 18
	ChannelMappingGBR ChannelMapping = 24
	ChannelMappingRGBW ChannelMapping = 27
	ChannelMappingRGWB ChannelMapping = 30
	ChannelMappingRBGW ChannelMapping = 39
	ChannelMappingRBWG ChannelMapping = 45
	ChannelMappingRWGB ChannelMapping = 54
	ChannelMappingRWBG ChannelMapping = 57
	ChannelMappingGRWB ChannelMapping = 78
	ChannelMappingGRBW ChannelMapping = 75
	ChannelMappingGBWR ChannelMapping = 108
	ChannelMappingGBRW ChannelMapping = 99
	ChannelMappingGWBR ChannelMapping = 120
	ChannelMappingGWRB ChannelMapping = 114
	ChannelMappingBRGW ChannelMapping = 135
	ChannelMappingBRWG ChannelMapping = 141
	ChannelMappingBGRW ChannelMapping = 147
	ChannelMappingBGWR ChannelMapping = 156
	ChannelMappingBWRG ChannelMapping = 177
	ChannelMappingBWGR ChannelMapping = 180
	ChannelMappingWRBG ChannelMapping = 201
	ChannelMappingWRGB ChannelMapping = 198
	ChannelMappingWGBR ChannelMapping = 216
	ChannelMappingWGRB ChannelMapping = 210
	ChannelMappingWBGR ChannelMapping = 228
	ChannelMappingWBRG ChannelMapping = 225
)

type LEDStripBricklet struct {
	device Device
}
const DeviceIdentifier = 231
const DeviceDisplayName = "LED Strip Bricklet"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (LEDStripBricklet, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{ 2,0,3 }, uid, &internalIPCon, 0, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return LEDStripBricklet{}, err
	}
	dev.ResponseExpected[FunctionSetRGBValues] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetRGBValues] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetFrameDuration] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetFrameDuration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetSupplyVoltage] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetClockFrequency] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetClockFrequency] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetChipType] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetChipType] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetRGBWValues] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetRGBWValues] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetChannelMapping] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetChannelMapping] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionEnableFrameRenderedCallback] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionDisableFrameRenderedCallback] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionIsFrameRenderedCallbackEnabled] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetIdentity] = ResponseExpectedFlagAlwaysTrue;
	return LEDStripBricklet{dev}, nil
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
func (device *LEDStripBricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *LEDStripBricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *LEDStripBricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *LEDStripBricklet) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// This callback is triggered directly after a new frame is rendered. The
// parameter is the number of RGB or RGBW LEDs in that frame.
// 
// You should send the data for the next frame directly after this callback
// was triggered.
// 
// For an explanation of the general approach see SetRGBValues.
func (device *LEDStripBricklet) RegisterFrameRenderedCallback(fn func(uint16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 10 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var length uint16
		binary.Read(buf, binary.LittleEndian, &length)
		fn(length)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackFrameRendered), wrapper)
}

// Remove a registered Frame Rendered callback.
func (device *LEDStripBricklet) DeregisterFrameRenderedCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackFrameRendered), registrationId)
}


// Sets *length* RGB values for the LEDs starting from *index*.
// 
// To make the colors show correctly you need to configure the chip type
// (SetChipType) and a 3-channel channel mapping (SetChannelMapping)
// according to the connected LEDs.
// 
// Example: If you set
// 
// * index to 5,
// * length to 3,
// * r to [255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
// * g to [0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0] and
// * b to [0, 0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
// 
// the LED with index 5 will be red, 6 will be green and 7 will be blue.
// 
// Note Depending on the LED circuitry colors can be permuted.
// 
// The colors will be transfered to actual LEDs when the next
// frame duration ends, see SetFrameDuration.
// 
// Generic approach:
// 
// * Set the frame duration to a value that represents
//   the number of frames per second you want to achieve.
// * Set all of the LED colors for one frame.
// * Wait for the RegisterFrameRenderedCallback callback.
// * Set all of the LED colors for next frame.
// * Wait for the RegisterFrameRenderedCallback callback.
// * and so on.
// 
// This approach ensures that you can change the LED colors with
// a fixed frame rate.
// 
// The actual number of controllable LEDs depends on the number of free
// Bricklet ports. See `here <led_strip_bricklet_ram_constraints>` for more
// information. A call of SetRGBValues with index + length above the
// bounds is ignored completely.
func (device *LEDStripBricklet) SetRGBValues(index uint16, length uint8, r [16]uint8, g [16]uint8, b [16]uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, index);
	binary.Write(&buf, binary.LittleEndian, length);
	binary.Write(&buf, binary.LittleEndian, r);
	binary.Write(&buf, binary.LittleEndian, g);
	binary.Write(&buf, binary.LittleEndian, b);

	resultBytes, err := device.device.Set(uint8(FunctionSetRGBValues), buf.Bytes())
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

// Returns *length* R, G and B values starting from the
// given LED *index*.
// 
// The values are the last values that were set by SetRGBValues.
func (device *LEDStripBricklet) GetRGBValues(index uint16, length uint8) (r [16]uint8, g [16]uint8, b [16]uint8, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, index);
	binary.Write(&buf, binary.LittleEndian, length);

	resultBytes, err := device.device.Get(uint8(FunctionGetRGBValues), buf.Bytes())
	if err != nil {
		return r, g, b, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 56 {
			return r, g, b, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 56)
		}

		if header.ErrorCode != 0 {
			return r, g, b, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &r)
		binary.Read(resultBuf, binary.LittleEndian, &g)
		binary.Read(resultBuf, binary.LittleEndian, &b)

	}

	return r, g, b, nil
}

// Sets the frame duration.
// 
// Example: If you want to achieve 20 frames per second, you should
// set the frame duration to 50ms (50ms * 20 = 1 second).
// 
// For an explanation of the general approach see SetRGBValues.
func (device *LEDStripBricklet) SetFrameDuration(duration uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, duration);

	resultBytes, err := device.device.Set(uint8(FunctionSetFrameDuration), buf.Bytes())
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

// Returns the frame duration as set by SetFrameDuration.
func (device *LEDStripBricklet) GetFrameDuration() (duration uint16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetFrameDuration), buf.Bytes())
	if err != nil {
		return duration, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 10 {
			return duration, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		if header.ErrorCode != 0 {
			return duration, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &duration)

	}

	return duration, nil
}

// Returns the current supply voltage of the LEDs.
func (device *LEDStripBricklet) GetSupplyVoltage() (voltage uint16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetSupplyVoltage), buf.Bytes())
	if err != nil {
		return voltage, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 10 {
			return voltage, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		if header.ErrorCode != 0 {
			return voltage, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &voltage)

	}

	return voltage, nil
}

// Sets the frequency of the clock.
// 
// The Bricklet will choose the nearest achievable frequency, which may
// be off by a few Hz. You can get the exact frequency that is used by
// calling GetClockFrequency.
// 
// If you have problems with flickering LEDs, they may be bits flipping. You
// can fix this by either making the connection between the LEDs and the
// Bricklet shorter or by reducing the frequency.
// 
// With a decreasing frequency your maximum frames per second will decrease
// too.
// 
// Note
//  The frequency in firmware version 2.0.0 is fixed at 2MHz.
// 
// .. versionadded:: 2.0.1$nbsp;(Plugin)
func (device *LEDStripBricklet) SetClockFrequency(frequency uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, frequency);

	resultBytes, err := device.device.Set(uint8(FunctionSetClockFrequency), buf.Bytes())
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

// Returns the currently used clock frequency as set by SetClockFrequency.
// 
// .. versionadded:: 2.0.1$nbsp;(Plugin)
func (device *LEDStripBricklet) GetClockFrequency() (frequency uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetClockFrequency), buf.Bytes())
	if err != nil {
		return frequency, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 12 {
			return frequency, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		if header.ErrorCode != 0 {
			return frequency, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &frequency)

	}

	return frequency, nil
}

// Sets the type of the LED driver chip. We currently support the chips
// 
// * WS2801,
// * WS2811,
// * WS2812 / SK6812 / NeoPixel RGB,
// * SK6812RGBW / NeoPixel RGBW (Chip Type = WS2812),
// * LPD8806 and
// * APA102 / DotStar.
// 
// .. versionadded:: 2.0.2$nbsp;(Plugin)
//
// Associated constants:
//
//	* ChipTypeWS2801
//	* ChipTypeWS2811
//	* ChipTypeWS2812
//	* ChipTypeLPD8806
//	* ChipTypeAPA102
func (device *LEDStripBricklet) SetChipType(chip ChipType) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, chip);

	resultBytes, err := device.device.Set(uint8(FunctionSetChipType), buf.Bytes())
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

// Returns the currently used chip type as set by SetChipType.
// 
// .. versionadded:: 2.0.2$nbsp;(Plugin)
//
// Associated constants:
//
//	* ChipTypeWS2801
//	* ChipTypeWS2811
//	* ChipTypeWS2812
//	* ChipTypeLPD8806
//	* ChipTypeAPA102
func (device *LEDStripBricklet) GetChipType() (chip ChipType, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetChipType), buf.Bytes())
	if err != nil {
		return chip, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 10 {
			return chip, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		if header.ErrorCode != 0 {
			return chip, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &chip)

	}

	return chip, nil
}

// Sets *length* RGBW values for the LEDs starting from *index*.
// 
// To make the colors show correctly you need to configure the chip type
// (SetChipType) and a 4-channel channel mapping (SetChannelMapping)
// according to the connected LEDs.
// 
// The maximum length is 12, the index goes from 0 to 239 and the rgbw values
// have 8 bits each.
// 
// Example: If you set
// 
// * index to 5,
// * length to 4,
// * r to [255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
// * g to [0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
// * b to [0, 0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0] and
// * w to [0, 0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0]
// 
// the LED with index 5 will be red, 6 will be green, 7 will be blue and 8 will be white.
// 
// Note Depending on the LED circuitry colors can be permuted.
// 
// The colors will be transfered to actual LEDs when the next
// frame duration ends, see SetFrameDuration.
// 
// Generic approach:
// 
// * Set the frame duration to a value that represents
//   the number of frames per second you want to achieve.
// * Set all of the LED colors for one frame.
// * Wait for the RegisterFrameRenderedCallback callback.
// * Set all of the LED colors for next frame.
// * Wait for the RegisterFrameRenderedCallback callback.
// * and so on.
// 
// This approach ensures that you can change the LED colors with
// a fixed frame rate.
// 
// The actual number of controllable LEDs depends on the number of free
// Bricklet ports. See `here <led_strip_bricklet_ram_constraints>` for more
// information. A call of SetRGBWValues with index + length above the
// bounds is ignored completely.
// 
// The LPD8806 LED driver chips have 7-bit channels for RGB. Internally the LED
// Strip Bricklets divides the 8-bit values set using this function by 2 to make
// them 7-bit. Therefore, you can just use the normal value range (0-255) for
// LPD8806 LEDs.
// 
// The brightness channel of the APA102 LED driver chips has 5-bit. Internally the
// LED Strip Bricklets divides the 8-bit values set using this function by 8 to make
// them 5-bit. Therefore, you can just use the normal value range (0-255) for
// the brightness channel of APA102 LEDs.
// 
// .. versionadded:: 2.0.6$nbsp;(Plugin)
func (device *LEDStripBricklet) SetRGBWValues(index uint16, length uint8, r [12]uint8, g [12]uint8, b [12]uint8, w [12]uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, index);
	binary.Write(&buf, binary.LittleEndian, length);
	binary.Write(&buf, binary.LittleEndian, r);
	binary.Write(&buf, binary.LittleEndian, g);
	binary.Write(&buf, binary.LittleEndian, b);
	binary.Write(&buf, binary.LittleEndian, w);

	resultBytes, err := device.device.Set(uint8(FunctionSetRGBWValues), buf.Bytes())
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

// Returns *length* RGBW values starting from the given *index*.
// 
// The values are the last values that were set by SetRGBWValues.
// 
// .. versionadded:: 2.0.6$nbsp;(Plugin)
func (device *LEDStripBricklet) GetRGBWValues(index uint16, length uint8) (r [12]uint8, g [12]uint8, b [12]uint8, w [12]uint8, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, index);
	binary.Write(&buf, binary.LittleEndian, length);

	resultBytes, err := device.device.Get(uint8(FunctionGetRGBWValues), buf.Bytes())
	if err != nil {
		return r, g, b, w, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 56 {
			return r, g, b, w, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 56)
		}

		if header.ErrorCode != 0 {
			return r, g, b, w, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &r)
		binary.Read(resultBuf, binary.LittleEndian, &g)
		binary.Read(resultBuf, binary.LittleEndian, &b)
		binary.Read(resultBuf, binary.LittleEndian, &w)

	}

	return r, g, b, w, nil
}

// Sets the channel mapping for the connected LEDs.
// 
// SetRGBValues and SetRGBWValues take the data in RGB(W) order.
// But the connected LED driver chips might have their 3 or 4 channels in a
// different order. For example, the WS2801 chips typically use BGR order, the
// WS2812 chips typically use GRB order and the APA102 chips typically use WBGR
// order.
// 
// The APA102 chips are special. They have three 8-bit channels for RGB
// and an additional 5-bit channel for the overall brightness of the RGB LED
// making them 4-channel chips. Internally the brightness channel is the first
// channel, therefore one of the Wxyz channel mappings should be used. Then
// the W channel controls the brightness.
// 
// If a 3-channel mapping is selected then SetRGBValues has to be used.
// Calling SetRGBWValues with a 3-channel mapping will produce incorrect
// results. Vice-versa if a 4-channel mapping is selected then
// SetRGBWValues has to be used. Calling SetRGBValues with a
// 4-channel mapping will produce incorrect results.
// 
// .. versionadded:: 2.0.6$nbsp;(Plugin)
//
// Associated constants:
//
//	* ChannelMappingRGB
//	* ChannelMappingRBG
//	* ChannelMappingBRG
//	* ChannelMappingBGR
//	* ChannelMappingGRB
//	* ChannelMappingGBR
//	* ChannelMappingRGBW
//	* ChannelMappingRGWB
//	* ChannelMappingRBGW
//	* ChannelMappingRBWG
//	* ChannelMappingRWGB
//	* ChannelMappingRWBG
//	* ChannelMappingGRWB
//	* ChannelMappingGRBW
//	* ChannelMappingGBWR
//	* ChannelMappingGBRW
//	* ChannelMappingGWBR
//	* ChannelMappingGWRB
//	* ChannelMappingBRGW
//	* ChannelMappingBRWG
//	* ChannelMappingBGRW
//	* ChannelMappingBGWR
//	* ChannelMappingBWRG
//	* ChannelMappingBWGR
//	* ChannelMappingWRBG
//	* ChannelMappingWRGB
//	* ChannelMappingWGBR
//	* ChannelMappingWGRB
//	* ChannelMappingWBGR
//	* ChannelMappingWBRG
func (device *LEDStripBricklet) SetChannelMapping(mapping ChannelMapping) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, mapping);

	resultBytes, err := device.device.Set(uint8(FunctionSetChannelMapping), buf.Bytes())
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

// Returns the currently used channel mapping as set by SetChannelMapping.
// 
// .. versionadded:: 2.0.6$nbsp;(Plugin)
//
// Associated constants:
//
//	* ChannelMappingRGB
//	* ChannelMappingRBG
//	* ChannelMappingBRG
//	* ChannelMappingBGR
//	* ChannelMappingGRB
//	* ChannelMappingGBR
//	* ChannelMappingRGBW
//	* ChannelMappingRGWB
//	* ChannelMappingRBGW
//	* ChannelMappingRBWG
//	* ChannelMappingRWGB
//	* ChannelMappingRWBG
//	* ChannelMappingGRWB
//	* ChannelMappingGRBW
//	* ChannelMappingGBWR
//	* ChannelMappingGBRW
//	* ChannelMappingGWBR
//	* ChannelMappingGWRB
//	* ChannelMappingBRGW
//	* ChannelMappingBRWG
//	* ChannelMappingBGRW
//	* ChannelMappingBGWR
//	* ChannelMappingBWRG
//	* ChannelMappingBWGR
//	* ChannelMappingWRBG
//	* ChannelMappingWRGB
//	* ChannelMappingWGBR
//	* ChannelMappingWGRB
//	* ChannelMappingWBGR
//	* ChannelMappingWBRG
func (device *LEDStripBricklet) GetChannelMapping() (mapping ChannelMapping, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetChannelMapping), buf.Bytes())
	if err != nil {
		return mapping, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return mapping, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return mapping, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &mapping)

	}

	return mapping, nil
}

// Enables the RegisterFrameRenderedCallback callback.
// 
// By default the callback is enabled.
// 
// .. versionadded:: 2.0.6$nbsp;(Plugin)
func (device *LEDStripBricklet) EnableFrameRenderedCallback() (err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Set(uint8(FunctionEnableFrameRenderedCallback), buf.Bytes())
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

// Disables the RegisterFrameRenderedCallback callback.
// 
// By default the callback is enabled.
// 
// .. versionadded:: 2.0.6$nbsp;(Plugin)
func (device *LEDStripBricklet) DisableFrameRenderedCallback() (err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Set(uint8(FunctionDisableFrameRenderedCallback), buf.Bytes())
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

// Returns *true* if the RegisterFrameRenderedCallback callback is enabled, *false* otherwise.
// 
// .. versionadded:: 2.0.6$nbsp;(Plugin)
func (device *LEDStripBricklet) IsFrameRenderedCallbackEnabled() (enabled bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionIsFrameRenderedCallbackEnabled), buf.Bytes())
	if err != nil {
		return enabled, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return enabled, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return enabled, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &enabled)

	}

	return enabled, nil
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
func (device *LEDStripBricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
