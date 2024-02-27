/* ***********************************************************
 * This file was automatically generated on 2024-02-27.      *
 *                                                           *
 * Go Bindings Version 2.0.15                                *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/

// Silently drives one bipolar stepper motor with up to 46V and 1.6A per phase.
//
//
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/SilentStepperV2_Bricklet_Go.html.
package silent_stepper_v2_bricklet

import (
	"bytes"
	"encoding/binary"
	"fmt"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionSetMaxVelocity                 Function = 1
	FunctionGetMaxVelocity                 Function = 2
	FunctionGetCurrentVelocity             Function = 3
	FunctionSetSpeedRamping                Function = 4
	FunctionGetSpeedRamping                Function = 5
	FunctionFullBrake                      Function = 6
	FunctionSetCurrentPosition             Function = 7
	FunctionGetCurrentPosition             Function = 8
	FunctionSetTargetPosition              Function = 9
	FunctionGetTargetPosition              Function = 10
	FunctionSetSteps                       Function = 11
	FunctionGetSteps                       Function = 12
	FunctionGetRemainingSteps              Function = 13
	FunctionSetStepConfiguration           Function = 14
	FunctionGetStepConfiguration           Function = 15
	FunctionDriveForward                   Function = 16
	FunctionDriveBackward                  Function = 17
	FunctionStop                           Function = 18
	FunctionGetInputVoltage                Function = 19
	FunctionSetMotorCurrent                Function = 22
	FunctionGetMotorCurrent                Function = 23
	FunctionSetEnabled                     Function = 24
	FunctionGetEnabled                     Function = 25
	FunctionSetBasicConfiguration          Function = 26
	FunctionGetBasicConfiguration          Function = 27
	FunctionSetSpreadcycleConfiguration    Function = 28
	FunctionGetSpreadcycleConfiguration    Function = 29
	FunctionSetStealthConfiguration        Function = 30
	FunctionGetStealthConfiguration        Function = 31
	FunctionSetCoolstepConfiguration       Function = 32
	FunctionGetCoolstepConfiguration       Function = 33
	FunctionSetMiscConfiguration           Function = 34
	FunctionGetMiscConfiguration           Function = 35
	FunctionSetErrorLEDConfig              Function = 36
	FunctionGetErrorLEDConfig              Function = 37
	FunctionGetDriverStatus                Function = 38
	FunctionSetMinimumVoltage              Function = 39
	FunctionGetMinimumVoltage              Function = 40
	FunctionSetTimeBase                    Function = 43
	FunctionGetTimeBase                    Function = 44
	FunctionGetAllData                     Function = 45
	FunctionSetAllCallbackConfiguration    Function = 46
	FunctionGetAllDataCallbackConfiguraton Function = 47
	FunctionSetGPIOConfiguration           Function = 48
	FunctionGetGPIOConfiguration           Function = 49
	FunctionSetGPIOAction                  Function = 50
	FunctionGetGPIOAction                  Function = 51
	FunctionGetGPIOState                   Function = 52
	FunctionGetSPITFPErrorCount            Function = 234
	FunctionSetBootloaderMode              Function = 235
	FunctionGetBootloaderMode              Function = 236
	FunctionSetWriteFirmwarePointer        Function = 237
	FunctionWriteFirmware                  Function = 238
	FunctionSetStatusLEDConfig             Function = 239
	FunctionGetStatusLEDConfig             Function = 240
	FunctionGetChipTemperature             Function = 242
	FunctionReset                          Function = 243
	FunctionWriteUID                       Function = 248
	FunctionReadUID                        Function = 249
	FunctionGetIdentity                    Function = 255
	FunctionCallbackUnderVoltage           Function = 41
	FunctionCallbackPositionReached        Function = 42
	FunctionCallbackAllData                Function = 53
	FunctionCallbackNewState               Function = 54
	FunctionCallbackGPIOState              Function = 55
)

type StepResolution = uint8

const (
	StepResolution1   StepResolution = 8
	StepResolution2   StepResolution = 7
	StepResolution4   StepResolution = 6
	StepResolution8   StepResolution = 5
	StepResolution16  StepResolution = 4
	StepResolution32  StepResolution = 3
	StepResolution64  StepResolution = 2
	StepResolution128 StepResolution = 1
	StepResolution256 StepResolution = 0
)

type ChopperMode = uint8

const (
	ChopperModeSpreadCycle ChopperMode = 0
	ChopperModeFastDecay   ChopperMode = 1
)

type FreewheelMode = uint8

const (
	FreewheelModeNormal       FreewheelMode = 0
	FreewheelModeFreewheeling FreewheelMode = 1
	FreewheelModeCoilShortLS  FreewheelMode = 2
	FreewheelModeCoilShortHS  FreewheelMode = 3
)

type CurrentUpStepIncrement = uint8

const (
	CurrentUpStepIncrement1 CurrentUpStepIncrement = 0
	CurrentUpStepIncrement2 CurrentUpStepIncrement = 1
	CurrentUpStepIncrement4 CurrentUpStepIncrement = 2
	CurrentUpStepIncrement8 CurrentUpStepIncrement = 3
)

type CurrentDownStepDecrement = uint8

const (
	CurrentDownStepDecrement1  CurrentDownStepDecrement = 0
	CurrentDownStepDecrement2  CurrentDownStepDecrement = 1
	CurrentDownStepDecrement8  CurrentDownStepDecrement = 2
	CurrentDownStepDecrement32 CurrentDownStepDecrement = 3
)

type MinimumCurrent = uint8

const (
	MinimumCurrentHalf    MinimumCurrent = 0
	MinimumCurrentQuarter MinimumCurrent = 1
)

type StallguardMode = uint8

const (
	StallguardModeStandard StallguardMode = 0
	StallguardModeFiltered StallguardMode = 1
)

type OpenLoad = uint8

const (
	OpenLoadNone    OpenLoad = 0
	OpenLoadPhaseA  OpenLoad = 1
	OpenLoadPhaseB  OpenLoad = 2
	OpenLoadPhaseAB OpenLoad = 3
)

type ShortToGround = uint8

const (
	ShortToGroundNone    ShortToGround = 0
	ShortToGroundPhaseA  ShortToGround = 1
	ShortToGroundPhaseB  ShortToGround = 2
	ShortToGroundPhaseAB ShortToGround = 3
)

type OverTemperature = uint8

const (
	OverTemperatureNone    OverTemperature = 0
	OverTemperatureWarning OverTemperature = 1
	OverTemperatureLimit   OverTemperature = 2
)

type State = uint8

const (
	StateStop                      State = 1
	StateAcceleration              State = 2
	StateRun                       State = 3
	StateDeacceleration            State = 4
	StateDirectionChangeToForward  State = 5
	StateDirectionChangeToBackward State = 6
)

type GPIOAction = uint32

const (
	GPIOActionNone                  GPIOAction = 0
	GPIOActionNormalStopRisingEdge  GPIOAction = 1
	GPIOActionNormalStopFallingEdge GPIOAction = 2
	GPIOActionFullBrakeRisingEdge   GPIOAction = 4
	GPIOActionFullBrakeFallingEdge  GPIOAction = 8
	GPIOActionCallbackRisingEdge    GPIOAction = 16
	GPIOActionCallbackFallingEdge   GPIOAction = 32
)

type ErrorLEDConfig = uint8

const (
	ErrorLEDConfigOff           ErrorLEDConfig = 0
	ErrorLEDConfigOn            ErrorLEDConfig = 1
	ErrorLEDConfigShowHeartbeat ErrorLEDConfig = 2
	ErrorLEDConfigShowError     ErrorLEDConfig = 3
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

type SilentStepperV2Bricklet struct {
	device Device
}

const DeviceIdentifier = 2166
const DeviceDisplayName = "Silent Stepper Bricklet 2.0"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (SilentStepperV2Bricklet, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{2, 0, 0}, uid, &internalIPCon, 0, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return SilentStepperV2Bricklet{}, err
	}
	dev.ResponseExpected[FunctionSetMaxVelocity] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetMaxVelocity] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetCurrentVelocity] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetSpeedRamping] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetSpeedRamping] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionFullBrake] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionSetCurrentPosition] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetCurrentPosition] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetTargetPosition] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetTargetPosition] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetSteps] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetSteps] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetRemainingSteps] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetStepConfiguration] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetStepConfiguration] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionDriveForward] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionDriveBackward] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionStop] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetInputVoltage] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetMotorCurrent] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetMotorCurrent] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetEnabled] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetEnabled] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetBasicConfiguration] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetBasicConfiguration] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetSpreadcycleConfiguration] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetSpreadcycleConfiguration] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetStealthConfiguration] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetStealthConfiguration] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetCoolstepConfiguration] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetCoolstepConfiguration] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetMiscConfiguration] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetMiscConfiguration] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetErrorLEDConfig] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetErrorLEDConfig] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetDriverStatus] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetMinimumVoltage] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionGetMinimumVoltage] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetTimeBase] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetTimeBase] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetAllData] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetAllCallbackConfiguration] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionGetAllDataCallbackConfiguraton] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetGPIOConfiguration] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetGPIOConfiguration] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetGPIOAction] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetGPIOAction] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetGPIOState] = ResponseExpectedFlagAlwaysTrue
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
	return SilentStepperV2Bricklet{dev}, nil
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
func (device *SilentStepperV2Bricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *SilentStepperV2Bricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *SilentStepperV2Bricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *SilentStepperV2Bricklet) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// This callback is triggered when the input voltage drops below the value set by
// SetMinimumVoltage. The parameter is the current voltage.
func (device *SilentStepperV2Bricklet) RegisterUnderVoltageCallback(fn func(uint16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 10 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var voltage uint16
		binary.Read(buf, binary.LittleEndian, &voltage)
		fn(voltage)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackUnderVoltage), wrapper)
}

// Remove a registered Under Voltage callback.
func (device *SilentStepperV2Bricklet) DeregisterUnderVoltageCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackUnderVoltage), registrationId)
}

// This callback is triggered when a position set by SetSteps or
// SetTargetPosition is reached.
//
// Note
//  Since we can't get any feedback from the stepper motor, this only works if the
//  acceleration (see SetSpeedRamping) is set smaller or equal to the
//  maximum acceleration of the motor. Otherwise the motor will lag behind the
//  control value and the callback will be triggered too early.
func (device *SilentStepperV2Bricklet) RegisterPositionReachedCallback(fn func(int32)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 12 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var position int32
		binary.Read(buf, binary.LittleEndian, &position)
		fn(position)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackPositionReached), wrapper)
}

// Remove a registered Position Reached callback.
func (device *SilentStepperV2Bricklet) DeregisterPositionReachedCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackPositionReached), registrationId)
}

// This callback is triggered periodically with the period that is set by
// SetAllCallbackConfiguration. The parameters are: the current velocity,
// the current position, the remaining steps, the stack voltage, the external
// voltage and the current consumption of the stepper motor.
func (device *SilentStepperV2Bricklet) RegisterAllDataCallback(fn func(uint16, int32, int32, uint16, uint16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 22 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var currentVelocity uint16
		var currentPosition int32
		var remainingSteps int32
		var inputVoltage uint16
		var currentConsumption uint16
		binary.Read(buf, binary.LittleEndian, &currentVelocity)
		binary.Read(buf, binary.LittleEndian, &currentPosition)
		binary.Read(buf, binary.LittleEndian, &remainingSteps)
		binary.Read(buf, binary.LittleEndian, &inputVoltage)
		binary.Read(buf, binary.LittleEndian, &currentConsumption)
		fn(currentVelocity, currentPosition, remainingSteps, inputVoltage, currentConsumption)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackAllData), wrapper)
}

// Remove a registered All Data callback.
func (device *SilentStepperV2Bricklet) DeregisterAllDataCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackAllData), registrationId)
}

// This callback is triggered whenever the Silent Stepper Bricklet 2.0 enters a new state.
// It returns the new state as well as the previous state.
func (device *SilentStepperV2Bricklet) RegisterNewStateCallback(fn func(State, State)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 10 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var stateNew State
		var statePrevious State
		binary.Read(buf, binary.LittleEndian, &stateNew)
		binary.Read(buf, binary.LittleEndian, &statePrevious)
		fn(stateNew, statePrevious)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackNewState), wrapper)
}

// Remove a registered New State callback.
func (device *SilentStepperV2Bricklet) DeregisterNewStateCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackNewState), registrationId)
}

// This callback is triggered by GPIO changes if it is activated through SetGPIOAction.
func (device *SilentStepperV2Bricklet) RegisterGPIOStateCallback(fn func([2]bool)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 9 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var gpioState [2]bool
		copy(gpioState[:], ByteSliceToBoolSlice(buf.Next(1)))
		fn(gpioState)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackGPIOState), wrapper)
}

// Remove a registered GPIO State callback.
func (device *SilentStepperV2Bricklet) DeregisterGPIOStateCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackGPIOState), registrationId)
}

// Sets the maximum velocity of the stepper motor.
// This function does *not* start the motor, it merely sets the maximum
// velocity the stepper motor is accelerated to. To get the motor running use
// either SetTargetPosition, SetSteps, DriveForward or
// DriveBackward.
func (device *SilentStepperV2Bricklet) SetMaxVelocity(velocity uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, velocity)

	resultBytes, err := device.device.Set(uint8(FunctionSetMaxVelocity), buf.Bytes())
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

// Returns the velocity as set by SetMaxVelocity.
func (device *SilentStepperV2Bricklet) GetMaxVelocity() (velocity uint16, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetMaxVelocity), buf.Bytes())
	if err != nil {
		return velocity, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return velocity, DeviceError(header.ErrorCode)
		}

		if header.Length != 10 {
			return velocity, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &velocity)

	}

	return velocity, nil
}

// Returns the *current* velocity of the stepper motor.
func (device *SilentStepperV2Bricklet) GetCurrentVelocity() (velocity uint16, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetCurrentVelocity), buf.Bytes())
	if err != nil {
		return velocity, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return velocity, DeviceError(header.ErrorCode)
		}

		if header.Length != 10 {
			return velocity, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &velocity)

	}

	return velocity, nil
}

// Sets the acceleration and deacceleration of the stepper motor.
// An acceleration of 1000 means, that
// every second the velocity is increased by 1000 *steps/s*.
//
// For example: If the current velocity is 0 and you want to accelerate to a
// velocity of 8000 *steps/s* in 10 seconds, you should set an acceleration
// of 800 *steps/s²*.
//
// An acceleration/deacceleration of 0 means instantaneous
// acceleration/deacceleration (not recommended)
func (device *SilentStepperV2Bricklet) SetSpeedRamping(acceleration uint16, deacceleration uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, acceleration)
	binary.Write(&buf, binary.LittleEndian, deacceleration)

	resultBytes, err := device.device.Set(uint8(FunctionSetSpeedRamping), buf.Bytes())
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

// Returns the acceleration and deacceleration as set by
// SetSpeedRamping.
func (device *SilentStepperV2Bricklet) GetSpeedRamping() (acceleration uint16, deacceleration uint16, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetSpeedRamping), buf.Bytes())
	if err != nil {
		return acceleration, deacceleration, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return acceleration, deacceleration, DeviceError(header.ErrorCode)
		}

		if header.Length != 12 {
			return acceleration, deacceleration, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &acceleration)
		binary.Read(resultBuf, binary.LittleEndian, &deacceleration)

	}

	return acceleration, deacceleration, nil
}

// Executes an active full brake.
//
// Warning
//  This function is for emergency purposes,
//  where an immediate brake is necessary. Depending on the current velocity and
//  the strength of the motor, a full brake can be quite violent.
//
// Call Stop if you just want to stop the motor.
func (device *SilentStepperV2Bricklet) FullBrake() (err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Set(uint8(FunctionFullBrake), buf.Bytes())
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

// Sets the current steps of the internal step counter. This can be used to
// set the current position to 0 when some kind of starting position
// is reached (e.g. when a CNC machine reaches a corner).
func (device *SilentStepperV2Bricklet) SetCurrentPosition(position int32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, position)

	resultBytes, err := device.device.Set(uint8(FunctionSetCurrentPosition), buf.Bytes())
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

// Returns the current position of the stepper motor in steps. On startup
// the position is 0. The steps are counted with all possible driving
// functions (SetTargetPosition, SetSteps, DriveForward or
// DriveBackward). It also is possible to reset the steps to 0 or
// set them to any other desired value with SetCurrentPosition.
func (device *SilentStepperV2Bricklet) GetCurrentPosition() (position int32, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetCurrentPosition), buf.Bytes())
	if err != nil {
		return position, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return position, DeviceError(header.ErrorCode)
		}

		if header.Length != 12 {
			return position, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &position)

	}

	return position, nil
}

// Sets the target position of the stepper motor in steps. For example,
// if the current position of the motor is 500 and SetTargetPosition is
// called with 1000, the stepper motor will drive 500 steps forward. It will
// use the velocity, acceleration and deacceleration as set by
// SetMaxVelocity and SetSpeedRamping.
//
// A call of SetTargetPosition with the parameter *x* is equivalent to
// a call of SetSteps with the parameter
// (*x* - GetCurrentPosition).
func (device *SilentStepperV2Bricklet) SetTargetPosition(position int32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, position)

	resultBytes, err := device.device.Set(uint8(FunctionSetTargetPosition), buf.Bytes())
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

// Returns the last target position as set by SetTargetPosition.
func (device *SilentStepperV2Bricklet) GetTargetPosition() (position int32, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetTargetPosition), buf.Bytes())
	if err != nil {
		return position, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return position, DeviceError(header.ErrorCode)
		}

		if header.Length != 12 {
			return position, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &position)

	}

	return position, nil
}

// Sets the number of steps the stepper motor should run. Positive values
// will drive the motor forward and negative values backward.
// The velocity, acceleration and deacceleration as set by
// SetMaxVelocity and SetSpeedRamping will be used.
func (device *SilentStepperV2Bricklet) SetSteps(steps int32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, steps)

	resultBytes, err := device.device.Set(uint8(FunctionSetSteps), buf.Bytes())
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

// Returns the last steps as set by SetSteps.
func (device *SilentStepperV2Bricklet) GetSteps() (steps int32, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetSteps), buf.Bytes())
	if err != nil {
		return steps, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return steps, DeviceError(header.ErrorCode)
		}

		if header.Length != 12 {
			return steps, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &steps)

	}

	return steps, nil
}

// Returns the remaining steps of the last call of SetSteps.
// For example, if SetSteps is called with 2000 and
// GetRemainingSteps is called after the motor has run for 500 steps,
// it will return 1500.
func (device *SilentStepperV2Bricklet) GetRemainingSteps() (steps int32, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetRemainingSteps), buf.Bytes())
	if err != nil {
		return steps, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return steps, DeviceError(header.ErrorCode)
		}

		if header.Length != 12 {
			return steps, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &steps)

	}

	return steps, nil
}

// Sets the step resolution from full-step up to 1/256-step.
//
// If interpolation is turned on, the Silent Stepper Bricklet 2.0 will always interpolate
// your step inputs as 1/256-step. If you use full-step mode with interpolation, each
// step will generate 256 1/256 steps.
//
// For maximum torque use full-step without interpolation. For maximum resolution use
// 1/256-step. Turn interpolation on to make the Stepper driving less noisy.
//
// If you often change the speed with high acceleration you should turn the
// interpolation off.
//
// Associated constants:
//
//	* StepResolution1
//	* StepResolution2
//	* StepResolution4
//	* StepResolution8
//	* StepResolution16
//	* StepResolution32
//	* StepResolution64
//	* StepResolution128
//	* StepResolution256
func (device *SilentStepperV2Bricklet) SetStepConfiguration(stepResolution StepResolution, interpolation bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, stepResolution)
	binary.Write(&buf, binary.LittleEndian, interpolation)

	resultBytes, err := device.device.Set(uint8(FunctionSetStepConfiguration), buf.Bytes())
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

// Returns the step mode as set by SetStepConfiguration.
//
// Associated constants:
//
//	* StepResolution1
//	* StepResolution2
//	* StepResolution4
//	* StepResolution8
//	* StepResolution16
//	* StepResolution32
//	* StepResolution64
//	* StepResolution128
//	* StepResolution256
func (device *SilentStepperV2Bricklet) GetStepConfiguration() (stepResolution StepResolution, interpolation bool, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetStepConfiguration), buf.Bytes())
	if err != nil {
		return stepResolution, interpolation, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return stepResolution, interpolation, DeviceError(header.ErrorCode)
		}

		if header.Length != 10 {
			return stepResolution, interpolation, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &stepResolution)
		binary.Read(resultBuf, binary.LittleEndian, &interpolation)

	}

	return stepResolution, interpolation, nil
}

// Drives the stepper motor forward until DriveBackward or
// Stop is called. The velocity, acceleration and deacceleration as
// set by SetMaxVelocity and SetSpeedRamping will be used.
func (device *SilentStepperV2Bricklet) DriveForward() (err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Set(uint8(FunctionDriveForward), buf.Bytes())
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

// Drives the stepper motor backward until DriveForward or
// Stop is triggered. The velocity, acceleration and deacceleration as
// set by SetMaxVelocity and SetSpeedRamping will be used.
func (device *SilentStepperV2Bricklet) DriveBackward() (err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Set(uint8(FunctionDriveBackward), buf.Bytes())
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

// Stops the stepper motor with the deacceleration as set by
// SetSpeedRamping.
func (device *SilentStepperV2Bricklet) Stop() (err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Set(uint8(FunctionStop), buf.Bytes())
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

// Returns the external input voltage. The external input voltage is
// given via the black power input connector on the Silent Stepper Bricklet 2.0.
//
// If there is an external input voltage and a stack input voltage, the motor
// will be driven by the external input voltage. If there is only a stack
// voltage present, the motor will be driven by this voltage.
//
// Warning
//  This means, if you have a high stack voltage and a low external voltage,
//  the motor will be driven with the low external voltage. If you then remove
//  the external connection, it will immediately be driven by the high
//  stack voltage
func (device *SilentStepperV2Bricklet) GetInputVoltage() (voltage uint16, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetInputVoltage), buf.Bytes())
	if err != nil {
		return voltage, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return voltage, DeviceError(header.ErrorCode)
		}

		if header.Length != 10 {
			return voltage, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &voltage)

	}

	return voltage, nil
}

// Sets the current with which the motor will be driven.
//
// Warning
//  Do not set this value above the specifications of your stepper motor.
//  Otherwise it may damage your motor.
func (device *SilentStepperV2Bricklet) SetMotorCurrent(current uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, current)

	resultBytes, err := device.device.Set(uint8(FunctionSetMotorCurrent), buf.Bytes())
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

// Returns the current as set by SetMotorCurrent.
func (device *SilentStepperV2Bricklet) GetMotorCurrent() (current uint16, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetMotorCurrent), buf.Bytes())
	if err != nil {
		return current, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return current, DeviceError(header.ErrorCode)
		}

		if header.Length != 10 {
			return current, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &current)

	}

	return current, nil
}

// Enables/Disables the driver chip. The driver parameters can be configured (maximum velocity,
// acceleration, etc) before it is enabled.
//
// Warning
//  Disabling the driver chip while the motor is still turning can damage the
//  driver chip. The motor should be stopped calling Stop function
//  before disabling the motor power. The Stop function will **not**
//  wait until the motor is actually stopped. You have to explicitly wait for the
//  appropriate time after calling the Stop function before calling
//  the SetEnabled with false function.
func (device *SilentStepperV2Bricklet) SetEnabled(enabled bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, enabled)

	resultBytes, err := device.device.Set(uint8(FunctionSetEnabled), buf.Bytes())
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

// Returns *true* if the stepper driver is enabled, *false* otherwise.
func (device *SilentStepperV2Bricklet) GetEnabled() (enabled bool, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetEnabled), buf.Bytes())
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

// Sets the basic configuration parameters for the different modes (Stealth, Coolstep, Classic).
//
// * Standstill Current: This value can be used to lower the current during stand still. This might
//   be reasonable to reduce the heating of the motor and the Bricklet 2.0. When the motor is in standstill
//   the configured motor phase current will be driven until the configured
//   Power Down Time is elapsed. After that the phase current will be reduced to the standstill
//   current. The elapsed time for this reduction can be configured with the Standstill Delay Time.
//   The maximum allowed value is the configured maximum motor current
//   (see SetMotorCurrent).
//
// * Motor Run Current: The value sets the motor current when the motor is running.
//   Use a value of at least one half of the global maximum motor current for a good
//   microstep performance. The maximum allowed value is the current
//   motor current. The API maps the entered value to 1/32 ... 32/32 of the maximum
//   motor current. This value should be used to change the motor current during motor movement,
//   whereas the global maximum motor current should not be changed while the motor is moving
//   (see SetMotorCurrent).
//
// * Standstill Delay Time: Controls the duration for motor power down after a motion
//   as soon as standstill is detected and the Power Down Time is expired. A high Standstill Delay
//   Time results in a smooth transition that avoids motor jerk during power down.
//
// * Power Down Time: Sets the delay time after a stand still.
//
// * Stealth Threshold: Sets the upper threshold for Stealth mode.
//   If the velocity of the motor goes above this value, Stealth mode is turned
//   off. Otherwise it is turned on. In Stealth mode the torque declines with high speed.
//
// * Coolstep Threshold: Sets the lower threshold for Coolstep mode.
//   The Coolstep Threshold needs to be above the Stealth Threshold.
//
// * Classic Threshold: Sets the lower threshold for classic mode.
//   In classic mode the stepper becomes more noisy, but the torque is maximized.
//
// * High Velocity Chopper Mode: If High Velocity Chopper Mode is enabled, the stepper control
//   is optimized to run the stepper motors at high velocities.
//
// If you want to use all three thresholds make sure that
// Stealth Threshold < Coolstep Threshold < Classic Threshold.
func (device *SilentStepperV2Bricklet) SetBasicConfiguration(standstillCurrent uint16, motorRunCurrent uint16, standstillDelayTime uint16, powerDownTime uint16, stealthThreshold uint16, coolstepThreshold uint16, classicThreshold uint16, highVelocityChopperMode bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, standstillCurrent)
	binary.Write(&buf, binary.LittleEndian, motorRunCurrent)
	binary.Write(&buf, binary.LittleEndian, standstillDelayTime)
	binary.Write(&buf, binary.LittleEndian, powerDownTime)
	binary.Write(&buf, binary.LittleEndian, stealthThreshold)
	binary.Write(&buf, binary.LittleEndian, coolstepThreshold)
	binary.Write(&buf, binary.LittleEndian, classicThreshold)
	binary.Write(&buf, binary.LittleEndian, highVelocityChopperMode)

	resultBytes, err := device.device.Set(uint8(FunctionSetBasicConfiguration), buf.Bytes())
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

// Returns the configuration as set by SetBasicConfiguration.
func (device *SilentStepperV2Bricklet) GetBasicConfiguration() (standstillCurrent uint16, motorRunCurrent uint16, standstillDelayTime uint16, powerDownTime uint16, stealthThreshold uint16, coolstepThreshold uint16, classicThreshold uint16, highVelocityChopperMode bool, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetBasicConfiguration), buf.Bytes())
	if err != nil {
		return standstillCurrent, motorRunCurrent, standstillDelayTime, powerDownTime, stealthThreshold, coolstepThreshold, classicThreshold, highVelocityChopperMode, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return standstillCurrent, motorRunCurrent, standstillDelayTime, powerDownTime, stealthThreshold, coolstepThreshold, classicThreshold, highVelocityChopperMode, DeviceError(header.ErrorCode)
		}

		if header.Length != 23 {
			return standstillCurrent, motorRunCurrent, standstillDelayTime, powerDownTime, stealthThreshold, coolstepThreshold, classicThreshold, highVelocityChopperMode, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 23)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &standstillCurrent)
		binary.Read(resultBuf, binary.LittleEndian, &motorRunCurrent)
		binary.Read(resultBuf, binary.LittleEndian, &standstillDelayTime)
		binary.Read(resultBuf, binary.LittleEndian, &powerDownTime)
		binary.Read(resultBuf, binary.LittleEndian, &stealthThreshold)
		binary.Read(resultBuf, binary.LittleEndian, &coolstepThreshold)
		binary.Read(resultBuf, binary.LittleEndian, &classicThreshold)
		binary.Read(resultBuf, binary.LittleEndian, &highVelocityChopperMode)

	}

	return standstillCurrent, motorRunCurrent, standstillDelayTime, powerDownTime, stealthThreshold, coolstepThreshold, classicThreshold, highVelocityChopperMode, nil
}

// Note: If you don't know what any of this means you can very likely keep all of
// the values as default!
//
// Sets the Spreadcycle configuration parameters. Spreadcycle is a chopper algorithm which actively
// controls the motor current flow. More information can be found in the TMC2130 datasheet on page
// 47 (7 spreadCycle and Classic Chopper).
//
// * Slow Decay Duration: Controls duration of off time setting of slow decay phase.
//   0 = driver disabled, all bridges off. Use 1 only with Comparator Blank time >= 2.
//
// * Enable Random Slow Decay: Set to false to fix chopper off time as set by Slow Decay Duration.
//   If you set it to true, Decay Duration is randomly modulated.
//
// * Fast Decay Duration: Sets the fast decay duration. This parameters is
//   only used if the Chopper Mode is set to Fast Decay.
//
// * Hysteresis Start Value: Sets the hysteresis start value. This parameter is
//   only used if the Chopper Mode is set to Spread Cycle.
//
// * Hysteresis End Value: Sets the hysteresis end value. This parameter is
//   only used if the Chopper Mode is set to Spread Cycle.
//
// * Sine Wave Offset: Sets the sine wave offset. This parameters is
//   only used if the Chopper Mode is set to Fast Decay. 1/512 of the value becomes added to the absolute
//   value of the sine wave.
//
// * Chopper Mode: 0 = Spread Cycle, 1 = Fast Decay.
//
// * Comparator Blank Time: Sets the blank time of the comparator. Available values are
//
//   * 0 = 16 clocks,
//   * 1 = 24 clocks,
//   * 2 = 36 clocks and
//   * 3 = 54 clocks.
//
//   A value of 1 or 2 is recommended for most applications.
//
// * Fast Decay Without Comparator: If set to true the current comparator usage for termination of the
//   fast decay cycle is disabled.
//
// Associated constants:
//
//	* ChopperModeSpreadCycle
//	* ChopperModeFastDecay
func (device *SilentStepperV2Bricklet) SetSpreadcycleConfiguration(slowDecayDuration uint8, enableRandomSlowDecay bool, fastDecayDuration uint8, hysteresisStartValue uint8, hysteresisEndValue int8, sineWaveOffset int8, chopperMode ChopperMode, comparatorBlankTime uint8, fastDecayWithoutComparator bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, slowDecayDuration)
	binary.Write(&buf, binary.LittleEndian, enableRandomSlowDecay)
	binary.Write(&buf, binary.LittleEndian, fastDecayDuration)
	binary.Write(&buf, binary.LittleEndian, hysteresisStartValue)
	binary.Write(&buf, binary.LittleEndian, hysteresisEndValue)
	binary.Write(&buf, binary.LittleEndian, sineWaveOffset)
	binary.Write(&buf, binary.LittleEndian, chopperMode)
	binary.Write(&buf, binary.LittleEndian, comparatorBlankTime)
	binary.Write(&buf, binary.LittleEndian, fastDecayWithoutComparator)

	resultBytes, err := device.device.Set(uint8(FunctionSetSpreadcycleConfiguration), buf.Bytes())
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

// Returns the configuration as set by SetBasicConfiguration.
//
// Associated constants:
//
//	* ChopperModeSpreadCycle
//	* ChopperModeFastDecay
func (device *SilentStepperV2Bricklet) GetSpreadcycleConfiguration() (slowDecayDuration uint8, enableRandomSlowDecay bool, fastDecayDuration uint8, hysteresisStartValue uint8, hysteresisEndValue int8, sineWaveOffset int8, chopperMode ChopperMode, comparatorBlankTime uint8, fastDecayWithoutComparator bool, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetSpreadcycleConfiguration), buf.Bytes())
	if err != nil {
		return slowDecayDuration, enableRandomSlowDecay, fastDecayDuration, hysteresisStartValue, hysteresisEndValue, sineWaveOffset, chopperMode, comparatorBlankTime, fastDecayWithoutComparator, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return slowDecayDuration, enableRandomSlowDecay, fastDecayDuration, hysteresisStartValue, hysteresisEndValue, sineWaveOffset, chopperMode, comparatorBlankTime, fastDecayWithoutComparator, DeviceError(header.ErrorCode)
		}

		if header.Length != 17 {
			return slowDecayDuration, enableRandomSlowDecay, fastDecayDuration, hysteresisStartValue, hysteresisEndValue, sineWaveOffset, chopperMode, comparatorBlankTime, fastDecayWithoutComparator, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 17)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &slowDecayDuration)
		binary.Read(resultBuf, binary.LittleEndian, &enableRandomSlowDecay)
		binary.Read(resultBuf, binary.LittleEndian, &fastDecayDuration)
		binary.Read(resultBuf, binary.LittleEndian, &hysteresisStartValue)
		binary.Read(resultBuf, binary.LittleEndian, &hysteresisEndValue)
		binary.Read(resultBuf, binary.LittleEndian, &sineWaveOffset)
		binary.Read(resultBuf, binary.LittleEndian, &chopperMode)
		binary.Read(resultBuf, binary.LittleEndian, &comparatorBlankTime)
		binary.Read(resultBuf, binary.LittleEndian, &fastDecayWithoutComparator)

	}

	return slowDecayDuration, enableRandomSlowDecay, fastDecayDuration, hysteresisStartValue, hysteresisEndValue, sineWaveOffset, chopperMode, comparatorBlankTime, fastDecayWithoutComparator, nil
}

// Note: If you don't know what any of this means you can very likely keep all of
// the values as default!
//
// Sets the configuration relevant for Stealth mode.
//
// * Enable Stealth: If set to true the stealth mode is enabled, if set to false the
//   stealth mode is disabled, even if the speed is below the threshold set in SetBasicConfiguration.
//
// * Amplitude: If autoscale is disabled, the PWM amplitude is scaled by this value. If autoscale is enabled,
//   this value defines the maximum PWM amplitude change per half wave.
//
// * Gradient: If autoscale is disabled, the PWM gradient is scaled by this value. If autoscale is enabled,
//   this value defines the maximum PWM gradient. With autoscale a value above 64 is recommended,
//   otherwise the regulation might not be able to measure the current.
//
// * Enable Autoscale: If set to true, automatic current control is used. Otherwise the user defined
//   amplitude and gradient are used.
//
// * Force Symmetric: If true, A symmetric PWM cycle is enforced. Otherwise the PWM value may change within each
//   PWM cycle.
//
// * Freewheel Mode: The freewheel mode defines the behavior in stand still if the Standstill Current
//   (see SetBasicConfiguration) is set to 0.
//
// Associated constants:
//
//	* FreewheelModeNormal
//	* FreewheelModeFreewheeling
//	* FreewheelModeCoilShortLS
//	* FreewheelModeCoilShortHS
func (device *SilentStepperV2Bricklet) SetStealthConfiguration(enableStealth bool, amplitude uint8, gradient uint8, enableAutoscale bool, forceSymmetric bool, freewheelMode FreewheelMode) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, enableStealth)
	binary.Write(&buf, binary.LittleEndian, amplitude)
	binary.Write(&buf, binary.LittleEndian, gradient)
	binary.Write(&buf, binary.LittleEndian, enableAutoscale)
	binary.Write(&buf, binary.LittleEndian, forceSymmetric)
	binary.Write(&buf, binary.LittleEndian, freewheelMode)

	resultBytes, err := device.device.Set(uint8(FunctionSetStealthConfiguration), buf.Bytes())
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

// Returns the configuration as set by SetStealthConfiguration.
//
// Associated constants:
//
//	* FreewheelModeNormal
//	* FreewheelModeFreewheeling
//	* FreewheelModeCoilShortLS
//	* FreewheelModeCoilShortHS
func (device *SilentStepperV2Bricklet) GetStealthConfiguration() (enableStealth bool, amplitude uint8, gradient uint8, enableAutoscale bool, forceSymmetric bool, freewheelMode FreewheelMode, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetStealthConfiguration), buf.Bytes())
	if err != nil {
		return enableStealth, amplitude, gradient, enableAutoscale, forceSymmetric, freewheelMode, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return enableStealth, amplitude, gradient, enableAutoscale, forceSymmetric, freewheelMode, DeviceError(header.ErrorCode)
		}

		if header.Length != 14 {
			return enableStealth, amplitude, gradient, enableAutoscale, forceSymmetric, freewheelMode, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 14)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &enableStealth)
		binary.Read(resultBuf, binary.LittleEndian, &amplitude)
		binary.Read(resultBuf, binary.LittleEndian, &gradient)
		binary.Read(resultBuf, binary.LittleEndian, &enableAutoscale)
		binary.Read(resultBuf, binary.LittleEndian, &forceSymmetric)
		binary.Read(resultBuf, binary.LittleEndian, &freewheelMode)

	}

	return enableStealth, amplitude, gradient, enableAutoscale, forceSymmetric, freewheelMode, nil
}

// Note: If you don't know what any of this means you can very likely keep all of
// the values as default!
//
// Sets the configuration relevant for Coolstep.
//
// * Minimum Stallguard Value: If the Stallguard result falls below this value*32, the motor current
//   is increased to reduce motor load angle. A value of 0 turns Coolstep off.
//
// * Maximum Stallguard Value: If the Stallguard result goes above
//   (Min Stallguard Value + Max Stallguard Value + 1) * 32, the motor current is decreased to save
//   energy.
//
// * Current Up Step Width: Sets the up step increment per Stallguard value. The value range is 0-3,
//   corresponding to the increments 1, 2, 4 and 8.
//
// * Current Down Step Width: Sets the down step decrement per Stallguard value. The value range is 0-3,
//   corresponding to the decrements 1, 2, 8 and 16.
//
// * Minimum Current: Sets the minimum current for Coolstep current control. You can choose between
//   half and quarter of the run current.
//
// * Stallguard Threshold Value: Sets the level for stall output (see GetDriverStatus).
//   A lower value gives a higher sensitivity. You have to find a suitable value for your
//   motor by trial and error, 0 works for most motors.
//
// * Stallguard Mode: Set to 0 for standard resolution or 1 for filtered mode. In filtered mode the Stallguard
//   signal will be updated every four full-steps.
//
// Associated constants:
//
//	* CurrentUpStepIncrement1
//	* CurrentUpStepIncrement2
//	* CurrentUpStepIncrement4
//	* CurrentUpStepIncrement8
//	* CurrentDownStepDecrement1
//	* CurrentDownStepDecrement2
//	* CurrentDownStepDecrement8
//	* CurrentDownStepDecrement32
//	* MinimumCurrentHalf
//	* MinimumCurrentQuarter
//	* StallguardModeStandard
//	* StallguardModeFiltered
func (device *SilentStepperV2Bricklet) SetCoolstepConfiguration(minimumStallguardValue uint8, maximumStallguardValue uint8, currentUpStepWidth CurrentUpStepIncrement, currentDownStepWidth CurrentDownStepDecrement, minimumCurrent MinimumCurrent, stallguardThresholdValue int8, stallguardMode StallguardMode) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, minimumStallguardValue)
	binary.Write(&buf, binary.LittleEndian, maximumStallguardValue)
	binary.Write(&buf, binary.LittleEndian, currentUpStepWidth)
	binary.Write(&buf, binary.LittleEndian, currentDownStepWidth)
	binary.Write(&buf, binary.LittleEndian, minimumCurrent)
	binary.Write(&buf, binary.LittleEndian, stallguardThresholdValue)
	binary.Write(&buf, binary.LittleEndian, stallguardMode)

	resultBytes, err := device.device.Set(uint8(FunctionSetCoolstepConfiguration), buf.Bytes())
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

// Returns the configuration as set by SetCoolstepConfiguration.
//
// Associated constants:
//
//	* CurrentUpStepIncrement1
//	* CurrentUpStepIncrement2
//	* CurrentUpStepIncrement4
//	* CurrentUpStepIncrement8
//	* CurrentDownStepDecrement1
//	* CurrentDownStepDecrement2
//	* CurrentDownStepDecrement8
//	* CurrentDownStepDecrement32
//	* MinimumCurrentHalf
//	* MinimumCurrentQuarter
//	* StallguardModeStandard
//	* StallguardModeFiltered
func (device *SilentStepperV2Bricklet) GetCoolstepConfiguration() (minimumStallguardValue uint8, maximumStallguardValue uint8, currentUpStepWidth CurrentUpStepIncrement, currentDownStepWidth CurrentDownStepDecrement, minimumCurrent MinimumCurrent, stallguardThresholdValue int8, stallguardMode StallguardMode, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetCoolstepConfiguration), buf.Bytes())
	if err != nil {
		return minimumStallguardValue, maximumStallguardValue, currentUpStepWidth, currentDownStepWidth, minimumCurrent, stallguardThresholdValue, stallguardMode, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return minimumStallguardValue, maximumStallguardValue, currentUpStepWidth, currentDownStepWidth, minimumCurrent, stallguardThresholdValue, stallguardMode, DeviceError(header.ErrorCode)
		}

		if header.Length != 15 {
			return minimumStallguardValue, maximumStallguardValue, currentUpStepWidth, currentDownStepWidth, minimumCurrent, stallguardThresholdValue, stallguardMode, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 15)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &minimumStallguardValue)
		binary.Read(resultBuf, binary.LittleEndian, &maximumStallguardValue)
		binary.Read(resultBuf, binary.LittleEndian, &currentUpStepWidth)
		binary.Read(resultBuf, binary.LittleEndian, &currentDownStepWidth)
		binary.Read(resultBuf, binary.LittleEndian, &minimumCurrent)
		binary.Read(resultBuf, binary.LittleEndian, &stallguardThresholdValue)
		binary.Read(resultBuf, binary.LittleEndian, &stallguardMode)

	}

	return minimumStallguardValue, maximumStallguardValue, currentUpStepWidth, currentDownStepWidth, minimumCurrent, stallguardThresholdValue, stallguardMode, nil
}

// Note: If you don't know what any of this means you can very likely keep all of
// the values as default!
//
// Sets miscellaneous configuration parameters.
//
// * Disable Short To Ground Protection: Set to false to enable short to ground protection, otherwise
//   it is disabled.
//
// * Synchronize Phase Frequency: With this parameter you can synchronize the chopper for both phases
//   of a two phase motor to avoid the occurrence of a beat. The value range is 0-15. If set to 0,
//   the synchronization is turned off. Otherwise the synchronization is done through the formula
//   f_sync = f_clk/(value*64). In Classic Mode the synchronization is automatically switched off.
//   f_clk is 12.8MHz.
func (device *SilentStepperV2Bricklet) SetMiscConfiguration(disableShortToGroundProtection bool, synchronizePhaseFrequency uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, disableShortToGroundProtection)
	binary.Write(&buf, binary.LittleEndian, synchronizePhaseFrequency)

	resultBytes, err := device.device.Set(uint8(FunctionSetMiscConfiguration), buf.Bytes())
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

// Returns the configuration as set by SetMiscConfiguration.
func (device *SilentStepperV2Bricklet) GetMiscConfiguration() (disableShortToGroundProtection bool, synchronizePhaseFrequency uint8, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetMiscConfiguration), buf.Bytes())
	if err != nil {
		return disableShortToGroundProtection, synchronizePhaseFrequency, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return disableShortToGroundProtection, synchronizePhaseFrequency, DeviceError(header.ErrorCode)
		}

		if header.Length != 10 {
			return disableShortToGroundProtection, synchronizePhaseFrequency, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &disableShortToGroundProtection)
		binary.Read(resultBuf, binary.LittleEndian, &synchronizePhaseFrequency)

	}

	return disableShortToGroundProtection, synchronizePhaseFrequency, nil
}

// Configures the error LED to be either turned off, turned on, blink in
// heartbeat mode or show an error.
//
// If the LED is configured to show errors it has three different states:
//
// * Off: No error present.
// * 250ms interval blink: Overtemperature warning.
// * 1s interval blink: Input voltage too small.
// * full red: motor disabled because of short to ground in phase a or b or because of overtemperature.
//
// Associated constants:
//
//	* ErrorLEDConfigOff
//	* ErrorLEDConfigOn
//	* ErrorLEDConfigShowHeartbeat
//	* ErrorLEDConfigShowError
func (device *SilentStepperV2Bricklet) SetErrorLEDConfig(config ErrorLEDConfig) (err error) {
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

// Returns the LED configuration as set by SetErrorLEDConfig
//
// Associated constants:
//
//	* ErrorLEDConfigOff
//	* ErrorLEDConfigOn
//	* ErrorLEDConfigShowHeartbeat
//	* ErrorLEDConfigShowError
func (device *SilentStepperV2Bricklet) GetErrorLEDConfig() (config ErrorLEDConfig, err error) {
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

// Returns the current driver status.
//
// * Open Load: Indicates if an open load is present on phase A, B or both. This could mean that there is a problem
//   with the wiring of the motor. False detection can occur in fast motion as well as during stand still.
//
// * Short To Ground: Indicates if a short to ground is present on phase A, B or both. If this is detected the driver
//   automatically becomes disabled and stays disabled until it is enabled again manually.
//
// * Over Temperature: The over temperature indicator switches to Warning if the driver IC warms up. The warning flag
//   is expected during long duration stepper uses. If the temperature limit is reached the indicator switches
//   to Limit. In this case the driver becomes disabled until it cools down again.
//
// * Motor Stalled: Is true if a motor stall was detected.
//
// * Actual Motor Current: Indicates the actual current control scaling as used in Coolstep mode.
//   It represents a multiplier of 1/32 to 32/32 of the
//   ``Motor Run Current`` as set by SetBasicConfiguration. Example: If a ``Motor Run Current``
//   of 1000mA was set and the returned value is 15, the ``Actual Motor Current`` is 16/32*1000mA = 500mA.
//
// * Stallguard Result: Indicates the load of the motor. A lower value signals a higher load. Per trial and error
//   you can find out which value corresponds to a suitable torque for the velocity used in your application.
//   After that you can use this threshold value to find out if a motor stall becomes probable and react on it (e.g.
//   decrease velocity).
//   During stand still this value can not be used for stall detection, it shows the chopper on-time for motor coil A.
//
// * Stealth Voltage Amplitude: Shows the actual PWM scaling. In Stealth mode it can be used to detect motor load and
//   stall if autoscale is enabled (see SetStealthConfiguration).
//
// Associated constants:
//
//	* OpenLoadNone
//	* OpenLoadPhaseA
//	* OpenLoadPhaseB
//	* OpenLoadPhaseAB
//	* ShortToGroundNone
//	* ShortToGroundPhaseA
//	* ShortToGroundPhaseB
//	* ShortToGroundPhaseAB
//	* OverTemperatureNone
//	* OverTemperatureWarning
//	* OverTemperatureLimit
func (device *SilentStepperV2Bricklet) GetDriverStatus() (openLoad OpenLoad, shortToGround ShortToGround, overTemperature OverTemperature, motorStalled bool, actualMotorCurrent uint8, fullStepActive bool, stallguardResult uint8, stealthVoltageAmplitude uint8, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetDriverStatus), buf.Bytes())
	if err != nil {
		return openLoad, shortToGround, overTemperature, motorStalled, actualMotorCurrent, fullStepActive, stallguardResult, stealthVoltageAmplitude, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return openLoad, shortToGround, overTemperature, motorStalled, actualMotorCurrent, fullStepActive, stallguardResult, stealthVoltageAmplitude, DeviceError(header.ErrorCode)
		}

		if header.Length != 16 {
			return openLoad, shortToGround, overTemperature, motorStalled, actualMotorCurrent, fullStepActive, stallguardResult, stealthVoltageAmplitude, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 16)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &openLoad)
		binary.Read(resultBuf, binary.LittleEndian, &shortToGround)
		binary.Read(resultBuf, binary.LittleEndian, &overTemperature)
		binary.Read(resultBuf, binary.LittleEndian, &motorStalled)
		binary.Read(resultBuf, binary.LittleEndian, &actualMotorCurrent)
		binary.Read(resultBuf, binary.LittleEndian, &fullStepActive)
		binary.Read(resultBuf, binary.LittleEndian, &stallguardResult)
		binary.Read(resultBuf, binary.LittleEndian, &stealthVoltageAmplitude)

	}

	return openLoad, shortToGround, overTemperature, motorStalled, actualMotorCurrent, fullStepActive, stallguardResult, stealthVoltageAmplitude, nil
}

// Sets the minimum voltage, below which the RegisterUnderVoltageCallback callback
// is triggered. The minimum possible value that works with the Silent Stepper
// Bricklet 2.0 is 8V.
// You can use this function to detect the discharge of a battery that is used
// to drive the stepper motor. If you have a fixed power supply, you likely do
// not need this functionality.
func (device *SilentStepperV2Bricklet) SetMinimumVoltage(voltage uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, voltage)

	resultBytes, err := device.device.Set(uint8(FunctionSetMinimumVoltage), buf.Bytes())
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

// Returns the minimum voltage as set by SetMinimumVoltage.
func (device *SilentStepperV2Bricklet) GetMinimumVoltage() (voltage uint16, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetMinimumVoltage), buf.Bytes())
	if err != nil {
		return voltage, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return voltage, DeviceError(header.ErrorCode)
		}

		if header.Length != 10 {
			return voltage, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &voltage)

	}

	return voltage, nil
}

// Sets the time base of the velocity and the acceleration of the Silent Stepper
// Bricklet 2.0.
//
// For example, if you want to make one step every 1.5 seconds, you can set
// the time base to 15 and the velocity to 10. Now the velocity is
// 10steps/15s = 1steps/1.5s.
func (device *SilentStepperV2Bricklet) SetTimeBase(timeBase uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, timeBase)

	resultBytes, err := device.device.Set(uint8(FunctionSetTimeBase), buf.Bytes())
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

// Returns the time base as set by SetTimeBase.
func (device *SilentStepperV2Bricklet) GetTimeBase() (timeBase uint32, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetTimeBase), buf.Bytes())
	if err != nil {
		return timeBase, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return timeBase, DeviceError(header.ErrorCode)
		}

		if header.Length != 12 {
			return timeBase, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &timeBase)

	}

	return timeBase, nil
}

// Returns the following parameters: The current velocity,
// the current position, the remaining steps, the stack voltage, the external
// voltage and the current consumption of the stepper motor.
//
// The current consumption is calculated by multiplying the ``Actual Motor Current``
// value (see SetBasicConfiguration) with the ``Motor Run Current``
// (see GetDriverStatus). This is an internal calculation of the
// driver, not an independent external measurement.
//
// The current consumption calculation was broken up to firmware 2.0.1, it is fixed
// since firmware 2.0.2.
//
// There is also a callback for this function, see RegisterAllDataCallback callback.
func (device *SilentStepperV2Bricklet) GetAllData() (currentVelocity uint16, currentPosition int32, remainingSteps int32, inputVoltage uint16, currentConsumption uint16, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetAllData), buf.Bytes())
	if err != nil {
		return currentVelocity, currentPosition, remainingSteps, inputVoltage, currentConsumption, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return currentVelocity, currentPosition, remainingSteps, inputVoltage, currentConsumption, DeviceError(header.ErrorCode)
		}

		if header.Length != 22 {
			return currentVelocity, currentPosition, remainingSteps, inputVoltage, currentConsumption, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 22)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &currentVelocity)
		binary.Read(resultBuf, binary.LittleEndian, &currentPosition)
		binary.Read(resultBuf, binary.LittleEndian, &remainingSteps)
		binary.Read(resultBuf, binary.LittleEndian, &inputVoltage)
		binary.Read(resultBuf, binary.LittleEndian, &currentConsumption)

	}

	return currentVelocity, currentPosition, remainingSteps, inputVoltage, currentConsumption, nil
}

// Sets the period with which the RegisterAllDataCallback callback is triggered
// periodically. A value of 0 turns the callback off.
func (device *SilentStepperV2Bricklet) SetAllCallbackConfiguration(period uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period)

	resultBytes, err := device.device.Set(uint8(FunctionSetAllCallbackConfiguration), buf.Bytes())
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

// Returns the period as set by SetAllCallbackConfiguration.
func (device *SilentStepperV2Bricklet) GetAllDataCallbackConfiguraton() (period uint32, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetAllDataCallbackConfiguraton), buf.Bytes())
	if err != nil {
		return period, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return period, DeviceError(header.ErrorCode)
		}

		if header.Length != 12 {
			return period, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &period)

	}

	return period, nil
}

// Sets the GPIO configuration for the given channel.
// You can configure a debounce and the deceleration that is used if the action is
// configured as ``normal stop``. See SetGPIOAction.
func (device *SilentStepperV2Bricklet) SetGPIOConfiguration(channel uint8, debounce uint16, stopDeceleration uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, channel)
	binary.Write(&buf, binary.LittleEndian, debounce)
	binary.Write(&buf, binary.LittleEndian, stopDeceleration)

	resultBytes, err := device.device.Set(uint8(FunctionSetGPIOConfiguration), buf.Bytes())
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

// Returns the GPIO configuration for a channel as set by SetGPIOConfiguration.
func (device *SilentStepperV2Bricklet) GetGPIOConfiguration(channel uint8) (debounce uint16, stopDeceleration uint16, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, channel)

	resultBytes, err := device.device.Get(uint8(FunctionGetGPIOConfiguration), buf.Bytes())
	if err != nil {
		return debounce, stopDeceleration, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return debounce, stopDeceleration, DeviceError(header.ErrorCode)
		}

		if header.Length != 12 {
			return debounce, stopDeceleration, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &debounce)
		binary.Read(resultBuf, binary.LittleEndian, &stopDeceleration)

	}

	return debounce, stopDeceleration, nil
}

// Sets the GPIO action for the given channel.
//
// The action can be a normal stop, a full brake or a callback. Each for a rising
// edge or falling edge. The actions are a bitmask they can be used at the same time.
// You can for example trigger a full brake and a callback at the same time or for
// rising and falling edge.
//
// The deceleration speed for the normal stop can be configured with
// SetGPIOConfiguration.
//
// Associated constants:
//
//	* GPIOActionNone
//	* GPIOActionNormalStopRisingEdge
//	* GPIOActionNormalStopFallingEdge
//	* GPIOActionFullBrakeRisingEdge
//	* GPIOActionFullBrakeFallingEdge
//	* GPIOActionCallbackRisingEdge
//	* GPIOActionCallbackFallingEdge
func (device *SilentStepperV2Bricklet) SetGPIOAction(channel uint8, action GPIOAction) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, channel)
	binary.Write(&buf, binary.LittleEndian, action)

	resultBytes, err := device.device.Set(uint8(FunctionSetGPIOAction), buf.Bytes())
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

// Returns the GPIO action for a channel as set by SetGPIOAction.
//
// Associated constants:
//
//	* GPIOActionNone
//	* GPIOActionNormalStopRisingEdge
//	* GPIOActionNormalStopFallingEdge
//	* GPIOActionFullBrakeRisingEdge
//	* GPIOActionFullBrakeFallingEdge
//	* GPIOActionCallbackRisingEdge
//	* GPIOActionCallbackFallingEdge
func (device *SilentStepperV2Bricklet) GetGPIOAction(channel uint8) (action GPIOAction, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, channel)

	resultBytes, err := device.device.Get(uint8(FunctionGetGPIOAction), buf.Bytes())
	if err != nil {
		return action, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return action, DeviceError(header.ErrorCode)
		}

		if header.Length != 12 {
			return action, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &action)

	}

	return action, nil
}

// Returns the GPIO state for both channels. True if the state is ``high`` and
// false if the state is ``low``.
func (device *SilentStepperV2Bricklet) GetGPIOState() (gpioState [2]bool, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetGPIOState), buf.Bytes())
	if err != nil {
		return gpioState, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return gpioState, DeviceError(header.ErrorCode)
		}

		if header.Length != 9 {
			return gpioState, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		copy(gpioState[:], ByteSliceToBoolSlice(resultBuf.Next(1)))

	}

	return gpioState, nil
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
func (device *SilentStepperV2Bricklet) GetSPITFPErrorCount() (errorCountAckChecksum uint32, errorCountMessageChecksum uint32, errorCountFrame uint32, errorCountOverflow uint32, err error) {
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
func (device *SilentStepperV2Bricklet) SetBootloaderMode(mode BootloaderMode) (status BootloaderStatus, err error) {
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
func (device *SilentStepperV2Bricklet) GetBootloaderMode() (mode BootloaderMode, err error) {
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
func (device *SilentStepperV2Bricklet) SetWriteFirmwarePointer(pointer uint32) (err error) {
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
func (device *SilentStepperV2Bricklet) WriteFirmware(data [64]uint8) (status uint8, err error) {
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
func (device *SilentStepperV2Bricklet) SetStatusLEDConfig(config StatusLEDConfig) (err error) {
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
func (device *SilentStepperV2Bricklet) GetStatusLEDConfig() (config StatusLEDConfig, err error) {
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
func (device *SilentStepperV2Bricklet) GetChipTemperature() (temperature int16, err error) {
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
func (device *SilentStepperV2Bricklet) Reset() (err error) {
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
func (device *SilentStepperV2Bricklet) WriteUID(uid uint32) (err error) {
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
func (device *SilentStepperV2Bricklet) ReadUID() (uid uint32, err error) {
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
func (device *SilentStepperV2Bricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
