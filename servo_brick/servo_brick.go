/* ***********************************************************
 * This file was automatically generated on 2022-08-08.      *
 *                                                           *
 * Go Bindings Version 2.0.13                                *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/

// Drives up to 7 RC Servos with up to 3A‍.
//
//
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricks/Servo_Brick_Go.html.
package servo_brick

import (
	"bytes"
	"encoding/binary"
	"fmt"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionEnable                           Function = 1
	FunctionDisable                          Function = 2
	FunctionIsEnabled                        Function = 3
	FunctionSetPosition                      Function = 4
	FunctionGetPosition                      Function = 5
	FunctionGetCurrentPosition               Function = 6
	FunctionSetVelocity                      Function = 7
	FunctionGetVelocity                      Function = 8
	FunctionGetCurrentVelocity               Function = 9
	FunctionSetAcceleration                  Function = 10
	FunctionGetAcceleration                  Function = 11
	FunctionSetOutputVoltage                 Function = 12
	FunctionGetOutputVoltage                 Function = 13
	FunctionSetPulseWidth                    Function = 14
	FunctionGetPulseWidth                    Function = 15
	FunctionSetDegree                        Function = 16
	FunctionGetDegree                        Function = 17
	FunctionSetPeriod                        Function = 18
	FunctionGetPeriod                        Function = 19
	FunctionGetServoCurrent                  Function = 20
	FunctionGetOverallCurrent                Function = 21
	FunctionGetStackInputVoltage             Function = 22
	FunctionGetExternalInputVoltage          Function = 23
	FunctionSetMinimumVoltage                Function = 24
	FunctionGetMinimumVoltage                Function = 25
	FunctionEnablePositionReachedCallback    Function = 29
	FunctionDisablePositionReachedCallback   Function = 30
	FunctionIsPositionReachedCallbackEnabled Function = 31
	FunctionEnableVelocityReachedCallback    Function = 32
	FunctionDisableVelocityReachedCallback   Function = 33
	FunctionIsVelocityReachedCallbackEnabled Function = 34
	FunctionSetSPITFPBaudrateConfig          Function = 231
	FunctionGetSPITFPBaudrateConfig          Function = 232
	FunctionGetSendTimeoutCount              Function = 233
	FunctionSetSPITFPBaudrate                Function = 234
	FunctionGetSPITFPBaudrate                Function = 235
	FunctionGetSPITFPErrorCount              Function = 237
	FunctionEnableStatusLED                  Function = 238
	FunctionDisableStatusLED                 Function = 239
	FunctionIsStatusLEDEnabled               Function = 240
	FunctionGetProtocol1BrickletName         Function = 241
	FunctionGetChipTemperature               Function = 242
	FunctionReset                            Function = 243
	FunctionWriteBrickletPlugin              Function = 246
	FunctionReadBrickletPlugin               Function = 247
	FunctionGetIdentity                      Function = 255
	FunctionCallbackUnderVoltage             Function = 26
	FunctionCallbackPositionReached          Function = 27
	FunctionCallbackVelocityReached          Function = 28
)

type CommunicationMethod = uint8

const (
	CommunicationMethodNone     CommunicationMethod = 0
	CommunicationMethodUSB      CommunicationMethod = 1
	CommunicationMethodSPIStack CommunicationMethod = 2
	CommunicationMethodChibi    CommunicationMethod = 3
	CommunicationMethodRS485    CommunicationMethod = 4
	CommunicationMethodWIFI     CommunicationMethod = 5
	CommunicationMethodEthernet CommunicationMethod = 6
	CommunicationMethodWIFIV2   CommunicationMethod = 7
)

type ServoBrick struct {
	device Device
}

const DeviceIdentifier = 14
const DeviceDisplayName = "Servo Brick"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (ServoBrick, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{2, 0, 4}, uid, &internalIPCon, 0, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return ServoBrick{}, err
	}
	dev.ResponseExpected[FunctionEnable] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionDisable] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionIsEnabled] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetPosition] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetPosition] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetCurrentPosition] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetVelocity] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetVelocity] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetCurrentVelocity] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetAcceleration] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetAcceleration] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetOutputVoltage] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetOutputVoltage] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetPulseWidth] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetPulseWidth] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetDegree] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetDegree] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetPeriod] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetPeriod] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetServoCurrent] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetOverallCurrent] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetStackInputVoltage] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetExternalInputVoltage] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetMinimumVoltage] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionGetMinimumVoltage] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionEnablePositionReachedCallback] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionDisablePositionReachedCallback] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionIsPositionReachedCallbackEnabled] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionEnableVelocityReachedCallback] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionDisableVelocityReachedCallback] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionIsVelocityReachedCallbackEnabled] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetSPITFPBaudrateConfig] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetSPITFPBaudrateConfig] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetSendTimeoutCount] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetSPITFPBaudrate] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetSPITFPBaudrate] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetSPITFPErrorCount] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionEnableStatusLED] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionDisableStatusLED] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionIsStatusLEDEnabled] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetProtocol1BrickletName] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetChipTemperature] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionReset] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionWriteBrickletPlugin] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionReadBrickletPlugin] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetIdentity] = ResponseExpectedFlagAlwaysTrue
	return ServoBrick{dev}, nil
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
func (device *ServoBrick) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *ServoBrick) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *ServoBrick) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *ServoBrick) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// This callback is triggered when the input voltage drops below the value set by
// SetMinimumVoltage. The parameter is the current voltage.
func (device *ServoBrick) RegisterUnderVoltageCallback(fn func(uint16)) uint64 {
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
func (device *ServoBrick) DeregisterUnderVoltageCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackUnderVoltage), registrationId)
}

// This callback is triggered when a position set by SetPosition
// is reached. If the new position matches the current position then the
// callback is not triggered, because the servo didn't move.
// The parameters are the servo and the position that is reached.
//
// You can enable this callback with EnablePositionReachedCallback.
//
// Note
//  Since we can't get any feedback from the servo, this only works if the
//  velocity (see SetVelocity) is set smaller or equal to the
//  maximum velocity of the servo. Otherwise the servo will lag behind the
//  control value and the callback will be triggered too early.
func (device *ServoBrick) RegisterPositionReachedCallback(fn func(uint8, int16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 11 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var servoNum uint8
		var position int16
		binary.Read(buf, binary.LittleEndian, &servoNum)
		binary.Read(buf, binary.LittleEndian, &position)
		fn(servoNum, position)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackPositionReached), wrapper)
}

// Remove a registered Position Reached callback.
func (device *ServoBrick) DeregisterPositionReachedCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackPositionReached), registrationId)
}

// This callback is triggered when a velocity set by SetVelocity
// is reached. The parameters are the servo and the velocity that is reached.
//
// You can enable this callback with EnableVelocityReachedCallback.
//
// Note
//  Since we can't get any feedback from the servo, this only works if the
//  acceleration (see SetAcceleration) is set smaller or equal to the
//  maximum acceleration of the servo. Otherwise the servo will lag behind the
//  control value and the callback will be triggered too early.
func (device *ServoBrick) RegisterVelocityReachedCallback(fn func(uint8, int16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 11 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var servoNum uint8
		var velocity int16
		binary.Read(buf, binary.LittleEndian, &servoNum)
		binary.Read(buf, binary.LittleEndian, &velocity)
		fn(servoNum, velocity)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackVelocityReached), wrapper)
}

// Remove a registered Velocity Reached callback.
func (device *ServoBrick) DeregisterVelocityReachedCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackVelocityReached), registrationId)
}

// Enables a servo (0 to 6). If a servo is enabled, the configured position,
// velocity, acceleration, etc. are applied immediately.
func (device *ServoBrick) Enable(servoNum uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, servoNum)

	resultBytes, err := device.device.Set(uint8(FunctionEnable), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])

	}

	return nil
}

// Disables a servo (0 to 6). Disabled servos are not driven at all, i.e. a
// disabled servo will not hold its position if a load is applied.
func (device *ServoBrick) Disable(servoNum uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, servoNum)

	resultBytes, err := device.device.Set(uint8(FunctionDisable), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])

	}

	return nil
}

// Returns *true* if the specified servo is enabled, *false* otherwise.
func (device *ServoBrick) IsEnabled(servoNum uint8) (enabled bool, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, servoNum)

	resultBytes, err := device.device.Get(uint8(FunctionIsEnabled), buf.Bytes())
	if err != nil {
		return enabled, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return enabled, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return enabled, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &enabled)

	}

	return enabled, nil
}

// Sets the position for the specified servo.
//
// The default range of the position is -9000 to 9000, but it can be specified
// according to your servo with SetDegree.
//
// If you want to control a linear servo or RC brushless motor controller or
// similar with the Servo Brick, you can also define lengths or speeds with
// SetDegree.
func (device *ServoBrick) SetPosition(servoNum uint8, position int16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, servoNum)
	binary.Write(&buf, binary.LittleEndian, position)

	resultBytes, err := device.device.Set(uint8(FunctionSetPosition), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])

	}

	return nil
}

// Returns the position of the specified servo as set by SetPosition.
func (device *ServoBrick) GetPosition(servoNum uint8) (position int16, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, servoNum)

	resultBytes, err := device.device.Get(uint8(FunctionGetPosition), buf.Bytes())
	if err != nil {
		return position, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 10 {
			return position, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		if header.ErrorCode != 0 {
			return position, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &position)

	}

	return position, nil
}

// Returns the *current* position of the specified servo. This may not be the
// value of SetPosition if the servo is currently approaching a
// position goal.
func (device *ServoBrick) GetCurrentPosition(servoNum uint8) (position int16, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, servoNum)

	resultBytes, err := device.device.Get(uint8(FunctionGetCurrentPosition), buf.Bytes())
	if err != nil {
		return position, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 10 {
			return position, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		if header.ErrorCode != 0 {
			return position, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &position)

	}

	return position, nil
}

// Sets the maximum velocity of the specified servo. The velocity
// is accelerated according to the value set by SetAcceleration.
//
// The minimum velocity is 0 (no movement) and the maximum velocity is 65535.
// With a value of 65535 the position will be set immediately (no velocity).
func (device *ServoBrick) SetVelocity(servoNum uint8, velocity uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, servoNum)
	binary.Write(&buf, binary.LittleEndian, velocity)

	resultBytes, err := device.device.Set(uint8(FunctionSetVelocity), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])

	}

	return nil
}

// Returns the velocity of the specified servo as set by SetVelocity.
func (device *ServoBrick) GetVelocity(servoNum uint8) (velocity uint16, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, servoNum)

	resultBytes, err := device.device.Get(uint8(FunctionGetVelocity), buf.Bytes())
	if err != nil {
		return velocity, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 10 {
			return velocity, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		if header.ErrorCode != 0 {
			return velocity, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &velocity)

	}

	return velocity, nil
}

// Returns the *current* velocity of the specified servo. This may not be the
// value of SetVelocity if the servo is currently approaching a
// velocity goal.
func (device *ServoBrick) GetCurrentVelocity(servoNum uint8) (velocity uint16, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, servoNum)

	resultBytes, err := device.device.Get(uint8(FunctionGetCurrentVelocity), buf.Bytes())
	if err != nil {
		return velocity, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 10 {
			return velocity, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		if header.ErrorCode != 0 {
			return velocity, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &velocity)

	}

	return velocity, nil
}

// Sets the acceleration of the specified servo.
//
// The minimum acceleration is 1 and the maximum acceleration is 65535.
// With a value of 65535 the velocity will be set immediately (no acceleration).
func (device *ServoBrick) SetAcceleration(servoNum uint8, acceleration uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, servoNum)
	binary.Write(&buf, binary.LittleEndian, acceleration)

	resultBytes, err := device.device.Set(uint8(FunctionSetAcceleration), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])

	}

	return nil
}

// Returns the acceleration for the specified servo as set by
// SetAcceleration.
func (device *ServoBrick) GetAcceleration(servoNum uint8) (acceleration uint16, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, servoNum)

	resultBytes, err := device.device.Get(uint8(FunctionGetAcceleration), buf.Bytes())
	if err != nil {
		return acceleration, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 10 {
			return acceleration, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		if header.ErrorCode != 0 {
			return acceleration, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &acceleration)

	}

	return acceleration, nil
}

// Sets the output voltages with which the servos are driven.
//
// Note
//  We recommend that you set this value to the maximum voltage that is
//  specified for your servo, most servos achieve their maximum force only
//  with high voltages.
func (device *ServoBrick) SetOutputVoltage(voltage uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, voltage)

	resultBytes, err := device.device.Set(uint8(FunctionSetOutputVoltage), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])

	}

	return nil
}

// Returns the output voltage as specified by SetOutputVoltage.
func (device *ServoBrick) GetOutputVoltage() (voltage uint16, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetOutputVoltage), buf.Bytes())
	if err != nil {
		return voltage, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 10 {
			return voltage, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		if header.ErrorCode != 0 {
			return voltage, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &voltage)

	}

	return voltage, nil
}

// Sets the minimum and maximum pulse width of the specified servo.
//
// Usually, servos are controlled with a
// https://en.wikipedia.org/wiki/Pulse-width_modulation, whereby the
// length of the pulse controls the position of the servo. Every servo has
// different minimum and maximum pulse widths, these can be specified with
// this function.
//
// If you have a datasheet for your servo that specifies the minimum and
// maximum pulse width, you should set the values accordingly. If your servo
// comes without any datasheet you have to find the values via trial and error.
//
// The minimum must be smaller than the maximum.
func (device *ServoBrick) SetPulseWidth(servoNum uint8, min uint16, max uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, servoNum)
	binary.Write(&buf, binary.LittleEndian, min)
	binary.Write(&buf, binary.LittleEndian, max)

	resultBytes, err := device.device.Set(uint8(FunctionSetPulseWidth), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])

	}

	return nil
}

// Returns the minimum and maximum pulse width for the specified servo as set by
// SetPulseWidth.
func (device *ServoBrick) GetPulseWidth(servoNum uint8) (min uint16, max uint16, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, servoNum)

	resultBytes, err := device.device.Get(uint8(FunctionGetPulseWidth), buf.Bytes())
	if err != nil {
		return min, max, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 12 {
			return min, max, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		if header.ErrorCode != 0 {
			return min, max, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &min)
		binary.Read(resultBuf, binary.LittleEndian, &max)

	}

	return min, max, nil
}

// Sets the minimum and maximum degree for the specified servo (by default
// given as °/100).
//
// This only specifies the abstract values between which the minimum and maximum
// pulse width is scaled. For example: If you specify a pulse width of 1000µs
// to 2000µs and a degree range of -90° to 90°, a call of SetPosition
// with 0 will result in a pulse width of 1500µs
// (-90° = 1000µs, 90° = 2000µs, etc.).
//
// Possible usage:
//
// * The datasheet of your servo specifies a range of 200° with the middle position
//   at 110°. In this case you can set the minimum to -9000 and the maximum to 11000.
// * You measure a range of 220° on your servo and you don't have or need a middle
//   position. In this case you can set the minimum to 0 and the maximum to 22000.
// * You have a linear servo with a drive length of 20cm, In this case you could
//   set the minimum to 0 and the maximum to 20000. Now you can set the Position
//   with SetPosition with a resolution of cm/100. Also the velocity will
//   have a resolution of cm/100s and the acceleration will have a resolution of
//   cm/100s².
// * You don't care about units and just want the highest possible resolution. In
//   this case you should set the minimum to -32767 and the maximum to 32767.
// * You have a brushless motor with a maximum speed of 10000 rpm and want to
//   control it with a RC brushless motor controller. In this case you can set the
//   minimum to 0 and the maximum to 10000. SetPosition now controls the rpm.
//
// The minimum must be smaller than the maximum.
func (device *ServoBrick) SetDegree(servoNum uint8, min int16, max int16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, servoNum)
	binary.Write(&buf, binary.LittleEndian, min)
	binary.Write(&buf, binary.LittleEndian, max)

	resultBytes, err := device.device.Set(uint8(FunctionSetDegree), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])

	}

	return nil
}

// Returns the minimum and maximum degree for the specified servo as set by
// SetDegree.
func (device *ServoBrick) GetDegree(servoNum uint8) (min int16, max int16, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, servoNum)

	resultBytes, err := device.device.Get(uint8(FunctionGetDegree), buf.Bytes())
	if err != nil {
		return min, max, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 12 {
			return min, max, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		if header.ErrorCode != 0 {
			return min, max, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &min)
		binary.Read(resultBuf, binary.LittleEndian, &max)

	}

	return min, max, nil
}

// Sets the period of the specified servo.
//
// Usually, servos are controlled with a
// https://en.wikipedia.org/wiki/Pulse-width_modulation. Different
// servos expect PWMs with different periods. Most servos run well with a
// period of about 20ms.
//
// If your servo comes with a datasheet that specifies a period, you should
// set it accordingly. If you don't have a datasheet and you have no idea
// what the correct period is, the default value will most likely
// work fine.
func (device *ServoBrick) SetPeriod(servoNum uint8, period uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, servoNum)
	binary.Write(&buf, binary.LittleEndian, period)

	resultBytes, err := device.device.Set(uint8(FunctionSetPeriod), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])

	}

	return nil
}

// Returns the period for the specified servo as set by SetPeriod.
func (device *ServoBrick) GetPeriod(servoNum uint8) (period uint16, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, servoNum)

	resultBytes, err := device.device.Get(uint8(FunctionGetPeriod), buf.Bytes())
	if err != nil {
		return period, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 10 {
			return period, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		if header.ErrorCode != 0 {
			return period, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &period)

	}

	return period, nil
}

// Returns the current consumption of the specified servo.
func (device *ServoBrick) GetServoCurrent(servoNum uint8) (current uint16, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, servoNum)

	resultBytes, err := device.device.Get(uint8(FunctionGetServoCurrent), buf.Bytes())
	if err != nil {
		return current, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 10 {
			return current, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		if header.ErrorCode != 0 {
			return current, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &current)

	}

	return current, nil
}

// Returns the current consumption of all servos together.
func (device *ServoBrick) GetOverallCurrent() (current uint16, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetOverallCurrent), buf.Bytes())
	if err != nil {
		return current, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 10 {
			return current, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		if header.ErrorCode != 0 {
			return current, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &current)

	}

	return current, nil
}

// Returns the stack input voltage. The stack input voltage is the
// voltage that is supplied via the stack, i.e. it is given by a
// Step-Down or Step-Up Power Supply.
func (device *ServoBrick) GetStackInputVoltage() (voltage uint16, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetStackInputVoltage), buf.Bytes())
	if err != nil {
		return voltage, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 10 {
			return voltage, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		if header.ErrorCode != 0 {
			return voltage, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &voltage)

	}

	return voltage, nil
}

// Returns the external input voltage. The external input voltage is
// given via the black power input connector on the Servo Brick.
//
// If there is an external input voltage and a stack input voltage, the motors
// will be driven by the external input voltage. If there is only a stack
// voltage present, the motors will be driven by this voltage.
//
// Warning
//  This means, if you have a high stack voltage and a low external voltage,
//  the motors will be driven with the low external voltage. If you then remove
//  the external connection, it will immediately be driven by the high
//  stack voltage
func (device *ServoBrick) GetExternalInputVoltage() (voltage uint16, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetExternalInputVoltage), buf.Bytes())
	if err != nil {
		return voltage, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 10 {
			return voltage, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		if header.ErrorCode != 0 {
			return voltage, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &voltage)

	}

	return voltage, nil
}

// Sets the minimum voltage, below which the RegisterUnderVoltageCallback callback
// is triggered. The minimum possible value that works with the Servo Brick is 5V.
// You can use this function to detect the discharge of a battery that is used
// to drive the stepper motor. If you have a fixed power supply, you likely do
// not need this functionality.
func (device *ServoBrick) SetMinimumVoltage(voltage uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, voltage)

	resultBytes, err := device.device.Set(uint8(FunctionSetMinimumVoltage), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])

	}

	return nil
}

// Returns the minimum voltage as set by SetMinimumVoltage
func (device *ServoBrick) GetMinimumVoltage() (voltage uint16, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetMinimumVoltage), buf.Bytes())
	if err != nil {
		return voltage, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 10 {
			return voltage, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		if header.ErrorCode != 0 {
			return voltage, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &voltage)

	}

	return voltage, nil
}

// Enables the RegisterPositionReachedCallback callback.
//
// Default is disabled.
//
// .. versionadded:: 2.0.1$nbsp;(Firmware)
func (device *ServoBrick) EnablePositionReachedCallback() (err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Set(uint8(FunctionEnablePositionReachedCallback), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])

	}

	return nil
}

// Disables the RegisterPositionReachedCallback callback.
//
// .. versionadded:: 2.0.1$nbsp;(Firmware)
func (device *ServoBrick) DisablePositionReachedCallback() (err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Set(uint8(FunctionDisablePositionReachedCallback), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])

	}

	return nil
}

// Returns *true* if RegisterPositionReachedCallback callback is enabled, *false* otherwise.
//
// .. versionadded:: 2.0.1$nbsp;(Firmware)
func (device *ServoBrick) IsPositionReachedCallbackEnabled() (enabled bool, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionIsPositionReachedCallbackEnabled), buf.Bytes())
	if err != nil {
		return enabled, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return enabled, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return enabled, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &enabled)

	}

	return enabled, nil
}

// Enables the RegisterVelocityReachedCallback callback.
//
// Default is disabled.
//
// .. versionadded:: 2.0.1$nbsp;(Firmware)
func (device *ServoBrick) EnableVelocityReachedCallback() (err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Set(uint8(FunctionEnableVelocityReachedCallback), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])

	}

	return nil
}

// Disables the RegisterVelocityReachedCallback callback.
//
// Default is disabled.
//
// .. versionadded:: 2.0.1$nbsp;(Firmware)
func (device *ServoBrick) DisableVelocityReachedCallback() (err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Set(uint8(FunctionDisableVelocityReachedCallback), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])

	}

	return nil
}

// Returns *true* if RegisterVelocityReachedCallback callback is enabled, *false* otherwise.
//
// .. versionadded:: 2.0.1$nbsp;(Firmware)
func (device *ServoBrick) IsVelocityReachedCallbackEnabled() (enabled bool, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionIsVelocityReachedCallbackEnabled), buf.Bytes())
	if err != nil {
		return enabled, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return enabled, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return enabled, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &enabled)

	}

	return enabled, nil
}

// The SPITF protocol can be used with a dynamic baudrate. If the dynamic baudrate is
// enabled, the Brick will try to adapt the baudrate for the communication
// between Bricks and Bricklets according to the amount of data that is transferred.
//
// The baudrate will be increased exponentially if lots of data is sent/received and
// decreased linearly if little data is sent/received.
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
// .. versionadded:: 2.3.4$nbsp;(Firmware)
func (device *ServoBrick) SetSPITFPBaudrateConfig(enableDynamicBaudrate bool, minimumDynamicBaudrate uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, enableDynamicBaudrate)
	binary.Write(&buf, binary.LittleEndian, minimumDynamicBaudrate)

	resultBytes, err := device.device.Set(uint8(FunctionSetSPITFPBaudrateConfig), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])

	}

	return nil
}

// Returns the baudrate config, see SetSPITFPBaudrateConfig.
//
// .. versionadded:: 2.3.4$nbsp;(Firmware)
func (device *ServoBrick) GetSPITFPBaudrateConfig() (enableDynamicBaudrate bool, minimumDynamicBaudrate uint32, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetSPITFPBaudrateConfig), buf.Bytes())
	if err != nil {
		return enableDynamicBaudrate, minimumDynamicBaudrate, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 13 {
			return enableDynamicBaudrate, minimumDynamicBaudrate, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 13)
		}

		if header.ErrorCode != 0 {
			return enableDynamicBaudrate, minimumDynamicBaudrate, DeviceError(header.ErrorCode)
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
// .. versionadded:: 2.3.2$nbsp;(Firmware)
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
func (device *ServoBrick) GetSendTimeoutCount(communicationMethod CommunicationMethod) (timeoutCount uint32, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, communicationMethod)

	resultBytes, err := device.device.Get(uint8(FunctionGetSendTimeoutCount), buf.Bytes())
	if err != nil {
		return timeoutCount, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 12 {
			return timeoutCount, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		if header.ErrorCode != 0 {
			return timeoutCount, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &timeoutCount)

	}

	return timeoutCount, nil
}

// Sets the baudrate for a specific Bricklet port.
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
// or similar is necessary in your applications we recommend to not change
// the baudrate.
//
// .. versionadded:: 2.3.2$nbsp;(Firmware)
func (device *ServoBrick) SetSPITFPBaudrate(brickletPort rune, baudrate uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, brickletPort)
	binary.Write(&buf, binary.LittleEndian, baudrate)

	resultBytes, err := device.device.Set(uint8(FunctionSetSPITFPBaudrate), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])

	}

	return nil
}

// Returns the baudrate for a given Bricklet port, see SetSPITFPBaudrate.
//
// .. versionadded:: 2.3.2$nbsp;(Firmware)
func (device *ServoBrick) GetSPITFPBaudrate(brickletPort rune) (baudrate uint32, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, brickletPort)

	resultBytes, err := device.device.Get(uint8(FunctionGetSPITFPBaudrate), buf.Bytes())
	if err != nil {
		return baudrate, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 12 {
			return baudrate, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		if header.ErrorCode != 0 {
			return baudrate, DeviceError(header.ErrorCode)
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
// .. versionadded:: 2.3.2$nbsp;(Firmware)
func (device *ServoBrick) GetSPITFPErrorCount(brickletPort rune) (errorCountACKChecksum uint32, errorCountMessageChecksum uint32, errorCountFrame uint32, errorCountOverflow uint32, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, brickletPort)

	resultBytes, err := device.device.Get(uint8(FunctionGetSPITFPErrorCount), buf.Bytes())
	if err != nil {
		return errorCountACKChecksum, errorCountMessageChecksum, errorCountFrame, errorCountOverflow, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 24 {
			return errorCountACKChecksum, errorCountMessageChecksum, errorCountFrame, errorCountOverflow, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 24)
		}

		if header.ErrorCode != 0 {
			return errorCountACKChecksum, errorCountMessageChecksum, errorCountFrame, errorCountOverflow, DeviceError(header.ErrorCode)
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
func (device *ServoBrick) EnableStatusLED() (err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Set(uint8(FunctionEnableStatusLED), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
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
func (device *ServoBrick) DisableStatusLED() (err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Set(uint8(FunctionDisableStatusLED), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])

	}

	return nil
}

// Returns *true* if the status LED is enabled, *false* otherwise.
//
// .. versionadded:: 2.3.1$nbsp;(Firmware)
func (device *ServoBrick) IsStatusLEDEnabled() (enabled bool, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionIsStatusLEDEnabled), buf.Bytes())
	if err != nil {
		return enabled, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return enabled, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return enabled, DeviceError(header.ErrorCode)
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
func (device *ServoBrick) GetProtocol1BrickletName(port rune) (protocolVersion uint8, firmwareVersion [3]uint8, name string, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, port)

	resultBytes, err := device.device.Get(uint8(FunctionGetProtocol1BrickletName), buf.Bytes())
	if err != nil {
		return protocolVersion, firmwareVersion, name, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 52 {
			return protocolVersion, firmwareVersion, name, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 52)
		}

		if header.ErrorCode != 0 {
			return protocolVersion, firmwareVersion, name, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &protocolVersion)
		binary.Read(resultBuf, binary.LittleEndian, &firmwareVersion)
		name = ByteSliceToString(resultBuf.Next(40))

	}

	return protocolVersion, firmwareVersion, name, nil
}

// Returns the temperature as measured inside the microcontroller. The
// value returned is not the ambient temperature!
//
// The temperature is only proportional to the real temperature and it has an
// accuracy of ±15%. Practically it is only useful as an indicator for
// temperature changes.
func (device *ServoBrick) GetChipTemperature() (temperature int16, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetChipTemperature), buf.Bytes())
	if err != nil {
		return temperature, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 10 {
			return temperature, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		if header.ErrorCode != 0 {
			return temperature, DeviceError(header.ErrorCode)
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
func (device *ServoBrick) Reset() (err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Set(uint8(FunctionReset), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])

	}

	return nil
}

// Writes 32 bytes of firmware to the bricklet attached at the given port.
// The bytes are written to the position offset * 32.
//
// This function is used by Brick Viewer during flashing. It should not be
// necessary to call it in a normal user program.
func (device *ServoBrick) WriteBrickletPlugin(port rune, offset uint8, chunk [32]uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, port)
	binary.Write(&buf, binary.LittleEndian, offset)
	binary.Write(&buf, binary.LittleEndian, chunk)

	resultBytes, err := device.device.Set(uint8(FunctionWriteBrickletPlugin), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])

	}

	return nil
}

// Reads 32 bytes of firmware from the bricklet attached at the given port.
// The bytes are read starting at the position offset * 32.
//
// This function is used by Brick Viewer during flashing. It should not be
// necessary to call it in a normal user program.
func (device *ServoBrick) ReadBrickletPlugin(port rune, offset uint8) (chunk [32]uint8, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, port)
	binary.Write(&buf, binary.LittleEndian, offset)

	resultBytes, err := device.device.Get(uint8(FunctionReadBrickletPlugin), buf.Bytes())
	if err != nil {
		return chunk, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 40 {
			return chunk, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 40)
		}

		if header.ErrorCode != 0 {
			return chunk, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &chunk)

	}

	return chunk, nil
}

// Returns the UID, the UID where the Brick is connected to,
// the position, the hardware and firmware version as well as the
// device identifier.
//
// The position is the position in the stack from '0' (bottom) to '8' (top).
//
// The device identifier numbers can be found `here <device_identifier>`.
// |device_identifier_constant|
func (device *ServoBrick) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetIdentity), buf.Bytes())
	if err != nil {
		return uid, connectedUid, position, hardwareVersion, firmwareVersion, deviceIdentifier, err
	}
	if len(resultBytes) > 0 {
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

	}

	return uid, connectedUid, position, hardwareVersion, firmwareVersion, deviceIdentifier, nil
}
