package endian

import (
	"testing"
)

func TestTypeUnmarshalText(t *testing.T) {

	tests := []struct{
		Value []byte
		Expected Type
	}{
		{
			Value: []byte("endianness-none"),
			Expected: None(),
		},
		{
			Value: []byte("endianness-unknown"),
			Expected: Unknown(),
		},
		{
			Value: []byte("little-endian"),
			Expected: Little(),
		},
		{
			Value: []byte("big-endian"),
			Expected: Big(),
		},
	}


	for testNumber, test := range tests {

		var actual Type

		err := actual.UnmarshalText(test.Value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %q", testNumber, err, err)
			continue
		}

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d, expected %q, but actually got %q.", testNumber, expected, actual)
			continue
		}
	}
}
