package field

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/loganwishartcraig/go-exif/tags"
)

type Field struct {
	Tag         tags.Id
	Type        tags.Type
	Label       string
	Count       int
	ValueOffset int
	FieldOffset int
	ValueLen    int
}

type rawField struct {
	Tag         uint16
	Type        uint16
	Count       uint32
	ValueOffset uint32
}

const FieldSize = 12

func (f *rawField) ToField(fieldOffset int) *Field {

	t := tags.Type(f.Type)
	c := int(f.Count)
	id := tags.Id(f.Tag)

	return &Field{
		tags.Id(id),
		t,
		tags.Label(id),
		c,
		int(f.ValueOffset),
		fieldOffset,
		tags.GetByteLength(t, c),
	}

}

func (f *rawField) String() string {
	return fmt.Sprintf("Raw Field: Tag 0x%x | Type %d | Count %d | Value Offset %x", f.Tag, f.Type, f.Count, f.ValueOffset)
}

func ReadCount(r io.Reader, count int, order binary.ByteOrder, baseOffset int) ([]*Field, error) {

	fields := make([]*Field, 0, count)

	for i := 0; i < count; i++ {

		raw := &rawField{}
		if err := binary.Read(r, order, raw); err != nil {
			return nil, err
		}

		fields = append(fields, raw.ToField(baseOffset+i*FieldSize))

	}

	return fields, nil

}

func (f *Field) String() string {
	return fmt.Sprintf("Field: %s 0x%x", f.Label, f.Tag)
}
