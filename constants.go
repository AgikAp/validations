package validations

const (
	emailPattern        = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	alphaPattern        = `^[a-zA-Z]+$`
	alphaNumericPattern = `^[a-zA-Z0-9]+$`
	numberPattern       = `^[0-9]+$`
)

var (
	jpegImageSignature     = []byte{0xFF, 0xD8, 0xFF, 0xE0}
	jpegExifImageSignature = []byte{0xFF, 0xD8, 0xFF, 0xE1}
	pngImageSignature      = []byte{0x89, 0x50, 0x4E, 0x47}
	gifImageSignature      = []byte{0x47, 0x49, 0x46, 0x38}
	webPImageSignature     = []byte{0x52, 0x49, 0x46, 0x46}
	tiffImageSignature     = []byte{0x49, 0x49, 0x2A, 0x00}
	svgImageSignature      = []byte{0x3C, 0x3F, 0x78, 0x6D, 0x6C}
)

const imageExtension = "png,jpg,webp,tiff,svg"
