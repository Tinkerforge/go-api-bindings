/* ***********************************************************
 * This file was automatically generated on 2021-01-15.      *
 *                                                           *
 * Go Bindings Version 2.0.10                                *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


// Measures power, DC voltage and DC current up to 720W/36V/20A‍.
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/VoltageCurrent_Bricklet_Go.html.
package voltage_current_bricklet

import (
	"encoding/binary"
	"bytes"
	"fmt"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionGetCurrent Function = 1
	FunctionGetVoltage Function = 2
	FunctionGetPower Function = 3
	FunctionSetConfiguration Function = 4
	FunctionGetConfiguration Function = 5
	FunctionSetCalibration Function = 6
	FunctionGetCalibration Function = 7
	FunctionSetCurrentCallbackPeriod Function = 8
	FunctionGetCurrentCallbackPeriod Function = 9
	FunctionSetVoltageCallbackPeriod Function = 10
	FunctionGetVoltageCallbackPeriod Function = 11
	FunctionSetPowerCallbackPeriod Function = 12
	FunctionGetPowerCallbackPeriod Function = 13
	FunctionSetCurrentCallbackThreshold Function = 14
	FunctionGetCurrentCallbackThreshold Function = 15
	FunctionSetVoltageCallbackThreshold Function = 16
	FunctionGetVoltageCallbackThreshold Function = 17
	FunctionSetPowerCallbackThreshold Function = 18
	FunctionGetPowerCallbackThreshold Function = 19
	FunctionSetDebouncePeriod Function = 20
	FunctionGetDebouncePeriod Function = 21
	FunctionGetIdentity Function = 255
	FunctionCallbackCurrent Function = 22
	FunctionCallbackVoltage Function = 23
	FunctionCallbackPower Function = 24
	FunctionCallbackCurrentReached Function = 25
	FunctionCallbackVoltageReached Function = 26
	FunctionCallbackPowerReached Function = 27
)

type Averaging = uint8

const (
	Averaging1 Averaging = 0
	Averaging4 Averaging = 1
	Averaging16 Averaging = 2
	Averaging64 Averaging = 3
	Averaging128 Averaging = 4
	Averaging256 Averaging = 5
	Averaging512 Averaging = 6
	Averaging1024 Averaging = 7
)

type ThresholdOption = rune

const (
	ThresholdOptionOff ThresholdOption = 'x'
	ThresholdOptionOutside ThresholdOption = 'o'
	ThresholdOptionInside ThresholdOption = 'i'
	ThresholdOptionSmaller ThresholdOption = '<'
	ThresholdOptionGreater ThresholdOption = '>'
)

type ConversionTime = uint8

const (
	ConversionTime140us ConversionTime = 0
	ConversionTime204us ConversionTime = 1
	ConversionTime332us ConversionTime = 2
	ConversionTime588us ConversionTime = 3
	ConversionTime1_1ms ConversionTime = 4
	ConversionTime2_116ms ConversionTime = 5
	ConversionTime4_156ms ConversionTime = 6
	ConversionTime8_244ms ConversionTime = 7
)

type VoltageCurrentBricklet struct {
	device Device
}
const DeviceIdentifier = 227
const DeviceDisplayName = "Voltage/Current Bricklet"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (VoltageCurrentBricklet, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{ 2,0,0 }, uid, &internalIPCon, 0, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return VoltageCurrentBricklet{}, err
	}
	dev.ResponseExpected[FunctionGetCurrent] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetVoltage] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetPower] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetConfiguration] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetCalibration] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetCalibration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetCurrentCallbackPeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetCurrentCallbackPeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetVoltageCallbackPeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetVoltageCallbackPeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetPowerCallbackPeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetPowerCallbackPeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetCurrentCallbackThreshold] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetCurrentCallbackThreshold] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetVoltageCallbackThreshold] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetVoltageCallbackThreshold] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetPowerCallbackThreshold] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetPowerCallbackThreshold] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetDebouncePeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetDebouncePeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetIdentity] = ResponseExpectedFlagAlwaysTrue;
	return VoltageCurrentBricklet{dev}, nil
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
func (device *VoltageCurrentBricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *VoltageCurrentBricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *VoltageCurrentBricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *VoltageCurrentBricklet) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// This callback is triggered periodically with the period that is set by
// SetCurrentCallbackPeriod. The parameter is the current of the
// sensor.
// 
// The RegisterCurrentCallback callback is only triggered if the current has changed since
// the last triggering.
func (device *VoltageCurrentBricklet) RegisterCurrentCallback(fn func(int32)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 12 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var current int32
		binary.Read(buf, binary.LittleEndian, &current)
		fn(current)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackCurrent), wrapper)
}

// Remove a registered Current callback.
func (device *VoltageCurrentBricklet) DeregisterCurrentCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackCurrent), registrationId)
}


// This callback is triggered periodically with the period that is set by
// SetVoltageCallbackPeriod. The parameter is the voltage of
// the sensor.
// 
// The RegisterVoltageCallback callback is only triggered if the voltage has changed since
// the last triggering.
func (device *VoltageCurrentBricklet) RegisterVoltageCallback(fn func(int32)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 12 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var voltage int32
		binary.Read(buf, binary.LittleEndian, &voltage)
		fn(voltage)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackVoltage), wrapper)
}

// Remove a registered Voltage callback.
func (device *VoltageCurrentBricklet) DeregisterVoltageCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackVoltage), registrationId)
}


// This callback is triggered periodically with the period that is set by
// SetPowerCallbackPeriod. The parameter is the power of the
// sensor.
// 
// The RegisterPowerCallback callback is only triggered if the power has changed since the
// last triggering.
func (device *VoltageCurrentBricklet) RegisterPowerCallback(fn func(int32)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 12 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var power int32
		binary.Read(buf, binary.LittleEndian, &power)
		fn(power)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackPower), wrapper)
}

// Remove a registered Power callback.
func (device *VoltageCurrentBricklet) DeregisterPowerCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackPower), registrationId)
}


// This callback is triggered when the threshold as set by
// SetCurrentCallbackThreshold is reached.
// The parameter is the current of the sensor.
// 
// If the threshold keeps being reached, the callback is triggered periodically
// with the period as set by SetDebouncePeriod.
func (device *VoltageCurrentBricklet) RegisterCurrentReachedCallback(fn func(int32)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 12 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var current int32
		binary.Read(buf, binary.LittleEndian, &current)
		fn(current)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackCurrentReached), wrapper)
}

// Remove a registered Current Reached callback.
func (device *VoltageCurrentBricklet) DeregisterCurrentReachedCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackCurrentReached), registrationId)
}


// This callback is triggered when the threshold as set by
// SetVoltageCallbackThreshold is reached.
// The parameter is the voltage of the sensor.
// 
// If the threshold keeps being reached, the callback is triggered periodically
// with the period as set by SetDebouncePeriod.
func (device *VoltageCurrentBricklet) RegisterVoltageReachedCallback(fn func(int32)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 12 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var voltage int32
		binary.Read(buf, binary.LittleEndian, &voltage)
		fn(voltage)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackVoltageReached), wrapper)
}

// Remove a registered Voltage Reached callback.
func (device *VoltageCurrentBricklet) DeregisterVoltageReachedCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackVoltageReached), registrationId)
}


// This callback is triggered when the threshold as set by
// SetPowerCallbackThreshold is reached.
// The parameter is the power of the sensor.
// 
// If the threshold keeps being reached, the callback is triggered periodically
// with the period as set by SetDebouncePeriod.
func (device *VoltageCurrentBricklet) RegisterPowerReachedCallback(fn func(int32)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 12 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var power int32
		binary.Read(buf, binary.LittleEndian, &power)
		fn(power)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackPowerReached), wrapper)
}

// Remove a registered Power Reached callback.
func (device *VoltageCurrentBricklet) DeregisterPowerReachedCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackPowerReached), registrationId)
}


// Returns the current.
// 
// If you want to get the current periodically, it is recommended to use the
// RegisterCurrentCallback callback and set the period with
// SetCurrentCallbackPeriod.
func (device *VoltageCurrentBricklet) GetCurrent() (current int32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetCurrent), buf.Bytes())
	if err != nil {
		return current, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 12 {
			return current, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		if header.ErrorCode != 0 {
			return current, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &current)

	}

	return current, nil
}

// Returns the voltage.
// 
// If you want to get the voltage periodically, it is recommended to use the
// RegisterVoltageCallback callback and set the period with
// SetVoltageCallbackPeriod.
func (device *VoltageCurrentBricklet) GetVoltage() (voltage int32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetVoltage), buf.Bytes())
	if err != nil {
		return voltage, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 12 {
			return voltage, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		if header.ErrorCode != 0 {
			return voltage, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &voltage)

	}

	return voltage, nil
}

// Returns the power.
// 
// If you want to get the power periodically, it is recommended to use the
// RegisterPowerCallback callback and set the period with
// SetPowerCallbackPeriod.
func (device *VoltageCurrentBricklet) GetPower() (power int32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetPower), buf.Bytes())
	if err != nil {
		return power, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 12 {
			return power, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		if header.ErrorCode != 0 {
			return power, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &power)

	}

	return power, nil
}

// Sets the configuration of the Voltage/Current Bricklet. It is
// possible to configure number of averages as well as
// voltage and current conversion time.
//
// Associated constants:
//
//	* Averaging1
//	* Averaging4
//	* Averaging16
//	* Averaging64
//	* Averaging128
//	* Averaging256
//	* Averaging512
//	* Averaging1024
//	* ConversionTime140us
//	* ConversionTime204us
//	* ConversionTime332us
//	* ConversionTime588us
//	* ConversionTime11ms
//	* ConversionTime2116ms
//	* ConversionTime4156ms
//	* ConversionTime8244ms
func (device *VoltageCurrentBricklet) SetConfiguration(averaging Averaging, voltageConversionTime ConversionTime, currentConversionTime ConversionTime) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, averaging);
	binary.Write(&buf, binary.LittleEndian, voltageConversionTime);
	binary.Write(&buf, binary.LittleEndian, currentConversionTime);

	resultBytes, err := device.device.Set(uint8(FunctionSetConfiguration), buf.Bytes())
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

// Returns the configuration as set by SetConfiguration.
//
// Associated constants:
//
//	* Averaging1
//	* Averaging4
//	* Averaging16
//	* Averaging64
//	* Averaging128
//	* Averaging256
//	* Averaging512
//	* Averaging1024
//	* ConversionTime140us
//	* ConversionTime204us
//	* ConversionTime332us
//	* ConversionTime588us
//	* ConversionTime11ms
//	* ConversionTime2116ms
//	* ConversionTime4156ms
//	* ConversionTime8244ms
func (device *VoltageCurrentBricklet) GetConfiguration() (averaging Averaging, voltageConversionTime ConversionTime, currentConversionTime ConversionTime, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetConfiguration), buf.Bytes())
	if err != nil {
		return averaging, voltageConversionTime, currentConversionTime, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 11 {
			return averaging, voltageConversionTime, currentConversionTime, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 11)
		}

		if header.ErrorCode != 0 {
			return averaging, voltageConversionTime, currentConversionTime, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &averaging)
		binary.Read(resultBuf, binary.LittleEndian, &voltageConversionTime)
		binary.Read(resultBuf, binary.LittleEndian, &currentConversionTime)

	}

	return averaging, voltageConversionTime, currentConversionTime, nil
}

// Since the shunt resistor that is used to measure the current is not
// perfectly precise, it needs to be calibrated by a multiplier and
// divisor if a very precise reading is needed.
// 
// For example, if you are expecting a measurement of 1000mA and you
// are measuring 1023mA, you can calibrate the Voltage/Current Bricklet
// by setting the multiplier to 1000 and the divisor to 1023.
func (device *VoltageCurrentBricklet) SetCalibration(gainMultiplier uint16, gainDivisor uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, gainMultiplier);
	binary.Write(&buf, binary.LittleEndian, gainDivisor);

	resultBytes, err := device.device.Set(uint8(FunctionSetCalibration), buf.Bytes())
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

// Returns the calibration as set by SetCalibration.
func (device *VoltageCurrentBricklet) GetCalibration() (gainMultiplier uint16, gainDivisor uint16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetCalibration), buf.Bytes())
	if err != nil {
		return gainMultiplier, gainDivisor, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 12 {
			return gainMultiplier, gainDivisor, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		if header.ErrorCode != 0 {
			return gainMultiplier, gainDivisor, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &gainMultiplier)
		binary.Read(resultBuf, binary.LittleEndian, &gainDivisor)

	}

	return gainMultiplier, gainDivisor, nil
}

// Sets the period with which the RegisterCurrentCallback callback is triggered
// periodically. A value of 0 turns the callback off.
// 
// The RegisterCurrentCallback callback is only triggered if the current has changed since
// the last triggering.
func (device *VoltageCurrentBricklet) SetCurrentCallbackPeriod(period uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);

	resultBytes, err := device.device.Set(uint8(FunctionSetCurrentCallbackPeriod), buf.Bytes())
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

// Returns the period as set by SetCurrentCallbackPeriod.
func (device *VoltageCurrentBricklet) GetCurrentCallbackPeriod() (period uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetCurrentCallbackPeriod), buf.Bytes())
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

// Sets the period with which the RegisterVoltageCallback callback is triggered
// periodically. A value of 0 turns the callback off.
// 
// The RegisterVoltageCallback callback is only triggered if the voltage has changed since
// the last triggering.
func (device *VoltageCurrentBricklet) SetVoltageCallbackPeriod(period uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);

	resultBytes, err := device.device.Set(uint8(FunctionSetVoltageCallbackPeriod), buf.Bytes())
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

// Returns the period as set by SetVoltageCallbackPeriod.
func (device *VoltageCurrentBricklet) GetVoltageCallbackPeriod() (period uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetVoltageCallbackPeriod), buf.Bytes())
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

// Sets the period with which the RegisterPowerCallback callback is triggered
// periodically. A value of 0 turns the callback off.
// 
// The RegisterPowerCallback callback is only triggered if the power has changed since the
// last triggering.
func (device *VoltageCurrentBricklet) SetPowerCallbackPeriod(period uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);

	resultBytes, err := device.device.Set(uint8(FunctionSetPowerCallbackPeriod), buf.Bytes())
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

// Returns the period as set by GetPowerCallbackPeriod.
func (device *VoltageCurrentBricklet) GetPowerCallbackPeriod() (period uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetPowerCallbackPeriod), buf.Bytes())
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

// Sets the thresholds for the RegisterCurrentReachedCallback callback.
// 
// The following options are possible:
// 
//  Option| Description
//  --- | --- 
//  'x'|    Callback is turned off
//  'o'|    Callback is triggered when the current is *outside* the min and max values
//  'i'|    Callback is triggered when the current is *inside* the min and max values
//  '<'|    Callback is triggered when the current is smaller than the min value (max is ignored)
//  '>'|    Callback is triggered when the current is greater than the min value (max is ignored)
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *VoltageCurrentBricklet) SetCurrentCallbackThreshold(option ThresholdOption, min int32, max int32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, option);
	binary.Write(&buf, binary.LittleEndian, min);
	binary.Write(&buf, binary.LittleEndian, max);

	resultBytes, err := device.device.Set(uint8(FunctionSetCurrentCallbackThreshold), buf.Bytes())
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

// Returns the threshold as set by SetCurrentCallbackThreshold.
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *VoltageCurrentBricklet) GetCurrentCallbackThreshold() (option ThresholdOption, min int32, max int32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetCurrentCallbackThreshold), buf.Bytes())
	if err != nil {
		return option, min, max, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 17 {
			return option, min, max, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 17)
		}

		if header.ErrorCode != 0 {
			return option, min, max, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &option)
		binary.Read(resultBuf, binary.LittleEndian, &min)
		binary.Read(resultBuf, binary.LittleEndian, &max)

	}

	return option, min, max, nil
}

// Sets the thresholds for the RegisterVoltageReachedCallback callback.
// 
// The following options are possible:
// 
//  Option| Description
//  --- | --- 
//  'x'|    Callback is turned off
//  'o'|    Callback is triggered when the voltage is *outside* the min and max values
//  'i'|    Callback is triggered when the voltage is *inside* the min and max values
//  '<'|    Callback is triggered when the voltage is smaller than the min value (max is ignored)
//  '>'|    Callback is triggered when the voltage is greater than the min value (max is ignored)
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *VoltageCurrentBricklet) SetVoltageCallbackThreshold(option ThresholdOption, min int32, max int32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, option);
	binary.Write(&buf, binary.LittleEndian, min);
	binary.Write(&buf, binary.LittleEndian, max);

	resultBytes, err := device.device.Set(uint8(FunctionSetVoltageCallbackThreshold), buf.Bytes())
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

// Returns the threshold as set by SetVoltageCallbackThreshold.
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *VoltageCurrentBricklet) GetVoltageCallbackThreshold() (option ThresholdOption, min int32, max int32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetVoltageCallbackThreshold), buf.Bytes())
	if err != nil {
		return option, min, max, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 17 {
			return option, min, max, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 17)
		}

		if header.ErrorCode != 0 {
			return option, min, max, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &option)
		binary.Read(resultBuf, binary.LittleEndian, &min)
		binary.Read(resultBuf, binary.LittleEndian, &max)

	}

	return option, min, max, nil
}

// Sets the thresholds for the RegisterPowerReachedCallback callback.
// 
// The following options are possible:
// 
//  Option| Description
//  --- | --- 
//  'x'|    Callback is turned off
//  'o'|    Callback is triggered when the power is *outside* the min and max values
//  'i'|    Callback is triggered when the power is *inside* the min and max values
//  '<'|    Callback is triggered when the power is smaller than the min value (max is ignored)
//  '>'|    Callback is triggered when the power is greater than the min value (max is ignored)
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *VoltageCurrentBricklet) SetPowerCallbackThreshold(option ThresholdOption, min int32, max int32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, option);
	binary.Write(&buf, binary.LittleEndian, min);
	binary.Write(&buf, binary.LittleEndian, max);

	resultBytes, err := device.device.Set(uint8(FunctionSetPowerCallbackThreshold), buf.Bytes())
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

// Returns the threshold as set by SetPowerCallbackThreshold.
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *VoltageCurrentBricklet) GetPowerCallbackThreshold() (option ThresholdOption, min int32, max int32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetPowerCallbackThreshold), buf.Bytes())
	if err != nil {
		return option, min, max, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 17 {
			return option, min, max, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 17)
		}

		if header.ErrorCode != 0 {
			return option, min, max, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &option)
		binary.Read(resultBuf, binary.LittleEndian, &min)
		binary.Read(resultBuf, binary.LittleEndian, &max)

	}

	return option, min, max, nil
}

// Sets the period with which the threshold callbacks
// 
// * RegisterCurrentReachedCallback,
// * RegisterVoltageReachedCallback,
// * RegisterPowerReachedCallback
// 
// are triggered, if the thresholds
// 
// * SetCurrentCallbackThreshold,
// * SetVoltageCallbackThreshold,
// * SetPowerCallbackThreshold
// 
// keep being reached.
func (device *VoltageCurrentBricklet) SetDebouncePeriod(debounce uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, debounce);

	resultBytes, err := device.device.Set(uint8(FunctionSetDebouncePeriod), buf.Bytes())
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

// Returns the debounce period as set by SetDebouncePeriod.
func (device *VoltageCurrentBricklet) GetDebouncePeriod() (debounce uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetDebouncePeriod), buf.Bytes())
	if err != nil {
		return debounce, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 12 {
			return debounce, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		if header.ErrorCode != 0 {
			return debounce, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &debounce)

	}

	return debounce, nil
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
func (device *VoltageCurrentBricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
