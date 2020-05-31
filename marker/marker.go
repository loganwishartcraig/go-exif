package marker

import (
	"encoding/binary"
	"errors"
	"fmt"
	"os"
)

type Marker struct {
	Id       []byte
	Length   uint16
	Offset   int64
	Contents []byte
}

func NewMarker(f *os.File, offset int64) (*Marker, error) {

	markerMeta := make([]byte, 4)

	if _, err := f.ReadAt(markerMeta, offset); err != nil {
		return nil, err
	}

	id := markerMeta[0:2]
	lengthBytes := markerMeta[2:4]
	length := binary.BigEndian.Uint16(lengthBytes)

	if !isValidAppMarkerId(id) {
		return nil, errors.New(fmt.Sprintf("Invalid AppMarker ID <0x%x> at given offset %d", id, offset))
	}

	markerContents := make([]byte, (int)(length)-len(lengthBytes))

	if _, err := f.ReadAt(markerContents, offset+(int64)(len(markerMeta))); err != nil {
		return nil, err
	}

	return &Marker{
		id,
		length,
		offset,
		markerContents,
	}, nil

}

func isValidAppMarkerId(bytePair []byte) bool {

	if len(bytePair) != 2 {
		return false
	}

	return bytePair[0] == 0xff && bytePair[1] >= 0xe0 && bytePair[1] <= 0xe9

}

func (m *Marker) String() string {
	return fmt.Sprintf("Marker: <0x%x> - length %d - offset %d", m.Id, m.Length, m.Offset)
}
