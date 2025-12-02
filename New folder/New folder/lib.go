package main

import (
	"container/list"
	"fmt"
	"strings"
	"time"
)
	type Book struct {
	Bname string
	Available bool
	Date time.Time
	BorrowedBy string

}
type Shelf struct{
	books []Book
}  
func (list *shelf) AddBook(title string){
	book := Book{
		Bname : bname,
		available: true, 
		//will be used    BorrowedBy : "",
	}
	list.books = append(list.books,book) 
}
func (list *shelf) borrowBook (title string , borrower string, days int){
	for i := range list.books {
		if strings.EqualFold(list.books[i].Title,title){
			if list.books[i].available{
				list.books[i].available= false
				list.books[i].BorrowedBy = borrower
				list.books[i].Date= time.Now().AddDate(0,0,days)
			}else {
				fmt.Printf("book'%s' is borrowed by %s.\n",title,borrower)
			}
			return
		}
	}
	fmt.Printf("we dont have Book'%s' in the library.\n",title)
}
func (list *Shelf)listBooks(){
	if len(list.books)==0{
		fmt.Println("the shelfs are empty")
		return
	}
	fmt.Println("this books are in library:")
	for _ , book := range list.books {
		status := "Available"
		if !book.Available{
			status = fmt.Printf("borrowed by %s until %s",book.BorrowedBy,book.Date.Format("2000-01-01"))
		}
		fmt.Printf("-> %s (%s)\n",book.Title,status)
	}
}






func main (){
	shelf := Shelf{}

	librarybooks := []string{
		"zzzzzzzzzzzzzzz"
		"aaaaaaaaaaaaaaaa"
		"sssssssssssssss"
		"fffffffffffffffff"
	}
	for_,book := range liblibrarybooks{
		shelf.addBooshelf(book)
	}





















choice int
title,bo

















}