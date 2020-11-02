/* ***********************************************************
 * This file was automatically generated on 2020-11-02.      *
 *                                                           *
 * Go Bindings Version 2.0.9                                 *
 *                                                           *
 * If you have a bugfix for this file and want to commit it, *
 * please fix the bug in the generator. You can find a link  *
 * to the generators git repository on tinkerforge.com       *
 *************************************************************/


// Basis to build stacks and has 4 Bricklet ports.
// 
// 
// See also the documentation here: https://www.tinkerforge.com/en/doc/Software/Bricks/Master_Brick_Go.html.
package master_brick

import (
	"encoding/binary"
	"bytes"
	"fmt"
	. "github.com/Tinkerforge/go-api-bindings/internal"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
)

type Function = uint8

const (
	FunctionGetStackVoltage Function = 1
	FunctionGetStackCurrent Function = 2
	FunctionSetExtensionType Function = 3
	FunctionGetExtensionType Function = 4
	FunctionIsChibiPresent Function = 5
	FunctionSetChibiAddress Function = 6
	FunctionGetChibiAddress Function = 7
	FunctionSetChibiMasterAddress Function = 8
	FunctionGetChibiMasterAddress Function = 9
	FunctionSetChibiSlaveAddress Function = 10
	FunctionGetChibiSlaveAddress Function = 11
	FunctionGetChibiSignalStrength Function = 12
	FunctionGetChibiErrorLog Function = 13
	FunctionSetChibiFrequency Function = 14
	FunctionGetChibiFrequency Function = 15
	FunctionSetChibiChannel Function = 16
	FunctionGetChibiChannel Function = 17
	FunctionIsRS485Present Function = 18
	FunctionSetRS485Address Function = 19
	FunctionGetRS485Address Function = 20
	FunctionSetRS485SlaveAddress Function = 21
	FunctionGetRS485SlaveAddress Function = 22
	FunctionGetRS485ErrorLog Function = 23
	FunctionSetRS485Configuration Function = 24
	FunctionGetRS485Configuration Function = 25
	FunctionIsWifiPresent Function = 26
	FunctionSetWifiConfiguration Function = 27
	FunctionGetWifiConfiguration Function = 28
	FunctionSetWifiEncryption Function = 29
	FunctionGetWifiEncryption Function = 30
	FunctionGetWifiStatus Function = 31
	FunctionRefreshWifiStatus Function = 32
	FunctionSetWifiCertificate Function = 33
	FunctionGetWifiCertificate Function = 34
	FunctionSetWifiPowerMode Function = 35
	FunctionGetWifiPowerMode Function = 36
	FunctionGetWifiBufferInfo Function = 37
	FunctionSetWifiRegulatoryDomain Function = 38
	FunctionGetWifiRegulatoryDomain Function = 39
	FunctionGetUSBVoltage Function = 40
	FunctionSetLongWifiKey Function = 41
	FunctionGetLongWifiKey Function = 42
	FunctionSetWifiHostname Function = 43
	FunctionGetWifiHostname Function = 44
	FunctionSetStackCurrentCallbackPeriod Function = 45
	FunctionGetStackCurrentCallbackPeriod Function = 46
	FunctionSetStackVoltageCallbackPeriod Function = 47
	FunctionGetStackVoltageCallbackPeriod Function = 48
	FunctionSetUSBVoltageCallbackPeriod Function = 49
	FunctionGetUSBVoltageCallbackPeriod Function = 50
	FunctionSetStackCurrentCallbackThreshold Function = 51
	FunctionGetStackCurrentCallbackThreshold Function = 52
	FunctionSetStackVoltageCallbackThreshold Function = 53
	FunctionGetStackVoltageCallbackThreshold Function = 54
	FunctionSetUSBVoltageCallbackThreshold Function = 55
	FunctionGetUSBVoltageCallbackThreshold Function = 56
	FunctionSetDebouncePeriod Function = 57
	FunctionGetDebouncePeriod Function = 58
	FunctionIsEthernetPresent Function = 65
	FunctionSetEthernetConfiguration Function = 66
	FunctionGetEthernetConfiguration Function = 67
	FunctionGetEthernetStatus Function = 68
	FunctionSetEthernetHostname Function = 69
	FunctionSetEthernetMACAddress Function = 70
	FunctionSetEthernetWebsocketConfiguration Function = 71
	FunctionGetEthernetWebsocketConfiguration Function = 72
	FunctionSetEthernetAuthenticationSecret Function = 73
	FunctionGetEthernetAuthenticationSecret Function = 74
	FunctionSetWifiAuthenticationSecret Function = 75
	FunctionGetWifiAuthenticationSecret Function = 76
	FunctionGetConnectionType Function = 77
	FunctionIsWifi2Present Function = 78
	FunctionStartWifi2Bootloader Function = 79
	FunctionWriteWifi2SerialPort Function = 80
	FunctionReadWifi2SerialPort Function = 81
	FunctionSetWifi2AuthenticationSecret Function = 82
	FunctionGetWifi2AuthenticationSecret Function = 83
	FunctionSetWifi2Configuration Function = 84
	FunctionGetWifi2Configuration Function = 85
	FunctionGetWifi2Status Function = 86
	FunctionSetWifi2ClientConfiguration Function = 87
	FunctionGetWifi2ClientConfiguration Function = 88
	FunctionSetWifi2ClientHostname Function = 89
	FunctionGetWifi2ClientHostname Function = 90
	FunctionSetWifi2ClientPassword Function = 91
	FunctionGetWifi2ClientPassword Function = 92
	FunctionSetWifi2APConfiguration Function = 93
	FunctionGetWifi2APConfiguration Function = 94
	FunctionSetWifi2APPassword Function = 95
	FunctionGetWifi2APPassword Function = 96
	FunctionSaveWifi2Configuration Function = 97
	FunctionGetWifi2FirmwareVersion Function = 98
	FunctionEnableWifi2StatusLED Function = 99
	FunctionDisableWifi2StatusLED Function = 100
	FunctionIsWifi2StatusLEDEnabled Function = 101
	FunctionSetWifi2MeshConfiguration Function = 102
	FunctionGetWifi2MeshConfiguration Function = 103
	FunctionSetWifi2MeshRouterSSID Function = 104
	FunctionGetWifi2MeshRouterSSID Function = 105
	FunctionSetWifi2MeshRouterPassword Function = 106
	FunctionGetWifi2MeshRouterPassword Function = 107
	FunctionGetWifi2MeshCommonStatus Function = 108
	FunctionGetWifi2MeshClientStatus Function = 109
	FunctionGetWifi2MeshAPStatus Function = 110
	FunctionSetBrickletXMCFlashConfig Function = 111
	FunctionSetBrickletXMCFlashData Function = 112
	FunctionSetBrickletsEnabled Function = 113
	FunctionGetBrickletsEnabled Function = 114
	FunctionSetSPITFPBaudrateConfig Function = 231
	FunctionGetSPITFPBaudrateConfig Function = 232
	FunctionGetSendTimeoutCount Function = 233
	FunctionSetSPITFPBaudrate Function = 234
	FunctionGetSPITFPBaudrate Function = 235
	FunctionGetSPITFPErrorCount Function = 237
	FunctionEnableStatusLED Function = 238
	FunctionDisableStatusLED Function = 239
	FunctionIsStatusLEDEnabled Function = 240
	FunctionGetProtocol1BrickletName Function = 241
	FunctionGetChipTemperature Function = 242
	FunctionReset Function = 243
	FunctionWriteBrickletPlugin Function = 246
	FunctionReadBrickletPlugin Function = 247
	FunctionGetIdentity Function = 255
	FunctionCallbackStackCurrent Function = 59
	FunctionCallbackStackVoltage Function = 60
	FunctionCallbackUSBVoltage Function = 61
	FunctionCallbackStackCurrentReached Function = 62
	FunctionCallbackStackVoltageReached Function = 63
	FunctionCallbackUSBVoltageReached Function = 64
)

type ExtensionType = uint32

const (
	ExtensionTypeChibi ExtensionType = 1
	ExtensionTypeRS485 ExtensionType = 2
	ExtensionTypeWifi ExtensionType = 3
	ExtensionTypeEthernet ExtensionType = 4
	ExtensionTypeWifi2 ExtensionType = 5
)

type ChibiFrequency = uint8

const (
	ChibiFrequencyOQPSK868MHz ChibiFrequency = 0
	ChibiFrequencyOQPSK915MHz ChibiFrequency = 1
	ChibiFrequencyOQPSK780MHz ChibiFrequency = 2
	//Deprecated: Use BPSK40_915MHz instead.
	ChibiFrequencyBPSK40915MHz ChibiFrequency = 3
	ChibiFrequencyBPSK40_915MHz ChibiFrequency = 3
)

type RS485Parity = rune

const (
	RS485ParityNone RS485Parity = 'n'
	RS485ParityEven RS485Parity = 'e'
	RS485ParityOdd RS485Parity = 'o'
)

type WifiConnection = uint8

const (
	WifiConnectionDHCP WifiConnection = 0
	WifiConnectionStaticIP WifiConnection = 1
	WifiConnectionAccessPointDHCP WifiConnection = 2
	WifiConnectionAccessPointStaticIP WifiConnection = 3
	WifiConnectionAdHocDHCP WifiConnection = 4
	WifiConnectionAdHocStaticIP WifiConnection = 5
)

type WifiEncryption = uint8

const (
	WifiEncryptionWPAWPA2 WifiEncryption = 0
	WifiEncryptionWPAEnterprise WifiEncryption = 1
	WifiEncryptionWEP WifiEncryption = 2
	WifiEncryptionNoEncryption WifiEncryption = 3
)

type WifiEAPOption = uint8

const (
	WifiEAPOptionOuterAuthEAPFAST WifiEAPOption = 0
	WifiEAPOptionOuterAuthEAPTLS WifiEAPOption = 1
	WifiEAPOptionOuterAuthEAPTTLS WifiEAPOption = 2
	WifiEAPOptionOuterAuthEAPPEAP WifiEAPOption = 3
	WifiEAPOptionInnerAuthEAPMSCHAP WifiEAPOption = 0
	WifiEAPOptionInnerAuthEAPGTC WifiEAPOption = 4
	WifiEAPOptionCertTypeCACert WifiEAPOption = 0
	WifiEAPOptionCertTypeClientCert WifiEAPOption = 8
	WifiEAPOptionCertTypePrivateKey WifiEAPOption = 16
)

type WifiState = uint8

const (
	WifiStateDisassociated WifiState = 0
	WifiStateAssociated WifiState = 1
	WifiStateAssociating WifiState = 2
	WifiStateError WifiState = 3
	WifiStateNotInitializedYet WifiState = 255
)

type WifiPowerMode = uint8

const (
	WifiPowerModeFullSpeed WifiPowerMode = 0
	WifiPowerModeLowPower WifiPowerMode = 1
)

type WifiDomain = uint8

const (
	WifiDomainChannel1To11 WifiDomain = 0
	WifiDomainChannel1To13 WifiDomain = 1
	WifiDomainChannel1To14 WifiDomain = 2
)

type ThresholdOption = rune

const (
	ThresholdOptionOff ThresholdOption = 'x'
	ThresholdOptionOutside ThresholdOption = 'o'
	ThresholdOptionInside ThresholdOption = 'i'
	ThresholdOptionSmaller ThresholdOption = '<'
	ThresholdOptionGreater ThresholdOption = '>'
)

type EthernetConnection = uint8

const (
	EthernetConnectionDHCP EthernetConnection = 0
	EthernetConnectionStaticIP EthernetConnection = 1
)

type ConnectionType = uint8

const (
	ConnectionTypeNone ConnectionType = 0
	ConnectionTypeUSB ConnectionType = 1
	ConnectionTypeSPIStack ConnectionType = 2
	ConnectionTypeChibi ConnectionType = 3
	ConnectionTypeRS485 ConnectionType = 4
	ConnectionTypeWifi ConnectionType = 5
	ConnectionTypeEthernet ConnectionType = 6
	ConnectionTypeWifi2 ConnectionType = 7
)

type Wifi2PHYMode = uint8

const (
	Wifi2PHYModeB Wifi2PHYMode = 0
	Wifi2PHYModeG Wifi2PHYMode = 1
	Wifi2PHYModeN Wifi2PHYMode = 2
)

type Wifi2ClientStatus = uint8

const (
	Wifi2ClientStatusIdle Wifi2ClientStatus = 0
	Wifi2ClientStatusConnecting Wifi2ClientStatus = 1
	Wifi2ClientStatusWrongPassword Wifi2ClientStatus = 2
	Wifi2ClientStatusNoAPFound Wifi2ClientStatus = 3
	Wifi2ClientStatusConnectFailed Wifi2ClientStatus = 4
	Wifi2ClientStatusGotIP Wifi2ClientStatus = 5
	Wifi2ClientStatusUnknown Wifi2ClientStatus = 255
)

type Wifi2APEncryption = uint8

const (
	Wifi2APEncryptionOpen Wifi2APEncryption = 0
	Wifi2APEncryptionWEP Wifi2APEncryption = 1
	Wifi2APEncryptionWPAPSK Wifi2APEncryption = 2
	Wifi2APEncryptionWPA2PSK Wifi2APEncryption = 3
	Wifi2APEncryptionWPAWPA2PSK Wifi2APEncryption = 4
)

type Wifi2MeshStatus = uint8

const (
	Wifi2MeshStatusDisabled Wifi2MeshStatus = 0
	Wifi2MeshStatusWIFIConnecting Wifi2MeshStatus = 1
	Wifi2MeshStatusGotIP Wifi2MeshStatus = 2
	Wifi2MeshStatusMeshLocal Wifi2MeshStatus = 3
	Wifi2MeshStatusMeshOnline Wifi2MeshStatus = 4
	Wifi2MeshStatusAPAvailable Wifi2MeshStatus = 5
	Wifi2MeshStatusAPSetup Wifi2MeshStatus = 6
	Wifi2MeshStatusLeafAvailable Wifi2MeshStatus = 7
)

type CommunicationMethod = uint8

const (
	CommunicationMethodNone CommunicationMethod = 0
	CommunicationMethodUSB CommunicationMethod = 1
	CommunicationMethodSPIStack CommunicationMethod = 2
	CommunicationMethodChibi CommunicationMethod = 3
	CommunicationMethodRS485 CommunicationMethod = 4
	CommunicationMethodWIFI CommunicationMethod = 5
	CommunicationMethodEthernet CommunicationMethod = 6
	CommunicationMethodWIFIV2 CommunicationMethod = 7
)

type MasterBrick struct {
	device Device
}
const DeviceIdentifier = 13
const DeviceDisplayName = "Master Brick"

