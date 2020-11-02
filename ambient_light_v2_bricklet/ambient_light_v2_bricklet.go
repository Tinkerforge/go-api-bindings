/* ***********************************************************
 * This file was automatically generated on 2020-11-02.      *
 *                                                           *
 * Go Bindings Version 2.0.9                                 *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


// Measures ambient light up to 64000lux.
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/AmbientLightV2_Bricklet_Go.html.
package ambient_light_v2_bricklet

import (
	"encoding/binary"
	"bytes"
	"fmt"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionGetIlluminance Function = 1
	FunctionSetIlluminanceCallbackPeriod Function = 2
	FunctionGetIlluminanceCallbackPeriod Function = 3
	FunctionSetIlluminanceCallbackThreshold Function = 4
	FunctionGetIlluminanceCallbackThreshold Function = 5
	FunctionSetDebouncePeriod Function = 6
	FunctionGetDebouncePeriod Function = 7
	FunctionSetConfiguration Function = 8
	FunctionGetConfiguration Function = 9
	FunctionGetIdentity Function = 255
	FunctionCallbackIlluminance Function = 10
	FunctionCallbackIlluminanceReached Function = 11
)

type ThresholdOption = rune

const (
	ThresholdOptionOff ThresholdOption = 'x'
	ThresholdOptionOutside ThresholdOption = 'o'
	ThresholdOptionInside ThresholdOption = 'i'
	ThresholdOptionSmaller ThresholdOption = '<'
	ThresholdOptionGreater ThresholdOption = '>'
)

type IlluminanceRange = uint8

const (
	IlluminanceRangeUnlimited IlluminanceRange = 6
	IlluminanceRange64000Lux IlluminanceRange = 0
	IlluminanceRange32000Lux IlluminanceRange = 1
	IlluminanceRange16000Lux IlluminanceRange = 2
	IlluminanceRange8000Lux IlluminanceRange = 3
	IlluminanceRange1300Lux IlluminanceRange = 4
	IlluminanceRange600Lux IlluminanceRange = 5
)

type IntegrationTime = uint8

const (
	IntegrationTime50ms IntegrationTime = 0
	IntegrationTime100ms IntegrationTime = 1
	IntegrationTime150ms IntegrationTime = 2
	IntegrationTime200ms IntegrationTime = 3
	IntegrationTime250ms IntegrationTime = 4
	IntegrationTime300ms IntegrationTime = 5
	IntegrationTime350ms IntegrationTime = 6
	IntegrationTime400ms IntegrationTime = 7
)

type AmbientLightV2Bricklet struct {
	device Device
}
const DeviceIdentifier = 259
const DeviceDisplayName = "Ambient Light Bricklet 2.0"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (AmbientLightV2Bricklet, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{ 2,0,1 }, uid, &internalIPCon, 0, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return AmbientLightV2Bricklet{}, err
	}
	dev.ResponseExpected[FunctionGetIlluminance] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetIlluminanceCallbackPeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetIlluminanceCallbackPeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetIlluminanceCallbackThreshold] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetIlluminanceCallbackThreshold] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetDebouncePeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetDebouncePeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetConfiguration] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetIdentity] = ResponseExpectedFlagAlwaysTrue;
	return AmbientLightV2Bricklet{dev}, nil
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
func (device *AmbientLightV2Bricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *AmbientLightV2Bricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *AmbientLightV2Bricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *AmbientLightV2Bricklet) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// This callback is triggered periodically with the period that is set by
// SetIlluminanceCallbackPeriod. The parameter is the illuminance of the
// ambient light sensor.
// 
// The RegisterIlluminanceCallback callback is only triggered if the illuminance has changed since the
// last triggering.
func (device *AmbientLightV2Bricklet) RegisterIlluminanceCallback(fn func(uint32)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 12 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var illuminance uint32
		binary.Read(buf, binary.LittleEndian, &illuminance)
		fn(illuminance)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackIlluminance), wrapper)
}

// Remove a registered Illuminance callback.
func (device *AmbientLightV2Bricklet) DeregisterIlluminanceCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackIlluminance), registrationId)
}


// This callback is triggered when the threshold as set by
// SetIlluminanceCallbackThreshold is reached.
// The parameter is the illuminance of the ambient light sensor.
// 
// If the threshold keeps being reached, the callback is triggered periodically
// with the period as set by SetDebouncePeriod.
func (device *AmbientLightV2Bricklet) RegisterIlluminanceReachedCallback(fn func(uint32)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 12 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var illuminance uint32
		binary.Read(buf, binary.LittleEndian, &illuminance)
		fn(illuminance)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackIlluminanceReached), wrapper)
}

// Remove a registered Illuminance Reached callback.
func (device *AmbientLightV2Bricklet) DeregisterIlluminanceReachedCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackIlluminanceReached), registrationId)
}


// Returns the illuminance of the ambient light sensor. The measurement range goes
// up to about 100000lux, but above 64000lux the precision starts to drop.
// 
// .. versionchanged:: 2.0.2$nbsp;(Plugin)
//   An illuminance of 0lux indicates that the sensor is saturated and the
//   configuration should be modified, see SetConfiguration.
// 
// If you want to get the illuminance periodically, it is recommended to use the
// RegisterIlluminanceCallback callback and set the period with
// SetIlluminanceCallbackPeriod.
func (device *AmbientLightV2Bricklet) GetIlluminance() (illuminance uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetIlluminance), buf.Bytes())
	if err != nil {
		return illuminance, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 12 {
			return illuminance, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		if header.ErrorCode != 0 {
			return illuminance, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &illuminance)

	}

	return illuminance, nil
}

// Sets the period with which the RegisterIlluminanceCallback callback is triggered
// periodically. A value of 0 turns the callback off.
// 
// The RegisterIlluminanceCallback callback is only triggered if the illuminance has changed
// since the last triggering.
func (device *AmbientLightV2Bricklet) SetIlluminanceCallbackPeriod(period uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);

	resultBytes, err := device.device.Set(uint8(FunctionSetIlluminanceCallbackPeriod), buf.Bytes())
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

// Returns the period as set by SetIlluminanceCallbackPeriod.
func (device *AmbientLightV2Bricklet) GetIlluminanceCallbackPeriod() (period uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetIlluminanceCallbackPeriod), buf.Bytes())
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

// Sets the thresholds for the RegisterIlluminanceReachedCallback callback.
// 
// The following options are possible:
// 
//  Option| Description
//  --- | --- 
//  'x'|    Callback is turned off
//  'o'|    Callback is triggered when the illuminance is *outside* the min and max values
//  'i'|    Callback is triggered when the illuminance is *inside* the min and max values
//  '<'|    Callback is triggered when the illuminance is smaller than the min value (max is ignored)
//  '>'|    Callback is triggered when the illuminance is greater than the min value (max is ignored)
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *AmbientLightV2Bricklet) SetIlluminanceCallbackThreshold(option ThresholdOption, min uint32, max uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, option);
	binary.Write(&buf, binary.LittleEndian, min);
	binary.Write(&buf, binary.LittleEndian, max);

	resultBytes, err := device.device.Set(uint8(FunctionSetIlluminanceCallbackThreshold), buf.Bytes())
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

// Returns the threshold as set by SetIlluminanceCallbackThreshold.
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *AmbientLightV2Bricklet) GetIlluminanceCallbackThreshold() (option ThresholdOption, min uint32, max uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetIlluminanceCallbackThreshold), buf.Bytes())
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
// * RegisterIlluminanceReachedCallback,
// 
// are triggered, if the thresholds
// 
// * SetIlluminanceCallbackThreshold,
// 
// keep being reached.
func (device *AmbientLightV2Bricklet) SetDebouncePeriod(debounce uint32) (err error) {
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
func (device *AmbientLightV2Bricklet) GetDebouncePeriod() (debounce uint32, err error) {
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

// Sets the configuration. It is possible to configure an illuminance range
// between 0-600lux and 0-64000lux and an integration time between 50ms and 400ms.
// 
// .. versionadded:: 2.0.2$nbsp;(Plugin)
//   The unlimited illuminance range allows to measure up to about 100000lux, but
//   above 64000lux the precision starts to drop.
// 
// A smaller illuminance range increases the resolution of the data. A longer
// integration time will result in less noise on the data.
// 
// .. versionchanged:: 2.0.2$nbsp;(Plugin)
//   If the actual measure illuminance is out-of-range then the current illuminance
//   range maximum +0.01lux is reported by GetIlluminance and the
//   RegisterIlluminanceCallback callback. For example, 800001 for the 0-8000lux range.
// 
// .. versionchanged:: 2.0.2$nbsp;(Plugin)
//   With a long integration time the sensor might be saturated before the measured
//   value reaches the maximum of the selected illuminance range. In this case 0lux
//   is reported by GetIlluminance and the RegisterIlluminanceCallback callback.
// 
// If the measurement is out-of-range or the sensor is saturated then you should
// configure the next higher illuminance range. If the highest range is already
// in use, then start to reduce the integration time.
//
// Associated constants:
//
//	* IlluminanceRangeUnlimited
//	* IlluminanceRange64000Lux
//	* IlluminanceRange32000Lux
//	* IlluminanceRange16000Lux
//	* IlluminanceRange8000Lux
//	* IlluminanceRange1300Lux
//	* IlluminanceRange600Lux
//	* IntegrationTime50ms
//	* IntegrationTime100ms
//	* IntegrationTime150ms
//	* IntegrationTime200ms
//	* IntegrationTime250ms
//	* IntegrationTime300ms
//	* IntegrationTime350ms
//	* IntegrationTime400ms
func (device *AmbientLightV2Bricklet) SetConfiguration(illuminanceRange IlluminanceRange, integrationTime IntegrationTime) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, illuminanceRange);
	binary.Write(&buf, binary.LittleEndian, integrationTime);

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
//	* IlluminanceRangeUnlimited
//	* IlluminanceRange64000Lux
//	* IlluminanceRange32000Lux
//	* IlluminanceRange16000Lux
//	* IlluminanceRange8000Lux
//	* IlluminanceRange1300Lux
//	* IlluminanceRange600Lux
//	* IntegrationTime50ms
//	* IntegrationTime100ms
//	* IntegrationTime150ms
//	* IntegrationTime200ms
//	* IntegrationTime250ms
//	* IntegrationTime300ms
//	* IntegrationTime350ms
//	* IntegrationTime400ms
func (device *AmbientLightV2Bricklet) GetConfiguration() (illuminanceRange IlluminanceRange, integrationTime IntegrationTime, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetConfiguration), buf.Bytes())
	if err != nil {
		return illuminanceRange, integrationTime, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 10 {
			return illuminanceRange, integrationTime, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		if header.ErrorCode != 0 {
			return illuminanceRange, integrationTime, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &illuminanceRange)
		binary.Read(resultBuf, binary.LittleEndian, &integrationTime)

	}

	return illuminanceRange, integrationTime, nil
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
func (device *AmbientLightV2Bricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
