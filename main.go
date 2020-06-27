package main

import (
	"fmt"
	"log"
	"os"

	exif "github.com/loganwishartcraig/go-exif/exif/reader"
	"github.com/loganwishartcraig/go-exif/ifd"
	"github.com/loganwishartcraig/go-exif/reader/jpeg"
)

const filename = "./Canon_40D.jpg"

func check(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {

	f, err := os.Open(filename)
	check(err)

	reader, err := jpeg.NewJpegReader(f)
	check(err)

	appMarker, err := reader.LoadApp1Marker()
	check(err)

	fmt.Println(appMarker)

	exifReader, err := exif.NewBasicExifReader(appMarker)
	check(err)

	fmt.Println(exifReader)

	ifdField, err := ifd.NewIfd(exifReader, exifReader.ByteOrder, exifReader.ZerothIfdOffset)
	check(err)

	fmt.Println(ifdField)
	// _, err = exifParser.ParseAll()
	// check(err)

}
