/* ***********************************************************
 * This file was automatically generated on 2019-01-29.      *
 *                                                           *
 * Go Bindings Version 2.0.2                                 *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


//Generates configurable DC voltage between 0V and 5V‍.
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/AnalogOut_Bricklet_Go.html.
package analog_out_bricklet

import (
	"encoding/binary"
	"bytes"
    . "github.com/Tinkerforge/go-api-bindings/internal"
    "github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function uint8

const (
    FunctionSetVoltage Function = 1
	FunctionGetVoltage Function = 2
	FunctionSetMode Function = 3
	FunctionGetMode Function = 4
	FunctionGetIdentity Function = 255
)

type Mode uint8

const (
    ModeAnalogValue Mode = 0
	Mode1kToGround Mode = 1
	Mode100kToGround Mode = 2
	Mode500kToGround Mode = 3
)

type AnalogOutBricklet struct{
	device Device
}
const DeviceIdentifier = 220
const DeviceDisplayName = "Analog Out Bricklet"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (AnalogOutBricklet, error) {
    internalIPCon := ipcon.GetInternalHandle().(IPConnection)
    dev, err := NewDevice([3]uint8{ 2,0,0 }, uid, &internalIPCon, 0)
    if err != nil {
        return AnalogOutBricklet{}, err
    }
    dev.ResponseExpected[FunctionSetVoltage] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetVoltage] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetMode] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetMode] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetIdentity] = ResponseExpectedFlagAlwaysTrue;
    return AnalogOutBricklet{dev}, nil
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
func (device *AnalogOutBricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *AnalogOutBricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
    return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *AnalogOutBricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *AnalogOutBricklet) GetAPIVersion() [3]uint8 {
    return device.device.GetAPIVersion()
}

// Sets the voltage in mV. The possible range is 0V to 5V (0-5000).
	// Calling this function will set the mode to 0 (see SetMode).
	// 
	// The default value is 0 (with mode 1).
func (device *AnalogOutBricklet) SetVoltage(voltage uint16) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, voltage);

    resultBytes, err := device.device.Set(uint8(FunctionSetVoltage), buf.Bytes())
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

// Returns the voltage as set by SetVoltage.
func (device *AnalogOutBricklet) GetVoltage() (voltage uint16, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetVoltage), buf.Bytes())
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

// Sets the mode of the analog value. Possible modes:
	// 
	// * 0: Normal Mode (Analog value as set by SetVoltage is applied)
	// * 1: 1k Ohm resistor to ground
	// * 2: 100k Ohm resistor to ground
	// * 3: 500k Ohm resistor to ground
	// 
	// Setting the mode to 0 will result in an output voltage of 0. You can jump
	// to a higher output voltage directly by calling SetVoltage.
	// 
	// The default mode is 1.
//
// Associated constants:
//
//	* ModeAnalogValue
//	* Mode1kToGround
//	* Mode100kToGround
//	* Mode500kToGround
func (device *AnalogOutBricklet) SetMode(mode Mode) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, mode);

    resultBytes, err := device.device.Set(uint8(FunctionSetMode), buf.Bytes())
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

// Returns the mode as set by SetMode.
//
// Associated constants:
//
//	* ModeAnalogValue
//	* Mode1kToGround
//	* Mode100kToGround
//	* Mode500kToGround
func (device *AnalogOutBricklet) GetMode() (mode Mode, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetMode), buf.Bytes())
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

// Returns the UID, the UID where the Bricklet is connected to,
	// the position, the hardware and firmware version as well as the
	// device identifier.
	// 
	// The position can be 'a', 'b', 'c' or 'd'.
	// 
	// The device identifier numbers can be found `here <device_identifier>`.
	// |device_identifier_constant|
func (device *AnalogOutBricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {    
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
