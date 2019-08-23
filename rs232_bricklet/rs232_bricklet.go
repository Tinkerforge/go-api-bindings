/* ***********************************************************
 * This file was automatically generated on 2019-08-23.      *
 *                                                           *
 * Go Bindings Version 2.0.4                                 *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


// Communicates with RS232 devices.
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/RS232_Bricklet_Go.html.
package rs232_bricklet

import (
	"encoding/binary"
	"bytes"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionWrite Function = 1
	FunctionRead Function = 2
	FunctionEnableReadCallback Function = 3
	FunctionDisableReadCallback Function = 4
	FunctionIsReadCallbackEnabled Function = 5
	FunctionSetConfiguration Function = 6
	FunctionGetConfiguration Function = 7
	FunctionSetBreakCondition Function = 10
	FunctionGetIdentity Function = 255
	FunctionCallbackRead Function = 8
	FunctionCallbackError Function = 9
)

type Baudrate = uint8

const (
	Baudrate300 Baudrate = 0
	Baudrate600 Baudrate = 1
	Baudrate1200 Baudrate = 2
	Baudrate2400 Baudrate = 3
	Baudrate4800 Baudrate = 4
	Baudrate9600 Baudrate = 5
	Baudrate14400 Baudrate = 6
	Baudrate19200 Baudrate = 7
	Baudrate28800 Baudrate = 8
	Baudrate38400 Baudrate = 9
	Baudrate57600 Baudrate = 10
	Baudrate115200 Baudrate = 11
	Baudrate230400 Baudrate = 12
)

type Parity = uint8

const (
	ParityNone Parity = 0
	ParityOdd Parity = 1
	ParityEven Parity = 2
	ParityForcedParity1 Parity = 3
	ParityForcedParity0 Parity = 4
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

type HardwareFlowcontrol = uint8

const (
	HardwareFlowcontrolOff HardwareFlowcontrol = 0
	HardwareFlowcontrolOn HardwareFlowcontrol = 1
)

type SoftwareFlowcontrol = uint8

const (
	SoftwareFlowcontrolOff SoftwareFlowcontrol = 0
	SoftwareFlowcontrolOn SoftwareFlowcontrol = 1
)

type Error = uint8

const (
	ErrorOverrun Error = 1
	ErrorParity Error = 2
	ErrorFraming Error = 4
)

type RS232Bricklet struct {
	device Device
}
const DeviceIdentifier = 254
const DeviceDisplayName = "RS232 Bricklet"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (RS232Bricklet, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{ 2,0,2 }, uid, &internalIPCon, 0)
	if err != nil {
		return RS232Bricklet{}, err
	}
	dev.ResponseExpected[FunctionWrite] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionRead] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionEnableReadCallback] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionDisableReadCallback] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionIsReadCallbackEnabled] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetConfiguration] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetBreakCondition] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetIdentity] = ResponseExpectedFlagAlwaysTrue;
	return RS232Bricklet{dev}, nil
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
func (device *RS232Bricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *RS232Bricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *RS232Bricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *RS232Bricklet) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// This callback is called if new data is available. The message has
// a maximum size of 60 characters. The actual length of the message
// is given in addition.
// 
// To enable this callback, use EnableReadCallback.
func (device *RS232Bricklet) RegisterReadCallback(fn func([60]rune, uint8)) uint64 {
	wrapper := func(byteSlice []byte) {
		buf := bytes.NewBuffer(byteSlice[8:])
		var message [60]rune
		var length uint8
		copy(message[:], ByteSliceToRuneSlice(buf.Next(60)))
		binary.Read(buf, binary.LittleEndian, &length)
		fn(message, length)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackRead), wrapper)
}

// Remove a registered Read callback.
func (device *RS232Bricklet) DeregisterReadCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackRead), registrationId)
}


// This callback is called if an error occurs.
// Possible errors are overrun, parity or framing error.
// 
// .. versionadded:: 2.0.1$nbsp;(Plugin)
func (device *RS232Bricklet) RegisterErrorCallback(fn func(Error)) uint64 {
	wrapper := func(byteSlice []byte) {
		buf := bytes.NewBuffer(byteSlice[8:])
		var error Error
		binary.Read(buf, binary.LittleEndian, &error)
		fn(error)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackError), wrapper)
}

// Remove a registered Error callback.
func (device *RS232Bricklet) DeregisterErrorCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackError), registrationId)
}


// Writes a string of up to 60 characters to the RS232 interface. The string
// can be binary data, ASCII or similar is not necessary.
// 
// The length of the string has to be given as an additional parameter.
// 
// The return value is the number of bytes that could be written.
// 
// See SetConfiguration for configuration possibilities
// regarding baudrate, parity and so on.
func (device *RS232Bricklet) Write(message [60]rune, length uint8) (written uint8, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, message);
	binary.Write(&buf, binary.LittleEndian, length);

	resultBytes, err := device.device.Get(uint8(FunctionWrite), buf.Bytes())
	if err != nil {
		return written, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return written, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &written)

	}

	return written, nil
}

// Returns the currently buffered message. The maximum length
// of message is 60. If the length is given as 0, there was no
// new data available.
// 
// Instead of polling with this function, you can also use
// callbacks. See EnableReadCallback and RegisterReadCallback callback.
func (device *RS232Bricklet) Read() (message [60]rune, length uint8, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionRead), buf.Bytes())
	if err != nil {
		return message, length, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return message, length, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		copy(message[:], ByteSliceToRuneSlice(resultBuf.Next(60)))
	binary.Read(resultBuf, binary.LittleEndian, &length)

	}

	return message, length, nil
}

// Enables the RegisterReadCallback callback.
// 
// By default the callback is disabled.
func (device *RS232Bricklet) EnableReadCallback() (err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Set(uint8(FunctionEnableReadCallback), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
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
func (device *RS232Bricklet) DisableReadCallback() (err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Set(uint8(FunctionDisableReadCallback), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])
		
	}

	return nil
}

// Returns *true* if the RegisterReadCallback callback is enabled,
// *false* otherwise.
func (device *RS232Bricklet) IsReadCallbackEnabled() (enabled bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionIsReadCallbackEnabled), buf.Bytes())
	if err != nil {
		return enabled, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return enabled, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &enabled)

	}

	return enabled, nil
}

// Sets the configuration for the RS232 communication. Available options:
// 
// * Baudrate between 300 and 230400 baud.
// * Parity of none, odd, even or forced parity.
// * Stopbits can be 1 or 2.
// * Word length of 5 to 8.
// * Hard-/Software flow control can either be on or off but not both simultaneously on.
// 
// The default is: 115200 baud, parity none, 1 stop bit, word length 8, hard-/software flow control off.
//
// Associated constants:
//
//	* Baudrate300
//	* Baudrate600
//	* Baudrate1200
//	* Baudrate2400
//	* Baudrate4800
//	* Baudrate9600
//	* Baudrate14400
//	* Baudrate19200
//	* Baudrate28800
//	* Baudrate38400
//	* Baudrate57600
//	* Baudrate115200
//	* Baudrate230400
//	* ParityNone
//	* ParityOdd
//	* ParityEven
//	* ParityForcedParity1
//	* ParityForcedParity0
//	* Stopbits1
//	* Stopbits2
//	* Wordlength5
//	* Wordlength6
//	* Wordlength7
//	* Wordlength8
//	* HardwareFlowcontrolOff
//	* HardwareFlowcontrolOn
//	* SoftwareFlowcontrolOff
//	* SoftwareFlowcontrolOn
func (device *RS232Bricklet) SetConfiguration(baudrate Baudrate, parity Parity, stopbits Stopbits, wordlength Wordlength, hardwareFlowcontrol HardwareFlowcontrol, softwareFlowcontrol SoftwareFlowcontrol) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, baudrate);
	binary.Write(&buf, binary.LittleEndian, parity);
	binary.Write(&buf, binary.LittleEndian, stopbits);
	binary.Write(&buf, binary.LittleEndian, wordlength);
	binary.Write(&buf, binary.LittleEndian, hardwareFlowcontrol);
	binary.Write(&buf, binary.LittleEndian, softwareFlowcontrol);

	resultBytes, err := device.device.Set(uint8(FunctionSetConfiguration), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
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
//	* Baudrate300
//	* Baudrate600
//	* Baudrate1200
//	* Baudrate2400
//	* Baudrate4800
//	* Baudrate9600
//	* Baudrate14400
//	* Baudrate19200
//	* Baudrate28800
//	* Baudrate38400
//	* Baudrate57600
//	* Baudrate115200
//	* Baudrate230400
//	* ParityNone
//	* ParityOdd
//	* ParityEven
//	* ParityForcedParity1
//	* ParityForcedParity0
//	* Stopbits1
//	* Stopbits2
//	* Wordlength5
//	* Wordlength6
//	* Wordlength7
//	* Wordlength8
//	* HardwareFlowcontrolOff
//	* HardwareFlowcontrolOn
//	* SoftwareFlowcontrolOff
//	* SoftwareFlowcontrolOn
func (device *RS232Bricklet) GetConfiguration() (baudrate Baudrate, parity Parity, stopbits Stopbits, wordlength Wordlength, hardwareFlowcontrol HardwareFlowcontrol, softwareFlowcontrol SoftwareFlowcontrol, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetConfiguration), buf.Bytes())
	if err != nil {
		return baudrate, parity, stopbits, wordlength, hardwareFlowcontrol, softwareFlowcontrol, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return baudrate, parity, stopbits, wordlength, hardwareFlowcontrol, softwareFlowcontrol, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &baudrate)
	binary.Read(resultBuf, binary.LittleEndian, &parity)
	binary.Read(resultBuf, binary.LittleEndian, &stopbits)
	binary.Read(resultBuf, binary.LittleEndian, &wordlength)
	binary.Read(resultBuf, binary.LittleEndian, &hardwareFlowcontrol)
	binary.Read(resultBuf, binary.LittleEndian, &softwareFlowcontrol)

	}

	return baudrate, parity, stopbits, wordlength, hardwareFlowcontrol, softwareFlowcontrol, nil
}

// Sets a break condition (the TX output is forced to a logic 0 state).
// The parameter sets the hold-time of the break condition (in ms).
// 
// .. versionadded:: 2.0.2$nbsp;(Plugin)
func (device *RS232Bricklet) SetBreakCondition(breakTime uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, breakTime);

	resultBytes, err := device.device.Set(uint8(FunctionSetBreakCondition), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
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
func (device *RS232Bricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
