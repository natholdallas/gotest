package main

import "fmt"

// LCG 参数（必须满足可逆条件）
const (
	a = 15325    // 乘数
	c = 0 // 增量
	m = 1 << 32    // 模数（2^32）
)

// 计算模逆元（a^-1 mod m）
func modInverse(a, m int64) int64 {
	// 使用扩展欧几里得算法求逆元
	var x, y int64
	g := extendedGCD(a, m, &x, &y)
	if g != 1 {
		panic("a 和 m 必须互质")
	}
	return (x%m + m) % m // 确保返回正数
}

// 扩展欧几里得算法
func extendedGCD(a, b int64, x, y *int64) int64 {
	if b == 0 {
		*x = 1
		*y = 0
		return a
	}
	var x1, y1 int64
	gcd := extendedGCD(b, a%b, &x1, &y1)
	*x = y1
	*y = x1 - (a/b)*y1
	return gcd
}

// 正向计算下一个随机数
func nextRandom(seed int64) int64 {
	return (a*seed + c) % m
}

// 逆向计算上一个随机数
func prevRandom(current int64) int64 {
	aInv := modInverse(a, m)
	return (aInv * (current - c)) % m
}

func main() {
	seed := int64(400) // 初始种子

	// 正向生成随机数
	r1 := nextRandom(seed)
	r2 := nextRandom(r1)
	r3 := nextRandom(r2)

	fmt.Println("正向序列:")
	fmt.Println("r1 =", r1) // 1664525*1 + 1013904223 mod 2^32
	fmt.Println("r2 =", r2)
	fmt.Println("r3 =", r3)

	// 逆向推算
	p2 := prevRandom(r3)
	p1 := prevRandom(p2)
	p0 := prevRandom(p1)

	fmt.Println("\n逆向推算:")
	fmt.Println("p2 =", p2, "(应该是 r2)")
	fmt.Println("p1 =", p1, "(应该是 r1)")
	fmt.Println("p0 =", p0, "(应该是 seed=1)")
}
