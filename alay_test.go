package alay

import "testing"

type testCase struct {
	in  string
	out string
}

var toAlayTests = []testCase{
	{"saya keren", "s4y4 k3r3n"},
	{"minta donk!", "m1nt4 d0nk!"},
}

var normalizeTests []testCase

func TestToAlay(t *testing.T) {
	for _, tc := range toAlayTests {
		res := ToAlay(tc.in)
		if res != tc.out {
			t.Fatalf("ToAlay('%s') = '%s', but expecting '%s'", tc.in, res, tc.out)
		}
	}
}

func reverseTests() {
	if normalizeTests == nil {
		normalizeTests = make([]testCase, len(toAlayTests))
		for i, tc := range toAlayTests {
			normalizeTests[i] = testCase{tc.out, tc.in}
		}
	}
}

func TestNormalize(t *testing.T) {
	reverseTests()
	for _, tc := range normalizeTests {
		res := Normalize(tc.in)
		if res != tc.out {
			t.Fatalf("Normalize('%s') = '%s', but expecting '%s'", tc.in, res, tc.out)
		}
	}
}
