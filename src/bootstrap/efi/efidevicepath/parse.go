package efidevicepath

import (
	"bootstrap/efi/common"
	"encoding"
	"fmt"
)

func ParsePath[T encoding.BinaryUnmarshaler](data []byte) (T, error) {
	t := common.New[T]()
	if err := t.UnmarshalBinary(data); err != nil {
		return common.Nil[T](), fmt.Errorf("failed to unmarshal data: %w", err)
	}

	return t, nil
}
