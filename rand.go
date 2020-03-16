package testutils

import (
	"fmt"
	"math/rand"
	"time"
)

func RandInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}

func RandInt64(min, max int64) int64 {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Int63n(max-min)
}

func RandFloat64(min, max float64) float64 {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Float64()*(max+min)
}

func RandomIP() string {
	return fmt.Sprintf(
		"%d.%d.%d.%d",
		RandInt(1, 255),
		RandInt(1, 255),
		RandInt(1, 255),
		RandInt(1, 255),
	)
}

func RandomDate(min, max time.Time) time.Time {
	ts := RandInt64(min.Unix(), max.Unix())
	return time.Unix(ts, 0)
}

var (
	set    = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890~!@#$%^&*()-_+={}[]\\|<,>.?/\"';:`")
	setLen = len(set) - 1
)

func RandomString(n int) string {
	rand.Seed(time.Now().UnixNano())
	randStr := make([]byte, n)

	for i := 0; i < n; i++ {
		r := rand.Intn(setLen)
		randStr[i] = set[r]
	}

	return string(randStr)
}
