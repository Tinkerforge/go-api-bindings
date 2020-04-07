/* ***********************************************************
 * This file was automatically generated on 2020-04-07.      *
 *                                                           *
 * Go Bindings Version 2.0.6                                 *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


// Executes user programs and controls other Bricks/Bricklets standalone.
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricks/RED_Brick_Go.html.
package red_brick

import (
	"encoding/binary"
	"bytes"
	"fmt"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionCreateSession Function = 1
	FunctionExpireSession Function = 2
	FunctionExpireSessionUnchecked Function = 3
	FunctionKeepSessionAlive Function = 4
	FunctionReleaseObject Function = 5
	FunctionReleaseObjectUnchecked Function = 6
	FunctionAllocateString Function = 7
	FunctionTruncateString Function = 8
	FunctionGetStringLength Function = 9
	FunctionSetStringChunk Function = 10
	FunctionGetStringChunk Function = 11
	FunctionAllocateList Function = 12
	FunctionGetListLength Function = 13
	FunctionGetListItem Function = 14
	FunctionAppendToList Function = 15
	FunctionRemoveFromList Function = 16
	FunctionOpenFile Function = 17
	FunctionCreatePipe Function = 18
	FunctionGetFileInfo Function = 19
	FunctionReadFile Function = 20
	FunctionReadFileAsync Function = 21
	FunctionAbortAsyncFileRead Function = 22
	FunctionWriteFile Function = 23
	FunctionWriteFileUnchecked Function = 24
	FunctionWriteFileAsync Function = 25
	FunctionSetFilePosition Function = 26
	FunctionGetFilePosition Function = 27
	FunctionSetFileEvents Function = 28
	FunctionGetFileEvents Function = 29
	FunctionOpenDirectory Function = 33
	FunctionGetDirectoryName Function = 34
	FunctionGetNextDirectoryEntry Function = 35
	FunctionRewindDirectory Function = 36
	FunctionCreateDirectory Function = 37
	FunctionGetProcesses Function = 38
	FunctionSpawnProcess Function = 39
	FunctionKillProcess Function = 40
	FunctionGetProcessCommand Function = 41
	FunctionGetProcessIdentity Function = 42
	FunctionGetProcessStdio Function = 43
	FunctionGetProcessState Function = 44
	FunctionGetPrograms Function = 46
	FunctionDefineProgram Function = 47
	FunctionPurgeProgram Function = 48
	FunctionGetProgramIdentifier Function = 49
	FunctionGetProgramRootDirectory Function = 50
	FunctionSetProgramCommand Function = 51
	FunctionGetProgramCommand Function = 52
	FunctionSetProgramStdioRedirection Function = 53
	FunctionGetProgramStdioRedirection Function = 54
	FunctionSetProgramSchedule Function = 55
	FunctionGetProgramSchedule Function = 56
	FunctionGetProgramSchedulerState Function = 57
	FunctionContinueProgramSchedule Function = 58
	FunctionStartProgram Function = 59
	FunctionGetLastSpawnedProgramProcess Function = 60
	FunctionGetCustomProgramOptionNames Function = 61
	FunctionSetCustomProgramOptionValue Function = 62
	FunctionGetCustomProgramOptionValue Function = 63
	FunctionRemoveCustomProgramOption Function = 64
	FunctionGetIdentity Function = 255
	FunctionCallbackAsyncFileRead Function = 30
	FunctionCallbackAsyncFileWrite Function = 31
	FunctionCallbackFileEventsOccurred Function = 32
	FunctionCallbackProcessStateChanged Function = 45
	FunctionCallbackProgramSchedulerStateChanged Function = 65
	FunctionCallbackProgramProcessSpawned Function = 66
)

type ErrorCode = uint8

const (
	ErrorCodeSuccess ErrorCode = 0
	ErrorCodeUnknownError ErrorCode = 1
	ErrorCodeInvalidOperation ErrorCode = 2
	ErrorCodeOperationAborted ErrorCode = 3
	ErrorCodeInternalError ErrorCode = 4
	ErrorCodeUnknownSessionId ErrorCode = 5
	ErrorCodeNoFreeSessionId ErrorCode = 6
	ErrorCodeUnknownObjectId ErrorCode = 7
	ErrorCodeNoFreeObjectId ErrorCode = 8
	ErrorCodeObjectIsLocked ErrorCode = 9
	ErrorCodeNoMoreData ErrorCode = 10
	ErrorCodeWrongListItemType ErrorCode = 11
	ErrorCodeProgramIsPurged ErrorCode = 12
	ErrorCodeInvalidParameter ErrorCode = 128
	ErrorCodeNoFreeMemory ErrorCode = 129
	ErrorCodeNoFreeSpace ErrorCode = 130
	ErrorCodeAccessDenied ErrorCode = 121
	ErrorCodeAlreadyExists ErrorCode = 132
	ErrorCodeDoesNotExist ErrorCode = 133
	ErrorCodeInterrupted ErrorCode = 134
	ErrorCodeIsDirectory ErrorCode = 135
	ErrorCodeNotADirectory ErrorCode = 136
	ErrorCodeWouldBlock ErrorCode = 137
	ErrorCodeOverflow ErrorCode = 138
	ErrorCodeBadFileDescriptor ErrorCode = 139
	ErrorCodeOutOfRange ErrorCode = 140
	ErrorCodeNameTooLong ErrorCode = 141
	ErrorCodeInvalidSeek ErrorCode = 142
	ErrorCodeNotSupported ErrorCode = 143
	ErrorCodeTooManyOpenFiles ErrorCode = 144
)

type ObjectType = uint8

const (
	ObjectTypeString ObjectType = 0
	ObjectTypeList ObjectType = 1
	ObjectTypeFile ObjectType = 2
	ObjectTypeDirectory ObjectType = 3
	ObjectTypeProcess ObjectType = 4
	ObjectTypeProgram ObjectType = 5
)

type FileFlag = uint32

const (
	FileFlagReadOnly FileFlag = 1
	FileFlagWriteOnly FileFlag = 2
	FileFlagReadWrite FileFlag = 4
	FileFlagAppend FileFlag = 8
	FileFlagCreate FileFlag = 16
	FileFlagExclusive FileFlag = 32
	FileFlagNonBlocking FileFlag = 64
	FileFlagTruncate FileFlag = 128
	FileFlagTemporary FileFlag = 256
	FileFlagReplace FileFlag = 512
)

type FilePermission = uint16

const (
	FilePermissionUserAll FilePermission = 448
	FilePermissionUserRead FilePermission = 256
	FilePermissionUserWrite FilePermission = 128
	FilePermissionUserExecute FilePermission = 64
	FilePermissionGroupAll FilePermission = 56
	FilePermissionGroupRead FilePermission = 32
	FilePermissionGroupWrite FilePermission = 16
	FilePermissionGroupExecute FilePermission = 8
	FilePermissionOthersAll FilePermission = 7
	FilePermissionOthersRead FilePermission = 4
	FilePermissionOthersWrite FilePermission = 2
	FilePermissionOthersExecute FilePermission = 1
)

type PipeFlag = uint32

const (
	PipeFlagNonBlockingRead PipeFlag = 1
	PipeFlagNonBlockingWrite PipeFlag = 2
)

type FileType = uint8

const (
	FileTypeUnknown FileType = 0
	FileTypeRegular FileType = 1
	FileTypeDirectory FileType = 2
	FileTypeCharacter FileType = 3
	FileTypeBlock FileType = 4
	FileTypeFIFO FileType = 5
	FileTypeSymlink FileType = 6
	FileTypeSocket FileType = 7
	FileTypePipe FileType = 8
)

type FileOrigin = uint8

const (
	FileOriginBeginning FileOrigin = 0
	FileOriginCurrent FileOrigin = 1
	FileOriginEnd FileOrigin = 2
)

type FileEvent = uint16

const (
	FileEventReadable FileEvent = 1
	FileEventWritable FileEvent = 2
)

type DirectoryEntryType = uint8

const (
	DirectoryEntryTypeUnknown DirectoryEntryType = 0
	DirectoryEntryTypeRegular DirectoryEntryType = 1
	DirectoryEntryTypeDirectory DirectoryEntryType = 2
	DirectoryEntryTypeCharacter DirectoryEntryType = 3
	DirectoryEntryTypeBlock DirectoryEntryType = 4
	DirectoryEntryTypeFIFO DirectoryEntryType = 5
	DirectoryEntryTypeSymlink DirectoryEntryType = 6
	DirectoryEntryTypeSocket DirectoryEntryType = 7
)

type DirectoryFlag = uint32

const (
	DirectoryFlagRecursive DirectoryFlag = 1
	DirectoryFlagExclusive DirectoryFlag = 2
)

type ProcessSignal = uint8

const (
	ProcessSignalInterrupt ProcessSignal = 2
	ProcessSignalQuit ProcessSignal = 3
	ProcessSignalAbort ProcessSignal = 6
	ProcessSignalKill ProcessSignal = 9
	ProcessSignalUser1 ProcessSignal = 10
	ProcessSignalUser2 ProcessSignal = 12
	ProcessSignalTerminate ProcessSignal = 15
	ProcessSignalContinue ProcessSignal = 18
	ProcessSignalStop ProcessSignal = 19
)

type ProcessState = uint8

const (
	ProcessStateUnknown ProcessState = 0
	ProcessStateRunning ProcessState = 1
	ProcessStateError ProcessState = 2
	ProcessStateExited ProcessState = 3
	ProcessStateKilled ProcessState = 4
	ProcessStateStopped ProcessState = 5
)

type ProgramStdioRedirection = uint8

const (
	ProgramStdioRedirectionDevNull ProgramStdioRedirection = 0
	ProgramStdioRedirectionPipe ProgramStdioRedirection = 1
	ProgramStdioRedirectionFile ProgramStdioRedirection = 2
	ProgramStdioRedirectionIndividualLog ProgramStdioRedirection = 3
	ProgramStdioRedirectionContinuousLog ProgramStdioRedirection = 4
	ProgramStdioRedirectionStdout ProgramStdioRedirection = 5
)

type ProgramStartMode = uint8

const (
	ProgramStartModeNever ProgramStartMode = 0
	ProgramStartModeAlways ProgramStartMode = 1
	ProgramStartModeInterval ProgramStartMode = 2
	ProgramStartModeCron ProgramStartMode = 3
)

type ProgramSchedulerState = uint8

const (
	ProgramSchedulerStateStopped ProgramSchedulerState = 0
	ProgramSchedulerStateRunning ProgramSchedulerState = 1
)

type REDBrick struct {
	device Device
}
const DeviceIdentifier = 17
const DeviceDisplayName = "RED Brick"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (REDBrick, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{ 2,0,0 }, uid, &internalIPCon, 0, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return REDBrick{}, err
	}
	dev.ResponseExpected[FunctionCreateSession] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionExpireSession] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionExpireSessionUnchecked] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionKeepSessionAlive] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionReleaseObject] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionReleaseObjectUnchecked] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionAllocateString] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionTruncateString] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetStringLength] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetStringChunk] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetStringChunk] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionAllocateList] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetListLength] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetListItem] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionAppendToList] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionRemoveFromList] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionOpenFile] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionCreatePipe] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetFileInfo] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionReadFile] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionReadFileAsync] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionAbortAsyncFileRead] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionWriteFile] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionWriteFileUnchecked] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionWriteFileAsync] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionSetFilePosition] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetFilePosition] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetFileEvents] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetFileEvents] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionOpenDirectory] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetDirectoryName] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetNextDirectoryEntry] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionRewindDirectory] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionCreateDirectory] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetProcesses] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSpawnProcess] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionKillProcess] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetProcessCommand] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetProcessIdentity] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetProcessStdio] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetProcessState] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetPrograms] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionDefineProgram] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionPurgeProgram] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetProgramIdentifier] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetProgramRootDirectory] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetProgramCommand] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetProgramCommand] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetProgramStdioRedirection] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetProgramStdioRedirection] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetProgramSchedule] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetProgramSchedule] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetProgramSchedulerState] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionContinueProgramSchedule] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionStartProgram] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetLastSpawnedProgramProcess] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetCustomProgramOptionNames] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetCustomProgramOptionValue] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetCustomProgramOptionValue] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionRemoveCustomProgramOption] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetIdentity] = ResponseExpectedFlagAlwaysTrue;
	return REDBrick{dev}, nil
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
func (device *REDBrick) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *REDBrick) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *REDBrick) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *REDBrick) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// This callback reports the result of a call to the ReadFileAsync
// function.
func (device *REDBrick) RegisterAsyncFileReadCallback(fn func(uint16, ErrorCode, [60]uint8, uint8)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 72 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var fileId uint16
		var errorCode ErrorCode
		var buffer [60]uint8
		var lengthRead uint8
		binary.Read(buf, binary.LittleEndian, &fileId)
		binary.Read(buf, binary.LittleEndian, &errorCode)
		binary.Read(buf, binary.LittleEndian, &buffer)
		binary.Read(buf, binary.LittleEndian, &lengthRead)
		fn(fileId, errorCode, buffer, lengthRead)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackAsyncFileRead), wrapper)
}

// Remove a registered Async File Read callback.
func (device *REDBrick) DeregisterAsyncFileReadCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackAsyncFileRead), registrationId)
}


// This callback reports the result of a call to the WriteFileAsync
// function.
func (device *REDBrick) RegisterAsyncFileWriteCallback(fn func(uint16, ErrorCode, uint8)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 12 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var fileId uint16
		var errorCode ErrorCode
		var lengthWritten uint8
		binary.Read(buf, binary.LittleEndian, &fileId)
		binary.Read(buf, binary.LittleEndian, &errorCode)
		binary.Read(buf, binary.LittleEndian, &lengthWritten)
		fn(fileId, errorCode, lengthWritten)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackAsyncFileWrite), wrapper)
}

// Remove a registered Async File Write callback.
func (device *REDBrick) DeregisterAsyncFileWriteCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackAsyncFileWrite), registrationId)
}


// 
func (device *REDBrick) RegisterFileEventsOccurredCallback(fn func(uint16, FileEvent)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 12 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var fileId uint16
		var events FileEvent
		binary.Read(buf, binary.LittleEndian, &fileId)
		binary.Read(buf, binary.LittleEndian, &events)
		fn(fileId, events)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackFileEventsOccurred), wrapper)
}

// Remove a registered File Events Occurred callback.
func (device *REDBrick) DeregisterFileEventsOccurredCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackFileEventsOccurred), registrationId)
}


// 
func (device *REDBrick) RegisterProcessStateChangedCallback(fn func(uint16, ProcessState, uint64, uint8)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 20 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var processId uint16
		var state ProcessState
		var timestamp uint64
		var exitCode uint8
		binary.Read(buf, binary.LittleEndian, &processId)
		binary.Read(buf, binary.LittleEndian, &state)
		binary.Read(buf, binary.LittleEndian, &timestamp)
		binary.Read(buf, binary.LittleEndian, &exitCode)
		fn(processId, state, timestamp, exitCode)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackProcessStateChanged), wrapper)
}

// Remove a registered Process State Changed callback.
func (device *REDBrick) DeregisterProcessStateChangedCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackProcessStateChanged), registrationId)
}


// 
func (device *REDBrick) RegisterProgramSchedulerStateChangedCallback(fn func(uint16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 10 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var programId uint16
		binary.Read(buf, binary.LittleEndian, &programId)
		fn(programId)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackProgramSchedulerStateChanged), wrapper)
}

// Remove a registered Program Scheduler State Changed callback.
func (device *REDBrick) DeregisterProgramSchedulerStateChangedCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackProgramSchedulerStateChanged), registrationId)
}


// 
func (device *REDBrick) RegisterProgramProcessSpawnedCallback(fn func(uint16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 10 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var programId uint16
		binary.Read(buf, binary.LittleEndian, &programId)
		fn(programId)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackProgramProcessSpawned), wrapper)
}

// Remove a registered Program Process Spawned callback.
func (device *REDBrick) DeregisterProgramProcessSpawnedCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackProgramProcessSpawned), registrationId)
}


// 
//
// Associated constants:
//
//	* ErrorCodeSuccess
//	* ErrorCodeUnknownError
//	* ErrorCodeInvalidOperation
//	* ErrorCodeOperationAborted
//	* ErrorCodeInternalError
//	* ErrorCodeUnknownSessionId
//	* ErrorCodeNoFreeSessionId
//	* ErrorCodeUnknownObjectId
//	* ErrorCodeNoFreeObjectId
//	* ErrorCodeObjectIsLocked
//	* ErrorCodeNoMoreData
//	* ErrorCodeWrongListItemType
//	* ErrorCodeProgramIsPurged
//	* ErrorCodeInvalidParameter
//	* ErrorCodeNoFreeMemory
//	* ErrorCodeNoFreeSpace
//	* ErrorCodeAccessDenied
//	* ErrorCodeAlreadyExists
//	* ErrorCodeDoesNotExist
//	* ErrorCodeInterrupted
//	* ErrorCodeIsDirectory
//	* ErrorCodeNotADirectory
//	* ErrorCodeWouldBlock
//	* ErrorCodeOverflow
//	* ErrorCodeBadFileDescriptor
//	* ErrorCodeOutOfRange
//	* ErrorCodeNameTooLong
//	* ErrorCodeInvalidSeek
//	* ErrorCodeNotSupported
//	* ErrorCodeTooManyOpenFiles
func (device *REDBrick) CreateSession(lifetime uint32) (errorCode ErrorCode, sessionId uint16, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, lifetime);

	resultBytes, err := device.device.Get(uint8(FunctionCreateSession), buf.Bytes())
	if err != nil {
		return errorCode, sessionId, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 11 {
		return errorCode, sessionId, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 11)
	}


	if header.ErrorCode != 0 {
		return errorCode, sessionId, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &errorCode)
	binary.Read(resultBuf, binary.LittleEndian, &sessionId)


	return errorCode, sessionId, nil
}

// 
//
// Associated constants:
//
//	* ErrorCodeSuccess
//	* ErrorCodeUnknownError
//	* ErrorCodeInvalidOperation
//	* ErrorCodeOperationAborted
//	* ErrorCodeInternalError
//	* ErrorCodeUnknownSessionId
//	* ErrorCodeNoFreeSessionId
//	* ErrorCodeUnknownObjectId
//	* ErrorCodeNoFreeObjectId
//	* ErrorCodeObjectIsLocked
//	* ErrorCodeNoMoreData
//	* ErrorCodeWrongListItemType
//	* ErrorCodeProgramIsPurged
//	* ErrorCodeInvalidParameter
//	* ErrorCodeNoFreeMemory
//	* ErrorCodeNoFreeSpace
//	* ErrorCodeAccessDenied
//	* ErrorCodeAlreadyExists
//	* ErrorCodeDoesNotExist
//	* ErrorCodeInterrupted
//	* ErrorCodeIsDirectory
//	* ErrorCodeNotADirectory
//	* ErrorCodeWouldBlock
//	* ErrorCodeOverflow
//	* ErrorCodeBadFileDescriptor
//	* ErrorCodeOutOfRange
//	* ErrorCodeNameTooLong
//	* ErrorCodeInvalidSeek
//	* ErrorCodeNotSupported
//	* ErrorCodeTooManyOpenFiles
func (device *REDBrick) ExpireSession(sessionId uint16) (errorCode ErrorCode, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, sessionId);

	resultBytes, err := device.device.Get(uint8(FunctionExpireSession), buf.Bytes())
	if err != nil {
		return errorCode, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 9 {
		return errorCode, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
	}


	if header.ErrorCode != 0 {
		return errorCode, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &errorCode)


	return errorCode, nil
}

// 
func (device *REDBrick) ExpireSessionUnchecked(sessionId uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, sessionId);

	resultBytes, err := device.device.Set(uint8(FunctionExpireSessionUnchecked), buf.Bytes())
	if err != nil {
		return err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 8 {
		return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
	}


	if header.ErrorCode != 0 {
		return DeviceError(header.ErrorCode)
	}

	bytes.NewBuffer(resultBytes[8:])
	

	return nil
}

// 
//
// Associated constants:
//
//	* ErrorCodeSuccess
//	* ErrorCodeUnknownError
//	* ErrorCodeInvalidOperation
//	* ErrorCodeOperationAborted
//	* ErrorCodeInternalError
//	* ErrorCodeUnknownSessionId
//	* ErrorCodeNoFreeSessionId
//	* ErrorCodeUnknownObjectId
//	* ErrorCodeNoFreeObjectId
//	* ErrorCodeObjectIsLocked
//	* ErrorCodeNoMoreData
//	* ErrorCodeWrongListItemType
//	* ErrorCodeProgramIsPurged
//	* ErrorCodeInvalidParameter
//	* ErrorCodeNoFreeMemory
//	* ErrorCodeNoFreeSpace
//	* ErrorCodeAccessDenied
//	* ErrorCodeAlreadyExists
//	* ErrorCodeDoesNotExist
//	* ErrorCodeInterrupted
//	* ErrorCodeIsDirectory
//	* ErrorCodeNotADirectory
//	* ErrorCodeWouldBlock
//	* ErrorCodeOverflow
//	* ErrorCodeBadFileDescriptor
//	* ErrorCodeOutOfRange
//	* ErrorCodeNameTooLong
//	* ErrorCodeInvalidSeek
//	* ErrorCodeNotSupported
//	* ErrorCodeTooManyOpenFiles
func (device *REDBrick) KeepSessionAlive(sessionId uint16, lifetime uint32) (errorCode ErrorCode, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, sessionId);
	binary.Write(&buf, binary.LittleEndian, lifetime);

	resultBytes, err := device.device.Get(uint8(FunctionKeepSessionAlive), buf.Bytes())
	if err != nil {
		return errorCode, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 9 {
		return errorCode, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
	}


	if header.ErrorCode != 0 {
		return errorCode, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &errorCode)


	return errorCode, nil
}

// Decreases the reference count of an object by one and returns the resulting
// error code. If the reference count reaches zero the object gets destroyed.
//
// Associated constants:
//
//	* ErrorCodeSuccess
//	* ErrorCodeUnknownError
//	* ErrorCodeInvalidOperation
//	* ErrorCodeOperationAborted
//	* ErrorCodeInternalError
//	* ErrorCodeUnknownSessionId
//	* ErrorCodeNoFreeSessionId
//	* ErrorCodeUnknownObjectId
//	* ErrorCodeNoFreeObjectId
//	* ErrorCodeObjectIsLocked
//	* ErrorCodeNoMoreData
//	* ErrorCodeWrongListItemType
//	* ErrorCodeProgramIsPurged
//	* ErrorCodeInvalidParameter
//	* ErrorCodeNoFreeMemory
//	* ErrorCodeNoFreeSpace
//	* ErrorCodeAccessDenied
//	* ErrorCodeAlreadyExists
//	* ErrorCodeDoesNotExist
//	* ErrorCodeInterrupted
//	* ErrorCodeIsDirectory
//	* ErrorCodeNotADirectory
//	* ErrorCodeWouldBlock
//	* ErrorCodeOverflow
//	* ErrorCodeBadFileDescriptor
//	* ErrorCodeOutOfRange
//	* ErrorCodeNameTooLong
//	* ErrorCodeInvalidSeek
//	* ErrorCodeNotSupported
//	* ErrorCodeTooManyOpenFiles
func (device *REDBrick) ReleaseObject(objectId uint16, sessionId uint16) (errorCode ErrorCode, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, objectId);
	binary.Write(&buf, binary.LittleEndian, sessionId);

	resultBytes, err := device.device.Get(uint8(FunctionReleaseObject), buf.Bytes())
	if err != nil {
		return errorCode, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 9 {
		return errorCode, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
	}


	if header.ErrorCode != 0 {
		return errorCode, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &errorCode)


	return errorCode, nil
}

// 
func (device *REDBrick) ReleaseObjectUnchecked(objectId uint16, sessionId uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, objectId);
	binary.Write(&buf, binary.LittleEndian, sessionId);

	resultBytes, err := device.device.Set(uint8(FunctionReleaseObjectUnchecked), buf.Bytes())
	if err != nil {
		return err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 8 {
		return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
	}


	if header.ErrorCode != 0 {
		return DeviceError(header.ErrorCode)
	}

	bytes.NewBuffer(resultBytes[8:])
	

	return nil
}

// Allocates a new string object, reserves ``length_to_reserve`` bytes memory
// for it and sets up to the first 60 bytes. Set ``length_to_reserve`` to the
// length of the string that should be stored in the string object.
// 
// Returns the object ID of the new string object and the resulting error code.
//
// Associated constants:
//
//	* ErrorCodeSuccess
//	* ErrorCodeUnknownError
//	* ErrorCodeInvalidOperation
//	* ErrorCodeOperationAborted
//	* ErrorCodeInternalError
//	* ErrorCodeUnknownSessionId
//	* ErrorCodeNoFreeSessionId
//	* ErrorCodeUnknownObjectId
//	* ErrorCodeNoFreeObjectId
//	* ErrorCodeObjectIsLocked
//	* ErrorCodeNoMoreData
//	* ErrorCodeWrongListItemType
//	* ErrorCodeProgramIsPurged
//	* ErrorCodeInvalidParameter
//	* ErrorCodeNoFreeMemory
//	* ErrorCodeNoFreeSpace
//	* ErrorCodeAccessDenied
//	* ErrorCodeAlreadyExists
//	* ErrorCodeDoesNotExist
//	* ErrorCodeInterrupted
//	* ErrorCodeIsDirectory
//	* ErrorCodeNotADirectory
//	* ErrorCodeWouldBlock
//	* ErrorCodeOverflow
//	* ErrorCodeBadFileDescriptor
//	* ErrorCodeOutOfRange
//	* ErrorCodeNameTooLong
//	* ErrorCodeInvalidSeek
//	* ErrorCodeNotSupported
//	* ErrorCodeTooManyOpenFiles
func (device *REDBrick) AllocateString(lengthToReserve uint32, buffer string, sessionId uint16) (errorCode ErrorCode, stringId uint16, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, lengthToReserve);
	buffer_byte_slice, err := StringToByteSlice(buffer, 58)
	if err != nil { return }
	buf.Write(buffer_byte_slice)
	binary.Write(&buf, binary.LittleEndian, sessionId);

	resultBytes, err := device.device.Get(uint8(FunctionAllocateString), buf.Bytes())
	if err != nil {
		return errorCode, stringId, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 11 {
		return errorCode, stringId, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 11)
	}


	if header.ErrorCode != 0 {
		return errorCode, stringId, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &errorCode)
	binary.Read(resultBuf, binary.LittleEndian, &stringId)


	return errorCode, stringId, nil
}

// Truncates a string object to ``length`` bytes and returns the resulting
// error code.
//
// Associated constants:
//
//	* ErrorCodeSuccess
//	* ErrorCodeUnknownError
//	* ErrorCodeInvalidOperation
//	* ErrorCodeOperationAborted
//	* ErrorCodeInternalError
//	* ErrorCodeUnknownSessionId
//	* ErrorCodeNoFreeSessionId
//	* ErrorCodeUnknownObjectId
//	* ErrorCodeNoFreeObjectId
//	* ErrorCodeObjectIsLocked
//	* ErrorCodeNoMoreData
//	* ErrorCodeWrongListItemType
//	* ErrorCodeProgramIsPurged
//	* ErrorCodeInvalidParameter
//	* ErrorCodeNoFreeMemory
//	* ErrorCodeNoFreeSpace
//	* ErrorCodeAccessDenied
//	* ErrorCodeAlreadyExists
//	* ErrorCodeDoesNotExist
//	* ErrorCodeInterrupted
//	* ErrorCodeIsDirectory
//	* ErrorCodeNotADirectory
//	* ErrorCodeWouldBlock
//	* ErrorCodeOverflow
//	* ErrorCodeBadFileDescriptor
//	* ErrorCodeOutOfRange
//	* ErrorCodeNameTooLong
//	* ErrorCodeInvalidSeek
//	* ErrorCodeNotSupported
//	* ErrorCodeTooManyOpenFiles
func (device *REDBrick) TruncateString(stringId uint16, length uint32) (errorCode ErrorCode, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, stringId);
	binary.Write(&buf, binary.LittleEndian, length);

	resultBytes, err := device.device.Get(uint8(FunctionTruncateString), buf.Bytes())
	if err != nil {
		return errorCode, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 9 {
		return errorCode, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
	}


	if header.ErrorCode != 0 {
		return errorCode, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &errorCode)


	return errorCode, nil
}

// Returns the length of a string object and the resulting error code.
//
// Associated constants:
//
//	* ErrorCodeSuccess
//	* ErrorCodeUnknownError
//	* ErrorCodeInvalidOperation
//	* ErrorCodeOperationAborted
//	* ErrorCodeInternalError
//	* ErrorCodeUnknownSessionId
//	* ErrorCodeNoFreeSessionId
//	* ErrorCodeUnknownObjectId
//	* ErrorCodeNoFreeObjectId
//	* ErrorCodeObjectIsLocked
//	* ErrorCodeNoMoreData
//	* ErrorCodeWrongListItemType
//	* ErrorCodeProgramIsPurged
//	* ErrorCodeInvalidParameter
//	* ErrorCodeNoFreeMemory
//	* ErrorCodeNoFreeSpace
//	* ErrorCodeAccessDenied
//	* ErrorCodeAlreadyExists
//	* ErrorCodeDoesNotExist
//	* ErrorCodeInterrupted
//	* ErrorCodeIsDirectory
//	* ErrorCodeNotADirectory
//	* ErrorCodeWouldBlock
//	* ErrorCodeOverflow
//	* ErrorCodeBadFileDescriptor
//	* ErrorCodeOutOfRange
//	* ErrorCodeNameTooLong
//	* ErrorCodeInvalidSeek
//	* ErrorCodeNotSupported
//	* ErrorCodeTooManyOpenFiles
func (device *REDBrick) GetStringLength(stringId uint16) (errorCode ErrorCode, length uint32, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, stringId);

	resultBytes, err := device.device.Get(uint8(FunctionGetStringLength), buf.Bytes())
	if err != nil {
		return errorCode, length, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 13 {
		return errorCode, length, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 13)
	}


	if header.ErrorCode != 0 {
		return errorCode, length, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &errorCode)
	binary.Read(resultBuf, binary.LittleEndian, &length)


	return errorCode, length, nil
}

// Sets a chunk of up to 58 bytes in a string object beginning at ``offset``.
// 
// Returns the resulting error code.
//
// Associated constants:
//
//	* ErrorCodeSuccess
//	* ErrorCodeUnknownError
//	* ErrorCodeInvalidOperation
//	* ErrorCodeOperationAborted
//	* ErrorCodeInternalError
//	* ErrorCodeUnknownSessionId
//	* ErrorCodeNoFreeSessionId
//	* ErrorCodeUnknownObjectId
//	* ErrorCodeNoFreeObjectId
//	* ErrorCodeObjectIsLocked
//	* ErrorCodeNoMoreData
//	* ErrorCodeWrongListItemType
//	* ErrorCodeProgramIsPurged
//	* ErrorCodeInvalidParameter
//	* ErrorCodeNoFreeMemory
//	* ErrorCodeNoFreeSpace
//	* ErrorCodeAccessDenied
//	* ErrorCodeAlreadyExists
//	* ErrorCodeDoesNotExist
//	* ErrorCodeInterrupted
//	* ErrorCodeIsDirectory
//	* ErrorCodeNotADirectory
//	* ErrorCodeWouldBlock
//	* ErrorCodeOverflow
//	* ErrorCodeBadFileDescriptor
//	* ErrorCodeOutOfRange
//	* ErrorCodeNameTooLong
//	* ErrorCodeInvalidSeek
//	* ErrorCodeNotSupported
//	* ErrorCodeTooManyOpenFiles
func (device *REDBrick) SetStringChunk(stringId uint16, offset uint32, buffer string) (errorCode ErrorCode, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, stringId);
	binary.Write(&buf, binary.LittleEndian, offset);
	buffer_byte_slice, err := StringToByteSlice(buffer, 58)
	if err != nil { return }
	buf.Write(buffer_byte_slice)

	resultBytes, err := device.device.Get(uint8(FunctionSetStringChunk), buf.Bytes())
	if err != nil {
		return errorCode, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 9 {
		return errorCode, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
	}


	if header.ErrorCode != 0 {
		return errorCode, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &errorCode)


	return errorCode, nil
}

// Returns a chunk up to 63 bytes from a string object beginning at ``offset`` and
// returns the resulting error code.
//
// Associated constants:
//
//	* ErrorCodeSuccess
//	* ErrorCodeUnknownError
//	* ErrorCodeInvalidOperation
//	* ErrorCodeOperationAborted
//	* ErrorCodeInternalError
//	* ErrorCodeUnknownSessionId
//	* ErrorCodeNoFreeSessionId
//	* ErrorCodeUnknownObjectId
//	* ErrorCodeNoFreeObjectId
//	* ErrorCodeObjectIsLocked
//	* ErrorCodeNoMoreData
//	* ErrorCodeWrongListItemType
//	* ErrorCodeProgramIsPurged
//	* ErrorCodeInvalidParameter
//	* ErrorCodeNoFreeMemory
//	* ErrorCodeNoFreeSpace
//	* ErrorCodeAccessDenied
//	* ErrorCodeAlreadyExists
//	* ErrorCodeDoesNotExist
//	* ErrorCodeInterrupted
//	* ErrorCodeIsDirectory
//	* ErrorCodeNotADirectory
//	* ErrorCodeWouldBlock
//	* ErrorCodeOverflow
//	* ErrorCodeBadFileDescriptor
//	* ErrorCodeOutOfRange
//	* ErrorCodeNameTooLong
//	* ErrorCodeInvalidSeek
//	* ErrorCodeNotSupported
//	* ErrorCodeTooManyOpenFiles
func (device *REDBrick) GetStringChunk(stringId uint16, offset uint32) (errorCode ErrorCode, buffer string, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, stringId);
	binary.Write(&buf, binary.LittleEndian, offset);

	resultBytes, err := device.device.Get(uint8(FunctionGetStringChunk), buf.Bytes())
	if err != nil {
		return errorCode, buffer, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 72 {
		return errorCode, buffer, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 72)
	}


	if header.ErrorCode != 0 {
		return errorCode, buffer, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &errorCode)
	buffer = ByteSliceToString(resultBuf.Next(63))


	return errorCode, buffer, nil
}

// Allocates a new list object and reserves memory for ``length_to_reserve``
// items. Set ``length_to_reserve`` to the number of items that should be stored
// in the list object.
// 
// Returns the object ID of the new list object and the resulting error code.
// 
// When a list object gets destroyed then the reference count of each object in
// the list object is decreased by one.
//
// Associated constants:
//
//	* ErrorCodeSuccess
//	* ErrorCodeUnknownError
//	* ErrorCodeInvalidOperation
//	* ErrorCodeOperationAborted
//	* ErrorCodeInternalError
//	* ErrorCodeUnknownSessionId
//	* ErrorCodeNoFreeSessionId
//	* ErrorCodeUnknownObjectId
//	* ErrorCodeNoFreeObjectId
//	* ErrorCodeObjectIsLocked
//	* ErrorCodeNoMoreData
//	* ErrorCodeWrongListItemType
//	* ErrorCodeProgramIsPurged
//	* ErrorCodeInvalidParameter
//	* ErrorCodeNoFreeMemory
//	* ErrorCodeNoFreeSpace
//	* ErrorCodeAccessDenied
//	* ErrorCodeAlreadyExists
//	* ErrorCodeDoesNotExist
//	* ErrorCodeInterrupted
//	* ErrorCodeIsDirectory
//	* ErrorCodeNotADirectory
//	* ErrorCodeWouldBlock
//	* ErrorCodeOverflow
//	* ErrorCodeBadFileDescriptor
//	* ErrorCodeOutOfRange
//	* ErrorCodeNameTooLong
//	* ErrorCodeInvalidSeek
//	* ErrorCodeNotSupported
//	* ErrorCodeTooManyOpenFiles
func (device *REDBrick) AllocateList(lengthToReserve uint16, sessionId uint16) (errorCode ErrorCode, listId uint16, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, lengthToReserve);
	binary.Write(&buf, binary.LittleEndian, sessionId);

	resultBytes, err := device.device.Get(uint8(FunctionAllocateList), buf.Bytes())
	if err != nil {
		return errorCode, listId, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 11 {
		return errorCode, listId, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 11)
	}


	if header.ErrorCode != 0 {
		return errorCode, listId, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &errorCode)
	binary.Read(resultBuf, binary.LittleEndian, &listId)


	return errorCode, listId, nil
}

// Returns the length of a list object in items and the resulting error code.
//
// Associated constants:
//
//	* ErrorCodeSuccess
//	* ErrorCodeUnknownError
//	* ErrorCodeInvalidOperation
//	* ErrorCodeOperationAborted
//	* ErrorCodeInternalError
//	* ErrorCodeUnknownSessionId
//	* ErrorCodeNoFreeSessionId
//	* ErrorCodeUnknownObjectId
//	* ErrorCodeNoFreeObjectId
//	* ErrorCodeObjectIsLocked
//	* ErrorCodeNoMoreData
//	* ErrorCodeWrongListItemType
//	* ErrorCodeProgramIsPurged
//	* ErrorCodeInvalidParameter
//	* ErrorCodeNoFreeMemory
//	* ErrorCodeNoFreeSpace
//	* ErrorCodeAccessDenied
//	* ErrorCodeAlreadyExists
//	* ErrorCodeDoesNotExist
//	* ErrorCodeInterrupted
//	* ErrorCodeIsDirectory
//	* ErrorCodeNotADirectory
//	* ErrorCodeWouldBlock
//	* ErrorCodeOverflow
//	* ErrorCodeBadFileDescriptor
//	* ErrorCodeOutOfRange
//	* ErrorCodeNameTooLong
//	* ErrorCodeInvalidSeek
//	* ErrorCodeNotSupported
//	* ErrorCodeTooManyOpenFiles
func (device *REDBrick) GetListLength(listId uint16) (errorCode ErrorCode, length uint16, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, listId);

	resultBytes, err := device.device.Get(uint8(FunctionGetListLength), buf.Bytes())
	if err != nil {
		return errorCode, length, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 11 {
		return errorCode, length, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 11)
	}


	if header.ErrorCode != 0 {
		return errorCode, length, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &errorCode)
	binary.Read(resultBuf, binary.LittleEndian, &length)


	return errorCode, length, nil
}

// Returns the object ID and type of the object stored at ``index`` in a list
// object and returns the resulting error code.
// 
// Possible object types are:
// 
// * String = 0
// * List = 1
// * File = 2
// * Directory = 3
// * Process = 4
// * Program = 5
//
// Associated constants:
//
//	* ErrorCodeSuccess
//	* ErrorCodeUnknownError
//	* ErrorCodeInvalidOperation
//	* ErrorCodeOperationAborted
//	* ErrorCodeInternalError
//	* ErrorCodeUnknownSessionId
//	* ErrorCodeNoFreeSessionId
//	* ErrorCodeUnknownObjectId
//	* ErrorCodeNoFreeObjectId
//	* ErrorCodeObjectIsLocked
//	* ErrorCodeNoMoreData
//	* ErrorCodeWrongListItemType
//	* ErrorCodeProgramIsPurged
//	* ErrorCodeInvalidParameter
//	* ErrorCodeNoFreeMemory
//	* ErrorCodeNoFreeSpace
//	* ErrorCodeAccessDenied
//	* ErrorCodeAlreadyExists
//	* ErrorCodeDoesNotExist
//	* ErrorCodeInterrupted
//	* ErrorCodeIsDirectory
//	* ErrorCodeNotADirectory
//	* ErrorCodeWouldBlock
//	* ErrorCodeOverflow
//	* ErrorCodeBadFileDescriptor
//	* ErrorCodeOutOfRange
//	* ErrorCodeNameTooLong
//	* ErrorCodeInvalidSeek
//	* ErrorCodeNotSupported
//	* ErrorCodeTooManyOpenFiles
//	* ObjectTypeString
//	* ObjectTypeList
//	* ObjectTypeFile
//	* ObjectTypeDirectory
//	* ObjectTypeProcess
//	* ObjectTypeProgram
func (device *REDBrick) GetListItem(listId uint16, index uint16, sessionId uint16) (errorCode ErrorCode, itemObjectId uint16, type_ ObjectType, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, listId);
	binary.Write(&buf, binary.LittleEndian, index);
	binary.Write(&buf, binary.LittleEndian, sessionId);

	resultBytes, err := device.device.Get(uint8(FunctionGetListItem), buf.Bytes())
	if err != nil {
		return errorCode, itemObjectId, type_, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 12 {
		return errorCode, itemObjectId, type_, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
	}


	if header.ErrorCode != 0 {
		return errorCode, itemObjectId, type_, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &errorCode)
	binary.Read(resultBuf, binary.LittleEndian, &itemObjectId)
	binary.Read(resultBuf, binary.LittleEndian, &type_)


	return errorCode, itemObjectId, type_, nil
}

// Appends an object to a list object and increases the reference count of the
// appended object by one.
// 
// Returns the resulting error code.
//
// Associated constants:
//
//	* ErrorCodeSuccess
//	* ErrorCodeUnknownError
//	* ErrorCodeInvalidOperation
//	* ErrorCodeOperationAborted
//	* ErrorCodeInternalError
//	* ErrorCodeUnknownSessionId
//	* ErrorCodeNoFreeSessionId
//	* ErrorCodeUnknownObjectId
//	* ErrorCodeNoFreeObjectId
//	* ErrorCodeObjectIsLocked
//	* ErrorCodeNoMoreData
//	* ErrorCodeWrongListItemType
//	* ErrorCodeProgramIsPurged
//	* ErrorCodeInvalidParameter
//	* ErrorCodeNoFreeMemory
//	* ErrorCodeNoFreeSpace
//	* ErrorCodeAccessDenied
//	* ErrorCodeAlreadyExists
//	* ErrorCodeDoesNotExist
//	* ErrorCodeInterrupted
//	* ErrorCodeIsDirectory
//	* ErrorCodeNotADirectory
//	* ErrorCodeWouldBlock
//	* ErrorCodeOverflow
//	* ErrorCodeBadFileDescriptor
//	* ErrorCodeOutOfRange
//	* ErrorCodeNameTooLong
//	* ErrorCodeInvalidSeek
//	* ErrorCodeNotSupported
//	* ErrorCodeTooManyOpenFiles
func (device *REDBrick) AppendToList(listId uint16, itemObjectId uint16) (errorCode ErrorCode, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, listId);
	binary.Write(&buf, binary.LittleEndian, itemObjectId);

	resultBytes, err := device.device.Get(uint8(FunctionAppendToList), buf.Bytes())
	if err != nil {
		return errorCode, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 9 {
		return errorCode, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
	}


	if header.ErrorCode != 0 {
		return errorCode, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &errorCode)


	return errorCode, nil
}

// Removes the object stored at ``index`` from a list object and decreases the
// reference count of the removed object by one.
// 
// Returns the resulting error code.
//
// Associated constants:
//
//	* ErrorCodeSuccess
//	* ErrorCodeUnknownError
//	* ErrorCodeInvalidOperation
//	* ErrorCodeOperationAborted
//	* ErrorCodeInternalError
//	* ErrorCodeUnknownSessionId
//	* ErrorCodeNoFreeSessionId
//	* ErrorCodeUnknownObjectId
//	* ErrorCodeNoFreeObjectId
//	* ErrorCodeObjectIsLocked
//	* ErrorCodeNoMoreData
//	* ErrorCodeWrongListItemType
//	* ErrorCodeProgramIsPurged
//	* ErrorCodeInvalidParameter
//	* ErrorCodeNoFreeMemory
//	* ErrorCodeNoFreeSpace
//	* ErrorCodeAccessDenied
//	* ErrorCodeAlreadyExists
//	* ErrorCodeDoesNotExist
//	* ErrorCodeInterrupted
//	* ErrorCodeIsDirectory
//	* ErrorCodeNotADirectory
//	* ErrorCodeWouldBlock
//	* ErrorCodeOverflow
//	* ErrorCodeBadFileDescriptor
//	* ErrorCodeOutOfRange
//	* ErrorCodeNameTooLong
//	* ErrorCodeInvalidSeek
//	* ErrorCodeNotSupported
//	* ErrorCodeTooManyOpenFiles
func (device *REDBrick) RemoveFromList(listId uint16, index uint16) (errorCode ErrorCode, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, listId);
	binary.Write(&buf, binary.LittleEndian, index);

	resultBytes, err := device.device.Get(uint8(FunctionRemoveFromList), buf.Bytes())
	if err != nil {
		return errorCode, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 9 {
		return errorCode, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
	}


	if header.ErrorCode != 0 {
		return errorCode, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &errorCode)


	return errorCode, nil
}

// Opens an existing file or creates a new file and allocates a new file object
// for it.
// 
// FIXME: name has to be absolute
// 
// The reference count of the name string object is increased by one. When the
// file object gets destroyed then the reference count of the name string object is
// decreased by one. Also the name string object is locked and cannot be modified
// while the file object holds a reference to it.
// 
// The ``flags`` parameter takes a ORed combination of the following possible file
// flags (in hexadecimal notation):
// 
// * ReadOnly = 0x0001 (O_RDONLY)
// * WriteOnly = 0x0002 (O_WRONLY)
// * ReadWrite = 0x0004 (O_RDWR)
// * Append = 0x0008 (O_APPEND)
// * Create = 0x0010 (O_CREAT)
// * Exclusive = 0x0020 (O_EXCL)
// * NonBlocking = 0x0040 (O_NONBLOCK)
// * Truncate = 0x0080 (O_TRUNC)
// * Temporary = 0x0100
// * Replace = 0x0200
// 
// FIXME: explain *Temporary* and *Replace* flag
// 
// The ``permissions`` parameter takes a ORed combination of the following
// possible file permissions (in octal notation) that match the common UNIX
// permission bits:
// 
// * UserRead = 00400
// * UserWrite = 00200
// * UserExecute = 00100
// * GroupRead = 00040
// * GroupWrite = 00020
// * GroupExecute = 00010
// * OthersRead = 00004
// * OthersWrite = 00002
// * OthersExecute = 00001
// 
// Returns the object ID of the new file object and the resulting error code.
//
// Associated constants:
//
//	* FileFlagReadOnly
//	* FileFlagWriteOnly
//	* FileFlagReadWrite
//	* FileFlagAppend
//	* FileFlagCreate
//	* FileFlagExclusive
//	* FileFlagNonBlocking
//	* FileFlagTruncate
//	* FileFlagTemporary
//	* FileFlagReplace
//	* FilePermissionUserAll
//	* FilePermissionUserRead
//	* FilePermissionUserWrite
//	* FilePermissionUserExecute
//	* FilePermissionGroupAll
//	* FilePermissionGroupRead
//	* FilePermissionGroupWrite
//	* FilePermissionGroupExecute
//	* FilePermissionOthersAll
//	* FilePermissionOthersRead
//	* FilePermissionOthersWrite
//	* FilePermissionOthersExecute
//	* ErrorCodeSuccess
//	* ErrorCodeUnknownError
//	* ErrorCodeInvalidOperation
//	* ErrorCodeOperationAborted
//	* ErrorCodeInternalError
//	* ErrorCodeUnknownSessionId
//	* ErrorCodeNoFreeSessionId
//	* ErrorCodeUnknownObjectId
//	* ErrorCodeNoFreeObjectId
//	* ErrorCodeObjectIsLocked
//	* ErrorCodeNoMoreData
//	* ErrorCodeWrongListItemType
//	* ErrorCodeProgramIsPurged
//	* ErrorCodeInvalidParameter
//	* ErrorCodeNoFreeMemory
//	* ErrorCodeNoFreeSpace
//	* ErrorCodeAccessDenied
//	* ErrorCodeAlreadyExists
//	* ErrorCodeDoesNotExist
//	* ErrorCodeInterrupted
//	* ErrorCodeIsDirectory
//	* ErrorCodeNotADirectory
//	* ErrorCodeWouldBlock
//	* ErrorCodeOverflow
//	* ErrorCodeBadFileDescriptor
//	* ErrorCodeOutOfRange
//	* ErrorCodeNameTooLong
//	* ErrorCodeInvalidSeek
//	* ErrorCodeNotSupported
//	* ErrorCodeTooManyOpenFiles
func (device *REDBrick) OpenFile(nameStringId uint16, flags FileFlag, permissions FilePermission, uid uint32, gid uint32, sessionId uint16) (errorCode ErrorCode, fileId uint16, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, nameStringId);
	binary.Write(&buf, binary.LittleEndian, flags);
	binary.Write(&buf, binary.LittleEndian, permissions);
	binary.Write(&buf, binary.LittleEndian, uid);
	binary.Write(&buf, binary.LittleEndian, gid);
	binary.Write(&buf, binary.LittleEndian, sessionId);

	resultBytes, err := device.device.Get(uint8(FunctionOpenFile), buf.Bytes())
	if err != nil {
		return errorCode, fileId, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 11 {
		return errorCode, fileId, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 11)
	}


	if header.ErrorCode != 0 {
		return errorCode, fileId, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &errorCode)
	binary.Read(resultBuf, binary.LittleEndian, &fileId)


	return errorCode, fileId, nil
}

// Creates a new pipe and allocates a new file object for it.
// 
// The ``flags`` parameter takes a ORed combination of the following possible
// pipe flags (in hexadecimal notation):
// 
// * NonBlockingRead = 0x0001
// * NonBlockingWrite = 0x0002
// 
// The length of the pipe buffer can be specified with the ``length`` parameter
// in bytes. If length is set to zero, then the default pipe buffer length is used.
// 
// Returns the object ID of the new file object and the resulting error code.
//
// Associated constants:
//
//	* PipeFlagNonBlockingRead
//	* PipeFlagNonBlockingWrite
//	* ErrorCodeSuccess
//	* ErrorCodeUnknownError
//	* ErrorCodeInvalidOperation
//	* ErrorCodeOperationAborted
//	* ErrorCodeInternalError
//	* ErrorCodeUnknownSessionId
//	* ErrorCodeNoFreeSessionId
//	* ErrorCodeUnknownObjectId
//	* ErrorCodeNoFreeObjectId
//	* ErrorCodeObjectIsLocked
//	* ErrorCodeNoMoreData
//	* ErrorCodeWrongListItemType
//	* ErrorCodeProgramIsPurged
//	* ErrorCodeInvalidParameter
//	* ErrorCodeNoFreeMemory
//	* ErrorCodeNoFreeSpace
//	* ErrorCodeAccessDenied
//	* ErrorCodeAlreadyExists
//	* ErrorCodeDoesNotExist
//	* ErrorCodeInterrupted
//	* ErrorCodeIsDirectory
//	* ErrorCodeNotADirectory
//	* ErrorCodeWouldBlock
//	* ErrorCodeOverflow
//	* ErrorCodeBadFileDescriptor
//	* ErrorCodeOutOfRange
//	* ErrorCodeNameTooLong
//	* ErrorCodeInvalidSeek
//	* ErrorCodeNotSupported
//	* ErrorCodeTooManyOpenFiles
func (device *REDBrick) CreatePipe(flags PipeFlag, length uint64, sessionId uint16) (errorCode ErrorCode, fileId uint16, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, flags);
	binary.Write(&buf, binary.LittleEndian, length);
	binary.Write(&buf, binary.LittleEndian, sessionId);

	resultBytes, err := device.device.Get(uint8(FunctionCreatePipe), buf.Bytes())
	if err != nil {
		return errorCode, fileId, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 11 {
		return errorCode, fileId, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 11)
	}


	if header.ErrorCode != 0 {
		return errorCode, fileId, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &errorCode)
	binary.Read(resultBuf, binary.LittleEndian, &fileId)


	return errorCode, fileId, nil
}

// Returns various information about a file and the resulting error code.
// 
// Possible file types are:
// 
// * Unknown = 0
// * Regular = 1
// * Directory = 2
// * Character = 3
// * Block = 4
// * FIFO = 5
// * Symlink = 6
// * Socket = 7
// * Pipe = 8
// 
// If the file type is *Pipe* then the returned name string object is invalid,
// because a pipe has no name. Otherwise the returned name string object was used
// to open or create the file object, as passed to OpenFile.
// 
// The returned flags were used to open or create the file object, as passed to
// OpenFile or CreatePipe. See the respective function for a list
// of possible file and pipe flags.
// 
// FIXME: everything except flags and length is invalid if file type is *Pipe*
//
// Associated constants:
//
//	* ErrorCodeSuccess
//	* ErrorCodeUnknownError
//	* ErrorCodeInvalidOperation
//	* ErrorCodeOperationAborted
//	* ErrorCodeInternalError
//	* ErrorCodeUnknownSessionId
//	* ErrorCodeNoFreeSessionId
//	* ErrorCodeUnknownObjectId
//	* ErrorCodeNoFreeObjectId
//	* ErrorCodeObjectIsLocked
//	* ErrorCodeNoMoreData
//	* ErrorCodeWrongListItemType
//	* ErrorCodeProgramIsPurged
//	* ErrorCodeInvalidParameter
//	* ErrorCodeNoFreeMemory
//	* ErrorCodeNoFreeSpace
//	* ErrorCodeAccessDenied
//	* ErrorCodeAlreadyExists
//	* ErrorCodeDoesNotExist
//	* ErrorCodeInterrupted
//	* ErrorCodeIsDirectory
//	* ErrorCodeNotADirectory
//	* ErrorCodeWouldBlock
//	* ErrorCodeOverflow
//	* ErrorCodeBadFileDescriptor
//	* ErrorCodeOutOfRange
//	* ErrorCodeNameTooLong
//	* ErrorCodeInvalidSeek
//	* ErrorCodeNotSupported
//	* ErrorCodeTooManyOpenFiles
//	* FileTypeUnknown
//	* FileTypeRegular
//	* FileTypeDirectory
//	* FileTypeCharacter
//	* FileTypeBlock
//	* FileTypeFIFO
//	* FileTypeSymlink
//	* FileTypeSocket
//	* FileTypePipe
//	* FileFlagReadOnly
//	* FileFlagWriteOnly
//	* FileFlagReadWrite
//	* FileFlagAppend
//	* FileFlagCreate
//	* FileFlagExclusive
//	* FileFlagNonBlocking
//	* FileFlagTruncate
//	* FileFlagTemporary
//	* FileFlagReplace
//	* FilePermissionUserAll
//	* FilePermissionUserRead
//	* FilePermissionUserWrite
//	* FilePermissionUserExecute
//	* FilePermissionGroupAll
//	* FilePermissionGroupRead
//	* FilePermissionGroupWrite
//	* FilePermissionGroupExecute
//	* FilePermissionOthersAll
//	* FilePermissionOthersRead
//	* FilePermissionOthersWrite
//	* FilePermissionOthersExecute
func (device *REDBrick) GetFileInfo(fileId uint16, sessionId uint16) (errorCode ErrorCode, type_ FileType, nameStringId uint16, flags FileFlag, permissions FilePermission, uid uint32, gid uint32, length uint64, accessTimestamp uint64, modificationTimestamp uint64, statusChangeTimestamp uint64, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, fileId);
	binary.Write(&buf, binary.LittleEndian, sessionId);

	resultBytes, err := device.device.Get(uint8(FunctionGetFileInfo), buf.Bytes())
	if err != nil {
		return errorCode, type_, nameStringId, flags, permissions, uid, gid, length, accessTimestamp, modificationTimestamp, statusChangeTimestamp, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 58 {
		return errorCode, type_, nameStringId, flags, permissions, uid, gid, length, accessTimestamp, modificationTimestamp, statusChangeTimestamp, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 58)
	}


	if header.ErrorCode != 0 {
		return errorCode, type_, nameStringId, flags, permissions, uid, gid, length, accessTimestamp, modificationTimestamp, statusChangeTimestamp, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &errorCode)
	binary.Read(resultBuf, binary.LittleEndian, &type_)
	binary.Read(resultBuf, binary.LittleEndian, &nameStringId)
	binary.Read(resultBuf, binary.LittleEndian, &flags)
	binary.Read(resultBuf, binary.LittleEndian, &permissions)
	binary.Read(resultBuf, binary.LittleEndian, &uid)
	binary.Read(resultBuf, binary.LittleEndian, &gid)
	binary.Read(resultBuf, binary.LittleEndian, &length)
	binary.Read(resultBuf, binary.LittleEndian, &accessTimestamp)
	binary.Read(resultBuf, binary.LittleEndian, &modificationTimestamp)
	binary.Read(resultBuf, binary.LittleEndian, &statusChangeTimestamp)


	return errorCode, type_, nameStringId, flags, permissions, uid, gid, length, accessTimestamp, modificationTimestamp, statusChangeTimestamp, nil
}

// Reads up to 62 bytes from a file object.
// 
// Returns the bytes read, the actual number of bytes read and the resulting
// error code.
// 
// If there is not data to be read, either because the file position reached
// end-of-file or because there is not data in the pipe, then zero bytes are
// returned.
// 
// If the file object was created by OpenFile without the *NonBlocking*
// flag or by CreatePipe without the *NonBlockingRead* flag then the
// error code *NotSupported* is returned.
//
// Associated constants:
//
//	* ErrorCodeSuccess
//	* ErrorCodeUnknownError
//	* ErrorCodeInvalidOperation
//	* ErrorCodeOperationAborted
//	* ErrorCodeInternalError
//	* ErrorCodeUnknownSessionId
//	* ErrorCodeNoFreeSessionId
//	* ErrorCodeUnknownObjectId
//	* ErrorCodeNoFreeObjectId
//	* ErrorCodeObjectIsLocked
//	* ErrorCodeNoMoreData
//	* ErrorCodeWrongListItemType
//	* ErrorCodeProgramIsPurged
//	* ErrorCodeInvalidParameter
//	* ErrorCodeNoFreeMemory
//	* ErrorCodeNoFreeSpace
//	* ErrorCodeAccessDenied
//	* ErrorCodeAlreadyExists
//	* ErrorCodeDoesNotExist
//	* ErrorCodeInterrupted
//	* ErrorCodeIsDirectory
//	* ErrorCodeNotADirectory
//	* ErrorCodeWouldBlock
//	* ErrorCodeOverflow
//	* ErrorCodeBadFileDescriptor
//	* ErrorCodeOutOfRange
//	* ErrorCodeNameTooLong
//	* ErrorCodeInvalidSeek
//	* ErrorCodeNotSupported
//	* ErrorCodeTooManyOpenFiles
func (device *REDBrick) ReadFile(fileId uint16, lengthToRead uint8) (errorCode ErrorCode, buffer [62]uint8, lengthRead uint8, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, fileId);
	binary.Write(&buf, binary.LittleEndian, lengthToRead);

	resultBytes, err := device.device.Get(uint8(FunctionReadFile), buf.Bytes())
	if err != nil {
		return errorCode, buffer, lengthRead, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 72 {
		return errorCode, buffer, lengthRead, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 72)
	}


	if header.ErrorCode != 0 {
		return errorCode, buffer, lengthRead, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &errorCode)
	binary.Read(resultBuf, binary.LittleEndian, &buffer)
	binary.Read(resultBuf, binary.LittleEndian, &lengthRead)


	return errorCode, buffer, lengthRead, nil
}

// Reads up to 2\ :sup:`63`\  - 1 bytes from a file object asynchronously.
// 
// Reports the bytes read (in 60 byte chunks), the actual number of bytes read and
// the resulting error code via the RegisterAsyncFileReadCallback callback.
// 
// If there is not data to be read, either because the file position reached
// end-of-file or because there is not data in the pipe, then zero bytes are
// reported.
// 
// If the file object was created by OpenFile without the *NonBlocking*
// flag or by CreatePipe without the *NonBlockingRead* flag then the error
// code *NotSupported* is reported via the RegisterAsyncFileReadCallback callback.
func (device *REDBrick) ReadFileAsync(fileId uint16, lengthToRead uint64) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, fileId);
	binary.Write(&buf, binary.LittleEndian, lengthToRead);

	resultBytes, err := device.device.Set(uint8(FunctionReadFileAsync), buf.Bytes())
	if err != nil {
		return err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 8 {
		return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
	}


	if header.ErrorCode != 0 {
		return DeviceError(header.ErrorCode)
	}

	bytes.NewBuffer(resultBytes[8:])
	

	return nil
}

// Aborts a ReadFileAsync operation in progress.
// 
// Returns the resulting error code.
// 
// On success the RegisterAsyncFileReadCallback callback will report *OperationAborted*.
//
// Associated constants:
//
//	* ErrorCodeSuccess
//	* ErrorCodeUnknownError
//	* ErrorCodeInvalidOperation
//	* ErrorCodeOperationAborted
//	* ErrorCodeInternalError
//	* ErrorCodeUnknownSessionId
//	* ErrorCodeNoFreeSessionId
//	* ErrorCodeUnknownObjectId
//	* ErrorCodeNoFreeObjectId
//	* ErrorCodeObjectIsLocked
//	* ErrorCodeNoMoreData
//	* ErrorCodeWrongListItemType
//	* ErrorCodeProgramIsPurged
//	* ErrorCodeInvalidParameter
//	* ErrorCodeNoFreeMemory
//	* ErrorCodeNoFreeSpace
//	* ErrorCodeAccessDenied
//	* ErrorCodeAlreadyExists
//	* ErrorCodeDoesNotExist
//	* ErrorCodeInterrupted
//	* ErrorCodeIsDirectory
//	* ErrorCodeNotADirectory
//	* ErrorCodeWouldBlock
//	* ErrorCodeOverflow
//	* ErrorCodeBadFileDescriptor
//	* ErrorCodeOutOfRange
//	* ErrorCodeNameTooLong
//	* ErrorCodeInvalidSeek
//	* ErrorCodeNotSupported
//	* ErrorCodeTooManyOpenFiles
func (device *REDBrick) AbortAsyncFileRead(fileId uint16) (errorCode ErrorCode, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, fileId);

	resultBytes, err := device.device.Get(uint8(FunctionAbortAsyncFileRead), buf.Bytes())
	if err != nil {
		return errorCode, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 9 {
		return errorCode, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
	}


	if header.ErrorCode != 0 {
		return errorCode, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &errorCode)


	return errorCode, nil
}

// Writes up to 61 bytes to a file object.
// 
// Returns the actual number of bytes written and the resulting error code.
// 
// If the file object was created by OpenFile without the *NonBlocking*
// flag or by CreatePipe without the *NonBlockingWrite* flag then the
// error code *NotSupported* is returned.
//
// Associated constants:
//
//	* ErrorCodeSuccess
//	* ErrorCodeUnknownError
//	* ErrorCodeInvalidOperation
//	* ErrorCodeOperationAborted
//	* ErrorCodeInternalError
//	* ErrorCodeUnknownSessionId
//	* ErrorCodeNoFreeSessionId
//	* ErrorCodeUnknownObjectId
//	* ErrorCodeNoFreeObjectId
//	* ErrorCodeObjectIsLocked
//	* ErrorCodeNoMoreData
//	* ErrorCodeWrongListItemType
//	* ErrorCodeProgramIsPurged
//	* ErrorCodeInvalidParameter
//	* ErrorCodeNoFreeMemory
//	* ErrorCodeNoFreeSpace
//	* ErrorCodeAccessDenied
//	* ErrorCodeAlreadyExists
//	* ErrorCodeDoesNotExist
//	* ErrorCodeInterrupted
//	* ErrorCodeIsDirectory
//	* ErrorCodeNotADirectory
//	* ErrorCodeWouldBlock
//	* ErrorCodeOverflow
//	* ErrorCodeBadFileDescriptor
//	* ErrorCodeOutOfRange
//	* ErrorCodeNameTooLong
//	* ErrorCodeInvalidSeek
//	* ErrorCodeNotSupported
//	* ErrorCodeTooManyOpenFiles
func (device *REDBrick) WriteFile(fileId uint16, buffer [61]uint8, lengthToWrite uint8) (errorCode ErrorCode, lengthWritten uint8, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, fileId);
	binary.Write(&buf, binary.LittleEndian, buffer);
	binary.Write(&buf, binary.LittleEndian, lengthToWrite);

	resultBytes, err := device.device.Get(uint8(FunctionWriteFile), buf.Bytes())
	if err != nil {
		return errorCode, lengthWritten, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 10 {
		return errorCode, lengthWritten, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
	}


	if header.ErrorCode != 0 {
		return errorCode, lengthWritten, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &errorCode)
	binary.Read(resultBuf, binary.LittleEndian, &lengthWritten)


	return errorCode, lengthWritten, nil
}

// Writes up to 61 bytes to a file object.
// 
// Does neither report the actual number of bytes written nor the resulting error
// code.
// 
// If the file object was created by OpenFile without the *NonBlocking*
// flag or by CreatePipe without the *NonBlockingWrite* flag then the
// write operation will fail silently.
func (device *REDBrick) WriteFileUnchecked(fileId uint16, buffer [61]uint8, lengthToWrite uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, fileId);
	binary.Write(&buf, binary.LittleEndian, buffer);
	binary.Write(&buf, binary.LittleEndian, lengthToWrite);

	resultBytes, err := device.device.Set(uint8(FunctionWriteFileUnchecked), buf.Bytes())
	if err != nil {
		return err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 8 {
		return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
	}


	if header.ErrorCode != 0 {
		return DeviceError(header.ErrorCode)
	}

	bytes.NewBuffer(resultBytes[8:])
	

	return nil
}

// Writes up to 61 bytes to a file object.
// 
// Reports the actual number of bytes written and the resulting error code via the
// RegisterAsyncFileWriteCallback callback.
// 
// If the file object was created by OpenFile without the *NonBlocking*
// flag or by CreatePipe without the *NonBlockingWrite* flag then the
// error code *NotSupported* is reported via the RegisterAsyncFileWriteCallback callback.
func (device *REDBrick) WriteFileAsync(fileId uint16, buffer [61]uint8, lengthToWrite uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, fileId);
	binary.Write(&buf, binary.LittleEndian, buffer);
	binary.Write(&buf, binary.LittleEndian, lengthToWrite);

	resultBytes, err := device.device.Set(uint8(FunctionWriteFileAsync), buf.Bytes())
	if err != nil {
		return err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 8 {
		return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
	}


	if header.ErrorCode != 0 {
		return DeviceError(header.ErrorCode)
	}

	bytes.NewBuffer(resultBytes[8:])
	

	return nil
}

// Set the current seek position of a file object relative to ``origin``.
// 
// Possible file origins are:
// 
// * Beginning = 0
// * Current = 1
// * End = 2
// 
// Returns the resulting absolute seek position and error code.
// 
// If the file object was created by CreatePipe then it has no seek
// position and the error code *InvalidSeek* is returned.
//
// Associated constants:
//
//	* FileOriginBeginning
//	* FileOriginCurrent
//	* FileOriginEnd
//	* ErrorCodeSuccess
//	* ErrorCodeUnknownError
//	* ErrorCodeInvalidOperation
//	* ErrorCodeOperationAborted
//	* ErrorCodeInternalError
//	* ErrorCodeUnknownSessionId
//	* ErrorCodeNoFreeSessionId
//	* ErrorCodeUnknownObjectId
//	* ErrorCodeNoFreeObjectId
//	* ErrorCodeObjectIsLocked
//	* ErrorCodeNoMoreData
//	* ErrorCodeWrongListItemType
//	* ErrorCodeProgramIsPurged
//	* ErrorCodeInvalidParameter
//	* ErrorCodeNoFreeMemory
//	* ErrorCodeNoFreeSpace
//	* ErrorCodeAccessDenied
//	* ErrorCodeAlreadyExists
//	* ErrorCodeDoesNotExist
//	* ErrorCodeInterrupted
//	* ErrorCodeIsDirectory
//	* ErrorCodeNotADirectory
//	* ErrorCodeWouldBlock
//	* ErrorCodeOverflow
//	* ErrorCodeBadFileDescriptor
//	* ErrorCodeOutOfRange
//	* ErrorCodeNameTooLong
//	* ErrorCodeInvalidSeek
//	* ErrorCodeNotSupported
//	* ErrorCodeTooManyOpenFiles
func (device *REDBrick) SetFilePosition(fileId uint16, offset int64, origin FileOrigin) (errorCode ErrorCode, position uint64, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, fileId);
	binary.Write(&buf, binary.LittleEndian, offset);
	binary.Write(&buf, binary.LittleEndian, origin);

	resultBytes, err := device.device.Get(uint8(FunctionSetFilePosition), buf.Bytes())
	if err != nil {
		return errorCode, position, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 17 {
		return errorCode, position, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 17)
	}


	if header.ErrorCode != 0 {
		return errorCode, position, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &errorCode)
	binary.Read(resultBuf, binary.LittleEndian, &position)


	return errorCode, position, nil
}

// Returns the current seek position of a file object and returns the
// resulting error code.
// 
// If the file object was created by CreatePipe then it has no seek
// position and the error code *InvalidSeek* is returned.
//
// Associated constants:
//
//	* ErrorCodeSuccess
//	* ErrorCodeUnknownError
//	* ErrorCodeInvalidOperation
//	* ErrorCodeOperationAborted
//	* ErrorCodeInternalError
//	* ErrorCodeUnknownSessionId
//	* ErrorCodeNoFreeSessionId
//	* ErrorCodeUnknownObjectId
//	* ErrorCodeNoFreeObjectId
//	* ErrorCodeObjectIsLocked
//	* ErrorCodeNoMoreData
//	* ErrorCodeWrongListItemType
//	* ErrorCodeProgramIsPurged
//	* ErrorCodeInvalidParameter
//	* ErrorCodeNoFreeMemory
//	* ErrorCodeNoFreeSpace
//	* ErrorCodeAccessDenied
//	* ErrorCodeAlreadyExists
//	* ErrorCodeDoesNotExist
//	* ErrorCodeInterrupted
//	* ErrorCodeIsDirectory
//	* ErrorCodeNotADirectory
//	* ErrorCodeWouldBlock
//	* ErrorCodeOverflow
//	* ErrorCodeBadFileDescriptor
//	* ErrorCodeOutOfRange
//	* ErrorCodeNameTooLong
//	* ErrorCodeInvalidSeek
//	* ErrorCodeNotSupported
//	* ErrorCodeTooManyOpenFiles
func (device *REDBrick) GetFilePosition(fileId uint16) (errorCode ErrorCode, position uint64, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, fileId);

	resultBytes, err := device.device.Get(uint8(FunctionGetFilePosition), buf.Bytes())
	if err != nil {
		return errorCode, position, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 17 {
		return errorCode, position, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 17)
	}


	if header.ErrorCode != 0 {
		return errorCode, position, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &errorCode)
	binary.Read(resultBuf, binary.LittleEndian, &position)


	return errorCode, position, nil
}

// 
//
// Associated constants:
//
//	* FileEventReadable
//	* FileEventWritable
//	* ErrorCodeSuccess
//	* ErrorCodeUnknownError
//	* ErrorCodeInvalidOperation
//	* ErrorCodeOperationAborted
//	* ErrorCodeInternalError
//	* ErrorCodeUnknownSessionId
//	* ErrorCodeNoFreeSessionId
//	* ErrorCodeUnknownObjectId
//	* ErrorCodeNoFreeObjectId
//	* ErrorCodeObjectIsLocked
//	* ErrorCodeNoMoreData
//	* ErrorCodeWrongListItemType
//	* ErrorCodeProgramIsPurged
//	* ErrorCodeInvalidParameter
//	* ErrorCodeNoFreeMemory
//	* ErrorCodeNoFreeSpace
//	* ErrorCodeAccessDenied
//	* ErrorCodeAlreadyExists
//	* ErrorCodeDoesNotExist
//	* ErrorCodeInterrupted
//	* ErrorCodeIsDirectory
//	* ErrorCodeNotADirectory
//	* ErrorCodeWouldBlock
//	* ErrorCodeOverflow
//	* ErrorCodeBadFileDescriptor
//	* ErrorCodeOutOfRange
//	* ErrorCodeNameTooLong
//	* ErrorCodeInvalidSeek
//	* ErrorCodeNotSupported
//	* ErrorCodeTooManyOpenFiles
func (device *REDBrick) SetFileEvents(fileId uint16, events FileEvent) (errorCode ErrorCode, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, fileId);
	binary.Write(&buf, binary.LittleEndian, events);

	resultBytes, err := device.device.Get(uint8(FunctionSetFileEvents), buf.Bytes())
	if err != nil {
		return errorCode, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 9 {
		return errorCode, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
	}


	if header.ErrorCode != 0 {
		return errorCode, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &errorCode)


	return errorCode, nil
}

// 
//
// Associated constants:
//
//	* ErrorCodeSuccess
//	* ErrorCodeUnknownError
//	* ErrorCodeInvalidOperation
//	* ErrorCodeOperationAborted
//	* ErrorCodeInternalError
//	* ErrorCodeUnknownSessionId
//	* ErrorCodeNoFreeSessionId
//	* ErrorCodeUnknownObjectId
//	* ErrorCodeNoFreeObjectId
//	* ErrorCodeObjectIsLocked
//	* ErrorCodeNoMoreData
//	* ErrorCodeWrongListItemType
//	* ErrorCodeProgramIsPurged
//	* ErrorCodeInvalidParameter
//	* ErrorCodeNoFreeMemory
//	* ErrorCodeNoFreeSpace
//	* ErrorCodeAccessDenied
//	* ErrorCodeAlreadyExists
//	* ErrorCodeDoesNotExist
//	* ErrorCodeInterrupted
//	* ErrorCodeIsDirectory
//	* ErrorCodeNotADirectory
//	* ErrorCodeWouldBlock
//	* ErrorCodeOverflow
//	* ErrorCodeBadFileDescriptor
//	* ErrorCodeOutOfRange
//	* ErrorCodeNameTooLong
//	* ErrorCodeInvalidSeek
//	* ErrorCodeNotSupported
//	* ErrorCodeTooManyOpenFiles
//	* FileEventReadable
//	* FileEventWritable
func (device *REDBrick) GetFileEvents(fileId uint16) (errorCode ErrorCode, events FileEvent, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, fileId);

	resultBytes, err := device.device.Get(uint8(FunctionGetFileEvents), buf.Bytes())
	if err != nil {
		return errorCode, events, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 11 {
		return errorCode, events, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 11)
	}


	if header.ErrorCode != 0 {
		return errorCode, events, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &errorCode)
	binary.Read(resultBuf, binary.LittleEndian, &events)


	return errorCode, events, nil
}

// Opens an existing directory and allocates a new directory object for it.
// 
// FIXME: name has to be absolute
// 
// The reference count of the name string object is increased by one. When the
// directory object is destroyed then the reference count of the name string
// object is decreased by one. Also the name string object is locked and cannot be
// modified while the directory object holds a reference to it.
// 
// Returns the object ID of the new directory object and the resulting error code.
//
// Associated constants:
//
//	* ErrorCodeSuccess
//	* ErrorCodeUnknownError
//	* ErrorCodeInvalidOperation
//	* ErrorCodeOperationAborted
//	* ErrorCodeInternalError
//	* ErrorCodeUnknownSessionId
//	* ErrorCodeNoFreeSessionId
//	* ErrorCodeUnknownObjectId
//	* ErrorCodeNoFreeObjectId
//	* ErrorCodeObjectIsLocked
//	* ErrorCodeNoMoreData
//	* ErrorCodeWrongListItemType
//	* ErrorCodeProgramIsPurged
//	* ErrorCodeInvalidParameter
//	* ErrorCodeNoFreeMemory
//	* ErrorCodeNoFreeSpace
//	* ErrorCodeAccessDenied
//	* ErrorCodeAlreadyExists
//	* ErrorCodeDoesNotExist
//	* ErrorCodeInterrupted
//	* ErrorCodeIsDirectory
//	* ErrorCodeNotADirectory
//	* ErrorCodeWouldBlock
//	* ErrorCodeOverflow
//	* ErrorCodeBadFileDescriptor
//	* ErrorCodeOutOfRange
//	* ErrorCodeNameTooLong
//	* ErrorCodeInvalidSeek
//	* ErrorCodeNotSupported
//	* ErrorCodeTooManyOpenFiles
func (device *REDBrick) OpenDirectory(nameStringId uint16, sessionId uint16) (errorCode ErrorCode, directoryId uint16, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, nameStringId);
	binary.Write(&buf, binary.LittleEndian, sessionId);

	resultBytes, err := device.device.Get(uint8(FunctionOpenDirectory), buf.Bytes())
	if err != nil {
		return errorCode, directoryId, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 11 {
		return errorCode, directoryId, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 11)
	}


	if header.ErrorCode != 0 {
		return errorCode, directoryId, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &errorCode)
	binary.Read(resultBuf, binary.LittleEndian, &directoryId)


	return errorCode, directoryId, nil
}

// Returns the name of a directory object, as passed to OpenDirectory, and
// the resulting error code.
//
// Associated constants:
//
//	* ErrorCodeSuccess
//	* ErrorCodeUnknownError
//	* ErrorCodeInvalidOperation
//	* ErrorCodeOperationAborted
//	* ErrorCodeInternalError
//	* ErrorCodeUnknownSessionId
//	* ErrorCodeNoFreeSessionId
//	* ErrorCodeUnknownObjectId
//	* ErrorCodeNoFreeObjectId
//	* ErrorCodeObjectIsLocked
//	* ErrorCodeNoMoreData
//	* ErrorCodeWrongListItemType
//	* ErrorCodeProgramIsPurged
//	* ErrorCodeInvalidParameter
//	* ErrorCodeNoFreeMemory
//	* ErrorCodeNoFreeSpace
//	* ErrorCodeAccessDenied
//	* ErrorCodeAlreadyExists
//	* ErrorCodeDoesNotExist
//	* ErrorCodeInterrupted
//	* ErrorCodeIsDirectory
//	* ErrorCodeNotADirectory
//	* ErrorCodeWouldBlock
//	* ErrorCodeOverflow
//	* ErrorCodeBadFileDescriptor
//	* ErrorCodeOutOfRange
//	* ErrorCodeNameTooLong
//	* ErrorCodeInvalidSeek
//	* ErrorCodeNotSupported
//	* ErrorCodeTooManyOpenFiles
func (device *REDBrick) GetDirectoryName(directoryId uint16, sessionId uint16) (errorCode ErrorCode, nameStringId uint16, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, directoryId);
	binary.Write(&buf, binary.LittleEndian, sessionId);

	resultBytes, err := device.device.Get(uint8(FunctionGetDirectoryName), buf.Bytes())
	if err != nil {
		return errorCode, nameStringId, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 11 {
		return errorCode, nameStringId, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 11)
	}


	if header.ErrorCode != 0 {
		return errorCode, nameStringId, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &errorCode)
	binary.Read(resultBuf, binary.LittleEndian, &nameStringId)


	return errorCode, nameStringId, nil
}

// Returns the next entry in a directory object and the resulting error code.
// 
// If there is not next entry then error code *NoMoreData* is returned. To rewind
// a directory object call RewindDirectory.
// 
// Possible directory entry types are:
// 
// * Unknown = 0
// * Regular = 1
// * Directory = 2
// * Character = 3
// * Block = 4
// * FIFO = 5
// * Symlink = 6
// * Socket = 7
//
// Associated constants:
//
//	* ErrorCodeSuccess
//	* ErrorCodeUnknownError
//	* ErrorCodeInvalidOperation
//	* ErrorCodeOperationAborted
//	* ErrorCodeInternalError
//	* ErrorCodeUnknownSessionId
//	* ErrorCodeNoFreeSessionId
//	* ErrorCodeUnknownObjectId
//	* ErrorCodeNoFreeObjectId
//	* ErrorCodeObjectIsLocked
//	* ErrorCodeNoMoreData
//	* ErrorCodeWrongListItemType
//	* ErrorCodeProgramIsPurged
//	* ErrorCodeInvalidParameter
//	* ErrorCodeNoFreeMemory
//	* ErrorCodeNoFreeSpace
//	* ErrorCodeAccessDenied
//	* ErrorCodeAlreadyExists
//	* ErrorCodeDoesNotExist
//	* ErrorCodeInterrupted
//	* ErrorCodeIsDirectory
//	* ErrorCodeNotADirectory
//	* ErrorCodeWouldBlock
//	* ErrorCodeOverflow
//	* ErrorCodeBadFileDescriptor
//	* ErrorCodeOutOfRange
//	* ErrorCodeNameTooLong
//	* ErrorCodeInvalidSeek
//	* ErrorCodeNotSupported
//	* ErrorCodeTooManyOpenFiles
//	* DirectoryEntryTypeUnknown
//	* DirectoryEntryTypeRegular
//	* DirectoryEntryTypeDirectory
//	* DirectoryEntryTypeCharacter
//	* DirectoryEntryTypeBlock
//	* DirectoryEntryTypeFIFO
//	* DirectoryEntryTypeSymlink
//	* DirectoryEntryTypeSocket
func (device *REDBrick) GetNextDirectoryEntry(directoryId uint16, sessionId uint16) (errorCode ErrorCode, nameStringId uint16, type_ DirectoryEntryType, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, directoryId);
	binary.Write(&buf, binary.LittleEndian, sessionId);

	resultBytes, err := device.device.Get(uint8(FunctionGetNextDirectoryEntry), buf.Bytes())
	if err != nil {
		return errorCode, nameStringId, type_, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 12 {
		return errorCode, nameStringId, type_, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
	}


	if header.ErrorCode != 0 {
		return errorCode, nameStringId, type_, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &errorCode)
	binary.Read(resultBuf, binary.LittleEndian, &nameStringId)
	binary.Read(resultBuf, binary.LittleEndian, &type_)


	return errorCode, nameStringId, type_, nil
}

// Rewinds a directory object and returns the resulting error code.
//
// Associated constants:
//
//	* ErrorCodeSuccess
//	* ErrorCodeUnknownError
//	* ErrorCodeInvalidOperation
//	* ErrorCodeOperationAborted
//	* ErrorCodeInternalError
//	* ErrorCodeUnknownSessionId
//	* ErrorCodeNoFreeSessionId
//	* ErrorCodeUnknownObjectId
//	* ErrorCodeNoFreeObjectId
//	* ErrorCodeObjectIsLocked
//	* ErrorCodeNoMoreData
//	* ErrorCodeWrongListItemType
//	* ErrorCodeProgramIsPurged
//	* ErrorCodeInvalidParameter
//	* ErrorCodeNoFreeMemory
//	* ErrorCodeNoFreeSpace
//	* ErrorCodeAccessDenied
//	* ErrorCodeAlreadyExists
//	* ErrorCodeDoesNotExist
//	* ErrorCodeInterrupted
//	* ErrorCodeIsDirectory
//	* ErrorCodeNotADirectory
//	* ErrorCodeWouldBlock
//	* ErrorCodeOverflow
//	* ErrorCodeBadFileDescriptor
//	* ErrorCodeOutOfRange
//	* ErrorCodeNameTooLong
//	* ErrorCodeInvalidSeek
//	* ErrorCodeNotSupported
//	* ErrorCodeTooManyOpenFiles
func (device *REDBrick) RewindDirectory(directoryId uint16) (errorCode ErrorCode, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, directoryId);

	resultBytes, err := device.device.Get(uint8(FunctionRewindDirectory), buf.Bytes())
	if err != nil {
		return errorCode, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 9 {
		return errorCode, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
	}


	if header.ErrorCode != 0 {
		return errorCode, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &errorCode)


	return errorCode, nil
}

// FIXME: name has to be absolute
//
// Associated constants:
//
//	* DirectoryFlagRecursive
//	* DirectoryFlagExclusive
//	* FilePermissionUserAll
//	* FilePermissionUserRead
//	* FilePermissionUserWrite
//	* FilePermissionUserExecute
//	* FilePermissionGroupAll
//	* FilePermissionGroupRead
//	* FilePermissionGroupWrite
//	* FilePermissionGroupExecute
//	* FilePermissionOthersAll
//	* FilePermissionOthersRead
//	* FilePermissionOthersWrite
//	* FilePermissionOthersExecute
//	* ErrorCodeSuccess
//	* ErrorCodeUnknownError
//	* ErrorCodeInvalidOperation
//	* ErrorCodeOperationAborted
//	* ErrorCodeInternalError
//	* ErrorCodeUnknownSessionId
//	* ErrorCodeNoFreeSessionId
//	* ErrorCodeUnknownObjectId
//	* ErrorCodeNoFreeObjectId
//	* ErrorCodeObjectIsLocked
//	* ErrorCodeNoMoreData
//	* ErrorCodeWrongListItemType
//	* ErrorCodeProgramIsPurged
//	* ErrorCodeInvalidParameter
//	* ErrorCodeNoFreeMemory
//	* ErrorCodeNoFreeSpace
//	* ErrorCodeAccessDenied
//	* ErrorCodeAlreadyExists
//	* ErrorCodeDoesNotExist
//	* ErrorCodeInterrupted
//	* ErrorCodeIsDirectory
//	* ErrorCodeNotADirectory
//	* ErrorCodeWouldBlock
//	* ErrorCodeOverflow
//	* ErrorCodeBadFileDescriptor
//	* ErrorCodeOutOfRange
//	* ErrorCodeNameTooLong
//	* ErrorCodeInvalidSeek
//	* ErrorCodeNotSupported
//	* ErrorCodeTooManyOpenFiles
func (device *REDBrick) CreateDirectory(nameStringId uint16, flags DirectoryFlag, permissions FilePermission, uid uint32, gid uint32) (errorCode ErrorCode, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, nameStringId);
	binary.Write(&buf, binary.LittleEndian, flags);
	binary.Write(&buf, binary.LittleEndian, permissions);
	binary.Write(&buf, binary.LittleEndian, uid);
	binary.Write(&buf, binary.LittleEndian, gid);

	resultBytes, err := device.device.Get(uint8(FunctionCreateDirectory), buf.Bytes())
	if err != nil {
		return errorCode, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 9 {
		return errorCode, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
	}


	if header.ErrorCode != 0 {
		return errorCode, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &errorCode)


	return errorCode, nil
}

// 
//
// Associated constants:
//
//	* ErrorCodeSuccess
//	* ErrorCodeUnknownError
//	* ErrorCodeInvalidOperation
//	* ErrorCodeOperationAborted
//	* ErrorCodeInternalError
//	* ErrorCodeUnknownSessionId
//	* ErrorCodeNoFreeSessionId
//	* ErrorCodeUnknownObjectId
//	* ErrorCodeNoFreeObjectId
//	* ErrorCodeObjectIsLocked
//	* ErrorCodeNoMoreData
//	* ErrorCodeWrongListItemType
//	* ErrorCodeProgramIsPurged
//	* ErrorCodeInvalidParameter
//	* ErrorCodeNoFreeMemory
//	* ErrorCodeNoFreeSpace
//	* ErrorCodeAccessDenied
//	* ErrorCodeAlreadyExists
//	* ErrorCodeDoesNotExist
//	* ErrorCodeInterrupted
//	* ErrorCodeIsDirectory
//	* ErrorCodeNotADirectory
//	* ErrorCodeWouldBlock
//	* ErrorCodeOverflow
//	* ErrorCodeBadFileDescriptor
//	* ErrorCodeOutOfRange
//	* ErrorCodeNameTooLong
//	* ErrorCodeInvalidSeek
//	* ErrorCodeNotSupported
//	* ErrorCodeTooManyOpenFiles
func (device *REDBrick) GetProcesses(sessionId uint16) (errorCode ErrorCode, processesListId uint16, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, sessionId);

	resultBytes, err := device.device.Get(uint8(FunctionGetProcesses), buf.Bytes())
	if err != nil {
		return errorCode, processesListId, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 11 {
		return errorCode, processesListId, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 11)
	}


	if header.ErrorCode != 0 {
		return errorCode, processesListId, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &errorCode)
	binary.Read(resultBuf, binary.LittleEndian, &processesListId)


	return errorCode, processesListId, nil
}

// 
//
// Associated constants:
//
//	* ErrorCodeSuccess
//	* ErrorCodeUnknownError
//	* ErrorCodeInvalidOperation
//	* ErrorCodeOperationAborted
//	* ErrorCodeInternalError
//	* ErrorCodeUnknownSessionId
//	* ErrorCodeNoFreeSessionId
//	* ErrorCodeUnknownObjectId
//	* ErrorCodeNoFreeObjectId
//	* ErrorCodeObjectIsLocked
//	* ErrorCodeNoMoreData
//	* ErrorCodeWrongListItemType
//	* ErrorCodeProgramIsPurged
//	* ErrorCodeInvalidParameter
//	* ErrorCodeNoFreeMemory
//	* ErrorCodeNoFreeSpace
//	* ErrorCodeAccessDenied
//	* ErrorCodeAlreadyExists
//	* ErrorCodeDoesNotExist
//	* ErrorCodeInterrupted
//	* ErrorCodeIsDirectory
//	* ErrorCodeNotADirectory
//	* ErrorCodeWouldBlock
//	* ErrorCodeOverflow
//	* ErrorCodeBadFileDescriptor
//	* ErrorCodeOutOfRange
//	* ErrorCodeNameTooLong
//	* ErrorCodeInvalidSeek
//	* ErrorCodeNotSupported
//	* ErrorCodeTooManyOpenFiles
func (device *REDBrick) SpawnProcess(executableStringId uint16, argumentsListId uint16, environmentListId uint16, workingDirectoryStringId uint16, uid uint32, gid uint32, stdinFileId uint16, stdoutFileId uint16, stderrFileId uint16, sessionId uint16) (errorCode ErrorCode, processId uint16, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, executableStringId);
	binary.Write(&buf, binary.LittleEndian, argumentsListId);
	binary.Write(&buf, binary.LittleEndian, environmentListId);
	binary.Write(&buf, binary.LittleEndian, workingDirectoryStringId);
	binary.Write(&buf, binary.LittleEndian, uid);
	binary.Write(&buf, binary.LittleEndian, gid);
	binary.Write(&buf, binary.LittleEndian, stdinFileId);
	binary.Write(&buf, binary.LittleEndian, stdoutFileId);
	binary.Write(&buf, binary.LittleEndian, stderrFileId);
	binary.Write(&buf, binary.LittleEndian, sessionId);

	resultBytes, err := device.device.Get(uint8(FunctionSpawnProcess), buf.Bytes())
	if err != nil {
		return errorCode, processId, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 11 {
		return errorCode, processId, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 11)
	}


	if header.ErrorCode != 0 {
		return errorCode, processId, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &errorCode)
	binary.Read(resultBuf, binary.LittleEndian, &processId)


	return errorCode, processId, nil
}

// Sends a UNIX signal to a process object and returns the resulting error code.
// 
// Possible UNIX signals are:
// 
// * Interrupt = 2
// * Quit = 3
// * Abort = 6
// * Kill = 9
// * User1 = 10
// * User2 = 12
// * Terminate = 15
// * Continue =  18
// * Stop = 19
//
// Associated constants:
//
//	* ProcessSignalInterrupt
//	* ProcessSignalQuit
//	* ProcessSignalAbort
//	* ProcessSignalKill
//	* ProcessSignalUser1
//	* ProcessSignalUser2
//	* ProcessSignalTerminate
//	* ProcessSignalContinue
//	* ProcessSignalStop
//	* ErrorCodeSuccess
//	* ErrorCodeUnknownError
//	* ErrorCodeInvalidOperation
//	* ErrorCodeOperationAborted
//	* ErrorCodeInternalError
//	* ErrorCodeUnknownSessionId
//	* ErrorCodeNoFreeSessionId
//	* ErrorCodeUnknownObjectId
//	* ErrorCodeNoFreeObjectId
//	* ErrorCodeObjectIsLocked
//	* ErrorCodeNoMoreData
//	* ErrorCodeWrongListItemType
//	* ErrorCodeProgramIsPurged
//	* ErrorCodeInvalidParameter
//	* ErrorCodeNoFreeMemory
//	* ErrorCodeNoFreeSpace
//	* ErrorCodeAccessDenied
//	* ErrorCodeAlreadyExists
//	* ErrorCodeDoesNotExist
//	* ErrorCodeInterrupted
//	* ErrorCodeIsDirectory
//	* ErrorCodeNotADirectory
//	* ErrorCodeWouldBlock
//	* ErrorCodeOverflow
//	* ErrorCodeBadFileDescriptor
//	* ErrorCodeOutOfRange
//	* ErrorCodeNameTooLong
//	* ErrorCodeInvalidSeek
//	* ErrorCodeNotSupported
//	* ErrorCodeTooManyOpenFiles
func (device *REDBrick) KillProcess(processId uint16, signal ProcessSignal) (errorCode ErrorCode, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, processId);
	binary.Write(&buf, binary.LittleEndian, signal);

	resultBytes, err := device.device.Get(uint8(FunctionKillProcess), buf.Bytes())
	if err != nil {
		return errorCode, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 9 {
		return errorCode, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
	}


	if header.ErrorCode != 0 {
		return errorCode, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &errorCode)


	return errorCode, nil
}

// Returns the executable, arguments, environment and working directory used to
// spawn a process object, as passed to SpawnProcess, and the resulting
// error code.
//
// Associated constants:
//
//	* ErrorCodeSuccess
//	* ErrorCodeUnknownError
//	* ErrorCodeInvalidOperation
//	* ErrorCodeOperationAborted
//	* ErrorCodeInternalError
//	* ErrorCodeUnknownSessionId
//	* ErrorCodeNoFreeSessionId
//	* ErrorCodeUnknownObjectId
//	* ErrorCodeNoFreeObjectId
//	* ErrorCodeObjectIsLocked
//	* ErrorCodeNoMoreData
//	* ErrorCodeWrongListItemType
//	* ErrorCodeProgramIsPurged
//	* ErrorCodeInvalidParameter
//	* ErrorCodeNoFreeMemory
//	* ErrorCodeNoFreeSpace
//	* ErrorCodeAccessDenied
//	* ErrorCodeAlreadyExists
//	* ErrorCodeDoesNotExist
//	* ErrorCodeInterrupted
//	* ErrorCodeIsDirectory
//	* ErrorCodeNotADirectory
//	* ErrorCodeWouldBlock
//	* ErrorCodeOverflow
//	* ErrorCodeBadFileDescriptor
//	* ErrorCodeOutOfRange
//	* ErrorCodeNameTooLong
//	* ErrorCodeInvalidSeek
//	* ErrorCodeNotSupported
//	* ErrorCodeTooManyOpenFiles
func (device *REDBrick) GetProcessCommand(processId uint16, sessionId uint16) (errorCode ErrorCode, executableStringId uint16, argumentsListId uint16, environmentListId uint16, workingDirectoryStringId uint16, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, processId);
	binary.Write(&buf, binary.LittleEndian, sessionId);

	resultBytes, err := device.device.Get(uint8(FunctionGetProcessCommand), buf.Bytes())
	if err != nil {
		return errorCode, executableStringId, argumentsListId, environmentListId, workingDirectoryStringId, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 17 {
		return errorCode, executableStringId, argumentsListId, environmentListId, workingDirectoryStringId, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 17)
	}


	if header.ErrorCode != 0 {
		return errorCode, executableStringId, argumentsListId, environmentListId, workingDirectoryStringId, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &errorCode)
	binary.Read(resultBuf, binary.LittleEndian, &executableStringId)
	binary.Read(resultBuf, binary.LittleEndian, &argumentsListId)
	binary.Read(resultBuf, binary.LittleEndian, &environmentListId)
	binary.Read(resultBuf, binary.LittleEndian, &workingDirectoryStringId)


	return errorCode, executableStringId, argumentsListId, environmentListId, workingDirectoryStringId, nil
}

// Returns the process ID and the user and group ID used to spawn a process object,
// as passed to SpawnProcess, and the resulting error code.
// 
// The process ID is only valid if the state is *Running* or *Stopped*, see
// GetProcessState.
//
// Associated constants:
//
//	* ErrorCodeSuccess
//	* ErrorCodeUnknownError
//	* ErrorCodeInvalidOperation
//	* ErrorCodeOperationAborted
//	* ErrorCodeInternalError
//	* ErrorCodeUnknownSessionId
//	* ErrorCodeNoFreeSessionId
//	* ErrorCodeUnknownObjectId
//	* ErrorCodeNoFreeObjectId
//	* ErrorCodeObjectIsLocked
//	* ErrorCodeNoMoreData
//	* ErrorCodeWrongListItemType
//	* ErrorCodeProgramIsPurged
//	* ErrorCodeInvalidParameter
//	* ErrorCodeNoFreeMemory
//	* ErrorCodeNoFreeSpace
//	* ErrorCodeAccessDenied
//	* ErrorCodeAlreadyExists
//	* ErrorCodeDoesNotExist
//	* ErrorCodeInterrupted
//	* ErrorCodeIsDirectory
//	* ErrorCodeNotADirectory
//	* ErrorCodeWouldBlock
//	* ErrorCodeOverflow
//	* ErrorCodeBadFileDescriptor
//	* ErrorCodeOutOfRange
//	* ErrorCodeNameTooLong
//	* ErrorCodeInvalidSeek
//	* ErrorCodeNotSupported
//	* ErrorCodeTooManyOpenFiles
func (device *REDBrick) GetProcessIdentity(processId uint16) (errorCode ErrorCode, pid uint32, uid uint32, gid uint32, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, processId);

	resultBytes, err := device.device.Get(uint8(FunctionGetProcessIdentity), buf.Bytes())
	if err != nil {
		return errorCode, pid, uid, gid, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 21 {
		return errorCode, pid, uid, gid, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 21)
	}


	if header.ErrorCode != 0 {
		return errorCode, pid, uid, gid, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &errorCode)
	binary.Read(resultBuf, binary.LittleEndian, &pid)
	binary.Read(resultBuf, binary.LittleEndian, &uid)
	binary.Read(resultBuf, binary.LittleEndian, &gid)


	return errorCode, pid, uid, gid, nil
}

// Returns the stdin, stdout and stderr files used to spawn a process object, as
// passed to SpawnProcess, and the resulting error code.
//
// Associated constants:
//
//	* ErrorCodeSuccess
//	* ErrorCodeUnknownError
//	* ErrorCodeInvalidOperation
//	* ErrorCodeOperationAborted
//	* ErrorCodeInternalError
//	* ErrorCodeUnknownSessionId
//	* ErrorCodeNoFreeSessionId
//	* ErrorCodeUnknownObjectId
//	* ErrorCodeNoFreeObjectId
//	* ErrorCodeObjectIsLocked
//	* ErrorCodeNoMoreData
//	* ErrorCodeWrongListItemType
//	* ErrorCodeProgramIsPurged
//	* ErrorCodeInvalidParameter
//	* ErrorCodeNoFreeMemory
//	* ErrorCodeNoFreeSpace
//	* ErrorCodeAccessDenied
//	* ErrorCodeAlreadyExists
//	* ErrorCodeDoesNotExist
//	* ErrorCodeInterrupted
//	* ErrorCodeIsDirectory
//	* ErrorCodeNotADirectory
//	* ErrorCodeWouldBlock
//	* ErrorCodeOverflow
//	* ErrorCodeBadFileDescriptor
//	* ErrorCodeOutOfRange
//	* ErrorCodeNameTooLong
//	* ErrorCodeInvalidSeek
//	* ErrorCodeNotSupported
//	* ErrorCodeTooManyOpenFiles
func (device *REDBrick) GetProcessStdio(processId uint16, sessionId uint16) (errorCode ErrorCode, stdinFileId uint16, stdoutFileId uint16, stderrFileId uint16, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, processId);
	binary.Write(&buf, binary.LittleEndian, sessionId);

	resultBytes, err := device.device.Get(uint8(FunctionGetProcessStdio), buf.Bytes())
	if err != nil {
		return errorCode, stdinFileId, stdoutFileId, stderrFileId, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 15 {
		return errorCode, stdinFileId, stdoutFileId, stderrFileId, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 15)
	}


	if header.ErrorCode != 0 {
		return errorCode, stdinFileId, stdoutFileId, stderrFileId, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &errorCode)
	binary.Read(resultBuf, binary.LittleEndian, &stdinFileId)
	binary.Read(resultBuf, binary.LittleEndian, &stdoutFileId)
	binary.Read(resultBuf, binary.LittleEndian, &stderrFileId)


	return errorCode, stdinFileId, stdoutFileId, stderrFileId, nil
}

// Returns the current state, timestamp and exit code of a process object, and
// the resulting error code.
// 
// Possible process states are:
// 
// * Unknown = 0
// * Running = 1
// * Error = 2
// * Exited = 3
// * Killed = 4
// * Stopped = 5
// 
// The timestamp represents the UNIX time since when the process is in its current
// state.
// 
// The exit code is only valid if the state is *Error*, *Exited*, *Killed* or
// *Stopped* and has different meanings depending on the state:
// 
// * Error: error code for error occurred while spawning the process (see below)
// * Exited: exit status of the process
// * Killed: UNIX signal number used to kill the process
// * Stopped: UNIX signal number used to stop the process
// 
// Possible exit/error codes in *Error* state are:
// 
// * InternalError = 125
// * CannotExecute = 126
// * DoesNotExist = 127
// 
// The *CannotExecute* error can be caused by the executable being opened for
// writing.
//
// Associated constants:
//
//	* ErrorCodeSuccess
//	* ErrorCodeUnknownError
//	* ErrorCodeInvalidOperation
//	* ErrorCodeOperationAborted
//	* ErrorCodeInternalError
//	* ErrorCodeUnknownSessionId
//	* ErrorCodeNoFreeSessionId
//	* ErrorCodeUnknownObjectId
//	* ErrorCodeNoFreeObjectId
//	* ErrorCodeObjectIsLocked
//	* ErrorCodeNoMoreData
//	* ErrorCodeWrongListItemType
//	* ErrorCodeProgramIsPurged
//	* ErrorCodeInvalidParameter
//	* ErrorCodeNoFreeMemory
//	* ErrorCodeNoFreeSpace
//	* ErrorCodeAccessDenied
//	* ErrorCodeAlreadyExists
//	* ErrorCodeDoesNotExist
//	* ErrorCodeInterrupted
//	* ErrorCodeIsDirectory
//	* ErrorCodeNotADirectory
//	* ErrorCodeWouldBlock
//	* ErrorCodeOverflow
//	* ErrorCodeBadFileDescriptor
//	* ErrorCodeOutOfRange
//	* ErrorCodeNameTooLong
//	* ErrorCodeInvalidSeek
//	* ErrorCodeNotSupported
//	* ErrorCodeTooManyOpenFiles
//	* ProcessStateUnknown
//	* ProcessStateRunning
//	* ProcessStateError
//	* ProcessStateExited
//	* ProcessStateKilled
//	* ProcessStateStopped
func (device *REDBrick) GetProcessState(processId uint16) (errorCode ErrorCode, state ProcessState, timestamp uint64, exitCode uint8, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, processId);

	resultBytes, err := device.device.Get(uint8(FunctionGetProcessState), buf.Bytes())
	if err != nil {
		return errorCode, state, timestamp, exitCode, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 19 {
		return errorCode, state, timestamp, exitCode, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 19)
	}


	if header.ErrorCode != 0 {
		return errorCode, state, timestamp, exitCode, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &errorCode)
	binary.Read(resultBuf, binary.LittleEndian, &state)
	binary.Read(resultBuf, binary.LittleEndian, &timestamp)
	binary.Read(resultBuf, binary.LittleEndian, &exitCode)


	return errorCode, state, timestamp, exitCode, nil
}

// 
//
// Associated constants:
//
//	* ErrorCodeSuccess
//	* ErrorCodeUnknownError
//	* ErrorCodeInvalidOperation
//	* ErrorCodeOperationAborted
//	* ErrorCodeInternalError
//	* ErrorCodeUnknownSessionId
//	* ErrorCodeNoFreeSessionId
//	* ErrorCodeUnknownObjectId
//	* ErrorCodeNoFreeObjectId
//	* ErrorCodeObjectIsLocked
//	* ErrorCodeNoMoreData
//	* ErrorCodeWrongListItemType
//	* ErrorCodeProgramIsPurged
//	* ErrorCodeInvalidParameter
//	* ErrorCodeNoFreeMemory
//	* ErrorCodeNoFreeSpace
//	* ErrorCodeAccessDenied
//	* ErrorCodeAlreadyExists
//	* ErrorCodeDoesNotExist
//	* ErrorCodeInterrupted
//	* ErrorCodeIsDirectory
//	* ErrorCodeNotADirectory
//	* ErrorCodeWouldBlock
//	* ErrorCodeOverflow
//	* ErrorCodeBadFileDescriptor
//	* ErrorCodeOutOfRange
//	* ErrorCodeNameTooLong
//	* ErrorCodeInvalidSeek
//	* ErrorCodeNotSupported
//	* ErrorCodeTooManyOpenFiles
func (device *REDBrick) GetPrograms(sessionId uint16) (errorCode ErrorCode, programsListId uint16, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, sessionId);

	resultBytes, err := device.device.Get(uint8(FunctionGetPrograms), buf.Bytes())
	if err != nil {
		return errorCode, programsListId, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 11 {
		return errorCode, programsListId, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 11)
	}


	if header.ErrorCode != 0 {
		return errorCode, programsListId, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &errorCode)
	binary.Read(resultBuf, binary.LittleEndian, &programsListId)


	return errorCode, programsListId, nil
}

// 
//
// Associated constants:
//
//	* ErrorCodeSuccess
//	* ErrorCodeUnknownError
//	* ErrorCodeInvalidOperation
//	* ErrorCodeOperationAborted
//	* ErrorCodeInternalError
//	* ErrorCodeUnknownSessionId
//	* ErrorCodeNoFreeSessionId
//	* ErrorCodeUnknownObjectId
//	* ErrorCodeNoFreeObjectId
//	* ErrorCodeObjectIsLocked
//	* ErrorCodeNoMoreData
//	* ErrorCodeWrongListItemType
//	* ErrorCodeProgramIsPurged
//	* ErrorCodeInvalidParameter
//	* ErrorCodeNoFreeMemory
//	* ErrorCodeNoFreeSpace
//	* ErrorCodeAccessDenied
//	* ErrorCodeAlreadyExists
//	* ErrorCodeDoesNotExist
//	* ErrorCodeInterrupted
//	* ErrorCodeIsDirectory
//	* ErrorCodeNotADirectory
//	* ErrorCodeWouldBlock
//	* ErrorCodeOverflow
//	* ErrorCodeBadFileDescriptor
//	* ErrorCodeOutOfRange
//	* ErrorCodeNameTooLong
//	* ErrorCodeInvalidSeek
//	* ErrorCodeNotSupported
//	* ErrorCodeTooManyOpenFiles
func (device *REDBrick) DefineProgram(identifierStringId uint16, sessionId uint16) (errorCode ErrorCode, programId uint16, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, identifierStringId);
	binary.Write(&buf, binary.LittleEndian, sessionId);

	resultBytes, err := device.device.Get(uint8(FunctionDefineProgram), buf.Bytes())
	if err != nil {
		return errorCode, programId, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 11 {
		return errorCode, programId, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 11)
	}


	if header.ErrorCode != 0 {
		return errorCode, programId, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &errorCode)
	binary.Read(resultBuf, binary.LittleEndian, &programId)


	return errorCode, programId, nil
}

// 
//
// Associated constants:
//
//	* ErrorCodeSuccess
//	* ErrorCodeUnknownError
//	* ErrorCodeInvalidOperation
//	* ErrorCodeOperationAborted
//	* ErrorCodeInternalError
//	* ErrorCodeUnknownSessionId
//	* ErrorCodeNoFreeSessionId
//	* ErrorCodeUnknownObjectId
//	* ErrorCodeNoFreeObjectId
//	* ErrorCodeObjectIsLocked
//	* ErrorCodeNoMoreData
//	* ErrorCodeWrongListItemType
//	* ErrorCodeProgramIsPurged
//	* ErrorCodeInvalidParameter
//	* ErrorCodeNoFreeMemory
//	* ErrorCodeNoFreeSpace
//	* ErrorCodeAccessDenied
//	* ErrorCodeAlreadyExists
//	* ErrorCodeDoesNotExist
//	* ErrorCodeInterrupted
//	* ErrorCodeIsDirectory
//	* ErrorCodeNotADirectory
//	* ErrorCodeWouldBlock
//	* ErrorCodeOverflow
//	* ErrorCodeBadFileDescriptor
//	* ErrorCodeOutOfRange
//	* ErrorCodeNameTooLong
//	* ErrorCodeInvalidSeek
//	* ErrorCodeNotSupported
//	* ErrorCodeTooManyOpenFiles
func (device *REDBrick) PurgeProgram(programId uint16, cookie uint32) (errorCode ErrorCode, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, programId);
	binary.Write(&buf, binary.LittleEndian, cookie);

	resultBytes, err := device.device.Get(uint8(FunctionPurgeProgram), buf.Bytes())
	if err != nil {
		return errorCode, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 9 {
		return errorCode, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
	}


	if header.ErrorCode != 0 {
		return errorCode, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &errorCode)


	return errorCode, nil
}

// 
//
// Associated constants:
//
//	* ErrorCodeSuccess
//	* ErrorCodeUnknownError
//	* ErrorCodeInvalidOperation
//	* ErrorCodeOperationAborted
//	* ErrorCodeInternalError
//	* ErrorCodeUnknownSessionId
//	* ErrorCodeNoFreeSessionId
//	* ErrorCodeUnknownObjectId
//	* ErrorCodeNoFreeObjectId
//	* ErrorCodeObjectIsLocked
//	* ErrorCodeNoMoreData
//	* ErrorCodeWrongListItemType
//	* ErrorCodeProgramIsPurged
//	* ErrorCodeInvalidParameter
//	* ErrorCodeNoFreeMemory
//	* ErrorCodeNoFreeSpace
//	* ErrorCodeAccessDenied
//	* ErrorCodeAlreadyExists
//	* ErrorCodeDoesNotExist
//	* ErrorCodeInterrupted
//	* ErrorCodeIsDirectory
//	* ErrorCodeNotADirectory
//	* ErrorCodeWouldBlock
//	* ErrorCodeOverflow
//	* ErrorCodeBadFileDescriptor
//	* ErrorCodeOutOfRange
//	* ErrorCodeNameTooLong
//	* ErrorCodeInvalidSeek
//	* ErrorCodeNotSupported
//	* ErrorCodeTooManyOpenFiles
func (device *REDBrick) GetProgramIdentifier(programId uint16, sessionId uint16) (errorCode ErrorCode, identifierStringId uint16, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, programId);
	binary.Write(&buf, binary.LittleEndian, sessionId);

	resultBytes, err := device.device.Get(uint8(FunctionGetProgramIdentifier), buf.Bytes())
	if err != nil {
		return errorCode, identifierStringId, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 11 {
		return errorCode, identifierStringId, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 11)
	}


	if header.ErrorCode != 0 {
		return errorCode, identifierStringId, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &errorCode)
	binary.Read(resultBuf, binary.LittleEndian, &identifierStringId)


	return errorCode, identifierStringId, nil
}

// FIXME: root directory is absolute: <home>/programs/<identifier>
//
// Associated constants:
//
//	* ErrorCodeSuccess
//	* ErrorCodeUnknownError
//	* ErrorCodeInvalidOperation
//	* ErrorCodeOperationAborted
//	* ErrorCodeInternalError
//	* ErrorCodeUnknownSessionId
//	* ErrorCodeNoFreeSessionId
//	* ErrorCodeUnknownObjectId
//	* ErrorCodeNoFreeObjectId
//	* ErrorCodeObjectIsLocked
//	* ErrorCodeNoMoreData
//	* ErrorCodeWrongListItemType
//	* ErrorCodeProgramIsPurged
//	* ErrorCodeInvalidParameter
//	* ErrorCodeNoFreeMemory
//	* ErrorCodeNoFreeSpace
//	* ErrorCodeAccessDenied
//	* ErrorCodeAlreadyExists
//	* ErrorCodeDoesNotExist
//	* ErrorCodeInterrupted
//	* ErrorCodeIsDirectory
//	* ErrorCodeNotADirectory
//	* ErrorCodeWouldBlock
//	* ErrorCodeOverflow
//	* ErrorCodeBadFileDescriptor
//	* ErrorCodeOutOfRange
//	* ErrorCodeNameTooLong
//	* ErrorCodeInvalidSeek
//	* ErrorCodeNotSupported
//	* ErrorCodeTooManyOpenFiles
func (device *REDBrick) GetProgramRootDirectory(programId uint16, sessionId uint16) (errorCode ErrorCode, rootDirectoryStringId uint16, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, programId);
	binary.Write(&buf, binary.LittleEndian, sessionId);

	resultBytes, err := device.device.Get(uint8(FunctionGetProgramRootDirectory), buf.Bytes())
	if err != nil {
		return errorCode, rootDirectoryStringId, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 11 {
		return errorCode, rootDirectoryStringId, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 11)
	}


	if header.ErrorCode != 0 {
		return errorCode, rootDirectoryStringId, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &errorCode)
	binary.Read(resultBuf, binary.LittleEndian, &rootDirectoryStringId)


	return errorCode, rootDirectoryStringId, nil
}

// FIXME: working directory is relative to <home>/programs/<identifier>/bin
//
// Associated constants:
//
//	* ErrorCodeSuccess
//	* ErrorCodeUnknownError
//	* ErrorCodeInvalidOperation
//	* ErrorCodeOperationAborted
//	* ErrorCodeInternalError
//	* ErrorCodeUnknownSessionId
//	* ErrorCodeNoFreeSessionId
//	* ErrorCodeUnknownObjectId
//	* ErrorCodeNoFreeObjectId
//	* ErrorCodeObjectIsLocked
//	* ErrorCodeNoMoreData
//	* ErrorCodeWrongListItemType
//	* ErrorCodeProgramIsPurged
//	* ErrorCodeInvalidParameter
//	* ErrorCodeNoFreeMemory
//	* ErrorCodeNoFreeSpace
//	* ErrorCodeAccessDenied
//	* ErrorCodeAlreadyExists
//	* ErrorCodeDoesNotExist
//	* ErrorCodeInterrupted
//	* ErrorCodeIsDirectory
//	* ErrorCodeNotADirectory
//	* ErrorCodeWouldBlock
//	* ErrorCodeOverflow
//	* ErrorCodeBadFileDescriptor
//	* ErrorCodeOutOfRange
//	* ErrorCodeNameTooLong
//	* ErrorCodeInvalidSeek
//	* ErrorCodeNotSupported
//	* ErrorCodeTooManyOpenFiles
func (device *REDBrick) SetProgramCommand(programId uint16, executableStringId uint16, argumentsListId uint16, environmentListId uint16, workingDirectoryStringId uint16) (errorCode ErrorCode, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, programId);
	binary.Write(&buf, binary.LittleEndian, executableStringId);
	binary.Write(&buf, binary.LittleEndian, argumentsListId);
	binary.Write(&buf, binary.LittleEndian, environmentListId);
	binary.Write(&buf, binary.LittleEndian, workingDirectoryStringId);

	resultBytes, err := device.device.Get(uint8(FunctionSetProgramCommand), buf.Bytes())
	if err != nil {
		return errorCode, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 9 {
		return errorCode, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
	}


	if header.ErrorCode != 0 {
		return errorCode, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &errorCode)


	return errorCode, nil
}

// FIXME: working directory is relative to <home>/programs/<identifier>/bin
//
// Associated constants:
//
//	* ErrorCodeSuccess
//	* ErrorCodeUnknownError
//	* ErrorCodeInvalidOperation
//	* ErrorCodeOperationAborted
//	* ErrorCodeInternalError
//	* ErrorCodeUnknownSessionId
//	* ErrorCodeNoFreeSessionId
//	* ErrorCodeUnknownObjectId
//	* ErrorCodeNoFreeObjectId
//	* ErrorCodeObjectIsLocked
//	* ErrorCodeNoMoreData
//	* ErrorCodeWrongListItemType
//	* ErrorCodeProgramIsPurged
//	* ErrorCodeInvalidParameter
//	* ErrorCodeNoFreeMemory
//	* ErrorCodeNoFreeSpace
//	* ErrorCodeAccessDenied
//	* ErrorCodeAlreadyExists
//	* ErrorCodeDoesNotExist
//	* ErrorCodeInterrupted
//	* ErrorCodeIsDirectory
//	* ErrorCodeNotADirectory
//	* ErrorCodeWouldBlock
//	* ErrorCodeOverflow
//	* ErrorCodeBadFileDescriptor
//	* ErrorCodeOutOfRange
//	* ErrorCodeNameTooLong
//	* ErrorCodeInvalidSeek
//	* ErrorCodeNotSupported
//	* ErrorCodeTooManyOpenFiles
func (device *REDBrick) GetProgramCommand(programId uint16, sessionId uint16) (errorCode ErrorCode, executableStringId uint16, argumentsListId uint16, environmentListId uint16, workingDirectoryStringId uint16, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, programId);
	binary.Write(&buf, binary.LittleEndian, sessionId);

	resultBytes, err := device.device.Get(uint8(FunctionGetProgramCommand), buf.Bytes())
	if err != nil {
		return errorCode, executableStringId, argumentsListId, environmentListId, workingDirectoryStringId, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 17 {
		return errorCode, executableStringId, argumentsListId, environmentListId, workingDirectoryStringId, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 17)
	}


	if header.ErrorCode != 0 {
		return errorCode, executableStringId, argumentsListId, environmentListId, workingDirectoryStringId, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &errorCode)
	binary.Read(resultBuf, binary.LittleEndian, &executableStringId)
	binary.Read(resultBuf, binary.LittleEndian, &argumentsListId)
	binary.Read(resultBuf, binary.LittleEndian, &environmentListId)
	binary.Read(resultBuf, binary.LittleEndian, &workingDirectoryStringId)


	return errorCode, executableStringId, argumentsListId, environmentListId, workingDirectoryStringId, nil
}

// FIXME: stdio file names are relative to <home>/programs/<identifier>/bin
//
// Associated constants:
//
//	* ProgramStdioRedirectionDevNull
//	* ProgramStdioRedirectionPipe
//	* ProgramStdioRedirectionFile
//	* ProgramStdioRedirectionIndividualLog
//	* ProgramStdioRedirectionContinuousLog
//	* ProgramStdioRedirectionStdout
//	* ErrorCodeSuccess
//	* ErrorCodeUnknownError
//	* ErrorCodeInvalidOperation
//	* ErrorCodeOperationAborted
//	* ErrorCodeInternalError
//	* ErrorCodeUnknownSessionId
//	* ErrorCodeNoFreeSessionId
//	* ErrorCodeUnknownObjectId
//	* ErrorCodeNoFreeObjectId
//	* ErrorCodeObjectIsLocked
//	* ErrorCodeNoMoreData
//	* ErrorCodeWrongListItemType
//	* ErrorCodeProgramIsPurged
//	* ErrorCodeInvalidParameter
//	* ErrorCodeNoFreeMemory
//	* ErrorCodeNoFreeSpace
//	* ErrorCodeAccessDenied
//	* ErrorCodeAlreadyExists
//	* ErrorCodeDoesNotExist
//	* ErrorCodeInterrupted
//	* ErrorCodeIsDirectory
//	* ErrorCodeNotADirectory
//	* ErrorCodeWouldBlock
//	* ErrorCodeOverflow
//	* ErrorCodeBadFileDescriptor
//	* ErrorCodeOutOfRange
//	* ErrorCodeNameTooLong
//	* ErrorCodeInvalidSeek
//	* ErrorCodeNotSupported
//	* ErrorCodeTooManyOpenFiles
func (device *REDBrick) SetProgramStdioRedirection(programId uint16, stdinRedirection ProgramStdioRedirection, stdinFileNameStringId uint16, stdoutRedirection ProgramStdioRedirection, stdoutFileNameStringId uint16, stderrRedirection ProgramStdioRedirection, stderrFileNameStringId uint16) (errorCode ErrorCode, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, programId);
	binary.Write(&buf, binary.LittleEndian, stdinRedirection);
	binary.Write(&buf, binary.LittleEndian, stdinFileNameStringId);
	binary.Write(&buf, binary.LittleEndian, stdoutRedirection);
	binary.Write(&buf, binary.LittleEndian, stdoutFileNameStringId);
	binary.Write(&buf, binary.LittleEndian, stderrRedirection);
	binary.Write(&buf, binary.LittleEndian, stderrFileNameStringId);

	resultBytes, err := device.device.Get(uint8(FunctionSetProgramStdioRedirection), buf.Bytes())
	if err != nil {
		return errorCode, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 9 {
		return errorCode, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
	}


	if header.ErrorCode != 0 {
		return errorCode, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &errorCode)


	return errorCode, nil
}

// FIXME: stdio file names are relative to <home>/programs/<identifier>/bin
//
// Associated constants:
//
//	* ErrorCodeSuccess
//	* ErrorCodeUnknownError
//	* ErrorCodeInvalidOperation
//	* ErrorCodeOperationAborted
//	* ErrorCodeInternalError
//	* ErrorCodeUnknownSessionId
//	* ErrorCodeNoFreeSessionId
//	* ErrorCodeUnknownObjectId
//	* ErrorCodeNoFreeObjectId
//	* ErrorCodeObjectIsLocked
//	* ErrorCodeNoMoreData
//	* ErrorCodeWrongListItemType
//	* ErrorCodeProgramIsPurged
//	* ErrorCodeInvalidParameter
//	* ErrorCodeNoFreeMemory
//	* ErrorCodeNoFreeSpace
//	* ErrorCodeAccessDenied
//	* ErrorCodeAlreadyExists
//	* ErrorCodeDoesNotExist
//	* ErrorCodeInterrupted
//	* ErrorCodeIsDirectory
//	* ErrorCodeNotADirectory
//	* ErrorCodeWouldBlock
//	* ErrorCodeOverflow
//	* ErrorCodeBadFileDescriptor
//	* ErrorCodeOutOfRange
//	* ErrorCodeNameTooLong
//	* ErrorCodeInvalidSeek
//	* ErrorCodeNotSupported
//	* ErrorCodeTooManyOpenFiles
//	* ProgramStdioRedirectionDevNull
//	* ProgramStdioRedirectionPipe
//	* ProgramStdioRedirectionFile
//	* ProgramStdioRedirectionIndividualLog
//	* ProgramStdioRedirectionContinuousLog
//	* ProgramStdioRedirectionStdout
func (device *REDBrick) GetProgramStdioRedirection(programId uint16, sessionId uint16) (errorCode ErrorCode, stdinRedirection ProgramStdioRedirection, stdinFileNameStringId uint16, stdoutRedirection ProgramStdioRedirection, stdoutFileNameStringId uint16, stderrRedirection ProgramStdioRedirection, stderrFileNameStringId uint16, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, programId);
	binary.Write(&buf, binary.LittleEndian, sessionId);

	resultBytes, err := device.device.Get(uint8(FunctionGetProgramStdioRedirection), buf.Bytes())
	if err != nil {
		return errorCode, stdinRedirection, stdinFileNameStringId, stdoutRedirection, stdoutFileNameStringId, stderrRedirection, stderrFileNameStringId, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 18 {
		return errorCode, stdinRedirection, stdinFileNameStringId, stdoutRedirection, stdoutFileNameStringId, stderrRedirection, stderrFileNameStringId, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 18)
	}


	if header.ErrorCode != 0 {
		return errorCode, stdinRedirection, stdinFileNameStringId, stdoutRedirection, stdoutFileNameStringId, stderrRedirection, stderrFileNameStringId, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &errorCode)
	binary.Read(resultBuf, binary.LittleEndian, &stdinRedirection)
	binary.Read(resultBuf, binary.LittleEndian, &stdinFileNameStringId)
	binary.Read(resultBuf, binary.LittleEndian, &stdoutRedirection)
	binary.Read(resultBuf, binary.LittleEndian, &stdoutFileNameStringId)
	binary.Read(resultBuf, binary.LittleEndian, &stderrRedirection)
	binary.Read(resultBuf, binary.LittleEndian, &stderrFileNameStringId)


	return errorCode, stdinRedirection, stdinFileNameStringId, stdoutRedirection, stdoutFileNameStringId, stderrRedirection, stderrFileNameStringId, nil
}

// 
//
// Associated constants:
//
//	* ProgramStartModeNever
//	* ProgramStartModeAlways
//	* ProgramStartModeInterval
//	* ProgramStartModeCron
//	* ErrorCodeSuccess
//	* ErrorCodeUnknownError
//	* ErrorCodeInvalidOperation
//	* ErrorCodeOperationAborted
//	* ErrorCodeInternalError
//	* ErrorCodeUnknownSessionId
//	* ErrorCodeNoFreeSessionId
//	* ErrorCodeUnknownObjectId
//	* ErrorCodeNoFreeObjectId
//	* ErrorCodeObjectIsLocked
//	* ErrorCodeNoMoreData
//	* ErrorCodeWrongListItemType
//	* ErrorCodeProgramIsPurged
//	* ErrorCodeInvalidParameter
//	* ErrorCodeNoFreeMemory
//	* ErrorCodeNoFreeSpace
//	* ErrorCodeAccessDenied
//	* ErrorCodeAlreadyExists
//	* ErrorCodeDoesNotExist
//	* ErrorCodeInterrupted
//	* ErrorCodeIsDirectory
//	* ErrorCodeNotADirectory
//	* ErrorCodeWouldBlock
//	* ErrorCodeOverflow
//	* ErrorCodeBadFileDescriptor
//	* ErrorCodeOutOfRange
//	* ErrorCodeNameTooLong
//	* ErrorCodeInvalidSeek
//	* ErrorCodeNotSupported
//	* ErrorCodeTooManyOpenFiles
func (device *REDBrick) SetProgramSchedule(programId uint16, startMode ProgramStartMode, continueAfterError bool, startInterval uint32, startFieldsStringId uint16) (errorCode ErrorCode, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, programId);
	binary.Write(&buf, binary.LittleEndian, startMode);
	binary.Write(&buf, binary.LittleEndian, continueAfterError);
	binary.Write(&buf, binary.LittleEndian, startInterval);
	binary.Write(&buf, binary.LittleEndian, startFieldsStringId);

	resultBytes, err := device.device.Get(uint8(FunctionSetProgramSchedule), buf.Bytes())
	if err != nil {
		return errorCode, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 9 {
		return errorCode, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
	}


	if header.ErrorCode != 0 {
		return errorCode, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &errorCode)


	return errorCode, nil
}

// 
//
// Associated constants:
//
//	* ErrorCodeSuccess
//	* ErrorCodeUnknownError
//	* ErrorCodeInvalidOperation
//	* ErrorCodeOperationAborted
//	* ErrorCodeInternalError
//	* ErrorCodeUnknownSessionId
//	* ErrorCodeNoFreeSessionId
//	* ErrorCodeUnknownObjectId
//	* ErrorCodeNoFreeObjectId
//	* ErrorCodeObjectIsLocked
//	* ErrorCodeNoMoreData
//	* ErrorCodeWrongListItemType
//	* ErrorCodeProgramIsPurged
//	* ErrorCodeInvalidParameter
//	* ErrorCodeNoFreeMemory
//	* ErrorCodeNoFreeSpace
//	* ErrorCodeAccessDenied
//	* ErrorCodeAlreadyExists
//	* ErrorCodeDoesNotExist
//	* ErrorCodeInterrupted
//	* ErrorCodeIsDirectory
//	* ErrorCodeNotADirectory
//	* ErrorCodeWouldBlock
//	* ErrorCodeOverflow
//	* ErrorCodeBadFileDescriptor
//	* ErrorCodeOutOfRange
//	* ErrorCodeNameTooLong
//	* ErrorCodeInvalidSeek
//	* ErrorCodeNotSupported
//	* ErrorCodeTooManyOpenFiles
//	* ProgramStartModeNever
//	* ProgramStartModeAlways
//	* ProgramStartModeInterval
//	* ProgramStartModeCron
func (device *REDBrick) GetProgramSchedule(programId uint16, sessionId uint16) (errorCode ErrorCode, startMode ProgramStartMode, continueAfterError bool, startInterval uint32, startFieldsStringId uint16, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, programId);
	binary.Write(&buf, binary.LittleEndian, sessionId);

	resultBytes, err := device.device.Get(uint8(FunctionGetProgramSchedule), buf.Bytes())
	if err != nil {
		return errorCode, startMode, continueAfterError, startInterval, startFieldsStringId, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 17 {
		return errorCode, startMode, continueAfterError, startInterval, startFieldsStringId, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 17)
	}


	if header.ErrorCode != 0 {
		return errorCode, startMode, continueAfterError, startInterval, startFieldsStringId, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &errorCode)
	binary.Read(resultBuf, binary.LittleEndian, &startMode)
	binary.Read(resultBuf, binary.LittleEndian, &continueAfterError)
	binary.Read(resultBuf, binary.LittleEndian, &startInterval)
	binary.Read(resultBuf, binary.LittleEndian, &startFieldsStringId)


	return errorCode, startMode, continueAfterError, startInterval, startFieldsStringId, nil
}

// FIXME: message is currently valid in error-occurred state only
//
// Associated constants:
//
//	* ErrorCodeSuccess
//	* ErrorCodeUnknownError
//	* ErrorCodeInvalidOperation
//	* ErrorCodeOperationAborted
//	* ErrorCodeInternalError
//	* ErrorCodeUnknownSessionId
//	* ErrorCodeNoFreeSessionId
//	* ErrorCodeUnknownObjectId
//	* ErrorCodeNoFreeObjectId
//	* ErrorCodeObjectIsLocked
//	* ErrorCodeNoMoreData
//	* ErrorCodeWrongListItemType
//	* ErrorCodeProgramIsPurged
//	* ErrorCodeInvalidParameter
//	* ErrorCodeNoFreeMemory
//	* ErrorCodeNoFreeSpace
//	* ErrorCodeAccessDenied
//	* ErrorCodeAlreadyExists
//	* ErrorCodeDoesNotExist
//	* ErrorCodeInterrupted
//	* ErrorCodeIsDirectory
//	* ErrorCodeNotADirectory
//	* ErrorCodeWouldBlock
//	* ErrorCodeOverflow
//	* ErrorCodeBadFileDescriptor
//	* ErrorCodeOutOfRange
//	* ErrorCodeNameTooLong
//	* ErrorCodeInvalidSeek
//	* ErrorCodeNotSupported
//	* ErrorCodeTooManyOpenFiles
//	* ProgramSchedulerStateStopped
//	* ProgramSchedulerStateRunning
func (device *REDBrick) GetProgramSchedulerState(programId uint16, sessionId uint16) (errorCode ErrorCode, state ProgramSchedulerState, timestamp uint64, messageStringId uint16, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, programId);
	binary.Write(&buf, binary.LittleEndian, sessionId);

	resultBytes, err := device.device.Get(uint8(FunctionGetProgramSchedulerState), buf.Bytes())
	if err != nil {
		return errorCode, state, timestamp, messageStringId, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 20 {
		return errorCode, state, timestamp, messageStringId, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 20)
	}


	if header.ErrorCode != 0 {
		return errorCode, state, timestamp, messageStringId, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &errorCode)
	binary.Read(resultBuf, binary.LittleEndian, &state)
	binary.Read(resultBuf, binary.LittleEndian, &timestamp)
	binary.Read(resultBuf, binary.LittleEndian, &messageStringId)


	return errorCode, state, timestamp, messageStringId, nil
}

// 
//
// Associated constants:
//
//	* ErrorCodeSuccess
//	* ErrorCodeUnknownError
//	* ErrorCodeInvalidOperation
//	* ErrorCodeOperationAborted
//	* ErrorCodeInternalError
//	* ErrorCodeUnknownSessionId
//	* ErrorCodeNoFreeSessionId
//	* ErrorCodeUnknownObjectId
//	* ErrorCodeNoFreeObjectId
//	* ErrorCodeObjectIsLocked
//	* ErrorCodeNoMoreData
//	* ErrorCodeWrongListItemType
//	* ErrorCodeProgramIsPurged
//	* ErrorCodeInvalidParameter
//	* ErrorCodeNoFreeMemory
//	* ErrorCodeNoFreeSpace
//	* ErrorCodeAccessDenied
//	* ErrorCodeAlreadyExists
//	* ErrorCodeDoesNotExist
//	* ErrorCodeInterrupted
//	* ErrorCodeIsDirectory
//	* ErrorCodeNotADirectory
//	* ErrorCodeWouldBlock
//	* ErrorCodeOverflow
//	* ErrorCodeBadFileDescriptor
//	* ErrorCodeOutOfRange
//	* ErrorCodeNameTooLong
//	* ErrorCodeInvalidSeek
//	* ErrorCodeNotSupported
//	* ErrorCodeTooManyOpenFiles
func (device *REDBrick) ContinueProgramSchedule(programId uint16) (errorCode ErrorCode, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, programId);

	resultBytes, err := device.device.Get(uint8(FunctionContinueProgramSchedule), buf.Bytes())
	if err != nil {
		return errorCode, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 9 {
		return errorCode, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
	}


	if header.ErrorCode != 0 {
		return errorCode, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &errorCode)


	return errorCode, nil
}

// 
//
// Associated constants:
//
//	* ErrorCodeSuccess
//	* ErrorCodeUnknownError
//	* ErrorCodeInvalidOperation
//	* ErrorCodeOperationAborted
//	* ErrorCodeInternalError
//	* ErrorCodeUnknownSessionId
//	* ErrorCodeNoFreeSessionId
//	* ErrorCodeUnknownObjectId
//	* ErrorCodeNoFreeObjectId
//	* ErrorCodeObjectIsLocked
//	* ErrorCodeNoMoreData
//	* ErrorCodeWrongListItemType
//	* ErrorCodeProgramIsPurged
//	* ErrorCodeInvalidParameter
//	* ErrorCodeNoFreeMemory
//	* ErrorCodeNoFreeSpace
//	* ErrorCodeAccessDenied
//	* ErrorCodeAlreadyExists
//	* ErrorCodeDoesNotExist
//	* ErrorCodeInterrupted
//	* ErrorCodeIsDirectory
//	* ErrorCodeNotADirectory
//	* ErrorCodeWouldBlock
//	* ErrorCodeOverflow
//	* ErrorCodeBadFileDescriptor
//	* ErrorCodeOutOfRange
//	* ErrorCodeNameTooLong
//	* ErrorCodeInvalidSeek
//	* ErrorCodeNotSupported
//	* ErrorCodeTooManyOpenFiles
func (device *REDBrick) StartProgram(programId uint16) (errorCode ErrorCode, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, programId);

	resultBytes, err := device.device.Get(uint8(FunctionStartProgram), buf.Bytes())
	if err != nil {
		return errorCode, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 9 {
		return errorCode, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
	}


	if header.ErrorCode != 0 {
		return errorCode, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &errorCode)


	return errorCode, nil
}

// 
//
// Associated constants:
//
//	* ErrorCodeSuccess
//	* ErrorCodeUnknownError
//	* ErrorCodeInvalidOperation
//	* ErrorCodeOperationAborted
//	* ErrorCodeInternalError
//	* ErrorCodeUnknownSessionId
//	* ErrorCodeNoFreeSessionId
//	* ErrorCodeUnknownObjectId
//	* ErrorCodeNoFreeObjectId
//	* ErrorCodeObjectIsLocked
//	* ErrorCodeNoMoreData
//	* ErrorCodeWrongListItemType
//	* ErrorCodeProgramIsPurged
//	* ErrorCodeInvalidParameter
//	* ErrorCodeNoFreeMemory
//	* ErrorCodeNoFreeSpace
//	* ErrorCodeAccessDenied
//	* ErrorCodeAlreadyExists
//	* ErrorCodeDoesNotExist
//	* ErrorCodeInterrupted
//	* ErrorCodeIsDirectory
//	* ErrorCodeNotADirectory
//	* ErrorCodeWouldBlock
//	* ErrorCodeOverflow
//	* ErrorCodeBadFileDescriptor
//	* ErrorCodeOutOfRange
//	* ErrorCodeNameTooLong
//	* ErrorCodeInvalidSeek
//	* ErrorCodeNotSupported
//	* ErrorCodeTooManyOpenFiles
func (device *REDBrick) GetLastSpawnedProgramProcess(programId uint16, sessionId uint16) (errorCode ErrorCode, processId uint16, timestamp uint64, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, programId);
	binary.Write(&buf, binary.LittleEndian, sessionId);

	resultBytes, err := device.device.Get(uint8(FunctionGetLastSpawnedProgramProcess), buf.Bytes())
	if err != nil {
		return errorCode, processId, timestamp, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 19 {
		return errorCode, processId, timestamp, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 19)
	}


	if header.ErrorCode != 0 {
		return errorCode, processId, timestamp, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &errorCode)
	binary.Read(resultBuf, binary.LittleEndian, &processId)
	binary.Read(resultBuf, binary.LittleEndian, &timestamp)


	return errorCode, processId, timestamp, nil
}

// 
//
// Associated constants:
//
//	* ErrorCodeSuccess
//	* ErrorCodeUnknownError
//	* ErrorCodeInvalidOperation
//	* ErrorCodeOperationAborted
//	* ErrorCodeInternalError
//	* ErrorCodeUnknownSessionId
//	* ErrorCodeNoFreeSessionId
//	* ErrorCodeUnknownObjectId
//	* ErrorCodeNoFreeObjectId
//	* ErrorCodeObjectIsLocked
//	* ErrorCodeNoMoreData
//	* ErrorCodeWrongListItemType
//	* ErrorCodeProgramIsPurged
//	* ErrorCodeInvalidParameter
//	* ErrorCodeNoFreeMemory
//	* ErrorCodeNoFreeSpace
//	* ErrorCodeAccessDenied
//	* ErrorCodeAlreadyExists
//	* ErrorCodeDoesNotExist
//	* ErrorCodeInterrupted
//	* ErrorCodeIsDirectory
//	* ErrorCodeNotADirectory
//	* ErrorCodeWouldBlock
//	* ErrorCodeOverflow
//	* ErrorCodeBadFileDescriptor
//	* ErrorCodeOutOfRange
//	* ErrorCodeNameTooLong
//	* ErrorCodeInvalidSeek
//	* ErrorCodeNotSupported
//	* ErrorCodeTooManyOpenFiles
func (device *REDBrick) GetCustomProgramOptionNames(programId uint16, sessionId uint16) (errorCode ErrorCode, namesListId uint16, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, programId);
	binary.Write(&buf, binary.LittleEndian, sessionId);

	resultBytes, err := device.device.Get(uint8(FunctionGetCustomProgramOptionNames), buf.Bytes())
	if err != nil {
		return errorCode, namesListId, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 11 {
		return errorCode, namesListId, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 11)
	}


	if header.ErrorCode != 0 {
		return errorCode, namesListId, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &errorCode)
	binary.Read(resultBuf, binary.LittleEndian, &namesListId)


	return errorCode, namesListId, nil
}

// 
//
// Associated constants:
//
//	* ErrorCodeSuccess
//	* ErrorCodeUnknownError
//	* ErrorCodeInvalidOperation
//	* ErrorCodeOperationAborted
//	* ErrorCodeInternalError
//	* ErrorCodeUnknownSessionId
//	* ErrorCodeNoFreeSessionId
//	* ErrorCodeUnknownObjectId
//	* ErrorCodeNoFreeObjectId
//	* ErrorCodeObjectIsLocked
//	* ErrorCodeNoMoreData
//	* ErrorCodeWrongListItemType
//	* ErrorCodeProgramIsPurged
//	* ErrorCodeInvalidParameter
//	* ErrorCodeNoFreeMemory
//	* ErrorCodeNoFreeSpace
//	* ErrorCodeAccessDenied
//	* ErrorCodeAlreadyExists
//	* ErrorCodeDoesNotExist
//	* ErrorCodeInterrupted
//	* ErrorCodeIsDirectory
//	* ErrorCodeNotADirectory
//	* ErrorCodeWouldBlock
//	* ErrorCodeOverflow
//	* ErrorCodeBadFileDescriptor
//	* ErrorCodeOutOfRange
//	* ErrorCodeNameTooLong
//	* ErrorCodeInvalidSeek
//	* ErrorCodeNotSupported
//	* ErrorCodeTooManyOpenFiles
func (device *REDBrick) SetCustomProgramOptionValue(programId uint16, nameStringId uint16, valueStringId uint16) (errorCode ErrorCode, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, programId);
	binary.Write(&buf, binary.LittleEndian, nameStringId);
	binary.Write(&buf, binary.LittleEndian, valueStringId);

	resultBytes, err := device.device.Get(uint8(FunctionSetCustomProgramOptionValue), buf.Bytes())
	if err != nil {
		return errorCode, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 9 {
		return errorCode, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
	}


	if header.ErrorCode != 0 {
		return errorCode, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &errorCode)


	return errorCode, nil
}

// 
//
// Associated constants:
//
//	* ErrorCodeSuccess
//	* ErrorCodeUnknownError
//	* ErrorCodeInvalidOperation
//	* ErrorCodeOperationAborted
//	* ErrorCodeInternalError
//	* ErrorCodeUnknownSessionId
//	* ErrorCodeNoFreeSessionId
//	* ErrorCodeUnknownObjectId
//	* ErrorCodeNoFreeObjectId
//	* ErrorCodeObjectIsLocked
//	* ErrorCodeNoMoreData
//	* ErrorCodeWrongListItemType
//	* ErrorCodeProgramIsPurged
//	* ErrorCodeInvalidParameter
//	* ErrorCodeNoFreeMemory
//	* ErrorCodeNoFreeSpace
//	* ErrorCodeAccessDenied
//	* ErrorCodeAlreadyExists
//	* ErrorCodeDoesNotExist
//	* ErrorCodeInterrupted
//	* ErrorCodeIsDirectory
//	* ErrorCodeNotADirectory
//	* ErrorCodeWouldBlock
//	* ErrorCodeOverflow
//	* ErrorCodeBadFileDescriptor
//	* ErrorCodeOutOfRange
//	* ErrorCodeNameTooLong
//	* ErrorCodeInvalidSeek
//	* ErrorCodeNotSupported
//	* ErrorCodeTooManyOpenFiles
func (device *REDBrick) GetCustomProgramOptionValue(programId uint16, nameStringId uint16, sessionId uint16) (errorCode ErrorCode, valueStringId uint16, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, programId);
	binary.Write(&buf, binary.LittleEndian, nameStringId);
	binary.Write(&buf, binary.LittleEndian, sessionId);

	resultBytes, err := device.device.Get(uint8(FunctionGetCustomProgramOptionValue), buf.Bytes())
	if err != nil {
		return errorCode, valueStringId, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 11 {
		return errorCode, valueStringId, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 11)
	}


	if header.ErrorCode != 0 {
		return errorCode, valueStringId, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &errorCode)
	binary.Read(resultBuf, binary.LittleEndian, &valueStringId)


	return errorCode, valueStringId, nil
}

// 
//
// Associated constants:
//
//	* ErrorCodeSuccess
//	* ErrorCodeUnknownError
//	* ErrorCodeInvalidOperation
//	* ErrorCodeOperationAborted
//	* ErrorCodeInternalError
//	* ErrorCodeUnknownSessionId
//	* ErrorCodeNoFreeSessionId
//	* ErrorCodeUnknownObjectId
//	* ErrorCodeNoFreeObjectId
//	* ErrorCodeObjectIsLocked
//	* ErrorCodeNoMoreData
//	* ErrorCodeWrongListItemType
//	* ErrorCodeProgramIsPurged
//	* ErrorCodeInvalidParameter
//	* ErrorCodeNoFreeMemory
//	* ErrorCodeNoFreeSpace
//	* ErrorCodeAccessDenied
//	* ErrorCodeAlreadyExists
//	* ErrorCodeDoesNotExist
//	* ErrorCodeInterrupted
//	* ErrorCodeIsDirectory
//	* ErrorCodeNotADirectory
//	* ErrorCodeWouldBlock
//	* ErrorCodeOverflow
//	* ErrorCodeBadFileDescriptor
//	* ErrorCodeOutOfRange
//	* ErrorCodeNameTooLong
//	* ErrorCodeInvalidSeek
//	* ErrorCodeNotSupported
//	* ErrorCodeTooManyOpenFiles
func (device *REDBrick) RemoveCustomProgramOption(programId uint16, nameStringId uint16) (errorCode ErrorCode, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, programId);
	binary.Write(&buf, binary.LittleEndian, nameStringId);

	resultBytes, err := device.device.Get(uint8(FunctionRemoveCustomProgramOption), buf.Bytes())
	if err != nil {
		return errorCode, err
	}

	var header PacketHeader
	header.FillFromBytes(resultBytes)

	if header.Length != 9 {
		return errorCode, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
	}


	if header.ErrorCode != 0 {
		return errorCode, DeviceError(header.ErrorCode)
	}

	resultBuf := bytes.NewBuffer(resultBytes[8:])
	binary.Read(resultBuf, binary.LittleEndian, &errorCode)


	return errorCode, nil
}

// Returns the UID, the UID where the Brick is connected to,
// the position, the hardware and firmware version as well as the
// device identifier.
// 
// The position is the position in the stack from '0' (bottom) to '8' (top).
// 
// The device identifier numbers can be found `here <device_identifier>`.
// |device_identifier_constant|
func (device *REDBrick) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetIdentity), buf.Bytes())
	if err != nil {
		return uid, connectedUid, position, hardwareVersion, firmwareVersion, deviceIdentifier, err
	}

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


	return uid, connectedUid, position, hardwareVersion, firmwareVersion, deviceIdentifier, nil
}
