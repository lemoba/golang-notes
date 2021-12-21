package trans

import "math"

var Pi float64

// 1. 每个包只能包含一个init函数
// 2. 自动调用,优先级比main函数高
func init() {
	Pi = 4 * math.Atan(1)
}
