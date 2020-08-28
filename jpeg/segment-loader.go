package jpeg

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/loganwishartcraig/go-exif/segment"
)

var (
	app0MarkerId   = []byte{0xff, 0xe0}
	app1MarkerId   = []byte{0xff, 0xe1}
	exifIdentifier = []byte{0x45, 0x78, 0x69, 0x66, 0x00, 0x00}
)

const exifSegmentHeaderLength = 0xA

type SegmentLoader struct {
	contents []byte
}

func NewSegmentLoader(contents []byte) *SegmentLoader {

	// Loader content will just ignore SOI
	return &SegmentLoader{contents[2:]}

}

func (l *SegmentLoader) LoadExifSegmentBody() (*segment.Segment, error) {

	offset, err := l.loadMarkerOffset(app1MarkerId)

	if err != nil {
		return nil, err
	}

	return l.loadFromOffset(offset)

}

func (l *SegmentLoader) loadMarkerOffset(id []byte) (int64, error) {

	reader := bytes.NewReader(l.contents)

	currentOffset := int64(0)
	currentMarkerInfo := make([]byte, 4)

	for {

		_, err := reader.ReadAt(currentMarkerInfo, currentOffset)

		if err != nil {
			return -1, err
		} else if err := validateMarkerId(currentMarkerInfo[:2]); err != nil {
			return -1, err
		} else if bytes.Equal(currentMarkerInfo[:2], id) {
			return currentOffset, nil
		}

		segmentLength := binary.BigEndian.Uint16(currentMarkerInfo[2:])

		// +2 to account for ID field
		currentOffset += int64(segmentLength) + 2

	}

}

func (l *SegmentLoader) loadFromOffset(offset int64) (*segment.Segment, error) {

	if err := validateExifSegmentHeader(l.contents[offset:]); err != nil {
		return nil, err
	}

	segmentLength := binary.BigEndian.Uint16(l.contents[offset+2 : offset+4])
	bodyOffset := offset + exifSegmentHeaderLength

	return &segment.Segment{
		Body: l.contents[bodyOffset : bodyOffset+int64(segmentLength)],
	}, nil

}

func validateMarkerId(id []byte) error {

	if !bytes.Equal(id, app0MarkerId) && !bytes.Equal(id, app1MarkerId) {
		return errors.New(fmt.Sprintf("Unknown marker ID %x", id))
	}

	return nil

}

func validateExifSegmentHeader(segment []byte) error {

	if len(segment) < exifSegmentHeaderLength {
		return errors.New(fmt.Sprintf("Exif header segment too small (len: %d)", len(segment)))
	} else if !bytes.Equal(segment[:2], app1MarkerId) {
		return errors.New(fmt.Sprintf("Invalid app1 marker ID [ % x ]", segment[:2]))
	} else if !bytes.Equal(segment[4:10], exifIdentifier) {
		return errors.New(fmt.Sprintf("Invalid exif identifier [ % x ]", segment[4:10]))
	}

	return nil

}
