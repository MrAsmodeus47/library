package manager

import (
	"strings"
	"time"
)

type Book struct {
	Title      string
	Available  bool
	BorrowedBy string
	DueDate    time.Time
}
type Node struct {
	Book  Book
	Left  *Node
	Right *Node
}

type Library struct {
	Root *Node
	insert
}

func insert(node *Node, book Book) *Node {

	if node == nil {
		return &Node{Book: book}
	}
	if strings.ToLower(book.Title) < strings.ToLower(node.Book.Title) {
		node.Left = insert(node.Left, book)
	} else {
		node.Right = insert(node.Right, book)
	}
	return node
}
