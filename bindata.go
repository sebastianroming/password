package main

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

func frontend_assets_application_coffee() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x94, 0x55,
		0x4d, 0x6f, 0xdb, 0x30, 0x0c, 0xbd, 0xeb, 0x57, 0x08, 0x69, 0x36, 0xdb,
		0x43, 0xe2, 0x7e, 0xec, 0x56, 0x20, 0xdd, 0x65, 0x97, 0x0e, 0xdd, 0x5a,
		0x2c, 0xbb, 0x0d, 0x3b, 0xa8, 0x36, 0xed, 0x08, 0x55, 0x6c, 0x43, 0x92,
		0x63, 0x14, 0x45, 0xfe, 0xfb, 0x48, 0x59, 0x8e, 0xed, 0xc6, 0x69, 0xb7,
		0x5e, 0x6a, 0x91, 0x8f, 0x8f, 0xe4, 0x23, 0xa5, 0xc8, 0x8c, 0x37, 0xb2,
		0x48, 0xcb, 0x26, 0x5e, 0xdb, 0x52, 0x8b, 0x1c, 0xb8, 0x28, 0xd2, 0xce,
		0xf4, 0x6d, 0x7d, 0xff, 0x83, 0xf1, 0xee, 0x34, 0x37, 0x1e, 0xb1, 0xe2,
		0xe1, 0x13, 0x3c, 0x47, 0x7c, 0x79, 0x83, 0x4e, 0xce, 0x0d, 0xd8, 0x6b,
		0x1e, 0xee, 0x84, 0xaa, 0xe1, 0x60, 0xe3, 0x5c, 0x95, 0x89, 0x50, 0x9e,
		0x33, 0x46, 0xc8, 0xad, 0x85, 0x2d, 0x85, 0x2d, 0x38, 0xb1, 0xc6, 0xc6,
		0x6a, 0x59, 0xe4, 0x32, 0x7b, 0xf6, 0x81, 0x91, 0x0b, 0xcb, 0x89, 0xea,
		0x40, 0x21, 0x31, 0x04, 0x93, 0x8d, 0x98, 0xf2, 0x9e, 0x29, 0xf2, 0x30,
		0xc7, 0x57, 0x09, 0x6d, 0x20, 0xa4, 0x88, 0x88, 0x63, 0x4f, 0xf4, 0xc1,
		0xd8, 0xbc, 0xe5, 0xf2, 0xf5, 0x2b, 0x61, 0xec, 0x5d, 0x29, 0x52, 0xa4,
		0x2c, 0xca, 0x26, 0x8c, 0x7a, 0x8f, 0x86, 0x4c, 0x83, 0xd9, 0x3c, 0x08,
		0x63, 0x9a, 0x52, 0x13, 0xe0, 0xf3, 0x05, 0xfe, 0xf5, 0x00, 0x2b, 0x93,
		0x27, 0xd0, 0xb7, 0x85, 0x05, 0x8d, 0xe5, 0xa2, 0xff, 0x0a, 0xbd, 0xaf,
		0xdd, 0x68, 0xa6, 0x3e, 0x3b, 0x10, 0x19, 0x17, 0xd3, 0x04, 0x18, 0x39,
		0x0f, 0x83, 0xb3, 0xac, 0x4c, 0x6a, 0x03, 0xe9, 0x6d, 0x51, 0xd5, 0x36,
		0x88, 0xe2, 0x47, 0x84, 0xf2, 0x20, 0x51, 0x08, 0x0d, 0x16, 0x1c, 0xc5,
		0xae, 0x7e, 0xb6, 0x75, 0xbd, 0x05, 0x7f, 0x54, 0xb5, 0x46, 0x34, 0xc2,
		0xac, 0xd0, 0x76, 0x1c, 0x50, 0x56, 0x56, 0x96, 0xc5, 0x5a, 0xec, 0xe0,
		0x98, 0x1d, 0x8d, 0xf7, 0xce, 0x6d, 0xa8, 0x0f, 0x85, 0xb2, 0xf8, 0xa3,
		0xd3, 0x85, 0xce, 0x9d, 0x1a, 0x68, 0x60, 0xa8, 0x17, 0x4d, 0xdd, 0x8f,
		0xd7, 0x29, 0x08, 0x0d, 0xff, 0x2a, 0x2c, 0x38, 0x78, 0x4a, 0x63, 0xf9,
		0x25, 0xb7, 0x74, 0x62, 0x83, 0xc2, 0x07, 0x31, 0x89, 0x02, 0x71, 0xe8,
		0x3f, 0x1c, 0xa9, 0x12, 0x4d, 0xf7, 0x67, 0x40, 0x41, 0x62, 0x1d, 0x7f,
		0x26, 0x94, 0x01, 0xc6, 0xc6, 0x4d, 0x0e, 0xc8, 0xdf, 0x1d, 0xf0, 0x7f,
		0x0f, 0x88, 0x21, 0xf0, 0x41, 0x97, 0x39, 0x26, 0x32, 0x94, 0xa8, 0x02,
		0x9d, 0xf8, 0x64, 0x58, 0x6a, 0x5c, 0x79, 0xd7, 0xf2, 0x51, 0x68, 0x2c,
		0x35, 0x31, 0x26, 0x0c, 0x1a, 0x99, 0xda, 0x0d, 0x4a, 0x3b, 0x3b, 0x7b,
		0x21, 0xf4, 0xfe, 0xc3, 0x0c, 0xc5, 0x18, 0x0a, 0x39, 0x28, 0xb8, 0x9d,
		0x8c, 0x71, 0xcb, 0x3d, 0x56, 0x7e, 0x4e, 0x52, 0xf2, 0xd9, 0xf9, 0xee,
		0xf2, 0x1c, 0x3f, 0xba, 0xd0, 0x2f, 0x0a, 0x8a, 0xdc, 0x6e, 0x56, 0x67,
		0x2f, 0x3e, 0x12, 0xf7, 0xbd, 0xf5, 0xdc, 0x39, 0xc7, 0xfe, 0xa3, 0xa9,
		0x20, 0x91, 0x42, 0x0d, 0x10, 0xa8, 0xe5, 0xba, 0x35, 0xee, 0x67, 0x0b,
		0x1e, 0xa6, 0xc2, 0x8a, 0xc3, 0xf5, 0x9c, 0x90, 0x9b, 0xc6, 0xe2, 0x30,
		0x0e, 0x70, 0x4a, 0x51, 0x36, 0x58, 0x9c, 0xa9, 0x7e, 0x5c, 0xf0, 0xb8,
		0xb6, 0x6b, 0x97, 0x6d, 0x6c, 0x6b, 0x19, 0x7c, 0xd6, 0x36, 0x63, 0x5f,
		0x6e, 0x1b, 0xd0, 0x9f, 0x3b, 0xf0, 0xef, 0x8b, 0x3f, 0x71, 0xb2, 0x01,
		0x9c, 0x53, 0xca, 0x8e, 0x5f, 0xa5, 0x30, 0x58, 0x43, 0x52, 0x6b, 0xe8,
		0x24, 0xf3, 0x45, 0xba, 0x45, 0xb2, 0xa1, 0x2f, 0xb0, 0x5b, 0x35, 0x34,
		0x59, 0x7c, 0x7e, 0xcc, 0xf7, 0x32, 0x15, 0x0a, 0x21, 0x5b, 0xfa, 0x1f,
		0x06, 0x1b, 0x99, 0xe2, 0x55, 0x99, 0xd8, 0xff, 0xc1, 0x94, 0x26, 0xa7,
		0xf8, 0xcf, 0xb5, 0xe0, 0x4c, 0x5d, 0xbf, 0xf8, 0x44, 0x1d, 0xa2, 0x57,
		0xbc, 0x2e, 0x52, 0xc8, 0x64, 0x81, 0x7d, 0x91, 0x12, 0x63, 0x31, 0x8f,
		0xe5, 0xbc, 0xba, 0xf0, 0x8e, 0xa1, 0x64, 0xed, 0x15, 0xe1, 0xef, 0x69,
		0x3d, 0xbd, 0x3e, 0x9d, 0x2c, 0x6f, 0x6a, 0x8e, 0x6d, 0x1e, 0xaf, 0x56,
		0xaf, 0x02, 0x63, 0x74, 0x85, 0x86, 0xaf, 0x84, 0xcc, 0xb2, 0x6e, 0x6f,
		0xf8, 0xf2, 0xf5, 0x46, 0x21, 0x80, 0x6e, 0x09, 0xe1, 0x4f, 0x3c, 0xc2,
		0x4b, 0xc7, 0x10, 0xf1, 0xf3, 0x53, 0xaf, 0xf4, 0x27, 0x7e, 0xe9, 0xde,
		0xe8, 0xe1, 0x55, 0x25, 0xce, 0x56, 0x5e, 0x97, 0xfe, 0x66, 0x75, 0x22,
		0x98, 0xb5, 0x3f, 0x51, 0xa3, 0x29, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff,
		0x19, 0x94, 0x55, 0x68, 0x0b, 0x07, 0x00, 0x00,
	},
		"frontend/assets/application.coffee",
	)
}

