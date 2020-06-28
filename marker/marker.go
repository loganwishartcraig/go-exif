package marker

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

type Id []byte

var (
	AppMarker0Id Id = []byte{0xff, 0xe0}
	AppMarker1Id    = []byte{0xff, 0xe1}
)

type Marker struct {
	Id       Id
	Contents []byte
}

func (m *Marker) String() string {
	return fmt.Sprintf("Marker ID: %x \t Length: %d", m.Id, len(m.Contents))
}

type Loader interface {
	LoadAll() ([]*Marker, error)
	Load(id Id) (*Marker, error)
}

type JpegLoader struct {
	contents []byte
}

func NewJpegLoader(contents []byte) *JpegLoader {

	// Loader content will just ignore SOI
	return &JpegLoader{contents[2:]}

}

func (l *JpegLoader) LoadAll() ([]*Marker, error) {
	return nil, errors.New("Not implemented")
}

func (l *JpegLoader) Load(id Id) (*Marker, error) {

	markerOffset, err := l.findMarkerOffset(id)

	if err != nil {
		return nil, err
	}

	return l.loadAtOffset(markerOffset)

}

func (l *JpegLoader) findMarkerOffset(id Id) (int, error) {

	currentOffset := 0
	var currentMarkerId []byte

	for currentOffset < len(l.contents)-2 {

		currentMarkerId = l.contents[currentOffset : currentOffset+2]

		fmt.Printf("Finding Marker... current offset %d - ID: %x\n", currentOffset, currentMarkerId)

		if bytes.Equal(currentMarkerId, id) {
			return currentOffset, nil
		}

		// Skip ID field
		currentOffset += 2

		// Get marker length
		markerLen := l.contents[currentOffset : currentOffset+2]

		// Skip marker
		currentOffset += int(binary.BigEndian.Uint16(markerLen))

	}

	return -1, errors.New(fmt.Sprintf("Marker ID %x not found", id))

}

func (l *JpegLoader) loadAtOffset(offset int) (*Marker, error) {

	reader := bytes.NewReader(l.contents[offset:])

	markerId := make([]byte, 2)
	var contentLength uint16

	if _, err := reader.Read(markerId); err != nil {
		return nil, err
	} else if err := binary.Read(reader, binary.BigEndian, &contentLength); err != nil {
		return nil, err
	}

	contents := l.contents[offset+4 : offset+4+int(contentLength)]

	return &Marker{
		Id:       l.contents[offset : offset+2],
		Contents: contents,
	}, nil

}
