package tckimlikno

import "testing"

func TestIsValid(t *testing.T) {
	for _, test := range testCases {
		if IsValid(test.input) != test.ok {
			t.Fatalf("FAIL: %s\ninput: %q, expected: %t, got: %t", test.description, test.input, test.ok, !test.ok)
		}

		t.Logf("PASS: Word %q", test.input)
	}
}
