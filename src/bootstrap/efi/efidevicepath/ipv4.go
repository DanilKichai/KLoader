package efidevicepath

// https://uefi.org/specs/UEFI/2.10/10_Protocols_Device_Path_Protocol.html#ipv4-device-path

const IPv4Type = 3 + 12*0x100

type IPv4 struct {
	LocalIPAddress   [4]byte
	RemoteIPAddress  [4]byte
	LocalPort        uint16
	RemotePort       uint16
	Protocol         [2]byte // RFC 3232
	StaticIPAddress  bool
	GatewayIPAddress [4]byte
	SubnetMask       [4]byte
}
