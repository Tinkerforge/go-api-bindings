/* ***********************************************************
 * This file was automatically generated on 2018-12-18.      *
 *                                                           *
 * Go Bindings Version 2.0.0                                 *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


//Silently drives one bipolar stepper motor with up to 46V and 1.6A per phase.
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricks/SilentStepper_Brick_Go.html.
package silent_stepper_brick

import (
	"encoding/binary"
	"bytes"
    . "github.com/tinkerforge/go-api-bindings/internal"
    "github.com/tinkerforge/go-api-bindings/ipconnection"
)

type Function uint8

const (
    FunctionSetMaxVelocity Function = 1
	FunctionGetMaxVelocity Function = 2
	FunctionGetCurrentVelocity Function = 3
	FunctionSetSpeedRamping Function = 4
	FunctionGetSpeedRamping Function = 5
	FunctionFullBrake Function = 6
	FunctionSetCurrentPosition Function = 7
	FunctionGetCurrentPosition Function = 8
	FunctionSetTargetPosition Function = 9
	FunctionGetTargetPosition Function = 10
	FunctionSetSteps Function = 11
	FunctionGetSteps Function = 12
	FunctionGetRemainingSteps Function = 13
	FunctionSetStepConfiguration Function = 14
	FunctionGetStepConfiguration Function = 15
	FunctionDriveForward Function = 16
	FunctionDriveBackward Function = 17
	FunctionStop Function = 18
	FunctionGetStackInputVoltage Function = 19
	FunctionGetExternalInputVoltage Function = 20
	FunctionSetMotorCurrent Function = 22
	FunctionGetMotorCurrent Function = 23
	FunctionEnable Function = 24
	FunctionDisable Function = 25
	FunctionIsEnabled Function = 26
	FunctionSetBasicConfiguration Function = 27
	FunctionGetBasicConfiguration Function = 28
	FunctionSetSpreadcycleConfiguration Function = 29
	FunctionGetSpreadcycleConfiguration Function = 30
	FunctionSetStealthConfiguration Function = 31
	FunctionGetStealthConfiguration Function = 32
	FunctionSetCoolstepConfiguration Function = 33
	FunctionGetCoolstepConfiguration Function = 34
	FunctionSetMiscConfiguration Function = 35
	FunctionGetMiscConfiguration Function = 36
	FunctionGetDriverStatus Function = 37
	FunctionSetMinimumVoltage Function = 38
	FunctionGetMinimumVoltage Function = 39
	FunctionSetTimeBase Function = 42
	FunctionGetTimeBase Function = 43
	FunctionGetAllData Function = 44
	FunctionSetAllDataPeriod Function = 45
	FunctionGetAllDataPeriod Function = 46
	FunctionSetSPITFPBaudrateConfig Function = 231
	FunctionGetSPITFPBaudrateConfig Function = 232
	FunctionGetSendTimeoutCount Function = 233
	FunctionSetSPITFPBaudrate Function = 234
	FunctionGetSPITFPBaudrate Function = 235
	FunctionGetSPITFPErrorCount Function = 237
	FunctionEnableStatusLED Function = 238
	FunctionDisableStatusLED Function = 239
	FunctionIsStatusLEDEnabled Function = 240
	FunctionGetProtocol1BrickletName Function = 241
	FunctionGetChipTemperature Function = 242
	FunctionReset Function = 243
	FunctionGetIdentity Function = 255
	FunctionCallbackUnderVoltage Function = 40
	FunctionCallbackPositionReached Function = 41
	FunctionCallbackAllData Function = 47
	FunctionCallbackNewState Function = 48
)

type StepResolution uint8

const (
    StepResolution1 StepResolution = 8
	StepResolution2 StepResolution = 7
	StepResolution4 StepResolution = 6
	StepResolution8 StepResolution = 5
	StepResolution16 StepResolution = 4
	StepResolution32 StepResolution = 3
	StepResolution64 StepResolution = 2
	StepResolution128 StepResolution = 1
	StepResolution256 StepResolution = 0
)

type ChopperMode uint8

const (
    ChopperModeSpreadCycle ChopperMode = 0
	ChopperModeFastDecay ChopperMode = 1
)

type FreewheelMode uint8

const (
    FreewheelModeNormal FreewheelMode = 0
	FreewheelModeFreewheeling FreewheelMode = 1
	FreewheelModeCoilShortLS FreewheelMode = 2
	FreewheelModeCoilShortHS FreewheelMode = 3
)

type CurrentUpStepIncrement uint8

const (
    CurrentUpStepIncrement1 CurrentUpStepIncrement = 0
	CurrentUpStepIncrement2 CurrentUpStepIncrement = 1
	CurrentUpStepIncrement4 CurrentUpStepIncrement = 2
	CurrentUpStepIncrement8 CurrentUpStepIncrement = 3
)

type CurrentDownStepDecrement uint8

const (
    CurrentDownStepDecrement1 CurrentDownStepDecrement = 0
	CurrentDownStepDecrement2 CurrentDownStepDecrement = 1
	CurrentDownStepDecrement8 CurrentDownStepDecrement = 2
	CurrentDownStepDecrement32 CurrentDownStepDecrement = 3
)

type MinimumCurrent uint8

const (
    MinimumCurrentHalf MinimumCurrent = 0
	MinimumCurrentQuarter MinimumCurrent = 1
)

type StallguardMode uint8

const (
    StallguardModeStandard StallguardMode = 0
	StallguardModeFiltered StallguardMode = 1
)

type OpenLoad uint8

const (
    OpenLoadNone OpenLoad = 0
	OpenLoadPhaseA OpenLoad = 1
	OpenLoadPhaseB OpenLoad = 2
	OpenLoadPhaseAB OpenLoad = 3
)

type ShortToGround uint8

const (
    ShortToGroundNone ShortToGround = 0
	ShortToGroundPhaseA ShortToGround = 1
	ShortToGroundPhaseB ShortToGround = 2
	ShortToGroundPhaseAB ShortToGround = 3
)

type OverTemperature uint8

const (
    OverTemperatureNone OverTemperature = 0
	OverTemperatureWarning OverTemperature = 1
	OverTemperatureLimit OverTemperature = 2
)

type State uint8

const (
    StateStop State = 1
	StateAcceleration State = 2
	StateRun State = 3
	StateDeacceleration State = 4
	StateDirectionChangeToForward State = 5
	StateDirectionChangeToBackward State = 6
)

type CommunicationMethod uint8

const (
    CommunicationMethodNone CommunicationMethod = 0
	CommunicationMethodUSB CommunicationMethod = 1
	CommunicationMethodSPIStack CommunicationMethod = 2
	CommunicationMethodChibi CommunicationMethod = 3
	CommunicationMethodRS485 CommunicationMethod = 4
	CommunicationMethodWIFI CommunicationMethod = 5
	CommunicationMethodEthernet CommunicationMethod = 6
	CommunicationMethodWIFIV2 CommunicationMethod = 7
)

type SilentStepperBrick struct{
	device Device
}
const DeviceIdentifier = 19
const DeviceDisplayName = "Silent Stepper Brick"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (SilentStepperBrick, error) {
    internalIPCon := ipcon.GetInternalHandle().(IPConnection)
    dev, err := NewDevice([3]uint8{ 2,0,1 }, uid, &internalIPCon, 0)
    if err != nil {
        return SilentStepperBrick{}, err
    }
    dev.ResponseExpected[FunctionSetMaxVelocity] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetMaxVelocity] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetCurrentVelocity] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetSpeedRamping] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetSpeedRamping] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionFullBrake] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionSetCurrentPosition] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetCurrentPosition] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetTargetPosition] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetTargetPosition] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetSteps] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetSteps] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetRemainingSteps] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetStepConfiguration] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetStepConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionDriveForward] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionDriveBackward] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionStop] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetStackInputVoltage] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetExternalInputVoltage] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetMotorCurrent] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetMotorCurrent] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionEnable] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionDisable] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionIsEnabled] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetBasicConfiguration] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetBasicConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetSpreadcycleConfiguration] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetSpreadcycleConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetStealthConfiguration] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetStealthConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetCoolstepConfiguration] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetCoolstepConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetMiscConfiguration] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetMiscConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetDriverStatus] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetMinimumVoltage] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetMinimumVoltage] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetTimeBase] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetTimeBase] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetAllData] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetAllDataPeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetAllDataPeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetSPITFPBaudrateConfig] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetSPITFPBaudrateConfig] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetSendTimeoutCount] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetSPITFPBaudrate] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetSPITFPBaudrate] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetSPITFPErrorCount] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionEnableStatusLED] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionDisableStatusLED] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionIsStatusLEDEnabled] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetProtocol1BrickletName] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetChipTemperature] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionReset] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetIdentity] = ResponseExpectedFlagAlwaysTrue;
    return SilentStepperBrick{dev}, nil
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
func (device *SilentStepperBrick) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *SilentStepperBrick) SetResponseExpected(functionID Function, responseExpected bool) error {
    return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *SilentStepperBrick) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *SilentStepperBrick) GetAPIVersion() [3]uint8 {
    return device.device.GetAPIVersion()
}

// This callback is triggered when the input voltage drops below the value set by
	// SetMinimumVoltage. The parameter is the current voltage given
	// in mV.
func (device *SilentStepperBrick) RegisterUnderVoltageCallback(fn func(uint16)) uint64 {
            wrapper := func(byteSlice []byte) {
                buf := bytes.NewBuffer(byteSlice[8:])
                var voltage uint16
                binary.Read(buf, binary.LittleEndian, &voltage)
                fn(voltage)
            }
    return device.device.RegisterCallback(uint8(FunctionCallbackUnderVoltage), wrapper)
}

//Remove a registered Under Voltage callback.
func (device *SilentStepperBrick) DeregisterUnderVoltageCallback(callbackID uint64) {
    device.device.DeregisterCallback(uint8(FunctionCallbackUnderVoltage), callbackID)
}


// This callback is triggered when a position set by SetSteps or
	// SetTargetPosition is reached.
	// 
	// Note
	//  Since we can't get any feedback from the stepper motor, this only works if the
	//  acceleration (see SetSpeedRamping) is set smaller or equal to the
	//  maximum acceleration of the motor. Otherwise the motor will lag behind the
	//  control value and the callback will be triggered too early.
func (device *SilentStepperBrick) RegisterPositionReachedCallback(fn func(int32)) uint64 {
            wrapper := func(byteSlice []byte) {
                buf := bytes.NewBuffer(byteSlice[8:])
                var position int32
                binary.Read(buf, binary.LittleEndian, &position)
                fn(position)
            }
    return device.device.RegisterCallback(uint8(FunctionCallbackPositionReached), wrapper)
}

//Remove a registered Position Reached callback.
func (device *SilentStepperBrick) DeregisterPositionReachedCallback(callbackID uint64) {
    device.device.DeregisterCallback(uint8(FunctionCallbackPositionReached), callbackID)
}


// This callback is triggered periodically with the period that is set by
	// SetAllDataPeriod. The parameters are: the current velocity,
	// the current position, the remaining steps, the stack voltage, the external
	// voltage and the current consumption of the stepper motor.
func (device *SilentStepperBrick) RegisterAllDataCallback(fn func(uint16, int32, int32, uint16, uint16, uint16)) uint64 {
            wrapper := func(byteSlice []byte) {
                buf := bytes.NewBuffer(byteSlice[8:])
                var currentVelocity uint16
var currentPosition int32
var remainingSteps int32
var stackVoltage uint16
var externalVoltage uint16
var currentConsumption uint16
                binary.Read(buf, binary.LittleEndian, &currentVelocity)
binary.Read(buf, binary.LittleEndian, &currentPosition)
binary.Read(buf, binary.LittleEndian, &remainingSteps)
binary.Read(buf, binary.LittleEndian, &stackVoltage)
binary.Read(buf, binary.LittleEndian, &externalVoltage)
binary.Read(buf, binary.LittleEndian, &currentConsumption)
                fn(currentVelocity, currentPosition, remainingSteps, stackVoltage, externalVoltage, currentConsumption)
            }
    return device.device.RegisterCallback(uint8(FunctionCallbackAllData), wrapper)
}

//Remove a registered All Data callback.
func (device *SilentStepperBrick) DeregisterAllDataCallback(callbackID uint64) {
    device.device.DeregisterCallback(uint8(FunctionCallbackAllData), callbackID)
}


// This callback is triggered whenever the Slient Stepper Brick enters a new state.
	// It returns the new state as well as the previous state.
func (device *SilentStepperBrick) RegisterNewStateCallback(fn func(State, State)) uint64 {
            wrapper := func(byteSlice []byte) {
                buf := bytes.NewBuffer(byteSlice[8:])
                var stateNew State
var statePrevious State
                binary.Read(buf, binary.LittleEndian, &stateNew)
binary.Read(buf, binary.LittleEndian, &statePrevious)
                fn(stateNew, statePrevious)
            }
    return device.device.RegisterCallback(uint8(FunctionCallbackNewState), wrapper)
}

//Remove a registered New State callback.
func (device *SilentStepperBrick) DeregisterNewStateCallback(callbackID uint64) {
    device.device.DeregisterCallback(uint8(FunctionCallbackNewState), callbackID)
}


// Sets the maximum velocity of the stepper motor in steps per second.
	// This function does *not* start the motor, it merely sets the maximum
	// velocity the stepper motor is accelerated to. To get the motor running use
	// either SetTargetPosition, SetSteps, DriveForward or
	// DriveBackward.
func (device *SilentStepperBrick) SetMaxVelocity(velocity uint16) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, velocity);

    resultBytes, err := device.device.Set(uint8(FunctionSetMaxVelocity), buf.Bytes())
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

// Returns the velocity as set by SetMaxVelocity.
func (device *SilentStepperBrick) GetMaxVelocity() (velocity uint16, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetMaxVelocity), buf.Bytes())
    if err != nil {
        return velocity, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return velocity, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &velocity)

    }
    
    return velocity, nil
}

// Returns the *current* velocity of the stepper motor in steps per second.
func (device *SilentStepperBrick) GetCurrentVelocity() (velocity uint16, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetCurrentVelocity), buf.Bytes())
    if err != nil {
        return velocity, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return velocity, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &velocity)

    }
    
    return velocity, nil
}

// Sets the acceleration and deacceleration of the stepper motor. The values
	// are given in *steps/s²*. An acceleration of 1000 means, that
	// every second the velocity is increased by 1000 *steps/s*.
	// 
	// For example: If the current velocity is 0 and you want to accelerate to a
	// velocity of 8000 *steps/s* in 10 seconds, you should set an acceleration
	// of 800 *steps/s²*.
	// 
	// An acceleration/deacceleration of 0 means instantaneous
	// acceleration/deacceleration (not recommended)
	// 
	// The default value is 1000 for both
func (device *SilentStepperBrick) SetSpeedRamping(acceleration uint16, deacceleration uint16) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, acceleration);
	binary.Write(&buf, binary.LittleEndian, deacceleration);

    resultBytes, err := device.device.Set(uint8(FunctionSetSpeedRamping), buf.Bytes())
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

// Returns the acceleration and deacceleration as set by
	// SetSpeedRamping.
func (device *SilentStepperBrick) GetSpeedRamping() (acceleration uint16, deacceleration uint16, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetSpeedRamping), buf.Bytes())
    if err != nil {
        return acceleration, deacceleration, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return acceleration, deacceleration, BrickletError(header.ErrorCode)
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
func (device *SilentStepperBrick) FullBrake() (err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Set(uint8(FunctionFullBrake), buf.Bytes())
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

// Sets the current steps of the internal step counter. This can be used to
	// set the current position to 0 when some kind of starting position
	// is reached (e.g. when a CNC machine reaches a corner).
func (device *SilentStepperBrick) SetCurrentPosition(position int32) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, position);

    resultBytes, err := device.device.Set(uint8(FunctionSetCurrentPosition), buf.Bytes())
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

// Returns the current position of the stepper motor in steps. On startup
	// the position is 0. The steps are counted with all possible driving
	// functions (SetTargetPosition, SetSteps, DriveForward or
	// DriveBackward). It also is possible to reset the steps to 0 or
	// set them to any other desired value with SetCurrentPosition.
func (device *SilentStepperBrick) GetCurrentPosition() (position int32, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetCurrentPosition), buf.Bytes())
    if err != nil {
        return position, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return position, BrickletError(header.ErrorCode)
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
func (device *SilentStepperBrick) SetTargetPosition(position int32) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, position);

    resultBytes, err := device.device.Set(uint8(FunctionSetTargetPosition), buf.Bytes())
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

// Returns the last target position as set by SetTargetPosition.
func (device *SilentStepperBrick) GetTargetPosition() (position int32, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetTargetPosition), buf.Bytes())
    if err != nil {
        return position, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return position, BrickletError(header.ErrorCode)
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
func (device *SilentStepperBrick) SetSteps(steps int32) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, steps);

    resultBytes, err := device.device.Set(uint8(FunctionSetSteps), buf.Bytes())
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

// Returns the last steps as set by SetSteps.
func (device *SilentStepperBrick) GetSteps() (steps int32, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetSteps), buf.Bytes())
    if err != nil {
        return steps, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return steps, BrickletError(header.ErrorCode)
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
func (device *SilentStepperBrick) GetRemainingSteps() (steps int32, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetRemainingSteps), buf.Bytes())
    if err != nil {
        return steps, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return steps, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &steps)

    }
    
    return steps, nil
}

// Sets the step resolution from full-step up to 1/256-step.
	// 
	// If interpolation is turned on, the Silent Stepper Brick will always interpolate
	// your step inputs as 1/256-step. If you use full-step mode with interpolation, each
	// step will generate 256 1/256 steps.
	// 
	// For maximum torque use full-step without interpolation. For maximum resolution use
	// 1/256-step. Turn interpolation on to make the Stepper driving less noisy.
	// 
	// If you often change the speed with high acceleration you should turn the
	// interpolation off.
	// 
	// The default is 1/256-step with interpolation on.
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
func (device *SilentStepperBrick) SetStepConfiguration(stepResolution StepResolution, interpolation bool) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, stepResolution);
	binary.Write(&buf, binary.LittleEndian, interpolation);

    resultBytes, err := device.device.Set(uint8(FunctionSetStepConfiguration), buf.Bytes())
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
func (device *SilentStepperBrick) GetStepConfiguration() (stepResolution StepResolution, interpolation bool, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetStepConfiguration), buf.Bytes())
    if err != nil {
        return stepResolution, interpolation, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return stepResolution, interpolation, BrickletError(header.ErrorCode)
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
func (device *SilentStepperBrick) DriveForward() (err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Set(uint8(FunctionDriveForward), buf.Bytes())
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

// Drives the stepper motor backward until DriveForward or
	// Stop is triggered. The velocity, acceleration and deacceleration as
	// set by SetMaxVelocity and SetSpeedRamping will be used.
func (device *SilentStepperBrick) DriveBackward() (err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Set(uint8(FunctionDriveBackward), buf.Bytes())
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

// Stops the stepper motor with the deacceleration as set by
	// SetSpeedRamping.
func (device *SilentStepperBrick) Stop() (err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Set(uint8(FunctionStop), buf.Bytes())
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

// Returns the stack input voltage in mV. The stack input voltage is the
	// voltage that is supplied via the stack, i.e. it is given by a
	// Step-Down or Step-Up Power Supply.
func (device *SilentStepperBrick) GetStackInputVoltage() (voltage uint16, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetStackInputVoltage), buf.Bytes())
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

// Returns the external input voltage in mV. The external input voltage is
	// given via the black power input connector on the Slient Stepper Brick.
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
func (device *SilentStepperBrick) GetExternalInputVoltage() (voltage uint16, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetExternalInputVoltage), buf.Bytes())
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

// Sets the current in mA with which the motor will be driven.
	// The minimum value is 360mA, the maximum value 1640mA and the
	// default value is 800mA.
	// 
	// Warning
	//  Do not set this value above the specifications of your stepper motor.
	//  Otherwise it may damage your motor.
func (device *SilentStepperBrick) SetMotorCurrent(current uint16) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, current);

    resultBytes, err := device.device.Set(uint8(FunctionSetMotorCurrent), buf.Bytes())
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

// Returns the current as set by SetMotorCurrent.
func (device *SilentStepperBrick) GetMotorCurrent() (current uint16, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetMotorCurrent), buf.Bytes())
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

// Enables the driver chip. The driver parameters can be configured (maximum velocity,
	// acceleration, etc) before it is enabled.
func (device *SilentStepperBrick) Enable() (err error) {    
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

// Disables the driver chip. The configurations are kept (maximum velocity,
	// acceleration, etc) but the motor is not driven until it is enabled again.
func (device *SilentStepperBrick) Disable() (err error) {    
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

// Returns *true* if the driver chip is enabled, *false* otherwise.
func (device *SilentStepperBrick) IsEnabled() (enabled bool, err error) {    
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

// Sets the basic configuration parameters for the different modes (Stealth, Coolstep, Classic).
	// 
	// * Standstill Current: This value can be used to lower the current during stand still. This might
	//   be reasonable to reduce the heating of the motor and the Brick. When the motor is in standstill
	//   the configured motor phase current will be driven until the configured
	//   Power Down Time is elapsed. After that the phase current will be reduced to the standstill
	//   current. The elapsed time for this reduction can be configured with the Standstill Delay Time.
	//   The unit is in mA and the maximum allowed value is the configured maximum motor current
	//   (see SetMotorCurrent).
	// 
	// * Motor Run Current: The value sets the motor current when the motor is running.
	//   Use a value of at least one half of the global maximum motor current for a good
	//   microstep performance. The unit is in mA and the maximum allowed value is the current
	//   motor current. The API maps the entered value to 1/32 ... 32/32 of the maximum
	//   motor current. This value should be used to change the motor current during motor movement,
	//   whereas the global maximum motor current should not be changed while the motor is moving
	//   (see SetMotorCurrent).
	// 
	// * Standstill Delay Time: Controls the duration for motor power down after a motion
	//   as soon as standstill is detected and the Power Down Time is expired. A high Standstill Delay
	//   Time results in a smooth transition that avoids motor jerk during power down.
	//   The value range is 0 to 307ms
	// 
	// * Power Down Time: Sets the delay time after a stand still.
	//   The value range is 0 to 5222ms.
	// 
	// * Stealth Threshold: Sets the upper threshold for Stealth mode in steps/s. The value range is
	//   0-65536 steps/s. If the velocity of the motor goes above this value, Stealth mode is turned
	//   off. Otherwise it is turned on. In Stealth mode the torque declines with high speed.
	// 
	// * Coolstep Threshold: Sets the lower threshold for Coolstep mode in steps/s. The value range is
	//   0-65536 steps/s. The Coolstep Threshold needs to be above the Stealth Threshold.
	// 
	// * Classic Threshold: Sets the lower threshold for classic mode. The value range is
	//   0-65536 steps/s. In classic mode the stepper becomes more noisy, but the torque is maximized.
	// 
	// * High Velocity Shopper Mode: If High Velocity Shopper Mode is enabled, the stepper control
	//   is optimized to run the stepper motors at high velocities.
	// 
	// If you want to use all three thresholds make sure that
	// Stealth Threshold < Coolstep Threshold < Classic Threshold.
	// 
	// The default values are:
	// 
	// * Standstill Current: 200
	// * Motor Run Current: 800
	// * Standstill Delay Time: 0
	// * Power Down Time: 1000
	// * Stealth Threshold: 500
	// * Coolstep Threshold: 500
	// * Classic Threshold: 1000
	// * High Velocity Shopper Mode: false
func (device *SilentStepperBrick) SetBasicConfiguration(standstillCurrent uint16, motorRunCurrent uint16, standstillDelayTime uint16, powerDownTime uint16, stealthThreshold uint16, coolstepThreshold uint16, classicThreshold uint16, highVelocityChopperMode bool) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, standstillCurrent);
	binary.Write(&buf, binary.LittleEndian, motorRunCurrent);
	binary.Write(&buf, binary.LittleEndian, standstillDelayTime);
	binary.Write(&buf, binary.LittleEndian, powerDownTime);
	binary.Write(&buf, binary.LittleEndian, stealthThreshold);
	binary.Write(&buf, binary.LittleEndian, coolstepThreshold);
	binary.Write(&buf, binary.LittleEndian, classicThreshold);
	binary.Write(&buf, binary.LittleEndian, highVelocityChopperMode);

    resultBytes, err := device.device.Set(uint8(FunctionSetBasicConfiguration), buf.Bytes())
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

// Returns the configuration as set by SetBasicConfiguration.
func (device *SilentStepperBrick) GetBasicConfiguration() (standstillCurrent uint16, motorRunCurrent uint16, standstillDelayTime uint16, powerDownTime uint16, stealthThreshold uint16, coolstepThreshold uint16, classicThreshold uint16, highVelocityChopperMode bool, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetBasicConfiguration), buf.Bytes())
    if err != nil {
        return standstillCurrent, motorRunCurrent, standstillDelayTime, powerDownTime, stealthThreshold, coolstepThreshold, classicThreshold, highVelocityChopperMode, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return standstillCurrent, motorRunCurrent, standstillDelayTime, powerDownTime, stealthThreshold, coolstepThreshold, classicThreshold, highVelocityChopperMode, BrickletError(header.ErrorCode)
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
	// * Slow Decay Duration: Controls duration of off time setting of slow decay phase. The value
	//   range is 0-15. 0 = driver disabled, all bridges off. Use 1 only with Comparator Blank time >= 2.
	// 
	// * Enable Random Slow Decay: Set to false to fix chopper off time as set by Slow Decay Duration.
	//   If you set it to true, Decay Duration is randomly modulated.
	// 
	// * Fast Decay Duration: Sets the fast decay duration. The value range is 0-15. This parameters is
	//   only used if the Chopper Mode is set to Fast Decay.
	// 
	// * Hysteresis Start Value: Sets the hysteresis start value. The value range is 0-7. This parameter is
	//   only used if the Chopper Mode is set to Spread Cycle.
	// 
	// * Hysteresis End Value: Sets the hysteresis end value. The value range is -3 to 12. This parameter is
	//   only used if the Chopper Mode is set to Spread Cycle.
	// 
	// * Sine Wave Offset: Sets the sine wave offset. The value range is -3 to 12. This parameters is
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
	// The default values are:
	// 
	// * Slow Decay Duration: 4
	// * Enable Random Slow Decay: 0
	// * Fast Decay Duration: 0
	// * Hysteresis Start Value: 0
	// * Hysteresis End Value: 0
	// * Sine Wave Offset: 0
	// * Chopper Mode: 0
	// * Comparator Blank Time: 1
	// * Fast Decay Without Comparator: false
//
// Associated constants:
//
//	* ChopperModeSpreadCycle
//	* ChopperModeFastDecay
func (device *SilentStepperBrick) SetSpreadcycleConfiguration(slowDecayDuration uint8, enableRandomSlowDecay bool, fastDecayDuration uint8, hysteresisStartValue uint8, hysteresisEndValue int8, sineWaveOffset int8, chopperMode ChopperMode, comparatorBlankTime uint8, fastDecayWithoutComparator bool) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, slowDecayDuration);
	binary.Write(&buf, binary.LittleEndian, enableRandomSlowDecay);
	binary.Write(&buf, binary.LittleEndian, fastDecayDuration);
	binary.Write(&buf, binary.LittleEndian, hysteresisStartValue);
	binary.Write(&buf, binary.LittleEndian, hysteresisEndValue);
	binary.Write(&buf, binary.LittleEndian, sineWaveOffset);
	binary.Write(&buf, binary.LittleEndian, chopperMode);
	binary.Write(&buf, binary.LittleEndian, comparatorBlankTime);
	binary.Write(&buf, binary.LittleEndian, fastDecayWithoutComparator);

    resultBytes, err := device.device.Set(uint8(FunctionSetSpreadcycleConfiguration), buf.Bytes())
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

// Returns the configuration as set by SetBasicConfiguration.
//
// Associated constants:
//
//	* ChopperModeSpreadCycle
//	* ChopperModeFastDecay
func (device *SilentStepperBrick) GetSpreadcycleConfiguration() (slowDecayDuration uint8, enableRandomSlowDecay bool, fastDecayDuration uint8, hysteresisStartValue uint8, hysteresisEndValue int8, sineWaveOffset int8, chopperMode ChopperMode, comparatorBlankTime uint8, fastDecayWithoutComparator bool, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetSpreadcycleConfiguration), buf.Bytes())
    if err != nil {
        return slowDecayDuration, enableRandomSlowDecay, fastDecayDuration, hysteresisStartValue, hysteresisEndValue, sineWaveOffset, chopperMode, comparatorBlankTime, fastDecayWithoutComparator, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return slowDecayDuration, enableRandomSlowDecay, fastDecayDuration, hysteresisStartValue, hysteresisEndValue, sineWaveOffset, chopperMode, comparatorBlankTime, fastDecayWithoutComparator, BrickletError(header.ErrorCode)
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
	//   this value defines the maximum PWM amplitude change per half wave. The value range is 0-255.
	// 
	// * Gradient: If autoscale is disabled, the PWM gradient is scaled by this value. If autoscale is enabled,
	//   this value defines the maximum PWM gradient. With autoscale a value above 64 is recommended,
	//   otherwise the regulation might not be able to measure the current. The value range is 0-255.
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
	// The default values are:
	// 
	// * Enable Stealth: true
	// * Amplitude: 128
	// * Gradient: 4
	// * Enable Autoscale: true
	// * Force Symmetric: false
	// * Freewheel Mode: 0 (Normal)
//
// Associated constants:
//
//	* FreewheelModeNormal
//	* FreewheelModeFreewheeling
//	* FreewheelModeCoilShortLS
//	* FreewheelModeCoilShortHS
func (device *SilentStepperBrick) SetStealthConfiguration(enableStealth bool, amplitude uint8, gradient uint8, enableAutoscale bool, forceSymmetric bool, freewheelMode FreewheelMode) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, enableStealth);
	binary.Write(&buf, binary.LittleEndian, amplitude);
	binary.Write(&buf, binary.LittleEndian, gradient);
	binary.Write(&buf, binary.LittleEndian, enableAutoscale);
	binary.Write(&buf, binary.LittleEndian, forceSymmetric);
	binary.Write(&buf, binary.LittleEndian, freewheelMode);

    resultBytes, err := device.device.Set(uint8(FunctionSetStealthConfiguration), buf.Bytes())
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

// Returns the configuration as set by SetStealthConfiguration.
//
// Associated constants:
//
//	* FreewheelModeNormal
//	* FreewheelModeFreewheeling
//	* FreewheelModeCoilShortLS
//	* FreewheelModeCoilShortHS
func (device *SilentStepperBrick) GetStealthConfiguration() (enableStealth bool, amplitude uint8, gradient uint8, enableAutoscale bool, forceSymmetric bool, freewheelMode FreewheelMode, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetStealthConfiguration), buf.Bytes())
    if err != nil {
        return enableStealth, amplitude, gradient, enableAutoscale, forceSymmetric, freewheelMode, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return enableStealth, amplitude, gradient, enableAutoscale, forceSymmetric, freewheelMode, BrickletError(header.ErrorCode)
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
	//   is increased to reduce motor load angle. The value range is 0-15. A value of 0 turns Coolstep off.
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
	// * Stallguard Threshold Value: Sets the level for stall output (see GetDriverStatus). The value
	//   range is -64 to +63. A lower value gives a higher sensitivity. You have to find a suitable value for your
	//   motor by trial and error, 0 works for most motors.
	// 
	// * Stallguard Mode: Set to 0 for standard resolution or 1 for filtered mode. In filtered mode the Stallguard
	//   signal will be updated every four full-steps.
	// 
	// The default values are:
	// 
	// * Minimum Stallguard Value: 2
	// * Maximum Stallguard Value: 10
	// * Current Up Step Width: 0
	// * Current Down Step Width: 0
	// * Minimum Current: 0
	// * Stallguard Threshold Value: 0
	// * Stallguard Mode: 0
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
func (device *SilentStepperBrick) SetCoolstepConfiguration(minimumStallguardValue uint8, maximumStallguardValue uint8, currentUpStepWidth CurrentUpStepIncrement, currentDownStepWidth CurrentDownStepDecrement, minimumCurrent MinimumCurrent, stallguardThresholdValue int8, stallguardMode StallguardMode) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, minimumStallguardValue);
	binary.Write(&buf, binary.LittleEndian, maximumStallguardValue);
	binary.Write(&buf, binary.LittleEndian, currentUpStepWidth);
	binary.Write(&buf, binary.LittleEndian, currentDownStepWidth);
	binary.Write(&buf, binary.LittleEndian, minimumCurrent);
	binary.Write(&buf, binary.LittleEndian, stallguardThresholdValue);
	binary.Write(&buf, binary.LittleEndian, stallguardMode);

    resultBytes, err := device.device.Set(uint8(FunctionSetCoolstepConfiguration), buf.Bytes())
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
func (device *SilentStepperBrick) GetCoolstepConfiguration() (minimumStallguardValue uint8, maximumStallguardValue uint8, currentUpStepWidth CurrentUpStepIncrement, currentDownStepWidth CurrentDownStepDecrement, minimumCurrent MinimumCurrent, stallguardThresholdValue int8, stallguardMode StallguardMode, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetCoolstepConfiguration), buf.Bytes())
    if err != nil {
        return minimumStallguardValue, maximumStallguardValue, currentUpStepWidth, currentDownStepWidth, minimumCurrent, stallguardThresholdValue, stallguardMode, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return minimumStallguardValue, maximumStallguardValue, currentUpStepWidth, currentDownStepWidth, minimumCurrent, stallguardThresholdValue, stallguardMode, BrickletError(header.ErrorCode)
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
	// 
	// The default values are:
	// 
	// * Disable Short To Ground Protection: 0
	// * Synchronize Phase Frequency: 0
func (device *SilentStepperBrick) SetMiscConfiguration(disableShortToGroundProtection bool, synchronizePhaseFrequency uint8) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, disableShortToGroundProtection);
	binary.Write(&buf, binary.LittleEndian, synchronizePhaseFrequency);

    resultBytes, err := device.device.Set(uint8(FunctionSetMiscConfiguration), buf.Bytes())
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

// Returns the configuration as set by SetMiscConfiguration.
func (device *SilentStepperBrick) GetMiscConfiguration() (disableShortToGroundProtection bool, synchronizePhaseFrequency uint8, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetMiscConfiguration), buf.Bytes())
    if err != nil {
        return disableShortToGroundProtection, synchronizePhaseFrequency, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return disableShortToGroundProtection, synchronizePhaseFrequency, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &disableShortToGroundProtection)
	binary.Read(resultBuf, binary.LittleEndian, &synchronizePhaseFrequency)

    }
    
    return disableShortToGroundProtection, synchronizePhaseFrequency, nil
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
	//   The returned value is between 0 and 31. It represents a multiplier of 1/32 to 32/32 of the
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
func (device *SilentStepperBrick) GetDriverStatus() (openLoad OpenLoad, shortToGround ShortToGround, overTemperature OverTemperature, motorStalled bool, actualMotorCurrent uint8, fullStepActive bool, stallguardResult uint8, stealthVoltageAmplitude uint8, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetDriverStatus), buf.Bytes())
    if err != nil {
        return openLoad, shortToGround, overTemperature, motorStalled, actualMotorCurrent, fullStepActive, stallguardResult, stealthVoltageAmplitude, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return openLoad, shortToGround, overTemperature, motorStalled, actualMotorCurrent, fullStepActive, stallguardResult, stealthVoltageAmplitude, BrickletError(header.ErrorCode)
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

// Sets the minimum voltage in mV, below which the RegisterUnderVoltageCallback callback
	// is triggered. The minimum possible value that works with the Slient Stepper
	// Brick is 8V.
	// You can use this function to detect the discharge of a battery that is used
	// to drive the stepper motor. If you have a fixed power supply, you likely do
	// not need this functionality.
	// 
	// The default value is 8V.
func (device *SilentStepperBrick) SetMinimumVoltage(voltage uint16) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, voltage);

    resultBytes, err := device.device.Set(uint8(FunctionSetMinimumVoltage), buf.Bytes())
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

// Returns the minimum voltage as set by SetMinimumVoltage.
func (device *SilentStepperBrick) GetMinimumVoltage() (voltage uint16, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetMinimumVoltage), buf.Bytes())
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

// Sets the time base of the velocity and the acceleration of the Silent Stepper
	// Brick (in seconds).
	// 
	// For example, if you want to make one step every 1.5 seconds, you can set
	// the time base to 15 and the velocity to 10. Now the velocity is
	// 10steps/15s = 1steps/1.5s.
	// 
	// The default value is 1.
func (device *SilentStepperBrick) SetTimeBase(timeBase uint32) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, timeBase);

    resultBytes, err := device.device.Set(uint8(FunctionSetTimeBase), buf.Bytes())
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

// Returns the time base as set by SetTimeBase.
func (device *SilentStepperBrick) GetTimeBase() (timeBase uint32, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetTimeBase), buf.Bytes())
    if err != nil {
        return timeBase, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return timeBase, BrickletError(header.ErrorCode)
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
func (device *SilentStepperBrick) GetAllData() (currentVelocity uint16, currentPosition int32, remainingSteps int32, stackVoltage uint16, externalVoltage uint16, currentConsumption uint16, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetAllData), buf.Bytes())
    if err != nil {
        return currentVelocity, currentPosition, remainingSteps, stackVoltage, externalVoltage, currentConsumption, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return currentVelocity, currentPosition, remainingSteps, stackVoltage, externalVoltage, currentConsumption, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &currentVelocity)
	binary.Read(resultBuf, binary.LittleEndian, &currentPosition)
	binary.Read(resultBuf, binary.LittleEndian, &remainingSteps)
	binary.Read(resultBuf, binary.LittleEndian, &stackVoltage)
	binary.Read(resultBuf, binary.LittleEndian, &externalVoltage)
	binary.Read(resultBuf, binary.LittleEndian, &currentConsumption)

    }
    
    return currentVelocity, currentPosition, remainingSteps, stackVoltage, externalVoltage, currentConsumption, nil
}

// Sets the period in ms with which the RegisterAllDataCallback callback is triggered
	// periodically. A value of 0 turns the callback off.
func (device *SilentStepperBrick) SetAllDataPeriod(period uint32) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, period);

    resultBytes, err := device.device.Set(uint8(FunctionSetAllDataPeriod), buf.Bytes())
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

// Returns the period as set by SetAllDataPeriod.
func (device *SilentStepperBrick) GetAllDataPeriod() (period uint32, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetAllDataPeriod), buf.Bytes())
    if err != nil {
        return period, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return period, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &period)

    }
    
    return period, nil
}

// The SPITF protocol can be used with a dynamic baudrate. If the dynamic baudrate is
	// enabled, the Brick will try to adapt the baudrate for the communication
	// between Bricks and Bricklets according to the amount of data that is transferred.
	// 
	// The baudrate will be increased exponentially if lots of data is send/received and
	// decreased linearly if little data is send/received.
	// 
	// This lowers the baudrate in applications where little data is transferred (e.g.
	// a weather station) and increases the robustness. If there is lots of data to transfer
	// (e.g. Thermal Imaging Bricklet) it automatically increases the baudrate as needed.
	// 
	// In cases where some data has to transferred as fast as possible every few seconds
	// (e.g. RS485 Bricklet with a high baudrate but small payload) you may want to turn
	// the dynamic baudrate off to get the highest possible performance.
	// 
	// The maximum value of the baudrate can be set per port with the function
	// SetSPITFPBaudrate. If the dynamic baudrate is disabled, the baudrate
	// as set by SetSPITFPBaudrate will be used statically.
	// 
	// The minimum dynamic baudrate has a value range of 400000 to 2000000 baud.
	// 
	// By default dynamic baudrate is enabled and the minimum dynamic baudrate is 400000.
	// 
	// .. versionadded:: 2.0.4$nbsp;(Firmware)
func (device *SilentStepperBrick) SetSPITFPBaudrateConfig(enableDynamicBaudrate bool, minimumDynamicBaudrate uint32) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, enableDynamicBaudrate);
	binary.Write(&buf, binary.LittleEndian, minimumDynamicBaudrate);

    resultBytes, err := device.device.Set(uint8(FunctionSetSPITFPBaudrateConfig), buf.Bytes())
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

// Returns the baudrate config, see SetSPITFPBaudrateConfig.
	// 
	// .. versionadded:: 2.0.4$nbsp;(Firmware)
func (device *SilentStepperBrick) GetSPITFPBaudrateConfig() (enableDynamicBaudrate bool, minimumDynamicBaudrate uint32, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetSPITFPBaudrateConfig), buf.Bytes())
    if err != nil {
        return enableDynamicBaudrate, minimumDynamicBaudrate, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return enableDynamicBaudrate, minimumDynamicBaudrate, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &enableDynamicBaudrate)
	binary.Read(resultBuf, binary.LittleEndian, &minimumDynamicBaudrate)

    }
    
    return enableDynamicBaudrate, minimumDynamicBaudrate, nil
}

// Returns the timeout count for the different communication methods.
	// 
	// The methods 0-2 are available for all Bricks, 3-7 only for Master Bricks.
	// 
	// This function is mostly used for debugging during development, in normal operation
	// the counters should nearly always stay at 0.
//
// Associated constants:
//
//	* CommunicationMethodNone
//	* CommunicationMethodUSB
//	* CommunicationMethodSPIStack
//	* CommunicationMethodChibi
//	* CommunicationMethodRS485
//	* CommunicationMethodWIFI
//	* CommunicationMethodEthernet
//	* CommunicationMethodWIFIV2
func (device *SilentStepperBrick) GetSendTimeoutCount(communicationMethod CommunicationMethod) (timeoutCount uint32, err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, communicationMethod);

    resultBytes, err := device.device.Get(uint8(FunctionGetSendTimeoutCount), buf.Bytes())
    if err != nil {
        return timeoutCount, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return timeoutCount, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &timeoutCount)

    }
    
    return timeoutCount, nil
}

// Sets the baudrate for a specific Bricklet port ('a' - 'd'). The
	// baudrate can be in the range 400000 to 2000000.
	// 
	// If you want to increase the throughput of Bricklets you can increase
	// the baudrate. If you get a high error count because of high
	// interference (see GetSPITFPErrorCount) you can decrease the
	// baudrate.
	// 
	// If the dynamic baudrate feature is enabled, the baudrate set by this
	// function corresponds to the maximum baudrate (see SetSPITFPBaudrateConfig).
	// 
	// Regulatory testing is done with the default baudrate. If CE compatibility
	// or similar is necessary in you applications we recommend to not change
	// the baudrate.
	// 
	// The default baudrate for all ports is 1400000.
func (device *SilentStepperBrick) SetSPITFPBaudrate(brickletPort rune, baudrate uint32) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, brickletPort);
	binary.Write(&buf, binary.LittleEndian, baudrate);

    resultBytes, err := device.device.Set(uint8(FunctionSetSPITFPBaudrate), buf.Bytes())
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

// Returns the baudrate for a given Bricklet port, see SetSPITFPBaudrate.
func (device *SilentStepperBrick) GetSPITFPBaudrate(brickletPort rune) (baudrate uint32, err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, brickletPort);

    resultBytes, err := device.device.Get(uint8(FunctionGetSPITFPBaudrate), buf.Bytes())
    if err != nil {
        return baudrate, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return baudrate, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &baudrate)

    }
    
    return baudrate, nil
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
	// The errors counts are for errors that occur on the Brick side. All
	// Bricklets have a similar function that returns the errors on the Bricklet side.
func (device *SilentStepperBrick) GetSPITFPErrorCount(brickletPort rune) (errorCountACKChecksum uint32, errorCountMessageChecksum uint32, errorCountFrame uint32, errorCountOverflow uint32, err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, brickletPort);

    resultBytes, err := device.device.Get(uint8(FunctionGetSPITFPErrorCount), buf.Bytes())
    if err != nil {
        return errorCountACKChecksum, errorCountMessageChecksum, errorCountFrame, errorCountOverflow, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return errorCountACKChecksum, errorCountMessageChecksum, errorCountFrame, errorCountOverflow, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &errorCountACKChecksum)
	binary.Read(resultBuf, binary.LittleEndian, &errorCountMessageChecksum)
	binary.Read(resultBuf, binary.LittleEndian, &errorCountFrame)
	binary.Read(resultBuf, binary.LittleEndian, &errorCountOverflow)

    }
    
    return errorCountACKChecksum, errorCountMessageChecksum, errorCountFrame, errorCountOverflow, nil
}

// Enables the status LED.
	// 
	// The status LED is the blue LED next to the USB connector. If enabled is is
	// on and it flickers if data is transfered. If disabled it is always off.
	// 
	// The default state is enabled.
func (device *SilentStepperBrick) EnableStatusLED() (err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Set(uint8(FunctionEnableStatusLED), buf.Bytes())
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

// Disables the status LED.
	// 
	// The status LED is the blue LED next to the USB connector. If enabled is is
	// on and it flickers if data is transfered. If disabled it is always off.
	// 
	// The default state is enabled.
func (device *SilentStepperBrick) DisableStatusLED() (err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Set(uint8(FunctionDisableStatusLED), buf.Bytes())
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

// Returns *true* if the status LED is enabled, *false* otherwise.
func (device *SilentStepperBrick) IsStatusLEDEnabled() (enabled bool, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionIsStatusLEDEnabled), buf.Bytes())
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

// Returns the firmware and protocol version and the name of the Bricklet for a
	// given port.
	// 
	// This functions sole purpose is to allow automatic flashing of v1.x.y Bricklet
	// plugins.
func (device *SilentStepperBrick) GetProtocol1BrickletName(port rune) (protocolVersion uint8, firmwareVersion [3]uint8, name string, err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, port);

    resultBytes, err := device.device.Get(uint8(FunctionGetProtocol1BrickletName), buf.Bytes())
    if err != nil {
        return protocolVersion, firmwareVersion, name, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return protocolVersion, firmwareVersion, name, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &protocolVersion)
	binary.Read(resultBuf, binary.LittleEndian, &firmwareVersion)
	name = ByteSliceToString(resultBuf.Next(40))

    }
    
    return protocolVersion, firmwareVersion, name, nil
}

// Returns the temperature in °C/10 as measured inside the microcontroller. The
	// value returned is not the ambient temperature!
	// 
	// The temperature is only proportional to the real temperature and it has an
	// accuracy of +-15%. Practically it is only useful as an indicator for
	// temperature changes.
func (device *SilentStepperBrick) GetChipTemperature() (temperature int16, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetChipTemperature), buf.Bytes())
    if err != nil {
        return temperature, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return temperature, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &temperature)

    }
    
    return temperature, nil
}

// Calling this function will reset the Brick. Calling this function
	// on a Brick inside of a stack will reset the whole stack.
	// 
	// After a reset you have to create new device objects,
	// calling functions on the existing ones will result in
	// undefined behavior!
func (device *SilentStepperBrick) Reset() (err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Set(uint8(FunctionReset), buf.Bytes())
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

// Returns the UID, the UID where the Brick is connected to,
	// the position, the hardware and firmware version as well as the
	// device identifier.
	// 
	// The position can be '0'-'8' (stack position).
	// 
	// The device identifier numbers can be found `here <device_identifier>`.
	// |device_identifier_constant|
func (device *SilentStepperBrick) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {    
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
