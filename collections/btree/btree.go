package btree

type Node struct {
	Val int

	Left  *Node
	Right *Node
}
