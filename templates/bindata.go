package templates

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

func templates_archives_html() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x4c, 0x8f,
		0xb1, 0x4e, 0xc4, 0x30, 0x0c, 0x86, 0xf7, 0x3e, 0x85, 0x15, 0x75, 0x84,
		0x54, 0xac, 0x28, 0x8d, 0x74, 0x82, 0x81, 0xfd, 0xd8, 0x51, 0xd4, 0xb8,
		0xad, 0xa5, 0x5c, 0x2a, 0xa5, 0x3e, 0x86, 0x33, 0x79, 0x77, 0x5c, 0xca,
		0x15, 0xa6, 0xc8, 0xbf, 0xbe, 0xcf, 0xce, 0x2f, 0x12, 0x71, 0xa4, 0x8c,
		0x60, 0x98, 0x38, 0xa1, 0xa9, 0xb5, 0x39, 0x95, 0x61, 0xa6, 0x4f, 0x5c,
		0xe1, 0x0b, 0x44, 0xec, 0x99, 0x18, 0xed, 0xcb, 0x92, 0x47, 0x9a, 0xec,
		0xfb, 0x86, 0x28, 0x21, 0x82, 0x39, 0xea, 0xdb, 0xfc, 0xd9, 0xc3, 0x92,
		0x19, 0x33, 0x6f, 0xbe, 0x9b, 0x9f, 0xfc, 0x7d, 0x87, 0xeb, 0x74, 0x68,
		0x5c, 0x4c, 0x30, 0xa4, 0xb0, 0xae, 0xbd, 0x89, 0xe9, 0x71, 0x5e, 0x0a,
		0xdd, 0x14, 0x0f, 0xc9, 0xf8, 0x06, 0xf4, 0x46, 0x09, 0x79, 0x42, 0x68,
		0x3f, 0x1e, 0xa0, 0x0d, 0x85, 0x69, 0x48, 0x08, 0xcf, 0x3d, 0xec, 0x97,
		0x4f, 0x7b, 0xb0, 0xea, 0x5e, 0x00, 0x17, 0xd9, 0x8b, 0xb4, 0xf6, 0xed,
		0x7a, 0x09, 0x99, 0x6e, 0xf8, 0x1a, 0x18, 0x99, 0x2e, 0x78, 0x78, 0x76,
		0x4b, 0x6a, 0x75, 0x9d, 0x82, 0x3f, 0x7c, 0xf4, 0x2e, 0xc0, 0x5c, 0x70,
		0xec, 0x8d, 0x8a, 0x77, 0xea, 0x9c, 0xae, 0x53, 0xad, 0xc6, 0xff, 0x8b,
		0x7e, 0xab, 0xb9, 0x2e, 0x78, 0xb5, 0xe3, 0xfe, 0xb1, 0xbd, 0xa5, 0xce,
		0xc9, 0x1f, 0x9d, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x3f, 0x64, 0x0b,
		0x14, 0x31, 0x01, 0x00, 0x00,
	},
		"templates/archives.html",
	)
}

