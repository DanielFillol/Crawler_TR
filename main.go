package main

import (
	"fmt"
	"github.com/Darklabel91/Crawler_TR/CSV"
	"github.com/Darklabel91/Crawler_TR/Crawler"
)

func main() {
	bookNames, err := CSV.ReadCsvFile("CSV/book4.csv")
	if err != nil {
		fmt.Println(err)
	}

	var data []Crawler.Book
	for i := 0; i < len(bookNames); i++ {
		book, err := Crawler.Craw(bookNames[i])
		if err != nil {
			fmt.Println(err)
		}

		data = append(data, book)
	}

	err = CSV.WriteCSV("Livros TR", "Result", data)
	if err != nil {
		fmt.Println(err)
	}
}
