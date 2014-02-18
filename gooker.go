package gooker

type Book struct {
	Title  string
	Author string
	Pages  int
}

func (this *Book) ScreamLoud() string {
	return "LOUD!"
}

func GetSomeBook() Book {
	return Book{"Rainbow Six", "Tom Clancy", 1354}
}
