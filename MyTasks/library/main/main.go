package main

import (
	"bufio"
	"fmt"

	//TODO use go modules ,
	"gogogo/MyTasks/library/bookinfo"
	"os"
	"strings"
)

func main() {
	//TODO use constructor for entity struct//ithink done
	//TODO use Library struct//done
	library := bookinfo.NewLibrary()
	reader := bufio.NewReader(os.Stdin)

	//Add default books
	defaultBooks := []string{
		"The Great Gatsby",
		"To Kill a Mockingbird",
		"1984",
		"Pride and Prejudice",
		"The Hobbit",
		"aaaaa",
	}

	for _, book := range defaultBooks {
		library.AddingBook(&bookinfo.Book{
			Title:     book,
			Available: true,
		})
	}

	/////////
	var choice int
	var title, borrower string
	var days int

	for {
		fmt.Println("\nWelcome to the Library Management System!")
		fmt.Println("1. Add a new book")
		fmt.Println("2. Search for a book")
		fmt.Println("3. Remove a book")
		fmt.Println("4. Borrow a book")
		fmt.Println("5. Exit")
		fmt.Print("\nEnter your choice (1-5): ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter the title of the book to add: ")
			title, _ = reader.ReadString('\n')
			title = strings.TrimSpace(title)
			library.AddingBook(&bookinfo.Book{
				Title: title,
			})
		case 2:
			fmt.Print("Enter the title or part of the title to search: ")
			title, _ = reader.ReadString('\n')
			title = strings.TrimSpace(title)
			_, msg := library.SearchBook(title)
			fmt.Println(msg)
		case 3:
			fmt.Print("Enter the title of the book to remove: ")
			title, _ = reader.ReadString('\n')
			title = strings.TrimSpace(title)
			_, msg := library.RemoveBook(title) //ok is deleted
			fmt.Println(msg)
		case 4:
			fmt.Print("Enter the title of the book to borrow: ")
			title, _ = reader.ReadString('\n')
			title = strings.TrimSpace(title)
			fmt.Print("Enter your name: ")
			borrower, _ = reader.ReadString('\n')
			borrower = strings.TrimSpace(borrower)
			fmt.Print("Enter number of days to borrow: ")
			fmt.Scanln(&days)
			_, msg := library.BorrowBook(title, borrower, days)
			fmt.Println(msg)
		case 5:
			fmt.Println("Thank you for using the Library Management System!")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
