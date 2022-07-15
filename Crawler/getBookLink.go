package Crawler

import (
	"errors"
	"github.com/tebeka/selenium"
	"strconv"
	"strings"
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

func getBookLink(driver selenium.WebDriver, bookOpenLink string, productSpecificationTR string, bookName string) (bookLink, error) {
	oneBook, err := getLink(driver, bookOpenLink, bookName)
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

func getLink(driver selenium.WebDriver, bookOpenLink string, bookName string) (bookLink, error) {
	links, err := driver.FindElements(selenium.ByXPATH, bookOpenLink)
	if err != nil {
		return bookLink{}, err
	}

	if len(links) > 1 {
		link, err := getSpecificLink(driver, links, bookName)
		if err != nil {
			return bookLink{}, err
		}
		return bookLink{
			HasLink:    true,
			LinkSearch: link,
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

func getSpecificLink(driver selenium.WebDriver, links []selenium.WebElement, bookName string) (selenium.WebElement, error) {
	for link := range links {
		elemBookTitle, err := driver.FindElement(selenium.ByXPATH, "//*[@id=\"smarthint-search-products\"]/section/div[2]/div/div/ul/div["+strconv.Itoa(link+1)+"]/li/div/div/div[1]/div[3]/h3/a/span")
		if err != nil {
			return nil, err
		}

		bookTitle, err := elemBookTitle.Text()
		if err != nil {
			return nil, err
		}

		if strings.Contains(bookTitle, bookName) {
			elementLinkSearch, err := driver.FindElement(selenium.ByXPATH, "//*[@id=\"smarthint-search-products\"]/section/div[2]/div/div/ul/div["+strconv.Itoa(link+1)+"]/li/div/div/div[1]/div[3]/h3/a")
			if err != nil {
				return nil, err
			}

			return elementLinkSearch, nil
		}
	}

	return nil, errors.New("bookName not found in search result list")
}
