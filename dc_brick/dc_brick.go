/* ***********************************************************
 * This file was automatically generated on 2022-08-08.      *
 *                                                           *
 * Go Bindings Version 2.0.13                                *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/

// Drives one brushed DC motor with up to 28V and 5A (peak).
//
//
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricks/DC_Brick_Go.html.
package dc_brick

import (
	"bytes"
	"encoding/binary"
	"fmt"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionSetVelocity               Function = 1
	FunctionGetVelocity               Function = 2
	FunctionGetCurrentVelocity        Function = 3
	FunctionSetAcceleration           Function = 4
	FunctionGetAcceleration           Function = 5
	FunctionSetPWMFrequency           Function = 6
	FunctionGetPWMFrequency           Function = 7
	FunctionFullBrake                 Function = 8
	FunctionGetStackInputVoltage      Function = 9
	FunctionGetExternalInputVoltage   Function = 10
	FunctionGetCurrentConsumption     Function = 11
	FunctionEnable                    Function = 12
	FunctionDisable                   Function = 13
	FunctionIsEnabled                 Function = 14
	FunctionSetMinimumVoltage         Function = 15
	FunctionGetMinimumVoltage         Function = 16
	FunctionSetDriveMode              Function = 17
	FunctionGetDriveMode              Function = 18
	FunctionSetCurrentVelocityPeriod  Function = 19
	FunctionGetCurrentVelocityPeriod  Function = 20
	FunctionSetSPITFPBaudrateConfig   Function = 231
	FunctionGetSPITFPBaudrateConfig   Function = 232
	FunctionGetSendTimeoutCount       Function = 233
	FunctionSetSPITFPBaudrate         Function = 234
	FunctionGetSPITFPBaudrate         Function = 235
	FunctionGetSPITFPErrorCount       Function = 237
	FunctionEnableStatusLED           Function = 238
	FunctionDisableStatusLED          Function = 239
	FunctionIsStatusLEDEnabled        Function = 240
	FunctionGetProtocol1BrickletName  Function = 241
	FunctionGetChipTemperature        Function = 242
	FunctionReset                     Function = 243
	FunctionWriteBrickletPlugin       Function = 246
	FunctionReadBrickletPlugin        Function = 247
	FunctionGetIdentity               Function = 255
	FunctionCallbackUnderVoltage      Function = 21
	FunctionCallbackEmergencyShutdown Function = 22
	FunctionCallbackVelocityReached   Function = 23
	FunctionCallbackCurrentVelocity   Function = 24
)

type DriveMode = uint8

const (
	DriveModeDriveBrake DriveMode = 0
	DriveModeDriveCoast DriveMode = 1
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

type DCBrick struct {
	device Device
}

const DeviceIdentifier = 11
const DeviceDisplayName = "DC Brick"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (DCBrick, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{2, 0, 3}, uid, &internalIPCon, 0, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return DCBrick{}, err
	}
	dev.ResponseExpected[FunctionSetVelocity] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetVelocity] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetCurrentVelocity] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetAcceleration] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetAcceleration] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetPWMFrequency] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetPWMFrequency] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionFullBrake] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetStackInputVoltage] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetExternalInputVoltage] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetCurrentConsumption] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionEnable] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionDisable] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionIsEnabled] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetMinimumVoltage] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionGetMinimumVoltage] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetDriveMode] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetDriveMode] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetCurrentVelocityPeriod] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionGetCurrentVelocityPeriod] = ResponseExpectedFlagAlwaysTrue
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
	return DCBrick{dev}, nil
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
func (device *DCBrick) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *DCBrick) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *DCBrick) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *DCBrick) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// This callback is triggered when the input voltage drops below the value set by
// SetMinimumVoltage. The parameter is the current voltage.
func (device *DCBrick) RegisterUnderVoltageCallback(fn func(uint16)) uint64 {
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
func (device *DCBrick) DeregisterUnderVoltageCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackUnderVoltage), registrationId)
}

// This callback is triggered if either the current consumption
// is too high (above 5A) or the temperature of the driver chip is too high
// (above 175°C). These two possibilities are essentially the same, since the
// temperature will reach this threshold immediately if the motor consumes too
// much current. In case of a voltage below 3.3V (external or stack) this
// callback is triggered as well.
//
// If this callback is triggered, the driver chip gets disabled at the same time.
// That means, Enable has to be called to drive the motor again.
//
// Note
//  This callback only works in Drive/Brake mode (see SetDriveMode). In
//  Drive/Coast mode it is unfortunately impossible to reliably read the
//  overcurrent/overtemperature signal from the driver chip.
func (device *DCBrick) RegisterEmergencyShutdownCallback(fn func()) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 8 {
			return
		}

		fn()
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackEmergencyShutdown), wrapper)
}

// Remove a registered Emergency Shutdown callback.
func (device *DCBrick) DeregisterEmergencyShutdownCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackEmergencyShutdown), registrationId)
}

// This callback is triggered whenever a set velocity is reached. For example:
// If a velocity of 0 is present, acceleration is set to 5000 and velocity
// to 10000, the RegisterVelocityReachedCallback callback will be triggered after about
// 2 seconds, when the set velocity is actually reached.
//
// Note
//  Since we can't get any feedback from the DC motor, this only works if the
//  acceleration (see SetAcceleration) is set smaller or equal to the
//  maximum acceleration of the motor. Otherwise the motor will lag behind the
//  control value and the callback will be triggered too early.
func (device *DCBrick) RegisterVelocityReachedCallback(fn func(int16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 10 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var velocity int16
		binary.Read(buf, binary.LittleEndian, &velocity)
		fn(velocity)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackVelocityReached), wrapper)
}

// Remove a registered Velocity Reached callback.
func (device *DCBrick) DeregisterVelocityReachedCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackVelocityReached), registrationId)
}

// This callback is triggered with the period that is set by
// SetCurrentVelocityPeriod. The parameter is the *current*
// velocity used by the motor.
//
// The RegisterCurrentVelocityCallback callback is only triggered after the set period
// if there is a change in the velocity.
func (device *DCBrick) RegisterCurrentVelocityCallback(fn func(int16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 10 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var velocity int16
		binary.Read(buf, binary.LittleEndian, &velocity)
		fn(velocity)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackCurrentVelocity), wrapper)
}

// Remove a registered Current Velocity callback.
func (device *DCBrick) DeregisterCurrentVelocityCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackCurrentVelocity), registrationId)
}

// Sets the velocity of the motor. Whereas -32767 is full speed backward,
// 0 is stop and 32767 is full speed forward. Depending on the
// acceleration (see SetAcceleration), the motor is not immediately
// brought to the velocity but smoothly accelerated.
//
// The velocity describes the duty cycle of the PWM with which the motor is
// controlled, e.g. a velocity of 3277 sets a PWM with a 10% duty cycle.
// You can not only control the duty cycle of the PWM but also the frequency,
// see SetPWMFrequency.
func (device *DCBrick) SetVelocity(velocity int16) (err error) {
	var buf bytes.Buffer
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

// Returns the velocity as set by SetVelocity.
func (device *DCBrick) GetVelocity() (velocity int16, err error) {
	var buf bytes.Buffer

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

// Returns the *current* velocity of the motor. This value is different
// from GetVelocity whenever the motor is currently accelerating
// to a goal set by SetVelocity.
func (device *DCBrick) GetCurrentVelocity() (velocity int16, err error) {
	var buf bytes.Buffer

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

// Sets the acceleration of the motor. It is given in *velocity/s*. An
// acceleration of 10000 means, that every second the velocity is increased
// by 10000 (or about 30% duty cycle).
//
// For example: If the current velocity is 0 and you want to accelerate to a
// velocity of 16000 (about 50% duty cycle) in 10 seconds, you should set
// an acceleration of 1600.
//
// If acceleration is set to 0, there is no speed ramping, i.e. a new velocity
// is immediately given to the motor.
func (device *DCBrick) SetAcceleration(acceleration uint16) (err error) {
	var buf bytes.Buffer
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

// Returns the acceleration as set by SetAcceleration.
func (device *DCBrick) GetAcceleration() (acceleration uint16, err error) {
	var buf bytes.Buffer

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

// Sets the frequency of the PWM with which the motor is driven.
// Often a high frequency
// is less noisy and the motor runs smoother. However, with a low frequency
// there are less switches and therefore fewer switching losses. Also with
// most motors lower frequencies enable higher torque.
//
// If you have no idea what all this means, just ignore this function and use
// the default frequency, it will very likely work fine.
func (device *DCBrick) SetPWMFrequency(frequency uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, frequency)

	resultBytes, err := device.device.Set(uint8(FunctionSetPWMFrequency), buf.Bytes())
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

// Returns the PWM frequency as set by SetPWMFrequency.
func (device *DCBrick) GetPWMFrequency() (frequency uint16, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetPWMFrequency), buf.Bytes())
	if err != nil {
		return frequency, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 10 {
			return frequency, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		if header.ErrorCode != 0 {
			return frequency, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &frequency)

	}

	return frequency, nil
}

// Executes an active full brake.
//
// Warning
//  This function is for emergency purposes,
//  where an immediate brake is necessary. Depending on the current velocity and
//  the strength of the motor, a full brake can be quite violent.
//
// Call SetVelocity with 0 if you just want to stop the motor.
func (device *DCBrick) FullBrake() (err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Set(uint8(FunctionFullBrake), buf.Bytes())
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

// Returns the stack input voltage. The stack input voltage is the
// voltage that is supplied via the stack, i.e. it is given by a
// Step-Down or Step-Up Power Supply.
func (device *DCBrick) GetStackInputVoltage() (voltage uint16, err error) {
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
// given via the black power input connector on the DC Brick.
//
// If there is an external input voltage and a stack input voltage, the motor
// will be driven by the external input voltage. If there is only a stack
// voltage present, the motor will be driven by this voltage.
//
// Warning
//  This means, if you have a high stack voltage and a low external voltage,
//  the motor will be driven with the low external voltage. If you then remove
//  the external connection, it will immediately be driven by the high
//  stack voltage.
func (device *DCBrick) GetExternalInputVoltage() (voltage uint16, err error) {
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

// Returns the current consumption of the motor.
func (device *DCBrick) GetCurrentConsumption() (voltage uint16, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetCurrentConsumption), buf.Bytes())
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

// Enables the driver chip. The driver parameters can be configured (velocity,
// acceleration, etc) before it is enabled.
func (device *DCBrick) Enable() (err error) {
	var buf bytes.Buffer

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

// Disables the driver chip. The configurations are kept (velocity,
// acceleration, etc) but the motor is not driven until it is enabled again.
//
// Warning
//  Disabling the driver chip while the motor is still turning can damage the
//  driver chip. The motor should be stopped calling SetVelocity with 0
//  before disabling the motor power. The SetVelocity function will **not**
//  wait until the motor is actually stopped. You have to explicitly wait for the
//  appropriate time after calling the SetVelocity function before calling
//  the Disable function.
func (device *DCBrick) Disable() (err error) {
	var buf bytes.Buffer

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

// Returns *true* if the driver chip is enabled, *false* otherwise.
func (device *DCBrick) IsEnabled() (enabled bool, err error) {
	var buf bytes.Buffer

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

// Sets the minimum voltage, below which the RegisterUnderVoltageCallback callback
// is triggered. The minimum possible value that works with the DC Brick is 6V.
// You can use this function to detect the discharge of a battery that is used
// to drive the motor. If you have a fixed power supply, you likely do not need
// this functionality.
func (device *DCBrick) SetMinimumVoltage(voltage uint16) (err error) {
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
func (device *DCBrick) GetMinimumVoltage() (voltage uint16, err error) {
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

// Sets the drive mode. Possible modes are:
//
// * 0 = Drive/Brake
// * 1 = Drive/Coast
//
// These modes are different kinds of motor controls.
//
// In Drive/Brake mode, the motor is always either driving or braking. There
// is no freewheeling. Advantages are: A more linear correlation between
// PWM and velocity, more exact accelerations and the possibility to drive
// with slower velocities.
//
// In Drive/Coast mode, the motor is always either driving or freewheeling.
// Advantages are: Less current consumption and less demands on the motor and
// driver chip.
//
// Associated constants:
//
//	* DriveModeDriveBrake
//	* DriveModeDriveCoast
func (device *DCBrick) SetDriveMode(mode DriveMode) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, mode)

	resultBytes, err := device.device.Set(uint8(FunctionSetDriveMode), buf.Bytes())
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

// Returns the drive mode, as set by SetDriveMode.
//
// Associated constants:
//
//	* DriveModeDriveBrake
//	* DriveModeDriveCoast
func (device *DCBrick) GetDriveMode() (mode DriveMode, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetDriveMode), buf.Bytes())
	if err != nil {
		return mode, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return mode, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return mode, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &mode)

	}

	return mode, nil
}

// Sets a period with which the RegisterCurrentVelocityCallback callback is triggered.
// A period of 0 turns the callback off.
func (device *DCBrick) SetCurrentVelocityPeriod(period uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period)

	resultBytes, err := device.device.Set(uint8(FunctionSetCurrentVelocityPeriod), buf.Bytes())
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

// Returns the period as set by SetCurrentVelocityPeriod.
func (device *DCBrick) GetCurrentVelocityPeriod() (period uint16, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetCurrentVelocityPeriod), buf.Bytes())
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
// .. versionadded:: 2.3.5$nbsp;(Firmware)
func (device *DCBrick) SetSPITFPBaudrateConfig(enableDynamicBaudrate bool, minimumDynamicBaudrate uint32) (err error) {
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
// .. versionadded:: 2.3.5$nbsp;(Firmware)
func (device *DCBrick) GetSPITFPBaudrateConfig() (enableDynamicBaudrate bool, minimumDynamicBaudrate uint32, err error) {
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
// .. versionadded:: 2.3.3$nbsp;(Firmware)
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
func (device *DCBrick) GetSendTimeoutCount(communicationMethod CommunicationMethod) (timeoutCount uint32, err error) {
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
// .. versionadded:: 2.3.3$nbsp;(Firmware)
func (device *DCBrick) SetSPITFPBaudrate(brickletPort rune, baudrate uint32) (err error) {
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
// .. versionadded:: 2.3.3$nbsp;(Firmware)
func (device *DCBrick) GetSPITFPBaudrate(brickletPort rune) (baudrate uint32, err error) {
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
// .. versionadded:: 2.3.3$nbsp;(Firmware)
func (device *DCBrick) GetSPITFPErrorCount(brickletPort rune) (errorCountACKChecksum uint32, errorCountMessageChecksum uint32, errorCountFrame uint32, errorCountOverflow uint32, err error) {
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
func (device *DCBrick) EnableStatusLED() (err error) {
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
func (device *DCBrick) DisableStatusLED() (err error) {
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
func (device *DCBrick) IsStatusLEDEnabled() (enabled bool, err error) {
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
func (device *DCBrick) GetProtocol1BrickletName(port rune) (protocolVersion uint8, firmwareVersion [3]uint8, name string, err error) {
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
func (device *DCBrick) GetChipTemperature() (temperature int16, err error) {
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
func (device *DCBrick) Reset() (err error) {
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
func (device *DCBrick) WriteBrickletPlugin(port rune, offset uint8, chunk [32]uint8) (err error) {
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
func (device *DCBrick) ReadBrickletPlugin(port rune, offset uint8) (chunk [32]uint8, err error) {
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
func (device *DCBrick) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