func frontend_assets_application_js() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xa4, 0x56,
		0x4d, 0x6f, 0xdb, 0x38, 0x10, 0xbd, 0xfb, 0x57, 0x10, 0x5e, 0x6f, 0x24,
		0xed, 0x2a, 0xb2, 0x9d, 0x3d, 0x2c, 0x10, 0xc3, 0xbb, 0x87, 0x16, 0x28,
		0x52, 0xa4, 0x4d, 0x50, 0xf7, 0x56, 0xf4, 0x40, 0x4b, 0x94, 0x4d, 0x44,
		0x11, 0x05, 0x92, 0xb6, 0x11, 0x14, 0xfa, 0xef, 0x1d, 0x7e, 0x88, 0x22,
		0xe5, 0x8f, 0xf4, 0x23, 0x17, 0x8b, 0xe4, 0x70, 0xe6, 0xf1, 0xbd, 0xc7,
		0x61, 0xa6, 0x53, 0xf4, 0x8e, 0xd4, 0x84, 0x63, 0x49, 0x0a, 0xb4, 0x7e,
		0x41, 0x6f, 0x58, 0x59, 0x12, 0xb2, 0xca, 0x39, 0x6d, 0x24, 0x9a, 0x67,
		0xff, 0x66, 0xf3, 0x51, 0x5c, 0xee, 0xea, 0x5c, 0x52, 0x56, 0xc7, 0x09,
		0xfa, 0x36, 0x42, 0x68, 0x8f, 0x39, 0xaa, 0x18, 0x2e, 0x1e, 0x1a, 0x35,
		0x29, 0x52, 0x3d, 0x78, 0xc4, 0x42, 0x1c, 0x18, 0x2f, 0x52, 0x54, 0xb3,
		0x43, 0x8a, 0x38, 0x11, 0x12, 0x73, 0xf9, 0x89, 0x94, 0xf0, 0xb5, 0x4d,
		0x91, 0xc0, 0x7b, 0xe2, 0xe2, 0x05, 0x91, 0x8f, 0x9c, 0x6d, 0x60, 0x45,
		0x0d, 0x24, 0x6b, 0x5c, 0x98, 0xa4, 0xf9, 0xd3, 0x62, 0x04, 0x35, 0x68,
		0x89, 0xe2, 0x03, 0xad, 0x0b, 0x76, 0xc8, 0x56, 0x92, 0x71, 0xbc, 0x21,
		0xe8, 0xea, 0x0a, 0xd9, 0x99, 0xf7, 0xab, 0x87, 0x8f, 0x06, 0x0a, 0xea,
		0xa6, 0x26, 0xc2, 0x46, 0x2d, 0x91, 0x43, 0xfb, 0x44, 0x5e, 0xba, 0x28,
		0x04, 0x80, 0xe4, 0x8e, 0xd7, 0x6e, 0x88, 0x14, 0x88, 0xdb, 0x3e, 0x76,
		0x8f, 0xab, 0x1d, 0x49, 0xbc, 0x65, 0xb7, 0xa3, 0x62, 0x39, 0xae, 0x2c,
		0x86, 0x0c, 0x36, 0xdd, 0x49, 0xf2, 0xac, 0x52, 0xa7, 0x48, 0xc1, 0xc8,
		0x84, 0xe4, 0xb4, 0xde, 0xd0, 0xf2, 0xc5, 0xa6, 0x48, 0x16, 0x2e, 0x45,
		0x9b, 0xba, 0xcf, 0x4d, 0x50, 0x2c, 0xac, 0xa3, 0xe8, 0xa4, 0x90, 0x74,
		0xe1, 0xcd, 0xa9, 0x31, 0x1c, 0x25, 0xa8, 0xbd, 0xe9, 0x6b, 0x27, 0x41,
		0x2c, 0x50, 0xa5, 0xe2, 0xc3, 0xac, 0x0e, 0xbf, 0x46, 0xd9, 0x60, 0x2e,
		0x88, 0x89, 0xf2, 0xb7, 0xb6, 0xa3, 0xe1, 0x57, 0x6b, 0x96, 0xf5, 0x4f,
		0xab, 0x84, 0x98, 0x0c, 0xd5, 0x77, 0x94, 0x57, 0x58, 0xc8, 0x7b, 0x50,
		0x1e, 0x70, 0x82, 0xe4, 0xb1, 0x4d, 0x6c, 0x17, 0xb9, 0x51, 0xb4, 0x73,
		0x05, 0xc4, 0xfc, 0x33, 0x83, 0xbf, 0x20, 0x46, 0x89, 0x4d, 0xf8, 0x5d,
		0x2d, 0x09, 0x07, 0xee, 0x20, 0xe4, 0xe6, 0x64, 0x00, 0x2c, 0x28, 0xda,
		0x6d, 0x58, 0xac, 0x26, 0xd3, 0xd3, 0x29, 0x2c, 0x84, 0x49, 0x1c, 0xfd,
		0x51, 0xb2, 0x7c, 0x27, 0x48, 0x71, 0x57, 0x37, 0x3b, 0x19, 0x25, 0xd9,
		0x1a, 0xc2, 0xe3, 0x28, 0xaf, 0x20, 0x3c, 0x0a, 0x2c, 0xf7, 0xda, 0x96,
		0x75, 0xb5, 0xe3, 0xd1, 0xd0, 0xce, 0xde, 0x26, 0xa6, 0x3d, 0xbd, 0x02,
		0x77, 0x1f, 0x57, 0xe9, 0x2d, 0x6f, 0x37, 0x78, 0x97, 0xa6, 0xa3, 0xcb,
		0x99, 0xac, 0xbf, 0x41, 0x66, 0xa9, 0x4d, 0xf4, 0x3d, 0x00, 0x66, 0x7d,
		0x4b, 0x77, 0x12, 0x28, 0xcf, 0x14, 0x26, 0x83, 0xe6, 0x9f, 0x1c, 0xd0,
		0x5b, 0xb8, 0xc1, 0x83, 0xac, 0x85, 0xf2, 0xcc, 0x67, 0xfa, 0x6c, 0xe7,
		0x5b, 0x9d, 0xd1, 0x3b, 0xfc, 0xa9, 0xcc, 0x79, 0x45, 0xb0, 0x23, 0x34,
		0x0e, 0x68, 0x3e, 0xcf, 0x95, 0x20, 0x15, 0xc9, 0xe5, 0xa0, 0x7a, 0x89,
		0x2b, 0x41, 0xfa, 0xba, 0x21, 0x85, 0xa7, 0x4a, 0x5f, 0xf4, 0x95, 0x4d,
		0xfa, 0x3b, 0xce, 0xb0, 0xe7, 0xef, 0x9b, 0x8f, 0x0f, 0xa2, 0x21, 0x3c,
		0xef, 0x80, 0xd8, 0x5a, 0x70, 0xd0, 0xac, 0xb1, 0xa1, 0xd7, 0x6b, 0xcc,
		0xe1, 0xa0, 0xb9, 0x10, 0x71, 0x74, 0xa0, 0x85, 0xdc, 0x82, 0xc0, 0xe3,
		0x31, 0xfa, 0x1b, 0xa9, 0x7d, 0xf0, 0x33, 0xfe, 0x73, 0xec, 0xd5, 0xf0,
		0xe5, 0x3c, 0x27, 0x9f, 0x71, 0x8e, 0x30, 0xa7, 0xb3, 0x03, 0x7d, 0xe5,
		0xcf, 0x79, 0x64, 0xa2, 0xd4, 0x8c, 0xc7, 0xd3, 0xfd, 0x7c, 0x0a, 0x1f,
		0x5d, 0xfa, 0xff, 0x2b, 0x52, 0x6f, 0xe4, 0x76, 0xa9, 0xb0, 0xd8, 0x2c,
		0x70, 0xdb, 0xcd, 0xda, 0xbd, 0x5e, 0x52, 0xe8, 0xae, 0x44, 0x43, 0x72,
		0x8a, 0xab, 0x20, 0x0c, 0x14, 0x5c, 0x99, 0xe9, 0xb4, 0x87, 0x58, 0x60,
		0x89, 0xfb, 0x4e, 0x72, 0x42, 0x6b, 0xc5, 0xb4, 0x0e, 0x5a, 0x84, 0x9d,
		0xf5, 0xa2, 0x7a, 0xad, 0x2f, 0x40, 0x7f, 0x2f, 0x7e, 0x92, 0x9b, 0x0e,
		0x55, 0x78, 0xbe, 0x5b, 0x8d, 0x32, 0x9c, 0x33, 0xf9, 0x2d, 0xda, 0xa4,
		0xeb, 0xc3, 0xfd, 0x81, 0xcd, 0x9e, 0x7e, 0xdc, 0xc5, 0x7f, 0x99, 0x7d,
		0xcd, 0xf2, 0x2d, 0x01, 0xeb, 0x14, 0x23, 0xaf, 0x1d, 0x0e, 0x1e, 0x99,
		0x38, 0x5a, 0x91, 0x7c, 0xc7, 0x49, 0xa7, 0x81, 0x3d, 0x8d, 0xbe, 0x08,
		0x32, 0x66, 0xc1, 0x9d, 0x57, 0x75, 0x60, 0x56, 0xc2, 0x23, 0x21, 0x3e,
		0xb0, 0x02, 0x57, 0x10, 0xf5, 0xac, 0x7e, 0xe3, 0x68, 0x4b, 0x0b, 0xe8,
		0x1a, 0xaf, 0xf5, 0x01, 0x67, 0xa8, 0x5f, 0xe3, 0xec, 0x87, 0xa1, 0x2b,
		0x73, 0x59, 0x30, 0xea, 0x45, 0x71, 0x19, 0x96, 0x4b, 0xb4, 0x67, 0xb4,
		0x40, 0xb3, 0xde, 0x15, 0xc7, 0x8a, 0x1c, 0x6b, 0x72, 0x33, 0xeb, 0x5f,
		0x3f, 0x9f, 0x77, 0xdd, 0x15, 0x06, 0x8f, 0x8d, 0x23, 0xea, 0x82, 0x88,
		0xa7, 0xbd, 0xed, 0x91, 0x7c, 0x51, 0x4c, 0x80, 0x7a, 0xec, 0xfa, 0x80,
		0x79, 0x8f, 0x3e, 0x43, 0xb9, 0x6a, 0x20, 0x67, 0x5b, 0x2f, 0x2d, 0xcb,
		0x54, 0x5f, 0x7d, 0xdb, 0x83, 0x61, 0xdc, 0x19, 0x1e, 0x5d, 0x0f, 0xaf,
		0x82, 0x89, 0xd1, 0x8d, 0x62, 0xe9, 0xfe, 0xab, 0x19, 0x3e, 0x90, 0xd7,
		0x3a, 0x49, 0x82, 0xa6, 0xe7, 0x5e, 0xd0, 0xbf, 0xd0, 0xbc, 0x7b, 0x1e,
		0xbd, 0x16, 0x66, 0xfa, 0x56, 0x2f, 0x9b, 0x46, 0xf2, 0xdf, 0xf2, 0x4c,
		0x92, 0xa3, 0x7f, 0x87, 0x8e, 0xfd, 0x66, 0xe4, 0x50, 0x14, 0xb4, 0xd0,
		0xee, 0x70, 0x05, 0x6d, 0x75, 0x4b, 0x95, 0x99, 0xbf, 0x07, 0x00, 0x00,
		0xff, 0xff, 0x75, 0xeb, 0x78, 0x8c, 0x29, 0x0a, 0x00, 0x00,
	},
		"frontend/assets/application.js",
	)
}

