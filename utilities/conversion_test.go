package utilities

import "testing"

type testpair struct {
	value          bool
	expectedOutput string
}

var tests = []testpair{
	{true, "true"},
	{false, "false"},
}

func TestBtos(t *testing.T) {
	var sampleBool string
	sampleBool = Btos(true)
	if sampleBool != "true" {
		t.Error("Expected 'true', got ", sampleBool)
	}

	for _, pair := range tests {
		v := Btos(pair.value)
		if v != pair.expectedOutput {
			t.Error("for ", pair.value, " expected ", pair.expectedOutput, " got ", v)

		}
	}

}
