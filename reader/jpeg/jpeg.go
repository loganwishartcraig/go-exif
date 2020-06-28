package jpeg

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
)

var soiMarker = []byte{0xff, 0xd8}

func ReadFile(filename string) ([]byte, error) {

	fBytes, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	} else if !bytes.Equal(soiMarker, fBytes[0:2]) {
		return nil, errors.New(fmt.Sprintf("Invalid SOI marker <0x%x>", soiMarker))
	}

	return fBytes, nil

}
