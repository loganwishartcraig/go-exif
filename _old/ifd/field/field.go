package field

type IfdMarkerId []byte

var (
	AppMarker0Id IfdMarkerId = []byte{0xff, 0xe0}
	AppMarker1Id             = []byte{0xff, 0xe1}
)

type IfdFieldValue interface{}

// type Reader interface {
// 	Read(b []byte) (int, error)
// 	ReadAt(b []byte, offset int64) (int, error)
// 	Seek(offset int64, whence int) (int64, error)
// }