// Creates an object with the unique device ID `uid`. This object can then be used after the IP Connection `ipcon` is connected.
func New(uid string, ipcon *ipconnection.IPConnection) (MasterBrick, error) {
	internalIPCon := ipcon.GetInternalHandle().(IPConnection)
	dev, err := NewDevice([3]uint8{ 2,0,10 }, uid, &internalIPCon, 0, DeviceIdentifier, DeviceDisplayName)
	if err != nil {
		return MasterBrick{}, err
	}
	dev.ResponseExpected[FunctionGetStackVoltage] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetStackCurrent] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetExtensionType] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetExtensionType] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionIsChibiPresent] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetChibiAddress] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetChibiAddress] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetChibiMasterAddress] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetChibiMasterAddress] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetChibiSlaveAddress] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetChibiSlaveAddress] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetChibiSignalStrength] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetChibiErrorLog] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetChibiFrequency] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetChibiFrequency] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetChibiChannel] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetChibiChannel] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionIsRS485Present] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetRS485Address] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetRS485Address] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetRS485SlaveAddress] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetRS485SlaveAddress] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetRS485ErrorLog] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetRS485Configuration] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetRS485Configuration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionIsWifiPresent] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetWifiConfiguration] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetWifiConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetWifiEncryption] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetWifiEncryption] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetWifiStatus] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionRefreshWifiStatus] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionSetWifiCertificate] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetWifiCertificate] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetWifiPowerMode] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetWifiPowerMode] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetWifiBufferInfo] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetWifiRegulatoryDomain] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetWifiRegulatoryDomain] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetUSBVoltage] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetLongWifiKey] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetLongWifiKey] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetWifiHostname] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetWifiHostname] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetStackCurrentCallbackPeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetStackCurrentCallbackPeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetStackVoltageCallbackPeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetStackVoltageCallbackPeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetUSBVoltageCallbackPeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetUSBVoltageCallbackPeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetStackCurrentCallbackThreshold] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetStackCurrentCallbackThreshold] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetStackVoltageCallbackThreshold] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetStackVoltageCallbackThreshold] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetUSBVoltageCallbackThreshold] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetUSBVoltageCallbackThreshold] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetDebouncePeriod] = ResponseExpectedFlagTrue;
	dev.ResponseExpected[FunctionGetDebouncePeriod] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionIsEthernetPresent] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetEthernetConfiguration] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetEthernetConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetEthernetStatus] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetEthernetHostname] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionSetEthernetMACAddress] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionSetEthernetWebsocketConfiguration] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetEthernetWebsocketConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetEthernetAuthenticationSecret] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetEthernetAuthenticationSecret] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetWifiAuthenticationSecret] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetWifiAuthenticationSecret] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetConnectionType] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionIsWifi2Present] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionStartWifi2Bootloader] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionWriteWifi2SerialPort] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionReadWifi2SerialPort] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetWifi2AuthenticationSecret] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetWifi2AuthenticationSecret] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetWifi2Configuration] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetWifi2Configuration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetWifi2Status] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetWifi2ClientConfiguration] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetWifi2ClientConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetWifi2ClientHostname] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetWifi2ClientHostname] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetWifi2ClientPassword] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetWifi2ClientPassword] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetWifi2APConfiguration] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetWifi2APConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetWifi2APPassword] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetWifi2APPassword] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSaveWifi2Configuration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetWifi2FirmwareVersion] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionEnableWifi2StatusLED] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionDisableWifi2StatusLED] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionIsWifi2StatusLEDEnabled] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetWifi2MeshConfiguration] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetWifi2MeshConfiguration] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetWifi2MeshRouterSSID] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetWifi2MeshRouterSSID] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetWifi2MeshRouterPassword] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetWifi2MeshRouterPassword] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetWifi2MeshCommonStatus] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetWifi2MeshClientStatus] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetWifi2MeshAPStatus] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetBrickletXMCFlashConfig] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetBrickletXMCFlashData] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetBrickletsEnabled] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetBrickletsEnabled] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetSPITFPBaudrateConfig] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetSPITFPBaudrateConfig] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetSendTimeoutCount] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionSetSPITFPBaudrate] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionGetSPITFPBaudrate] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetSPITFPErrorCount] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionEnableStatusLED] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionDisableStatusLED] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionIsStatusLEDEnabled] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetProtocol1BrickletName] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetChipTemperature] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionReset] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionWriteBrickletPlugin] = ResponseExpectedFlagFalse;
	dev.ResponseExpected[FunctionReadBrickletPlugin] = ResponseExpectedFlagAlwaysTrue;
	dev.ResponseExpected[FunctionGetIdentity] = ResponseExpectedFlagAlwaysTrue;
	return MasterBrick{dev}, nil
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
func (device *MasterBrick) GetResponseExpected(functionID Function) (bool, error) {
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
func (device *MasterBrick) SetResponseExpected(functionID Function, responseExpected bool) error {
	return device.device.SetResponseExpected(uint8(functionID), responseExpected)
}

// Changes the response expected flag for all setter and callback configuration functions of this device at once.
func (device *MasterBrick) SetResponseExpectedAll(responseExpected bool) {
	device.device.SetResponseExpectedAll(responseExpected)
}

// Returns the version of the API definition (major, minor, revision) implemented by this API bindings. This is neither the release version of this API bindings nor does it tell you anything about the represented Brick or Bricklet.
func (device *MasterBrick) GetAPIVersion() [3]uint8 {
	return device.device.GetAPIVersion()
}

// This callback is triggered periodically with the period that is set by
// SetStackCurrentCallbackPeriod. The parameter is the current
// of the sensor.
// 
// The RegisterStackCurrentCallback callback is only triggered if the current has changed
// since the last triggering.
// 
// .. versionadded:: 2.0.5$nbsp;(Firmware)
func (device *MasterBrick) RegisterStackCurrentCallback(fn func(uint16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 10 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var current uint16
		binary.Read(buf, binary.LittleEndian, &current)
		fn(current)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackStackCurrent), wrapper)
}

// Remove a registered Stack Current callback.
func (device *MasterBrick) DeregisterStackCurrentCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackStackCurrent), registrationId)
}


// This callback is triggered periodically with the period that is set by
// SetStackVoltageCallbackPeriod. The parameter is the voltage
// of the sensor.
// 
// The RegisterStackVoltageCallback callback is only triggered if the voltage has changed
// since the last triggering.
// 
// .. versionadded:: 2.0.5$nbsp;(Firmware)
func (device *MasterBrick) RegisterStackVoltageCallback(fn func(uint16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 10 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var voltage uint16
		binary.Read(buf, binary.LittleEndian, &voltage)
		fn(voltage)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackStackVoltage), wrapper)
}

// Remove a registered Stack Voltage callback.
func (device *MasterBrick) DeregisterStackVoltageCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackStackVoltage), registrationId)
}


// This callback is triggered periodically with the period that is set by
// SetUSBVoltageCallbackPeriod. The parameter is the USB
// voltage.
// 
// The RegisterUSBVoltageCallback callback is only triggered if the USB voltage has changed
// since the last triggering.
// 
// Does not work with hardware version 2.1.
// 
// .. versionadded:: 2.0.5$nbsp;(Firmware)
func (device *MasterBrick) RegisterUSBVoltageCallback(fn func(uint16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 10 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var voltage uint16
		binary.Read(buf, binary.LittleEndian, &voltage)
		fn(voltage)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackUSBVoltage), wrapper)
}

// Remove a registered USB Voltage callback.
func (device *MasterBrick) DeregisterUSBVoltageCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackUSBVoltage), registrationId)
}


// This callback is triggered when the threshold as set by
// SetStackCurrentCallbackThreshold is reached.
// The parameter is the stack current.
// 
// If the threshold keeps being reached, the callback is triggered periodically
// with the period as set by SetDebouncePeriod.
// 
// .. versionadded:: 2.0.5$nbsp;(Firmware)
func (device *MasterBrick) RegisterStackCurrentReachedCallback(fn func(uint16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 10 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var current uint16
		binary.Read(buf, binary.LittleEndian, &current)
		fn(current)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackStackCurrentReached), wrapper)
}

// Remove a registered Stack Current Reached callback.
func (device *MasterBrick) DeregisterStackCurrentReachedCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackStackCurrentReached), registrationId)
}


// This callback is triggered when the threshold as set by
// SetStackVoltageCallbackThreshold is reached.
// The parameter is the stack voltage.
// 
// If the threshold keeps being reached, the callback is triggered periodically
// with the period as set by SetDebouncePeriod.
// 
// .. versionadded:: 2.0.5$nbsp;(Firmware)
func (device *MasterBrick) RegisterStackVoltageReachedCallback(fn func(uint16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 10 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var voltage uint16
		binary.Read(buf, binary.LittleEndian, &voltage)
		fn(voltage)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackStackVoltageReached), wrapper)
}

// Remove a registered Stack Voltage Reached callback.
func (device *MasterBrick) DeregisterStackVoltageReachedCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackStackVoltageReached), registrationId)
}


// This callback is triggered when the threshold as set by
// SetUSBVoltageCallbackThreshold is reached.
// The parameter is the voltage of the sensor.
// 
// If the threshold keeps being reached, the callback is triggered periodically
// with the period as set by SetDebouncePeriod.
// 
// .. versionadded:: 2.0.5$nbsp;(Firmware)
func (device *MasterBrick) RegisterUSBVoltageReachedCallback(fn func(uint16)) uint64 {
	wrapper := func(byteSlice []byte) {
		var header PacketHeader

		header.FillFromBytes(byteSlice)
		if header.Length != 10 {
			return
		}
		buf := bytes.NewBuffer(byteSlice[8:])
		var voltage uint16
		binary.Read(buf, binary.LittleEndian, &voltage)
		fn(voltage)
	}
	return device.device.RegisterCallback(uint8(FunctionCallbackUSBVoltageReached), wrapper)
}

// Remove a registered USB Voltage Reached callback.
func (device *MasterBrick) DeregisterUSBVoltageReachedCallback(registrationId uint64) {
	device.device.DeregisterCallback(uint8(FunctionCallbackUSBVoltageReached), registrationId)
}


// Returns the stack voltage. The stack voltage is the
// voltage that is supplied via the stack, i.e. it is given by a
// Step-Down or Step-Up Power Supply.
// 
// Note
//  It is not possible to measure voltages supplied per PoE or USB with this function.
func (device *MasterBrick) GetStackVoltage() (voltage uint16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetStackVoltage), buf.Bytes())
	if err != nil {
		return voltage, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 10 {
			return voltage, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		if header.ErrorCode != 0 {
			return voltage, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &voltage)

	}

	return voltage, nil
}

// Returns the stack current. The stack current is the
// current that is drawn via the stack, i.e. it is given by a
// Step-Down or Step-Up Power Supply.
// 
// Note
//  It is not possible to measure the current drawn via PoE or USB with this function.
func (device *MasterBrick) GetStackCurrent() (current uint16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetStackCurrent), buf.Bytes())
	if err != nil {
		return current, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 10 {
			return current, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		if header.ErrorCode != 0 {
			return current, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &current)

	}

	return current, nil
}

// Writes the extension type to the EEPROM of a specified extension.
// The extension is either 0 or 1 (0 is the lower one, 1 is the upper one,
// if only one extension is present use 0).
// 
// Possible extension types:
// 
//  Type| Description
//  --- | --- 
//  1|    Chibi
//  2|    RS485
//  3|    WIFI
//  4|    Ethernet
//  5|    WIFI 2.0
// 
// The extension type is already set when bought and it can be set with the
// Brick Viewer, it is unlikely that you need this function.
//
// Associated constants:
//
//	* ExtensionTypeChibi
//	* ExtensionTypeRS485
//	* ExtensionTypeWifi
//	* ExtensionTypeEthernet
//	* ExtensionTypeWifi2
func (device *MasterBrick) SetExtensionType(extension uint8, exttype ExtensionType) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, extension);
	binary.Write(&buf, binary.LittleEndian, exttype);

	resultBytes, err := device.device.Set(uint8(FunctionSetExtensionType), buf.Bytes())
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

// Returns the type for a given extension as set by SetExtensionType.
//
// Associated constants:
//
//	* ExtensionTypeChibi
//	* ExtensionTypeRS485
//	* ExtensionTypeWifi
//	* ExtensionTypeEthernet
//	* ExtensionTypeWifi2
func (device *MasterBrick) GetExtensionType(extension uint8) (exttype ExtensionType, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, extension);

	resultBytes, err := device.device.Get(uint8(FunctionGetExtensionType), buf.Bytes())
	if err != nil {
		return exttype, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 12 {
			return exttype, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		if header.ErrorCode != 0 {
			return exttype, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &exttype)

	}

	return exttype, nil
}

// Returns *true* if the Master Brick is at position 0 in the stack and a Chibi
// Extension is available.
func (device *MasterBrick) IsChibiPresent() (present bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionIsChibiPresent), buf.Bytes())
	if err != nil {
		return present, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return present, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return present, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &present)

	}

	return present, nil
}

// Sets the address belonging to the Chibi Extension.
// 
// It is possible to set the address with the Brick Viewer and it will be
// saved in the EEPROM of the Chibi Extension, it does not
// have to be set on every startup.
func (device *MasterBrick) SetChibiAddress(address uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, address);

	resultBytes, err := device.device.Set(uint8(FunctionSetChibiAddress), buf.Bytes())
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

// Returns the address as set by SetChibiAddress.
func (device *MasterBrick) GetChibiAddress() (address uint8, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetChibiAddress), buf.Bytes())
	if err != nil {
		return address, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return address, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return address, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &address)

	}

	return address, nil
}

// Sets the address of the Chibi Master. This address is used if the
// Chibi Extension is used as slave (i.e. it does not have a USB connection).
// 
// It is possible to set the address with the Brick Viewer and it will be
// saved in the EEPROM of the Chibi Extension, it does not
// have to be set on every startup.
func (device *MasterBrick) SetChibiMasterAddress(address uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, address);

	resultBytes, err := device.device.Set(uint8(FunctionSetChibiMasterAddress), buf.Bytes())
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

// Returns the address as set by SetChibiMasterAddress.
func (device *MasterBrick) GetChibiMasterAddress() (address uint8, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetChibiMasterAddress), buf.Bytes())
	if err != nil {
		return address, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return address, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return address, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &address)

	}

	return address, nil
}

// Sets up to 254 slave addresses. 0 has a
// special meaning, it is used as list terminator and not allowed as normal slave
// address. The address numeration (via \c num parameter) has to be used
// ascending from 0. For example: If you use the Chibi Extension in Master mode
// (i.e. the stack has an USB connection) and you want to talk to three other
// Chibi stacks with the slave addresses 17, 23, and 42, you should call with
// ``(0, 17)``, ``(1, 23)``, ``(2, 42)`` and ``(3, 0)``. The last call with
// ``(3, 0)`` is a list terminator and indicates that the Chibi slave address
// list contains 3 addresses in this case.
// 
// It is possible to set the addresses with the Brick Viewer, that will take care
// of correct address numeration and list termination.
// 
// The slave addresses will be saved in the EEPROM of the Chibi Extension, they
// don't have to be set on every startup.
func (device *MasterBrick) SetChibiSlaveAddress(num uint8, address uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, num);
	binary.Write(&buf, binary.LittleEndian, address);

	resultBytes, err := device.device.Set(uint8(FunctionSetChibiSlaveAddress), buf.Bytes())
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

// Returns the slave address for a given \c num as set by
// SetChibiSlaveAddress.
func (device *MasterBrick) GetChibiSlaveAddress(num uint8) (address uint8, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, num);

	resultBytes, err := device.device.Get(uint8(FunctionGetChibiSlaveAddress), buf.Bytes())
	if err != nil {
		return address, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return address, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return address, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &address)

	}

	return address, nil
}

// Returns the signal strength in dBm. The signal strength updates every time a
// packet is received.
func (device *MasterBrick) GetChibiSignalStrength() (signalStrength uint8, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetChibiSignalStrength), buf.Bytes())
	if err != nil {
		return signalStrength, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return signalStrength, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return signalStrength, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &signalStrength)

	}

	return signalStrength, nil
}

// Returns underrun, CRC error, no ACK and overflow error counts of the Chibi
// communication. If these errors start rising, it is likely that either the
// distance between two Chibi stacks is becoming too big or there are
// interferences.
func (device *MasterBrick) GetChibiErrorLog() (underrun uint16, crcError uint16, noAck uint16, overflow uint16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetChibiErrorLog), buf.Bytes())
	if err != nil {
		return underrun, crcError, noAck, overflow, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 16 {
			return underrun, crcError, noAck, overflow, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 16)
		}

		if header.ErrorCode != 0 {
			return underrun, crcError, noAck, overflow, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &underrun)
		binary.Read(resultBuf, binary.LittleEndian, &crcError)
		binary.Read(resultBuf, binary.LittleEndian, &noAck)
		binary.Read(resultBuf, binary.LittleEndian, &overflow)

	}

	return underrun, crcError, noAck, overflow, nil
}

// Sets the Chibi frequency range for the Chibi Extension. Possible values are:
// 
//  Type| Description
//  --- | --- 
//  0|    OQPSK 868MHz (Europe)
//  1|    OQPSK 915MHz (US)
//  2|    OQPSK 780MHz (China)
//  3|    BPSK40 915MHz
// 
// It is possible to set the frequency with the Brick Viewer and it will be
// saved in the EEPROM of the Chibi Extension, it does not
// have to be set on every startup.
//
// Associated constants:
//
//	* ChibiFrequencyOQPSK868MHz
//	* ChibiFrequencyOQPSK915MHz
//	* ChibiFrequencyOQPSK780MHz
//	* ChibiFrequencyBPSK40915MHz
func (device *MasterBrick) SetChibiFrequency(frequency ChibiFrequency) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, frequency);

	resultBytes, err := device.device.Set(uint8(FunctionSetChibiFrequency), buf.Bytes())
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

