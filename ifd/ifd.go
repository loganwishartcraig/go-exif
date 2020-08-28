package ifd

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/loganwishartcraig/go-exif/ifd/block"
	"github.com/loganwishartcraig/go-exif/segment"
)

type Table struct {
	BaseOffset int
	ByteOrder  binary.ByteOrder
	Blocks     []*block.Block
}

type header struct {
	FortyTwo        uint16
	ZerothIfdOffset uint32
}

const (
	byteOrderOffset       = 0x0
	fortyTwoOffset        = 0x2
	zerothIfdOffsetOffset = 0x4
)

const (
	littleEndianFlag = "II"
	bigEndianFlag    = "MM"
)

type ExifSegmentLoader interface {
	LoadExifSegmentBody() (*segment.Segment, error)
}

func NewTable(l ExifSegmentLoader) (*Table, error) {

	segment, err := l.LoadExifSegmentBody()

	if err != nil {
		return nil, err
	}

	order, err := readSegmentByteOrder(segment)

	if err != nil {
		return nil, err
	}

	zerothIfdOffset, err := readZerothIfdOffset(segment, order)

	if err != nil {
		return nil, err
	}

	currentOffset := zerothIfdOffset
	blocks := make([]*block.Block, 0)

	for currentOffset != 0 {
		block, err := block.NewBlock(segment, currentOffset, order)
		if err != nil {
			return nil, err
		}
		blocks = append(blocks, block)
		currentOffset = block.NextIfdOffset
	}

	return &Table{
		BaseOffset: 0,
		ByteOrder:  order,
		Blocks:     blocks,
	}, nil

}

func readSegmentByteOrder(s *segment.Segment) (binary.ByteOrder, error) {

	order := string(s.Body[:2])

	if order == littleEndianFlag {
		return binary.LittleEndian, nil
	} else if order == bigEndianFlag {
		return binary.BigEndian, nil
	}

	return nil, fmt.Errorf("Unable to parse unknown byte order flag [ %X ]", order)

}

func readZerothIfdOffset(s *segment.Segment, order binary.ByteOrder) (int, error) {

	var rawOffset uint32

	if err := binary.Read(bytes.NewReader(s.Body[zerothIfdOffsetOffset:]), order, &rawOffset); err != nil {
		return -1, err
	}

	return int(rawOffset), nil

}