func templates_article_html() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x8c, 0x54,
		0x51, 0x6f, 0xa3, 0x38, 0x10, 0x7e, 0xcf, 0xaf, 0x18, 0xa1, 0x4a, 0x10,
		0xf5, 0x0a, 0x77, 0x7d, 0x6c, 0x09, 0xa7, 0x5e, 0x7a, 0x52, 0xef, 0x74,
		0xb7, 0x5d, 0xa9, 0x79, 0x5b, 0xad, 0x2a, 0x03, 0x03, 0xb8, 0x35, 0x36,
		0xc5, 0x26, 0x59, 0x9a, 0xe6, 0xbf, 0xef, 0x18, 0x93, 0x36, 0x49, 0x53,
		0x75, 0x5f, 0x42, 0x98, 0x99, 0xef, 0x9b, 0xcf, 0xe3, 0x6f, 0x58, 0xaf,
		0x73, 0x2c, 0xb8, 0x44, 0xf0, 0x0c, 0x37, 0x02, 0xbd, 0xcd, 0x66, 0xb2,
		0x5e, 0x87, 0x57, 0xad, 0xe1, 0x99, 0xc0, 0x70, 0x61, 0x63, 0x9b, 0x0d,
		0xbc, 0x00, 0x05, 0xef, 0xb8, 0xc1, 0x70, 0xae, 0x64, 0xc1, 0xcb, 0x6d,
		0x82, 0x6a, 0x51, 0xe6, 0xf4, 0x9c, 0xbc, 0xf1, 0x64, 0x4a, 0x1a, 0x94,
		0xc6, 0x32, 0xc5, 0xd5, 0x1f, 0xc9, 0x7b, 0xb6, 0x38, 0xa2, 0xf0, 0x64,
		0x12, 0x37, 0x09, 0xa1, 0x78, 0x01, 0x7b, 0xc4, 0x57, 0x9d, 0xa9, 0x54,
		0x0b, 0x16, 0xac, 0x1b, 0x26, 0x21, 0x13, 0x4c, 0xeb, 0x99, 0x27, 0x58,
		0x8a, 0x02, 0x86, 0xdf, 0x33, 0x6a, 0xc4, 0x3a, 0x61, 0xbc, 0x24, 0xed,
		0x0f, 0x65, 0x39, 0xb4, 0x6d, 0x61, 0xc1, 0xc9, 0xae, 0x3c, 0xdb, 0x68,
		0x2b, 0xe4, 0x9a, 0x19, 0xfc, 0xa5, 0x0e, 0x44, 0x7f, 0xd3, 0xd5, 0x4c,
		0xf2, 0x67, 0xb4, 0x18, 0xc3, 0x6b, 0x3c, 0x64, 0xf9, 0xa4, 0xd5, 0x9c,
		0x8a, 0x4a, 0xd5, 0xf6, 0xb6, 0x1d, 0x3b, 0xd6, 0x4b, 0x77, 0x59, 0x86,
		0x5a, 0x7b, 0x50, 0xb5, 0x58, 0xcc, 0xbc, 0x28, 0x1b, 0x01, 0xd1, 0xce,
		0xe0, 0xde, 0x48, 0xc2, 0xca, 0xd4, 0xc2, 0x4b, 0x8e, 0xe6, 0xe2, 0x88,
		0x7d, 0x28, 0x63, 0xc1, 0x4a, 0x3d, 0x5c, 0x57, 0xcb, 0x64, 0x89, 0x70,
		0x72, 0xff, 0x1b, 0x9c, 0x18, 0x56, 0xc2, 0xc5, 0xec, 0x5d, 0xcd, 0x71,
		0x99, 0x5c, 0x16, 0xea, 0x55, 0x23, 0x21, 0x49, 0x9e, 0x25, 0xd8, 0x51,
		0xe4, 0x5e, 0xf7, 0x44, 0x6c, 0x9f, 0x71, 0x44, 0x77, 0xbd, 0xeb, 0xab,
		0x9b, 0xc5, 0xff, 0xff, 0xed, 0xcc, 0xc8, 0x39, 0x66, 0x00, 0x1c, 0xfa,
		0xe1, 0x9a, 0xeb, 0xa7, 0x4e, 0xdb, 0x88, 0x64, 0xf5, 0x70, 0x69, 0xd5,
		0x79, 0x32, 0x57, 0x75, 0x4d, 0x00, 0x4d, 0x4e, 0x3a, 0x4f, 0x26, 0x71,
		0xce, 0x97, 0xc0, 0xf3, 0x99, 0x97, 0x0f, 0xb5, 0xf7, 0x86, 0x54, 0xb2,
		0xdc, 0x4b, 0xe2, 0x88, 0x12, 0x94, 0xd6, 0x59, 0xcb, 0x1b, 0x03, 0xa6,
		0x6f, 0x70, 0xe6, 0x19, 0xfc, 0x61, 0xa2, 0x07, 0xb6, 0x64, 0x2e, 0xea,
		0x25, 0x13, 0x80, 0x25, 0x6b, 0x61, 0xc4, 0x6a, 0xb2, 0x8f, 0xb1, 0x9d,
		0x60, 0x06, 0xfe, 0x81, 0xb9, 0x0e, 0xa5, 0xf8, 0x97, 0x84, 0x8d, 0x22,
		0x58, 0xdc, 0x5e, 0xdf, 0x42, 0xc0, 0x4a, 0x25, 0x9f, 0x99, 0xc0, 0xe7,
		0x56, 0x4d, 0x2f, 0x80, 0x14, 0x36, 0xcc, 0xf0, 0x94, 0x0b, 0x6e, 0x7a,
		0x58, 0x71, 0x53, 0x41, 0xdd, 0x83, 0x12, 0x39, 0xa4, 0x42, 0x95, 0xc0,
		0xdc, 0xc9, 0xb5, 0x23, 0xf8, 0x07, 0x72, 0x25, 0x7d, 0x03, 0x8f, 0x52,
		0xad, 0x80, 0x06, 0xc0, 0x0d, 0x21, 0x84, 0x80, 0x95, 0x6a, 0x1f, 0xa1,
		0xa0, 0x6d, 0x50, 0xa6, 0xc2, 0x16, 0x3a, 0x8d, 0xad, 0xde, 0x97, 0xcb,
		0x73, 0x1a, 0x03, 0x2f, 0x38, 0x65, 0x9d, 0xde, 0xed, 0x48, 0xef, 0x44,
		0x57, 0x0e, 0x0a, 0xa9, 0x3e, 0x28, 0x3a, 0x99, 0x19, 0xae, 0x64, 0x30,
		0x85, 0x35, 0xbd, 0x8f, 0x0c, 0xfa, 0x89, 0x30, 0xb9, 0xca, 0x3a, 0x3b,
		0xca, 0x30, 0xa3, 0x99, 0x19, 0xfc, 0x5b, 0xa0, 0x7d, 0x0b, 0x7c, 0x37,
		0x1d, 0x7f, 0x7a, 0x69, 0xeb, 0x42, 0x3b, 0x3a, 0xdb, 0xe0, 0x60, 0x78,
		0xbe, 0xcb, 0x32, 0xdd, 0xcb, 0x8c, 0xd2, 0xa6, 0xed, 0xf0, 0x72, 0xe0,
		0xb7, 0x51, 0xdd, 0xda, 0x98, 0x1f, 0x45, 0x3e, 0x9c, 0xbe, 0x9f, 0xee,
		0x29, 0xf8, 0xa1, 0x0b, 0x86, 0x99, 0xaa, 0x23, 0xac, 0x53, 0xcc, 0xc3,
		0x07, 0xed, 0x3b, 0x7c, 0xf0, 0x2a, 0xab, 0x44, 0x33, 0x6a, 0xd2, 0x7f,
		0xf5, 0x64, 0xd0, 0x2f, 0x04, 0x0e, 0xfc, 0x8a, 0xee, 0xd7, 0x9f, 0x7e,
		0xfb, 0xfd, 0x3b, 0xbc, 0xbc, 0xc0, 0x27, 0xb5, 0xa9, 0xca, 0xfb, 0xa1,
		0x76, 0x1a, 0xb2, 0xa6, 0x21, 0x3f, 0xce, 0x2b, 0x2e, 0xf2, 0x80, 0x24,
		0x4e, 0x6d, 0xb3, 0xcd, 0x34, 0xa0, 0x27, 0x2d, 0xf1, 0x70, 0x22, 0x72,
		0x8b, 0x54, 0xe3, 0xdf, 0xaf, 0x02, 0x99, 0x46, 0xa0, 0xcb, 0x4e, 0x05,
		0xc2, 0xbf, 0x74, 0xea, 0xbb, 0xd1, 0x48, 0x0a, 0x96, 0x1c, 0x57, 0x40,
		0x97, 0x02, 0xb4, 0x2c, 0x6e, 0x27, 0x2a, 0x63, 0x9a, 0x8b, 0x28, 0xda,
		0x39, 0xd3, 0x9f, 0x14, 0xbf, 0xdf, 0xb2, 0x79, 0x49, 0x36, 0x7a, 0x16,
		0x1a, 0xb5, 0xc2, 0x16, 0xc9, 0x08, 0x3d, 0x38, 0x47, 0x85, 0x76, 0x67,
		0xe2, 0xe8, 0xb5, 0xf1, 0xe4, 0x63, 0x52, 0x6f, 0xbb, 0x9a, 0x24, 0xff,
		0x2c, 0x6d, 0x05, 0x97, 0x8f, 0xc7, 0x99, 0xf7, 0x3f, 0x6d, 0xaa, 0x54,
		0x67, 0x8e, 0xc4, 0x4b, 0x5c, 0xcf, 0xf1, 0xab, 0x75, 0x74, 0x5b, 0x7f,
		0x06, 0x00, 0x00, 0xff, 0xff, 0xe3, 0x49, 0x85, 0x94, 0x13, 0x06, 0x00,
		0x00,
	},
		"templates/article.html",
	)
}

