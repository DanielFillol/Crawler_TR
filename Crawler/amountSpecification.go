package Crawler

import (
	"github.com/tebeka/selenium"
)

type bookLink struct {
	HasLink bool
	Link    selenium.WebElement
}

const (
	https = "https:"
)

func onlyOneBook(driver selenium.WebDriver, bookOpenLink string) (bookLink, error) {
	links, err := driver.FindElements(selenium.ByXPATH, bookOpenLink)
	if err != nil {
		return bookLink{}, err
	}

	if len(links) > 1 {
		//TODO: need to parse all books
		return bookLink{
			HasLink: true,
			Link:    nil,
		}, nil
	} else if len(links) != 0 {
		return bookLink{
			HasLink: true,
			Link:    links[0],
		}, nil
	} else {
		return bookLink{
			HasLink: false,
			Link:    nil,
		}, nil
	}
}

func amountSpecification(driver selenium.WebDriver, bookOpenLink string, productSpecificationTR string) (int, error) {
	oneBook, err := onlyOneBook(driver, bookOpenLink)
	if err != nil {
		return 0, err
	}

	if oneBook.HasLink {
		link, err := oneBook.Link.GetAttribute("href")
		if err != nil {
			return 0, err
		}

		err = driver.Get(https + link)
		if err != nil {
			return 0, err
		}

		specification, err := driver.FindElements(selenium.ByXPATH, productSpecificationTR)
		if err != nil {
			return 0, err
		}

		if len(specification) != 0 {
			return len(specification), nil
		}

	}

	return 0, nil

}