// Returns the frequency value as set by SetChibiFrequency.
//
// Associated constants:
//
//	* ChibiFrequencyOQPSK868MHz
//	* ChibiFrequencyOQPSK915MHz
//	* ChibiFrequencyOQPSK780MHz
//	* ChibiFrequencyBPSK40915MHz
func (device *MasterBrick) GetChibiFrequency() (frequency ChibiFrequency, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetChibiFrequency), buf.Bytes())
	if err != nil {
		return frequency, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return frequency, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return frequency, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &frequency)

	}

	return frequency, nil
}

// Sets the channel used by the Chibi Extension. Possible channels are
// different for different frequencies:
// 
//  Frequency| Possible Channels
//  --- | --- 
//  OQPSK 868MHz (Europe)| 0
//  OQPSK 915MHz (US)|     1| 2| 3| 4| 5| 6| 7| 8| 9| 10
//  OQPSK 780MHz (China)|  0| 1| 2| 3
//  BPSK40 915MHz|         1| 2| 3| 4| 5| 6| 7| 8| 9| 10
// 
// It is possible to set the channel with the Brick Viewer and it will be
// saved in the EEPROM of the Chibi Extension, it does not
// have to be set on every startup.
func (device *MasterBrick) SetChibiChannel(channel uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, channel);

	resultBytes, err := device.device.Set(uint8(FunctionSetChibiChannel), buf.Bytes())
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

// Returns the channel as set by SetChibiChannel.
func (device *MasterBrick) GetChibiChannel() (channel uint8, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetChibiChannel), buf.Bytes())
	if err != nil {
		return channel, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return channel, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return channel, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &channel)

	}

	return channel, nil
}

// Returns *true* if the Master Brick is at position 0 in the stack and a RS485
// Extension is available.
func (device *MasterBrick) IsRS485Present() (present bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionIsRS485Present), buf.Bytes())
	if err != nil {
		return present, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return present, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return present, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &present)

	}

	return present, nil
}

// Sets the address (0-255) belonging to the RS485 Extension.
// 
// Set to 0 if the RS485 Extension should be the RS485 Master (i.e.
// connected to a PC via USB).
// 
// It is possible to set the address with the Brick Viewer and it will be
// saved in the EEPROM of the RS485 Extension, it does not
// have to be set on every startup.
func (device *MasterBrick) SetRS485Address(address uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, address);

	resultBytes, err := device.device.Set(uint8(FunctionSetRS485Address), buf.Bytes())
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

// Returns the address as set by SetRS485Address.
func (device *MasterBrick) GetRS485Address() (address uint8, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetRS485Address), buf.Bytes())
	if err != nil {
		return address, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return address, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return address, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &address)

	}

	return address, nil
}

// Sets up to 255 slave addresses. Valid addresses are in range 1-255. 0 has a
// special meaning, it is used as list terminator and not allowed as normal slave
// address. The address numeration (via ``num`` parameter) has to be used
// ascending from 0. For example: If you use the RS485 Extension in Master mode
// (i.e. the stack has an USB connection) and you want to talk to three other
// RS485 stacks with the addresses 17, 23, and 42, you should call with
// ``(0, 17)``, ``(1, 23)``, ``(2, 42)`` and ``(3, 0)``. The last call with
// ``(3, 0)`` is a list terminator and indicates that the RS485 slave address list
// contains 3 addresses in this case.
// 
// It is possible to set the addresses with the Brick Viewer, that will take care
// of correct address numeration and list termination.
// 
// The slave addresses will be saved in the EEPROM of the Chibi Extension, they
// don't have to be set on every startup.
func (device *MasterBrick) SetRS485SlaveAddress(num uint8, address uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, num);
	binary.Write(&buf, binary.LittleEndian, address);

	resultBytes, err := device.device.Set(uint8(FunctionSetRS485SlaveAddress), buf.Bytes())
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

// Returns the slave address for a given ``num`` as set by
// SetRS485SlaveAddress.
func (device *MasterBrick) GetRS485SlaveAddress(num uint8) (address uint8, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, num);

	resultBytes, err := device.device.Get(uint8(FunctionGetRS485SlaveAddress), buf.Bytes())
	if err != nil {
		return address, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return address, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return address, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &address)

	}

	return address, nil
}

// Returns CRC error counts of the RS485 communication.
// If this counter starts rising, it is likely that the distance
// between the RS485 nodes is too big or there is some kind of
// interference.
func (device *MasterBrick) GetRS485ErrorLog() (crcError uint16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetRS485ErrorLog), buf.Bytes())
	if err != nil {
		return crcError, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 10 {
			return crcError, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		if header.ErrorCode != 0 {
			return crcError, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &crcError)

	}

	return crcError, nil
}

// Sets the configuration of the RS485 Extension. The
// Master Brick will try to match the given baud rate as exactly as possible.
// The maximum recommended baud rate is 2000000 (2MBd).
// Possible values for parity are 'n' (none), 'e' (even) and 'o' (odd).
// 
// If your RS485 is unstable (lost messages etc.), the first thing you should
// try is to decrease the speed. On very large bus (e.g. 1km), you probably
// should use a value in the range of 100000 (100kBd).
// 
// The values are stored in the EEPROM and only applied on startup. That means
// you have to restart the Master Brick after configuration.
//
// Associated constants:
//
//	* RS485ParityNone
//	* RS485ParityEven
//	* RS485ParityOdd
func (device *MasterBrick) SetRS485Configuration(speed uint32, parity RS485Parity, stopbits uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, speed);
	binary.Write(&buf, binary.LittleEndian, parity);
	binary.Write(&buf, binary.LittleEndian, stopbits);

	resultBytes, err := device.device.Set(uint8(FunctionSetRS485Configuration), buf.Bytes())
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

// Returns the configuration as set by SetRS485Configuration.
//
// Associated constants:
//
//	* RS485ParityNone
//	* RS485ParityEven
//	* RS485ParityOdd
func (device *MasterBrick) GetRS485Configuration() (speed uint32, parity RS485Parity, stopbits uint8, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetRS485Configuration), buf.Bytes())
	if err != nil {
		return speed, parity, stopbits, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 14 {
			return speed, parity, stopbits, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 14)
		}

		if header.ErrorCode != 0 {
			return speed, parity, stopbits, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &speed)
		binary.Read(resultBuf, binary.LittleEndian, &parity)
		binary.Read(resultBuf, binary.LittleEndian, &stopbits)

	}

	return speed, parity, stopbits, nil
}

// Returns *true* if the Master Brick is at position 0 in the stack and a WIFI
// Extension is available.
func (device *MasterBrick) IsWifiPresent() (present bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionIsWifiPresent), buf.Bytes())
	if err != nil {
		return present, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return present, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return present, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &present)

	}

	return present, nil
}

// Sets the configuration of the WIFI Extension. The ``ssid`` can have a max length
// of 32 characters. Possible values for ``connection`` are:
// 
//  Value| Description
//  --- | --- 
//  0| DHCP
//  1| Static IP
//  2| Access Point: DHCP
//  3| Access Point: Static IP
//  4| Ad Hoc: DHCP
//  5| Ad Hoc: Static IP
// 
// If you set ``connection`` to one of the static IP options then you have to
// supply ``ip``, ``subnet_mask`` and ``gateway`` as an array of size 4 (first
// element of the array is the least significant byte of the address). If
// ``connection`` is set to one of the DHCP options then ``ip``, ``subnet_mask``
// and ``gateway`` are ignored, you can set them to 0.
// 
// The last parameter is the port that your program will connect to.
// 
// The values are stored in the EEPROM and only applied on startup. That means
// you have to restart the Master Brick after configuration.
// 
// It is recommended to use the Brick Viewer to set the WIFI configuration.
//
// Associated constants:
//
//	* WifiConnectionDHCP
//	* WifiConnectionStaticIP
//	* WifiConnectionAccessPointDHCP
//	* WifiConnectionAccessPointStaticIP
//	* WifiConnectionAdHocDHCP
//	* WifiConnectionAdHocStaticIP
func (device *MasterBrick) SetWifiConfiguration(ssid string, connection WifiConnection, ip [4]uint8, subnetMask [4]uint8, gateway [4]uint8, port uint16) (err error) {
	var buf bytes.Buffer
	ssid_byte_slice, err := StringToByteSlice(ssid, 32)
	if err != nil { return }
	buf.Write(ssid_byte_slice)
	binary.Write(&buf, binary.LittleEndian, connection);
	binary.Write(&buf, binary.LittleEndian, ip);
	binary.Write(&buf, binary.LittleEndian, subnetMask);
	binary.Write(&buf, binary.LittleEndian, gateway);
	binary.Write(&buf, binary.LittleEndian, port);

	resultBytes, err := device.device.Set(uint8(FunctionSetWifiConfiguration), buf.Bytes())
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

// Returns the configuration as set by SetWifiConfiguration.
//
// Associated constants:
//
//	* WifiConnectionDHCP
//	* WifiConnectionStaticIP
//	* WifiConnectionAccessPointDHCP
//	* WifiConnectionAccessPointStaticIP
//	* WifiConnectionAdHocDHCP
//	* WifiConnectionAdHocStaticIP
func (device *MasterBrick) GetWifiConfiguration() (ssid string, connection WifiConnection, ip [4]uint8, subnetMask [4]uint8, gateway [4]uint8, port uint16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetWifiConfiguration), buf.Bytes())
	if err != nil {
		return ssid, connection, ip, subnetMask, gateway, port, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 55 {
			return ssid, connection, ip, subnetMask, gateway, port, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 55)
		}

		if header.ErrorCode != 0 {
			return ssid, connection, ip, subnetMask, gateway, port, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		ssid = ByteSliceToString(resultBuf.Next(32))
		binary.Read(resultBuf, binary.LittleEndian, &connection)
		binary.Read(resultBuf, binary.LittleEndian, &ip)
		binary.Read(resultBuf, binary.LittleEndian, &subnetMask)
		binary.Read(resultBuf, binary.LittleEndian, &gateway)
		binary.Read(resultBuf, binary.LittleEndian, &port)

	}

	return ssid, connection, ip, subnetMask, gateway, port, nil
}

// Sets the encryption of the WIFI Extension. The first parameter is the
// type of the encryption. Possible values are:
// 
//  Value| Description
//  --- | --- 
//  0| WPA/WPA2
//  1| WPA Enterprise (EAP-FAST| EAP-TLS| EAP-TTLS| PEAP)
//  2| WEP
//  3| No Encryption
// 
// The ``key`` has a max length of 50 characters and is used if ``encryption``
// is set to 0 or 2 (WPA/WPA2 or WEP). Otherwise the value is ignored.
// 
// For WPA/WPA2 the key has to be at least 8 characters long. If you want to set
// a key with more than 50 characters, see SetLongWifiKey.
// 
// For WEP the key has to be either 10 or 26 hexadecimal digits long. It is
// possible to set the WEP ``key_index`` (1-4). If you don't know your
// ``key_index``, it is likely 1.
// 
// If you choose WPA Enterprise as encryption, you have to set ``eap_options`` and
// the length of the certificates (for other encryption types these parameters
// are ignored). The certificates
// themselves can be set with SetWifiCertificate. ``eap_options`` consist
// of the outer authentication (bits 1-2), inner authentication (bit 3) and
// certificate type (bits 4-5):
// 
//  Option| Bits| Description
//  --- | --- | --- 
//  outer authentication| 1-2| 0=EAP-FAST| 1=EAP-TLS| 2=EAP-TTLS| 3=EAP-PEAP
//  inner authentication| 3| 0=EAP-MSCHAP| 1=EAP-GTC
//  certificate type| 4-5| 0=CA Certificate| 1=Client Certificate| 2=Private Key
// 
// Example for EAP-TTLS + EAP-GTC + Private Key: ``option = 2 | (1 << 2) | (2 << 3)``.
// 
// The values are stored in the EEPROM and only applied on startup. That means
// you have to restart the Master Brick after configuration.
// 
// It is recommended to use the Brick Viewer to set the Wi-Fi encryption.
//
// Associated constants:
//
//	* WifiEncryptionWPAWPA2
//	* WifiEncryptionWPAEnterprise
//	* WifiEncryptionWEP
//	* WifiEncryptionNoEncryption
//	* WifiEAPOptionOuterAuthEAPFAST
//	* WifiEAPOptionOuterAuthEAPTLS
//	* WifiEAPOptionOuterAuthEAPTTLS
//	* WifiEAPOptionOuterAuthEAPPEAP
//	* WifiEAPOptionInnerAuthEAPMSCHAP
//	* WifiEAPOptionInnerAuthEAPGTC
//	* WifiEAPOptionCertTypeCACert
//	* WifiEAPOptionCertTypeClientCert
//	* WifiEAPOptionCertTypePrivateKey
func (device *MasterBrick) SetWifiEncryption(encryption WifiEncryption, key string, keyIndex uint8, eapOptions WifiEAPOption, caCertificateLength uint16, clientCertificateLength uint16, privateKeyLength uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, encryption);
	key_byte_slice, err := StringToByteSlice(key, 50)
	if err != nil { return }
	buf.Write(key_byte_slice)
	binary.Write(&buf, binary.LittleEndian, keyIndex);
	binary.Write(&buf, binary.LittleEndian, eapOptions);
	binary.Write(&buf, binary.LittleEndian, caCertificateLength);
	binary.Write(&buf, binary.LittleEndian, clientCertificateLength);
	binary.Write(&buf, binary.LittleEndian, privateKeyLength);

	resultBytes, err := device.device.Set(uint8(FunctionSetWifiEncryption), buf.Bytes())
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

// Returns the encryption as set by SetWifiEncryption.
// 
// Note
//  Since Master Brick Firmware version 2.4.4 the key is not returned anymore.
//
// Associated constants:
//
//	* WifiEncryptionWPAWPA2
//	* WifiEncryptionWPAEnterprise
//	* WifiEncryptionWEP
//	* WifiEncryptionNoEncryption
//	* WifiEAPOptionOuterAuthEAPFAST
//	* WifiEAPOptionOuterAuthEAPTLS
//	* WifiEAPOptionOuterAuthEAPTTLS
//	* WifiEAPOptionOuterAuthEAPPEAP
//	* WifiEAPOptionInnerAuthEAPMSCHAP
//	* WifiEAPOptionInnerAuthEAPGTC
//	* WifiEAPOptionCertTypeCACert
//	* WifiEAPOptionCertTypeClientCert
//	* WifiEAPOptionCertTypePrivateKey
func (device *MasterBrick) GetWifiEncryption() (encryption WifiEncryption, key string, keyIndex uint8, eapOptions WifiEAPOption, caCertificateLength uint16, clientCertificateLength uint16, privateKeyLength uint16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetWifiEncryption), buf.Bytes())
	if err != nil {
		return encryption, key, keyIndex, eapOptions, caCertificateLength, clientCertificateLength, privateKeyLength, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 67 {
			return encryption, key, keyIndex, eapOptions, caCertificateLength, clientCertificateLength, privateKeyLength, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 67)
		}

		if header.ErrorCode != 0 {
			return encryption, key, keyIndex, eapOptions, caCertificateLength, clientCertificateLength, privateKeyLength, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &encryption)
		key = ByteSliceToString(resultBuf.Next(50))
		binary.Read(resultBuf, binary.LittleEndian, &keyIndex)
		binary.Read(resultBuf, binary.LittleEndian, &eapOptions)
		binary.Read(resultBuf, binary.LittleEndian, &caCertificateLength)
		binary.Read(resultBuf, binary.LittleEndian, &clientCertificateLength)
		binary.Read(resultBuf, binary.LittleEndian, &privateKeyLength)

	}

	return encryption, key, keyIndex, eapOptions, caCertificateLength, clientCertificateLength, privateKeyLength, nil
}

