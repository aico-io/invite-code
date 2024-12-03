package invite

import (
	"errors"
	"strings"
)

const CHARSET = "97FEMpQdLjq2ca3yGU5ZrHB84bDznYkWeRSgKoXmJh6itCuNvATsPxwVf"

type number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Generator[T number] struct {
	charset      string
	length       int
	coprime      int
	decodeFactor uint16
	maxSupport   uint64
	base         T
}

func NewGenerator[T number](charset string, length uint8) *Generator[T] {
	base := T(len(charset))
	if len(charset) == 0 || length <= 0 {
		panic("invalid create generator params")
	}

	return &Generator[T]{
		charset:      charset,
		length:       int(length),
		coprime:      int(minCoprime(uint64(length))),
		decodeFactor: uint16(base) * uint16(length),
		maxSupport:   pow(uint64(base), uint64(length)) - 1,
		base:         base,
	}
}

func (g *Generator[T]) MaxSupportID() uint64 {
	return g.maxSupport
}

// Encode 通过id获取指定code（进制法+扩散+混淆）
func (g *Generator[T]) Encode(id T) (string, error) {
	if id < 0 || uint64(id) > g.maxSupport {
		return "", errors.New("id out of range")
	}

	idx := make([]uint16, g.length)

	// 扩散
	for i := 0; i < g.length; i++ {
		idx[i] = uint16(id % g.base)
		idx[i] = (idx[i] + uint16(i)*idx[0]) % uint16(g.base)
		id /= g.base
	}

	// 混淆
	var buf strings.Builder
	buf.Grow(g.length)
	for i := 0; i < g.length; i++ {
		n := i * g.coprime % g.length
		buf.WriteByte(g.charset[idx[n]])
	}

	return buf.String(), nil
}

// Decode 通过code反推id
func (g *Generator[T]) Decode(code string) T {
	idx := make([]uint16, g.length)
	for i, c := range code {
		idx[i*g.coprime%g.length] = uint16(strings.IndexRune(g.charset, c)) // 反推下标数组
	}

	var id T
	for i := g.length - 1; i >= 0; i-- {
		id *= g.base
		idx[i] = (idx[i] + g.decodeFactor - idx[0]*uint16(i)) % uint16(g.base)
		id += T(idx[i])
	}

	return id
}

// 求uint64类型n的最小互质数
func minCoprime(n uint64) uint64 {
	// 如果n是1，那么最小互质数是2
	if n == 1 {
		return 2
	}
	// 从2开始遍历，找到第一个和n互质的数
	for i := uint64(2); i < n; i++ {
		// 如果i和n的最大公约数是1，那么i就是最小互质数
		if isCoprime(i, n) {
			return i
		}
	}
	// 如果没有找到，那么返回n+1，因为n+1一定和n互质
	return n + 1
}

// 判断两个数是否互质
func isCoprime(n, m uint64) bool {
	// 求最大公因数
	return gcd(n, m) == 1
}

// 辗转相除法求最大公因数
func gcd(n, m uint64) uint64 {
	if m == 0 {
		return n
	}
	return gcd(m, n%m)
}

// 求n的m次方
func pow(n, m uint64) uint64 {
	sum := n
	for i := uint64(1); i < m; i++ {
		sum *= n
	}
	return sum
}
