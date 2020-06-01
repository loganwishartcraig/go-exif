package field

import "github.com/loganwishartcraig/go-exif/ifd/tags"

type IfdMarkerId []byte

var (
	AppMarker0Id IfdMarkerId = []byte{0xff, 0xe0}
	AppMarker1Id             = []byte{0xff, 0xe1}
)

type IfdFieldType uint16
type IfdFieldValue interface{}

// const (
// 	Byte      IfdFieldType = 1
// 	Ascii                  = 2
// 	Short                  = 3
// 	Long                   = 4
// 	Rational               = 5
// 	Undefined              = 7
// 	SLong                  = 9
// 	SRational              = 10
// )

type IfdField struct {
	TagId       tags.TagId
	Type        IfdFieldType
	TypeLabel   string
	Count       uint32
	ValueOffset uint32
	BaseOffset  uint32
	Contents    []byte
}

type IfdSet []IfdField

// TODO - Might be better to follow same as binary.read?
type IfdFieldInterpreter interface {
	Value() IfdFieldValue
	Marshal(buf *IfdFieldValue)
}
