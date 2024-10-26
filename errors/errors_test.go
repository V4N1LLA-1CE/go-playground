package timeparser

import "testing"

func TestParseTime(t *testing.T) {
	testcases := []struct {
		time string
		ok   bool
	}{
		{"00:00:00", true},
		{"24:00:00", false},
		{"23:59:59", true},
		{"aa:bb:cc", false},
		{"", false},
		{"11:12", false},
		{"11:12:", false},
		{"11:12: ", false},
		{"-0:00:00", false},
		{"-12:23:00", false},
		{"0:59:00", true},
	}

	for _, test := range testcases {
		_, err := ParseTime(test.time)
		if test.ok && err != nil {
			t.Errorf("%v: %v, error should be nil", test.time, err)
		}
	}
}
