package squashfs

import (
	"encoding/binary"
	"fmt"
	"os"
)

type superBlock struct {
	Smagic              uint32
	Inodes              uint32
	MkfsTime            uint32
	BlockSize           uint32
	Fragments           uint32
	Compression         uint16
	BlockLog            uint16
	Flags               uint16
	NoIds               uint16
	SMajor              uint16
	SMinor              uint16
	RootInode           uint64
	BytesUsed           int64
	IdTableStart        int64
	XattrIdTableStart   int64
	InodeTableStart     int64
	DirectoryTableStart int64
	FragmentTableStart  int64
	LookupTableStart    int64
}

type Squashfs struct {
	Sb superBlock
}

func NewFromFile(path string) (*Squashfs, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	fs := &Squashfs{}
	if err := binary.Read(f, binary.LittleEndian, &fs.Sb); err != nil {
		return nil, err
	}
	if fs.Sb.Smagic != 0x73717368 {
		return nil, fmt.Errorf("unexpected super block magic: %x", fs.Sb.Smagic)
	}

	return fs, nil
}
