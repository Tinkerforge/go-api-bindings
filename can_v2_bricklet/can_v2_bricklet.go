/* ***********************************************************
 * This file was automatically generated on 2021-05-06.      *
 *                                                           *
 * Go Bindings Version 2.0.11                                *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


// Communicates with CAN bus devices.
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/CANV2_Bricklet_Go.html.
package can_v2_bricklet

import (
	"encoding/binary"
	"bytes"
	"fmt"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionWriteFrameLowLevel Function = 1
	FunctionReadFrameLowLevel Function = 2
	FunctionSetFrameReadCallbackConfiguration Function = 3
	FunctionGetFrameReadCallbackConfiguration Function = 4
	FunctionSetTransceiverConfiguration Function = 5
	FunctionGetTransceiverConfiguration Function = 6
	FunctionSetQueueConfigurationLowLevel Function = 7
	FunctionGetQueueConfigurationLowLevel Function = 8
	FunctionSetReadFilterConfiguration Function = 9
	FunctionGetReadFilterConfiguration Function = 10
	FunctionGetErrorLogLowLevel Function = 11
	FunctionSetCommunicationLEDConfig Function = 12
	FunctionGetCommunicationLEDConfig Function = 13
	FunctionSetErrorLEDConfig Function = 14
	FunctionGetErrorLEDConfig Function = 15
	FunctionSetFrameReadableCallbackConfiguration Function = 17
	FunctionGetFrameReadableCallbackConfiguration Function = 18
	FunctionSetErrorOccurredCallbackConfiguration Function = 20
	FunctionGetErrorOccurredCallbackConfiguration Function = 21
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
	FunctionCallbackFrameReadLowLevel Function = 16
	FunctionCallbackFrameReadable Function = 19
	FunctionCallbackErrorOccurred Function = 22
)

type FrameType = uint8

const (
	FrameTypeStandardData FrameType = 0
	FrameTypeStandardRemote FrameType = 1
	FrameTypeExtendedData FrameType = 2
	FrameTypeExtendedRemote FrameType = 3
)

type TransceiverMode = uint8

const (
	TransceiverModeNormal TransceiverMode = 0
	TransceiverModeLoopback TransceiverMode = 1
	TransceiverModeReadOnly TransceiverMode = 2
)

type FilterMode = uint8

const (
	FilterModeAcceptAll FilterMode = 0
	FilterModeMatchStandardOnly FilterMode = 1
	FilterModeMatchExtendedOnly FilterMode = 2
	FilterModeMatchStandardAndExtended FilterMode = 3
)

type TransceiverState = uint8

const (
	TransceiverStateActive TransceiverState = 0
	TransceiverStatePassive TransceiverState = 1
	TransceiverStateDisabled TransceiverState = 2
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
	ErrorLEDConfigShowTransceiverState ErrorLEDConfig = 3
	ErrorLEDConfigShowError ErrorLEDConfig = 4
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

type CANV2Bricklet struct {
	device Device
}
const DeviceIdentifier = 2107
const DeviceDisplayName = "CAN Bricklet 2.0"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (CANV2Bricklet, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{ 2,0,1 }, uid, &internalIPCon, 6, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return CANV2Bricklet{}, err
	}
	dev.ResponseExpected[FunctionWriteFrameLowLevel] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionReadFrameLowLevel] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetFrameReadCallbackConfiguration] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetFrameReadCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetTransceiverConfiguration] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetTransceiverConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetQueueConfigurationLowLevel] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetQueueConfigurationLowLevel] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetReadFilterConfiguration] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetReadFilterConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetErrorLogLowLevel] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetCommunicationLEDConfig] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetCommunicationLEDConfig] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetErrorLEDConfig] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetErrorLEDConfig] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetFrameReadableCallbackConfiguration] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetFrameReadableCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetErrorOccurredCallbackConfiguration] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetErrorOccurredCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue;
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
	return CANV2Bricklet{dev}, nil
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
func (device *CANV2Bricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *CANV2Bricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *CANV2Bricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *CANV2Bricklet) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// See RegisterFrameReadCallback
func (device *CANV2Bricklet) RegisterFrameReadLowLevelCallback(fn func(FrameType, uint32, uint8, [15]uint8)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 29 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var frameType FrameType
		var identifier uint32
		var dataLength uint8
		var dataData [15]uint8
		binary.Read(buf, binary.LittleEndian, &frameType)
		binary.Read(buf, binary.LittleEndian, &identifier)
		binary.Read(buf, binary.LittleEndian, &dataLength)
		copy(dataData[:], ByteSliceToUint8Slice(buf.Next(8 * 15/8)))
		fn(frameType, identifier, dataLength, dataData)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackFrameReadLowLevel), wrapper)
}

// Remove a registered Frame Read Low Level callback.
func (device *CANV2Bricklet) DeregisterFrameReadLowLevelCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackFrameReadLowLevel), registrationId)
}


// This callback is triggered if a data or remote frame was received by the CAN
// transceiver.
// 
// The ``identifier`` return value follows the identifier format described for
// WriteFrame.
// 
// For details on the ``data`` return value see ReadFrame.
// 
// A configurable read filter can be used to define which frames should be
// received by the CAN transceiver and put into the read queue (see
// SetReadFilterConfiguration).
// 
// To enable this callback, use SetFrameReadCallbackConfiguration.
// 
// Note
//  If reconstructing the value fails, the callback is triggered with nil for data.
func (device *CANV2Bricklet) RegisterFrameReadCallback(fn func(FrameType, uint32, []uint8)) uint64 {
	buf := make([]uint8, 0)
	wrapper := func(frameType FrameType, identifier uint32, dataLength uint8, dataData [15]uint8)  {
		if int(0) != len(buf) {
			if len(buf) != 0 {
				buf = nil
				fn(frameType, identifier, buf)
				buf = make([]uint8, 0)
			}
			if 0 != 0 {
				return
			}
		}
		toRead := MinU(uint64(dataLength-0), uint64(len(dataData[:])))
		buf = append(buf, dataData[:toRead]...)
		if len(buf) >= int(dataLength) {
			fn(frameType, identifier, buf)
			buf = make([]uint8, 0)
		}
	}
	return device.RegisterFrameReadLowLevelCallback(wrapper)
}

// Remove a registered Frame Read Low Level callback.
func (device *CANV2Bricklet) DeregisterFrameReadCallback(registrationId uint64) {
	device.DeregisterFrameReadLowLevelCallback(registrationId)
}


// This callback is triggered if a data or remote frame was received by the CAN
// transceiver. The received frame can be read with ReadFrame.
// If additional frames are received, but ReadFrame was not called yet, the callback
// will not trigger again.
// 
// A configurable read filter can be used to define which frames should be
// received by the CAN transceiver and put into the read queue (see
// SetReadFilterConfiguration).
// 
// To enable this callback, use SetFrameReadableCallbackConfiguration.
// 
// .. versionadded:: 2.0.3$nbsp;(Plugin)
func (device *CANV2Bricklet) RegisterFrameReadableCallback(fn func()) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 8 {
			return
		}
		
		
		
		fn()
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackFrameReadable), wrapper)
}

// Remove a registered Frame Readable callback.
func (device *CANV2Bricklet) DeregisterFrameReadableCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackFrameReadable), registrationId)
}


// This callback is triggered if any error occurred while writing, reading or transmitting CAN frames.
// 
// The callback is only triggered once until GetErrorLog is called. That function will return
// details abount the error(s) occurred.
// 
// To enable this callback, use SetErrorOccurredCallbackConfiguration.
// 
// .. versionadded:: 2.0.3$nbsp;(Plugin)
func (device *CANV2Bricklet) RegisterErrorOccurredCallback(fn func()) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 8 {
			return
		}
		
		
		
		fn()
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackErrorOccurred), wrapper)
}

// Remove a registered Error Occurred callback.
func (device *CANV2Bricklet) DeregisterErrorOccurredCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackErrorOccurred), registrationId)
}


// Writes a data or remote frame to the write queue to be transmitted over the
// CAN transceiver.
// 
// The Bricklet supports the standard 11-bit (CAN 2.0A) and the additional extended
// 29-bit (CAN 2.0B) identifiers. For standard frames the Bricklet uses bit 0 to 10
// from the ``identifier`` parameter as standard 11-bit identifier. For extended
// frames the Bricklet uses bit 0 to 28 from the ``identifier`` parameter as
// extended 29-bit identifier.
// 
// The ``data`` parameter can be up to 15 bytes long. For data frames up to 8 bytes
// will be used as the actual data. The length (DLC) field in the data or remote
// frame will be set to the actual length of the ``data`` parameter. This allows
// to transmit data and remote frames with excess length. For remote frames only
// the length of the ``data`` parameter is used. The actual ``data`` bytes are
// ignored.
// 
// Returns *true* if the frame was successfully added to the write queue. Returns
// *false* if the frame could not be added because write queue is already full or
// because the write buffer or the write backlog are configured with a size of
// zero (see SetQueueConfiguration).
// 
// The write queue can overflow if frames are written to it at a higher rate
// than the Bricklet can transmitted them over the CAN transceiver. This may
// happen if the CAN transceiver is configured as read-only or is using a low baud
// rate (see SetTransceiverConfiguration). It can also happen if the CAN
// bus is congested and the frame cannot be transmitted because it constantly loses
// arbitration or because the CAN transceiver is currently disabled due to a high
// write error level (see GetErrorLog).
//
// Associated constants:
//
//	* FrameTypeStandardData
//	* FrameTypeStandardRemote
//	* FrameTypeExtendedData
//	* FrameTypeExtendedRemote
func (device *CANV2Bricklet) WriteFrameLowLevel(frameType FrameType, identifier uint32, dataLength uint8, dataData [15]uint8) (success bool, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, frameType);
	binary.Write(&buf, binary.LittleEndian, identifier);
	binary.Write(&buf, binary.LittleEndian, dataLength);
	buf.Write(Uint8SliceToByteSlice(dataData[:]))

	resultBytes, err := device.device.Get(uint8(FunctionWriteFrameLowLevel), buf.Bytes())
	if err != nil {
		return success, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return success, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return success, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &success)

	}

	return success, nil
}

// Writes a data or remote frame to the write queue to be transmitted over the
// CAN transceiver.
// 
// The Bricklet supports the standard 11-bit (CAN 2.0A) and the additional extended
// 29-bit (CAN 2.0B) identifiers. For standard frames the Bricklet uses bit 0 to 10
// from the ``identifier`` parameter as standard 11-bit identifier. For extended
// frames the Bricklet uses bit 0 to 28 from the ``identifier`` parameter as
// extended 29-bit identifier.
// 
// The ``data`` parameter can be up to 15 bytes long. For data frames up to 8 bytes
// will be used as the actual data. The length (DLC) field in the data or remote
// frame will be set to the actual length of the ``data`` parameter. This allows
// to transmit data and remote frames with excess length. For remote frames only
// the length of the ``data`` parameter is used. The actual ``data`` bytes are
// ignored.
// 
// Returns *true* if the frame was successfully added to the write queue. Returns
// *false* if the frame could not be added because write queue is already full or
// because the write buffer or the write backlog are configured with a size of
// zero (see SetQueueConfiguration).
// 
// The write queue can overflow if frames are written to it at a higher rate
// than the Bricklet can transmitted them over the CAN transceiver. This may
// happen if the CAN transceiver is configured as read-only or is using a low baud
// rate (see SetTransceiverConfiguration). It can also happen if the CAN
// bus is congested and the frame cannot be transmitted because it constantly loses
// arbitration or because the CAN transceiver is currently disabled due to a high
// write error level (see GetErrorLog).
	func (device *CANV2Bricklet) WriteFrame(frameType FrameType, identifier uint32, data []uint8) (success bool, err error) {
		lowLevelResult, err := device.device.SetHighLevel(func(dataLength uint64, chunkOffset uint64, dataData []byte) (LowLevelWriteResult, error) {
			arr := [15]uint8{}
			copy(arr[:], ByteSliceToUint8Slice(dataData))

			success, err := device.WriteFrameLowLevel(frameType, identifier, uint8(dataLength), arr)

			var lowLevelResults bytes.Buffer
			binary.Write(&lowLevelResults, binary.LittleEndian, success);

			return LowLevelWriteResult{
				uint64(15),
				lowLevelResults.Bytes()}, err
		}, 0, 8, 120, Uint8SliceToByteSlice(data))

		if err != nil {
			return
		}

		resultBuf := bytes.NewBuffer(lowLevelResult.Result)

		binary.Read(resultBuf, binary.LittleEndian, &success)
		
		return
	}

// Tries to read the next data or remote frame from the read queue and returns it.
// If a frame was successfully read, then the ``success`` return value is set to
// *true* and the other return values contain the frame. If the read queue is
// empty and no frame could be read, then the ``success`` return value is set to
// *false* and the other return values contain invalid data.
// 
// The ``identifier`` return value follows the identifier format described for
// WriteFrame.
// 
// The ``data`` return value can be up to 15 bytes long. For data frames up to the
// first 8 bytes are the actual received data. All bytes after the 8th byte are
// always zero and only there to indicate the length of a data or remote frame
// with excess length. For remote frames the length of the ``data`` return value
// represents the requested length. The actual ``data`` bytes are always zero.
// 
// A configurable read filter can be used to define which frames should be
// received by the CAN transceiver and put into the read queue (see
// SetReadFilterConfiguration).
// 
// Instead of polling with this function, you can also use callbacks. See the
// SetFrameReadCallbackConfiguration function and the RegisterFrameReadCallback
// callback.
//
// Associated constants:
//
//	* FrameTypeStandardData
//	* FrameTypeStandardRemote
//	* FrameTypeExtendedData
//	* FrameTypeExtendedRemote
func (device *CANV2Bricklet) ReadFrameLowLevel() (success bool, frameType FrameType, identifier uint32, dataLength uint8, dataData [15]uint8, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionReadFrameLowLevel), buf.Bytes())
	if err != nil {
		return success, frameType, identifier, dataLength, dataData, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 30 {
			return success, frameType, identifier, dataLength, dataData, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 30)
		}

		if header.ErrorCode != 0 {
			return success, frameType, identifier, dataLength, dataData, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &success)
		binary.Read(resultBuf, binary.LittleEndian, &frameType)
		binary.Read(resultBuf, binary.LittleEndian, &identifier)
		binary.Read(resultBuf, binary.LittleEndian, &dataLength)
		copy(dataData[:], ByteSliceToUint8Slice(resultBuf.Next(8 * 15/8)))

	}

	return success, frameType, identifier, dataLength, dataData, nil
}

// Tries to read the next data or remote frame from the read queue and returns it.
// If a frame was successfully read, then the ``success`` return value is set to
// *true* and the other return values contain the frame. If the read queue is
// empty and no frame could be read, then the ``success`` return value is set to
// *false* and the other return values contain invalid data.
// 
// The ``identifier`` return value follows the identifier format described for
// WriteFrame.
// 
// The ``data`` return value can be up to 15 bytes long. For data frames up to the
// first 8 bytes are the actual received data. All bytes after the 8th byte are
// always zero and only there to indicate the length of a data or remote frame
// with excess length. For remote frames the length of the ``data`` return value
// represents the requested length. The actual ``data`` bytes are always zero.
// 
// A configurable read filter can be used to define which frames should be
// received by the CAN transceiver and put into the read queue (see
// SetReadFilterConfiguration).
// 
// Instead of polling with this function, you can also use callbacks. See the
// SetFrameReadCallbackConfiguration function and the RegisterFrameReadCallback
// callback.
	func (device *CANV2Bricklet) ReadFrame() (data []uint8, success bool, frameType FrameType, identifier uint32, err error) {
		buf, result, err := device.device.GetHighLevel(func() (LowLevelResult, error) {
			success, frameType, identifier, dataLength, dataData, err := device.ReadFrameLowLevel()

			if err != nil {
				return LowLevelResult{}, err
			}

			var lowLevelResults bytes.Buffer
			binary.Write(&lowLevelResults, binary.LittleEndian, success);
	binary.Write(&lowLevelResults, binary.LittleEndian, frameType);
	binary.Write(&lowLevelResults, binary.LittleEndian, identifier);

			return LowLevelResult{
				uint64(dataLength),
				uint64(dataLength),
				Uint8SliceToByteSlice(dataData[:]),
				lowLevelResults.Bytes()}, nil
		},
			1,
			8)
		if err != nil {
			return ByteSliceToUint8Slice(buf), success, frameType, identifier, err
		}
		resultBuf := bytes.NewBuffer(result)
		binary.Read(resultBuf, binary.LittleEndian, &success)
	binary.Read(resultBuf, binary.LittleEndian, &frameType)
	binary.Read(resultBuf, binary.LittleEndian, &identifier)
		return ByteSliceToUint8Slice(buf), success, frameType, identifier, nil
	}

// Enables and disables the RegisterFrameReadCallback callback.
// 
// By default the callback is disabled. Enabling this callback will disable the RegisterFrameReadableCallback callback.
func (device *CANV2Bricklet) SetFrameReadCallbackConfiguration(enabled bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, enabled);

	resultBytes, err := device.device.Set(uint8(FunctionSetFrameReadCallbackConfiguration), buf.Bytes())
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

// Returns *true* if the RegisterFrameReadCallback callback is enabled, *false* otherwise.
func (device *CANV2Bricklet) GetFrameReadCallbackConfiguration() (enabled bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetFrameReadCallbackConfiguration), buf.Bytes())
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

// Sets the transceiver configuration for the CAN bus communication.
// 
// The CAN transceiver has three different modes:
// 
// * Normal: Reads from and writes to the CAN bus and performs active bus
//   error detection and acknowledgement.
// * Loopback: All reads and writes are performed internally. The transceiver
//   is disconnected from the actual CAN bus.
// * Read-Only: Only reads from the CAN bus, but does neither active bus error
//   detection nor acknowledgement. Only the receiving part of the transceiver
//   is connected to the CAN bus.
//
// Associated constants:
//
//	* TransceiverModeNormal
//	* TransceiverModeLoopback
//	* TransceiverModeReadOnly
func (device *CANV2Bricklet) SetTransceiverConfiguration(baudRate uint32, samplePoint uint16, transceiverMode TransceiverMode) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, baudRate);
	binary.Write(&buf, binary.LittleEndian, samplePoint);
	binary.Write(&buf, binary.LittleEndian, transceiverMode);

	resultBytes, err := device.device.Set(uint8(FunctionSetTransceiverConfiguration), buf.Bytes())
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

// Returns the configuration as set by SetTransceiverConfiguration.
//
// Associated constants:
//
//	* TransceiverModeNormal
//	* TransceiverModeLoopback
//	* TransceiverModeReadOnly
func (device *CANV2Bricklet) GetTransceiverConfiguration() (baudRate uint32, samplePoint uint16, transceiverMode TransceiverMode, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetTransceiverConfiguration), buf.Bytes())
	if err != nil {
		return baudRate, samplePoint, transceiverMode, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 15 {
			return baudRate, samplePoint, transceiverMode, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 15)
		}

		if header.ErrorCode != 0 {
			return baudRate, samplePoint, transceiverMode, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &baudRate)
		binary.Read(resultBuf, binary.LittleEndian, &samplePoint)
		binary.Read(resultBuf, binary.LittleEndian, &transceiverMode)

	}

	return baudRate, samplePoint, transceiverMode, nil
}

// Sets the write and read queue configuration.
// 
// The CAN transceiver has 32 buffers in total in hardware for transmitting and
// receiving frames. Additionally, the Bricklet has a backlog for 768 frames in
// total in software. The buffers and the backlog can be freely assigned to the
// write and read queues.
// 
// WriteFrame writes a frame into the write backlog. The Bricklet moves
// the frame from the backlog into a free write buffer. The CAN transceiver then
// transmits the frame from the write buffer to the CAN bus. If there are no
// write buffers (``write_buffer_size`` is zero) or there is no write backlog
// (``write_backlog_size`` is zero) then no frames can be transmitted and
// WriteFrame returns always *false*.
// 
// The CAN transceiver receives a frame from the CAN bus and stores it into a
// free read buffer. The Bricklet moves the frame from the read buffer into the
// read backlog. ReadFrame reads the frame from the read backlog and
// returns it. If there are no read buffers (``read_buffer_sizes`` is empty) or
// there is no read backlog (``read_backlog_size`` is zero) then no frames can be
// received and ReadFrame returns always *false*.
// 
// There can be multiple read buffers, because the CAN transceiver cannot receive
// data and remote frames into the same read buffer. A positive read buffer size
// represents a data frame read buffer and a negative read buffer size represents
// a remote frame read buffer. A read buffer size of zero is not allowed. By
// default the first read buffer is configured for data frames and the second read
// buffer is configured for remote frame. There can be up to 32 different read
// buffers, assuming that no write buffer is used. Each read buffer has its own
// filter configuration (see SetReadFilterConfiguration).
// 
// A valid queue configuration fulfills these conditions::
// 
//  write_buffer_size + abs(read_buffer_size_0) + abs(read_buffer_size_1) + ... + abs(read_buffer_size_31) <= 32
//  write_backlog_size + read_backlog_size <= 768
// 
// The write buffer timeout has three different modes that define how a failed
// frame transmission should be handled:
// 
// * Single-Shot (< 0): Only one transmission attempt will be made. If the
//   transmission fails then the frame is discarded.
// * Infinite (= 0): Infinite transmission attempts will be made. The frame will
//   never be discarded.
// * Milliseconds (> 0): A limited number of transmission attempts will be made.
//   If the frame could not be transmitted successfully after the configured
//   number of milliseconds then the frame is discarded.
// 
// The current content of the queues is lost when this function is called.
func (device *CANV2Bricklet) SetQueueConfigurationLowLevel(writeBufferSize uint8, writeBufferTimeout int32, writeBacklogSize uint16, readBufferSizesLength uint8, readBufferSizesData [32]int8, readBacklogSize uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, writeBufferSize);
	binary.Write(&buf, binary.LittleEndian, writeBufferTimeout);
	binary.Write(&buf, binary.LittleEndian, writeBacklogSize);
	binary.Write(&buf, binary.LittleEndian, readBufferSizesLength);
	buf.Write(Int8SliceToByteSlice(readBufferSizesData[:]))
	binary.Write(&buf, binary.LittleEndian, readBacklogSize);

	resultBytes, err := device.device.Set(uint8(FunctionSetQueueConfigurationLowLevel), buf.Bytes())
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

// Sets the write and read queue configuration.
// 
// The CAN transceiver has 32 buffers in total in hardware for transmitting and
// receiving frames. Additionally, the Bricklet has a backlog for 768 frames in
// total in software. The buffers and the backlog can be freely assigned to the
// write and read queues.
// 
// WriteFrame writes a frame into the write backlog. The Bricklet moves
// the frame from the backlog into a free write buffer. The CAN transceiver then
// transmits the frame from the write buffer to the CAN bus. If there are no
// write buffers (``write_buffer_size`` is zero) or there is no write backlog
// (``write_backlog_size`` is zero) then no frames can be transmitted and
// WriteFrame returns always *false*.
// 
// The CAN transceiver receives a frame from the CAN bus and stores it into a
// free read buffer. The Bricklet moves the frame from the read buffer into the
// read backlog. ReadFrame reads the frame from the read backlog and
// returns it. If there are no read buffers (``read_buffer_sizes`` is empty) or
// there is no read backlog (``read_backlog_size`` is zero) then no frames can be
// received and ReadFrame returns always *false*.
// 
// There can be multiple read buffers, because the CAN transceiver cannot receive
// data and remote frames into the same read buffer. A positive read buffer size
// represents a data frame read buffer and a negative read buffer size represents
// a remote frame read buffer. A read buffer size of zero is not allowed. By
// default the first read buffer is configured for data frames and the second read
// buffer is configured for remote frame. There can be up to 32 different read
// buffers, assuming that no write buffer is used. Each read buffer has its own
// filter configuration (see SetReadFilterConfiguration).
// 
// A valid queue configuration fulfills these conditions::
// 
//  write_buffer_size + abs(read_buffer_size_0) + abs(read_buffer_size_1) + ... + abs(read_buffer_size_31) <= 32
//  write_backlog_size + read_backlog_size <= 768
// 
// The write buffer timeout has three different modes that define how a failed
// frame transmission should be handled:
// 
// * Single-Shot (< 0): Only one transmission attempt will be made. If the
//   transmission fails then the frame is discarded.
// * Infinite (= 0): Infinite transmission attempts will be made. The frame will
//   never be discarded.
// * Milliseconds (> 0): A limited number of transmission attempts will be made.
//   If the frame could not be transmitted successfully after the configured
//   number of milliseconds then the frame is discarded.
// 
// The current content of the queues is lost when this function is called.
	func (device *CANV2Bricklet) SetQueueConfiguration(writeBufferSize uint8, writeBufferTimeout int32, writeBacklogSize uint16, readBacklogSize uint16, readBufferSizes []int8) (err error) {
		_, err = device.device.SetHighLevel(func(readBufferSizesLength uint64, chunkOffset uint64, readBufferSizesData []byte) (LowLevelWriteResult, error) {
			arr := [32]int8{}
			copy(arr[:], ByteSliceToInt8Slice(readBufferSizesData))

			err := device.SetQueueConfigurationLowLevel(writeBufferSize, writeBufferTimeout, writeBacklogSize, uint8(readBufferSizesLength), arr, readBacklogSize)

			var lowLevelResults bytes.Buffer
			

			return LowLevelWriteResult{
				uint64(32),
				lowLevelResults.Bytes()}, err
		}, 2, 8, 256, Int8SliceToByteSlice(readBufferSizes))

		if err != nil {
			return
		}

		
		
		
		return
	}

// Returns the queue configuration as set by SetQueueConfiguration.
func (device *CANV2Bricklet) GetQueueConfigurationLowLevel() (writeBufferSize uint8, writeBufferTimeout int32, writeBacklogSize uint16, readBufferSizesLength uint8, readBufferSizesData [32]int8, readBacklogSize uint16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetQueueConfigurationLowLevel), buf.Bytes())
	if err != nil {
		return writeBufferSize, writeBufferTimeout, writeBacklogSize, readBufferSizesLength, readBufferSizesData, readBacklogSize, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 50 {
			return writeBufferSize, writeBufferTimeout, writeBacklogSize, readBufferSizesLength, readBufferSizesData, readBacklogSize, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 50)
		}

		if header.ErrorCode != 0 {
			return writeBufferSize, writeBufferTimeout, writeBacklogSize, readBufferSizesLength, readBufferSizesData, readBacklogSize, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &writeBufferSize)
		binary.Read(resultBuf, binary.LittleEndian, &writeBufferTimeout)
		binary.Read(resultBuf, binary.LittleEndian, &writeBacklogSize)
		binary.Read(resultBuf, binary.LittleEndian, &readBufferSizesLength)
		copy(readBufferSizesData[:], ByteSliceToInt8Slice(resultBuf.Next(8 * 32/8)))
		binary.Read(resultBuf, binary.LittleEndian, &readBacklogSize)

	}

	return writeBufferSize, writeBufferTimeout, writeBacklogSize, readBufferSizesLength, readBufferSizesData, readBacklogSize, nil
}

// Returns the queue configuration as set by SetQueueConfiguration.
	func (device *CANV2Bricklet) GetQueueConfiguration() (readBufferSizes []int8, writeBufferSize uint8, writeBufferTimeout int32, writeBacklogSize uint16, readBacklogSize uint16, err error) {
		buf, result, err := device.device.GetHighLevel(func() (LowLevelResult, error) {
			writeBufferSize, writeBufferTimeout, writeBacklogSize, readBufferSizesLength, readBufferSizesData, readBacklogSize, err := device.GetQueueConfigurationLowLevel()

			if err != nil {
				return LowLevelResult{}, err
			}

			var lowLevelResults bytes.Buffer
			binary.Write(&lowLevelResults, binary.LittleEndian, writeBufferSize);
	binary.Write(&lowLevelResults, binary.LittleEndian, writeBufferTimeout);
	binary.Write(&lowLevelResults, binary.LittleEndian, writeBacklogSize);
	binary.Write(&lowLevelResults, binary.LittleEndian, readBacklogSize);

			return LowLevelResult{
				uint64(readBufferSizesLength),
				uint64(readBufferSizesLength),
				Int8SliceToByteSlice(readBufferSizesData[:]),
				lowLevelResults.Bytes()}, nil
		},
			3,
			8)
		if err != nil {
			return ByteSliceToInt8Slice(buf), writeBufferSize, writeBufferTimeout, writeBacklogSize, readBacklogSize, err
		}
		resultBuf := bytes.NewBuffer(result)
		binary.Read(resultBuf, binary.LittleEndian, &writeBufferSize)
	binary.Read(resultBuf, binary.LittleEndian, &writeBufferTimeout)
	binary.Read(resultBuf, binary.LittleEndian, &writeBacklogSize)
	binary.Read(resultBuf, binary.LittleEndian, &readBacklogSize)
		return ByteSliceToInt8Slice(buf), writeBufferSize, writeBufferTimeout, writeBacklogSize, readBacklogSize, nil
	}

// Set the read filter configuration for the given read buffer index. This can be
// used to define which frames should be received by the CAN transceiver and put
// into the read buffer.
// 
// The read filter has four different modes that define if and how the filter mask
// and the filter identifier are applied:
// 
// * Accept-All: All frames are received.
// * Match-Standard-Only: Only standard frames with a matching identifier are
//   received.
// * Match-Extended-Only: Only extended frames with a matching identifier are
//   received.
// * Match-Standard-And-Extended: Standard and extended frames with a matching
//   identifier are received.
// 
// The filter mask and filter identifier are used as bit masks. Their usage
// depends on the mode:
// 
// * Accept-All: Mask and identifier are ignored.
// * Match-Standard-Only: Bit 0 to 10 (11 bits) of filter mask and filter
//   identifier are used to match the 11-bit identifier of standard frames.
// * Match-Extended-Only: Bit 0 to 28 (29 bits) of filter mask and filter
//   identifier are used to match the 29-bit identifier of extended frames.
// * Match-Standard-And-Extended: Bit 18 to 28 (11 bits) of filter mask and filter
//   identifier are used to match the 11-bit identifier of standard frames, bit 0
//   to 17 (18 bits) are ignored in this case. Bit 0 to 28 (29 bits) of filter
//   mask and filter identifier are used to match the 29-bit identifier of extended
//   frames.
// 
// The filter mask and filter identifier are applied in this way: The filter mask
// is used to select the frame identifier bits that should be compared to the
// corresponding filter identifier bits. All unselected bits are automatically
// accepted. All selected bits have to match the filter identifier to be accepted.
// If all bits for the selected mode are accepted then the frame is accepted and
// is added to the read buffer.
// 
//  Filter Mask Bit| Filter Identifier Bit| Frame Identifier Bit| Result
//  --- | --- | --- | --- 
//  0| X| X| Accept
//  1| 0| 0| Accept
//  1| 0| 1| Reject
//  1| 1| 0| Reject
//  1| 1| 1| Accept
// 
// For example, to receive standard frames with identifier 0x123 only, the mode
// can be set to Match-Standard-Only with 0x7FF as mask and 0x123 as identifier.
// The mask of 0x7FF selects all 11 identifier bits for matching so that the
// identifier has to be exactly 0x123 to be accepted.
// 
// To accept identifier 0x123 and identifier 0x456 at the same time, just set
// filter 2 to 0x456 and keep mask and filter 1 unchanged.
// 
// There can be up to 32 different read filters configured at the same time,
// because there can be up to 32 read buffer (see SetQueueConfiguration).
// 
// The default mode is accept-all for all read buffers.
//
// Associated constants:
//
//	* FilterModeAcceptAll
//	* FilterModeMatchStandardOnly
//	* FilterModeMatchExtendedOnly
//	* FilterModeMatchStandardAndExtended
func (device *CANV2Bricklet) SetReadFilterConfiguration(bufferIndex uint8, filterMode FilterMode, filterMask uint32, filterIdentifier uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, bufferIndex);
	binary.Write(&buf, binary.LittleEndian, filterMode);
	binary.Write(&buf, binary.LittleEndian, filterMask);
	binary.Write(&buf, binary.LittleEndian, filterIdentifier);

	resultBytes, err := device.device.Set(uint8(FunctionSetReadFilterConfiguration), buf.Bytes())
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

// Returns the read filter configuration as set by SetReadFilterConfiguration.
//
// Associated constants:
//
//	* FilterModeAcceptAll
//	* FilterModeMatchStandardOnly
//	* FilterModeMatchExtendedOnly
//	* FilterModeMatchStandardAndExtended
func (device *CANV2Bricklet) GetReadFilterConfiguration(bufferIndex uint8) (filterMode FilterMode, filterMask uint32, filterIdentifier uint32, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, bufferIndex);

	resultBytes, err := device.device.Get(uint8(FunctionGetReadFilterConfiguration), buf.Bytes())
	if err != nil {
		return filterMode, filterMask, filterIdentifier, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 17 {
			return filterMode, filterMask, filterIdentifier, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 17)
		}

		if header.ErrorCode != 0 {
			return filterMode, filterMask, filterIdentifier, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &filterMode)
		binary.Read(resultBuf, binary.LittleEndian, &filterMask)
		binary.Read(resultBuf, binary.LittleEndian, &filterIdentifier)

	}

	return filterMode, filterMask, filterIdentifier, nil
}

// Returns information about different kinds of errors.
// 
// The write and read error levels indicate the current level of stuffing, form,
// acknowledgement, bit and checksum errors during CAN bus write and read
// operations. For each of this error kinds there is also an individual counter.
// 
// When the write error level extends 255 then the CAN transceiver gets disabled
// and no frames can be transmitted or received anymore. The CAN transceiver will
// automatically be activated again after the CAN bus is idle for a while.
// 
// The write buffer timeout, read buffer and backlog overflow counts represents the
// number of these errors:
// 
// * A write buffer timeout occurs if a frame could not be transmitted before the
//   configured write buffer timeout expired (see SetQueueConfiguration).
// * A read buffer overflow occurs if a read buffer of the CAN transceiver
//   still contains the last received frame when the next frame arrives. In this
//   case the last received frame is lost. This happens if the CAN transceiver
//   receives more frames than the Bricklet can handle. Using the read filter
//   (see SetReadFilterConfiguration) can help to reduce the amount of
//   received frames. This count is not exact, but a lower bound, because the
//   Bricklet might not able detect all overflows if they occur in rapid succession.
// * A read backlog overflow occurs if the read backlog of the Bricklet is already
//   full when the next frame should be read from a read buffer of the CAN
//   transceiver. In this case the frame in the read buffer is lost. This
//   happens if the CAN transceiver receives more frames to be added to the read
//   backlog than are removed from the read backlog using the ReadFrame
//   function. Using the RegisterFrameReadCallback callback ensures that the read backlog
//   can not overflow.
// 
// The read buffer overflow counter counts the overflows of all configured read
// buffers. Which read buffer exactly suffered from an overflow can be figured
// out from the read buffer overflow occurrence list
// (``read_buffer_overflow_error_occurred``). Reading the error log clears the
// occurence list.
//
// Associated constants:
//
//	* TransceiverStateActive
//	* TransceiverStatePassive
//	* TransceiverStateDisabled
func (device *CANV2Bricklet) GetErrorLogLowLevel() (transceiverState TransceiverState, transceiverWriteErrorLevel uint8, transceiverReadErrorLevel uint8, transceiverStuffingErrorCount uint32, transceiverFormatErrorCount uint32, transceiverACKErrorCount uint32, transceiverBit1ErrorCount uint32, transceiverBit0ErrorCount uint32, transceiverCRCErrorCount uint32, writeBufferTimeoutErrorCount uint32, readBufferOverflowErrorCount uint32, readBufferOverflowErrorOccurredLength uint8, readBufferOverflowErrorOccurredData [32]bool, readBacklogOverflowErrorCount uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetErrorLogLowLevel), buf.Bytes())
	if err != nil {
		return transceiverState, transceiverWriteErrorLevel, transceiverReadErrorLevel, transceiverStuffingErrorCount, transceiverFormatErrorCount, transceiverACKErrorCount, transceiverBit1ErrorCount, transceiverBit0ErrorCount, transceiverCRCErrorCount, writeBufferTimeoutErrorCount, readBufferOverflowErrorCount, readBufferOverflowErrorOccurredLength, readBufferOverflowErrorOccurredData, readBacklogOverflowErrorCount, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 52 {
			return transceiverState, transceiverWriteErrorLevel, transceiverReadErrorLevel, transceiverStuffingErrorCount, transceiverFormatErrorCount, transceiverACKErrorCount, transceiverBit1ErrorCount, transceiverBit0ErrorCount, transceiverCRCErrorCount, writeBufferTimeoutErrorCount, readBufferOverflowErrorCount, readBufferOverflowErrorOccurredLength, readBufferOverflowErrorOccurredData, readBacklogOverflowErrorCount, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 52)
		}

		if header.ErrorCode != 0 {
			return transceiverState, transceiverWriteErrorLevel, transceiverReadErrorLevel, transceiverStuffingErrorCount, transceiverFormatErrorCount, transceiverACKErrorCount, transceiverBit1ErrorCount, transceiverBit0ErrorCount, transceiverCRCErrorCount, writeBufferTimeoutErrorCount, readBufferOverflowErrorCount, readBufferOverflowErrorOccurredLength, readBufferOverflowErrorOccurredData, readBacklogOverflowErrorCount, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &transceiverState)
		binary.Read(resultBuf, binary.LittleEndian, &transceiverWriteErrorLevel)
		binary.Read(resultBuf, binary.LittleEndian, &transceiverReadErrorLevel)
		binary.Read(resultBuf, binary.LittleEndian, &transceiverStuffingErrorCount)
		binary.Read(resultBuf, binary.LittleEndian, &transceiverFormatErrorCount)
		binary.Read(resultBuf, binary.LittleEndian, &transceiverACKErrorCount)
		binary.Read(resultBuf, binary.LittleEndian, &transceiverBit1ErrorCount)
		binary.Read(resultBuf, binary.LittleEndian, &transceiverBit0ErrorCount)
		binary.Read(resultBuf, binary.LittleEndian, &transceiverCRCErrorCount)
		binary.Read(resultBuf, binary.LittleEndian, &writeBufferTimeoutErrorCount)
		binary.Read(resultBuf, binary.LittleEndian, &readBufferOverflowErrorCount)
		binary.Read(resultBuf, binary.LittleEndian, &readBufferOverflowErrorOccurredLength)
		copy(readBufferOverflowErrorOccurredData[:], ByteSliceToBoolSlice(resultBuf.Next(1 * 32/8)))
		binary.Read(resultBuf, binary.LittleEndian, &readBacklogOverflowErrorCount)

	}

	return transceiverState, transceiverWriteErrorLevel, transceiverReadErrorLevel, transceiverStuffingErrorCount, transceiverFormatErrorCount, transceiverACKErrorCount, transceiverBit1ErrorCount, transceiverBit0ErrorCount, transceiverCRCErrorCount, writeBufferTimeoutErrorCount, readBufferOverflowErrorCount, readBufferOverflowErrorOccurredLength, readBufferOverflowErrorOccurredData, readBacklogOverflowErrorCount, nil
}

// Returns information about different kinds of errors.
// 
// The write and read error levels indicate the current level of stuffing, form,
// acknowledgement, bit and checksum errors during CAN bus write and read
// operations. For each of this error kinds there is also an individual counter.
// 
// When the write error level extends 255 then the CAN transceiver gets disabled
// and no frames can be transmitted or received anymore. The CAN transceiver will
// automatically be activated again after the CAN bus is idle for a while.
// 
// The write buffer timeout, read buffer and backlog overflow counts represents the
// number of these errors:
// 
// * A write buffer timeout occurs if a frame could not be transmitted before the
//   configured write buffer timeout expired (see SetQueueConfiguration).
// * A read buffer overflow occurs if a read buffer of the CAN transceiver
//   still contains the last received frame when the next frame arrives. In this
//   case the last received frame is lost. This happens if the CAN transceiver
//   receives more frames than the Bricklet can handle. Using the read filter
//   (see SetReadFilterConfiguration) can help to reduce the amount of
//   received frames. This count is not exact, but a lower bound, because the
//   Bricklet might not able detect all overflows if they occur in rapid succession.
// * A read backlog overflow occurs if the read backlog of the Bricklet is already
//   full when the next frame should be read from a read buffer of the CAN
//   transceiver. In this case the frame in the read buffer is lost. This
//   happens if the CAN transceiver receives more frames to be added to the read
//   backlog than are removed from the read backlog using the ReadFrame
//   function. Using the RegisterFrameReadCallback callback ensures that the read backlog
//   can not overflow.
// 
// The read buffer overflow counter counts the overflows of all configured read
// buffers. Which read buffer exactly suffered from an overflow can be figured
// out from the read buffer overflow occurrence list
// (``read_buffer_overflow_error_occurred``). Reading the error log clears the
// occurence list.
	func (device *CANV2Bricklet) GetErrorLog() (readBufferOverflowErrorOccurred []bool, transceiverState TransceiverState, transceiverWriteErrorLevel uint8, transceiverReadErrorLevel uint8, transceiverStuffingErrorCount uint32, transceiverFormatErrorCount uint32, transceiverACKErrorCount uint32, transceiverBit1ErrorCount uint32, transceiverBit0ErrorCount uint32, transceiverCRCErrorCount uint32, writeBufferTimeoutErrorCount uint32, readBufferOverflowErrorCount uint32, readBacklogOverflowErrorCount uint32, err error) {
		buf, result, err := device.device.GetHighLevel(func() (LowLevelResult, error) {
			transceiverState, transceiverWriteErrorLevel, transceiverReadErrorLevel, transceiverStuffingErrorCount, transceiverFormatErrorCount, transceiverACKErrorCount, transceiverBit1ErrorCount, transceiverBit0ErrorCount, transceiverCRCErrorCount, writeBufferTimeoutErrorCount, readBufferOverflowErrorCount, readBufferOverflowErrorOccurredLength, readBufferOverflowErrorOccurredData, readBacklogOverflowErrorCount, err := device.GetErrorLogLowLevel()

			if err != nil {
				return LowLevelResult{}, err
			}

			var lowLevelResults bytes.Buffer
			binary.Write(&lowLevelResults, binary.LittleEndian, transceiverState);
	binary.Write(&lowLevelResults, binary.LittleEndian, transceiverWriteErrorLevel);
	binary.Write(&lowLevelResults, binary.LittleEndian, transceiverReadErrorLevel);
	binary.Write(&lowLevelResults, binary.LittleEndian, transceiverStuffingErrorCount);
	binary.Write(&lowLevelResults, binary.LittleEndian, transceiverFormatErrorCount);
	binary.Write(&lowLevelResults, binary.LittleEndian, transceiverACKErrorCount);
	binary.Write(&lowLevelResults, binary.LittleEndian, transceiverBit1ErrorCount);
	binary.Write(&lowLevelResults, binary.LittleEndian, transceiverBit0ErrorCount);
	binary.Write(&lowLevelResults, binary.LittleEndian, transceiverCRCErrorCount);
	binary.Write(&lowLevelResults, binary.LittleEndian, writeBufferTimeoutErrorCount);
	binary.Write(&lowLevelResults, binary.LittleEndian, readBufferOverflowErrorCount);
	binary.Write(&lowLevelResults, binary.LittleEndian, readBacklogOverflowErrorCount);

			return LowLevelResult{
				uint64(readBufferOverflowErrorOccurredLength),
				uint64(readBufferOverflowErrorOccurredLength),
				BoolSliceToByteSlice(readBufferOverflowErrorOccurredData[:]),
				lowLevelResults.Bytes()}, nil
		},
			4,
			1)
		if err != nil {
			return ByteSliceToBoolSlice(buf), transceiverState, transceiverWriteErrorLevel, transceiverReadErrorLevel, transceiverStuffingErrorCount, transceiverFormatErrorCount, transceiverACKErrorCount, transceiverBit1ErrorCount, transceiverBit0ErrorCount, transceiverCRCErrorCount, writeBufferTimeoutErrorCount, readBufferOverflowErrorCount, readBacklogOverflowErrorCount, err
		}
		resultBuf := bytes.NewBuffer(result)
		binary.Read(resultBuf, binary.LittleEndian, &transceiverState)
	binary.Read(resultBuf, binary.LittleEndian, &transceiverWriteErrorLevel)
	binary.Read(resultBuf, binary.LittleEndian, &transceiverReadErrorLevel)
	binary.Read(resultBuf, binary.LittleEndian, &transceiverStuffingErrorCount)
	binary.Read(resultBuf, binary.LittleEndian, &transceiverFormatErrorCount)
	binary.Read(resultBuf, binary.LittleEndian, &transceiverACKErrorCount)
	binary.Read(resultBuf, binary.LittleEndian, &transceiverBit1ErrorCount)
	binary.Read(resultBuf, binary.LittleEndian, &transceiverBit0ErrorCount)
	binary.Read(resultBuf, binary.LittleEndian, &transceiverCRCErrorCount)
	binary.Read(resultBuf, binary.LittleEndian, &writeBufferTimeoutErrorCount)
	binary.Read(resultBuf, binary.LittleEndian, &readBufferOverflowErrorCount)
	binary.Read(resultBuf, binary.LittleEndian, &readBacklogOverflowErrorCount)
		return ByteSliceToBoolSlice(buf), transceiverState, transceiverWriteErrorLevel, transceiverReadErrorLevel, transceiverStuffingErrorCount, transceiverFormatErrorCount, transceiverACKErrorCount, transceiverBit1ErrorCount, transceiverBit0ErrorCount, transceiverCRCErrorCount, writeBufferTimeoutErrorCount, readBufferOverflowErrorCount, readBacklogOverflowErrorCount, nil
	}

// Sets the communication LED configuration. By default the LED shows
// CAN-Bus traffic, it flickers once for every 40 transmitted or received frames.
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
func (device *CANV2Bricklet) SetCommunicationLEDConfig(config CommunicationLEDConfig) (err error) {
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
func (device *CANV2Bricklet) GetCommunicationLEDConfig() (config CommunicationLEDConfig, err error) {
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
// By default (show-transceiver-state) the error LED turns on if the CAN
// transceiver is passive or disabled state (see GetErrorLog). If
// the CAN transceiver is in active state the LED turns off.
// 
// If the LED is configured as show-error then the error LED turns on if any error
// occurs. If you call this function with the show-error option again, the LED will
// turn off until the next error occurs.
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
//	* ErrorLEDConfigShowTransceiverState
//	* ErrorLEDConfigShowError
func (device *CANV2Bricklet) SetErrorLEDConfig(config ErrorLEDConfig) (err error) {
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
//	* ErrorLEDConfigShowTransceiverState
//	* ErrorLEDConfigShowError
func (device *CANV2Bricklet) GetErrorLEDConfig() (config ErrorLEDConfig, err error) {
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

// Enables and disables the RegisterFrameReadableCallback callback.
// 
// By default the callback is disabled. Enabling this callback will disable the RegisterFrameReadCallback callback.
// 
// .. versionadded:: 2.0.3$nbsp;(Plugin)
func (device *CANV2Bricklet) SetFrameReadableCallbackConfiguration(enabled bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, enabled);

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

// Returns *true* if the RegisterFrameReadableCallback callback is enabled, *false* otherwise.
// 
// .. versionadded:: 2.0.3$nbsp;(Plugin)
func (device *CANV2Bricklet) GetFrameReadableCallbackConfiguration() (enabled bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetFrameReadableCallbackConfiguration), buf.Bytes())
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

// Enables and disables the RegisterErrorOccurredCallback callback.
// 
// By default the callback is disabled.
// 
// .. versionadded:: 2.0.3$nbsp;(Plugin)
func (device *CANV2Bricklet) SetErrorOccurredCallbackConfiguration(enabled bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, enabled);

	resultBytes, err := device.device.Set(uint8(FunctionSetErrorOccurredCallbackConfiguration), buf.Bytes())
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

// Returns *true* if the RegisterErrorOccurredCallback callback is enabled, *false* otherwise.
// 
// .. versionadded:: 2.0.3$nbsp;(Plugin)
func (device *CANV2Bricklet) GetErrorOccurredCallbackConfiguration() (enabled bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetErrorOccurredCallbackConfiguration), buf.Bytes())
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
func (device *CANV2Bricklet) GetSPITFPErrorCount() (errorCountAckChecksum uint32, errorCountMessageChecksum uint32, errorCountFrame uint32, errorCountOverflow uint32, err error) {
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
func (device *CANV2Bricklet) SetBootloaderMode(mode BootloaderMode) (status BootloaderStatus, err error) {
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
func (device *CANV2Bricklet) GetBootloaderMode() (mode BootloaderMode, err error) {
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
func (device *CANV2Bricklet) SetWriteFirmwarePointer(pointer uint32) (err error) {
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
func (device *CANV2Bricklet) WriteFirmware(data [64]uint8) (status uint8, err error) {
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
func (device *CANV2Bricklet) SetStatusLEDConfig(config StatusLEDConfig) (err error) {
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
func (device *CANV2Bricklet) GetStatusLEDConfig() (config StatusLEDConfig, err error) {
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
func (device *CANV2Bricklet) GetChipTemperature() (temperature int16, err error) {
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
func (device *CANV2Bricklet) Reset() (err error) {
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
func (device *CANV2Bricklet) WriteUID(uid uint32) (err error) {
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
func (device *CANV2Bricklet) ReadUID() (uid uint32, err error) {
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
func (device *CANV2Bricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
