package Crawler

import (
	"github.com/tebeka/selenium"
	"strconv"
)

const (
	xpathTitleInit = "//*[@id=\"caracteristicas\"]/div[2]/div["
	xpathTitleEnd  = "]/div[1]"
	ISBN           = "Código ISBN"
	Date           = "Data de disponibilidade"
	Pages          = "Número de páginas"
	Year           = "Ano de publicação"
)

func title(driver selenium.WebDriver, i int) (string, error) {
	xpath := xpathTitleInit + strconv.Itoa(i) + xpathTitleEnd

	titles, err := driver.FindElements(selenium.ByXPATH, xpath)
	if err != nil {
		return "", err
	}

	if len(titles) != 0 {
		textTitle, err := titles[0].Text()
		if err != nil {
			return "", err
		}

		return textTitle, nil
	}

	return "", nil
}

func bookFounded(driver selenium.WebDriver, bookName string) (Book, error) {
	var isbn string
	var dtDis string
	var pgs string
	var yr string

	bookLink, err := amountSpecification(driver, bookOpenLink, productSpecificationTR)
	if err != nil {
		return Book{}, err
	}

	for i := 0; i < bookLink.Amount; i++ {
		specification, err := title(driver, i)
		if err != nil {
			return Book{}, err
		}

		if specification != "" {
			//TODO: does this makes sense?
			position := i + 1
			elem, err := driver.FindElement(selenium.ByXPATH, "//*[@id=\"caracteristicas\"]/div[2]/div["+strconv.Itoa(position)+"]/div[2]")
			if err != nil {
				return Book{}, err
			}

			switch specification {
			case ISBN:
				elem, err = driver.FindElement(selenium.ByXPATH, "//*[@id=\"caracteristicas\"]/div[2]/div["+strconv.Itoa(position-1)+"]/div[2]")
				if err != nil {
					return Book{}, err
				}

				text, err := elem.Text()
				if err != nil {
					return Book{}, err
				}

				isbn = text
			case Date:
				elem, err = driver.FindElement(selenium.ByXPATH, "//*[@id=\"caracteristicas\"]/div[2]/div["+strconv.Itoa(position-1)+"]/div[2]")
				if err != nil {
					return Book{}, err
				}

				text, err := elem.Text()
				if err != nil {
					return Book{}, err
				}

				dtDis = text
			case Pages:
				elem, err = driver.FindElement(selenium.ByXPATH, "//*[@id=\"caracteristicas\"]/div[2]/div["+strconv.Itoa(position-1)+"]/div[2]")
				if err != nil {
					return Book{}, err
				}

				text, err := elem.Text()
				if err != nil {
					return Book{}, err
				}

				pgs = text
			case Year:
				elem, err = driver.FindElement(selenium.ByXPATH, "//*[@id=\"caracteristicas\"]/div[2]/div["+strconv.Itoa(position-1)+"]/div[2]")
				if err != nil {
					return Book{}, err
				}

				text, err := elem.Text()
				if err != nil {
					return Book{}, err
				}

				yr = text
			}
		}
	}

	return Book{
		SearchName:    bookName,
		ISBN:          isbn,
		AvailableDate: dtDis,
		Pages:         pgs,
		PubYear:       yr,
		Link:          bookLink.Link,
	}, nil

}
