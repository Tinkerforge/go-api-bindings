/* ***********************************************************
 * This file was automatically generated on 2018-12-18.      *
 *                                                           *
 * Go Bindings Version 2.0.0                                 *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


//Battery-backed real-time clock.
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/RealTimeClockV2_Bricklet_Go.html.
package real_time_clock_v2_bricklet

import (
	"encoding/binary"
	"bytes"
    . "github.com/Tinkerforge/go-api-bindings/internal"
    "github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function uint8

const (
    FunctionSetDateTime Function = 1
	FunctionGetDateTime Function = 2
	FunctionGetTimestamp Function = 3
	FunctionSetOffset Function = 4
	FunctionGetOffset Function = 5
	FunctionSetDateTimeCallbackConfiguration Function = 6
	FunctionGetDateTimeCallbackConfiguration Function = 7
	FunctionSetAlarm Function = 8
	FunctionGetAlarm Function = 9
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
	FunctionCallbackDateTime Function = 10
	FunctionCallbackAlarm Function = 11
)

type Weekday uint8

const (
    WeekdayMonday Weekday = 1
	WeekdayTuesday Weekday = 2
	WeekdayWednesday Weekday = 3
	WeekdayThursday Weekday = 4
	WeekdayFriday Weekday = 5
	WeekdaySaturday Weekday = 6
	WeekdaySunday Weekday = 7
)

type AlarmMatch int8

const (
    AlarmMatchDisabled AlarmMatch = -1
)

type AlarmInterval int32

const (
    AlarmIntervalDisabled AlarmInterval = -1
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

type RealTimeClockV2Bricklet struct{
	device Device
}
const DeviceIdentifier = 2106
const DeviceDisplayName = "Real-Time Clock Bricklet 2.0"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (RealTimeClockV2Bricklet, error) {
    internalIPCon := ipcon.GetInternalHandle().(IPConnection)
    dev, err := NewDevice([3]uint8{ 2,0,0 }, uid, &internalIPCon, 0)
    if err != nil {
        return RealTimeClockV2Bricklet{}, err
    }
    dev.ResponseExpected[FunctionSetDateTime] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetDateTime] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetTimestamp] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetOffset] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetOffset] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetDateTimeCallbackConfiguration] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetDateTimeCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetAlarm] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetAlarm] = ResponseExpectedFlagAlwaysTrue;
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
    return RealTimeClockV2Bricklet{dev}, nil
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
func (device *RealTimeClockV2Bricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *RealTimeClockV2Bricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
    return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *RealTimeClockV2Bricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *RealTimeClockV2Bricklet) GetAPIVersion() [3]uint8 {
    return device.device.GetAPIVersion()
}

// This callback is triggered periodically with the period that is set by
	// SetDateTimeCallbackConfiguration. The parameters are the
	// same as for GetDateTime.
func (device *RealTimeClockV2Bricklet) RegisterDateTimeCallback(fn func(uint16, uint8, uint8, uint8, uint8, uint8, uint8, Weekday, int64)) uint64 {
            wrapper := func(byteSlice []byte) {
                buf := bytes.NewBuffer(byteSlice[8:])
                var year uint16
var month uint8
var day uint8
var hour uint8
var minute uint8
var second uint8
var centisecond uint8
var weekday Weekday
var timestamp int64
                binary.Read(buf, binary.LittleEndian, &year)
binary.Read(buf, binary.LittleEndian, &month)
binary.Read(buf, binary.LittleEndian, &day)
binary.Read(buf, binary.LittleEndian, &hour)
binary.Read(buf, binary.LittleEndian, &minute)
binary.Read(buf, binary.LittleEndian, &second)
binary.Read(buf, binary.LittleEndian, &centisecond)
binary.Read(buf, binary.LittleEndian, &weekday)
binary.Read(buf, binary.LittleEndian, &timestamp)
                fn(year, month, day, hour, minute, second, centisecond, weekday, timestamp)
            }
    return device.device.RegisterCallback(uint8(FunctionCallbackDateTime), wrapper)
}

//Remove a registered Date Time callback.
func (device *RealTimeClockV2Bricklet) DeregisterDateTimeCallback(callbackID uint64) {
    device.device.DeregisterCallback(uint8(FunctionCallbackDateTime), callbackID)
}


// This callback is triggered every time the current date and time matches the
	// configured alarm (see SetAlarm). The parameters are the same
	// as for GetDateTime.
func (device *RealTimeClockV2Bricklet) RegisterAlarmCallback(fn func(uint16, uint8, uint8, uint8, uint8, uint8, uint8, Weekday, int64)) uint64 {
            wrapper := func(byteSlice []byte) {
                buf := bytes.NewBuffer(byteSlice[8:])
                var year uint16
var month uint8
var day uint8
var hour uint8
var minute uint8
var second uint8
var centisecond uint8
var weekday Weekday
var timestamp int64
                binary.Read(buf, binary.LittleEndian, &year)
binary.Read(buf, binary.LittleEndian, &month)
binary.Read(buf, binary.LittleEndian, &day)
binary.Read(buf, binary.LittleEndian, &hour)
binary.Read(buf, binary.LittleEndian, &minute)
binary.Read(buf, binary.LittleEndian, &second)
binary.Read(buf, binary.LittleEndian, &centisecond)
binary.Read(buf, binary.LittleEndian, &weekday)
binary.Read(buf, binary.LittleEndian, &timestamp)
                fn(year, month, day, hour, minute, second, centisecond, weekday, timestamp)
            }
    return device.device.RegisterCallback(uint8(FunctionCallbackAlarm), wrapper)
}

//Remove a registered Alarm callback.
func (device *RealTimeClockV2Bricklet) DeregisterAlarmCallback(callbackID uint64) {
    device.device.DeregisterCallback(uint8(FunctionCallbackAlarm), callbackID)
}


// Sets the current date (including weekday) and the current time with hundredths
	// of a second resolution.
	// 
	// Possible value ranges:
	// 
	// * Year: 2000 to 2099
	// * Month: 1 to 12 (January to December)
	// * Day: 1 to 31
	// * Hour: 0 to 23
	// * Minute: 0 to 59
	// * Second: 0 to 59
	// * Centisecond: 0 to 99
	// * Weekday: 1 to 7 (Monday to Sunday)
	// 
	// If the backup battery is installed then the real-time clock keeps date and
	// time even if the Bricklet is not powered by a Brick.
	// 
	// The real-time clock handles leap year and inserts the 29th of February
	// accordingly. But leap seconds, time zones and daylight saving time are not
	// handled.
//
// Associated constants:
//
//	* WeekdayMonday
//	* WeekdayTuesday
//	* WeekdayWednesday
//	* WeekdayThursday
//	* WeekdayFriday
//	* WeekdaySaturday
//	* WeekdaySunday
func (device *RealTimeClockV2Bricklet) SetDateTime(year uint16, month uint8, day uint8, hour uint8, minute uint8, second uint8, centisecond uint8, weekday Weekday) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, year);
	binary.Write(&buf, binary.LittleEndian, month);
	binary.Write(&buf, binary.LittleEndian, day);
	binary.Write(&buf, binary.LittleEndian, hour);
	binary.Write(&buf, binary.LittleEndian, minute);
	binary.Write(&buf, binary.LittleEndian, second);
	binary.Write(&buf, binary.LittleEndian, centisecond);
	binary.Write(&buf, binary.LittleEndian, weekday);

    resultBytes, err := device.device.Set(uint8(FunctionSetDateTime), buf.Bytes())
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

// Returns the current date (including weekday) and the current time of the
	// real-time clock with hundredths of a second resolution.
	// 
	// The timestamp represents the current date and the the current time of the
	// real-time clock converted to milliseconds.
//
// Associated constants:
//
//	* WeekdayMonday
//	* WeekdayTuesday
//	* WeekdayWednesday
//	* WeekdayThursday
//	* WeekdayFriday
//	* WeekdaySaturday
//	* WeekdaySunday
func (device *RealTimeClockV2Bricklet) GetDateTime() (year uint16, month uint8, day uint8, hour uint8, minute uint8, second uint8, centisecond uint8, weekday Weekday, timestamp int64, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetDateTime), buf.Bytes())
    if err != nil {
        return year, month, day, hour, minute, second, centisecond, weekday, timestamp, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return year, month, day, hour, minute, second, centisecond, weekday, timestamp, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &year)
	binary.Read(resultBuf, binary.LittleEndian, &month)
	binary.Read(resultBuf, binary.LittleEndian, &day)
	binary.Read(resultBuf, binary.LittleEndian, &hour)
	binary.Read(resultBuf, binary.LittleEndian, &minute)
	binary.Read(resultBuf, binary.LittleEndian, &second)
	binary.Read(resultBuf, binary.LittleEndian, &centisecond)
	binary.Read(resultBuf, binary.LittleEndian, &weekday)
	binary.Read(resultBuf, binary.LittleEndian, &timestamp)

    }
    
    return year, month, day, hour, minute, second, centisecond, weekday, timestamp, nil
}

// Returns the current date and the time of the real-time clock converted to
	// milliseconds. The timestamp has an effective resolution of hundredths of a
	// second.
func (device *RealTimeClockV2Bricklet) GetTimestamp() (timestamp int64, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetTimestamp), buf.Bytes())
    if err != nil {
        return timestamp, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return timestamp, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &timestamp)

    }
    
    return timestamp, nil
}

// Sets the offset the real-time clock should compensate for in 2.17 ppm steps
	// between -277.76 ppm (-128) and +275.59 ppm (127).
	// 
	// The real-time clock time can deviate from the actual time due to the frequency
	// deviation of its 32.768 kHz crystal. Even without compensation (factory
	// default) the resulting time deviation should be at most ±20 ppm (±52.6
	// seconds per month).
	// 
	// This deviation can be calculated by comparing the same duration measured by the
	// real-time clock (``rtc_duration``) an accurate reference clock
	// (``ref_duration``).
	// 
	// For best results the configured offset should be set to 0 ppm first and then a
	// duration of at least 6 hours should be measured.
	// 
	// The new offset (``new_offset``) can be calculated from the currently configured
	// offset (``current_offset``) and the measured durations as follow::
	// 
	//   new_offset = current_offset - round(1000000 * (rtc_duration - ref_duration) / rtc_duration / 2.17)
	// 
	// If you want to calculate the offset, then we recommend using the calibration
	// dialog in Brick Viewer, instead of doing it manually.
	// 
	// The offset is saved in the EEPROM of the Bricklet and only needs to be
	// configured once.
func (device *RealTimeClockV2Bricklet) SetOffset(offset int8) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, offset);

    resultBytes, err := device.device.Set(uint8(FunctionSetOffset), buf.Bytes())
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

// Returns the offset as set by SetOffset.
func (device *RealTimeClockV2Bricklet) GetOffset() (offset int8, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetOffset), buf.Bytes())
    if err != nil {
        return offset, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return offset, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &offset)

    }
    
    return offset, nil
}

// Sets the period in ms with which the RegisterDateTimeCallback callback is triggered
	// periodically. A value of 0 turns the callback off.
	// 
	// The default value is 0.
func (device *RealTimeClockV2Bricklet) SetDateTimeCallbackConfiguration(period uint32) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, period);

    resultBytes, err := device.device.Set(uint8(FunctionSetDateTimeCallbackConfiguration), buf.Bytes())
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

// Returns the period as set by SetDateTimeCallbackConfiguration.
func (device *RealTimeClockV2Bricklet) GetDateTimeCallbackConfiguration() (period uint32, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetDateTimeCallbackConfiguration), buf.Bytes())
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

// Configures a repeatable alarm. The RegisterAlarmCallback callback is triggered if the
	// current date and time matches the configured alarm.
	// 
	// Setting a parameter to -1 means that it should be disabled and doesn't take part
	// in the match. Setting all parameters to -1 disables the alarm completely.
	// 
	// For example, to make the alarm trigger every day at 7:30 AM it can be
	// configured as (-1, -1, 7, 30, -1, -1, -1). The hour is set to match 7 and the
	// minute is set to match 30. The alarm is triggered if all enabled parameters
	// match.
	// 
	// The interval has a special role. It allows to make the alarm reconfigure itself.
	// This is useful if you need a repeated alarm that cannot be expressed by matching
	// the current date and time. For example, to make the alarm trigger every 23
	// seconds it can be configured as (-1, -1, -1, -1, -1, -1, 23). Internally the
	// Bricklet will take the current date and time, add 23 seconds to it and set the
	// result as its alarm. The first alarm will be triggered 23 seconds after the
	// call. Because the interval is not -1, the Bricklet will do the same again
	// internally, take the current date and time, add 23 seconds to it and set that
	// as its alarm. This results in a repeated alarm that triggers every 23 seconds.
	// 
	// The interval can also be used in combination with the other parameters. For
	// example, configuring the alarm as (-1, -1, 7, 30, -1, -1, 300) results in an
	// alarm that triggers every day at 7:30 AM and is then repeated every 5 minutes.
//
// Associated constants:
//
//	* AlarmMatchDisabled
//	* AlarmIntervalDisabled
func (device *RealTimeClockV2Bricklet) SetAlarm(month AlarmMatch, day AlarmMatch, hour AlarmMatch, minute AlarmMatch, second AlarmMatch, weekday AlarmMatch, interval AlarmInterval) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, month);
	binary.Write(&buf, binary.LittleEndian, day);
	binary.Write(&buf, binary.LittleEndian, hour);
	binary.Write(&buf, binary.LittleEndian, minute);
	binary.Write(&buf, binary.LittleEndian, second);
	binary.Write(&buf, binary.LittleEndian, weekday);
	binary.Write(&buf, binary.LittleEndian, interval);

    resultBytes, err := device.device.Set(uint8(FunctionSetAlarm), buf.Bytes())
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

// Returns the alarm configuration as set by SetAlarm.
//
// Associated constants:
//
//	* AlarmMatchDisabled
//	* AlarmIntervalDisabled
func (device *RealTimeClockV2Bricklet) GetAlarm() (month AlarmMatch, day AlarmMatch, hour AlarmMatch, minute AlarmMatch, second AlarmMatch, weekday AlarmMatch, interval AlarmInterval, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetAlarm), buf.Bytes())
    if err != nil {
        return month, day, hour, minute, second, weekday, interval, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return month, day, hour, minute, second, weekday, interval, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &month)
	binary.Read(resultBuf, binary.LittleEndian, &day)
	binary.Read(resultBuf, binary.LittleEndian, &hour)
	binary.Read(resultBuf, binary.LittleEndian, &minute)
	binary.Read(resultBuf, binary.LittleEndian, &second)
	binary.Read(resultBuf, binary.LittleEndian, &weekday)
	binary.Read(resultBuf, binary.LittleEndian, &interval)

    }
    
    return month, day, hour, minute, second, weekday, interval, nil
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
func (device *RealTimeClockV2Bricklet) GetSPITFPErrorCount() (errorCountAckChecksum uint32, errorCountMessageChecksum uint32, errorCountFrame uint32, errorCountOverflow uint32, err error) {    
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
func (device *RealTimeClockV2Bricklet) SetBootloaderMode(mode BootloaderMode) (status BootloaderStatus, err error) {    
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
func (device *RealTimeClockV2Bricklet) GetBootloaderMode() (mode BootloaderMode, err error) {    
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
func (device *RealTimeClockV2Bricklet) SetWriteFirmwarePointer(pointer uint32) (err error) {    
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
func (device *RealTimeClockV2Bricklet) WriteFirmware(data [64]uint8) (status uint8, err error) {    
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
func (device *RealTimeClockV2Bricklet) SetStatusLEDConfig(config StatusLEDConfig) (err error) {    
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
func (device *RealTimeClockV2Bricklet) GetStatusLEDConfig() (config StatusLEDConfig, err error) {    
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
func (device *RealTimeClockV2Bricklet) GetChipTemperature() (temperature int16, err error) {    
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
func (device *RealTimeClockV2Bricklet) Reset() (err error) {    
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
func (device *RealTimeClockV2Bricklet) WriteUID(uid uint32) (err error) {    
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
func (device *RealTimeClockV2Bricklet) ReadUID() (uid uint32, err error) {    
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
func (device *RealTimeClockV2Bricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {    
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
