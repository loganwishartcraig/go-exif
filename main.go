package main

import (
	"fmt"
	"log"
	"os"

	exif "github.com/loganwishartcraig/go-exif/exif/reader"
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

	exifReader, err := exif.NewBasicExifReader(appMarker)
	check(err)

	fmt.Println(exifReader)

	// _, err = exifParser.ParseAll()
	// check(err)

}
