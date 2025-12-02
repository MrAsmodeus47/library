package main

import (
	"fmt"
	"time"
)

type Title string //book title
type Name string  //members name
type LendingHstry struct {
	checkOut time.Time
	checkin  time.Time
}
type Member struct {
	name  Name
	books map[Title]LendingHstry
}
type BookEntry struct {
	total  int // book that exist in library
	lended int // total books lended
}
type Library struct {
	members map[Name]Member
	books   map[Title]BookEntry
}

func printMemberHstry(member *Member) { ///////////////////
	for title, history := range member.books {
		var returnTime string
		if history.checkin.IsZero() {
			returnTime = "[has not return yet]"
		} else {
			returnTime = history.checkin.String()
		}
		fmt.Println(member.name, ":", title, ":", history.checkOut.String(), "by", returnTime)
	}

}
func printMemberHstrys(library *Library) {
	for _, member := range library.members {
		printMemberHstry(&member)
	}
}

func printLabraryBooks(library *Library) {
	fmt.Println()
	for title, book := range library.books {
		fmt.Println(title, "/ totle", book.total, "/ lended", book.lended)
	}
	fmt.Println()
}
func ckeckOutBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title]
	if !found {
		fmt.Println("book not founded")
		return false
	}
	if book.lended == book.total {
		fmt.Println("wood needed")
		return false

	}
	book.lended += 1
	library.books[title] = book
	member.books[title] = LendingHstry{checkOut: time.Now()}
	return true
}
func returnBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title]
	if !found {
		fmt.Println("we dont have such thing")
		return false
	}
	history, found := member.books[title]
	if !found {
		fmt.Println("book is missing")
		return false
	}
	book.lended -= 1
	library.books[title] = book
	history.checkin = time.Now()
	member.books[title] = history
	return true
}

func main() {
	library := Library{
		books:   make(map[Title]BookEntry),
		members: make(map[Name]Member),
	}
	library.books["webapps in go"] = BookEntry{
		total:  4,
		lended: 0,
	}
	library.books["witcher"] = BookEntry{
		total:  3,
		lended: 0,
	}
	library.books["hp love craft"] = BookEntry{
		total:  2,
		lended: 0,
	}
	library.books["deltora belt"] = BookEntry{
		total:  1,
		lended: 0,
	}
	library.members["jamshid"] = Member{"jamshid", make(map[Title]LendingHstry)}
	library.members["oscar"] = Member{"oscar", make(map[Title]LendingHstry)}
	library.members["hamed"] = Member{"hamed", make(map[Title]LendingHstry)}
	fmt.Println("\ninitial:")
	printLabraryBooks(&library)
	printMemberHstrys(&library)

	member := library.members["jamshid"]
	checkedOut := ckeckOutBook(&library, "witcher", &member)
	fmt.Println("\nchech out a book :")
	if checkedOut {
		printLabraryBooks(&library)
		printMemberHstrys(&library)
	}
	returned := returnBook(&library, "witcher", &member)
	fmt.Println("\nchech out a book:")
	if returned {
		printLabraryBooks(&library)
		printMemberHstrys(&library)
	}
}