// Returns the status of the WIFI Extension. The ``state`` is updated automatically,
// all of the other parameters are updated on startup and every time
// RefreshWifiStatus is called.
// 
// Possible states are:
// 
//  State| Description
//  --- | --- 
//  0| Disassociated
//  1| Associated
//  2| Associating
//  3| Error
//  255| Not initialized yet
//
// Associated constants:
//
//	* WifiStateDisassociated
//	* WifiStateAssociated
//	* WifiStateAssociating
//	* WifiStateError
//	* WifiStateNotInitializedYet
func (device *MasterBrick) GetWifiStatus() (macAddress [6]uint8, bssid [6]uint8, channel uint8, rssi int16, ip [4]uint8, subnetMask [4]uint8, gateway [4]uint8, rxCount uint32, txCount uint32, state WifiState, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetWifiStatus), buf.Bytes())
	if err != nil {
		return macAddress, bssid, channel, rssi, ip, subnetMask, gateway, rxCount, txCount, state, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 44 {
			return macAddress, bssid, channel, rssi, ip, subnetMask, gateway, rxCount, txCount, state, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 44)
		}

		if header.ErrorCode != 0 {
			return macAddress, bssid, channel, rssi, ip, subnetMask, gateway, rxCount, txCount, state, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &macAddress)
		binary.Read(resultBuf, binary.LittleEndian, &bssid)
		binary.Read(resultBuf, binary.LittleEndian, &channel)
		binary.Read(resultBuf, binary.LittleEndian, &rssi)
		binary.Read(resultBuf, binary.LittleEndian, &ip)
		binary.Read(resultBuf, binary.LittleEndian, &subnetMask)
		binary.Read(resultBuf, binary.LittleEndian, &gateway)
		binary.Read(resultBuf, binary.LittleEndian, &rxCount)
		binary.Read(resultBuf, binary.LittleEndian, &txCount)
		binary.Read(resultBuf, binary.LittleEndian, &state)

	}

	return macAddress, bssid, channel, rssi, ip, subnetMask, gateway, rxCount, txCount, state, nil
}

// Refreshes the Wi-Fi status (see GetWifiStatus). To read the status
// of the Wi-Fi module, the Master Brick has to change from data mode to
// command mode and back. This transaction and the readout itself is
// unfortunately time consuming. This means, that it might take some ms
// until the stack with attached WIFI Extension reacts again after this
// function is called.
func (device *MasterBrick) RefreshWifiStatus() (err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Set(uint8(FunctionRefreshWifiStatus), buf.Bytes())
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

// This function is used to set the certificate as well as password and username
// for WPA Enterprise. To set the username use index 0xFFFF,
// to set the password use index 0xFFFE. The max length of username and
// password is 32.
// 
// The certificate is written in chunks of size 32 and the index is used as
// the index of the chunk. ``data_length`` should nearly always be 32. Only
// the last chunk can have a length that is not equal to 32.
// 
// The starting index of the CA Certificate is 0, of the Client Certificate
// 10000 and for the Private Key 20000. Maximum sizes are 1312, 1312 and
// 4320 byte respectively.
// 
// The values are stored in the EEPROM and only applied on startup. That means
// you have to restart the Master Brick after uploading the certificate.
// 
// It is recommended to use the Brick Viewer to set the certificate, username
// and password.
func (device *MasterBrick) SetWifiCertificate(index uint16, data [32]uint8, dataLength uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, index);
	binary.Write(&buf, binary.LittleEndian, data);
	binary.Write(&buf, binary.LittleEndian, dataLength);

	resultBytes, err := device.device.Set(uint8(FunctionSetWifiCertificate), buf.Bytes())
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

// Returns the certificate for a given index as set by SetWifiCertificate.
func (device *MasterBrick) GetWifiCertificate(index uint16) (data [32]uint8, dataLength uint8, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, index);

	resultBytes, err := device.device.Get(uint8(FunctionGetWifiCertificate), buf.Bytes())
	if err != nil {
		return data, dataLength, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 41 {
			return data, dataLength, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 41)
		}

		if header.ErrorCode != 0 {
			return data, dataLength, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &data)
		binary.Read(resultBuf, binary.LittleEndian, &dataLength)

	}

	return data, dataLength, nil
}

// Sets the power mode of the WIFI Extension. Possible modes are:
// 
//  Mode| Description
//  --- | --- 
//  0| Full Speed (high power consumption| high throughput)
//  1| Low Power (low power consumption| low throughput)
//
// Associated constants:
//
//	* WifiPowerModeFullSpeed
//	* WifiPowerModeLowPower
func (device *MasterBrick) SetWifiPowerMode(mode WifiPowerMode) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, mode);

	resultBytes, err := device.device.Set(uint8(FunctionSetWifiPowerMode), buf.Bytes())
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

// Returns the power mode as set by SetWifiPowerMode.
//
// Associated constants:
//
//	* WifiPowerModeFullSpeed
//	* WifiPowerModeLowPower
func (device *MasterBrick) GetWifiPowerMode() (mode WifiPowerMode, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetWifiPowerMode), buf.Bytes())
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

// Returns informations about the Wi-Fi receive buffer. The Wi-Fi
// receive buffer has a max size of 1500 byte and if data is transfered
// too fast, it might overflow.
// 
// The return values are the number of overflows, the low watermark
// (i.e. the smallest number of bytes that were free in the buffer) and
// the bytes that are currently used.
// 
// You should always try to keep the buffer empty, otherwise you will
// have a permanent latency. A good rule of thumb is, that you can transfer
// 1000 messages per second without problems.
// 
// Try to not send more then 50 messages at a time without any kind of
// break between them.
func (device *MasterBrick) GetWifiBufferInfo() (overflow uint32, lowWatermark uint16, used uint16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetWifiBufferInfo), buf.Bytes())
	if err != nil {
		return overflow, lowWatermark, used, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 16 {
			return overflow, lowWatermark, used, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 16)
		}

		if header.ErrorCode != 0 {
			return overflow, lowWatermark, used, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &overflow)
		binary.Read(resultBuf, binary.LittleEndian, &lowWatermark)
		binary.Read(resultBuf, binary.LittleEndian, &used)

	}

	return overflow, lowWatermark, used, nil
}

// Sets the regulatory domain of the WIFI Extension. Possible domains are:
// 
//  Domain| Description
//  --- | --- 
//  0| FCC: Channel 1-11 (N/S America| Australia| New Zealand)
//  1| ETSI: Channel 1-13 (Europe| Middle East| Africa)
//  2| TELEC: Channel 1-14 (Japan)
//
// Associated constants:
//
//	* WifiDomainChannel1To11
//	* WifiDomainChannel1To13
//	* WifiDomainChannel1To14
func (device *MasterBrick) SetWifiRegulatoryDomain(domain WifiDomain) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, domain);

	resultBytes, err := device.device.Set(uint8(FunctionSetWifiRegulatoryDomain), buf.Bytes())
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

// Returns the regulatory domain as set by SetWifiRegulatoryDomain.
//
// Associated constants:
//
//	* WifiDomainChannel1To11
//	* WifiDomainChannel1To13
//	* WifiDomainChannel1To14
func (device *MasterBrick) GetWifiRegulatoryDomain() (domain WifiDomain, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetWifiRegulatoryDomain), buf.Bytes())
	if err != nil {
		return domain, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return domain, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return domain, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &domain)

	}

	return domain, nil
}

// Returns the USB voltage. Does not work with hardware version 2.1.
func (device *MasterBrick) GetUSBVoltage() (voltage uint16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetUSBVoltage), buf.Bytes())
	if err != nil {
		return voltage, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 10 {
			return voltage, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 10)
		}

		if header.ErrorCode != 0 {
			return voltage, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &voltage)

	}

	return voltage, nil
}

// Sets a long Wi-Fi key (up to 63 chars, at least 8 chars) for WPA encryption.
// This key will be used
// if the key in SetWifiEncryption is set to -. In the old protocol,
// a payload of size 63 was not possible, so the maximum key length was 50 chars.
// 
// With the new protocol this is possible, since we didn't want to break API,
// this function was added additionally.
// 
// .. versionadded:: 2.0.2$nbsp;(Firmware)
func (device *MasterBrick) SetLongWifiKey(key string) (err error) {
	var buf bytes.Buffer
	key_byte_slice, err := StringToByteSlice(key, 64)
	if err != nil { return }
	buf.Write(key_byte_slice)

	resultBytes, err := device.device.Set(uint8(FunctionSetLongWifiKey), buf.Bytes())
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

// Returns the encryption key as set by SetLongWifiKey.
// 
// Note
//  Since Master Brick firmware version 2.4.4 the key is not returned anymore.
// 
// .. versionadded:: 2.0.2$nbsp;(Firmware)
func (device *MasterBrick) GetLongWifiKey() (key string, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetLongWifiKey), buf.Bytes())
	if err != nil {
		return key, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 72 {
			return key, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 72)
		}

		if header.ErrorCode != 0 {
			return key, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		key = ByteSliceToString(resultBuf.Next(64))

	}

	return key, nil
}

// Sets the hostname of the WIFI Extension. The hostname will be displayed
// by access points as the hostname in the DHCP clients table.
// 
// Setting an empty String will restore the default hostname.
// 
// .. versionadded:: 2.0.5$nbsp;(Firmware)
func (device *MasterBrick) SetWifiHostname(hostname string) (err error) {
	var buf bytes.Buffer
	hostname_byte_slice, err := StringToByteSlice(hostname, 16)
	if err != nil { return }
	buf.Write(hostname_byte_slice)

	resultBytes, err := device.device.Set(uint8(FunctionSetWifiHostname), buf.Bytes())
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

// Returns the hostname as set by SetWifiHostname.
// 
// An empty String means, that the default hostname is used.
// 
// .. versionadded:: 2.0.5$nbsp;(Firmware)
func (device *MasterBrick) GetWifiHostname() (hostname string, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetWifiHostname), buf.Bytes())
	if err != nil {
		return hostname, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 24 {
			return hostname, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 24)
		}

		if header.ErrorCode != 0 {
			return hostname, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		hostname = ByteSliceToString(resultBuf.Next(16))

	}

	return hostname, nil
}

// Sets the period with which the RegisterStackCurrentCallback callback is triggered
// periodically. A value of 0 turns the callback off.
// 
// The RegisterStackCurrentCallback callback is only triggered if the current has changed
// since the last triggering.
// 
// .. versionadded:: 2.0.5$nbsp;(Firmware)
func (device *MasterBrick) SetStackCurrentCallbackPeriod(period uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);

	resultBytes, err := device.device.Set(uint8(FunctionSetStackCurrentCallbackPeriod), buf.Bytes())
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

// Returns the period as set by SetStackCurrentCallbackPeriod.
// 
// .. versionadded:: 2.0.5$nbsp;(Firmware)
func (device *MasterBrick) GetStackCurrentCallbackPeriod() (period uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetStackCurrentCallbackPeriod), buf.Bytes())
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

// Sets the period with which the RegisterStackVoltageCallback callback is triggered
// periodically. A value of 0 turns the callback off.
// 
// The RegisterStackVoltageCallback callback is only triggered if the voltage has changed
// since the last triggering.
// 
// .. versionadded:: 2.0.5$nbsp;(Firmware)
func (device *MasterBrick) SetStackVoltageCallbackPeriod(period uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);

	resultBytes, err := device.device.Set(uint8(FunctionSetStackVoltageCallbackPeriod), buf.Bytes())
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

// Returns the period as set by SetStackVoltageCallbackPeriod.
// 
// .. versionadded:: 2.0.5$nbsp;(Firmware)
func (device *MasterBrick) GetStackVoltageCallbackPeriod() (period uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetStackVoltageCallbackPeriod), buf.Bytes())
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

// Sets the period with which the RegisterUSBVoltageCallback callback is triggered
// periodically. A value of 0 turns the callback off.
// 
// The RegisterUSBVoltageCallback callback is only triggered if the voltage has changed
// since the last triggering.
// 
// .. versionadded:: 2.0.5$nbsp;(Firmware)
func (device *MasterBrick) SetUSBVoltageCallbackPeriod(period uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, period);

	resultBytes, err := device.device.Set(uint8(FunctionSetUSBVoltageCallbackPeriod), buf.Bytes())
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

// Returns the period as set by SetUSBVoltageCallbackPeriod.
// 
// .. versionadded:: 2.0.5$nbsp;(Firmware)
func (device *MasterBrick) GetUSBVoltageCallbackPeriod() (period uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetUSBVoltageCallbackPeriod), buf.Bytes())
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

// Sets the thresholds for the RegisterStackCurrentReachedCallback callback.
// 
// The following options are possible:
// 
//  Option| Description
//  --- | --- 
//  'x'|    Callback is turned off
//  'o'|    Callback is triggered when the current is *outside* the min and max values
//  'i'|    Callback is triggered when the current is *inside* the min and max values
//  '<'|    Callback is triggered when the current is smaller than the min value (max is ignored)
//  '>'|    Callback is triggered when the current is greater than the min value (max is ignored)
// 
// .. versionadded:: 2.0.5$nbsp;(Firmware)
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *MasterBrick) SetStackCurrentCallbackThreshold(option ThresholdOption, min uint16, max uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, option);
	binary.Write(&buf, binary.LittleEndian, min);
	binary.Write(&buf, binary.LittleEndian, max);

	resultBytes, err := device.device.Set(uint8(FunctionSetStackCurrentCallbackThreshold), buf.Bytes())
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

// Returns the threshold as set by SetStackCurrentCallbackThreshold.
// 
// .. versionadded:: 2.0.5$nbsp;(Firmware)
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *MasterBrick) GetStackCurrentCallbackThreshold() (option ThresholdOption, min uint16, max uint16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetStackCurrentCallbackThreshold), buf.Bytes())
	if err != nil {
		return option, min, max, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 13 {
			return option, min, max, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 13)
		}

		if header.ErrorCode != 0 {
			return option, min, max, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &option)
		binary.Read(resultBuf, binary.LittleEndian, &min)
		binary.Read(resultBuf, binary.LittleEndian, &max)

	}

	return option, min, max, nil
}

// Sets the thresholds for the RegisterStackVoltageReachedCallback callback.
// 
// The following options are possible:
// 
//  Option| Description
//  --- | --- 
//  'x'|    Callback is turned off
//  'o'|    Callback is triggered when the voltage is *outside* the min and max values
//  'i'|    Callback is triggered when the voltage is *inside* the min and max values
//  '<'|    Callback is triggered when the voltage is smaller than the min value (max is ignored)
//  '>'|    Callback is triggered when the voltage is greater than the min value (max is ignored)
// 
// .. versionadded:: 2.0.5$nbsp;(Firmware)
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *MasterBrick) SetStackVoltageCallbackThreshold(option ThresholdOption, min uint16, max uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, option);
	binary.Write(&buf, binary.LittleEndian, min);
	binary.Write(&buf, binary.LittleEndian, max);

	resultBytes, err := device.device.Set(uint8(FunctionSetStackVoltageCallbackThreshold), buf.Bytes())
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

// Returns the threshold as set by SetStackVoltageCallbackThreshold.
// 
// .. versionadded:: 2.0.5$nbsp;(Firmware)
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *MasterBrick) GetStackVoltageCallbackThreshold() (option ThresholdOption, min uint16, max uint16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetStackVoltageCallbackThreshold), buf.Bytes())
	if err != nil {
		return option, min, max, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 13 {
			return option, min, max, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 13)
		}

		if header.ErrorCode != 0 {
			return option, min, max, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &option)
		binary.Read(resultBuf, binary.LittleEndian, &min)
		binary.Read(resultBuf, binary.LittleEndian, &max)

	}

	return option, min, max, nil
}

