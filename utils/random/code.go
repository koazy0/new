package random

import (
	"fmt"
	"math/rand"
)

func Code(length int) string {
	// go 1.20之后 内置了随机种子 因此不需要每次执行时初始化
	// rand.Seed(time.Now().UnixNano())

	return fmt.Sprintf("%4v", rand.Intn(10*length))
}
