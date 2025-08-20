/* ***********************************************************
 * This file was automatically generated on 2025-08-20.      *
 *                                                           *
 * Go Bindings Version 2.0.16                                *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/

// Determine position, velocity and altitude using GPS‍.
//
//
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/GPSV2_Bricklet_Go.html.
package gps_v2_bricklet

import (
	"bytes"
	"encoding/binary"
	"fmt"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionGetCoordinates                   Function = 1
	FunctionGetStatus                        Function = 2
	FunctionGetAltitude                      Function = 3
	FunctionGetMotion                        Function = 4
	FunctionGetDateTime                      Function = 5
	FunctionRestart                          Function = 6
	FunctionGetSatelliteSystemStatusLowLevel Function = 7
	FunctionGetSatelliteStatus               Function = 8
	FunctionSetFixLEDConfig                  Function = 9
	FunctionGetFixLEDConfig                  Function = 10
	FunctionSetCoordinatesCallbackPeriod     Function = 11
	FunctionGetCoordinatesCallbackPeriod     Function = 12
	FunctionSetStatusCallbackPeriod          Function = 13
	FunctionGetStatusCallbackPeriod          Function = 14
	FunctionSetAltitudeCallbackPeriod        Function = 15
	FunctionGetAltitudeCallbackPeriod        Function = 16
	FunctionSetMotionCallbackPeriod          Function = 17
	FunctionGetMotionCallbackPeriod          Function = 18
	FunctionSetDateTimeCallbackPeriod        Function = 19
	FunctionGetDateTimeCallbackPeriod        Function = 20
	FunctionSetSBASConfig                    Function = 27
	FunctionGetSBASConfig                    Function = 28
	FunctionGetSPITFPErrorCount              Function = 234
	FunctionSetBootloaderMode                Function = 235
	FunctionGetBootloaderMode                Function = 236
	FunctionSetWriteFirmwarePointer          Function = 237
	FunctionWriteFirmware                    Function = 238
	FunctionSetStatusLEDConfig               Function = 239
	FunctionGetStatusLEDConfig               Function = 240
	FunctionGetChipTemperature               Function = 242
	FunctionReset                            Function = 243
	FunctionWriteUID                         Function = 248
	FunctionReadUID                          Function = 249
	FunctionGetIdentity                      Function = 255
	FunctionCallbackPulsePerSecond           Function = 21
	FunctionCallbackCoordinates              Function = 22
	FunctionCallbackStatus                   Function = 23
	FunctionCallbackAltitude                 Function = 24
	FunctionCallbackMotion                   Function = 25
	FunctionCallbackDateTime                 Function = 26
)

type RestartType = uint8

const (
	RestartTypeHotStart     RestartType = 0
	RestartTypeWarmStart    RestartType = 1
	RestartTypeColdStart    RestartType = 2
	RestartTypeFactoryReset RestartType = 3
)

type SatelliteSystem = uint8

const (
	SatelliteSystemGPS     SatelliteSystem = 0
	SatelliteSystemGLONASS SatelliteSystem = 1
	SatelliteSystemGalileo SatelliteSystem = 2
)

type Fix = uint8

const (
	FixNoFix Fix = 1
	Fix2DFix Fix = 2
	Fix3DFix Fix = 3
)

type FixLEDConfig = uint8

const (
	FixLEDConfigOff           FixLEDConfig = 0
	FixLEDConfigOn            FixLEDConfig = 1
	FixLEDConfigShowHeartbeat FixLEDConfig = 2
	FixLEDConfigShowFix       FixLEDConfig = 3
	FixLEDConfigShowPPS       FixLEDConfig = 4
)

type SBAS = uint8

const (
	SBASEnabled  SBAS = 0
	SBASDisabled SBAS = 1
)

type BootloaderMode = uint8

const (
	BootloaderModeBootloader                    BootloaderMode = 0
	BootloaderModeFirmware                      BootloaderMode = 1
	BootloaderModeBootloaderWaitForReboot       BootloaderMode = 2
	BootloaderModeFirmwareWaitForReboot         BootloaderMode = 3
	BootloaderModeFirmwareWaitForEraseAndReboot BootloaderMode = 4
)

type BootloaderStatus = uint8

const (
	BootloaderStatusOK                        BootloaderStatus = 0
	BootloaderStatusInvalidMode               BootloaderStatus = 1
	BootloaderStatusNoChange                  BootloaderStatus = 2
	BootloaderStatusEntryFunctionNotPresent   BootloaderStatus = 3
	BootloaderStatusDeviceIdentifierIncorrect BootloaderStatus = 4
	BootloaderStatusCRCMismatch               BootloaderStatus = 5
)

type StatusLEDConfig = uint8

const (
	StatusLEDConfigOff           StatusLEDConfig = 0
	StatusLEDConfigOn            StatusLEDConfig = 1
	StatusLEDConfigShowHeartbeat StatusLEDConfig = 2
	StatusLEDConfigShowStatus    StatusLEDConfig = 3
)

type GPSV2Bricklet struct {
	device Device
}

const DeviceIdentifier = 276
const DeviceDisplayName = "GPS Bricklet 2.0"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (GPSV2Bricklet, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{2, 0, 1}, uid, &internalIPCon, 1, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return GPSV2Bricklet{}, err
	}
	dev.ResponseExpected[FunctionGetCoordinates] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetStatus] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetAltitude] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetMotion] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetDateTime] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionRestart] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetSatelliteSystemStatusLowLevel] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetSatelliteStatus] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetFixLEDConfig] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetFixLEDConfig] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetCoordinatesCallbackPeriod] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionGetCoordinatesCallbackPeriod] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetStatusCallbackPeriod] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionGetStatusCallbackPeriod] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetAltitudeCallbackPeriod] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionGetAltitudeCallbackPeriod] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetMotionCallbackPeriod] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionGetMotionCallbackPeriod] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetDateTimeCallbackPeriod] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionGetDateTimeCallbackPeriod] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetSBASConfig] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetSBASConfig] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetSPITFPErrorCount] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetBootloaderMode] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetBootloaderMode] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetWriteFirmwarePointer] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionWriteFirmware] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetStatusLEDConfig] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetStatusLEDConfig] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetChipTemperature] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionReset] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionWriteUID] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionReadUID] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetIdentity] = ResponseExpectedFlagAlwaysTrue
	return GPSV2Bricklet{dev}, nil
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
func (device *GPSV2Bricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *GPSV2Bricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *GPSV2Bricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *GPSV2Bricklet) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// This callback is triggered precisely once per second,
// see https://en.wikipedia.org/wiki/Pulse-per-second_signal.
//
// The precision of two subsequent pulses will be skewed because
// of the latency in the USB/RS485/Ethernet connection. But in the
// long run this will be very precise. For example a count of
// 3600 pulses will take exactly 1 hour.
func (device *GPSV2Bricklet) RegisterPulsePerSecondCallback(fn func()) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 8 {
			return
		}

		fn()
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackPulsePerSecond), wrapper)
}

// Remove a registered Pulse Per Second callback.
func (device *GPSV2Bricklet) DeregisterPulsePerSecondCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackPulsePerSecond), registrationId)
}

// This callback is triggered periodically with the period that is set by
// SetCoordinatesCallbackPeriod. The parameters are the same
// as for GetCoordinates.
//
// The RegisterCoordinatesCallback callback is only triggered if the coordinates changed
// since the last triggering and if there is currently a fix as indicated by
// GetStatus.
func (device *GPSV2Bricklet) RegisterCoordinatesCallback(fn func(uint32, rune, uint32, rune)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 18 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var latitude uint32
		var ns rune
		var longitude uint32
		var ew rune
		binary.Read(buf, binary.LittleEndian, &latitude)
		ns = rune(buf.Next(1)[0])
		binary.Read(buf, binary.LittleEndian, &longitude)
		ew = rune(buf.Next(1)[0])
		fn(latitude, ns, longitude, ew)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackCoordinates), wrapper)
}

// Remove a registered Coordinates callback.
func (device *GPSV2Bricklet) DeregisterCoordinatesCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackCoordinates), registrationId)
}

// This callback is triggered periodically with the period that is set by
// SetStatusCallbackPeriod. The parameters are the same
// as for GetStatus.
//
// The RegisterStatusCallback callback is only triggered if the status changed since the
// last triggering.
func (device *GPSV2Bricklet) RegisterStatusCallback(fn func(bool, uint8)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 10 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var hasFix bool
		var satellitesView uint8
		binary.Read(buf, binary.LittleEndian, &hasFix)
		binary.Read(buf, binary.LittleEndian, &satellitesView)
		fn(hasFix, satellitesView)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackStatus), wrapper)
}

// Remove a registered Status callback.
func (device *GPSV2Bricklet) DeregisterStatusCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackStatus), registrationId)
}

// This callback is triggered periodically with the period that is set by
// SetAltitudeCallbackPeriod. The parameters are the same
// as for GetAltitude.
//
// The RegisterAltitudeCallback callback is only triggered if the altitude changed since the
// last triggering and if there is currently a fix as indicated by
// GetStatus.
func (device *GPSV2Bricklet) RegisterAltitudeCallback(fn func(int32, int32)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 16 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var altitude int32
		var geoidalSeparation int32
		binary.Read(buf, binary.LittleEndian, &altitude)
		binary.Read(buf, binary.LittleEndian, &geoidalSeparation)
		fn(altitude, geoidalSeparation)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackAltitude), wrapper)
}

// Remove a registered Altitude callback.
func (device *GPSV2Bricklet) DeregisterAltitudeCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackAltitude), registrationId)
}

// This callback is triggered periodically with the period that is set by
// SetMotionCallbackPeriod. The parameters are the same
// as for GetMotion.
//
// The RegisterMotionCallback callback is only triggered if the motion changed since the
// last triggering and if there is currently a fix as indicated by
// GetStatus.
func (device *GPSV2Bricklet) RegisterMotionCallback(fn func(uint32, uint32)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 16 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var course uint32
		var speed uint32
		binary.Read(buf, binary.LittleEndian, &course)
		binary.Read(buf, binary.LittleEndian, &speed)
		fn(course, speed)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackMotion), wrapper)
}

// Remove a registered Motion callback.
func (device *GPSV2Bricklet) DeregisterMotionCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackMotion), registrationId)
}

// This callback is triggered periodically with the period that is set by
// SetDateTimeCallbackPeriod. The parameters are the same
// as for GetDateTime.
//
// The RegisterDateTimeCallback callback is only triggered if the date or time changed
// since the last triggering.
func (device *GPSV2Bricklet) RegisterDateTimeCallback(fn func(uint32, uint32)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 16 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var date uint32
		var time uint32
		binary.Read(buf, binary.LittleEndian, &date)
		binary.Read(buf, binary.LittleEndian, &time)
		fn(date, time)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackDateTime), wrapper)
}

// Remove a registered Date Time callback.
func (device *GPSV2Bricklet) DeregisterDateTimeCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackDateTime), registrationId)
}

// Returns the GPS coordinates. Latitude and longitude are given in the
// ``DD.dddddd°`` format, the value 57123468 means 57.123468°.
// The parameter ``ns`` and ``ew`` are the cardinal directions for
// latitude and longitude. Possible values for ``ns`` and ``ew`` are 'N', 'S', 'E'
// and 'W' (north, south, east and west).
//
// This data is only valid if there is currently a fix as indicated by
// GetStatus.
func (device *GPSV2Bricklet) GetCoordinates() (latitude uint32, ns rune, longitude uint32, ew rune, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetCoordinates), buf.Bytes())
	if err != nil {
		return latitude, ns, longitude, ew, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return latitude, ns, longitude, ew, DeviceError(header.ErrorCode)
		}

		if header.Length != 18 {
			return latitude, ns, longitude, ew, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 18)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &latitude)
		ns = rune(resultBuf.Next(1)[0])
		binary.Read(resultBuf, binary.LittleEndian, &longitude)
		ew = rune(resultBuf.Next(1)[0])

	}

	return latitude, ns, longitude, ew, nil
}

// Returns if a fix is currently available as well as the number of
// satellites that are in view.
//
// There is also a `green LED <gps_v2_bricklet_fix_led>` on the Bricklet that
// indicates the fix status.
func (device *GPSV2Bricklet) GetStatus() (hasFix bool, satellitesView uint8, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetStatus), buf.Bytes())
	if err != nil {
		return hasFix, satellitesView, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return hasFix, satellitesView, DeviceError(header.ErrorCode)
		}

		if header.Length != 10 {
			return hasFix, satellitesView, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &hasFix)
		binary.Read(resultBuf, binary.LittleEndian, &satellitesView)

	}

	return hasFix, satellitesView, nil
}

// Returns the current altitude and corresponding geoidal separation.
//
// This data is only valid if there is currently a fix as indicated by
// GetStatus.
func (device *GPSV2Bricklet) GetAltitude() (altitude int32, geoidalSeparation int32, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetAltitude), buf.Bytes())
	if err != nil {
		return altitude, geoidalSeparation, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return altitude, geoidalSeparation, DeviceError(header.ErrorCode)
		}

		if header.Length != 16 {
			return altitude, geoidalSeparation, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 16)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &altitude)
		binary.Read(resultBuf, binary.LittleEndian, &geoidalSeparation)

	}

	return altitude, geoidalSeparation, nil
}

// Returns the current course and speed. A course of 0° means the Bricklet is
// traveling north bound and 90° means it is traveling east bound.
//
// Please note that this only returns useful values if an actual movement
// is present.
//
// This data is only valid if there is currently a fix as indicated by
// GetStatus.
func (device *GPSV2Bricklet) GetMotion() (course uint32, speed uint32, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetMotion), buf.Bytes())
	if err != nil {
		return course, speed, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return course, speed, DeviceError(header.ErrorCode)
		}

		if header.Length != 16 {
			return course, speed, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 16)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &course)
		binary.Read(resultBuf, binary.LittleEndian, &speed)

	}

	return course, speed, nil
}

// Returns the current date and time. The date is
// given in the format ``ddmmyy`` and the time is given
// in the format ``hhmmss.sss``. For example, 140713 means
// 14.07.13 as date and 195923568 means 19:59:23.568 as time.
func (device *GPSV2Bricklet) GetDateTime() (date uint32, time uint32, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetDateTime), buf.Bytes())
	if err != nil {
		return date, time, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return date, time, DeviceError(header.ErrorCode)
		}

		if header.Length != 16 {
			return date, time, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 16)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &date)
		binary.Read(resultBuf, binary.LittleEndian, &time)

	}

	return date, time, nil
}

// Restarts the GPS Bricklet, the following restart types are available:
//
//  Value| Description
//  --- | ---
//  0| Hot start (use all available data in the NV store)
//  1| Warm start (don't use ephemeris at restart)
//  2| Cold start (don't use time| position| almanacs and ephemeris at restart)
//  3| Factory reset (clear all system/user configurations at restart)
//
// Associated constants:
//
//	* RestartTypeHotStart
//	* RestartTypeWarmStart
//	* RestartTypeColdStart
//	* RestartTypeFactoryReset
func (device *GPSV2Bricklet) Restart(restartType RestartType) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, restartType)

	resultBytes, err := device.device.Set(uint8(FunctionRestart), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		bytes.NewBuffer(resultBytes[8:])

	}

	return nil
}

// Returns the
//
// * satellite numbers list (up to 12 items)
// * fix value,
// * PDOP value,
// * HDOP value and
// * VDOP value
//
// for a given satellite system. Currently GPS and GLONASS are supported, Galileo
// is not yet supported.
//
// The GPS and GLONASS satellites have unique numbers and the satellite list gives
// the numbers of the satellites that are currently utilized. The number 0 is not
// a valid satellite number and can be ignored in the list.
//
// Associated constants:
//
//	* SatelliteSystemGPS
//	* SatelliteSystemGLONASS
//	* SatelliteSystemGalileo
//	* FixNoFix
//	* Fix2DFix
//	* Fix3DFix
func (device *GPSV2Bricklet) GetSatelliteSystemStatusLowLevel(satelliteSystem SatelliteSystem) (satelliteNumbersLength uint8, satelliteNumbersData [12]uint8, fix Fix, pdop uint16, hdop uint16, vdop uint16, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, satelliteSystem)

	resultBytes, err := device.device.Get(uint8(FunctionGetSatelliteSystemStatusLowLevel), buf.Bytes())
	if err != nil {
		return satelliteNumbersLength, satelliteNumbersData, fix, pdop, hdop, vdop, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return satelliteNumbersLength, satelliteNumbersData, fix, pdop, hdop, vdop, DeviceError(header.ErrorCode)
		}

		if header.Length != 28 {
			return satelliteNumbersLength, satelliteNumbersData, fix, pdop, hdop, vdop, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 28)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &satelliteNumbersLength)
		copy(satelliteNumbersData[:], ByteSliceToUint8Slice(resultBuf.Next(8*12/8)))
		binary.Read(resultBuf, binary.LittleEndian, &fix)
		binary.Read(resultBuf, binary.LittleEndian, &pdop)
		binary.Read(resultBuf, binary.LittleEndian, &hdop)
		binary.Read(resultBuf, binary.LittleEndian, &vdop)

	}

	return satelliteNumbersLength, satelliteNumbersData, fix, pdop, hdop, vdop, nil
}

// Returns the
//
// * satellite numbers list (up to 12 items)
// * fix value,
// * PDOP value,
// * HDOP value and
// * VDOP value
//
// for a given satellite system. Currently GPS and GLONASS are supported, Galileo
// is not yet supported.
//
// The GPS and GLONASS satellites have unique numbers and the satellite list gives
// the numbers of the satellites that are currently utilized. The number 0 is not
// a valid satellite number and can be ignored in the list.
func (device *GPSV2Bricklet) GetSatelliteSystemStatus(satelliteSystem SatelliteSystem) (satelliteNumbers []uint8, fix Fix, pdop uint16, hdop uint16, vdop uint16, err error) {
	buf, result, err := device.device.GetHighLevel(func() (LowLevelResult, error) {
		satelliteNumbersLength, satelliteNumbersData, fix, pdop, hdop, vdop, err := device.GetSatelliteSystemStatusLowLevel(satelliteSystem)

		if err != nil {
			return LowLevelResult{}, err
		}

		var lowLevelResults bytes.Buffer
		binary.Write(&lowLevelResults, binary.LittleEndian, fix)
		binary.Write(&lowLevelResults, binary.LittleEndian, pdop)
		binary.Write(&lowLevelResults, binary.LittleEndian, hdop)
		binary.Write(&lowLevelResults, binary.LittleEndian, vdop)

		return LowLevelResult{
			uint64(satelliteNumbersLength),
			uint64(0),
			Uint8SliceToByteSlice(satelliteNumbersData[:]),
			lowLevelResults.Bytes()}, nil
	},
		0,
		8)
	if err != nil {
		return ByteSliceToUint8Slice(buf), fix, pdop, hdop, vdop, err
	}
	resultBuf := bytes.NewBuffer(result)
	binary.Read(resultBuf, binary.LittleEndian, &fix)
	binary.Read(resultBuf, binary.LittleEndian, &pdop)
	binary.Read(resultBuf, binary.LittleEndian, &hdop)
	binary.Read(resultBuf, binary.LittleEndian, &vdop)
	return ByteSliceToUint8Slice(buf), fix, pdop, hdop, vdop, nil
}

// Returns the current elevation, azimuth and SNR
// for a given satellite and satellite system.
//
// The satellite number here always goes from 1 to 32. For GLONASS it corresponds to
// the satellites 65-96.
//
// Galileo is not yet supported.
//
// Associated constants:
//
//	* SatelliteSystemGPS
//	* SatelliteSystemGLONASS
//	* SatelliteSystemGalileo
func (device *GPSV2Bricklet) GetSatelliteStatus(satelliteSystem SatelliteSystem, satelliteNumber uint8) (elevation int16, azimuth int16, snr int16, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, satelliteSystem)
	binary.Write(&buf, binary.LittleEndian, satelliteNumber)

	resultBytes, err := device.device.Get(uint8(FunctionGetSatelliteStatus), buf.Bytes())
	if err != nil {
		return elevation, azimuth, snr, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return elevation, azimuth, snr, DeviceError(header.ErrorCode)
		}

		if header.Length != 14 {
			return elevation, azimuth, snr, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 14)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &elevation)
		binary.Read(resultBuf, binary.LittleEndian, &azimuth)
		binary.Read(resultBuf, binary.LittleEndian, &snr)

	}

	return elevation, azimuth, snr, nil
}

// Sets the fix LED configuration. By default the LED shows if
// the Bricklet got a GPS fix yet. If a fix is established the LED turns on.
// If there is no fix then the LED is turned off.
//
// You can also turn the LED permanently on/off, show a heartbeat or let it blink
// in sync with the PPS (pulse per second) output of the GPS module.
//
// If the Bricklet is in bootloader mode, the LED is off.
//
// Associated constants:
//
//	* FixLEDConfigOff
//	* FixLEDConfigOn
//	* FixLEDConfigShowHeartbeat
//	* FixLEDConfigShowFix
//	* FixLEDConfigShowPPS
func (device *GPSV2Bricklet) SetFixLEDConfig(config FixLEDConfig) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, config)

	resultBytes, err := device.device.Set(uint8(FunctionSetFixLEDConfig), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		bytes.NewBuffer(resultBytes[8:])

	}

	return nil
}

// Returns the configuration as set by SetFixLEDConfig
//
// Associated constants:
//
//	* FixLEDConfigOff
//	* FixLEDConfigOn
//	* FixLEDConfigShowHeartbeat
//	* FixLEDConfigShowFix
//	* FixLEDConfigShowPPS
func (device *GPSV2Bricklet) GetFixLEDConfig() (config FixLEDConfig, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetFixLEDConfig), buf.Bytes())
	if err != nil {
		return config, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return config, DeviceError(header.ErrorCode)
		}

		if header.Length != 9 {
			return config, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &config)

	}

	return config, nil
}

// Sets the period with which the RegisterCoordinatesCallback callback is triggered
// periodically. A value of 0 turns the callback off.
//
// The RegisterCoordinatesCallback callback is only triggered if the coordinates changed
// since the last triggering.
func (device *GPSV2Bricklet) SetCoordinatesCallbackPeriod(period uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period)

	resultBytes, err := device.device.Set(uint8(FunctionSetCoordinatesCallbackPeriod), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		bytes.NewBuffer(resultBytes[8:])

	}

	return nil
}

// Returns the period as set by SetCoordinatesCallbackPeriod.
func (device *GPSV2Bricklet) GetCoordinatesCallbackPeriod() (period uint32, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetCoordinatesCallbackPeriod), buf.Bytes())
	if err != nil {
		return period, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return period, DeviceError(header.ErrorCode)
		}

		if header.Length != 12 {
			return period, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &period)

	}

	return period, nil
}

// Sets the period with which the RegisterStatusCallback callback is triggered
// periodically. A value of 0 turns the callback off.
//
// The RegisterStatusCallback callback is only triggered if the status changed since the
// last triggering.
func (device *GPSV2Bricklet) SetStatusCallbackPeriod(period uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period)

	resultBytes, err := device.device.Set(uint8(FunctionSetStatusCallbackPeriod), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		bytes.NewBuffer(resultBytes[8:])

	}

	return nil
}

// Returns the period as set by SetStatusCallbackPeriod.
func (device *GPSV2Bricklet) GetStatusCallbackPeriod() (period uint32, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetStatusCallbackPeriod), buf.Bytes())
	if err != nil {
		return period, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return period, DeviceError(header.ErrorCode)
		}

		if header.Length != 12 {
			return period, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &period)

	}

	return period, nil
}

// Sets the period with which the RegisterAltitudeCallback callback is triggered
// periodically. A value of 0 turns the callback off.
//
// The RegisterAltitudeCallback callback is only triggered if the altitude changed since the
// last triggering.
func (device *GPSV2Bricklet) SetAltitudeCallbackPeriod(period uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period)

	resultBytes, err := device.device.Set(uint8(FunctionSetAltitudeCallbackPeriod), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		bytes.NewBuffer(resultBytes[8:])

	}

	return nil
}

// Returns the period as set by SetAltitudeCallbackPeriod.
func (device *GPSV2Bricklet) GetAltitudeCallbackPeriod() (period uint32, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetAltitudeCallbackPeriod), buf.Bytes())
	if err != nil {
		return period, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return period, DeviceError(header.ErrorCode)
		}

		if header.Length != 12 {
			return period, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &period)

	}

	return period, nil
}

// Sets the period with which the RegisterMotionCallback callback is triggered
// periodically. A value of 0 turns the callback off.
//
// The RegisterMotionCallback callback is only triggered if the motion changed since the
// last triggering.
func (device *GPSV2Bricklet) SetMotionCallbackPeriod(period uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period)

	resultBytes, err := device.device.Set(uint8(FunctionSetMotionCallbackPeriod), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		bytes.NewBuffer(resultBytes[8:])

	}

	return nil
}

// Returns the period as set by SetMotionCallbackPeriod.
func (device *GPSV2Bricklet) GetMotionCallbackPeriod() (period uint32, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetMotionCallbackPeriod), buf.Bytes())
	if err != nil {
		return period, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return period, DeviceError(header.ErrorCode)
		}

		if header.Length != 12 {
			return period, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &period)

	}

	return period, nil
}

// Sets the period with which the RegisterDateTimeCallback callback is triggered
// periodically. A value of 0 turns the callback off.
//
// The RegisterDateTimeCallback callback is only triggered if the date or time changed
// since the last triggering.
func (device *GPSV2Bricklet) SetDateTimeCallbackPeriod(period uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period)

	resultBytes, err := device.device.Set(uint8(FunctionSetDateTimeCallbackPeriod), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		bytes.NewBuffer(resultBytes[8:])

	}

	return nil
}

// Returns the period as set by SetDateTimeCallbackPeriod.
func (device *GPSV2Bricklet) GetDateTimeCallbackPeriod() (period uint32, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetDateTimeCallbackPeriod), buf.Bytes())
	if err != nil {
		return period, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return period, DeviceError(header.ErrorCode)
		}

		if header.Length != 12 {
			return period, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &period)

	}

	return period, nil
}

// If https://en.wikipedia.org/wiki/GNSS_augmentation#Satellite-based_augmentation_system is enabled,
// the position accuracy increases (if SBAS satellites are in view),
// but the update rate is limited to 5Hz. With SBAS disabled the update rate is increased to 10Hz.
//
// .. versionadded:: 2.0.2$nbsp;(Plugin)
//
// Associated constants:
//
//	* SBASEnabled
//	* SBASDisabled
func (device *GPSV2Bricklet) SetSBASConfig(sbasConfig SBAS) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, sbasConfig)

	resultBytes, err := device.device.Set(uint8(FunctionSetSBASConfig), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		bytes.NewBuffer(resultBytes[8:])

	}

	return nil
}

// Returns the SBAS configuration as set by SetSBASConfig
//
// .. versionadded:: 2.0.2$nbsp;(Plugin)
//
// Associated constants:
//
//	* SBASEnabled
//	* SBASDisabled
func (device *GPSV2Bricklet) GetSBASConfig() (sbasConfig SBAS, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetSBASConfig), buf.Bytes())
	if err != nil {
		return sbasConfig, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return sbasConfig, DeviceError(header.ErrorCode)
		}

		if header.Length != 9 {
			return sbasConfig, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &sbasConfig)

	}

	return sbasConfig, nil
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
func (device *GPSV2Bricklet) GetSPITFPErrorCount() (errorCountAckChecksum uint32, errorCountMessageChecksum uint32, errorCountFrame uint32, errorCountOverflow uint32, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetSPITFPErrorCount), buf.Bytes())
	if err != nil {
		return errorCountAckChecksum, errorCountMessageChecksum, errorCountFrame, errorCountOverflow, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return errorCountAckChecksum, errorCountMessageChecksum, errorCountFrame, errorCountOverflow, DeviceError(header.ErrorCode)
		}

		if header.Length != 24 {
			return errorCountAckChecksum, errorCountMessageChecksum, errorCountFrame, errorCountOverflow, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 24)
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
func (device *GPSV2Bricklet) SetBootloaderMode(mode BootloaderMode) (status BootloaderStatus, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, mode)

	resultBytes, err := device.device.Get(uint8(FunctionSetBootloaderMode), buf.Bytes())
	if err != nil {
		return status, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return status, DeviceError(header.ErrorCode)
		}

		if header.Length != 9 {
			return status, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
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
func (device *GPSV2Bricklet) GetBootloaderMode() (mode BootloaderMode, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetBootloaderMode), buf.Bytes())
	if err != nil {
		return mode, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return mode, DeviceError(header.ErrorCode)
		}

		if header.Length != 9 {
			return mode, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
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
func (device *GPSV2Bricklet) SetWriteFirmwarePointer(pointer uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, pointer)

	resultBytes, err := device.device.Set(uint8(FunctionSetWriteFirmwarePointer), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
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
func (device *GPSV2Bricklet) WriteFirmware(data [64]uint8) (status uint8, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, data)

	resultBytes, err := device.device.Get(uint8(FunctionWriteFirmware), buf.Bytes())
	if err != nil {
		return status, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return status, DeviceError(header.ErrorCode)
		}

		if header.Length != 9 {
			return status, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
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
func (device *GPSV2Bricklet) SetStatusLEDConfig(config StatusLEDConfig) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, config)

	resultBytes, err := device.device.Set(uint8(FunctionSetStatusLEDConfig), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
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
func (device *GPSV2Bricklet) GetStatusLEDConfig() (config StatusLEDConfig, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetStatusLEDConfig), buf.Bytes())
	if err != nil {
		return config, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return config, DeviceError(header.ErrorCode)
		}

		if header.Length != 9 {
			return config, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
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
func (device *GPSV2Bricklet) GetChipTemperature() (temperature int16, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetChipTemperature), buf.Bytes())
	if err != nil {
		return temperature, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return temperature, DeviceError(header.ErrorCode)
		}

		if header.Length != 10 {
			return temperature, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
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
func (device *GPSV2Bricklet) Reset() (err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Set(uint8(FunctionReset), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
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
func (device *GPSV2Bricklet) WriteUID(uid uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, uid)

	resultBytes, err := device.device.Set(uint8(FunctionWriteUID), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		if header.Length != 8 {
			return fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 8)
		}

		bytes.NewBuffer(resultBytes[8:])

	}

	return nil
}

// Returns the current UID as an integer. Encode as
// Base58 to get the usual string version.
func (device *GPSV2Bricklet) ReadUID() (uid uint32, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionReadUID), buf.Bytes())
	if err != nil {
		return uid, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return uid, DeviceError(header.ErrorCode)
		}

		if header.Length != 12 {
			return uid, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
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
func (device *GPSV2Bricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetIdentity), buf.Bytes())
	if err != nil {
		return uid, connectedUid, position, hardwareVersion, firmwareVersion, deviceIdentifier, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return uid, connectedUid, position, hardwareVersion, firmwareVersion, deviceIdentifier, DeviceError(header.ErrorCode)
		}

		if header.Length != 33 {
			return uid, connectedUid, position, hardwareVersion, firmwareVersion, deviceIdentifier, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 33)
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
