package lesson1

//PlusOne：加一
/*parameter
digits: 整数型数组
return: 加一后的整数型数组
*/
func PlusOne(digits []int) []int {

	bits := len(digits)

	//从后往前计算
	for i := bits - 1; i >= 0; i-- {

		//遇9置0，否则加1
		if digits[i] == 9 {

			digits[i] = 0
		} else {

			digits[i]++

			return digits
		}
	}

	//处理进位导致新数组比原数组长的问题
	result := make([]int, bits+1)
	result[0] = 1

	return result
}
