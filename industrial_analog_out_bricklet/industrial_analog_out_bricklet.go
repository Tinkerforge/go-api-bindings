/* ***********************************************************
 * This file was automatically generated on 2019-01-29.      *
 *                                                           *
 * Go Bindings Version 2.0.2                                 *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


//Generates configurable DC voltage and current, 0V to 10V and 4mA to 20mA‍.
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/IndustrialAnalogOut_Bricklet_Go.html.
package industrial_analog_out_bricklet

import (
	"encoding/binary"
	"bytes"
    . "github.com/Tinkerforge/go-api-bindings/internal"
    "github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function uint8

const (
    FunctionEnable Function = 1
	FunctionDisable Function = 2
	FunctionIsEnabled Function = 3
	FunctionSetVoltage Function = 4
	FunctionGetVoltage Function = 5
	FunctionSetCurrent Function = 6
	FunctionGetCurrent Function = 7
	FunctionSetConfiguration Function = 8
	FunctionGetConfiguration Function = 9
	FunctionGetIdentity Function = 255
)

type VoltageRange uint8

const (
    VoltageRange0To5V VoltageRange = 0
	VoltageRange0To10V VoltageRange = 1
)

type CurrentRange uint8

const (
    CurrentRange4To20mA CurrentRange = 0
	CurrentRange0To20mA CurrentRange = 1
	CurrentRange0To24ma CurrentRange = 2
)

type IndustrialAnalogOutBricklet struct{
	device Device
}
const DeviceIdentifier = 258
const DeviceDisplayName = "Industrial Analog Out Bricklet"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (IndustrialAnalogOutBricklet, error) {
    internalIPCon := ipcon.GetInternalHandle().(IPConnection)
    dev, err := NewDevice([3]uint8{ 2,0,0 }, uid, &internalIPCon, 0)
    if err != nil {
        return IndustrialAnalogOutBricklet{}, err
    }
    dev.ResponseExpected[FunctionEnable] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionDisable] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionIsEnabled] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetVoltage] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetVoltage] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetCurrent] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetCurrent] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetConfiguration] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetIdentity] = ResponseExpectedFlagAlwaysTrue;
    return IndustrialAnalogOutBricklet{dev}, nil
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
func (device *IndustrialAnalogOutBricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *IndustrialAnalogOutBricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
    return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *IndustrialAnalogOutBricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *IndustrialAnalogOutBricklet) GetAPIVersion() [3]uint8 {
    return device.device.GetAPIVersion()
}

// Enables the output of voltage and current.
	// 
	// The default is disabled.
func (device *IndustrialAnalogOutBricklet) Enable() (err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Set(uint8(FunctionEnable), buf.Bytes())
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

// Disables the output of voltage and current.
	// 
	// The default is disabled.
func (device *IndustrialAnalogOutBricklet) Disable() (err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Set(uint8(FunctionDisable), buf.Bytes())
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

// Returns *true* if output of voltage and current is enabled, *false* otherwise.
func (device *IndustrialAnalogOutBricklet) IsEnabled() (enabled bool, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionIsEnabled), buf.Bytes())
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

// Sets the output voltage in mV.
	// 
	// The output voltage and output current are linked. Changing the output voltage
	// also changes the output current.
func (device *IndustrialAnalogOutBricklet) SetVoltage(voltage uint16) (err error) {    
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
func (device *IndustrialAnalogOutBricklet) GetVoltage() (voltage uint16, err error) {    
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

// Sets the output current in µA.
	// 
	// The output current and output voltage are linked. Changing the output current
	// also changes the output voltage.
func (device *IndustrialAnalogOutBricklet) SetCurrent(current uint16) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, current);

    resultBytes, err := device.device.Set(uint8(FunctionSetCurrent), buf.Bytes())
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

// Returns the current as set by SetCurrent.
func (device *IndustrialAnalogOutBricklet) GetCurrent() (current uint16, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetCurrent), buf.Bytes())
    if err != nil {
        return current, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return current, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &current)

    }
    
    return current, nil
}

// Configures the voltage and current range.
	// 
	// Possible voltage ranges are:
	// 
	// * 0V to 5V
	// * 0V to 10V (default)
	// 
	// Possible current ranges are:
	// 
	// * 4mA to 20mA (default)
	// * 0mA to 20mA
	// * 0mA to 24mA
	// 
	// The resolution will always be 12 bit. This means, that the
	// precision is higher with a smaller range.
//
// Associated constants:
//
//	* VoltageRange0To5V
//	* VoltageRange0To10V
//	* CurrentRange4To20mA
//	* CurrentRange0To20mA
//	* CurrentRange0To24ma
func (device *IndustrialAnalogOutBricklet) SetConfiguration(voltageRange VoltageRange, currentRange CurrentRange) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, voltageRange);
	binary.Write(&buf, binary.LittleEndian, currentRange);

    resultBytes, err := device.device.Set(uint8(FunctionSetConfiguration), buf.Bytes())
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

// Returns the configuration as set by SetConfiguration.
//
// Associated constants:
//
//	* VoltageRange0To5V
//	* VoltageRange0To10V
//	* CurrentRange4To20mA
//	* CurrentRange0To20mA
//	* CurrentRange0To24ma
func (device *IndustrialAnalogOutBricklet) GetConfiguration() (voltageRange VoltageRange, currentRange CurrentRange, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetConfiguration), buf.Bytes())
    if err != nil {
        return voltageRange, currentRange, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return voltageRange, currentRange, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &voltageRange)
	binary.Read(resultBuf, binary.LittleEndian, &currentRange)

    }
    
    return voltageRange, currentRange, nil
}

// Returns the UID, the UID where the Bricklet is connected to,
	// the position, the hardware and firmware version as well as the
	// device identifier.
	// 
	// The position can be 'a', 'b', 'c' or 'd'.
	// 
	// The device identifier numbers can be found `here <device_identifier>`.
	// |device_identifier_constant|
func (device *IndustrialAnalogOutBricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {    
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
