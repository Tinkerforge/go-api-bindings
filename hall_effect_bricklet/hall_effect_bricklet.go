/* ***********************************************************
 * This file was automatically generated on 2020-05-19.      *
 *                                                           *
 * Go Bindings Version 2.0.8                                 *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


// Detects presence of magnetic field.
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricklets/HallEffect_Bricklet_Go.html.
package hall_effect_bricklet

import (
	"encoding/binary"
	"bytes"
	"fmt"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionGetValue Function = 1
	FunctionGetEdgeCount Function = 2
	FunctionSetEdgeCountConfig Function = 3
	FunctionGetEdgeCountConfig Function = 4
	FunctionSetEdgeInterrupt Function = 5
	FunctionGetEdgeInterrupt Function = 6
	FunctionSetEdgeCountCallbackPeriod Function = 7
	FunctionGetEdgeCountCallbackPeriod Function = 8
	FunctionEdgeInterrupt Function = 9
	FunctionGetIdentity Function = 255
	FunctionCallbackEdgeCount Function = 10
)

type EdgeType = uint8

const (
	EdgeTypeRising EdgeType = 0
	EdgeTypeFalling EdgeType = 1
	EdgeTypeBoth EdgeType = 2
)

type HallEffectBricklet struct {
	device Device
}
const DeviceIdentifier = 240
const DeviceDisplayName = "Hall Effect Bricklet"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (HallEffectBricklet, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{ 2,0,0 }, uid, &internalIPCon, 0, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return HallEffectBricklet{}, err
	}
	dev.ResponseExpected[FunctionGetValue] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetEdgeCount] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetEdgeCountConfig] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetEdgeCountConfig] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetEdgeInterrupt] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetEdgeInterrupt] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetEdgeCountCallbackPeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetEdgeCountCallbackPeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionEdgeInterrupt] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetIdentity] = ResponseExpectedFlagAlwaysTrue;
	return HallEffectBricklet{dev}, nil
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
func (device *HallEffectBricklet) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *HallEffectBricklet) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *HallEffectBricklet) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *HallEffectBricklet) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// This callback is triggered periodically with the period that is set by
// SetEdgeCountCallbackPeriod. The parameters are the
// current count and the current value (see GetValue and
// GetEdgeCount).
// 
// The RegisterEdgeCountCallback callback is only triggered if the count or value changed
// since the last triggering.
func (device *HallEffectBricklet) RegisterEdgeCountCallback(fn func(uint32, bool)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 13 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var count uint32
		var value bool
		binary.Read(buf, binary.LittleEndian, &count)
		binary.Read(buf, binary.LittleEndian, &value)
		fn(count, value)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackEdgeCount), wrapper)
}

// Remove a registered Edge Count callback.
func (device *HallEffectBricklet) DeregisterEdgeCountCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackEdgeCount), registrationId)
}


// Returns *true* if a magnetic field of 3.5 millitesla or greater is detected.
func (device *HallEffectBricklet) GetValue() (value bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetValue), buf.Bytes())
	if err != nil {
		return value, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return value, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return value, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &value)

	}

	return value, nil
}

// Returns the current value of the edge counter. You can configure
// edge type (rising, falling, both) that is counted with
// SetEdgeCountConfig.
// 
// If you set the reset counter to *true*, the count is set back to 0
// directly after it is read.
func (device *HallEffectBricklet) GetEdgeCount(resetCounter bool) (count uint32, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, resetCounter);

	resultBytes, err := device.device.Get(uint8(FunctionGetEdgeCount), buf.Bytes())
	if err != nil {
		return count, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 12 {
			return count, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		if header.ErrorCode != 0 {
			return count, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &count)

	}

	return count, nil
}

// The edge type parameter configures if rising edges, falling edges or
// both are counted. Possible edge types are:
// 
// * 0 = rising
// * 1 = falling
// * 2 = both
// 
// A magnetic field of 3.5 millitesla or greater causes a falling edge and a
// magnetic field of 2.5 millitesla or smaller causes a rising edge.
// 
// If a magnet comes near the Bricklet the signal goes low (falling edge), if
// a magnet is removed from the vicinity the signal goes high (rising edge).
// 
// Configuring an edge counter resets its value to 0.
// 
// If you don't know what any of this means, just leave it at default. The
// default configuration is very likely OK for you.
//
// Associated constants:
//
//	* EdgeTypeRising
//	* EdgeTypeFalling
//	* EdgeTypeBoth
func (device *HallEffectBricklet) SetEdgeCountConfig(edgeType EdgeType, debounce uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, edgeType);
	binary.Write(&buf, binary.LittleEndian, debounce);

	resultBytes, err := device.device.Set(uint8(FunctionSetEdgeCountConfig), buf.Bytes())
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

// Returns the edge type and debounce time as set by SetEdgeCountConfig.
//
// Associated constants:
//
//	* EdgeTypeRising
//	* EdgeTypeFalling
//	* EdgeTypeBoth
func (device *HallEffectBricklet) GetEdgeCountConfig() (edgeType EdgeType, debounce uint8, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetEdgeCountConfig), buf.Bytes())
	if err != nil {
		return edgeType, debounce, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 10 {
			return edgeType, debounce, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		if header.ErrorCode != 0 {
			return edgeType, debounce, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &edgeType)
		binary.Read(resultBuf, binary.LittleEndian, &debounce)

	}

	return edgeType, debounce, nil
}

// Sets the number of edges until an interrupt is invoked.
// 
// If *edges* is set to n, an interrupt is invoked for every n-th detected edge.
// 
// If *edges* is set to 0, the interrupt is disabled.
func (device *HallEffectBricklet) SetEdgeInterrupt(edges uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, edges);

	resultBytes, err := device.device.Set(uint8(FunctionSetEdgeInterrupt), buf.Bytes())
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

// Returns the edges as set by SetEdgeInterrupt.
func (device *HallEffectBricklet) GetEdgeInterrupt() (edges uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetEdgeInterrupt), buf.Bytes())
	if err != nil {
		return edges, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 12 {
			return edges, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		if header.ErrorCode != 0 {
			return edges, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &edges)

	}

	return edges, nil
}

// Sets the period with which the RegisterEdgeCountCallback callback is triggered
// periodically. A value of 0 turns the callback off.
// 
// The RegisterEdgeCountCallback callback is only triggered if the edge count has changed
// since the last triggering.
func (device *HallEffectBricklet) SetEdgeCountCallbackPeriod(period uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);

	resultBytes, err := device.device.Set(uint8(FunctionSetEdgeCountCallbackPeriod), buf.Bytes())
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

// Returns the period as set by SetEdgeCountCallbackPeriod.
func (device *HallEffectBricklet) GetEdgeCountCallbackPeriod() (period uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetEdgeCountCallbackPeriod), buf.Bytes())
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

// This callback is triggered every n-th count, as configured with
// SetEdgeInterrupt. The parameters are the
// current count and the current value (see GetValue and
// GetEdgeCount).
func (device *HallEffectBricklet) EdgeInterrupt() (count uint32, value bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionEdgeInterrupt), buf.Bytes())
	if err != nil {
		return count, value, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 13 {
			return count, value, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 13)
		}

		if header.ErrorCode != 0 {
			return count, value, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &count)
		binary.Read(resultBuf, binary.LittleEndian, &value)

	}

	return count, value, nil
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
func (device *HallEffectBricklet) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
