package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
请你将两个数相加，并以相同形式返回一个表示和的链表。
你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
输出：[8,9,9,9,0,0,0,1]
**/

func main() {
	var t *two_num_sum
	t.test()
}

type two_num_sum struct{}

type ListNode struct {
	num  int
	Next *ListNode
}

func (t *two_num_sum) sum(l1 *ListNode, l2 *ListNode) *ListNode {
	l1stack := []*ListNode{}
	l2stack := []*ListNode{}
	l3stack := []*ListNode{}
	for l1 != nil {
		l1stack = append(l1stack, l1)
		l1 = l1.Next
	}
	for l2 != nil {
		l2stack = append(l2stack, l2)
		l2 = l2.Next
	}
	lenl1 := len(l1stack)
	fmt.Printf("sum_f:%v\n", lenl1)
	lenl2 := len(l2stack)
	fmt.Printf("sum_f:%v\n", lenl2)
	jin := 0
	for _i, _j := 0, 0; _i < lenl1 || _j < lenl2; {
		sum := jin
		if _i < lenl1 {
			sum += l1stack[_i].num
		}
		if _j < lenl2 {
			sum += l2stack[_j].num
		}
		if sum > 9 {
			sum = sum - 10
			jin = 1
		} else {
			jin = 0
		}
		node := new(ListNode)
		node.num = sum
		l3stack = append(l3stack, node)
		_i++
		_j++
	}
	if jin > 0 {
		node := new(ListNode)
		node.num = jin
		l3stack = append(l3stack, node)
	}
	head := new(ListNode)
	pre := head
	for _, v := range l3stack {
		head.Next = v
		head = head.Next
	}
	return pre.Next
}

func (t *two_num_sum) test() {
	head1 := new(ListNode)
	pre1 := head1
	rand.Seed(time.Now().UnixNano())
	count := rand.Intn(10)
	fmt.Printf("l1:\n")
	for i := 0; i < count; i++ {
		node := new(ListNode)
		node.num = rand.Intn(10)
		head1.Next = node
		head1 = head1.Next
		fmt.Printf("%d=>", node.num)
	}
	head2 := new(ListNode)
	pre2 := head2
	rand.Seed(time.Now().UnixNano())
	count2 := rand.Intn(10)
	fmt.Printf("\nl2:\n")
	for i := 0; i < count2; i++ {
		node := new(ListNode)
		node.num = rand.Intn(10)
		head2.Next = node
		head2 = head2.Next
		fmt.Printf("%d=>", node.num)
	}
	fmt.Printf("\n结果：\n")
	h := t.sum1(pre1.Next, pre2.Next)
	for h != nil {
		fmt.Printf("%d=>", h.num)
		h = h.Next
	}
}

func (t *two_num_sum) sum1(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	tail := new(ListNode)
	pre := tail
	for l1 != nil || l2 != nil {
		sum := carry
		if l1 != nil {
			sum += l1.num
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.num
			l2 = l2.Next
		}
		carry = sum / 10
		sum = sum % 10
		node := new(ListNode)
		node.num = sum
		tail.Next = node
		tail = tail.Next
	}
	if carry > 0 {
		node := new(ListNode)
		node.num = carry
		tail.Next = node
	}
	return pre.Next
}
