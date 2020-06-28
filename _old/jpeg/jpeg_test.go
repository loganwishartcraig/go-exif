package jpeg

import (
	"io/ioutil"
	"log"
	"testing"
)

func check(e error) {
	if e != nil {
		log.Fatal(e.Error())
	}
}

func TestNewJpegReader(t *testing.T) {

	normalFile, err := ioutil.TempFile("", "normal.jpg")
	check(err)
	_, err = normalFile.Write([]byte{0xff, 0xd8})

	if _, err := NewJpegReader(normalFile); err != nil {
		t.Errorf("Expected reader to have no error when opening normal file. '%s'", err.Error())
	}

	corruptFile, err := ioutil.TempFile("", "corrupt.jpg")
	check(err)
	_, err = corruptFile.Write([]byte{0x33, 0x12})

	if _, err := NewJpegReader(normalFile); err == nil {
		t.Errorf("Expected reader to have error when reading corrupt file. '%s'", err.Error())
	}

}
