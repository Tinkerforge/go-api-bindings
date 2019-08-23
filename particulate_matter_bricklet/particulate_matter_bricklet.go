/* ***********************************************************
 * This file was automatically generated on 2019-08-23.      *
 *                                                           *
 * Go Bindings Version 2.0.4                                 *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


// Measures Particulate Matter concentration (PM1.0, PM2.5 and PM10).
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/ParticulateMatter_Bricklet_Go.html.
package particulate_matter_bricklet

import (
	"encoding/binary"
	"bytes"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionGetPMConcentration Function = 1
	FunctionGetPMCount Function = 2
	FunctionSetEnable Function = 3
	FunctionGetEnable Function = 4
	FunctionGetSensorInfo Function = 5
	FunctionSetPMConcentrationCallbackConfiguration Function = 6
	FunctionGetPMConcentrationCallbackConfiguration Function = 7
	FunctionSetPMCountCallbackConfiguration Function = 8
	FunctionGetPMCountCallbackConfiguration Function = 9
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
	FunctionCallbackPMConcentration Function = 10
	FunctionCallbackPMCount Function = 11
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

type ParticulateMatterBricklet struct {
	device Device
}
const DeviceIdentifier = 2110
const DeviceDisplayName = "Particulate Matter Bricklet"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (ParticulateMatterBricklet, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{ 2,0,0 }, uid, &internalIPCon, 0)
	if err != nil {
		return ParticulateMatterBricklet{}, err
	}
	dev.ResponseExpected[FunctionGetPMConcentration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetPMCount] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetEnable] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetEnable] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetSensorInfo] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetPMConcentrationCallbackConfiguration] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetPMConcentrationCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetPMCountCallbackConfiguration] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetPMCountCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue;
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
	return ParticulateMatterBricklet{dev}, nil
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
func (device *ParticulateMatterBricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *ParticulateMatterBricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *ParticulateMatterBricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *ParticulateMatterBricklet) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// This callback is triggered periodically according to the configuration set by
// SetPMConcentrationCallbackConfiguration.
// 
// The parameters are the same as GetPMConcentration.
func (device *ParticulateMatterBricklet) RegisterPMConcentrationCallback(fn func(uint16, uint16, uint16)) uint64 {
	wrapper := func(byteSlice []byte) {
		buf := bytes.NewBuffer(byteSlice[8:])
		var pm10 uint16
		var pm25 uint16
		var pm100 uint16
		binary.Read(buf, binary.LittleEndian, &pm10)
		binary.Read(buf, binary.LittleEndian, &pm25)
		binary.Read(buf, binary.LittleEndian, &pm100)
		fn(pm10, pm25, pm100)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackPMConcentration), wrapper)
}

// Remove a registered PM Concentration callback.
func (device *ParticulateMatterBricklet) DeregisterPMConcentrationCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackPMConcentration), registrationId)
}


// This callback is triggered periodically according to the configuration set by
// SetPMCountCallbackConfiguration.
// 
// The parameters are the same as GetPMCount.
func (device *ParticulateMatterBricklet) RegisterPMCountCallback(fn func(uint16, uint16, uint16, uint16, uint16, uint16)) uint64 {
	wrapper := func(byteSlice []byte) {
		buf := bytes.NewBuffer(byteSlice[8:])
		var greater03um uint16
		var greater05um uint16
		var greater10um uint16
		var greater25um uint16
		var greater50um uint16
		var greater100um uint16
		binary.Read(buf, binary.LittleEndian, &greater03um)
		binary.Read(buf, binary.LittleEndian, &greater05um)
		binary.Read(buf, binary.LittleEndian, &greater10um)
		binary.Read(buf, binary.LittleEndian, &greater25um)
		binary.Read(buf, binary.LittleEndian, &greater50um)
		binary.Read(buf, binary.LittleEndian, &greater100um)
		fn(greater03um, greater05um, greater10um, greater25um, greater50um, greater100um)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackPMCount), wrapper)
}

// Remove a registered PM Count callback.
func (device *ParticulateMatterBricklet) DeregisterPMCountCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackPMCount), registrationId)
}


// Returns the particulate matter concentration in µg/m³, broken down as:
// 
// * PM\ :sub:`1.0`\ ,
// * PM\ :sub:`2.5`\  and
// * PM\ :sub:`10.0`\ .
// 
// If the sensor is disabled (see SetEnable) then the last known good
// values from the sensor are returned.
func (device *ParticulateMatterBricklet) GetPMConcentration() (pm10 uint16, pm25 uint16, pm100 uint16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetPMConcentration), buf.Bytes())
	if err != nil {
		return pm10, pm25, pm100, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return pm10, pm25, pm100, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &pm10)
	binary.Read(resultBuf, binary.LittleEndian, &pm25)
	binary.Read(resultBuf, binary.LittleEndian, &pm100)

	}

	return pm10, pm25, pm100, nil
}

// Returns the number of particulates in 100 ml of air, broken down by their
// diameter:
// 
// * greater 0.3µm,
// * greater 0.5µm,
// * greater 1.0µm,
// * greater 2.5µm,
// * greater 5.0µm and
// * greater 10.0µm.
// 
// If the sensor is disabled (see SetEnable) then the last known good
// value from the sensor is returned.
func (device *ParticulateMatterBricklet) GetPMCount() (greater03um uint16, greater05um uint16, greater10um uint16, greater25um uint16, greater50um uint16, greater100um uint16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetPMCount), buf.Bytes())
	if err != nil {
		return greater03um, greater05um, greater10um, greater25um, greater50um, greater100um, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return greater03um, greater05um, greater10um, greater25um, greater50um, greater100um, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &greater03um)
	binary.Read(resultBuf, binary.LittleEndian, &greater05um)
	binary.Read(resultBuf, binary.LittleEndian, &greater10um)
	binary.Read(resultBuf, binary.LittleEndian, &greater25um)
	binary.Read(resultBuf, binary.LittleEndian, &greater50um)
	binary.Read(resultBuf, binary.LittleEndian, &greater100um)

	}

	return greater03um, greater05um, greater10um, greater25um, greater50um, greater100um, nil
}

// Enables/Disables the fan and the laser diode of the sensors. The sensor is
// enabled by default.
// 
// The sensor takes about 30 seconds after it is enabled to settle and produce stable
// values.
// 
// The laser diode has a lifetime of about 8000 hours. If you want to measure in
// an interval with a long idle time (e.g. hourly) you should turn the
// laser diode off between the measurements.
func (device *ParticulateMatterBricklet) SetEnable(enable bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, enable);

	resultBytes, err := device.device.Set(uint8(FunctionSetEnable), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])
		
	}

	return nil
}

// Returns the state of the sensor as set by SetEnable.
func (device *ParticulateMatterBricklet) GetEnable() (enable bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetEnable), buf.Bytes())
	if err != nil {
		return enable, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return enable, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &enable)

	}

	return enable, nil
}

// Returns information about the sensor:
// 
// * the sensor version number,
// * the last error code reported by the sensor (0 means no error) and
// * the number of framing and checksum errors that occurred in the communication
//   with the sensor.
func (device *ParticulateMatterBricklet) GetSensorInfo() (sensorVersion uint8, lastErrorCode uint8, framingErrorCount uint8, checksumErrorCount uint8, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetSensorInfo), buf.Bytes())
	if err != nil {
		return sensorVersion, lastErrorCode, framingErrorCount, checksumErrorCount, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return sensorVersion, lastErrorCode, framingErrorCount, checksumErrorCount, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &sensorVersion)
	binary.Read(resultBuf, binary.LittleEndian, &lastErrorCode)
	binary.Read(resultBuf, binary.LittleEndian, &framingErrorCount)
	binary.Read(resultBuf, binary.LittleEndian, &checksumErrorCount)

	}

	return sensorVersion, lastErrorCode, framingErrorCount, checksumErrorCount, nil
}

// The period in ms is the period with which the RegisterPMConcentrationCallback
// callback is triggered periodically. A value of 0 turns the callback off.
// 
// If the `value has to change`-parameter is set to true, the callback is only
// triggered after the value has changed. If the value didn't change within the
// period, the callback is triggered immediately on change.
// 
// If it is set to false, the callback is continuously triggered with the period,
// independent of the value.
// 
// The default value is (0, false).
func (device *ParticulateMatterBricklet) SetPMConcentrationCallbackConfiguration(period uint32, valueHasToChange bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);
	binary.Write(&buf, binary.LittleEndian, valueHasToChange);

	resultBytes, err := device.device.Set(uint8(FunctionSetPMConcentrationCallbackConfiguration), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])
		
	}

	return nil
}

// Returns the callback configuration as set by
// SetPMConcentrationCallbackConfiguration.
func (device *ParticulateMatterBricklet) GetPMConcentrationCallbackConfiguration() (period uint32, valueHasToChange bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetPMConcentrationCallbackConfiguration), buf.Bytes())
	if err != nil {
		return period, valueHasToChange, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return period, valueHasToChange, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &period)
	binary.Read(resultBuf, binary.LittleEndian, &valueHasToChange)

	}

	return period, valueHasToChange, nil
}

// The period in ms is the period with which the RegisterPMCountCallback callback
// is triggered periodically. A value of 0 turns the callback off.
// 
// If the `value has to change`-parameter is set to true, the callback is only
// triggered after the value has changed. If the value didn't change within the
// period, the callback is triggered immediately on change.
// 
// If it is set to false, the callback is continuously triggered with the period,
// independent of the value.
// 
// The default value is (0, false).
func (device *ParticulateMatterBricklet) SetPMCountCallbackConfiguration(period uint32, valueHasToChange bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);
	binary.Write(&buf, binary.LittleEndian, valueHasToChange);

	resultBytes, err := device.device.Set(uint8(FunctionSetPMCountCallbackConfiguration), buf.Bytes())
	if err != nil {
		return err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])
		
	}

	return nil
}

// Returns the callback configuration as set by
// SetPMCountCallbackConfiguration.
func (device *ParticulateMatterBricklet) GetPMCountCallbackConfiguration() (period uint32, valueHasToChange bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetPMCountCallbackConfiguration), buf.Bytes())
	if err != nil {
		return period, valueHasToChange, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)
		if header.ErrorCode != 0 {
			return period, valueHasToChange, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &period)
	binary.Read(resultBuf, binary.LittleEndian, &valueHasToChange)

	}

	return period, valueHasToChange, nil
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
func (device *ParticulateMatterBricklet) GetSPITFPErrorCount() (errorCountAckChecksum uint32, errorCountMessageChecksum uint32, errorCountFrame uint32, errorCountOverflow uint32, err error) {
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
func (device *ParticulateMatterBricklet) SetBootloaderMode(mode BootloaderMode) (status BootloaderStatus, err error) {
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
func (device *ParticulateMatterBricklet) GetBootloaderMode() (mode BootloaderMode, err error) {
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
func (device *ParticulateMatterBricklet) SetWriteFirmwarePointer(pointer uint32) (err error) {
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
func (device *ParticulateMatterBricklet) WriteFirmware(data [64]uint8) (status uint8, err error) {
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
func (device *ParticulateMatterBricklet) SetStatusLEDConfig(config StatusLEDConfig) (err error) {
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
func (device *ParticulateMatterBricklet) GetStatusLEDConfig() (config StatusLEDConfig, err error) {
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
func (device *ParticulateMatterBricklet) GetChipTemperature() (temperature int16, err error) {
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
func (device *ParticulateMatterBricklet) Reset() (err error) {
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

		bytes.NewBuffer(resultBytes[8:])
		
	}

	return nil
}

// Writes a new UID into flash. If you want to set a new UID
// you have to decode the Base58 encoded UID string into an
// integer first.
// 
// We recommend that you use Brick Viewer to change the UID.
func (device *ParticulateMatterBricklet) WriteUID(uid uint32) (err error) {
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
			return DeviceError(header.ErrorCode)
		}

		bytes.NewBuffer(resultBytes[8:])
		
	}

	return nil
}

// Returns the current UID as an integer. Encode as
// Base58 to get the usual string version.
func (device *ParticulateMatterBricklet) ReadUID() (uid uint32, err error) {
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
func (device *ParticulateMatterBricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
