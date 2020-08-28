package main

import (
	"fmt"
	"log"
	"os"

	"github.com/loganwishartcraig/go-exif/ifd"
	"github.com/loganwishartcraig/go-exif/ifd/field/reader"
	"github.com/loganwishartcraig/go-exif/jpeg"
)

func check(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {

	filename := os.Args[1]

	content, err := jpeg.ReadFile(filename)
	check(err)

	markerLoader := jpeg.NewSegmentLoader(content)

	segment, err := markerLoader.LoadExifSegmentBody()
	check(err)

	table, err := ifd.NewTable(markerLoader)
	check(err)

	r := reader.NewFieldReader()

	for _, b := range table.Blocks {
		fmt.Println(b)
		for _, f := range b.Fields {
			r.Print(f, segment)
		}
	}

}

// Open jpeg as a byte array

// -- IF JPEG --
// [x] - validate SOI marker

// [x] ingest APP0 if present

// [x] ingest APP1
//     - Marker
//     - Length
//     - Identifier
//     - slice of entire APP1 body contents

// [ ] parse APP1 body into App1Description
// 	   - ByteOrder
//	   - 0th IFD Offset
//     - slice of body containing all content after header (starting at 0th ifd offset)
// -- END IF JPEG --

// [x] parse App1Body into IFDTable(s)
//     - TableOffset
//     - ByteOrder
//     - fields
//         - TagId
//         - Type
//         - Count
//         - ValueOffset

// [x] process IFDTable fields in conjunction with App1Body
// to produce an ExifDataTable
//      - map[TagId]ExifField
//          - TagId
//          - Name
//          - Value ({}interface)
// [x] ExifField implemented based on TagId and Type.
//		- ExifByteField
//		- ExifStringField
//		- Etc...
