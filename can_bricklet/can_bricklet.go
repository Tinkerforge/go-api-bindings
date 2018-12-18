/* ***********************************************************
 * This file was automatically generated on 2018-12-18.      *
 *                                                           *
 * Go Bindings Version 2.0.0                                 *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


//Communicates with CAN bus devices.
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/CAN_Bricklet_Go.html.
package can_bricklet

import (
	"encoding/binary"
	"bytes"
    . "github.com/tinkerforge/go-api-bindings/internal"
    "github.com/tinkerforge/go-api-bindings/ipconnection"
)

type Function uint8

const (
    FunctionWriteFrame Function = 1
	FunctionReadFrame Function = 2
	FunctionEnableFrameReadCallback Function = 3
	FunctionDisableFrameReadCallback Function = 4
	FunctionIsFrameReadCallbackEnabled Function = 5
	FunctionSetConfiguration Function = 6
	FunctionGetConfiguration Function = 7
	FunctionSetReadFilter Function = 8
	FunctionGetReadFilter Function = 9
	FunctionGetErrorLog Function = 10
	FunctionGetIdentity Function = 255
	FunctionCallbackFrameRead Function = 11
)

type FrameType uint8

const (
    FrameTypeStandardData FrameType = 0
	FrameTypeStandardRemote FrameType = 1
	FrameTypeExtendedData FrameType = 2
	FrameTypeExtendedRemote FrameType = 3
)

type BaudRate uint8

const (
    BaudRate10kbps BaudRate = 0
	BaudRate20kbps BaudRate = 1
	BaudRate50kbps BaudRate = 2
	BaudRate125kbps BaudRate = 3
	BaudRate250kbps BaudRate = 4
	BaudRate500kbps BaudRate = 5
	BaudRate800kbps BaudRate = 6
	BaudRate1000kbps BaudRate = 7
)

type TransceiverMode uint8

const (
    TransceiverModeNormal TransceiverMode = 0
	TransceiverModeLoopback TransceiverMode = 1
	TransceiverModeReadOnly TransceiverMode = 2
)

type FilterMode uint8

const (
    FilterModeDisabled FilterMode = 0
	FilterModeAcceptAll FilterMode = 1
	FilterModeMatchStandard FilterMode = 2
	FilterModeMatchStandardAndData FilterMode = 3
	FilterModeMatchExtended FilterMode = 4
)

type CANBricklet struct{
	device Device
}
const DeviceIdentifier = 270
const DeviceDisplayName = "CAN Bricklet"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (CANBricklet, error) {
    internalIPCon := ipcon.GetInternalHandle().(IPConnection)
    dev, err := NewDevice([3]uint8{ 2,0,0 }, uid, &internalIPCon, 0)
    if err != nil {
        return CANBricklet{}, err
    }
    dev.ResponseExpected[FunctionWriteFrame] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionReadFrame] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionEnableFrameReadCallback] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionDisableFrameReadCallback] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionIsFrameReadCallbackEnabled] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetConfiguration] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetReadFilter] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetReadFilter] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetErrorLog] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetIdentity] = ResponseExpectedFlagAlwaysTrue;
    return CANBricklet{dev}, nil
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
func (device *CANBricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *CANBricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
    return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *CANBricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *CANBricklet) GetAPIVersion() [3]uint8 {
    return device.device.GetAPIVersion()
}

// This callback is triggered if a data or remote frame was received by the CAN
	// transceiver.
	// 
	// The ``identifier`` return value follows the identifier format described for
	// WriteFrame.
	// 
	// For remote frames the ``data`` return value always contains invalid values.
	// 
	// A configurable read filter can be used to define which frames should be
	// received by the CAN transceiver at all (see SetReadFilter).
	// 
	// To enable this callback, use EnableFrameReadCallback.
func (device *CANBricklet) RegisterFrameReadCallback(fn func(FrameType, uint32, [8]uint8, uint8)) uint64 {
            wrapper := func(byteSlice []byte) {
                buf := bytes.NewBuffer(byteSlice[8:])
                var frameType FrameType
var identifier uint32
var data [8]uint8
var length uint8
                binary.Read(buf, binary.LittleEndian, &frameType)
binary.Read(buf, binary.LittleEndian, &identifier)
binary.Read(buf, binary.LittleEndian, &data)
binary.Read(buf, binary.LittleEndian, &length)
                fn(frameType, identifier, data, length)
            }
    return device.device.RegisterCallback(uint8(FunctionCallbackFrameRead), wrapper)
}

//Remove a registered Frame Read callback.
func (device *CANBricklet) DeregisterFrameReadCallback(callbackID uint64) {
    device.device.DeregisterCallback(uint8(FunctionCallbackFrameRead), callbackID)
}


// Writes a data or remote frame to the write buffer to be transmitted over the
	// CAN transceiver.
	// 
	// The Bricklet supports the standard 11-bit (CAN 2.0A) and the additional extended
	// 18-bit (CAN 2.0B) identifiers. For standard frames the Bricklet uses bit 0 to 10
	// from the ``identifier`` parameter as standard 11-bit identifier. For extended
	// frames the Bricklet additionally uses bit 11 to 28 from the ``identifier``
	// parameter as extended 18-bit identifier.
	// 
	// For remote frames the ``data`` parameter is ignored.
	// 
	// Returns *true* if the frame was successfully added to the write buffer. Returns
	// *false* if the frame could not be added because write buffer is already full.
	// 
	// The write buffer can overflow if frames are written to it at a higher rate
	// than the Bricklet can transmitted them over the CAN transceiver. This may
	// happen if the CAN transceiver is configured as read-only or is using a low baud
	// rate (see SetConfiguration). It can also happen if the CAN bus is
	// congested and the frame cannot be transmitted because it constantly loses
	// arbitration or because the CAN transceiver is currently disabled due to a high
	// write error level (see GetErrorLog).
//
// Associated constants:
//
//	* FrameTypeStandardData
//	* FrameTypeStandardRemote
//	* FrameTypeExtendedData
//	* FrameTypeExtendedRemote
func (device *CANBricklet) WriteFrame(frameType FrameType, identifier uint32, data [8]uint8, length uint8) (success bool, err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, frameType);
	binary.Write(&buf, binary.LittleEndian, identifier);
	binary.Write(&buf, binary.LittleEndian, data);
	binary.Write(&buf, binary.LittleEndian, length);

    resultBytes, err := device.device.Get(uint8(FunctionWriteFrame), buf.Bytes())
    if err != nil {
        return success, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return success, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &success)

    }
    
    return success, nil
}

// Tries to read the next data or remote frame from the read buffer and return it.
	// If a frame was successfully read, then the ``success`` return value is set to
	// *true* and the other return values contain the frame. If the read buffer is
	// empty and no frame could be read, then the ``success`` return value is set to
	// *false* and the other return values contain invalid data.
	// 
	// The ``identifier`` return value follows the identifier format described for
	// WriteFrame.
	// 
	// For remote frames the ``data`` return value always contains invalid data.
	// 
	// A configurable read filter can be used to define which frames should be
	// received by the CAN transceiver and put into the read buffer (see
	// SetReadFilter).
	// 
	// Instead of polling with this function, you can also use callbacks. See the
	// EnableFrameReadCallback function and the RegisterFrameReadCallback callback.
//
// Associated constants:
//
//	* FrameTypeStandardData
//	* FrameTypeStandardRemote
//	* FrameTypeExtendedData
//	* FrameTypeExtendedRemote
func (device *CANBricklet) ReadFrame() (success bool, frameType FrameType, identifier uint32, data [8]uint8, length uint8, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionReadFrame), buf.Bytes())
    if err != nil {
        return success, frameType, identifier, data, length, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return success, frameType, identifier, data, length, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &success)
	binary.Read(resultBuf, binary.LittleEndian, &frameType)
	binary.Read(resultBuf, binary.LittleEndian, &identifier)
	binary.Read(resultBuf, binary.LittleEndian, &data)
	binary.Read(resultBuf, binary.LittleEndian, &length)

    }
    
    return success, frameType, identifier, data, length, nil
}

// Enables the RegisterFrameReadCallback callback.
	// 
	// By default the callback is disabled.
func (device *CANBricklet) EnableFrameReadCallback() (err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Set(uint8(FunctionEnableFrameReadCallback), buf.Bytes())
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

// Disables the RegisterFrameReadCallback callback.
	// 
	// By default the callback is disabled.
func (device *CANBricklet) DisableFrameReadCallback() (err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Set(uint8(FunctionDisableFrameReadCallback), buf.Bytes())
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

// Returns *true* if the RegisterFrameReadCallback callback is enabled, *false* otherwise.
func (device *CANBricklet) IsFrameReadCallbackEnabled() (enabled bool, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionIsFrameReadCallbackEnabled), buf.Bytes())
    if err != nil {
        return enabled, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return enabled, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &enabled)

    }
    
    return enabled, nil
}

// Sets the configuration for the CAN bus communication.
	// 
	// The baud rate can be configured in steps between 10 and 1000 kbit/s.
	// 
	// The CAN transceiver has three different modes:
	// 
	// * Normal: Reads from and writes to the CAN bus and performs active bus
	//   error detection and acknowledgement.
	// * Loopback: All reads and writes are performed internally. The transceiver
	//   is disconnected from the actual CAN bus.
	// * Read-Only: Only reads from the CAN bus, but does neither active bus error
	//   detection nor acknowledgement. Only the receiving part of the transceiver
	//   is connected to the CAN bus.
	// 
	// The write timeout has three different modes that define how a failed frame
	// transmission should be handled:
	// 
	// * One-Shot (< 0): Only one transmission attempt will be made. If the
	//   transmission fails then the frame is discarded.
	// * Infinite (= 0): Infinite transmission attempts will be made. The frame will
	//   never be discarded.
	// * Milliseconds (> 0): A limited number of transmission attempts will be made.
	//   If the frame could not be transmitted successfully after the configured
	//   number of milliseconds then the frame is discarded.
	// 
	// The default is: 125 kbit/s, normal transceiver mode and infinite write timeout.
//
// Associated constants:
//
//	* BaudRate10kbps
//	* BaudRate20kbps
//	* BaudRate50kbps
//	* BaudRate125kbps
//	* BaudRate250kbps
//	* BaudRate500kbps
//	* BaudRate800kbps
//	* BaudRate1000kbps
//	* TransceiverModeNormal
//	* TransceiverModeLoopback
//	* TransceiverModeReadOnly
func (device *CANBricklet) SetConfiguration(baudRate BaudRate, transceiverMode TransceiverMode, writeTimeout int32) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, baudRate);
	binary.Write(&buf, binary.LittleEndian, transceiverMode);
	binary.Write(&buf, binary.LittleEndian, writeTimeout);

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
//	* BaudRate10kbps
//	* BaudRate20kbps
//	* BaudRate50kbps
//	* BaudRate125kbps
//	* BaudRate250kbps
//	* BaudRate500kbps
//	* BaudRate800kbps
//	* BaudRate1000kbps
//	* TransceiverModeNormal
//	* TransceiverModeLoopback
//	* TransceiverModeReadOnly
func (device *CANBricklet) GetConfiguration() (baudRate BaudRate, transceiverMode TransceiverMode, writeTimeout int32, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetConfiguration), buf.Bytes())
    if err != nil {
        return baudRate, transceiverMode, writeTimeout, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return baudRate, transceiverMode, writeTimeout, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &baudRate)
	binary.Read(resultBuf, binary.LittleEndian, &transceiverMode)
	binary.Read(resultBuf, binary.LittleEndian, &writeTimeout)

    }
    
    return baudRate, transceiverMode, writeTimeout, nil
}

// Set the read filter configuration. This can be used to define which frames
	// should be received by the CAN transceiver and put into the read buffer.
	// 
	// The read filter has five different modes that define if and how the mask and
	// the two filters are applied:
	// 
	// * Disabled: No filtering is applied at all. All frames are received even
	//   incomplete and defective frames. This mode should be used for debugging only.
	// * Accept-All: All complete and error-free frames are received.
	// * Match-Standard: Only standard frames with a matching identifier are received.
	// * Match-Standard-and-Data: Only standard frames with matching identifier and
	//   data bytes are received.
	// * Match-Extended: Only extended frames with a matching identifier are received.
	// 
	// The mask and filters are used as bit masks. Their usage depends on the mode:
	// 
	// * Disabled: Mask and filters are ignored.
	// * Accept-All: Mask and filters are ignored.
	// * Match-Standard: Bit 0 to 10 (11 bits) of mask and filters are used to match
	//   the 11-bit identifier of standard frames.
	// * Match-Standard-and-Data: Bit 0 to 10 (11 bits) of mask and filters are used
	//   to match the 11-bit identifier of standard frames. Bit 11 to 18 (8 bits) and
	//   bit 19 to 26 (8 bits) of mask and filters are used to match the first and
	//   second data byte (if present) of standard frames.
	// * Match-Extended: Bit 0 to 10 (11 bits) of mask and filters are used
	//   to match the standard 11-bit identifier part of extended frames. Bit 11 to 28
	//   (18 bits) of mask and filters are used to match the extended 18-bit identifier
	//   part of extended frames.
	// 
	// The mask and filters are applied in this way: The mask is used to select the
	// identifier and data bits that should be compared to the corresponding filter
	// bits. All unselected bits are automatically accepted. All selected bits have
	// to match one of the filters to be accepted. If all bits for the selected mode
	// are accepted then the frame is accepted and is added to the read buffer.
	// 
	//  Mask Bit| Filter Bit| Identifier/Data Bit| Result
	//  --- | --- | --- | --- 
	//  0| X| X| Accept
	//  1| 0| 0| Accept
	//  1| 0| 1| Reject
	//  1| 1| 0| Reject
	//  1| 1| 1| Accept
	// 
	// For example, to receive standard frames with identifier 0x123 only the mode can
	// be set to Match-Standard with 0x7FF as mask and 0x123 as filter 1 and filter 2.
	// The mask of 0x7FF selects all 11 identifier bits for matching so that the
	// identifier has to be exactly 0x123 to be accepted.
	// 
	// To accept identifier 0x123 and identifier 0x456 at the same time, just set
	// filter 2 to 0x456 and keep mask and filter 1 unchanged.
	// 
	// The default mode is accept-all.
//
// Associated constants:
//
//	* FilterModeDisabled
//	* FilterModeAcceptAll
//	* FilterModeMatchStandard
//	* FilterModeMatchStandardAndData
//	* FilterModeMatchExtended
func (device *CANBricklet) SetReadFilter(mode FilterMode, mask uint32, filter1 uint32, filter2 uint32) (err error) {    
        var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, mode);
	binary.Write(&buf, binary.LittleEndian, mask);
	binary.Write(&buf, binary.LittleEndian, filter1);
	binary.Write(&buf, binary.LittleEndian, filter2);

    resultBytes, err := device.device.Set(uint8(FunctionSetReadFilter), buf.Bytes())
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

// Returns the read filter as set by SetReadFilter.
//
// Associated constants:
//
//	* FilterModeDisabled
//	* FilterModeAcceptAll
//	* FilterModeMatchStandard
//	* FilterModeMatchStandardAndData
//	* FilterModeMatchExtended
func (device *CANBricklet) GetReadFilter() (mode FilterMode, mask uint32, filter1 uint32, filter2 uint32, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetReadFilter), buf.Bytes())
    if err != nil {
        return mode, mask, filter1, filter2, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return mode, mask, filter1, filter2, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &mode)
	binary.Read(resultBuf, binary.LittleEndian, &mask)
	binary.Read(resultBuf, binary.LittleEndian, &filter1)
	binary.Read(resultBuf, binary.LittleEndian, &filter2)

    }
    
    return mode, mask, filter1, filter2, nil
}

// Returns information about different kinds of errors.
	// 
	// The write and read error levels indicate the current level of checksum,
	// acknowledgement, form, bit and stuffing errors during CAN bus write and read
	// operations.
	// 
	// When the write error level extends 255 then the CAN transceiver gets disabled
	// and no frames can be transmitted or received anymore. The CAN transceiver will
	// automatically be activated again after the CAN bus is idle for a while.
	// 
	// The write and read error levels are not available in read-only transceiver mode
	// (see SetConfiguration) and are reset to 0 as a side effect of changing
	// the configuration or the read filter.
	// 
	// The write timeout, read register and buffer overflow counts represents the
	// number of these errors:
	// 
	// * A write timeout occurs if a frame could not be transmitted before the
	//   configured write timeout expired (see SetConfiguration).
	// * A read register overflow occurs if the read register of the CAN transceiver
	//   still contains the last received frame when the next frame arrives. In this
	//   case the newly arrived frame is lost. This happens if the CAN transceiver
	//   receives more frames than the Bricklet can handle. Using the read filter
	//   (see SetReadFilter) can help to reduce the amount of received frames.
	//   This count is not exact, but a lower bound, because the Bricklet might not
	//   able detect all overflows if they occur in rapid succession.
	// * A read buffer overflow occurs if the read buffer of the Bricklet is already
	//   full when the next frame should be read from the read register of the CAN
	//   transceiver. In this case the frame in the read register is lost. This
	//   happens if the CAN transceiver receives more frames to be added to the read
	//   buffer than are removed from the read buffer using the ReadFrame
	//   function. Using the RegisterFrameReadCallback callback ensures that the read buffer
	//   can not overflow.
func (device *CANBricklet) GetErrorLog() (writeErrorLevel uint8, readErrorLevel uint8, transceiverDisabled bool, writeTimeoutCount uint32, readRegisterOverflowCount uint32, readBufferOverflowCount uint32, err error) {    
        var buf bytes.Buffer
    
    resultBytes, err := device.device.Get(uint8(FunctionGetErrorLog), buf.Bytes())
    if err != nil {
        return writeErrorLevel, readErrorLevel, transceiverDisabled, writeTimeoutCount, readRegisterOverflowCount, readBufferOverflowCount, err
    }
    if len(resultBytes) > 0 {
        var header PacketHeader
        
        header.FillFromBytes(resultBytes)
        if header.ErrorCode != 0 {
            return writeErrorLevel, readErrorLevel, transceiverDisabled, writeTimeoutCount, readRegisterOverflowCount, readBufferOverflowCount, BrickletError(header.ErrorCode)
        }

        resultBuf := bytes.NewBuffer(resultBytes[8:])
        binary.Read(resultBuf, binary.LittleEndian, &writeErrorLevel)
	binary.Read(resultBuf, binary.LittleEndian, &readErrorLevel)
	binary.Read(resultBuf, binary.LittleEndian, &transceiverDisabled)
	binary.Read(resultBuf, binary.LittleEndian, &writeTimeoutCount)
	binary.Read(resultBuf, binary.LittleEndian, &readRegisterOverflowCount)
	binary.Read(resultBuf, binary.LittleEndian, &readBufferOverflowCount)

    }
    
    return writeErrorLevel, readErrorLevel, transceiverDisabled, writeTimeoutCount, readRegisterOverflowCount, readBufferOverflowCount, nil
}

// Returns the UID, the UID where the Bricklet is connected to,
	// the position, the hardware and firmware version as well as the
	// device identifier.
	// 
	// The position can be 'a', 'b', 'c' or 'd'.
	// 
	// The device identifier numbers can be found `here <device_identifier>`.
	// |device_identifier_constant|
func (device *CANBricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {    
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
