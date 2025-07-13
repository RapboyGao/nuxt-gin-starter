package utils

// 获取百分比, 结果为0-100之间的值
func Percentage(value float32, min float32, max float32) (raw float32, safe float32) {
	diff := max - min
	if diff == 0 {
		return 0, 0
	} else {
		raw = (value - min) * 100 / diff
		if raw < 0 {
			safe = 0
		} else if raw > 100 {
			safe = 100
		} else {
			safe = raw
		}
		return raw, safe
	}
}
