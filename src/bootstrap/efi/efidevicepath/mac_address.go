package efidevicepath

// https://uefi.org/specs/UEFI/2.10/10_Protocols_Device_Path_Protocol.html#mac-address-device-path

const MACAddressType = 3 + 11*0x100

type MACAddress struct {
	MACAddress [32]byte
	IfType     byte
}
