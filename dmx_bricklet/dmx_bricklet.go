/* ***********************************************************
 * This file was automatically generated on 2021-01-15.      *
 *                                                           *
 * Go Bindings Version 2.0.10                                *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


// DMX master and slave.
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/DMX_Bricklet_Go.html.
package dmx_bricklet

import (
	"encoding/binary"
	"bytes"
	"fmt"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionSetDMXMode Function = 1
	FunctionGetDMXMode Function = 2
	FunctionWriteFrameLowLevel Function = 3
	FunctionReadFrameLowLevel Function = 4
	FunctionSetFrameDuration Function = 5
	FunctionGetFrameDuration Function = 6
	FunctionGetFrameErrorCount Function = 7
	FunctionSetCommunicationLEDConfig Function = 8
	FunctionGetCommunicationLEDConfig Function = 9
	FunctionSetErrorLEDConfig Function = 10
	FunctionGetErrorLEDConfig Function = 11
	FunctionSetFrameCallbackConfig Function = 12
	FunctionGetFrameCallbackConfig Function = 13
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
	FunctionCallbackFrameStarted Function = 14
	FunctionCallbackFrameAvailable Function = 15
	FunctionCallbackFrameLowLevel Function = 16
	FunctionCallbackFrameErrorCount Function = 17
)

type DMXMode = uint8

const (
	DMXModeMaster DMXMode = 0
	DMXModeSlave DMXMode = 1
)

type CommunicationLEDConfig = uint8

const (
	CommunicationLEDConfigOff CommunicationLEDConfig = 0
	CommunicationLEDConfigOn CommunicationLEDConfig = 1
	CommunicationLEDConfigShowHeartbeat CommunicationLEDConfig = 2
	CommunicationLEDConfigShowCommunication CommunicationLEDConfig = 3
)

type ErrorLEDConfig = uint8

const (
	ErrorLEDConfigOff ErrorLEDConfig = 0
	ErrorLEDConfigOn ErrorLEDConfig = 1
	ErrorLEDConfigShowHeartbeat ErrorLEDConfig = 2
	ErrorLEDConfigShowError ErrorLEDConfig = 3
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

type DMXBricklet struct {
	device Device
}
const DeviceIdentifier = 285
const DeviceDisplayName = "DMX Bricklet"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (DMXBricklet, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{ 2,0,0 }, uid, &internalIPCon, 3, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return DMXBricklet{}, err
	}
	dev.ResponseExpected[FunctionSetDMXMode] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetDMXMode] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionWriteFrameLowLevel] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionReadFrameLowLevel] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetFrameDuration] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetFrameDuration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetFrameErrorCount] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetCommunicationLEDConfig] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetCommunicationLEDConfig] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetErrorLEDConfig] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetErrorLEDConfig] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetFrameCallbackConfig] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetFrameCallbackConfig] = ResponseExpectedFlagAlwaysTrue;
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
	return DMXBricklet{dev}, nil
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
func (device *DMXBricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *DMXBricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *DMXBricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *DMXBricklet) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// This callback is triggered as soon as a new frame write is started.
// You should send the data for the next frame directly after this callback
// was triggered.
// 
// For an explanation of the general approach see WriteFrame.
// 
// This callback can be enabled via SetFrameCallbackConfig.
// 
// This callback can only be triggered in master mode.
func (device *DMXBricklet) RegisterFrameStartedCallback(fn func()) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 8 {
			return
		}
		
		
		
		fn()
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackFrameStarted), wrapper)
}

// Remove a registered Frame Started callback.
func (device *DMXBricklet) DeregisterFrameStartedCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackFrameStarted), registrationId)
}


// This callback is triggered in slave mode when a new frame was received from the DMX master
// and it can be read out. You have to read the frame before the master has written
// the next frame, see ReadFrame for more details.
// 
// The parameter is the frame number, it is increased by one with each received frame.
// 
// This callback can be enabled via SetFrameCallbackConfig.
// 
// This callback can only be triggered in slave mode.
func (device *DMXBricklet) RegisterFrameAvailableCallback(fn func(uint32)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 12 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var frameNumber uint32
		binary.Read(buf, binary.LittleEndian, &frameNumber)
		fn(frameNumber)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackFrameAvailable), wrapper)
}

// Remove a registered Frame Available callback.
func (device *DMXBricklet) DeregisterFrameAvailableCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackFrameAvailable), registrationId)
}


// See RegisterFrameCallback
func (device *DMXBricklet) RegisterFrameLowLevelCallback(fn func(uint16, uint16, [56]uint8, uint32)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 72 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var frameLength uint16
		var frameChunkOffset uint16
		var frameChunkData [56]uint8
		var frameNumber uint32
		binary.Read(buf, binary.LittleEndian, &frameLength)
		binary.Read(buf, binary.LittleEndian, &frameChunkOffset)
		copy(frameChunkData[:], ByteSliceToUint8Slice(buf.Next(8 * 56/8)))
		binary.Read(buf, binary.LittleEndian, &frameNumber)
		fn(frameLength, frameChunkOffset, frameChunkData, frameNumber)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackFrameLowLevel), wrapper)
}

// Remove a registered Frame Low Level callback.
func (device *DMXBricklet) DeregisterFrameLowLevelCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackFrameLowLevel), registrationId)
}


// This callback is called as soon as a new frame is available
// (written by the DMX master).
// 
// The size of the array is equivalent to the number of channels in
// the frame. Each byte represents one channel.
// 
// This callback can be enabled via SetFrameCallbackConfig.
// 
// This callback can only be triggered in slave mode.
// 
// Note
//  If reconstructing the value fails, the callback is triggered with nil for frame.
func (device *DMXBricklet) RegisterFrameCallback(fn func(uint32, []uint8)) uint64 {
	buf := make([]uint8, 0)
	wrapper := func(frameLength uint16, frameChunkOffset uint16, frameChunkData [56]uint8, frameNumber uint32)  {
		if int(frameChunkOffset) != len(buf) {
			if len(buf) != 0 {
				buf = nil
				fn(frameNumber, buf)
				buf = make([]uint8, 0)
			}
			if frameChunkOffset != 0 {
				return
			}
		}
		toRead := MinU(uint64(frameLength-frameChunkOffset), uint64(len(frameChunkData[:])))
		buf = append(buf, frameChunkData[:toRead]...)
		if len(buf) >= int(frameLength) {
			fn(frameNumber, buf)
			buf = make([]uint8, 0)
		}
	}
	return device.RegisterFrameLowLevelCallback(wrapper)
}

// Remove a registered Frame Low Level callback.
func (device *DMXBricklet) DeregisterFrameCallback(registrationId uint64) {
	device.DeregisterFrameLowLevelCallback(registrationId)
}


// This callback is called if a new error occurs. It returns
// the current overrun and framing error count.
func (device *DMXBricklet) RegisterFrameErrorCountCallback(fn func(uint32, uint32)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 16 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var overrunErrorCount uint32
		var framingErrorCount uint32
		binary.Read(buf, binary.LittleEndian, &overrunErrorCount)
		binary.Read(buf, binary.LittleEndian, &framingErrorCount)
		fn(overrunErrorCount, framingErrorCount)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackFrameErrorCount), wrapper)
}

// Remove a registered Frame Error Count callback.
func (device *DMXBricklet) DeregisterFrameErrorCountCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackFrameErrorCount), registrationId)
}


// Sets the DMX mode to either master or slave.
// 
// Calling this function sets frame number to 0.
//
// Associated constants:
//
//	* DMXModeMaster
//	* DMXModeSlave
func (device *DMXBricklet) SetDMXMode(dmxMode DMXMode) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, dmxMode);

	resultBytes, err := device.device.Set(uint8(FunctionSetDMXMode), buf.Bytes())
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

// Returns the DMX mode, as set by SetDMXMode.
//
// Associated constants:
//
//	* DMXModeMaster
//	* DMXModeSlave
func (device *DMXBricklet) GetDMXMode() (dmxMode DMXMode, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetDMXMode), buf.Bytes())
	if err != nil {
		return dmxMode, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return dmxMode, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return dmxMode, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &dmxMode)

	}

	return dmxMode, nil
}

// Writes a DMX frame. The maximum frame size is 512 byte. Each byte represents one channel.
// 
// The next frame can be written after the RegisterFrameStartedCallback callback was called. The frame
// is double buffered, so a new frame can be written as soon as the writing of the prior frame
// starts.
// 
// The data will be transfered when the next frame duration ends, see SetFrameDuration.
// 
// Generic approach:
// 
// * Set the frame duration to a value that represents the number of frames per second you want to achieve.
// * Set channels for first frame.
// * Wait for the RegisterFrameStartedCallback callback.
// * Set channels for next frame.
// * Wait for the RegisterFrameStartedCallback callback.
// * and so on.
// 
// This approach ensures that you can set new DMX data with a fixed frame rate.
// 
// This function can only be called in master mode.
func (device *DMXBricklet) WriteFrameLowLevel(frameLength uint16, frameChunkOffset uint16, frameChunkData [60]uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, frameLength);
	binary.Write(&buf, binary.LittleEndian, frameChunkOffset);
	buf.Write(Uint8SliceToByteSlice(frameChunkData[:]))

	resultBytes, err := device.device.Set(uint8(FunctionWriteFrameLowLevel), buf.Bytes())
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

// Writes a DMX frame. The maximum frame size is 512 byte. Each byte represents one channel.
// 
// The next frame can be written after the RegisterFrameStartedCallback callback was called. The frame
// is double buffered, so a new frame can be written as soon as the writing of the prior frame
// starts.
// 
// The data will be transfered when the next frame duration ends, see SetFrameDuration.
// 
// Generic approach:
// 
// * Set the frame duration to a value that represents the number of frames per second you want to achieve.
// * Set channels for first frame.
// * Wait for the RegisterFrameStartedCallback callback.
// * Set channels for next frame.
// * Wait for the RegisterFrameStartedCallback callback.
// * and so on.
// 
// This approach ensures that you can set new DMX data with a fixed frame rate.
// 
// This function can only be called in master mode.
	func (device *DMXBricklet) WriteFrame(frame []uint8) (err error) {
		_, err = device.device.SetHighLevel(func(frameLength uint64, frameChunkOffset uint64, frameChunkData []byte) (LowLevelWriteResult, error) {
			arr := [60]uint8{}
			copy(arr[:], ByteSliceToUint8Slice(frameChunkData))

			err := device.WriteFrameLowLevel(uint16(frameLength), uint16(frameChunkOffset), arr)

			var lowLevelResults bytes.Buffer
			

			return LowLevelWriteResult{
				uint64(60),
				lowLevelResults.Bytes()}, err
		}, 0, 8, 480, Uint8SliceToByteSlice(frame))

		if err != nil {
			return
		}

		
		
		
		return
	}

// Returns the last frame that was written by the DMX master. The size of the array
// is equivalent to the number of channels in the frame. Each byte represents one channel.
// 
// The next frame is available after the RegisterFrameAvailableCallback callback was called.
// 
// Generic approach:
// 
// * Call ReadFrame to get first frame.
// * Wait for the RegisterFrameAvailableCallback callback.
// * Call ReadFrame to get second frame.
// * Wait for the RegisterFrameAvailableCallback callback.
// * and so on.
// 
// Instead of polling this function you can also use the RegisterFrameCallback callback.
// You can enable it with SetFrameCallbackConfig.
// 
// The frame number starts at 0 and it is increased by one with each received frame.
// 
// This function can only be called in slave mode.
func (device *DMXBricklet) ReadFrameLowLevel() (frameLength uint16, frameChunkOffset uint16, frameChunkData [56]uint8, frameNumber uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionReadFrameLowLevel), buf.Bytes())
	if err != nil {
		return frameLength, frameChunkOffset, frameChunkData, frameNumber, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 72 {
			return frameLength, frameChunkOffset, frameChunkData, frameNumber, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 72)
		}

		if header.ErrorCode != 0 {
			return frameLength, frameChunkOffset, frameChunkData, frameNumber, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &frameLength)
		binary.Read(resultBuf, binary.LittleEndian, &frameChunkOffset)
		copy(frameChunkData[:], ByteSliceToUint8Slice(resultBuf.Next(8 * 56/8)))
		binary.Read(resultBuf, binary.LittleEndian, &frameNumber)

	}

	return frameLength, frameChunkOffset, frameChunkData, frameNumber, nil
}

// Returns the last frame that was written by the DMX master. The size of the array
// is equivalent to the number of channels in the frame. Each byte represents one channel.
// 
// The next frame is available after the RegisterFrameAvailableCallback callback was called.
// 
// Generic approach:
// 
// * Call ReadFrame to get first frame.
// * Wait for the RegisterFrameAvailableCallback callback.
// * Call ReadFrame to get second frame.
// * Wait for the RegisterFrameAvailableCallback callback.
// * and so on.
// 
// Instead of polling this function you can also use the RegisterFrameCallback callback.
// You can enable it with SetFrameCallbackConfig.
// 
// The frame number starts at 0 and it is increased by one with each received frame.
// 
// This function can only be called in slave mode.
	func (device *DMXBricklet) ReadFrame() (frame []uint8, frameNumber uint32, err error) {
		buf, result, err := device.device.GetHighLevel(func() (LowLevelResult, error) {
			frameLength, frameChunkOffset, frameChunkData, frameNumber, err := device.ReadFrameLowLevel()

			if err != nil {
				return LowLevelResult{}, err
			}

			var lowLevelResults bytes.Buffer
			binary.Write(&lowLevelResults, binary.LittleEndian, frameNumber);

			return LowLevelResult{
				uint64(frameLength),
				uint64(frameChunkOffset),
				Uint8SliceToByteSlice(frameChunkData[:]),
				lowLevelResults.Bytes()}, nil
		},
			1,
			8)
		if err != nil {
			return ByteSliceToUint8Slice(buf), frameNumber, err
		}
		resultBuf := bytes.NewBuffer(result)
		binary.Read(resultBuf, binary.LittleEndian, &frameNumber)
		return ByteSliceToUint8Slice(buf), frameNumber, nil
	}

// Sets the duration of a frame.
// 
// Example: If you want to achieve 20 frames per second, you should
// set the frame duration to 50ms (50ms * 20 = 1 second).
// 
// If you always want to send a frame as fast as possible you can set
// this value to 0.
// 
// This setting is only used in master mode.
func (device *DMXBricklet) SetFrameDuration(frameDuration uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, frameDuration);

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
func (device *DMXBricklet) GetFrameDuration() (frameDuration uint16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetFrameDuration), buf.Bytes())
	if err != nil {
		return frameDuration, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 10 {
			return frameDuration, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		if header.ErrorCode != 0 {
			return frameDuration, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &frameDuration)

	}

	return frameDuration, nil
}

// Returns the current number of overrun and framing errors.
func (device *DMXBricklet) GetFrameErrorCount() (overrunErrorCount uint32, framingErrorCount uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetFrameErrorCount), buf.Bytes())
	if err != nil {
		return overrunErrorCount, framingErrorCount, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 16 {
			return overrunErrorCount, framingErrorCount, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 16)
		}

		if header.ErrorCode != 0 {
			return overrunErrorCount, framingErrorCount, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &overrunErrorCount)
		binary.Read(resultBuf, binary.LittleEndian, &framingErrorCount)

	}

	return overrunErrorCount, framingErrorCount, nil
}

// Sets the communication LED configuration. By default the LED shows
// communication traffic, it flickers once for every 10 received data packets.
// 
// You can also turn the LED permanently on/off or show a heartbeat.
// 
// If the Bricklet is in bootloader mode, the LED is off.
//
// Associated constants:
//
//	* CommunicationLEDConfigOff
//	* CommunicationLEDConfigOn
//	* CommunicationLEDConfigShowHeartbeat
//	* CommunicationLEDConfigShowCommunication
func (device *DMXBricklet) SetCommunicationLEDConfig(config CommunicationLEDConfig) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, config);

	resultBytes, err := device.device.Set(uint8(FunctionSetCommunicationLEDConfig), buf.Bytes())
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

// Returns the configuration as set by SetCommunicationLEDConfig
//
// Associated constants:
//
//	* CommunicationLEDConfigOff
//	* CommunicationLEDConfigOn
//	* CommunicationLEDConfigShowHeartbeat
//	* CommunicationLEDConfigShowCommunication
func (device *DMXBricklet) GetCommunicationLEDConfig() (config CommunicationLEDConfig, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetCommunicationLEDConfig), buf.Bytes())
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

// Sets the error LED configuration.
// 
// By default the error LED turns on if there is any error (see RegisterFrameErrorCountCallback
// callback). If you call this function with the Show-Error option again, the LED
// will turn off until the next error occurs.
// 
// You can also turn the LED permanently on/off or show a heartbeat.
// 
// If the Bricklet is in bootloader mode, the LED is off.
//
// Associated constants:
//
//	* ErrorLEDConfigOff
//	* ErrorLEDConfigOn
//	* ErrorLEDConfigShowHeartbeat
//	* ErrorLEDConfigShowError
func (device *DMXBricklet) SetErrorLEDConfig(config ErrorLEDConfig) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, config);

	resultBytes, err := device.device.Set(uint8(FunctionSetErrorLEDConfig), buf.Bytes())
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

// Returns the configuration as set by SetErrorLEDConfig.
//
// Associated constants:
//
//	* ErrorLEDConfigOff
//	* ErrorLEDConfigOn
//	* ErrorLEDConfigShowHeartbeat
//	* ErrorLEDConfigShowError
func (device *DMXBricklet) GetErrorLEDConfig() (config ErrorLEDConfig, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetErrorLEDConfig), buf.Bytes())
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

// Enables/Disables the different callbacks. By default the
// RegisterFrameStartedCallback callback and RegisterFrameAvailableCallback callback are enabled while
// the RegisterFrameCallback callback and RegisterFrameErrorCountCallback callback are disabled.
// 
// If you want to use the RegisterFrameCallback callback you can enable it and disable
// the cb:`Frame Available` callback at the same time. It becomes redundant in
// this case.
func (device *DMXBricklet) SetFrameCallbackConfig(frameStartedCallbackEnabled bool, frameAvailableCallbackEnabled bool, frameCallbackEnabled bool, frameErrorCountCallbackEnabled bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, frameStartedCallbackEnabled);
	binary.Write(&buf, binary.LittleEndian, frameAvailableCallbackEnabled);
	binary.Write(&buf, binary.LittleEndian, frameCallbackEnabled);
	binary.Write(&buf, binary.LittleEndian, frameErrorCountCallbackEnabled);

	resultBytes, err := device.device.Set(uint8(FunctionSetFrameCallbackConfig), buf.Bytes())
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

// Returns the frame callback config as set by SetFrameCallbackConfig.
func (device *DMXBricklet) GetFrameCallbackConfig() (frameStartedCallbackEnabled bool, frameAvailableCallbackEnabled bool, frameCallbackEnabled bool, frameErrorCountCallbackEnabled bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetFrameCallbackConfig), buf.Bytes())
	if err != nil {
		return frameStartedCallbackEnabled, frameAvailableCallbackEnabled, frameCallbackEnabled, frameErrorCountCallbackEnabled, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 12 {
			return frameStartedCallbackEnabled, frameAvailableCallbackEnabled, frameCallbackEnabled, frameErrorCountCallbackEnabled, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		if header.ErrorCode != 0 {
			return frameStartedCallbackEnabled, frameAvailableCallbackEnabled, frameCallbackEnabled, frameErrorCountCallbackEnabled, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &frameStartedCallbackEnabled)
		binary.Read(resultBuf, binary.LittleEndian, &frameAvailableCallbackEnabled)
		binary.Read(resultBuf, binary.LittleEndian, &frameCallbackEnabled)
		binary.Read(resultBuf, binary.LittleEndian, &frameErrorCountCallbackEnabled)

	}

	return frameStartedCallbackEnabled, frameAvailableCallbackEnabled, frameCallbackEnabled, frameErrorCountCallbackEnabled, nil
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
func (device *DMXBricklet) GetSPITFPErrorCount() (errorCountAckChecksum uint32, errorCountMessageChecksum uint32, errorCountFrame uint32, errorCountOverflow uint32, err error) {
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
func (device *DMXBricklet) SetBootloaderMode(mode BootloaderMode) (status BootloaderStatus, err error) {
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
func (device *DMXBricklet) GetBootloaderMode() (mode BootloaderMode, err error) {
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
func (device *DMXBricklet) SetWriteFirmwarePointer(pointer uint32) (err error) {
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
func (device *DMXBricklet) WriteFirmware(data [64]uint8) (status uint8, err error) {
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
func (device *DMXBricklet) SetStatusLEDConfig(config StatusLEDConfig) (err error) {
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
func (device *DMXBricklet) GetStatusLEDConfig() (config StatusLEDConfig, err error) {
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
func (device *DMXBricklet) GetChipTemperature() (temperature int16, err error) {
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
func (device *DMXBricklet) Reset() (err error) {
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
func (device *DMXBricklet) WriteUID(uid uint32) (err error) {
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
func (device *DMXBricklet) ReadUID() (uid uint32, err error) {
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
func (device *DMXBricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
