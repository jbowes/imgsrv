package assets

import (
	"bytes"
	"compress/gzip"
	"io"
)

// TagCloudJs returns raw, uncompressed file data.
func TagCloudJs() []byte {
	gz, err := gzip.NewReader(bytes.NewBuffer([]byte{
0x1f,0x8b,0x08,0x00,0x00,0x09,0x6e,0x88,0x00,0xff,0x94,0x55,
0x4d,0x6f,0xe3,0x36,0x10,0xbd,0xfb,0x57,0x4c,0x55,0x1f,0xa4,
0xac,0x22,0xc5,0xd9,0xdd,0xa2,0xb0,0xea,0x16,0x69,0xb0,0xd8,
0xa6,0xd8,0x16,0x6d,0x53,0xa0,0x87,0x6c,0x0e,0xb4,0x4c,0xcb,
0x4a,0x65,0xd2,0x25,0xa9,0x24,0xae,0xa1,0xff,0xde,0x19,0x4a,
0xa2,0x28,0x35,0x39,0x2c,0xe0,0x48,0xf1,0xf0,0xbd,0x99,0xe1,
0x9b,0x0f,0xa7,0x67,0x5f,0xcd,0xe0,0x0c,0x1e,0xfe,0xa9,0xb9,
0x3a,0x26,0x86,0x15,0x79,0x25,0xeb,0x4d,0xf2,0xa0,0xc9,0x7a,
0x05,0xb7,0xe5,0xfe,0x50,0x71,0xf8,0x93,0x15,0x70,0x4d,0x07,
0xf0,0x5b,0x55,0x17,0xa5,0x80,0xad,0x54,0xf0,0xf3,0xef,0xc4,
0x41,0x1c,0x41,0x77,0xc6,0x1c,0xf4,0x32,0x4d,0x8b,0xd2,0xec,
0xea,0x75,0x92,0xcb,0x7d,0xca,0x36,0x9b,0xe3,0x13,0x3d,0xd2,
0x97,0xbd,0xe7,0x8a,0x33,0xc3,0x37,0xb0,0x3e,0xc2,0xd5,0x86,
0xed,0xe1,0xa3,0x92,0x8f,0x5c,0x77,0xfe,0x3e,0x95,0x39,0x17,
0x9a,0x2f,0xbf,0xd4,0x71,0xba,0xae,0xe4,0x3a,0xdd,0x33,0x6d,
0xb8,0x4a,0x3f,0xdd,0x5c,0x7f,0xf8,0xf5,0xf6,0x03,0xfa,0x4b,
0x67,0xe1,0xb6,0x16,0xb9,0x29,0xa5,0x08,0xe7,0x11,0x9c,0x66,
0x33,0x80,0xf4,0xac,0x40,0x2c,0xab,0xe0,0xc1,0x5e,0x04,0x31,
0x00,0x41,0xad,0x39,0x68,0xa3,0xca,0xdc,0x04,0x19,0x81,0x1e,
0x99,0x02,0x8c,0x79,0x60,0x8a,0xff,0xc5,0xcb,0x62,0x67,0x34,
0xac,0xc0,0xb9,0x62,0x31,0xac,0x23,0x44,0x9d,0xf0,0x0f,0x40,
0x71,0x53,0x2b,0x01,0x0c,0xce,0x61,0x9d,0xa1,0xa5,0xb1,0x1e,
0xd2,0x14,0xae,0xa5,0x78,0xe4,0x0a,0xa9,0x3b,0xfe,0x0c,0x46,
0x02,0x13,0xf0,0xc7,0xc7,0x1f,0x81,0x29,0xc5,0x8e,0x5d,0x0c,
0x23,0xc9,0xe2,0xb9,0xce,0xe5,0x86,0x47,0x9d,0xe3,0x72,0x0b,
0xf6,0x7b,0x52,0x71,0x51,0x98,0x1d,0xac,0x56,0x2b,0x78,0xd7,
0x1f,0x02,0xd0,0x11,0x52,0x2d,0x42,0xf1,0x43,0xc5,0x72,0x1e,
0xa6,0xe1,0xe7,0xa7,0xa8,0xff,0x43,0xf1,0x62,0x08,0x3e,0xcf,
0x17,0xf6,0x73,0x69,0x3f,0x6f,0xf1,0x13,0x44,0x99,0x75,0xd1,
0xd8,0x27,0xa5,0x41,0x09,0xae,0x80,0xc8,0xa7,0xcb,0x26,0x1a,
0xbf,0xd2,0x84,0x3f,0xf3,0xbc,0x4d,0x2c,0xf3,0x2f,0x7c,0x87,
0xea,0x68,0x7e,0x23,0x4c,0x88,0xf4,0xbb,0xc5,0x7d,0x0c,0x8b,
0x6f,0xa2,0x18,0x46,0xd6,0xcb,0x17,0xad,0x6f,0x5b,0xeb,0xfd,
0xcb,0x6a,0xf9,0x32,0x91,0x6c,0xc8,0x70,0x6a,0xfd,0x64,0x13,
0x1d,0x0a,0xa1,0x8e,0xd1,0xb8,0x0a,0xc1,0xd7,0x01,0xbc,0xe9,
0x4a,0x9b,0xec,0xd9,0x81,0x20,0xf1,0x40,0x28,0x07,0xf9,0x86,
0x7b,0x43,0x99,0x18,0x79,0x8b,0xd5,0x17,0x45,0x88,0x69,0x65,
0x1d,0xa0,0x3d,0xa4,0x84,0xfd,0x02,0x2c,0x22,0xf8,0x01,0x82,
0x0b,0x8a,0x42,0x80,0x25,0x3d,0x7b,0x46,0x97,0x83,0xb3,0x34,
0x51,0xf2,0x20,0x4b,0x11,0x06,0xad,0xe2,0x8d,0xd7,0x5a,0x95,
0x54,0x37,0x02,0x47,0x61,0xcf,0x85,0x19,0xd7,0x1f,0x4f,0x62,
0x50,0x4c,0x14,0x7c,0x72,0x35,0xef,0x52,0xb6,0x6d,0x5a,0x6c,
0xc2,0xc5,0x26,0xf2,0x6e,0x28,0x62,0xf0,0x2e,0xd9,0x51,0x43,
0x81,0xcd,0xe9,0x93,0xb4,0x61,0xca,0x44,0x77,0xe5,0x7d,0x94,
0xda,0x50,0x7d,0xbe,0xe3,0x34,0x71,0xc0,0xae,0x09,0xfe,0x42,
0x82,0x65,0x9f,0x7c,0x0c,0x4f,0x76,0x40,0x50,0xbc,0x3e,0x2c,
0x51,0x55,0xb1,0x46,0xd6,0x2b,0x29,0xb7,0xd1,0x5f,0x4d,0xda,
0xf2,0xf9,0x16,0xf9,0xbf,0x30,0xb3,0x4b,0x94,0xac,0xc5,0x06,
0x6f,0xf0,0x06,0x42,0x17,0x15,0x33,0xc7,0x6d,0x31,0x44,0x76,
0x45,0xa3,0x99,0x21,0xee,0xf7,0x70,0xf9,0xfe,0xfd,0xe0,0x12,
0x3a,0x87,0x68,0xec,0x91,0x0d,0xf0,0x0a,0xa7,0x7e,0x40,0xf4,
0xd4,0xef,0xe0,0xc2,0x27,0xf6,0xd4,0x8b,0xcc,0x99,0x9a,0xd9,
0xf8,0xdd,0xa9,0x8c,0x38,0x5f,0x47,0x67,0xb7,0x6d,0x1b,0xa2,
0x22,0x9e,0xbc,0xf3,0x64,0x2b,0xdc,0x02,0xf3,0xf5,0x95,0x07,
0x7a,0xe9,0x6e,0x59,0xb5,0x62,0xa0,0x8d,0xf6,0xcf,0x1c,0x27,
0xd1,0x60,0xb9,0xc3,0x53,0x13,0x8f,0x1d,0x24,0x1b,0xbe,0x65,
0x75,0x65,0x74,0x0c,0x3d,0x3f,0x73,0x6c,0x04,0x0d,0x3b,0xcc,
0xec,0x4a,0x6d,0xeb,0xe1,0x02,0x46,0x93,0x5e,0x99,0x87,0x84,
0x89,0x12,0x66,0x8c,0x0a,0x03,0xc5,0x2b,0xb7,0x2e,0xba,0xf7,
0xc8,0x9f,0xab,0xf0,0xdf,0xfc,0x8a,0x06,0x36,0x1c,0x4e,0xa3,
0x44,0x4b,0x65,0xc2,0xf1,0x1a,0xf5,0xf2,0xaa,0xe4,0x13,0xd7,
0xd4,0xfc,0x03,0xe5,0xee,0xe2,0x7e,0x38,0xdf,0xa1,0x69,0x0a,
0x48,0x0e,0xf2,0x10,0x7a,0x3e,0x6c,0xf3,0x22,0xa2,0xc7,0x9e,
0x77,0x5e,0xb3,0x6e,0x81,0x86,0x1d,0x00,0x07,0x97,0x8a,0xda,
0xc3,0x17,0x59,0x5b,0x3a,0x5c,0x3b,0xb7,0xe5,0xbf,0xf4,0xeb,
0xd3,0xf9,0xdb,0x4a,0x61,0x68,0x2e,0xe3,0x61,0x44,0x33,0xb7,
0x8b,0xa9,0x0c,0x89,0x46,0xfc,0xd0,0x1e,0x3d,0x9e,0x16,0x85,
0x3b,0xa6,0x91,0xc4,0x4c,0x86,0xef,0x6d,0xbf,0x8f,0x26,0xad,
0x0f,0x6f,0xe7,0x4b,0x8f,0x43,0xd8,0xd0,0xfe,0x9e,0xef,0x32,
0xb1,0xcb,0x7e,0xb4,0x38,0x3c,0x7c,0xbf,0x33,0x7c,0xff,0x7d,
0xff,0x51,0xcd,0x39,0xcb,0x77,0x5e,0xd1,0x47,0xc3,0xe6,0xe6,
0x88,0x9a,0xec,0xff,0xe5,0x9f,0xa8,0xfa,0xaa,0x18,0xe0,0xc8,
0xb9,0xd6,0xe1,0x29,0x20,0x71,0xce,0x09,0x12,0x2c,0xa7,0x62,
0xd0,0x30,0x0f,0x51,0xcf,0x9c,0x8e,0x11,0x1e,0x0c,0xd0,0x5a,
0x94,0xa6,0x71,0xa3,0xdd,0x4c,0xa3,0x4f,0x74,0x9a,0x86,0xb7,
0xc7,0x18,0xba,0xdf,0x62,0x23,0xb1,0x9c,0x90,0xfe,0xfe,0x9a,
0xc6,0x6a,0x5e,0x1b,0x58,0x37,0x6f,0x28,0x58,0x1b,0x9e,0xf2,
0x5d,0xc2,0xc9,0x5e,0x6e,0x09,0x8b,0x77,0x31,0x60,0x13,0xe0,
0x3f,0xdf,0xc6,0x40,0xb7,0x58,0x42,0x70,0x30,0x41,0xd3,0x79,
0xc3,0xdf,0xd6,0x76,0x6e,0xd0,0xff,0x7f,0x01,0x00,0x00,0xff,
0xff,0x07,0x83,0x9c,0xd6,0x8a,0x09,0x00,0x00,
	}))

	if err != nil {
		panic("Decompression failed: " + err.Error())
	}

	var b bytes.Buffer
	io.Copy(&b, gz)
	gz.Close()

	return b.Bytes()
}