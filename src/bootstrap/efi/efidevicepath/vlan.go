package efidevicepath

// https://uefi.org/specs/UEFI/2.10/10_Protocols_Device_Path_Protocol.html#vlan-device-path-node

const VlanType = 3 + 20*0x100

type Vlan struct {
	Vlanid uint16
}
