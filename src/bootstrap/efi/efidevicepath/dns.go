package efidevicepath

// https://uefi.org/specs/UEFI/2.10/10_Protocols_Device_Path_Protocol.html#dns-device-path

const DNSType = 3 + 31*0x100

type DNS struct {
	IsIPv6    bool
	Instances []byte
}
