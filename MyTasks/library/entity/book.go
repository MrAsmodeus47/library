package entity

import "time"

type Book struct {
	Title     string
	Available bool
	DueDate   time.Time

	//TODO convert it to User entity User{id , firstName , lastName}//mybe done
	//TODO what is foreign key and primary key and cardinality for example one to one , one to many and so on//done
}
type User struct {
	id         int
	FirstName  string
	LastName   string
	BorrowedBy string
}
