/* ***********************************************************
 * This file was automatically generated on 2022-05-11.      *
 *                                                           *
 * Go Bindings Version 2.0.12                                *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


// Drives up to 10 RC Servos.
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/ServoV2_Bricklet_Go.html.
package servo_v2_bricklet

import (
	"encoding/binary"
	"bytes"
	"fmt"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionGetStatus Function = 1
	FunctionSetEnable Function = 2
	FunctionGetEnabled Function = 3
	FunctionSetPosition Function = 4
	FunctionGetPosition Function = 5
	FunctionGetCurrentPosition Function = 6
	FunctionGetCurrentVelocity Function = 7
	FunctionSetMotionConfiguration Function = 8
	FunctionGetMotionConfiguration Function = 9
	FunctionSetPulseWidth Function = 10
	FunctionGetPulseWidth Function = 11
	FunctionSetDegree Function = 12
	FunctionGetDegree Function = 13
	FunctionSetPeriod Function = 14
	FunctionGetPeriod Function = 15
	FunctionGetServoCurrent Function = 16
	FunctionSetServoCurrentConfiguration Function = 17
	FunctionGetServoCurrentConfiguration Function = 18
	FunctionSetInputVoltageConfiguration Function = 19
	FunctionGetInputVoltageConfiguration Function = 20
	FunctionGetOverallCurrent Function = 21
	FunctionGetInputVoltage Function = 22
	FunctionSetCurrentCalibration Function = 23
	FunctionGetCurrentCalibration Function = 24
	FunctionSetPositionReachedCallbackConfiguration Function = 25
	FunctionGetPositionReachedCallbackConfiguration Function = 26
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
	FunctionCallbackPositionReached Function = 27
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

type ServoV2Bricklet struct {
	device Device
}
const DeviceIdentifier = 2157
const DeviceDisplayName = "Servo Bricklet 2.0"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (ServoV2Bricklet, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{ 2,0,0 }, uid, &internalIPCon, 0, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return ServoV2Bricklet{}, err
	}
	dev.ResponseExpected[FunctionGetStatus] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetEnable] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetEnabled] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetPosition] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetPosition] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetCurrentPosition] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetCurrentVelocity] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetMotionConfiguration] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetMotionConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetPulseWidth] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetPulseWidth] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetDegree] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetDegree] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetPeriod] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetPeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetServoCurrent] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetServoCurrentConfiguration] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetServoCurrentConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetInputVoltageConfiguration] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetInputVoltageConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetOverallCurrent] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetInputVoltage] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetCurrentCalibration] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetCurrentCalibration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetPositionReachedCallbackConfiguration] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetPositionReachedCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue;
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
	return ServoV2Bricklet{dev}, nil
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
func (device *ServoV2Bricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *ServoV2Bricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *ServoV2Bricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *ServoV2Bricklet) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// This callback is triggered when a position set by SetPosition
// is reached. If the new position matches the current position then the
// callback is not triggered, because the servo didn't move.
// The parameters are the servo and the position that is reached.
// 
// You can enable this callback with SetPositionReachedCallbackConfiguration.
// 
// Note
//  Since we can't get any feedback from the servo, this only works if the
//  velocity (see SetMotionConfiguration) is set smaller or equal to the
//  maximum velocity of the servo. Otherwise the servo will lag behind the
//  control value and the callback will be triggered too early.
func (device *ServoV2Bricklet) RegisterPositionReachedCallback(fn func(uint16, int16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 12 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var servoChannel uint16
		var position int16
		binary.Read(buf, binary.LittleEndian, &servoChannel)
		binary.Read(buf, binary.LittleEndian, &position)
		fn(servoChannel, position)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackPositionReached), wrapper)
}

// Remove a registered Position Reached callback.
func (device *ServoV2Bricklet) DeregisterPositionReachedCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackPositionReached), registrationId)
}


// Returns the status information of the Servo Bricklet 2.0.
// 
// The status includes
// 
// * for each channel if it is enabled or disabled,
// * for each channel the current position,
// * for each channel the current velocity,
// * for each channel the current usage and
// * the input voltage.
// 
// Please note that the position and the velocity is a snapshot of the
// current position and velocity of the servo in motion.
func (device *ServoV2Bricklet) GetStatus() (enabled [10]bool, currentPosition [10]int16, currentVelocity [10]int16, current [10]uint16, inputVoltage uint16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetStatus), buf.Bytes())
	if err != nil {
		return enabled, currentPosition, currentVelocity, current, inputVoltage, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 72 {
			return enabled, currentPosition, currentVelocity, current, inputVoltage, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 72)
		}

		if header.ErrorCode != 0 {
			return enabled, currentPosition, currentVelocity, current, inputVoltage, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &enabled)
		binary.Read(resultBuf, binary.LittleEndian, &currentPosition)
		binary.Read(resultBuf, binary.LittleEndian, &currentVelocity)
		binary.Read(resultBuf, binary.LittleEndian, &current)
		binary.Read(resultBuf, binary.LittleEndian, &inputVoltage)

	}

	return enabled, currentPosition, currentVelocity, current, inputVoltage, nil
}

// Enables a servo channel (0 to 9). If a servo is enabled, the configured position,
// velocity, acceleration, etc. are applied immediately.
func (device *ServoV2Bricklet) SetEnable(servoChannel uint16, enable bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, servoChannel);
	binary.Write(&buf, binary.LittleEndian, enable);

	resultBytes, err := device.device.Set(uint8(FunctionSetEnable), buf.Bytes())
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

// Returns *true* if the specified servo channel is enabled, *false* otherwise.
func (device *ServoV2Bricklet) GetEnabled(servoChannel uint16) (enable bool, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, servoChannel);

	resultBytes, err := device.device.Get(uint8(FunctionGetEnabled), buf.Bytes())
	if err != nil {
		return enable, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return enable, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return enable, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &enable)

	}

	return enable, nil
}

// Sets the position in °/100 for the specified servo channel.
// 
// The default range of the position is -9000 to 9000, but it can be specified
// according to your servo with SetDegree.
// 
// If you want to control a linear servo or RC brushless motor controller or
// similar with the Servo Brick, you can also define lengths or speeds with
// SetDegree.
func (device *ServoV2Bricklet) SetPosition(servoChannel uint16, position int16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, servoChannel);
	binary.Write(&buf, binary.LittleEndian, position);

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

// Returns the position of the specified servo channel as set by SetPosition.
func (device *ServoV2Bricklet) GetPosition(servoChannel uint16) (position int16, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, servoChannel);

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

// Returns the *current* position of the specified servo channel. This may not be the
// value of SetPosition if the servo is currently approaching a
// position goal.
func (device *ServoV2Bricklet) GetCurrentPosition(servoChannel uint16) (position int16, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, servoChannel);

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

// Returns the *current* velocity of the specified servo channel. This may not be the
// velocity specified by SetMotionConfiguration. if the servo is
// currently approaching a velocity goal.
func (device *ServoV2Bricklet) GetCurrentVelocity(servoChannel uint16) (velocity uint16, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, servoChannel);

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

// Sets the maximum velocity of the specified servo channel in °/100s as well as
// the acceleration and deceleration in °/100s²
// 
// With a velocity of 0 °/100s the position will be set immediately (no velocity).
// 
// With an acc-/deceleration of 0 °/100s² the velocity will be set immediately
// (no acc-/deceleration).
func (device *ServoV2Bricklet) SetMotionConfiguration(servoChannel uint16, velocity uint32, acceleration uint32, deceleration uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, servoChannel);
	binary.Write(&buf, binary.LittleEndian, velocity);
	binary.Write(&buf, binary.LittleEndian, acceleration);
	binary.Write(&buf, binary.LittleEndian, deceleration);

	resultBytes, err := device.device.Set(uint8(FunctionSetMotionConfiguration), buf.Bytes())
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

// Returns the motion configuration as set by SetMotionConfiguration.
func (device *ServoV2Bricklet) GetMotionConfiguration(servoChannel uint16) (velocity uint32, acceleration uint32, deceleration uint32, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, servoChannel);

	resultBytes, err := device.device.Get(uint8(FunctionGetMotionConfiguration), buf.Bytes())
	if err != nil {
		return velocity, acceleration, deceleration, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 20 {
			return velocity, acceleration, deceleration, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 20)
		}

		if header.ErrorCode != 0 {
			return velocity, acceleration, deceleration, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &velocity)
		binary.Read(resultBuf, binary.LittleEndian, &acceleration)
		binary.Read(resultBuf, binary.LittleEndian, &deceleration)

	}

	return velocity, acceleration, deceleration, nil
}

// Sets the minimum and maximum pulse width of the specified servo channel in µs.
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
// Both values have a range from 1 to 65535 (unsigned 16-bit integer). The
// minimum must be smaller than the maximum.
// 
// The default values are 1000µs (1ms) and 2000µs (2ms) for minimum and
// maximum pulse width.
func (device *ServoV2Bricklet) SetPulseWidth(servoChannel uint16, min uint32, max uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, servoChannel);
	binary.Write(&buf, binary.LittleEndian, min);
	binary.Write(&buf, binary.LittleEndian, max);

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

// Returns the minimum and maximum pulse width for the specified servo channel as set by
// SetPulseWidth.
func (device *ServoV2Bricklet) GetPulseWidth(servoChannel uint16) (min uint32, max uint32, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, servoChannel);

	resultBytes, err := device.device.Get(uint8(FunctionGetPulseWidth), buf.Bytes())
	if err != nil {
		return min, max, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 16 {
			return min, max, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 16)
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

// Sets the minimum and maximum degree for the specified servo channel (by default
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
// Both values have a possible range from -32767 to 32767
// (signed 16-bit integer). The minimum must be smaller than the maximum.
// 
// The default values are -9000 and 9000 for the minimum and maximum degree.
func (device *ServoV2Bricklet) SetDegree(servoChannel uint16, min int16, max int16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, servoChannel);
	binary.Write(&buf, binary.LittleEndian, min);
	binary.Write(&buf, binary.LittleEndian, max);

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

// Returns the minimum and maximum degree for the specified servo channel as set by
// SetDegree.
func (device *ServoV2Bricklet) GetDegree(servoChannel uint16) (min int16, max int16, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, servoChannel);

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

// Sets the period of the specified servo channel in µs.
// 
// Usually, servos are controlled with a
// https://en.wikipedia.org/wiki/Pulse-width_modulation. Different
// servos expect PWMs with different periods. Most servos run well with a
// period of about 20ms.
// 
// If your servo comes with a datasheet that specifies a period, you should
// set it accordingly. If you don't have a datasheet and you have no idea
// what the correct period is, the default value (19.5ms) will most likely
// work fine.
// 
// The minimum possible period is 1µs and the maximum is 1000000µs.
// 
// The default value is 19.5ms (19500µs).
func (device *ServoV2Bricklet) SetPeriod(servoChannel uint16, period uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, servoChannel);
	binary.Write(&buf, binary.LittleEndian, period);

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

// Returns the period for the specified servo channel as set by SetPeriod.
func (device *ServoV2Bricklet) GetPeriod(servoChannel uint16) (period uint32, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, servoChannel);

	resultBytes, err := device.device.Get(uint8(FunctionGetPeriod), buf.Bytes())
	if err != nil {
		return period, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 12 {
			return period, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		if header.ErrorCode != 0 {
			return period, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &period)

	}

	return period, nil
}

// Returns the current consumption of the specified servo channel in mA.
func (device *ServoV2Bricklet) GetServoCurrent(servoChannel uint16) (current uint16, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, servoChannel);

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

// Sets the averaging duration of the current measurement for the specified servo channel in ms.
func (device *ServoV2Bricklet) SetServoCurrentConfiguration(servoChannel uint16, averagingDuration uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, servoChannel);
	binary.Write(&buf, binary.LittleEndian, averagingDuration);

	resultBytes, err := device.device.Set(uint8(FunctionSetServoCurrentConfiguration), buf.Bytes())
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

// Returns the servo current configuration for the specified servo channel as set
// by SetServoCurrentConfiguration.
func (device *ServoV2Bricklet) GetServoCurrentConfiguration(servoChannel uint16) (averagingDuration uint8, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, servoChannel);

	resultBytes, err := device.device.Get(uint8(FunctionGetServoCurrentConfiguration), buf.Bytes())
	if err != nil {
		return averagingDuration, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return averagingDuration, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return averagingDuration, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &averagingDuration)

	}

	return averagingDuration, nil
}

// Sets the averaging duration of the input voltage measurement for the specified servo channel in ms.
func (device *ServoV2Bricklet) SetInputVoltageConfiguration(averagingDuration uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, averagingDuration);

	resultBytes, err := device.device.Set(uint8(FunctionSetInputVoltageConfiguration), buf.Bytes())
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

// Returns the input voltage configuration as set by SetInputVoltageConfiguration.
func (device *ServoV2Bricklet) GetInputVoltageConfiguration() (averagingDuration uint8, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetInputVoltageConfiguration), buf.Bytes())
	if err != nil {
		return averagingDuration, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return averagingDuration, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return averagingDuration, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &averagingDuration)

	}

	return averagingDuration, nil
}

// Returns the current consumption of all servos together in mA.
func (device *ServoV2Bricklet) GetOverallCurrent() (current uint16, err error) {
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

// Returns the input voltage in mV. The input voltage is
// given via the black power input connector on the Servo Brick.
func (device *ServoV2Bricklet) GetInputVoltage() (voltage uint16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetInputVoltage), buf.Bytes())
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

// Sets an offset value (in mA) for each channel.
// 
// Note: On delivery the Servo Bricklet 2.0 is already calibrated.
func (device *ServoV2Bricklet) SetCurrentCalibration(offset [10]int16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, offset);

	resultBytes, err := device.device.Set(uint8(FunctionSetCurrentCalibration), buf.Bytes())
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

// Returns the current calibration as set by SetCurrentCalibration.
func (device *ServoV2Bricklet) GetCurrentCalibration() (offset [10]int16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetCurrentCalibration), buf.Bytes())
	if err != nil {
		return offset, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 28 {
			return offset, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 28)
		}

		if header.ErrorCode != 0 {
			return offset, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &offset)

	}

	return offset, nil
}

// Enable/Disable RegisterPositionReachedCallback callback.
func (device *ServoV2Bricklet) SetPositionReachedCallbackConfiguration(servoChannel uint16, enabled bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, servoChannel);
	binary.Write(&buf, binary.LittleEndian, enabled);

	resultBytes, err := device.device.Set(uint8(FunctionSetPositionReachedCallbackConfiguration), buf.Bytes())
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
// SetPositionReachedCallbackConfiguration.
func (device *ServoV2Bricklet) GetPositionReachedCallbackConfiguration(servoChannel uint16) (enabled bool, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, servoChannel);

	resultBytes, err := device.device.Get(uint8(FunctionGetPositionReachedCallbackConfiguration), buf.Bytes())
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
func (device *ServoV2Bricklet) GetSPITFPErrorCount() (errorCountAckChecksum uint32, errorCountMessageChecksum uint32, errorCountFrame uint32, errorCountOverflow uint32, err error) {
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
func (device *ServoV2Bricklet) SetBootloaderMode(mode BootloaderMode) (status BootloaderStatus, err error) {
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
func (device *ServoV2Bricklet) GetBootloaderMode() (mode BootloaderMode, err error) {
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
func (device *ServoV2Bricklet) SetWriteFirmwarePointer(pointer uint32) (err error) {
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
func (device *ServoV2Bricklet) WriteFirmware(data [64]uint8) (status uint8, err error) {
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
func (device *ServoV2Bricklet) SetStatusLEDConfig(config StatusLEDConfig) (err error) {
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
func (device *ServoV2Bricklet) GetStatusLEDConfig() (config StatusLEDConfig, err error) {
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
func (device *ServoV2Bricklet) GetChipTemperature() (temperature int16, err error) {
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
func (device *ServoV2Bricklet) Reset() (err error) {
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
func (device *ServoV2Bricklet) WriteUID(uid uint32) (err error) {
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
func (device *ServoV2Bricklet) ReadUID() (uid uint32, err error) {
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
func (device *ServoV2Bricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
