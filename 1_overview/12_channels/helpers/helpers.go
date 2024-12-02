package helpers

import "math/rand"

func RandomNumber(n int) int {
	randomVal := rand.Intn(n)
	return int(randomVal)
}
