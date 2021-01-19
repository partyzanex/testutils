package testutils

import "github.com/stretchr/testify/assert"

func AssertEqual(t Tester, _ string, exp, got interface{}) {
	assert.Equal(t, exp, got)
}

func AssertEqualFatal(t Tester, name string, exp, got interface{}) {
	if !assert.Equal(t, exp, got) {
		t.Fatalf("wrong %s: expected %v, got %v", name, exp, got)
	}
}