// Sets the thresholds for the RegisterUSBVoltageReachedCallback callback.
// 
// The following options are possible:
// 
//  Option| Description
//  --- | --- 
//  'x'|    Callback is turned off
//  'o'|    Callback is triggered when the voltage is *outside* the min and max values
//  'i'|    Callback is triggered when the voltage is *inside* the min and max values
//  '<'|    Callback is triggered when the voltage is smaller than the min value (max is ignored)
//  '>'|    Callback is triggered when the voltage is greater than the min value (max is ignored)
// 
// .. versionadded:: 2.0.5$nbsp;(Firmware)
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *MasterBrick) SetUSBVoltageCallbackThreshold(option ThresholdOption, min uint16, max uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, option);
	binary.Write(&buf, binary.LittleEndian, min);
	binary.Write(&buf, binary.LittleEndian, max);

	resultBytes, err := device.device.Set(uint8(FunctionSetUSBVoltageCallbackThreshold), buf.Bytes())
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

// Returns the threshold as set by SetUSBVoltageCallbackThreshold.
// 
// .. versionadded:: 2.0.5$nbsp;(Firmware)
//
// Associated constants:
//
//	* ThresholdOptionOff
//	* ThresholdOptionOutside
//	* ThresholdOptionInside
//	* ThresholdOptionSmaller
//	* ThresholdOptionGreater
func (device *MasterBrick) GetUSBVoltageCallbackThreshold() (option ThresholdOption, min uint16, max uint16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetUSBVoltageCallbackThreshold), buf.Bytes())
	if err != nil {
		return option, min, max, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 13 {
			return option, min, max, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 13)
		}

		if header.ErrorCode != 0 {
			return option, min, max, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &option)
		binary.Read(resultBuf, binary.LittleEndian, &min)
		binary.Read(resultBuf, binary.LittleEndian, &max)

	}

	return option, min, max, nil
}

// Sets the period with which the threshold callbacks
// 
// * RegisterStackCurrentReachedCallback,
// * RegisterStackVoltageReachedCallback,
// * RegisterUSBVoltageReachedCallback
// 
// are triggered, if the thresholds
// 
// * SetStackCurrentCallbackThreshold,
// * SetStackVoltageCallbackThreshold,
// * SetUSBVoltageCallbackThreshold
// 
// keep being reached.
// 
// .. versionadded:: 2.0.5$nbsp;(Firmware)
func (device *MasterBrick) SetDebouncePeriod(debounce uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, debounce);

	resultBytes, err := device.device.Set(uint8(FunctionSetDebouncePeriod), buf.Bytes())
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

// Returns the debounce period as set by SetDebouncePeriod.
// 
// .. versionadded:: 2.0.5$nbsp;(Firmware)
func (device *MasterBrick) GetDebouncePeriod() (debounce uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetDebouncePeriod), buf.Bytes())
	if err != nil {
		return debounce, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 12 {
			return debounce, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		if header.ErrorCode != 0 {
			return debounce, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &debounce)

	}

	return debounce, nil
}

// Returns *true* if the Master Brick is at position 0 in the stack and an Ethernet
// Extension is available.
// 
// .. versionadded:: 2.1.0$nbsp;(Firmware)
func (device *MasterBrick) IsEthernetPresent() (present bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionIsEthernetPresent), buf.Bytes())
	if err != nil {
		return present, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return present, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return present, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &present)

	}

	return present, nil
}

// Sets the configuration of the Ethernet Extension. Possible values for
// ``connection`` are:
// 
//  Value| Description
//  --- | --- 
//  0| DHCP
//  1| Static IP
// 
// If you set ``connection`` to static IP options then you have to supply ``ip``,
// ``subnet_mask`` and ``gateway`` as an array of size 4 (first element of the
// array is the least significant byte of the address). If ``connection`` is set
// to the DHCP options then ``ip``, ``subnet_mask`` and ``gateway`` are ignored,
// you can set them to 0.
// 
// The last parameter is the port that your program will connect to.
// 
// The values are stored in the EEPROM and only applied on startup. That means
// you have to restart the Master Brick after configuration.
// 
// It is recommended to use the Brick Viewer to set the Ethernet configuration.
// 
// .. versionadded:: 2.1.0$nbsp;(Firmware)
//
// Associated constants:
//
//	* EthernetConnectionDHCP
//	* EthernetConnectionStaticIP
func (device *MasterBrick) SetEthernetConfiguration(connection EthernetConnection, ip [4]uint8, subnetMask [4]uint8, gateway [4]uint8, port uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, connection);
	binary.Write(&buf, binary.LittleEndian, ip);
	binary.Write(&buf, binary.LittleEndian, subnetMask);
	binary.Write(&buf, binary.LittleEndian, gateway);
	binary.Write(&buf, binary.LittleEndian, port);

	resultBytes, err := device.device.Set(uint8(FunctionSetEthernetConfiguration), buf.Bytes())
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

// Returns the configuration as set by SetEthernetConfiguration.
// 
// .. versionadded:: 2.1.0$nbsp;(Firmware)
//
// Associated constants:
//
//	* EthernetConnectionDHCP
//	* EthernetConnectionStaticIP
func (device *MasterBrick) GetEthernetConfiguration() (connection EthernetConnection, ip [4]uint8, subnetMask [4]uint8, gateway [4]uint8, port uint16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetEthernetConfiguration), buf.Bytes())
	if err != nil {
		return connection, ip, subnetMask, gateway, port, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 23 {
			return connection, ip, subnetMask, gateway, port, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 23)
		}

		if header.ErrorCode != 0 {
			return connection, ip, subnetMask, gateway, port, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &connection)
		binary.Read(resultBuf, binary.LittleEndian, &ip)
		binary.Read(resultBuf, binary.LittleEndian, &subnetMask)
		binary.Read(resultBuf, binary.LittleEndian, &gateway)
		binary.Read(resultBuf, binary.LittleEndian, &port)

	}

	return connection, ip, subnetMask, gateway, port, nil
}

// Returns the status of the Ethernet Extension.
// 
// ``mac_address``, ``ip``, ``subnet_mask`` and ``gateway`` are given as an array.
// The first element of the array is the least significant byte of the address.
// 
// ``rx_count`` and ``tx_count`` are the number of bytes that have been
// received/send since last restart.
// 
// ``hostname`` is the currently used hostname.
// 
// .. versionadded:: 2.1.0$nbsp;(Firmware)
func (device *MasterBrick) GetEthernetStatus() (macAddress [6]uint8, ip [4]uint8, subnetMask [4]uint8, gateway [4]uint8, rxCount uint32, txCount uint32, hostname string, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetEthernetStatus), buf.Bytes())
	if err != nil {
		return macAddress, ip, subnetMask, gateway, rxCount, txCount, hostname, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 66 {
			return macAddress, ip, subnetMask, gateway, rxCount, txCount, hostname, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 66)
		}

		if header.ErrorCode != 0 {
			return macAddress, ip, subnetMask, gateway, rxCount, txCount, hostname, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &macAddress)
		binary.Read(resultBuf, binary.LittleEndian, &ip)
		binary.Read(resultBuf, binary.LittleEndian, &subnetMask)
		binary.Read(resultBuf, binary.LittleEndian, &gateway)
		binary.Read(resultBuf, binary.LittleEndian, &rxCount)
		binary.Read(resultBuf, binary.LittleEndian, &txCount)
		hostname = ByteSliceToString(resultBuf.Next(32))

	}

	return macAddress, ip, subnetMask, gateway, rxCount, txCount, hostname, nil
}

// Sets the hostname of the Ethernet Extension. The hostname will be displayed
// by access points as the hostname in the DHCP clients table.
// 
// Setting an empty String will restore the default hostname.
// 
// The current hostname can be discovered with GetEthernetStatus.
// 
// .. versionadded:: 2.1.0$nbsp;(Firmware)
func (device *MasterBrick) SetEthernetHostname(hostname string) (err error) {
	var buf bytes.Buffer
	hostname_byte_slice, err := StringToByteSlice(hostname, 32)
	if err != nil { return }
	buf.Write(hostname_byte_slice)

	resultBytes, err := device.device.Set(uint8(FunctionSetEthernetHostname), buf.Bytes())
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

// Sets the MAC address of the Ethernet Extension. The Ethernet Extension should
// come configured with a valid MAC address, that is also written on a
// sticker of the extension itself.
// 
// The MAC address can be read out again with GetEthernetStatus.
// 
// .. versionadded:: 2.1.0$nbsp;(Firmware)
func (device *MasterBrick) SetEthernetMACAddress(macAddress [6]uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, macAddress);

	resultBytes, err := device.device.Set(uint8(FunctionSetEthernetMACAddress), buf.Bytes())
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

// Sets the Ethernet WebSocket configuration. The first parameter sets the number of socket
// connections that are reserved for WebSockets. The range is 0-7. The connections
// are shared with the plain sockets. Example: If you set the connections to 3,
// there will be 3 WebSocket and 4 plain socket connections available.
// 
// The second parameter is the port for the WebSocket connections. The port can
// not be the same as the port for the plain socket connections.
// 
// The values are stored in the EEPROM and only applied on startup. That means
// you have to restart the Master Brick after configuration.
// 
// It is recommended to use the Brick Viewer to set the Ethernet configuration.
// 
// .. versionadded:: 2.2.0$nbsp;(Firmware)
func (device *MasterBrick) SetEthernetWebsocketConfiguration(sockets uint8, port uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, sockets);
	binary.Write(&buf, binary.LittleEndian, port);

	resultBytes, err := device.device.Set(uint8(FunctionSetEthernetWebsocketConfiguration), buf.Bytes())
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

// Returns the configuration as set by SetEthernetConfiguration.
// 
// .. versionadded:: 2.2.0$nbsp;(Firmware)
func (device *MasterBrick) GetEthernetWebsocketConfiguration() (sockets uint8, port uint16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetEthernetWebsocketConfiguration), buf.Bytes())
	if err != nil {
		return sockets, port, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 11 {
			return sockets, port, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 11)
		}

		if header.ErrorCode != 0 {
			return sockets, port, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &sockets)
		binary.Read(resultBuf, binary.LittleEndian, &port)

	}

	return sockets, port, nil
}

// Sets the Ethernet authentication secret. The secret can be a string of up to 64
// characters. An empty string disables the authentication.
// 
// See the `authentication tutorial <tutorial_authentication>` for more
// information.
// 
// The secret is stored in the EEPROM and only applied on startup. That means
// you have to restart the Master Brick after configuration.
// 
// It is recommended to use the Brick Viewer to set the Ethernet authentication secret.
// 
// The default value is an empty string (authentication disabled).
// 
// .. versionadded:: 2.2.0$nbsp;(Firmware)
func (device *MasterBrick) SetEthernetAuthenticationSecret(secret string) (err error) {
	var buf bytes.Buffer
	secret_byte_slice, err := StringToByteSlice(secret, 64)
	if err != nil { return }
	buf.Write(secret_byte_slice)

	resultBytes, err := device.device.Set(uint8(FunctionSetEthernetAuthenticationSecret), buf.Bytes())
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

// Returns the authentication secret as set by
// SetEthernetAuthenticationSecret.
// 
// .. versionadded:: 2.2.0$nbsp;(Firmware)
func (device *MasterBrick) GetEthernetAuthenticationSecret() (secret string, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetEthernetAuthenticationSecret), buf.Bytes())
	if err != nil {
		return secret, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 72 {
			return secret, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 72)
		}

		if header.ErrorCode != 0 {
			return secret, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		secret = ByteSliceToString(resultBuf.Next(64))

	}

	return secret, nil
}

// Sets the WIFI authentication secret. The secret can be a string of up to 64
// characters. An empty string disables the authentication.
// 
// See the `authentication tutorial <tutorial_authentication>` for more
// information.
// 
// The secret is stored in the EEPROM and only applied on startup. That means
// you have to restart the Master Brick after configuration.
// 
// It is recommended to use the Brick Viewer to set the WIFI authentication secret.
// 
// The default value is an empty string (authentication disabled).
// 
// .. versionadded:: 2.2.0$nbsp;(Firmware)
func (device *MasterBrick) SetWifiAuthenticationSecret(secret string) (err error) {
	var buf bytes.Buffer
	secret_byte_slice, err := StringToByteSlice(secret, 64)
	if err != nil { return }
	buf.Write(secret_byte_slice)

	resultBytes, err := device.device.Set(uint8(FunctionSetWifiAuthenticationSecret), buf.Bytes())
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

// Returns the authentication secret as set by
// SetWifiAuthenticationSecret.
// 
// .. versionadded:: 2.2.0$nbsp;(Firmware)
func (device *MasterBrick) GetWifiAuthenticationSecret() (secret string, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetWifiAuthenticationSecret), buf.Bytes())
	if err != nil {
		return secret, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 72 {
			return secret, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 72)
		}

		if header.ErrorCode != 0 {
			return secret, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		secret = ByteSliceToString(resultBuf.Next(64))

	}

	return secret, nil
}

// Returns the type of the connection over which this function was called.
// 
// .. versionadded:: 2.4.0$nbsp;(Firmware)
//
// Associated constants:
//
//	* ConnectionTypeNone
//	* ConnectionTypeUSB
//	* ConnectionTypeSPIStack
//	* ConnectionTypeChibi
//	* ConnectionTypeRS485
//	* ConnectionTypeWifi
//	* ConnectionTypeEthernet
//	* ConnectionTypeWifi2
func (device *MasterBrick) GetConnectionType() (connectionType ConnectionType, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetConnectionType), buf.Bytes())
	if err != nil {
		return connectionType, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return connectionType, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return connectionType, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &connectionType)

	}

	return connectionType, nil
}

// Returns *true* if the Master Brick is at position 0 in the stack and a WIFI
// Extension 2.0 is available.
// 
// .. versionadded:: 2.4.0$nbsp;(Firmware)
func (device *MasterBrick) IsWifi2Present() (present bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionIsWifi2Present), buf.Bytes())
	if err != nil {
		return present, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return present, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return present, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &present)

	}

	return present, nil
}

// Starts the bootloader of the WIFI Extension 2.0. Returns 0 on success.
// Afterwards the WriteWifi2SerialPort and ReadWifi2SerialPort
// functions can be used to communicate with the bootloader to flash a new
// firmware.
// 
// The bootloader should only be started over a USB connection. It cannot be
// started over a WIFI2 connection, see the GetConnectionType function.
// 
// It is recommended to use the Brick Viewer to update the firmware of the WIFI
// Extension 2.0.
// 
// .. versionadded:: 2.4.0$nbsp;(Firmware)
func (device *MasterBrick) StartWifi2Bootloader() (result int8, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionStartWifi2Bootloader), buf.Bytes())
	if err != nil {
		return result, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return result, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return result, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &result)

	}

	return result, nil
}

// Writes up to 60 bytes (number of bytes to be written specified by ``length``)
// to the serial port of the bootloader of the WIFI Extension 2.0. Returns 0 on
// success.
// 
// Before this function can be used the bootloader has to be started using the
// StartWifi2Bootloader function.
// 
// It is recommended to use the Brick Viewer to update the firmware of the WIFI
// Extension 2.0.
// 
// .. versionadded:: 2.4.0$nbsp;(Firmware)
func (device *MasterBrick) WriteWifi2SerialPort(data [60]uint8, length uint8) (result int8, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, data);
	binary.Write(&buf, binary.LittleEndian, length);

	resultBytes, err := device.device.Get(uint8(FunctionWriteWifi2SerialPort), buf.Bytes())
	if err != nil {
		return result, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return result, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return result, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &result)

	}

	return result, nil
}

// Reads up to 60 bytes (number of bytes to be read specified by ``length``)
// from the serial port of the bootloader of the WIFI Extension 2.0.
// Returns the number of actually read bytes.
// 
// Before this function can be used the bootloader has to be started using the
// StartWifi2Bootloader function.
// 
// It is recommended to use the Brick Viewer to update the firmware of the WIFI
// Extension 2.0.
// 
// .. versionadded:: 2.4.0$nbsp;(Firmware)
func (device *MasterBrick) ReadWifi2SerialPort(length uint8) (data [60]uint8, result uint8, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, length);

	resultBytes, err := device.device.Get(uint8(FunctionReadWifi2SerialPort), buf.Bytes())
	if err != nil {
		return data, result, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 69 {
			return data, result, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 69)
		}

		if header.ErrorCode != 0 {
			return data, result, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &data)
		binary.Read(resultBuf, binary.LittleEndian, &result)

	}

	return data, result, nil
}

