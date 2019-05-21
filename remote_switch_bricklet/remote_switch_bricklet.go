/* ***********************************************************
 * This file was automatically generated on 2019-05-21.      *
 *                                                           *
 * Go Bindings Version 2.0.3                                 *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


//Controls remote mains switches.
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/RemoteSwitch_Bricklet_Go.html.
package remote_switch_bricklet

import (
	"encoding/binary"
	"bytes"
    . "github.com/Tinkerforge/go-api-bindings/internal"
    "github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function uint8

const (
    FunctionSwitchSocket Function = 1
	FunctionGetSwitchingState Function = 2
	FunctionSetRepeats Function = 4
	FunctionGetRepeats Function = 5
	FunctionSwitchSocketA Function = 6
	FunctionSwitchSocketB Function = 7
	FunctionDimSocketB Function = 8
	FunctionSwitchSocketC Function = 9
	FunctionGetIdentity Function = 255
	FunctionCallbackSwitchingDone Function = 3
)

type SwitchTo uint8

const (
    SwitchToOff SwitchTo = 0
	SwitchToOn SwitchTo = 1
)

type SwitchingState uint8

const (
    SwitchingStateReady SwitchingState = 0
	SwitchingStateBusy SwitchingState = 1
)

type RemoteSwitchBricklet struct{
	device Device
}
const DeviceIdentifier = 235
const DeviceDisplayName = "Remote Switch Bricklet"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (RemoteSwitchBricklet, error) {
    internalIPCon := ipcon.GetInternalHandle().(IPConnection)
    dev, err := NewDevice([3]uint8{ 2,0,1 }, uid, &internalIPCon, 0)
    if err != nil {
        return RemoteSwitchBricklet{}, err
    }
    dev.ResponseExpected[FunctionSwitchSocket] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetSwitchingState] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetRepeats] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetRepeats] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSwitchSocketA] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionSwitchSocketB] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionDimSocketB] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionSwitchSocketC] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetIdentity] = ResponseExpectedFlagAlwaysTrue;
    return RemoteSwitchBricklet{dev}, nil
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
func (device *RemoteSwitchBricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *RemoteSwitchBricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
    return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *RemoteSwitchBricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *RemoteSwitchBricklet) GetAPIVersion() [3]uint8 {
    return device.device.GetAPIVersion()
}

// This callback is triggered whenever the switching state changes
// from busy to ready, see GetSwitchingState.
func (device *RemoteSwitchBricklet) RegisterSwitchingDoneCallback(fn func()) uint64 {
            wrapper := func(byteSlice []byte) {
                
                
                
                fn()
            }
    return device.device.RegisterCallback(uint8(FunctionCallbackSwitchingDone), wrapper)
}

//Remove a registered Switching Done callback.
func (device *RemoteSwitchBricklet) DeregisterSwitchingDoneCallback(registrationID uint64) {
    device.device.DeregisterCallback(uint8(FunctionCallbackSwitchingDone), registrationID)
}


// This function is deprecated, use SwitchSocketA instead.
//
// Associated constants:
//
//	* SwitchToOff
//	* SwitchToOn
func (device *RemoteSwitchBricklet) SwitchSocket(houseCode uint8, receiverCode uint8, switchTo SwitchTo) (err error) {
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, houseCode);
	binary.Write(&buf, binary.LittleEndian, receiverCode);
	binary.Write(&buf, binary.LittleEndian, switchTo);

    resultBytes, err := device.device.Set(uint8(FunctionSwitchSocket), buf.Bytes())
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

// Returns the current switching state. If the current state is busy, the
// Bricklet is currently sending a code to switch a socket. It will not
// accept any calls of SwitchSocket until the state changes to ready.
// 
// How long the switching takes is dependent on the number of repeats, see
// SetRepeats.
//
// Associated constants:
//
//	* SwitchingStateReady
//	* SwitchingStateBusy
func (device *RemoteSwitchBricklet) GetSwitchingState() (state SwitchingState, err error) {
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetSwitchingState), buf.Bytes())
    if err != nil {
        return state, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader

        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return state, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &state)

    }

    return state, nil
}

// Sets the number of times the code is send when of the SwitchSocket
// functions is called. The repeats basically correspond to the amount of time
// that a button of the remote is pressed.
// 
// Some dimmers are controlled by the length of a button pressed,
// this can be simulated by increasing the repeats.
// 
// The default value is 5.
func (device *RemoteSwitchBricklet) SetRepeats(repeats uint8) (err error) {
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, repeats);

    resultBytes, err := device.device.Set(uint8(FunctionSetRepeats), buf.Bytes())
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

// Returns the number of repeats as set by SetRepeats.
func (device *RemoteSwitchBricklet) GetRepeats() (repeats uint8, err error) {
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetRepeats), buf.Bytes())
    if err != nil {
        return repeats, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader

        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return repeats, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &repeats)

    }

    return repeats, nil
}

// To switch a type A socket you have to give the house code, receiver code and the
// state (on or off) you want to switch to.
// 
// The house code and receiver code have a range of 0 to 31 (5bit).
// 
// A detailed description on how you can figure out the house and receiver code
// can be found `here <remote_switch_bricklet_type_a_house_and_receiver_code>`.
// 
// .. versionadded:: 2.0.1$nbsp;(Plugin)
//
// Associated constants:
//
//	* SwitchToOff
//	* SwitchToOn
func (device *RemoteSwitchBricklet) SwitchSocketA(houseCode uint8, receiverCode uint8, switchTo SwitchTo) (err error) {
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, houseCode);
	binary.Write(&buf, binary.LittleEndian, receiverCode);
	binary.Write(&buf, binary.LittleEndian, switchTo);

    resultBytes, err := device.device.Set(uint8(FunctionSwitchSocketA), buf.Bytes())
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

// To switch a type B socket you have to give the address, unit and the state
// (on or off) you want to switch to.
// 
// The address has a range of 0 to 67108863 (26bit) and the unit has a range
// of 0 to 15 (4bit). To switch all devices with the same address use 255 for
// the unit.
// 
// A detailed description on how you can teach a socket the address and unit can
// be found `here <remote_switch_bricklet_type_b_address_and_unit>`.
// 
// .. versionadded:: 2.0.1$nbsp;(Plugin)
//
// Associated constants:
//
//	* SwitchToOff
//	* SwitchToOn
func (device *RemoteSwitchBricklet) SwitchSocketB(address uint32, unit uint8, switchTo SwitchTo) (err error) {
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, address);
	binary.Write(&buf, binary.LittleEndian, unit);
	binary.Write(&buf, binary.LittleEndian, switchTo);

    resultBytes, err := device.device.Set(uint8(FunctionSwitchSocketB), buf.Bytes())
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

// To control a type B dimmer you have to give the address, unit and the
// dim value you want to set the dimmer to.
// 
// The address has a range of 0 to 67108863 (26bit), the unit and the dim value
// has a range of 0 to 15 (4bit).
// 
// A detailed description on how you can teach a dimmer the address and unit can
// be found `here <remote_switch_bricklet_type_b_address_and_unit>`.
// 
// .. versionadded:: 2.0.1$nbsp;(Plugin)
func (device *RemoteSwitchBricklet) DimSocketB(address uint32, unit uint8, dimValue uint8) (err error) {
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, address);
	binary.Write(&buf, binary.LittleEndian, unit);
	binary.Write(&buf, binary.LittleEndian, dimValue);

    resultBytes, err := device.device.Set(uint8(FunctionDimSocketB), buf.Bytes())
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

// To switch a type C socket you have to give the system code, device code and the
// state (on or off) you want to switch to.
// 
// The system code has a range of 'A' to 'P' (4bit) and the device code has a
// range of 1 to 16 (4bit).
// 
// A detailed description on how you can figure out the system and device code
// can be found `here <remote_switch_bricklet_type_c_system_and_device_code>`.
// 
// .. versionadded:: 2.0.1$nbsp;(Plugin)
//
// Associated constants:
//
//	* SwitchToOff
//	* SwitchToOn
func (device *RemoteSwitchBricklet) SwitchSocketC(systemCode rune, deviceCode uint8, switchTo SwitchTo) (err error) {
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, systemCode);
	binary.Write(&buf, binary.LittleEndian, deviceCode);
	binary.Write(&buf, binary.LittleEndian, switchTo);

    resultBytes, err := device.device.Set(uint8(FunctionSwitchSocketC), buf.Bytes())
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

// Returns the UID, the UID where the Bricklet is connected to,
// the position, the hardware and firmware version as well as the
// device identifier.
// 
// The position can be 'a', 'b', 'c' or 'd'.
// 
// The device identifier numbers can be found `here <device_identifier>`.
// |device_identifier_constant|
func (device *RemoteSwitchBricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
