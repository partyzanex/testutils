package testutils_test

import (
	"github.com/partyzanex/testutils"
	"testing"
)

func TestRandomString(t *testing.T) {
	str := testutils.RandomString(100)
	testutils.AssertEqual(t, "len", 100, len(str))
}

func BenchmarkRandomString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testutils.RandomString(100)
	}
}
