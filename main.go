package main

import (
	"math/rand"
	"time"

	t "github.com/natholdallas/gotest/pkg/tools"
)

// 生成指定位数的伪随机数（可逆）
func generateReversibleRandom(digits, seed int) (int, int) {
	rand.Seed(time.Now().UnixNano())
	min := pow10(digits - 1)
	max := pow10(digits) - 1
	randomNum := min + rand.Intn(max-min+1)
	mask := 0xABCD1234             // 可自定义掩码（密钥）
	obfuscated := randomNum ^ mask // XOR 运算可逆
	return randomNum, obfuscated
}

// 逆向解码
func reverseObfuscated(obfuscated, mask int) int {
	return obfuscated ^ mask // 再次 XOR 即可还原
}

// 计算 10^n
func pow10(n int) int {
	if n == 0 {
		return 1
	}
	return 10 * pow10(n-1)
}

func main() {
	// digits := 6        // 6 位数
	// mask := 0xABCD1234 // 掩码（需保密）
	//
	// // 1. 生成伪随机数 + 可逆变换
	// original, obfuscated := generateReversibleRandom(digits, mask)
	// fmt.Printf("Original (%d-digit): %d\n", digits, original)
	// fmt.Printf("Obfuscated: %d\n", obfuscated)
	//
	// // 2. 逆向解码
	// recovered := reverseObfuscated(obfuscated, mask)
	// fmt.Printf("Recovered: %d\n", recovered)

	a := []int{1, 2, 3, 4, 5}
	a = a[1:]
	t.PrintJSON(a)
}
