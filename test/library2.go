package main

import (
	"fmt"
	"time"
)

type Title string //book name
type Name string  //member name

type Landaudit struct {
	ckeckOut time.Time
	ckeckIn  time.Time
}
type Member struct {
	name  Name
	books map[Title]Landaudit
}
type BookEntry struct {
	total  int //owned by library
	lended int //lended by library

}
type Library struct {
	members map[Name]Member
	books   map[Title]BookEntry
}

func printMemberAudit(member *Member) {
	for title, audit := range member.books {
		var returnTime string
		if audit.ckeckIn.IsZero() {
			returnTime = "[not returned yet]"
		} else {
			returnTime = audit.ckeckOut.String()
		}
		fmt.Println(member.name, ":", title, ":", audit.ckeckOut.String(), "through", returnTime)
	}
}
func printMemberAudits(library *Library) {
	for _, member := range library.members {
		printMemberAudit(&member)
	}
}
func printLibrarybooks(library *Library) {
	fmt.Println()
	for title, book := range library.books {
		fmt.Println(title, "/ total:", book.total, "/ lended", book.lended)
	}
	println()
}

func chechoutBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title]
	if !found {
		fmt.Println("book is not part of library")
		return false
	}
	if book.lended == book.total {
		fmt.Println("no more books available to lend")
		return false
	}
	book.lended += 1
	library.books[title] = book
	member.books[title] = Landaudit{ckeckOut: time.Now()}
	return true
}
func returnBook(library *Library, titel Title, member *Member) bool {
	book, found := library.books[titel]
	if !found {
		fmt.Println("book is not part of libraryw")
		return false
	}
	audit, found := member.books[titel]
	if !found {
		fmt.Println(("member did not ckeck out this book"))
		return false

	}
	book.lended -= 1
	library.books[titel] = book

	audit.ckeckIn = time.Now()
	member.books[titel] = audit
	return true
}

func main() {
	library := Library{
		books:   make(map[Title]BookEntry),
		members: make(map[Name]Member),
	}
	library.books["webapp in go"] = BookEntry{
		total:  4,
		lended: 0,
	}
	library.books["the little go book "] = BookEntry{
		total:  3,
		lended: 0,
	}
	library.books["lets learn go"] = BookEntry{
		total:  2,
		lended: 0,
	}
	library.books["go bootcamp"] = BookEntry{
		total:  1,
		lended: 0,
	}
	library.members["aaa"] = Member{"aaa", make(map[Title]Landaudit)}
	library.members["ccc"] = Member{"ccc", make(map[Title]Landaudit)}
	library.members["bbb"] = Member{"bbb", make(map[Title]Landaudit)}

	fmt.Println("\ninitial:")
	printLibrarybooks(&library)
	printMemberAudits(&library)

	member := library.members["aaa"]
	ckeckOut := chechoutBook(&library, "go bootcamp", &member)
	fmt.Println("\nCheck out a book:")
	if ckeckOut {
		printLibrarybooks(&library)
		printMemberAudits(&library)

	}
	returned := returnBook(&library, "go bootcamp", &member)
	fmt.Println("\n check in a book:")
	if returned {
		printLibrarybooks(&library)
		printMemberAudits(&library)

	}

}
