/* ***********************************************************
 * This file was automatically generated on 2022-08-08.      *
 *                                                           *
 * Go Bindings Version 2.0.13                                *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/

// Measures Sound Pressure Level in dB(A/B/C/D/Z).
//
//
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/SoundPressureLevel_Bricklet_Go.html.
package sound_pressure_level_bricklet

import (
	"bytes"
	"encoding/binary"
	"fmt"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionGetDecibel                       Function = 1
	FunctionSetDecibelCallbackConfiguration  Function = 2
	FunctionGetDecibelCallbackConfiguration  Function = 3
	FunctionGetSpectrumLowLevel              Function = 5
	FunctionSetSpectrumCallbackConfiguration Function = 6
	FunctionGetSpectrumCallbackConfiguration Function = 7
	FunctionSetConfiguration                 Function = 9
	FunctionGetConfiguration                 Function = 10
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
	FunctionCallbackDecibel                  Function = 4
	FunctionCallbackSpectrumLowLevel         Function = 8
)

type ThresholdOption = rune

const (
	ThresholdOptionOff     ThresholdOption = 'x'
	ThresholdOptionOutside ThresholdOption = 'o'
	ThresholdOptionInside  ThresholdOption = 'i'
	ThresholdOptionSmaller ThresholdOption = '<'
	ThresholdOptionGreater ThresholdOption = '>'
)

type FFTSize = uint8

const (
	FFTSize128  FFTSize = 0
	FFTSize256  FFTSize = 1
	FFTSize512  FFTSize = 2
	FFTSize1024 FFTSize = 3
)

type Weighting = uint8

const (
	WeightingA       Weighting = 0
	WeightingB       Weighting = 1
	WeightingC       Weighting = 2
	WeightingD       Weighting = 3
	WeightingZ       Weighting = 4
	WeightingITUR468 Weighting = 5
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

type SoundPressureLevelBricklet struct {
	device Device
}

const DeviceIdentifier = 290
const DeviceDisplayName = "Sound Pressure Level Bricklet"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (SoundPressureLevelBricklet, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{2, 0, 0}, uid, &internalIPCon, 2, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return SoundPressureLevelBricklet{}, err
	}
	dev.ResponseExpected[FunctionGetDecibel] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetDecibelCallbackConfiguration] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionGetDecibelCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionGetSpectrumLowLevel] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetSpectrumCallbackConfiguration] = ResponseExpectedFlagTrue
	dev.ResponseExpected[FunctionGetSpectrumCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue
	dev.ResponseExpected[FunctionSetConfiguration] = ResponseExpectedFlagFalse
	dev.ResponseExpected[FunctionGetConfiguration] = ResponseExpectedFlagAlwaysTrue
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
	return SoundPressureLevelBricklet{dev}, nil
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
func (device *SoundPressureLevelBricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *SoundPressureLevelBricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *SoundPressureLevelBricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *SoundPressureLevelBricklet) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// This callback is triggered periodically according to the configuration set by
// SetDecibelCallbackConfiguration.
//
// The parameter is the same as GetDecibel.
func (device *SoundPressureLevelBricklet) RegisterDecibelCallback(fn func(uint16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 10 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var decibel uint16
		binary.Read(buf, binary.LittleEndian, &decibel)
		fn(decibel)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackDecibel), wrapper)
}

// Remove a registered Decibel callback.
func (device *SoundPressureLevelBricklet) DeregisterDecibelCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackDecibel), registrationId)
}

// See RegisterSpectrumCallback
func (device *SoundPressureLevelBricklet) RegisterSpectrumLowLevelCallback(fn func(uint16, uint16, [30]uint16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 72 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var spectrumLength uint16
		var spectrumChunkOffset uint16
		var spectrumChunkData [30]uint16
		binary.Read(buf, binary.LittleEndian, &spectrumLength)
		binary.Read(buf, binary.LittleEndian, &spectrumChunkOffset)
		copy(spectrumChunkData[:], ByteSliceToUint16Slice(buf.Next(16*30/8)))
		fn(spectrumLength, spectrumChunkOffset, spectrumChunkData)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackSpectrumLowLevel), wrapper)
}

// Remove a registered Spectrum Low Level callback.
func (device *SoundPressureLevelBricklet) DeregisterSpectrumLowLevelCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackSpectrumLowLevel), registrationId)
}

// This callback is triggered periodically according to the configuration set by
// SetSpectrumCallbackConfiguration.
//
// The parameter is the same as GetSpectrum.
//
// Note
//  If reconstructing the value fails, the callback is triggered with nil for spectrum.
func (device *SoundPressureLevelBricklet) RegisterSpectrumCallback(fn func([]uint16)) uint64 {
	buf := make([]uint16, 0)
	wrapper := func(spectrumLength uint16, spectrumChunkOffset uint16, spectrumChunkData [30]uint16) {
		if int(spectrumChunkOffset) != len(buf) {
			if len(buf) != 0 {
				buf = nil
				fn(buf)
				buf = make([]uint16, 0)
			}
			if spectrumChunkOffset != 0 {
				return
			}
		}
		toRead := MinU(uint64(spectrumLength-spectrumChunkOffset), uint64(len(spectrumChunkData[:])))
		buf = append(buf, spectrumChunkData[:toRead]...)
		if len(buf) >= int(spectrumLength) {
			fn(buf)
			buf = make([]uint16, 0)
		}
	}
	return device.RegisterSpectrumLowLevelCallback(wrapper)
}

// Remove a registered Spectrum Low Level callback.
func (device *SoundPressureLevelBricklet) DeregisterSpectrumCallback(registrationId uint64) {
	device.DeregisterSpectrumLowLevelCallback(registrationId)
}

// Returns the measured sound pressure in decibels.
//
// The Bricklet supports the weighting standards dB(A), dB(B), dB(C), dB(D),
// dB(Z) and ITU-R 468. You can configure the weighting with SetConfiguration.
//
// By default dB(A) will be used.
//
//
// If you want to get the value periodically, it is recommended to use the
// RegisterDecibelCallback callback. You can set the callback configuration
// with SetDecibelCallbackConfiguration.
func (device *SoundPressureLevelBricklet) GetDecibel() (decibel uint16, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetDecibel), buf.Bytes())
	if err != nil {
		return decibel, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 10 {
			return decibel, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		if header.ErrorCode != 0 {
			return decibel, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &decibel)

	}

	return decibel, nil
}

// The period is the period with which the RegisterDecibelCallback callback is triggered
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
// The `option`-parameter together with min/max sets a threshold for the RegisterDecibelCallback callback.
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
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *SoundPressureLevelBricklet) SetDecibelCallbackConfiguration(period uint32, valueHasToChange bool, option ThresholdOption, min uint16, max uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period)
	binary.Write(&buf, binary.LittleEndian, valueHasToChange)
	binary.Write(&buf, binary.LittleEndian, option)
	binary.Write(&buf, binary.LittleEndian, min)
	binary.Write(&buf, binary.LittleEndian, max)

	resultBytes, err := device.device.Set(uint8(FunctionSetDecibelCallbackConfiguration), buf.Bytes())
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

// Returns the callback configuration as set by SetDecibelCallbackConfiguration.
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *SoundPressureLevelBricklet) GetDecibelCallbackConfiguration() (period uint32, valueHasToChange bool, option ThresholdOption, min uint16, max uint16, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetDecibelCallbackConfiguration), buf.Bytes())
	if err != nil {
		return period, valueHasToChange, option, min, max, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 18 {
			return period, valueHasToChange, option, min, max, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 18)
		}

		if header.ErrorCode != 0 {
			return period, valueHasToChange, option, min, max, DeviceError(header.ErrorCode)
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

// Returns the frequency spectrum. The length of the spectrum is between
// 512 (FFT size 1024) and 64 (FFT size 128). See SetConfiguration.
//
// Each array element is one bin of the FFT. The first bin is always the
// DC offset and the other bins have a size between 40Hz (FFT size 1024) and
// 320Hz (FFT size 128).
//
// In sum the frequency of the spectrum always has a range from 0 to
// 20480Hz (the FFT is applied to samples with a frequency of 40960Hz).
//
// The returned data is already equalized, which means that the microphone
// frequency response is compensated and the weighting function is applied
// (see SetConfiguration for the available weighting standards). Use
// dB(Z) if you need the unaltered spectrum.
//
// The values are not in dB form yet. If you want a proper dB scale of the
// spectrum you have to apply the formula f(x) = 20*log10(max(1, x/sqrt(2)))
// on each value.
func (device *SoundPressureLevelBricklet) GetSpectrumLowLevel() (spectrumLength uint16, spectrumChunkOffset uint16, spectrumChunkData [30]uint16, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetSpectrumLowLevel), buf.Bytes())
	if err != nil {
		return spectrumLength, spectrumChunkOffset, spectrumChunkData, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 72 {
			return spectrumLength, spectrumChunkOffset, spectrumChunkData, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 72)
		}

		if header.ErrorCode != 0 {
			return spectrumLength, spectrumChunkOffset, spectrumChunkData, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &spectrumLength)
		binary.Read(resultBuf, binary.LittleEndian, &spectrumChunkOffset)
		copy(spectrumChunkData[:], ByteSliceToUint16Slice(resultBuf.Next(16*30/8)))

	}

	return spectrumLength, spectrumChunkOffset, spectrumChunkData, nil
}

// Returns the frequency spectrum. The length of the spectrum is between
// 512 (FFT size 1024) and 64 (FFT size 128). See SetConfiguration.
//
// Each array element is one bin of the FFT. The first bin is always the
// DC offset and the other bins have a size between 40Hz (FFT size 1024) and
// 320Hz (FFT size 128).
//
// In sum the frequency of the spectrum always has a range from 0 to
// 20480Hz (the FFT is applied to samples with a frequency of 40960Hz).
//
// The returned data is already equalized, which means that the microphone
// frequency response is compensated and the weighting function is applied
// (see SetConfiguration for the available weighting standards). Use
// dB(Z) if you need the unaltered spectrum.
//
// The values are not in dB form yet. If you want a proper dB scale of the
// spectrum you have to apply the formula f(x) = 20*log10(max(1, x/sqrt(2)))
// on each value.
func (device *SoundPressureLevelBricklet) GetSpectrum() (spectrum []uint16, err error) {
	buf, _, err := device.device.GetHighLevel(func() (LowLevelResult, error) {
		spectrumLength, spectrumChunkOffset, spectrumChunkData, err := device.GetSpectrumLowLevel()

		if err != nil {
			return LowLevelResult{}, err
		}

		var lowLevelResults bytes.Buffer

		return LowLevelResult{
			uint64(spectrumLength),
			uint64(spectrumChunkOffset),
			Uint16SliceToByteSlice(spectrumChunkData[:]),
			lowLevelResults.Bytes()}, nil
	},
		0,
		16)
	if err != nil {
		return ByteSliceToUint16Slice(buf), err
	}

	return ByteSliceToUint16Slice(buf), nil
}

// The period is the period with which the RegisterSpectrumCallback callback is
// triggered periodically. A value of 0 turns the callback off.
//
// Every new measured spectrum will be send at most once. Set the period to 1 to
// make sure that you get every spectrum.
func (device *SoundPressureLevelBricklet) SetSpectrumCallbackConfiguration(period uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period)

	resultBytes, err := device.device.Set(uint8(FunctionSetSpectrumCallbackConfiguration), buf.Bytes())
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

// Returns the callback configuration as set by
// GetSpectrumCallbackConfiguration.
func (device *SoundPressureLevelBricklet) GetSpectrumCallbackConfiguration() (period uint32, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetSpectrumCallbackConfiguration), buf.Bytes())
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

// Sets the Sound Pressure Level Bricklet configuration.
//
// With different FFT sizes the Bricklet has a different
// amount of samples per second and the size of the FFT bins
// changes. The higher the FFT size the more precise is the result
// of the dB(X) calculation.
//
// Available FFT sizes are:
//
// * 1024: 512 bins, 10 samples per second, each bin has size 40Hz
// * 512: 256 bins, 20 samples per second, each bin has size 80Hz
// * 256: 128 bins, 40 samples per second, each bin has size 160Hz
// * 128: 64 bins, 80 samples per second, each bin has size 320Hz
//
// The Bricklet supports different weighting functions. You can choose
// between dB(A), dB(B), dB(C), dB(D), dB(Z) and ITU-R 468.
//
// dB(A/B/C/D) are the standard dB weighting curves. dB(A) is
// often used to measure volumes at concerts etc. dB(Z) has a
// flat response, no weighting is applied. ITU-R 468 is an ITU
// weighting standard mostly used in the UK and Europe.
//
// Associated constants:
//
//	* FFTSize128
//	* FFTSize256
//	* FFTSize512
//	* FFTSize1024
//	* WeightingA
//	* WeightingB
//	* WeightingC
//	* WeightingD
//	* WeightingZ
//	* WeightingITUR468
func (device *SoundPressureLevelBricklet) SetConfiguration(fftSize FFTSize, weighting Weighting) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, fftSize)
	binary.Write(&buf, binary.LittleEndian, weighting)

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
//	* FFTSize128
//	* FFTSize256
//	* FFTSize512
//	* FFTSize1024
//	* WeightingA
//	* WeightingB
//	* WeightingC
//	* WeightingD
//	* WeightingZ
//	* WeightingITUR468
func (device *SoundPressureLevelBricklet) GetConfiguration() (fftSize FFTSize, weighting Weighting, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetConfiguration), buf.Bytes())
	if err != nil {
		return fftSize, weighting, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 10 {
			return fftSize, weighting, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		if header.ErrorCode != 0 {
			return fftSize, weighting, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &fftSize)
		binary.Read(resultBuf, binary.LittleEndian, &weighting)

	}

	return fftSize, weighting, nil
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
func (device *SoundPressureLevelBricklet) GetSPITFPErrorCount() (errorCountAckChecksum uint32, errorCountMessageChecksum uint32, errorCountFrame uint32, errorCountOverflow uint32, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetSPITFPErrorCount), buf.Bytes())
	if err != nil {
		return errorCountAckChecksum, errorCountMessageChecksum, errorCountFrame, errorCountOverflow, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 24 {
			return errorCountAckChecksum, errorCountMessageChecksum, errorCountFrame, errorCountOverflow, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 24)
		}

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
func (device *SoundPressureLevelBricklet) SetBootloaderMode(mode BootloaderMode) (status BootloaderStatus, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, mode)

	resultBytes, err := device.device.Get(uint8(FunctionSetBootloaderMode), buf.Bytes())
	if err != nil {
		return status, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return status, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

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
func (device *SoundPressureLevelBricklet) GetBootloaderMode() (mode BootloaderMode, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetBootloaderMode), buf.Bytes())
	if err != nil {
		return mode, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return mode, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

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
func (device *SoundPressureLevelBricklet) SetWriteFirmwarePointer(pointer uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, pointer)

	resultBytes, err := device.device.Set(uint8(FunctionSetWriteFirmwarePointer), buf.Bytes())
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

// Writes 64 Bytes of firmware at the position as written by
// SetWriteFirmwarePointer before. The firmware is written
// to flash every 4 chunks.
//
// You can only write firmware in bootloader mode.
//
// This function is used by Brick Viewer during flashing. It should not be
// necessary to call it in a normal user program.
func (device *SoundPressureLevelBricklet) WriteFirmware(data [64]uint8) (status uint8, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, data)

	resultBytes, err := device.device.Get(uint8(FunctionWriteFirmware), buf.Bytes())
	if err != nil {
		return status, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return status, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

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
func (device *SoundPressureLevelBricklet) SetStatusLEDConfig(config StatusLEDConfig) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, config)

	resultBytes, err := device.device.Set(uint8(FunctionSetStatusLEDConfig), buf.Bytes())
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

// Returns the configuration as set by SetStatusLEDConfig
//
// Associated constants:
//
//	* StatusLEDConfigOff
//	* StatusLEDConfigOn
//	* StatusLEDConfigShowHeartbeat
//	* StatusLEDConfigShowStatus
func (device *SoundPressureLevelBricklet) GetStatusLEDConfig() (config StatusLEDConfig, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetStatusLEDConfig), buf.Bytes())
	if err != nil {
		return config, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return config, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return config, DeviceError(header.ErrorCode)
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
func (device *SoundPressureLevelBricklet) GetChipTemperature() (temperature int16, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionGetChipTemperature), buf.Bytes())
	if err != nil {
		return temperature, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 10 {
			return temperature, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

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
func (device *SoundPressureLevelBricklet) Reset() (err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Set(uint8(FunctionReset), buf.Bytes())
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

// Writes a new UID into flash. If you want to set a new UID
// you have to decode the Base58 encoded UID string into an
// integer first.
//
// We recommend that you use Brick Viewer to change the UID.
func (device *SoundPressureLevelBricklet) WriteUID(uid uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, uid)

	resultBytes, err := device.device.Set(uint8(FunctionWriteUID), buf.Bytes())
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

// Returns the current UID as an integer. Encode as
// Base58 to get the usual string version.
func (device *SoundPressureLevelBricklet) ReadUID() (uid uint32, err error) {
	var buf bytes.Buffer

	resultBytes, err := device.device.Get(uint8(FunctionReadUID), buf.Bytes())
	if err != nil {
		return uid, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 12 {
			return uid, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

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
// The position can be 'a', 'b', 'c', 'd', 'e', 'f', 'g' or 'h' (Bricklet Port).
// A Bricklet connected to an `Isolator Bricklet <isolator_bricklet>` is always at
// position 'z'.
//
// The device identifier numbers can be found `here <device_identifier>`.
// |device_identifier_constant|
func (device *SoundPressureLevelBricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