func templates_atom_xml() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x6c, 0x92,
		0x4d, 0x4f, 0xc3, 0x20, 0x18, 0xc7, 0xef, 0xfd, 0x14, 0x84, 0xec, 0xa8,
		0x30, 0x35, 0x26, 0x66, 0xa1, 0x5d, 0x16, 0x3d, 0x78, 0x98, 0x17, 0xb7,
		0x9d, 0x0d, 0x59, 0x9f, 0x76, 0xc4, 0x16, 0x16, 0x4a, 0xad, 0xa6, 0xe1,
		0xbb, 0x3b, 0xe8, 0xe8, 0x9b, 0xde, 0xc8, 0xf3, 0x7f, 0xe1, 0xf7, 0x10,
		0xd8, 0xfa, 0xbb, 0x2c, 0xd0, 0x17, 0xe8, 0x4a, 0x28, 0x19, 0xe3, 0x3b,
		0xb2, 0xc4, 0x08, 0xe4, 0x51, 0xa5, 0x42, 0xe6, 0x31, 0xae, 0x4d, 0x76,
		0xfb, 0x84, 0xd7, 0x49, 0x14, 0xb1, 0x0c, 0x20, 0x45, 0x17, 0xaf, 0xac,
		0x62, 0x7c, 0x32, 0xe6, 0xbc, 0xa2, 0xb4, 0x69, 0x1a, 0xd2, 0x3c, 0x10,
		0xa5, 0x73, 0x7a, 0xbf, 0x5c, 0x3e, 0xd2, 0x8d, 0x51, 0x25, 0x4e, 0x22,
		0x84, 0x98, 0x11, 0xa6, 0x80, 0xa4, 0x6d, 0xc9, 0x4e, 0x18, 0x20, 0xcf,
		0x4a, 0x66, 0x22, 0x27, 0x7b, 0x37, 0xb4, 0x96, 0xd1, 0x4e, 0x75, 0xbe,
		0x42, 0xc8, 0x4f, 0x74, 0xd2, 0x90, 0xc5, 0x78, 0x66, 0x3e, 0xbc, 0x6f,
		0xad, 0xc5, 0x88, 0x7a, 0x5b, 0x7d, 0x4e, 0xb9, 0x81, 0xd4, 0x15, 0x1e,
		0xba, 0xa3, 0xab, 0x09, 0xd3, 0xe8, 0x62, 0x69, 0x5b, 0x91, 0x21, 0xb2,
		0xd1, 0x46, 0x1c, 0x0b, 0xa8, 0xac, 0xf5, 0x23, 0xcd, 0x65, 0x0e, 0x68,
		0xf1, 0x71, 0x83, 0x16, 0xbc, 0x53, 0xd0, 0x2a, 0x9e, 0xb9, 0x18, 0x48,
		0xa3, 0x7f, 0xdc, 0x2d, 0x23, 0xec, 0x60, 0xff, 0x07, 0x79, 0x0e, 0xbd,
		0xf8, 0x4b, 0x3d, 0xca, 0xef, 0x8a, 0x3a, 0xef, 0xd7, 0xb8, 0x52, 0xf6,
		0xe2, 0xcb, 0x85, 0xde, 0x33, 0x4c, 0x36, 0x9c, 0xc9, 0xa3, 0x35, 0xbb,
		0x06, 0x90, 0x69, 0x08, 0x55, 0x75, 0x59, 0xf2, 0x00, 0xef, 0xb4, 0x3e,
		0xfb, 0xba, 0x7f, 0xdb, 0x0e, 0x17, 0xed, 0x3a, 0x5f, 0x88, 0xd1, 0x49,
		0x8e, 0xf1, 0xda, 0x9c, 0x94, 0x0e, 0x25, 0x4c, 0xf2, 0xd2, 0x3f, 0xc1,
		0x64, 0xaf, 0x8d, 0xf7, 0x38, 0x1a, 0x2f, 0x5f, 0x7b, 0x86, 0x24, 0xa3,
		0xfd, 0x33, 0x0e, 0x84, 0xe1, 0xc4, 0xa8, 0xfb, 0x3b, 0x49, 0xf4, 0x1b,
		0x00, 0x00, 0xff, 0xff, 0x41, 0x57, 0x24, 0xb9, 0x6b, 0x02, 0x00, 0x00,
	},
		"templates/atom.xml",
	)
}

func templates_base_analytics_html() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x84, 0x92,
		0x51, 0x8f, 0xd3, 0x30, 0x0c, 0xc7, 0xdf, 0xf9, 0x14, 0x56, 0x5f, 0xd2,
		0x8a, 0x53, 0xca, 0xf3, 0xf5, 0x76, 0xe8, 0x0e, 0x10, 0xe2, 0xe5, 0x84,
		0x04, 0x6f, 0xd3, 0x34, 0x99, 0xd4, 0xcd, 0xb2, 0x75, 0x49, 0x49, 0xdc,
		0x41, 0xd5, 0xed, 0xbb, 0x93, 0xa6, 0x9d, 0x98, 0x10, 0xd2, 0x3d, 0xc5,
		0xfe, 0xfb, 0x57, 0xdb, 0xff, 0xca, 0xe3, 0x58, 0x53, 0x63, 0x2c, 0x41,
		0x86, 0x16, 0xdb, 0x81, 0x8d, 0x0a, 0xd9, 0xe5, 0xf2, 0x06, 0x60, 0x1c,
		0x4d, 0x03, 0xf2, 0x9b, 0x61, 0x92, 0x1f, 0x9c, 0x6d, 0x8c, 0x96, 0x9f,
		0x9d, 0xd3, 0x2d, 0x3d, 0x5d, 0xb1, 0x2f, 0x1f, 0x13, 0xf7, 0x10, 0x94,
		0x37, 0x1d, 0x03, 0x0f, 0x1d, 0xad, 0x32, 0xa6, 0xdf, 0x5c, 0xee, 0xf1,
		0x84, 0xb3, 0x9a, 0x3d, 0x46, 0x02, 0xe0, 0x84, 0x1e, 0xb6, 0x1a, 0x7f,
		0xc2, 0x6a, 0x7e, 0xce, 0x67, 0x58, 0x6f, 0xaa, 0x54, 0x9a, 0x72, 0xd9,
		0xf5, 0x61, 0x97, 0xaf, 0xc5, 0x36, 0x10, 0x3f, 0x29, 0xe5, 0x7a, 0xcb,
		0xe2, 0x0e, 0xc4, 0x38, 0xbe, 0x36, 0x5e, 0x6c, 0x8a, 0xff, 0x74, 0x61,
		0x8f, 0xea, 0xf0, 0x15, 0x35, 0x9d, 0x0c, 0xfd, 0x4a, 0x48, 0x62, 0xf2,
		0xa6, 0xb7, 0x8a, 0x8d, 0xb3, 0x79, 0x01, 0x63, 0x52, 0xe6, 0xc5, 0x34,
		0xc6, 0xb5, 0x6a, 0xa7, 0xfa, 0x23, 0x59, 0x96, 0xca, 0x13, 0x32, 0x7d,
		0x6a, 0x69, 0xca, 0x72, 0x31, 0xdb, 0x10, 0x45, 0x15, 0x31, 0x39, 0x59,
		0x8c, 0xac, 0xf8, 0xc7, 0xa4, 0x48, 0x45, 0x0c, 0x83, 0x55, 0xb1, 0xca,
		0xbe, 0xa7, 0x6a, 0x69, 0x1f, 0xe5, 0xe0, 0x27, 0x31, 0x17, 0x3b, 0xe6,
		0x2e, 0xdc, 0x0b, 0x58, 0xdd, 0xcc, 0x6a, 0x9d, 0xc2, 0x69, 0x21, 0xd9,
		0x79, 0xc7, 0x4e, 0xb9, 0x16, 0xde, 0xc3, 0x02, 0x96, 0xa5, 0x80, 0xfb,
		0x39, 0x99, 0xe2, 0x02, 0xde, 0x82, 0x08, 0x8c, 0x1c, 0xa4, 0x96, 0xb5,
		0xeb, 0x7f, 0xb4, 0xa4, 0x5a, 0xa3, 0x0e, 0xd2, 0x12, 0x97, 0xb5, 0x92,
		0xfb, 0x20, 0xaa, 0x1b, 0x4b, 0xe1, 0xd6, 0x91, 0x26, 0x5e, 0xec, 0x84,
		0xe7, 0xe1, 0x3b, 0xea, 0x17, 0x3c, 0xd2, 0x5f, 0x63, 0xeb, 0x77, 0x9b,
		0x0a, 0x82, 0xec, 0xd0, 0x47, 0xe0, 0xc5, 0xd5, 0x24, 0x8d, 0x0d, 0xe4,
		0xf9, 0x99, 0x1a, 0xe7, 0x29, 0xd7, 0x78, 0x07, 0x61, 0xf9, 0xc9, 0x97,
		0x22, 0x4f, 0xd1, 0x43, 0x39, 0x7f, 0xfc, 0x98, 0xae, 0x84, 0x6c, 0x1d,
		0xef, 0xe0, 0xfa, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0xea, 0xed, 0x60,
		0x78, 0x4f, 0x02, 0x00, 0x00,
	},
		"templates/base/analytics.html",
	)
}

