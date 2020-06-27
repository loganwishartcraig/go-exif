package ifd

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"

	exif "github.com/loganwishartcraig/go-exif/exif/reader"
	"github.com/loganwishartcraig/go-exif/ifd/tags"
)

// Image File Directory (ifd)
type Ifd struct {
	FieldCount uint16
	NextOffset uint32
	ByteOrder  binary.ByteOrder

	contentReader *bytes.Reader
	fieldSize     uint16

	// TODO: implement
	fields map[tags.TagId]*Field
	// fields field.Field[] - basic field information, DOES NOT include value.
	// Then can have interface FieldReader { Read(field: Field, valueContent: []byte, interface{}) }
}

type Reader interface {
	Read(b []byte) (int, error)
	ReadAt(b []byte, offset int64) (int, error)
	ReadField(b []byte, tagId tags.TagId) error
	Seek(offset int64, whence int) (int64, error)
}

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

type Field struct {
	TagId       tags.TagId
	Type        IfdFieldType
	Count       uint32
	ValueOffset uint32
}

const (
	FieldSize uint16 = 12

	fieldCountFieldSize = 2
	nextOffsetFieldSize = 2
)

func NewIfd(reader exif.Reader, byteOrder binary.ByteOrder, offset int64) (*Ifd, error) {

	var fieldCount uint16

	reader.Seek(offset, 0)

	if err := binary.Read(reader, byteOrder, &fieldCount); err != nil {
		return nil, err
	}

	contentBuffer := make([]byte, fieldCount*FieldSize)

	if _, err := reader.Read(contentBuffer); err != nil {
		return nil, err
	}

	var nextOffset uint32

	if err := binary.Read(reader, byteOrder, &nextOffset); err != nil {
		return nil, err
	}

	contentReader := bytes.NewReader(contentBuffer)
	fields := make(map[tags.TagId]*Field)

	// TODO: Extract this into goroutine?
	for i := uint16(0); i < fieldCount; i++ {
		field, err := NewField(contentReader, byteOrder, int64(i*FieldSize))
		if err != nil {
			return nil, err
		}
		fields[field.TagId] = field
		fmt.Println(field)
		test := make([]byte, field.Count)
		reader.ReadAt(test, int64(field.ValueOffset+6))
		fmt.Printf("Test data output: %s\n\n", test)
	}

	return &Ifd{
		fieldCount,
		nextOffset,
		byteOrder,
		contentReader,
		FieldSize,
		fields,
	}, nil

}

func (ifd *Ifd) String() string {
	contentPreview := make([]byte, 25)
	n, err := ifd.contentReader.Read(contentPreview)
	if err != nil {
		ifd.contentReader.Seek(int64(-1*n), 1)
	}
	return fmt.Sprintf("Ifd - Count %d - Content Partial (len: %d bytes) <0x%x...> - Next IFD Offset <0x%x>", ifd.FieldCount, ifd.contentReader.Size(), contentPreview, ifd.NextOffset)
}

func (ifd *Ifd) Read(b []byte) (int, error) {
	return ifd.contentReader.Read(b)
}
func (ifd *Ifd) ReadAt(b []byte, offset int64) (int, error) {
	return ifd.contentReader.ReadAt(b, offset)
}
func (ifd *Ifd) Seek(offset int64, whence int) (int64, error) {
	return ifd.contentReader.Seek(offset, whence)
}

func NewField(r io.ReadSeeker, byteOrder binary.ByteOrder, offset int64) (*Field, error) {

	r.Seek(offset, 0)

	field := &Field{}

	err := binary.Read(r, byteOrder, field)

	if err != nil {
		return nil, err
	}

	return field, nil

}

func (f *Field) String() string {
	return fmt.Sprintf("Ifd.Field: Tag <0x%x>\tType %d\tCount %d\tValueOffset <0x%x>", f.TagId, f.Type, f.Count, f.ValueOffset)
}
