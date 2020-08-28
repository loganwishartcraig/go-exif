package tags

import (
	"fmt"
)

type Id uint16

type Type uint16

const (
	Byte      Type = 1
	Ascii          = 2
	Short          = 3
	Long           = 4
	Rational       = 5
	Undefined      = 7
	SLong          = 9
	SRational      = 10
)

const (

	// TIFF Tags
	ImageWidth                  Id = 0x100
	ImageLength                    = 0x101
	BitsPerSample                  = 0x102
	Compression                    = 0x103
	PhotometricInterpretation      = 0x106
	ImageDescription               = 0x10E
	Make                           = 0x10F
	Model                          = 0x110
	StripOffsets                   = 0x111
	Orientation                    = 0x112
	SamplesPerPixel                = 0x115
	RowsPerStrip                   = 0x116
	StripByteCounts                = 0x117
	XResolution                    = 0x11A
	YResolution                    = 0x11B
	PlanarConfiguration            = 0x11C
	ResolutionUnit                 = 0x128
	TransferFunction               = 0x12D
	Software                       = 0x131
	DateTime                       = 0x132
	Artist                         = 0x13B
	WhitePoint                     = 0x13E
	PrimaryChromaticities          = 0x13F
	JpegInterchangeFormat          = 0x201
	JpegInterchangeFormatLength    = 0x202
	YCbCrCoefficients              = 0x211
	YCbCrSubSampling               = 0x212
	YCbCrPositioning               = 0x213
	ReferenceBlackWhite            = 0x214
	Copyright                      = 0x8298 // Going to require special processing
	ExifIfdPointer                 = 0x8769
	GpsInfoIfdPointer              = 0x8825

	// Exif Private Tags
	ExposureTime               = 0x829A
	FNumber                    = 0x829D
	ExposureProgram            = 0x8822
	SpectralSensitivity        = 0x8824
	ISOSpeedRatings            = 0x8827
	OECF                       = 0x8828
	ExifVersion                = 0x9000
	DateTimeOriginal           = 0x9003
	DateTimeDigitized          = 0x9004
	ComponentsConfiguration    = 0x9101
	CompressedBitsPerPixel     = 0x9102
	ShutterSpeedValue          = 0x9201
	ApertureValue              = 0x9202
	BrightnessValue            = 0x9203
	ExposureBiasValue          = 0x9204
	MaxApertureValue           = 0x9205
	SubjectDistance            = 0x9206
	MeteringMode               = 0x9207
	LightSource                = 0x9208
	Flash                      = 0x9209
	FocalLength                = 0x920A
	SubjectArea                = 0x9214
	MakerNote                  = 0x927C // Requires special processing
	UserComment                = 0x9286 // Requires special processing
	SubsecTime                 = 0x9290
	SubsecTimeOriginal         = 0x9291
	SubsectimeDigital          = 0x9292
	FlashpixVersion            = 0xA000
	ColorSpace                 = 0xA001 // Special types
	PixelXDimension            = 0xA002
	PixelYDimension            = 0xA003
	RelatedSoundFile           = 0xA004
	InteroperabilityIfdPointer = 0xA005
	FlashEnergy                = 0xA20B
	SpatialFrequencyResponse   = 0xA20C
	FocalPlaneXResolution      = 0xA20E
	FocalPlaneYResolution      = 0xA20F
	FocalPlaneResolutionUnit   = 0xA210
	SubjectLocation            = 0xA214
	ExposureIndex              = 0xA215
	SensingMethod              = 0xA217
	FileSource                 = 0xA300
	SceneType                  = 0xA301
	CFAPattern                 = 0xA302
	CustomRendered             = 0xA401
	ExposureMode               = 0xA402
	WhiteBalance               = 0xA403
	DigitalZoomRatio           = 0xA404
	FocalLengthIn35mmFilm      = 0xA405
	SceneCaptureType           = 0xA406
	GainControl                = 0xA407
	Contrast                   = 0xA408
	Saturation                 = 0xA409
	Sharpness                  = 0xA40A
	DeviceSettingDescription   = 0xA40B
	SubjectDistanceRange       = 0xA40C
	ImageUniqueId              = 0xA420

	// GPS Info
	GpsVersionId        = 0x0
	GpsLatitudeRef      = 0x1
	GpsLatitude         = 0x2
	GpsLongitudeRef     = 0x3
	GpsLongitude        = 0x4
	GpsAltitudeRef      = 0x5
	GpsAltitude         = 0x6
	GpsTimeStamp        = 0x7
	GpsSatelites        = 0x8
	GpsStatus           = 0x9
	GpsMeasureMode      = 0xA
	GpsDop              = 0xB
	GpsSpeedRef         = 0xC
	GpsSpeed            = 0xD
	GpsTrackRef         = 0xE
	GpsTrack            = 0xF
	GpsImgDirectionRef  = 0x10
	GpsImgDirection     = 0x11
	GpsMapDatum         = 0x12
	GpsDestLatitudeRef  = 0x13
	GpsDestLatitude     = 0x14
	GpsDestLongitudeRef = 0x15
	GpsDestLongitude    = 0x16
	GpsDestBearingRef   = 0x17
	GpsDestBearing      = 0x18
	GpsDestDistanceRef  = 0x19
	GpsDestDistance     = 0x1A
	GpsProcessingMethod = 0x1B
	GpsAreaInformation  = 0x1C
	GpsDateStamp        = 0x1D
	GpsDifferential     = 0x1E

	// Interoperability Tag
	InteroperabilityIndex = 0x1
)

