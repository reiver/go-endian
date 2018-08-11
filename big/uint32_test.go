package bigendian_test

import (
	"github.com/reiver/go-endian/big"

	"bytes"
	"math/rand"
	"time"

	"testing"
)

func TestWriteUint32To(t *testing.T) {

	tests := []struct{
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
	}


	{
		randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))

		const max = 50
		for i:=0; i<max; i++ {
			x64 := randomness.Int63n(0xffffffff)

			u32 := uint32(x64)

			var b0 byte = uint8((0xff000000 & u32) >> 24)
			var b1 byte = uint8((0x00ff0000 & u32) >> 16)
			var b2 byte = uint8((0x0000ff00 & u32) >> 8)
			var b3 byte = uint8( 0x000000ff & u32      )

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

		n64, err := bigendian.WriteUint32To(&buffer, test.Value)
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