// Sets the WIFI authentication secret. The secret can be a string of up to 64
// characters. An empty string disables the authentication. The default value is
// an empty string (authentication disabled).
// 
// See the `authentication tutorial <tutorial_authentication>` for more
// information.
// 
// To apply configuration changes to the WIFI Extension 2.0 the
// SaveWifi2Configuration function has to be called and the Master Brick
// has to be restarted afterwards.
// 
// It is recommended to use the Brick Viewer to configure the WIFI Extension 2.0.
// 
// .. versionadded:: 2.4.0$nbsp;(Firmware)
func (device *MasterBrick) SetWifi2AuthenticationSecret(secret string) (err error) {
	var buf bytes.Buffer
	secret_byte_slice, err := StringToByteSlice(secret, 64)
	if err != nil { return }
	buf.Write(secret_byte_slice)

	resultBytes, err := device.device.Set(uint8(FunctionSetWifi2AuthenticationSecret), buf.Bytes())
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

// Returns the WIFI authentication secret as set by
// SetWifi2AuthenticationSecret.
// 
// .. versionadded:: 2.4.0$nbsp;(Firmware)
func (device *MasterBrick) GetWifi2AuthenticationSecret() (secret string, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetWifi2AuthenticationSecret), buf.Bytes())
	if err != nil {
		return secret, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 72 {
			return secret, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 72)
		}

		if header.ErrorCode != 0 {
			return secret, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		secret = ByteSliceToString(resultBuf.Next(64))

	}

	return secret, nil
}

// Sets the general configuration of the WIFI Extension 2.0.
// 
// The ``port`` parameter sets the port number that your programm will connect
// to.
// 
// The ``websocket_port`` parameter sets the WebSocket port number that your
// JavaScript programm will connect to.
// 
// The ``website_port`` parameter sets the port number for the website of the
// WIFI Extension 2.0.
// 
// The ``phy_mode`` parameter sets the specific wireless network mode to be used.
// Possible values are B, G and N.
// 
// The ``sleep_mode`` parameter is currently unused.
// 
// The ``website`` parameter is used to enable or disable the web interface of
// the WIFI Extension 2.0, which is available from firmware version 2.0.1. Note
// that, for firmware version 2.0.3 and older, to disable the the web interface
// the ``website_port`` parameter must be set to 1 and greater than 1 to enable
// the web interface. For firmware version 2.0.4 and later, setting this parameter
// to 1 will enable the web interface and setting it to 0 will disable the web
// interface.
// 
// To apply configuration changes to the WIFI Extension 2.0 the
// SaveWifi2Configuration function has to be called and the Master Brick
// has to be restarted afterwards.
// 
// It is recommended to use the Brick Viewer to configure the WIFI Extension 2.0.
// 
// .. versionadded:: 2.4.0$nbsp;(Firmware)
//
// Associated constants:
//
//	* Wifi2PHYModeB
//	* Wifi2PHYModeG
//	* Wifi2PHYModeN
func (device *MasterBrick) SetWifi2Configuration(port uint16, websocketPort uint16, websitePort uint16, phyMode Wifi2PHYMode, sleepMode uint8, website uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, port);
	binary.Write(&buf, binary.LittleEndian, websocketPort);
	binary.Write(&buf, binary.LittleEndian, websitePort);
	binary.Write(&buf, binary.LittleEndian, phyMode);
	binary.Write(&buf, binary.LittleEndian, sleepMode);
	binary.Write(&buf, binary.LittleEndian, website);

	resultBytes, err := device.device.Set(uint8(FunctionSetWifi2Configuration), buf.Bytes())
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

// Returns the general configuration as set by SetWifi2Configuration.
// 
// .. versionadded:: 2.4.0$nbsp;(Firmware)
//
// Associated constants:
//
//	* Wifi2PHYModeB
//	* Wifi2PHYModeG
//	* Wifi2PHYModeN
func (device *MasterBrick) GetWifi2Configuration() (port uint16, websocketPort uint16, websitePort uint16, phyMode Wifi2PHYMode, sleepMode uint8, website uint8, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetWifi2Configuration), buf.Bytes())
	if err != nil {
		return port, websocketPort, websitePort, phyMode, sleepMode, website, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 17 {
			return port, websocketPort, websitePort, phyMode, sleepMode, website, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 17)
		}

		if header.ErrorCode != 0 {
			return port, websocketPort, websitePort, phyMode, sleepMode, website, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &port)
		binary.Read(resultBuf, binary.LittleEndian, &websocketPort)
		binary.Read(resultBuf, binary.LittleEndian, &websitePort)
		binary.Read(resultBuf, binary.LittleEndian, &phyMode)
		binary.Read(resultBuf, binary.LittleEndian, &sleepMode)
		binary.Read(resultBuf, binary.LittleEndian, &website)

	}

	return port, websocketPort, websitePort, phyMode, sleepMode, website, nil
}

// Returns the client and access point status of the WIFI Extension 2.0.
// 
// .. versionadded:: 2.4.0$nbsp;(Firmware)
//
// Associated constants:
//
//	* Wifi2ClientStatusIdle
//	* Wifi2ClientStatusConnecting
//	* Wifi2ClientStatusWrongPassword
//	* Wifi2ClientStatusNoAPFound
//	* Wifi2ClientStatusConnectFailed
//	* Wifi2ClientStatusGotIP
//	* Wifi2ClientStatusUnknown
func (device *MasterBrick) GetWifi2Status() (clientEnabled bool, clientStatus Wifi2ClientStatus, clientIP [4]uint8, clientSubnetMask [4]uint8, clientGateway [4]uint8, clientMACAddress [6]uint8, clientRXCount uint32, clientTXCount uint32, clientRSSI int8, apEnabled bool, apIP [4]uint8, apSubnetMask [4]uint8, apGateway [4]uint8, apMACAddress [6]uint8, apRXCount uint32, apTXCount uint32, apConnectedCount uint8, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetWifi2Status), buf.Bytes())
	if err != nil {
		return clientEnabled, clientStatus, clientIP, clientSubnetMask, clientGateway, clientMACAddress, clientRXCount, clientTXCount, clientRSSI, apEnabled, apIP, apSubnetMask, apGateway, apMACAddress, apRXCount, apTXCount, apConnectedCount, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 65 {
			return clientEnabled, clientStatus, clientIP, clientSubnetMask, clientGateway, clientMACAddress, clientRXCount, clientTXCount, clientRSSI, apEnabled, apIP, apSubnetMask, apGateway, apMACAddress, apRXCount, apTXCount, apConnectedCount, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 65)
		}

		if header.ErrorCode != 0 {
			return clientEnabled, clientStatus, clientIP, clientSubnetMask, clientGateway, clientMACAddress, clientRXCount, clientTXCount, clientRSSI, apEnabled, apIP, apSubnetMask, apGateway, apMACAddress, apRXCount, apTXCount, apConnectedCount, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &clientEnabled)
		binary.Read(resultBuf, binary.LittleEndian, &clientStatus)
		binary.Read(resultBuf, binary.LittleEndian, &clientIP)
		binary.Read(resultBuf, binary.LittleEndian, &clientSubnetMask)
		binary.Read(resultBuf, binary.LittleEndian, &clientGateway)
		binary.Read(resultBuf, binary.LittleEndian, &clientMACAddress)
		binary.Read(resultBuf, binary.LittleEndian, &clientRXCount)
		binary.Read(resultBuf, binary.LittleEndian, &clientTXCount)
		binary.Read(resultBuf, binary.LittleEndian, &clientRSSI)
		binary.Read(resultBuf, binary.LittleEndian, &apEnabled)
		binary.Read(resultBuf, binary.LittleEndian, &apIP)
		binary.Read(resultBuf, binary.LittleEndian, &apSubnetMask)
		binary.Read(resultBuf, binary.LittleEndian, &apGateway)
		binary.Read(resultBuf, binary.LittleEndian, &apMACAddress)
		binary.Read(resultBuf, binary.LittleEndian, &apRXCount)
		binary.Read(resultBuf, binary.LittleEndian, &apTXCount)
		binary.Read(resultBuf, binary.LittleEndian, &apConnectedCount)

	}

	return clientEnabled, clientStatus, clientIP, clientSubnetMask, clientGateway, clientMACAddress, clientRXCount, clientTXCount, clientRSSI, apEnabled, apIP, apSubnetMask, apGateway, apMACAddress, apRXCount, apTXCount, apConnectedCount, nil
}

// Sets the client specific configuration of the WIFI Extension 2.0.
// 
// The ``enable`` parameter enables or disables the client part of the
// WIFI Extension 2.0.
// 
// The ``ssid`` parameter sets the SSID (up to 32 characters) of the access point
// to connect to.
// 
// If the ``ip`` parameter is set to all zero then ``subnet_mask`` and ``gateway``
// parameters are also set to all zero and DHCP is used for IP address configuration.
// Otherwise those three parameters can be used to configure a static IP address.
// The default configuration is DHCP.
// 
// If the ``mac_address`` parameter is set to all zero then the factory MAC
// address is used. Otherwise this parameter can be used to set a custom MAC
// address.
// 
// If the ``bssid`` parameter is set to all zero then WIFI Extension 2.0 will
// connect to any access point that matches the configured SSID. Otherwise this
// parameter can be used to make the WIFI Extension 2.0 only connect to an
// access point if SSID and BSSID match.
// 
// To apply configuration changes to the WIFI Extension 2.0 the
// SaveWifi2Configuration function has to be called and the Master Brick
// has to be restarted afterwards.
// 
// It is recommended to use the Brick Viewer to configure the WIFI Extension 2.0.
// 
// .. versionadded:: 2.4.0$nbsp;(Firmware)
func (device *MasterBrick) SetWifi2ClientConfiguration(enable bool, ssid string, ip [4]uint8, subnetMask [4]uint8, gateway [4]uint8, macAddress [6]uint8, bssid [6]uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, enable);
	ssid_byte_slice, err := StringToByteSlice(ssid, 32)
	if err != nil { return }
	buf.Write(ssid_byte_slice)
	binary.Write(&buf, binary.LittleEndian, ip);
	binary.Write(&buf, binary.LittleEndian, subnetMask);
	binary.Write(&buf, binary.LittleEndian, gateway);
	binary.Write(&buf, binary.LittleEndian, macAddress);
	binary.Write(&buf, binary.LittleEndian, bssid);

	resultBytes, err := device.device.Set(uint8(FunctionSetWifi2ClientConfiguration), buf.Bytes())
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

// Returns the client configuration as set by SetWifi2ClientConfiguration.
// 
// .. versionadded:: 2.4.0$nbsp;(Firmware)
func (device *MasterBrick) GetWifi2ClientConfiguration() (enable bool, ssid string, ip [4]uint8, subnetMask [4]uint8, gateway [4]uint8, macAddress [6]uint8, bssid [6]uint8, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetWifi2ClientConfiguration), buf.Bytes())
	if err != nil {
		return enable, ssid, ip, subnetMask, gateway, macAddress, bssid, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 65 {
			return enable, ssid, ip, subnetMask, gateway, macAddress, bssid, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 65)
		}

		if header.ErrorCode != 0 {
			return enable, ssid, ip, subnetMask, gateway, macAddress, bssid, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &enable)
		ssid = ByteSliceToString(resultBuf.Next(32))
		binary.Read(resultBuf, binary.LittleEndian, &ip)
		binary.Read(resultBuf, binary.LittleEndian, &subnetMask)
		binary.Read(resultBuf, binary.LittleEndian, &gateway)
		binary.Read(resultBuf, binary.LittleEndian, &macAddress)
		binary.Read(resultBuf, binary.LittleEndian, &bssid)

	}

	return enable, ssid, ip, subnetMask, gateway, macAddress, bssid, nil
}

// Sets the client hostname (up to 32 characters) of the WIFI Extension 2.0. The
// hostname will be displayed by access points as the hostname in the DHCP clients
// table.
// 
// To apply configuration changes to the WIFI Extension 2.0 the
// SaveWifi2Configuration function has to be called and the Master Brick
// has to be restarted afterwards.
// 
// It is recommended to use the Brick Viewer to configure the WIFI Extension 2.0.
// 
// .. versionadded:: 2.4.0$nbsp;(Firmware)
func (device *MasterBrick) SetWifi2ClientHostname(hostname string) (err error) {
	var buf bytes.Buffer
	hostname_byte_slice, err := StringToByteSlice(hostname, 32)
	if err != nil { return }
	buf.Write(hostname_byte_slice)

	resultBytes, err := device.device.Set(uint8(FunctionSetWifi2ClientHostname), buf.Bytes())
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

// Returns the client hostname as set by SetWifi2ClientHostname.
// 
// .. versionadded:: 2.4.0$nbsp;(Firmware)
func (device *MasterBrick) GetWifi2ClientHostname() (hostname string, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetWifi2ClientHostname), buf.Bytes())
	if err != nil {
		return hostname, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 40 {
			return hostname, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 40)
		}

		if header.ErrorCode != 0 {
			return hostname, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		hostname = ByteSliceToString(resultBuf.Next(32))

	}

	return hostname, nil
}

// Sets the client password (up to 63 chars) for WPA/WPA2 encryption.
// 
// To apply configuration changes to the WIFI Extension 2.0 the
// SaveWifi2Configuration function has to be called and the Master Brick
// has to be restarted afterwards.
// 
// It is recommended to use the Brick Viewer to configure the WIFI Extension 2.0.
// 
// .. versionadded:: 2.4.0$nbsp;(Firmware)
func (device *MasterBrick) SetWifi2ClientPassword(password string) (err error) {
	var buf bytes.Buffer
	password_byte_slice, err := StringToByteSlice(password, 64)
	if err != nil { return }
	buf.Write(password_byte_slice)

	resultBytes, err := device.device.Set(uint8(FunctionSetWifi2ClientPassword), buf.Bytes())
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

// Returns the client password as set by SetWifi2ClientPassword.
// 
// Note
//  Since WIFI Extension 2.0 firmware version 2.1.3 the password is not
//  returned anymore.
// 
// .. versionadded:: 2.4.0$nbsp;(Firmware)
func (device *MasterBrick) GetWifi2ClientPassword() (password string, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetWifi2ClientPassword), buf.Bytes())
	if err != nil {
		return password, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 72 {
			return password, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 72)
		}

		if header.ErrorCode != 0 {
			return password, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		password = ByteSliceToString(resultBuf.Next(64))

	}

	return password, nil
}

// Sets the access point specific configuration of the WIFI Extension 2.0.
// 
// The ``enable`` parameter enables or disables the access point part of the
// WIFI Extension 2.0.
// 
// The ``ssid`` parameter sets the SSID (up to 32 characters) of the access point.
// 
// If the ``ip`` parameter is set to all zero then ``subnet_mask`` and ``gateway``
// parameters are also set to all zero and DHCP is used for IP address configuration.
// Otherwise those three parameters can be used to configure a static IP address.
// The default configuration is DHCP.
// 
// The ``encryption`` parameter sets the encryption mode to be used. Possible
// values are Open (no encryption), WEP or WPA/WPA2 PSK.
// Use the SetWifi2APPassword function to set the encryption
// password.
// 
// The ``hidden`` parameter makes the access point hide or show its SSID.
// 
// The ``channel`` parameter sets the channel (1 to 13) of the access point.
// 
// If the ``mac_address`` parameter is set to all zero then the factory MAC
// address is used. Otherwise this parameter can be used to set a custom MAC
// address.
// 
// To apply configuration changes to the WIFI Extension 2.0 the
// SaveWifi2Configuration function has to be called and the Master Brick
// has to be restarted afterwards.
// 
// It is recommended to use the Brick Viewer to configure the WIFI Extension 2.0.
// 
// .. versionadded:: 2.4.0$nbsp;(Firmware)
//
// Associated constants:
//
//	* Wifi2APEncryptionOpen
//	* Wifi2APEncryptionWEP
//	* Wifi2APEncryptionWPAPSK
//	* Wifi2APEncryptionWPA2PSK
//	* Wifi2APEncryptionWPAWPA2PSK
func (device *MasterBrick) SetWifi2APConfiguration(enable bool, ssid string, ip [4]uint8, subnetMask [4]uint8, gateway [4]uint8, encryption Wifi2APEncryption, hidden bool, channel uint8, macAddress [6]uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, enable);
	ssid_byte_slice, err := StringToByteSlice(ssid, 32)
	if err != nil { return }
	buf.Write(ssid_byte_slice)
	binary.Write(&buf, binary.LittleEndian, ip);
	binary.Write(&buf, binary.LittleEndian, subnetMask);
	binary.Write(&buf, binary.LittleEndian, gateway);
	binary.Write(&buf, binary.LittleEndian, encryption);
	binary.Write(&buf, binary.LittleEndian, hidden);
	binary.Write(&buf, binary.LittleEndian, channel);
	binary.Write(&buf, binary.LittleEndian, macAddress);

	resultBytes, err := device.device.Set(uint8(FunctionSetWifi2APConfiguration), buf.Bytes())
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