func templates_base_base_html() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xd4, 0x57,
		0xcd, 0x8e, 0xe4, 0x34, 0x10, 0xbe, 0xef, 0x53, 0x18, 0x33, 0x07, 0x56,
		0x6c, 0x92, 0x1d, 0x4e, 0x2b, 0x94, 0x44, 0x1a, 0x9a, 0x5d, 0x89, 0x03,
		0xda, 0x95, 0xa6, 0x91, 0xe0, 0x84, 0xaa, 0x13, 0x27, 0xb1, 0x70, 0xec,
		0xe0, 0x38, 0x3d, 0xb4, 0xa2, 0x79, 0x77, 0xca, 0x89, 0x9d, 0xbf, 0xce,
		0x0c, 0x3b, 0x42, 0x8b, 0xc4, 0x29, 0xb6, 0xeb, 0xcf, 0x55, 0xdf, 0xe7,
		0xea, 0xea, 0xbe, 0xcf, 0x59, 0xc1, 0x25, 0x23, 0xf4, 0x04, 0x2d, 0xa3,
		0x8f, 0x8f, 0xaf, 0xe2, 0xaf, 0x7e, 0xfc, 0x78, 0x38, 0xfe, 0xf6, 0xe9,
		0x3d, 0xa9, 0x4c, 0x2d, 0xd2, 0x57, 0xf1, 0xf8, 0x21, 0x24, 0xae, 0x18,
		0xe4, 0x76, 0x81, 0xcb, 0x9a, 0x19, 0x20, 0x59, 0x05, 0xba, 0x65, 0x26,
		0xa1, 0x9d, 0x29, 0x82, 0x77, 0x74, 0x29, 0xaa, 0x8c, 0x69, 0x02, 0xf6,
		0x67, 0xc7, 0xcf, 0x09, 0xfd, 0x35, 0xf8, 0xe5, 0x2e, 0x38, 0xa8, 0xba,
		0x01, 0xc3, 0x4f, 0x82, 0x51, 0x92, 0x29, 0x69, 0x98, 0x44, 0xbb, 0x9f,
		0xde, 0x27, 0x2c, 0x2f, 0xd9, 0xca, 0x52, 0x42, 0xcd, 0x12, 0x7a, 0xe6,
		0xec, 0xa1, 0x51, 0xda, 0x2c, 0x94, 0x1f, 0x78, 0x6e, 0xaa, 0x24, 0x67,
		0x67, 0x9e, 0xb1, 0x60, 0xd8, 0xbc, 0x21, 0x5c, 0x72, 0xc3, 0x41, 0x04,
		0x6d, 0x06, 0x82, 0x25, 0xb7, 0xe1, 0x5b, 0x74, 0x35, 0xfa, 0x32, 0xdc,
		0x08, 0x96, 0xf6, 0xbd, 0x61, 0x75, 0x23, 0xc0, 0x60, 0x7e, 0xc3, 0x09,
		0x25, 0xe1, 0xe3, 0x63, 0x1c, 0x8d, 0x52, 0xa7, 0x2a, 0xb8, 0xfc, 0x83,
		0x54, 0x9a, 0x15, 0x09, 0x8d, 0x0a, 0xc6, 0xf2, 0x36, 0x02, 0x21, 0x42,
		0x30, 0xaa, 0x0e, 0xff, 0xaa, 0x05, 0x25, 0xe6, 0xd2, 0xe0, 0x8d, 0xa0,
		0x69, 0x04, 0xcf, 0x30, 0x05, 0x25, 0x23, 0x2b, 0xfb, 0x76, 0x90, 0x69,
		0x26, 0x50, 0x24, 0x0c, 0xd3, 0x12, 0x83, 0xa0, 0xae, 0x75, 0x9c, 0xd0,
		0x7b, 0x6e, 0x18, 0x5e, 0x91, 0x91, 0xbb, 0xe3, 0xc7, 0x9f, 0xc9, 0x07,
		0x74, 0xea, 0x72, 0xec, 0x7b, 0x5e, 0x90, 0xd0, 0x8a, 0xc3, 0x83, 0x92,
		0x05, 0x2f, 0xc3, 0x0f, 0x80, 0x09, 0x29, 0x89, 0x85, 0x9f, 0x2f, 0x33,
		0x78, 0x6d, 0x2b, 0xcc, 0x3f, 0xeb, 0x0c, 0xb1, 0x62, 0xea, 0x2e, 0xd8,
		0xf7, 0xfb, 0xc6, 0x94, 0x44, 0x3e, 0x00, 0x93, 0x39, 0x3a, 0x73, 0x9b,
		0x39, 0x7d, 0x0b, 0x1e, 0xd3, 0x43, 0xfe, 0x16, 0xcb, 0x68, 0x04, 0xd3,
		0x2e, 0x4f, 0x2a, 0xbf, 0x38, 0x04, 0x72, 0x7e, 0x26, 0x99, 0x80, 0xb6,
		0x4d, 0xa8, 0x84, 0xf3, 0x09, 0x34, 0x19, 0x3f, 0x01, 0x92, 0x04, 0x3a,
		0x61, 0xfc, 0xb6, 0x35, 0x58, 0x88, 0x2c, 0x30, 0xaa, 0xc1, 0x12, 0x28,
		0x9b, 0x31, 0x0a, 0x78, 0x39, 0x54, 0xc7, 0x65, 0xba, 0xf6, 0x66, 0x51,
		0x04, 0xa4, 0x99, 0x9e, 0xa4, 0x7b, 0xd1, 0x02, 0x77, 0xc9, 0x59, 0xc7,
		0x5e, 0xaf, 0x33, 0x46, 0x49, 0x87, 0xc2, 0xb8, 0xa1, 0x1b, 0x33, 0xa3,
		0xca, 0xd2, 0x62, 0x9b, 0x83, 0x01, 0xb7, 0xb1, 0x21, 0x85, 0x80, 0xa6,
		0x9d, 0x8e, 0x41, 0x97, 0x96, 0xab, 0xa1, 0xb3, 0x99, 0xc4, 0xcb, 0x60,
		0x18, 0xae, 0x6d, 0x40, 0x7a, 0xf7, 0xad, 0x0e, 0x94, 0x14, 0x17, 0x9a,
		0x1e, 0x07, 0x9f, 0x64, 0x4e, 0x32, 0x8e, 0xac, 0xde, 0x33, 0xa6, 0x16,
		0x96, 0x00, 0xe3, 0xd0, 0xf4, 0xbf, 0x52, 0x8d, 0xa3, 0xb1, 0x38, 0xab,
		0x33, 0xd8, 0x54, 0xea, 0xa4, 0x41, 0xe6, 0x9e, 0x4c, 0x11, 0x4d, 0x37,
		0x7c, 0x3a, 0x5a, 0xfa, 0xda, 0x07, 0x02, 0x0b, 0x98, 0x22, 0xc4, 0x69,
		0xde, 0x0e, 0x0c, 0x56, 0xda, 0x91, 0xf8, 0x13, 0x94, 0xac, 0x25, 0xdf,
		0xa0, 0x53, 0x77, 0x70, 0x84, 0xb2, 0x5d, 0x13, 0xfc, 0xbe, 0x52, 0x0f,
		0xf6, 0xf4, 0xf5, 0x52, 0xed, 0x80, 0x9c, 0x2c, 0x95, 0xe6, 0x6c, 0x47,
		0x79, 0x96, 0xad, 0x4c, 0xee, 0x34, 0x52, 0x4e, 0xec, 0x19, 0xdc, 0xe9,
		0xac, 0xe2, 0x67, 0xf6, 0xda, 0xbd, 0xa1, 0x27, 0xb8, 0xe5, 0x01, 0x27,
		0xbb, 0xc8, 0xff, 0x43, 0x5a, 0x2f, 0x0a, 0x8e, 0xe1, 0x3b, 0xb1, 0x88,
		0xee, 0x5f, 0x0d, 0x7e, 0x36, 0x6c, 0xbb, 0x6e, 0x07, 0x0b, 0x97, 0x2b,
		0x8f, 0x43, 0x6b, 0x48, 0x11, 0x4f, 0x07, 0x1d, 0x8c, 0x3a, 0x6d, 0x68,
		0xbb, 0x33, 0x4d, 0x9d, 0x49, 0x6b, 0x91, 0x8b, 0x23, 0xd4, 0xdc, 0x84,
		0x19, 0x9b, 0xc2, 0xfa, 0x0c, 0xa9, 0x50, 0x32, 0x72, 0xf3, 0xfb, 0x1b,
		0x72, 0xd3, 0x60, 0xbe, 0xe4, 0xfb, 0x64, 0x99, 0xfd, 0x73, 0xd1, 0xfb,
		0x7e, 0xb0, 0x08, 0xef, 0x45, 0x57, 0x62, 0xef, 0x49, 0xfd, 0x7e, 0x41,
		0x9f, 0xcf, 0xb9, 0x44, 0x1c, 0x75, 0x62, 0x8d, 0x01, 0x41, 0x15, 0xb2,
		0xd2, 0xf1, 0xb8, 0xfc, 0x3b, 0xee, 0x3c, 0xc3, 0xca, 0xcf, 0xc3, 0xcd,
		0x2f, 0x35, 0x2f, 0x2b, 0xb3, 0x07, 0xe2, 0x4b, 0xaf, 0x77, 0x5d, 0x5e,
		0x1f, 0x38, 0xd7, 0xaa, 0xc9, 0xd5, 0x83, 0xdc, 0x36, 0x26, 0x5f, 0xfc,
		0xaf, 0x37, 0x7d, 0x6e, 0xd6, 0x5f, 0x44, 0x8e, 0x4f, 0x53, 0xe7, 0x05,
		0xcd, 0x8c, 0xed, 0x17, 0xa7, 0x74, 0xf5, 0xac, 0x37, 0xe9, 0x7a, 0x2f,
		0x41, 0xcd, 0x64, 0xb7, 0x09, 0xbd, 0x21, 0x4b, 0x36, 0x86, 0xb9, 0xcc,
		0x84, 0x79, 0x32, 0x2d, 0x6b, 0x79, 0x03, 0xfe, 0xed, 0xa0, 0xfe, 0xcd,
		0xf4, 0x92, 0x7e, 0xb8, 0x1c, 0xbc, 0x9f, 0xc9, 0xe3, 0x95, 0x75, 0xbc,
		0x25, 0xd1, 0xa2, 0x0e, 0x91, 0xb7, 0x8a, 0x30, 0xc6, 0xec, 0xc1, 0x3d,
		0x88, 0xd5, 0xd9, 0xba, 0x89, 0x9e, 0x00, 0x67, 0x0e, 0xd2, 0x74, 0x42,
		0x78, 0x38, 0xfb, 0x5e, 0x30, 0x49, 0xa6, 0x8b, 0x5a, 0x06, 0x0f, 0xdd,
		0xf5, 0xaa, 0x60, 0x96, 0xb3, 0xd7, 0x57, 0xda, 0x7b, 0x5e, 0x5b, 0x6e,
		0xef, 0x59, 0xee, 0x3f, 0xcb, 0x15, 0x99, 0x9e, 0x26, 0xee, 0x97, 0x20,
		0xd0, 0x10, 0xed, 0x4b, 0x52, 0xc7, 0x40, 0x39, 0xb3, 0x66, 0x27, 0x8b,
		0x67, 0xf9, 0x82, 0xfa, 0x83, 0x87, 0x97, 0xb1, 0x04, 0x0d, 0x2c, 0x41,
		0x06, 0xbb, 0x99, 0x1b, 0xc3, 0xf6, 0x7f, 0x43, 0x8b, 0xeb, 0x46, 0xb9,
		0xd6, 0xb8, 0xfa, 0x89, 0x5e, 0x8a, 0x17, 0x42, 0xb7, 0xbc, 0x9a, 0xf5,
		0xa6, 0xe9, 0xcc, 0xfe, 0x36, 0x06, 0x75, 0x1e, 0xdc, 0xbe, 0xf5, 0x2b,
		0x55, 0x14, 0x38, 0xe2, 0x07, 0xb7, 0xc3, 0x5e, 0x94, 0xc1, 0x3b, 0xbf,
		0x70, 0x82, 0xef, 0x26, 0xc0, 0x97, 0xc3, 0xa6, 0x1b, 0xda, 0xfd, 0xb4,
		0xb9, 0xba, 0xc4, 0x52, 0xaf, 0x50, 0xca, 0xcc, 0x43, 0xe9, 0x5a, 0x06,
		0x12, 0xc4, 0x05, 0x2b, 0xdf, 0xce, 0x33, 0xeb, 0x38, 0xa9, 0xe2, 0xec,
		0x3a, 0xfc, 0x23, 0xf1, 0x69, 0xfe, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xb7,
		0x56, 0x17, 0x76, 0xc3, 0x0c, 0x00, 0x00,
	},
		"templates/base/base.html",
	)
}

