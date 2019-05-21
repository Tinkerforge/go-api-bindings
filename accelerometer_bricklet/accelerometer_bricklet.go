/* ***********************************************************
 * This file was automatically generated on 2019-05-21.      *
 *                                                           *
 * Go Bindings Version 2.0.3                                 *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


//Measures acceleration in three axis.
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/Accelerometer_Bricklet_Go.html.
package accelerometer_bricklet

import (
	"encoding/binary"
	"bytes"
    . "github.com/Tinkerforge/go-api-bindings/internal"
    "github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function uint8

const (
    FunctionGetAcceleration Function = 1
	FunctionSetAccelerationCallbackPeriod Function = 2
	FunctionGetAccelerationCallbackPeriod Function = 3
	FunctionSetAccelerationCallbackThreshold Function = 4
	FunctionGetAccelerationCallbackThreshold Function = 5
	FunctionSetDebouncePeriod Function = 6
	FunctionGetDebouncePeriod Function = 7
	FunctionGetTemperature Function = 8
	FunctionSetConfiguration Function = 9
	FunctionGetConfiguration Function = 10
	FunctionLEDOn Function = 11
	FunctionLEDOff Function = 12
	FunctionIsLEDOn Function = 13
	FunctionGetIdentity Function = 255
	FunctionCallbackAcceleration Function = 14
	FunctionCallbackAccelerationReached Function = 15
)

type ThresholdOption rune

const (
    ThresholdOptionOff ThresholdOption = 'x'
	ThresholdOptionOutside ThresholdOption = 'o'
	ThresholdOptionInside ThresholdOption = 'i'
	ThresholdOptionSmaller ThresholdOption = '<'
	ThresholdOptionGreater ThresholdOption = '>'
)

type DataRate uint8

const (
    DataRateOff DataRate = 0
	DataRate3Hz DataRate = 1
	DataRate6Hz DataRate = 2
	DataRate12Hz DataRate = 3
	DataRate25Hz DataRate = 4
	DataRate50Hz DataRate = 5
	DataRate100Hz DataRate = 6
	DataRate400Hz DataRate = 7
	DataRate800Hz DataRate = 8
	DataRate1600Hz DataRate = 9
)

type FullScale uint8

const (
    FullScale2g FullScale = 0
	FullScale4g FullScale = 1
	FullScale6g FullScale = 2
	FullScale8g FullScale = 3
	FullScale16g FullScale = 4
)

type FilterBandwidth uint8

const (
    FilterBandwidth800Hz FilterBandwidth = 0
	FilterBandwidth400Hz FilterBandwidth = 1
	FilterBandwidth200Hz FilterBandwidth = 2
	FilterBandwidth50Hz FilterBandwidth = 3
)

type AccelerometerBricklet struct{
	device Device
}
const DeviceIdentifier = 250
const DeviceDisplayName = "Accelerometer Bricklet"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (AccelerometerBricklet, error) {
    internalIPCon := ipcon.GetInternalHandle().(IPConnection)
    dev, err := NewDevice([3]uint8{ 2,0,1 }, uid, &internalIPCon, 0)
    if err != nil {
        return AccelerometerBricklet{}, err
    }
    dev.ResponseExpected[FunctionGetAcceleration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetAccelerationCallbackPeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetAccelerationCallbackPeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetAccelerationCallbackThreshold] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetAccelerationCallbackThreshold] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetDebouncePeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetDebouncePeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetTemperature] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetConfiguration] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionLEDOn] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionLEDOff] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionIsLEDOn] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetIdentity] = ResponseExpectedFlagAlwaysTrue;
    return AccelerometerBricklet{dev}, nil
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
func (device *AccelerometerBricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *AccelerometerBricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
    return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *AccelerometerBricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *AccelerometerBricklet) GetAPIVersion() [3]uint8 {
    return device.device.GetAPIVersion()
}

// This callback is triggered periodically with the period that is set by
// SetAccelerationCallbackPeriod. The parameters are the
// X, Y and Z acceleration.
// 
// The RegisterAccelerationCallback callback is only triggered if the acceleration has
// changed since the last triggering.
func (device *AccelerometerBricklet) RegisterAccelerationCallback(fn func(int16, int16, int16)) uint64 {
            wrapper := func(byteSlice []byte) {
                buf := bytes.NewBuffer(byteSlice[8:])
                var x int16
var y int16
var z int16
                binary.Read(buf, binary.LittleEndian, &x)
binary.Read(buf, binary.LittleEndian, &y)
binary.Read(buf, binary.LittleEndian, &z)
                fn(x, y, z)
            }
    return device.device.RegisterCallback(uint8(FunctionCallbackAcceleration), wrapper)
}

//Remove a registered Acceleration callback.
func (device *AccelerometerBricklet) DeregisterAccelerationCallback(registrationID uint64) {
    device.device.DeregisterCallback(uint8(FunctionCallbackAcceleration), registrationID)
}


// This callback is triggered when the threshold as set by
// SetAccelerationCallbackThreshold is reached.
// The parameters are the X, Y and Z acceleration.
// 
// If the threshold keeps being reached, the callback is triggered periodically
// with the period as set by SetDebouncePeriod.
func (device *AccelerometerBricklet) RegisterAccelerationReachedCallback(fn func(int16, int16, int16)) uint64 {
            wrapper := func(byteSlice []byte) {
                buf := bytes.NewBuffer(byteSlice[8:])
                var x int16
var y int16
var z int16
                binary.Read(buf, binary.LittleEndian, &x)
binary.Read(buf, binary.LittleEndian, &y)
binary.Read(buf, binary.LittleEndian, &z)
                fn(x, y, z)
            }
    return device.device.RegisterCallback(uint8(FunctionCallbackAccelerationReached), wrapper)
}

//Remove a registered Acceleration Reached callback.
func (device *AccelerometerBricklet) DeregisterAccelerationReachedCallback(registrationID uint64) {
    device.device.DeregisterCallback(uint8(FunctionCallbackAccelerationReached), registrationID)
}


// Returns the acceleration in x, y and z direction. The values
// are given in g/1000 (1g = 9.80665m/s²), not to be confused with grams.
// 
// If you want to get the acceleration periodically, it is recommended
// to use the RegisterAccelerationCallback callback and set the period with
// SetAccelerationCallbackPeriod.
func (device *AccelerometerBricklet) GetAcceleration() (x int16, y int16, z int16, err error) {
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetAcceleration), buf.Bytes())
    if err != nil {
        return x, y, z, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader

        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return x, y, z, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &x)
	binary.Read(resultBuf, binary.LittleEndian, &y)
	binary.Read(resultBuf, binary.LittleEndian, &z)

    }

    return x, y, z, nil
}

// Sets the period in ms with which the RegisterAccelerationCallback callback is triggered
// periodically. A value of 0 turns the callback off.
// 
// The RegisterAccelerationCallback callback is only triggered if the acceleration has
// changed since the last triggering.
// 
// The default value is 0.
func (device *AccelerometerBricklet) SetAccelerationCallbackPeriod(period uint32) (err error) {
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, period);

    resultBytes, err := device.device.Set(uint8(FunctionSetAccelerationCallbackPeriod), buf.Bytes())
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

// Returns the period as set by SetAccelerationCallbackPeriod.
func (device *AccelerometerBricklet) GetAccelerationCallbackPeriod() (period uint32, err error) {
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetAccelerationCallbackPeriod), buf.Bytes())
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

// Sets the thresholds for the RegisterAccelerationReachedCallback callback.
// 
// The following options are possible:
// 
//  Option| Description
//  --- | --- 
//  'x'|    Callback is turned off
//  'o'|    Callback is triggered when the acceleration is *outside* the min and max values
//  'i'|    Callback is triggered when the acceleration is *inside* the min and max values
//  '<'|    Callback is triggered when the acceleration is smaller than the min value (max is ignored)
//  '>'|    Callback is triggered when the acceleration is greater than the min value (max is ignored)
// 
// The default value is ('x', 0, 0, 0, 0, 0, 0).
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *AccelerometerBricklet) SetAccelerationCallbackThreshold(option ThresholdOption, minX int16, maxX int16, minY int16, maxY int16, minZ int16, maxZ int16) (err error) {
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, option);
	binary.Write(&buf, binary.LittleEndian, minX);
	binary.Write(&buf, binary.LittleEndian, maxX);
	binary.Write(&buf, binary.LittleEndian, minY);
	binary.Write(&buf, binary.LittleEndian, maxY);
	binary.Write(&buf, binary.LittleEndian, minZ);
	binary.Write(&buf, binary.LittleEndian, maxZ);

    resultBytes, err := device.device.Set(uint8(FunctionSetAccelerationCallbackThreshold), buf.Bytes())
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

// Returns the threshold as set by SetAccelerationCallbackThreshold.
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *AccelerometerBricklet) GetAccelerationCallbackThreshold() (option ThresholdOption, minX int16, maxX int16, minY int16, maxY int16, minZ int16, maxZ int16, err error) {
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetAccelerationCallbackThreshold), buf.Bytes())
    if err != nil {
        return option, minX, maxX, minY, maxY, minZ, maxZ, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader

        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return option, minX, maxX, minY, maxY, minZ, maxZ, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &option)
	binary.Read(resultBuf, binary.LittleEndian, &minX)
	binary.Read(resultBuf, binary.LittleEndian, &maxX)
	binary.Read(resultBuf, binary.LittleEndian, &minY)
	binary.Read(resultBuf, binary.LittleEndian, &maxY)
	binary.Read(resultBuf, binary.LittleEndian, &minZ)
	binary.Read(resultBuf, binary.LittleEndian, &maxZ)

    }

    return option, minX, maxX, minY, maxY, minZ, maxZ, nil
}

// Sets the period in ms with which the threshold callback
// 
// * RegisterAccelerationReachedCallback
// 
// is triggered, if the threshold
// 
// * SetAccelerationCallbackThreshold
// 
// keeps being reached.
// 
// The default value is 100.
func (device *AccelerometerBricklet) SetDebouncePeriod(debounce uint32) (err error) {
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, debounce);

    resultBytes, err := device.device.Set(uint8(FunctionSetDebouncePeriod), buf.Bytes())
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

// Returns the debounce period as set by SetDebouncePeriod.
func (device *AccelerometerBricklet) GetDebouncePeriod() (debounce uint32, err error) {
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetDebouncePeriod), buf.Bytes())
    if err != nil {
        return debounce, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader

        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return debounce, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &debounce)

    }

    return debounce, nil
}

// Returns the temperature of the accelerometer in °C.
func (device *AccelerometerBricklet) GetTemperature() (temperature int16, err error) {
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetTemperature), buf.Bytes())
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

// Configures the data rate, full scale range and filter bandwidth.
// Possible values are:
// 
// * Data rate of 0Hz to 1600Hz.
// * Full scale range of -2G to +2G up to -16G to +16G.
// * Filter bandwidth between 50Hz and 800Hz.
// 
// Decreasing data rate or full scale range will also decrease the noise on
// the data.
// 
// The default values are 100Hz data rate, -4G to +4G range and 200Hz
// filter bandwidth.
//
// Associated constants:
//
//	* DataRateOff
//	* DataRate3Hz
//	* DataRate6Hz
//	* DataRate12Hz
//	* DataRate25Hz
//	* DataRate50Hz
//	* DataRate100Hz
//	* DataRate400Hz
//	* DataRate800Hz
//	* DataRate1600Hz
//	* FullScale2g
//	* FullScale4g
//	* FullScale6g
//	* FullScale8g
//	* FullScale16g
//	* FilterBandwidth800Hz
//	* FilterBandwidth400Hz
//	* FilterBandwidth200Hz
//	* FilterBandwidth50Hz
func (device *AccelerometerBricklet) SetConfiguration(dataRate DataRate, fullScale FullScale, filterBandwidth FilterBandwidth) (err error) {
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, dataRate);
	binary.Write(&buf, binary.LittleEndian, fullScale);
	binary.Write(&buf, binary.LittleEndian, filterBandwidth);

    resultBytes, err := device.device.Set(uint8(FunctionSetConfiguration), buf.Bytes())
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

// Returns the configuration as set by SetConfiguration.
//
// Associated constants:
//
//	* DataRateOff
//	* DataRate3Hz
//	* DataRate6Hz
//	* DataRate12Hz
//	* DataRate25Hz
//	* DataRate50Hz
//	* DataRate100Hz
//	* DataRate400Hz
//	* DataRate800Hz
//	* DataRate1600Hz
//	* FullScale2g
//	* FullScale4g
//	* FullScale6g
//	* FullScale8g
//	* FullScale16g
//	* FilterBandwidth800Hz
//	* FilterBandwidth400Hz
//	* FilterBandwidth200Hz
//	* FilterBandwidth50Hz
func (device *AccelerometerBricklet) GetConfiguration() (dataRate DataRate, fullScale FullScale, filterBandwidth FilterBandwidth, err error) {
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetConfiguration), buf.Bytes())
    if err != nil {
        return dataRate, fullScale, filterBandwidth, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader

        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return dataRate, fullScale, filterBandwidth, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &dataRate)
	binary.Read(resultBuf, binary.LittleEndian, &fullScale)
	binary.Read(resultBuf, binary.LittleEndian, &filterBandwidth)

    }

    return dataRate, fullScale, filterBandwidth, nil
}

// Enables the LED on the Bricklet.
func (device *AccelerometerBricklet) LEDOn() (err error) {
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Set(uint8(FunctionLEDOn), buf.Bytes())
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

// Disables the LED on the Bricklet.
func (device *AccelerometerBricklet) LEDOff() (err error) {
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Set(uint8(FunctionLEDOff), buf.Bytes())
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

// Returns *true* if the LED is enabled, *false* otherwise.
func (device *AccelerometerBricklet) IsLEDOn() (on bool, err error) {
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionIsLEDOn), buf.Bytes())
    if err != nil {
        return on, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader

        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return on, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &on)

    }

    return on, nil
}

// Returns the UID, the UID where the Bricklet is connected to,
// the position, the hardware and firmware version as well as the
// device identifier.
// 
// The position can be 'a', 'b', 'c' or 'd'.
// 
// The device identifier numbers can be found `here <device_identifier>`.
// |device_identifier_constant|
func (device *AccelerometerBricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
