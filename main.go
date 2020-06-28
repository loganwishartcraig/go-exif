package main

import (
	"fmt"
	"log"

	// exif "github.com/loganwishartcraig/go-exif/exif/reader"
	// "github.com/loganwishartcraig/go-exif/ifd"
	"github.com/loganwishartcraig/go-exif/marker"
	"github.com/loganwishartcraig/go-exif/reader/jpeg"
)

const filename = "./Canon_40D.jpg"

func check(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {

	content, err := jpeg.ReadFile(filename)
	check(err)

	markerLoader := marker.NewJpegLoader(content)

	marker, err := markerLoader.Load(marker.AppMarker1Id)
	check(err)

	fmt.Println(marker)

	// f, err := os.Open(filename)

	// reader, err := jpeg.NewJpegReader(f)
	// check(err)

	// appMarker, err := reader.LoadApp1Marker()
	// check(err)

	// fmt.Println(appMarker)

	// exifReader, err := exif.NewBasicExifReader(appMarker)
	// check(err)

	// fmt.Println(exifReader)

	// ifdField, err := ifd.NewIfd(exifReader, exifReader.ByteOrder, exifReader.ZerothIfdOffset)
	// check(err)

	// fmt.Println(ifdField)
	// _, err = exifParser.ParseAll()
	// check(err)

}

// Open jpeg as a byte array

// -- IF JPEG --
// [x] - validate SOI marker

// [ ] ingest APP0 if present

// [ ] ingest APP1
//     - Marker
//     - Length
//     - Identifier
//     - slice of entire APP1 body contents

// [ ] parse APP1 body into App1Body
// 	   - ByteOrder
//	   - 0th IFD Offset
//     - slice of body containing all content after header (starting at 0th ifd offset)
// -- END IF JPEG --

// [ ] parse App1Body into IFDTable(s)
//     - TableOffset
//     - ByteOrder
//     - fields
//         - TagId
//         - Type
//         - Count
//         - ValueOffset

// [ ] process IFDTable fields in conjunction with App1Body
// to produce an ExifDataTable
//      - map[TagId]ExifField
//          - TagId
//          - Name
//          - Value ({}interface)
// [ ] ExifField implemented based on TagId and Type.
//		- ExifByteField
//		- ExifStringField
//		- Etc...
