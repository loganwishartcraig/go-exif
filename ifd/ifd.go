package ifd

import (
	"bytes"
	"encoding/binary"
	"fmt"

	exif "github.com/loganwishartcraig/go-exif/exif/reader"
)

// Image File Directory (ifd)
type Ifd struct {
	FieldCount uint16
	NextOffset uint32
	ByteOrder  binary.ByteOrder

	contentReader *bytes.Reader
	fieldSize     uint16
}

const (
	defaultFieldSize uint16 = 12

	fieldCountFieldSize = 2
	nextOffsetFieldSize = 2
)

func NewIfd(reader exif.Reader, byteOrder binary.ByteOrder) (*Ifd, error) {

	var fieldCount uint16

	fmt.Println("Need to seek to beginning of reader")

	if err := binary.Read(reader, byteOrder, &fieldCount); err != nil {
		return nil, err
	}

	contentBuffer := make([]byte, fieldCount*defaultFieldSize)

	if _, err := reader.Read(contentBuffer); err != nil {
		return nil, err
	}

	var nextOffset uint32

	if err := binary.Read(reader, byteOrder, &nextOffset); err != nil {
		return nil, err
	}

	return &Ifd{
		fieldCount,
		nextOffset,
		byteOrder,
		bytes.NewReader(contentBuffer),
		defaultFieldSize,
	}, nil

}

func (ifd *Ifd) String() string {
	return fmt.Sprintf("Ifd - Count %d", ifd.FieldCount)
}

// func (ifd *Ifd) LoadField(index uint16, byteOrder binary.ByteOrder) (*IfdField, error) {

// 	if index > ifd.FieldCount {
// 		return nil, errors.New(fmt.Sprintf("Index out of bounds [0, %d]. Received %d", ifd.FieldCount-1, index))
// 	}

// 	offset := ifd.fieldSize * index
// 	return NewIfdField(ifd.Content[offset:offset+ifd.fieldSize], byteOrder, ifd.BaseOffset)

// }

// func (ifd *Ifd) LoadAllFields(byteOrder binary.ByteOrder) ([](*IfdField), error) {

// 	fields := make([](*IfdField), ifd.FieldCount)

// 	var i uint16

// 	for i = 0; i < ifd.FieldCount; i++ {

// 		field, err := ifd.LoadField(i, byteOrder)

// 		if err != nil {
// 			return nil, err
// 		}

// 		fields[i] = field

// 	}

// 	return fields, nil

// }

// var fieldSize = map[IfdFieldType]uint32{
// 	Byte:      1,
// 	Ascii:     1,
// 	Short:     2,
// 	Long:      4,
// 	Rational:  8,
// 	Undefined: 1,
// 	SLong:     4,
// 	SRational: 8,
// }

// var typeLabel = map[IfdFieldType]string{
// 	Byte:      "Byte",
// 	Ascii:     "Ascii",
// 	Short:     "Short",
// 	Long:      "Long",
// 	Rational:  "Rational",
// 	Undefined: "Undefined",
// 	SLong:     "SLong",
// 	SRational: "SRational",
// }

// func NewIfdField(b []byte, byteOrder binary.ByteOrder, offset uint32) (*IfdField, error) {

// 	reader := bytes.NewReader(b)
// 	rawIfdField := &RawIfdField{}

// 	if err := binary.Read(reader, byteOrder, rawIfdField); err != nil {
// 		return nil, err
// 	}

// 	fmt.Println("NewIfdField - Should validate field type")

// 	size, sizeInMap := fieldSize[rawIfdField.Type]
// 	typeLabel, typeLabelInMap := typeLabel[rawIfdField.Type]

// 	if !sizeInMap {
// 		size = 1
// 	}
// 	if !typeLabelInMap {
// 		typeLabel = "UNKNOWN"
// 	}

// 	return &IfdField{
// 		TagId:       rawIfdField.TagId,
// 		Type:        rawIfdField.Type,
// 		TypeLabel:   typeLabel,
// 		Count:       rawIfdField.Count,
// 		Size:        rawIfdField.Count * size,
// 		Offset:      offset,
// 		ValueOffset: rawIfdField.ValueOffset,
// 	}, nil

// }

// func (field *IfdField) String() string {
// 	return fmt.Sprintf("IfdField \t ID <0x%x> \t %s x %d \t ValueOffset <0x%x>", field.TagId, field.TypeLabel, field.Count, field.ValueOffset)
// }

// func (field *IfdField) NewValueBuffer() IfdFieldValue {

// 	switch field.Type {
// 	case Byte:
// 		if field.Count > 1 {
// 			return make([]byte, field.Count)
// 		} else {
// 			return new(byte)
// 		}
// 	case Ascii:
// 		return new(string)
// 	case Short:
// 		if field.Count > 1 {
// 			return make([]uint16, field.Count)
// 		} else {
// 			return new(uint16)
// 		}
// 	case Long:
// 		if field.Count > 1 {
// 			return make([]uint32, field.Count)
// 		} else {
// 			return new(uint32)
// 		}
// 	case Rational:
// 		if field.Count > 1 {
// 			return make([]float64, field.Count)
// 		} else {
// 			return new(float64)
// 		}
// 	case Undefined:
// 		if field.Count > 1 {
// 			return make([]byte, field.Count)
// 		} else {
// 			return new(byte)
// 		}
// 	case SLong:
// 		if field.Count > 1 {
// 			return make([]int32, field.Count)
// 		} else {
// 			return new(int32)
// 		}
// 	case SRational:
// 		if field.Count > 1 {
// 			return make([]float64, field.Count)
// 		} else {
// 			return new(float64)
// 		}
// 	default:
// 		if field.Count > 1 {
// 			return make([]byte, field.Count)
// 		} else {
// 			return new(byte)
// 		}
// 	}

// }

// func (field *IfdField) Marshal(valueBuf *IfdFieldValue) error {

// 	valueStart := field.Offset + field.ValueOffset

// 	rawValueBytes = field.

// 	// WANT TO SPLIT THE IfdField into seperate types that implement
// 	// a common interface

// }
