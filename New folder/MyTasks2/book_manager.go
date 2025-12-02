package main

import (
	"bufio"
	"fmt"
	"os"
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
	Book  *Book
	Left  *Node
	Right *Node
}
type BooksByTitleTree struct {
	root *Node
}
type Library struct {
	books BooksByTitleTree
}

func (b *BooksByTitleTree) AddBookToTree(book *Book) {
	b.root = insertNode(b.root, book)
	// Add Book To Trre
}

func insertNode(node *Node, book *Book) *Node {

	if node == nil {
		return &Node{Book: book}
	}
	if strings.ToLower(book.Title) < strings.ToLower(node.Book.Title) {
		node.Left = insertNode(node.Left, book)
	} else {
		node.Right = insertNode(node.Right, book)
	}
	return node
}

// func insertnode(node *Node, book *Book) *Node {

// 	if node == nil {
// 		return &Node{Book: book}
// 	}
// 	if strings.ToLower(book.Title) < strings.ToLower(node.Book.Title) {
// 		node.Left = insertnode(node.Left, book)
// 	} else {
// 		node.Right = insertnode(node.Right, book)
// 	}
// 	return node
// }

// /////////////////////////////////////////////////////////////////////////////search
func (b *BooksByTitleTree) SearchBook(title string) (bool, *Book) {

	// Search Book In Tree
	return false, nil
}
func (b BooksByTitleTree) Searchaction(title string) { //new one
	node := searchNode(b.root, title)
	if node != nil {
		status := "available"
		if !node.Book.Available {
			status = fmt.Sprintf("borrowed by %s untill %s", node.Book.BorrowedBy, node.Book.DueDate.Format("2006-01-02"))
		}
		fmt.Printf("Book '%s' exists in the library.status: %s\n", node.Book.Title, status)
	} else {
		fmt.Printf("Book '%s' not found in the library.\n", title)
	}
}
func searchNode(node *Node, title string) *Node {
	if node == nil {
		return nil
	}
	if strings.EqualFold(node.Book.Title, title) {
		return node
	}
	if strings.ToLower(title) < strings.ToLower(node.Book.Title) {
		return searchNode(node.Left, title)

	}
	return searchNode(node.Right, title)
}

// ///////////////////////////remove
func (b *BooksByTitleTree) RemoveBook(title string) bool {

	// Search Book In Tree
	return false
}
func removeNode(node *Node, title string) (*Node, bool) { //third if need to study
	if node == nil {
		return nil, false
	}
	lt := strings.ToLower(title)
	cur := strings.ToLower(node.Book.Title)

	if lt < cur {
		var removed bool
		node.Left, removed = removeNode(node.Left, title)
		return node, removed
	} else if lt > cur {
		var removed bool
		node.Right, removed = removeNode(node.Right, title)
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
		succ := finMin(node.Right)
		node.Book = succ.Book
		node.Right, _ = removeNode(node.Right, succ.Book.Title)
		return node, true
	}
}
func finMin(node *Node) *Node {
	cur := node
	for cur.Left != nil {
		cur = cur.Left
	}
	return cur
}
func (b *BooksByTitleTree) Removeaction(title string) {
	var removed bool
	b.root, removed = removeNode(b.root, title)
	if removed {
		fmt.Printf("Book '%s' has been removed from library.\n", title)
	} else {
		fmt.Printf("threr is no by name '%s' in the library ", title)
	}
}

// //////////////////////////////////////////////////////
func (b *BooksByTitleTree) AddBook(title string) { //new one (Add book)
	book := Book{}
	root := AddBook(book)
}

// ////////////////////////////////////////////////////////////////////////////
func (l *Library) BorrowBook(title string, borrower string, days int) {
	node := searchNode(l.books.root, title)
	if node == nil {
		fmt.Printf("book '%s' not found in the library.\n", title)
		return
	}
	if node.Book.Available {
		node.Book.Available = false
		node.Book.BorrowedBy = borrower
		node.Book.DueDate = time.Now().AddDate(0, 0, days)
		fmt.Printf("book '%s' has been borrowed by %s.DUe date: %s\n ", node.Book.Title, borrower, node.Book.DueDate.Format("2006-0102"))
	} else {
		fmt.Printf("book '%s' is already borrowed by %s until %s.\n ", node.Book.Title, node.Book.BorrowedBy, node.Book.DueDate.Format("2006-01-02"))
	}
}

// ///////////////////////////////////////////////////////////////////////////

//////////////////////////////////////////////////////////////////////////////////////////////////

func inorder(node *Node) {
	if node == nil {
		return
	}

	inorder(node.Left)
	status := "Available"
	if !node.Book.Available {
		status = fmt.Sprintf("borrowed by %s untile %s,", node.Book.BorrowedBy, node.Book.DueDate.Format("2006-01-02")) //neeed to changee time formart
	}
	fmt.Printf("- %s (%s)\n", node.Book.Title, status)
	inorder(node.Right)
}

func (b BooksByTitleTree) ListBooks() {
	if b.root == nil {
		fmt.Println("The library is empty.")
		return
	}
	fmt.Println("Current books in the library:")
	inorder(b.root)
}

//////////////////////////////////////////////////////////////////////////////////////////////////

// //////////////////////////////////////////////////////////////////////////////////////////////
func main() {
	library := BooksByTitleTree{}
	reader := bufio.NewReader(os.Stdin)

	// Add default books //need updaate
	defaultBooks := []string{
		"The Great Gatsby",
		"To Kill a Mockingbird",
		"1984",
		"Pride and Prejudice",
		"The Hobbit",
		"aaaaa",
	}

	for _, book := range defaultBooks {
		library.AddBook(book)
	}
	var choice int
	var title, borrower string
	var days int

	for {
		fmt.Println("\nWelcome to the Library Management System!")
		fmt.Println("1. Add a new book")
		fmt.Println("2. Borrow a book")
		fmt.Println("3. Search for a book")
		fmt.Println("4. List all books")
		fmt.Println("5. Remove a book")
		fmt.Println("6. Exit")
		fmt.Print("\nEnter your choice (1-6): ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter the title of the book to add: ")
			title, _ = reader.ReadString('\n')
			title = strings.TrimSpace(title)
			library.AddBook(title)
		case 2:
			fmt.Print("Enter the title of the book to borrow: ")
			title, _ = reader.ReadString('\n')
			title = strings.TrimSpace(title)
			fmt.Print("Enter your name: ")
			borrower, _ = reader.ReadString('\n')
			borrower = strings.TrimSpace(borrower)
			fmt.Print("Enter number of days to borrow: ")
			fmt.Scanln(&days)
			library.BorrowBook(title, borrower, days)
		case 3:
			fmt.Print("Enter the title or part of the title to search: ")
			title, _ = reader.ReadString('\n')
			title = strings.TrimSpace(title)
			library.SearchBook(title)
		case 4:
			BooksByTitleTree.ListBooks()
		case 5:
			fmt.Print("Enter the title of the book to remove: ")
			title, _ = reader.ReadString('\n')
			title = strings.TrimSpace(title)
			library.RemoveBook(title)
		case 6:
			fmt.Println("Thank you for using the Library Management System!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
