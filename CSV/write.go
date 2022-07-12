package CSV

import (
	"encoding/csv"
	"github.com/Darklabel91/Crawler_TR/Crawler"
	"os"
	"path/filepath"
	"strings"
)

//writeCSV exports a csv to a given folder, with a given name from a collection of Books
func WriteCSV(fileName string, folderName string, cnjRows []Crawler.Book) error {
	var rows [][]string

	rows = append(rows, generateHeaders())

	for _, cnj := range cnjRows {
		rows = append(rows, generateRow(cnj))
	}

	cf, err := createFile(folderName + "/" + fileName + ".csv")
	if err != nil {
		return err
	}

	defer cf.Close()

	w := csv.NewWriter(cf)

	err = w.WriteAll(rows)
	if err != nil {
		return err
	}

	return nil
}

// create csv file from operating system
func createFile(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.Create(p)
}

// generate the necessary headers for csv file
func generateHeaders() []string {
	return []string{
		"Nome Pesquisado",
		"ISBN",
		"Data de Disponibilização",
		"Número de páginas",
		"Ano de publicação",
	}
}

// returns a []string that compose the row in the csv file
func generateRow(books Crawler.Book) []string {
	var cnjReturn string
	return []string{
		books.SearchName,
		books.ISBN,
		books.AvailableDate,
		books.Pages,
		books.PubYear,
		strings.TrimSpace(cnjReturn),
	}
}
