package efivarfs

import (
	"encoding/binary"
	"fmt"
	"unicode/utf16"
)

// https://uefi.org/specs/UEFI/2.10/10_Protocols_Device_Path_Protocol.html#efi-device-path-protocol

type FilePath struct {
	Type uint16
	Data []byte
}

type FilePathList []FilePath

func (fpl *FilePathList) AppendFilePathFromBinary(data []byte) ([]byte, error) {
	if len(data) == 0 {
		return nil, nil
	}

	if len(data) < 4 {
		return nil, fmt.Errorf("unmarshal data is too short")
	}

	t := binary.LittleEndian.Uint16(data[0:2])
	l := binary.LittleEndian.Uint16(data[2:4])

	if len(data) < int(l) {
		return nil, fmt.Errorf("unmarshal data is too short")
	}

	fp := FilePath{
		Type: t,
		Data: data[4:l],
	}

	*fpl = append(*fpl, fp)

	return data[l:], nil
}

func (fpl *FilePathList) UnmarshalBinary(data []byte) (err error) {
	for len(data) > 0 {
		if data, err = fpl.AppendFilePathFromBinary(data); err != nil {
			return fmt.Errorf("the \"FilePathList\" data parse error occured: %w", err)
		}
	}

	return nil
}

// https://uefi.org/specs/UEFI/2.10/03_Boot_Manager.html#load-options
type LoadOption struct {
	Attributes         uint32
	FilePathListLength uint16
	Description        string
	FilePathList       FilePathList
	OptionalData       []uint8
}

func (lo *LoadOption) UnmarshalBinary(data []byte) error {
	if len(data) < 10 {
		return fmt.Errorf("unmarshal data is too short")
	}

	lo.Attributes = binary.LittleEndian.Uint32(data[4:8])
	lo.FilePathListLength = binary.LittleEndian.Uint16(data[8:10])

	term := false
	var offset int
	var desc []uint16

	for offset = 10; offset < len(data)-1; offset += 2 {
		char := binary.LittleEndian.Uint16(data[offset : offset+2])

		if char == 0 {
			term = true
			offset += 2

			break
		}

		desc = append(desc, char)
	}

	if !term {
		return fmt.Errorf("the \"Description\" data doesn't have a null-terminated")
	}

	lo.Description = string(utf16.Decode(desc))

	if len(data) < offset+int(lo.FilePathListLength) {
		return fmt.Errorf("unmarshal data is too short")
	}

	err := lo.FilePathList.UnmarshalBinary(data[offset : offset+int(lo.FilePathListLength)])
	if err != nil {
		return fmt.Errorf("the \"FilePathList\" data parse error occured: %w", err)
	}

	lo.OptionalData = data[offset+int(lo.FilePathListLength):]

	return nil
}
