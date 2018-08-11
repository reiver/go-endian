package bigendian_test

import (
	"github.com/reiver/go-endian/big"

	"bytes"
	"math/rand"
	"time"

	"testing"
)

func TestWriteUint16To(t *testing.T) {

	tests := []struct{
		Value uint16
		Expected []byte
	}{
		{
			Value:              0x0000,
			Expected: []byte{0x00,0x00},

		},
		{
			Value:              0x0001,
			Expected: []byte{0x00,0x01},

		},
		{
			Value:              0x0002,
			Expected: []byte{0x00,0x02},

		},
		{
			Value:              0x0003,
			Expected: []byte{0x00,0x03},

		},
		{
			Value:              0x0004,
			Expected: []byte{0x00,0x04},

		},
		{
			Value:              0x0005,
			Expected: []byte{0x00,0x05},

		},
		{
			Value:              0x0006,
			Expected: []byte{0x00,0x06},

		},
		{
			Value:              0x0007,
			Expected: []byte{0x00,0x07},

		},
		{
			Value:              0x0008,
			Expected: []byte{0x00,0x08},

		},
		{
			Value:              0x0009,
			Expected: []byte{0x00,0x09},

		},
		{
			Value:              0x000a,
			Expected: []byte{0x00,0x0a},

		},
		{
			Value:              0x000b,
			Expected: []byte{0x00,0x0b},

		},
		{
			Value:              0x000c,
			Expected: []byte{0x00,0x0c},

		},
		{
			Value:              0x000d,
			Expected: []byte{0x00,0x0d},

		},
		{
			Value:              0x000e,
			Expected: []byte{0x00,0x0e},

		},
		{
			Value:              0x000f,
			Expected: []byte{0x00,0x0f},

		},
		{
			Value:              0x0010,
			Expected: []byte{0x00,0x10},

		},
		{
			Value:              0x0011,
			Expected: []byte{0x00,0x11},

		},
		{
			Value:              0x0012,
			Expected: []byte{0x00,0x12},

		},
		{
			Value:              0x0013,
			Expected: []byte{0x00,0x13},

		},

		{
			Value:              0x00ef,
			Expected: []byte{0x00,0xef},

		},
		{
			Value:              0x00f0,
			Expected: []byte{0x00,0xf0},

		},
		{
			Value:              0x00f1,
			Expected: []byte{0x00,0xf1},

		},
		{
			Value:              0x00f2,
			Expected: []byte{0x00,0xf2},

		},
		{
			Value:              0x00f3,
			Expected: []byte{0x00,0xf3},

		},
		{
			Value:              0x00f4,
			Expected: []byte{0x00,0xf4},

		},
		{
			Value:              0x00f5,
			Expected: []byte{0x00,0xf5},

		},
		{
			Value:              0x00f6,
			Expected: []byte{0x00,0xf6},

		},
		{
			Value:              0x00f7,
			Expected: []byte{0x00,0xf7},

		},
		{
			Value:              0x00f8,
			Expected: []byte{0x00,0xf8},

		},
		{
			Value:              0x00f9,
			Expected: []byte{0x00,0xf9},

		},
		{
			Value:              0x00fa,
			Expected: []byte{0x00,0xfa},

		},
		{
			Value:              0x00fb,
			Expected: []byte{0x00,0xfb},

		},
		{
			Value:              0x00fc,
			Expected: []byte{0x00,0xfc},

		},
		{
			Value:              0x00fd,
			Expected: []byte{0x00,0xfd},

		},
		{
			Value:              0x00fe,
			Expected: []byte{0x00,0xfe},

		},
		{
			Value:              0x00ff,
			Expected: []byte{0x00,0xff},

		},
		{
			Value:              0x0100,
			Expected: []byte{0x01,0x00},

		},
		{
			Value:              0x0101,
			Expected: []byte{0x01,0x01},

		},
		{
			Value:              0x0102,
			Expected: []byte{0x01,0x02},

		},
		{
			Value:              0x0103,
			Expected: []byte{0x01,0x03},

		},

		{
			Value:              0x0245,
			Expected: []byte{0x02,0x45},

		},

		{
			Value:              0x37ae,
			Expected: []byte{0x37,0xae},

		},

		{
			Value:              0xfffa,
			Expected: []byte{0xff,0xfa},

		},
		{
			Value:              0xfffb,
			Expected: []byte{0xff,0xfb},

		},
		{
			Value:              0xfffc,
			Expected: []byte{0xff,0xfc},

		},
		{
			Value:              0xfffd,
			Expected: []byte{0xff,0xfd},

		},
		{
			Value:              0xfffe,
			Expected: []byte{0xff,0xfe},

		},
		{
			Value:              0xffff,
			Expected: []byte{0xff,0xff},

		},
	}


	{
		randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))

		const max = 50
		for i:=0; i<max; i++ {
			x64 := randomness.Int63n(0xffff)

			u16 := uint16(x64)

			var b0 byte = uint8((0xff00 & u16) >> 8)
			var b1 byte = uint8( 0x00ff & u16      )

			test := struct{
				Value uint16
				Expected []byte
			}{
				Value:           u16,
				Expected: []byte{b0, b1},
			}

			tests = append(tests, test)
		}
	}


	for testNumber, test := range tests {

		var buffer bytes.Buffer

		n64, err := bigendian.WriteUint16To(&buffer, test.Value)
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
