package library

type Book struct {
	Title  string `bson:"title"`
	Author string `bson:"author"`
	Year   int    `bson:"year"`
}

type PrintedBook struct {
	Book  Book `bson:",inline"`
	Pages int  `bson:"pages"`
}

// GetDetails implements LibraryItem.
func (p PrintedBook) GetDetails() string {
	panic("unimplemented")
}

type Ebook struct {
	Book       Book `bson:",inline"`
	FileSizeMB int  `bson:"fileSizeMB"`
}

// GetDetails implements LibraryItem.
func (e Ebook) GetDetails() string {
	panic("unimplemented")
}
