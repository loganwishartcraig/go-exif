package jpeg

import (
	"bytes"
	"errors"
	"fmt"
	"os"

	"github.com/loganwishartcraig/go-exif/ifd/field"
	"github.com/loganwishartcraig/go-exif/marker"
)

var soiMarker = [2]byte{0xff, 0xd8}

type JpegReaderErrorCode string

const (
	MalformedJpeg = "MalformedJpeg"
)

type JpegReaderError struct {
	code    JpegReaderErrorCode
	message string
}

func newJpegReaderError(code JpegReaderErrorCode, message string) *JpegReaderError {
	return &JpegReaderError{
		code,
		message,
	}
}

const (
	BaseAppMarkerOffset int64 = 0x2
)

func (e *JpegReaderError) Error() string {
	return fmt.Sprintf("Code: %s, message: %s", e.code, e.message)
}

type JpegReader struct {
	file *os.File
}

func NewJpegReader(f *os.File) (*JpegReader, error) {

	if err := validateSoi(f); err != nil {
		return nil, err
	}

	return &JpegReader{f}, nil

}

func validateSoi(f *os.File) error {

	soiBuf := make([]byte, 2)

	_, err := f.ReadAt(soiBuf, 0)

	if err != nil {
		return err
	}

	if soiBuf[0] != soiMarker[0] || soiBuf[1] != soiMarker[1] {
		return newJpegReaderError(
			MalformedJpeg,
			fmt.Sprintf("Malformed SOI Marker, received <0x%x>, expected <0x%x>", soiBuf, soiMarker),
		)
	}

	return nil

}

func (r *JpegReader) LoadApp1Marker() (*marker.Marker, error) {

	appMarker, err := marker.NewMarker(r.file, BaseAppMarkerOffset)

	if err != nil {
		return nil, err
	}

	if bytes.Equal(appMarker.Id, field.AppMarker1Id) {
		return appMarker, nil
	}

	// Assume it's APP0 if not APP1
	appMarker, err = marker.NewMarker(r.file, BaseAppMarkerOffset+(int64)(appMarker.Length)+(int64)(len(appMarker.Id)))

	if err != nil {
		return nil, err
	}

	if bytes.Equal(appMarker.Id, field.AppMarker1Id) {
		return appMarker, nil
	}

	return nil, errors.New("Unable to locate APP1 marker")

}
