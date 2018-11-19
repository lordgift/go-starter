package removeduplicate_test

import (
	"removeduplicate"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicate(t *testing.T) {
	tcs := []struct {
		input  []string
		output []string
	}{
		{[]string{"a", "a", "b"}, []string{"a", "b"}},
		{[]string{"c", "a", "c"}, []string{"c", "a"}},
	}
	for _, tc := range tcs {
		r := removeduplicate.RemoveDuplicate(tc.input)
		if !assert.Equal(t, r, tc.output) {
			t.Errorf("Expect %v but got %v", tc.output, r)
		}
	}
}

func BenchmarkRemoveDuplicate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		removeduplicate.RemoveDuplicate([]string{"c", "a", "c", "d", "e"})
	}
}
