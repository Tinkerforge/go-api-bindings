/* ***********************************************************
 * This file was automatically generated on 2022-08-22.      *
 *                                                           *
 * Go Bindings Version 2.0.14                                *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/

// Controls remote mains switches and receives signals from remotes.
//
//
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/RemoteSwitchV2_Bricklet_Go.html.
package remote_switch_v2_bricklet

import (
	"bytes"
	"encoding/binary"
	"fmt"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionGetSwitchingState       Function = 1
	FunctionSetRepeats              Function = 3
	FunctionGetRepeats              Function = 4
	FunctionSwitchSocketA           Function = 5
	FunctionSwitchSocketB           Function = 6
	FunctionDimSocketB              Function = 7
	FunctionSwitchSocketC           Function = 8
	FunctionSetRemoteConfiguration  Function = 9
	FunctionGetRemoteConfiguration  Function = 10
	FunctionGetRemoteStatusA        Function = 11
	FunctionGetRemoteStatusB        Function = 12
	FunctionGetRemoteStatusC        Function = 13
	FunctionGetSPITFPErrorCount     Function = 234
	FunctionSetBootloaderMode       Function = 235
	FunctionGetBootloaderMode       Function = 236
	FunctionSetWriteFirmwarePointer Function = 237
	FunctionWriteFirmware           Function = 238
	FunctionSetStatusLEDConfig      Function = 239
	FunctionGetStatusLEDConfig      Function = 240
	FunctionGetChipTemperature      Function = 242
	FunctionReset                   Function = 243
	FunctionWriteUID                Function = 248
	FunctionReadUID                 Function = 249
	FunctionGetIdentity             Function = 255
	FunctionCallbackSwitchingDone   Function = 2
	FunctionCallbackRemoteStatusA   Function = 14
	FunctionCallbackRemoteStatusB   Function = 15
	FunctionCallbackRemoteStatusC   Function = 16
)

type SwitchingState = uint8

const (
	SwitchingStateReady SwitchingState = 0
	SwitchingStateBusy  SwitchingState = 1
)

type SwitchTo = uint8

const (
	SwitchToOff SwitchTo = 0
	SwitchToOn  SwitchTo = 1
)

type RemoteType = uint8

const (
	RemoteTypeA RemoteType = 0
	RemoteTypeB RemoteType = 1
	RemoteTypeC RemoteType = 2
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

type RemoteSwitchV2Bricklet struct {
	device Device
}

const DeviceIdentifier = 289
const DeviceDisplayName = "Remote Switch Bricklet 2.0"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (RemoteSwitchV2Bricklet, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{2, 0, 0}, uid, &internalIPCon, 0, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return RemoteSwitchV2Bricklet{}, err
	}
	dev.ResponseExpected[FunctionGetSwitchingState] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetRepeats] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetRepeats] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSwitchSocketA] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionSwitchSocketB] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionDimSocketB] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionSwitchSocketC] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionSetRemoteConfiguration] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetRemoteConfiguration] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetRemoteStatusA] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetRemoteStatusB] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetRemoteStatusC] = ResponseExpectedFlagAlwaysTrue
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
	return RemoteSwitchV2Bricklet{dev}, nil
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
func (device *RemoteSwitchV2Bricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *RemoteSwitchV2Bricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *RemoteSwitchV2Bricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *RemoteSwitchV2Bricklet) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// This callback is triggered whenever the switching state changes
// from busy to ready, see GetSwitchingState.
func (device *RemoteSwitchV2Bricklet) RegisterSwitchingDoneCallback(fn func()) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 8 {
			return
		}

		fn()
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackSwitchingDone), wrapper)
}

// Remove a registered Switching Done callback.
func (device *RemoteSwitchV2Bricklet) DeregisterSwitchingDoneCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackSwitchingDone), registrationId)
}

// Returns the house code, receiver code, switch state (on/off) and number of repeats for
// remote type A.
//
// The repeats are the number of received identical data packets. The longer the button is pressed,
// the higher the repeat number. The callback is triggered with every repeat.
//
// You have to enable the callback with SetRemoteConfiguration. The number
// of repeats that you can set in the configuration is the minimum number of repeats that have
// to be seen before the callback is triggered for the first time.
func (device *RemoteSwitchV2Bricklet) RegisterRemoteStatusACallback(fn func(uint8, uint8, SwitchTo, uint16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 13 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var houseCode uint8
		var receiverCode uint8
		var switchTo SwitchTo
		var repeats uint16
		binary.Read(buf, binary.LittleEndian, &houseCode)
		binary.Read(buf, binary.LittleEndian, &receiverCode)
		binary.Read(buf, binary.LittleEndian, &switchTo)
		binary.Read(buf, binary.LittleEndian, &repeats)
		fn(houseCode, receiverCode, switchTo, repeats)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackRemoteStatusA), wrapper)
}

// Remove a registered Remote Status A callback.
func (device *RemoteSwitchV2Bricklet) DeregisterRemoteStatusACallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackRemoteStatusA), registrationId)
}

// Returns the address (unique per remote), unit (button number), switch state (on/off) and number of repeats for
// remote type B.
//
// If the remote supports dimming the dim value is used instead of the switch state.
//
// The repeats are the number of received identical data packets. The longer the button is pressed,
// the higher the repeat number. The callback is triggered with every repeat.
//
// You have to enable the callback with SetRemoteConfiguration. The number
// of repeats that you can set in the configuration is the minimum number of repeats that have
// to be seen before the callback is triggered for the first time.
func (device *RemoteSwitchV2Bricklet) RegisterRemoteStatusBCallback(fn func(uint32, uint8, SwitchTo, uint8, uint16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 17 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var address uint32
		var unit uint8
		var switchTo SwitchTo
		var dimValue uint8
		var repeats uint16
		binary.Read(buf, binary.LittleEndian, &address)
		binary.Read(buf, binary.LittleEndian, &unit)
		binary.Read(buf, binary.LittleEndian, &switchTo)
		binary.Read(buf, binary.LittleEndian, &dimValue)
		binary.Read(buf, binary.LittleEndian, &repeats)
		fn(address, unit, switchTo, dimValue, repeats)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackRemoteStatusB), wrapper)
}

// Remove a registered Remote Status B callback.
func (device *RemoteSwitchV2Bricklet) DeregisterRemoteStatusBCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackRemoteStatusB), registrationId)
}

// Returns the system code, device code, switch state (on/off) and number of repeats for
// remote type C.
//
// The repeats are the number of received identical data packets. The longer the button is pressed,
// the higher the repeat number. The callback is triggered with every repeat.
//
// You have to enable the callback with SetRemoteConfiguration. The number
// of repeats that you can set in the configuration is the minimum number of repeats that have
// to be seen before the callback is triggered for the first time.
func (device *RemoteSwitchV2Bricklet) RegisterRemoteStatusCCallback(fn func(rune, uint8, SwitchTo, uint16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 13 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var systemCode rune
		var deviceCode uint8
		var switchTo SwitchTo
		var repeats uint16
		systemCode = rune(buf.Next(1)[0])
		binary.Read(buf, binary.LittleEndian, &deviceCode)
		binary.Read(buf, binary.LittleEndian, &switchTo)
		binary.Read(buf, binary.LittleEndian, &repeats)
		fn(systemCode, deviceCode, switchTo, repeats)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackRemoteStatusC), wrapper)
}

// Remove a registered Remote Status C callback.
func (device *RemoteSwitchV2Bricklet) DeregisterRemoteStatusCCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackRemoteStatusC), registrationId)
}

// Returns the current switching state. If the current state is busy, the
// Bricklet is currently sending a code to switch a socket. It will not
// accept any calls of switch socket functions until the state changes to ready.
//
// How long the switching takes is dependent on the number of repeats, see
// SetRepeats.
//
// Associated constants:
//
//	* SwitchingStateReady
//	* SwitchingStateBusy
func (device *RemoteSwitchV2Bricklet) GetSwitchingState() (state SwitchingState, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetSwitchingState), buf.Bytes())
	if err != nil {
		return state, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return state, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return state, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &state)

	}

	return state, nil
}

// Sets the number of times the code is sent when one of the Switch Socket
// functions is called. The repeats basically correspond to the amount of time
// that a button of the remote is pressed.
//
// Some dimmers are controlled by the length of a button pressed,
// this can be simulated by increasing the repeats.
func (device *RemoteSwitchV2Bricklet) SetRepeats(repeats uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, repeats)

	resultBytes, err := device.device.Set(uint8(FunctionSetRepeats), buf.Bytes())
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

// Returns the number of repeats as set by SetRepeats.
func (device *RemoteSwitchV2Bricklet) GetRepeats() (repeats uint8, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetRepeats), buf.Bytes())
	if err != nil {
		return repeats, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return repeats, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return repeats, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &repeats)

	}

	return repeats, nil
}

// To switch a type A socket you have to give the house code, receiver code and the
// state (on or off) you want to switch to.
//
// A detailed description on how you can figure out the house and receiver code
// can be found `here <remote_switch_bricklet_type_a_house_and_receiver_code>`.
//
// Associated constants:
//
//	* SwitchToOff
//	* SwitchToOn
func (device *RemoteSwitchV2Bricklet) SwitchSocketA(houseCode uint8, receiverCode uint8, switchTo SwitchTo) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, houseCode)
	binary.Write(&buf, binary.LittleEndian, receiverCode)
	binary.Write(&buf, binary.LittleEndian, switchTo)

	resultBytes, err := device.device.Set(uint8(FunctionSwitchSocketA), buf.Bytes())
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

// To switch a type B socket you have to give the address, unit and the state
// (on or off) you want to switch to.
//
// To switch all devices with the same address use 255 for the unit.
//
// A detailed description on how you can teach a socket the address and unit can
// be found `here <remote_switch_bricklet_type_b_address_and_unit>`.
//
// Associated constants:
//
//	* SwitchToOff
//	* SwitchToOn
func (device *RemoteSwitchV2Bricklet) SwitchSocketB(address uint32, unit uint8, switchTo SwitchTo) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, address)
	binary.Write(&buf, binary.LittleEndian, unit)
	binary.Write(&buf, binary.LittleEndian, switchTo)

	resultBytes, err := device.device.Set(uint8(FunctionSwitchSocketB), buf.Bytes())
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

// To control a type B dimmer you have to give the address, unit and the
// dim value you want to set the dimmer to.
//
// A detailed description on how you can teach a dimmer the address and unit can
// be found `here <remote_switch_bricklet_type_b_address_and_unit>`.
func (device *RemoteSwitchV2Bricklet) DimSocketB(address uint32, unit uint8, dimValue uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, address)
	binary.Write(&buf, binary.LittleEndian, unit)
	binary.Write(&buf, binary.LittleEndian, dimValue)

	resultBytes, err := device.device.Set(uint8(FunctionDimSocketB), buf.Bytes())
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

// To switch a type C socket you have to give the system code, device code and the
// state (on or off) you want to switch to.
//
// A detailed description on how you can figure out the system and device code
// can be found `here <remote_switch_bricklet_type_c_system_and_device_code>`.
//
// Associated constants:
//
//	* SwitchToOff
//	* SwitchToOn
func (device *RemoteSwitchV2Bricklet) SwitchSocketC(systemCode rune, deviceCode uint8, switchTo SwitchTo) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, systemCode)
	binary.Write(&buf, binary.LittleEndian, deviceCode)
	binary.Write(&buf, binary.LittleEndian, switchTo)

	resultBytes, err := device.device.Set(uint8(FunctionSwitchSocketC), buf.Bytes())
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

// Sets the configuration for **receiving** data from a remote of type A, B or C.
//
// * Remote Type: A, B or C depending on the type of remote you want to receive.
// * Minimum Repeats: The minimum number of repeated data packets until the callback
//   is triggered (if enabled).
// * Callback Enabled: Enable or disable callback (see RegisterRemoteStatusACallback callback,
//   RegisterRemoteStatusBCallback callback and RegisterRemoteStatusCCallback callback).
//
// Associated constants:
//
//	* RemoteTypeA
//	* RemoteTypeB
//	* RemoteTypeC
func (device *RemoteSwitchV2Bricklet) SetRemoteConfiguration(remoteType RemoteType, minimumRepeats uint16, callbackEnabled bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, remoteType)
	binary.Write(&buf, binary.LittleEndian, minimumRepeats)
	binary.Write(&buf, binary.LittleEndian, callbackEnabled)

	resultBytes, err := device.device.Set(uint8(FunctionSetRemoteConfiguration), buf.Bytes())
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

// Returns the remote configuration as set by SetRemoteConfiguration
//
// Associated constants:
//
//	* RemoteTypeA
//	* RemoteTypeB
//	* RemoteTypeC
func (device *RemoteSwitchV2Bricklet) GetRemoteConfiguration() (remoteType RemoteType, minimumRepeats uint16, callbackEnabled bool, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetRemoteConfiguration), buf.Bytes())
	if err != nil {
		return remoteType, minimumRepeats, callbackEnabled, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 12 {
			return remoteType, minimumRepeats, callbackEnabled, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		if header.ErrorCode != 0 {
			return remoteType, minimumRepeats, callbackEnabled, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &remoteType)
		binary.Read(resultBuf, binary.LittleEndian, &minimumRepeats)
		binary.Read(resultBuf, binary.LittleEndian, &callbackEnabled)

	}

	return remoteType, minimumRepeats, callbackEnabled, nil
}

// Returns the house code, receiver code, switch state (on/off) and number of
// repeats for remote type A.
//
// Repeats == 0 means there was no button press. Repeats >= 1 means there
// was a button press with the specified house/receiver code. The repeats are the
// number of received identical data packets. The longer the button is pressed,
// the higher the repeat number.
//
// Use the callback to get this data automatically when a button is pressed,
// see SetRemoteConfiguration and RegisterRemoteStatusACallback callback.
//
// Associated constants:
//
//	* SwitchToOff
//	* SwitchToOn
func (device *RemoteSwitchV2Bricklet) GetRemoteStatusA() (houseCode uint8, receiverCode uint8, switchTo SwitchTo, repeats uint16, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetRemoteStatusA), buf.Bytes())
	if err != nil {
		return houseCode, receiverCode, switchTo, repeats, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 13 {
			return houseCode, receiverCode, switchTo, repeats, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 13)
		}

		if header.ErrorCode != 0 {
			return houseCode, receiverCode, switchTo, repeats, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &houseCode)
		binary.Read(resultBuf, binary.LittleEndian, &receiverCode)
		binary.Read(resultBuf, binary.LittleEndian, &switchTo)
		binary.Read(resultBuf, binary.LittleEndian, &repeats)

	}

	return houseCode, receiverCode, switchTo, repeats, nil
}

// Returns the address (unique per remote), unit (button number), switch state
// (on/off) and number of repeats for remote type B.
//
// If the remote supports dimming the dim value is used instead of the switch state.
//
// If repeats=0 there was no button press. If repeats >= 1 there
// was a button press with the specified address/unit. The repeats are the number of received
// identical data packets. The longer the button is pressed, the higher the repeat number.
//
// Use the callback to get this data automatically when a button is pressed,
// see SetRemoteConfiguration and RegisterRemoteStatusBCallback callback.
//
// Associated constants:
//
//	* SwitchToOff
//	* SwitchToOn
func (device *RemoteSwitchV2Bricklet) GetRemoteStatusB() (address uint32, unit uint8, switchTo SwitchTo, dimValue uint8, repeats uint16, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetRemoteStatusB), buf.Bytes())
	if err != nil {
		return address, unit, switchTo, dimValue, repeats, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 17 {
			return address, unit, switchTo, dimValue, repeats, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 17)
		}

		if header.ErrorCode != 0 {
			return address, unit, switchTo, dimValue, repeats, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &address)
		binary.Read(resultBuf, binary.LittleEndian, &unit)
		binary.Read(resultBuf, binary.LittleEndian, &switchTo)
		binary.Read(resultBuf, binary.LittleEndian, &dimValue)
		binary.Read(resultBuf, binary.LittleEndian, &repeats)

	}

	return address, unit, switchTo, dimValue, repeats, nil
}

// Returns the system code, device code, switch state (on/off) and number of repeats for
// remote type C.
//
// If repeats=0 there was no button press. If repeats >= 1 there
// was a button press with the specified system/device code. The repeats are the number of received
// identical data packets. The longer the button is pressed, the higher the repeat number.
//
// Use the callback to get this data automatically when a button is pressed,
// see SetRemoteConfiguration and RegisterRemoteStatusCCallback callback.
//
// Associated constants:
//
//	* SwitchToOff
//	* SwitchToOn
func (device *RemoteSwitchV2Bricklet) GetRemoteStatusC() (systemCode rune, deviceCode uint8, switchTo SwitchTo, repeats uint16, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetRemoteStatusC), buf.Bytes())
	if err != nil {
		return systemCode, deviceCode, switchTo, repeats, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 13 {
			return systemCode, deviceCode, switchTo, repeats, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 13)
		}

		if header.ErrorCode != 0 {
			return systemCode, deviceCode, switchTo, repeats, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		systemCode = rune(resultBuf.Next(1)[0])
		binary.Read(resultBuf, binary.LittleEndian, &deviceCode)
		binary.Read(resultBuf, binary.LittleEndian, &switchTo)
		binary.Read(resultBuf, binary.LittleEndian, &repeats)

	}

	return systemCode, deviceCode, switchTo, repeats, nil
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
func (device *RemoteSwitchV2Bricklet) GetSPITFPErrorCount() (errorCountAckChecksum uint32, errorCountMessageChecksum uint32, errorCountFrame uint32, errorCountOverflow uint32, err error) {
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
func (device *RemoteSwitchV2Bricklet) SetBootloaderMode(mode BootloaderMode) (status BootloaderStatus, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, mode)

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
func (device *RemoteSwitchV2Bricklet) GetBootloaderMode() (mode BootloaderMode, err error) {
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
func (device *RemoteSwitchV2Bricklet) SetWriteFirmwarePointer(pointer uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, pointer)

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
func (device *RemoteSwitchV2Bricklet) WriteFirmware(data [64]uint8) (status uint8, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, data)

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
func (device *RemoteSwitchV2Bricklet) SetStatusLEDConfig(config StatusLEDConfig) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, config)

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
func (device *RemoteSwitchV2Bricklet) GetStatusLEDConfig() (config StatusLEDConfig, err error) {
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
func (device *RemoteSwitchV2Bricklet) GetChipTemperature() (temperature int16, err error) {
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
func (device *RemoteSwitchV2Bricklet) Reset() (err error) {
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
func (device *RemoteSwitchV2Bricklet) WriteUID(uid uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, uid)

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
func (device *RemoteSwitchV2Bricklet) ReadUID() (uid uint32, err error) {
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
func (device *RemoteSwitchV2Bricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