func frontend_index_html() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xbc, 0x58,
		0xed, 0x72, 0xdc, 0xb6, 0xd5, 0xfe, 0xef, 0xab, 0x38, 0xa6, 0xe7, 0x9d,
		0x24, 0x1e, 0x93, 0xb4, 0x24, 0xe7, 0xad, 0x2b, 0xef, 0x2a, 0xa3, 0xda,
		0xae, 0xad, 0x8c, 0x13, 0x3b, 0x91, 0x33, 0x6d, 0x26, 0x93, 0x1f, 0x20,
		0x79, 0x48, 0x42, 0x02, 0x01, 0x18, 0x00, 0x77, 0xb5, 0xcd, 0xf4, 0x3a,
		0x7a, 0x41, 0xbd, 0xb1, 0x1e, 0x80, 0xe4, 0x2e, 0x77, 0x97, 0x2b, 0x4b,
		0x49, 0xa7, 0xe3, 0x91, 0x97, 0x04, 0x0f, 0xce, 0xc7, 0x83, 0xf3, 0x89,
		0xd9, 0xc3, 0x57, 0xef, 0x5f, 0x7e, 0xfc, 0xf9, 0xc3, 0x6b, 0xa8, 0x5d,
		0x23, 0xce, 0x1e, 0xcc, 0xfc, 0x0f, 0x08, 0x26, 0xab, 0x79, 0x84, 0x32,
		0x3a, 0x7b, 0x00, 0x30, 0xab, 0x91, 0x15, 0xfe, 0x81, 0x1e, 0x1b, 0x74,
		0x0c, 0xf2, 0x9a, 0x19, 0x8b, 0x6e, 0x1e, 0xb5, 0xae, 0x8c, 0x9f, 0x47,
		0xe3, 0x4f, 0xb5, 0x73, 0x3a, 0xc6, 0x4f, 0x2d, 0x5f, 0xcc, 0xa3, 0xbf,
		0xc7, 0x3f, 0x9d, 0xc7, 0x2f, 0x55, 0xa3, 0x99, 0xe3, 0x99, 0xc0, 0x08,
		0x72, 0x25, 0x1d, 0x4a, 0xda, 0x77, 0xf1, 0x7a, 0x8e, 0x45, 0x85, 0x5b,
		0x3b, 0x25, 0x6b, 0x70, 0x1e, 0x2d, 0x38, 0x2e, 0xb5, 0x32, 0x6e, 0x44,
		0xbc, 0xe4, 0x85, 0xab, 0xe7, 0x05, 0x2e, 0x78, 0x8e, 0x71, 0x78, 0x79,
		0x02, 0x5c, 0x72, 0xc7, 0x99, 0x88, 0x6d, 0xce, 0x04, 0xce, 0x8f, 0x9e,
		0x40, 0xc3, 0x6e, 0x78, 0xd3, 0x36, 0x9b, 0x85, 0xd6, 0xa2, 0x09, 0x6f,
		0x8c, 0x44, 0xcf, 0xa5, 0x1a, 0x64, 0x3d, 0x8c, 0x63, 0xf8, 0x58, 0x23,
		0xb0, 0x4c, 0x2d, 0x10, 0x4e, 0x20, 0xc8, 0x76, 0xac, 0xb2, 0xf0, 0xb8,
		0x69, 0xad, 0x7b, 0x4c, 0x72, 0x1b, 0x84, 0x92, 0x1b, 0xeb, 0x48, 0x0a,
		0x38, 0x22, 0xf5, 0xe6, 0xbf, 0x00, 0x26, 0x57, 0xa0, 0xe8, 0xd5, 0x84,
		0xf7, 0x41, 0x3d, 0xf0, 0x9b, 0xba, 0x3d, 0x8f, 0x59, 0xe9, 0xd0, 0x3c,
		0xf6, 0x5b, 0x2c, 0x76, 0x2c, 0xe3, 0xb8, 0x97, 0xea, 0xb8, 0x13, 0x78,
		0x76, 0x89, 0x79, 0x6b, 0x10, 0x3e, 0x30, 0x6b, 0x97, 0xca, 0x14, 0xb3,
		0xb4, 0x5b, 0x7e, 0xb0, 0xd1, 0xec, 0x2f, 0x4a, 0x39, 0xeb, 0x0c, 0xd3,
		0x9b, 0xad, 0x82, 0xcb, 0x6b, 0x30, 0x28, 0xe6, 0x91, 0x75, 0x2b, 0x81,
		0xb6, 0x46, 0x24, 0x74, 0x6a, 0x83, 0xe5, 0x3c, 0xf2, 0x68, 0xdb, 0xd3,
		0x34, 0x25, 0xeb, 0xf3, 0x42, 0x26, 0xd9, 0xb0, 0xdb, 0xbf, 0x90, 0x4e,
		0xe9, 0x7a, 0x21, 0x3d, 0x49, 0x4e, 0x92, 0x67, 0x69, 0x6e, 0xed, 0x66,
		0x2d, 0x69, 0x38, 0x51, 0x59, 0x1b, 0xdd, 0x43, 0x50, 0xd8, 0xbc, 0x64,
		0x2e, 0xaf, 0x03, 0x7f, 0xcd, 0x34, 0x9a, 0x49, 0x8e, 0x1b, 0x93, 0xde,
		0x7e, 0xfc, 0xee, 0xdd, 0xd7, 0x60, 0x6b, 0xde, 0x10, 0x84, 0x05, 0xfc,
		0x88, 0x56, 0x2b, 0x59, 0x24, 0x57, 0x16, 0x4a, 0x65, 0xe0, 0xe2, 0xf5,
		0x73, 0xb0, 0xad, 0xf6, 0x47, 0x0e, 0xaa, 0xec, 0x89, 0x51, 0x60, 0x43,
		0xd8, 0xda, 0xb0, 0xa1, 0xc1, 0x82, 0x33, 0xf8, 0xd4, 0xa2, 0xe1, 0x38,
		0x42, 0xd4, 0xb3, 0xfe, 0xdb, 0xf9, 0x8f, 0xdf, 0x5f, 0x7c, 0xff, 0xe6,
		0x74, 0xcc, 0xb4, 0x50, 0x68, 0xe5, 0x17, 0x0e, 0x08, 0xe0, 0x6b, 0xe0,
		0x25, 0xac, 0x54, 0x0b, 0xde, 0xa9, 0xc2, 0x49, 0x6a, 0x56, 0x21, 0xbd,
		0x31, 0x3a, 0x5f, 0x81, 0x64, 0xce, 0x16, 0xbb, 0x5f, 0x88, 0x5a, 0x38,
		0xd2, 0x08, 0xfe, 0xfc, 0x6b, 0xb7, 0x4a, 0xeb, 0x36, 0x37, 0x5c, 0x3b,
		0xb0, 0x26, 0xdf, 0x80, 0xa0, 0xac, 0x4d, 0x7a, 0xc4, 0x3d, 0x08, 0x3e,
		0x60, 0xbe, 0x26, 0xfb, 0x16, 0x04, 0xf2, 0x9f, 0x92, 0xe3, 0xcd, 0x7b,
		0x80, 0xe3, 0x8a, 0xd0, 0x98, 0xa5, 0x1d, 0x9b, 0xfb, 0x70, 0x35, 0x9d,
		0x49, 0xe9, 0x51, 0xf2, 0x8c, 0x78, 0xf6, 0x6f, 0x07, 0x38, 0xce, 0x1e,
		0xfe, 0x82, 0xb2, 0xe0, 0xe5, 0xaf, 0x6b, 0x73, 0xc2, 0x09, 0x0e, 0xe2,
		0x32, 0x55, 0xac, 0xe0, 0x37, 0xc8, 0x58, 0x7e, 0x5d, 0x19, 0xd5, 0xca,
		0x22, 0xce, 0x95, 0x50, 0xe6, 0x14, 0x1e, 0xe1, 0xb1, 0xff, 0xf7, 0x02,
		0xfe, 0xd9, 0x53, 0x26, 0xcb, 0x9a, 0x3b, 0x9c, 0xa6, 0x2d, 0xcb, 0x72,
		0x44, 0x98, 0xd3, 0xf9, 0x50, 0x28, 0xfc, 0x06, 0x0e, 0x6f, 0x5c, 0xcc,
		0x04, 0xaf, 0xe4, 0x29, 0x74, 0x8b, 0x23, 0xaa, 0x92, 0x7d, 0x22, 0x92,
		0x86, 0x99, 0x8a, 0xcb, 0xd8, 0x29, 0x7d, 0x0a, 0x27, 0xd8, 0x8c, 0xbe,
		0x6b, 0xa3, 0x2a, 0x32, 0xcd, 0x6e, 0x88, 0x32, 0xe5, 0x9c, 0x6a, 0x4e,
		0xe1, 0xa9, 0xbe, 0x19, 0xd1, 0x5d, 0xb5, 0x0d, 0x7d, 0x30, 0x4a, 0x82,
		0x9e, 0x20, 0x1d, 0x08, 0x09, 0x94, 0xc1, 0xec, 0x59, 0x3a, 0xe4, 0xac,
		0x99, 0xb7, 0xbe, 0x47, 0x45, 0xb2, 0x05, 0xe4, 0x82, 0xa2, 0x6f, 0x1e,
		0xd1, 0x63, 0xc6, 0x0c, 0x74, 0x3f, 0x71, 0x81, 0x25, 0x6b, 0x85, 0x8b,
		0xd6, 0x07, 0x54, 0xf0, 0x35, 0xa5, 0x0f, 0x74, 0xc6, 0x25, 0x65, 0x93,
		0x52, 0xb4, 0xbc, 0x58, 0xd3, 0x0c, 0x11, 0x6b, 0xbc, 0x93, 0xfa, 0x3f,
		0xa7, 0xaa, 0x4a, 0x20, 0x54, 0xe8, 0xc0, 0x43, 0xa7, 0xb1, 0x08, 0x1e,
		0x9e, 0xa1, 0xf3, 0x48, 0x35, 0x2a, 0x23, 0x9f, 0x83, 0x82, 0x5b, 0x2d,
		0xd8, 0x6a, 0xed, 0x78, 0xbb, 0xd2, 0x7a, 0x85, 0xbc, 0xf6, 0x68, 0x46,
		0xb2, 0xbc, 0x21, 0x2d, 0xd9, 0x4b, 0x19, 0x69, 0xa5, 0x29, 0x4f, 0x76,
		0x2f, 0xd1, 0xce, 0xb6, 0x4e, 0x85, 0x08, 0x0a, 0xe6, 0x58, 0xff, 0x32,
		0x8f, 0x1a, 0x55, 0x30, 0x31, 0xac, 0x11, 0x74, 0x3e, 0x71, 0x3f, 0xa2,
		0xf4, 0xed, 0xb8, 0xac, 0xec, 0x77, 0xe1, 0xe3, 0x58, 0x8e, 0xf7, 0x1e,
		0xcd, 0xe4, 0xc0, 0xb9, 0x12, 0x2b, 0x5d, 0x73, 0x02, 0x01, 0xd6, 0x4f,
		0xe4, 0x12, 0x55, 0x04, 0xcc, 0x70, 0x16, 0xd7, 0xbc, 0x28, 0x50, 0xce,
		0x23, 0x67, 0x5a, 0x0c, 0x5e, 0x49, 0x3b, 0xb7, 0x94, 0x4e, 0x3b, 0x45,
		0xb7, 0xd6, 0xd8, 0x8e, 0xd6, 0x99, 0x87, 0x70, 0xc8, 0x34, 0x8f, 0xa2,
		0xfd, 0x34, 0xc9, 0x46, 0x58, 0xa5, 0x04, 0xd6, 0x34, 0x74, 0xe4, 0xa8,
		0x82, 0x69, 0x4a, 0xbd, 0x3d, 0xdb, 0xe1, 0x3d, 0x02, 0x5e, 0x10, 0x60,
		0x36, 0xc6, 0x1b, 0xd6, 0x68, 0x81, 0xf1, 0xce, 0xf7, 0xf8, 0x68, 0x1b,
		0xe7, 0x56, 0x8c, 0xf4, 0x1b, 0x98, 0xd1, 0xcf, 0x2e, 0x4a, 0x82, 0x9f,
		0x91, 0x29, 0xdb, 0xf9, 0xb1, 0xe2, 0xae, 0x6e, 0xb3, 0x10, 0xc0, 0xef,
		0xda, 0x7f, 0xf0, 0x92, 0xb2, 0x23, 0x13, 0xa5, 0xc1, 0x22, 0xd6, 0x4b,
		0x2a, 0x78, 0x32, 0xcd, 0x84, 0xca, 0x28, 0x5d, 0x5b, 0x72, 0x89, 0x74,
		0x30, 0xf0, 0x0d, 0x92, 0x77, 0x31, 0xa7, 0x4c, 0xd2, 0xd1, 0xfa, 0xd4,
		0x55, 0x0a, 0xb5, 0xfc, 0xc6, 0xb0, 0xe5, 0xbc, 0x43, 0x76, 0x4b, 0x32,
		0xc0, 0x79, 0xa0, 0x83, 0x63, 0x78, 0x7d, 0x43, 0x45, 0xc8, 0x72, 0x25,
		0x3d, 0x48, 0xb3, 0x94, 0x54, 0xfa, 0x5d, 0x3a, 0xea, 0x5e, 0x93, 0x47,
		0x94, 0x1b, 0x63, 0xa6, 0xf9, 0xbe, 0xc0, 0x0f, 0x17, 0xf0, 0x4a, 0xe5,
		0xad, 0x4f, 0xcb, 0x54, 0xcd, 0xa7, 0xe5, 0xcd, 0xd2, 0x56, 0xdc, 0x09,
		0xc9, 0xe1, 0xd1, 0xf0, 0xaa, 0x76, 0x13, 0xb0, 0xee, 0x08, 0xbf, 0xd5,
		0xf3, 0x33, 0x27, 0x81, 0xfe, 0x86, 0x10, 0x1e, 0x58, 0xd3, 0xd2, 0x1f,
		0x0d, 0x83, 0x20, 0xfa, 0x0f, 0x86, 0x02, 0x5c, 0xf6, 0x02, 0x76, 0x4d,
		0x9a, 0x88, 0x0b, 0xbf, 0x7a, 0x1b, 0xa0, 0x5b, 0xae, 0x3f, 0x7a, 0x99,
		0xa5, 0x64, 0xf3, 0x78, 0x5b, 0xb7, 0x3a, 0x95, 0xc2, 0x26, 0x13, 0x9c,
		0x51, 0xcb, 0xe8, 0x60, 0x44, 0xc5, 0x37, 0x36, 0x3e, 0x3a, 0x06, 0xff,
		0x64, 0x9b, 0xf8, 0xf9, 0xf0, 0xa0, 0xca, 0x92, 0xa0, 0x8b, 0xbb, 0x0f,
		0x4d, 0x11, 0xff, 0xff, 0xf0, 0xd0, 0x7f, 0x38, 0x19, 0x5a, 0x80, 0x7d,
		0xa6, 0xeb, 0x5c, 0xbe, 0x7b, 0xf0, 0x7a, 0x2d, 0x36, 0x94, 0x91, 0xbd,
		0xd3, 0xb8, 0x80, 0x25, 0x17, 0x82, 0xf2, 0x6b, 0x88, 0x16, 0xea, 0xe0,
		0x40, 0x52, 0x59, 0xb7, 0x5d, 0xb2, 0x18, 0x3c, 0x18, 0x70, 0x81, 0x66,
		0x05, 0x27, 0x4f, 0xfd, 0x07, 0xaa, 0x98, 0x36, 0x81, 0x6f, 0x43, 0x87,
		0x26, 0x78, 0x7e, 0xdd, 0xf7, 0x00, 0x3d, 0xa5, 0x53, 0x70, 0x8d, 0xa8,
		0xc3, 0x22, 0xf1, 0x30, 0xbe, 0x9b, 0x53, 0x12, 0x13, 0xd8, 0x39, 0x13,
		0xbd, 0x7d, 0x24, 0x01, 0xf7, 0x03, 0xb6, 0xd1, 0x91, 0xa3, 0x80, 0xf0,
		0xff, 0x5e, 0x4d, 0x39, 0x40, 0x1d, 0x52, 0x3d, 0x79, 0xc8, 0x9e, 0xb9,
		0x3f, 0xab, 0xd6, 0xac, 0x95, 0xdd, 0x51, 0x69, 0x2b, 0x07, 0x4e, 0xb3,
		0xf5, 0x65, 0x6f, 0x8f, 0xe7, 0x98, 0x8c, 0x6a, 0x53, 0x13, 0x87, 0x42,
		0x35, 0xe5, 0xf7, 0x5c, 0xea, 0xd6, 0x6d, 0x91, 0x7a, 0x27, 0x32, 0x4a,
		0xf4, 0x45, 0xbe, 0x4b, 0xaa, 0x25, 0xa5, 0x04, 0x8b, 0xc5, 0x85, 0x27,
		0x8e, 0xfa, 0xf8, 0xf4, 0x0d, 0x41, 0x04, 0x0b, 0x26, 0x5a, 0x7a, 0x49,
		0x92, 0x64, 0x5f, 0x89, 0x09, 0xf5, 0xef, 0x62, 0x51, 0x49, 0x6d, 0xe6,
		0x84, 0x5b, 0x6c, 0x11, 0xf6, 0xfd, 0xc4, 0x94, 0x45, 0x13, 0x64, 0x31,
		0x65, 0x8a, 0x08, 0x42, 0xd3, 0xd0, 0x8f, 0x19, 0xa7, 0x70, 0xf4, 0xf4,
		0xe9, 0xff, 0xbd, 0xf0, 0xf1, 0xbb, 0xa7, 0xd1, 0xdd, 0x34, 0xff, 0xdf,
		0x39, 0xc8, 0x5f, 0xcf, 0x7f, 0xf8, 0x6f, 0xb9, 0xc5, 0x76, 0xde, 0xee,
		0x17, 0x7d, 0xed, 0xb0, 0x3e, 0x50, 0xab, 0xb3, 0x57, 0x2a, 0xf4, 0xd1,
		0x96, 0x6a, 0x14, 0x42, 0xb3, 0x5a, 0xbb, 0xe5, 0x37, 0xbe, 0xe5, 0x0a,
		0x14, 0xb3, 0xcc, 0xec, 0xb3, 0x20, 0x17, 0x86, 0x02, 0x73, 0x5e, 0x20,
		0x50, 0x63, 0x99, 0xd7, 0x3e, 0xbc, 0x7c, 0xd8, 0x39, 0x76, 0x8d, 0xa1,
		0x65, 0x5a, 0x62, 0x17, 0xd3, 0x52, 0xb9, 0x9e, 0x79, 0x98, 0xb2, 0xca,
		0x10, 0x93, 0x43, 0xa0, 0x17, 0x6b, 0x71, 0x14, 0xcc, 0x1f, 0x6b, 0x46,
		0xbd, 0x3d, 0x35, 0x50, 0x5e, 0x9f, 0x9c, 0x72, 0x73, 0x86, 0x34, 0x3f,
		0xd0, 0x46, 0x25, 0x45, 0xb7, 0x78, 0x2d, 0xd5, 0xb2, 0x17, 0xb7, 0x0e,
		0xf5, 0x40, 0x5c, 0x2b, 0x8b, 0xc9, 0xbe, 0x99, 0x7b, 0x45, 0x73, 0xd7,
		0xf8, 0xb7, 0x6a, 0x9d, 0x65, 0x18, 0xfd, 0x75, 0xe3, 0xdd, 0x5a, 0xa5,
		0xcf, 0x40, 0x70, 0x2e, 0xc4, 0x86, 0x36, 0xec, 0xcf, 0x6b, 0xcc, 0xaf,
		0xc9, 0x28, 0x56, 0x51, 0x4a, 0xa6, 0xcc, 0xc4, 0x88, 0x79, 0x98, 0x7d,
		0x4c, 0x4b, 0x43, 0x97, 0x87, 0x87, 0x2a, 0x7a, 0xdb, 0x09, 0x5a, 0x85,
		0x71, 0x26, 0xc0, 0xd3, 0xe7, 0x70, 0x62, 0xe6, 0xdb, 0x49, 0x69, 0x7b,
		0x13, 0x97, 0xaa, 0x15, 0x34, 0x23, 0x79, 0x40, 0xb9, 0x03, 0x64, 0x76,
		0x15, 0x9a, 0xce, 0xe0, 0xe1, 0xac, 0x09, 0xec, 0xaa, 0xd6, 0x77, 0xd8,
		0xe3, 0xd4, 0xf7, 0x7b, 0x50, 0x78, 0x49, 0x58, 0x5f, 0xd0, 0x78, 0x48,
		0xa3, 0xaf, 0x67, 0x65, 0x29, 0x3b, 0xe5, 0x64, 0x8c, 0x2a, 0xf0, 0x33,
		0x08, 0xbc, 0x2f, 0x89, 0xaa, 0x35, 0x84, 0x59, 0x7f, 0x64, 0x09, 0x5c,
		0xb8, 0x2f, 0x2c, 0xe8, 0x36, 0x13, 0x9c, 0x86, 0xcc, 0x82, 0x8e, 0x0e,
		0xee, 0xd3, 0xa4, 0x44, 0x67, 0x6f, 0xb8, 0x7b, 0xdb, 0x66, 0xbe, 0x05,
		0xf1, 0x4d, 0x37, 0x86, 0xa1, 0x7c, 0x49, 0x1b, 0x08, 0x4b, 0x87, 0x04,
		0xa9, 0x6d, 0xfd, 0x14, 0xb3, 0x85, 0xe4, 0xce, 0x21, 0x58, 0x56, 0xde,
		0xd1, 0x17, 0x76, 0x9b, 0x9a, 0x3b, 0x84, 0xfc, 0xe1, 0x42, 0xbd, 0xa1,
		0x1b, 0x47, 0x66, 0xdf, 0x99, 0xf8, 0x84, 0x3a, 0xdd, 0x8f, 0xec, 0x11,
		0xc7, 0x34, 0x12, 0x0b, 0x55, 0x1d, 0x28, 0xdb, 0x1d, 0x49, 0x7f, 0x41,
		0xb1, 0xdd, 0xdb, 0xee, 0x51, 0x4d, 0x0c, 0x1a, 0x9f, 0x69, 0xb8, 0x72,
		0xa1, 0xec, 0x30, 0x62, 0xd0, 0x34, 0xd3, 0xf0, 0x91, 0x09, 0x13, 0x8d,
		0xd0, 0xbf, 0xff, 0x75, 0xa0, 0xd7, 0xa9, 0x9f, 0x6d, 0x6b, 0x12, 0xee,
		0x42, 0x7c, 0xf3, 0xdf, 0x21, 0x40, 0x13, 0xdc, 0xb3, 0x09, 0x90, 0x6f,
		0xb3, 0x65, 0x22, 0xb7, 0xdd, 0xb1, 0xe0, 0xcd, 0x04, 0xcb, 0x28, 0x2f,
		0x13, 0x85, 0xcf, 0x93, 0x9d, 0xa3, 0xbc, 0x43, 0x59, 0xb9, 0xfa, 0xbd,
		0xf6, 0xdd, 0x6e, 0x34, 0xee, 0xa4, 0xa8, 0x08, 0xc6, 0x81, 0x3e, 0x3a,
		0x1b, 0x3a, 0x78, 0x10, 0x81, 0xf8, 0x94, 0x1c, 0xc8, 0x7f, 0xd8, 0x63,
		0xdf, 0x55, 0xd3, 0x0e, 0x4e, 0x49, 0xfd, 0x8f, 0xaf, 0x9f, 0x13, 0xb5,
		0x75, 0x5d, 0x36, 0x8f, 0x9f, 0x76, 0xee, 0x30, 0xa9, 0xcb, 0xbd, 0x4a,
		0xe7, 0x67, 0x8d, 0x3e, 0x58, 0xfb, 0x3b, 0x6d, 0x43, 0xb6, 0xca, 0xd4,
		0x4d, 0xa7, 0x0f, 0x15, 0xfb, 0x4b, 0x4d, 0x09, 0x9d, 0x89, 0x41, 0x17,
		0xf8, 0x89, 0x02, 0xdb, 0x76, 0x6b, 0xe1, 0x5e, 0x90, 0xe5, 0x94, 0x9e,
		0xf6, 0xdb, 0xdd, 0x09, 0x59, 0x87, 0xe2, 0xe8, 0xb6, 0x23, 0x9e, 0xec,
		0x01, 0xee, 0x33, 0x1f, 0x4c, 0x3b, 0xee, 0xd9, 0x4b, 0xef, 0xd4, 0x07,
		0x1c, 0xf5, 0x2e, 0xdc, 0xb5, 0xe1, 0x0d, 0x33, 0xab, 0x0e, 0x25, 0x15,
		0xb0, 0xb9, 0x64, 0x0b, 0xef, 0xce, 0xf4, 0xbf, 0xc7, 0x45, 0x56, 0x68,
		0x27, 0x87, 0xe1, 0x9d, 0x79, 0xf6, 0x2e, 0xa9, 0xc3, 0x5f, 0x3b, 0x5c,
		0xfd, 0xd0, 0xfa, 0x3e, 0xf7, 0x4b, 0x89, 0x39, 0x65, 0x77, 0x12, 0x1d,
		0xd2, 0xfe, 0xfa, 0xfa, 0x90, 0xf2, 0xeb, 0xb7, 0x6c, 0xc1, 0x2e, 0xbb,
		0x7b, 0x26, 0x2d, 0xda, 0x8a, 0x6a, 0xcd, 0x57, 0x9b, 0xeb, 0xae, 0xa9,
		0x0b, 0x28, 0x76, 0xc5, 0x6e, 0x92, 0x4a, 0x29, 0x9a, 0x96, 0x68, 0x04,
		0xb4, 0x21, 0xf7, 0xfa, 0x35, 0x4a, 0x8a, 0x99, 0x4d, 0xaf, 0xfc, 0x0d,
		0xdc, 0x2a, 0x3d, 0x4a, 0x8e, 0x8e, 0x92, 0xe3, 0xfe, 0xed, 0xe0, 0x75,
		0x14, 0x29, 0x78, 0x21, 0x73, 0xd1, 0x52, 0xe1, 0x67, 0x54, 0x00, 0x89,
		0x93, 0xe6, 0xc2, 0x17, 0xf2, 0x4e, 0x11, 0xf8, 0x92, 0x9c, 0x41, 0x2d,
		0xbf, 0x7a, 0x02, 0xa4, 0x33, 0xef, 0x09, 0xb9, 0x24, 0x13, 0x79, 0xd1,
		0x92, 0x1f, 0xf9, 0xdb, 0x39, 0xca, 0xd3, 0x54, 0xf8, 0x10, 0x0b, 0xda,
		0x76, 0xab, 0xda, 0x77, 0xbd, 0xfb, 0xbc, 0xda, 0xbd, 0xfa, 0x9c, 0x50,
		0x7c, 0xcc, 0x9f, 0xce, 0x17, 0x9d, 0x4d, 0x99, 0xd6, 0x34, 0x3a, 0x84,
		0x99, 0x77, 0x6f, 0x07, 0x1d, 0x68, 0xb8, 0x5c, 0x9a, 0xa5, 0xdd, 0xdd,
		0xf9, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa2, 0x3b, 0xe3, 0xe1, 0x4c,
		0x17, 0x00, 0x00,
	},
		"frontend/index.html",
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
	"frontend/assets/application.coffee": frontend_assets_application_coffee,
	"frontend/assets/application.js": frontend_assets_application_js,
	"frontend/index.html": frontend_index_html,
}
// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"frontend": &_bintree_t{nil, map[string]*_bintree_t{
		"assets": &_bintree_t{nil, map[string]*_bintree_t{
			"application.coffee": &_bintree_t{frontend_assets_application_coffee, map[string]*_bintree_t{
			}},
			"application.js": &_bintree_t{frontend_assets_application_js, map[string]*_bintree_t{
			}},
		}},
		"index.html": &_bintree_t{frontend_index_html, map[string]*_bintree_t{
		}},
	}},
}}