func templates_base_footer_html() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x74, 0x92,
		0xdd, 0x8e, 0xa3, 0x30, 0x0c, 0x85, 0xef, 0xf7, 0x29, 0xac, 0x5c, 0x56,
		0x82, 0x94, 0xbd, 0x5a, 0xad, 0x28, 0xda, 0x87, 0xd8, 0x17, 0x08, 0x60,
		0x20, 0x28, 0x24, 0xd9, 0xd8, 0xdd, 0xfe, 0xa0, 0xbe, 0xfb, 0x18, 0x5a,
		0xa6, 0x33, 0x53, 0xcd, 0x0d, 0x3a, 0x8e, 0xc3, 0xe7, 0x73, 0xc0, 0xf3,
		0xdc, 0x62, 0x67, 0x3d, 0x82, 0xea, 0x42, 0x60, 0x4c, 0xea, 0x76, 0xfb,
		0x01, 0x30, 0xcf, 0x7a, 0x07, 0x7f, 0x07, 0x4b, 0x40, 0x43, 0x38, 0xba,
		0x16, 0x6a, 0x04, 0x8a, 0xce, 0x32, 0x04, 0x0f, 0x1e, 0x4f, 0x90, 0xd0,
		0xa1, 0x21, 0x24, 0xd8, 0xe9, 0xf5, 0x85, 0xb2, 0xb5, 0xff, 0xa1, 0x71,
		0x86, 0xe8, 0xa0, 0x9a, 0xe0, 0xd9, 0x08, 0x32, 0x41, 0x13, 0x5c, 0x36,
		0xb5, 0x59, 0xb1, 0xdf, 0x54, 0xe8, 0x3a, 0x42, 0xce, 0x8a, 0xb5, 0x76,
		0x7d, 0xf6, 0x6b, 0x13, 0x8f, 0xc6, 0x4f, 0x55, 0x09, 0xec, 0x33, 0x8e,
		0xf1, 0xcc, 0x59, 0x83, 0x7e, 0x31, 0x77, 0xef, 0x4a, 0x3f, 0x6e, 0x4a,
		0xb4, 0x81, 0x21, 0x61, 0x77, 0x50, 0x03, 0x73, 0xfc, 0xad, 0x75, 0x6f,
		0x79, 0x38, 0xd6, 0x79, 0x13, 0x26, 0x6d, 0xfa, 0xe0, 0xaf, 0xc6, 0xe1,
		0x35, 0x05, 0x1d, 0x83, 0x0b, 0xaa, 0x5a, 0x9e, 0xa5, 0x36, 0x15, 0x48,
		0xb4, 0xc9, 0xb4, 0x08, 0x27, 0xb9, 0xfd, 0x44, 0x51, 0x34, 0x7e, 0x9b,
		0xdb, 0xbb, 0x4b, 0x1c, 0xac, 0x84, 0x81, 0x77, 0x95, 0x0d, 0x68, 0x12,
		0xab, 0xaa, 0xd4, 0xcb, 0xc5, 0x0a, 0xea, 0xcb, 0xb7, 0x2e, 0x58, 0xc0,
		0xe2, 0xf8, 0xab, 0x0d, 0x55, 0xfd, 0xf9, 0x50, 0x2d, 0x4e, 0xb6, 0x44,
		0xfa, 0x11, 0xa9, 0xd4, 0x92, 0x7d, 0x91, 0x4f, 0x41, 0x4d, 0xb2, 0x91,
		0x81, 0x52, 0x73, 0xc7, 0x93, 0xf0, 0xcd, 0x68, 0xce, 0x79, 0x1f, 0x42,
		0x2f, 0xff, 0x21, 0x5a, 0xba, 0xcf, 0x91, 0x33, 0xed, 0x6c, 0x4d, 0x7a,
		0xfc, 0x77, 0xc4, 0x74, 0xd1, 0x45, 0x5e, 0x14, 0xf9, 0xfe, 0x51, 0xe5,
		0x93, 0xf5, 0xf9, 0x48, 0xab, 0xfb, 0x15, 0xf8, 0xc2, 0x96, 0x6f, 0x87,
		0x5c, 0xcb, 0x1a, 0x10, 0x27, 0x13, 0x57, 0x64, 0x6b, 0x89, 0xf5, 0x48,
		0xfa, 0x79, 0xfa, 0x4a, 0x99, 0x67, 0xf4, 0xad, 0x6c, 0xc1, 0x5b, 0x00,
		0x00, 0x00, 0xff, 0xff, 0xa9, 0xef, 0x18, 0x09, 0x4b, 0x02, 0x00, 0x00,
	},
		"templates/base/footer.html",
	)
}

