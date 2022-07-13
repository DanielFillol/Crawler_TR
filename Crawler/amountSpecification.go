package Crawler

import (
	"github.com/tebeka/selenium"
)

type bookLink struct {
	HasLink    bool
	LinkSearch selenium.WebElement
	Link       string
	Amount     int
}

const (
	https     = "https:"
	attribute = "href"
)

func onlyOneBook(driver selenium.WebDriver, bookOpenLink string) (bookLink, error) {
	links, err := driver.FindElements(selenium.ByXPATH, bookOpenLink)
	if err != nil {
		return bookLink{}, err
	}

	if len(links) > 1 {
		//TODO: need to parse all books
		return bookLink{
			HasLink:    true,
			LinkSearch: links[0],
		}, nil
	} else if len(links) != 0 {
		return bookLink{
			HasLink:    true,
			LinkSearch: links[0],
		}, nil
	} else {
		return bookLink{
			HasLink:    false,
			LinkSearch: nil,
		}, nil
	}
}

func amountSpecification(driver selenium.WebDriver, bookOpenLink string, productSpecificationTR string) (bookLink, error) {
	oneBook, err := onlyOneBook(driver, bookOpenLink)
	if err != nil {
		return bookLink{}, err
	}

	if oneBook.HasLink {
		href, err := oneBook.LinkSearch.GetAttribute(attribute)
		if err != nil {
			return bookLink{}, err
		}

		link := https + href
		driver.Get(link)

		specification, err := driver.FindElements(selenium.ByXPATH, productSpecificationTR)
		if err != nil {
			return bookLink{}, err
		}

		if len(specification) != 0 {
			return bookLink{
				HasLink:    oneBook.HasLink,
				LinkSearch: oneBook.LinkSearch,
				Link:       link,
				Amount:     len(specification),
			}, nil
		}

	}

	return bookLink{}, nil

}
