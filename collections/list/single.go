package list

type Node struct {
	Val  int
	Next *Node
}

func SingleFromSlice(a ...int) *Node {
	if len(a) == 0 {
		return nil
	}
	head := &Node{Val: a[0]}
	pre := head
	for i := 1; i < len(a); i++ {
		node := &Node{Val: a[i]}
		pre.Next = node
		pre = node
	}
	return head
}

func (node *Node) Slice() []int {
	var slice []int
	h := node
	for h != nil {
		slice = append(slice, h.Val)
		h = h.Next
	}
	return slice
}

func (node *Node) Len() int {
	h := node
	var l int
	for h != nil {
		l++
		h = h.Next
	}
	return l
}

func (node *Node) Tail() *Node {
	if node == nil {
		return nil
	}
	h := node
	for h.Next != nil {
		h = h.Next
	}
	return h

}