func templates_base_header_html() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xb4, 0xcf,
		0x31, 0x0e, 0x02, 0x21, 0x10, 0x85, 0xe1, 0xde, 0x53, 0x90, 0xe9, 0x05,
		0x93, 0x6d, 0x5d, 0xef, 0x82, 0xcc, 0xdb, 0x40, 0x84, 0xc1, 0x30, 0xd3,
		0x18, 0xb2, 0x77, 0x77, 0x2b, 0xf5, 0x02, 0xb6, 0x7f, 0x5e, 0xf2, 0xe5,
		0xcd, 0xc9, 0xd8, 0x8a, 0xc0, 0x51, 0x46, 0x64, 0x0c, 0xda, 0xf7, 0x93,
		0x73, 0xd7, 0x5a, 0xe4, 0xe1, 0x06, 0xea, 0x4a, 0x6a, 0xaf, 0x0a, 0xcd,
		0x80, 0x91, 0xcb, 0x03, 0xdb, 0x4a, 0x21, 0x08, 0x8c, 0x25, 0xfa, 0x7b,
		0xef, 0xa6, 0x36, 0xe2, 0x33, 0xb1, 0xf8, 0xd4, 0x5b, 0xf8, 0x84, 0xb0,
		0xf8, 0x8b, 0x5f, 0x42, 0x52, 0xfd, 0x36, 0xdf, 0xca, 0xb1, 0x52, 0xa5,
		0xdb, 0x9f, 0x80, 0xb3, 0x65, 0x34, 0xfc, 0x30, 0x73, 0x42, 0xf8, 0xf8,
		0xf3, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x6a, 0x24, 0xeb, 0x71, 0xe4, 0x00,
		0x00, 0x00,
	},
		"templates/base/header.html",
	)
}

func templates_category_html() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x54, 0x8e,
		0x41, 0xab, 0xc3, 0x20, 0x10, 0x84, 0xef, 0xfe, 0x8a, 0x45, 0x72, 0x7c,
		0x28, 0xef, 0xfa, 0xd8, 0x08, 0x8f, 0xfc, 0x84, 0xf4, 0x5e, 0x24, 0xd9,
		0x18, 0x41, 0x0c, 0xd8, 0xcd, 0xa1, 0x58, 0xff, 0x7b, 0x4d, 0xd2, 0xd0,
		0xf6, 0xa4, 0x33, 0xcc, 0x7e, 0x33, 0x39, 0x8f, 0x34, 0xf9, 0x48, 0x20,
		0xd9, 0x73, 0x20, 0x59, 0x8a, 0xc8, 0x59, 0x75, 0x96, 0xc9, 0x2d, 0xe9,
		0x5e, 0x0a, 0x3c, 0xa0, 0xea, 0xde, 0x33, 0xa9, 0x6e, 0x89, 0x93, 0x77,
		0xea, 0xb2, 0xe5, 0xf6, 0x18, 0xc5, 0xb1, 0xbe, 0xe2, 0x8d, 0x18, 0x96,
		0xc8, 0x14, 0x79, 0x83, 0xe0, 0xfc, 0x6b, 0xbe, 0x40, 0xa8, 0xab, 0x23,
		0x70, 0x0d, 0x46, 0x40, 0x45, 0x26, 0x1b, 0x1d, 0x41, 0x73, 0xfd, 0x81,
		0xc6, 0x26, 0xf6, 0x43, 0x20, 0xf8, 0x6b, 0x41, 0xfd, 0x1f, 0xff, 0x5b,
		0x25, 0x00, 0x60, 0xf0, 0x06, 0x2d, 0xcc, 0x89, 0xa6, 0x56, 0xe6, 0x7c,
		0x06, 0x55, 0x1f, 0x56, 0x57, 0x8a, 0x34, 0x1f, 0xd6, 0x6b, 0x14, 0x6a,
		0x6b, 0x50, 0xd7, 0xb3, 0xbd, 0xe3, 0xd8, 0x87, 0x7a, 0xeb, 0x3c, 0xd5,
		0x33, 0x00, 0x00, 0xff, 0xff, 0x03, 0x42, 0x6c, 0xbd, 0xf0, 0x00, 0x00,
		0x00,
	},
		"templates/category.html",
	)
}

