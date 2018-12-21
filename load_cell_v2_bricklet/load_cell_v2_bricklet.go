/* ***********************************************************
 * This file was automatically generated on 2018-12-21.      *
 *                                                           *
 * Go Bindings Version 2.0.0                                 *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


//Measures weight with a load cell.
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/LoadCellV2_Bricklet_Go.html.
package load_cell_v2_bricklet

import (
	"encoding/binary"
	"bytes"
    . "github.com/Tinkerforge/go-api-bindings/internal"
    "github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function uint8

const (
    FunctionGetWeight Function = 1
	FunctionSetWeightCallbackConfiguration Function = 2
	FunctionGetWeightCallbackConfiguration Function = 3
	FunctionSetMovingAverage Function = 5
	FunctionGetMovingAverage Function = 6
	FunctionSetInfoLEDConfig Function = 7
	FunctionGetInfoLEDConfig Function = 8
	FunctionCalibrate Function = 9
	FunctionTare Function = 10
	FunctionSetConfiguration Function = 11
	FunctionGetConfiguration Function = 12
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
	FunctionCallbackWeight Function = 4
)

type ThresholdOption rune

const (
    ThresholdOptionOff ThresholdOption = 'x'
	ThresholdOptionOutside ThresholdOption = 'o'
	ThresholdOptionInside ThresholdOption = 'i'
	ThresholdOptionSmaller ThresholdOption = '<'
	ThresholdOptionGreater ThresholdOption = '>'
)

type InfoLEDConfig uint8

const (
    InfoLEDConfigOff InfoLEDConfig = 0
	InfoLEDConfigOn InfoLEDConfig = 1
	InfoLEDConfigShowHeartbeat InfoLEDConfig = 2
)

type Rate uint8

const (
    Rate10Hz Rate = 0
	Rate80Hz Rate = 1
)

type Gain uint8

const (
    Gain128x Gain = 0
	Gain64x Gain = 1
	Gain32x Gain = 2
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

type LoadCellV2Bricklet struct{
	device Device
}
const DeviceIdentifier = 2104
const DeviceDisplayName = "Load Cell Bricklet 2.0"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (LoadCellV2Bricklet, error) {
    internalIPCon := ipcon.GetInternalHandle().(IPConnection)
    dev, err := NewDevice([3]uint8{ 2,0,0 }, uid, &internalIPCon, 0)
    if err != nil {
        return LoadCellV2Bricklet{}, err
    }
    dev.ResponseExpected[FunctionGetWeight] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetWeightCallbackConfiguration] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetWeightCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetMovingAverage] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetMovingAverage] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetInfoLEDConfig] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetInfoLEDConfig] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionCalibrate] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionTare] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionSetConfiguration] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetConfiguration] = ResponseExpectedFlagAlwaysTrue;
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
    return LoadCellV2Bricklet{dev}, nil
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
func (device *LoadCellV2Bricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *LoadCellV2Bricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
    return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *LoadCellV2Bricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *LoadCellV2Bricklet) GetAPIVersion() [3]uint8 {
    return device.device.GetAPIVersion()
}

// This callback is triggered periodically according to the configuration set by
	// SetWeightCallbackConfiguration.
	// 
	// The parameter is the same as GetWeight.
func (device *LoadCellV2Bricklet) RegisterWeightCallback(fn func(int32)) uint64 {
            wrapper := func(byteSlice []byte) {
                buf := bytes.NewBuffer(byteSlice[8:])
                var weight int32
                binary.Read(buf, binary.LittleEndian, &weight)
                fn(weight)
            }
    return device.device.RegisterCallback(uint8(FunctionCallbackWeight), wrapper)
}

//Remove a registered Weight callback.
func (device *LoadCellV2Bricklet) DeregisterWeightCallback(callbackID uint64) {
    device.device.DeregisterCallback(uint8(FunctionCallbackWeight), callbackID)
}


// Returns the currently measured weight in grams.
	// 
	// 
	// If you want to get the value periodically, it is recommended to use the
	// RegisterWeightCallback callback. You can set the callback configuration
	// with SetWeightCallbackConfiguration.
func (device *LoadCellV2Bricklet) GetWeight() (weight int32, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetWeight), buf.Bytes())
    if err != nil {
        return weight, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return weight, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &weight)

    }
    
    return weight, nil
}

// The period in ms is the period with which the RegisterWeightCallback callback is triggered
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
	// The `option`-parameter together with min/max sets a threshold for the RegisterWeightCallback callback.
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
func (device *LoadCellV2Bricklet) SetWeightCallbackConfiguration(period uint32, valueHasToChange bool, option ThresholdOption, min int32, max int32) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, period);
	binary.Write(&buf, binary.LittleEndian, valueHasToChange);
	binary.Write(&buf, binary.LittleEndian, option);
	binary.Write(&buf, binary.LittleEndian, min);
	binary.Write(&buf, binary.LittleEndian, max);

    resultBytes, err := device.device.Set(uint8(FunctionSetWeightCallbackConfiguration), buf.Bytes())
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

// Returns the callback configuration as set by SetWeightCallbackConfiguration.
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *LoadCellV2Bricklet) GetWeightCallbackConfiguration() (period uint32, valueHasToChange bool, option ThresholdOption, min int32, max int32, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetWeightCallbackConfiguration), buf.Bytes())
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

// Sets the length of a https://en.wikipedia.org/wiki/Moving_average
	// for the weight value.
	// 
	// Setting the length to 1 will turn the averaging off. With less
	// averaging, there is more noise on the data.
	// 
	// The range for the averaging is 1-100.
	// 
	// The default value is 4.
func (device *LoadCellV2Bricklet) SetMovingAverage(average uint16) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, average);

    resultBytes, err := device.device.Set(uint8(FunctionSetMovingAverage), buf.Bytes())
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

// Returns the length moving average as set by SetMovingAverage.
func (device *LoadCellV2Bricklet) GetMovingAverage() (average uint16, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetMovingAverage), buf.Bytes())
    if err != nil {
        return average, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return average, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &average)

    }
    
    return average, nil
}

// Configures the info LED to be either turned off, turned on, or blink in
	// heartbeat mode.
//
// Associated constants:
//
//	* InfoLEDConfigOff
//	* InfoLEDConfigOn
//	* InfoLEDConfigShowHeartbeat
func (device *LoadCellV2Bricklet) SetInfoLEDConfig(config InfoLEDConfig) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, config);

    resultBytes, err := device.device.Set(uint8(FunctionSetInfoLEDConfig), buf.Bytes())
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

// Returns the LED configuration as set by SetInfoLEDConfig
//
// Associated constants:
//
//	* InfoLEDConfigOff
//	* InfoLEDConfigOn
//	* InfoLEDConfigShowHeartbeat
func (device *LoadCellV2Bricklet) GetInfoLEDConfig() (config InfoLEDConfig, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetInfoLEDConfig), buf.Bytes())
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

// To calibrate your Load Cell Bricklet 2.0 you have to
	// 
	// * empty the scale and call this function with 0 and
	// * add a known weight to the scale and call this function with the weight in
	//   grams.
	// 
	// The calibration is saved in the flash of the Bricklet and only
	// needs to be done once.
	// 
	// We recommend to use the Brick Viewer for calibration, you don't need
	// to call this function in your source code.
func (device *LoadCellV2Bricklet) Calibrate(weight uint32) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, weight);

    resultBytes, err := device.device.Set(uint8(FunctionCalibrate), buf.Bytes())
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

// Sets the currently measured weight as tare weight.
func (device *LoadCellV2Bricklet) Tare() (err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Set(uint8(FunctionTare), buf.Bytes())
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
	// The default rate is 10Hz and the default gain is 128x.
//
// Associated constants:
//
//	* Rate10Hz
//	* Rate80Hz
//	* Gain128x
//	* Gain64x
//	* Gain32x
func (device *LoadCellV2Bricklet) SetConfiguration(rate Rate, gain Gain) (err error) {    
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
//	* Rate10Hz
//	* Rate80Hz
//	* Gain128x
//	* Gain64x
//	* Gain32x
func (device *LoadCellV2Bricklet) GetConfiguration() (rate Rate, gain Gain, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetConfiguration), buf.Bytes())
    if err != nil {
        return rate, gain, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return rate, gain, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &rate)
	binary.Read(resultBuf, binary.LittleEndian, &gain)

    }
    
    return rate, gain, nil
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
func (device *LoadCellV2Bricklet) GetSPITFPErrorCount() (errorCountAckChecksum uint32, errorCountMessageChecksum uint32, errorCountFrame uint32, errorCountOverflow uint32, err error) {    
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
func (device *LoadCellV2Bricklet) SetBootloaderMode(mode BootloaderMode) (status BootloaderStatus, err error) {    
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
func (device *LoadCellV2Bricklet) GetBootloaderMode() (mode BootloaderMode, err error) {    
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
func (device *LoadCellV2Bricklet) SetWriteFirmwarePointer(pointer uint32) (err error) {    
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
func (device *LoadCellV2Bricklet) WriteFirmware(data [64]uint8) (status uint8, err error) {    
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
func (device *LoadCellV2Bricklet) SetStatusLEDConfig(config StatusLEDConfig) (err error) {    
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
func (device *LoadCellV2Bricklet) GetStatusLEDConfig() (config StatusLEDConfig, err error) {    
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
func (device *LoadCellV2Bricklet) GetChipTemperature() (temperature int16, err error) {    
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
func (device *LoadCellV2Bricklet) Reset() (err error) {    
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
func (device *LoadCellV2Bricklet) WriteUID(uid uint32) (err error) {    
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
func (device *LoadCellV2Bricklet) ReadUID() (uid uint32, err error) {    
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
func (device *LoadCellV2Bricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {    
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
