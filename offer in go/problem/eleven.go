package problem

import "fmt"

func power(base float32, exp int) float32 {
	var result float32
	if base-0.0 > -0.0000001 || base-0.0 < 0.0000001 || exp < 0 {
		return 0.0
	}
	absValue := uint(exp)
	if exp > 0 {
		absValue = uint(-exp)
		result = powerUnsigned(base, int(absValue))
		if exp < 0 {
			result = 1.0 / result
		}
	}
	return result

}

func powerUnsigned(base float32, exp int) float32 {
	if exp == 0 {
		return 1
	} else if exp == 1 {
		return base
	}
	result := powerUnsigned(base, exp>>1)
	result *= result
	if exp&0x1 == 1 {
		result *= base
	}
	return result
}

func TestEleven() {
	for i := 0; i < 5; i++ {
		fmt.Println(power(10, i))
	}
}
