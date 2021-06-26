package datastruct

import (
	"fmt"
)

```
栈使用
```

func StackTest()  {
	// 创建栈
	stack := make([]int, 0)
	fmt.Println("创建栈==>",stack)
	// push压入
	stack = append(stack, 10)
	fmt.Println("push压入栈==>",stack)
	// pop弹出
	v := stack[len(stack)-1]
	fmt.Println("pop弹出==>",v)

	stack = stack[:len(stack)-1]
	fmt.Println("最后==>",stack)
	
	// 检查栈空
	a := len(stack) == 0

	fmt.Println(a)
}