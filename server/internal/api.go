package internal

import "errors"

type API struct{}

func (a *API) GetDB(title string, reply *[]Book) error {
	*reply = Database
	return nil
}

func (a *API) GetBookByName(title string, reply *Book) error {
	var getBook Book

	for _, val := range Database {
		if val.Title == title {
			getBook = val
			*reply = getBook
			return nil
		}
	}

	return errors.New("Can't find the book!")
}

func (a *API) AddNewBook(book Book, reply *Book) error {
	Database = append(Database, book)
	*reply = book
	return nil
}

func (a *API) EditBook(edit Book, reply *Book) error {
	var changed Book

	for idx, val := range Database {
		if val.Title == edit.Title {
			Database[idx] = edit
			changed = Database[idx]
			*reply = changed
			return nil
		}
	}

	return errors.New("Can't find the book!")
}

func (a *API) DeleteBook(book Book, reply *Book) error {
	var del Book

	for idx, val := range Database {
		if val == book {
			Database = append(Database[:idx], Database[idx+1:]...)
			del = book
			*reply = del
			return nil
		}
	}
	return errors.New("Can't find the book!")
}
