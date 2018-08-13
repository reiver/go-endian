package nativeendian_test

import (
	"github.com/reiver/go-endian"
	"github.com/reiver/go-endian/native"

	"bytes"
	"math/rand"
	"time"

	"testing"
)

func TestWriteUint32To(t *testing.T) {

	var tests []struct{
		Value uint32
		Expected []byte
	}


	var endianness endian.Type = endian.NativeEndianness()

	switch endianness {
	default:
		panic(endianness)
	case endian.Little():
		tests = append(tests, []struct{
			Value uint32
			Expected []byte
		}{
			{
				Value:     0x00000000,
				Expected: []byte{0x00,0x00,0x00,0x00},
			},
			{
				Value:     0x00000001,
				Expected: []byte{0x01,0x00,0x00,0x00},
			},
			{
				Value:     0x00000002,
				Expected: []byte{0x02,0x00,0x00,0x00},
			},
			{
				Value:     0x00000003,
				Expected: []byte{0x03,0x00,0x00,0x00},
			},
			{
				Value:     0x00000004,
				Expected: []byte{0x04,0x00,0x00,0x00},
			},
			{
				Value:     0x00000005,
				Expected: []byte{0x05,0x00,0x00,0x00},
			},
			{
				Value:     0x00000006,
				Expected: []byte{0x06,0x00,0x00,0x00},
			},
			{
				Value:     0x00000007,
				Expected: []byte{0x07,0x00,0x00,0x00},
			},
			{
				Value:     0x00000008,
				Expected: []byte{0x08,0x00,0x00,0x00},
			},
			{
				Value:     0x00000009,
				Expected: []byte{0x09,0x00,0x00,0x00},
			},
			{
				Value:     0x0000000a,
				Expected: []byte{0x0a,0x00,0x00,0x00},
			},
			{
				Value:     0x0000000b,
				Expected: []byte{0x0b,0x00,0x00,0x00},
			},
			{
				Value:     0x0000000c,
				Expected: []byte{0x0c,0x00,0x00,0x00},
			},
			{
				Value:     0x0000000d,
				Expected: []byte{0x0d,0x00,0x00,0x00},
			},
			{
				Value:     0x0000000e,
				Expected: []byte{0x0e,0x00,0x00,0x00},
			},
			{
				Value:     0x0000000f,
				Expected: []byte{0x0f,0x00,0x00,0x00},
			},
			{
				Value:     0x00000010,
				Expected: []byte{0x10,0x00,0x00,0x00},
			},
			{
				Value:     0x00000011,
				Expected: []byte{0x11,0x00,0x00,0x00},
			},
			{
				Value:     0x00000012,
				Expected: []byte{0x12,0x00,0x00,0x00},
			},
			{
				Value:     0x00000013,
				Expected: []byte{0x13,0x00,0x00,0x00},
			},

			{
				Value:     0x000000ef,
				Expected: []byte{0xef,0x00,0x00,0x00},
			},
			{
				Value:     0x000000f0,
				Expected: []byte{0xf0,0x00,0x00,0x00},
			},
			{
				Value:     0x000000f1,
				Expected: []byte{0xf1,0x00,0x00,0x00},
			},
			{
				Value:     0x000000f2,
				Expected: []byte{0xf2,0x00,0x00,0x00},
			},
			{
				Value:     0x000000f3,
				Expected: []byte{0xf3,0x00,0x00,0x00},
			},
			{
				Value:     0x000000f4,
				Expected: []byte{0xf4,0x00,0x00,0x00},
			},
			{
				Value:     0x000000f5,
				Expected: []byte{0xf5,0x00,0x00,0x00},
			},
			{
				Value:     0x000000f6,
				Expected: []byte{0xf6,0x00,0x00,0x00},
			},
			{
				Value:     0x000000f7,
				Expected: []byte{0xf7,0x00,0x00,0x00},
			},
			{
				Value:     0x000000f8,
				Expected: []byte{0xf8,0x00,0x00,0x00},
			},
			{
				Value:     0x000000f9,
				Expected: []byte{0xf9,0x00,0x00,0x00},
			},
			{
				Value:     0x000000fa,
				Expected: []byte{0xfa,0x00,0x00,0x00},
			},
			{
				Value:     0x000000fb,
				Expected: []byte{0xfb,0x00,0x00,0x00},
			},
			{
				Value:     0x000000fc,
				Expected: []byte{0xfc,0x00,0x00,0x00},
			},
			{
				Value:     0x000000fd,
				Expected: []byte{0xfd,0x00,0x00,0x00},
			},
			{
				Value:     0x000000fe,
				Expected: []byte{0xfe,0x00,0x00,0x00},
			},
			{
				Value:     0x000000ff,
				Expected: []byte{0xff,0x00,0x00,0x00},
			},
			{
				Value:     0x00000100,
				Expected: []byte{0x00,0x01,0x00,0x00},
			},
			{
				Value:     0x00000101,
				Expected: []byte{0x01,0x01,0x00,0x00},
			},
			{
				Value:     0x00000102,
				Expected: []byte{0x02,0x01,0x00,0x00},
			},
			{
				Value:     0x00000103,
				Expected: []byte{0x03,0x01,0x00,0x00},
			},

			{
				Value:     0x00000245,
				Expected: []byte{0x45,0x02,0x00,0x00},
			},

			{
				Value:     0x000037ae,
				Expected: []byte{0xae,0x37,0x00,0x00},
			},

			{
				Value:     0x0000fffa,
				Expected: []byte{0xfa,0xff,0x00,0x00},
			},
			{
				Value:     0x0000fffb,
				Expected: []byte{0xfb,0xff,0x00,0x00},
			},
			{
				Value:     0x0000fffc,
				Expected: []byte{0xfc,0xff,0x00,0x00},
			},
			{
				Value:     0x0000fffd,
				Expected: []byte{0xfd,0xff,0x00,0x00},
			},
			{
				Value:     0x0000fffe,
				Expected: []byte{0xfe,0xff,0x00,0x00},
			},
			{
				Value:     0x0000ffff,
				Expected: []byte{0xff,0xff,0x00,0x00},
			},
			{
				Value:     0x00010000,
				Expected: []byte{0x00,0x00,0x01,0x00},
			},
			{
				Value:     0x00010001,
				Expected: []byte{0x01,0x00,0x01,0x00},
			},
			{
				Value:     0x00010002,
				Expected: []byte{0x02,0x00,0x01,0x00},
			},
			{
				Value:     0x00010003,
				Expected: []byte{0x03,0x00,0x01,0x00},
			},

			{
				Value:     0x1357bd24,
				Expected: []byte{0x24,0xbd,0x57,0x13},
			},

			{
				Value:     0xffffffff,
				Expected: []byte{0xff,0xff,0xff,0xff},
			},
		}...)
	case endian.Big():
		tests = append(tests, []struct{
			Value uint32
			Expected []byte
		}{
			{
				Value:                    0x00000000,
				Expected: []byte{0x00,0x00,0x00,0x00},
			},
			{
				Value:                    0x00000001,
				Expected: []byte{0x00,0x00,0x00,0x01},
			},
			{
				Value:                    0x00000002,
				Expected: []byte{0x00,0x00,0x00,0x02},
			},
			{
				Value:                    0x00000003,
				Expected: []byte{0x00,0x00,0x00,0x03},
			},
			{
				Value:                    0x00000004,
				Expected: []byte{0x00,0x00,0x00,0x04},
			},
			{
				Value:                    0x00000005,
				Expected: []byte{0x00,0x00,0x00,0x05},
			},
			{
				Value:                    0x00000006,
				Expected: []byte{0x00,0x00,0x00,0x06},
			},
			{
				Value:                    0x00000007,
				Expected: []byte{0x00,0x00,0x00,0x07},
			},
			{
				Value:                    0x00000008,
				Expected: []byte{0x00,0x00,0x00,0x08},
			},
			{
				Value:                    0x00000009,
				Expected: []byte{0x00,0x00,0x00,0x09},
			},
			{
				Value:                    0x0000000a,
				Expected: []byte{0x00,0x00,0x00,0x0a},
			},
			{
				Value:                    0x0000000b,
				Expected: []byte{0x00,0x00,0x00,0x0b},
			},
			{
				Value:                    0x0000000c,
				Expected: []byte{0x00,0x00,0x00,0x0c},
			},
			{
				Value:                    0x0000000d,
				Expected: []byte{0x00,0x00,0x00,0x0d},
			},
			{
				Value:                    0x0000000e,
				Expected: []byte{0x00,0x00,0x00,0x0e},
			},
			{
				Value:                    0x0000000f,
				Expected: []byte{0x00,0x00,0x00,0x0f},
			},
			{
				Value:                    0x00000010,
				Expected: []byte{0x00,0x00,0x00,0x10},
			},
			{
				Value:                    0x00000011,
				Expected: []byte{0x00,0x00,0x00,0x11},
			},
			{
				Value:                    0x00000012,
				Expected: []byte{0x00,0x00,0x00,0x12},
			},
			{
				Value:                    0x00000013,
				Expected: []byte{0x00,0x00,0x00,0x13},
			},

			{
				Value:                    0x000000ef,
				Expected: []byte{0x00,0x00,0x00,0xef},
			},
			{
				Value:                    0x000000f0,
				Expected: []byte{0x00,0x00,0x00,0xf0},
			},
			{
				Value:                    0x000000f1,
				Expected: []byte{0x00,0x00,0x00,0xf1},
			},
			{
				Value:                    0x000000f2,
				Expected: []byte{0x00,0x00,0x00,0xf2},
			},
			{
				Value:                    0x000000f3,
				Expected: []byte{0x00,0x00,0x00,0xf3},
			},
			{
				Value:                    0x000000f4,
				Expected: []byte{0x00,0x00,0x00,0xf4},
			},
			{
				Value:                    0x000000f5,
				Expected: []byte{0x00,0x00,0x00,0xf5},
			},
			{
				Value:                    0x000000f6,
				Expected: []byte{0x00,0x00,0x00,0xf6},
			},
			{
				Value:                    0x000000f7,
				Expected: []byte{0x00,0x00,0x00,0xf7},
			},
			{
				Value:                    0x000000f8,
				Expected: []byte{0x00,0x00,0x00,0xf8},
			},
			{
				Value:                    0x000000f9,
				Expected: []byte{0x00,0x00,0x00,0xf9},
			},
			{
				Value:                    0x000000fa,
				Expected: []byte{0x00,0x00,0x00,0xfa},
			},
			{
				Value:                    0x000000fb,
				Expected: []byte{0x00,0x00,0x00,0xfb},
			},
			{
				Value:                    0x000000fc,
				Expected: []byte{0x00,0x00,0x00,0xfc},
			},
			{
				Value:                    0x000000fd,
				Expected: []byte{0x00,0x00,0x00,0xfd},
			},
			{
				Value:                    0x000000fe,
				Expected: []byte{0x00,0x00,0x00,0xfe},
			},
			{
				Value:                    0x000000ff,
				Expected: []byte{0x00,0x00,0x00,0xff},
			},
			{
				Value:                    0x00000100,
				Expected: []byte{0x00,0x00,0x01,0x00},
			},
			{
				Value:                    0x00000101,
				Expected: []byte{0x00,0x00,0x01,0x01},
			},
			{
				Value:                    0x00000102,
				Expected: []byte{0x00,0x00,0x01,0x02},
			},
			{
				Value:                    0x00000103,
				Expected: []byte{0x00,0x00,0x01,0x03},
			},

			{
				Value:                    0x00000245,
				Expected: []byte{0x00,0x00,0x02,0x45},
			},

			{
				Value:                    0x000037ae,
				Expected: []byte{0x00,0x00,0x37,0xae},
			},

			{
				Value:                    0x0000fffa,
				Expected: []byte{0x00,0x00,0x0ff,0xfa},
			},
			{
				Value:                    0x0000fffb,
				Expected: []byte{0x00,0x00,0xff,0xfb},
			},
			{
				Value:                    0x0000fffc,
				Expected: []byte{0x00,0x00,0xff,0xfc},
			},
			{
				Value:                    0x0000fffd,
				Expected: []byte{0x00,0x00,0xff,0xfd},
			},
			{
				Value:                    0x0000fffe,
				Expected: []byte{0x00,0x00,0xff,0xfe},
			},
			{
				Value:                    0x0000ffff,
				Expected: []byte{0x00,0x00,0xff,0xff},
			},
			{
				Value:                    0x00010000,
				Expected: []byte{0x00,0x01,0x00,0x00},
			},
			{
				Value:                    0x00010001,
				Expected: []byte{0x00,0x01,0x00,0x01},
			},
			{
				Value:                    0x00010002,
				Expected: []byte{0x00,0x01,0x00,0x02},
			},
			{
				Value:                    0x00010003,
				Expected: []byte{0x00,0x01,0x00,0x03},
			},

			{
				Value:                    0x1357bd24,
				Expected: []byte{0x13,0x57,0xbd,0x24},
			},

			{
				Value:                    0xffffffff,
				Expected: []byte{0xff,0xff,0xff,0xff},
			},
		}...)
	}


	{
		randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))

		const max = 50
		for i:=0; i<max; i++ {
			x64 := randomness.Int63n(0xffffffff)

			u32 := uint32(x64)

			var b0 byte
			var b1 byte
			var b2 byte
			var b3 byte

			switch endianness {
			default:
                                panic(endianness)
			case endian.Little():
				b0 = uint8( 0x000000ff & u32      )
				b1 = uint8((0x0000ff00 & u32) >> 8)
				b2 = uint8((0x00ff0000 & u32) >> 16)
				b3 = uint8((0xff000000 & u32) >> 24)
			case endian.Big():
				b0 = uint8((0xff000000 & u32) >> 24)
				b1 = uint8((0x00ff0000 & u32) >> 16)
				b2 = uint8((0x0000ff00 & u32) >> 8)
				b3 = uint8( 0x000000ff & u32      )
                        }

			test := struct{
				Value uint32
				Expected []byte
			}{
				Value:           u32,
				Expected: []byte{b0, b1, b2, b3},
			}

			tests = append(tests, test)
		}
	}


	for testNumber, test := range tests {

		var buffer bytes.Buffer

		n64, err := nativeendian.WriteUint32To(&buffer, test.Value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually go one: (%T) %q", testNumber, err, err)
			continue
		}

		if expected, actual := int64(len(test.Expected)), n64; expected != actual {
			t.Errorf("For test #%d, expected %d, but actually got %d.", testNumber, expected, actual)
			continue
		}

		if expected, actual := test.Expected, buffer.Bytes(); !bytes.Equal(expected, actual) {
			t.Errorf("For test #%d,....", testNumber)
			t.Errorf("\tEXPECTED: % v", expected)
			t.Errorf("\tACTUAL:   % v", actual)
			continue
		}
	}
}

