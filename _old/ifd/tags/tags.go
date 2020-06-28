package tags

type MarkerId []byte

var (
	AppMarker0Id MarkerId = []byte{0xff, 0xe0}
	AppMarker1Id          = []byte{0xff, 0xe1}
)

type TagId uint16

const (

	// TIFF Tags
	ImageWidth                  TagId = 0x100
	ImageLength                       = 0x101
	BitsPerSample                     = 0x102
	Compression                       = 0x103
	PhotometricInterpretation         = 0x106
	ImageDescription                  = 0x10E
	Make                              = 0x10F
	Model                             = 0x110
	StripOffsets                      = 0x111
	Orientation                       = 0x112
	SamplesPerPixel                   = 0x115
	RowsPerStrip                      = 0x116
	StripByteCounts                   = 0x117
	XResolution                       = 0x11A
	YResolution                       = 0x11B
	PlanarConfiguration               = 0x11C
	ResolutionUnit                    = 0x128
	TransferFunction                  = 0x12D
	Software                          = 0x131
	DateTime                          = 0x132
	Artist                            = 0x13B
	WhitePoint                        = 0x13E
	PrimaryChromaticities             = 0x13F
	JpegInterchangeFormat             = 0x201
	JpegInterchangeFormatLength       = 0x202
	YCbCrCoefficients                 = 0x211
	YCbCrSubSampling                  = 0x212
	YCbCrPositioning                  = 0x213
	ReferenceBlackWhite               = 0x214
	Copyright                         = 0x8298 // Going to require special processing
	ExifIfdPointer                    = 0x8769
	GpsInfoIfdPointer                 = 0x8825

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
