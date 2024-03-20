package efidevicepath

// https://uefi.org/specs/UEFI/2.10/10_Protocols_Device_Path_Protocol.html#ipv6-device-path

const IPv6Type = 3 + 13*0x100

const (
	IPv6ManualOrigin        = 0x00
	IPv6StatelessAutoOrigin = 0x01
	IPv6StatefullAutoOrigin = 0x02
)

type IPv6 struct {
	LocalIPAddress   [16]byte
	RemoteIPAddress  [16]byte
	LocalPort        uint16
	RemotePort       uint16
	Protocol         [2]byte // RFC 3232
	IPAddressOrigin  byte
	PrefixLength     uint8
	GatewayIPAddress [16]byte
}
