package Crawler

import (
	"github.com/tebeka/selenium"
)

func hasFoundBooks(driver selenium.WebDriver, searchResultTR string, notFoundTR string) (bool, error) {
	searchReturn, err := hasSearchReturn(driver, searchResultTR)
	if err != nil {
		return false, err
	}

	if searchReturn {
		resultElem, err := driver.FindElements(selenium.ByXPATH, searchResultTR)
		if err != nil {
			return false, err
		}

		foundBooks, err := findBooks(resultElem, notFoundTR)
		if err != nil {
			return false, err
		}

		if foundBooks {
			return true, nil
		}
	}

	return false, nil
}

func findBooks(resultElem []selenium.WebElement, notFoundTR string) (bool, error) {
	resultText, err := resultElem[0].Text()
	if err != nil {
		return false, err
	}

	if resultText != notFoundTR {
		return true, nil
	}

	return false, nil
}

func hasSearchReturn(driver selenium.WebDriver, searchResultTR string) (bool, error) {
	resultElem, err := driver.FindElements(selenium.ByXPATH, searchResultTR)
	if err != nil {
		return false, err
	}

	if len(resultElem) != 0 {
		return true, nil
	}

	return false, nil
}
