package datastruct

import "fmt"

type ValNode struct {
	col int
	row int
	val int
}

func TestSparsearray() {
	// 1.先创建一个原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1 //黑子
	chessMap[2][3] = 2 //白子
	// 2.输出原始的数组
	for _,v := range chessMap{
		for _,v2 := range v{
			fmt.Printf("%d\t",v2)
		}
		fmt.Println()
	}
	// 3.转成稀疏数组
	// 思路
	// (1).遍历chessMap，如果发现有一个元素的值不为0，我们就创建一个node节点，
	// (2).将其值放入到对应的数组中
	var sparseArr []ValNode
	// 标准的稀疏数组应该还含有一个表示记录原始的二维数组的规模和默认值(行，列，默认值)
	valNode := ValNode{
		col: 11,
		row: 11,
		val: 0,
	}
	sparseArr = append(sparseArr,valNode)
	for i,v := range chessMap{
		for j,v2 := range v{
			if v2 != 0{
				// 创建一个valNode值节点
				valNode := ValNode{
					row: i,
					col: j,
					val: v2,
				}
				sparseArr = append(sparseArr,valNode)
			}
		}
	}
	// 输出稀疏数组
	fmt.Println("当前的稀疏数组是：：：：")
	for i,valNode := range sparseArr {
		fmt.Printf("%d: %d %d %d\n",i,valNode.row,valNode.col,valNode.val)
	}
	// 恢复成原始数组
	// 先定义一个原始数组
	var chessMap2 [11][11]int
	// 遍历稀疏数组
	for i,valNode := range sparseArr{
		// 跳过第一行记录数组维度的数据
		if i != 0{
			chessMap2[valNode.row][valNode.col] = valNode.val
		}
	}
	// 看看chessmap2是不是恢复了
	for _,v := range chessMap{
		for _,v2 := range v{
			fmt.Printf("%d\t",v2)
		}
		fmt.Println()
	}
}

