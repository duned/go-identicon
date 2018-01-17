package identicon

import (
	"crypto/md5"
	"fmt"
)

type Identicon struct {
	grid [25]byte
}

func Generate(key string) Identicon {
	hash := md5.Sum([]byte(key))
	fmt.Println(hash)
	return Identicon{
	}
}

func generatePatternFromHash(sum [16]byte) []byte {
	p := make([]byte, 25)

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			jCount := j

			if j > 2 {
				jCount = 4 - j
			}

			p[5 * i + j] = sum[3 * i + jCount]
		}
	}

	return p
}

func convertPatternToBinarySwitch(pattern []byte) []byte {
	b := make([]byte, 25)
	for i, v := range pattern {
		if v%2 == 0 {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}

	return b
}
