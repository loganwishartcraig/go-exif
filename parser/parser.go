package parser

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/loganwishartcraig/go-exif/ifd"
	"github.com/loganwishartcraig/go-exif/marker"
	"github.com/loganwishartcraig/go-exif/tags"
)

var (
	exifIdentifierCode     = []byte{0x45, 0x78, 0x69, 0x66}
	littleEndianIdentifier = []byte{0x49, 0x49}
	bigEndianIdentifier    = []byte{0x4D, 0x4D}
	fortyTwoIdentifier     = []byte{0x0, 0x2A}

	exifIdentiferOffset = 0x0

	headerDataOffset = 0x6
	byteOrderOffset  = 0x0
	fortyTwoOffset   = 0x2
	zerothIfdOffset  = 0x4
)

type ExifParser interface {
	ParseAll() ([](*tags.ExifTag), error)
	ParseTag(tag tags.ExifTagId) (*tags.ExifTag, error)
}

type BasicExifReader struct {
	marker          *marker.Marker
	byteOrder       binary.ByteOrder
	zerothIfdOffset uint32
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

	return &BasicExifReader{m, byteOrder, uint32(headerDataOffset) + zIdOffset}, nil

}

func (r *BasicExifReader) ParseAll() ([](*tags.ExifTag), error) {

	ifd, err := ifd.NewIfd(r.marker.Contents[r.zerothIfdOffset:], r.byteOrder, r.zerothIfdOffset)

	if err != nil {
		return nil, err
	}

	fmt.Println(ifd)
	// field, err := ifd.LoadField(0, r.byteOrder)
	fields, err := ifd.LoadAllFields(r.byteOrder)

	if err != nil {
		return nil, err
	}

	for _, r := range fields {
		fmt.Println(r)
	}

	return nil, errors.New("ParseAll - Not Implemented")
}

func (p *BasicExifReader) ParseTag() (*tags.ExifTag, error) {
	return nil, errors.New("ParseTag - Not Implemented")
}

func validateExifHeader(m *marker.Marker) error {

	offset := exifIdentiferOffset

	if !bytes.Equal(m.Contents[offset:offset+4], exifIdentifierCode) {
		return errors.New(fmt.Sprintf("Invalid EXIF marker. Expected <0x%x> received <0x%x>", exifIdentifierCode, m.Contents[offset:offset+4]))
	} else if err := validateByteOrder(m); err != nil {
		return err
	}

	return nil

}

func validateByteOrder(m *marker.Marker) error {

	offset := headerDataOffset + byteOrderOffset
	flag := m.Contents[offset : offset+2]

	isLittle := bytes.Equal(flag, littleEndianIdentifier)
	isBig := bytes.Equal(flag, bigEndianIdentifier)

	if !isLittle && !isBig {
		return errors.New(fmt.Sprintf("Unable to parse endianness. Expected <0x%x> or <0x%x>. Received <0x%x>", littleEndianIdentifier, bigEndianIdentifier, flag))
	}

	return nil

}

func parseByteOrder(m *marker.Marker) binary.ByteOrder {

	offset := headerDataOffset + byteOrderOffset
	value := m.Contents[offset : offset+2]

	if bytes.Equal(value, littleEndianIdentifier) {
		return binary.LittleEndian
	}

	return binary.BigEndian

}

func parseZerothIfdOffset(m *marker.Marker, byteOrder binary.ByteOrder) (uint32, error) {

	valueOffset := headerDataOffset + zerothIfdOffset
	value := m.Contents[valueOffset : valueOffset+4]

	reader := bytes.NewReader(value)
	var offset uint32 = 0
	err := binary.Read(reader, byteOrder, &offset)

	return offset, err

}