const (
	byteSize      = 1
	asciiSize     = 1
	shortSize     = 2
	longSize      = 4
	rationalSize  = 8
	undefinedSize = 1
	sLongSize     = 4
	sRationalSize = 8
)

func GetByteLength(t Type, count int) int {
	switch t {
	case Byte:
		return byteSize * count
	case Ascii:
		return asciiSize * count
	case Short:
		return shortSize * count
	case Long:
		return longSize * count
	case Rational:
		return rationalSize * count
	case Undefined:
		return undefinedSize * count
	case SLong:
		return sLongSize * count
	case SRational:
		return sRationalSize * count
	default:
		return 0
	}
}

func ValidateBufferType(t Type, buff interface{}) error {

	switch t {
	case Byte:
		if _, ok := buff.(byte); !ok {
			return fmt.Errorf("Expected 'byte' buffer type")
		}
	case Ascii:
		if _, ok := buff.(string); !ok {
			return fmt.Errorf("Expected 'string' buffer type")
		}
	case Short:
		if _, ok := buff.(int); !ok {
			return fmt.Errorf("Expected 'int' buffer type")
		}
	case Long:
		if _, ok := buff.(int); !ok {
			return fmt.Errorf("Expected 'int' buffer type")
		}
	case Rational:
		if _, ok := buff.(float64); !ok {
			return fmt.Errorf("Expected 'float64' buffer type")
		}
	case Undefined:
		if _, ok := buff.([]byte); !ok {
			return fmt.Errorf("Expected '[]byte' buffer type")
		}
	case SLong:
		if _, ok := buff.(int); !ok {
			return fmt.Errorf("Expected 'int' buffer type")
		}
	case SRational:
		if _, ok := buff.(float64); !ok {
			return fmt.Errorf("Expected 'float64' buffer type")
		}
	default:
		return fmt.Errorf("Unknown type %d", t)
	}

	return nil
}

