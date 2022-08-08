/* ***********************************************************
 * This file was automatically generated on 2022-08-08.      *
 *                                                           *
 * Go Bindings Version 2.0.13                                *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/

// Reads and writes NFC and RFID tags.
//
//
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/NFCRFID_Bricklet_Go.html.
package nfc_rfid_bricklet

import (
	"bytes"
	"encoding/binary"
	"fmt"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionRequestTagID                  Function = 1
	FunctionGetTagID                      Function = 2
	FunctionGetState                      Function = 3
	FunctionAuthenticateMifareClassicPage Function = 4
	FunctionWritePage                     Function = 5
	FunctionRequestPage                   Function = 6
	FunctionGetPage                       Function = 7
	FunctionGetIdentity                   Function = 255
	FunctionCallbackStateChanged          Function = 8
)

type TagType = uint8

const (
	TagTypeMifareClassic TagType = 0
	TagTypeType1         TagType = 1
	TagTypeType2         TagType = 2
)

type State = uint8

const (
	StateInitialization                       State = 0
	StateIdle                                 State = 128
	StateError                                State = 192
	StateRequestTagID                         State = 2
	StateRequestTagIDReady                    State = 130
	StateRequestTagIDError                    State = 194
	StateAuthenticatingMifareClassicPage      State = 3
	StateAuthenticatingMifareClassicPageReady State = 131
	StateAuthenticatingMifareClassicPageError State = 195
	StateWritePage                            State = 4
	StateWritePageReady                       State = 132
	StateWritePageError                       State = 196
	StateRequestPage                          State = 5
	StateRequestPageReady                     State = 133
	StateRequestPageError                     State = 197
)

type Key = uint8

const (
	KeyA Key = 0
	KeyB Key = 1
)

type NFCRFIDBricklet struct {
	device Device
}

const DeviceIdentifier = 246
const DeviceDisplayName = "NFC/RFID Bricklet"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (NFCRFIDBricklet, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{2, 0, 0}, uid, &internalIPCon, 0, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return NFCRFIDBricklet{}, err
	}
	dev.ResponseExpected[FunctionRequestTagID] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetTagID] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetState] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionAuthenticateMifareClassicPage] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionWritePage] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionRequestPage] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetPage] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetIdentity] = ResponseExpectedFlagAlwaysTrue
	return NFCRFIDBricklet{dev}, nil
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
func (device *NFCRFIDBricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *NFCRFIDBricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *NFCRFIDBricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *NFCRFIDBricklet) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// This callback is called if the state of the NFC/RFID Bricklet changes.
// See GetState for more information about the possible states.
func (device *NFCRFIDBricklet) RegisterStateChangedCallback(fn func(State, bool)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 10 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var state State
		var idle bool
		binary.Read(buf, binary.LittleEndian, &state)
		binary.Read(buf, binary.LittleEndian, &idle)
		fn(state, idle)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackStateChanged), wrapper)
}

// Remove a registered State Changed callback.
func (device *NFCRFIDBricklet) DeregisterStateChangedCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackStateChanged), registrationId)
}

// To read or write a tag that is in proximity of the NFC/RFID Bricklet you
// first have to call this function with the expected tag type as parameter.
// It is no problem if you don't know the tag type. You can cycle through
// the available tag types until the tag gives an answer to the request.
//
// Currently the following tag types are supported:
//
// * Mifare Classic
// * NFC Forum Type 1
// * NFC Forum Type 2
//
// After you call RequestTagID the NFC/RFID Bricklet will try to read
// the tag ID from the tag. After this process is done the state will change.
// You can either register the RegisterStateChangedCallback callback or you can poll
// GetState to find out about the state change.
//
// If the state changes to *RequestTagIDError* it means that either there was
// no tag present or that the tag is of an incompatible type. If the state
// changes to *RequestTagIDReady* it means that a compatible tag was found
// and that the tag ID could be read out. You can now get the tag ID by
// calling GetTagID.
//
// If two tags are in the proximity of the NFC/RFID Bricklet, this
// function will cycle through the tags. To select a specific tag you have
// to call RequestTagID until the correct tag id is found.
//
// In case of any *Error* state the selection is lost and you have to
// start again by calling RequestTagID.
//
// Associated constants:
//
//	* TagTypeMifareClassic
//	* TagTypeType1
//	* TagTypeType2
func (device *NFCRFIDBricklet) RequestTagID(tagType TagType) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, tagType)

	resultBytes, err := device.device.Set(uint8(FunctionRequestTagID), buf.Bytes())
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

// Returns the tag type, tag ID and the length of the tag ID
// (4 or 7 bytes are possible length). This function can only be called if the
// NFC/RFID is currently in one of the *Ready* states. The returned ID
// is the ID that was saved through the last call of RequestTagID.
//
// To get the tag ID of a tag the approach is as follows:
//
// 1. Call RequestTagID
// 2. Wait for state to change to *RequestTagIDReady* (see GetState or
//    RegisterStateChangedCallback callback)
// 3. Call GetTagID
//
// Associated constants:
//
//	* TagTypeMifareClassic
//	* TagTypeType1
//	* TagTypeType2
func (device *NFCRFIDBricklet) GetTagID() (tagType TagType, tidLength uint8, tid [7]uint8, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetTagID), buf.Bytes())
	if err != nil {
		return tagType, tidLength, tid, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 17 {
			return tagType, tidLength, tid, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 17)
		}

		if header.ErrorCode != 0 {
			return tagType, tidLength, tid, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &tagType)
		binary.Read(resultBuf, binary.LittleEndian, &tidLength)
		binary.Read(resultBuf, binary.LittleEndian, &tid)

	}

	return tagType, tidLength, tid, nil
}

// Returns the current state of the NFC/RFID Bricklet.
//
// On startup the Bricklet will be in the *Initialization* state. The
// initialization will only take about 20ms. After that it changes to *Idle*.
//
// The functions of this Bricklet can be called in the *Idle* state and all of
// the *Ready* and *Error* states.
//
// Example: If you call RequestPage, the state will change to
// *RequestPage* until the reading of the page is finished. Then it will change
// to either *RequestPageReady* if it worked or to *RequestPageError* if it
// didn't. If the request worked you can get the page by calling GetPage.
//
// The same approach is used analogously for the other API functions.
//
// Associated constants:
//
//	* StateInitialization
//	* StateIdle
//	* StateError
//	* StateRequestTagID
//	* StateRequestTagIDReady
//	* StateRequestTagIDError
//	* StateAuthenticatingMifareClassicPage
//	* StateAuthenticatingMifareClassicPageReady
//	* StateAuthenticatingMifareClassicPageError
//	* StateWritePage
//	* StateWritePageReady
//	* StateWritePageError
//	* StateRequestPage
//	* StateRequestPageReady
//	* StateRequestPageError
func (device *NFCRFIDBricklet) GetState() (state State, idle bool, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetState), buf.Bytes())
	if err != nil {
		return state, idle, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 10 {
			return state, idle, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		if header.ErrorCode != 0 {
			return state, idle, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &state)
		binary.Read(resultBuf, binary.LittleEndian, &idle)

	}

	return state, idle, nil
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
// 1. Call RequestTagID
// 2. Wait for state to change to *RequestTagIDReady* (see GetState
//    or RegisterStateChangedCallback callback)
// 3. If looking for a specific tag then call GetTagID and check if the
//    expected tag was found, if it was not found go back to step 1
// 4. Call AuthenticateMifareClassicPage with page and key for the page
// 5. Wait for state to change to *AuthenticatingMifareClassicPageReady* (see
//    GetState or RegisterStateChangedCallback callback)
// 6. Call RequestPage or WritePage to read/write page
//
// Associated constants:
//
//	* KeyA
//	* KeyB
func (device *NFCRFIDBricklet) AuthenticateMifareClassicPage(page uint16, keyNumber Key, key [6]uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, page)
	binary.Write(&buf, binary.LittleEndian, keyNumber)
	binary.Write(&buf, binary.LittleEndian, key)

	resultBytes, err := device.device.Set(uint8(FunctionAuthenticateMifareClassicPage), buf.Bytes())
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

// Writes 16 bytes starting from the given page. How many pages are written
// depends on the tag type. The page sizes are as follows:
//
// * Mifare Classic page size: 16 byte (one page is written)
// * NFC Forum Type 1 page size: 8 byte (two pages are written)
// * NFC Forum Type 2 page size: 4 byte (four pages are written)
//
// The general approach for writing to a tag is as follows:
//
// 1. Call RequestTagID
// 2. Wait for state to change to *RequestTagIDReady* (see GetState or
//    RegisterStateChangedCallback callback)
// 3. If looking for a specific tag then call GetTagID and check if the
//    expected tag was found, if it was not found got back to step 1
// 4. Call WritePage with page number and data
// 5. Wait for state to change to *WritePageReady* (see GetState or
//    RegisterStateChangedCallback callback)
//
// If you use a Mifare Classic tag you have to authenticate a page before you
// can write to it. See AuthenticateMifareClassicPage.
func (device *NFCRFIDBricklet) WritePage(page uint16, data [16]uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, page)
	binary.Write(&buf, binary.LittleEndian, data)

	resultBytes, err := device.device.Set(uint8(FunctionWritePage), buf.Bytes())
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

// Reads 16 bytes starting from the given page and stores them into a buffer.
// The buffer can then be read out with GetPage.
// How many pages are read depends on the tag type. The page sizes are
// as follows:
//
// * Mifare Classic page size: 16 byte (one page is read)
// * NFC Forum Type 1 page size: 8 byte (two pages are read)
// * NFC Forum Type 2 page size: 4 byte (four pages are read)
//
// The general approach for reading a tag is as follows:
//
// 1. Call RequestTagID
// 2. Wait for state to change to *RequestTagIDReady* (see GetState
//    or RegisterStateChangedCallback callback)
// 3. If looking for a specific tag then call GetTagID and check if the
//    expected tag was found, if it was not found got back to step 1
// 4. Call RequestPage with page number
// 5. Wait for state to change to *RequestPageReady* (see GetState
//    or RegisterStateChangedCallback callback)
// 6. Call GetPage to retrieve the page from the buffer
//
// If you use a Mifare Classic tag you have to authenticate a page before you
// can read it. See AuthenticateMifareClassicPage.
func (device *NFCRFIDBricklet) RequestPage(page uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, page)

	resultBytes, err := device.device.Set(uint8(FunctionRequestPage), buf.Bytes())
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

// Returns 16 bytes of data from an internal buffer. To fill the buffer
// with specific pages you have to call RequestPage beforehand.
func (device *NFCRFIDBricklet) GetPage() (data [16]uint8, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetPage), buf.Bytes())
	if err != nil {
		return data, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 24 {
			return data, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 24)
		}

		if header.ErrorCode != 0 {
			return data, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &data)

	}

	return data, nil
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
func (device *NFCRFIDBricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
