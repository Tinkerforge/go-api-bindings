/* ***********************************************************
 * This file was automatically generated on 2019-01-07.      *
 *                                                           *
 * Go Bindings Version 2.0.1                                 *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


//Measures DC voltage between 0V and 42V‍.
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/AnalogInV3_Bricklet_Go.html.
package analog_in_v3_bricklet

import (
	"encoding/binary"
	"bytes"
    . "github.com/Tinkerforge/go-api-bindings/internal"
    "github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function uint8

const (
    FunctionGetVoltage Function = 1
	FunctionSetVoltageCallbackConfiguration Function = 2
	FunctionGetVoltageCallbackConfiguration Function = 3
	FunctionSetOversampling Function = 5
	FunctionGetOversampling Function = 6
	FunctionSetCalibration Function = 7
	FunctionGetCalibration Function = 8
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
	FunctionCallbackVoltage Function = 4
)

type ThresholdOption rune

const (
    ThresholdOptionOff ThresholdOption = 'x'
	ThresholdOptionOutside ThresholdOption = 'o'
	ThresholdOptionInside ThresholdOption = 'i'
	ThresholdOptionSmaller ThresholdOption = '<'
	ThresholdOptionGreater ThresholdOption = '>'
)

type Oversampling uint8

const (
    Oversampling32 Oversampling = 0
	Oversampling64 Oversampling = 1
	Oversampling128 Oversampling = 2
	Oversampling256 Oversampling = 3
	Oversampling512 Oversampling = 4
	Oversampling1024 Oversampling = 5
	Oversampling2048 Oversampling = 6
	Oversampling4096 Oversampling = 7
	Oversampling8192 Oversampling = 8
	Oversampling16384 Oversampling = 9
)

type BootloaderMode uint8

const (
    BootloaderModeBootloader BootloaderMode = 0
	BootloaderModeFirmware BootloaderMode = 1
	BootloaderModeBootloaderWaitForReboot BootloaderMode = 2
	BootloaderModeFirmwareWaitForReboot BootloaderMode = 3
	BootloaderModeFirmwareWaitForEraseAndReboot BootloaderMode = 4
)

type BootloaderStatus uint8

const (
    BootloaderStatusOK BootloaderStatus = 0
	BootloaderStatusInvalidMode BootloaderStatus = 1
	BootloaderStatusNoChange BootloaderStatus = 2
	BootloaderStatusEntryFunctionNotPresent BootloaderStatus = 3
	BootloaderStatusDeviceIdentifierIncorrect BootloaderStatus = 4
	BootloaderStatusCRCMismatch BootloaderStatus = 5
)

type StatusLEDConfig uint8

const (
    StatusLEDConfigOff StatusLEDConfig = 0
	StatusLEDConfigOn StatusLEDConfig = 1
	StatusLEDConfigShowHeartbeat StatusLEDConfig = 2
	StatusLEDConfigShowStatus StatusLEDConfig = 3
)

type AnalogInV3Bricklet struct{
	device Device
}
const DeviceIdentifier = 295
const DeviceDisplayName = "Analog In Bricklet 3.0"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (AnalogInV3Bricklet, error) {
    internalIPCon := ipcon.GetInternalHandle().(IPConnection)
    dev, err := NewDevice([3]uint8{ 2,0,0 }, uid, &internalIPCon, 0)
    if err != nil {
        return AnalogInV3Bricklet{}, err
    }
    dev.ResponseExpected[FunctionGetVoltage] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetVoltageCallbackConfiguration] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetVoltageCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetOversampling] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetOversampling] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetCalibration] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetCalibration] = ResponseExpectedFlagAlwaysTrue;
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
    return AnalogInV3Bricklet{dev}, nil
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
func (device *AnalogInV3Bricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *AnalogInV3Bricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
    return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *AnalogInV3Bricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *AnalogInV3Bricklet) GetAPIVersion() [3]uint8 {
    return device.device.GetAPIVersion()
}

// This callback is triggered periodically according to the configuration set by
	// SetVoltageCallbackConfiguration.
	// 
	// The parameter is the same as GetVoltage.
func (device *AnalogInV3Bricklet) RegisterVoltageCallback(fn func(uint16)) uint64 {
            wrapper := func(byteSlice []byte) {
                buf := bytes.NewBuffer(byteSlice[8:])
                var voltage uint16
                binary.Read(buf, binary.LittleEndian, &voltage)
                fn(voltage)
            }
    return device.device.RegisterCallback(uint8(FunctionCallbackVoltage), wrapper)
}

//Remove a registered Voltage callback.
func (device *AnalogInV3Bricklet) DeregisterVoltageCallback(callbackID uint64) {
    device.device.DeregisterCallback(uint8(FunctionCallbackVoltage), callbackID)
}


// Returns the measured voltage. The value is in mV and
	// between 0V and 42V. The resolution is approximately 10mV to 1mV
	// depending on the oversampling configuration (SetOversampling).
	// 
	// 
	// If you want to get the value periodically, it is recommended to use the
	// RegisterVoltageCallback callback. You can set the callback configuration
	// with SetVoltageCallbackConfiguration.
func (device *AnalogInV3Bricklet) GetVoltage() (voltage uint16, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetVoltage), buf.Bytes())
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

// The period in ms is the period with which the RegisterVoltageCallback callback is triggered
	// periodically. A value of 0 turns the callback off.
	// 
	// If the `value has to change`-parameter is set to true, the callback is only
	// triggered after the value has changed. If the value didn't change
	// within the period, the callback is triggered immediately on change.
	// 
	// If it is set to false, the callback is continuously triggered with the period,
	// independent of the value.
	// 
	// It is furthermore possible to constrain the callback with thresholds.
	// 
	// The `option`-parameter together with min/max sets a threshold for the RegisterVoltageCallback callback.
	// 
	// The following options are possible:
	// 
	//  Option| Description
	//  --- | --- 
	//  'x'|    Threshold is turned off
	//  'o'|    Threshold is triggered when the value is *outside* the min and max values
	//  'i'|    Threshold is triggered when the value is *inside* or equal to the min and max values
	//  '<'|    Threshold is triggered when the value is smaller than the min value (max is ignored)
	//  '>'|    Threshold is triggered when the value is greater than the min value (max is ignored)
	// 
	// If the option is set to 'x' (threshold turned off) the callback is triggered with the fixed period.
	// 
	// The default value is (0, false, 'x', 0, 0).
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *AnalogInV3Bricklet) SetVoltageCallbackConfiguration(period uint32, valueHasToChange bool, option ThresholdOption, min uint16, max uint16) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, period);
	binary.Write(&buf, binary.LittleEndian, valueHasToChange);
	binary.Write(&buf, binary.LittleEndian, option);
	binary.Write(&buf, binary.LittleEndian, min);
	binary.Write(&buf, binary.LittleEndian, max);

    resultBytes, err := device.device.Set(uint8(FunctionSetVoltageCallbackConfiguration), buf.Bytes())
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

// Returns the callback configuration as set by SetVoltageCallbackConfiguration.
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *AnalogInV3Bricklet) GetVoltageCallbackConfiguration() (period uint32, valueHasToChange bool, option ThresholdOption, min uint16, max uint16, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetVoltageCallbackConfiguration), buf.Bytes())
    if err != nil {
        return period, valueHasToChange, option, min, max, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return period, valueHasToChange, option, min, max, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &period)
	binary.Read(resultBuf, binary.LittleEndian, &valueHasToChange)
	binary.Read(resultBuf, binary.LittleEndian, &option)
	binary.Read(resultBuf, binary.LittleEndian, &min)
	binary.Read(resultBuf, binary.LittleEndian, &max)

    }
    
    return period, valueHasToChange, option, min, max, nil
}

// Sets the oversampling between 32x and 16384x. The Bricklet
	// takes one 12bit sample every 17.5µs. Thus an oversampling
	// of 32x is equivalent to an integration time of 0.56ms and
	// a oversampling of 16384x is equivalent to an integration
	// time of 286ms.
	// 
	// The oversampling uses the moving average principle. A
	// new value is always calculated once per millisecond.
	// 
	// With increased oversampling the noise decreases. With decreased
	// oversampling the reaction time increases (changes in voltage will be
	// measured faster).
	// 
	// The default value is 4096x.
//
// Associated constants:
//
//	* Oversampling32
//	* Oversampling64
//	* Oversampling128
//	* Oversampling256
//	* Oversampling512
//	* Oversampling1024
//	* Oversampling2048
//	* Oversampling4096
//	* Oversampling8192
//	* Oversampling16384
func (device *AnalogInV3Bricklet) SetOversampling(oversampling Oversampling) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, oversampling);

    resultBytes, err := device.device.Set(uint8(FunctionSetOversampling), buf.Bytes())
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

// Returns the oversampling value as set by SetOversampling.
//
// Associated constants:
//
//	* Oversampling32
//	* Oversampling64
//	* Oversampling128
//	* Oversampling256
//	* Oversampling512
//	* Oversampling1024
//	* Oversampling2048
//	* Oversampling4096
//	* Oversampling8192
//	* Oversampling16384
func (device *AnalogInV3Bricklet) GetOversampling() (oversampling Oversampling, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetOversampling), buf.Bytes())
    if err != nil {
        return oversampling, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return oversampling, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &oversampling)

    }
    
    return oversampling, nil
}

// Sets a calibration for the measured voltage value.
	// The formula for the calibration is as follows::
	// 
	//  Calibrated Value = (Value + Offset) * Multiplier / Divisor
	// 
	// We recommend that you use the Brick Viewer to calibrate
	// the Bricklet. The calibration will be saved internally and only
	// has to be done once.
func (device *AnalogInV3Bricklet) SetCalibration(offset int16, multiplier uint16, divisor uint16) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, offset);
	binary.Write(&buf, binary.LittleEndian, multiplier);
	binary.Write(&buf, binary.LittleEndian, divisor);

    resultBytes, err := device.device.Set(uint8(FunctionSetCalibration), buf.Bytes())
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

// Returns the calibration as set by SetCalibration.
func (device *AnalogInV3Bricklet) GetCalibration() (offset int16, multiplier uint16, divisor uint16, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetCalibration), buf.Bytes())
    if err != nil {
        return offset, multiplier, divisor, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return offset, multiplier, divisor, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &offset)
	binary.Read(resultBuf, binary.LittleEndian, &multiplier)
	binary.Read(resultBuf, binary.LittleEndian, &divisor)

    }
    
    return offset, multiplier, divisor, nil
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
func (device *AnalogInV3Bricklet) GetSPITFPErrorCount() (errorCountAckChecksum uint32, errorCountMessageChecksum uint32, errorCountFrame uint32, errorCountOverflow uint32, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetSPITFPErrorCount), buf.Bytes())
    if err != nil {
        return errorCountAckChecksum, errorCountMessageChecksum, errorCountFrame, errorCountOverflow, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return errorCountAckChecksum, errorCountMessageChecksum, errorCountFrame, errorCountOverflow, BrickletError(header.ErrorCode)
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
func (device *AnalogInV3Bricklet) SetBootloaderMode(mode BootloaderMode) (status BootloaderStatus, err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, mode);

    resultBytes, err := device.device.Get(uint8(FunctionSetBootloaderMode), buf.Bytes())
    if err != nil {
        return status, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return status, BrickletError(header.ErrorCode)
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
func (device *AnalogInV3Bricklet) GetBootloaderMode() (mode BootloaderMode, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetBootloaderMode), buf.Bytes())
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

// Sets the firmware pointer for WriteFirmware. The pointer has
	// to be increased by chunks of size 64. The data is written to flash
	// every 4 chunks (which equals to one page of size 256).
	// 
	// This function is used by Brick Viewer during flashing. It should not be
	// necessary to call it in a normal user program.
func (device *AnalogInV3Bricklet) SetWriteFirmwarePointer(pointer uint32) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, pointer);

    resultBytes, err := device.device.Set(uint8(FunctionSetWriteFirmwarePointer), buf.Bytes())
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

// Writes 64 Bytes of firmware at the position as written by
	// SetWriteFirmwarePointer before. The firmware is written
	// to flash every 4 chunks.
	// 
	// You can only write firmware in bootloader mode.
	// 
	// This function is used by Brick Viewer during flashing. It should not be
	// necessary to call it in a normal user program.
func (device *AnalogInV3Bricklet) WriteFirmware(data [64]uint8) (status uint8, err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, data);

    resultBytes, err := device.device.Get(uint8(FunctionWriteFirmware), buf.Bytes())
    if err != nil {
        return status, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return status, BrickletError(header.ErrorCode)
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
func (device *AnalogInV3Bricklet) SetStatusLEDConfig(config StatusLEDConfig) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, config);

    resultBytes, err := device.device.Set(uint8(FunctionSetStatusLEDConfig), buf.Bytes())
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

// Returns the configuration as set by SetStatusLEDConfig
//
// Associated constants:
//
//	* StatusLEDConfigOff
//	* StatusLEDConfigOn
//	* StatusLEDConfigShowHeartbeat
//	* StatusLEDConfigShowStatus
func (device *AnalogInV3Bricklet) GetStatusLEDConfig() (config StatusLEDConfig, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetStatusLEDConfig), buf.Bytes())
    if err != nil {
        return config, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return config, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &config)

    }
    
    return config, nil
}

// Returns the temperature in °C as measured inside the microcontroller. The
	// value returned is not the ambient temperature!
	// 
	// The temperature is only proportional to the real temperature and it has bad
	// accuracy. Practically it is only useful as an indicator for
	// temperature changes.
func (device *AnalogInV3Bricklet) GetChipTemperature() (temperature int16, err error) {    
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

// Calling this function will reset the Bricklet. All configurations
	// will be lost.
	// 
	// After a reset you have to create new device objects,
	// calling functions on the existing ones will result in
	// undefined behavior!
func (device *AnalogInV3Bricklet) Reset() (err error) {    
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

// Writes a new UID into flash. If you want to set a new UID
	// you have to decode the Base58 encoded UID string into an
	// integer first.
	// 
	// We recommend that you use Brick Viewer to change the UID.
func (device *AnalogInV3Bricklet) WriteUID(uid uint32) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, uid);

    resultBytes, err := device.device.Set(uint8(FunctionWriteUID), buf.Bytes())
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

// Returns the current UID as an integer. Encode as
	// Base58 to get the usual string version.
func (device *AnalogInV3Bricklet) ReadUID() (uid uint32, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionReadUID), buf.Bytes())
    if err != nil {
        return uid, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return uid, BrickletError(header.ErrorCode)
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
	// The position can be 'a', 'b', 'c' or 'd'.
	// 
	// The device identifier numbers can be found `here <device_identifier>`.
	// |device_identifier_constant|
func (device *AnalogInV3Bricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {    
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
