package exif

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/loganwishartcraig/go-exif/marker"
)

type Reader interface {
	Read(b []byte) (int, error)
	ReadAt(b []byte, offset int64) (int, error)
	Seek(offset int64, whence int) (int64, error)
}

var (
	exifIdentifierCode     = []byte{0x45, 0x78, 0x69, 0x66}
	littleEndianIdentifier = []byte{0x49, 0x49}
	bigEndianIdentifier    = []byte{0x4D, 0x4D}
	fortyTwoIdentifier     = []byte{0x0, 0x2A}
)

const (
	exifIdentiferOffset int64 = 0x0
	headerDataOffset    int64 = 0x6
	byteOrderOffset     int64 = 0x0
	fortyTwoOffset      int64 = 0x2
	zerothIfdOffset     int64 = 0x4
)

type BasicExifReader struct {
	markerReader    marker.Reader
	ByteOrder       binary.ByteOrder
	zerothIfdOffset int64
}

func NewBasicExifReader(m *marker.Marker) (*BasicExifReader, error) {

	if err := validateExifHeader(m); err != nil {
		return nil, err
	}

	byteOrder := parseByteOrder(m)
	zIdOffset, err := parseZerothIfdOffset(m, byteOrder)

	if err != nil {
		return nil, err
	}

	return &BasicExifReader{m, byteOrder, headerDataOffset + zIdOffset}, nil

}

func validateExifHeader(m *marker.Marker) error {

	offset := exifIdentiferOffset
	idBuffer := make([]byte, 4)

	if _, err := m.ReadAt(idBuffer, offset); err != nil {
		return err
	} else if !bytes.Equal(idBuffer, exifIdentifierCode) {
		return errors.New(fmt.Sprintf("Invalid EXIF marker. Expected <0x%x> received <0x%x>", exifIdentifierCode, idBuffer))
	} else if err := validateByteOrder(m); err != nil {
		return err
	}

	return nil

}

func validateByteOrder(m *marker.Marker) error {

	offset := headerDataOffset + byteOrderOffset
	flagBuffer := make([]byte, 2)

	if _, err := m.ReadAt(flagBuffer, offset); err != nil {
		return err
	}

	isLittle := bytes.Equal(flagBuffer, littleEndianIdentifier)
	isBig := bytes.Equal(flagBuffer, bigEndianIdentifier)

	if !isLittle && !isBig {
		return errors.New(fmt.Sprintf("Unable to parse endianness. Expected <0x%x> or <0x%x>. Received <0x%x>", littleEndianIdentifier, bigEndianIdentifier, flagBuffer))
	}

	return nil

}

func parseByteOrder(m *marker.Marker) binary.ByteOrder {

	offset := headerDataOffset + byteOrderOffset
	byteOrderBuffer := make([]byte, 2)

	m.ReadAt(byteOrderBuffer, offset)

	if bytes.Equal(byteOrderBuffer, littleEndianIdentifier) {
		return binary.LittleEndian
	}

	return binary.BigEndian

}

func parseZerothIfdOffset(m *marker.Marker, byteOrder binary.ByteOrder) (int64, error) {

	valueOffset := headerDataOffset + zerothIfdOffset
	valueBuffer := make([]byte, 4)

	if _, err := m.ReadAt(valueBuffer, valueOffset); err != nil {
		return 0, err
	}

	reader := bytes.NewReader(valueBuffer)
	var offset uint32 = 0
	err := binary.Read(reader, byteOrder, &offset)

	return int64(offset), err

}

func (r *BasicExifReader) String() string {
	return fmt.Sprintf("ExifReader - 0th ID offset <0x%x>", r.zerothIfdOffset)
}

func (r *BasicExifReader) Read(b []byte) (int, error) {
	return r.markerReader.ReadAt(b, r.zerothIfdOffset)
}

func (r *BasicExifReader) ReadAt(b []byte, offset int64) (int, error) {
	fmt.Println("WARN - Need to make sure negative offsets work correctly")
	return r.markerReader.ReadAt(b, r.zerothIfdOffset+offset)
}

func (r *BasicExifReader) Seek(offset int64, whence int) (int64, error) {
	return r.markerReader.Seek(offset, whence)
}