func TestReadUint32From(t *testing.T) {

	var tests []struct{
		Value []byte
		Expected uint32
	}

	var endianness endian.Type = endian.NativeEndianness()

	switch endianness {
	default:
		panic(endianness)
	case endian.Little():
		tests = append(tests, []struct{
			Value []byte
			Expected uint32
		}{
			{
				Value:   []byte{0x00,0x00,0x00,0x00},
				Expected: 0x00000000,
			},
			{
				Value:   []byte{0x01,0x00,0x00,0x00},
				Expected: 0x00000001,
			},
			{
				Value:   []byte{0x02,0x00,0x00,0x00},
				Expected: 0x00000002,
			},
			{
				Value:   []byte{0x03,0x00,0x00,0x00},
				Expected: 0x00000003,
			},
			{
				Value:   []byte{0x04,0x00,0x00,0x00},
				Expected: 0x00000004,
			},
			{
				Value:   []byte{0x05,0x00,0x00,0x00},
				Expected: 0x00000005,
			},
			{
				Value:   []byte{0x06,0x00,0x00,0x00},
				Expected: 0x00000006,
			},
			{
				Value:   []byte{0x07,0x00,0x00,0x00},
				Expected: 0x00000007,
			},
			{
				Value:   []byte{0x08,0x00,0x00,0x00},
				Expected: 0x00000008,
			},
			{
				Value:   []byte{0x09,0x00,0x00,0x00},
				Expected: 0x00000009,
			},
			{
				Value:   []byte{0x0a,0x00,0x00,0x00},
				Expected: 0x0000000a,
			},
			{
				Value:   []byte{0x0b,0x00,0x00,0x00},
				Expected: 0x0000000b,
			},
			{
				Value:   []byte{0x0c,0x00,0x00,0x00},
				Expected: 0x0000000c,
			},
			{
				Value:   []byte{0x0d,0x00,0x00,0x00},
				Expected: 0x0000000d,
			},
			{
				Value:   []byte{0x0e,0x00,0x00,0x00},
				Expected: 0x0000000e,
			},
			{
				Value:   []byte{0x0f,0x00,0x00,0x00},
				Expected: 0x0000000f,
			},
				{
				Value:   []byte{0x10,0x00,0x00,0x00},
				Expected: 0x00000010,
			},
			{
				Value:   []byte{0x11,0x00,0x00,0x00},
				Expected: 0x00000011,
			},
			{
				Value:   []byte{0x12,0x00,0x00,0x00},
				Expected: 0x00000012,
			},
			{
				Value:   []byte{0x13,0x00,0x00,0x00},
				Expected: 0x00000013,
			},

			{
				Value:   []byte{0xef,0x00,0x00,0x00},
				Expected: 0x000000ef,
			},
			{
				Value:   []byte{0xf0,0x00,0x00,0x00},
				Expected: 0x000000f0,
			},
			{
				Value:   []byte{0xf1,0x00,0x00,0x00},
				Expected: 0x000000f1,
			},
			{
				Value:   []byte{0xf2,0x00,0x00,0x00},
				Expected: 0x000000f2,
			},
			{
				Value:   []byte{0xf3,0x00,0x00,0x00},
				Expected: 0x000000f3,
			},
			{
				Value:   []byte{0xf4,0x00,0x00,0x00},
				Expected: 0x000000f4,
			},
			{
				Value:   []byte{0xf5,0x00,0x00,0x00},
				Expected: 0x000000f5,
			},
			{
				Value:   []byte{0xf6,0x00,0x00,0x00},
				Expected: 0x000000f6,
			},
			{
				Value:   []byte{0xf7,0x00,0x00,0x00},
				Expected: 0x000000f7,
			},
			{
				Value:   []byte{0xf8,0x00,0x00,0x00},
				Expected: 0x000000f8,
			},
			{
				Value:   []byte{0xf9,0x00,0x00,0x00},
				Expected: 0x000000f9,
			},
			{
				Value:   []byte{0xfa,0x00,0x00,0x00},
				Expected: 0x000000fa,
			},
			{
				Value:   []byte{0xfb,0x00,0x00,0x00},
				Expected: 0x000000fb,
			},
			{
				Value:   []byte{0xfc,0x00,0x00,0x00},
				Expected: 0x000000fc,
			},
			{
				Value:   []byte{0xfd,0x00,0x00,0x00},
				Expected: 0x000000fd,
			},
			{
				Value:   []byte{0xfe,0x00,0x00,0x00},
				Expected: 0x000000fe,
			},
			{
				Value:   []byte{0xff,0x00,0x00,0x00},
				Expected: 0x000000ff,
			},
			{
				Value:   []byte{0x00,0x01,0x00,0x00},
				Expected: 0x00000100,
			},
			{
				Value:   []byte{0x01,0x01,0x00,0x00},
				Expected: 0x00000101,
			},
			{
				Value:   []byte{0x02,0x01,0x00,0x00},
				Expected: 0x00000102,
			},
			{
				Value:   []byte{0x03,0x01,0x00,0x00},
				Expected: 0x00000103,
			},

			{
				Value:   []byte{0x45,0x02,0x00,0x00},
				Expected: 0x00000245,
			},

			{
				Value:   []byte{0xae,0x37,0x00,0x00},
				Expected: 0x000037ae,
			},

			{
				Value:   []byte{0xfa,0xff,0x00,0x00},
				Expected: 0x0000fffa,
			},
			{
				Value:   []byte{0xfb,0xff,0x00,0x00},
				Expected: 0x0000fffb,
			},
			{
				Value:   []byte{0xfc,0xff,0x00,0x00},
				Expected: 0x0000fffc,
			},
			{
				Value:   []byte{0xfd,0xff,0x00,0x00},
				Expected: 0x0000fffd,
			},
			{
				Value:   []byte{0xfe,0xff,0x00,0x00},
				Expected: 0x0000fffe,
			},
			{
				Value:   []byte{0xff,0xff,0x00,0x00},
				Expected: 0x0000ffff,
			},
			{
				Value:   []byte{0x00,0x00,0x01,0x00},
				Expected: 0x00010000,
			},
			{
				Value:   []byte{0x01,0x00,0x01,0x00},
				Expected: 0x00010001,
			},
			{
				Value:   []byte{0x02,0x00,0x01,0x00},
				Expected: 0x00010002,
			},
			{
				Value:   []byte{0x03,0x00,0x01,0x00},
				Expected: 0x00010003,
			},

			{
				Value:   []byte{0x24,0xbd,0x57,0x13},
				Expected: 0x1357bd24,
			},

			{
				Value:   []byte{0xff,0xff,0xff,0xff},
				Expected: 0xffffffff,
			},
		}...)
	case endian.Big():
		tests = append(tests, []struct{
			Value []byte
			Expected uint32
		}{
			{
				Value: []byte{0x00,0x00,0x00,0x00},
				Expected:              0x00000000,
			},
			{
				Value: []byte{0x00,0x00,0x00,0x01},
				Expected:              0x00000001,
			},
			{
				Value: []byte{0x00,0x00,0x00,0x02},
				Expected:              0x00000002,
			},
			{
				Value: []byte{0x00,0x00,0x00,0x03},
				Expected:              0x00000003,
			},
			{
				Value: []byte{0x00,0x00,0x00,0x04},
				Expected:              0x00000004,
			},
			{
				Value: []byte{0x00,0x00,0x00,0x05},
				Expected:              0x00000005,
			},
			{
				Value: []byte{0x00,0x00,0x00,0x06},
				Expected:              0x00000006,
			},
			{
				Value: []byte{0x00,0x00,0x00,0x07},
				Expected:              0x00000007,
			},
			{
				Value: []byte{0x00,0x00,0x00,0x08},
				Expected:              0x00000008,
			},
			{
				Value: []byte{0x00,0x00,0x00,0x09},
				Expected:              0x00000009,
			},
			{
				Value: []byte{0x00,0x00,0x00,0x0a},
				Expected:              0x0000000a,
			},
			{
				Value: []byte{0x00,0x00,0x00,0x0b},
				Expected:              0x0000000b,
			},
			{
				Value: []byte{0x00,0x00,0x00,0x0c},
				Expected:              0x0000000c,
			},
				{
				Value: []byte{0x00,0x00,0x00,0x0d},
				Expected:              0x0000000d,
			},
			{
				Value: []byte{0x00,0x00,0x00,0x0e},
				Expected:              0x0000000e,
			},
			{
				Value: []byte{0x00,0x00,0x00,0x0f},
				Expected:              0x0000000f,
			},
			{
				Value: []byte{0x00,0x00,0x00,0x10},
				Expected:              0x00000010,
			},
			{
				Value: []byte{0x00,0x00,0x00,0x11},
				Expected:              0x00000011,
			},
			{
				Value: []byte{0x00,0x00,0x00,0x12},
				Expected:              0x00000012,
			},
			{
				Value: []byte{0x00,0x00,0x00,0x13},
				Expected:              0x00000013,
			},

			{
				Value: []byte{0x00,0x00,0x00,0xef},
				Expected:              0x000000ef,
			},
			{
				Value: []byte{0x00,0x00,0x00,0xf0},
				Expected:              0x000000f0,
			},
			{
				Value: []byte{0x00,0x00,0x00,0xf1},
				Expected:              0x000000f1,
			},
			{
				Value: []byte{0x00,0x00,0x00,0xf2},
				Expected:              0x000000f2,
			},
			{
				Value: []byte{0x00,0x00,0x00,0xf3},
				Expected:              0x000000f3,
			},
			{
				Value: []byte{0x00,0x00,0x00,0xf4},
				Expected:              0x000000f4,
			},
			{
				Value: []byte{0x00,0x00,0x00,0xf5},
				Expected:              0x000000f5,
			},
			{
				Value: []byte{0x00,0x00,0x00,0xf6},
				Expected:              0x000000f6,
			},
			{
				Value: []byte{0x00,0x00,0x00,0xf7},
				Expected:              0x000000f7,
			},
			{
				Value: []byte{0x00,0x00,0x00,0xf8},
				Expected:              0x000000f8,
			},
			{
				Value: []byte{0x00,0x00,0x00,0xf9},
				Expected:              0x000000f9,
			},
			{
				Value: []byte{0x00,0x00,0x00,0xfa},
				Expected:              0x000000fa,
			},
			{
				Value: []byte{0x00,0x00,0x00,0xfb},
				Expected:              0x000000fb,
			},
			{
				Value: []byte{0x00,0x00,0x00,0xfc},
				Expected:              0x000000fc,
			},
			{
				Value: []byte{0x00,0x00,0x00,0xfd},
				Expected:              0x000000fd,
			},
			{
				Value: []byte{0x00,0x00,0x00,0xfe},
				Expected:              0x000000fe,
			},
			{
				Value: []byte{0x00,0x00,0x00,0xff},
				Expected:              0x000000ff,
			},
			{
				Value: []byte{0x00,0x00,0x01,0x00},
				Expected:              0x00000100,
			},
			{
				Value: []byte{0x00,0x00,0x01,0x01},
				Expected:              0x00000101,
			},
			{
				Value: []byte{0x00,0x00,0x01,0x02},
				Expected:              0x00000102,
			},
			{
				Value: []byte{0x00,0x00,0x01,0x03},
				Expected:              0x00000103,
			},

			{
				Value: []byte{0x00,0x00,0x02,0x45},
				Expected:              0x00000245,
			},

			{
				Value: []byte{0x00,0x00,0x37,0xae},
				Expected:              0x000037ae,
			},

			{
				Value: []byte{0x00,0x00,0xff,0xfa},
				Expected:              0x0000fffa,
			},
			{
				Value: []byte{0x00,0x00,0xff,0xfb},
				Expected:              0x0000fffb,
			},
			{
				Value: []byte{0x00,0x00,0xff,0xfc},
				Expected:              0x0000fffc,
			},
			{
				Value: []byte{0x00,0x00,0xff,0xfd},
				Expected:              0x0000fffd,
			},
			{
				Value: []byte{0x00,0x00,0xff,0xfe},
				Expected:              0x0000fffe,
			},
			{
				Value: []byte{0x00,0x00,0xff,0xff},
				Expected:              0x0000ffff,
			},
			{
				Value: []byte{0x00,0x01,0x00,0x00},
				Expected:              0x00010000,
			},
			{
				Value: []byte{0x00,0x01,0x00,0x01},
				Expected:              0x00010001,
			},
			{
				Value: []byte{0x00,0x01,0x00,0x02},
				Expected:              0x00010002,
			},
			{
				Value: []byte{0x00,0x01,0x00,0x03},
				Expected:              0x00010003,
			},

			{
				Value: []byte{0x13,0x57,0xbd,0x24},
				Expected:              0x1357bd24,
			},

			{
				Value: []byte{0xff,0xff,0xff,0xff},
				Expected:              0xffffffff,
			},
		}...)
	}


	{
		randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))

		const max = 50
		for i:=0; i<max; i++ {
			x64 := randomness.Int63n(0xffffffff)

			u32 := uint32(x64)

			var b0 byte
			var b1 byte
			var b2 byte
			var b3 byte

			switch endianness {
			default:
                                panic(endianness)
			case endian.Little():
				b0 = uint8( 0x000000ff & u32       )
				b1 = uint8((0x0000ff00 & u32) >>  8)
				b2 = uint8((0x00ff0000 & u32) >> 16)
				b3 = uint8((0xff000000 & u32) >> 24)
			case endian.Big():
				b0 = uint8((0xff000000 & u32) >> 24)
				b1 = uint8((0x00ff0000 & u32) >> 16)
				b2 = uint8((0x0000ff00 & u32) >>  8)
				b3 = uint8( 0x000000ff & u32       )
                        }

			test := struct{
				Value []byte
				Expected uint32
			}{
				Value: []byte{b0,b1,b2,b3},
				Expected:     u32,
			}

			tests = append(tests, test)
		}
	}


	for testNumber, test := range tests {

		var buffer bytes.Buffer

		buffer.Write(test.Value)

		var actual uint32

		n64, err := nativeendian.ReadUint32From(&buffer, &actual)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually go one: (%T) %q", testNumber, err, err)
			continue
		}

		if expected, actual := int64(len(test.Value)), n64; expected != actual {
			t.Errorf("For test #%d, expected %d, but actually got %d.", testNumber, expected, actual)
			continue
		}

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d, expected 0x%X, but actually got 0x%X.", testNumber, expected, actual)
			continue
		}
	}
}
