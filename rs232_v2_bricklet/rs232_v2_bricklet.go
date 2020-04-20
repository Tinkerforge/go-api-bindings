/* ***********************************************************
 * This file was automatically generated on 2020-04-20.      *
 *                                                           *
 * Go Bindings Version 2.0.7                                 *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


// Communicates with RS232 devices.
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/RS232V2_Bricklet_Go.html.
package rs232_v2_bricklet

import (
	"encoding/binary"
	"bytes"
	"fmt"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionWriteLowLevel Function = 1
	FunctionReadLowLevel Function = 2
	FunctionEnableReadCallback Function = 3
	FunctionDisableReadCallback Function = 4
	FunctionIsReadCallbackEnabled Function = 5
	FunctionSetConfiguration Function = 6
	FunctionGetConfiguration Function = 7
	FunctionSetBufferConfig Function = 8
	FunctionGetBufferConfig Function = 9
	FunctionGetBufferStatus Function = 10
	FunctionGetErrorCount Function = 11
	FunctionSetFrameReadableCallbackConfiguration Function = 14
	FunctionGetFrameReadableCallbackConfiguration Function = 15
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
	FunctionCallbackReadLowLevel Function = 12
	FunctionCallbackErrorCount Function = 13
	FunctionCallbackFrameReadable Function = 16
)

type Parity = uint8

const (
	ParityNone Parity = 0
	ParityOdd Parity = 1
	ParityEven Parity = 2
)

type Stopbits = uint8

const (
	Stopbits1 Stopbits = 1
	Stopbits2 Stopbits = 2
)

type Wordlength = uint8

const (
	Wordlength5 Wordlength = 5
	Wordlength6 Wordlength = 6
	Wordlength7 Wordlength = 7
	Wordlength8 Wordlength = 8
)

type Flowcontrol = uint8

const (
	FlowcontrolOff Flowcontrol = 0
	FlowcontrolSoftware Flowcontrol = 1
	FlowcontrolHardware Flowcontrol = 2
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

type RS232V2Bricklet struct {
	device Device
}
const DeviceIdentifier = 2108
const DeviceDisplayName = "RS232 Bricklet 2.0"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (RS232V2Bricklet, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{ 2,0,1 }, uid, &internalIPCon, 3, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return RS232V2Bricklet{}, err
	}
	dev.ResponseExpected[FunctionWriteLowLevel] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionReadLowLevel] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionEnableReadCallback] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionDisableReadCallback] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionIsReadCallbackEnabled] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetConfiguration] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetBufferConfig] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetBufferConfig] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetBufferStatus] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetErrorCount] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetFrameReadableCallbackConfiguration] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetFrameReadableCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue;
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
	return RS232V2Bricklet{dev}, nil
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
func (device *RS232V2Bricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *RS232V2Bricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *RS232V2Bricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *RS232V2Bricklet) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// See RegisterReadCallback
func (device *RS232V2Bricklet) RegisterReadLowLevelCallback(fn func(uint16, uint16, [60]rune)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 72 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var messageLength uint16
		var messageChunkOffset uint16
		var messageChunkData [60]rune
		binary.Read(buf, binary.LittleEndian, &messageLength)
		binary.Read(buf, binary.LittleEndian, &messageChunkOffset)
		copy(messageChunkData[:], ByteSliceToRuneSlice(buf.Next(8 * 60/8)))
		fn(messageLength, messageChunkOffset, messageChunkData)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackReadLowLevel), wrapper)
}

// Remove a registered Read Low Level callback.
func (device *RS232V2Bricklet) DeregisterReadLowLevelCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackReadLowLevel), registrationId)
}


// This callback is called if new data is available.
// 
// To enable this callback, use EnableReadCallback.
// 
// Note
//  If reconstructing the value fails, the callback is triggered with nil for message.
func (device *RS232V2Bricklet) RegisterReadCallback(fn func([]rune)) uint64 {
	buf := make([]rune, 0)
	wrapper := func(messageLength uint16, messageChunkOffset uint16, messageChunkData [60]rune)  {
		if int(messageChunkOffset) != len(buf) {
			if len(buf) != 0 {
				buf = nil
				fn(buf)
				buf = make([]rune, 0)
			}
			if messageChunkOffset != 0 {
				return
			}
		}
		toRead := MinU(uint64(messageLength-messageChunkOffset), uint64(len(messageChunkData[:])))
		buf = append(buf, messageChunkData[:toRead]...)
		if len(buf) >= int(messageLength) {
			fn(buf)
			buf = make([]rune, 0)
		}
	}
	return device.RegisterReadLowLevelCallback(wrapper)
}

// Remove a registered Read Low Level callback.
func (device *RS232V2Bricklet) DeregisterReadCallback(registrationId uint64) {
	device.DeregisterReadLowLevelCallback(registrationId)
}


// This callback is called if a new error occurs. It returns
// the current overrun and parity error count.
func (device *RS232V2Bricklet) RegisterErrorCountCallback(fn func(uint32, uint32)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 16 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var errorCountOverrun uint32
		var errorCountParity uint32
		binary.Read(buf, binary.LittleEndian, &errorCountOverrun)
		binary.Read(buf, binary.LittleEndian, &errorCountParity)
		fn(errorCountOverrun, errorCountParity)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackErrorCount), wrapper)
}

// Remove a registered Error Count callback.
func (device *RS232V2Bricklet) DeregisterErrorCountCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackErrorCount), registrationId)
}


// This callback is called if at least one frame of data is readable. The frame size is configured with SetFrameReadableCallbackConfiguration.
// The frame count parameter is the number of frames that can be read.
// This callback is triggered only once until Read is called. This means, that if you have configured a frame size of X bytes,
// you can read exactly X bytes using the Read function, every time the callback triggers without checking the frame count parameter.
// 
// .. versionadded:: 2.0.3$nbsp;(Plugin)
func (device *RS232V2Bricklet) RegisterFrameReadableCallback(fn func(uint16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 10 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var frameCount uint16
		binary.Read(buf, binary.LittleEndian, &frameCount)
		fn(frameCount)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackFrameReadable), wrapper)
}

// Remove a registered Frame Readable callback.
func (device *RS232V2Bricklet) DeregisterFrameReadableCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackFrameReadable), registrationId)
}


// Writes characters to the RS232 interface. The characters can be binary data,
// ASCII or similar is not necessary.
// 
// The return value is the number of characters that were written.
// 
// See SetConfiguration for configuration possibilities
// regarding baud rate, parity and so on.
func (device *RS232V2Bricklet) WriteLowLevel(messageLength uint16, messageChunkOffset uint16, messageChunkData [60]rune) (messageChunkWritten uint8, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, messageLength);
	binary.Write(&buf, binary.LittleEndian, messageChunkOffset);
	buf.Write(RuneSliceToByteSlice(messageChunkData[:]))

	resultBytes, err := device.device.Get(uint8(FunctionWriteLowLevel), buf.Bytes())
	if err != nil {
		return messageChunkWritten, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return messageChunkWritten, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return messageChunkWritten, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &messageChunkWritten)

	}

	return messageChunkWritten, nil
}

// Writes characters to the RS232 interface. The characters can be binary data,
// ASCII or similar is not necessary.
// 
// The return value is the number of characters that were written.
// 
// See SetConfiguration for configuration possibilities
// regarding baud rate, parity and so on.
	func (device *RS232V2Bricklet) Write(message []rune) (messageChunkWritten uint64, err error) {
		lowLevelResult, err := device.device.SetHighLevel(func(messageLength uint64, messageChunkOffset uint64, messageChunkData []byte) (LowLevelWriteResult, error) {
			arr := [60]rune{}
			copy(arr[:], ByteSliceToRuneSlice(messageChunkData))

			messageChunkWritten, err := device.WriteLowLevel(uint16(messageLength), uint16(messageChunkOffset), arr)

			var lowLevelResults bytes.Buffer
			

			return LowLevelWriteResult{
				uint64(messageChunkWritten),
				lowLevelResults.Bytes()}, err
		}, 0, 8, 480, RuneSliceToByteSlice(message))

		if err != nil {
			return
		}

		
		
		messageChunkWritten = lowLevelResult.Written
		return
	}

// Returns up to *length* characters from receive buffer.
// 
// Instead of polling with this function, you can also use
// callbacks. But note that this function will return available
// data only when the read callback is disabled.
// See EnableReadCallback and RegisterReadCallback callback.
func (device *RS232V2Bricklet) ReadLowLevel(length uint16) (messageLength uint16, messageChunkOffset uint16, messageChunkData [60]rune, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, length);

	resultBytes, err := device.device.Get(uint8(FunctionReadLowLevel), buf.Bytes())
	if err != nil {
		return messageLength, messageChunkOffset, messageChunkData, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 72 {
			return messageLength, messageChunkOffset, messageChunkData, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 72)
		}

		if header.ErrorCode != 0 {
			return messageLength, messageChunkOffset, messageChunkData, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &messageLength)
		binary.Read(resultBuf, binary.LittleEndian, &messageChunkOffset)
		copy(messageChunkData[:], ByteSliceToRuneSlice(resultBuf.Next(8 * 60/8)))

	}

	return messageLength, messageChunkOffset, messageChunkData, nil
}

// Returns up to *length* characters from receive buffer.
// 
// Instead of polling with this function, you can also use
// callbacks. But note that this function will return available
// data only when the read callback is disabled.
// See EnableReadCallback and RegisterReadCallback callback.
	func (device *RS232V2Bricklet) Read(length uint16) (message []rune, err error) {
		buf, _, err := device.device.GetHighLevel(func() (LowLevelResult, error) {
			messageLength, messageChunkOffset, messageChunkData, err := device.ReadLowLevel(length)

			if err != nil {
				return LowLevelResult{}, err
			}

			var lowLevelResults bytes.Buffer
			

			return LowLevelResult{
				uint64(messageLength),
				uint64(messageChunkOffset),
				RuneSliceToByteSlice(messageChunkData[:]),
				lowLevelResults.Bytes()}, nil
		},
			1,
			8)
		if err != nil {
			return ByteSliceToRuneSlice(buf), err
		}
		
		
		return ByteSliceToRuneSlice(buf), nil
	}

// Enables the RegisterReadCallback callback. This will disable the RegisterFrameReadableCallback callback.
// 
// By default the callback is disabled.
func (device *RS232V2Bricklet) EnableReadCallback() (err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Set(uint8(FunctionEnableReadCallback), buf.Bytes())
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

// Disables the RegisterReadCallback callback.
// 
// By default the callback is disabled.
func (device *RS232V2Bricklet) DisableReadCallback() (err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Set(uint8(FunctionDisableReadCallback), buf.Bytes())
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

// Returns *true* if the RegisterReadCallback callback is enabled,
// *false* otherwise.
func (device *RS232V2Bricklet) IsReadCallbackEnabled() (enabled bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionIsReadCallbackEnabled), buf.Bytes())
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

// Sets the configuration for the RS232 communication.
//
// Associated constants:
//
//	* ParityNone
//	* ParityOdd
//	* ParityEven
//	* Stopbits1
//	* Stopbits2
//	* Wordlength5
//	* Wordlength6
//	* Wordlength7
//	* Wordlength8
//	* FlowcontrolOff
//	* FlowcontrolSoftware
//	* FlowcontrolHardware
func (device *RS232V2Bricklet) SetConfiguration(baudrate uint32, parity Parity, stopbits Stopbits, wordlength Wordlength, flowcontrol Flowcontrol) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, baudrate);
	binary.Write(&buf, binary.LittleEndian, parity);
	binary.Write(&buf, binary.LittleEndian, stopbits);
	binary.Write(&buf, binary.LittleEndian, wordlength);
	binary.Write(&buf, binary.LittleEndian, flowcontrol);

	resultBytes, err := device.device.Set(uint8(FunctionSetConfiguration), buf.Bytes())
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

// Returns the configuration as set by SetConfiguration.
//
// Associated constants:
//
//	* ParityNone
//	* ParityOdd
//	* ParityEven
//	* Stopbits1
//	* Stopbits2
//	* Wordlength5
//	* Wordlength6
//	* Wordlength7
//	* Wordlength8
//	* FlowcontrolOff
//	* FlowcontrolSoftware
//	* FlowcontrolHardware
func (device *RS232V2Bricklet) GetConfiguration() (baudrate uint32, parity Parity, stopbits Stopbits, wordlength Wordlength, flowcontrol Flowcontrol, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetConfiguration), buf.Bytes())
	if err != nil {
		return baudrate, parity, stopbits, wordlength, flowcontrol, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 16 {
			return baudrate, parity, stopbits, wordlength, flowcontrol, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 16)
		}

		if header.ErrorCode != 0 {
			return baudrate, parity, stopbits, wordlength, flowcontrol, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &baudrate)
		binary.Read(resultBuf, binary.LittleEndian, &parity)
		binary.Read(resultBuf, binary.LittleEndian, &stopbits)
		binary.Read(resultBuf, binary.LittleEndian, &wordlength)
		binary.Read(resultBuf, binary.LittleEndian, &flowcontrol)

	}

	return baudrate, parity, stopbits, wordlength, flowcontrol, nil
}

// Sets the send and receive buffer size in byte. In total the buffers have to be
// 10240 byte (10KiB) in size, the minimum buffer size is 1024 byte (1KiB) for each.
// 
// The current buffer content is lost if this function is called.
// 
// The send buffer holds data that is given by Write and
// can not be written yet. The receive buffer holds data that is
// received through RS232 but could not yet be send to the
// user, either by Read or through RegisterReadCallback callback.
func (device *RS232V2Bricklet) SetBufferConfig(sendBufferSize uint16, receiveBufferSize uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, sendBufferSize);
	binary.Write(&buf, binary.LittleEndian, receiveBufferSize);

	resultBytes, err := device.device.Set(uint8(FunctionSetBufferConfig), buf.Bytes())
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

// Returns the buffer configuration as set by SetBufferConfig.
func (device *RS232V2Bricklet) GetBufferConfig() (sendBufferSize uint16, receiveBufferSize uint16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetBufferConfig), buf.Bytes())
	if err != nil {
		return sendBufferSize, receiveBufferSize, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 12 {
			return sendBufferSize, receiveBufferSize, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		if header.ErrorCode != 0 {
			return sendBufferSize, receiveBufferSize, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &sendBufferSize)
		binary.Read(resultBuf, binary.LittleEndian, &receiveBufferSize)

	}

	return sendBufferSize, receiveBufferSize, nil
}

// Returns the currently used bytes for the send and received buffer.
// 
// See SetBufferConfig for buffer size configuration.
func (device *RS232V2Bricklet) GetBufferStatus() (sendBufferUsed uint16, receiveBufferUsed uint16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetBufferStatus), buf.Bytes())
	if err != nil {
		return sendBufferUsed, receiveBufferUsed, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 12 {
			return sendBufferUsed, receiveBufferUsed, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		if header.ErrorCode != 0 {
			return sendBufferUsed, receiveBufferUsed, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &sendBufferUsed)
		binary.Read(resultBuf, binary.LittleEndian, &receiveBufferUsed)

	}

	return sendBufferUsed, receiveBufferUsed, nil
}

// Returns the current number of overrun and parity errors.
func (device *RS232V2Bricklet) GetErrorCount() (errorCountOverrun uint32, errorCountParity uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetErrorCount), buf.Bytes())
	if err != nil {
		return errorCountOverrun, errorCountParity, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 16 {
			return errorCountOverrun, errorCountParity, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 16)
		}

		if header.ErrorCode != 0 {
			return errorCountOverrun, errorCountParity, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &errorCountOverrun)
		binary.Read(resultBuf, binary.LittleEndian, &errorCountParity)

	}

	return errorCountOverrun, errorCountParity, nil
}

// Configures the RegisterFrameReadableCallback callback. The frame size is the number of bytes, that have to be readable to trigger the callback.
// A frame size of 0 disables the callback. A frame size greater than 0 enables the callback and disables the RegisterReadCallback callback.
// 
// By default the callback is disabled.
// 
// .. versionadded:: 2.0.3$nbsp;(Plugin)
func (device *RS232V2Bricklet) SetFrameReadableCallbackConfiguration(frameSize uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, frameSize);

	resultBytes, err := device.device.Set(uint8(FunctionSetFrameReadableCallbackConfiguration), buf.Bytes())
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

// Returns the callback configuration as set by SetFrameReadableCallbackConfiguration.
// 
// .. versionadded:: 2.0.3$nbsp;(Plugin)
func (device *RS232V2Bricklet) GetFrameReadableCallbackConfiguration() (frameSize uint16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetFrameReadableCallbackConfiguration), buf.Bytes())
	if err != nil {
		return frameSize, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 10 {
			return frameSize, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		if header.ErrorCode != 0 {
			return frameSize, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &frameSize)

	}

	return frameSize, nil
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
func (device *RS232V2Bricklet) GetSPITFPErrorCount() (errorCountAckChecksum uint32, errorCountMessageChecksum uint32, errorCountFrame uint32, errorCountOverflow uint32, err error) {
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
func (device *RS232V2Bricklet) SetBootloaderMode(mode BootloaderMode) (status BootloaderStatus, err error) {
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
func (device *RS232V2Bricklet) GetBootloaderMode() (mode BootloaderMode, err error) {
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
func (device *RS232V2Bricklet) SetWriteFirmwarePointer(pointer uint32) (err error) {
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
func (device *RS232V2Bricklet) WriteFirmware(data [64]uint8) (status uint8, err error) {
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
func (device *RS232V2Bricklet) SetStatusLEDConfig(config StatusLEDConfig) (err error) {
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
func (device *RS232V2Bricklet) GetStatusLEDConfig() (config StatusLEDConfig, err error) {
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
func (device *RS232V2Bricklet) GetChipTemperature() (temperature int16, err error) {
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
func (device *RS232V2Bricklet) Reset() (err error) {
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
func (device *RS232V2Bricklet) WriteUID(uid uint32) (err error) {
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
func (device *RS232V2Bricklet) ReadUID() (uid uint32, err error) {
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
func (device *RS232V2Bricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
