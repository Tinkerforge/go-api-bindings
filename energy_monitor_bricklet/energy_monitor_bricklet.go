/* ***********************************************************
 * This file was automatically generated on 2020-05-19.      *
 *                                                           *
 * Go Bindings Version 2.0.8                                 *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


// Measures Voltage, Current, Energy, Real/Apparent/Reactive Power, Power Factor and Frequency.
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/EnergyMonitor_Bricklet_Go.html.
package energy_monitor_bricklet

import (
	"encoding/binary"
	"bytes"
	"fmt"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionGetEnergyData Function = 1
	FunctionResetEnergy Function = 2
	FunctionGetWaveformLowLevel Function = 3
	FunctionGetTransformerStatus Function = 4
	FunctionSetTransformerCalibration Function = 5
	FunctionGetTransformerCalibration Function = 6
	FunctionCalibrateOffset Function = 7
	FunctionSetEnergyDataCallbackConfiguration Function = 8
	FunctionGetEnergyDataCallbackConfiguration Function = 9
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
	FunctionCallbackEnergyData Function = 10
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

type EnergyMonitorBricklet struct {
	device Device
}
const DeviceIdentifier = 2152
const DeviceDisplayName = "Energy Monitor Bricklet"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (EnergyMonitorBricklet, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{ 2,0,0 }, uid, &internalIPCon, 1, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return EnergyMonitorBricklet{}, err
	}
	dev.ResponseExpected[FunctionGetEnergyData] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionResetEnergy] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetWaveformLowLevel] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetTransformerStatus] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetTransformerCalibration] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetTransformerCalibration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionCalibrateOffset] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionSetEnergyDataCallbackConfiguration] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetEnergyDataCallbackConfiguration] = ResponseExpectedFlagAlwaysTrue;
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
	return EnergyMonitorBricklet{dev}, nil
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
func (device *EnergyMonitorBricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *EnergyMonitorBricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *EnergyMonitorBricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *EnergyMonitorBricklet) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// This callback is triggered periodically according to the configuration set by
// SetEnergyDataCallbackConfiguration.
// 
// The parameters are the same as GetEnergyData.
func (device *EnergyMonitorBricklet) RegisterEnergyDataCallback(fn func(int32, int32, int32, int32, int32, int32, uint16, uint16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 36 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var voltage int32
		var current int32
		var energy int32
		var realPower int32
		var apparentPower int32
		var reactivePower int32
		var powerFactor uint16
		var frequency uint16
		binary.Read(buf, binary.LittleEndian, &voltage)
		binary.Read(buf, binary.LittleEndian, &current)
		binary.Read(buf, binary.LittleEndian, &energy)
		binary.Read(buf, binary.LittleEndian, &realPower)
		binary.Read(buf, binary.LittleEndian, &apparentPower)
		binary.Read(buf, binary.LittleEndian, &reactivePower)
		binary.Read(buf, binary.LittleEndian, &powerFactor)
		binary.Read(buf, binary.LittleEndian, &frequency)
		fn(voltage, current, energy, realPower, apparentPower, reactivePower, powerFactor, frequency)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackEnergyData), wrapper)
}

// Remove a registered Energy Data callback.
func (device *EnergyMonitorBricklet) DeregisterEnergyDataCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackEnergyData), registrationId)
}


// Returns all of the measurements that are done by the Energy Monitor Bricklet.
// 
// * Voltage RMS
// * Current RMS
// * Energy (integrated over time)
// * Real Power
// * Apparent Power
// * Reactive Power
// * Power Factor
// * Frequency (AC Frequency of the mains voltage)
// 
// The frequency is recalculated every 6 seconds.
// 
// All other values are integrated over 10 zero-crossings of the voltage sine wave.
// With a standard AC mains voltage frequency of 50Hz this results in a 5 measurements
// per second (or an integration time of 200ms per measurement).
// 
// If no voltage transformer is connected, the Bricklet will use the current waveform
// to calculate the frequency and it will use an integration time of
// 10 zero-crossings of the current waveform.
func (device *EnergyMonitorBricklet) GetEnergyData() (voltage int32, current int32, energy int32, realPower int32, apparentPower int32, reactivePower int32, powerFactor uint16, frequency uint16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetEnergyData), buf.Bytes())
	if err != nil {
		return voltage, current, energy, realPower, apparentPower, reactivePower, powerFactor, frequency, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 36 {
			return voltage, current, energy, realPower, apparentPower, reactivePower, powerFactor, frequency, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 36)
		}

		if header.ErrorCode != 0 {
			return voltage, current, energy, realPower, apparentPower, reactivePower, powerFactor, frequency, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &voltage)
		binary.Read(resultBuf, binary.LittleEndian, &current)
		binary.Read(resultBuf, binary.LittleEndian, &energy)
		binary.Read(resultBuf, binary.LittleEndian, &realPower)
		binary.Read(resultBuf, binary.LittleEndian, &apparentPower)
		binary.Read(resultBuf, binary.LittleEndian, &reactivePower)
		binary.Read(resultBuf, binary.LittleEndian, &powerFactor)
		binary.Read(resultBuf, binary.LittleEndian, &frequency)

	}

	return voltage, current, energy, realPower, apparentPower, reactivePower, powerFactor, frequency, nil
}

// Sets the energy value (see GetEnergyData) back to 0Wh.
func (device *EnergyMonitorBricklet) ResetEnergy() (err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Set(uint8(FunctionResetEnergy), buf.Bytes())
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

// Returns a snapshot of the voltage and current waveform. The values
// in the returned array alternate between voltage and current. The data from
// one getter call contains 768 data points for voltage and current, which
// correspond to about 3 full sine waves.
// 
// The voltage is given with a resolution of 100mV and the current is given
// with a resolution of 10mA.
// 
// This data is meant to be used for a non-realtime graphical representation of
// the voltage and current waveforms.
func (device *EnergyMonitorBricklet) GetWaveformLowLevel() (waveformChunkOffset uint16, waveformChunkData [30]int16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetWaveformLowLevel), buf.Bytes())
	if err != nil {
		return waveformChunkOffset, waveformChunkData, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 70 {
			return waveformChunkOffset, waveformChunkData, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 70)
		}

		if header.ErrorCode != 0 {
			return waveformChunkOffset, waveformChunkData, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &waveformChunkOffset)
		copy(waveformChunkData[:], ByteSliceToInt16Slice(resultBuf.Next(16 * 30/8)))

	}

	return waveformChunkOffset, waveformChunkData, nil
}

// Returns a snapshot of the voltage and current waveform. The values
// in the returned array alternate between voltage and current. The data from
// one getter call contains 768 data points for voltage and current, which
// correspond to about 3 full sine waves.
// 
// The voltage is given with a resolution of 100mV and the current is given
// with a resolution of 10mA.
// 
// This data is meant to be used for a non-realtime graphical representation of
// the voltage and current waveforms.
	func (device *EnergyMonitorBricklet) GetWaveform() (waveform []int16, err error) {
		buf, _, err := device.device.GetHighLevel(func() (LowLevelResult, error) {
			waveformChunkOffset, waveformChunkData, err := device.GetWaveformLowLevel()

			if err != nil {
				return LowLevelResult{}, err
			}

			var lowLevelResults bytes.Buffer
			

			return LowLevelResult{
				uint64(1536),
				uint64(waveformChunkOffset),
				Int16SliceToByteSlice(waveformChunkData[:]),
				lowLevelResults.Bytes()}, nil
		},
			0,
			16)
		if err != nil {
			return ByteSliceToInt16Slice(buf), err
		}
		
		
		return ByteSliceToInt16Slice(buf), nil
	}

// Returns *true* if a voltage/current transformer is connected to the Bricklet.
func (device *EnergyMonitorBricklet) GetTransformerStatus() (voltageTransformerConnected bool, currentTransformerConnected bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetTransformerStatus), buf.Bytes())
	if err != nil {
		return voltageTransformerConnected, currentTransformerConnected, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 10 {
			return voltageTransformerConnected, currentTransformerConnected, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		if header.ErrorCode != 0 {
			return voltageTransformerConnected, currentTransformerConnected, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &voltageTransformerConnected)
		binary.Read(resultBuf, binary.LittleEndian, &currentTransformerConnected)

	}

	return voltageTransformerConnected, currentTransformerConnected, nil
}

// Sets the transformer ratio for the voltage and current transformer in 1/100 form.
// 
// Example: If your mains voltage is 230V, you use 9V voltage transformer and a
// 1V:30A current clamp your voltage ratio is 230/9 = 25.56 and your current ratio
// is 30/1 = 30.
// 
// In this case you have to set the values 2556 and 3000 for voltage ratio and current
// ratio.
// 
// The calibration is saved in non-volatile memory, you only have to set it once.
// 
// Set the phase shift to 0. It is for future use and currently not supported by the Bricklet.
func (device *EnergyMonitorBricklet) SetTransformerCalibration(voltageRatio uint16, currentRatio uint16, phaseShift int16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, voltageRatio);
	binary.Write(&buf, binary.LittleEndian, currentRatio);
	binary.Write(&buf, binary.LittleEndian, phaseShift);

	resultBytes, err := device.device.Set(uint8(FunctionSetTransformerCalibration), buf.Bytes())
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

// Returns the transformer calibration as set by SetTransformerCalibration.
func (device *EnergyMonitorBricklet) GetTransformerCalibration() (voltageRatio uint16, currentRatio uint16, phaseShift int16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetTransformerCalibration), buf.Bytes())
	if err != nil {
		return voltageRatio, currentRatio, phaseShift, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 14 {
			return voltageRatio, currentRatio, phaseShift, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 14)
		}

		if header.ErrorCode != 0 {
			return voltageRatio, currentRatio, phaseShift, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &voltageRatio)
		binary.Read(resultBuf, binary.LittleEndian, &currentRatio)
		binary.Read(resultBuf, binary.LittleEndian, &phaseShift)

	}

	return voltageRatio, currentRatio, phaseShift, nil
}

// Calling this function will start an offset calibration. The offset calibration will
// integrate the voltage and current waveform over a longer time period to find the 0
// transition point in the sine wave.
// 
// The Bricklet comes with a factory-calibrated offset value, you should not have to
// call this function.
// 
// If you want to re-calibrate the offset we recommend that you connect a load that
// has a smooth sinusoidal voltage and current waveform. Alternatively you can also
// short both inputs.
// 
// The calibration is saved in non-volatile memory, you only have to set it once.
func (device *EnergyMonitorBricklet) CalibrateOffset() (err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Set(uint8(FunctionCalibrateOffset), buf.Bytes())
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

// The period is the period with which the RegisterEnergyDataCallback
// callback is triggered periodically. A value of 0 turns the callback off.
// 
// If the `value has to change`-parameter is set to true, the callback is only
// triggered after the value has changed. If the value didn't change within the
// period, the callback is triggered immediately on change.
// 
// If it is set to false, the callback is continuously triggered with the period,
// independent of the value.
func (device *EnergyMonitorBricklet) SetEnergyDataCallbackConfiguration(period uint32, valueHasToChange bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);
	binary.Write(&buf, binary.LittleEndian, valueHasToChange);

	resultBytes, err := device.device.Set(uint8(FunctionSetEnergyDataCallbackConfiguration), buf.Bytes())
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
// SetEnergyDataCallbackConfiguration.
func (device *EnergyMonitorBricklet) GetEnergyDataCallbackConfiguration() (period uint32, valueHasToChange bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetEnergyDataCallbackConfiguration), buf.Bytes())
	if err != nil {
		return period, valueHasToChange, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 13 {
			return period, valueHasToChange, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 13)
		}

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
func (device *EnergyMonitorBricklet) GetSPITFPErrorCount() (errorCountAckChecksum uint32, errorCountMessageChecksum uint32, errorCountFrame uint32, errorCountOverflow uint32, err error) {
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
func (device *EnergyMonitorBricklet) SetBootloaderMode(mode BootloaderMode) (status BootloaderStatus, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, mode);

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
func (device *EnergyMonitorBricklet) GetBootloaderMode() (mode BootloaderMode, err error) {
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
func (device *EnergyMonitorBricklet) SetWriteFirmwarePointer(pointer uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, pointer);

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
func (device *EnergyMonitorBricklet) WriteFirmware(data [64]uint8) (status uint8, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, data);

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
func (device *EnergyMonitorBricklet) SetStatusLEDConfig(config StatusLEDConfig) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, config);

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
func (device *EnergyMonitorBricklet) GetStatusLEDConfig() (config StatusLEDConfig, err error) {
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
func (device *EnergyMonitorBricklet) GetChipTemperature() (temperature int16, err error) {
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
func (device *EnergyMonitorBricklet) Reset() (err error) {
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
func (device *EnergyMonitorBricklet) WriteUID(uid uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, uid);

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
func (device *EnergyMonitorBricklet) ReadUID() (uid uint32, err error) {
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
// The Raspberry Pi HAT (Zero) Brick is always at position 'i' and the Bricklet
// connected to an `Isolator Bricklet <isolator_bricklet>` is always as
// position 'z'.
// 
// The device identifier numbers can be found `here <device_identifier>`.
// |device_identifier_constant|
func (device *EnergyMonitorBricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
