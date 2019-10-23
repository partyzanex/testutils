# testutils

Test toolkit for simplify the development of tests on Go

This package contains universal interfaces and utilities for handling errors in tests and benchmarks

### Example

```go
package some_test

import (
	"gitlab.com/partyzan65/testutils"
	"testing"
	"time"
)

func TestSomeTestName(t *testing.T) {
	// supported mysql and postgres
	db := testutils.NewSqlDB(t, "postgres", "TEST_DSN")

	rows, err := db.Query(`select id, name from tbl where dt > ?`, time.Now())
	testutils.FatalErr(t, "db.Query", err)

	// ...
}
```

Run:
```bash
export TEST_DSN="user=postgres password=password sslmode=disable dbname=testdb"

go test ./path/to/some -v
```