// Returns the access point configuration as set by SetWifi2APConfiguration.
// 
// .. versionadded:: 2.4.0$nbsp;(Firmware)
//
// Associated constants:
//
//	* Wifi2APEncryptionOpen
//	* Wifi2APEncryptionWEP
//	* Wifi2APEncryptionWPAPSK
//	* Wifi2APEncryptionWPA2PSK
//	* Wifi2APEncryptionWPAWPA2PSK
func (device *MasterBrick) GetWifi2APConfiguration() (enable bool, ssid string, ip [4]uint8, subnetMask [4]uint8, gateway [4]uint8, encryption Wifi2APEncryption, hidden bool, channel uint8, macAddress [6]uint8, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetWifi2APConfiguration), buf.Bytes())
	if err != nil {
		return enable, ssid, ip, subnetMask, gateway, encryption, hidden, channel, macAddress, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 62 {
			return enable, ssid, ip, subnetMask, gateway, encryption, hidden, channel, macAddress, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 62)
		}

		if header.ErrorCode != 0 {
			return enable, ssid, ip, subnetMask, gateway, encryption, hidden, channel, macAddress, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &enable)
		ssid = ByteSliceToString(resultBuf.Next(32))
		binary.Read(resultBuf, binary.LittleEndian, &ip)
		binary.Read(resultBuf, binary.LittleEndian, &subnetMask)
		binary.Read(resultBuf, binary.LittleEndian, &gateway)
		binary.Read(resultBuf, binary.LittleEndian, &encryption)
		binary.Read(resultBuf, binary.LittleEndian, &hidden)
		binary.Read(resultBuf, binary.LittleEndian, &channel)
		binary.Read(resultBuf, binary.LittleEndian, &macAddress)

	}

	return enable, ssid, ip, subnetMask, gateway, encryption, hidden, channel, macAddress, nil
}

// Sets the access point password (at least 8 and up to 63 chars) for the configured encryption
// mode, see SetWifi2APConfiguration.
// 
// To apply configuration changes to the WIFI Extension 2.0 the
// SaveWifi2Configuration function has to be called and the Master Brick
// has to be restarted afterwards.
// 
// It is recommended to use the Brick Viewer to configure the WIFI Extension 2.0.
// 
// .. versionadded:: 2.4.0$nbsp;(Firmware)
func (device *MasterBrick) SetWifi2APPassword(password string) (err error) {
	var buf bytes.Buffer
	password_byte_slice, err := StringToByteSlice(password, 64)
	if err != nil { return }
	buf.Write(password_byte_slice)

	resultBytes, err := device.device.Set(uint8(FunctionSetWifi2APPassword), buf.Bytes())
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

// Returns the access point password as set by SetWifi2APPassword.
// 
// Note
//  Since WIFI Extension 2.0 firmware version 2.1.3 the password is not
//  returned anymore.
// 
// .. versionadded:: 2.4.0$nbsp;(Firmware)
func (device *MasterBrick) GetWifi2APPassword() (password string, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetWifi2APPassword), buf.Bytes())
	if err != nil {
		return password, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 72 {
			return password, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 72)
		}

		if header.ErrorCode != 0 {
			return password, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		password = ByteSliceToString(resultBuf.Next(64))

	}

	return password, nil
}

// All configuration functions for the WIFI Extension 2.0 do not change the
// values permanently. After configuration this function has to be called to
// permanently store the values.
// 
// The values are stored in the EEPROM and only applied on startup. That means
// you have to restart the Master Brick after configuration.
// 
// .. versionadded:: 2.4.0$nbsp;(Firmware)
func (device *MasterBrick) SaveWifi2Configuration() (result uint8, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionSaveWifi2Configuration), buf.Bytes())
	if err != nil {
		return result, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return result, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return result, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &result)

	}

	return result, nil
}

// Returns the current version of the WIFI Extension 2.0 firmware.
// 
// .. versionadded:: 2.4.0$nbsp;(Firmware)
func (device *MasterBrick) GetWifi2FirmwareVersion() (firmwareVersion [3]uint8, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetWifi2FirmwareVersion), buf.Bytes())
	if err != nil {
		return firmwareVersion, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 11 {
			return firmwareVersion, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 11)
		}

		if header.ErrorCode != 0 {
			return firmwareVersion, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &firmwareVersion)

	}

	return firmwareVersion, nil
}

// Turns the green status LED of the WIFI Extension 2.0 on.
// 
// .. versionadded:: 2.4.0$nbsp;(Firmware)
func (device *MasterBrick) EnableWifi2StatusLED() (err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Set(uint8(FunctionEnableWifi2StatusLED), buf.Bytes())
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

// Turns the green status LED of the WIFI Extension 2.0 off.
// 
// .. versionadded:: 2.4.0$nbsp;(Firmware)
func (device *MasterBrick) DisableWifi2StatusLED() (err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Set(uint8(FunctionDisableWifi2StatusLED), buf.Bytes())
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

// Returns *true* if the green status LED of the WIFI Extension 2.0 is turned on.
// 
// .. versionadded:: 2.4.0$nbsp;(Firmware)
func (device *MasterBrick) IsWifi2StatusLEDEnabled() (enabled bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionIsWifi2StatusLEDEnabled), buf.Bytes())
	if err != nil {
		return enabled, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return enabled, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return enabled, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &enabled)

	}

	return enabled, nil
}

// Requires WIFI Extension 2.0 firmware 2.1.0.
// 
// Sets the mesh specific configuration of the WIFI Extension 2.0.
// 
// The ``enable`` parameter enables or disables the mesh part of the
// WIFI Extension 2.0. The mesh part cannot be
// enabled together with the client and access-point part.
// 
// If the ``root_ip`` parameter is set to all zero then ``root_subnet_mask``
// and ``root_gateway`` parameters are also set to all zero and DHCP is used for
// IP address configuration. Otherwise those three parameters can be used to
// configure a static IP address. The default configuration is DHCP.
// 
// If the ``router_bssid`` parameter is set to all zero then the information is
// taken from Wi-Fi scan when connecting the SSID as set by
// SetWifi2MeshRouterSSID. This only works if the the SSID is not hidden.
// In case the router has hidden SSID this parameter must be specified, otherwise
// the node will not be able to reach the mesh router.
// 
// The ``group_id`` and the ``group_ssid_prefix`` parameters identifies a
// particular mesh network and nodes configured with same ``group_id`` and the
// ``group_ssid_prefix`` are considered to be in the same mesh network.
// 
// The ``gateway_ip`` and the ``gateway_port`` parameters specifies the location
// of the brickd that supports mesh feature.
// 
// To apply configuration changes to the WIFI Extension 2.0 the
// SaveWifi2Configuration function has to be called and the Master Brick
// has to be restarted afterwards.
// 
// It is recommended to use the Brick Viewer to configure the WIFI Extension 2.0.
// 
// .. versionadded:: 2.4.2$nbsp;(Firmware)
func (device *MasterBrick) SetWifi2MeshConfiguration(enable bool, rootIP [4]uint8, rootSubnetMask [4]uint8, rootGateway [4]uint8, routerBSSID [6]uint8, groupID [6]uint8, groupSSIDPrefix string, gatewayIP [4]uint8, gatewayPort uint16) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, enable);
	binary.Write(&buf, binary.LittleEndian, rootIP);
	binary.Write(&buf, binary.LittleEndian, rootSubnetMask);
	binary.Write(&buf, binary.LittleEndian, rootGateway);
	binary.Write(&buf, binary.LittleEndian, routerBSSID);
	binary.Write(&buf, binary.LittleEndian, groupID);
	groupSSIDPrefix_byte_slice, err := StringToByteSlice(groupSSIDPrefix, 16)
	if err != nil { return }
	buf.Write(groupSSIDPrefix_byte_slice)
	binary.Write(&buf, binary.LittleEndian, gatewayIP);
	binary.Write(&buf, binary.LittleEndian, gatewayPort);

	resultBytes, err := device.device.Set(uint8(FunctionSetWifi2MeshConfiguration), buf.Bytes())
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

// Requires WIFI Extension 2.0 firmware 2.1.0.
// 
// Returns the mesh configuration as set by SetWifi2MeshConfiguration.
// 
// .. versionadded:: 2.4.2$nbsp;(Firmware)
func (device *MasterBrick) GetWifi2MeshConfiguration() (enable bool, rootIP [4]uint8, rootSubnetMask [4]uint8, rootGateway [4]uint8, routerBSSID [6]uint8, groupID [6]uint8, groupSSIDPrefix string, gatewayIP [4]uint8, gatewayPort uint16, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetWifi2MeshConfiguration), buf.Bytes())
	if err != nil {
		return enable, rootIP, rootSubnetMask, rootGateway, routerBSSID, groupID, groupSSIDPrefix, gatewayIP, gatewayPort, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 55 {
			return enable, rootIP, rootSubnetMask, rootGateway, routerBSSID, groupID, groupSSIDPrefix, gatewayIP, gatewayPort, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 55)
		}

		if header.ErrorCode != 0 {
			return enable, rootIP, rootSubnetMask, rootGateway, routerBSSID, groupID, groupSSIDPrefix, gatewayIP, gatewayPort, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &enable)
		binary.Read(resultBuf, binary.LittleEndian, &rootIP)
		binary.Read(resultBuf, binary.LittleEndian, &rootSubnetMask)
		binary.Read(resultBuf, binary.LittleEndian, &rootGateway)
		binary.Read(resultBuf, binary.LittleEndian, &routerBSSID)
		binary.Read(resultBuf, binary.LittleEndian, &groupID)
		groupSSIDPrefix = ByteSliceToString(resultBuf.Next(16))
		binary.Read(resultBuf, binary.LittleEndian, &gatewayIP)
		binary.Read(resultBuf, binary.LittleEndian, &gatewayPort)

	}

	return enable, rootIP, rootSubnetMask, rootGateway, routerBSSID, groupID, groupSSIDPrefix, gatewayIP, gatewayPort, nil
}

// Requires WIFI Extension 2.0 firmware 2.1.0.
// 
// Sets the mesh router SSID of the WIFI Extension 2.0.
// It is used to specify the mesh router to connect to.
// 
// Note that even though in the argument of this function a 32 characters long SSID
// is allowed, in practice valid SSID should have a maximum of 31 characters. This
// is due to a bug in the mesh library that we use in the firmware of the extension.
// 
// To apply configuration changes to the WIFI Extension 2.0 the
// SaveWifi2Configuration function has to be called and the Master Brick
// has to be restarted afterwards.
// 
// It is recommended to use the Brick Viewer to configure the WIFI Extension 2.0.
// 
// .. versionadded:: 2.4.2$nbsp;(Firmware)
func (device *MasterBrick) SetWifi2MeshRouterSSID(ssid string) (err error) {
	var buf bytes.Buffer
	ssid_byte_slice, err := StringToByteSlice(ssid, 32)
	if err != nil { return }
	buf.Write(ssid_byte_slice)

	resultBytes, err := device.device.Set(uint8(FunctionSetWifi2MeshRouterSSID), buf.Bytes())
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

// Requires WIFI Extension 2.0 firmware 2.1.0.
// 
// Returns the mesh router SSID as set by SetWifi2MeshRouterSSID.
// 
// .. versionadded:: 2.4.2$nbsp;(Firmware)
func (device *MasterBrick) GetWifi2MeshRouterSSID() (ssid string, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetWifi2MeshRouterSSID), buf.Bytes())
	if err != nil {
		return ssid, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 40 {
			return ssid, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 40)
		}

		if header.ErrorCode != 0 {
			return ssid, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		ssid = ByteSliceToString(resultBuf.Next(32))

	}

	return ssid, nil
}

// Requires WIFI Extension 2.0 firmware 2.1.0.
// 
// Sets the mesh router password (up to 64 characters) for WPA/WPA2 encryption.
// The password will be used to connect to the mesh router.
// 
// To apply configuration changes to the WIFI Extension 2.0 the
// SaveWifi2Configuration function has to be called and the Master Brick
// has to be restarted afterwards.
// 
// It is recommended to use the Brick Viewer to configure the WIFI Extension 2.0.
// 
// .. versionadded:: 2.4.2$nbsp;(Firmware)
func (device *MasterBrick) SetWifi2MeshRouterPassword(password string) (err error) {
	var buf bytes.Buffer
	password_byte_slice, err := StringToByteSlice(password, 64)
	if err != nil { return }
	buf.Write(password_byte_slice)

	resultBytes, err := device.device.Set(uint8(FunctionSetWifi2MeshRouterPassword), buf.Bytes())
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

// Requires WIFI Extension 2.0 firmware 2.1.0.
// 
// Returns the mesh router password as set by SetWifi2MeshRouterPassword.
// 
// .. versionadded:: 2.4.2$nbsp;(Firmware)
func (device *MasterBrick) GetWifi2MeshRouterPassword() (password string, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetWifi2MeshRouterPassword), buf.Bytes())
	if err != nil {
		return password, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 72 {
			return password, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 72)
		}

		if header.ErrorCode != 0 {
			return password, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		password = ByteSliceToString(resultBuf.Next(64))

	}

	return password, nil
}

// Requires WIFI Extension 2.0 firmware 2.1.0.
// 
// Returns the common mesh status of the WIFI Extension 2.0.
// 
// .. versionadded:: 2.4.2$nbsp;(Firmware)
//
// Associated constants:
//
//	* Wifi2MeshStatusDisabled
//	* Wifi2MeshStatusWIFIConnecting
//	* Wifi2MeshStatusGotIP
//	* Wifi2MeshStatusMeshLocal
//	* Wifi2MeshStatusMeshOnline
//	* Wifi2MeshStatusAPAvailable
//	* Wifi2MeshStatusAPSetup
//	* Wifi2MeshStatusLeafAvailable
func (device *MasterBrick) GetWifi2MeshCommonStatus() (status Wifi2MeshStatus, rootNode bool, rootCandidate bool, connectedNodes uint16, rxCount uint32, txCount uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetWifi2MeshCommonStatus), buf.Bytes())
	if err != nil {
		return status, rootNode, rootCandidate, connectedNodes, rxCount, txCount, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 21 {
			return status, rootNode, rootCandidate, connectedNodes, rxCount, txCount, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 21)
		}

		if header.ErrorCode != 0 {
			return status, rootNode, rootCandidate, connectedNodes, rxCount, txCount, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &status)
		binary.Read(resultBuf, binary.LittleEndian, &rootNode)
		binary.Read(resultBuf, binary.LittleEndian, &rootCandidate)
		binary.Read(resultBuf, binary.LittleEndian, &connectedNodes)
		binary.Read(resultBuf, binary.LittleEndian, &rxCount)
		binary.Read(resultBuf, binary.LittleEndian, &txCount)

	}

	return status, rootNode, rootCandidate, connectedNodes, rxCount, txCount, nil
}

