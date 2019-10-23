package testutils

type Tester interface {
	Skip(args ...interface{})
	Skipf(format string, args ...interface{})

	Log(args ...interface{})
	Logf(format string, args ...interface{})

	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})

	Error(args ...interface{})
	Errorf(format string, args ...interface{})
}
