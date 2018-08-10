package littleendian_test

import (
	"github.com/reiver/go-endian/little"

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
			Value:         0x0000,
			Expected: []byte{0x00,0x00},

		},
		{
			Value:         0x0001,
			Expected: []byte{0x01,0x00},

		},
		{
			Value:         0x0002,
			Expected: []byte{0x02,0x00},

		},
		{
			Value:         0x0003,
			Expected: []byte{0x03,0x00},

		},
		{
			Value:         0x0004,
			Expected: []byte{0x04,0x00},

		},
		{
			Value:         0x0005,
			Expected: []byte{0x05,0x00},

		},
		{
			Value:         0x0006,
			Expected: []byte{0x06,0x00},

		},
		{
			Value:         0x0007,
			Expected: []byte{0x07,0x00},

		},
		{
			Value:         0x0008,
			Expected: []byte{0x08,0x00},

		},
		{
			Value:         0x0009,
			Expected: []byte{0x09,0x00},

		},
		{
			Value:         0x000a,
			Expected: []byte{0x0a,0x00},

		},
		{
			Value:         0x000b,
			Expected: []byte{0x0b,0x00},

		},
		{
			Value:         0x000c,
			Expected: []byte{0x0c,0x00},

		},
		{
			Value:         0x000d,
			Expected: []byte{0x0d,0x00},

		},
		{
			Value:         0x000e,
			Expected: []byte{0x0e,0x00},

		},
		{
			Value:         0x000f,
			Expected: []byte{0x0f,0x00},

		},
		{
			Value:         0x0010,
			Expected: []byte{0x10,0x00},

		},
		{
			Value:         0x0011,
			Expected: []byte{0x11,0x00},

		},
		{
			Value:         0x0012,
			Expected: []byte{0x12,0x00},

		},
		{
			Value:         0x0013,
			Expected: []byte{0x13,0x00},

		},

		{
			Value:         0x00ef,
			Expected: []byte{0xef,0x00},

		},
		{
			Value:         0x00f0,
			Expected: []byte{0xf0,0x00},

		},
		{
			Value:         0x00f1,
			Expected: []byte{0xf1,0x00},

		},
		{
			Value:         0x00f2,
			Expected: []byte{0xf2,0x00},

		},
		{
			Value:         0x00f3,
			Expected: []byte{0xf3,0x00},

		},
		{
			Value:         0x00f4,
			Expected: []byte{0xf4,0x00},

		},
		{
			Value:         0x00f5,
			Expected: []byte{0xf5,0x00},

		},
		{
			Value:         0x00f6,
			Expected: []byte{0xf6,0x00},

		},
		{
			Value:         0x00f7,
			Expected: []byte{0xf7,0x00},

		},
		{
			Value:         0x00f8,
			Expected: []byte{0xf8,0x00},

		},
		{
			Value:         0x00f9,
			Expected: []byte{0xf9,0x00},

		},
		{
			Value:         0x00fa,
			Expected: []byte{0xfa,0x00},

		},
		{
			Value:         0x00fb,
			Expected: []byte{0xfb,0x00},

		},
		{
			Value:         0x00fc,
			Expected: []byte{0xfc,0x00},

		},
		{
			Value:         0x00fd,
			Expected: []byte{0xfd,0x00},

		},
		{
			Value:         0x00fe,
			Expected: []byte{0xfe,0x00},

		},
		{
			Value:         0x00ff,
			Expected: []byte{0xff,0x00},

		},
		{
			Value:         0x0100,
			Expected: []byte{0x00,0x01},

		},
		{
			Value:         0x0101,
			Expected: []byte{0x01,0x01},

		},
		{
			Value:         0x0102,
			Expected: []byte{0x02,0x01},

		},
		{
			Value:         0x0103,
			Expected: []byte{0x03,0x01},

		},

		{
			Value:         0x0245,
			Expected: []byte{0x45,0x02},

		},

		{
			Value:         0x37ae,
			Expected: []byte{0xae,0x37},

		},

		{
			Value:         0xfffa,
			Expected: []byte{0xfa,0xff},

		},
		{
			Value:         0xfffb,
			Expected: []byte{0xfb,0xff},

		},
		{
			Value:         0xfffc,
			Expected: []byte{0xfc,0xff},

		},
		{
			Value:         0xfffd,
			Expected: []byte{0xfd,0xff},

		},
		{
			Value:         0xfffe,
			Expected: []byte{0xfe,0xff},

		},
		{
			Value:         0xffff,
			Expected: []byte{0xff,0xff},

		},
	}


	{
		randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))

		const max = 50
		for i:=0; i<max; i++ {
			x64 := randomness.Int63n(0xffff)

			u16 := uint16(x64)

			var b0 byte = uint8( 0x00ff & u16      )
			var b1 byte = uint8((0xff00 & u16) >> 8)

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

		n64, err := littleendian.WriteUint16To(&buffer, test.Value)
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
