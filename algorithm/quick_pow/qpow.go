package main

// x 的 n 次方, x、n 都是非负整数
func Qpow(x, n int64) int64 {
	if x == 0 {
		return 0
	}
	// 用滚雪球的方式计算幂
	// 雪球初始值是 1
	var result int64 = 1
	// 处理所有的二进制位
	for n > 0 {
		// 如果不是1的话这位可以不做运算
		if n&1 == 0 {
			result *= x
		}
		/**
		 * 每轮结束时将滚动因子自乘, 因为每行进一轮, 指数都翻倍, 整体结果就是自乘
		 * 比如本轮因子是 2**4, 下一轮就是 2**8, 2**8 = 2**(4+4) = 2**4 * 2**4
		 */
		x *= x
		n >>= 1
	}
	return result
}
