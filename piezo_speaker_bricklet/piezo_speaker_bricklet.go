/* ***********************************************************
 * This file was automatically generated on 2018-12-18.      *
 *                                                           *
 * Go Bindings Version 2.0.0                                 *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


//Creates beep with configurable frequency.
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/PiezoSpeaker_Bricklet_Go.html.
package piezo_speaker_bricklet

import (
	"encoding/binary"
	"bytes"
    . "github.com/tinkerforge/go-api-bindings/internal"
    "github.com/tinkerforge/go-api-bindings/ipconnection"
)

type Function uint8

const (
    FunctionBeep Function = 1
	FunctionMorseCode Function = 2
	FunctionCalibrate Function = 3
	FunctionGetIdentity Function = 255
	FunctionCallbackBeepFinished Function = 4
	FunctionCallbackMorseCodeFinished Function = 5
)

type BeepDuration uint32

const (
    BeepDurationOff BeepDuration = 0
	BeepDurationInfinite BeepDuration = 4294967295
)

type PiezoSpeakerBricklet struct{
	device Device
}
const DeviceIdentifier = 242
const DeviceDisplayName = "Piezo Speaker Bricklet"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (PiezoSpeakerBricklet, error) {
    internalIPCon := ipcon.GetInternalHandle().(IPConnection)
    dev, err := NewDevice([3]uint8{ 2,0,0 }, uid, &internalIPCon, 0)
    if err != nil {
        return PiezoSpeakerBricklet{}, err
    }
    dev.ResponseExpected[FunctionBeep] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionMorseCode] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionCalibrate] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetIdentity] = ResponseExpectedFlagAlwaysTrue;
    return PiezoSpeakerBricklet{dev}, nil
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
func (device *PiezoSpeakerBricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *PiezoSpeakerBricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
    return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *PiezoSpeakerBricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *PiezoSpeakerBricklet) GetAPIVersion() [3]uint8 {
    return device.device.GetAPIVersion()
}

// This callback is triggered if a beep set by Beep is finished
func (device *PiezoSpeakerBricklet) RegisterBeepFinishedCallback(fn func()) uint64 {
            wrapper := func(byteSlice []byte) {
                
                
                
                fn()
            }
    return device.device.RegisterCallback(uint8(FunctionCallbackBeepFinished), wrapper)
}

//Remove a registered Beep Finished callback.
func (device *PiezoSpeakerBricklet) DeregisterBeepFinishedCallback(callbackID uint64) {
    device.device.DeregisterCallback(uint8(FunctionCallbackBeepFinished), callbackID)
}


// This callback is triggered if the playback of the morse code set by
	// MorseCode is finished.
func (device *PiezoSpeakerBricklet) RegisterMorseCodeFinishedCallback(fn func()) uint64 {
            wrapper := func(byteSlice []byte) {
                
                
                
                fn()
            }
    return device.device.RegisterCallback(uint8(FunctionCallbackMorseCodeFinished), wrapper)
}

//Remove a registered Morse Code Finished callback.
func (device *PiezoSpeakerBricklet) DeregisterMorseCodeFinishedCallback(callbackID uint64) {
    device.device.DeregisterCallback(uint8(FunctionCallbackMorseCodeFinished), callbackID)
}


// Beeps with the given frequency for the duration in ms. For example:
	// If you set a duration of 1000, with a frequency value of 2000
	// the piezo buzzer will beep for one second with a frequency of
	// approximately 2 kHz.
	// 
	// .. versionchanged:: 2.0.2$nbsp;(Plugin)
	//    A duration of 0 stops the current beep if any, the frequency parameter is
	//    ignored. A duration of 4294967295 results in an infinite beep.
	// 
	// The *frequency* parameter can be set between 585 and 7100.
	// 
	// The Piezo Speaker Bricklet can only approximate the frequency, it will play
	// the best possible match by applying the calibration (see Calibrate).
//
// Associated constants:
//
//	* BeepDurationOff
//	* BeepDurationInfinite
func (device *PiezoSpeakerBricklet) Beep(duration BeepDuration, frequency uint16) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, duration);
	binary.Write(&buf, binary.LittleEndian, frequency);

    resultBytes, err := device.device.Set(uint8(FunctionBeep), buf.Bytes())
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

// Sets morse code that will be played by the piezo buzzer. The morse code
	// is given as a string consisting of . (dot), - (minus) and   (space)
	// for *dits*, *dahs* and *pauses*. Every other character is ignored.
	// The second parameter is the frequency (see Beep).
	// 
	// For example: If you set the string ...---..., the piezo buzzer will beep
	// nine times with the durations short short short long long long short
	// short short.
	// 
	// The maximum string size is 60.
func (device *PiezoSpeakerBricklet) MorseCode(morse string, frequency uint16) (err error) {    
        var buf bytes.Buffer
    morse_byte_slice, err := StringToByteSlice(morse, 60)
	if err != nil { return }
	buf.Write(morse_byte_slice)
	binary.Write(&buf, binary.LittleEndian, frequency);

    resultBytes, err := device.device.Set(uint8(FunctionMorseCode), buf.Bytes())
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

// The Piezo Speaker Bricklet can play 512 different tones. This function
	// plays each tone and measures the exact frequency back. The result is a
	// mapping between setting value and frequency. This mapping is stored
	// in the EEPROM and loaded on startup.
	// 
	// The Bricklet should come calibrated, you only need to call this
	// function (once) every time you reflash the Bricklet plugin.
	// 
	// Returns *true* after the calibration finishes.
func (device *PiezoSpeakerBricklet) Calibrate() (calibration bool, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionCalibrate), buf.Bytes())
    if err != nil {
        return calibration, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return calibration, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &calibration)

    }
    
    return calibration, nil
}

// Returns the UID, the UID where the Bricklet is connected to,
	// the position, the hardware and firmware version as well as the
	// device identifier.
	// 
	// The position can be 'a', 'b', 'c' or 'd'.
	// 
	// The device identifier numbers can be found `here <device_identifier>`.
	// |device_identifier_constant|
func (device *PiezoSpeakerBricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {    
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
