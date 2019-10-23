package testutils

func AssertEqual(t Tester, name string, exp, got interface{}) {
	if exp != got {
		t.Errorf("wrong %s: expected %v, got %v", name, exp, got)
	}
}

func AssertEqualFatal(t Tester, name string, exp, got interface{}) {
	if exp != got {
		t.Fatalf("wrong %s: expected %v, got %v", name, exp, got)
	}
}
