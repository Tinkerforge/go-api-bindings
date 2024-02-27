/* ***********************************************************
 * This file was automatically generated on 2024-02-27.      *
 *                                                           *
 * Go Bindings Version 2.0.15                                *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/

// 433MHz receiver for outdoor weather station.
//
//
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/OutdoorWeather_Bricklet_Go.html.
package outdoor_weather_bricklet

import (
	"bytes"
	"encoding/binary"
	"fmt"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionGetStationIdentifiersLowLevel   Function = 1
	FunctionGetSensorIdentifiersLowLevel    Function = 2
	FunctionGetStationData                  Function = 3
	FunctionGetSensorData                   Function = 4
	FunctionSetStationCallbackConfiguration Function = 5
	FunctionGetStationCallbackConfiguration Function = 6
	FunctionSetSensorCallbackConfiguration  Function = 7
	FunctionGetSensorCallbackConfiguration  Function = 8
	FunctionGetSPITFPErrorCount             Function = 234
	FunctionSetBootloaderMode               Function = 235
	FunctionGetBootloaderMode               Function = 236
	FunctionSetWriteFirmwarePointer         Function = 237
	FunctionWriteFirmware                   Function = 238
	FunctionSetStatusLEDConfig              Function = 239
	FunctionGetStatusLEDConfig              Function = 240
	FunctionGetChipTemperature              Function = 242
	FunctionReset                           Function = 243
	FunctionWriteUID                        Function = 248
	FunctionReadUID                         Function = 249
	FunctionGetIdentity                     Function = 255
	FunctionCallbackStationData             Function = 9
	FunctionCallbackSensorData              Function = 10
)

type WindDirection = uint8

const (
	WindDirectionN     WindDirection = 0
	WindDirectionNNE   WindDirection = 1
	WindDirectionNE    WindDirection = 2
	WindDirectionENE   WindDirection = 3
	WindDirectionE     WindDirection = 4
	WindDirectionESE   WindDirection = 5
	WindDirectionSE    WindDirection = 6
	WindDirectionSSE   WindDirection = 7
	WindDirectionS     WindDirection = 8
	WindDirectionSSW   WindDirection = 9
	WindDirectionSW    WindDirection = 10
	WindDirectionWSW   WindDirection = 11
	WindDirectionW     WindDirection = 12
	WindDirectionWNW   WindDirection = 13
	WindDirectionNW    WindDirection = 14
	WindDirectionNNW   WindDirection = 15
	WindDirectionError WindDirection = 255
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

type OutdoorWeatherBricklet struct {
	device Device
}

const DeviceIdentifier = 288
const DeviceDisplayName = "Outdoor Weather Bricklet"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (OutdoorWeatherBricklet, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{2, 0, 0}, uid, &internalIPCon, 2, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return OutdoorWeatherBricklet{}, err
	}
	dev.ResponseExpected[FunctionGetStationIdentifiersLowLevel] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetSensorIdentifiersLowLevel] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetStationData] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetSensorData] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetStationCallbackConfiguration] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionGetStationCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetSensorCallbackConfiguration] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionGetSensorCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue
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
	return OutdoorWeatherBricklet{dev}, nil
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
func (device *OutdoorWeatherBricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *OutdoorWeatherBricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *OutdoorWeatherBricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *OutdoorWeatherBricklet) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// Reports the station data every time a new data packet is received.
// See GetStationData for information about the data.
//
// For each station the callback will be triggered about every 45 seconds.
//
// Turn the callback on/off with SetStationCallbackConfiguration
// (by default it is turned off).
func (device *OutdoorWeatherBricklet) RegisterStationDataCallback(fn func(uint8, int16, uint8, uint32, uint32, uint32, WindDirection, bool)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 26 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var identifier uint8
		var temperature int16
		var humidity uint8
		var windSpeed uint32
		var gustSpeed uint32
		var rain uint32
		var windDirection WindDirection
		var batteryLow bool
		binary.Read(buf, binary.LittleEndian, &identifier)
		binary.Read(buf, binary.LittleEndian, &temperature)
		binary.Read(buf, binary.LittleEndian, &humidity)
		binary.Read(buf, binary.LittleEndian, &windSpeed)
		binary.Read(buf, binary.LittleEndian, &gustSpeed)
		binary.Read(buf, binary.LittleEndian, &rain)
		binary.Read(buf, binary.LittleEndian, &windDirection)
		binary.Read(buf, binary.LittleEndian, &batteryLow)
		fn(identifier, temperature, humidity, windSpeed, gustSpeed, rain, windDirection, batteryLow)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackStationData), wrapper)
}

// Remove a registered Station Data callback.
func (device *OutdoorWeatherBricklet) DeregisterStationDataCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackStationData), registrationId)
}

// Reports the sensor data every time a new data packet is received.
// See GetSensorData for information about the data.
//
// For each sensor the callback will be called about every 45 seconds.
//
// Turn the callback on/off with SetSensorCallbackConfiguration
// (by default it is turned off).
func (device *OutdoorWeatherBricklet) RegisterSensorDataCallback(fn func(uint8, int16, uint8)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 12 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var identifier uint8
		var temperature int16
		var humidity uint8
		binary.Read(buf, binary.LittleEndian, &identifier)
		binary.Read(buf, binary.LittleEndian, &temperature)
		binary.Read(buf, binary.LittleEndian, &humidity)
		fn(identifier, temperature, humidity)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackSensorData), wrapper)
}

// Remove a registered Sensor Data callback.
func (device *OutdoorWeatherBricklet) DeregisterSensorDataCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackSensorData), registrationId)
}

// Returns the identifiers (number between 0 and 255) of all `stations
// <https://www.tinkerforge.com/en/shop/accessories/sensors/outdoor-weather-station-ws-6147.html>`__
// that have been seen since the startup of the Bricklet.
//
// Each station gives itself a random identifier on first startup.
//
// Since firmware version 2.0.2 a station is removed from the list if no data was received for
// 12 hours.
func (device *OutdoorWeatherBricklet) GetStationIdentifiersLowLevel() (identifiersLength uint16, identifiersChunkOffset uint16, identifiersChunkData [60]uint8, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetStationIdentifiersLowLevel), buf.Bytes())
	if err != nil {
		return identifiersLength, identifiersChunkOffset, identifiersChunkData, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return identifiersLength, identifiersChunkOffset, identifiersChunkData, DeviceError(header.ErrorCode)
		}

		if header.Length != 72 {
			return identifiersLength, identifiersChunkOffset, identifiersChunkData, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 72)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &identifiersLength)
		binary.Read(resultBuf, binary.LittleEndian, &identifiersChunkOffset)
		copy(identifiersChunkData[:], ByteSliceToUint8Slice(resultBuf.Next(8*60/8)))

	}

	return identifiersLength, identifiersChunkOffset, identifiersChunkData, nil
}

// Returns the identifiers (number between 0 and 255) of all `stations
// <https://www.tinkerforge.com/en/shop/accessories/sensors/outdoor-weather-station-ws-6147.html>`__
// that have been seen since the startup of the Bricklet.
//
// Each station gives itself a random identifier on first startup.
//
// Since firmware version 2.0.2 a station is removed from the list if no data was received for
// 12 hours.
func (device *OutdoorWeatherBricklet) GetStationIdentifiers() (identifiers []uint8, err error) {
	buf, _, err := device.device.GetHighLevel(func() (LowLevelResult, error) {
		identifiersLength, identifiersChunkOffset, identifiersChunkData, err := device.GetStationIdentifiersLowLevel()

		if err != nil {
			return LowLevelResult{}, err
		}

		var lowLevelResults bytes.Buffer

		return LowLevelResult{
			uint64(identifiersLength),
			uint64(identifiersChunkOffset),
			Uint8SliceToByteSlice(identifiersChunkData[:]),
			lowLevelResults.Bytes()}, nil
	},
		0,
		8)
	if err != nil {
		return ByteSliceToUint8Slice(buf), err
	}

	return ByteSliceToUint8Slice(buf), nil
}

// Returns the identifiers (number between 0 and 255) of all `sensors
// <https://www.tinkerforge.com/en/shop/accessories/sensors/temperature-humidity-sensor-th-6148.html>`__
// that have been seen since the startup of the Bricklet.
//
// Each sensor gives itself a random identifier on first startup.
//
// Since firmware version 2.0.2 a sensor is removed from the list if no data was received for
// 12 hours.
func (device *OutdoorWeatherBricklet) GetSensorIdentifiersLowLevel() (identifiersLength uint16, identifiersChunkOffset uint16, identifiersChunkData [60]uint8, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetSensorIdentifiersLowLevel), buf.Bytes())
	if err != nil {
		return identifiersLength, identifiersChunkOffset, identifiersChunkData, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return identifiersLength, identifiersChunkOffset, identifiersChunkData, DeviceError(header.ErrorCode)
		}

		if header.Length != 72 {
			return identifiersLength, identifiersChunkOffset, identifiersChunkData, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 72)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &identifiersLength)
		binary.Read(resultBuf, binary.LittleEndian, &identifiersChunkOffset)
		copy(identifiersChunkData[:], ByteSliceToUint8Slice(resultBuf.Next(8*60/8)))

	}

	return identifiersLength, identifiersChunkOffset, identifiersChunkData, nil
}

// Returns the identifiers (number between 0 and 255) of all `sensors
// <https://www.tinkerforge.com/en/shop/accessories/sensors/temperature-humidity-sensor-th-6148.html>`__
// that have been seen since the startup of the Bricklet.
//
// Each sensor gives itself a random identifier on first startup.
//
// Since firmware version 2.0.2 a sensor is removed from the list if no data was received for
// 12 hours.
func (device *OutdoorWeatherBricklet) GetSensorIdentifiers() (identifiers []uint8, err error) {
	buf, _, err := device.device.GetHighLevel(func() (LowLevelResult, error) {
		identifiersLength, identifiersChunkOffset, identifiersChunkData, err := device.GetSensorIdentifiersLowLevel()

		if err != nil {
			return LowLevelResult{}, err
		}

		var lowLevelResults bytes.Buffer

		return LowLevelResult{
			uint64(identifiersLength),
			uint64(identifiersChunkOffset),
			Uint8SliceToByteSlice(identifiersChunkData[:]),
			lowLevelResults.Bytes()}, nil
	},
		1,
		8)
	if err != nil {
		return ByteSliceToUint8Slice(buf), err
	}

	return ByteSliceToUint8Slice(buf), nil
}

// Returns the last received data for a station with the given identifier.
// Call GetStationIdentifiers for a list of all available identifiers.
//
// The return values are:
//
// * Temperature,
// * Humidity,
// * Wind Speed,
// * Gust Speed,
// * Rain Fall (accumulated since station power-up),
// * Wind Direction,
// * Battery Low (true if battery is low) and
// * Last Change (seconds since the reception of this data).
//
// Associated constants:
//
//	* WindDirectionN
//	* WindDirectionNNE
//	* WindDirectionNE
//	* WindDirectionENE
//	* WindDirectionE
//	* WindDirectionESE
//	* WindDirectionSE
//	* WindDirectionSSE
//	* WindDirectionS
//	* WindDirectionSSW
//	* WindDirectionSW
//	* WindDirectionWSW
//	* WindDirectionW
//	* WindDirectionWNW
//	* WindDirectionNW
//	* WindDirectionNNW
//	* WindDirectionError
func (device *OutdoorWeatherBricklet) GetStationData(identifier uint8) (temperature int16, humidity uint8, windSpeed uint32, gustSpeed uint32, rain uint32, windDirection WindDirection, batteryLow bool, lastChange uint16, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, identifier)

	resultBytes, err := device.device.Get(uint8(FunctionGetStationData), buf.Bytes())
	if err != nil {
		return temperature, humidity, windSpeed, gustSpeed, rain, windDirection, batteryLow, lastChange, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return temperature, humidity, windSpeed, gustSpeed, rain, windDirection, batteryLow, lastChange, DeviceError(header.ErrorCode)
		}

		if header.Length != 27 {
			return temperature, humidity, windSpeed, gustSpeed, rain, windDirection, batteryLow, lastChange, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 27)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &temperature)
		binary.Read(resultBuf, binary.LittleEndian, &humidity)
		binary.Read(resultBuf, binary.LittleEndian, &windSpeed)
		binary.Read(resultBuf, binary.LittleEndian, &gustSpeed)
		binary.Read(resultBuf, binary.LittleEndian, &rain)
		binary.Read(resultBuf, binary.LittleEndian, &windDirection)
		binary.Read(resultBuf, binary.LittleEndian, &batteryLow)
		binary.Read(resultBuf, binary.LittleEndian, &lastChange)

	}

	return temperature, humidity, windSpeed, gustSpeed, rain, windDirection, batteryLow, lastChange, nil
}

// Returns the last measured data for a sensor with the given identifier.
// Call GetSensorIdentifiers for a list of all available identifiers.
//
// The return values are:
//
// * Temperature,
// * Humidity and
// * Last Change (seconds since the last reception of data).
func (device *OutdoorWeatherBricklet) GetSensorData(identifier uint8) (temperature int16, humidity uint8, lastChange uint16, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, identifier)

	resultBytes, err := device.device.Get(uint8(FunctionGetSensorData), buf.Bytes())
	if err != nil {
		return temperature, humidity, lastChange, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return temperature, humidity, lastChange, DeviceError(header.ErrorCode)
		}

		if header.Length != 13 {
			return temperature, humidity, lastChange, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 13)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &temperature)
		binary.Read(resultBuf, binary.LittleEndian, &humidity)
		binary.Read(resultBuf, binary.LittleEndian, &lastChange)

	}

	return temperature, humidity, lastChange, nil
}

// Turns callback for station data on or off.
func (device *OutdoorWeatherBricklet) SetStationCallbackConfiguration(enableCallback bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, enableCallback)

	resultBytes, err := device.device.Set(uint8(FunctionSetStationCallbackConfiguration), buf.Bytes())
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

// Returns the configuration as set by SetStationCallbackConfiguration.
func (device *OutdoorWeatherBricklet) GetStationCallbackConfiguration() (enableCallback bool, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetStationCallbackConfiguration), buf.Bytes())
	if err != nil {
		return enableCallback, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return enableCallback, DeviceError(header.ErrorCode)
		}

		if header.Length != 9 {
			return enableCallback, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &enableCallback)

	}

	return enableCallback, nil
}

// Turns callback for sensor data on or off.
func (device *OutdoorWeatherBricklet) SetSensorCallbackConfiguration(enableCallback bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, enableCallback)

	resultBytes, err := device.device.Set(uint8(FunctionSetSensorCallbackConfiguration), buf.Bytes())
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

// Returns the configuration as set by SetSensorCallbackConfiguration.
func (device *OutdoorWeatherBricklet) GetSensorCallbackConfiguration() (enableCallback bool, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetSensorCallbackConfiguration), buf.Bytes())
	if err != nil {
		return enableCallback, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return enableCallback, DeviceError(header.ErrorCode)
		}

		if header.Length != 9 {
			return enableCallback, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &enableCallback)

	}

	return enableCallback, nil
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
func (device *OutdoorWeatherBricklet) GetSPITFPErrorCount() (errorCountAckChecksum uint32, errorCountMessageChecksum uint32, errorCountFrame uint32, errorCountOverflow uint32, err error) {
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
func (device *OutdoorWeatherBricklet) SetBootloaderMode(mode BootloaderMode) (status BootloaderStatus, err error) {
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
func (device *OutdoorWeatherBricklet) GetBootloaderMode() (mode BootloaderMode, err error) {
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
func (device *OutdoorWeatherBricklet) SetWriteFirmwarePointer(pointer uint32) (err error) {
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
func (device *OutdoorWeatherBricklet) WriteFirmware(data [64]uint8) (status uint8, err error) {
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
func (device *OutdoorWeatherBricklet) SetStatusLEDConfig(config StatusLEDConfig) (err error) {
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
func (device *OutdoorWeatherBricklet) GetStatusLEDConfig() (config StatusLEDConfig, err error) {
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
func (device *OutdoorWeatherBricklet) GetChipTemperature() (temperature int16, err error) {
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
func (device *OutdoorWeatherBricklet) Reset() (err error) {
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
func (device *OutdoorWeatherBricklet) WriteUID(uid uint32) (err error) {
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
func (device *OutdoorWeatherBricklet) ReadUID() (uid uint32, err error) {
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
func (device *OutdoorWeatherBricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
