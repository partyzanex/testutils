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

func TestRandomCase(t *testing.T) {
	cases := []interface{}{"1", "abc", "qwerty", "2", "3", "0"}
	res := testutils.RandomCase(cases...)
	t.Log(res)
}

func BenchmarkRandomCase(b *testing.B) {
	cases := []interface{}{"1", "abc", "qwerty", "2", "3", "0"}

	for i := 0; i < b.N; i++ {
		testutils.RandomCase(cases...)
	}
}
