/* ***********************************************************
 * This file was automatically generated on 2020-11-02.      *
 *                                                           *
 * Go Bindings Version 2.0.9                                 *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


// Determine position, velocity and altitude using GPS‍.
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/GPS_Bricklet_Go.html.
package gps_bricklet

import (
	"encoding/binary"
	"bytes"
	"fmt"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionGetCoordinates Function = 1
	FunctionGetStatus Function = 2
	FunctionGetAltitude Function = 3
	FunctionGetMotion Function = 4
	FunctionGetDateTime Function = 5
	FunctionRestart Function = 6
	FunctionSetCoordinatesCallbackPeriod Function = 7
	FunctionGetCoordinatesCallbackPeriod Function = 8
	FunctionSetStatusCallbackPeriod Function = 9
	FunctionGetStatusCallbackPeriod Function = 10
	FunctionSetAltitudeCallbackPeriod Function = 11
	FunctionGetAltitudeCallbackPeriod Function = 12
	FunctionSetMotionCallbackPeriod Function = 13
	FunctionGetMotionCallbackPeriod Function = 14
	FunctionSetDateTimeCallbackPeriod Function = 15
	FunctionGetDateTimeCallbackPeriod Function = 16
	FunctionGetIdentity Function = 255
	FunctionCallbackCoordinates Function = 17
	FunctionCallbackStatus Function = 18
	FunctionCallbackAltitude Function = 19
	FunctionCallbackMotion Function = 20
	FunctionCallbackDateTime Function = 21
)

type Fix = uint8

const (
	FixNoFix Fix = 1
	Fix2DFix Fix = 2
	Fix3DFix Fix = 3
)

type RestartType = uint8

const (
	RestartTypeHotStart RestartType = 0
	RestartTypeWarmStart RestartType = 1
	RestartTypeColdStart RestartType = 2
	RestartTypeFactoryReset RestartType = 3
)

type GPSBricklet struct {
	device Device
}
const DeviceIdentifier = 222
const DeviceDisplayName = "GPS Bricklet"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (GPSBricklet, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{ 2,0,1 }, uid, &internalIPCon, 0, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return GPSBricklet{}, err
	}
	dev.ResponseExpected[FunctionGetCoordinates] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetStatus] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetAltitude] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetMotion] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetDateTime] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionRestart] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionSetCoordinatesCallbackPeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetCoordinatesCallbackPeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetStatusCallbackPeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetStatusCallbackPeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetAltitudeCallbackPeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetAltitudeCallbackPeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetMotionCallbackPeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetMotionCallbackPeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetDateTimeCallbackPeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetDateTimeCallbackPeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetIdentity] = ResponseExpectedFlagAlwaysTrue;
	return GPSBricklet{dev}, nil
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
func (device *GPSBricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *GPSBricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *GPSBricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *GPSBricklet) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// This callback is triggered periodically with the period that is set by
// SetCoordinatesCallbackPeriod. The parameters are the same
// as for GetCoordinates.
// 
// The RegisterCoordinatesCallback callback is only triggered if the coordinates changed
// since the last triggering and if there is currently a fix as indicated by
// GetStatus.
func (device *GPSBricklet) RegisterCoordinatesCallback(fn func(uint32, rune, uint32, rune, uint16, uint16, uint16, uint16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 26 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var latitude uint32
		var ns rune
		var longitude uint32
		var ew rune
		var pdop uint16
		var hdop uint16
		var vdop uint16
		var epe uint16
		binary.Read(buf, binary.LittleEndian, &latitude)
		ns = rune(buf.Next(1)[0])
		binary.Read(buf, binary.LittleEndian, &longitude)
		ew = rune(buf.Next(1)[0])
		binary.Read(buf, binary.LittleEndian, &pdop)
		binary.Read(buf, binary.LittleEndian, &hdop)
		binary.Read(buf, binary.LittleEndian, &vdop)
		binary.Read(buf, binary.LittleEndian, &epe)
		fn(latitude, ns, longitude, ew, pdop, hdop, vdop, epe)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackCoordinates), wrapper)
}

// Remove a registered Coordinates callback.
func (device *GPSBricklet) DeregisterCoordinatesCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackCoordinates), registrationId)
}


// This callback is triggered periodically with the period that is set by
// SetStatusCallbackPeriod. The parameters are the same
// as for GetStatus.
// 
// The RegisterStatusCallback callback is only triggered if the status changed since the
// last triggering.
func (device *GPSBricklet) RegisterStatusCallback(fn func(Fix, uint8, uint8)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 11 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var fix Fix
		var satellitesView uint8
		var satellitesUsed uint8
		binary.Read(buf, binary.LittleEndian, &fix)
		binary.Read(buf, binary.LittleEndian, &satellitesView)
		binary.Read(buf, binary.LittleEndian, &satellitesUsed)
		fn(fix, satellitesView, satellitesUsed)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackStatus), wrapper)
}

// Remove a registered Status callback.
func (device *GPSBricklet) DeregisterStatusCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackStatus), registrationId)
}


// This callback is triggered periodically with the period that is set by
// SetAltitudeCallbackPeriod. The parameters are the same
// as for GetAltitude.
// 
// The RegisterAltitudeCallback callback is only triggered if the altitude changed since
// the last triggering and if there is currently a fix as indicated by
// GetStatus.
func (device *GPSBricklet) RegisterAltitudeCallback(fn func(int32, int32)) uint64 {
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
func (device *GPSBricklet) DeregisterAltitudeCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackAltitude), registrationId)
}


// This callback is triggered periodically with the period that is set by
// SetMotionCallbackPeriod. The parameters are the same
// as for GetMotion.
// 
// The RegisterMotionCallback callback is only triggered if the motion changed since the
// last triggering and if there is currently a fix as indicated by
// GetStatus.
func (device *GPSBricklet) RegisterMotionCallback(fn func(uint32, uint32)) uint64 {
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
func (device *GPSBricklet) DeregisterMotionCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackMotion), registrationId)
}


// This callback is triggered periodically with the period that is set by
// SetDateTimeCallbackPeriod. The parameters are the same
// as for GetDateTime.
// 
// The RegisterDateTimeCallback callback is only triggered if the date or time changed
// since the last triggering.
func (device *GPSBricklet) RegisterDateTimeCallback(fn func(uint32, uint32)) uint64 {
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
func (device *GPSBricklet) DeregisterDateTimeCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackDateTime), registrationId)
}


// Returns the GPS coordinates. Latitude and longitude are given in the
// https://en.wikipedia.org/wiki/Dilution_of_precision_(GPS)
// for more information.
// 
// EPE is the Estimated Position Error. This is not the
// absolute maximum error, it is the error with a specific confidence. See
// https://www.nps.gov/gis/gps/WhatisEPE.html for more information.
// 
// This data is only valid if there is currently a fix as indicated by
// GetStatus.
func (device *GPSBricklet) GetCoordinates() (latitude uint32, ns rune, longitude uint32, ew rune, pdop uint16, hdop uint16, vdop uint16, epe uint16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetCoordinates), buf.Bytes())
	if err != nil {
		return latitude, ns, longitude, ew, pdop, hdop, vdop, epe, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 26 {
			return latitude, ns, longitude, ew, pdop, hdop, vdop, epe, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 26)
		}

		if header.ErrorCode != 0 {
			return latitude, ns, longitude, ew, pdop, hdop, vdop, epe, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &latitude)
		ns = rune(resultBuf.Next(1)[0])
		binary.Read(resultBuf, binary.LittleEndian, &longitude)
		ew = rune(resultBuf.Next(1)[0])
		binary.Read(resultBuf, binary.LittleEndian, &pdop)
		binary.Read(resultBuf, binary.LittleEndian, &hdop)
		binary.Read(resultBuf, binary.LittleEndian, &vdop)
		binary.Read(resultBuf, binary.LittleEndian, &epe)

	}

	return latitude, ns, longitude, ew, pdop, hdop, vdop, epe, nil
}

// Returns the current fix status, the number of satellites that are in view and
// the number of satellites that are currently used.
// 
// Possible fix status values can be:
// 
//  Value| Description
//  --- | --- 
//  1| No Fix| GetCoordinates| GetAltitude and GetMotion return invalid data
//  2| 2D Fix| only GetCoordinates and GetMotion return valid data
//  3| 3D Fix| GetCoordinates| GetAltitude and GetMotion return valid data
// 
// There is also a `blue LED <gps_bricklet_fix_led>` on the Bricklet that
// indicates the fix status.
//
// Associated constants:
//
//	* FixNoFix
//	* Fix2DFix
//	* Fix3DFix
func (device *GPSBricklet) GetStatus() (fix Fix, satellitesView uint8, satellitesUsed uint8, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetStatus), buf.Bytes())
	if err != nil {
		return fix, satellitesView, satellitesUsed, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 11 {
			return fix, satellitesView, satellitesUsed, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 11)
		}

		if header.ErrorCode != 0 {
			return fix, satellitesView, satellitesUsed, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &fix)
		binary.Read(resultBuf, binary.LittleEndian, &satellitesView)
		binary.Read(resultBuf, binary.LittleEndian, &satellitesUsed)

	}

	return fix, satellitesView, satellitesUsed, nil
}

// Returns the current altitude and corresponding geoidal separation.
// 
// This data is only valid if there is currently a fix as indicated by
// GetStatus.
func (device *GPSBricklet) GetAltitude() (altitude int32, geoidalSeparation int32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetAltitude), buf.Bytes())
	if err != nil {
		return altitude, geoidalSeparation, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 16 {
			return altitude, geoidalSeparation, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 16)
		}

		if header.ErrorCode != 0 {
			return altitude, geoidalSeparation, DeviceError(header.ErrorCode)
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
func (device *GPSBricklet) GetMotion() (course uint32, speed uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetMotion), buf.Bytes())
	if err != nil {
		return course, speed, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 16 {
			return course, speed, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 16)
		}

		if header.ErrorCode != 0 {
			return course, speed, DeviceError(header.ErrorCode)
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
func (device *GPSBricklet) GetDateTime() (date uint32, time uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetDateTime), buf.Bytes())
	if err != nil {
		return date, time, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 16 {
			return date, time, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 16)
		}

		if header.ErrorCode != 0 {
			return date, time, DeviceError(header.ErrorCode)
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
func (device *GPSBricklet) Restart(restartType RestartType) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, restartType);

	resultBytes, err := device.device.Set(uint8(FunctionRestart), buf.Bytes())
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

// Sets the period with which the RegisterCoordinatesCallback callback is triggered
// periodically. A value of 0 turns the callback off.
// 
// The RegisterCoordinatesCallback callback is only triggered if the coordinates changed
// since the last triggering.
func (device *GPSBricklet) SetCoordinatesCallbackPeriod(period uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);

	resultBytes, err := device.device.Set(uint8(FunctionSetCoordinatesCallbackPeriod), buf.Bytes())
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

// Returns the period as set by SetCoordinatesCallbackPeriod.
func (device *GPSBricklet) GetCoordinatesCallbackPeriod() (period uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetCoordinatesCallbackPeriod), buf.Bytes())
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

// Sets the period with which the RegisterStatusCallback callback is triggered
// periodically. A value of 0 turns the callback off.
// 
// The RegisterStatusCallback callback is only triggered if the status changed since the
// last triggering.
func (device *GPSBricklet) SetStatusCallbackPeriod(period uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);

	resultBytes, err := device.device.Set(uint8(FunctionSetStatusCallbackPeriod), buf.Bytes())
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

// Returns the period as set by SetStatusCallbackPeriod.
func (device *GPSBricklet) GetStatusCallbackPeriod() (period uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetStatusCallbackPeriod), buf.Bytes())
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

// Sets the period with which the RegisterAltitudeCallback callback is triggered
// periodically. A value of 0 turns the callback off.
// 
// The RegisterAltitudeCallback callback is only triggered if the altitude changed since
// the last triggering.
func (device *GPSBricklet) SetAltitudeCallbackPeriod(period uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);

	resultBytes, err := device.device.Set(uint8(FunctionSetAltitudeCallbackPeriod), buf.Bytes())
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

// Returns the period as set by SetAltitudeCallbackPeriod.
func (device *GPSBricklet) GetAltitudeCallbackPeriod() (period uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetAltitudeCallbackPeriod), buf.Bytes())
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

// Sets the period with which the RegisterMotionCallback callback is triggered
// periodically. A value of 0 turns the callback off.
// 
// The RegisterMotionCallback callback is only triggered if the motion changed since the
// last triggering.
func (device *GPSBricklet) SetMotionCallbackPeriod(period uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);

	resultBytes, err := device.device.Set(uint8(FunctionSetMotionCallbackPeriod), buf.Bytes())
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

// Returns the period as set by SetMotionCallbackPeriod.
func (device *GPSBricklet) GetMotionCallbackPeriod() (period uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetMotionCallbackPeriod), buf.Bytes())
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

// Sets the period with which the RegisterDateTimeCallback callback is triggered
// periodically. A value of 0 turns the callback off.
// 
// The RegisterDateTimeCallback callback is only triggered if the date or time changed
// since the last triggering.
func (device *GPSBricklet) SetDateTimeCallbackPeriod(period uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);

	resultBytes, err := device.device.Set(uint8(FunctionSetDateTimeCallbackPeriod), buf.Bytes())
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

// Returns the period as set by SetDateTimeCallbackPeriod.
func (device *GPSBricklet) GetDateTimeCallbackPeriod() (period uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetDateTimeCallbackPeriod), buf.Bytes())
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
func (device *GPSBricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