// Requires WIFI Extension 2.0 firmware 2.1.0.
// 
// Returns the mesh client status of the WIFI Extension 2.0.
// 
// .. versionadded:: 2.4.2$nbsp;(Firmware)
func (device *MasterBrick) GetWifi2MeshClientStatus() (hostname string, ip [4]uint8, subnetMask [4]uint8, gateway [4]uint8, macAddress [6]uint8, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetWifi2MeshClientStatus), buf.Bytes())
	if err != nil {
		return hostname, ip, subnetMask, gateway, macAddress, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 58 {
			return hostname, ip, subnetMask, gateway, macAddress, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 58)
		}

		if header.ErrorCode != 0 {
			return hostname, ip, subnetMask, gateway, macAddress, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		hostname = ByteSliceToString(resultBuf.Next(32))
		binary.Read(resultBuf, binary.LittleEndian, &ip)
		binary.Read(resultBuf, binary.LittleEndian, &subnetMask)
		binary.Read(resultBuf, binary.LittleEndian, &gateway)
		binary.Read(resultBuf, binary.LittleEndian, &macAddress)

	}

	return hostname, ip, subnetMask, gateway, macAddress, nil
}

// Requires WIFI Extension 2.0 firmware 2.1.0.
// 
// Returns the mesh AP status of the WIFI Extension 2.0.
// 
// .. versionadded:: 2.4.2$nbsp;(Firmware)
func (device *MasterBrick) GetWifi2MeshAPStatus() (ssid string, ip [4]uint8, subnetMask [4]uint8, gateway [4]uint8, macAddress [6]uint8, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetWifi2MeshAPStatus), buf.Bytes())
	if err != nil {
		return ssid, ip, subnetMask, gateway, macAddress, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 58 {
			return ssid, ip, subnetMask, gateway, macAddress, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 58)
		}

		if header.ErrorCode != 0 {
			return ssid, ip, subnetMask, gateway, macAddress, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		ssid = ByteSliceToString(resultBuf.Next(32))
		binary.Read(resultBuf, binary.LittleEndian, &ip)
		binary.Read(resultBuf, binary.LittleEndian, &subnetMask)
		binary.Read(resultBuf, binary.LittleEndian, &gateway)
		binary.Read(resultBuf, binary.LittleEndian, &macAddress)

	}

	return ssid, ip, subnetMask, gateway, macAddress, nil
}

// This function is for internal use to flash the initial
// bootstrapper and bootloader to the Bricklets.
// 
// If you need to flash a boostrapper/bootloader (for exmaple
// because you made your own Bricklet from scratch) please
// take a look at our open source flash and test tool at
// https://github.com/Tinkerforge/flash-test
// 
// Don't use this function directly.
// 
// .. versionadded:: 2.5.0$nbsp;(Firmware)
func (device *MasterBrick) SetBrickletXMCFlashConfig(config uint32, parameter1 uint32, parameter2 uint32, data [52]uint8) (returnValue uint32, returnData [60]uint8, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, config);
	binary.Write(&buf, binary.LittleEndian, parameter1);
	binary.Write(&buf, binary.LittleEndian, parameter2);
	binary.Write(&buf, binary.LittleEndian, data);

	resultBytes, err := device.device.Get(uint8(FunctionSetBrickletXMCFlashConfig), buf.Bytes())
	if err != nil {
		return returnValue, returnData, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 72 {
			return returnValue, returnData, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 72)
		}

		if header.ErrorCode != 0 {
			return returnValue, returnData, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &returnValue)
		binary.Read(resultBuf, binary.LittleEndian, &returnData)

	}

	return returnValue, returnData, nil
}

// This function is for internal use to flash the initial
// bootstrapper and bootloader to the Bricklets.
// 
// If you need to flash a boostrapper/bootloader (for exmaple
// because you made your own Bricklet from scratch) please
// take a look at our open source flash and test tool at
// https://github.com/Tinkerforge/flash-test
// 
// Don't use this function directly.
// 
// .. versionadded:: 2.5.0$nbsp;(Firmware)
func (device *MasterBrick) SetBrickletXMCFlashData(data [64]uint8) (returnData uint32, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, data);

	resultBytes, err := device.device.Get(uint8(FunctionSetBrickletXMCFlashData), buf.Bytes())
	if err != nil {
		return returnData, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 12 {
			return returnData, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		if header.ErrorCode != 0 {
			return returnData, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &returnData)

	}

	return returnData, nil
}

// This function is only available in Master Brick hardware version >= 3.0.
// 
// Enables/disables all four Bricklets if set to true/false.
// 
// If you disable the Bricklets the power supply to the Bricklets will be disconnected.
// The Bricklets will lose all configurations if disabled.
// 
// .. versionadded:: 2.5.0$nbsp;(Firmware)
func (device *MasterBrick) SetBrickletsEnabled(brickletsEnabled bool) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, brickletsEnabled);

	resultBytes, err := device.device.Set(uint8(FunctionSetBrickletsEnabled), buf.Bytes())
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

// Returns *true* if the Bricklets are enabled, *false* otherwise.
// 
// .. versionadded:: 2.5.0$nbsp;(Firmware)
func (device *MasterBrick) GetBrickletsEnabled() (brickletsEnabled bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetBrickletsEnabled), buf.Bytes())
	if err != nil {
		return brickletsEnabled, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return brickletsEnabled, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return brickletsEnabled, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &brickletsEnabled)

	}

	return brickletsEnabled, nil
}

// The SPITF protocol can be used with a dynamic baudrate. If the dynamic baudrate is
// enabled, the Brick will try to adapt the baudrate for the communication
// between Bricks and Bricklets according to the amount of data that is transferred.
// 
// The baudrate will be increased exponentially if lots of data is sent/received and
// decreased linearly if little data is sent/received.
// 
// This lowers the baudrate in applications where little data is transferred (e.g.
// a weather station) and increases the robustness. If there is lots of data to transfer
// (e.g. Thermal Imaging Bricklet) it automatically increases the baudrate as needed.
// 
// In cases where some data has to transferred as fast as possible every few seconds
// (e.g. RS485 Bricklet with a high baudrate but small payload) you may want to turn
// the dynamic baudrate off to get the highest possible performance.
// 
// The maximum value of the baudrate can be set per port with the function
// SetSPITFPBaudrate. If the dynamic baudrate is disabled, the baudrate
// as set by SetSPITFPBaudrate will be used statically.
// 
// .. versionadded:: 2.4.6$nbsp;(Firmware)
func (device *MasterBrick) SetSPITFPBaudrateConfig(enableDynamicBaudrate bool, minimumDynamicBaudrate uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, enableDynamicBaudrate);
	binary.Write(&buf, binary.LittleEndian, minimumDynamicBaudrate);

	resultBytes, err := device.device.Set(uint8(FunctionSetSPITFPBaudrateConfig), buf.Bytes())
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

// Returns the baudrate config, see SetSPITFPBaudrateConfig.
// 
// .. versionadded:: 2.4.6$nbsp;(Firmware)
func (device *MasterBrick) GetSPITFPBaudrateConfig() (enableDynamicBaudrate bool, minimumDynamicBaudrate uint32, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionGetSPITFPBaudrateConfig), buf.Bytes())
	if err != nil {
		return enableDynamicBaudrate, minimumDynamicBaudrate, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 13 {
			return enableDynamicBaudrate, minimumDynamicBaudrate, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 13)
		}

		if header.ErrorCode != 0 {
			return enableDynamicBaudrate, minimumDynamicBaudrate, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &enableDynamicBaudrate)
		binary.Read(resultBuf, binary.LittleEndian, &minimumDynamicBaudrate)

	}

	return enableDynamicBaudrate, minimumDynamicBaudrate, nil
}

// Returns the timeout count for the different communication methods.
// 
// The methods 0-2 are available for all Bricks, 3-7 only for Master Bricks.
// 
// This function is mostly used for debugging during development, in normal operation
// the counters should nearly always stay at 0.
// 
// .. versionadded:: 2.4.3$nbsp;(Firmware)
//
// Associated constants:
//
//	* CommunicationMethodNone
//	* CommunicationMethodUSB
//	* CommunicationMethodSPIStack
//	* CommunicationMethodChibi
//	* CommunicationMethodRS485
//	* CommunicationMethodWIFI
//	* CommunicationMethodEthernet
//	* CommunicationMethodWIFIV2
func (device *MasterBrick) GetSendTimeoutCount(communicationMethod CommunicationMethod) (timeoutCount uint32, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, communicationMethod);

	resultBytes, err := device.device.Get(uint8(FunctionGetSendTimeoutCount), buf.Bytes())
	if err != nil {
		return timeoutCount, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 12 {
			return timeoutCount, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		if header.ErrorCode != 0 {
			return timeoutCount, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &timeoutCount)

	}

	return timeoutCount, nil
}

// Sets the baudrate for a specific Bricklet port.
// 
// If you want to increase the throughput of Bricklets you can increase
// the baudrate. If you get a high error count because of high
// interference (see GetSPITFPErrorCount) you can decrease the
// baudrate.
// 
// If the dynamic baudrate feature is enabled, the baudrate set by this
// function corresponds to the maximum baudrate (see SetSPITFPBaudrateConfig).
// 
// Regulatory testing is done with the default baudrate. If CE compatibility
// or similar is necessary in your applications we recommend to not change
// the baudrate.
// 
// .. versionadded:: 2.4.3$nbsp;(Firmware)
func (device *MasterBrick) SetSPITFPBaudrate(brickletPort rune, baudrate uint32) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, brickletPort);
	binary.Write(&buf, binary.LittleEndian, baudrate);

	resultBytes, err := device.device.Set(uint8(FunctionSetSPITFPBaudrate), buf.Bytes())
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

// Returns the baudrate for a given Bricklet port, see SetSPITFPBaudrate.
// 
// .. versionadded:: 2.4.3$nbsp;(Firmware)
func (device *MasterBrick) GetSPITFPBaudrate(brickletPort rune) (baudrate uint32, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, brickletPort);

	resultBytes, err := device.device.Get(uint8(FunctionGetSPITFPBaudrate), buf.Bytes())
	if err != nil {
		return baudrate, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 12 {
			return baudrate, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 12)
		}

		if header.ErrorCode != 0 {
			return baudrate, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &baudrate)

	}

	return baudrate, nil
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
// The errors counts are for errors that occur on the Brick side. All
// Bricklets have a similar function that returns the errors on the Bricklet side.
// 
// .. versionadded:: 2.4.3$nbsp;(Firmware)
func (device *MasterBrick) GetSPITFPErrorCount(brickletPort rune) (errorCountACKChecksum uint32, errorCountMessageChecksum uint32, errorCountFrame uint32, errorCountOverflow uint32, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, brickletPort);

	resultBytes, err := device.device.Get(uint8(FunctionGetSPITFPErrorCount), buf.Bytes())
	if err != nil {
		return errorCountACKChecksum, errorCountMessageChecksum, errorCountFrame, errorCountOverflow, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 24 {
			return errorCountACKChecksum, errorCountMessageChecksum, errorCountFrame, errorCountOverflow, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 24)
		}

		if header.ErrorCode != 0 {
			return errorCountACKChecksum, errorCountMessageChecksum, errorCountFrame, errorCountOverflow, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &errorCountACKChecksum)
		binary.Read(resultBuf, binary.LittleEndian, &errorCountMessageChecksum)
		binary.Read(resultBuf, binary.LittleEndian, &errorCountFrame)
		binary.Read(resultBuf, binary.LittleEndian, &errorCountOverflow)

	}

	return errorCountACKChecksum, errorCountMessageChecksum, errorCountFrame, errorCountOverflow, nil
}

// Enables the status LED.
// 
// The status LED is the blue LED next to the USB connector. If enabled is is
// on and it flickers if data is transfered. If disabled it is always off.
// 
// The default state is enabled.
// 
// .. versionadded:: 2.3.2$nbsp;(Firmware)
func (device *MasterBrick) EnableStatusLED() (err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Set(uint8(FunctionEnableStatusLED), buf.Bytes())
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

// Disables the status LED.
// 
// The status LED is the blue LED next to the USB connector. If enabled is is
// on and it flickers if data is transfered. If disabled it is always off.
// 
// The default state is enabled.
// 
// .. versionadded:: 2.3.2$nbsp;(Firmware)
func (device *MasterBrick) DisableStatusLED() (err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Set(uint8(FunctionDisableStatusLED), buf.Bytes())
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

// Returns *true* if the status LED is enabled, *false* otherwise.
// 
// .. versionadded:: 2.3.2$nbsp;(Firmware)
func (device *MasterBrick) IsStatusLEDEnabled() (enabled bool, err error) {
	var buf bytes.Buffer
	
	resultBytes, err := device.device.Get(uint8(FunctionIsStatusLEDEnabled), buf.Bytes())
	if err != nil {
		return enabled, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 9 {
			return enabled, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 9)
		}

		if header.ErrorCode != 0 {
			return enabled, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &enabled)

	}

	return enabled, nil
}

// Returns the firmware and protocol version and the name of the Bricklet for a
// given port.
// 
// This functions sole purpose is to allow automatic flashing of v1.x.y Bricklet
// plugins.
func (device *MasterBrick) GetProtocol1BrickletName(port rune) (protocolVersion uint8, firmwareVersion [3]uint8, name string, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, port);

	resultBytes, err := device.device.Get(uint8(FunctionGetProtocol1BrickletName), buf.Bytes())
	if err != nil {
		return protocolVersion, firmwareVersion, name, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 52 {
			return protocolVersion, firmwareVersion, name, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 52)
		}

		if header.ErrorCode != 0 {
			return protocolVersion, firmwareVersion, name, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &protocolVersion)
		binary.Read(resultBuf, binary.LittleEndian, &firmwareVersion)
		name = ByteSliceToString(resultBuf.Next(40))

	}

	return protocolVersion, firmwareVersion, name, nil
}

// Returns the temperature as measured inside the microcontroller. The
// value returned is not the ambient temperature!
// 
// The temperature is only proportional to the real temperature and it has an
// accuracy of ±15%. Practically it is only useful as an indicator for
// temperature changes.
func (device *MasterBrick) GetChipTemperature() (temperature int16, err error) {
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

// Calling this function will reset the Brick. Calling this function
// on a Brick inside of a stack will reset the whole stack.
// 
// After a reset you have to create new device objects,
// calling functions on the existing ones will result in
// undefined behavior!
func (device *MasterBrick) Reset() (err error) {
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

// Writes 32 bytes of firmware to the bricklet attached at the given port.
// The bytes are written to the position offset * 32.
// 
// This function is used by Brick Viewer during flashing. It should not be
// necessary to call it in a normal user program.
func (device *MasterBrick) WriteBrickletPlugin(port rune, offset uint8, chunk [32]uint8) (err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, port);
	binary.Write(&buf, binary.LittleEndian, offset);
	binary.Write(&buf, binary.LittleEndian, chunk);

	resultBytes, err := device.device.Set(uint8(FunctionWriteBrickletPlugin), buf.Bytes())
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

// Reads 32 bytes of firmware from the bricklet attached at the given port.
// The bytes are read starting at the position offset * 32.
// 
// This function is used by Brick Viewer during flashing. It should not be
// necessary to call it in a normal user program.
func (device *MasterBrick) ReadBrickletPlugin(port rune, offset uint8) (chunk [32]uint8, err error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, port);
	binary.Write(&buf, binary.LittleEndian, offset);

	resultBytes, err := device.device.Get(uint8(FunctionReadBrickletPlugin), buf.Bytes())
	if err != nil {
		return chunk, err
	}
	if len(resultBytes) > 0 {
		var header PacketHeader

		header.FillFromBytes(resultBytes)

		if header.Length != 40 {
			return chunk, fmt.Errorf("Received packet of unexpected size %d, instead of %d", header.Length, 40)
		}

		if header.ErrorCode != 0 {
			return chunk, DeviceError(header.ErrorCode)
		}

		resultBuf := bytes.NewBuffer(resultBytes[8:])
		binary.Read(resultBuf, binary.LittleEndian, &chunk)

	}

	return chunk, nil
}

// Returns the UID, the UID where the Brick is connected to,
// the position, the hardware and firmware version as well as the
// device identifier.
// 
// The position is the position in the stack from '0' (bottom) to '8' (top).
// 
// The device identifier numbers can be found `here <device_identifier>`.
// |device_identifier_constant|
func (device *MasterBrick) GetIdentity() (uid string, connectedUid string, position rune, hardwareVersion [3]uint8, firmwareVersion [3]uint8, deviceIdentifier uint16, err error) {
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
