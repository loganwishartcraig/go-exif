package byte

import (
	"errors"
	"fmt"

	"github.com/loganwishartcraig/go-exif/ifd/field"
)

const FieldType field.IfdFieldType = 1

type ByteFieldInterpreter struct {
	field       *field.IfdField
	cachedValue byte
}

func NewByteFieldInterpreter(field *field.IfdField) (*ByteFieldInterpreter, error) {

	if field.Type != FieldType {
		return nil, errors.New(fmt.Sprintf("Cannot create byte field interpretor for field type %d. Expected field type %d.", field.Type, FieldType))
	}

	return &ByteFieldInterpreter{
		field,
		0,
	}, nil

}

// func (i *ByteFieldInterpreter) Value() byte {

// }
// func (i *ByteFieldInterpreter) Marshal(buff *byte) {}
