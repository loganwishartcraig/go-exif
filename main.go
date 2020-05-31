package main

import (
	"log"
	"os"

	"github.com/loganwishartcraig/go-exif/parser"
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

	exifParser, err := parser.NewBasicExifReader(appMarker)
	check(err)

	_, err = exifParser.ParseAll()
	check(err)

}
