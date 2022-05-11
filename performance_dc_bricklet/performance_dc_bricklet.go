/* ***********************************************************
 * This file was automatically generated on 2022-05-11.      *
 *                                                           *
 * Go Bindings Version 2.0.12                                *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


// Drives one brushed DC motor with up to 36V and 10A‍.
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/PerformanceDC_Bricklet_Go.html.
package performance_dc_bricklet

import (
	"encoding/binary"
	"bytes"
	"fmt"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionSetEnabled Function = 1
	FunctionGetEnabled Function = 2
	FunctionSetVelocity Function = 3
	FunctionGetVelocity Function = 4
	FunctionGetCurrentVelocity Function = 5
	FunctionSetMotion Function = 6
	FunctionGetMotion Function = 7
	FunctionFullBrake Function = 8
	FunctionSetDriveMode Function = 9
	FunctionGetDriveMode Function = 10
	FunctionSetPWMFrequency Function = 11
	FunctionGetPWMFrequency Function = 12
	FunctionGetPowerStatistics Function = 13
	FunctionSetThermalShutdown Function = 14
	FunctionGetThermalShutdown Function = 15
	FunctionSetGPIOConfiguration Function = 16
	FunctionGetGPIOConfiguration Function = 17
	FunctionSetGPIOAction Function = 18
	FunctionGetGPIOAction Function = 19
	FunctionGetGPIOState Function = 20
	FunctionSetErrorLEDConfig Function = 21
	FunctionGetErrorLEDConfig Function = 22
	FunctionSetCWLEDConfig Function = 23
	FunctionGetCWLEDConfig Function = 24
	FunctionSetCCWLEDConfig Function = 25
	FunctionGetCCWLEDConfig Function = 26
	FunctionSetGPIOLEDConfig Function = 27
	FunctionGetGPIOLEDConfig Function = 28
	FunctionSetEmergencyShutdownCallbackConfiguration Function = 29
	FunctionGetEmergencyShutdownCallbackConfiguration Function = 30
	FunctionSetVelocityReachedCallbackConfiguration Function = 31
	FunctionGetVelocityReachedCallbackConfiguration Function = 32
	FunctionSetCurrentVelocityCallbackConfiguration Function = 33
	FunctionGetCurrentVelocityCallbackConfiguration Function = 34
	FunctionGetSPITFPErrorCount Function = 234
	FunctionSetBootloaderMode Function = 235
	FunctionGetBootloaderMode Function = 236
	FunctionSetWriteFirmwarePointer Function = 237
	FunctionWriteFirmware Function = 238
	FunctionSetStatusLEDConfig Function = 239
	FunctionGetStatusLEDConfig Function = 240
	FunctionGetChipTemperature Function = 242
	FunctionReset Function = 243
	FunctionWriteUID Function = 248
	FunctionReadUID Function = 249
	FunctionGetIdentity Function = 255
	FunctionCallbackEmergencyShutdown Function = 35
	FunctionCallbackVelocityReached Function = 36
	FunctionCallbackCurrentVelocity Function = 37
	FunctionCallbackGPIOState Function = 38
)

type DriveMode = uint8

const (
	DriveModeDriveBrake DriveMode = 0
	DriveModeDriveCoast DriveMode = 1
)

type GPIOAction = uint32

const (
	GPIOActionNone GPIOAction = 0
	GPIOActionNormalStopRisingEdge GPIOAction = 1
	GPIOActionNormalStopFallingEdge GPIOAction = 2
	GPIOActionFullBrakeRisingEdge GPIOAction = 4
	GPIOActionFullBrakeFallingEdge GPIOAction = 8
	GPIOActionCallbackRisingEdge GPIOAction = 16
	GPIOActionCallbackFallingEdge GPIOAction = 32
)

type ErrorLEDConfig = uint8

const (
	ErrorLEDConfigOff ErrorLEDConfig = 0
	ErrorLEDConfigOn ErrorLEDConfig = 1
	ErrorLEDConfigShowHeartbeat ErrorLEDConfig = 2
	ErrorLEDConfigShowError ErrorLEDConfig = 3
)

type CWLEDConfig = uint8

const (
	CWLEDConfigOff CWLEDConfig = 0
	CWLEDConfigOn CWLEDConfig = 1
	CWLEDConfigShowHeartbeat CWLEDConfig = 2
	CWLEDConfigShowCWAsForward CWLEDConfig = 3
	CWLEDConfigShowCWAsBackward CWLEDConfig = 4
)

type CCWLEDConfig = uint8

const (
	CCWLEDConfigOff CCWLEDConfig = 0
	CCWLEDConfigOn CCWLEDConfig = 1
	CCWLEDConfigShowHeartbeat CCWLEDConfig = 2
	CCWLEDConfigShowCCWAsForward CCWLEDConfig = 3
	CCWLEDConfigShowCCWAsBackward CCWLEDConfig = 4
)

type GPIOLEDConfig = uint8

const (
	GPIOLEDConfigOff GPIOLEDConfig = 0
	GPIOLEDConfigOn GPIOLEDConfig = 1
	GPIOLEDConfigShowHeartbeat GPIOLEDConfig = 2
	GPIOLEDConfigShowGPIOActiveHigh GPIOLEDConfig = 3
	GPIOLEDConfigShowGPIOActiveLow GPIOLEDConfig = 4
)

type BootloaderMode = uint8

const (
	BootloaderModeBootloader BootloaderMode = 0
	BootloaderModeFirmware BootloaderMode = 1
	BootloaderModeBootloaderWaitForReboot BootloaderMode = 2
	BootloaderModeFirmwareWaitForReboot BootloaderMode = 3
	BootloaderModeFirmwareWaitForEraseAndReboot BootloaderMode = 4
)

type BootloaderStatus = uint8

const (
	BootloaderStatusOK BootloaderStatus = 0
	BootloaderStatusInvalidMode BootloaderStatus = 1
	BootloaderStatusNoChange BootloaderStatus = 2
	BootloaderStatusEntryFunctionNotPresent BootloaderStatus = 3
	BootloaderStatusDeviceIdentifierIncorrect BootloaderStatus = 4
	BootloaderStatusCRCMismatch BootloaderStatus = 5
)

type StatusLEDConfig = uint8

const (
	StatusLEDConfigOff StatusLEDConfig = 0
	StatusLEDConfigOn StatusLEDConfig = 1
	StatusLEDConfigShowHeartbeat StatusLEDConfig = 2
	StatusLEDConfigShowStatus StatusLEDConfig = 3
)

type PerformanceDCBricklet struct {
	device Device
}
const DeviceIdentifier = 2156
const DeviceDisplayName = "Performance DC Bricklet"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (PerformanceDCBricklet, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{ 2,0,1 }, uid, &internalIPCon, 0, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return PerformanceDCBricklet{}, err
	}
	dev.ResponseExpected[FunctionSetEnabled] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetEnabled] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetVelocity] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetVelocity] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetCurrentVelocity] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetMotion] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetMotion] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionFullBrake] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionSetDriveMode] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetDriveMode] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetPWMFrequency] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetPWMFrequency] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetPowerStatistics] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetThermalShutdown] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetThermalShutdown] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetGPIOConfiguration] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetGPIOConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetGPIOAction] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetGPIOAction] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetGPIOState] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetErrorLEDConfig] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetErrorLEDConfig] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetCWLEDConfig] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetCWLEDConfig] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetCCWLEDConfig] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetCCWLEDConfig] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetGPIOLEDConfig] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetGPIOLEDConfig] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetEmergencyShutdownCallbackConfiguration] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetEmergencyShutdownCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetVelocityReachedCallbackConfiguration] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetVelocityReachedCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetCurrentVelocityCallbackConfiguration] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetCurrentVelocityCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetSPITFPErrorCount] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetBootloaderMode] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetBootloaderMode] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetWriteFirmwarePointer] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionWriteFirmware] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetStatusLEDConfig] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetStatusLEDConfig] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetChipTemperature] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionReset] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionWriteUID] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionReadUID] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetIdentity] = ResponseExpectedFlagAlwaysTrue;
	return PerformanceDCBricklet{dev}, nil
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
func (device *PerformanceDCBricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *PerformanceDCBricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *PerformanceDCBricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *PerformanceDCBricklet) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// This callback is triggered if either the current consumption
// is too high or the temperature of the driver chip is too high
// (above 150°C) or the user defined thermal shutdown is triggered (see SetThermalShutdown).
// n case of a voltage below 6V (input voltage) this
// callback is triggered as well.
// 
// If this callback is triggered, the driver chip gets disabled at the same time.
// That means, SetEnabled has to be called to drive the motor again.
func (device *PerformanceDCBricklet) RegisterEmergencyShutdownCallback(fn func()) uint64 {
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
func (device *PerformanceDCBricklet) DeregisterEmergencyShutdownCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackEmergencyShutdown), registrationId)
}


// This callback is triggered whenever a set velocity is reached. For example:
// If a velocity of 0 is present, acceleration is set to 5000 and velocity
// to 10000, the RegisterVelocityReachedCallback callback will be triggered after about
// 2 seconds, when the set velocity is actually reached.
// 
// Note
//  Since we can't get any feedback from the DC motor, this only works if the
//  acceleration (see SetMotion) is set smaller or equal to the
//  maximum acceleration of the motor. Otherwise the motor will lag behind the
//  control value and the callback will be triggered too early.
func (device *PerformanceDCBricklet) RegisterVelocityReachedCallback(fn func(int16)) uint64 {
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
func (device *PerformanceDCBricklet) DeregisterVelocityReachedCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackVelocityReached), registrationId)
}


// This callback is triggered with the period that is set by
// SetCurrentVelocityCallbackConfiguration. The parameter is the *current*
// velocity used by the motor.
// 
// The RegisterCurrentVelocityCallback callback is only triggered after the set period
// if there is a change in the velocity.
func (device *PerformanceDCBricklet) RegisterCurrentVelocityCallback(fn func(int16)) uint64 {
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
func (device *PerformanceDCBricklet) DeregisterCurrentVelocityCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackCurrentVelocity), registrationId)
}


// This callback is triggered by GPIO changes if it is activated through SetGPIOAction.
// 
// .. versionadded:: 2.0.1$nbsp;(Plugin)
func (device *PerformanceDCBricklet) RegisterGPIOStateCallback(fn func([2]bool)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 9 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var gpioState [2]bool
		binary.Read(buf, binary.LittleEndian, &gpioState)
		fn(gpioState)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackGPIOState), wrapper)
}

// Remove a registered GPIO State callback.
func (device *PerformanceDCBricklet) DeregisterGPIOStateCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackGPIOState), registrationId)
}


// Enables/Disables the driver chip. The driver parameters can be configured
// (velocity, acceleration, etc) before it is enabled.
func (device *PerformanceDCBricklet) SetEnabled(enabled bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, enabled);

	resultBytes, err := device.device.Set(uint8(FunctionSetEnabled), buf.Bytes())
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
func (device *PerformanceDCBricklet) GetEnabled() (enabled bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetEnabled), buf.Bytes())
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

// Sets the velocity of the motor. Whereas -32767 is full speed backward,
// 0 is stop and 32767 is full speed forward. Depending on the
// acceleration (see SetMotion), the motor is not immediately
// brought to the velocity but smoothly accelerated.
// 
// The velocity describes the duty cycle of the PWM with which the motor is
// controlled, e.g. a velocity of 3277 sets a PWM with a 10% duty cycle.
// You can not only control the duty cycle of the PWM but also the frequency,
// see SetPWMFrequency.
func (device *PerformanceDCBricklet) SetVelocity(velocity int16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, velocity);

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
func (device *PerformanceDCBricklet) GetVelocity() (velocity int16, err error) {
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
func (device *PerformanceDCBricklet) GetCurrentVelocity() (velocity int16, err error) {
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

// Sets the acceleration and deceleration of the motor. It is given in *velocity/s*.
// An acceleration of 10000 means, that every second the velocity is increased
// by 10000 (or about 30% duty cycle).
// 
// For example: If the current velocity is 0 and you want to accelerate to a
// velocity of 16000 (about 50% duty cycle) in 10 seconds, you should set
// an acceleration of 1600.
// 
// If acceleration and deceleration is set to 0, there is no speed ramping, i.e. a
// new velocity is immediately given to the motor.
func (device *PerformanceDCBricklet) SetMotion(acceleration uint16, deceleration uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, acceleration);
	binary.Write(&buf, binary.LittleEndian, deceleration);

	resultBytes, err := device.device.Set(uint8(FunctionSetMotion), buf.Bytes())
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

// Returns the acceleration/deceleration as set by SetMotion.
func (device *PerformanceDCBricklet) GetMotion() (acceleration uint16, deceleration uint16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetMotion), buf.Bytes())
	if err != nil {
		return acceleration, deceleration, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 12 {
			return acceleration, deceleration, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		if header.ErrorCode != 0 {
			return acceleration, deceleration, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &acceleration)
		binary.Read(resultBuf, binary.LittleEndian, &deceleration)

	}

	return acceleration, deceleration, nil
}

// Executes an active full brake.
// 
// Warning
//  This function is for emergency purposes,
//  where an immediate brake is necessary. Depending on the current velocity and
//  the strength of the motor, a full brake can be quite violent.
// 
// Call SetVelocity with 0 if you just want to stop the motor.
func (device *PerformanceDCBricklet) FullBrake() (err error) {
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
func (device *PerformanceDCBricklet) SetDriveMode(mode DriveMode) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, mode);

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
func (device *PerformanceDCBricklet) GetDriveMode() (mode DriveMode, err error) {
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

// Sets the frequency of the PWM with which the motor is driven.
// Often a high frequency
// is less noisy and the motor runs smoother. However, with a low frequency
// there are less switches and therefore fewer switching losses. Also with
// most motors lower frequencies enable higher torque.
// 
// If you have no idea what all this means, just ignore this function and use
// the default frequency, it will very likely work fine.
func (device *PerformanceDCBricklet) SetPWMFrequency(frequency uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, frequency);

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
func (device *PerformanceDCBricklet) GetPWMFrequency() (frequency uint16, err error) {
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

// Returns input voltage, current usage and temperature of the driver.
func (device *PerformanceDCBricklet) GetPowerStatistics() (voltage uint16, current uint16, temperature int16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetPowerStatistics), buf.Bytes())
	if err != nil {
		return voltage, current, temperature, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 14 {
			return voltage, current, temperature, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 14)
		}

		if header.ErrorCode != 0 {
			return voltage, current, temperature, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &voltage)
		binary.Read(resultBuf, binary.LittleEndian, &current)
		binary.Read(resultBuf, binary.LittleEndian, &temperature)

	}

	return voltage, current, temperature, nil
}

// Sets a temperature threshold that is used for thermal shutdown.
// 
// Additionally to this user defined threshold the driver chip will shut down at a
// temperature of 150°C.
// 
// If a thermal shutdown is triggered the driver is disabled and has to be
// explicitly re-enabled with SetEnabled.
func (device *PerformanceDCBricklet) SetThermalShutdown(temperature uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, temperature);

	resultBytes, err := device.device.Set(uint8(FunctionSetThermalShutdown), buf.Bytes())
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

// Returns the thermal shutdown temperature as set by SetThermalShutdown.
func (device *PerformanceDCBricklet) GetThermalShutdown() (temperature uint8, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetThermalShutdown), buf.Bytes())
	if err != nil {
		return temperature, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return temperature, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return temperature, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &temperature)

	}

	return temperature, nil
}

// Sets the GPIO configuration for the given channel.
// You can configure a debounce and the deceleration that is used if the action is
// configured as ``normal stop``. See SetGPIOAction.
func (device *PerformanceDCBricklet) SetGPIOConfiguration(channel uint8, debounce uint16, stopDeceleration uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, channel);
	binary.Write(&buf, binary.LittleEndian, debounce);
	binary.Write(&buf, binary.LittleEndian, stopDeceleration);

	resultBytes, err := device.device.Set(uint8(FunctionSetGPIOConfiguration), buf.Bytes())
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

// Returns the GPIO configuration for a channel as set by SetGPIOConfiguration.
func (device *PerformanceDCBricklet) GetGPIOConfiguration(channel uint8) (debounce uint16, stopDeceleration uint16, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, channel);

	resultBytes, err := device.device.Get(uint8(FunctionGetGPIOConfiguration), buf.Bytes())
	if err != nil {
		return debounce, stopDeceleration, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 12 {
			return debounce, stopDeceleration, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		if header.ErrorCode != 0 {
			return debounce, stopDeceleration, DeviceError(header.ErrorCode)
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
func (device *PerformanceDCBricklet) SetGPIOAction(channel uint8, action GPIOAction) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, channel);
	binary.Write(&buf, binary.LittleEndian, action);

	resultBytes, err := device.device.Set(uint8(FunctionSetGPIOAction), buf.Bytes())
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
func (device *PerformanceDCBricklet) GetGPIOAction(channel uint8) (action GPIOAction, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, channel);

	resultBytes, err := device.device.Get(uint8(FunctionGetGPIOAction), buf.Bytes())
	if err != nil {
		return action, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 12 {
			return action, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		if header.ErrorCode != 0 {
			return action, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &action)

	}

	return action, nil
}

// Returns the GPIO state for both channels. True if the state is ``high`` and
// false if the state is ``low``.
func (device *PerformanceDCBricklet) GetGPIOState() (gpioState [2]bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetGPIOState), buf.Bytes())
	if err != nil {
		return gpioState, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return gpioState, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return gpioState, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &gpioState)

	}

	return gpioState, nil
}

// Configures the error LED to be either turned off, turned on, blink in
// heartbeat mode or show an error.
// 
// If the LED is configured to show errors it has three different states:
// 
// * Off: No error present.
// * 1s interval blinking: Input voltage too low (below 6V).
// * 250ms interval blinking: Overtemperature or overcurrent.
//
// Associated constants:
//
//	* ErrorLEDConfigOff
//	* ErrorLEDConfigOn
//	* ErrorLEDConfigShowHeartbeat
//	* ErrorLEDConfigShowError
func (device *PerformanceDCBricklet) SetErrorLEDConfig(config ErrorLEDConfig) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, config);

	resultBytes, err := device.device.Set(uint8(FunctionSetErrorLEDConfig), buf.Bytes())
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

// Returns the LED configuration as set by SetErrorLEDConfig
//
// Associated constants:
//
//	* ErrorLEDConfigOff
//	* ErrorLEDConfigOn
//	* ErrorLEDConfigShowHeartbeat
//	* ErrorLEDConfigShowError
func (device *PerformanceDCBricklet) GetErrorLEDConfig() (config ErrorLEDConfig, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetErrorLEDConfig), buf.Bytes())
	if err != nil {
		return config, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return config, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return config, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &config)

	}

	return config, nil
}

// Configures the CW LED to be either turned off, turned on, blink in
// heartbeat mode or if the motor turn clockwise.
//
// Associated constants:
//
//	* CWLEDConfigOff
//	* CWLEDConfigOn
//	* CWLEDConfigShowHeartbeat
//	* CWLEDConfigShowCWAsForward
//	* CWLEDConfigShowCWAsBackward
func (device *PerformanceDCBricklet) SetCWLEDConfig(config CWLEDConfig) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, config);

	resultBytes, err := device.device.Set(uint8(FunctionSetCWLEDConfig), buf.Bytes())
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

// Returns the LED configuration as set by SetCWLEDConfig
//
// Associated constants:
//
//	* CWLEDConfigOff
//	* CWLEDConfigOn
//	* CWLEDConfigShowHeartbeat
//	* CWLEDConfigShowCWAsForward
//	* CWLEDConfigShowCWAsBackward
func (device *PerformanceDCBricklet) GetCWLEDConfig() (config CWLEDConfig, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetCWLEDConfig), buf.Bytes())
	if err != nil {
		return config, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return config, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return config, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &config)

	}

	return config, nil
}

// Configures the CCW LED to be either turned off, turned on, blink in
// heartbeat mode or if the motor turn counter-clockwise.
//
// Associated constants:
//
//	* CCWLEDConfigOff
//	* CCWLEDConfigOn
//	* CCWLEDConfigShowHeartbeat
//	* CCWLEDConfigShowCCWAsForward
//	* CCWLEDConfigShowCCWAsBackward
func (device *PerformanceDCBricklet) SetCCWLEDConfig(config CCWLEDConfig) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, config);

	resultBytes, err := device.device.Set(uint8(FunctionSetCCWLEDConfig), buf.Bytes())
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

// Returns the LED configuration as set by SetCCWLEDConfig
//
// Associated constants:
//
//	* CCWLEDConfigOff
//	* CCWLEDConfigOn
//	* CCWLEDConfigShowHeartbeat
//	* CCWLEDConfigShowCCWAsForward
//	* CCWLEDConfigShowCCWAsBackward
func (device *PerformanceDCBricklet) GetCCWLEDConfig() (config CCWLEDConfig, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetCCWLEDConfig), buf.Bytes())
	if err != nil {
		return config, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return config, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return config, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &config)

	}

	return config, nil
}

// Configures the GPIO LED to be either turned off, turned on, blink in
// heartbeat mode or the GPIO state.
// 
// The GPIO LED can be configured for both channels.
//
// Associated constants:
//
//	* GPIOLEDConfigOff
//	* GPIOLEDConfigOn
//	* GPIOLEDConfigShowHeartbeat
//	* GPIOLEDConfigShowGPIOActiveHigh
//	* GPIOLEDConfigShowGPIOActiveLow
func (device *PerformanceDCBricklet) SetGPIOLEDConfig(channel uint8, config GPIOLEDConfig) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, channel);
	binary.Write(&buf, binary.LittleEndian, config);

	resultBytes, err := device.device.Set(uint8(FunctionSetGPIOLEDConfig), buf.Bytes())
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

// Returns the LED configuration as set by SetGPIOLEDConfig
//
// Associated constants:
//
//	* GPIOLEDConfigOff
//	* GPIOLEDConfigOn
//	* GPIOLEDConfigShowHeartbeat
//	* GPIOLEDConfigShowGPIOActiveHigh
//	* GPIOLEDConfigShowGPIOActiveLow
func (device *PerformanceDCBricklet) GetGPIOLEDConfig(channel uint8) (config GPIOLEDConfig, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, channel);

	resultBytes, err := device.device.Get(uint8(FunctionGetGPIOLEDConfig), buf.Bytes())
	if err != nil {
		return config, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return config, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return config, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &config)

	}

	return config, nil
}

// Enable/Disable RegisterEmergencyShutdownCallback callback.
func (device *PerformanceDCBricklet) SetEmergencyShutdownCallbackConfiguration(enabled bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, enabled);

	resultBytes, err := device.device.Set(uint8(FunctionSetEmergencyShutdownCallbackConfiguration), buf.Bytes())
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

// Returns the callback configuration as set by
// SetEmergencyShutdownCallbackConfiguration.
func (device *PerformanceDCBricklet) GetEmergencyShutdownCallbackConfiguration() (enabled bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetEmergencyShutdownCallbackConfiguration), buf.Bytes())
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

// Enable/Disable RegisterVelocityReachedCallback callback.
func (device *PerformanceDCBricklet) SetVelocityReachedCallbackConfiguration(enabled bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, enabled);

	resultBytes, err := device.device.Set(uint8(FunctionSetVelocityReachedCallbackConfiguration), buf.Bytes())
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

// Returns the callback configuration as set by
// SetVelocityReachedCallbackConfiguration.
func (device *PerformanceDCBricklet) GetVelocityReachedCallbackConfiguration() (enabled bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetVelocityReachedCallbackConfiguration), buf.Bytes())
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

// The period is the period with which the RegisterCurrentVelocityCallback
// callback is triggered periodically. A value of 0 turns the callback off.
// 
// If the `value has to change`-parameter is set to true, the callback is only
// triggered after the value has changed. If the value didn't change within the
// period, the callback is triggered immediately on change.
// 
// If it is set to false, the callback is continuously triggered with the period,
// independent of the value.
func (device *PerformanceDCBricklet) SetCurrentVelocityCallbackConfiguration(period uint32, valueHasToChange bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);
	binary.Write(&buf, binary.LittleEndian, valueHasToChange);

	resultBytes, err := device.device.Set(uint8(FunctionSetCurrentVelocityCallbackConfiguration), buf.Bytes())
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

// Returns the callback configuration as set by
// SetCurrentVelocityCallbackConfiguration.
func (device *PerformanceDCBricklet) GetCurrentVelocityCallbackConfiguration() (period uint32, valueHasToChange bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetCurrentVelocityCallbackConfiguration), buf.Bytes())
	if err != nil {
		return period, valueHasToChange, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 13 {
			return period, valueHasToChange, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 13)
		}

		if header.ErrorCode != 0 {
			return period, valueHasToChange, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &period)
		binary.Read(resultBuf, binary.LittleEndian, &valueHasToChange)

	}

	return period, valueHasToChange, nil
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
func (device *PerformanceDCBricklet) GetSPITFPErrorCount() (errorCountAckChecksum uint32, errorCountMessageChecksum uint32, errorCountFrame uint32, errorCountOverflow uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetSPITFPErrorCount), buf.Bytes())
	if err != nil {
		return errorCountAckChecksum, errorCountMessageChecksum, errorCountFrame, errorCountOverflow, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 24 {
			return errorCountAckChecksum, errorCountMessageChecksum, errorCountFrame, errorCountOverflow, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 24)
		}

		if header.ErrorCode != 0 {
			return errorCountAckChecksum, errorCountMessageChecksum, errorCountFrame, errorCountOverflow, DeviceError(header.ErrorCode)
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
func (device *PerformanceDCBricklet) SetBootloaderMode(mode BootloaderMode) (status BootloaderStatus, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, mode);

	resultBytes, err := device.device.Get(uint8(FunctionSetBootloaderMode), buf.Bytes())
	if err != nil {
		return status, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return status, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return status, DeviceError(header.ErrorCode)
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
func (device *PerformanceDCBricklet) GetBootloaderMode() (mode BootloaderMode, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetBootloaderMode), buf.Bytes())
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

// Sets the firmware pointer for WriteFirmware. The pointer has
// to be increased by chunks of size 64. The data is written to flash
// every 4 chunks (which equals to one page of size 256).
// 
// This function is used by Brick Viewer during flashing. It should not be
// necessary to call it in a normal user program.
func (device *PerformanceDCBricklet) SetWriteFirmwarePointer(pointer uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, pointer);

	resultBytes, err := device.device.Set(uint8(FunctionSetWriteFirmwarePointer), buf.Bytes())
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

// Writes 64 Bytes of firmware at the position as written by
// SetWriteFirmwarePointer before. The firmware is written
// to flash every 4 chunks.
// 
// You can only write firmware in bootloader mode.
// 
// This function is used by Brick Viewer during flashing. It should not be
// necessary to call it in a normal user program.
func (device *PerformanceDCBricklet) WriteFirmware(data [64]uint8) (status uint8, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, data);

	resultBytes, err := device.device.Get(uint8(FunctionWriteFirmware), buf.Bytes())
	if err != nil {
		return status, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return status, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return status, DeviceError(header.ErrorCode)
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
func (device *PerformanceDCBricklet) SetStatusLEDConfig(config StatusLEDConfig) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, config);

	resultBytes, err := device.device.Set(uint8(FunctionSetStatusLEDConfig), buf.Bytes())
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

// Returns the configuration as set by SetStatusLEDConfig
//
// Associated constants:
//
//	* StatusLEDConfigOff
//	* StatusLEDConfigOn
//	* StatusLEDConfigShowHeartbeat
//	* StatusLEDConfigShowStatus
func (device *PerformanceDCBricklet) GetStatusLEDConfig() (config StatusLEDConfig, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetStatusLEDConfig), buf.Bytes())
	if err != nil {
		return config, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return config, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return config, DeviceError(header.ErrorCode)
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
func (device *PerformanceDCBricklet) GetChipTemperature() (temperature int16, err error) {
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

// Calling this function will reset the Bricklet. All configurations
// will be lost.
// 
// After a reset you have to create new device objects,
// calling functions on the existing ones will result in
// undefined behavior!
func (device *PerformanceDCBricklet) Reset() (err error) {
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

// Writes a new UID into flash. If you want to set a new UID
// you have to decode the Base58 encoded UID string into an
// integer first.
// 
// We recommend that you use Brick Viewer to change the UID.
func (device *PerformanceDCBricklet) WriteUID(uid uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, uid);

	resultBytes, err := device.device.Set(uint8(FunctionWriteUID), buf.Bytes())
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

// Returns the current UID as an integer. Encode as
// Base58 to get the usual string version.
func (device *PerformanceDCBricklet) ReadUID() (uid uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionReadUID), buf.Bytes())
	if err != nil {
		return uid, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 12 {
			return uid, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		if header.ErrorCode != 0 {
			return uid, DeviceError(header.ErrorCode)
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
func (device *PerformanceDCBricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