func templates_index_html() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x94, 0x55,
		0x4d, 0x6f, 0xa3, 0x3c, 0x10, 0xbe, 0xf3, 0x2b, 0x5c, 0xde, 0xe8, 0x3d,
		0x54, 0x5b, 0x50, 0xaf, 0xbb, 0x04, 0xa9, 0xea, 0x1e, 0x7a, 0xe8, 0x66,
		0x2b, 0x35, 0xf7, 0xca, 0x81, 0x01, 0x2c, 0x19, 0x93, 0xd8, 0x26, 0x6a,
		0x16, 0xe5, 0xbf, 0xef, 0xf8, 0x83, 0x80, 0xb3, 0x54, 0xdb, 0x8d, 0x94,
		0x18, 0xe6, 0xf3, 0x99, 0x99, 0x67, 0x9c, 0x61, 0x28, 0xa1, 0x62, 0x02,
		0x48, 0xac, 0x99, 0xe6, 0x10, 0x9f, 0xcf, 0xd1, 0x30, 0x24, 0xaf, 0x4c,
		0x43, 0xf2, 0xd8, 0x89, 0x8a, 0xd5, 0xc9, 0xd6, 0xc8, 0xad, 0x18, 0x44,
		0x89, 0x67, 0x34, 0xb9, 0x14, 0x9d, 0xd0, 0x20, 0xb4, 0x73, 0x62, 0x15,
		0x49, 0x1e, 0xa4, 0x66, 0x05, 0x07, 0xe5, 0xcc, 0x24, 0x15, 0x35, 0x90,
		0xd5, 0xdb, 0x17, 0xb2, 0xa2, 0x4e, 0x41, 0xbe, 0xae, 0x03, 0xa3, 0xcc,
		0xcb, 0xf3, 0x28, 0x6b, 0xee, 0xf3, 0x8c, 0x92, 0x46, 0x42, 0xb5, 0x8e,
		0x87, 0x61, 0x74, 0x48, 0x5e, 0x79, 0x5f, 0x9f, 0xcf, 0x71, 0x3e, 0x13,
		0x79, 0x40, 0x59, 0x4a, 0xf3, 0x2c, 0x45, 0xb7, 0x28, 0xca, 0xf6, 0xb9,
		0x03, 0xd0, 0x49, 0xb2, 0x0a, 0xc0, 0x3f, 0xf4, 0xba, 0x31, 0xc2, 0xd1,
		0xd7, 0xbd, 0x9b, 0xcc, 0x6a, 0x4f, 0x05, 0x29, 0x38, 0x55, 0x6a, 0x1d,
		0x73, 0xba, 0x03, 0x4e, 0xec, 0xef, 0x1d, 0xd6, 0x46, 0x7b, 0xae, 0xe3,
		0x7c, 0x77, 0x22, 0x36, 0xe6, 0x1f, 0xbe, 0x33, 0x28, 0x93, 0x08, 0xb8,
		0x02, 0xab, 0x5a, 0x48, 0x6f, 0xf5, 0xa6, 0x77, 0x59, 0x6a, 0xb2, 0xe6,
		0xf3, 0x56, 0xce, 0x13, 0x7c, 0xa7, 0x1a, 0x3e, 0x05, 0xcd, 0xa4, 0x79,
		0xea, 0x5b, 0x2a, 0xd8, 0x2f, 0x30, 0x4e, 0x9a, 0xb5, 0x70, 0x1d, 0xe6,
		0x2f, 0xb9, 0x1e, 0xd1, 0xa8, 0xee, 0xe4, 0xc9, 0x0e, 0x61, 0x29, 0x99,
		0xea, 0x8b, 0x02, 0x94, 0x8a, 0xfd, 0x48, 0xd2, 0xc2, 0x3b, 0xa4, 0xb3,
		0xf2, 0xa7, 0x20, 0x49, 0xa3, 0x5b, 0x1e, 0x4c, 0x69, 0xd2, 0x99, 0x41,
		0x7d, 0x04, 0x63, 0x4b, 0x6b, 0x65, 0xd9, 0x33, 0xe3, 0x8a, 0xa6, 0xb5,
		0xe1, 0xc9, 0xb5, 0xcd, 0x32, 0x4c, 0x26, 0xaa, 0xee, 0x82, 0x11, 0x3d,
		0x0d, 0x3c, 0x3c, 0xe6, 0x88, 0xec, 0x6b, 0x00, 0x62, 0x3c, 0xb3, 0x14,
		0x89, 0x13, 0xcd, 0x40, 0x3f, 0x6d, 0x7f, 0x3c, 0xcf, 0x7a, 0xe4, 0xe8,
		0x6d, 0x0c, 0xf7, 0x63, 0x6e, 0x0d, 0xef, 0xfa, 0x4e, 0xb2, 0xba, 0xc1,
		0x31, 0x4c, 0x88, 0x76, 0x5a, 0x10, 0xfc, 0x5e, 0x26, 0xf4, 0x11, 0x8f,
		0xff, 0x2b, 0x99, 0x3a, 0xf4, 0xea, 0x4d, 0xa3, 0x9e, 0x96, 0x71, 0xfe,
		0x0c, 0xf4, 0x08, 0x04, 0xc3, 0x74, 0x6d, 0x8b, 0x99, 0x6e, 0x1c, 0xa7,
		0x11, 0x14, 0x3e, 0x4c, 0x9b, 0x21, 0x83, 0xf6, 0x39, 0xa6, 0xd9, 0x85,
		0xd9, 0x74, 0xba, 0x61, 0xa2, 0x26, 0xba, 0x23, 0x0a, 0x80, 0x34, 0x20,
		0xe1, 0xc6, 0xad, 0xc4, 0xb5, 0x93, 0x59, 0x4e, 0x4b, 0xcc, 0x4d, 0xdf,
		0xee, 0x40, 0xfe, 0xac, 0x5e, 0x68, 0xed, 0x36, 0xb0, 0x64, 0xc7, 0xa0,
		0xb6, 0x02, 0x71, 0x80, 0xc4, 0xe2, 0x7a, 0x3e, 0xca, 0xf7, 0xb4, 0x66,
		0x82, 0x6a, 0xd6, 0x89, 0x18, 0xdb, 0x45, 0xf0, 0x93, 0x71, 0x36, 0x2a,
		0xed, 0xbb, 0x4d, 0x00, 0x07, 0x5c, 0x3e, 0x13, 0xd7, 0xe5, 0x20, 0xf7,
		0x18, 0xde, 0x28, 0xb1, 0x66, 0xba, 0xe3, 0x50, 0x7a, 0x4b, 0x87, 0xc9,
		0x3c, 0x63, 0x34, 0x1b, 0x6c, 0xb6, 0xf6, 0xc9, 0x8b, 0x84, 0x23, 0xeb,
		0x7a, 0x65, 0xfa, 0x15, 0x84, 0x33, 0x97, 0xc0, 0xff, 0x9c, 0x1e, 0xfa,
		0xee, 0x9b, 0x9d, 0xa5, 0xf5, 0x4c, 0x39, 0xf3, 0x88, 0x02, 0x06, 0x21,
		0xe0, 0xf1, 0xaa, 0x91, 0xf4, 0x34, 0x15, 0x3b, 0x42, 0xbf, 0xc0, 0xb5,
		0x86, 0x61, 0x16, 0x5f, 0x16, 0x2d, 0x34, 0x3b, 0x42, 0xec, 0xe1, 0x5e,
		0x01, 0x4d, 0x99, 0x28, 0xe1, 0xdd, 0x46, 0xc1, 0x6b, 0xd0, 0x45, 0xb9,
		0xb7, 0xdb, 0x6f, 0x1e, 0x2f, 0xeb, 0xee, 0x29, 0x68, 0x7d, 0x0d, 0x42,
		0xaf, 0xf5, 0xd0, 0x83, 0x1a, 0x66, 0xad, 0xc1, 0x33, 0xbd, 0x25, 0xae,
		0x9a, 0xdb, 0xd4, 0x8c, 0xef, 0xd3, 0x1d, 0x5f, 0x2d, 0xcf, 0xf8, 0x9f,
		0x87, 0xb0, 0x41, 0x22, 0x2c, 0x0f, 0x40, 0x2e, 0x0e, 0x20, 0x4b, 0x7b,
		0x8e, 0xbc, 0x45, 0x2a, 0x5d, 0x48, 0xe7, 0xaa, 0x40, 0x8c, 0xb6, 0x84,
		0x40, 0xe8, 0xff, 0x35, 0x9c, 0xe6, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff,
		0xc6, 0x53, 0x85, 0xab, 0x80, 0x06, 0x00, 0x00,
	},
		"templates/index.html",
	)
}

