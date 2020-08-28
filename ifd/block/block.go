package block

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/loganwishartcraig/go-exif/ifd/field"
	"github.com/loganwishartcraig/go-exif/segment"
)

type Block struct {
	BaseOffset    int
	Fields        []*field.Field
	NextIfdOffset int
}

func NewBlock(s *segment.Segment, baseOffset int, order binary.ByteOrder) (*Block, error) {

	var fieldCount uint16
	var nextIfdOffset uint32

	r := bytes.NewReader(s.Body[baseOffset:])

	if err := binary.Read(r, order, &fieldCount); err != nil {
		return nil, err
	}

	fields, err := field.ReadCount(r, int(fieldCount), order, baseOffset+2)

	if err != nil {
		return nil, err
	}

	if err := binary.Read(r, order, &nextIfdOffset); err != nil {
		return nil, err
	}

	return &Block{
		BaseOffset:    baseOffset,
		Fields:        fields,
		NextIfdOffset: int(nextIfdOffset),
	}, nil

}

func (b *Block) String() string {
	return fmt.Sprintf("Block: Offset [ % X ] | Field Count: %d | Next IFD Offset: [ % X ]", b.BaseOffset, len(b.Fields), b.NextIfdOffset)
}
