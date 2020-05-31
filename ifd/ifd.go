package ifd

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

// Image File Directory (ifd)

type Ifd struct {
	FieldCount uint16
	Content    []byte
	Offset     uint32
	NextOffset uint32
	fieldSize  uint16
}

var (
	defaultFieldSize uint16 = 12
)

func NewIfd(b []byte, byteOrder binary.ByteOrder, offset uint32) (*Ifd, error) {

	length := len(b)

	if length < 6 {
		return nil, errors.New("Invalid Ifd length")
	}

	fmt.Printf("Count %x\n", b[0:2])
	fmt.Printf("Next Offset %x\n", b[length-4:length])
	fmt.Println("NewIfd - Should parse into RawIfd struct")

	fcReader := bytes.NewReader(b[0:2])

	var fieldCount uint16 = 0
	var nextOffset uint32 = 0

	if err := binary.Read(fcReader, byteOrder, &fieldCount); err != nil {
		return nil, err
	}

	endOfContentOffset := (fieldCount * defaultFieldSize)
	content := b[2:endOfContentOffset]
	nextOffsetReader := bytes.NewReader(b[endOfContentOffset : endOfContentOffset+4])

	fmt.Printf("Content %x\n", content)

	if err := binary.Read(nextOffsetReader, byteOrder, &nextOffset); err != nil {
		return nil, err
	}

	return &Ifd{
		fieldCount,
		content,
		offset,
		nextOffset,
		defaultFieldSize,
	}, nil

}

func (ifd *Ifd) String() string {
	return fmt.Sprintf("Ifd - Count %d", ifd.FieldCount)
}

func (ifd *Ifd) LoadField(index uint16, byteOrder binary.ByteOrder) (*IfdField, error) {

	if index > ifd.FieldCount {
		return nil, errors.New(fmt.Sprintf("Index out of bounds [0, %d]. Received %d", ifd.FieldCount-1, index))
	}

	offset := ifd.fieldSize * index
	return NewIfdField(ifd.Content[offset:offset+ifd.fieldSize], byteOrder, ifd.Offset)

}

func (ifd *Ifd) LoadAllFields(byteOrder binary.ByteOrder) ([](*IfdField), error) {

	fields := make([](*IfdField), ifd.FieldCount)

	var i uint16

	for i = 0; i < ifd.FieldCount; i++ {

		field, err := ifd.LoadField(i, byteOrder)

		if err != nil {
			return nil, err
		}

		fields[i] = field

	}

	return fields, nil

}

type IfdMarkerId []byte

var (
	AppMarker0Id IfdMarkerId = []byte{0xff, 0xe0}
	AppMarker1Id             = []byte{0xff, 0xe1}
)

type IfdFieldType uint16

const (
	Byte      IfdFieldType = 1
	Ascii                  = 2
	Short                  = 3
	Long                   = 4
	Rational               = 5
	Undefined              = 7
	SLong                  = 9
	SRational              = 10
)

var fieldSize = map[IfdFieldType]uint32{
	Byte:      1,
	Ascii:     1,
	Short:     2,
	Long:      4,
	Rational:  8,
	Undefined: 1,
	SLong:     4,
	SRational: 8,
}

var typeLabel = map[IfdFieldType]string{
	Byte:      "Byte",
	Ascii:     "Ascii",
	Short:     "Short",
	Long:      "Long",
	Rational:  "Rational",
	Undefined: "Undefined",
	SLong:     "SLong",
	SRational: "SRational",
}

type RawIfdField struct {
	TagId       uint16
	Type        IfdFieldType
	Count       uint32
	ValueOffset uint32
}

type IfdFieldValue interface{}

type IfdField struct {
	TagId       uint16
	Type        IfdFieldType
	TypeLabel   string
	Count       uint32
	Size        uint32
	Offset      uint32
	ValueOffset uint32
	Value       IfdFieldValue
}

type IfdSet []IfdField

func NewIfdField(b []byte, byteOrder binary.ByteOrder, offset uint32) (*IfdField, error) {

	reader := bytes.NewReader(b)
	rawIfdField := &RawIfdField{}

	if err := binary.Read(reader, byteOrder, rawIfdField); err != nil {
		return nil, err
	}

	fmt.Println("NewIfdField - Should validate field type")

	size, sizeInMap := fieldSize[rawIfdField.Type]
	typeLabel, typeLabelInMap := typeLabel[rawIfdField.Type]

	if !sizeInMap {
		size = 1
	}
	if !typeLabelInMap {
		typeLabel = "UNKNOWN"
	}

	return &IfdField{
		TagId:       rawIfdField.TagId,
		Type:        rawIfdField.Type,
		TypeLabel:   typeLabel,
		Count:       rawIfdField.Count,
		Size:        rawIfdField.Count * size,
		Offset:      offset,
		ValueOffset: rawIfdField.ValueOffset,
	}, nil

}

func (field *IfdField) String() string {
	return fmt.Sprintf("IfdField \t ID <0x%x> \t %s x %d \t ValueOffset <0x%x>", field.TagId, field.TypeLabel, field.Count, field.ValueOffset)
}

func (field *IfdField) NewValueBuffer() IfdFieldValue {

	switch field.Type {
	case Byte:
		if field.Count > 1 {
			return make([]byte, field.Count)
		} else {
			return new(byte)
		}
	case Ascii:
		return new(string)
	case Short:
		if field.Count > 1 {
			return make([]uint16, field.Count)
		} else {
			return new(uint16)
		}
	case Long:
		if field.Count > 1 {
			return make([]uint32, field.Count)
		} else {
			return new(uint32)
		}
	case Rational:
		if field.Count > 1 {
			return make([]float64, field.Count)
		} else {
			return new(float64)
		}
	case Undefined:
		if field.Count > 1 {
			return make([]byte, field.Count)
		} else {
			return new(byte)
		}
	case SLong:
		if field.Count > 1 {
			return make([]int32, field.Count)
		} else {
			return new(int32)
		}
	case SRational:
		if field.Count > 1 {
			return make([]float64, field.Count)
		} else {
			return new(float64)
		}
	default:
		if field.Count > 1 {
			return make([]byte, field.Count)
		} else {
			return new(byte)
		}
	}

}
