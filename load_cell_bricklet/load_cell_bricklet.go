/* ***********************************************************
 * This file was automatically generated on 2020-05-19.      *
 *                                                           *
 * Go Bindings Version 2.0.8                                 *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


// Measures weight with a load cell.
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/LoadCell_Bricklet_Go.html.
package load_cell_bricklet

import (
	"encoding/binary"
	"bytes"
	"fmt"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionGetWeight Function = 1
	FunctionSetWeightCallbackPeriod Function = 2
	FunctionGetWeightCallbackPeriod Function = 3
	FunctionSetWeightCallbackThreshold Function = 4
	FunctionGetWeightCallbackThreshold Function = 5
	FunctionSetDebouncePeriod Function = 6
	FunctionGetDebouncePeriod Function = 7
	FunctionSetMovingAverage Function = 8
	FunctionGetMovingAverage Function = 9
	FunctionLEDOn Function = 10
	FunctionLEDOff Function = 11
	FunctionIsLEDOn Function = 12
	FunctionCalibrate Function = 13
	FunctionTare Function = 14
	FunctionSetConfiguration Function = 15
	FunctionGetConfiguration Function = 16
	FunctionGetIdentity Function = 255
	FunctionCallbackWeight Function = 17
	FunctionCallbackWeightReached Function = 18
)

type ThresholdOption = rune

const (
	ThresholdOptionOff ThresholdOption = 'x'
	ThresholdOptionOutside ThresholdOption = 'o'
	ThresholdOptionInside ThresholdOption = 'i'
	ThresholdOptionSmaller ThresholdOption = '<'
	ThresholdOptionGreater ThresholdOption = '>'
)

type Rate = uint8

const (
	Rate10Hz Rate = 0
	Rate80Hz Rate = 1
)

type Gain = uint8

const (
	Gain128x Gain = 0
	Gain64x Gain = 1
	Gain32x Gain = 2
)

type LoadCellBricklet struct {
	device Device
}
const DeviceIdentifier = 253
const DeviceDisplayName = "Load Cell Bricklet"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (LoadCellBricklet, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{ 2,0,0 }, uid, &internalIPCon, 0, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return LoadCellBricklet{}, err
	}
	dev.ResponseExpected[FunctionGetWeight] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetWeightCallbackPeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetWeightCallbackPeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetWeightCallbackThreshold] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetWeightCallbackThreshold] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetDebouncePeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetDebouncePeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetMovingAverage] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetMovingAverage] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionLEDOn] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionLEDOff] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionIsLEDOn] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionCalibrate] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionTare] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionSetConfiguration] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetIdentity] = ResponseExpectedFlagAlwaysTrue;
	return LoadCellBricklet{dev}, nil
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
func (device *LoadCellBricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *LoadCellBricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *LoadCellBricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *LoadCellBricklet) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// This callback is triggered periodically with the period that is set by
// SetWeightCallbackPeriod. The parameter is the weight
// as measured by the load cell.
// 
// The RegisterWeightCallback callback is only triggered if the weight has changed since the
// last triggering.
func (device *LoadCellBricklet) RegisterWeightCallback(fn func(int32)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 12 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var weight int32
		binary.Read(buf, binary.LittleEndian, &weight)
		fn(weight)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackWeight), wrapper)
}

// Remove a registered Weight callback.
func (device *LoadCellBricklet) DeregisterWeightCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackWeight), registrationId)
}


// This callback is triggered when the threshold as set by
// SetWeightCallbackThreshold is reached.
// The parameter is the weight as measured by the load cell.
// 
// If the threshold keeps being reached, the callback is triggered periodically
// with the period as set by SetDebouncePeriod.
func (device *LoadCellBricklet) RegisterWeightReachedCallback(fn func(int32)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 12 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var weight int32
		binary.Read(buf, binary.LittleEndian, &weight)
		fn(weight)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackWeightReached), wrapper)
}

// Remove a registered Weight Reached callback.
func (device *LoadCellBricklet) DeregisterWeightReachedCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackWeightReached), registrationId)
}


// Returns the currently measured weight.
// 
// If you want to get the weight periodically, it is recommended
// to use the RegisterWeightCallback callback and set the period with
// SetWeightCallbackPeriod.
func (device *LoadCellBricklet) GetWeight() (weight int32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetWeight), buf.Bytes())
	if err != nil {
		return weight, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 12 {
			return weight, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		if header.ErrorCode != 0 {
			return weight, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &weight)

	}

	return weight, nil
}

// Sets the period with which the RegisterWeightCallback callback is triggered
// periodically. A value of 0 turns the callback off.
// 
// The RegisterWeightCallback callback is only triggered if the weight has changed since the
// last triggering.
func (device *LoadCellBricklet) SetWeightCallbackPeriod(period uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);

	resultBytes, err := device.device.Set(uint8(FunctionSetWeightCallbackPeriod), buf.Bytes())
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

// Returns the period as set by SetWeightCallbackPeriod.
func (device *LoadCellBricklet) GetWeightCallbackPeriod() (period uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetWeightCallbackPeriod), buf.Bytes())
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

// Sets the thresholds for the RegisterWeightReachedCallback callback.
// 
// The following options are possible:
// 
//  Option| Description
//  --- | --- 
//  'x'|    Callback is turned off
//  'o'|    Callback is triggered when the weight is *outside* the min and max values
//  'i'|    Callback is triggered when the weight is *inside* the min and max values
//  '<'|    Callback is triggered when the weight is smaller than the min value (max is ignored)
//  '>'|    Callback is triggered when the weight is greater than the min value (max is ignored)
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *LoadCellBricklet) SetWeightCallbackThreshold(option ThresholdOption, min int32, max int32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, option);
	binary.Write(&buf, binary.LittleEndian, min);
	binary.Write(&buf, binary.LittleEndian, max);

	resultBytes, err := device.device.Set(uint8(FunctionSetWeightCallbackThreshold), buf.Bytes())
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

// Returns the threshold as set by SetWeightCallbackThreshold.
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *LoadCellBricklet) GetWeightCallbackThreshold() (option ThresholdOption, min int32, max int32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetWeightCallbackThreshold), buf.Bytes())
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

// Sets the period with which the threshold callback
// 
// * RegisterWeightReachedCallback
// 
// is triggered, if the threshold
// 
// * SetWeightCallbackThreshold
// 
// keeps being reached.
func (device *LoadCellBricklet) SetDebouncePeriod(debounce uint32) (err error) {
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
func (device *LoadCellBricklet) GetDebouncePeriod() (debounce uint32, err error) {
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

// Sets the length of a https://en.wikipedia.org/wiki/Moving_average
// for the weight value.
// 
// Setting the length to 1 will turn the averaging off. With less
// averaging, there is more noise on the data.
func (device *LoadCellBricklet) SetMovingAverage(average uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, average);

	resultBytes, err := device.device.Set(uint8(FunctionSetMovingAverage), buf.Bytes())
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

// Returns the length moving average as set by SetMovingAverage.
func (device *LoadCellBricklet) GetMovingAverage() (average uint8, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetMovingAverage), buf.Bytes())
	if err != nil {
		return average, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return average, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return average, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &average)

	}

	return average, nil
}

// Turns the LED on.
func (device *LoadCellBricklet) LEDOn() (err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Set(uint8(FunctionLEDOn), buf.Bytes())
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

// Turns the LED off.
func (device *LoadCellBricklet) LEDOff() (err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Set(uint8(FunctionLEDOff), buf.Bytes())
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

// Returns *true* if the led is on, *false* otherwise.
func (device *LoadCellBricklet) IsLEDOn() (on bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionIsLEDOn), buf.Bytes())
	if err != nil {
		return on, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return on, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return on, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &on)

	}

	return on, nil
}

// To calibrate your Load Cell Bricklet you have to
// 
// * empty the scale and call this function with 0 and
// * add a known weight to the scale and call this function with the weight.
// 
// The calibration is saved in the EEPROM of the Bricklet and only
// needs to be done once.
// 
// We recommend to use the Brick Viewer for calibration, you don't need
// to call this function in your source code.
func (device *LoadCellBricklet) Calibrate(weight uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, weight);

	resultBytes, err := device.device.Set(uint8(FunctionCalibrate), buf.Bytes())
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

// Sets the currently measured weight as tare weight.
func (device *LoadCellBricklet) Tare() (err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Set(uint8(FunctionTare), buf.Bytes())
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

// The measurement rate and gain are configurable.
// 
// The rate can be either 10Hz or 80Hz. A faster rate will produce more noise.
// It is additionally possible to add a moving average
// (see SetMovingAverage) to the measurements.
// 
// The gain can be 128x, 64x or 32x. It represents a measurement range of
// ±20mV, ±40mV and ±80mV respectively. The Load Cell Bricklet uses an
// excitation voltage of 5V and most load cells use an output of 2mV/V. That
// means the voltage range is ±15mV for most load cells (i.e. gain of 128x
// is best). If you don't know what all of this means you should keep it at
// 128x, it will most likely be correct.
// 
// The configuration is saved in the EEPROM of the Bricklet and only
// needs to be done once.
// 
// We recommend to use the Brick Viewer for configuration, you don't need
// to call this function in your source code.
//
// Associated constants:
//
//	* Rate10Hz
//	* Rate80Hz
//	* Gain128x
//	* Gain64x
//	* Gain32x
func (device *LoadCellBricklet) SetConfiguration(rate Rate, gain Gain) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, rate);
	binary.Write(&buf, binary.LittleEndian, gain);

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
//	* Rate10Hz
//	* Rate80Hz
//	* Gain128x
//	* Gain64x
//	* Gain32x
func (device *LoadCellBricklet) GetConfiguration() (rate Rate, gain Gain, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetConfiguration), buf.Bytes())
	if err != nil {
		return rate, gain, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 10 {
			return rate, gain, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		if header.ErrorCode != 0 {
			return rate, gain, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &rate)
		binary.Read(resultBuf, binary.LittleEndian, &gain)

	}

	return rate, gain, nil
}

// Returns the UID, the UID where the Bricklet is connected to,
// the position, the hardware and firmware version as well as the
// device identifier.
// 
// The position can be 'a', 'b', 'c', 'd', 'e', 'f', 'g' or 'h' (Bricklet Port).
// The Raspberry Pi HAT (Zero) Brick is always at position 'i' and the Bricklet
// connected to an `Isolator Bricklet <isolator_bricklet>` is always as
// position 'z'.
// 
// The device identifier numbers can be found `here <device_identifier>`.
// |device_identifier_constant|
func (device *LoadCellBricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
