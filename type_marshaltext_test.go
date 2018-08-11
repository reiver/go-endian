package endian

import (
	"bytes"

	"testing"
)

func TestTypeMarshalText(t *testing.T) {

	tests := []struct{
		Value Type
		Expected []byte
	}{
		{
			Value: None(),
			Expected: []byte("endianness-none"),
		},
		{
			Value: Unknown(),
			Expected: []byte("endianness-unknown"),
		},
		{
			Value: Little(),
			Expected: []byte("little-endian"),
		},
		{
			Value: Big(),
			Expected: []byte("big-endian"),
		},
	}


	for testNumber, test := range tests {

		actual, err := test.Value.MarshalText()
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %q", testNumber, err, err)
			continue
		}

		if expected := test.Expected; !bytes.Equal(expected, actual) {
			t.Errorf("For test #%d, expected %q, but actually got %q.", testNumber, expected, actual)
			continue
		}
	}
}
