package exif

import (
	exif "github.com/loganwishartcraig/go-exif/exif/reader"
	"github.com/loganwishartcraig/go-exif/ifd/field"
	"github.com/loganwishartcraig/go-exif/ifd/tags"
)

type Parser interface {
	ParseAll() ([](*field.IfdField), error)
	ParseTag(id tags.TagId) (*field.IfdField, error)
}

type BasicExifParser struct {
	reader exif.Reader
}

func NewBasicExifParser(reader exif.Reader)

// func (r *BasicExifReader) ParseAll() ([](*tags.ExifTag), error) {

// 	ifd, err := ifd.NewIfd(r.marker.Contents[r.zerothIfdOffset:], r.byteOrder, r.zerothIfdOffset)

// 	if err != nil {
// 		return nil, err
// 	}

// 	fmt.Println(ifd)
// 	// field, err := ifd.LoadField(0, r.byteOrder)
// 	fields, err := ifd.LoadAllFields(r.byteOrder)

// 	if err != nil {
// 		return nil, err
// 	}

// 	for _, r := range fields {
// 		fmt.Println(r)
// 	}

// 	return nil, errors.New("ParseAll - Not Implemented")
// }

// func (p *BasicExifReader) ParseTag() (*tags.ExifTag, error) {
// 	return nil, errors.New("ParseTag - Not Implemented")
// }

// func (p *BasicExifReader) Read([]byte) (int, error) {

// }
