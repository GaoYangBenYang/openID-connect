package test

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(val int) *ListNode {
	return &ListNode{
		Val:  val,
		Next: nil,
	}
}

// 在N节点后插入P节点
func InsertListNode(n, p *ListNode) {
	temp := n.Next
	n.Next = p
	p.Next = temp
}

// 删除节点N之后的首个节点
func DeletetListNode(n *ListNode) {
	if n.Next == nil {
		return
	}
	temp := n.Next
	n.Next = temp.Next
}

// 访问节点
func SelectListNode(head *ListNode, index int) *ListNode {
	for i := 0; i < index; i++ {
		fmt.Println(head.Val)
		if head.Next == nil {
			return nil
		}
		head = head.Next
	}
	return head
}

func TestRedis(t *testing.T) {
	// 初始化链表
	n0 := NewListNode(0)
	n1 := NewListNode(1)
	n2 := NewListNode(2)
	n3 := NewListNode(3)
	n4 := NewListNode(4)
	//构建引用指向
	n0.Next = n1
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4

	p := NewListNode(999)
	//插入节点
	InsertListNode(n2, p)
	//删除节点
	DeletetListNode(n2)
	DeletetListNode(n2)
	//访问节点
	SelectListNode(n0, 6)
	//查找节点
	FindListNode(n0, 3)
}

func FindListNode(head *ListNode, i int) int {
	index := 0
	for head != nil {
		if head.Val == i {
			return index
		}
		head = head.Next
		index++
	}
	return -1
}