func templates_page_html() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xaa, 0xae,
		0x4e, 0x49, 0x4d, 0xcb, 0xcc, 0x4b, 0x55, 0x50, 0x2a, 0xc9, 0x2c, 0xc9,
		0x49, 0x55, 0xaa, 0xad, 0xe5, 0xaa, 0xae, 0xd6, 0x0b, 0x48, 0x4c, 0x4f,
		0xd5, 0x0b, 0x01, 0x09, 0xd4, 0xd6, 0x2a, 0xd4, 0x28, 0x00, 0x45, 0x82,
		0x33, 0x4b, 0x52, 0xf5, 0x9c, 0xf3, 0xf3, 0xd2, 0x32, 0xd3, 0x61, 0x12,
		0x40, 0x85, 0xa9, 0x79, 0x29, 0x40, 0x9a, 0x0b, 0x61, 0x48, 0x72, 0x7e,
		0x5e, 0x49, 0x6a, 0x5e, 0x09, 0xc8, 0x18, 0x9b, 0x0c, 0x43, 0x3b, 0x34,
		0xa3, 0x6c, 0xf4, 0x81, 0x62, 0x70, 0xf3, 0x3d, 0x42, 0x7c, 0x7d, 0x14,
		0x20, 0x4c, 0x67, 0x88, 0x36, 0x24, 0x33, 0x01, 0x01, 0x00, 0x00, 0xff,
		0xff, 0x27, 0x7a, 0x43, 0x12, 0x98, 0x00, 0x00, 0x00,
	},
		"templates/page.html",
	)
}

func templates_tag_html() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x4c, 0x8e,
		0xc1, 0x0a, 0x83, 0x30, 0x10, 0x44, 0xef, 0xf9, 0x8a, 0x25, 0x78, 0x2c,
		0x09, 0xbd, 0x96, 0x35, 0x50, 0xfa, 0x09, 0x7a, 0x2f, 0x41, 0x57, 0x0d,
		0x84, 0x08, 0x36, 0x9e, 0xb6, 0xf9, 0xf7, 0xae, 0x5a, 0x69, 0x4f, 0xd9,
		0x09, 0xf3, 0x66, 0x86, 0xb9, 0xa7, 0x21, 0x24, 0x02, 0x9d, 0x43, 0x8e,
		0xa4, 0x4b, 0x51, 0xcc, 0xa6, 0xf5, 0x63, 0x29, 0xf0, 0x06, 0x39, 0x9b,
		0x90, 0xc9, 0x3c, 0xe6, 0x34, 0x84, 0xd1, 0xb4, 0x9b, 0x65, 0x77, 0x50,
		0xea, 0xe5, 0x55, 0x3f, 0xba, 0x9b, 0x53, 0xa6, 0x94, 0x37, 0x1e, 0xa7,
		0xab, 0x3b, 0x33, 0xd0, 0x8a, 0x50, 0xb8, 0x46, 0xa7, 0x40, 0xd2, 0x16,
		0x9f, 0x46, 0x82, 0xea, 0x79, 0x81, 0xca, 0x2f, 0x39, 0x74, 0x91, 0xe0,
		0x56, 0x83, 0xb9, 0x1f, 0xf7, 0x4b, 0x60, 0x00, 0x8c, 0xc1, 0xa1, 0x87,
		0x69, 0xa1, 0xa1, 0xd6, 0xcc, 0xa7, 0xd1, 0x34, 0x71, 0x95, 0x40, 0xed,
		0xfe, 0xbe, 0xbe, 0x7b, 0xd0, 0x7a, 0x87, 0x56, 0xb0, 0xbd, 0xe3, 0x98,
		0x86, 0x76, 0xeb, 0x3c, 0xd5, 0x27, 0x00, 0x00, 0xff, 0xff, 0xd0, 0xa4,
		0x2a, 0x22, 0xe6, 0x00, 0x00, 0x00,
	},
		"templates/tag.html",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() ([]byte, error){
	"templates/archives.html": templates_archives_html,
	"templates/article.html": templates_article_html,
	"templates/atom.xml": templates_atom_xml,
	"templates/base/analytics.html": templates_base_analytics_html,
	"templates/base/base.html": templates_base_base_html,
	"templates/base/footer.html": templates_base_footer_html,
	"templates/base/header.html": templates_base_header_html,
	"templates/category.html": templates_category_html,
	"templates/index.html": templates_index_html,
	"templates/page.html": templates_page_html,
	"templates/tag.html": templates_tag_html,
}