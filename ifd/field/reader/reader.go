package reader

import (
	"encoding/binary"
	"fmt"

	"github.com/loganwishartcraig/go-exif/ifd/field"
	"github.com/loganwishartcraig/go-exif/segment"
	"github.com/loganwishartcraig/go-exif/tags"
)

type FieldReader struct{}

func NewFieldReader() *FieldReader {
	return &FieldReader{}
}

func (r *FieldReader) Print(f *field.Field, s *segment.Segment) {

	switch f.Type {
	case tags.Byte:
		readByte(f, s)
	case tags.Ascii:
		readString(f, s)
	case tags.Short:
		readShort(f, s)
	case tags.Long:
		readLong(f, s)
	case tags.Rational:
		readRational(f, s)
	case tags.Undefined:
		readUndefined(f, s)
	case tags.SLong:
		readSLong(f, s)
	case tags.SRational:
		readSRational(f, s)
	default:
		fmt.Printf("Unable to read unknown type %d\n", f.Type)
	}

}

func readByte(f *field.Field, s *segment.Segment) {
	buff := []byte{}
	if f.ValueLen <= 4 {
		buff = append(buff, byte(f.ValueOffset))
	} else {
		buff = append(buff, s.Body[f.ValueOffset:f.ValueOffset+f.ValueLen]...)
	}
	fmt.Printf("[%s] - %s\n", f, buff)
}
func readString(f *field.Field, s *segment.Segment) {
	buff := s.Body[f.ValueOffset : f.ValueOffset+f.ValueLen]
	fmt.Printf("[%s] - %s\n", f, buff)
}
func readShort(f *field.Field, s *segment.Segment) {
	var buff uint16
	if f.ValueLen <= 4 {
		buff = uint16(f.ValueOffset)
	} else {
		buff = binary.BigEndian.Uint16(s.Body[f.ValueOffset : f.ValueOffset+f.ValueLen])
	}
	fmt.Printf("[%s] - %d\n", f, buff)
}
func readLong(f *field.Field, s *segment.Segment) {
	var buff uint32
	if f.ValueLen <= 4 {
		buff = uint32(f.ValueOffset)
	} else {
		buff = binary.BigEndian.Uint32(s.Body[f.ValueOffset : f.ValueOffset+f.ValueLen])
	}
	fmt.Printf("[%s] - %d\n", f, buff)
}

func readRational(f *field.Field, s *segment.Segment) {
	var buff float64
	buff = float64(binary.BigEndian.Uint32(s.Body[f.ValueOffset:f.ValueOffset+f.ValueLen])) / float64(binary.BigEndian.Uint32(s.Body[f.ValueOffset+4:f.ValueOffset+f.ValueLen+4]))
	fmt.Printf("[%s] - %f\n", f, buff)
}
func readUndefined(f *field.Field, s *segment.Segment) {
	fmt.Printf("[%s] - 'Undefined' field types not supported.", f)
}
func readSLong(f *field.Field, s *segment.Segment) {
	var buff int32
	if f.ValueLen <= 4 {
		buff = int32(f.ValueOffset)
	} else {
		buff = int32(binary.BigEndian.Uint32(s.Body[f.ValueOffset : f.ValueOffset+f.ValueLen]))
	}
	fmt.Printf("[%s] - %d\n", f, buff)
}
func readSRational(f *field.Field, s *segment.Segment) {
	var buff float64
	buff = float64(int32(binary.BigEndian.Uint32(s.Body[f.ValueOffset:f.ValueOffset+f.ValueLen]))) / float64(int32(binary.BigEndian.Uint32(s.Body[f.ValueOffset+4:f.ValueOffset+f.ValueLen+4])))
	fmt.Printf("[%s] - %f\n", f, buff)
}
