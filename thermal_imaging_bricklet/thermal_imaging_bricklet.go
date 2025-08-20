/* ***********************************************************
 * This file was automatically generated on 2025-08-20.      *
 *                                                           *
 * Go Bindings Version 2.0.16                                *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/

// 80x60 pixel thermal imaging camera.
//
//
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/ThermalImaging_Bricklet_Go.html.
package thermal_imaging_bricklet

import (
	"bytes"
	"encoding/binary"
	"fmt"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionGetHighContrastImageLowLevel      Function = 1
	FunctionGetTemperatureImageLowLevel       Function = 2
	FunctionGetStatistics                     Function = 3
	FunctionSetResolution                     Function = 4
	FunctionGetResolution                     Function = 5
	FunctionSetSpotmeterConfig                Function = 6
	FunctionGetSpotmeterConfig                Function = 7
	FunctionSetHighContrastConfig             Function = 8
	FunctionGetHighContrastConfig             Function = 9
	FunctionSetImageTransferConfig            Function = 10
	FunctionGetImageTransferConfig            Function = 11
	FunctionSetFluxLinearParameters           Function = 14
	FunctionGetFluxLinearParameters           Function = 15
	FunctionSetFFCShutterMode                 Function = 16
	FunctionGetFFCShutterMode                 Function = 17
	FunctionRunFFCNormalization               Function = 18
	FunctionGetSPITFPErrorCount               Function = 234
	FunctionSetBootloaderMode                 Function = 235
	FunctionGetBootloaderMode                 Function = 236
	FunctionSetWriteFirmwarePointer           Function = 237
	FunctionWriteFirmware                     Function = 238
	FunctionSetStatusLEDConfig                Function = 239
	FunctionGetStatusLEDConfig                Function = 240
	FunctionGetChipTemperature                Function = 242
	FunctionReset                             Function = 243
	FunctionWriteUID                          Function = 248
	FunctionReadUID                           Function = 249
	FunctionGetIdentity                       Function = 255
	FunctionCallbackHighContrastImageLowLevel Function = 12
	FunctionCallbackTemperatureImageLowLevel  Function = 13
)

type Resolution = uint8

const (
	Resolution0To6553Kelvin Resolution = 0
	Resolution0To655Kelvin  Resolution = 1
)

type FFCStatus = uint8

const (
	FFCStatusNeverCommanded FFCStatus = 0
	FFCStatusImminent       FFCStatus = 1
	FFCStatusInProgress     FFCStatus = 2
	FFCStatusComplete       FFCStatus = 3
)

type ImageTransfer = uint8

const (
	ImageTransferManualHighContrastImage   ImageTransfer = 0
	ImageTransferManualTemperatureImage    ImageTransfer = 1
	ImageTransferCallbackHighContrastImage ImageTransfer = 2
	ImageTransferCallbackTemperatureImage  ImageTransfer = 3
)

type ShutterMode = uint8

const (
	ShutterModeManual   ShutterMode = 0
	ShutterModeAuto     ShutterMode = 1
	ShutterModeExternal ShutterMode = 2
)

type ShutterLockout = uint8

const (
	ShutterLockoutInactive ShutterLockout = 0
	ShutterLockoutHigh     ShutterLockout = 1
	ShutterLockoutLow      ShutterLockout = 2
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

type ThermalImagingBricklet struct {
	device Device
}

const DeviceIdentifier = 278
const DeviceDisplayName = "Thermal Imaging Bricklet"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (ThermalImagingBricklet, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{2, 0, 2}, uid, &internalIPCon, 4, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return ThermalImagingBricklet{}, err
	}
	dev.ResponseExpected[FunctionGetHighContrastImageLowLevel] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetTemperatureImageLowLevel] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetStatistics] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetResolution] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetResolution] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetSpotmeterConfig] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetSpotmeterConfig] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetHighContrastConfig] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetHighContrastConfig] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetImageTransferConfig] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionGetImageTransferConfig] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetFluxLinearParameters] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetFluxLinearParameters] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetFFCShutterMode] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetFFCShutterMode] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionRunFFCNormalization] = ResponseExpectedFlagFalse
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
	return ThermalImagingBricklet{dev}, nil
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
func (device *ThermalImagingBricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *ThermalImagingBricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *ThermalImagingBricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *ThermalImagingBricklet) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// See RegisterHighContrastImageCallback
func (device *ThermalImagingBricklet) RegisterHighContrastImageLowLevelCallback(fn func(uint16, [62]uint8)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 72 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var imageChunkOffset uint16
		var imageChunkData [62]uint8
		binary.Read(buf, binary.LittleEndian, &imageChunkOffset)
		copy(imageChunkData[:], ByteSliceToUint8Slice(buf.Next(8*62/8)))
		fn(imageChunkOffset, imageChunkData)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackHighContrastImageLowLevel), wrapper)
}

// Remove a registered High Contrast Image Low Level callback.
func (device *ThermalImagingBricklet) DeregisterHighContrastImageLowLevelCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackHighContrastImageLowLevel), registrationId)
}

// This callback is triggered with every new high contrast image if the transfer image
// config is configured for high contrast callback (see SetImageTransferConfig).
//
// The data is organized as a 8-bit value 80x60 pixel matrix linearized in
// a one-dimensional array. The data is arranged line by line from top left to
// bottom right.
//
// Each 8-bit value represents one gray-scale image pixel that can directly be
// shown to a user on a display.
//
// Note
//  If reconstructing the value fails, the callback is triggered with nil for image.
func (device *ThermalImagingBricklet) RegisterHighContrastImageCallback(fn func([]uint8)) uint64 {
	buf := make([]uint8, 0)
	wrapper := func(imageChunkOffset uint16, imageChunkData [62]uint8) {
		if int(imageChunkOffset) != len(buf) {
			if len(buf) != 0 {
				buf = nil
				fn(buf)
				buf = make([]uint8, 0)
			}
			if imageChunkOffset != 0 {
				return
			}
		}
		toRead := MinU(uint64(4800-imageChunkOffset), uint64(len(imageChunkData[:])))
		buf = append(buf, imageChunkData[:toRead]...)
		if len(buf) >= int(4800) {
			fn(buf)
			buf = make([]uint8, 0)
		}
	}
	return device.RegisterHighContrastImageLowLevelCallback(wrapper)
}

// Remove a registered High Contrast Image Low Level callback.
func (device *ThermalImagingBricklet) DeregisterHighContrastImageCallback(registrationId uint64) {
	device.DeregisterHighContrastImageLowLevelCallback(registrationId)
}

// See RegisterTemperatureImageCallback
func (device *ThermalImagingBricklet) RegisterTemperatureImageLowLevelCallback(fn func(uint16, [31]uint16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 72 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var imageChunkOffset uint16
		var imageChunkData [31]uint16
		binary.Read(buf, binary.LittleEndian, &imageChunkOffset)
		copy(imageChunkData[:], ByteSliceToUint16Slice(buf.Next(16*31/8)))
		fn(imageChunkOffset, imageChunkData)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackTemperatureImageLowLevel), wrapper)
}

// Remove a registered Temperature Image Low Level callback.
func (device *ThermalImagingBricklet) DeregisterTemperatureImageLowLevelCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackTemperatureImageLowLevel), registrationId)
}

// This callback is triggered with every new temperature image if the transfer image
// config is configured for temperature callback (see SetImageTransferConfig).
//
// The data is organized as a 16-bit value 80x60 pixel matrix linearized in
// a one-dimensional array. The data is arranged line by line from top left to
// bottom right.
//
// Each 16-bit value represents one temperature measurement in either
// Kelvin/10 or Kelvin/100 (depending on the resolution set with SetResolution).
//
// Note
//  If reconstructing the value fails, the callback is triggered with nil for image.
func (device *ThermalImagingBricklet) RegisterTemperatureImageCallback(fn func([]uint16)) uint64 {
	buf := make([]uint16, 0)
	wrapper := func(imageChunkOffset uint16, imageChunkData [31]uint16) {
		if int(imageChunkOffset) != len(buf) {
			if len(buf) != 0 {
				buf = nil
				fn(buf)
				buf = make([]uint16, 0)
			}
			if imageChunkOffset != 0 {
				return
			}
		}
		toRead := MinU(uint64(4800-imageChunkOffset), uint64(len(imageChunkData[:])))
		buf = append(buf, imageChunkData[:toRead]...)
		if len(buf) >= int(4800) {
			fn(buf)
			buf = make([]uint16, 0)
		}
	}
	return device.RegisterTemperatureImageLowLevelCallback(wrapper)
}

// Remove a registered Temperature Image Low Level callback.
func (device *ThermalImagingBricklet) DeregisterTemperatureImageCallback(registrationId uint64) {
	device.DeregisterTemperatureImageLowLevelCallback(registrationId)
}

// Returns the current high contrast image. See https://www.tinkerforge.com/en/doc/Hardware/Bricklets/Thermal_Imaging.html#high-contrast-image-vs-temperature-image
// for the difference between
// High Contrast and Temperature Image. If you don't know what to use
// the High Contrast Image is probably right for you.
//
// The data is organized as a 8-bit value 80x60 pixel matrix linearized in
// a one-dimensional array. The data is arranged line by line from top left to
// bottom right.
//
// Each 8-bit value represents one gray-scale image pixel that can directly be
// shown to a user on a display.
//
// Before you can use this function you have to enable it with
// SetImageTransferConfig.
func (device *ThermalImagingBricklet) GetHighContrastImageLowLevel() (imageChunkOffset uint16, imageChunkData [62]uint8, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetHighContrastImageLowLevel), buf.Bytes())
	if err != nil {
		return imageChunkOffset, imageChunkData, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return imageChunkOffset, imageChunkData, DeviceError(header.ErrorCode)
		}

		if header.Length != 72 {
			return imageChunkOffset, imageChunkData, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 72)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &imageChunkOffset)
		copy(imageChunkData[:], ByteSliceToUint8Slice(resultBuf.Next(8*62/8)))

	}

	return imageChunkOffset, imageChunkData, nil
}

// Returns the current high contrast image. See https://www.tinkerforge.com/en/doc/Hardware/Bricklets/Thermal_Imaging.html#high-contrast-image-vs-temperature-image
// for the difference between
// High Contrast and Temperature Image. If you don't know what to use
// the High Contrast Image is probably right for you.
//
// The data is organized as a 8-bit value 80x60 pixel matrix linearized in
// a one-dimensional array. The data is arranged line by line from top left to
// bottom right.
//
// Each 8-bit value represents one gray-scale image pixel that can directly be
// shown to a user on a display.
//
// Before you can use this function you have to enable it with
// SetImageTransferConfig.
func (device *ThermalImagingBricklet) GetHighContrastImage() (image []uint8, err error) {
	buf, _, err := device.device.GetHighLevel(func() (LowLevelResult, error) {
		imageChunkOffset, imageChunkData, err := device.GetHighContrastImageLowLevel()

		if err != nil {
			return LowLevelResult{}, err
		}

		var lowLevelResults bytes.Buffer

		return LowLevelResult{
			uint64(4800),
			uint64(imageChunkOffset),
			Uint8SliceToByteSlice(imageChunkData[:]),
			lowLevelResults.Bytes()}, nil
	},
		0,
		8)
	if err != nil {
		return ByteSliceToUint8Slice(buf), err
	}

	return ByteSliceToUint8Slice(buf), nil
}

// Returns the current temperature image. See https://www.tinkerforge.com/en/doc/Hardware/Bricklets/Thermal_Imaging.html#high-contrast-image-vs-temperature-image
// for the difference between High Contrast and Temperature Image.
// If you don't know what to use the High Contrast Image is probably right for you.
//
// The data is organized as a 16-bit value 80x60 pixel matrix linearized in
// a one-dimensional array. The data is arranged line by line from top left to
// bottom right.
//
// Each 16-bit value represents one temperature measurement in either
// Kelvin/10 or Kelvin/100 (depending on the resolution set with SetResolution).
//
// Before you can use this function you have to enable it with
// SetImageTransferConfig.
func (device *ThermalImagingBricklet) GetTemperatureImageLowLevel() (imageChunkOffset uint16, imageChunkData [31]uint16, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetTemperatureImageLowLevel), buf.Bytes())
	if err != nil {
		return imageChunkOffset, imageChunkData, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return imageChunkOffset, imageChunkData, DeviceError(header.ErrorCode)
		}

		if header.Length != 72 {
			return imageChunkOffset, imageChunkData, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 72)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &imageChunkOffset)
		copy(imageChunkData[:], ByteSliceToUint16Slice(resultBuf.Next(16*31/8)))

	}

	return imageChunkOffset, imageChunkData, nil
}

// Returns the current temperature image. See https://www.tinkerforge.com/en/doc/Hardware/Bricklets/Thermal_Imaging.html#high-contrast-image-vs-temperature-image
// for the difference between High Contrast and Temperature Image.
// If you don't know what to use the High Contrast Image is probably right for you.
//
// The data is organized as a 16-bit value 80x60 pixel matrix linearized in
// a one-dimensional array. The data is arranged line by line from top left to
// bottom right.
//
// Each 16-bit value represents one temperature measurement in either
// Kelvin/10 or Kelvin/100 (depending on the resolution set with SetResolution).
//
// Before you can use this function you have to enable it with
// SetImageTransferConfig.
func (device *ThermalImagingBricklet) GetTemperatureImage() (image []uint16, err error) {
	buf, _, err := device.device.GetHighLevel(func() (LowLevelResult, error) {
		imageChunkOffset, imageChunkData, err := device.GetTemperatureImageLowLevel()

		if err != nil {
			return LowLevelResult{}, err
		}

		var lowLevelResults bytes.Buffer

		return LowLevelResult{
			uint64(4800),
			uint64(imageChunkOffset),
			Uint16SliceToByteSlice(imageChunkData[:]),
			lowLevelResults.Bytes()}, nil
	},
		1,
		16)
	if err != nil {
		return ByteSliceToUint16Slice(buf), err
	}

	return ByteSliceToUint16Slice(buf), nil
}

// Returns the spotmeter statistics, various temperatures, current resolution and status bits.
//
// The spotmeter statistics are:
//
// * Index 0: Mean Temperature.
// * Index 1: Maximum Temperature.
// * Index 2: Minimum Temperature.
// * Index 3: Pixel Count of spotmeter region of interest.
//
// The temperatures are:
//
// * Index 0: Focal Plain Array temperature.
// * Index 1: Focal Plain Array temperature at last FFC (Flat Field Correction).
// * Index 2: Housing temperature.
// * Index 3: Housing temperature at last FFC.
//
// The resolution is either `0 to 6553 Kelvin` or `0 to 655 Kelvin`. If the resolution is the former,
// the temperatures are in Kelvin/10, if it is the latter the temperatures are in Kelvin/100.
//
// FFC (Flat Field Correction) Status:
//
// * FFC Never Commanded: Only seen on startup before first FFC.
// * FFC Imminent: This state is entered 2 seconds prior to initiating FFC.
// * FFC In Progress: Flat field correction is started (shutter moves in front of lens and back). Takes about 1 second.
// * FFC Complete: Shutter is in waiting position again, FFC done.
//
// Temperature warning bits:
//
// * Index 0: Shutter lockout (if true shutter is locked out because temperature is outside -10°C to +65°C)
// * Index 1: Overtemperature shut down imminent (goes true 10 seconds before shutdown)
//
// Associated constants:
//
//	* Resolution0To6553Kelvin
//	* Resolution0To655Kelvin
//	* FFCStatusNeverCommanded
//	* FFCStatusImminent
//	* FFCStatusInProgress
//	* FFCStatusComplete
func (device *ThermalImagingBricklet) GetStatistics() (spotmeterStatistics [4]uint16, temperatures [4]uint16, resolution Resolution, ffcStatus FFCStatus, temperatureWarning [2]bool, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetStatistics), buf.Bytes())
	if err != nil {
		return spotmeterStatistics, temperatures, resolution, ffcStatus, temperatureWarning, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return spotmeterStatistics, temperatures, resolution, ffcStatus, temperatureWarning, DeviceError(header.ErrorCode)
		}

		if header.Length != 27 {
			return spotmeterStatistics, temperatures, resolution, ffcStatus, temperatureWarning, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 27)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &spotmeterStatistics)
		binary.Read(resultBuf, binary.LittleEndian, &temperatures)
		binary.Read(resultBuf, binary.LittleEndian, &resolution)
		binary.Read(resultBuf, binary.LittleEndian, &ffcStatus)
		copy(temperatureWarning[:], ByteSliceToBoolSlice(resultBuf.Next(1)))

	}

	return spotmeterStatistics, temperatures, resolution, ffcStatus, temperatureWarning, nil
}

// Sets the resolution. The Thermal Imaging Bricklet can either measure
//
// * from 0 to 6553 Kelvin (-273.15°C to +6279.85°C) with 0.1°C resolution or
// * from 0 to 655 Kelvin (-273.15°C to +381.85°C) with 0.01°C resolution.
//
// The accuracy is specified for -10°C to 450°C in the
// first range and -10°C and 140°C in the second range.
//
// Associated constants:
//
//	* Resolution0To6553Kelvin
//	* Resolution0To655Kelvin
func (device *ThermalImagingBricklet) SetResolution(resolution Resolution) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, resolution)

	resultBytes, err := device.device.Set(uint8(FunctionSetResolution), buf.Bytes())
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

// Returns the resolution as set by SetResolution.
//
// Associated constants:
//
//	* Resolution0To6553Kelvin
//	* Resolution0To655Kelvin
func (device *ThermalImagingBricklet) GetResolution() (resolution Resolution, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetResolution), buf.Bytes())
	if err != nil {
		return resolution, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return resolution, DeviceError(header.ErrorCode)
		}

		if header.Length != 9 {
			return resolution, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &resolution)

	}

	return resolution, nil
}

// Sets the spotmeter region of interest. The 4 values are
//
// * Index 0: Column start (has to be smaller than column end).
// * Index 1: Row start (has to be smaller than row end).
// * Index 2: Column end (has to be smaller than 80).
// * Index 3: Row end (has to be smaller than 60).
//
// The spotmeter statistics can be read out with GetStatistics.
func (device *ThermalImagingBricklet) SetSpotmeterConfig(regionOfInterest [4]uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, regionOfInterest)

	resultBytes, err := device.device.Set(uint8(FunctionSetSpotmeterConfig), buf.Bytes())
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

// Returns the spotmeter config as set by SetSpotmeterConfig.
func (device *ThermalImagingBricklet) GetSpotmeterConfig() (regionOfInterest [4]uint8, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetSpotmeterConfig), buf.Bytes())
	if err != nil {
		return regionOfInterest, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return regionOfInterest, DeviceError(header.ErrorCode)
		}

		if header.Length != 12 {
			return regionOfInterest, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &regionOfInterest)

	}

	return regionOfInterest, nil
}

// Sets the high contrast region of interest, dampening factor, clip limit and empty counts.
// This config is only used in high contrast mode (see SetImageTransferConfig).
//
// The high contrast region of interest consists of four values:
//
// * Index 0: Column start (has to be smaller than or equal to column end).
// * Index 1: Row start (has to be smaller than row end).
// * Index 2: Column end (has to be smaller than 80).
// * Index 3: Row end (has to be smaller than 60).
//
// The algorithm to generate the high contrast image is applied to this region.
//
// Dampening Factor: This parameter is the amount of temporal dampening applied to the HEQ
// (history equalization) transformation function. An IIR filter of the form::
//
//  (N / 256) * previous + ((256 - N) / 256) * current
//
// is applied, and the HEQ dampening factor
// represents the value N in the equation, i.e., a value that applies to the amount of
// influence the previous HEQ transformation function has on the current function. The
// lower the value of N the higher the influence of the current video frame whereas
// the higher the value of N the more influence the previous damped transfer function has.
//
// Clip Limit Index 0 (AGC HEQ Clip Limit High): This parameter defines the maximum number of pixels allowed
// to accumulate in any given histogram bin. Any additional pixels in a given bin are clipped.
// The effect of this parameter is to limit the influence of highly-populated bins on the
// resulting HEQ transformation function.
//
// Clip Limit Index 1 (AGC HEQ Clip Limit Low): This parameter defines an artificial population that is added to
// every non-empty histogram bin. In other words, if the Clip Limit Low is set to L, a bin
// with an actual population of X will have an effective population of L + X. Any empty bin
// that is nearby a populated bin will be given an artificial population of L. The effect of
// higher values is to provide a more linear transfer function; lower values provide a more
// non-linear (equalized) transfer function.
//
// Empty Counts: This parameter specifies the maximum number of pixels in a bin that will be
// interpreted as an empty bin. Histogram bins with this number of pixels or less will be
// processed as an empty bin.
func (device *ThermalImagingBricklet) SetHighContrastConfig(regionOfInterest [4]uint8, dampeningFactor uint16, clipLimit [2]uint16, emptyCounts uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, regionOfInterest)
	binary.Write(&buf, binary.LittleEndian, dampeningFactor)
	binary.Write(&buf, binary.LittleEndian, clipLimit)
	binary.Write(&buf, binary.LittleEndian, emptyCounts)

	resultBytes, err := device.device.Set(uint8(FunctionSetHighContrastConfig), buf.Bytes())
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

// Returns the high contrast config as set by SetHighContrastConfig.
func (device *ThermalImagingBricklet) GetHighContrastConfig() (regionOfInterest [4]uint8, dampeningFactor uint16, clipLimit [2]uint16, emptyCounts uint16, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetHighContrastConfig), buf.Bytes())
	if err != nil {
		return regionOfInterest, dampeningFactor, clipLimit, emptyCounts, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return regionOfInterest, dampeningFactor, clipLimit, emptyCounts, DeviceError(header.ErrorCode)
		}

		if header.Length != 20 {
			return regionOfInterest, dampeningFactor, clipLimit, emptyCounts, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 20)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &regionOfInterest)
		binary.Read(resultBuf, binary.LittleEndian, &dampeningFactor)
		binary.Read(resultBuf, binary.LittleEndian, &clipLimit)
		binary.Read(resultBuf, binary.LittleEndian, &emptyCounts)

	}

	return regionOfInterest, dampeningFactor, clipLimit, emptyCounts, nil
}

// The necessary bandwidth of this Bricklet is too high to use getter/callback or
// high contrast/temperature image at the same time. You have to configure the one
// you want to use, the Bricklet will optimize the internal configuration accordingly.
//
// Corresponding functions:
//
// * Manual High Contrast Image: GetHighContrastImage.
// * Manual Temperature Image: GetTemperatureImage.
// * Callback High Contrast Image: RegisterHighContrastImageCallback callback.
// * Callback Temperature Image: RegisterTemperatureImageCallback callback.
//
// Associated constants:
//
//	* ImageTransferManualHighContrastImage
//	* ImageTransferManualTemperatureImage
//	* ImageTransferCallbackHighContrastImage
//	* ImageTransferCallbackTemperatureImage
func (device *ThermalImagingBricklet) SetImageTransferConfig(config ImageTransfer) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, config)

	resultBytes, err := device.device.Set(uint8(FunctionSetImageTransferConfig), buf.Bytes())
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

// Returns the image transfer config, as set by SetImageTransferConfig.
//
// Associated constants:
//
//	* ImageTransferManualHighContrastImage
//	* ImageTransferManualTemperatureImage
//	* ImageTransferCallbackHighContrastImage
//	* ImageTransferCallbackTemperatureImage
func (device *ThermalImagingBricklet) GetImageTransferConfig() (config ImageTransfer, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetImageTransferConfig), buf.Bytes())
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

// Sets the flux linear parameters that can be used for radiometry calibration.
//
// See FLIR document 102-PS245-100-01 for more details.
//
// .. versionadded:: 2.0.5$nbsp;(Plugin)
func (device *ThermalImagingBricklet) SetFluxLinearParameters(sceneEmissivity uint16, temperatureBackground uint16, tauWindow uint16, temperaturWindow uint16, tauAtmosphere uint16, temperatureAtmosphere uint16, reflectionWindow uint16, temperatureReflection uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, sceneEmissivity)
	binary.Write(&buf, binary.LittleEndian, temperatureBackground)
	binary.Write(&buf, binary.LittleEndian, tauWindow)
	binary.Write(&buf, binary.LittleEndian, temperaturWindow)
	binary.Write(&buf, binary.LittleEndian, tauAtmosphere)
	binary.Write(&buf, binary.LittleEndian, temperatureAtmosphere)
	binary.Write(&buf, binary.LittleEndian, reflectionWindow)
	binary.Write(&buf, binary.LittleEndian, temperatureReflection)

	resultBytes, err := device.device.Set(uint8(FunctionSetFluxLinearParameters), buf.Bytes())
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

// Returns the flux linear parameters, as set by SetFluxLinearParameters.
//
// .. versionadded:: 2.0.5$nbsp;(Plugin)
func (device *ThermalImagingBricklet) GetFluxLinearParameters() (sceneEmissivity uint16, temperatureBackground uint16, tauWindow uint16, temperaturWindow uint16, tauAtmosphere uint16, temperatureAtmosphere uint16, reflectionWindow uint16, temperatureReflection uint16, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetFluxLinearParameters), buf.Bytes())
	if err != nil {
		return sceneEmissivity, temperatureBackground, tauWindow, temperaturWindow, tauAtmosphere, temperatureAtmosphere, reflectionWindow, temperatureReflection, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return sceneEmissivity, temperatureBackground, tauWindow, temperaturWindow, tauAtmosphere, temperatureAtmosphere, reflectionWindow, temperatureReflection, DeviceError(header.ErrorCode)
		}

		if header.Length != 24 {
			return sceneEmissivity, temperatureBackground, tauWindow, temperaturWindow, tauAtmosphere, temperatureAtmosphere, reflectionWindow, temperatureReflection, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 24)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &sceneEmissivity)
		binary.Read(resultBuf, binary.LittleEndian, &temperatureBackground)
		binary.Read(resultBuf, binary.LittleEndian, &tauWindow)
		binary.Read(resultBuf, binary.LittleEndian, &temperaturWindow)
		binary.Read(resultBuf, binary.LittleEndian, &tauAtmosphere)
		binary.Read(resultBuf, binary.LittleEndian, &temperatureAtmosphere)
		binary.Read(resultBuf, binary.LittleEndian, &reflectionWindow)
		binary.Read(resultBuf, binary.LittleEndian, &temperatureReflection)

	}

	return sceneEmissivity, temperatureBackground, tauWindow, temperaturWindow, tauAtmosphere, temperatureAtmosphere, reflectionWindow, temperatureReflection, nil
}

// Sets the FFC shutter mode parameters.
//
// See FLIR document 110-0144-03 4.5.15 for more details.
//
// .. versionadded:: 2.0.6$nbsp;(Plugin)
//
// Associated constants:
//
//	* ShutterModeManual
//	* ShutterModeAuto
//	* ShutterModeExternal
//	* ShutterLockoutInactive
//	* ShutterLockoutHigh
//	* ShutterLockoutLow
func (device *ThermalImagingBricklet) SetFFCShutterMode(shutterMode ShutterMode, tempLockoutState ShutterLockout, videoFreezeDuringFFC bool, ffcDesired bool, elapsedTimeSinceLastFFC uint32, desiredFFCPeriod uint32, explicitCmdToOpen bool, desiredFFCTempDelta uint16, imminentDelay uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, shutterMode)
	binary.Write(&buf, binary.LittleEndian, tempLockoutState)
	binary.Write(&buf, binary.LittleEndian, videoFreezeDuringFFC)
	binary.Write(&buf, binary.LittleEndian, ffcDesired)
	binary.Write(&buf, binary.LittleEndian, elapsedTimeSinceLastFFC)
	binary.Write(&buf, binary.LittleEndian, desiredFFCPeriod)
	binary.Write(&buf, binary.LittleEndian, explicitCmdToOpen)
	binary.Write(&buf, binary.LittleEndian, desiredFFCTempDelta)
	binary.Write(&buf, binary.LittleEndian, imminentDelay)

	resultBytes, err := device.device.Set(uint8(FunctionSetFFCShutterMode), buf.Bytes())
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

// Sets the FFC shutter mode parameters.
//
// See FLIR document 110-0144-03 4.5.15 for more details.
//
// .. versionadded:: 2.0.6$nbsp;(Plugin)
//
// Associated constants:
//
//	* ShutterModeManual
//	* ShutterModeAuto
//	* ShutterModeExternal
//	* ShutterLockoutInactive
//	* ShutterLockoutHigh
//	* ShutterLockoutLow
func (device *ThermalImagingBricklet) GetFFCShutterMode() (shutterMode ShutterMode, tempLockoutState ShutterLockout, videoFreezeDuringFFC bool, ffcDesired bool, elapsedTimeSinceLastFFC uint32, desiredFFCPeriod uint32, explicitCmdToOpen bool, desiredFFCTempDelta uint16, imminentDelay uint16, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetFFCShutterMode), buf.Bytes())
	if err != nil {
		return shutterMode, tempLockoutState, videoFreezeDuringFFC, ffcDesired, elapsedTimeSinceLastFFC, desiredFFCPeriod, explicitCmdToOpen, desiredFFCTempDelta, imminentDelay, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.ErrorCode != 0 {
			return shutterMode, tempLockoutState, videoFreezeDuringFFC, ffcDesired, elapsedTimeSinceLastFFC, desiredFFCPeriod, explicitCmdToOpen, desiredFFCTempDelta, imminentDelay, DeviceError(header.ErrorCode)
		}

		if header.Length != 25 {
			return shutterMode, tempLockoutState, videoFreezeDuringFFC, ffcDesired, elapsedTimeSinceLastFFC, desiredFFCPeriod, explicitCmdToOpen, desiredFFCTempDelta, imminentDelay, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 25)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &shutterMode)
		binary.Read(resultBuf, binary.LittleEndian, &tempLockoutState)
		binary.Read(resultBuf, binary.LittleEndian, &videoFreezeDuringFFC)
		binary.Read(resultBuf, binary.LittleEndian, &ffcDesired)
		binary.Read(resultBuf, binary.LittleEndian, &elapsedTimeSinceLastFFC)
		binary.Read(resultBuf, binary.LittleEndian, &desiredFFCPeriod)
		binary.Read(resultBuf, binary.LittleEndian, &explicitCmdToOpen)
		binary.Read(resultBuf, binary.LittleEndian, &desiredFFCTempDelta)
		binary.Read(resultBuf, binary.LittleEndian, &imminentDelay)

	}

	return shutterMode, tempLockoutState, videoFreezeDuringFFC, ffcDesired, elapsedTimeSinceLastFFC, desiredFFCPeriod, explicitCmdToOpen, desiredFFCTempDelta, imminentDelay, nil
}

// Starts the Flat-Field Correction (FFC) normalization.
//
// See FLIR document 110-0144-03 4.5.16 for more details.
//
// .. versionadded:: 2.0.6$nbsp;(Plugin)
func (device *ThermalImagingBricklet) RunFFCNormalization() (err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Set(uint8(FunctionRunFFCNormalization), buf.Bytes())
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
func (device *ThermalImagingBricklet) GetSPITFPErrorCount() (errorCountAckChecksum uint32, errorCountMessageChecksum uint32, errorCountFrame uint32, errorCountOverflow uint32, err error) {
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
func (device *ThermalImagingBricklet) SetBootloaderMode(mode BootloaderMode) (status BootloaderStatus, err error) {
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
func (device *ThermalImagingBricklet) GetBootloaderMode() (mode BootloaderMode, err error) {
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
func (device *ThermalImagingBricklet) SetWriteFirmwarePointer(pointer uint32) (err error) {
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
func (device *ThermalImagingBricklet) WriteFirmware(data [64]uint8) (status uint8, err error) {
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
func (device *ThermalImagingBricklet) SetStatusLEDConfig(config StatusLEDConfig) (err error) {
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
func (device *ThermalImagingBricklet) GetStatusLEDConfig() (config StatusLEDConfig, err error) {
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
func (device *ThermalImagingBricklet) GetChipTemperature() (temperature int16, err error) {
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
func (device *ThermalImagingBricklet) Reset() (err error) {
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
func (device *ThermalImagingBricklet) WriteUID(uid uint32) (err error) {
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
func (device *ThermalImagingBricklet) ReadUID() (uid uint32, err error) {
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
func (device *ThermalImagingBricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
