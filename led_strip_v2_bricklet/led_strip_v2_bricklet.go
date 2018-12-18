/* ***********************************************************
 * This file was automatically generated on 2018-12-18.      *
 *                                                           *
 * Go Bindings Version 2.0.0                                 *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


//Controls up to 2048 RGB(W) LEDs.
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/LEDStripV2_Bricklet_Go.html.
package led_strip_v2_bricklet

import (
	"encoding/binary"
	"bytes"
    . "github.com/Tinkerforge/go-api-bindings/internal"
    "github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function uint8

const (
    FunctionSetLEDValuesLowLevel Function = 1
	FunctionGetLEDValuesLowLevel Function = 2
	FunctionSetFrameDuration Function = 3
	FunctionGetFrameDuration Function = 4
	FunctionGetSupplyVoltage Function = 5
	FunctionSetClockFrequency Function = 7
	FunctionGetClockFrequency Function = 8
	FunctionSetChipType Function = 9
	FunctionGetChipType Function = 10
	FunctionSetChannelMapping Function = 11
	FunctionGetChannelMapping Function = 12
	FunctionSetFrameStartedCallbackConfiguration Function = 13
	FunctionGetFrameStartedCallbackConfiguration Function = 14
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
	FunctionCallbackFrameStarted Function = 6
)

type ChipType uint16

const (
    ChipTypeWS2801 ChipType = 2801
	ChipTypeWS2811 ChipType = 2811
	ChipTypeWS2812 ChipType = 2812
	ChipTypeLPD8806 ChipType = 8806
	ChipTypeAPA102 ChipType = 102
)

type ChannelMapping uint8

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

type BootloaderMode uint8

const (
    BootloaderModeBootloader BootloaderMode = 0
	BootloaderModeFirmware BootloaderMode = 1
	BootloaderModeBootloaderWaitForReboot BootloaderMode = 2
	BootloaderModeFirmwareWaitForReboot BootloaderMode = 3
	BootloaderModeFirmwareWaitForEraseAndReboot BootloaderMode = 4
)

type BootloaderStatus uint8

const (
    BootloaderStatusOK BootloaderStatus = 0
	BootloaderStatusInvalidMode BootloaderStatus = 1
	BootloaderStatusNoChange BootloaderStatus = 2
	BootloaderStatusEntryFunctionNotPresent BootloaderStatus = 3
	BootloaderStatusDeviceIdentifierIncorrect BootloaderStatus = 4
	BootloaderStatusCRCMismatch BootloaderStatus = 5
)

type StatusLEDConfig uint8

const (
    StatusLEDConfigOff StatusLEDConfig = 0
	StatusLEDConfigOn StatusLEDConfig = 1
	StatusLEDConfigShowHeartbeat StatusLEDConfig = 2
	StatusLEDConfigShowStatus StatusLEDConfig = 3
)

type LEDStripV2Bricklet struct{
	device Device
}
const DeviceIdentifier = 2103
const DeviceDisplayName = "LED Strip Bricklet 2.0"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (LEDStripV2Bricklet, error) {
    internalIPCon := ipcon.GetInternalHandle().(IPConnection)
    dev, err := NewDevice([3]uint8{ 2,0,0 }, uid, &internalIPCon, 2)
    if err != nil {
        return LEDStripV2Bricklet{}, err
    }
    dev.ResponseExpected[FunctionSetLEDValuesLowLevel] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetLEDValuesLowLevel] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetFrameDuration] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetFrameDuration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetSupplyVoltage] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetClockFrequency] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetClockFrequency] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetChipType] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetChipType] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetChannelMapping] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetChannelMapping] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetFrameStartedCallbackConfiguration] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetFrameStartedCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue;
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
    return LEDStripV2Bricklet{dev}, nil
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
// for this purpose. If this flag is disabled for a setter function then no response is send
// and errors are silently ignored, because they cannot be detected.
// 
// See SetResponseExpected for the list of function ID constants available for this function.
func (device *LEDStripV2Bricklet) GetResponseExpected(functionID Function) (bool, error) {
    return device.device.GetResponseExpected(uint8(functionID))
}

// Changes the response expected flag of the function specified by the function ID parameter.
// This flag can only be changed for setter (default value: false) and callback configuration
// functions (default value: true). For getter functions it is always enabled.
// 
// Enabling the response expected flag for a setter function allows to detect timeouts and
// other error conditions calls of this setter as well. The device will then send a response
// for this purpose. If this flag is disabled for a setter function then no response is send
// and errors are silently ignored, because they cannot be detected.
func (device *LEDStripV2Bricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
    return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *LEDStripV2Bricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *LEDStripV2Bricklet) GetAPIVersion() [3]uint8 {
    return device.device.GetAPIVersion()
}

// This callback is triggered directly after a new frame render is started.
	// The parameter is the number of LEDs in that frame.
	// 
	// You should send the data for the next frame directly after this callback
	// was triggered.
	// 
	// For an explanation of the general approach see SetLEDValues.
func (device *LEDStripV2Bricklet) RegisterFrameStartedCallback(fn func(uint16)) uint64 {
            wrapper := func(byteSlice []byte) {
                buf := bytes.NewBuffer(byteSlice[8:])
                var length uint16
                binary.Read(buf, binary.LittleEndian, &length)
                fn(length)
            }
    return device.device.RegisterCallback(uint8(FunctionCallbackFrameStarted), wrapper)
}

//Remove a registered Frame Started callback.
func (device *LEDStripV2Bricklet) DeregisterFrameStartedCallback(callbackID uint64) {
    device.device.DeregisterCallback(uint8(FunctionCallbackFrameStarted), callbackID)
}


// Sets the RGB(W) values for the LEDs starting from *index*.
	// You can set at most 2048 RGB values or 1536 RGBW values.
	// 
	// To make the colors show correctly you need to configure the chip type
	// (see SetChipType) and a channel mapping (see SetChannelMapping)
	// according to the connected LEDs.
	// 
	// If the channel mapping has 3 colors, you need to give the data in the sequence
	// RGBRGBRGB... if the channel mapping has 4 colors you need to give data in the
	// sequence RGBWRGBWRGBW...
	// 
	// The data is double buffered and the colors will be transfered to the
	// LEDs when the next frame duration ends (see SetFrameDuration).
	// 
	// Generic approach:
	// 
	// * Set the frame duration to a value that represents the number of frames per
	//   second you want to achieve.
	// * Set all of the LED colors for one frame.
	// * Wait for the RegisterFrameStartedCallback callback.
	// * Set all of the LED colors for next frame.
	// * Wait for the RegisterFrameStartedCallback callback.
	// * And so on.
	// 
	// This approach ensures that you can change the LED colors with a fixed frame rate.
func (device *LEDStripV2Bricklet) SetLEDValuesLowLevel(index uint16, valueLength uint16, valueChunkOffset uint16, valueChunkData [58]uint8) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, index);
	binary.Write(&buf, binary.LittleEndian, valueLength);
	binary.Write(&buf, binary.LittleEndian, valueChunkOffset);
	buf.Write(Uint8SliceToByteSlice(valueChunkData[:]))

    resultBytes, err := device.device.Set(uint8(FunctionSetLEDValuesLowLevel), buf.Bytes())
    if err != nil {
        return err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return BrickletError(header.ErrorCode)
        }

        bytes.NewBuffer(resultBytes[8:])
        
    }
    
    return nil
}

// Sets the RGB(W) values for the LEDs starting from *index*.
	// You can set at most 2048 RGB values or 1536 RGBW values.
	// 
	// To make the colors show correctly you need to configure the chip type
	// (see SetChipType) and a channel mapping (see SetChannelMapping)
	// according to the connected LEDs.
	// 
	// If the channel mapping has 3 colors, you need to give the data in the sequence
	// RGBRGBRGB... if the channel mapping has 4 colors you need to give data in the
	// sequence RGBWRGBWRGBW...
	// 
	// The data is double buffered and the colors will be transfered to the
	// LEDs when the next frame duration ends (see SetFrameDuration).
	// 
	// Generic approach:
	// 
	// * Set the frame duration to a value that represents the number of frames per
	//   second you want to achieve.
	// * Set all of the LED colors for one frame.
	// * Wait for the RegisterFrameStartedCallback callback.
	// * Set all of the LED colors for next frame.
	// * Wait for the RegisterFrameStartedCallback callback.
	// * And so on.
	// 
	// This approach ensures that you can change the LED colors with a fixed frame rate.
	func (device *LEDStripV2Bricklet) SetLEDValues(index uint16, value []uint8) (err error) {            
        _, err = device.device.SetHighLevel(func(valueLength uint64, valueChunkOffset uint64, valueChunkData []byte) (LowLevelWriteResult, error) {
            arr := [58]uint8{}
            copy(arr[:], ByteSliceToUint8Slice(valueChunkData))

            err := device.SetLEDValuesLowLevel(index, uint16(valueLength), uint16(valueChunkOffset), arr)

            var lowLevelResults bytes.Buffer
            

            return LowLevelWriteResult{
                uint64(58),
                lowLevelResults.Bytes()}, err
        }, 0, 8, 464, Uint8SliceToByteSlice(value))   

         if err != nil {
            return
        }

        
        
        
        return
    }

// Returns the RGB(W) values as set by SetLEDValues.
func (device *LEDStripV2Bricklet) GetLEDValuesLowLevel(index uint16, length uint16) (valueLength uint16, valueChunkOffset uint16, valueChunkData [60]uint8, err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, index);
	binary.Write(&buf, binary.LittleEndian, length);

    resultBytes, err := device.device.Get(uint8(FunctionGetLEDValuesLowLevel), buf.Bytes())
    if err != nil {
        return valueLength, valueChunkOffset, valueChunkData, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return valueLength, valueChunkOffset, valueChunkData, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &valueLength)
	binary.Read(resultBuf, binary.LittleEndian, &valueChunkOffset)
	copy(valueChunkData[:], ByteSliceToUint8Slice(resultBuf.Next(8 * 60/8)))

    }
    
    return valueLength, valueChunkOffset, valueChunkData, nil
}

// Returns the RGB(W) values as set by SetLEDValues.
	func (device *LEDStripV2Bricklet) GetLEDValues(index uint16, length uint16) (value []uint8, err error) {
        buf, _, err := device.device.GetHighLevel(func() (LowLevelResult, error) {
            valueLength, valueChunkOffset, valueChunkData, err := device.GetLEDValuesLowLevel(index, length)            

            if err != nil {
                return LowLevelResult{}, err
            }

            var lowLevelResults bytes.Buffer
            

            return LowLevelResult{
                uint64(valueLength),
                uint64(valueChunkOffset),
                Uint8SliceToByteSlice(valueChunkData[:]),
                lowLevelResults.Bytes()}, nil
        },
            1,
            8)
        if err != nil {
            return ByteSliceToUint8Slice(buf), err
        }
        
        
        return ByteSliceToUint8Slice(buf), nil
    }

// Sets the frame duration in ms.
	// 
	// Example: If you want to achieve 20 frames per second, you should
	// set the frame duration to 50ms (50ms * 20 = 1 second).
	// 
	// For an explanation of the general approach see SetLEDValues.
	// 
	// Default value: 100ms (10 frames per second).
func (device *LEDStripV2Bricklet) SetFrameDuration(duration uint16) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, duration);

    resultBytes, err := device.device.Set(uint8(FunctionSetFrameDuration), buf.Bytes())
    if err != nil {
        return err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return BrickletError(header.ErrorCode)
        }

        bytes.NewBuffer(resultBytes[8:])
        
    }
    
    return nil
}

// Returns the frame duration in ms as set by SetFrameDuration.
func (device *LEDStripV2Bricklet) GetFrameDuration() (duration uint16, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetFrameDuration), buf.Bytes())
    if err != nil {
        return duration, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return duration, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &duration)

    }
    
    return duration, nil
}

// Returns the current supply voltage of the LEDs. The voltage is given in mV.
func (device *LEDStripV2Bricklet) GetSupplyVoltage() (voltage uint16, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetSupplyVoltage), buf.Bytes())
    if err != nil {
        return voltage, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return voltage, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &voltage)

    }
    
    return voltage, nil
}

// Sets the frequency of the clock in Hz. The range is 10000Hz (10kHz) up to
	// 2000000Hz (2MHz).
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
	// The default value is 1.66MHz.
func (device *LEDStripV2Bricklet) SetClockFrequency(frequency uint32) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, frequency);

    resultBytes, err := device.device.Set(uint8(FunctionSetClockFrequency), buf.Bytes())
    if err != nil {
        return err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return BrickletError(header.ErrorCode)
        }

        bytes.NewBuffer(resultBytes[8:])
        
    }
    
    return nil
}

// Returns the currently used clock frequency as set by SetClockFrequency.
func (device *LEDStripV2Bricklet) GetClockFrequency() (frequency uint32, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetClockFrequency), buf.Bytes())
    if err != nil {
        return frequency, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return frequency, BrickletError(header.ErrorCode)
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
	// The default value is WS2801 (2801).
//
// Associated constants:
//
//	* ChipTypeWS2801
//	* ChipTypeWS2811
//	* ChipTypeWS2812
//	* ChipTypeLPD8806
//	* ChipTypeAPA102
func (device *LEDStripV2Bricklet) SetChipType(chip ChipType) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, chip);

    resultBytes, err := device.device.Set(uint8(FunctionSetChipType), buf.Bytes())
    if err != nil {
        return err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return BrickletError(header.ErrorCode)
        }

        bytes.NewBuffer(resultBytes[8:])
        
    }
    
    return nil
}

// Returns the currently used chip type as set by SetChipType.
//
// Associated constants:
//
//	* ChipTypeWS2801
//	* ChipTypeWS2811
//	* ChipTypeWS2812
//	* ChipTypeLPD8806
//	* ChipTypeAPA102
func (device *LEDStripV2Bricklet) GetChipType() (chip ChipType, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetChipType), buf.Bytes())
    if err != nil {
        return chip, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return chip, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &chip)

    }
    
    return chip, nil
}

// Sets the channel mapping for the connected LEDs.
	// 
	// If the mapping has 4 colors, the function SetLEDValues expects 4
	// values per pixel and if the mapping has 3 colors it expects 3 values per pixel.
	// 
	// The function always expects the order RGB(W). The connected LED driver chips
	// might have their 3 or 4 channels in a different order. For example, the WS2801
	// chips typically use BGR order, then WS2812 chips typically use GRB order and
	// the APA102 chips typically use WBGR order.
	// 
	// The APA102 chips are special. They have three 8-bit channels for RGB
	// and an additional 5-bit channel for the overall brightness of the RGB LED
	// making them 4-channel chips. Internally the brightness channel is the first
	// channel, therefore one of the Wxyz channel mappings should be used. Then
	// the W channel controls the brightness.
	// 
	// The default value is BGR (36).
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
func (device *LEDStripV2Bricklet) SetChannelMapping(mapping ChannelMapping) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, mapping);

    resultBytes, err := device.device.Set(uint8(FunctionSetChannelMapping), buf.Bytes())
    if err != nil {
        return err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return BrickletError(header.ErrorCode)
        }

        bytes.NewBuffer(resultBytes[8:])
        
    }
    
    return nil
}

// Returns the currently used channel mapping as set by SetChannelMapping.
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
func (device *LEDStripV2Bricklet) GetChannelMapping() (mapping ChannelMapping, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetChannelMapping), buf.Bytes())
    if err != nil {
        return mapping, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return mapping, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &mapping)

    }
    
    return mapping, nil
}

// Enables/disables the RegisterFrameStartedCallback callback.
	// 
	// By default the callback is enabled.
func (device *LEDStripV2Bricklet) SetFrameStartedCallbackConfiguration(enable bool) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, enable);

    resultBytes, err := device.device.Set(uint8(FunctionSetFrameStartedCallbackConfiguration), buf.Bytes())
    if err != nil {
        return err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return BrickletError(header.ErrorCode)
        }

        bytes.NewBuffer(resultBytes[8:])
        
    }
    
    return nil
}

// Returns the configuration as set by
	// SetFrameStartedCallbackConfiguration.
func (device *LEDStripV2Bricklet) GetFrameStartedCallbackConfiguration() (enable bool, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetFrameStartedCallbackConfiguration), buf.Bytes())
    if err != nil {
        return enable, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return enable, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &enable)

    }
    
    return enable, nil
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
func (device *LEDStripV2Bricklet) GetSPITFPErrorCount() (errorCountAckChecksum uint32, errorCountMessageChecksum uint32, errorCountFrame uint32, errorCountOverflow uint32, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetSPITFPErrorCount), buf.Bytes())
    if err != nil {
        return errorCountAckChecksum, errorCountMessageChecksum, errorCountFrame, errorCountOverflow, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return errorCountAckChecksum, errorCountMessageChecksum, errorCountFrame, errorCountOverflow, BrickletError(header.ErrorCode)
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
func (device *LEDStripV2Bricklet) SetBootloaderMode(mode BootloaderMode) (status BootloaderStatus, err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, mode);

    resultBytes, err := device.device.Get(uint8(FunctionSetBootloaderMode), buf.Bytes())
    if err != nil {
        return status, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return status, BrickletError(header.ErrorCode)
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
func (device *LEDStripV2Bricklet) GetBootloaderMode() (mode BootloaderMode, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetBootloaderMode), buf.Bytes())
    if err != nil {
        return mode, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return mode, BrickletError(header.ErrorCode)
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
func (device *LEDStripV2Bricklet) SetWriteFirmwarePointer(pointer uint32) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, pointer);

    resultBytes, err := device.device.Set(uint8(FunctionSetWriteFirmwarePointer), buf.Bytes())
    if err != nil {
        return err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return BrickletError(header.ErrorCode)
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
func (device *LEDStripV2Bricklet) WriteFirmware(data [64]uint8) (status uint8, err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, data);

    resultBytes, err := device.device.Get(uint8(FunctionWriteFirmware), buf.Bytes())
    if err != nil {
        return status, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return status, BrickletError(header.ErrorCode)
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
func (device *LEDStripV2Bricklet) SetStatusLEDConfig(config StatusLEDConfig) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, config);

    resultBytes, err := device.device.Set(uint8(FunctionSetStatusLEDConfig), buf.Bytes())
    if err != nil {
        return err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return BrickletError(header.ErrorCode)
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
func (device *LEDStripV2Bricklet) GetStatusLEDConfig() (config StatusLEDConfig, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetStatusLEDConfig), buf.Bytes())
    if err != nil {
        return config, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return config, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &config)

    }
    
    return config, nil
}

// Returns the temperature in °C as measured inside the microcontroller. The
	// value returned is not the ambient temperature!
	// 
	// The temperature is only proportional to the real temperature and it has bad
	// accuracy. Practically it is only useful as an indicator for
	// temperature changes.
func (device *LEDStripV2Bricklet) GetChipTemperature() (temperature int16, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetChipTemperature), buf.Bytes())
    if err != nil {
        return temperature, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return temperature, BrickletError(header.ErrorCode)
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
func (device *LEDStripV2Bricklet) Reset() (err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Set(uint8(FunctionReset), buf.Bytes())
    if err != nil {
        return err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return BrickletError(header.ErrorCode)
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
func (device *LEDStripV2Bricklet) WriteUID(uid uint32) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, uid);

    resultBytes, err := device.device.Set(uint8(FunctionWriteUID), buf.Bytes())
    if err != nil {
        return err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return BrickletError(header.ErrorCode)
        }

        bytes.NewBuffer(resultBytes[8:])
        
    }
    
    return nil
}

// Returns the current UID as an integer. Encode as
	// Base58 to get the usual string version.
func (device *LEDStripV2Bricklet) ReadUID() (uid uint32, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionReadUID), buf.Bytes())
    if err != nil {
        return uid, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return uid, BrickletError(header.ErrorCode)
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
	// The position can be 'a', 'b', 'c' or 'd'.
	// 
	// The device identifier numbers can be found `here <device_identifier>`.
	// |device_identifier_constant|
func (device *LEDStripV2Bricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetIdentity), buf.Bytes())
    if err != nil {
        return uid, connectedUid, position, hardwareVersion, firmwareVersion, deviceIdentifier, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return uid, connectedUid, position, hardwareVersion, firmwareVersion, deviceIdentifier, BrickletError(header.ErrorCode)
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
