package randnum

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Rand(min, max, num int) ([]int, error) {
	result := make([]int, num)
	for i := 0; i < num; i++ {
		result[i] = min + rand.Intn(max-min)
	}
	return result, nil
}

func RandPerm(min, max, num int) ([][]int, error) {
	result := make([][]int, num)

	for i := 0; i < num; i++ {
		result[i] = rand.Perm(max - min)
		for j := 0; j < max-min; j++ {
			result[i][j] = result[i][j] + min
		}
	}

	return result, nil
}

func RandPick(min, max, pick, num int) ([][]int, error) {
	if num > max-min {
		num = max - min
	}

	result := make([][]int, num)

	for i := 0; i < num; i++ {
		result[i] = make([]int, pick)
		v := rand.Perm(max - min)
		for j := 0; j < pick; j++ {
			result[i][j] = v[j] + min
		}
	}

	return result, nil
}
