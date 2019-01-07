/* ***********************************************************
 * This file was automatically generated on 2019-01-07.      *
 *                                                           *
 * Go Bindings Version 2.0.1                                 *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


//Drives one bipolar stepper motor with up to 38V and 2.5A per phase.
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricks/Stepper_Brick_Go.html.
package stepper_brick

import (
	"encoding/binary"
	"bytes"
    . "github.com/Tinkerforge/go-api-bindings/internal"
    "github.com/Tinkerforge/go-api-bindings/ipconnection"
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
	FunctionSetStepMode Function = 14
	FunctionGetStepMode Function = 15
	FunctionDriveForward Function = 16
	FunctionDriveBackward Function = 17
	FunctionStop Function = 18
	FunctionGetStackInputVoltage Function = 19
	FunctionGetExternalInputVoltage Function = 20
	FunctionGetCurrentConsumption Function = 21
	FunctionSetMotorCurrent Function = 22
	FunctionGetMotorCurrent Function = 23
	FunctionEnable Function = 24
	FunctionDisable Function = 25
	FunctionIsEnabled Function = 26
	FunctionSetDecay Function = 27
	FunctionGetDecay Function = 28
	FunctionSetMinimumVoltage Function = 29
	FunctionGetMinimumVoltage Function = 30
	FunctionSetSyncRect Function = 33
	FunctionIsSyncRect Function = 34
	FunctionSetTimeBase Function = 35
	FunctionGetTimeBase Function = 36
	FunctionGetAllData Function = 37
	FunctionSetAllDataPeriod Function = 38
	FunctionGetAllDataPeriod Function = 39
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
	FunctionCallbackUnderVoltage Function = 31
	FunctionCallbackPositionReached Function = 32
	FunctionCallbackAllData Function = 40
	FunctionCallbackNewState Function = 41
)

type StepMode uint8

const (
    StepModeFullStep StepMode = 1
	StepModeHalfStep StepMode = 2
	StepModeQuarterStep StepMode = 4
	StepModeEighthStep StepMode = 8
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

type StepperBrick struct{
	device Device
}
const DeviceIdentifier = 15
const DeviceDisplayName = "Stepper Brick"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (StepperBrick, error) {
    internalIPCon := ipcon.GetInternalHandle().(IPConnection)
    dev, err := NewDevice([3]uint8{ 2,0,4 }, uid, &internalIPCon, 0)
    if err != nil {
        return StepperBrick{}, err
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
	dev.ResponseExpected[FunctionSetStepMode] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetStepMode] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionDriveForward] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionDriveBackward] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionStop] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetStackInputVoltage] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetExternalInputVoltage] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetCurrentConsumption] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetMotorCurrent] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetMotorCurrent] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionEnable] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionDisable] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionIsEnabled] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetDecay] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetDecay] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetMinimumVoltage] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetMinimumVoltage] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetSyncRect] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionIsSyncRect] = ResponseExpectedFlagAlwaysTrue;
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
    return StepperBrick{dev}, nil
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
func (device *StepperBrick) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *StepperBrick) SetResponseExpected(functionID Function, responseExpected bool) error {
    return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *StepperBrick) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *StepperBrick) GetAPIVersion() [3]uint8 {
    return device.device.GetAPIVersion()
}

// This callback is triggered when the input voltage drops below the value set by
	// SetMinimumVoltage. The parameter is the current voltage given
	// in mV.
func (device *StepperBrick) RegisterUnderVoltageCallback(fn func(uint16)) uint64 {
            wrapper := func(byteSlice []byte) {
                buf := bytes.NewBuffer(byteSlice[8:])
                var voltage uint16
                binary.Read(buf, binary.LittleEndian, &voltage)
                fn(voltage)
            }
    return device.device.RegisterCallback(uint8(FunctionCallbackUnderVoltage), wrapper)
}

//Remove a registered Under Voltage callback.
func (device *StepperBrick) DeregisterUnderVoltageCallback(callbackID uint64) {
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
func (device *StepperBrick) RegisterPositionReachedCallback(fn func(int32)) uint64 {
            wrapper := func(byteSlice []byte) {
                buf := bytes.NewBuffer(byteSlice[8:])
                var position int32
                binary.Read(buf, binary.LittleEndian, &position)
                fn(position)
            }
    return device.device.RegisterCallback(uint8(FunctionCallbackPositionReached), wrapper)
}

//Remove a registered Position Reached callback.
func (device *StepperBrick) DeregisterPositionReachedCallback(callbackID uint64) {
    device.device.DeregisterCallback(uint8(FunctionCallbackPositionReached), callbackID)
}


// This callback is triggered periodically with the period that is set by
	// SetAllDataPeriod. The parameters are: the current velocity,
	// the current position, the remaining steps, the stack voltage, the external
	// voltage and the current consumption of the stepper motor.
func (device *StepperBrick) RegisterAllDataCallback(fn func(uint16, int32, int32, uint16, uint16, uint16)) uint64 {
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
func (device *StepperBrick) DeregisterAllDataCallback(callbackID uint64) {
    device.device.DeregisterCallback(uint8(FunctionCallbackAllData), callbackID)
}


// This callback is triggered whenever the Stepper Brick enters a new state.
	// It returns the new state as well as the previous state.
func (device *StepperBrick) RegisterNewStateCallback(fn func(State, State)) uint64 {
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
func (device *StepperBrick) DeregisterNewStateCallback(callbackID uint64) {
    device.device.DeregisterCallback(uint8(FunctionCallbackNewState), callbackID)
}


// Sets the maximum velocity of the stepper motor in steps per second.
	// This function does *not* start the motor, it merely sets the maximum
	// velocity the stepper motor is accelerated to. To get the motor running use
	// either SetTargetPosition, SetSteps, DriveForward or
	// DriveBackward.
func (device *StepperBrick) SetMaxVelocity(velocity uint16) (err error) {    
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
func (device *StepperBrick) GetMaxVelocity() (velocity uint16, err error) {    
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
func (device *StepperBrick) GetCurrentVelocity() (velocity uint16, err error) {    
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
func (device *StepperBrick) SetSpeedRamping(acceleration uint16, deacceleration uint16) (err error) {    
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
func (device *StepperBrick) GetSpeedRamping() (acceleration uint16, deacceleration uint16, err error) {    
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
func (device *StepperBrick) FullBrake() (err error) {    
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
func (device *StepperBrick) SetCurrentPosition(position int32) (err error) {    
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
func (device *StepperBrick) GetCurrentPosition() (position int32, err error) {    
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
func (device *StepperBrick) SetTargetPosition(position int32) (err error) {    
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
func (device *StepperBrick) GetTargetPosition() (position int32, err error) {    
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
func (device *StepperBrick) SetSteps(steps int32) (err error) {    
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
func (device *StepperBrick) GetSteps() (steps int32, err error) {    
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
func (device *StepperBrick) GetRemainingSteps() (steps int32, err error) {    
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

// Sets the step mode of the stepper motor. Possible values are:
	// 
	// * Full Step = 1
	// * Half Step = 2
	// * Quarter Step = 4
	// * Eighth Step = 8
	// 
	// A higher value will increase the resolution and
	// decrease the torque of the stepper motor.
	// 
	// The default value is 8 (Eighth Step).
//
// Associated constants:
//
//	* StepModeFullStep
//	* StepModeHalfStep
//	* StepModeQuarterStep
//	* StepModeEighthStep
func (device *StepperBrick) SetStepMode(mode StepMode) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, mode);

    resultBytes, err := device.device.Set(uint8(FunctionSetStepMode), buf.Bytes())
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

// Returns the step mode as set by SetStepMode.
//
// Associated constants:
//
//	* StepModeFullStep
//	* StepModeHalfStep
//	* StepModeQuarterStep
//	* StepModeEighthStep
func (device *StepperBrick) GetStepMode() (mode StepMode, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetStepMode), buf.Bytes())
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

// Drives the stepper motor forward until DriveBackward or
	// Stop is called. The velocity, acceleration and deacceleration as
	// set by SetMaxVelocity and SetSpeedRamping will be used.
func (device *StepperBrick) DriveForward() (err error) {    
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
func (device *StepperBrick) DriveBackward() (err error) {    
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
func (device *StepperBrick) Stop() (err error) {    
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
func (device *StepperBrick) GetStackInputVoltage() (voltage uint16, err error) {    
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
	// given via the black power input connector on the Stepper Brick.
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
func (device *StepperBrick) GetExternalInputVoltage() (voltage uint16, err error) {    
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

// Returns the current consumption of the motor in mA.
func (device *StepperBrick) GetCurrentConsumption() (current uint16, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetCurrentConsumption), buf.Bytes())
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

// Sets the current in mA with which the motor will be driven.
	// The minimum value is 100mA, the maximum value 2291mA and the
	// default value is 800mA.
	// 
	// Warning
	//  Do not set this value above the specifications of your stepper motor.
	//  Otherwise it may damage your motor.
func (device *StepperBrick) SetMotorCurrent(current uint16) (err error) {    
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
func (device *StepperBrick) GetMotorCurrent() (current uint16, err error) {    
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
func (device *StepperBrick) Enable() (err error) {    
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
func (device *StepperBrick) Disable() (err error) {    
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
func (device *StepperBrick) IsEnabled() (enabled bool, err error) {    
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

// Sets the decay mode of the stepper motor. The possible value range is
	// between 0 and 65535. A value of 0 sets the fast decay mode, a value of
	// 65535 sets the slow decay mode and a value in between sets the mixed
	// decay mode.
	// 
	// Changing the decay mode is only possible if synchronous rectification
	// is enabled (see :func:http://ebldc.com/?p=86/ blog post by Avayan.
	// 
	// A good decay mode is unfortunately different for every motor. The best
	// way to work out a good decay mode for your stepper motor, if you can't
	// measure the current with an oscilloscope, is to listen to the sound of
	// the motor. If the value is too low, you often hear a high pitched
	// sound and if it is too high you can often hear a humming sound.
	// 
	// Generally, fast decay mode (small value) will be noisier but also
	// allow higher motor speeds.
	// 
	// The default value is 10000.
	// 
	// Note
	//  There is unfortunately no formula to calculate a perfect decay
	//  mode for a given stepper motor. If you have problems with loud noises
	//  or the maximum motor speed is too slow, you should try to tinker with
	//  the decay value
func (device *StepperBrick) SetDecay(decay uint16) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, decay);

    resultBytes, err := device.device.Set(uint8(FunctionSetDecay), buf.Bytes())
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

// Returns the decay mode as set by SetDecay.
func (device *StepperBrick) GetDecay() (decay uint16, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetDecay), buf.Bytes())
    if err != nil {
        return decay, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return decay, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &decay)

    }
    
    return decay, nil
}

// Sets the minimum voltage in mV, below which the RegisterUnderVoltageCallback callback
	// is triggered. The minimum possible value that works with the Stepper Brick is 8V.
	// You can use this function to detect the discharge of a battery that is used
	// to drive the stepper motor. If you have a fixed power supply, you likely do
	// not need this functionality.
	// 
	// The default value is 8V.
func (device *StepperBrick) SetMinimumVoltage(voltage uint16) (err error) {    
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
func (device *StepperBrick) GetMinimumVoltage() (voltage uint16, err error) {    
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

// Turns synchronous rectification on or off (*true* or *false*).
	// 
	// With synchronous rectification on, the decay can be changed
	// (see :func:https://en.wikipedia.org/wiki/Active_rectification.
	// 
	// Warning
	//  If you want to use high speeds (> 10000 steps/s) for a large
	//  stepper motor with a large inductivity we strongly
	//  suggest that you disable synchronous rectification. Otherwise the
	//  Brick may not be able to cope with the load and overheat.
	// 
	// The default value is *false*.
func (device *StepperBrick) SetSyncRect(syncRect bool) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, syncRect);

    resultBytes, err := device.device.Set(uint8(FunctionSetSyncRect), buf.Bytes())
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

// Returns *true* if synchronous rectification is enabled, *false* otherwise.
func (device *StepperBrick) IsSyncRect() (syncRect bool, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionIsSyncRect), buf.Bytes())
    if err != nil {
        return syncRect, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return syncRect, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &syncRect)

    }
    
    return syncRect, nil
}

// Sets the time base of the velocity and the acceleration of the stepper brick
	// (in seconds).
	// 
	// For example, if you want to make one step every 1.5 seconds, you can set
	// the time base to 15 and the velocity to 10. Now the velocity is
	// 10steps/15s = 1steps/1.5s.
	// 
	// The default value is 1.
func (device *StepperBrick) SetTimeBase(timeBase uint32) (err error) {    
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
func (device *StepperBrick) GetTimeBase() (timeBase uint32, err error) {    
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
	// There is also a callback for this function, see RegisterAllDataCallback callback.
func (device *StepperBrick) GetAllData() (currentVelocity uint16, currentPosition int32, remainingSteps int32, stackVoltage uint16, externalVoltage uint16, currentConsumption uint16, err error) {    
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
func (device *StepperBrick) SetAllDataPeriod(period uint32) (err error) {    
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
func (device *StepperBrick) GetAllDataPeriod() (period uint32, err error) {    
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
	// .. versionadded:: 2.3.6$nbsp;(Firmware)
func (device *StepperBrick) SetSPITFPBaudrateConfig(enableDynamicBaudrate bool, minimumDynamicBaudrate uint32) (err error) {    
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
	// .. versionadded:: 2.3.6$nbsp;(Firmware)
func (device *StepperBrick) GetSPITFPBaudrateConfig() (enableDynamicBaudrate bool, minimumDynamicBaudrate uint32, err error) {    
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
	// .. versionadded:: 2.3.4$nbsp;(Firmware)
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
func (device *StepperBrick) GetSendTimeoutCount(communicationMethod CommunicationMethod) (timeoutCount uint32, err error) {    
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
	// 
	// .. versionadded:: 2.3.3$nbsp;(Firmware)
func (device *StepperBrick) SetSPITFPBaudrate(brickletPort rune, baudrate uint32) (err error) {    
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
	// 
	// .. versionadded:: 2.3.3$nbsp;(Firmware)
func (device *StepperBrick) GetSPITFPBaudrate(brickletPort rune) (baudrate uint32, err error) {    
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
	// 
	// .. versionadded:: 2.3.3$nbsp;(Firmware)
func (device *StepperBrick) GetSPITFPErrorCount(brickletPort rune) (errorCountACKChecksum uint32, errorCountMessageChecksum uint32, errorCountFrame uint32, errorCountOverflow uint32, err error) {    
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
	// 
	// .. versionadded:: 2.3.1$nbsp;(Firmware)
func (device *StepperBrick) EnableStatusLED() (err error) {    
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
	// 
	// .. versionadded:: 2.3.1$nbsp;(Firmware)
func (device *StepperBrick) DisableStatusLED() (err error) {    
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
	// 
	// .. versionadded:: 2.3.1$nbsp;(Firmware)
func (device *StepperBrick) IsStatusLEDEnabled() (enabled bool, err error) {    
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
func (device *StepperBrick) GetProtocol1BrickletName(port rune) (protocolVersion uint8, firmwareVersion [3]uint8, name string, err error) {    
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
func (device *StepperBrick) GetChipTemperature() (temperature int16, err error) {    
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
func (device *StepperBrick) Reset() (err error) {    
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
func (device *StepperBrick) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {    
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
