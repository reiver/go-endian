package bigendian_test

import (
	"github.com/reiver/go-endian/big"

	"bytes"
	"math/rand"
	"time"

	"testing"
)

func TestWriteUint8To(t *testing.T) {

	tests := []struct{
		Value uint8
		Expected []byte
	}{
		{
			Value:           0x00,
			Expected: []byte{0x00},

		},
		{
			Value:           0x01,
			Expected: []byte{0x01},

		},
		{
			Value:           0x02,
			Expected: []byte{0x02},

		},
		{
			Value:           0x03,
			Expected: []byte{0x03},

		},
		{
			Value:           0x04,
			Expected: []byte{0x04},

		},
		{
			Value:           0x05,
			Expected: []byte{0x05},

		},
		{
			Value:           0x06,
			Expected: []byte{0x06},

		},
		{
			Value:           0x07,
			Expected: []byte{0x07},

		},
		{
			Value:           0x08,
			Expected: []byte{0x08},

		},
		{
			Value:           0x09,
			Expected: []byte{0x09},

		},
		{
			Value:           0x0a,
			Expected: []byte{0x0a},

		},
		{
			Value:           0x0b,
			Expected: []byte{0x0b},

		},
		{
			Value:           0x0c,
			Expected: []byte{0x0c},

		},
		{
			Value:           0x0d,
			Expected: []byte{0x0d},

		},
		{
			Value:           0x0e,
			Expected: []byte{0x0e},

		},
		{
			Value:           0x0f,
			Expected: []byte{0x0f},

		},
		{
			Value:           0x10,
			Expected: []byte{0x10},

		},
		{
			Value:           0x11,
			Expected: []byte{0x11},

		},
		{
			Value:           0x12,
			Expected: []byte{0x12},

		},
		{
			Value:           0x13,
			Expected: []byte{0x13},

		},

		{
			Value:           0xef,
			Expected: []byte{0xef},

		},
		{
			Value:           0xf0,
			Expected: []byte{0xf0},

		},
		{
			Value:           0xf1,
			Expected: []byte{0xf1},

		},
		{
			Value:           0xf2,
			Expected: []byte{0xf2},

		},
		{
			Value:           0xf3,
			Expected: []byte{0xf3},

		},
		{
			Value:           0xf4,
			Expected: []byte{0xf4},

		},
		{
			Value:           0xf5,
			Expected: []byte{0xf5},

		},
		{
			Value:           0xf6,
			Expected: []byte{0xf6},

		},
		{
			Value:           0xf7,
			Expected: []byte{0xf7},

		},
		{
			Value:           0xf8,
			Expected: []byte{0xf8},

		},
		{
			Value:           0xf9,
			Expected: []byte{0xf9},

		},
		{
			Value:           0xfa,
			Expected: []byte{0xfa},

		},
		{
			Value:           0xfb,
			Expected: []byte{0xfb},

		},
		{
			Value:           0xfc,
			Expected: []byte{0xfc},

		},
		{
			Value:           0xfd,
			Expected: []byte{0xfd},

		},
		{
			Value:           0xfe,
			Expected: []byte{0xfe},

		},
		{
			Value:           0xff,
			Expected: []byte{0xff},

		},
	}


	{
		randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))

		const max = 50
		for i:=0; i<max; i++ {
			x64 := randomness.Int63n(0xffff)

			u8 := uint8(x64)

			var b0 byte = u8

			test := struct{
				Value uint8
				Expected []byte
			}{
				Value:           u8,
				Expected: []byte{b0},
			}

			tests = append(tests, test)
		}
	}


	for testNumber, test := range tests {

		var buffer bytes.Buffer

		n64, err := bigendian.WriteUint8To(&buffer, test.Value)
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