func Label(id Id) string {

	// TIFF Tags
	switch id {
	case ImageWidth:
		return "ImageWidth"
	case ImageLength:
		return "ImageLength"
	case BitsPerSample:
		return "BitsPerSample"
	case Compression:
		return "Compression"
	case PhotometricInterpretation:
		return "PhotometricInterpretation"
	case ImageDescription:
		return "ImageDescription"
	case Make:
		return "Make"
	case Model:
		return "Model"
	case StripOffsets:
		return "StripOffsets"
	case Orientation:
		return "Orientation"
	case SamplesPerPixel:
		return "SamplesPerPixel"
	case RowsPerStrip:
		return "RowsPerStrip"
	case StripByteCounts:
		return "StripByteCounts"
	case XResolution:
		return "XResolution"
	case YResolution:
		return "YResolution"
	case PlanarConfiguration:
		return "PlanarConfiguration"
	case ResolutionUnit:
		return "ResolutionUnit"
	case TransferFunction:
		return "TransferFunction"
	case Software:
		return "Software"
	case DateTime:
		return "DateTime"
	case Artist:
		return "Artist"
	case WhitePoint:
		return "WhitePoint"
	case PrimaryChromaticities:
		return "PrimaryChromaticities"
	case JpegInterchangeFormat:
		return "JpegInterchangeFormat"
	case JpegInterchangeFormatLength:
		return "JpegInterchangeFormatLength"
	case YCbCrCoefficients:
		return "YCbCrCoefficients"
	case YCbCrSubSampling:
		return "YCbCrSubSampling"
	case YCbCrPositioning:
		return "YCbCrPositioning"
	case ReferenceBlackWhite:
		return "ReferenceBlackWhite"
	case Copyright:
		return "Copyright"
	case ExifIfdPointer:
		return "ExifIfdPointer"
	case GpsInfoIfdPointer:
		return "GpsInfoIfdPointer"
	default:
		return "Unknown"
	case ExposureTime:
		return "ExposureTime"
	case FNumber:
		return "FNumber"
	case ExposureProgram:
		return "ExposureProgram"
	case SpectralSensitivity:
		return "SpectralSensitivity"
	case ISOSpeedRatings:
		return "ISOSpeedRatings"
	case OECF:
		return "OECF"
	case ExifVersion:
		return "ExifVersion"
	case DateTimeOriginal:
		return "DateTimeOriginal"
	case DateTimeDigitized:
		return "DateTimeDigitized"
	case ComponentsConfiguration:
		return "ComponentsConfiguration"
	case CompressedBitsPerPixel:
		return "CompressedBitsPerPixel"
	case ShutterSpeedValue:
		return "ShutterSpeedValue"
	case ApertureValue:
		return "ApertureValue"
	case BrightnessValue:
		return "BrightnessValue"
	case ExposureBiasValue:
		return "ExposureBiasValue"
	case MaxApertureValue:
		return "MaxApertureValue"
	case SubjectDistance:
		return "SubjectDistance"
	case MeteringMode:
		return "MeteringMode"
	case LightSource:
		return "LightSource"
	case Flash:
		return "Flash"
	case FocalLength:
		return "FocalLength"
	case SubjectArea:
		return "SubjectArea"
	case MakerNote:
		return "MakerNote"
	case UserComment:
		return "UserComment"
	case SubsecTime:
		return "SubsecTime"
	case SubsecTimeOriginal:
		return "SubsecTimeOriginal"
	case SubsectimeDigital:
		return "SubsectimeDigital"
	case FlashpixVersion:
		return "FlashpixVersion"
	case ColorSpace:
		return "ColorSpace"
	case PixelXDimension:
		return "PixelXDimension"
	case PixelYDimension:
		return "PixelYDimension"
	case RelatedSoundFile:
		return "RelatedSoundFile"
	case InteroperabilityIfdPointer:
		return "InteroperabilityIfd"
	case FlashEnergy:
		return "FlashEnergy"
	case SpatialFrequencyResponse:
		return "SpatialFrequencyResponse"
	case FocalPlaneXResolution:
		return "FocalPlaneXResolution"
	case FocalPlaneYResolution:
		return "FocalPlaneYResolution"
	case FocalPlaneResolutionUnit:
		return "FocalPlaneResolutionUnit"
	case SubjectLocation:
		return "SubjectLocation"
	case ExposureIndex:
		return "ExposureIndex"
	case SensingMethod:
		return "SensingMethod"
	case FileSource:
		return "FileSource"
	case SceneType:
		return "SceneType"
	case CFAPattern:
		return "CFAPattern"
	case CustomRendered:
		return "CustomRendered"
	case ExposureMode:
		return "ExposureMode"
	case WhiteBalance:
		return "WhiteBalance"
	case DigitalZoomRatio:
		return "DigitalZoomRatio"
	case FocalLengthIn35mmFilm:
		return "FocalLengthIn35mmFilm"
	case SceneCaptureType:
		return "SceneCaptureType"
	case GainControl:
		return "GainControl"
	case Contrast:
		return "Contrast"
	case Saturation:
		return "Saturation"
	case Sharpness:
		return "Sharpness"
	case DeviceSettingDescription:
		return "DeviceSettingDescription"
	case SubjectDistanceRange:
		return "SubjectDistanceRange"
	case ImageUniqueId:
		return "ImageUniqueId"
	case GpsVersionId:
		return "GpsVersionId"
	case GpsLatitudeRef:
		return "GpsLatitudeRef"
	case GpsLatitude:
		return "GpsLatitude"
	case GpsLongitudeRef:
		return "GpsLongitudeRef"
	case GpsLongitude:
		return "GpsLongitude"
	case GpsAltitudeRef:
		return "GpsAltitudeRef"
	case GpsAltitude:
		return "GpsAltitude"
	case GpsTimeStamp:
		return "GpsTimeStamp"
	case GpsSatelites:
		return "GpsSatelites"
	case GpsStatus:
		return "GpsStatus"
	case GpsMeasureMode:
		return "GpsMeasureMode"
	case GpsDop:
		return "GpsDop"
	case GpsSpeedRef:
		return "GpsSpeedRef"
	case GpsSpeed:
		return "GpsSpeed"
	case GpsTrackRef:
		return "GpsTrackRef"
	case GpsTrack:
		return "GpsTrack"
	case GpsImgDirectionRef:
		return "GpsImgDirection"
	case GpsImgDirection:
		return "GpsImgDirection"
	case GpsMapDatum:
		return "GpsMapDatum"
	case GpsDestLatitudeRef:
		return "GpsDestLatitude"
	case GpsDestLatitude:
		return "GpsDestLatitude"
	case GpsDestLongitudeRef:
		return "GpsDestLongitude"
	case GpsDestLongitude:
		return "GpsDestLongitude"
	case GpsDestBearingRef:
		return "GpsDestBearingRef"
	case GpsDestBearing:
		return "GpsDestBearing"
	case GpsDestDistanceRef:
		return "GpsDestDistance"
	case GpsDestDistance:
		return "GpsDestDistance"
	case GpsProcessingMethod:
		return "GpsProcessing"
	case GpsAreaInformation:
		return "GpsArea"
	case GpsDateStamp:
		return "GpsDateStamp"
	case GpsDifferential:
		return "GpsDifferential"
	}
}
