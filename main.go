package main

import (
	"fmt"
	"github.com/Darklabel91/Crawler_TR/CSV"
	"github.com/Darklabel91/Crawler_TR/Crawler"
	"github.com/tebeka/selenium"
	"strconv"
)

func main() {
	bookNames, err := CSV.ReadCsvFile("CSV/books.csv")
	if err != nil {
		fmt.Println(err)
	}

	var data []Crawler.Book
	var driver selenium.WebDriver
	for i := 0; i < len(bookNames); i++ {
		driver, err = Crawler.SeleniumWebDriver()
		if err != nil {
			fmt.Println(err)
		}

		book, err := Crawler.Craw(driver, bookNames[i])
		if err != nil {
			fmt.Println(err)
		}

		data = append(data, book)

		driver.Close()

		fmt.Println(strconv.Itoa(i+1) + "/" + strconv.Itoa(len(bookNames)))
	}

	defer driver.Close()

	err = CSV.WriteCSV("Livros TR", "Result", data)
	if err != nil {
		fmt.Println(err)
	}
}
