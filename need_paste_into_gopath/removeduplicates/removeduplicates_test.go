package removeduplicates

import "testing"

func TestRemoveDuplicate(t *testing.T) {
	tcs := []struct {
		input  []string
		output []string
	}{
		{[]string{"a", "a", "b"}, []string{"a", "b"}},
	}

	for _, tc := range tcs {
		r := RemoveDuplicate(tc.input)
		if !sliceStringEqual(r, tc.output) {
			t.Errorf("Expected %v but got %v", tc.input, tc.output)
		}
	}
}

func sliceStringEqual(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
