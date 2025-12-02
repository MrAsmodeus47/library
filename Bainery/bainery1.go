package main

import "fmt"

type Node struct {
	val   int
	right *Node
	left  *Node
}

func (n *Node) insert(val int) *Node {
	if n == nil {
		return &Node{val: val}
	}

	if val < n.val {
		n.left = n.left.insert(val)
	} else if val > n.val {
		n.right = n.right.insert(val)
	}
	return n
}

func (n *Node) search(val int) bool {
	if n == nil {
		return false
	}
	if val == n.val {
		return true
	}
	if val < n.val {
		return n.left.search(val)
	}
	return n.right.search(val)

}

func inorder(root *Node) []int {
	var res []int
	var dfs func(*Node)
	dfs = func(x *Node) {
		if x == nil {
			return
		}
		dfs(x.left)
		res = append(res, x.val)
		dfs(x.right)
	}
	dfs(root)
	return res

}

func main() {
	var root *Node
	for _, v := range []int{8, 3, 10, 1, 6, 14, 4, 7, 13} {
		root = root.insert(v)
	}
	fmt.Println(inorder(root))
	fmt.Println(root.search(6))
	fmt.Println(root.search(2))

}
