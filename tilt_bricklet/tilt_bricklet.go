/* ***********************************************************
 * This file was automatically generated on 2018-12-21.      *
 *                                                           *
 * Go Bindings Version 2.0.0                                 *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


//Detects inclination of Bricklet (tilt switch open/closed).
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/Tilt_Bricklet_Go.html.
package tilt_bricklet

import (
	"encoding/binary"
	"bytes"
    . "github.com/Tinkerforge/go-api-bindings/internal"
    "github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function uint8

const (
    FunctionGetTiltState Function = 1
	FunctionEnableTiltStateCallback Function = 2
	FunctionDisableTiltStateCallback Function = 3
	FunctionIsTiltStateCallbackEnabled Function = 4
	FunctionGetIdentity Function = 255
	FunctionCallbackTiltState Function = 5
)

type TiltState uint8

const (
    TiltStateClosed TiltState = 0
	TiltStateOpen TiltState = 1
	TiltStateClosedVibrating TiltState = 2
)

type TiltBricklet struct{
	device Device
}
const DeviceIdentifier = 239
const DeviceDisplayName = "Tilt Bricklet"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (TiltBricklet, error) {
    internalIPCon := ipcon.GetInternalHandle().(IPConnection)
    dev, err := NewDevice([3]uint8{ 2,0,0 }, uid, &internalIPCon, 0)
    if err != nil {
        return TiltBricklet{}, err
    }
    dev.ResponseExpected[FunctionGetTiltState] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionEnableTiltStateCallback] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionDisableTiltStateCallback] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionIsTiltStateCallbackEnabled] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetIdentity] = ResponseExpectedFlagAlwaysTrue;
    return TiltBricklet{dev}, nil
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
func (device *TiltBricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *TiltBricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
    return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *TiltBricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *TiltBricklet) GetAPIVersion() [3]uint8 {
    return device.device.GetAPIVersion()
}

// This callback provides the current tilt state. It is called every time the
	// state changes.
	// 
	// See GetTiltState for a description of the states.
func (device *TiltBricklet) RegisterTiltStateCallback(fn func(TiltState)) uint64 {
            wrapper := func(byteSlice []byte) {
                buf := bytes.NewBuffer(byteSlice[8:])
                var state TiltState
                binary.Read(buf, binary.LittleEndian, &state)
                fn(state)
            }
    return device.device.RegisterCallback(uint8(FunctionCallbackTiltState), wrapper)
}

//Remove a registered Tilt State callback.
func (device *TiltBricklet) DeregisterTiltStateCallback(callbackID uint64) {
    device.device.DeregisterCallback(uint8(FunctionCallbackTiltState), callbackID)
}


// Returns the current tilt state. The state can either be
	// 
	// * 0 = Closed: The ball in the tilt switch closes the circuit.
	// * 1 = Open: The ball in the tilt switch does not close the circuit.
	// * 2 = Closed Vibrating: The tilt switch is in motion (rapid change between open and close).
	// 
	// .. image:: /Images/Bricklets/bricklet_tilt_mechanics.jpg
	//    :scale: 100 %
	//    :alt: Tilt states
	//    :align: center
	//    :target: ../../_images/Bricklets/bricklet_tilt_mechanics.jpg
//
// Associated constants:
//
//	* TiltStateClosed
//	* TiltStateOpen
//	* TiltStateClosedVibrating
func (device *TiltBricklet) GetTiltState() (state TiltState, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetTiltState), buf.Bytes())
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

// Enables the RegisterTiltStateCallback callback.
func (device *TiltBricklet) EnableTiltStateCallback() (err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Set(uint8(FunctionEnableTiltStateCallback), buf.Bytes())
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

// Disables the RegisterTiltStateCallback callback.
func (device *TiltBricklet) DisableTiltStateCallback() (err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Set(uint8(FunctionDisableTiltStateCallback), buf.Bytes())
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

// Returns *true* if the RegisterTiltStateCallback callback is enabled.
func (device *TiltBricklet) IsTiltStateCallbackEnabled() (enabled bool, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionIsTiltStateCallbackEnabled), buf.Bytes())
    if err != nil {
        return enabled, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return enabled, BrickletError(header.ErrorCode)
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
	// The position can be 'a', 'b', 'c' or 'd'.
	// 
	// The device identifier numbers can be found `here <device_identifier>`.
	// |device_identifier_constant|
func (device *TiltBricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {    
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
