package bookinfo

import (
	"fmt"
	"strings"
	"time"
)

type Book struct {
	Title     string
	Available bool
	//TODO convert it to User entity User{id , firstName , lastName}
	//TODO what is foreign key and primary key and cardinality for example one to one , one to many and so on
	BorrowedBy string
	DueDate    time.Time
}
type Node struct {
	Book  *Book
	Left  *Node
	Right *Node
}
type BooksByTitleTree struct {
	root *Node
}

// TODO separate it into another file and move borrow fucntion into it
// TODO  Make it Pointer
type Library struct {
	books *BooksByTitleTree
}

func NewLibrary() *Library {
	l := &Library{
		books: &BooksByTitleTree{},
	}
	return l
} //i dont know that i did right or not// Add Book To Tree
func (b *BooksByTitleTree) AddBookToTree(book *Book) {
	b.root = b.insertNode(b.root, book)
}
func (l *Library) AddingBook(book *Book) {

	l.books.AddBookToTree(book)
}

func (b *BooksByTitleTree) insertNode(node *Node, book *Book) *Node {

	if node == nil {
		return &Node{Book: book}
	}
	if strings.ToLower(book.Title) < strings.ToLower(node.Book.Title) {
		node.Left = b.insertNode(node.Left, book)
	} else {
		node.Right = b.insertNode(node.Right, book)
	}
	return node
}

// ///////////////////////////////////////////////////////////////
func (b *BooksByTitleTree) SearchBookInTree(title string) (bool, string) {
	node := b.searchNode(b.root, title)
	if node != nil {
		status := "available"
		if !node.Book.Available {
			status = fmt.Sprintf("borrowed by %s untile %s", node.Book.BorrowedBy, node.Book.DueDate.Format("2006-01-02"))
		}
		return true, fmt.Sprintf("Book '%s' exists in the library.status: %s\n", node.Book.Title, status)
	}

	return false, fmt.Sprintf("Book '%s' not found in the library.\n", title)

}
func (l *Library) SearchBook(title string) (bool, string) {
	return l.books.SearchBookInTree(title)

}
func (b *BooksByTitleTree) searchNode(node *Node, title string) *Node {
	if node == nil {
		return nil
	}
	if strings.EqualFold(node.Book.Title, title) {
		return node
	}
	if strings.ToLower(title) < strings.ToLower(node.Book.Title) {
		return b.searchNode(node.Left, title)

	}
	return b.searchNode(node.Right, title)
}

// //////////////////////////////////////////////////////////
func (b *BooksByTitleTree) RemoveBookFromTree(title string) (bool, string) {
	var removed bool
	b.root, removed = b.removeNode(b.root, title)
	if removed {
		//TODO transfer to main function
		return true, fmt.Sprintf("Book '%s' has been removed from library.\n", title)
	} else {
		return false, fmt.Sprintf("threr is no by name '%s' in the library ", title)
	}
}
func (l *Library) RemoveBook(title string) (bool, string) {
	return l.books.RemoveBookFromTree(title)
}
func (b *BooksByTitleTree) removeNode(node *Node, title string) (*Node, bool) {
	if node == nil {
		return nil, false
	}
	lt := strings.ToLower(title)
	cur := strings.ToLower(node.Book.Title)

	if lt < cur {
		var removed bool
		node.Left, removed = b.removeNode(node.Left, title)
		return node, removed
	} else if lt > cur {
		var removed bool
		node.Right, removed = b.removeNode(node.Right, title)
		return node, removed
	} else {
		if node.Left == nil && node.Right == nil {
			return nil, true
		}
		if node.Left == nil {
			return node.Right, true
		}
		if node.Right == nil {
			return node.Left, true
		}
		succ := b.finMin(node.Right)
		node.Book = succ.Book
		node.Right, _ = b.removeNode(node.Right, succ.Book.Title)
		return node, true
	}
}
func (b *BooksByTitleTree) finMin(node *Node) *Node {
	cur := node
	for cur.Left != nil {
		cur = cur.Left
	}
	return cur
}
func (b *BooksByTitleTree) Removeaction(title string) {
	var removed bool
	b.root, removed = b.removeNode(b.root, title)
	if removed {
		fmt.Printf("Book '%s' has been removed from library.\n", title)
	} else {
		fmt.Printf("threr is no by name '%s' in the library ", title)
	}
}

// ////////////////////////////////////////////////////////////////////////////////////////////////
func (b *BooksByTitleTree) BorrowBookFromTree(title string, borrower string, days int) (bool, string) {
	node := b.searchNode(b.root, title)
	if node == nil {
		return false, fmt.Sprintf("book '%s' not found in the library.\n", title)
	}
	if node.Book.Available {
		node.Book.Available = false
		node.Book.BorrowedBy = borrower
		node.Book.DueDate = time.Now().AddDate(0, 0, days)
		return true, fmt.Sprintf("book '%s' has been borrowed by %s.DUe date: %s\n ", node.Book.Title, borrower, node.Book.DueDate.Format("2006-01-02"))

	}
	return false, fmt.Sprintf("book '%s' is already borrowed by %s until %s.\n ", node.Book.Title, node.Book.BorrowedBy, node.Book.DueDate.Format("2006-01-02"))

}
func (l *Library) BorrowBook(title, borrower string, days int) (bool, string) {
	return l.books.BorrowBookFromTree(title, borrower, days)
}
