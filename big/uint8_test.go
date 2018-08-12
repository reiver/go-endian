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

func TestReadUint8From(t *testing.T) {

	tests := []struct{
		Value []byte
		Expected uint8
	}{
		{
			Value: []byte{0x00},
			Expected:     0x00,
		},
		{
			Value: []byte{0x01},
			Expected:     0x01,
		},
		{
			Value: []byte{0x02},
			Expected:     0x02,
		},
		{
			Value: []byte{0x03},
			Expected:     0x03,
		},
		{
			Value: []byte{0x04},
			Expected:     0x04,
		},
		{
			Value: []byte{0x05},
			Expected:     0x05,
		},
		{
			Value: []byte{0x06},
			Expected:     0x06,
		},
		{
			Value: []byte{0x07},
			Expected:     0x07,
		},
		{
			Value: []byte{0x08},
			Expected:     0x08,
		},
		{
			Value: []byte{0x09},
			Expected:     0x09,
		},
		{
			Value: []byte{0x0a},
			Expected:     0x0a,
		},
		{
			Value: []byte{0x0b},
			Expected:     0x0b,
		},
		{
			Value: []byte{0x0c},
			Expected:     0x0c,
		},
		{
			Value: []byte{0x0d},
			Expected:     0x0d,
		},
		{
			Value: []byte{0x0e},
			Expected:     0x0e,
		},
		{
			Value: []byte{0x0f},
			Expected:     0x0f,
		},
		{
			Value: []byte{0x10},
			Expected:     0x10,
		},
		{
			Value: []byte{0x11},
			Expected:     0x11,
		},
		{
			Value: []byte{0x12},
			Expected:     0x12,
		},
		{
			Value: []byte{0x13},
			Expected:     0x13,
		},

		{
			Value: []byte{0xef},
			Expected:     0xef,
		},
		{
			Value: []byte{0xf0},
			Expected:     0xf0,
		},
		{
			Value: []byte{0xf1},
			Expected:     0xf1,
		},
		{
			Value: []byte{0xf2},
			Expected:     0xf2,
		},
		{
			Value: []byte{0xf3},
			Expected:     0xf3,
		},
		{
			Value: []byte{0xf4},
			Expected:     0xf4,
		},
		{
			Value: []byte{0xf5},
			Expected:     0xf5,
		},
		{
			Value: []byte{0xf6},
			Expected:     0xf6,
		},
		{
			Value: []byte{0xf7},
			Expected:     0xf7,
		},
		{
			Value: []byte{0xf8},
			Expected:     0xf8,
		},
		{
			Value: []byte{0xf9},
			Expected:     0xf9,
		},
		{
			Value: []byte{0xfa},
			Expected:     0xfa,
		},
		{
			Value: []byte{0xfb},
			Expected:     0xfb,
		},
		{
			Value: []byte{0xfc},
			Expected:     0xfc,
		},
		{
			Value: []byte{0xfd},
			Expected:     0xfd,
		},
		{
			Value: []byte{0xfe},
			Expected:     0xfe,
		},
		{
			Value: []byte{0xff},
			Expected:     0xff,
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
				Value []byte
				Expected uint8
			}{
				Value: []byte{b0},
				Expected:     u8,
			}

			tests = append(tests, test)
		}
	}


	for testNumber, test := range tests {

		var buffer bytes.Buffer

		buffer.Write(test.Value)

		var actual uint8

		n64, err := bigendian.ReadUint8From(&buffer, &actual)
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
