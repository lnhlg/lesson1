package lesson1

//MaximalRectangle：从给定只包含0与1的矩阵中找出只包含1的最大矩形并返回其面积
/*parameter
matrix: 只包含0与1的矩阵
return: 返回只包含1的最大矩形的面积
*/
func MaximalRectangle(matrix [][]byte) int {

	//获取高度直方图
	heights := getHeights(matrix)

	//计算面积并找出其中最大的一块
	area := 0

	for i := range heights {

		area = max(area, getMaxArea(heights[i]))
	}

	return area
}

//getHeights: 获取矩阵中每一行每一列以当前行为底的高度
/*parameter
matrix: 只包含0与1的矩阵
return: 高度直方图
*/
func getHeights(matrix [][]byte) [][]int {

	rows, cols := len(matrix), len(matrix[0])
	heights := make([][]int, rows)

	for rowIndex := range matrix {

		heights[rowIndex] = make([]int, cols)

		for colIndex := range matrix[rowIndex] {

			//当前行当前列为‘0’则高度置为0，否则取前一行的高度值加1
			if matrix[rowIndex][colIndex] == '0' {

				heights[rowIndex][colIndex] = 0
			} else {

				if rowIndex == 0 {

					heights[rowIndex][colIndex] = 1
				} else {

					heights[rowIndex][colIndex] = heights[rowIndex-1][colIndex] + 1
				}
			}
		}
	}

	return heights
}

//getMaxArea: 从指定行的高度直方图中取出最大面积
/*parameter
rowHeights: 行中每列的高度
return: 最大的面积值
*/
func getMaxArea(rowHeights []int) int {

	colNum := len(rowHeights)
	//边界数组（索引)
	up, down := make([]int, colNum), make([]int, colNum)
	//索引栈
	indexStack := make([]int, 0)
	//增加负索引以处理边界溢出的问题
	indexStack = append(indexStack, -1)
	stackSize := 1
	area := 0

	for colIndex := range rowHeights {

		down[colIndex] = colNum

		//当前列高度小于栈顶索引指示的高度，则出栈
		for indexStack[stackSize-1] >= 0 && rowHeights[indexStack[stackSize-1]] > rowHeights[colIndex] {

			down[indexStack[stackSize-1]] = colIndex
			indexStack = indexStack[:stackSize-1]
			stackSize--
		}

		up[colIndex] = indexStack[stackSize-1]
		indexStack = append(indexStack, colIndex)
		stackSize++
	}

	for i := 0; i < colNum; i++ {

		area = max(area, rowHeights[i]*(down[i]-up[i]-1))
	}

	return area
}

func max(x, y int) int {

	if x > y {

		return x
	}

	return y
}
