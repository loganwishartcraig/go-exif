package marker

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"os"
)

type Marker struct {
	Id     []byte
	Length uint16
	Offset int64
	reader *bytes.Reader
}

type Reader interface {
	Read(b []byte) (int, error)
	ReadAt(b []byte, offset int64) (int, error)
	Seek(offset int64, whence int) (int64, error)
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

	reader := bytes.NewReader(markerContents)

	return &Marker{
		id,
		length,
		offset,
		reader,
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

func (m *Marker) Read(b []byte) (int, error) {
	return m.reader.Read(b)
}

func (m *Marker) ReadAt(b []byte, offset int64) (int, error) {
	return m.reader.ReadAt(b, offset)
}

func (m *Marker) Seek(offset int64, whence int) (int64, error) {
	return m.reader.Seek(offset, whence)
}
