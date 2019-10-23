package testutils

func Err(t Tester, name string, err error) {
	if err != nil {
		t.Errorf("%s: %s", name, err)
	}
}

func FatalErr(t Tester, name string, err error) {
	if err != nil {
		t.Fatalf("%s: %s", name, err)
	}
}
