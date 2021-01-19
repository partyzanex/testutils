package testutils

import "github.com/stretchr/testify/assert"

func Err(t Tester, _ string, err error) {
	assert.Equal(t, nil, err)
}

func FatalErr(t Tester, name string, err error) {
	if !assert.Equal(t, nil, err) {
		t.Fatalf("%s: %s", name, err)
	}
}
