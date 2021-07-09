package datastruct

import (
	"fmt"
)


type LinkNode struct {
    Id     int64
	Name   string
    NextNode *LinkNode
}

//链表就是从一个数据指向另外一个数据,将数据和数据关联起来
// 链表实现：只支持顺序访问，在某些遍历操作中查询速度慢，但增删元素快。

// 1、单链表, 就是链表是单向的，像我们上面这个结构一样，可以一直往下找到下一个数据节点，它只有一个方向，它不能往回找。
// 2、双链表, 每个节点既可以找到它之前的节点，也可以找到之后的节点，是双向的。
// 3、循环链表, 就是它一直往下找数据节点，最后回到了自己那个节点，形成了一个回路。循环单链表和循环双链表的区别就是，一个只能一个方向走，一个两个方向都可以走。

//单链表
func SingleLink() {
	// datastruct.StackTest()
    // 新的节点
    node := new(LinkNode)
    node.Id = 2
    node.Name = "总经理"

    // 新的节点
    node1 := new(LinkNode)
    node1.Id = 3
    node1.Name = "主管"
    node.NextNode = node1 // node1 链接到 node 节点上

    // 新的节点
    node2 := new(LinkNode)
    node2.Id = 4
	node2.Name = "组长"
    node1.NextNode = node2 // node2 链接到 node1 节点上

    // 按顺序打印数据
    nowNode := node
    for {
        if nowNode != nil {
            // 打印节点值
            fmt.Println(nowNode.Id,nowNode.Name)
            // 获取下一个节点
            nowNode = nowNode.NextNode
            continue
        }
        // 如果下一个节点为空，表示链表结束了
        break
    }
}


//双向链表
func DoubleLink()  {
	
}



