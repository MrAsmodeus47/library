package entity

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
} // Add Book To Tree

func (l *Library) AddingBook(book *Book) {

	l.books.AddBookToTree(book)
}

func (l *Library) RemoveBook(title string) (bool, string) {
	return l.books.RemoveBookFromTree(title)
}
func (l *Library) BorrowBook(title, borrower string, days int) (bool, string) {
	return l.books.BorrowBookFromTree(title, borrower, days)
}

func (l *Library) SearchBook(title string) (bool, string) {
	return l.books.SearchBookInTree(title)

}

// //////////////////////////////////////////////////////////
