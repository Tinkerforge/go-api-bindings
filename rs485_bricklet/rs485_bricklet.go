/* ***********************************************************
 * This file was automatically generated on 2024-02-27.      *
 *                                                           *
 * Go Bindings Version 2.0.15                                *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/

// Communicates with RS485/Modbus devices with full- or half-duplex.
//
//
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/RS485_Bricklet_Go.html.
package rs485_bricklet

import (
	"bytes"
	"encoding/binary"
	"fmt"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionWriteLowLevel                                            Function = 1
	FunctionReadLowLevel                                             Function = 2
	FunctionEnableReadCallback                                       Function = 3
	FunctionDisableReadCallback                                      Function = 4
	FunctionIsReadCallbackEnabled                                    Function = 5
	FunctionSetRS485Configuration                                    Function = 6
	FunctionGetRS485Configuration                                    Function = 7
	FunctionSetModbusConfiguration                                   Function = 8
	FunctionGetModbusConfiguration                                   Function = 9
	FunctionSetMode                                                  Function = 10
	FunctionGetMode                                                  Function = 11
	FunctionSetCommunicationLEDConfig                                Function = 12
	FunctionGetCommunicationLEDConfig                                Function = 13
	FunctionSetErrorLEDConfig                                        Function = 14
	FunctionGetErrorLEDConfig                                        Function = 15
	FunctionSetBufferConfig                                          Function = 16
	FunctionGetBufferConfig                                          Function = 17
	FunctionGetBufferStatus                                          Function = 18
	FunctionEnableErrorCountCallback                                 Function = 19
	FunctionDisableErrorCountCallback                                Function = 20
	FunctionIsErrorCountCallbackEnabled                              Function = 21
	FunctionGetErrorCount                                            Function = 22
	FunctionGetModbusCommonErrorCount                                Function = 23
	FunctionModbusSlaveReportException                               Function = 24
	FunctionModbusSlaveAnswerReadCoilsRequestLowLevel                Function = 25
	FunctionModbusMasterReadCoils                                    Function = 26
	FunctionModbusSlaveAnswerReadHoldingRegistersRequestLowLevel     Function = 27
	FunctionModbusMasterReadHoldingRegisters                         Function = 28
	FunctionModbusSlaveAnswerWriteSingleCoilRequest                  Function = 29
	FunctionModbusMasterWriteSingleCoil                              Function = 30
	FunctionModbusSlaveAnswerWriteSingleRegisterRequest              Function = 31
	FunctionModbusMasterWriteSingleRegister                          Function = 32
	FunctionModbusSlaveAnswerWriteMultipleCoilsRequest               Function = 33
	FunctionModbusMasterWriteMultipleCoilsLowLevel                   Function = 34
	FunctionModbusSlaveAnswerWriteMultipleRegistersRequest           Function = 35
	FunctionModbusMasterWriteMultipleRegistersLowLevel               Function = 36
	FunctionModbusSlaveAnswerReadDiscreteInputsRequestLowLevel       Function = 37
	FunctionModbusMasterReadDiscreteInputs                           Function = 38
	FunctionModbusSlaveAnswerReadInputRegistersRequestLowLevel       Function = 39
	FunctionModbusMasterReadInputRegisters                           Function = 40
	FunctionSetFrameReadableCallbackConfiguration                    Function = 59
	FunctionGetFrameReadableCallbackConfiguration                    Function = 60
	FunctionGetSPITFPErrorCount                                      Function = 234
	FunctionSetBootloaderMode                                        Function = 235
	FunctionGetBootloaderMode                                        Function = 236
	FunctionSetWriteFirmwarePointer                                  Function = 237
	FunctionWriteFirmware                                            Function = 238
	FunctionSetStatusLEDConfig                                       Function = 239
	FunctionGetStatusLEDConfig                                       Function = 240
	FunctionGetChipTemperature                                       Function = 242
	FunctionReset                                                    Function = 243
	FunctionWriteUID                                                 Function = 248
	FunctionReadUID                                                  Function = 249
	FunctionGetIdentity                                              Function = 255
	FunctionCallbackReadLowLevel                                     Function = 41
	FunctionCallbackErrorCount                                       Function = 42
	FunctionCallbackModbusSlaveReadCoilsRequest                      Function = 43
	FunctionCallbackModbusMasterReadCoilsResponseLowLevel            Function = 44
	FunctionCallbackModbusSlaveReadHoldingRegistersRequest           Function = 45
	FunctionCallbackModbusMasterReadHoldingRegistersResponseLowLevel Function = 46
	FunctionCallbackModbusSlaveWriteSingleCoilRequest                Function = 47
	FunctionCallbackModbusMasterWriteSingleCoilResponse              Function = 48
	FunctionCallbackModbusSlaveWriteSingleRegisterRequest            Function = 49
	FunctionCallbackModbusMasterWriteSingleRegisterResponse          Function = 50
	FunctionCallbackModbusSlaveWriteMultipleCoilsRequestLowLevel     Function = 51
	FunctionCallbackModbusMasterWriteMultipleCoilsResponse           Function = 52
	FunctionCallbackModbusSlaveWriteMultipleRegistersRequestLowLevel Function = 53
	FunctionCallbackModbusMasterWriteMultipleRegistersResponse       Function = 54
	FunctionCallbackModbusSlaveReadDiscreteInputsRequest             Function = 55
	FunctionCallbackModbusMasterReadDiscreteInputsResponseLowLevel   Function = 56
	FunctionCallbackModbusSlaveReadInputRegistersRequest             Function = 57
	FunctionCallbackModbusMasterReadInputRegistersResponseLowLevel   Function = 58
	FunctionCallbackFrameReadable                                    Function = 61
)

type Parity = uint8

const (
	ParityNone Parity = 0
	ParityOdd  Parity = 1
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

type Duplex = uint8

const (
	DuplexHalf Duplex = 0
	DuplexFull Duplex = 1
)

type Mode = uint8

const (
	ModeRS485           Mode = 0
	ModeModbusMasterRTU Mode = 1
	ModeModbusSlaveRTU  Mode = 2
)

type CommunicationLEDConfig = uint8

const (
	CommunicationLEDConfigOff               CommunicationLEDConfig = 0
	CommunicationLEDConfigOn                CommunicationLEDConfig = 1
	CommunicationLEDConfigShowHeartbeat     CommunicationLEDConfig = 2
	CommunicationLEDConfigShowCommunication CommunicationLEDConfig = 3
)

type ErrorLEDConfig = uint8

const (
	ErrorLEDConfigOff           ErrorLEDConfig = 0
	ErrorLEDConfigOn            ErrorLEDConfig = 1
	ErrorLEDConfigShowHeartbeat ErrorLEDConfig = 2
	ErrorLEDConfigShowError     ErrorLEDConfig = 3
)

type ExceptionCode = int8

const (
	ExceptionCodeTimeout                            ExceptionCode = -1
	ExceptionCodeSuccess                            ExceptionCode = 0
	ExceptionCodeIllegalFunction                    ExceptionCode = 1
	ExceptionCodeIllegalDataAddress                 ExceptionCode = 2
	ExceptionCodeIllegalDataValue                   ExceptionCode = 3
	ExceptionCodeSlaveDeviceFailure                 ExceptionCode = 4
	ExceptionCodeAcknowledge                        ExceptionCode = 5
	ExceptionCodeSlaveDeviceBusy                    ExceptionCode = 6
	ExceptionCodeMemoryParityError                  ExceptionCode = 8
	ExceptionCodeGatewayPathUnavailable             ExceptionCode = 10
	ExceptionCodeGatewayTargetDeviceFailedToRespond ExceptionCode = 11
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

type RS485Bricklet struct {
	device Device
}

const DeviceIdentifier = 277
const DeviceDisplayName = "RS485 Bricklet"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (RS485Bricklet, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{2, 0, 1}, uid, &internalIPCon, 15, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return RS485Bricklet{}, err
	}
	dev.ResponseExpected[FunctionWriteLowLevel] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionReadLowLevel] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionEnableReadCallback] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionDisableReadCallback] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionIsReadCallbackEnabled] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetRS485Configuration] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetRS485Configuration] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetModbusConfiguration] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetModbusConfiguration] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetMode] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetMode] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetCommunicationLEDConfig] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetCommunicationLEDConfig] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetErrorLEDConfig] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetErrorLEDConfig] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetBufferConfig] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetBufferConfig] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetBufferStatus] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionEnableErrorCountCallback] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionDisableErrorCountCallback] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionIsErrorCountCallbackEnabled] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetErrorCount] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetModbusCommonErrorCount] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionModbusSlaveReportException] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionModbusSlaveAnswerReadCoilsRequestLowLevel] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionModbusMasterReadCoils] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionModbusSlaveAnswerReadHoldingRegistersRequestLowLevel] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionModbusMasterReadHoldingRegisters] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionModbusSlaveAnswerWriteSingleCoilRequest] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionModbusMasterWriteSingleCoil] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionModbusSlaveAnswerWriteSingleRegisterRequest] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionModbusMasterWriteSingleRegister] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionModbusSlaveAnswerWriteMultipleCoilsRequest] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionModbusMasterWriteMultipleCoilsLowLevel] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionModbusSlaveAnswerWriteMultipleRegistersRequest] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionModbusMasterWriteMultipleRegistersLowLevel] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionModbusSlaveAnswerReadDiscreteInputsRequestLowLevel] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionModbusMasterReadDiscreteInputs] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionModbusSlaveAnswerReadInputRegistersRequestLowLevel] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionModbusMasterReadInputRegisters] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetFrameReadableCallbackConfiguration] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionGetFrameReadableCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue
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
	return RS485Bricklet{dev}, nil
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
func (device *RS485Bricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *RS485Bricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *RS485Bricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *RS485Bricklet) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// See RegisterReadCallback
func (device *RS485Bricklet) RegisterReadLowLevelCallback(fn func(uint16, uint16, [60]rune)) uint64 {
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
		copy(messageChunkData[:], ByteSliceToRuneSlice(buf.Next(8*60/8)))
		fn(messageLength, messageChunkOffset, messageChunkData)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackReadLowLevel), wrapper)
}

// Remove a registered Read Low Level callback.
func (device *RS485Bricklet) DeregisterReadLowLevelCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackReadLowLevel), registrationId)
}

// This callback is called if new data is available.
//
// To enable this callback, use EnableReadCallback.
//
// Note
//  If reconstructing the value fails, the callback is triggered with nil for message.
func (device *RS485Bricklet) RegisterReadCallback(fn func([]rune)) uint64 {
	buf := make([]rune, 0)
	wrapper := func(messageLength uint16, messageChunkOffset uint16, messageChunkData [60]rune) {
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
func (device *RS485Bricklet) DeregisterReadCallback(registrationId uint64) {
	device.DeregisterReadLowLevelCallback(registrationId)
}

// This callback is called if a new error occurs. It returns
// the current overrun and parity error count.
func (device *RS485Bricklet) RegisterErrorCountCallback(fn func(uint32, uint32)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 16 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var overrunErrorCount uint32
		var parityErrorCount uint32
		binary.Read(buf, binary.LittleEndian, &overrunErrorCount)
		binary.Read(buf, binary.LittleEndian, &parityErrorCount)
		fn(overrunErrorCount, parityErrorCount)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackErrorCount), wrapper)
}

// Remove a registered Error Count callback.
func (device *RS485Bricklet) DeregisterErrorCountCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackErrorCount), registrationId)
}

// This callback is called only in Modbus slave mode when the slave receives a
// valid request from a Modbus master to read coils. The parameters are
// request ID of the request, the number of the first coil to be read and the number of coils to
// be read as received by the request. The number of the first coil is called starting address for backwards compatibility reasons.
// It is not an address, but instead a coil number in the range of 1 to 65536.
//
// To send a response of this request use ModbusSlaveAnswerReadCoilsRequest.
func (device *RS485Bricklet) RegisterModbusSlaveReadCoilsRequestCallback(fn func(uint8, uint32, uint16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 15 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var requestID uint8
		var startingAddress uint32
		var count uint16
		binary.Read(buf, binary.LittleEndian, &requestID)
		binary.Read(buf, binary.LittleEndian, &startingAddress)
		binary.Read(buf, binary.LittleEndian, &count)
		fn(requestID, startingAddress, count)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackModbusSlaveReadCoilsRequest), wrapper)
}

// Remove a registered Modbus Slave Read Coils Request callback.
func (device *RS485Bricklet) DeregisterModbusSlaveReadCoilsRequestCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackModbusSlaveReadCoilsRequest), registrationId)
}

// See RegisterModbusMasterReadCoilsResponseCallback
func (device *RS485Bricklet) RegisterModbusMasterReadCoilsResponseLowLevelCallback(fn func(uint8, ExceptionCode, uint16, uint16, [464]bool)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 72 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var requestID uint8
		var exceptionCode ExceptionCode
		var coilsLength uint16
		var coilsChunkOffset uint16
		var coilsChunkData [464]bool
		binary.Read(buf, binary.LittleEndian, &requestID)
		binary.Read(buf, binary.LittleEndian, &exceptionCode)
		binary.Read(buf, binary.LittleEndian, &coilsLength)
		binary.Read(buf, binary.LittleEndian, &coilsChunkOffset)
		copy(coilsChunkData[:], ByteSliceToBoolSlice(buf.Next(1*464/8)))
		fn(requestID, exceptionCode, coilsLength, coilsChunkOffset, coilsChunkData)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackModbusMasterReadCoilsResponseLowLevel), wrapper)
}

// Remove a registered Modbus Master Read Coils Response Low Level callback.
func (device *RS485Bricklet) DeregisterModbusMasterReadCoilsResponseLowLevelCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackModbusMasterReadCoilsResponseLowLevel), registrationId)
}

// This callback is called only in Modbus master mode when the master receives a
// valid response of a request to read coils.
//
// The parameters are request ID
// of the request, exception code of the response and the data as received by the
// response.
//
// Any non-zero exception code indicates a problem. If the exception code
// is greater than 0 then the number represents a Modbus exception code. If it is
// less than 0 then it represents other errors. For example, -1 indicates that
// the request timed out or that the master did not receive any valid response of the
// request within the master request timeout period as set by
// SetModbusConfiguration.
//
// Note
//  If reconstructing the value fails, the callback is triggered with nil for coils.
func (device *RS485Bricklet) RegisterModbusMasterReadCoilsResponseCallback(fn func(uint8, ExceptionCode, []bool)) uint64 {
	buf := make([]bool, 0)
	wrapper := func(requestID uint8, exceptionCode ExceptionCode, coilsLength uint16, coilsChunkOffset uint16, coilsChunkData [464]bool) {
		if int(coilsChunkOffset) != len(buf) {
			if len(buf) != 0 {
				buf = nil
				fn(requestID, exceptionCode, buf)
				buf = make([]bool, 0)
			}
			if coilsChunkOffset != 0 {
				return
			}
		}
		toRead := MinU(uint64(coilsLength-coilsChunkOffset), uint64(len(coilsChunkData[:])))
		buf = append(buf, coilsChunkData[:toRead]...)
		if len(buf) >= int(coilsLength) {
			fn(requestID, exceptionCode, buf)
			buf = make([]bool, 0)
		}
	}
	return device.RegisterModbusMasterReadCoilsResponseLowLevelCallback(wrapper)
}

// Remove a registered Modbus Master Read Coils Response Low Level callback.
func (device *RS485Bricklet) DeregisterModbusMasterReadCoilsResponseCallback(registrationId uint64) {
	device.DeregisterModbusMasterReadCoilsResponseLowLevelCallback(registrationId)
}

// This callback is called only in Modbus slave mode when the slave receives a
// valid request from a Modbus master to read holding registers. The parameters
// are request ID of the request, the number of the first holding register to be read and the number of holding
// registers to be read as received by the request. The number of the first holding register is called starting address for backwards compatibility reasons.
// It is not an address, but instead a holding register number in the range of 1 to 65536. The prefix digit 4 (for holding register) is omitted.
//
// To send a response of this request use ModbusSlaveAnswerReadHoldingRegistersRequest.
func (device *RS485Bricklet) RegisterModbusSlaveReadHoldingRegistersRequestCallback(fn func(uint8, uint32, uint16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 15 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var requestID uint8
		var startingAddress uint32
		var count uint16
		binary.Read(buf, binary.LittleEndian, &requestID)
		binary.Read(buf, binary.LittleEndian, &startingAddress)
		binary.Read(buf, binary.LittleEndian, &count)
		fn(requestID, startingAddress, count)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackModbusSlaveReadHoldingRegistersRequest), wrapper)
}

// Remove a registered Modbus Slave Read Holding Registers Request callback.
func (device *RS485Bricklet) DeregisterModbusSlaveReadHoldingRegistersRequestCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackModbusSlaveReadHoldingRegistersRequest), registrationId)
}

// See RegisterModbusMasterReadHoldingRegistersResponseCallback
func (device *RS485Bricklet) RegisterModbusMasterReadHoldingRegistersResponseLowLevelCallback(fn func(uint8, ExceptionCode, uint16, uint16, [29]uint16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 72 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var requestID uint8
		var exceptionCode ExceptionCode
		var holdingRegistersLength uint16
		var holdingRegistersChunkOffset uint16
		var holdingRegistersChunkData [29]uint16
		binary.Read(buf, binary.LittleEndian, &requestID)
		binary.Read(buf, binary.LittleEndian, &exceptionCode)
		binary.Read(buf, binary.LittleEndian, &holdingRegistersLength)
		binary.Read(buf, binary.LittleEndian, &holdingRegistersChunkOffset)
		copy(holdingRegistersChunkData[:], ByteSliceToUint16Slice(buf.Next(16*29/8)))
		fn(requestID, exceptionCode, holdingRegistersLength, holdingRegistersChunkOffset, holdingRegistersChunkData)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackModbusMasterReadHoldingRegistersResponseLowLevel), wrapper)
}

// Remove a registered Modbus Master Read Holding Registers Response Low Level callback.
func (device *RS485Bricklet) DeregisterModbusMasterReadHoldingRegistersResponseLowLevelCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackModbusMasterReadHoldingRegistersResponseLowLevel), registrationId)
}

// This callback is called only in Modbus master mode when the master receives a
// valid response of a request to read holding registers.
//
// The parameters are
// request ID of the request, exception code of the response and the data as received
// by the response.
//
// Any non-zero exception code indicates a problem. If the exception
// code is greater than 0 then the number represents a Modbus exception code. If
// it is less than 0 then it represents other errors. For example, -1 indicates that
// the request timed out or that the master did not receive any valid response of the
// request within the master request timeout period as set by
// SetModbusConfiguration.
//
// Note
//  If reconstructing the value fails, the callback is triggered with nil for holdingRegisters.
func (device *RS485Bricklet) RegisterModbusMasterReadHoldingRegistersResponseCallback(fn func(uint8, ExceptionCode, []uint16)) uint64 {
	buf := make([]uint16, 0)
	wrapper := func(requestID uint8, exceptionCode ExceptionCode, holdingRegistersLength uint16, holdingRegistersChunkOffset uint16, holdingRegistersChunkData [29]uint16) {
		if int(holdingRegistersChunkOffset) != len(buf) {
			if len(buf) != 0 {
				buf = nil
				fn(requestID, exceptionCode, buf)
				buf = make([]uint16, 0)
			}
			if holdingRegistersChunkOffset != 0 {
				return
			}
		}
		toRead := MinU(uint64(holdingRegistersLength-holdingRegistersChunkOffset), uint64(len(holdingRegistersChunkData[:])))
		buf = append(buf, holdingRegistersChunkData[:toRead]...)
		if len(buf) >= int(holdingRegistersLength) {
			fn(requestID, exceptionCode, buf)
			buf = make([]uint16, 0)
		}
	}
	return device.RegisterModbusMasterReadHoldingRegistersResponseLowLevelCallback(wrapper)
}

// Remove a registered Modbus Master Read Holding Registers Response Low Level callback.
func (device *RS485Bricklet) DeregisterModbusMasterReadHoldingRegistersResponseCallback(registrationId uint64) {
	device.DeregisterModbusMasterReadHoldingRegistersResponseLowLevelCallback(registrationId)
}

// This callback is called only in Modbus slave mode when the slave receives a
// valid request from a Modbus master to write a single coil. The parameters
// are request ID of the request, the number of the coil and the value of coil to be
// written as received by the request. The number of the coil is called coil address for backwards compatibility reasons.
// It is not an address, but instead a coil number in the range of 1 to 65536.
//
// To send a response of this request use ModbusSlaveAnswerWriteSingleCoilRequest.
func (device *RS485Bricklet) RegisterModbusSlaveWriteSingleCoilRequestCallback(fn func(uint8, uint32, bool)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 14 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var requestID uint8
		var coilAddress uint32
		var coilValue bool
		binary.Read(buf, binary.LittleEndian, &requestID)
		binary.Read(buf, binary.LittleEndian, &coilAddress)
		binary.Read(buf, binary.LittleEndian, &coilValue)
		fn(requestID, coilAddress, coilValue)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackModbusSlaveWriteSingleCoilRequest), wrapper)
}

// Remove a registered Modbus Slave Write Single Coil Request callback.
func (device *RS485Bricklet) DeregisterModbusSlaveWriteSingleCoilRequestCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackModbusSlaveWriteSingleCoilRequest), registrationId)
}

// This callback is called only in Modbus master mode when the master receives a
// valid response of a request to write a single coil.
//
// The parameters are
// request ID of the request and exception code of the response.
//
// Any non-zero exception code indicates a problem.
// If the exception code is greater than 0 then the number represents a Modbus
// exception code. If it is less than 0 then it represents other errors. For
// example, -1 indicates that the request timed out or that the master did not receive
// any valid response of the request within the master request timeout period as set
// by SetModbusConfiguration.
func (device *RS485Bricklet) RegisterModbusMasterWriteSingleCoilResponseCallback(fn func(uint8, ExceptionCode)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 10 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var requestID uint8
		var exceptionCode ExceptionCode
		binary.Read(buf, binary.LittleEndian, &requestID)
		binary.Read(buf, binary.LittleEndian, &exceptionCode)
		fn(requestID, exceptionCode)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackModbusMasterWriteSingleCoilResponse), wrapper)
}

// Remove a registered Modbus Master Write Single Coil Response callback.
func (device *RS485Bricklet) DeregisterModbusMasterWriteSingleCoilResponseCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackModbusMasterWriteSingleCoilResponse), registrationId)
}

// This callback is called only in Modbus slave mode when the slave receives a
// valid request from a Modbus master to write a single holding register. The parameters
// are request ID of the request, the number of the holding register and the register value to
// be written as received by the request. The number of the holding register is called starting address for backwards compatibility reasons.
// It is not an address, but instead a holding register number in the range of 1 to 65536. The prefix digit 4 (for holding register) is omitted.
//
// To send a response of this request use ModbusSlaveAnswerWriteSingleRegisterRequest.
func (device *RS485Bricklet) RegisterModbusSlaveWriteSingleRegisterRequestCallback(fn func(uint8, uint32, uint16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 15 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var requestID uint8
		var registerAddress uint32
		var registerValue uint16
		binary.Read(buf, binary.LittleEndian, &requestID)
		binary.Read(buf, binary.LittleEndian, &registerAddress)
		binary.Read(buf, binary.LittleEndian, &registerValue)
		fn(requestID, registerAddress, registerValue)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackModbusSlaveWriteSingleRegisterRequest), wrapper)
}

// Remove a registered Modbus Slave Write Single Register Request callback.
func (device *RS485Bricklet) DeregisterModbusSlaveWriteSingleRegisterRequestCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackModbusSlaveWriteSingleRegisterRequest), registrationId)
}

// This callback is called only in Modbus master mode when the master receives a
// valid response of a request to write a single register.
//
// The parameters are
// request ID of the request and exception code of the response.
//
// Any non-zero exception code
// indicates a problem. If the exception code is greater than 0 then the number
// represents a Modbus exception code. If it is less than 0 then it represents
// other errors. For example, -1 indicates that the request timed out or that the
// master did not receive any valid response of the request within the master request
// timeout period as set by SetModbusConfiguration.
func (device *RS485Bricklet) RegisterModbusMasterWriteSingleRegisterResponseCallback(fn func(uint8, ExceptionCode)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 10 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var requestID uint8
		var exceptionCode ExceptionCode
		binary.Read(buf, binary.LittleEndian, &requestID)
		binary.Read(buf, binary.LittleEndian, &exceptionCode)
		fn(requestID, exceptionCode)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackModbusMasterWriteSingleRegisterResponse), wrapper)
}

// Remove a registered Modbus Master Write Single Register Response callback.
func (device *RS485Bricklet) DeregisterModbusMasterWriteSingleRegisterResponseCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackModbusMasterWriteSingleRegisterResponse), registrationId)
}

// See RegisterModbusSlaveWriteMultipleCoilsRequestCallback
func (device *RS485Bricklet) RegisterModbusSlaveWriteMultipleCoilsRequestLowLevelCallback(fn func(uint8, uint32, uint16, uint16, [440]bool)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 72 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var requestID uint8
		var startingAddress uint32
		var coilsLength uint16
		var coilsChunkOffset uint16
		var coilsChunkData [440]bool
		binary.Read(buf, binary.LittleEndian, &requestID)
		binary.Read(buf, binary.LittleEndian, &startingAddress)
		binary.Read(buf, binary.LittleEndian, &coilsLength)
		binary.Read(buf, binary.LittleEndian, &coilsChunkOffset)
		copy(coilsChunkData[:], ByteSliceToBoolSlice(buf.Next(1*440/8)))
		fn(requestID, startingAddress, coilsLength, coilsChunkOffset, coilsChunkData)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackModbusSlaveWriteMultipleCoilsRequestLowLevel), wrapper)
}

// Remove a registered Modbus Slave Write Multiple Coils Request Low Level callback.
func (device *RS485Bricklet) DeregisterModbusSlaveWriteMultipleCoilsRequestLowLevelCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackModbusSlaveWriteMultipleCoilsRequestLowLevel), registrationId)
}

// This callback is called only in Modbus slave mode when the slave receives a
// valid request from a Modbus master to write multiple coils. The parameters
// are request ID of the request, the number of the first coil and the data to be written as
// received by the request. The number of the first coil is called starting address for backwards compatibility reasons.
// It is not an address, but instead a coil number in the range of 1 to 65536.
//
// To send a response of this request use ModbusSlaveAnswerWriteMultipleCoilsRequest.
//
// Note
//  If reconstructing the value fails, the callback is triggered with nil for coils.
func (device *RS485Bricklet) RegisterModbusSlaveWriteMultipleCoilsRequestCallback(fn func(uint8, uint32, []bool)) uint64 {
	buf := make([]bool, 0)
	wrapper := func(requestID uint8, startingAddress uint32, coilsLength uint16, coilsChunkOffset uint16, coilsChunkData [440]bool) {
		if int(coilsChunkOffset) != len(buf) {
			if len(buf) != 0 {
				buf = nil
				fn(requestID, startingAddress, buf)
				buf = make([]bool, 0)
			}
			if coilsChunkOffset != 0 {
				return
			}
		}
		toRead := MinU(uint64(coilsLength-coilsChunkOffset), uint64(len(coilsChunkData[:])))
		buf = append(buf, coilsChunkData[:toRead]...)
		if len(buf) >= int(coilsLength) {
			fn(requestID, startingAddress, buf)
			buf = make([]bool, 0)
		}
	}
	return device.RegisterModbusSlaveWriteMultipleCoilsRequestLowLevelCallback(wrapper)
}

// Remove a registered Modbus Slave Write Multiple Coils Request Low Level callback.
func (device *RS485Bricklet) DeregisterModbusSlaveWriteMultipleCoilsRequestCallback(registrationId uint64) {
	device.DeregisterModbusSlaveWriteMultipleCoilsRequestLowLevelCallback(registrationId)
}

// This callback is called only in Modbus master mode when the master receives a
// valid response of a request to read coils.
//
// The parameters are
// request ID of the request and exception code of the response.
//
// Any non-zero exception code
// indicates a problem. If the exception code is greater than 0 then the number
// represents a Modbus exception code. If it is less than 0 then it represents
// other errors. For example, -1 indicates that the request timedout or that the
// master did not receive any valid response of the request within the master request
// timeout period as set by SetModbusConfiguration.
func (device *RS485Bricklet) RegisterModbusMasterWriteMultipleCoilsResponseCallback(fn func(uint8, ExceptionCode)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 10 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var requestID uint8
		var exceptionCode ExceptionCode
		binary.Read(buf, binary.LittleEndian, &requestID)
		binary.Read(buf, binary.LittleEndian, &exceptionCode)
		fn(requestID, exceptionCode)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackModbusMasterWriteMultipleCoilsResponse), wrapper)
}

// Remove a registered Modbus Master Write Multiple Coils Response callback.
func (device *RS485Bricklet) DeregisterModbusMasterWriteMultipleCoilsResponseCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackModbusMasterWriteMultipleCoilsResponse), registrationId)
}

// See RegisterModbusSlaveWriteMultipleRegistersRequestCallback
func (device *RS485Bricklet) RegisterModbusSlaveWriteMultipleRegistersRequestLowLevelCallback(fn func(uint8, uint32, uint16, uint16, [27]uint16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 71 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var requestID uint8
		var startingAddress uint32
		var registersLength uint16
		var registersChunkOffset uint16
		var registersChunkData [27]uint16
		binary.Read(buf, binary.LittleEndian, &requestID)
		binary.Read(buf, binary.LittleEndian, &startingAddress)
		binary.Read(buf, binary.LittleEndian, &registersLength)
		binary.Read(buf, binary.LittleEndian, &registersChunkOffset)
		copy(registersChunkData[:], ByteSliceToUint16Slice(buf.Next(16*27/8)))
		fn(requestID, startingAddress, registersLength, registersChunkOffset, registersChunkData)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackModbusSlaveWriteMultipleRegistersRequestLowLevel), wrapper)
}

// Remove a registered Modbus Slave Write Multiple Registers Request Low Level callback.
func (device *RS485Bricklet) DeregisterModbusSlaveWriteMultipleRegistersRequestLowLevelCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackModbusSlaveWriteMultipleRegistersRequestLowLevel), registrationId)
}

// This callback is called only in Modbus slave mode when the slave receives a
// valid request from a Modbus master to write multiple holding registers. The parameters
// are request ID of the request, the number of the first holding register and the data to be written as
// received by the request. The number of the first holding register is called starting address for backwards compatibility reasons.
// It is not an address, but instead a holding register number in the range of 1 to 65536. The prefix digit 4 (for holding register) is omitted.
//
// To send a response of this request use ModbusSlaveAnswerWriteMultipleRegistersRequest.
//
// Note
//  If reconstructing the value fails, the callback is triggered with nil for registers.
func (device *RS485Bricklet) RegisterModbusSlaveWriteMultipleRegistersRequestCallback(fn func(uint8, uint32, []uint16)) uint64 {
	buf := make([]uint16, 0)
	wrapper := func(requestID uint8, startingAddress uint32, registersLength uint16, registersChunkOffset uint16, registersChunkData [27]uint16) {
		if int(registersChunkOffset) != len(buf) {
			if len(buf) != 0 {
				buf = nil
				fn(requestID, startingAddress, buf)
				buf = make([]uint16, 0)
			}
			if registersChunkOffset != 0 {
				return
			}
		}
		toRead := MinU(uint64(registersLength-registersChunkOffset), uint64(len(registersChunkData[:])))
		buf = append(buf, registersChunkData[:toRead]...)
		if len(buf) >= int(registersLength) {
			fn(requestID, startingAddress, buf)
			buf = make([]uint16, 0)
		}
	}
	return device.RegisterModbusSlaveWriteMultipleRegistersRequestLowLevelCallback(wrapper)
}

// Remove a registered Modbus Slave Write Multiple Registers Request Low Level callback.
func (device *RS485Bricklet) DeregisterModbusSlaveWriteMultipleRegistersRequestCallback(registrationId uint64) {
	device.DeregisterModbusSlaveWriteMultipleRegistersRequestLowLevelCallback(registrationId)
}

// This callback is called only in Modbus master mode when the master receives a
// valid response of a request to write multiple registers.
//
// The parameters
// are request ID of the request and exception code of the response.
//
// Any non-zero
// exception code indicates a problem. If the exception code is greater than 0 then
// the number represents a Modbus exception code. If it is less than 0 then it
// represents other errors. For example, -1 indicates that the request timedout or
// that the master did not receive any valid response of the request within the master
// request timeout period as set by SetModbusConfiguration.
func (device *RS485Bricklet) RegisterModbusMasterWriteMultipleRegistersResponseCallback(fn func(uint8, ExceptionCode)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 10 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var requestID uint8
		var exceptionCode ExceptionCode
		binary.Read(buf, binary.LittleEndian, &requestID)
		binary.Read(buf, binary.LittleEndian, &exceptionCode)
		fn(requestID, exceptionCode)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackModbusMasterWriteMultipleRegistersResponse), wrapper)
}

// Remove a registered Modbus Master Write Multiple Registers Response callback.
func (device *RS485Bricklet) DeregisterModbusMasterWriteMultipleRegistersResponseCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackModbusMasterWriteMultipleRegistersResponse), registrationId)
}

// This callback is called only in Modbus slave mode when the slave receives a
// valid request from a Modbus master to read discrete inputs. The parameters
// are request ID of the request, the number of the first discrete input and the number of discrete
// inputs to be read as received by the request. The number of the first discrete input is called starting address for backwards compatibility reasons.
// It is not an address, but instead a discrete input number in the range of 1 to 65536. The prefix digit 1 (for discrete input) is omitted.
//
// To send a response of this request use ModbusSlaveAnswerReadDiscreteInputsRequest.
func (device *RS485Bricklet) RegisterModbusSlaveReadDiscreteInputsRequestCallback(fn func(uint8, uint32, uint16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 15 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var requestID uint8
		var startingAddress uint32
		var count uint16
		binary.Read(buf, binary.LittleEndian, &requestID)
		binary.Read(buf, binary.LittleEndian, &startingAddress)
		binary.Read(buf, binary.LittleEndian, &count)
		fn(requestID, startingAddress, count)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackModbusSlaveReadDiscreteInputsRequest), wrapper)
}

// Remove a registered Modbus Slave Read Discrete Inputs Request callback.
func (device *RS485Bricklet) DeregisterModbusSlaveReadDiscreteInputsRequestCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackModbusSlaveReadDiscreteInputsRequest), registrationId)
}

// See RegisterModbusMasterReadDiscreteInputsResponseCallback
func (device *RS485Bricklet) RegisterModbusMasterReadDiscreteInputsResponseLowLevelCallback(fn func(uint8, ExceptionCode, uint16, uint16, [464]bool)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 72 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var requestID uint8
		var exceptionCode ExceptionCode
		var discreteInputsLength uint16
		var discreteInputsChunkOffset uint16
		var discreteInputsChunkData [464]bool
		binary.Read(buf, binary.LittleEndian, &requestID)
		binary.Read(buf, binary.LittleEndian, &exceptionCode)
		binary.Read(buf, binary.LittleEndian, &discreteInputsLength)
		binary.Read(buf, binary.LittleEndian, &discreteInputsChunkOffset)
		copy(discreteInputsChunkData[:], ByteSliceToBoolSlice(buf.Next(1*464/8)))
		fn(requestID, exceptionCode, discreteInputsLength, discreteInputsChunkOffset, discreteInputsChunkData)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackModbusMasterReadDiscreteInputsResponseLowLevel), wrapper)
}

// Remove a registered Modbus Master Read Discrete Inputs Response Low Level callback.
func (device *RS485Bricklet) DeregisterModbusMasterReadDiscreteInputsResponseLowLevelCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackModbusMasterReadDiscreteInputsResponseLowLevel), registrationId)
}

// This callback is called only in Modbus master mode when the master receives a
// valid response of a request to read discrete inputs.
//
// The parameters are
// request ID of the request, exception code of the response and the data as received
// by the response.
//
// Any non-zero exception code indicates a problem. If the exception
// code is greater than 0 then the number represents a Modbus exception code. If
// it is less than 0 then it represents other errors. For example, -1 indicates that
// the request timedout or that the master did not receive any valid response of the
// request within the master request timeout period as set by
// SetModbusConfiguration.
//
// Note
//  If reconstructing the value fails, the callback is triggered with nil for discreteInputs.
func (device *RS485Bricklet) RegisterModbusMasterReadDiscreteInputsResponseCallback(fn func(uint8, ExceptionCode, []bool)) uint64 {
	buf := make([]bool, 0)
	wrapper := func(requestID uint8, exceptionCode ExceptionCode, discreteInputsLength uint16, discreteInputsChunkOffset uint16, discreteInputsChunkData [464]bool) {
		if int(discreteInputsChunkOffset) != len(buf) {
			if len(buf) != 0 {
				buf = nil
				fn(requestID, exceptionCode, buf)
				buf = make([]bool, 0)
			}
			if discreteInputsChunkOffset != 0 {
				return
			}
		}
		toRead := MinU(uint64(discreteInputsLength-discreteInputsChunkOffset), uint64(len(discreteInputsChunkData[:])))
		buf = append(buf, discreteInputsChunkData[:toRead]...)
		if len(buf) >= int(discreteInputsLength) {
			fn(requestID, exceptionCode, buf)
			buf = make([]bool, 0)
		}
	}
	return device.RegisterModbusMasterReadDiscreteInputsResponseLowLevelCallback(wrapper)
}

// Remove a registered Modbus Master Read Discrete Inputs Response Low Level callback.
func (device *RS485Bricklet) DeregisterModbusMasterReadDiscreteInputsResponseCallback(registrationId uint64) {
	device.DeregisterModbusMasterReadDiscreteInputsResponseLowLevelCallback(registrationId)
}

// This callback is called only in Modbus slave mode when the slave receives a
// valid request from a Modbus master to read input registers. The parameters
// are request ID of the request, the number of the first input register and the number of input
// registers to be read as received by the request. The number of the first input register is called starting address for backwards compatibility reasons.
// It is not an address, but instead a input register number in the range of 1 to 65536. The prefix digit 3 (for input register) is omitted.
//
// To send a response of this request use ModbusSlaveAnswerReadInputRegistersRequest.
func (device *RS485Bricklet) RegisterModbusSlaveReadInputRegistersRequestCallback(fn func(uint8, uint32, uint16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 15 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var requestID uint8
		var startingAddress uint32
		var count uint16
		binary.Read(buf, binary.LittleEndian, &requestID)
		binary.Read(buf, binary.LittleEndian, &startingAddress)
		binary.Read(buf, binary.LittleEndian, &count)
		fn(requestID, startingAddress, count)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackModbusSlaveReadInputRegistersRequest), wrapper)
}

// Remove a registered Modbus Slave Read Input Registers Request callback.
func (device *RS485Bricklet) DeregisterModbusSlaveReadInputRegistersRequestCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackModbusSlaveReadInputRegistersRequest), registrationId)
}

// See RegisterModbusMasterReadInputRegistersResponseCallback
func (device *RS485Bricklet) RegisterModbusMasterReadInputRegistersResponseLowLevelCallback(fn func(uint8, ExceptionCode, uint16, uint16, [29]uint16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 72 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var requestID uint8
		var exceptionCode ExceptionCode
		var inputRegistersLength uint16
		var inputRegistersChunkOffset uint16
		var inputRegistersChunkData [29]uint16
		binary.Read(buf, binary.LittleEndian, &requestID)
		binary.Read(buf, binary.LittleEndian, &exceptionCode)
		binary.Read(buf, binary.LittleEndian, &inputRegistersLength)
		binary.Read(buf, binary.LittleEndian, &inputRegistersChunkOffset)
		copy(inputRegistersChunkData[:], ByteSliceToUint16Slice(buf.Next(16*29/8)))
		fn(requestID, exceptionCode, inputRegistersLength, inputRegistersChunkOffset, inputRegistersChunkData)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackModbusMasterReadInputRegistersResponseLowLevel), wrapper)
}

// Remove a registered Modbus Master Read Input Registers Response Low Level callback.
func (device *RS485Bricklet) DeregisterModbusMasterReadInputRegistersResponseLowLevelCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackModbusMasterReadInputRegistersResponseLowLevel), registrationId)
}

// This callback is called only in Modbus master mode when the master receives a
// valid response of a request to read input registers.
//
// The parameters are
// request ID of the request, exception code of the response and the data as received
// by the response.
//
// Any non-zero exception code indicates a problem. If the exception
// code is greater than 0 then the number represents a Modbus exception code. If
// it is less than 0 then it represents other errors. For example, -1 indicates that
// the request timedout or that the master did not receive any valid response of the
// request within the master request timeout period as set by
// SetModbusConfiguration.
//
// Note
//  If reconstructing the value fails, the callback is triggered with nil for inputRegisters.
func (device *RS485Bricklet) RegisterModbusMasterReadInputRegistersResponseCallback(fn func(uint8, ExceptionCode, []uint16)) uint64 {
	buf := make([]uint16, 0)
	wrapper := func(requestID uint8, exceptionCode ExceptionCode, inputRegistersLength uint16, inputRegistersChunkOffset uint16, inputRegistersChunkData [29]uint16) {
		if int(inputRegistersChunkOffset) != len(buf) {
			if len(buf) != 0 {
				buf = nil
				fn(requestID, exceptionCode, buf)
				buf = make([]uint16, 0)
			}
			if inputRegistersChunkOffset != 0 {
				return
			}
		}
		toRead := MinU(uint64(inputRegistersLength-inputRegistersChunkOffset), uint64(len(inputRegistersChunkData[:])))
		buf = append(buf, inputRegistersChunkData[:toRead]...)
		if len(buf) >= int(inputRegistersLength) {
			fn(requestID, exceptionCode, buf)
			buf = make([]uint16, 0)
		}
	}
	return device.RegisterModbusMasterReadInputRegistersResponseLowLevelCallback(wrapper)
}

// Remove a registered Modbus Master Read Input Registers Response Low Level callback.
func (device *RS485Bricklet) DeregisterModbusMasterReadInputRegistersResponseCallback(registrationId uint64) {
	device.DeregisterModbusMasterReadInputRegistersResponseLowLevelCallback(registrationId)
}

// This callback is called if at least one frame of data is readable. The frame size is configured with SetFrameReadableCallbackConfiguration.
// The frame count parameter is the number of frames that can be read.
// This callback is triggered only once until Read is called. This means, that if you have configured a frame size of X bytes,
// you can read exactly X bytes using the Read function, every time the callback triggers without checking the frame count parameter.
//
// .. versionadded:: 2.0.5$nbsp;(Plugin)
func (device *RS485Bricklet) RegisterFrameReadableCallback(fn func(uint16)) uint64 {
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
func (device *RS485Bricklet) DeregisterFrameReadableCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackFrameReadable), registrationId)
}

// Writes characters to the RS485 interface. The characters can be binary data,
// ASCII or similar is not necessary.
//
// The return value is the number of characters that were written.
//
// See SetRS485Configuration for configuration possibilities
// regarding baudrate, parity and so on.
func (device *RS485Bricklet) WriteLowLevel(messageLength uint16, messageChunkOffset uint16, messageChunkData [60]rune) (messageChunkWritten uint8, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, messageLength)
	binary.Write(&buf, binary.LittleEndian, messageChunkOffset)
	buf.Write(RuneSliceToByteSlice(messageChunkData[:]))

	resultBytes, err := device.device.Get(uint8(FunctionWriteLowLevel), buf.Bytes())
	if err != nil {
		return messageChunkWritten, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return messageChunkWritten, DeviceError(header.ErrorCode)
		}

		if header.Length != 9 {
			return messageChunkWritten, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &messageChunkWritten)

	}

	return messageChunkWritten, nil
}

// Writes characters to the RS485 interface. The characters can be binary data,
// ASCII or similar is not necessary.
//
// The return value is the number of characters that were written.
//
// See SetRS485Configuration for configuration possibilities
// regarding baudrate, parity and so on.
func (device *RS485Bricklet) Write(message []rune) (messageChunkWritten uint64, err error) {
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
func (device *RS485Bricklet) ReadLowLevel(length uint16) (messageLength uint16, messageChunkOffset uint16, messageChunkData [60]rune, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, length)

	resultBytes, err := device.device.Get(uint8(FunctionReadLowLevel), buf.Bytes())
	if err != nil {
		return messageLength, messageChunkOffset, messageChunkData, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return messageLength, messageChunkOffset, messageChunkData, DeviceError(header.ErrorCode)
		}

		if header.Length != 72 {
			return messageLength, messageChunkOffset, messageChunkData, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 72)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &messageLength)
		binary.Read(resultBuf, binary.LittleEndian, &messageChunkOffset)
		copy(messageChunkData[:], ByteSliceToRuneSlice(resultBuf.Next(8*60/8)))

	}

	return messageLength, messageChunkOffset, messageChunkData, nil
}

// Returns up to *length* characters from receive buffer.
//
// Instead of polling with this function, you can also use
// callbacks. But note that this function will return available
// data only when the read callback is disabled.
// See EnableReadCallback and RegisterReadCallback callback.
func (device *RS485Bricklet) Read(length uint16) (message []rune, err error) {
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
func (device *RS485Bricklet) EnableReadCallback() (err error) {
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

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		bytes.NewBuffer(resultBytes[8:])

	}

	return nil
}

// Disables the RegisterReadCallback callback.
//
// By default the callback is disabled.
func (device *RS485Bricklet) DisableReadCallback() (err error) {
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

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		bytes.NewBuffer(resultBytes[8:])

	}

	return nil
}

// Returns *true* if the RegisterReadCallback callback is enabled,
// *false* otherwise.
func (device *RS485Bricklet) IsReadCallbackEnabled() (enabled bool, err error) {
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

		if header.Length != 9 {
			return enabled, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &enabled)

	}

	return enabled, nil
}

// Sets the configuration for the RS485 communication.
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
//	* DuplexHalf
//	* DuplexFull
func (device *RS485Bricklet) SetRS485Configuration(baudrate uint32, parity Parity, stopbits Stopbits, wordlength Wordlength, duplex Duplex) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, baudrate)
	binary.Write(&buf, binary.LittleEndian, parity)
	binary.Write(&buf, binary.LittleEndian, stopbits)
	binary.Write(&buf, binary.LittleEndian, wordlength)
	binary.Write(&buf, binary.LittleEndian, duplex)

	resultBytes, err := device.device.Set(uint8(FunctionSetRS485Configuration), buf.Bytes())
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

// Returns the configuration as set by SetRS485Configuration.
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
//	* DuplexHalf
//	* DuplexFull
func (device *RS485Bricklet) GetRS485Configuration() (baudrate uint32, parity Parity, stopbits Stopbits, wordlength Wordlength, duplex Duplex, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetRS485Configuration), buf.Bytes())
	if err != nil {
		return baudrate, parity, stopbits, wordlength, duplex, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return baudrate, parity, stopbits, wordlength, duplex, DeviceError(header.ErrorCode)
		}

		if header.Length != 16 {
			return baudrate, parity, stopbits, wordlength, duplex, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 16)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &baudrate)
		binary.Read(resultBuf, binary.LittleEndian, &parity)
		binary.Read(resultBuf, binary.LittleEndian, &stopbits)
		binary.Read(resultBuf, binary.LittleEndian, &wordlength)
		binary.Read(resultBuf, binary.LittleEndian, &duplex)

	}

	return baudrate, parity, stopbits, wordlength, duplex, nil
}

// Sets the configuration for the RS485 Modbus communication. Available options:
//
// * Slave Address: Address to be used as the Modbus slave address in Modbus slave mode. Valid Modbus slave address range is 1 to 247.
// * Master Request Timeout: Specifies how long the master should wait for a response from a slave when in Modbus master mode.
func (device *RS485Bricklet) SetModbusConfiguration(slaveAddress uint8, masterRequestTimeout uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, slaveAddress)
	binary.Write(&buf, binary.LittleEndian, masterRequestTimeout)

	resultBytes, err := device.device.Set(uint8(FunctionSetModbusConfiguration), buf.Bytes())
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

// Returns the configuration as set by SetModbusConfiguration.
func (device *RS485Bricklet) GetModbusConfiguration() (slaveAddress uint8, masterRequestTimeout uint32, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetModbusConfiguration), buf.Bytes())
	if err != nil {
		return slaveAddress, masterRequestTimeout, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return slaveAddress, masterRequestTimeout, DeviceError(header.ErrorCode)
		}

		if header.Length != 13 {
			return slaveAddress, masterRequestTimeout, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 13)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &slaveAddress)
		binary.Read(resultBuf, binary.LittleEndian, &masterRequestTimeout)

	}

	return slaveAddress, masterRequestTimeout, nil
}

// Sets the mode of the Bricklet in which it operates. Available options are
//
// * RS485,
// * Modbus Master RTU and
// * Modbus Slave RTU.
//
// Associated constants:
//
//	* ModeRS485
//	* ModeModbusMasterRTU
//	* ModeModbusSlaveRTU
func (device *RS485Bricklet) SetMode(mode Mode) (err error) {
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

// Returns the configuration as set by SetMode.
//
// Associated constants:
//
//	* ModeRS485
//	* ModeModbusMasterRTU
//	* ModeModbusSlaveRTU
func (device *RS485Bricklet) GetMode() (mode Mode, err error) {
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

// Sets the communication LED configuration. By default the LED shows RS485
// communication traffic by flickering.
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
func (device *RS485Bricklet) SetCommunicationLEDConfig(config CommunicationLEDConfig) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, config)

	resultBytes, err := device.device.Set(uint8(FunctionSetCommunicationLEDConfig), buf.Bytes())
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

// Returns the configuration as set by SetCommunicationLEDConfig
//
// Associated constants:
//
//	* CommunicationLEDConfigOff
//	* CommunicationLEDConfigOn
//	* CommunicationLEDConfigShowHeartbeat
//	* CommunicationLEDConfigShowCommunication
func (device *RS485Bricklet) GetCommunicationLEDConfig() (config CommunicationLEDConfig, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetCommunicationLEDConfig), buf.Bytes())
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

// Sets the error LED configuration.
//
// By default the error LED turns on if there is any error (see RegisterErrorCountCallback
// callback). If you call this function with the SHOW ERROR option again, the LED
// will turn off until the next error occurs.
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
//	* ErrorLEDConfigShowError
func (device *RS485Bricklet) SetErrorLEDConfig(config ErrorLEDConfig) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, config)

	resultBytes, err := device.device.Set(uint8(FunctionSetErrorLEDConfig), buf.Bytes())
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

// Returns the configuration as set by SetErrorLEDConfig.
//
// Associated constants:
//
//	* ErrorLEDConfigOff
//	* ErrorLEDConfigOn
//	* ErrorLEDConfigShowHeartbeat
//	* ErrorLEDConfigShowError
func (device *RS485Bricklet) GetErrorLEDConfig() (config ErrorLEDConfig, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetErrorLEDConfig), buf.Bytes())
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

// Sets the send and receive buffer size in byte. In sum there is
// 10240 byte (10KiB) buffer available and the minimum buffer size
// is 1024 byte (1KiB) for both.
//
// The current buffer content is lost if this function is called.
//
// The send buffer holds data that was given by Write and
// could not be written yet. The receive buffer holds data that is
// received through RS485 but could not yet be send to the
// user, either by Read or through RegisterReadCallback callback.
func (device *RS485Bricklet) SetBufferConfig(sendBufferSize uint16, receiveBufferSize uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, sendBufferSize)
	binary.Write(&buf, binary.LittleEndian, receiveBufferSize)

	resultBytes, err := device.device.Set(uint8(FunctionSetBufferConfig), buf.Bytes())
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

// Returns the buffer configuration as set by SetBufferConfig.
func (device *RS485Bricklet) GetBufferConfig() (sendBufferSize uint16, receiveBufferSize uint16, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetBufferConfig), buf.Bytes())
	if err != nil {
		return sendBufferSize, receiveBufferSize, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return sendBufferSize, receiveBufferSize, DeviceError(header.ErrorCode)
		}

		if header.Length != 12 {
			return sendBufferSize, receiveBufferSize, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
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
func (device *RS485Bricklet) GetBufferStatus() (sendBufferUsed uint16, receiveBufferUsed uint16, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetBufferStatus), buf.Bytes())
	if err != nil {
		return sendBufferUsed, receiveBufferUsed, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return sendBufferUsed, receiveBufferUsed, DeviceError(header.ErrorCode)
		}

		if header.Length != 12 {
			return sendBufferUsed, receiveBufferUsed, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &sendBufferUsed)
		binary.Read(resultBuf, binary.LittleEndian, &receiveBufferUsed)

	}

	return sendBufferUsed, receiveBufferUsed, nil
}

// Enables the RegisterErrorCountCallback callback.
//
// By default the callback is disabled.
func (device *RS485Bricklet) EnableErrorCountCallback() (err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Set(uint8(FunctionEnableErrorCountCallback), buf.Bytes())
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

// Disables the RegisterErrorCountCallback callback.
//
// By default the callback is disabled.
func (device *RS485Bricklet) DisableErrorCountCallback() (err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Set(uint8(FunctionDisableErrorCountCallback), buf.Bytes())
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

// Returns *true* if the RegisterErrorCountCallback callback is enabled,
// *false* otherwise.
func (device *RS485Bricklet) IsErrorCountCallbackEnabled() (enabled bool, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionIsErrorCountCallbackEnabled), buf.Bytes())
	if err != nil {
		return enabled, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return enabled, DeviceError(header.ErrorCode)
		}

		if header.Length != 9 {
			return enabled, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &enabled)

	}

	return enabled, nil
}

// Returns the current number of overrun and parity errors.
func (device *RS485Bricklet) GetErrorCount() (overrunErrorCount uint32, parityErrorCount uint32, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetErrorCount), buf.Bytes())
	if err != nil {
		return overrunErrorCount, parityErrorCount, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return overrunErrorCount, parityErrorCount, DeviceError(header.ErrorCode)
		}

		if header.Length != 16 {
			return overrunErrorCount, parityErrorCount, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 16)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &overrunErrorCount)
		binary.Read(resultBuf, binary.LittleEndian, &parityErrorCount)

	}

	return overrunErrorCount, parityErrorCount, nil
}

// Returns the current number of errors occurred in Modbus mode.
//
// * Timeout Error Count: Number of timeouts occurred.
// * Checksum Error Count: Number of failures due to Modbus frame CRC16 checksum mismatch.
// * Frame Too Big Error Count: Number of times frames were rejected because they exceeded maximum Modbus frame size which is 256 bytes.
// * Illegal Function Error Count: Number of errors when an unimplemented or illegal function is requested. This corresponds to Modbus exception code 1.
// * Illegal Data Address Error Count: Number of errors due to invalid data address. This corresponds to Modbus exception code 2.
// * Illegal Data Value Error Count: Number of errors due to invalid data value. This corresponds to Modbus exception code 3.
// * Slave Device Failure Error Count: Number of errors occurred on the slave device which were unrecoverable. This corresponds to Modbus exception code 4.
func (device *RS485Bricklet) GetModbusCommonErrorCount() (timeoutErrorCount uint32, checksumErrorCount uint32, frameTooBigErrorCount uint32, illegalFunctionErrorCount uint32, illegalDataAddressErrorCount uint32, illegalDataValueErrorCount uint32, slaveDeviceFailureErrorCount uint32, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetModbusCommonErrorCount), buf.Bytes())
	if err != nil {
		return timeoutErrorCount, checksumErrorCount, frameTooBigErrorCount, illegalFunctionErrorCount, illegalDataAddressErrorCount, illegalDataValueErrorCount, slaveDeviceFailureErrorCount, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return timeoutErrorCount, checksumErrorCount, frameTooBigErrorCount, illegalFunctionErrorCount, illegalDataAddressErrorCount, illegalDataValueErrorCount, slaveDeviceFailureErrorCount, DeviceError(header.ErrorCode)
		}

		if header.Length != 36 {
			return timeoutErrorCount, checksumErrorCount, frameTooBigErrorCount, illegalFunctionErrorCount, illegalDataAddressErrorCount, illegalDataValueErrorCount, slaveDeviceFailureErrorCount, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 36)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &timeoutErrorCount)
		binary.Read(resultBuf, binary.LittleEndian, &checksumErrorCount)
		binary.Read(resultBuf, binary.LittleEndian, &frameTooBigErrorCount)
		binary.Read(resultBuf, binary.LittleEndian, &illegalFunctionErrorCount)
		binary.Read(resultBuf, binary.LittleEndian, &illegalDataAddressErrorCount)
		binary.Read(resultBuf, binary.LittleEndian, &illegalDataValueErrorCount)
		binary.Read(resultBuf, binary.LittleEndian, &slaveDeviceFailureErrorCount)

	}

	return timeoutErrorCount, checksumErrorCount, frameTooBigErrorCount, illegalFunctionErrorCount, illegalDataAddressErrorCount, illegalDataValueErrorCount, slaveDeviceFailureErrorCount, nil
}

// In Modbus slave mode this function can be used to report a Modbus exception for
// a Modbus master request.
//
// * Request ID: Request ID of the request received by the slave.
// * Exception Code: Modbus exception code to report to the Modbus master.
//
// Associated constants:
//
//	* ExceptionCodeTimeout
//	* ExceptionCodeSuccess
//	* ExceptionCodeIllegalFunction
//	* ExceptionCodeIllegalDataAddress
//	* ExceptionCodeIllegalDataValue
//	* ExceptionCodeSlaveDeviceFailure
//	* ExceptionCodeAcknowledge
//	* ExceptionCodeSlaveDeviceBusy
//	* ExceptionCodeMemoryParityError
//	* ExceptionCodeGatewayPathUnavailable
//	* ExceptionCodeGatewayTargetDeviceFailedToRespond
func (device *RS485Bricklet) ModbusSlaveReportException(requestID uint8, exceptionCode ExceptionCode) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, requestID)
	binary.Write(&buf, binary.LittleEndian, exceptionCode)

	resultBytes, err := device.device.Set(uint8(FunctionModbusSlaveReportException), buf.Bytes())
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

// In Modbus slave mode this function can be used to answer a master request to
// read coils.
//
// * Request ID: Request ID of the corresponding request that is being answered.
// * Coils: Data that is to be sent to the Modbus master for the corresponding request.
//
// This function must be called from the RegisterModbusSlaveReadCoilsRequestCallback callback
// with the Request ID as provided by the argument of the callback.
func (device *RS485Bricklet) ModbusSlaveAnswerReadCoilsRequestLowLevel(requestID uint8, coilsLength uint16, coilsChunkOffset uint16, coilsChunkData [472]bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, requestID)
	binary.Write(&buf, binary.LittleEndian, coilsLength)
	binary.Write(&buf, binary.LittleEndian, coilsChunkOffset)
	buf.Write(BoolSliceToByteSlice(coilsChunkData[:]))

	resultBytes, err := device.device.Set(uint8(FunctionModbusSlaveAnswerReadCoilsRequestLowLevel), buf.Bytes())
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

// In Modbus slave mode this function can be used to answer a master request to
// read coils.
//
// * Request ID: Request ID of the corresponding request that is being answered.
// * Coils: Data that is to be sent to the Modbus master for the corresponding request.
//
// This function must be called from the RegisterModbusSlaveReadCoilsRequestCallback callback
// with the Request ID as provided by the argument of the callback.
func (device *RS485Bricklet) ModbusSlaveAnswerReadCoilsRequest(requestID uint8, coils []bool) (err error) {
	_, err = device.device.SetHighLevel(func(coilsLength uint64, coilsChunkOffset uint64, coilsChunkData []byte) (LowLevelWriteResult, error) {
		arr := [472]bool{}
		copy(arr[:], ByteSliceToBoolSlice(coilsChunkData))

		err := device.ModbusSlaveAnswerReadCoilsRequestLowLevel(requestID, uint16(coilsLength), uint16(coilsChunkOffset), arr)

		var lowLevelResults bytes.Buffer

		return LowLevelWriteResult{
			uint64(472),
			lowLevelResults.Bytes()}, err
	}, 2, 1, 472, BoolSliceToByteSlice(coils))

	if err != nil {
		return
	}

	return
}

// In Modbus master mode this function can be used to read coils from a slave. This
// function creates a Modbus function code 1 request.
//
// * Slave Address: Address of the target Modbus slave.
// * Starting Address: Number of the first coil to read. For backwards compatibility reasons this parameter is called Starting Address. It is not an address, but instead a coil number in the range of 1 to 65536.
// * Count: Number of coils to read.
//
// Upon success the function will return a non-zero request ID which will represent
// the current request initiated by the Modbus master. In case of failure the returned
// request ID will be 0.
//
// When successful this function will also invoke the RegisterModbusMasterReadCoilsResponseCallback
// callback. In this callback the Request ID provided by the callback argument must be
// matched with the Request ID returned from this function to verify that the callback
// is indeed for a particular request.
func (device *RS485Bricklet) ModbusMasterReadCoils(slaveAddress uint8, startingAddress uint32, count uint16) (requestID uint8, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, slaveAddress)
	binary.Write(&buf, binary.LittleEndian, startingAddress)
	binary.Write(&buf, binary.LittleEndian, count)

	resultBytes, err := device.device.Get(uint8(FunctionModbusMasterReadCoils), buf.Bytes())
	if err != nil {
		return requestID, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return requestID, DeviceError(header.ErrorCode)
		}

		if header.Length != 9 {
			return requestID, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &requestID)

	}

	return requestID, nil
}

// In Modbus slave mode this function can be used to answer a master request to
// read holding registers.
//
// * Request ID: Request ID of the corresponding request that is being answered.
// * Holding Registers: Data that is to be sent to the Modbus master for the corresponding request.
//
// This function must be called from the RegisterModbusSlaveReadHoldingRegistersRequestCallback
// callback with the Request ID as provided by the argument of the callback.
func (device *RS485Bricklet) ModbusSlaveAnswerReadHoldingRegistersRequestLowLevel(requestID uint8, holdingRegistersLength uint16, holdingRegistersChunkOffset uint16, holdingRegistersChunkData [29]uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, requestID)
	binary.Write(&buf, binary.LittleEndian, holdingRegistersLength)
	binary.Write(&buf, binary.LittleEndian, holdingRegistersChunkOffset)
	buf.Write(Uint16SliceToByteSlice(holdingRegistersChunkData[:]))

	resultBytes, err := device.device.Set(uint8(FunctionModbusSlaveAnswerReadHoldingRegistersRequestLowLevel), buf.Bytes())
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

// In Modbus slave mode this function can be used to answer a master request to
// read holding registers.
//
// * Request ID: Request ID of the corresponding request that is being answered.
// * Holding Registers: Data that is to be sent to the Modbus master for the corresponding request.
//
// This function must be called from the RegisterModbusSlaveReadHoldingRegistersRequestCallback
// callback with the Request ID as provided by the argument of the callback.
func (device *RS485Bricklet) ModbusSlaveAnswerReadHoldingRegistersRequest(requestID uint8, holdingRegisters []uint16) (err error) {
	_, err = device.device.SetHighLevel(func(holdingRegistersLength uint64, holdingRegistersChunkOffset uint64, holdingRegistersChunkData []byte) (LowLevelWriteResult, error) {
		arr := [29]uint16{}
		copy(arr[:], ByteSliceToUint16Slice(holdingRegistersChunkData))

		err := device.ModbusSlaveAnswerReadHoldingRegistersRequestLowLevel(requestID, uint16(holdingRegistersLength), uint16(holdingRegistersChunkOffset), arr)

		var lowLevelResults bytes.Buffer

		return LowLevelWriteResult{
			uint64(29),
			lowLevelResults.Bytes()}, err
	}, 3, 16, 464, Uint16SliceToByteSlice(holdingRegisters))

	if err != nil {
		return
	}

	return
}

// In Modbus master mode this function can be used to read holding registers from a slave.
// This function creates a Modbus function code 3 request.
//
// * Slave Address: Address of the target Modbus slave.
// * Starting Address: Number of the first holding register to read. For backwards compatibility reasons this parameter is called Starting Address. It is not an address, but instead a holding register number in the range of 1 to 65536. The prefix digit 4 (for holding register) is implicit and must be omitted.
// * Count: Number of holding registers to read.
//
// Upon success the function will return a non-zero request ID which will represent
// the current request initiated by the Modbus master. In case of failure the returned
// request ID will be 0.
//
// When successful this function will also invoke the RegisterModbusMasterReadHoldingRegistersResponseCallback
// callback. In this callback the Request ID provided by the callback argument must be matched
// with the Request ID returned from this function to verify that the callback is indeed for a
// particular request.
func (device *RS485Bricklet) ModbusMasterReadHoldingRegisters(slaveAddress uint8, startingAddress uint32, count uint16) (requestID uint8, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, slaveAddress)
	binary.Write(&buf, binary.LittleEndian, startingAddress)
	binary.Write(&buf, binary.LittleEndian, count)

	resultBytes, err := device.device.Get(uint8(FunctionModbusMasterReadHoldingRegisters), buf.Bytes())
	if err != nil {
		return requestID, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return requestID, DeviceError(header.ErrorCode)
		}

		if header.Length != 9 {
			return requestID, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &requestID)

	}

	return requestID, nil
}

// In Modbus slave mode this function can be used to answer a master request to
// write a single coil.
//
// * Request ID: Request ID of the corresponding request that is being answered.
//
// This function must be called from the RegisterModbusSlaveWriteSingleCoilRequestCallback
// callback with the Request ID as provided by the arguments of the callback.
func (device *RS485Bricklet) ModbusSlaveAnswerWriteSingleCoilRequest(requestID uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, requestID)

	resultBytes, err := device.device.Set(uint8(FunctionModbusSlaveAnswerWriteSingleCoilRequest), buf.Bytes())
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

// In Modbus master mode this function can be used to write a single coil of a slave.
// This function creates a Modbus function code 5 request.
//
// * Slave Address: Address of the target Modbus slave.
// * Coil Address: Number of the coil to be written. For backwards compatibility reasons, this parameter is called Starting Address. It is not an address, but instead a coil number in the range of 1 to 65536.
// * Coil Value: Value to be written.
//
// Upon success the function will return a non-zero request ID which will represent
// the current request initiated by the Modbus master. In case of failure the returned
// request ID will be 0.
//
// When successful this function will also invoke the RegisterModbusMasterWriteSingleCoilResponseCallback
// callback. In this callback the Request ID provided by the callback argument must be matched
// with the Request ID returned from this function to verify that the callback is indeed for a
// particular request.
func (device *RS485Bricklet) ModbusMasterWriteSingleCoil(slaveAddress uint8, coilAddress uint32, coilValue bool) (requestID uint8, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, slaveAddress)
	binary.Write(&buf, binary.LittleEndian, coilAddress)
	binary.Write(&buf, binary.LittleEndian, coilValue)

	resultBytes, err := device.device.Get(uint8(FunctionModbusMasterWriteSingleCoil), buf.Bytes())
	if err != nil {
		return requestID, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return requestID, DeviceError(header.ErrorCode)
		}

		if header.Length != 9 {
			return requestID, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &requestID)

	}

	return requestID, nil
}

// In Modbus slave mode this function can be used to answer a master request to
// write a single register.
//
// * Request ID: Request ID of the corresponding request that is being answered.
//
// This function must be called from the RegisterModbusSlaveWriteSingleRegisterRequestCallback
// callback with the Request ID, Register Address and Register Value as provided by
// the arguments of the callback.
func (device *RS485Bricklet) ModbusSlaveAnswerWriteSingleRegisterRequest(requestID uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, requestID)

	resultBytes, err := device.device.Set(uint8(FunctionModbusSlaveAnswerWriteSingleRegisterRequest), buf.Bytes())
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

// In Modbus master mode this function can be used to write a single holding register of a
// slave. This function creates a Modbus function code 6 request.
//
// * Slave Address: Address of the target Modbus slave.
// * Register Address: Number of the holding register to be written. For backwards compatibility reasons, this parameter is called Starting Address. It is not an address, but instead a holding register number in the range of 1 to 65536. The prefix digit 4 (for holding register) is implicit and must be omitted.
// * Register Value: Value to be written.
//
// Upon success the function will return a non-zero request ID which will represent
// the current request initiated by the Modbus master. In case of failure the returned
// request ID will be 0.
//
// When successful this function will also invoke the RegisterModbusMasterWriteSingleRegisterResponseCallback
// callback. In this callback the Request ID provided by the callback argument must be matched
// with the Request ID returned from this function to verify that the callback is indeed for a
// particular request.
func (device *RS485Bricklet) ModbusMasterWriteSingleRegister(slaveAddress uint8, registerAddress uint32, registerValue uint16) (requestID uint8, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, slaveAddress)
	binary.Write(&buf, binary.LittleEndian, registerAddress)
	binary.Write(&buf, binary.LittleEndian, registerValue)

	resultBytes, err := device.device.Get(uint8(FunctionModbusMasterWriteSingleRegister), buf.Bytes())
	if err != nil {
		return requestID, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return requestID, DeviceError(header.ErrorCode)
		}

		if header.Length != 9 {
			return requestID, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &requestID)

	}

	return requestID, nil
}

// In Modbus slave mode this function can be used to answer a master request to
// write multiple coils.
//
// * Request ID: Request ID of the corresponding request that is being answered.
//
// This function must be called from the RegisterModbusSlaveWriteMultipleCoilsRequestCallback
// callback with the Request ID of the callback.
func (device *RS485Bricklet) ModbusSlaveAnswerWriteMultipleCoilsRequest(requestID uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, requestID)

	resultBytes, err := device.device.Set(uint8(FunctionModbusSlaveAnswerWriteMultipleCoilsRequest), buf.Bytes())
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

// In Modbus master mode this function can be used to write multiple coils of a slave.
// This function creates a Modbus function code 15 request.
//
// * Slave Address: Address of the target Modbus slave.
// * Starting Address: Number of the first coil to write. For backwards compatibility reasons, this parameter is called Starting Address.It is not an address, but instead a coil number in the range of 1 to 65536.
//
// Upon success the function will return a non-zero request ID which will represent
// the current request initiated by the Modbus master. In case of failure the returned
// request ID will be 0.
//
// When successful this function will also invoke the RegisterModbusMasterWriteMultipleCoilsResponseCallback
// callback. In this callback the Request ID provided by the callback argument must be matched
// with the Request ID returned from this function to verify that the callback is indeed for a
// particular request.
func (device *RS485Bricklet) ModbusMasterWriteMultipleCoilsLowLevel(slaveAddress uint8, startingAddress uint32, coilsLength uint16, coilsChunkOffset uint16, coilsChunkData [440]bool) (requestID uint8, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, slaveAddress)
	binary.Write(&buf, binary.LittleEndian, startingAddress)
	binary.Write(&buf, binary.LittleEndian, coilsLength)
	binary.Write(&buf, binary.LittleEndian, coilsChunkOffset)
	buf.Write(BoolSliceToByteSlice(coilsChunkData[:]))

	resultBytes, err := device.device.Get(uint8(FunctionModbusMasterWriteMultipleCoilsLowLevel), buf.Bytes())
	if err != nil {
		return requestID, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return requestID, DeviceError(header.ErrorCode)
		}

		if header.Length != 9 {
			return requestID, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &requestID)

	}

	return requestID, nil
}

// In Modbus master mode this function can be used to write multiple coils of a slave.
// This function creates a Modbus function code 15 request.
//
// * Slave Address: Address of the target Modbus slave.
// * Starting Address: Number of the first coil to write. For backwards compatibility reasons, this parameter is called Starting Address.It is not an address, but instead a coil number in the range of 1 to 65536.
//
// Upon success the function will return a non-zero request ID which will represent
// the current request initiated by the Modbus master. In case of failure the returned
// request ID will be 0.
//
// When successful this function will also invoke the RegisterModbusMasterWriteMultipleCoilsResponseCallback
// callback. In this callback the Request ID provided by the callback argument must be matched
// with the Request ID returned from this function to verify that the callback is indeed for a
// particular request.
func (device *RS485Bricklet) ModbusMasterWriteMultipleCoils(slaveAddress uint8, startingAddress uint32, coils []bool) (requestID uint8, err error) {
	lowLevelResult, err := device.device.SetHighLevel(func(coilsLength uint64, coilsChunkOffset uint64, coilsChunkData []byte) (LowLevelWriteResult, error) {
		arr := [440]bool{}
		copy(arr[:], ByteSliceToBoolSlice(coilsChunkData))

		requestID, err := device.ModbusMasterWriteMultipleCoilsLowLevel(slaveAddress, startingAddress, uint16(coilsLength), uint16(coilsChunkOffset), arr)

		var lowLevelResults bytes.Buffer
		binary.Write(&lowLevelResults, binary.LittleEndian, requestID)

		return LowLevelWriteResult{
			uint64(440),
			lowLevelResults.Bytes()}, err
	}, 4, 1, 440, BoolSliceToByteSlice(coils))

	if err != nil {
		return
	}

	resultBuf := bytes.NewBuffer(lowLevelResult.Result)

	binary.Read(resultBuf, binary.LittleEndian, &requestID)

	return
}

// In Modbus slave mode this function can be used to answer a master request to
// write multiple registers.
//
// * Request ID: Request ID of the corresponding request that is being answered.
//
// This function must be called from the RegisterModbusSlaveWriteMultipleRegistersRequestCallback
// callback with the Request ID of the callback.
func (device *RS485Bricklet) ModbusSlaveAnswerWriteMultipleRegistersRequest(requestID uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, requestID)

	resultBytes, err := device.device.Set(uint8(FunctionModbusSlaveAnswerWriteMultipleRegistersRequest), buf.Bytes())
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

// In Modbus master mode this function can be used to write multiple registers of a slave.
// This function creates a Modbus function code 16 request.
//
// * Slave Address: Address of the target Modbus slave.
// * Starting Address: Number of the first holding register to write. For backwards compatibility reasons, this parameter is called Starting Address. It is not an address, but instead a holding register number in the range of 1 to 65536. The prefix digit 4 (for holding register) is implicit and must be omitted.
//
// Upon success the function will return a non-zero request ID which will represent
// the current request initiated by the Modbus master. In case of failure the returned
// request ID will be 0.
//
// When successful this function will also invoke the RegisterModbusMasterWriteMultipleRegistersResponseCallback
// callback. In this callback the Request ID provided by the callback argument must be matched
// with the Request ID returned from this function to verify that the callback is indeed for a
// particular request.
func (device *RS485Bricklet) ModbusMasterWriteMultipleRegistersLowLevel(slaveAddress uint8, startingAddress uint32, registersLength uint16, registersChunkOffset uint16, registersChunkData [27]uint16) (requestID uint8, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, slaveAddress)
	binary.Write(&buf, binary.LittleEndian, startingAddress)
	binary.Write(&buf, binary.LittleEndian, registersLength)
	binary.Write(&buf, binary.LittleEndian, registersChunkOffset)
	buf.Write(Uint16SliceToByteSlice(registersChunkData[:]))

	resultBytes, err := device.device.Get(uint8(FunctionModbusMasterWriteMultipleRegistersLowLevel), buf.Bytes())
	if err != nil {
		return requestID, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return requestID, DeviceError(header.ErrorCode)
		}

		if header.Length != 9 {
			return requestID, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &requestID)

	}

	return requestID, nil
}

// In Modbus master mode this function can be used to write multiple registers of a slave.
// This function creates a Modbus function code 16 request.
//
// * Slave Address: Address of the target Modbus slave.
// * Starting Address: Number of the first holding register to write. For backwards compatibility reasons, this parameter is called Starting Address. It is not an address, but instead a holding register number in the range of 1 to 65536. The prefix digit 4 (for holding register) is implicit and must be omitted.
//
// Upon success the function will return a non-zero request ID which will represent
// the current request initiated by the Modbus master. In case of failure the returned
// request ID will be 0.
//
// When successful this function will also invoke the RegisterModbusMasterWriteMultipleRegistersResponseCallback
// callback. In this callback the Request ID provided by the callback argument must be matched
// with the Request ID returned from this function to verify that the callback is indeed for a
// particular request.
func (device *RS485Bricklet) ModbusMasterWriteMultipleRegisters(slaveAddress uint8, startingAddress uint32, registers []uint16) (requestID uint8, err error) {
	lowLevelResult, err := device.device.SetHighLevel(func(registersLength uint64, registersChunkOffset uint64, registersChunkData []byte) (LowLevelWriteResult, error) {
		arr := [27]uint16{}
		copy(arr[:], ByteSliceToUint16Slice(registersChunkData))

		requestID, err := device.ModbusMasterWriteMultipleRegistersLowLevel(slaveAddress, startingAddress, uint16(registersLength), uint16(registersChunkOffset), arr)

		var lowLevelResults bytes.Buffer
		binary.Write(&lowLevelResults, binary.LittleEndian, requestID)

		return LowLevelWriteResult{
			uint64(27),
			lowLevelResults.Bytes()}, err
	}, 5, 16, 432, Uint16SliceToByteSlice(registers))

	if err != nil {
		return
	}

	resultBuf := bytes.NewBuffer(lowLevelResult.Result)

	binary.Read(resultBuf, binary.LittleEndian, &requestID)

	return
}

// In Modbus slave mode this function can be used to answer a master request to
// read discrete inputs.
//
// * Request ID: Request ID of the corresponding request that is being answered.
// * Discrete Inputs: Data that is to be sent to the Modbus master for the corresponding request.
//
// This function must be called from the RegisterModbusSlaveReadDiscreteInputsRequestCallback
// callback with the Request ID as provided by the argument of the callback.
func (device *RS485Bricklet) ModbusSlaveAnswerReadDiscreteInputsRequestLowLevel(requestID uint8, discreteInputsLength uint16, discreteInputsChunkOffset uint16, discreteInputsChunkData [472]bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, requestID)
	binary.Write(&buf, binary.LittleEndian, discreteInputsLength)
	binary.Write(&buf, binary.LittleEndian, discreteInputsChunkOffset)
	buf.Write(BoolSliceToByteSlice(discreteInputsChunkData[:]))

	resultBytes, err := device.device.Set(uint8(FunctionModbusSlaveAnswerReadDiscreteInputsRequestLowLevel), buf.Bytes())
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

// In Modbus slave mode this function can be used to answer a master request to
// read discrete inputs.
//
// * Request ID: Request ID of the corresponding request that is being answered.
// * Discrete Inputs: Data that is to be sent to the Modbus master for the corresponding request.
//
// This function must be called from the RegisterModbusSlaveReadDiscreteInputsRequestCallback
// callback with the Request ID as provided by the argument of the callback.
func (device *RS485Bricklet) ModbusSlaveAnswerReadDiscreteInputsRequest(requestID uint8, discreteInputs []bool) (err error) {
	_, err = device.device.SetHighLevel(func(discreteInputsLength uint64, discreteInputsChunkOffset uint64, discreteInputsChunkData []byte) (LowLevelWriteResult, error) {
		arr := [472]bool{}
		copy(arr[:], ByteSliceToBoolSlice(discreteInputsChunkData))

		err := device.ModbusSlaveAnswerReadDiscreteInputsRequestLowLevel(requestID, uint16(discreteInputsLength), uint16(discreteInputsChunkOffset), arr)

		var lowLevelResults bytes.Buffer

		return LowLevelWriteResult{
			uint64(472),
			lowLevelResults.Bytes()}, err
	}, 6, 1, 472, BoolSliceToByteSlice(discreteInputs))

	if err != nil {
		return
	}

	return
}

// In Modbus master mode this function can be used to read discrete inputs from a slave.
// This function creates a Modbus function code 2 request.
//
// * Slave Address: Address of the target Modbus slave.
// * Starting Address: Number of the first discrete input to read. For backwards compatibility reasons, this parameter is called Starting Address. It is not an address, but instead a discrete input number in the range of 1 to 65536. The prefix digit 1 (for discrete input) is implicit and must be omitted.
// * Count: Number of discrete inputs to read.
//
// Upon success the function will return a non-zero request ID which will represent
// the current request initiated by the Modbus master. In case of failure the returned
// request ID will be 0.
//
// When successful this function will also invoke the RegisterModbusMasterReadDiscreteInputsResponseCallback
// callback. In this callback the Request ID provided by the callback argument must be matched
// with the Request ID returned from this function to verify that the callback is indeed for a
// particular request.
func (device *RS485Bricklet) ModbusMasterReadDiscreteInputs(slaveAddress uint8, startingAddress uint32, count uint16) (requestID uint8, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, slaveAddress)
	binary.Write(&buf, binary.LittleEndian, startingAddress)
	binary.Write(&buf, binary.LittleEndian, count)

	resultBytes, err := device.device.Get(uint8(FunctionModbusMasterReadDiscreteInputs), buf.Bytes())
	if err != nil {
		return requestID, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return requestID, DeviceError(header.ErrorCode)
		}

		if header.Length != 9 {
			return requestID, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &requestID)

	}

	return requestID, nil
}

// In Modbus slave mode this function can be used to answer a master request to
// read input registers.
//
// * Request ID: Request ID of the corresponding request that is being answered.
// * Input Registers: Data that is to be sent to the Modbus master for the corresponding request.
//
// This function must be called from the RegisterModbusSlaveReadInputRegistersRequestCallback callback
// with the Request ID as provided by the argument of the callback.
func (device *RS485Bricklet) ModbusSlaveAnswerReadInputRegistersRequestLowLevel(requestID uint8, inputRegistersLength uint16, inputRegistersChunkOffset uint16, inputRegistersChunkData [29]uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, requestID)
	binary.Write(&buf, binary.LittleEndian, inputRegistersLength)
	binary.Write(&buf, binary.LittleEndian, inputRegistersChunkOffset)
	buf.Write(Uint16SliceToByteSlice(inputRegistersChunkData[:]))

	resultBytes, err := device.device.Set(uint8(FunctionModbusSlaveAnswerReadInputRegistersRequestLowLevel), buf.Bytes())
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

// In Modbus slave mode this function can be used to answer a master request to
// read input registers.
//
// * Request ID: Request ID of the corresponding request that is being answered.
// * Input Registers: Data that is to be sent to the Modbus master for the corresponding request.
//
// This function must be called from the RegisterModbusSlaveReadInputRegistersRequestCallback callback
// with the Request ID as provided by the argument of the callback.
func (device *RS485Bricklet) ModbusSlaveAnswerReadInputRegistersRequest(requestID uint8, inputRegisters []uint16) (err error) {
	_, err = device.device.SetHighLevel(func(inputRegistersLength uint64, inputRegistersChunkOffset uint64, inputRegistersChunkData []byte) (LowLevelWriteResult, error) {
		arr := [29]uint16{}
		copy(arr[:], ByteSliceToUint16Slice(inputRegistersChunkData))

		err := device.ModbusSlaveAnswerReadInputRegistersRequestLowLevel(requestID, uint16(inputRegistersLength), uint16(inputRegistersChunkOffset), arr)

		var lowLevelResults bytes.Buffer

		return LowLevelWriteResult{
			uint64(29),
			lowLevelResults.Bytes()}, err
	}, 7, 16, 464, Uint16SliceToByteSlice(inputRegisters))

	if err != nil {
		return
	}

	return
}

// In Modbus master mode this function can be used to read input registers from a slave.
// This function creates a Modbus function code 4 request.
//
// * Slave Address: Address of the target Modbus slave.
// * Starting Address: Number of the first input register to read. For backwards compatibility reasons, this parameter is called Starting Address. It is not an address, but instead an input register number in the range of 1 to 65536. The prefix digit 3 (for input register) is implicit and must be omitted.
// * Count: Number of input registers to read.
//
// Upon success the function will return a non-zero request ID which will represent
// the current request initiated by the Modbus master. In case of failure the returned
// request ID will be 0.
//
// When successful this function will also invoke the RegisterModbusMasterReadInputRegistersResponseCallback
// callback. In this callback the Request ID provided by the callback argument must be matched
// with the Request ID returned from this function to verify that the callback is indeed for a
// particular request.
func (device *RS485Bricklet) ModbusMasterReadInputRegisters(slaveAddress uint8, startingAddress uint32, count uint16) (requestID uint8, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, slaveAddress)
	binary.Write(&buf, binary.LittleEndian, startingAddress)
	binary.Write(&buf, binary.LittleEndian, count)

	resultBytes, err := device.device.Get(uint8(FunctionModbusMasterReadInputRegisters), buf.Bytes())
	if err != nil {
		return requestID, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return requestID, DeviceError(header.ErrorCode)
		}

		if header.Length != 9 {
			return requestID, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &requestID)

	}

	return requestID, nil
}

// Configures the RegisterFrameReadableCallback callback. The frame size is the number of bytes, that have to be readable to trigger the callback.
// A frame size of 0 disables the callback. A frame size greater than 0 enables the callback and disables the RegisterReadCallback callback.
//
// By default the callback is disabled.
//
// .. versionadded:: 2.0.5$nbsp;(Plugin)
func (device *RS485Bricklet) SetFrameReadableCallbackConfiguration(frameSize uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, frameSize)

	resultBytes, err := device.device.Set(uint8(FunctionSetFrameReadableCallbackConfiguration), buf.Bytes())
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

// Returns the callback configuration as set by SetFrameReadableCallbackConfiguration.
//
// .. versionadded:: 2.0.5$nbsp;(Plugin)
func (device *RS485Bricklet) GetFrameReadableCallbackConfiguration() (frameSize uint16, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetFrameReadableCallbackConfiguration), buf.Bytes())
	if err != nil {
		return frameSize, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return frameSize, DeviceError(header.ErrorCode)
		}

		if header.Length != 10 {
			return frameSize, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
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
func (device *RS485Bricklet) GetSPITFPErrorCount() (errorCountAckChecksum uint32, errorCountMessageChecksum uint32, errorCountFrame uint32, errorCountOverflow uint32, err error) {
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
func (device *RS485Bricklet) SetBootloaderMode(mode BootloaderMode) (status BootloaderStatus, err error) {
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
func (device *RS485Bricklet) GetBootloaderMode() (mode BootloaderMode, err error) {
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
func (device *RS485Bricklet) SetWriteFirmwarePointer(pointer uint32) (err error) {
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
func (device *RS485Bricklet) WriteFirmware(data [64]uint8) (status uint8, err error) {
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
func (device *RS485Bricklet) SetStatusLEDConfig(config StatusLEDConfig) (err error) {
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
func (device *RS485Bricklet) GetStatusLEDConfig() (config StatusLEDConfig, err error) {
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
func (device *RS485Bricklet) GetChipTemperature() (temperature int16, err error) {
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
func (device *RS485Bricklet) Reset() (err error) {
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
func (device *RS485Bricklet) WriteUID(uid uint32) (err error) {
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
func (device *RS485Bricklet) ReadUID() (uid uint32, err error) {
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
func (device *RS485Bricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
