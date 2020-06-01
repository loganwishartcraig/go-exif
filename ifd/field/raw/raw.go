package raw

import (
	"encoding/binary"
	"errors"

	"github.com/loganwishartcraig/go-exif/ifd"
)

type RawIfdField struct {
	TagId uint16
	Type  uint16
	Count uint32
	// ValueOffset uint32
	Contents []byte
}

type ifdFieldBlob struct {
	TagId       uint16
	Type        uint16
	Count       uint32
	ValueOffset uint32
}

func NewRawIfdField(ifd *ifd.Ifd, byteOrder binary.ByteOrder) (*RawIfdField, error) {

	// reader := bytes.NewReader(contents)
	// blob := &ifdFieldBlob{}

	// if err := binary.Read(reader, byteOrder, blob); err != nil {
	// 	return nil, err
	// }

	// return rawIfdField, nil

	return nil, errors.New("Not implemented")

}
