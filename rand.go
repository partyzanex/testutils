package testutils

import (
	"fmt"
	"image"
	"image/color"
	"math/rand"
	"sync"
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
	set    = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890")
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

func RandomCase(args ...interface{}) interface{} {
	return args[RandInt(0, len(args))]
}

func RandomImage(width, height int) image.Image {
	imq := image.NewRGBA(image.Rectangle{
		Min: image.Point{X: 0, Y: 0},
		Max: image.Point{X: width, Y: height},
	})

	var r, g, b uint8
	clr := &color.RGBA{
		A: 0xff,
	}

	pool := &sync.Pool{
		New: func() interface{} {
			return clr
		},
	}

	for x := 0; x < width; x++ {
		r = uint8(RandInt(0, 255))
		g = uint8(RandInt(0, 255))
		b = uint8(RandInt(0, 255))

		c := pool.New().(*color.RGBA)

		c.R = r
		c.G = g
		c.B = b

		for y := 0; y < height; y++ {
			imq.Set(x, y, c)
		}
	}

	return imq
}
