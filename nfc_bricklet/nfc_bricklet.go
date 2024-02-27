/* ***********************************************************
 * This file was automatically generated on 2024-02-27.      *
 *                                                           *
 * Go Bindings Version 2.0.15                                *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/

// NFC tag read/write, NFC P2P and Card Emulation.
//
//
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/NFC_Bricklet_Go.html.
package nfc_bricklet

import (
	"bytes"
	"encoding/binary"
	"fmt"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionSetMode                             Function = 1
	FunctionGetMode                             Function = 2
	FunctionReaderRequestTagID                  Function = 3
	FunctionReaderGetTagIDLowLevel              Function = 4
	FunctionReaderGetState                      Function = 5
	FunctionReaderWriteNDEFLowLevel             Function = 6
	FunctionReaderRequestNDEF                   Function = 7
	FunctionReaderReadNDEFLowLevel              Function = 8
	FunctionReaderAuthenticateMifareClassicPage Function = 9
	FunctionReaderWritePageLowLevel             Function = 10
	FunctionReaderRequestPage                   Function = 11
	FunctionReaderReadPageLowLevel              Function = 12
	FunctionCardemuGetState                     Function = 14
	FunctionCardemuStartDiscovery               Function = 15
	FunctionCardemuWriteNDEFLowLevel            Function = 16
	FunctionCardemuStartTransfer                Function = 17
	FunctionP2PGetState                         Function = 19
	FunctionP2PStartDiscovery                   Function = 20
	FunctionP2PWriteNDEFLowLevel                Function = 21
	FunctionP2PStartTransfer                    Function = 22
	FunctionP2PReadNDEFLowLevel                 Function = 23
	FunctionSetDetectionLEDConfig               Function = 25
	FunctionGetDetectionLEDConfig               Function = 26
	FunctionSetMaximumTimeout                   Function = 27
	FunctionGetMaximumTimeout                   Function = 28
	FunctionSimpleGetTagIDLowLevel              Function = 29
	FunctionGetSPITFPErrorCount                 Function = 234
	FunctionSetBootloaderMode                   Function = 235
	FunctionGetBootloaderMode                   Function = 236
	FunctionSetWriteFirmwarePointer             Function = 237
	FunctionWriteFirmware                       Function = 238
	FunctionSetStatusLEDConfig                  Function = 239
	FunctionGetStatusLEDConfig                  Function = 240
	FunctionGetChipTemperature                  Function = 242
	FunctionReset                               Function = 243
	FunctionWriteUID                            Function = 248
	FunctionReadUID                             Function = 249
	FunctionGetIdentity                         Function = 255
	FunctionCallbackReaderStateChanged          Function = 13
	FunctionCallbackCardemuStateChanged         Function = 18
	FunctionCallbackP2PStateChanged             Function = 24
)

type Mode = uint8

const (
	ModeOff     Mode = 0
	ModeCardemu Mode = 1
	ModeP2P     Mode = 2
	ModeReader  Mode = 3
	ModeSimple  Mode = 4
)

type TagType = uint8

const (
	TagTypeMifareClassic TagType = 0
	TagTypeType1         TagType = 1
	TagTypeType2         TagType = 2
	TagTypeType3         TagType = 3
	TagTypeType4         TagType = 4
)

type ReaderState = uint8

const (
	ReaderStateInitialization                     ReaderState = 0
	ReaderStateIdle                               ReaderState = 128
	ReaderStateError                              ReaderState = 192
	ReaderStateRequestTagID                       ReaderState = 2
	ReaderStateRequestTagIDReady                  ReaderState = 130
	ReaderStateRequestTagIDError                  ReaderState = 194
	ReaderStateAuthenticateMifareClassicPage      ReaderState = 3
	ReaderStateAuthenticateMifareClassicPageReady ReaderState = 131
	ReaderStateAuthenticateMifareClassicPageError ReaderState = 195
	ReaderStateWritePage                          ReaderState = 4
	ReaderStateWritePageReady                     ReaderState = 132
	ReaderStateWritePageError                     ReaderState = 196
	ReaderStateRequestPage                        ReaderState = 5
	ReaderStateRequestPageReady                   ReaderState = 133
	ReaderStateRequestPageError                   ReaderState = 197
	ReaderStateWriteNDEF                          ReaderState = 6
	ReaderStateWriteNDEFReady                     ReaderState = 134
	ReaderStateWriteNDEFError                     ReaderState = 198
	ReaderStateRequestNDEF                        ReaderState = 7
	ReaderStateRequestNDEFReady                   ReaderState = 135
	ReaderStateRequestNDEFError                   ReaderState = 199
)

type Key = uint8

const (
	KeyA Key = 0
	KeyB Key = 1
)

type ReaderWrite = uint16

const (
	ReaderWriteType4CapabilityContainer ReaderWrite = 3
	ReaderWriteType4NDEF                ReaderWrite = 4
)

type ReaderRequest = uint16

const (
	ReaderRequestType4CapabilityContainer ReaderRequest = 3
	ReaderRequestType4NDEF                ReaderRequest = 4
)

type CardemuState = uint8

const (
	CardemuStateInitialization    CardemuState = 0
	CardemuStateIdle              CardemuState = 128
	CardemuStateError             CardemuState = 192
	CardemuStateDiscover          CardemuState = 2
	CardemuStateDiscoverReady     CardemuState = 130
	CardemuStateDiscoverError     CardemuState = 194
	CardemuStateTransferNDEF      CardemuState = 3
	CardemuStateTransferNDEFReady CardemuState = 131
	CardemuStateTransferNDEFError CardemuState = 195
)

type CardemuTransfer = uint8

const (
	CardemuTransferAbort CardemuTransfer = 0
	CardemuTransferWrite CardemuTransfer = 1
)

type P2PState = uint8

const (
	P2PStateInitialization    P2PState = 0
	P2PStateIdle              P2PState = 128
	P2PStateError             P2PState = 192
	P2PStateDiscover          P2PState = 2
	P2PStateDiscoverReady     P2PState = 130
	P2PStateDiscoverError     P2PState = 194
	P2PStateTransferNDEF      P2PState = 3
	P2PStateTransferNDEFReady P2PState = 131
	P2PStateTransferNDEFError P2PState = 195
)

type P2PTransfer = uint8

const (
	P2PTransferAbort P2PTransfer = 0
	P2PTransferWrite P2PTransfer = 1
	P2PTransferRead  P2PTransfer = 2
)

type DetectionLEDConfig = uint8

const (
	DetectionLEDConfigOff           DetectionLEDConfig = 0
	DetectionLEDConfigOn            DetectionLEDConfig = 1
	DetectionLEDConfigShowHeartbeat DetectionLEDConfig = 2
	DetectionLEDConfigShowDetection DetectionLEDConfig = 3
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

type NFCBricklet struct {
	device Device
}

const DeviceIdentifier = 286
const DeviceDisplayName = "NFC Bricklet"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (NFCBricklet, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{2, 0, 2}, uid, &internalIPCon, 9, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return NFCBricklet{}, err
	}
	dev.ResponseExpected[FunctionSetMode] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetMode] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionReaderRequestTagID] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionReaderGetTagIDLowLevel] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionReaderGetState] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionReaderWriteNDEFLowLevel] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionReaderRequestNDEF] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionReaderReadNDEFLowLevel] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionReaderAuthenticateMifareClassicPage] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionReaderWritePageLowLevel] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionReaderRequestPage] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionReaderReadPageLowLevel] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionCardemuGetState] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionCardemuStartDiscovery] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionCardemuWriteNDEFLowLevel] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionCardemuStartTransfer] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionP2PGetState] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionP2PStartDiscovery] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionP2PWriteNDEFLowLevel] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionP2PStartTransfer] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionP2PReadNDEFLowLevel] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetDetectionLEDConfig] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetDetectionLEDConfig] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetMaximumTimeout] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetMaximumTimeout] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSimpleGetTagIDLowLevel] = ResponseExpectedFlagAlwaysTrue
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
	return NFCBricklet{dev}, nil
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
func (device *NFCBricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *NFCBricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *NFCBricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *NFCBricklet) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// This callback is called if the reader state of the NFC Bricklet changes.
// See ReaderGetState for more information about the possible states.
func (device *NFCBricklet) RegisterReaderStateChangedCallback(fn func(ReaderState, bool)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 10 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var state ReaderState
		var idle bool
		binary.Read(buf, binary.LittleEndian, &state)
		binary.Read(buf, binary.LittleEndian, &idle)
		fn(state, idle)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackReaderStateChanged), wrapper)
}

// Remove a registered Reader State Changed callback.
func (device *NFCBricklet) DeregisterReaderStateChangedCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackReaderStateChanged), registrationId)
}

// This callback is called if the cardemu state of the NFC Bricklet changes.
// See CardemuGetState for more information about the possible states.
func (device *NFCBricklet) RegisterCardemuStateChangedCallback(fn func(CardemuState, bool)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 10 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var state CardemuState
		var idle bool
		binary.Read(buf, binary.LittleEndian, &state)
		binary.Read(buf, binary.LittleEndian, &idle)
		fn(state, idle)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackCardemuStateChanged), wrapper)
}

// Remove a registered Cardemu State Changed callback.
func (device *NFCBricklet) DeregisterCardemuStateChangedCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackCardemuStateChanged), registrationId)
}

// This callback is called if the P2P state of the NFC Bricklet changes.
// See P2PGetState for more information about the possible states.
func (device *NFCBricklet) RegisterP2PStateChangedCallback(fn func(P2PState, bool)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 10 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var state P2PState
		var idle bool
		binary.Read(buf, binary.LittleEndian, &state)
		binary.Read(buf, binary.LittleEndian, &idle)
		fn(state, idle)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackP2PStateChanged), wrapper)
}

// Remove a registered P2P State Changed callback.
func (device *NFCBricklet) DeregisterP2PStateChangedCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackP2PStateChanged), registrationId)
}

// Sets the mode. The NFC Bricklet supports four modes:
//
// * Off
// * Card Emulation (Cardemu): Emulates a tag for other readers
// * Peer to Peer (P2P): Exchange data with other readers
// * Reader: Reads and writes tags
// * Simple: Automatically reads tag IDs
//
// If you change a mode, the Bricklet will reconfigure the hardware for this mode.
// Therefore, you can only use functions corresponding to the current mode. For
// example, in Reader mode you can only use Reader functions.
//
// Associated constants:
//
//	* ModeOff
//	* ModeCardemu
//	* ModeP2P
//	* ModeReader
//	* ModeSimple
func (device *NFCBricklet) SetMode(mode Mode) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, mode)

	resultBytes, err := device.device.Set(uint8(FunctionSetMode), buf.Bytes())
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

// Returns the mode as set by SetMode.
//
// Associated constants:
//
//	* ModeOff
//	* ModeCardemu
//	* ModeP2P
//	* ModeReader
//	* ModeSimple
func (device *NFCBricklet) GetMode() (mode Mode, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetMode), buf.Bytes())
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

// After you call ReaderRequestTagID the NFC Bricklet will try to read
// the tag ID from the tag. After this process is done the state will change.
// You can either register the RegisterReaderStateChangedCallback callback or you can poll
// ReaderGetState to find out about the state change.
//
// If the state changes to *ReaderRequestTagIDError* it means that either there was
// no tag present or that the tag has an incompatible type. If the state
// changes to *ReaderRequestTagIDReady* it means that a compatible tag was found
// and that the tag ID has been saved. You can now read out the tag ID by
// calling ReaderGetTagID.
//
// If two tags are in the proximity of the NFC Bricklet, this
// function will cycle through the tags. To select a specific tag you have
// to call ReaderRequestTagID until the correct tag ID is found.
//
// In case of any *ReaderError* state the selection is lost and you have to
// start again by calling ReaderRequestTagID.
func (device *NFCBricklet) ReaderRequestTagID() (err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Set(uint8(FunctionReaderRequestTagID), buf.Bytes())
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

// Returns the tag type and the tag ID. This function can only be called if the
// NFC Bricklet is currently in one of the *ReaderReady* states. The returned tag ID
// is the tag ID that was saved through the last call of ReaderRequestTagID.
//
// To get the tag ID of a tag the approach is as follows:
//
// 1. Call ReaderRequestTagID
// 2. Wait for state to change to *ReaderRequestTagIDReady* (see ReaderGetState or
//    RegisterReaderStateChangedCallback callback)
// 3. Call ReaderGetTagID
//
// Associated constants:
//
//	* TagTypeMifareClassic
//	* TagTypeType1
//	* TagTypeType2
//	* TagTypeType3
//	* TagTypeType4
func (device *NFCBricklet) ReaderGetTagIDLowLevel() (tagType TagType, tagIDLength uint8, tagIDData [32]uint8, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionReaderGetTagIDLowLevel), buf.Bytes())
	if err != nil {
		return tagType, tagIDLength, tagIDData, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return tagType, tagIDLength, tagIDData, DeviceError(header.ErrorCode)
		}

		if header.Length != 42 {
			return tagType, tagIDLength, tagIDData, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 42)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &tagType)
		binary.Read(resultBuf, binary.LittleEndian, &tagIDLength)
		copy(tagIDData[:], ByteSliceToUint8Slice(resultBuf.Next(8*32/8)))

	}

	return tagType, tagIDLength, tagIDData, nil
}

// Returns the tag type and the tag ID. This function can only be called if the
// NFC Bricklet is currently in one of the *ReaderReady* states. The returned tag ID
// is the tag ID that was saved through the last call of ReaderRequestTagID.
//
// To get the tag ID of a tag the approach is as follows:
//
// 1. Call ReaderRequestTagID
// 2. Wait for state to change to *ReaderRequestTagIDReady* (see ReaderGetState or
//    RegisterReaderStateChangedCallback callback)
// 3. Call ReaderGetTagID
func (device *NFCBricklet) ReaderGetTagID() (tagID []uint8, tagType TagType, err error) {
	buf, result, err := device.device.GetHighLevel(func() (LowLevelResult, error) {
		tagType, tagIDLength, tagIDData, err := device.ReaderGetTagIDLowLevel()

		if err != nil {
			return LowLevelResult{}, err
		}

		var lowLevelResults bytes.Buffer
		binary.Write(&lowLevelResults, binary.LittleEndian, tagType)

		return LowLevelResult{
			uint64(tagIDLength),
			uint64(0),
			Uint8SliceToByteSlice(tagIDData[:]),
			lowLevelResults.Bytes()}, nil
	},
		0,
		8)
	if err != nil {
		return ByteSliceToUint8Slice(buf), tagType, err
	}
	resultBuf := bytes.NewBuffer(result)
	binary.Read(resultBuf, binary.LittleEndian, &tagType)
	return ByteSliceToUint8Slice(buf), tagType, nil
}

// Returns the current reader state of the NFC Bricklet.
//
// On startup the Bricklet will be in the *ReaderInitialization* state. The
// initialization will only take about 20ms. After that it changes to *ReaderIdle*.
//
// The Bricklet is also reinitialized if the mode is changed, see SetMode.
//
// The functions of this Bricklet can be called in the *ReaderIdle* state and all of
// the *ReaderReady* and *ReaderError* states.
//
// Example: If you call ReaderRequestPage, the state will change to
// *ReaderRequestPage* until the reading of the page is finished. Then it will change
// to either *ReaderRequestPageReady* if it worked or to *ReaderRequestPageError* if it
// didn't. If the request worked you can get the page by calling ReaderReadPage.
//
// The same approach is used analogously for the other API functions.
//
// Associated constants:
//
//	* ReaderStateInitialization
//	* ReaderStateIdle
//	* ReaderStateError
//	* ReaderStateRequestTagID
//	* ReaderStateRequestTagIDReady
//	* ReaderStateRequestTagIDError
//	* ReaderStateAuthenticateMifareClassicPage
//	* ReaderStateAuthenticateMifareClassicPageReady
//	* ReaderStateAuthenticateMifareClassicPageError
//	* ReaderStateWritePage
//	* ReaderStateWritePageReady
//	* ReaderStateWritePageError
//	* ReaderStateRequestPage
//	* ReaderStateRequestPageReady
//	* ReaderStateRequestPageError
//	* ReaderStateWriteNDEF
//	* ReaderStateWriteNDEFReady
//	* ReaderStateWriteNDEFError
//	* ReaderStateRequestNDEF
//	* ReaderStateRequestNDEFReady
//	* ReaderStateRequestNDEFError
func (device *NFCBricklet) ReaderGetState() (state ReaderState, idle bool, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionReaderGetState), buf.Bytes())
	if err != nil {
		return state, idle, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return state, idle, DeviceError(header.ErrorCode)
		}

		if header.Length != 10 {
			return state, idle, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &state)
		binary.Read(resultBuf, binary.LittleEndian, &idle)

	}

	return state, idle, nil
}

// Writes NDEF formated data.
//
// This function currently supports NFC Forum Type 2 and 4.
//
// The general approach for writing a NDEF message is as follows:
//
// 1. Call ReaderRequestTagID
// 2. Wait for state to change to *ReaderRequestTagIDReady* (see
//    ReaderGetState or RegisterReaderStateChangedCallback callback)
// 3. If looking for a specific tag then call ReaderGetTagID and check
//    if the expected tag was found, if it was not found got back to step 1
// 4. Call ReaderWriteNDEF with the NDEF message that you want to write
// 5. Wait for state to change to *ReaderWriteNDEFReady* (see ReaderGetState
//    or RegisterReaderStateChangedCallback callback)
func (device *NFCBricklet) ReaderWriteNDEFLowLevel(ndefLength uint16, ndefChunkOffset uint16, ndefChunkData [60]uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, ndefLength)
	binary.Write(&buf, binary.LittleEndian, ndefChunkOffset)
	buf.Write(Uint8SliceToByteSlice(ndefChunkData[:]))

	resultBytes, err := device.device.Set(uint8(FunctionReaderWriteNDEFLowLevel), buf.Bytes())
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

// Writes NDEF formated data.
//
// This function currently supports NFC Forum Type 2 and 4.
//
// The general approach for writing a NDEF message is as follows:
//
// 1. Call ReaderRequestTagID
// 2. Wait for state to change to *ReaderRequestTagIDReady* (see
//    ReaderGetState or RegisterReaderStateChangedCallback callback)
// 3. If looking for a specific tag then call ReaderGetTagID and check
//    if the expected tag was found, if it was not found got back to step 1
// 4. Call ReaderWriteNDEF with the NDEF message that you want to write
// 5. Wait for state to change to *ReaderWriteNDEFReady* (see ReaderGetState
//    or RegisterReaderStateChangedCallback callback)
func (device *NFCBricklet) ReaderWriteNDEF(ndef []uint8) (err error) {
	_, err = device.device.SetHighLevel(func(ndefLength uint64, ndefChunkOffset uint64, ndefChunkData []byte) (LowLevelWriteResult, error) {
		arr := [60]uint8{}
		copy(arr[:], ByteSliceToUint8Slice(ndefChunkData))

		err := device.ReaderWriteNDEFLowLevel(uint16(ndefLength), uint16(ndefChunkOffset), arr)

		var lowLevelResults bytes.Buffer

		return LowLevelWriteResult{
			uint64(60),
			lowLevelResults.Bytes()}, err
	}, 1, 8, 480, Uint8SliceToByteSlice(ndef))

	if err != nil {
		return
	}

	return
}

// Reads NDEF formated data from a tag.
//
// This function currently supports NFC Forum Type 1, 2, 3 and 4.
//
// The general approach for reading a NDEF message is as follows:
//
// 1. Call ReaderRequestTagID
// 2. Wait for state to change to *RequestTagIDReady* (see ReaderGetState
//    or RegisterReaderStateChangedCallback callback)
// 3. If looking for a specific tag then call ReaderGetTagID and check if the
//    expected tag was found, if it was not found got back to step 1
// 4. Call ReaderRequestNDEF
// 5. Wait for state to change to *ReaderRequestNDEFReady* (see ReaderGetState
//    or RegisterReaderStateChangedCallback callback)
// 6. Call ReaderReadNDEF to retrieve the NDEF message from the buffer
func (device *NFCBricklet) ReaderRequestNDEF() (err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Set(uint8(FunctionReaderRequestNDEF), buf.Bytes())
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

// Returns the NDEF data from an internal buffer. To fill the buffer
// with a NDEF message you have to call ReaderRequestNDEF beforehand.
func (device *NFCBricklet) ReaderReadNDEFLowLevel() (ndefLength uint16, ndefChunkOffset uint16, ndefChunkData [60]uint8, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionReaderReadNDEFLowLevel), buf.Bytes())
	if err != nil {
		return ndefLength, ndefChunkOffset, ndefChunkData, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return ndefLength, ndefChunkOffset, ndefChunkData, DeviceError(header.ErrorCode)
		}

		if header.Length != 72 {
			return ndefLength, ndefChunkOffset, ndefChunkData, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 72)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &ndefLength)
		binary.Read(resultBuf, binary.LittleEndian, &ndefChunkOffset)
		copy(ndefChunkData[:], ByteSliceToUint8Slice(resultBuf.Next(8*60/8)))

	}

	return ndefLength, ndefChunkOffset, ndefChunkData, nil
}

// Returns the NDEF data from an internal buffer. To fill the buffer
// with a NDEF message you have to call ReaderRequestNDEF beforehand.
func (device *NFCBricklet) ReaderReadNDEF() (ndef []uint8, err error) {
	buf, _, err := device.device.GetHighLevel(func() (LowLevelResult, error) {
		ndefLength, ndefChunkOffset, ndefChunkData, err := device.ReaderReadNDEFLowLevel()

		if err != nil {
			return LowLevelResult{}, err
		}

		var lowLevelResults bytes.Buffer

		return LowLevelResult{
			uint64(ndefLength),
			uint64(ndefChunkOffset),
			Uint8SliceToByteSlice(ndefChunkData[:]),
			lowLevelResults.Bytes()}, nil
	},
		2,
		8)
	if err != nil {
		return ByteSliceToUint8Slice(buf), err
	}

	return ByteSliceToUint8Slice(buf), nil
}

// Mifare Classic tags use authentication. If you want to read from or write to
// a Mifare Classic page you have to authenticate it beforehand.
// Each page can be authenticated with two keys: A (``key_number`` = 0) and B
// (``key_number`` = 1). A new Mifare Classic
// tag that has not yet been written to can be accessed with key A
// and the default key ``[0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF]``.
//
// The approach to read or write a Mifare Classic page is as follows:
//
// 1. Call ReaderRequestTagID
// 2. Wait for state to change to *ReaderRequestTagIDReady* (see ReaderGetState
//    or RegisterReaderStateChangedCallback callback)
// 3. If looking for a specific tag then call ReaderGetTagID and check if the
//    expected tag was found, if it was not found got back to step 1
// 4. Call ReaderAuthenticateMifareClassicPage with page and key for the page
// 5. Wait for state to change to *ReaderAuthenticatingMifareClassicPageReady* (see
//    ReaderGetState or RegisterReaderStateChangedCallback callback)
// 6. Call ReaderRequestPage or ReaderWritePage to read/write page
//
// The authentication will always work for one whole sector (4 pages).
//
// Associated constants:
//
//	* KeyA
//	* KeyB
func (device *NFCBricklet) ReaderAuthenticateMifareClassicPage(page uint16, keyNumber Key, key [6]uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, page)
	binary.Write(&buf, binary.LittleEndian, keyNumber)
	binary.Write(&buf, binary.LittleEndian, key)

	resultBytes, err := device.device.Set(uint8(FunctionReaderAuthenticateMifareClassicPage), buf.Bytes())
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

// Writes a maximum of 8192 bytes starting from the given page. How many pages are written
// depends on the tag type. The page sizes are as follows:
//
// * Mifare Classic page size: 16 byte
// * NFC Forum Type 1 page size: 8 byte
// * NFC Forum Type 2 page size: 4 byte
// * NFC Forum Type 3 page size: 16 byte
// * NFC Forum Type 4: No pages, page = file selection (CC or NDEF, see below)
//
// The general approach for writing to a tag is as follows:
//
// 1. Call ReaderRequestTagID
// 2. Wait for state to change to *ReaderRequestTagIDReady* (see ReaderGetState or
//    RegisterReaderStateChangedCallback callback)
// 3. If looking for a specific tag then call ReaderGetTagID and check if the
//    expected tag was found, if it was not found got back to step 1
// 4. Call ReaderWritePage with page number and data
// 5. Wait for state to change to *ReaderWritePageReady* (see ReaderGetState or
//    RegisterReaderStateChangedCallback callback)
//
// If you use a Mifare Classic tag you have to authenticate a page before you
// can write to it. See ReaderAuthenticateMifareClassicPage.
//
// NFC Forum Type 4 tags are not organized into pages but different files. We currently
// support two files: Capability Container file (CC) and NDEF file.
//
// Choose CC by setting page to 3 or NDEF by setting page to 4.
//
// Associated constants:
//
//	* ReaderWriteType4CapabilityContainer
//	* ReaderWriteType4NDEF
func (device *NFCBricklet) ReaderWritePageLowLevel(page ReaderWrite, dataLength uint16, dataChunkOffset uint16, dataChunkData [58]uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, page)
	binary.Write(&buf, binary.LittleEndian, dataLength)
	binary.Write(&buf, binary.LittleEndian, dataChunkOffset)
	buf.Write(Uint8SliceToByteSlice(dataChunkData[:]))

	resultBytes, err := device.device.Set(uint8(FunctionReaderWritePageLowLevel), buf.Bytes())
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

// Writes a maximum of 8192 bytes starting from the given page. How many pages are written
// depends on the tag type. The page sizes are as follows:
//
// * Mifare Classic page size: 16 byte
// * NFC Forum Type 1 page size: 8 byte
// * NFC Forum Type 2 page size: 4 byte
// * NFC Forum Type 3 page size: 16 byte
// * NFC Forum Type 4: No pages, page = file selection (CC or NDEF, see below)
//
// The general approach for writing to a tag is as follows:
//
// 1. Call ReaderRequestTagID
// 2. Wait for state to change to *ReaderRequestTagIDReady* (see ReaderGetState or
//    RegisterReaderStateChangedCallback callback)
// 3. If looking for a specific tag then call ReaderGetTagID and check if the
//    expected tag was found, if it was not found got back to step 1
// 4. Call ReaderWritePage with page number and data
// 5. Wait for state to change to *ReaderWritePageReady* (see ReaderGetState or
//    RegisterReaderStateChangedCallback callback)
//
// If you use a Mifare Classic tag you have to authenticate a page before you
// can write to it. See ReaderAuthenticateMifareClassicPage.
//
// NFC Forum Type 4 tags are not organized into pages but different files. We currently
// support two files: Capability Container file (CC) and NDEF file.
//
// Choose CC by setting page to 3 or NDEF by setting page to 4.
func (device *NFCBricklet) ReaderWritePage(page ReaderWrite, data []uint8) (err error) {
	_, err = device.device.SetHighLevel(func(dataLength uint64, dataChunkOffset uint64, dataChunkData []byte) (LowLevelWriteResult, error) {
		arr := [58]uint8{}
		copy(arr[:], ByteSliceToUint8Slice(dataChunkData))

		err := device.ReaderWritePageLowLevel(page, uint16(dataLength), uint16(dataChunkOffset), arr)

		var lowLevelResults bytes.Buffer

		return LowLevelWriteResult{
			uint64(58),
			lowLevelResults.Bytes()}, err
	}, 3, 8, 464, Uint8SliceToByteSlice(data))

	if err != nil {
		return
	}

	return
}

// Reads a maximum of 8192 bytes starting from the given page and stores them into a buffer.
// The buffer can then be read out with ReaderReadPage.
// How many pages are read depends on the tag type. The page sizes are
// as follows:
//
// * Mifare Classic page size: 16 byte
// * NFC Forum Type 1 page size: 8 byte
// * NFC Forum Type 2 page size: 4 byte
// * NFC Forum Type 3 page size: 16 byte
// * NFC Forum Type 4: No pages, page = file selection (CC or NDEF, see below)
//
// The general approach for reading a tag is as follows:
//
// 1. Call ReaderRequestTagID
// 2. Wait for state to change to *RequestTagIDReady* (see ReaderGetState
//    or RegisterReaderStateChangedCallback callback)
// 3. If looking for a specific tag then call ReaderGetTagID and check if the
//    expected tag was found, if it was not found got back to step 1
// 4. Call ReaderRequestPage with page number
// 5. Wait for state to change to *ReaderRequestPageReady* (see ReaderGetState
//    or RegisterReaderStateChangedCallback callback)
// 6. Call ReaderReadPage to retrieve the page from the buffer
//
// If you use a Mifare Classic tag you have to authenticate a page before you
// can read it. See ReaderAuthenticateMifareClassicPage.
//
// NFC Forum Type 4 tags are not organized into pages but different files. We currently
// support two files: Capability Container file (CC) and NDEF file.
//
// Choose CC by setting page to 3 or NDEF by setting page to 4.
//
// Associated constants:
//
//	* ReaderRequestType4CapabilityContainer
//	* ReaderRequestType4NDEF
func (device *NFCBricklet) ReaderRequestPage(page ReaderRequest, length uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, page)
	binary.Write(&buf, binary.LittleEndian, length)

	resultBytes, err := device.device.Set(uint8(FunctionReaderRequestPage), buf.Bytes())
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

// Returns the page data from an internal buffer. To fill the buffer
// with specific pages you have to call ReaderRequestPage beforehand.
func (device *NFCBricklet) ReaderReadPageLowLevel() (dataLength uint16, dataChunkOffset uint16, dataChunkData [60]uint8, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionReaderReadPageLowLevel), buf.Bytes())
	if err != nil {
		return dataLength, dataChunkOffset, dataChunkData, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return dataLength, dataChunkOffset, dataChunkData, DeviceError(header.ErrorCode)
		}

		if header.Length != 72 {
			return dataLength, dataChunkOffset, dataChunkData, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 72)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &dataLength)
		binary.Read(resultBuf, binary.LittleEndian, &dataChunkOffset)
		copy(dataChunkData[:], ByteSliceToUint8Slice(resultBuf.Next(8*60/8)))

	}

	return dataLength, dataChunkOffset, dataChunkData, nil
}

// Returns the page data from an internal buffer. To fill the buffer
// with specific pages you have to call ReaderRequestPage beforehand.
func (device *NFCBricklet) ReaderReadPage() (data []uint8, err error) {
	buf, _, err := device.device.GetHighLevel(func() (LowLevelResult, error) {
		dataLength, dataChunkOffset, dataChunkData, err := device.ReaderReadPageLowLevel()

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
		4,
		8)
	if err != nil {
		return ByteSliceToUint8Slice(buf), err
	}

	return ByteSliceToUint8Slice(buf), nil
}

// Returns the current cardemu state of the NFC Bricklet.
//
// On startup the Bricklet will be in the *CardemuInitialization* state. The
// initialization will only take about 20ms. After that it changes to *CardemuIdle*.
//
// The Bricklet is also reinitialized if the mode is changed, see SetMode.
//
// The functions of this Bricklet can be called in the *CardemuIdle* state and all of
// the *CardemuReady* and *CardemuError* states.
//
// Example: If you call CardemuStartDiscovery, the state will change to
// *CardemuDiscover* until the discovery is finished. Then it will change
// to either *CardemuDiscoverReady* if it worked or to *CardemuDiscoverError* if it
// didn't.
//
// The same approach is used analogously for the other API functions.
//
// Associated constants:
//
//	* CardemuStateInitialization
//	* CardemuStateIdle
//	* CardemuStateError
//	* CardemuStateDiscover
//	* CardemuStateDiscoverReady
//	* CardemuStateDiscoverError
//	* CardemuStateTransferNDEF
//	* CardemuStateTransferNDEFReady
//	* CardemuStateTransferNDEFError
func (device *NFCBricklet) CardemuGetState() (state CardemuState, idle bool, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionCardemuGetState), buf.Bytes())
	if err != nil {
		return state, idle, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return state, idle, DeviceError(header.ErrorCode)
		}

		if header.Length != 10 {
			return state, idle, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &state)
		binary.Read(resultBuf, binary.LittleEndian, &idle)

	}

	return state, idle, nil
}

// Starts the discovery process. If you call this function while a NFC
// reader device is near to the NFC Bricklet the state will change from
// *CardemuDiscovery* to *CardemuDiscoveryReady*.
//
// If no NFC reader device can be found or if there is an error during
// discovery the cardemu state will change to *CardemuDiscoveryError*. In this case you
// have to restart the discovery process.
//
// If the cardemu state changes to *CardemuDiscoveryReady* you can start the NDEF message
// transfer with CardemuWriteNDEF and CardemuStartTransfer.
func (device *NFCBricklet) CardemuStartDiscovery() (err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Set(uint8(FunctionCardemuStartDiscovery), buf.Bytes())
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

// Writes the NDEF message that is to be transferred to the NFC peer.
//
// The maximum supported NDEF message size in Cardemu mode is 255 byte.
//
// You can call this function at any time in Cardemu mode. The internal buffer
// will not be overwritten until you call this function again or change the
// mode.
func (device *NFCBricklet) CardemuWriteNDEFLowLevel(ndefLength uint16, ndefChunkOffset uint16, ndefChunkData [60]uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, ndefLength)
	binary.Write(&buf, binary.LittleEndian, ndefChunkOffset)
	buf.Write(Uint8SliceToByteSlice(ndefChunkData[:]))

	resultBytes, err := device.device.Set(uint8(FunctionCardemuWriteNDEFLowLevel), buf.Bytes())
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

// Writes the NDEF message that is to be transferred to the NFC peer.
//
// The maximum supported NDEF message size in Cardemu mode is 255 byte.
//
// You can call this function at any time in Cardemu mode. The internal buffer
// will not be overwritten until you call this function again or change the
// mode.
func (device *NFCBricklet) CardemuWriteNDEF(ndef []uint8) (err error) {
	_, err = device.device.SetHighLevel(func(ndefLength uint64, ndefChunkOffset uint64, ndefChunkData []byte) (LowLevelWriteResult, error) {
		arr := [60]uint8{}
		copy(arr[:], ByteSliceToUint8Slice(ndefChunkData))

		err := device.CardemuWriteNDEFLowLevel(uint16(ndefLength), uint16(ndefChunkOffset), arr)

		var lowLevelResults bytes.Buffer

		return LowLevelWriteResult{
			uint64(60),
			lowLevelResults.Bytes()}, err
	}, 5, 8, 480, Uint8SliceToByteSlice(ndef))

	if err != nil {
		return
	}

	return
}

// You can start the transfer of a NDEF message if the cardemu state is *CardemuDiscoveryReady*.
//
// Before you call this function to start a write transfer, the NDEF message that
// is to be transferred has to be written via CardemuWriteNDEF first.
//
// After you call this function the state will change to *CardemuTransferNDEF*. It will
// change to *CardemuTransferNDEFReady* if the transfer was successful or
// *CardemuTransferNDEFError* if it wasn't.
//
// Associated constants:
//
//	* CardemuTransferAbort
//	* CardemuTransferWrite
func (device *NFCBricklet) CardemuStartTransfer(transfer CardemuTransfer) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, transfer)

	resultBytes, err := device.device.Set(uint8(FunctionCardemuStartTransfer), buf.Bytes())
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

// Returns the current P2P state of the NFC Bricklet.
//
// On startup the Bricklet will be in the *P2PInitialization* state. The
// initialization will only take about 20ms. After that it changes to *P2PIdle*.
//
// The Bricklet is also reinitialized if the mode is changed, see SetMode.
//
// The functions of this Bricklet can be called in the *P2PIdle* state and all of
// the *P2PReady* and *P2PError* states.
//
// Example: If you call P2PStartDiscovery, the state will change to
// *P2PDiscover* until the discovery is finished. Then it will change
// to either P2PDiscoverReady* if it worked or to *P2PDiscoverError* if it
// didn't.
//
// The same approach is used analogously for the other API functions.
//
// Associated constants:
//
//	* P2PStateInitialization
//	* P2PStateIdle
//	* P2PStateError
//	* P2PStateDiscover
//	* P2PStateDiscoverReady
//	* P2PStateDiscoverError
//	* P2PStateTransferNDEF
//	* P2PStateTransferNDEFReady
//	* P2PStateTransferNDEFError
func (device *NFCBricklet) P2PGetState() (state P2PState, idle bool, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionP2PGetState), buf.Bytes())
	if err != nil {
		return state, idle, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return state, idle, DeviceError(header.ErrorCode)
		}

		if header.Length != 10 {
			return state, idle, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &state)
		binary.Read(resultBuf, binary.LittleEndian, &idle)

	}

	return state, idle, nil
}

// Starts the discovery process. If you call this function while another NFC
// P2P enabled device is near to the NFC Bricklet the state will change from
// *P2PDiscovery* to *P2PDiscoveryReady*.
//
// If no NFC P2P enabled device can be found or if there is an error during
// discovery the P2P state will change to *P2PDiscoveryError*. In this case you
// have to restart the discovery process.
//
// If the P2P state changes to *P2PDiscoveryReady* you can start the NDEF message
// transfer with P2PStartTransfer.
func (device *NFCBricklet) P2PStartDiscovery() (err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Set(uint8(FunctionP2PStartDiscovery), buf.Bytes())
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

// Writes the NDEF message that is to be transferred to the NFC peer.
//
// The maximum supported NDEF message size for P2P transfer is 255 byte.
//
// You can call this function at any time in P2P mode. The internal buffer
// will not be overwritten until you call this function again, change the
// mode or use P2P to read an NDEF messages.
func (device *NFCBricklet) P2PWriteNDEFLowLevel(ndefLength uint16, ndefChunkOffset uint16, ndefChunkData [60]uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, ndefLength)
	binary.Write(&buf, binary.LittleEndian, ndefChunkOffset)
	buf.Write(Uint8SliceToByteSlice(ndefChunkData[:]))

	resultBytes, err := device.device.Set(uint8(FunctionP2PWriteNDEFLowLevel), buf.Bytes())
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

// Writes the NDEF message that is to be transferred to the NFC peer.
//
// The maximum supported NDEF message size for P2P transfer is 255 byte.
//
// You can call this function at any time in P2P mode. The internal buffer
// will not be overwritten until you call this function again, change the
// mode or use P2P to read an NDEF messages.
func (device *NFCBricklet) P2PWriteNDEF(ndef []uint8) (err error) {
	_, err = device.device.SetHighLevel(func(ndefLength uint64, ndefChunkOffset uint64, ndefChunkData []byte) (LowLevelWriteResult, error) {
		arr := [60]uint8{}
		copy(arr[:], ByteSliceToUint8Slice(ndefChunkData))

		err := device.P2PWriteNDEFLowLevel(uint16(ndefLength), uint16(ndefChunkOffset), arr)

		var lowLevelResults bytes.Buffer

		return LowLevelWriteResult{
			uint64(60),
			lowLevelResults.Bytes()}, err
	}, 6, 8, 480, Uint8SliceToByteSlice(ndef))

	if err != nil {
		return
	}

	return
}

// You can start the transfer of a NDEF message if the P2P state is *P2PDiscoveryReady*.
//
// Before you call this function to start a write transfer, the NDEF message that
// is to be transferred has to be written via P2PWriteNDEF first.
//
// After you call this function the P2P state will change to *P2PTransferNDEF*. It will
// change to *P2PTransferNDEFReady* if the transfer was successfull or
// *P2PTransferNDEFError* if it wasn't.
//
// If you started a write transfer you are now done. If you started a read transfer
// you can now use P2PReadNDEF to read the NDEF message that was written
// by the NFC peer.
//
// Associated constants:
//
//	* P2PTransferAbort
//	* P2PTransferWrite
//	* P2PTransferRead
func (device *NFCBricklet) P2PStartTransfer(transfer P2PTransfer) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, transfer)

	resultBytes, err := device.device.Set(uint8(FunctionP2PStartTransfer), buf.Bytes())
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

// Returns the NDEF message that was written by a NFC peer in NFC P2P mode.
//
// The NDEF message is ready if you called P2PStartTransfer with a
// read transfer and the P2P state changed to *P2PTransferNDEFReady*.
func (device *NFCBricklet) P2PReadNDEFLowLevel() (ndefLength uint16, ndefChunkOffset uint16, ndefChunkData [60]uint8, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionP2PReadNDEFLowLevel), buf.Bytes())
	if err != nil {
		return ndefLength, ndefChunkOffset, ndefChunkData, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return ndefLength, ndefChunkOffset, ndefChunkData, DeviceError(header.ErrorCode)
		}

		if header.Length != 72 {
			return ndefLength, ndefChunkOffset, ndefChunkData, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 72)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &ndefLength)
		binary.Read(resultBuf, binary.LittleEndian, &ndefChunkOffset)
		copy(ndefChunkData[:], ByteSliceToUint8Slice(resultBuf.Next(8*60/8)))

	}

	return ndefLength, ndefChunkOffset, ndefChunkData, nil
}

// Returns the NDEF message that was written by a NFC peer in NFC P2P mode.
//
// The NDEF message is ready if you called P2PStartTransfer with a
// read transfer and the P2P state changed to *P2PTransferNDEFReady*.
func (device *NFCBricklet) P2PReadNDEF() (ndef []uint8, err error) {
	buf, _, err := device.device.GetHighLevel(func() (LowLevelResult, error) {
		ndefLength, ndefChunkOffset, ndefChunkData, err := device.P2PReadNDEFLowLevel()

		if err != nil {
			return LowLevelResult{}, err
		}

		var lowLevelResults bytes.Buffer

		return LowLevelResult{
			uint64(ndefLength),
			uint64(ndefChunkOffset),
			Uint8SliceToByteSlice(ndefChunkData[:]),
			lowLevelResults.Bytes()}, nil
	},
		7,
		8)
	if err != nil {
		return ByteSliceToUint8Slice(buf), err
	}

	return ByteSliceToUint8Slice(buf), nil
}

// Sets the detection LED configuration. By default the LED shows
// if a card/reader is detected.
//
// You can also turn the LED permanently on/off or show a heartbeat.
//
// If the Bricklet is in bootloader mode, the LED is off.
//
// Associated constants:
//
//	* DetectionLEDConfigOff
//	* DetectionLEDConfigOn
//	* DetectionLEDConfigShowHeartbeat
//	* DetectionLEDConfigShowDetection
func (device *NFCBricklet) SetDetectionLEDConfig(config DetectionLEDConfig) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, config)

	resultBytes, err := device.device.Set(uint8(FunctionSetDetectionLEDConfig), buf.Bytes())
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

// Returns the configuration as set by SetDetectionLEDConfig
//
// Associated constants:
//
//	* DetectionLEDConfigOff
//	* DetectionLEDConfigOn
//	* DetectionLEDConfigShowHeartbeat
//	* DetectionLEDConfigShowDetection
func (device *NFCBricklet) GetDetectionLEDConfig() (config DetectionLEDConfig, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetDetectionLEDConfig), buf.Bytes())
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

// Sets the maximum timeout.
//
// This is a global maximum used for all internal state timeouts. The timeouts depend heavily
// on the used tags etc. For example: If you use a Type 2 tag and you want to detect if
// it is present, you have to use ReaderRequestTagID and wait for the state
// to change to either the error state or the ready state.
//
// With the default configuration this takes 2-3 seconds. By setting the maximum timeout to
// 100ms you can reduce this time to ~150-200ms. For Type 2 this would also still work
// with a 20ms timeout (a Type 2 tag answers usually within 10ms). A type 4 tag can take
// up to 500ms in our tests.
//
// If you need a fast response time to discover if a tag is present or not you can find
// a good timeout value by trial and error for your specific tag.
//
// By default we use a very conservative timeout, to be sure that any tag can always
// answer in time.
//
// .. versionadded:: 2.0.1$nbsp;(Plugin)
func (device *NFCBricklet) SetMaximumTimeout(timeout uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, timeout)

	resultBytes, err := device.device.Set(uint8(FunctionSetMaximumTimeout), buf.Bytes())
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

// Returns the timeout as set by SetMaximumTimeout
//
// .. versionadded:: 2.0.1$nbsp;(Plugin)
func (device *NFCBricklet) GetMaximumTimeout() (timeout uint16, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetMaximumTimeout), buf.Bytes())
	if err != nil {
		return timeout, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return timeout, DeviceError(header.ErrorCode)
		}

		if header.Length != 10 {
			return timeout, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &timeout)

	}

	return timeout, nil
}

// .. versionadded:: 2.0.6$nbsp;(Plugin)
//
// Associated constants:
//
//	* TagTypeMifareClassic
//	* TagTypeType1
//	* TagTypeType2
//	* TagTypeType3
//	* TagTypeType4
func (device *NFCBricklet) SimpleGetTagIDLowLevel(index uint8) (tagType TagType, tagIDLength uint8, tagIDData [10]uint8, lastSeen uint32, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, index)

	resultBytes, err := device.device.Get(uint8(FunctionSimpleGetTagIDLowLevel), buf.Bytes())
	if err != nil {
		return tagType, tagIDLength, tagIDData, lastSeen, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return tagType, tagIDLength, tagIDData, lastSeen, DeviceError(header.ErrorCode)
		}

		if header.Length != 24 {
			return tagType, tagIDLength, tagIDData, lastSeen, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 24)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &tagType)
		binary.Read(resultBuf, binary.LittleEndian, &tagIDLength)
		copy(tagIDData[:], ByteSliceToUint8Slice(resultBuf.Next(8*10/8)))
		binary.Read(resultBuf, binary.LittleEndian, &lastSeen)

	}

	return tagType, tagIDLength, tagIDData, lastSeen, nil
}

// .. versionadded:: 2.0.6$nbsp;(Plugin)
func (device *NFCBricklet) SimpleGetTagID(index uint8) (tagID []uint8, tagType TagType, lastSeen uint32, err error) {
	buf, result, err := device.device.GetHighLevel(func() (LowLevelResult, error) {
		tagType, tagIDLength, tagIDData, lastSeen, err := device.SimpleGetTagIDLowLevel(index)

		if err != nil {
			return LowLevelResult{}, err
		}

		var lowLevelResults bytes.Buffer
		binary.Write(&lowLevelResults, binary.LittleEndian, tagType)
		binary.Write(&lowLevelResults, binary.LittleEndian, lastSeen)

		return LowLevelResult{
			uint64(tagIDLength),
			uint64(0),
			Uint8SliceToByteSlice(tagIDData[:]),
			lowLevelResults.Bytes()}, nil
	},
		8,
		8)
	if err != nil {
		return ByteSliceToUint8Slice(buf), tagType, lastSeen, err
	}
	resultBuf := bytes.NewBuffer(result)
	binary.Read(resultBuf, binary.LittleEndian, &tagType)
	binary.Read(resultBuf, binary.LittleEndian, &lastSeen)
	return ByteSliceToUint8Slice(buf), tagType, lastSeen, nil
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
func (device *NFCBricklet) GetSPITFPErrorCount() (errorCountAckChecksum uint32, errorCountMessageChecksum uint32, errorCountFrame uint32, errorCountOverflow uint32, err error) {
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
func (device *NFCBricklet) SetBootloaderMode(mode BootloaderMode) (status BootloaderStatus, err error) {
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
func (device *NFCBricklet) GetBootloaderMode() (mode BootloaderMode, err error) {
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
func (device *NFCBricklet) SetWriteFirmwarePointer(pointer uint32) (err error) {
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
func (device *NFCBricklet) WriteFirmware(data [64]uint8) (status uint8, err error) {
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
func (device *NFCBricklet) SetStatusLEDConfig(config StatusLEDConfig) (err error) {
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
func (device *NFCBricklet) GetStatusLEDConfig() (config StatusLEDConfig, err error) {
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
func (device *NFCBricklet) GetChipTemperature() (temperature int16, err error) {
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
func (device *NFCBricklet) Reset() (err error) {
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
func (device *NFCBricklet) WriteUID(uid uint32) (err error) {
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
func (device *NFCBricklet) ReadUID() (uid uint32, err error) {
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
func (device *NFCBricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
