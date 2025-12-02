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
	Book  Book
	Left  *Node
	Right *Node
}

type Library struct {
	Root *Node
	insert
}

//	func (l *Library) AddBook(title string) {       //old one
//		book := Book{
//			Title:      title,
//			Available:  true,
//			BorrowedBy: "",
//		}
//		l.Books = append(l.Books, book)
//		// Sort books after adding
//		sort.Slice(l.Books, func(i, j int) bool {
//			return strings.ToLower(l.Books[i].Title) < strings.ToLower(l.Books[j].Title)
//		})
//		fmt.Printf("Book '%s' has been added to the library.\n", title)
//	}
//
// //////////////////////////////////////////////////////
func (l *Library) AddBook(title string) { //new one (Add book)
	book := Book{
		Title:      title,
		Available:  true,
		BorrowedBy: "",
	}
	l.Root = insert(l.Root, book)
	fmt.Printf("book '%s' has been added to the library.\n", title)
}

// ////////////////////////////////////////////
// func (l *Library) BorrowBook(title string, borrower string, days int) {  // going to change
//
//		for i := range l.Books {
//			if strings.EqualFold(l.Books[i].Title, title) {
//				if l.Books[i].Available {
//					l.Books[i].Available = false
//					l.Books[i].BorrowedBy = borrower
//					l.Books[i].DueDate = time.Now().AddDate(0, 0, days)
//					fmt.Printf("Book '%s' has been borrowed by %s. Due date: %s\n",
//						title, borrower, l.Books[i].DueDate.Format("2006-01-02"))
//				} else {
//					fmt.Printf("Book '%s' is already borrowed.\n", title)
//				}
//				return
//			}
//		}
//		fmt.Printf("Book '%s' not found in the library.\n", title)
//	}
func (l *Library) BorrowBook(title string, borrower string, days int) {
	node := search(l.Root, title)
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

/////////////////////////////////////////////
// func (l *Library) SearchBook(title string) {    ////linear search
//
//		found := false
//		for _, book := range l.Books {
//			if strings.Contains(strings.ToLower(book.Title), strings.ToLower(title)) {
//				found = true
//				status := "available"
//				if !book.Available {
//					status = fmt.Sprintf("borrowed by %s until %s", book.BorrowedBy, book.DueDate.Format("2006-01-02"))
//				}
//				fmt.Printf("Book '%s' exists in the library. Status: %s\n", book.Title, status)
//			}
//		}
//		if !found {
//			fmt.Printf("No books found matching '%s' in the library.\n", title)
//		}
//	}

//////////////////////////////////////////////////////////////////////////////////////////////////
//func (l *Library) SortBooks() {       //need to delete
//sort.Slice(l.Books, func(i, j int) bool {
//	return strings.ToLower(l.Books[i].Title) < strings.ToLower(l.Books[j].Title)
//})
//fmt.Println("Books have been sorted alphabetically.")
//}
//////////////////////////////////////////////////
//old sht
//func (l *Library) ListBooks() {
//	if len(l.Books) == 0 {
//		fmt.Println("The library is empty.")
//		return
//	}
//	fmt.Println("Current books in the library:")
//	for _, book := range l.Books {
//		status := "Available"
//		if !book.Available {
//			status = fmt.Sprintf("Borrowed by %s until %s", book.BorrowedBy, book.DueDate.Format("2006-01-02"))
//		}
//		fmt.Printf("- %s (%s)\n", book.Title, status)
//	}
//}

// new one

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

func (l *Library) ListBooks() {
	if l.Root == nil {
		fmt.Println("The library is empty.")
		return
	}
	fmt.Println("Current books in the library:")
	inorder(l.Root)
}
func finMin(node *Node) *Node {
	cur := node
	for cur.Left != nil {
		cur = cur.Left
	}
	return cur
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

//////////////////////////////////////////////////////////////////////////////////////////////////

// func (l *Library) RemoveBook(title string) {   //old one
// 	for i, book := range l.Books {
// 		if strings.EqualFold(book.Title, title) {
// 			// Remove the book by creating a new slice without it
// 			l.Books = append(l.Books[:i], l.Books[i+1:]...)
// 			fmt.Printf("Book '%s' has been removed from the library.\n", title)
// 			return
// 		}
// 	}
// 	fmt.Printf("Book '%s' not found in the library.\n", title)
// }

func (l *Library) RemoveBook(title string) {
	var removed bool
	l.Root, removed = removeNode(l.Root, title)
	if removed {
		fmt.Printf("Book '%s' has been removed from library.\n", title)
	} else {
		fmt.Printf("threr is no by name '%s' in the library ", title)
	}
}

// //////////////////////////////////////////////////////////////////////////////////////////////
func main() {
	library := Library{}
	reader := bufio.NewReader(os.Stdin)

	// Add default books //need updaate
	defaultBooks := []string{ //have probleme in running works for search but not for print
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

	// Sort books after initialization   //need to change
	//sort.Slice(library.Books, func(i, j int) bool {
	//	return strings.ToLower(library.Books[i].Title) < strings.ToLower(library.Books[j].Title)
	//})

	var choice int
	var title, borrower string
	var days int

	for {
		fmt.Println("\nWelcome to the Library Management System!")
		fmt.Println("1. Add a new book")
		fmt.Println("2. Borrow a book")
		fmt.Println("3. Search for a book")
		//fmt.Println("4. Sort books alphabetically") //need to delete
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
			library.SearchBook(title) //need to delete or change
		case 4:
			library.ListBooks()
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
