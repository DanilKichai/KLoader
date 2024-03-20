package main

import (
	"bootstrap/efi/efidevicepath"
	"bootstrap/efi/efivarfs"
	"fmt"
)

func main() {
	current, err := efivarfs.ParseVar[*efivarfs.BootCurrent]("BootCurrent", efivarfs.GlobalVariable)
	if err != nil {
		fmt.Println(err.Error())

		return
	}

	entry, err := efivarfs.ParseVar[*efivarfs.LoadOption](fmt.Sprintf("Boot%04X", *current), efivarfs.GlobalVariable)
	if err != nil {
		fmt.Println(err.Error())

		return
	}

	fmt.Println(entry.Description)

	disk, err := efidevicepath.ParsePath[*efidevicepath.HardDrive](entry.FilePathList[0].Data)
	if err != nil {
		fmt.Println(err.Error())

		return
	}

	fmt.Printf("PartitionNumber: %d, PartitionStart: %d, PartitionSize: %d\n", disk.PartitionNumber, disk.PartitionStart, disk.PartitionSize)

	file, err := efidevicepath.ParsePath[*efidevicepath.FilePath](entry.FilePathList[1].Data)
	if err != nil {
		fmt.Println(err.Error())

		return
	}

	fmt.Println(file.PathName)
}
