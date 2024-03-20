package efidevicepath

import (
	"encoding/binary"
	"fmt"
	"unicode/utf16"
)

// https://uefi.org/specs/UEFI/2.10/10_Protocols_Device_Path_Protocol.html#file-path-media-device-path

const FilePathType = 4 + 4*0x100

type FilePath struct {
	PathName string
}

func (fp *FilePath) UnmarshalBinary(data []byte) error {
	if len(data) < 2 {
		return fmt.Errorf("unmarshal data is too short")
	}

	term := false
	var file []uint16

	for i := 0; i < len(data)-1; i += 2 {
		char := binary.LittleEndian.Uint16(data[i : i+2])

		if char == 0 {
			term = true

			break
		}

		file = append(file, char)
	}

	if !term {
		return fmt.Errorf("the \"PathName\" data doesn't have a null-terminated")
	}

	fp.PathName = string(utf16.Decode(file))

	return nil
}
