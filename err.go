package testutils

import "github.com/stretchr/testify/assert"

func Err(t Tester, name string, err error) {
	if !assert.Equal(t, nil, err) {
		t.Errorf("%s: %s", name, err)
	}
}

func FatalErr(t Tester, name string, err error) {
	if !assert.Equal(t, nil, err) {
		t.Fatalf("%s: %s", name, err)
	}
}
