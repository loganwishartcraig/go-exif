package marker

import (
	"fmt"
)

type Id []byte

var (
	AppMarker0Id Id = []byte{0xff, 0xe0}
	AppMarker1Id    = []byte{0xff, 0xe1}
)

type Marker struct {
	Id       Id
	Contents []byte
}

func (m *Marker) String() string {
	return fmt.Sprintf("Marker \t ID: %x \t Length: %d \t Content Preview: % x ", m.Id, len(m.Contents), m.Contents[:50])
}
